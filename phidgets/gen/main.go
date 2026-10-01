// Command gen writes the Go wrappers for the channel classes listed in spec.txt
// from the Phidget22 C header. Run it with `go generate` from the phidgets dir.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"os"
	"regexp"
	"sort"
	"strings"
)

var (
	header = flag.String("header", "/usr/include/phidget22.h", "path to phidget22.h")
	spec   = flag.String("spec", "gen/spec.txt", "classes to generate")
	out    = flag.String("out", ".", "output directory")
)

type param struct{ typ, name string }

type fn struct {
	name   string // e.g. getBackEMF
	params []param
}

var (
	funcRe = regexp.MustCompile(`PhidgetReturnCode (Phidget\w+?)_(\w+)\(((?:[^()]|\([^()]*\))*)\);`)
	cbRe   = regexp.MustCompile(`typedef void \( \*(Phidget\w+_On\w+Callback)\)\(([^)]*)\);`)
	enumRe = regexp.MustCompile(`typedef enum \{([^}]*)\} (\w+);`)
	valRe  = regexp.MustCompile(`(\w+) = `)
)

// scalar C types -> Go type and conversion helpers.
var scalars = map[string]struct{ goT, cT, get string }{
	"double":   {"float64", "C.double", "getDouble"},
	"uint32_t": {"uint32", "C.uint32_t", "getUint32"},
	"int":      {"int", "C.int", "getInt"},
	"int64_t":  {"int64", "C.int64_t", "getInt64"},
	"uint64_t": {"uint64", "C.uint64_t", "getUint64"},
}

// callback signatures (after handle, ctx) -> Go func type, C typedef and shim.
var callbacks = map[string]struct{ goT, typ, shim string }{
	"":                                     {"func()", "void", "cvoidcallback"},
	"double":                               {"func(float64)", "double", "ccallback"},
	"double,double":                        {"func(float64, float64)", "twofloat", "ctwofloatcallback"},
	"double,double,double":                 {"func(float64, float64, float64)", "threefloat", "cthreefloatcallback"},
	"double[3],double":                     {"func([]float64, float64)", "motion", "cmotioncallback"},
	"double[4],double":                     {"func([]float64, float64)", "quaternion", "cquaternioncallback"},
	"int,double,int":                       {"func(int, float64, bool)", "encoder", "cencodercallback"},
	"uint64_t,double":                      {"func(uint64, float64)", "count", "ccountcallback"},
	"double[3],double[3],double[3],double": {"func([]float64, []float64, []float64, float64)", "spatial", "cspatialcallback"},
}

func main() {
	flag.Parse()
	classes, bools := readSpec(*spec)
	raw, err := os.ReadFile(*header)
	check(err)
	text := strings.Join(strings.Fields(string(raw)), " ")

	enums := map[string]bool{}
	enumNames := map[string]string{}
	for _, m := range enumRe.FindAllStringSubmatch(text, -1) {
		enums[m[2]] = true
		enumNames[m[2]] = m[1]
	}
	cbs := map[string]string{}
	for _, m := range cbRe.FindAllStringSubmatch(text, -1) {
		cbs[m[1]] = m[2]
	}
	byClass := map[string][]fn{}
	for _, m := range funcRe.FindAllStringSubmatch(text, -1) {
		byClass[m[1]] = append(byClass[m[1]], fn{m[2], parseParams(m[3])})
	}

	usedEnums := map[string]bool{}
	var attach bytes.Buffer
	for _, c := range classes {
		var b bytes.Buffer
		g := &gen{b: &b, cls: "Phidget" + c, bools: bools, enums: enums, used: usedEnums, cbs: cbs}
		g.class(byClass["Phidget"+c])
		write(*out+"/"+strings.ToLower(c)+"_gen.go", b.Bytes())
		fmt.Fprintf(&attach, "\tcase C.PHIDCHCLASS_%s:\n\t\treturn &Phidget%s{phidget{handle: channel}, C.Phidget%sHandle(unsafe.Pointer(channel))}\n", strings.ToUpper(c), c, c)
	}
	var e bytes.Buffer
	e.WriteString(enumHeader)
	names := make([]string, 0, len(usedEnums))
	for name := range usedEnums {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		writeEnum(&e, name, valRe.FindAllStringSubmatch(enumNames[name], -1))
	}
	write(*out+"/enums_gen.go", e.Bytes())
	write(*out+"/attach_gen.go", []byte(fmt.Sprintf(attachTmpl, attach.String())))
}

func readSpec(path string) (classes []string, bools map[string]bool) {
	raw, err := os.ReadFile(path)
	check(err)
	bools = map[string]bool{}
	for _, l := range strings.Split(string(raw), "\n") {
		f := strings.Fields(l)
		if len(f) < 2 || f[0][0] == '#' {
			continue
		}
		switch f[0] {
		case "classes":
			classes = append(classes, f[1:]...)
		case "bool":
			for _, n := range f[1:] {
				bools[n] = true
			}
		}
	}
	return
}

var ptrArrRe = regexp.MustCompile(`(\w+)\s*\(\s*\*\s*(\w+)\s*\)\s*\[(\d+)\]`)

func parseParams(s string) []param {
	s = ptrArrRe.ReplaceAllString(strings.ReplaceAll(s, "*", " * "), "$1 $2[$3]") // double (*x)[3] -> double x[3]
	var ps []param
	for i, p := range strings.Split(s, ",") {
		p = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(p), "const "))
		if p == "" || p == "void" || i == 0 && strings.HasSuffix(strings.Fields(p)[0], "Handle") {
			continue
		}
		f := strings.Fields(p)
		name := f[len(f)-1]
		typ := strings.Join(f[:len(f)-1], " ")
		if j := strings.Index(name, "["); j >= 0 { // double x[3] -> "double[3]"
			typ += name[j:]
			name = name[:j]
		}
		ps = append(ps, param{strings.ReplaceAll(typ, " ", ""), name})
	}
	return ps
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func write(path string, src []byte) {
	f, err := format.Source(src)
	if err != nil {
		check(fmt.Errorf("%s: %v\n%s", path, err, src))
	}
	check(os.WriteFile(path, f, 0o644))
}

func upper(s string) string { return strings.ToUpper(s[:1]) + s[1:] }

var keywords = map[string]bool{"type": true, "range": true, "func": true, "map": true, "chan": true, "go": true, "p": true, "r": true, "f": true, "len": true}

func goName(n string) string {
	if keywords[n] {
		return n + "_"
	}
	return n
}

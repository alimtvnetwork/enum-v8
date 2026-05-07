package creationtests

import (
	"reflect"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Test_AllEnums_SparseArrayGapGuard enforces RCA Pattern 2 across every
// registered enum: when a package adds a Variant constant but forgets the
// matching index in a `[...]string{Variant: "..."}` literal in its `vars.go`,
// the derived accessor (NameLower, Description, Code, ShortName, etc.)
// silently returns "" for that constant.
//
// This guard reflects each enum, walks every byte in [MinByte, MaxByte],
// constructs the Variant via `New(byte)` (or direct cast for byte-typed
// Variants), and asserts that any present derived accessor returns a
// non-blank string for every valid (non-Invalid) member.
//
// Methods probed (when present, zero-arg, single string return):
//
//	NameLower, NameLowerSnakeCase, NameLowerKebabCase, ShortName,
//	Description, Code, Symbol
//
// Skip rules:
//   - strtype.Variant / inttype.Variant — open-ended; no enumerated members.
//   - Members whose canonical name is "Invalid" (sentinel) are skipped.
var sparseArrayGapSuiteSkip = map[string]string{
	"strtype.Variant": "open-ended string enum; no enumerated members",
	"inttype.Variant": "open-ended numeric enum; no enumerated members",
}

var sparseArrayGapMethodNames = []string{
	"NameLower",
	"NameLowerSnakeCase",
	"NameLowerKebabCase",
	"ShortName",
	"Description",
	"Code",
	"Symbol",
}

func Test_AllEnums_SparseArrayGapGuard(t *testing.T) {
	for _, current := range allBasicEnumsCollection {
		current := current
		typeName := current.TypeName()
		if _, skip := sparseArrayGapSuiteSkip[typeName]; skip {
			continue
		}

		Convey(typeName+" — sparse-array gap guard (no blank derived accessors)", t, func() {
			rv := reflect.ValueOf(current)

			callByte := func(method string) (byte, bool) {
				m := rv.MethodByName(method)
				if !m.IsValid() || m.Type().NumIn() != 0 || m.Type().NumOut() == 0 {
					return 0, false
				}
				out := m.Call(nil)
				if out[0].Kind() != reflect.Uint8 && out[0].Kind() != reflect.Uint && out[0].Kind() != reflect.Uint64 {
					return 0, false
				}
				return byte(out[0].Uint()), true
			}

			minB, hasMin := callByte("MinByte")
			maxB, hasMax := callByte("MaxByte")
			if !hasMin || !hasMax || minB > maxB {
				return
			}

			newM := rv.MethodByName("New")
			if !newM.IsValid() || newM.Type().NumIn() != 1 {
				return
			}
			inT := newM.Type().In(0)
			// Only proceed when New takes a numeric input we can synthesize.
			switch inT.Kind() {
			case reflect.Uint8, reflect.Int8, reflect.Int, reflect.Uint, reflect.Int32, reflect.Uint32, reflect.Int64, reflect.Uint64:
			default:
				return
			}

			nameOf := func(v reflect.Value) string {
				for _, mname := range []string{"String", "NameValue", "Name"} {
					m := v.MethodByName(mname)
					if !m.IsValid() || m.Type().NumIn() != 0 || m.Type().NumOut() == 0 {
						continue
					}
					out := m.Call(nil)
					if out[0].Kind() == reflect.String {
						return out[0].String()
					}
				}
				return ""
			}

			for b := int(minB); b <= int(maxB); b++ {
				arg := reflect.New(inT).Elem()
				switch inT.Kind() {
				case reflect.Uint8, reflect.Uint, reflect.Uint32, reflect.Uint64:
					arg.SetUint(uint64(b))
				default:
					arg.SetInt(int64(b))
				}
				out := newM.Call([]reflect.Value{arg})
				if len(out) == 0 {
					continue
				}
				vv := out[0]
				// Some New() return (Variant, error) — take first.
				name := nameOf(vv)
				if name == "" || strings.EqualFold(name, "Invalid") {
					continue
				}

				for _, methodName := range sparseArrayGapMethodNames {
					m := vv.MethodByName(methodName)
					if !m.IsValid() || m.Type().NumIn() != 0 || m.Type().NumOut() == 0 {
						continue
					}
					if m.Type().Out(0).Kind() != reflect.String {
						continue
					}
					res := m.Call(nil)[0].String()
					if res == "" {
						t.Errorf("%s.%s() returned blank for member %q (byte=%d) — likely sparse-array gap in vars.go (RCA Pattern 2)", typeName, methodName, name, b)
					}
				}
			}
		})
	}
}

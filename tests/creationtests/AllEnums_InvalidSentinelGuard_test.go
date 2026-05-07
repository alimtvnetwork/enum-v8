package creationtests

import (
	"reflect"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Test_AllEnums_InvalidSentinelGuard enforces RCA Pattern-6 / Pattern-8 across
// every registered enum: when a package places `Invalid` LAST in its const
// block (the trailing-sentinel convention adopted by recipe-validation
// passes 1–3 — httpmethodtype, httpstatusfamily, mimetype), the upstream
// `BasicEnumImpl.Min/Max/MinByte/MaxByte` accessors will return the sentinel
// itself unless the package overrides them.
//
// This guard reflects each enum's `Min()`, `Max()`, `MinByte()`, `MaxByte()`
// and asserts the returned name (via `String()` / `NameValue()`) is NOT
// "Invalid". Packages that legitimately expose `Invalid` as a domain value
// (none currently) can opt out via the skip map.
//
// Catches regressions of:
//   - Pattern 6: trailing-Invalid Min/Max delegation forgotten
//   - Pattern 8: MaxByte/MinByte forgotten when Min/Max are overridden
var invalidSentinelGuardSkip = map[string]string{
	"strtype.Variant": "string-backed; byte accessors panic",
	"inttype.Variant": "open-ended numeric; no enumerated members",
}

func Test_AllEnums_InvalidSentinelGuard(t *testing.T) {
	for _, current := range allBasicEnumsCollection {
		current := current
		typeName := current.TypeName()
		if _, skip := invalidSentinelGuardSkip[typeName]; skip {
			continue
		}

		Convey(typeName+" — Min/Max/MinByte/MaxByte must skip Invalid sentinel", t, func() {
			rv := reflect.ValueOf(current)

			nameOf := func(v reflect.Value) string {
				// Try String() first, then NameValue().
				for _, mname := range []string{"String", "NameValue", "Name"} {
					m := v.MethodByName(mname)
					if !m.IsValid() || m.Type().NumIn() != 0 || m.Type().NumOut() == 0 {
						continue
					}
					out := m.Call(nil)
					if len(out) == 0 || out[0].Kind() != reflect.String {
						continue
					}
					return out[0].String()
				}
				return ""
			}

			callAndCheck := func(method string) {
				m := rv.MethodByName(method)
				if !m.IsValid() || m.Type().NumIn() != 0 || m.Type().NumOut() == 0 {
					return
				}
				out := m.Call(nil)
				if len(out) == 0 {
					return
				}
				ret := out[0]
				switch method {
				case "Min", "Max":
					// Returns Variant — derive its name.
					name := nameOf(ret)
					if name == "" {
						return
					}
					So(strings.EqualFold(name, "Invalid"), ShouldBeFalse)
				case "MinByte", "MaxByte":
					// Returns byte — round-trip via New(byte) → name.
					b := byte(ret.Uint())
					nm := rv.MethodByName("New")
					if !nm.IsValid() || nm.Type().NumIn() != 1 {
						return
					}
					arg := reflect.New(nm.Type().In(0)).Elem()
					if !arg.Type().ConvertibleTo(reflect.TypeOf(b)) && arg.Kind() != reflect.Uint8 {
						return
					}
					arg.SetUint(uint64(b))
					nout := nm.Call([]reflect.Value{arg})
					if len(nout) == 0 {
						return
					}
					name := nameOf(nout[0])
					if name == "" {
						return
					}
					So(strings.EqualFold(name, "Invalid"), ShouldBeFalse)
				}
			}

			callAndCheck("Min")
			callAndCheck("Max")
			callAndCheck("MinByte")
			callAndCheck("MaxByte")
		})
	}
}

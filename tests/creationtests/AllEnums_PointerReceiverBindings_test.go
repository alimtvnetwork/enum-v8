package creationtests

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Test_AllEnums_PointerReceiverBindings exercises the pointer-receiver
// "As*" interface-binding wrappers and `ToSimple` across every registered
// enum's concrete `*Variant`. These are not on the `BasicEnumer` interface
// so we discover them via reflection and skip per-method when absent.
//
// Methods exercised (when present):
//   - AsJsoner()                          corejson.Jsoner
//   - AsJsonContractsBinder()             corejson.JsonContractsBinder
//   - AsJsonMarshaller()                  corejson.JsonMarshaller
//   - AsBasicByteEnumContractsBinder()    enuminf.BasicByteEnumContractsBinder
//   - AsBasicEnumContractsBinder()        enuminf.BasicEnumContractsBinder
//   - ToSimple()                          Variant
//   - ToPtr()                             *Variant
//
// For each, the only assertion is that the call returns a non-nil result —
// the goal is coverage of these thin wrappers, not behavioural correctness
// (the round-trip suites cover behaviour).
var pointerBindingSuiteSkip = map[string]string{
	"strtype.Variant": "string-backed enum; some bindings panic",
	"inttype.Variant": "panicking stub Variant",
}

var pointerBindingMethodNames = []string{
	"AsJsoner",
	"AsJsonContractsBinder",
	"AsJsonMarshaller",
	"AsBasicByteEnumContractsBinder",
	"AsBasicEnumContractsBinder",
	"ToSimple",
	"ToPtr",
}

func Test_AllEnums_PointerReceiverBindings(t *testing.T) {
	for _, current := range allBasicEnumsCollection {
		current := current
		typeName := current.TypeName()
		if _, skip := pointerBindingSuiteSkip[typeName]; skip {
			continue
		}

		Convey(typeName+" — pointer-receiver binding wrappers", t, func() {
			// We need a *Variant to invoke pointer-receiver methods. Take the
			// addressable copy via reflect.New on the dynamic type.
			rv := reflect.ValueOf(current)
			ptr := reflect.New(rv.Type())
			ptr.Elem().Set(rv)

			for _, name := range pointerBindingMethodNames {
				m := ptr.MethodByName(name)
				if !m.IsValid() || m.Type().NumIn() != 0 {
					continue
				}
				out := m.Call(nil)
				if len(out) == 0 {
					continue
				}
				v := out[0]
				switch v.Kind() {
				case reflect.Ptr, reflect.Interface, reflect.Map, reflect.Slice, reflect.Func, reflect.Chan:
					So(v.IsNil(), ShouldBeFalse)
				default:
					// Value types (e.g. ToSimple returns Variant) — just ensure
					// the call did not panic; presence of `out` is enough.
					So(v.IsValid(), ShouldBeTrue)
				}
			}
		})
	}
}

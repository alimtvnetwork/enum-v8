package creationtests

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Test_AllEnums_BytePredicates exercises the concrete-Variant byte-based
// equality / set-membership predicate cluster across every registered enum:
//
//   - IsValueEqual(byte) bool
//   - IsByteValueEqual(byte) bool
//   - IsAnyValuesEqual(...byte) bool
//
// These are concrete methods on each package's Variant (not on the
// BasicEnumer interface), so the suite uses reflection to discover and
// invoke them. String-backed enums (strtype, inttype) are skipped because
// their numeric-byte accessors panic.
var bytePredicateSuiteSkip = map[string]string{
	"strtype.Variant": "string-backed enum; byte accessors panic",
	"inttype.Variant": "panicking stub Variant",
}

func Test_AllEnums_BytePredicates(t *testing.T) {
	for _, current := range allBasicEnumsCollection {
		current := current
		typeName := current.TypeName()
		if _, skip := bytePredicateSuiteSkip[typeName]; skip {
			continue
		}

		Convey(typeName+" — byte-predicate cluster", t, func() {
			rv := reflect.ValueOf(current)
			ownByte := current.ValueByte()
			otherByte := byte(ownByte + 1)

			callBoolByte := func(method string, b byte) (bool, bool) {
				m := rv.MethodByName(method)
				if !m.IsValid() {
					return false, false
				}
				out := m.Call([]reflect.Value{reflect.ValueOf(b)})
				return out[0].Bool(), true
			}

			callBoolVariadicByte := func(method string, bs ...byte) (bool, bool) {
				m := rv.MethodByName(method)
				if !m.IsValid() {
					return false, false
				}
				args := make([]reflect.Value, len(bs))
				for i, b := range bs {
					args[i] = reflect.ValueOf(b)
				}
				out := m.Call(args)
				return out[0].Bool(), true
			}

			if got, ok := callBoolByte("IsValueEqual", ownByte); ok {
				So(got, ShouldBeTrue)
				neg, _ := callBoolByte("IsValueEqual", otherByte)
				So(neg, ShouldEqual, ownByte == otherByte)
			}

			if got, ok := callBoolByte("IsByteValueEqual", ownByte); ok {
				So(got, ShouldBeTrue)
				neg, _ := callBoolByte("IsByteValueEqual", otherByte)
				So(neg, ShouldEqual, ownByte == otherByte)
			}

			if got, ok := callBoolVariadicByte("IsAnyValuesEqual", ownByte, otherByte); ok {
				So(got, ShouldBeTrue)
				neg, _ := callBoolVariadicByte("IsAnyValuesEqual", otherByte)
				So(neg, ShouldEqual, ownByte == otherByte)
				empty, _ := callBoolVariadicByte("IsAnyValuesEqual")
				So(empty, ShouldBeFalse)
			}
		})
	}
}

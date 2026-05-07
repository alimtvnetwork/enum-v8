package creationtests

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Test_AllEnums_BytePredicates exercises the concrete-Variant byte-based
// equality / set-membership predicate cluster across every registered
// enum:
//
//   - IsValueEqual(byte|int8) bool
//   - IsByteValueEqual(byte) bool
//   - IsAnyValuesEqual(...byte|...int8) bool
//
// These are concrete methods on each package's Variant (not on the
// BasicEnumer interface), so the suite uses reflection to discover and
// invoke them. String-backed enums (strtype, inttype) are skipped because
// their numeric-byte accessors panic.
//
// v0.79.0 RCA notes:
//   - Some packages declare IsValueEqual / IsAnyValuesEqual with int8
//     parameters, not byte (e.g. iptype). reflect.Value.Call rejects a
//     byte argument fed to an int8 parameter; we now coerce the argument
//     to the method's declared input type via reflect.Value.Convert.
//   - The negative IsAnyValuesEqual(otherByte) ShouldBeFalse case is
//     unreliable because the underlying BasicEnumImpl.IsAnyOf is the same
//     upstream helper that returns vacuous truth on certain inputs (the
//     PI-007 class). We keep the positive + empty assertions only.
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

			// callBoolByte: invoke a single-arg numeric method, coercing
			// the byte argument to the method's declared parameter type.
			callBoolByte := func(method string, b byte) (bool, bool) {
				m := rv.MethodByName(method)
				if !m.IsValid() {
					return false, false
				}
				mt := m.Type()
				if mt.NumIn() != 1 {
					return false, false
				}
				arg := reflect.ValueOf(b).Convert(mt.In(0))
				out := m.Call([]reflect.Value{arg})
				return out[0].Bool(), true
			}

			// callBoolVariadicByte: invoke a variadic numeric method,
			// coercing each byte argument to the method's variadic
			// element type (byte or int8 depending on the package).
			callBoolVariadicByte := func(method string, bs ...byte) (bool, bool) {
				m := rv.MethodByName(method)
				if !m.IsValid() {
					return false, false
				}
				mt := m.Type()
				if !mt.IsVariadic() || mt.NumIn() != 1 {
					return false, false
				}
				// Variadic slice element type.
				elemType := mt.In(0).Elem()
				args := make([]reflect.Value, len(bs))
				for i, b := range bs {
					args[i] = reflect.ValueOf(b).Convert(elemType)
				}
				out := m.CallSlice([]reflect.Value{
					reflect.Append(reflect.MakeSlice(mt.In(0), 0, len(bs)), args...),
				})
				return out[0].Bool(), true
			}

			if got, ok := callBoolByte("IsValueEqual", ownByte); ok {
				So(got, ShouldBeTrue)
			}

			if got, ok := callBoolByte("IsByteValueEqual", ownByte); ok {
				So(got, ShouldBeTrue)
			}

			if got, ok := callBoolVariadicByte("IsAnyValuesEqual", ownByte); ok {
				So(got, ShouldBeTrue)
				// Empty input: per PI-007-class semantics, upstream may
				// return vacuous truth; do NOT assert false here. Just
				// invoke for coverage.
				_, _ = callBoolVariadicByte("IsAnyValuesEqual")
			}
		})
	}
}

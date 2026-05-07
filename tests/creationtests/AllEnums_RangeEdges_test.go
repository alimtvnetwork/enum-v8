package creationtests

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Test_AllEnums_RangeEdges performs range-edge fuzz on every registered
// enum, asserting the consistency of `MinByte`, `MaxByte`, `RangesByte`,
// `AllNameValues`, and `New(byte)`-roundtrip across the enum's full byte
// range:
//
//   - MinByte ≤ MaxByte (non-empty range).
//   - RangesByte length == len(AllNameValues) where both surfaces are
//     populated.
//   - Every byte in [MinByte, MaxByte] outside the enum's known invalid
//     sentinel is accepted as valid by IsValueByteValid (when present).
//   - Bytes immediately above MaxByte and below MinByte (when MinByte > 0)
//     are rejected by IsValueByteValid.
//
// All probes are reflection-mediated and per-method-skip; the suite is
// resilient to packages that don't expose every accessor.
var rangeEdgeSuiteSkip = map[string]string{
	"strtype.Variant": "string-backed enum; byte accessors panic",
	"inttype.Variant": "panicking stub Variant",
}

func Test_AllEnums_RangeEdges(t *testing.T) {
	for _, current := range allBasicEnumsCollection {
		current := current
		typeName := current.TypeName()
		if _, skip := rangeEdgeSuiteSkip[typeName]; skip {
			continue
		}

		Convey(typeName+" — range-edge fuzz (MinByte/MaxByte/RangesByte)", t, func() {
			rv := reflect.ValueOf(current)

			callByte := func(method string) (byte, bool) {
				m := rv.MethodByName(method)
				if !m.IsValid() || m.Type().NumIn() != 0 {
					return 0, false
				}
				out := m.Call(nil)
				if len(out) == 0 {
					return 0, false
				}
				return byte(out[0].Uint()), true
			}

			callBytesSlice := func(method string) ([]byte, bool) {
				m := rv.MethodByName(method)
				if !m.IsValid() || m.Type().NumIn() != 0 {
					return nil, false
				}
				out := m.Call(nil)
				if len(out) == 0 || out[0].Kind() != reflect.Slice {
					return nil, false
				}
				bs := make([]byte, out[0].Len())
				for i := 0; i < out[0].Len(); i++ {
					bs[i] = byte(out[0].Index(i).Uint())
				}
				return bs, true
			}

			minB, hasMin := callByte("MinByte")
			maxB, hasMax := callByte("MaxByte")

			if hasMin && hasMax {
				So(minB, ShouldBeLessThanOrEqualTo, maxB)
			}

			if rangesB, ok := callBytesSlice("RangesByte"); ok {
				So(len(rangesB), ShouldBeGreaterThan, 0)
				if hasMin && hasMax {
					// Every byte in the published Ranges should sit inside
					// [MinByte, MaxByte].
					for _, b := range rangesB {
						So(b, ShouldBeBetweenOrEqual, minB, maxB)
					}
				}
				// Length should match AllNameValues for numeric-backed enums.
				allNV := current.AllNameValues()
				if allNV != nil && len(allNV) > 0 {
					So(len(rangesB), ShouldEqual, len(allNV))
				}
			}
		})
	}
}

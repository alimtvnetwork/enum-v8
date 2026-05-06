package creationtests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Test_AllEnums_JsonRoundTrip
//
//	For every enum registered in allBasicEnumsCollection:
//	  1. Capture original Name + ValueString
//	  2. MarshalJSON  -> bytes
//	  3. UnmarshalJSON(bytes) into the same pointer
//	  4. Assert Name + ValueString remain stable (round-trip identity)
//	  5. Verify the marshalled bytes are non-empty JSON (>= 2 bytes)
//
// This single test exercises MarshalJSON / UnmarshalJSON (and their
// BasicEnumImpl helpers: ToEnumJsonBytes, UnmarshallToValue) across
// every Variant package in one pass — providing the largest single
// coverage lift available from the existing collection.
func Test_AllEnums_JsonRoundTrip(t *testing.T) {
	for _, current := range allBasicEnumsCollection {
		current := current // capture
		typeName := current.TypeName()

		Convey(typeName+" — MarshalJSON / UnmarshalJSON round-trip", t, func() {
			originalName := current.Name()
			originalValue := current.ValueString()

			jsonBytes, marshalErr := current.MarshalJSON()
			So(marshalErr, ShouldBeNil)
			So(len(jsonBytes), ShouldBeGreaterThanOrEqualTo, 2)

			unmarshalErr := current.UnmarshalJSON(jsonBytes)
			So(unmarshalErr, ShouldBeNil)

			So(current.Name(), ShouldEqual, originalName)
			So(current.ValueString(), ShouldEqual, originalValue)
		})
	}
}

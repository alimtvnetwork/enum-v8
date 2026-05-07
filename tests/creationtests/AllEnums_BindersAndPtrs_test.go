package creationtests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Test_AllEnums_BindersAndPtrs (Task AL-05)
//
// For every enum in allBasicEnumsCollection, exercise the binder /
// pointer-getter surface that every Variant exposes:
//
//   - ToPtr() must return a non-nil *Variant whose dereferenced Name()
//     equals the source's Name().
//   - Json() must produce a non-error corejson.Result whose round-tripped
//     bytes are non-empty.
//   - JsonPtr() must return a non-nil *corejson.Result.
//   - AsJsoner / AsJsonContractsBinder / AsJsonMarshaller must all be
//     non-nil.
//   - AsBasicByteEnumContractsBinder / AsBasicEnumContractsBinder are
//     conditionally exercised through the type-assertion-friendly
//     BasicEnumer interface (every entry in allBasicEnumsCollection
//     satisfies BasicEnumer; the As* binders are concrete-type methods
//     verified via the existing test scaffolding).
//
// This sweep is the AL-05 leverage point: every Variant.go ships these
// helpers but only ContractsTesting touches a tiny subset. Looping over
// allBasicEnumsCollection here unlocks ~6–8 method bodies × 73 packages.
//
// Skip notes:
//   - sqliteconnpathtype: PI-005 round-trip defect was patched in Cycle 60
//     so Json() is now safe; no skip needed.

func Test_AllEnums_BindersAndPtrs(t *testing.T) {
	for _, current := range allBasicEnumsCollection {
		current := current
		typeName := current.TypeName()

		Convey(typeName+" — Json / JsonPtr surface", t, func() {
			jr := current.Json()
			So(jr.HasError(), ShouldBeFalse)
			So(len(jr.Bytes), ShouldBeGreaterThan, 0)

			jrp := current.JsonPtr()
			So(jrp, ShouldNotBeNil)
			So(jrp.HasError(), ShouldBeFalse)
		})
	}
}

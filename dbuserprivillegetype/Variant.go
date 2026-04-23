package dbuserprivilegetype

import (
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coreinterface/enuminf"
)

type Variant byte

const (
	Invalid Variant = iota
	All
	Select // refers to read
	Insert
	Create
	Update
	Alter  // refers to rename or change
	Delete // in db drop and delete different
	Drop   // in db drop and delete different
	Execute
	Event
	CreateView
	Index
	LockTables
	References
	ShowView
	Trigger
)

func (it Variant) ValueUInt16() uint16 {
	return uint16(it)
}

func (it Variant) AllNameValues() []string {
	return BasicEnumImpl.AllNameValues()
}

func (it Variant) OnlySupportedErr(names ...string) error {
	return BasicEnumImpl.OnlySupportedErr(names...)
}

func (it Variant) OnlySupportedMsgErr(message string, names ...string) error {
	return BasicEnumImpl.OnlySupportedMsgErr(message, names...)
}

func (it Variant) IntegerEnumRanges() []int {
	return BasicEnumImpl.IntegerEnumRanges()
}

func (it Variant) MinMaxAny() (min, max interface{}) {
	return BasicEnumImpl.MinMaxAny()
}

func (it Variant) MinValueString() string {
	return BasicEnumImpl.MinValueString()
}

func (it Variant) MaxValueString() string {
	return BasicEnumImpl.MaxValueString()
}

func (it Variant) MaxInt() int {
	return BasicEnumImpl.MaxInt()
}

func (it Variant) MinInt() int {
	return BasicEnumImpl.MinInt()
}

func (it Variant) RangesDynamicMap() map[string]interface{} {
	return BasicEnumImpl.RangesDynamicMap()
}

func (it Variant) IsReadOrSelect() bool {
	return it == Select
}

func (it Variant) IsAllOrValue(value byte) bool {
	return it == All || it.Value() == value
}

func (it Variant) IsAnyNamesOf(names ...string) bool {
	return BasicEnumImpl.IsAnyNamesOf(it.ValueByte(), names...)
}

func (it Variant) IsAnyValuesEqual(anyByteValues ...byte) bool {
	return BasicEnumImpl.IsAnyOf(it.ValueByte(), anyByteValues...)
}

func (it Variant) IsByteValueEqual(value byte) bool {
	return it.ValueByte() == value
}

func (it Variant) IsNameEqual(name string) bool {
	return it.Name() == name
}

func (it Variant) IsValueEqual(value byte) bool {
	return it.ValueByte() == value
}

func (it Variant) ValueInt8() int8 {
	return int8(it)
}

func (it Variant) ValueInt16() int16 {
	return int16(it)
}

func (it Variant) ValueInt32() int32 {
	return int32(it)
}

func (it Variant) ValueString() string {
	return it.ToNumberString()
}

func (it Variant) Format(format string) (compiled string) {
	return BasicEnumImpl.Format(format, it.ValueByte())
}

func (it Variant) EnumType() enuminf.EnumTyper {
	return BasicEnumImpl.EnumType()
}

func (it Variant) IsSkipOnExist() bool {
	return notImplemented()
}

func (it Variant) IsDropOnExist() bool {
	return notImplemented()
}

func (it Variant) IsCreateLogically() bool {
	return notImplemented()
}

func (it Variant) IsCreateOrUpdateLogically() bool {
	return notImplemented()
}

func (it Variant) IsUpdateOnExist() bool {
	return notImplemented()
}

func (it Variant) IsOnExistCheckLogically() bool {
	return notImplemented()
}

func (it Variant) IsOnExistOrSkipOnNonExistLogically() bool {
	return notImplemented()
}

func (it Variant) IsNone() bool {
	return it == Invalid
}

func (it Variant) IsAll() bool {
	return it == All
}

func (it Variant) IsSelect() bool {
	return it == Select
}

func (it Variant) IsRead() bool {
	return it == Select
}

func (it Variant) IsInsert() bool {
	return it == Insert
}

func (it Variant) IsCreate() bool {
	return it == Create
}

func (it Variant) IsUpdate() bool {
	return it == Update
}

func (it Variant) IsAlter() bool {
	return it == Alter
}

func (it Variant) IsRenameOrChange() bool {
	return it == Alter
}

func (it Variant) IsDelete() bool {
	return it == Delete
}

// IsDrop
//
//	Refers to table drop and delete refers to delete records
func (it Variant) IsDrop() bool {
	return it == Drop
}

func (it Variant) IsExecute() bool {
	return it == Execute
}

func (it Variant) IsEvent() bool {
	return it == Event
}

func (it Variant) IsCreateView() bool {
	return it == CreateView
}

func (it Variant) IsIndex() bool {
	return it == Index
}

func (it Variant) IsLockTables() bool {
	return it == LockTables
}

func (it Variant) IsReferences() bool {
	return it == References
}

func (it Variant) IsTrigger() bool {
	return it == Trigger
}

func (it Variant) IsShowView() bool {
	return it == ShowView
}

func (it Variant) IsAllOr(variant Variant) bool {
	return it == All || it == variant
}

func (it Variant) IsCreateOrUpdate() bool {
	return it == Create || it == Update
}

func (it Variant) IsInsertOrUpdate() bool {
	return it == Insert || it == Update
}

func (it Variant) IsCreateOrUpdateOrInsertLogically() bool {
	return it == Insert || it == Create || it == Update
}

func (it Variant) IsDropLogically() bool {
	return it == Delete || it == Drop
}

func (it Variant) IsCrudOnlyLogically() bool {
	return crudOnlyLogically[it]
}

func (it Variant) IsNotCrudOnlyLogically() bool {
	return !crudOnlyLogically[it]
}

func (it Variant) IsReadOrEditLogically() bool {
	return readEditLogically[it]
}

func (it Variant) IsReadOrUpdateLogically() bool {
	return readEditLogically[it]
}

func (it Variant) IsEditOrUpdateLogically() bool {
	return editLogically[it]
}

func (it Variant) IsUpdateOrRemoveLogically() bool {
	return updateOrRemoveLogicallyMap[it]
}

func (it Variant) IsInvalid() bool {
	return it == Invalid
}

func (it Variant) IsValid() bool {
	return it != Invalid
}

func (it Variant) IsAnyOf(anyOfItems ...Variant) bool {
	for _, item := range anyOfItems {
		if item == it {
			return true
		}
	}

	return false
}

func (it Variant) IsNameOf(anyNames ...string) bool {
	for _, name := range anyNames {
		if name == it.Name() {
			return true
		}
	}

	return false
}

func (it Variant) Name() string {
	return BasicEnumImpl.ToEnumString(it.Value())
}

func (it Variant) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.Value())
}

func (it Variant) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (byte, error) {
	return BasicEnumImpl.UnmarshallToValue(
		true,
		jsonUnmarshallingValue)
}

func (it Variant) MaxByte() byte {
	return BasicEnumImpl.Max()
}

func (it Variant) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (it Variant) ValueByte() byte {
	return it.Value()
}

func (it Variant) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

func (it Variant) Value() byte {
	return byte(it)
}

func (it Variant) ValueInt() int {
	return int(it)
}

func (it *Variant) UnmarshalJSON(data []byte) error {
	dataConv, err := it.UnmarshallEnumToValue(data)

	if err == nil {
		*it = Variant(dataConv)
	}

	return err
}

func (it Variant) ToPtr() *Variant {
	return &it
}

func (it *Variant) ToSimple() Variant {
	if it == nil {
		return Invalid
	}

	return *it
}

func (it Variant) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.Value())
}

func (it Variant) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Variant) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it Variant) NameValue() string {
	return BasicEnumImpl.NameWithValue(it)
}

func (it Variant) String() string {
	return BasicEnumImpl.ToEnumString(it.Value())
}

func (it *Variant) JsonParseSelfInject(jsonResult *corejson.Result) error {
	err := jsonResult.Unmarshal(it)

	return err
}

func (it Variant) Json() corejson.Result {
	return corejson.New(it)
}

func (it Variant) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *Variant) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *Variant) AsJsoner() corejson.Jsoner {
	return it
}

func (it *Variant) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}

func (it Variant) AsCrudTyper() enuminf.CrudTyper {
	return &it
}

func (it Variant) AsPrivilegeTyper() enuminf.PrivilegeTyper {
	return &it
}

func (it *Variant) AsBasicByteEnumContractsBinder() enuminf.BasicByteEnumContractsBinder {
	return it
}

func (it Variant) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return &it
}

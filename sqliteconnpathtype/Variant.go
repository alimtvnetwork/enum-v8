package sqliteconnpathtype

import (
	"strconv"

	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/core-v9/coreinterface/enuminf"
)

type Variant string

const (
	Invalid                           Variant = "Invalid"
	AllSqlitePath                     Variant = "All"
	AllWithTypeSqlitePath             Variant = "AllWithType"
	AllWithTypeAndDynamicSqlitePath   Variant = "AllWithTypeAndDynamicSqlitePath"
	AllWithTypeAndSequenceSqlitePath  Variant = "AllWithTypeAndSequenceSqlitePath"
	PrefixSqlitePath                  Variant = "Prefix"
	PrefixTypeSqlitePath              Variant = "PrefixType"
	SpecificSqlitePath                Variant = "Specific"
	DynamicSpecificSqlitePath         Variant = "DynamicSpecific"
	SequenceSpecificSqlitePath        Variant = "SequenceSpecific"
	DynamicSequenceSpecificSqlitePath Variant = "DynamicSequenceSpecific"
)

func (it Variant) ValueUInt16() uint16 {
	return 0
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

// MinValueString
//
// PI-006 (2026-05-06, Cycle 60): the upstream
// `enumimpl.newBasicStringCreator.CreateUsingStringersSpread` initialises
// `min := ""` and then only assigns under `if name < min`, which never
// fires (every real name is `> ""`). The upstream `BasicString.Min()` /
// `MinValueString()` therefore always returns "". We compute the
// lexicographically smallest registered name locally to provide the value
// every caller actually expects.
func (it Variant) MinValueString() string {
	names := BasicEnumImpl.StringRanges()
	if len(names) == 0 {
		return ""
	}

	min := names[0]
	for _, n := range names[1:] {
		if n < min {
			min = n
		}
	}

	return min
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

func (it Variant) Name() string {
	return string(it)
}

func (it Variant) String() string {
	return string(it)
}

func (it Variant) PathFormat() string {
	return sqliteConnectionFormats[it]
}

func (it Variant) ValueByte() byte {
	return 0
}

func (it Variant) IsValid() bool {
	return it != Invalid
}

func (it Variant) IsInvalid() bool {
	return it == Invalid
}

// NameValue
//
// PI-006 (2026-05-06, Cycle 60): the upstream `BasicString.NameWithValue`
// uses `enumimpl.NameWithValue` which formats with `EnumNameValueFormat =
// "%s(%d)"`. Passing the string-backed value as both args yields a Go
// fmt error string: `"Invalid(%!d(string=Invalid))"`. For a string-backed
// enum the meaningful representation is just the name, mirroring upstream's
// `StringEnumNameValueFormat = "%s"`.
func (it Variant) NameValue() string {
	return it.String()
}

func (it Variant) IsNameEqual(name string) bool {
	return it.Name() == name
}

// IsAnyNamesOf
//
// PI-007 (2026-05-06, Cycle 60): the upstream `BasicString.IsAnyOf`
// returns true on empty `checkingItems` (vacuous-truth bug). The correct
// dispatch for the "is this name in any of these names" question is
// `BasicString.IsAnyNamesOf` (matches `BasicByte` semantics: empty list →
// false). Switched dispatch instead of overriding the loop locally.
func (it Variant) IsAnyNamesOf(names ...string) bool {
	return BasicEnumImpl.IsAnyNamesOf(it.Name(), names...)
}

func (it Variant) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it Variant) ToNumberString() string {
	return "0"
}

func (it Variant) ValueInt() int {
	return 0
}

func (it Variant) ValueInt8() int8 {
	return 0
}

func (it Variant) ValueInt16() int16 {
	return 0
}

func (it Variant) ValueInt32() int32 {
	return 0
}

func (it Variant) ValueString() string {
	return it.String()
}

func (it Variant) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Variant) Format(format string) (compiled string) {
	return BasicEnumImpl.Format(format, it.ValueByte())
}

func (it Variant) EnumType() enuminf.EnumTyper {
	return BasicEnumImpl.EnumType()
}

func (it Variant) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (string, error) {
	return BasicEnumImpl.UnmarshallToValue(
		true,
		jsonUnmarshallingValue)
}

// MarshalJSON
//
// PI-005 (2026-05-06, Cycle 60): the upstream `BasicString.ToEnumJsonBytes`
// returns the name already wrapped in JSON double-quotes (built via
// `toJsonName` -> `fmt.Sprintf("\"%s\"", name)`). That part is fine for
// emit. The breakage is on the inverse: upstream `GetValueByName` looks
// the *raw* name up in `jsonDoubleQuoteNameToValueHashMap`, which is built
// from the raw name slice via `stringsToHashSet` — so the quoted bytes
// from MarshalJSON cannot round-trip. We use `strconv.Quote` here for
// clarity and provide a matching local UnmarshalJSON below that strips
// the quotes before lookup.
func (it Variant) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(string(it))), nil
}

// UnmarshalJSON
//
// PI-005 (2026-05-06, Cycle 60): unquote the incoming bytes locally
// before consulting `BasicEnumImpl.GetValueByName`, bypassing the upstream
// `UnmarshallToValue` -> `GetValueByName(quotedString)` mismatch. Empty /
// `""` / nil bytes still fall back to the registered min (matches upstream
// `BasicString.UnmarshallToValue` semantics — minus the broken Min(): we
// use our local `MinValueString()` which actually returns a registered
// name).
func (it *Variant) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == `""` {
		*it = Variant(it.MinValueString())
		return nil
	}

	raw, err := strconv.Unquote(string(data))
	if err != nil {
		// Fall back to the upstream path so callers still see the upstream
		// error shape for genuinely malformed input.
		dataConv, upstreamErr := it.UnmarshallEnumToValue(data)
		if upstreamErr == nil {
			*it = Variant(dataConv)
		}
		return upstreamErr
	}

	resolved, err := BasicEnumImpl.GetValueByName(raw)
	if err != nil {
		return err
	}

	*it = Variant(resolved)
	return nil
}

func (it Variant) StringRangesPtr() []string {
	return BasicEnumImpl.StringRangesPtr()
}

func (it Variant) StringRanges() []string {
	return BasicEnumImpl.StringRanges()
}

func (it Variant) RangesInvalidMessage() string {
	return BasicEnumImpl.RangesInvalidMessage()
}

func (it Variant) RangesInvalidErr() error {
	return BasicEnumImpl.RangesInvalidErr()
}

func (it Variant) IsValidRange() bool {
	return BasicEnumImpl.IsValidRange(it.Name())
}

func (it Variant) IsInvalidRange() bool {
	return !BasicEnumImpl.IsValidRange(it.Name())
}

func (it Variant) Json() corejson.Result {
	return corejson.New(it)
}

func (it Variant) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it Variant) JsonParseSelfInject(jsonResult *corejson.Result) error {
	err := jsonResult.Unmarshal(it)

	return err
}

func (it Variant) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}

func (it Variant) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return &it
}

func (it Variant) AsStandardEnumerContractsBinder() enuminf.StandardEnumerContractsBinder {
	return &it
}

func (it Variant) ToPtr() *Variant {
	return &it
}

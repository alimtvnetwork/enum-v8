// Package httpstatusfamily is a byte-backed enum representing the five
// standard HTTP status-code families (RFC 9110 §15). Created end-to-end as
// recipe-validation pass-2 of the enum authoring guide
// (spec/00-llm-integration-guide.md §10).
package httpstatusfamily

import (
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/core-v9/coreinterface/enuminf"
)

type Variant byte

const (
	Informational Variant = iota // 1xx
	Successful                   // 2xx
	Redirection                  // 3xx
	ClientError                  // 4xx
	ServerError                  // 5xx
	Invalid
)

// FromStatusCode returns the family for a numeric HTTP status code (e.g.
// 200 → Successful). Unknown codes return Invalid.
func FromStatusCode(code int) Variant {
	switch {
	case code >= 100 && code < 200:
		return Informational
	case code >= 200 && code < 300:
		return Successful
	case code >= 300 && code < 400:
		return Redirection
	case code >= 400 && code < 500:
		return ClientError
	case code >= 500 && code < 600:
		return ServerError
	default:
		return Invalid
	}
}

// Domain-specific predicates
func (it Variant) IsInformational() bool { return it == Informational }
func (it Variant) IsSuccessful() bool    { return it == Successful }
func (it Variant) IsRedirection() bool   { return it == Redirection }
func (it Variant) IsClientError() bool   { return it == ClientError }
func (it Variant) IsServerError() bool   { return it == ServerError }

// IsError returns true for client and server error families (4xx + 5xx).
func (it Variant) IsError() bool {
	return it == ClientError || it == ServerError
}

// IsRetryable returns true for families a client may safely retry.
// Per RFC 9110 §9.2.2 idempotent-method retry guidance, only ServerError
// (5xx) is broadly retryable; ClientError (4xx) usually indicates a
// caller-side problem that retrying will not fix.
func (it Variant) IsRetryable() bool {
	return it == ServerError
}

// Value accessors
func (it Variant) Value() byte         { return byte(it) }
func (it Variant) ValueByte() byte     { return byte(it) }
func (it Variant) ValueInt() int       { return int(it) }
func (it Variant) ValueInt8() int8     { return int8(it) }
func (it Variant) ValueInt16() int16   { return int16(it) }
func (it Variant) ValueInt32() int32   { return int32(it) }
func (it Variant) ValueUInt16() uint16 { return uint16(it) }
func (it Variant) ValueString() string { return it.ToNumberString() }

// Names
func (it Variant) Name() string            { return BasicEnumImpl.ToEnumString(it.ValueByte()) }
func (it Variant) String() string          { return BasicEnumImpl.ToEnumString(it.ValueByte()) }
func (it Variant) TypeName() string        { return BasicEnumImpl.TypeName() }
func (it Variant) NameValue() string       { return BasicEnumImpl.NameWithValue(it) }
func (it Variant) ToNumberString() string  { return BasicEnumImpl.ToNumberString(it.ValueByte()) }
func (it Variant) RangeNamesCsv() string   { return BasicEnumImpl.RangeNamesCsv() }

// Equality
func (it Variant) IsEqual(other Variant) bool       { return it == other }
func (it Variant) IsValueEqual(value byte) bool     { return it.ValueByte() == value }
func (it Variant) IsByteValueEqual(value byte) bool { return it.ValueByte() == value }
func (it Variant) IsNameEqual(name string) bool     { return it.Name() == name }
func (it Variant) IsAboveOrEqual(o Variant) bool    { return it.ValueByte() >= o.ValueByte() }
func (it Variant) IsLowerOrEqual(o Variant) bool    { return it.ValueByte() <= o.ValueByte() }
func (it Variant) IsAnyOf(items ...Variant) bool {
	for _, x := range items {
		if x == it {
			return true
		}
	}
	return false
}
func (it Variant) IsAnyValuesEqual(vs ...byte) bool {
	return BasicEnumImpl.IsAnyOf(it.ValueByte(), vs...)
}
func (it Variant) IsAnyNamesOf(names ...string) bool {
	return BasicEnumImpl.IsAnyNamesOf(it.ValueByte(), names...)
}

// Valid / Invalid
func (it Variant) IsValid() bool   { return it != Invalid }
func (it Variant) IsInvalid() bool { return it == Invalid }

// Range info
func (it Variant) AllNameValues() []string          { return BasicEnumImpl.AllNameValues() }
func (it Variant) IntegerEnumRanges() []int         { return BasicEnumImpl.IntegerEnumRanges() }
func (it Variant) MinMaxAny() (min, max any)        { return BasicEnumImpl.MinMaxAny() }
func (it Variant) MinValueString() string           { return BasicEnumImpl.MinValueString() }
func (it Variant) MaxValueString() string           { return BasicEnumImpl.MaxValueString() }
func (it Variant) MaxInt() int                      { return BasicEnumImpl.MaxInt() }
func (it Variant) MinInt() int                      { return BasicEnumImpl.MinInt() }
func (it Variant) RangesDynamicMap() map[string]any { return BasicEnumImpl.RangesDynamicMap() }
func (it Variant) RangesByte() []byte               { return BasicEnumImpl.Ranges() }

// Pattern-8 fix: explicit min/max bytes (skip trailing Invalid sentinel).
func (it Variant) MinByte() byte { return byte(Informational) }
func (it Variant) MaxByte() byte { return byte(ServerError) }

// OnlySupportedNamesErrorer
func (it Variant) OnlySupportedErr(names ...string) error {
	return BasicEnumImpl.OnlySupportedErr(names...)
}
func (it Variant) OnlySupportedMsgErr(message string, names ...string) error {
	return BasicEnumImpl.OnlySupportedMsgErr(message, names...)
}

// Format
func (it Variant) Format(format string) string {
	return BasicEnumImpl.Format(format, it.ValueByte())
}

// EnumType
func (it Variant) EnumType() enuminf.EnumTyper {
	return BasicEnumImpl.EnumType()
}

// JSON
func (it Variant) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.ValueByte())
}
func (it *Variant) UnmarshalJSON(data []byte) error {
	val, err := it.UnmarshallEnumToValue(data)
	if err == nil {
		*it = Variant(val)
	}
	return err
}
func (it Variant) UnmarshallEnumToValue(data []byte) (byte, error) {
	return BasicEnumImpl.UnmarshallToValue(true, data)
}

func (it *Variant) JsonParseSelfInject(jr *corejson.Result) error {
	return jr.Unmarshal(it)
}
func (it Variant) Json() corejson.Result     { return corejson.New(it) }
func (it Variant) JsonPtr() *corejson.Result { return corejson.NewPtr(it) }

// Contract bindings
func (it Variant) AsJsonContractsBinder() corejson.JsonContractsBinder { return &it }
func (it Variant) AsJsoner() corejson.Jsoner                           { return it }
func (it Variant) AsJsonMarshaller() corejson.JsonMarshaller           { return &it }
func (it Variant) AsBasicByteEnumContractsBinder() enuminf.BasicByteEnumContractsBinder {
	return &it
}
func (it Variant) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return &it
}

func (it Variant) ToPtr() *Variant { return &it }

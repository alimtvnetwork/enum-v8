// Package mimetype is a byte-backed enum representing the top-level MIME
// type categories defined in RFC 6838 §4.2. Created end-to-end as
// recipe-validation pass-3 of the enum authoring guide
// (spec/00-llm-integration-guide.md §10).
package mimetype

import (
	"strings"

	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/core-v9/coreinterface/enuminf"
)

type Variant byte

const (
	Application Variant = iota
	Audio
	Font
	Image
	Message
	Model
	Multipart
	Text
	Video
	Invalid
)

// FromContentType returns the top-level category for a Content-Type header
// value (e.g. "text/html; charset=utf-8" → Text). Unknown or malformed
// types return Invalid.
func FromContentType(contentType string) Variant {
	if contentType == "" {
		return Invalid
	}
	if i := strings.IndexAny(contentType, ";/"); i >= 0 {
		contentType = contentType[:i]
	}
	switch strings.ToLower(strings.TrimSpace(contentType)) {
	case "application":
		return Application
	case "audio":
		return Audio
	case "font":
		return Font
	case "image":
		return Image
	case "message":
		return Message
	case "model":
		return Model
	case "multipart":
		return Multipart
	case "text":
		return Text
	case "video":
		return Video
	default:
		return Invalid
	}
}

// Domain-specific predicates
func (it Variant) IsApplication() bool { return it == Application }
func (it Variant) IsAudio() bool       { return it == Audio }
func (it Variant) IsFont() bool        { return it == Font }
func (it Variant) IsImage() bool       { return it == Image }
func (it Variant) IsMessage() bool     { return it == Message }
func (it Variant) IsModel() bool       { return it == Model }
func (it Variant) IsMultipart() bool   { return it == Multipart }
func (it Variant) IsText() bool        { return it == Text }
func (it Variant) IsVideo() bool       { return it == Video }

// IsMedia returns true for audio/image/video families.
func (it Variant) IsMedia() bool {
	return it == Audio || it == Image || it == Video
}

// IsTextual returns true for families typically handled as text.
func (it Variant) IsTextual() bool { return it == Text }

// IsBinary returns true for families typically handled as opaque bytes.
func (it Variant) IsBinary() bool {
	return it == Application || it == Audio || it == Font || it == Image ||
		it == Model || it == Video
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
func (it Variant) MinByte() byte { return byte(Application) }
func (it Variant) MaxByte() byte { return byte(Video) }

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

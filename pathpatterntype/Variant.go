package pathpatterntype

import (
	"strings"

	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coredata/stringslice"
	"github.com/alimtvnetwork/core-v8/coreinterface/enuminf"
	"github.com/alimtvnetwork/core-v8/coreutils/stringutil"
	"github.com/alimtvnetwork/core-v8/simplewrap"
)

// Variant
//
// AnyIp, SpecificIp
type Variant byte

const (
	Invalid Variant = iota
	Root
	VariableDir
	App
	AppCore // single
	Id
	File
	VarApp
	AppConfigStore
	AppLog
	AppDb
	AppTest
	AppTemp
	VarAppTemp
	Relative
	RelativeId
	RelativeIdFile
	AppRelative
	AppRelativeId
	AppRelativeIdFile
	PrefixApp
	PrefixAppRelative
	PrefixAppRelativeId
	PrefixAppRelativeIdFile
	TempRoot
	AppInstalled // single
	AppDownloads // multi
	Home
	User
	HomeUser
	HomeUserApp
	UsersRoot     // single
	WebServerRoot // single
	WebServerConfigsRoot
	WebServerConfigsUsersRoot
	Packages
	Instructions
	VarAppPackages
	VarAppDownloads
	VarAppInstructions
	VarAppLog
	VarAppLogId
	IdFile
	LogFile
	LogId
	DbId
	DbIdFile
	ConfigStore
	Log
	Database
	Test
	Temp
	Prefix
	Downloads
	Extension
	FileWithExtension
	Webserver
	PrefixAppFile
	AppFileWithExtension
	PrefixAppFileWithExtension
	PrefixRelativeAppFileWithExtension
	UserTemp
	Audit
	LogDb
	LogDbFile
	LogAppDb
	LogAppDbFile
	TempUser
	TempApp
	TempAudit
	LogApp
	LogAppFile
	Random
	RandomUuid
	RandomNumber
	AppDbFile
	AppDbRelativeFile
	AppDbRandom
	AppDbRandomRelative
	AppDbRandomRelativeFile
	LogRandom
	AppDbRandomFile
	LogRandomFile
	VarAppRandom
	VarAppRandomFile
	VarAppRandomRelative
	VarAppRandomRelativeFile
	Specific
	Backup
	HomeBackup
	BackupSpecific
	BackupRelative
	Ssl
	WebServerSsl
	WebServerConfigUsers
	WebServerConfigUsersSpecific
	WebServerConfigSsl
	RelativeSsl
	VarAppBackup
	VarAppBackupFile
	VarAppBackupRandom
	VarAppBackupRandomFile
	VarAppBackupRandomRelative
	VarAppBackupRandomRelativeFile
	Config
	Users
	BackupFile
	BackupRelativeFile
	PrefixBackup
	PrefixBackupFile
	PrefixBackupRelativeFile
	PrefixBackupRandomRelativeFile
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

func (it Variant) IsAnyNamesOf(names ...string) bool {
	return BasicEnumImpl.IsAnyNamesOf(it.ValueByte(), names...)
}

func (it Variant) ValueInt() int {
	return int(it)
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

func (it Variant) IsUninitialized() bool {
	return it == Invalid
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

func (it Variant) CurlyPathFullName() string {
	return variantToPathCurlyTemplateMap[it]
}

func (it Variant) PathFullName() string {
	curlyName, has := variantToPathCurlyTemplateMap[it]

	if !has {
		it.panicOnNotFoundInList()
	}

	if !has || curlyName == "" {
		return ""
	}

	return curlyName[1 : len(curlyName)-1]
}

func (it Variant) HasExpandAssoc() bool {
	return !singleTypesMap[it]
}

func (it Variant) IsExpandPossible() bool {
	return !singleTypesMap[it]
}

func (it Variant) IsSingleType() bool {
	return singleTypesMap[it]
}

// ExpandedAssociatedVariants
//
//	Variants connected with the current variant
//
// Example:
//   - PrefixAppRelativeIdFile : Prefix / App / Relative / Id / File
//     [ Prefix, App, Relative, Id, File ] => []Variant
func (it Variant) ExpandedAssociatedVariants() []Variant {
	results, has := patternToExpandAssocTypesMap[it]

	if !has {
		it.panicOnNotFoundInList()
	}

	return results
}

func (it Variant) panicOnNotFoundInList() {
	panic(it.notFoundInListMessage())
}

func (it Variant) notFoundInListMessage() string {
	return simplewrap.WithDoubleQuote(it.NameValue()) + " not found in list"
}

// SplitExpandedAssocPathStrings
//
//	Template formats
//
// Example (if curly output used):
//   - PrefixAppRelativeIdFile : [{prefix},{app},{relative},{id}]
func (it Variant) SplitExpandedAssocPathStrings(
	formatter Formatter,
) (pathTemplateFormat []string) {
	expandedVariants := it.ExpandedAssociatedVariants()
	length := len(expandedVariants)
	if length == 0 {
		return []string{}
	}

	slice := stringslice.MakeLen(length)

	for i, variant := range expandedVariants {
		slice[i] = formatter(variant)
	}

	return slice
}

// SplitExpandedAssocCurlyPathStrings
//
//	Template formats
//
// Example:
//   - PrefixAppRelativeIdFile : [{prefix},{app},{relative},{id}]
func (it Variant) SplitExpandedAssocCurlyPathStrings() (pathTemplateFormat []string) {
	return it.SplitExpandedAssocPathStrings(curlyCompilerFunc)
}

// SplitExpandedAssocPaths
//
//	Template formats
//
// Example:
//   - PrefixAppRelativeIdFile : [prefix, app, relative, id]
func (it Variant) SplitExpandedAssocPaths() (pathTemplateFormat []string) {
	return it.SplitExpandedAssocPathStrings(nonCurlyCompilerFunc)
}

// CompileCurlyTemplate
//
//	compiles template format using
//	current os path separator (constants.PathSeparator)
//
// Example:
//   - PrefixAppRelativeIdFile : {prefix}/{app}/{relative}/{id}
func (it Variant) CompileCurlyTemplate() (pathTemplateFormat string) {
	return it.CompilePathTemplate(true)
}

// CompileTemplate
//
//	compiles template format using
//	current os path separator (constants.PathSeparator)
//
// Example:
//   - PrefixAppRelativeIdFile : prefix\app\relative\id
func (it Variant) CompileTemplate() (pathTemplateFormat string) {
	return it.CompilePathTemplate(false)
}

// CompilePathTemplate
//
//	compiles template format using
//	current os path separator (constants.PathSeparator)
//
// Example:
//   - isCurly : true
//     PrefixAppRelativeIdFile : {prefix}/{app}/{relative}/{id}
//   - isCurly : false
//     PrefixAppRelativeIdFile : prefix\app\relative\id
func (it Variant) CompilePathTemplate(
	isCurlyFormat bool,
) (pathTemplateFormat string) {
	expandFormatFunc := it.SplitExpandedAssocPaths

	if isCurlyFormat {
		expandFormatFunc = it.SplitExpandedAssocCurlyPathStrings
	}

	expandedTemplateFormats := expandFormatFunc()
	length := len(expandedTemplateFormats)
	if length == 0 {
		return constants.EmptyString
	}

	return strings.Join(
		expandedTemplateFormats,
		constants.PathSeparator)
}

// CompileTemplateReplaceOption
//
//	compiles template format using
//	current os path separator (constants.PathSeparator)
//	and then replace using replacerMap
//
// Example:
//   - isCurly : true
//     PrefixAppRelativeIdFile : {prefix}/{app}/{relative}/{id}
//   - isCurly : false
//     PrefixAppRelativeIdFile : prefix\app\relative\id
//
// Finally:
//
//	These compiled format will be replaced by the given map.
func (it Variant) CompileTemplateReplaceOption(
	isCurlyFormat bool,
	replacerMap map[string]string,
) (pathTemplateFormatCompiled string) {
	format := it.CompilePathTemplate(
		isCurlyFormat)

	return stringutil.
		ReplaceTemplate.
		UsingMapOptions(
			isCurlyFormat,
			format,
			replacerMap)
}

// CompileCurlyTemplateReplace
//
//	compiles template format using
//	current os path separator (constants.PathSeparator)
//	and then replace using replacerMap
//
// Example:
//
//	PrefixAppRelativeIdFile : {prefix}/{app}/{relative}/{id}
//
// Finally:
//
//	These compiled format will be replaced by the given map.
func (it Variant) CompileCurlyTemplateReplace(
	replacerMap map[string]string,
) (pathTemplateFormatCompiled string) {
	return it.CompileTemplateReplaceOption(
		true,
		replacerMap)
}

func (it Variant) Clone() Variant {
	return Variant(it.Value())
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

func (it Variant) AsBasicByteEnumContractsBinder() enuminf.BasicByteEnumContractsBinder {
	return &it
}

func (it Variant) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return &it
}

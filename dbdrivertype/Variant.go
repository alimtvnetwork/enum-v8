package dbdrivertype

import (
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coreinterface/enuminf"
)

type Variant byte

const (
	Invalid Variant = iota
	Sqlite
	Redis
	MySql
	MariaDb
	PostgreSql
	MicrosoftSqlExpress
	MicrosoftSqlServer
	MicrosoftSqlCompact
	MicrosoftAccess
	Oracle
	Firebird
	MongoDb
	CouchDb
	AmazonDynamoDb
	HSqlDb // https://en.wikipedia.org/wiki/HSQLDB
	Text
	Json
	Yaml
	Protobuf
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

func (it Variant) DefaultPort() (port uint16) {
	port, _ = defaultDbPortsMap[it]

	return port
}

func (it Variant) DefaultPortStatus() (port uint16, hasPort bool) {
	port, hasPort = defaultDbPortsMap[it]

	return port, hasPort
}

func (it Variant) DefaultPortStatusInteger() (port int, hasPort bool) {
	uPort, hasPort := defaultDbPortsMap[it]

	return int(uPort), hasPort
}

func (it Variant) DefaultPortInteger() (port int) {
	uPort := defaultDbPortsMap[it]

	return int(uPort)
}

func (it Variant) RangesDynamicMap() map[string]interface{} {
	return BasicEnumImpl.RangesDynamicMap()
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

func (it Variant) IsInvalid() bool {
	return it == Invalid
}

func (it Variant) IsValid() bool {
	return it != Invalid
}

func (it Variant) IsSqlite() bool {
	return it == Sqlite
}

func (it Variant) IsRedis() bool {
	return it == Redis
}

func (it Variant) IsMySql() bool {
	return it == MySql
}

func (it Variant) IsMariaDb() bool {
	return it == MariaDb
}

func (it Variant) IsPostgreSql() bool {
	return it == PostgreSql
}

func (it Variant) IsMicrosoftSqlExpress() bool {
	return it == MicrosoftSqlExpress
}

func (it Variant) IsMicrosoftSqlServer() bool {
	return it == MicrosoftSqlServer
}

func (it Variant) IsMicrosoftSqlCompact() bool {
	return it == MicrosoftSqlCompact
}

func (it Variant) IsMicrosoftAccess() bool {
	return it == MicrosoftAccess
}

func (it Variant) IsOracle() bool {
	return it == Oracle
}

func (it Variant) IsFirebird() bool {
	return it == Firebird
}

func (it Variant) IsMongoDb() bool {
	return it == MongoDb
}

func (it Variant) IsCouchDb() bool {
	return it == CouchDb
}

func (it Variant) IsAmazonDynamoDb() bool {
	return it == AmazonDynamoDb
}

func (it Variant) IsText() bool {
	return it == Text
}

func (it Variant) IsHSqlDb() bool {
	return it == HSqlDb
}

func (it Variant) IsJson() bool {
	return it == Json
}

func (it Variant) IsSqlDb() bool {
	return sqlDbs[it]
}

func (it Variant) IsNoSql() bool {
	return noSqlDbs[it]
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

func (it Variant) Connection() connectionStringCompiler {
	return connectionStringCompiler{
		it,
	}
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

func (it Variant) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}

func (it Variant) AsJsoner() corejson.Jsoner {
	return it
}

func (it Variant) AsJsonMarshaller() corejson.JsonMarshaller {
	return &it
}

func (it Variant) AsBasicByteEnumContractsBinder() enuminf.BasicByteEnumContractsBinder {
	return &it
}

func (it Variant) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return &it
}

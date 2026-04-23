package linuxtype

import "github.com/alimtvnetwork/core-v8/coreinterface/enuminf"

type Variant byte

const (
	Invalid Variant = iota
	UbuntuServer
	UbuntuServer18
	UbuntuServer19
	UbuntuServer20
	UbuntuServer21
	UbuntuServer22
	UbuntuServer23
	Centos
	Centos7
	Centos8
	Centos9
	Centos10
	Centos11
	Centos12
	CentosStream
	DebianServer
	DebianServer7
	DebianServer8
	DebianServer9
	DebianServer10
	DebianServer11
	DebianServer12
	DebianServer13
	DebianServer14
	DebianDesktop
	Docker
	DockerUbuntuServer
	DockerUbuntuServer18
	DockerUbuntuServer19
	DockerUbuntuServer20
	DockerUbuntuServer21
	DockerUbuntuServer22
	DockerCentos7
	DockerCentos8
	DockerCentos9
	DockerCentos10
	Android
	UbuntuDesktop
)

func (it Variant) ValueUInt16() uint16 {
	return uint16(it)
}

func (it Variant) DebianMap() map[Variant]bool {
	return isDebianMap
}

func (it Variant) DockerMap() map[Variant]bool {
	return isDockerMap
}

func (it Variant) UbuntuMap() map[Variant]bool {
	return isUbuntuMap
}

func (it Variant) UbuntuServerMap() map[Variant]bool {
	return isUbuntuServerMap
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

func (it Variant) IsAnyOf(checkingItems ...byte) bool {
	return BasicEnumImpl.IsAnyOf(it.Value(), checkingItems...)
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

func (it Variant) NameValue() string {
	return BasicEnumImpl.NameWithValue(it)
}

func (it Variant) Value() byte {
	return byte(it)
}

func (it Variant) Name() string {
	return BasicEnumImpl.ToEnumString(it.Value())
}

func (it Variant) IsUnknown() bool {
	return it == Invalid
}

func (it Variant) IsUbuntuServer() bool {
	return it == UbuntuServer
}

func (it Variant) IsUbuntuServer18() bool {
	return it == UbuntuServer18
}

func (it Variant) IsUbuntuServer19() bool {
	return it == UbuntuServer19
}

func (it Variant) IsUbuntuServer20() bool {
	return it == UbuntuServer20
}

func (it Variant) IsUbuntuServer21() bool {
	return it == UbuntuServer21
}

func (it Variant) IsUbuntuDesktop() bool {
	return it == UbuntuDesktop
}

func (it Variant) IsCentos() bool {
	return it == Centos
}

func (it Variant) IsCentos7() bool {
	return it == Centos7
}

func (it Variant) IsCentos8() bool {
	return it == Centos8
}

func (it Variant) IsCentos9() bool {
	return it == Centos9
}

func (it Variant) IsDebianServer() bool {
	return it == DebianServer
}

func (it Variant) IsDebianDesktop() bool {
	return it == DebianDesktop
}

func (it Variant) IsDocker() bool {
	return it == Docker
}

func (it Variant) IsDockerUbuntuServer() bool {
	return it == DockerUbuntuServer
}

func (it Variant) IsDockerUbuntuServer20() bool {
	return it == DockerUbuntuServer20
}

func (it Variant) IsDockerUbuntuServer21() bool {
	return it == DockerUbuntuServer20
}

func (it Variant) IsDockerCentos9() bool {
	return it == DockerCentos9
}

func (it Variant) IsAndroid() bool {
	return it == Android
}

func (it Variant) String() string {
	return BasicEnumImpl.ToEnumString(it.Value())
}

func (it Variant) IsInvalid() bool {
	return it == Invalid
}

func (it Variant) IsValid() bool {
	return it != Invalid
}

func (it Variant) IsNameOf(anyNames ...string) bool {
	for _, name := range anyNames {
		if name == it.Name() {
			return true
		}
	}

	return false
}

func (it Variant) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.Value())
}

func (it Variant) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.Value())
}

func (it *Variant) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (byte, error) {
	return BasicEnumImpl.
		UnmarshallToValue(true, jsonUnmarshallingValue)
}

func (it Variant) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Variant) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it *Variant) UnmarshalJSON(data []byte) error {
	dataConv, err := it.UnmarshallEnumToValue(data)

	if err != nil {
		return err
	}

	*it = Variant(dataConv)

	return nil
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

func (it Variant) AsLinuxTyper() enuminf.LinuxTyper {
	return &it
}

func (it Variant) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return &it
}

func (it Variant) AsBasicByteEnumContractsBinder() enuminf.BasicByteEnumContractsBinder {
	return &it
}

func (it Variant) ToPtr() *Variant {
	return &it
}

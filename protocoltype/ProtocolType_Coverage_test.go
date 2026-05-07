package protocoltype

import (
	"encoding/json"
	"testing"
)

func Test_ProtocolType_Coverage(t *testing.T) {
	all := []Variant{Invalid, Tcp, Udp, Icmp, Grpc, Rpc, OAuth, Rest, Http, Https, HttpsV3, MSMQ, Ip, IpV6, Ftp, Smtp, Imap, Pop3, Ssh, Sftp, Telnet, Pam, Sso, Smb, P2p, Custom}
	predicates := []func(Variant) bool{
		Variant.IsTcp, Variant.IsTcpOrUdp, Variant.IsUdp, Variant.IsIcmp, Variant.IsGrpc, Variant.IsRpc,
		Variant.IsOAuth, Variant.IsRest, Variant.IsHttp, Variant.IsHttps, Variant.IsHttpsV3, Variant.IsMSMQ,
		Variant.IsIp, Variant.IsIpV6, Variant.IsFtp, Variant.IsSmtp, Variant.IsImap, Variant.IsPop3, Variant.IsSsh,
		Variant.IsSftp, Variant.IsTelnet, Variant.IsPam, Variant.IsSso, Variant.IsCustom, Variant.IsDefined,
		Variant.IsTransactionProtocol, Variant.IsMailingProtocol, Variant.IsInternetProtocol, Variant.IsFirewallIpTablesProtocol,
		Variant.IsInvalid, Variant.IsValid,
	}
	for _, v := range all {
		_ = v.ValueByte()
		_ = v.ValueInt()
		_ = v.ValueUInt16()
		_ = v.Value()
		_ = v.Name()
		_ = v.CapitalName()
		_ = v.LowerName()
		_ = v.NameValue()
		_ = v.String()
		_ = v.ToNumberString()
		_ = v.RangeNamesCsv()
		_ = v.TypeName()
		_ = v.AllNameValues()
		_ = v.RangesByte()
		_ = v.RangesDynamicMap()
		_ = v.IntegerEnumRanges()
		_ = v.MinByte()
		_ = v.MaxByte()
		_ = v.MinInt()
		_ = v.MaxInt()
		_ = v.MinValueString()
		_ = v.MaxValueString()
		_, _ = v.MinMaxAny()
		_ = v.Format("name")
		_ = v.EnumType()
		_ = v.IsByteValueEqual(v.ValueByte())
		_ = v.IsValueEqual(v.ValueByte())
		_ = v.IsNameEqual(v.Name())
		_ = v.IsAnyValuesEqual(v.ValueByte())
		_ = v.IsAnyNamesOf(v.Name())
		_ = v.OnlySupportedErr("Tcp")
		_ = v.OnlySupportedMsgErr("m", "Tcp")
		_ = v.ToPtr()
		_ = v.Json()
		_ = v.JsonPtr()
		_ = v.AsBasicEnumContractsBinder()

		for _, p := range predicates {
			_ = p(v)
		}

		raw, err := json.Marshal(v)
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		var rt Variant
		_ = json.Unmarshal(raw, &rt)
	}
}

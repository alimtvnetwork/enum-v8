package protocoltype

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid: "Invalid",
		Tcp:     "Tcp",
		Udp:     "Udp",
		Icmp:    "Icmp",
		Grpc:    "Grpc",
		Rpc:     "Rpc",
		OAuth:   "OAuth",
		Rest:    "Rest",
		Http:    "Http",
		Https:   "Https",
		HttpsV3: "HttpsV3",
		MSMQ:    "MSMQ",
		Ip:      "Ip",
		IpV6:    "IpV6",
		Ftp:     "Ftp",
		Smtp:    "Smtp",
		Imap:    "Imap",
		Pop3:    "Pop3",
		Sftp:    "Sftp",
		Ssh:     "Ssh",
		Telnet:  "Telnet",
		Pam:     "Pam",
		Sso:     "Sso",
		Smb:     "Smb",
		P2p:     "P2p",
		Custom:  "Custom",
	}

	iptablesProtocols = [...]bool{
		Tcp:  true,
		Udp:  true,
		Icmp: true,
	}

	transactionProtocols = [...]bool{
		Tcp:   true,
		Udp:   true,
		Grpc:  true,
		Rpc:   true,
		OAuth: true,
		Rest:  true,
	}

	mailProtocols = [...]bool{
		Smtp: true,
		Imap: true,
		Pop3: true,
	}

	signInProtocols = [...]bool{
		OAuth: true,
		Ssh:   true,
		Pam:   true,
		Sso:   true,
	}

	ipProtocols = [...]bool{
		Ip:   true,
		IpV6: true,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)

package ethernetcmdnames

import (
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:                 "Invalid",
		Help:                    "Help",
		EnableDhcp:              "EnableDhcp",
		Ip:                      "Ip",
		IpSpecific:              "IpSpecific",
		IpConfig:                "IpConfig",
		IpConfigAdapter:         "IpConfigAdapter",
		Revert:                  "Revert",
		SetIpConfig:             "SetIpConfig",
		SetIpConfigRevertOnFail: "SetIpConfigRevertOnFail",
		SetIp:                   "SetIp",
		SetIpAdapter:            "SetIpAdapter",
		SetSubnet:               "SetSubnet",
		SetSubnetAdapter:        "SetSubnetAdapter",
		SetDns:                  "SetDns",
		SetDnsAdapter:           "SetDnsAdapter",
		SetGateway:              "SetGateway",
		SetGatewayAdapter:       "SetGatewayAdapter",
		Bridge:                  "Bridge",
		List:                    "List",
		ListIps:                 "ListIps",
		ListAdapters:            "ListAdapters",
		ListSpecificAdapter:     "ListSpecificAdapter",
		ListNetworks:            "ListNetworks",
		SetHostname:             "SetHostname",
		Hostname:                "Hostname",
		ListHistories:           "ListHistories",
		ListDetailedHistories:   "ListDetailedHistories",
		StateChange:             "StateChange",
		Backup:                  "Backup",
		Import:                  "Import",
	}

	aliasMap = map[string]byte{
		"ipconfig":  IpConfig.Value(),
		"ip-config": IpConfig.Value(),

		"set-ip":                SetIp.Value(),
		"set-ipconfig":          SetIpConfig.Value(),
		"set-ip-config":         SetIpConfig.Value(),
		"set-ip-revert-on-fail": SetIpConfigRevertOnFail.Value(),
		"set-ip-revert":         SetIpConfigRevertOnFail.Value(),
		"set-ip-r":              SetIpConfigRevertOnFail.Value(),
		"set-ip-adap":           SetIpAdapter.Value(),
		"set-ip-adapter":        SetIpAdapter.Value(),
		"set-adapter-ip":        SetIpAdapter.Value(),
		"set-subnet":            SetSubnet.Value(),
		"set-subnet-adapter":    SetSubnetAdapter.Value(),
		"set-adapter-subnet":    SetSubnetAdapter.Value(),
		"set-dns":               SetDns.Value(),
		"set-dns-adapter":       SetDnsAdapter.Value(),
		"set-adapter-dns":       SetDnsAdapter.Value(),
		"set-gateway":           SetGateway.Value(),
		"set-adapter-gateway":   SetGatewayAdapter.Value(),
		"set-gateway-adapter":   SetGatewayAdapter.Value(),
		"enable-dhcp":           EnableDhcp.Value(),
		"list-ips":              ListIps.Value(),
		"ips-list":              ListIps.Value(),
		"ip-a":                  List.Value(),
		"list-all":              List.Value(),
		"set-hostname":          SetHostname.Value(),
	}

	BasicEnumImpl = enumimpl.New.BasicByte.DefaultWithAliasMapAllCases(
		Invalid,
		Ranges[:],
		aliasMap)
)

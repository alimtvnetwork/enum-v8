package iptype

import (
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid: "Invalid",
		V4:      "IpV4",
		V6:      "IpV6",
	}

	aliasMap = map[string]byte{
		"Ipv4":     V4.ValueByte(),
		"Ipv6":     V6.ValueByte(),
		"v4":       V4.ValueByte(),
		"v6":       V6.ValueByte(),
		"ver4":     V4.ValueByte(),
		"ver6":     V6.ValueByte(),
		"version4": V4.ValueByte(),
		"version6": V6.ValueByte(),
		"ipv4":     V4.ValueByte(),
		"ipv6":     V6.ValueByte(),
		"ipV4":     V4.ValueByte(),
		"ipV6":     V6.ValueByte(),
		"IpV4":     V4.ValueByte(),
		"IpV6":     V6.ValueByte(),
	}

	BasicEnumImpl = enumimpl.New.BasicByte.DefaultWithAliasMap(
		Invalid,
		Ranges[:],
		aliasMap)
)

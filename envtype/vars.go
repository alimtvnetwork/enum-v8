package envtype

import (
	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Uninitialized: "Uninitialized",
		Development:   "Development",
		Development1:  "Development1",
		Development2:  "Development2",
		Test:          "Test",
		Test1:         "Test1",
		Test2:         "Test2",
		Production:    "Production",
		Production1:   "Production1",
		Production2:   "Production2",
	}

	keyNameMap = map[Variant]string{
		Uninitialized: "",
		Development:   "dev",
		Development1:  "dev-1",
		Development2:  "dev-2",
		Test:          "test",
		Test1:         "test-2",
		Test2:         "test-2",
		Production:    "prod",
		Production1:   "prod-1",
		Production2:   "prod-2",
	}

	curlyKeyNameMap = map[Variant]string{
		Uninitialized: "",
		Development:   "{dev}",
		Development1:  "{dev-1}",
		Development2:  "{dev-2}",
		Test:          "{test}",
		Test1:         "{test-2}",
		Test2:         "{test-2}",
		Production:    "{prod}",
		Production1:   "{prod-1}",
		Production2:   "{prod-2}",
	}

	rootMapping = map[Variant]Variant{
		Uninitialized: Uninitialized,
		Development:   Development,
		Development1:  Development,
		Development2:  Development,
		Test:          Test,
		Test1:         Test,
		Test2:         Test,
		Production:    Production,
		Production1:   Production,
		Production2:   Production,
	}

	envVersionNumber = map[Variant]int{
		Uninitialized: constants.InvalidValue,
		Development:   constants.Zero,
		Development1:  constants.One,
		Development2:  constants.Two,
		Test:          constants.Zero,
		Test1:         constants.One,
		Test2:         constants.Two,
		Production:    constants.Zero,
		Production1:   constants.One,
		Production2:   constants.Two,
	}

	devEnvMap = map[Variant]bool{
		Development:  true,
		Development1: true,
		Development2: true,
	}

	testEnvMap = map[Variant]bool{
		Test:  true,
		Test1: true,
		Test2: true,
	}

	productionEnvMap = map[Variant]bool{
		Production:  true,
		Production1: true,
		Production2: true,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Uninitialized),
		Ranges[:])
)

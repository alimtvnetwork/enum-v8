package creationtests

import (
	"github.com/alimtvnetwork/core-v9/coreinterface/enuminf"
	"github.com/alimtvnetwork/core-v9/enums/stringcompareas"
	"github.com/alimtvnetwork/core-v9/issetter"
	"github.com/alimtvnetwork/core-v9/reqtype"
	"github.com/alimtvnetwork/enum-v3/accesstype"
	"github.com/alimtvnetwork/enum-v3/brackets"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/compresscmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/configcmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/crontabscmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/decompresscmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/dnscmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/dockercmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/downloadcmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/envpathcmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/envvarscmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/ethernetcmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/fail2bancmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/firewallcmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/ftpcmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/hostingplancmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/macrocmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/operatingsystemcmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/packagecmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/rootcmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/servicescmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/snapshotcmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/sshcmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/sslcmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/toolingcmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/usercmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/userrolecmdnames"
	"github.com/alimtvnetwork/enum-v3/cmdenumtypes/webservercmdnames"
	"github.com/alimtvnetwork/enum-v3/completionstate"
	"github.com/alimtvnetwork/enum-v3/configfilestate"
	"github.com/alimtvnetwork/enum-v3/conntrackstate"
	"github.com/alimtvnetwork/enum-v3/dbaction"
	"github.com/alimtvnetwork/enum-v3/dbdrivertype"
	"github.com/alimtvnetwork/enum-v3/dbexposetype"
	"github.com/alimtvnetwork/enum-v3/dbuserprivillegetype"
	"github.com/alimtvnetwork/enum-v3/eventtype"
	"github.com/alimtvnetwork/enum-v3/instructiontype"
	"github.com/alimtvnetwork/enum-v3/inttype"
	"github.com/alimtvnetwork/enum-v3/iptype"
	"github.com/alimtvnetwork/enum-v3/leveltype"
	"github.com/alimtvnetwork/enum-v3/licensetype"
	"github.com/alimtvnetwork/enum-v3/linescomparetype"
	"github.com/alimtvnetwork/enum-v3/linuxservicestate"
	"github.com/alimtvnetwork/enum-v3/linuxtype"
	"github.com/alimtvnetwork/enum-v3/linuxvendortype"
	"github.com/alimtvnetwork/enum-v3/logtype"
	"github.com/alimtvnetwork/enum-v3/nginxlogtype"
	"github.com/alimtvnetwork/enum-v3/onofftype"
	"github.com/alimtvnetwork/enum-v3/osarchs"
	"github.com/alimtvnetwork/enum-v3/osdetect"
	"github.com/alimtvnetwork/enum-v3/osgroupexecution"
	"github.com/alimtvnetwork/enum-v3/overwritetype"
	"github.com/alimtvnetwork/enum-v3/packageinstallmethod"
	"github.com/alimtvnetwork/enum-v3/pathpatterntype"
	"github.com/alimtvnetwork/enum-v3/protocoltype"
	"github.com/alimtvnetwork/enum-v3/querymethodtype"
	"github.com/alimtvnetwork/enum-v3/quotes"
	"github.com/alimtvnetwork/enum-v3/resauthtype"
	"github.com/alimtvnetwork/enum-v3/revokereason"
	"github.com/alimtvnetwork/enum-v3/runtype"
	"github.com/alimtvnetwork/enum-v3/scripttype"
	"github.com/alimtvnetwork/enum-v3/servicestate"
	"github.com/alimtvnetwork/enum-v3/sqliteconnpathtype"
	"github.com/alimtvnetwork/enum-v3/sqljointype"
	"github.com/alimtvnetwork/enum-v3/strtype"
	"github.com/alimtvnetwork/enum-v3/taskcategory"
	"github.com/alimtvnetwork/enum-v3/taskpriority"
	"github.com/alimtvnetwork/enum-v3/timeunit"
	"github.com/alimtvnetwork/enum-v3/verifiertriggertype"
)

var simpleEnumCollectionTestCases = []enuminf.SimpleEnumer{
	issetter.Uninitialized,
	reqtype.Invalid,
	stringcompareas.Invalid.AsBasicEnumContractsBinder(),
	accesstype.Invalid,
	brackets.Invalid,
	
	completionstate.Invalid,
	
	configfilestate.Invalid,
	conntrackstate.Invalid,
	
	dbaction.Invalid,
	dbexposetype.Invalid,
	dbdrivertype.Invalid,
	
	dbuserprivilegetype.Invalid,
	eventtype.Invalid,
	
	instructiontype.Invalid,
	inttype.Invalid,
	iptype.Invalid,
	
	iptype.Invalid.AsBasicEnumContractsBinder(),
	
	leveltype.Invalid,
	licensetype.Invalid,
	linescomparetype.Invalid,
	linuxservicestate.Invalid,
	linuxtype.Invalid,
	linuxvendortype.Invalid,
	logtype.Invalid,
	
	nginxlogtype.Invalid.AsBasicEnumContractsBinder(),
	
	onofftype.Invalid,
	
	osarchs.Invalid,
	osgroupexecution.Invalid,
	osdetect.Invalid,
	overwritetype.Invalid,
	
	packageinstallmethod.Invalid,
	pathpatterntype.Invalid,
	protocoltype.Invalid,
	
	querymethodtype.Invalid,
	quotes.Invalid,
	
	resauthtype.Invalid,
	revokereason.Unspecified,
	runtype.Invalid,
	
	scripttype.Invalid,
	servicestate.Invalid,
	
	sqliteconnpathtype.Invalid,
	
	sqljointype.Invalid,
	strtype.Variant("Invalid"),
	taskcategory.Invalid,
	taskpriority.Invalid,
	
	timeunit.Invalid,
	verifiertriggertype.Invalid,
	
	compresscmdnames.Invalid.AsBasicEnumContractsBinder(),
	configcmdnames.Invalid.AsBasicEnumContractsBinder(),
	crontabscmdnames.Invalid.AsBasicEnumContractsBinder(),
	decompresscmdnames.Invalid.AsBasicEnumContractsBinder(),
	dnscmdnames.Invalid.AsBasicEnumContractsBinder(),
	dockercmdnames.Invalid.AsBasicEnumContractsBinder(),
	envpathcmdnames.Invalid.AsBasicEnumContractsBinder(),
	envvarscmdnames.Invalid.AsBasicEnumContractsBinder(),
	ethernetcmdnames.Invalid.AsBasicEnumContractsBinder(),
	downloadcmdnames.Invalid.AsBasicEnumContractsBinder(),
	ethernetcmdnames.Invalid.AsBasicEnumContractsBinder(),
	fail2bancmdnames.Invalid.AsBasicEnumContractsBinder(),
	firewallcmdnames.Invalid.AsBasicEnumContractsBinder(),
	ftpcmdnames.Invalid.AsBasicEnumContractsBinder(),
	hostingplancmdnames.Invalid.AsBasicEnumContractsBinder(),
	macrocmdnames.Invalid.AsBasicEnumContractsBinder(),
	operatingsystemcmdnames.Invalid.AsBasicEnumContractsBinder(),
	packagecmdnames.Invalid.AsBasicEnumContractsBinder(),
	rootcmdnames.Invalid.AsBasicEnumContractsBinder(),
	servicescmdnames.Invalid.AsBasicEnumContractsBinder(),
	snapshotcmdnames.Invalid.AsBasicEnumContractsBinder(),
	sshcmdnames.Invalid.AsBasicEnumContractsBinder(),
	sslcmdnames.Invalid.AsBasicEnumContractsBinder(),
	toolingcmdnames.Invalid.AsBasicEnumContractsBinder(),
	usercmdnames.Invalid.AsBasicEnumContractsBinder(),
	userrolecmdnames.Invalid.AsBasicEnumContractsBinder(),
	webservercmdnames.Invalid.AsBasicEnumContractsBinder(),
}

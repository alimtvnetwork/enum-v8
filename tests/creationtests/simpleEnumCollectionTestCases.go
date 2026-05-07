package creationtests

import (
	"github.com/alimtvnetwork/core-v9/coreinterface/enuminf"
	"github.com/alimtvnetwork/core-v9/enums/stringcompareas"
	"github.com/alimtvnetwork/core-v9/issetter"
	"github.com/alimtvnetwork/core-v9/reqtype"
	"github.com/alimtvnetwork/enum-v8/accesstype"
	"github.com/alimtvnetwork/enum-v8/brackets"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/compresscmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/configcmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/crontabscmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/decompresscmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/dnscmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/dockercmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/downloadcmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/envpathcmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/envvarscmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/ethernetcmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/fail2bancmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/firewallcmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/ftpcmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/hostingplancmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/macrocmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/operatingsystemcmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/packagecmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/rootcmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/servicescmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/snapshotcmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/sshcmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/sslcmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/toolingcmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/usercmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/userrolecmdnames"
	"github.com/alimtvnetwork/enum-v8/cmdenumtypes/webservercmdnames"
	"github.com/alimtvnetwork/enum-v8/completionstate"
	"github.com/alimtvnetwork/enum-v8/configfilestate"
	"github.com/alimtvnetwork/enum-v8/conntrackstate"
	"github.com/alimtvnetwork/enum-v8/dbaction"
	"github.com/alimtvnetwork/enum-v8/dbdrivertype"
	"github.com/alimtvnetwork/enum-v8/dbexposetype"
	"github.com/alimtvnetwork/enum-v8/dbuserprivillegetype"
	"github.com/alimtvnetwork/enum-v8/eventtype"
	"github.com/alimtvnetwork/enum-v8/instructiontype"
	"github.com/alimtvnetwork/enum-v8/inttype"
	"github.com/alimtvnetwork/enum-v8/iptype"
	"github.com/alimtvnetwork/enum-v8/leveltype"
	"github.com/alimtvnetwork/enum-v8/licensetype"
	"github.com/alimtvnetwork/enum-v8/linescomparetype"
	"github.com/alimtvnetwork/enum-v8/linuxservicestate"
	"github.com/alimtvnetwork/enum-v8/linuxtype"
	"github.com/alimtvnetwork/enum-v8/linuxvendortype"
	"github.com/alimtvnetwork/enum-v8/logtype"
	"github.com/alimtvnetwork/enum-v8/nginxlogtype"
	"github.com/alimtvnetwork/enum-v8/onofftype"
	"github.com/alimtvnetwork/enum-v8/osarchs"
	"github.com/alimtvnetwork/enum-v8/osdetect"
	"github.com/alimtvnetwork/enum-v8/osgroupexecution"
	"github.com/alimtvnetwork/enum-v8/overwritetype"
	"github.com/alimtvnetwork/enum-v8/packageinstallmethod"
	"github.com/alimtvnetwork/enum-v8/pathpatterntype"
	"github.com/alimtvnetwork/enum-v8/protocoltype"
	"github.com/alimtvnetwork/enum-v8/querymethodtype"
	"github.com/alimtvnetwork/enum-v8/quotes"
	"github.com/alimtvnetwork/enum-v8/resauthtype"
	"github.com/alimtvnetwork/enum-v8/revokereason"
	"github.com/alimtvnetwork/enum-v8/runtype"
	"github.com/alimtvnetwork/enum-v8/scripttype"
	"github.com/alimtvnetwork/enum-v8/servicestate"
	"github.com/alimtvnetwork/enum-v8/sqliteconnpathtype"
	"github.com/alimtvnetwork/enum-v8/sqljointype"
	"github.com/alimtvnetwork/enum-v8/strtype"
	"github.com/alimtvnetwork/enum-v8/taskcategory"
	"github.com/alimtvnetwork/enum-v8/taskpriority"
	"github.com/alimtvnetwork/enum-v8/timeunit"
	"github.com/alimtvnetwork/enum-v8/verifiertriggertype"
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

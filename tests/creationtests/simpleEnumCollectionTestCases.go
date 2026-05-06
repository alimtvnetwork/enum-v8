package creationtests

import (
	"github.com/alimtvnetwork/core-v9/coreinterface/enuminf"
	"github.com/alimtvnetwork/core-v9/enums/stringcompareas"
	"github.com/alimtvnetwork/core-v9/issetter"
	"github.com/alimtvnetwork/core-v9/reqtype"
	"github.com/alimtvnetwork/enum-v4/accesstype"
	"github.com/alimtvnetwork/enum-v4/brackets"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/compresscmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/configcmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/crontabscmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/decompresscmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/dnscmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/dockercmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/downloadcmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/envpathcmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/envvarscmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/ethernetcmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/fail2bancmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/firewallcmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/ftpcmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/hostingplancmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/macrocmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/operatingsystemcmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/packagecmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/rootcmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/servicescmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/snapshotcmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/sshcmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/sslcmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/toolingcmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/usercmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/userrolecmdnames"
	"github.com/alimtvnetwork/enum-v4/cmdenumtypes/webservercmdnames"
	"github.com/alimtvnetwork/enum-v4/completionstate"
	"github.com/alimtvnetwork/enum-v4/configfilestate"
	"github.com/alimtvnetwork/enum-v4/conntrackstate"
	"github.com/alimtvnetwork/enum-v4/dbaction"
	"github.com/alimtvnetwork/enum-v4/dbdrivertype"
	"github.com/alimtvnetwork/enum-v4/dbexposetype"
	"github.com/alimtvnetwork/enum-v4/dbuserprivillegetype"
	"github.com/alimtvnetwork/enum-v4/eventtype"
	"github.com/alimtvnetwork/enum-v4/instructiontype"
	"github.com/alimtvnetwork/enum-v4/inttype"
	"github.com/alimtvnetwork/enum-v4/iptype"
	"github.com/alimtvnetwork/enum-v4/leveltype"
	"github.com/alimtvnetwork/enum-v4/licensetype"
	"github.com/alimtvnetwork/enum-v4/linescomparetype"
	"github.com/alimtvnetwork/enum-v4/linuxservicestate"
	"github.com/alimtvnetwork/enum-v4/linuxtype"
	"github.com/alimtvnetwork/enum-v4/linuxvendortype"
	"github.com/alimtvnetwork/enum-v4/logtype"
	"github.com/alimtvnetwork/enum-v4/nginxlogtype"
	"github.com/alimtvnetwork/enum-v4/onofftype"
	"github.com/alimtvnetwork/enum-v4/osarchs"
	"github.com/alimtvnetwork/enum-v4/osdetect"
	"github.com/alimtvnetwork/enum-v4/osgroupexecution"
	"github.com/alimtvnetwork/enum-v4/overwritetype"
	"github.com/alimtvnetwork/enum-v4/packageinstallmethod"
	"github.com/alimtvnetwork/enum-v4/pathpatterntype"
	"github.com/alimtvnetwork/enum-v4/protocoltype"
	"github.com/alimtvnetwork/enum-v4/querymethodtype"
	"github.com/alimtvnetwork/enum-v4/quotes"
	"github.com/alimtvnetwork/enum-v4/resauthtype"
	"github.com/alimtvnetwork/enum-v4/revokereason"
	"github.com/alimtvnetwork/enum-v4/runtype"
	"github.com/alimtvnetwork/enum-v4/scripttype"
	"github.com/alimtvnetwork/enum-v4/servicestate"
	"github.com/alimtvnetwork/enum-v4/sqliteconnpathtype"
	"github.com/alimtvnetwork/enum-v4/sqljointype"
	"github.com/alimtvnetwork/enum-v4/strtype"
	"github.com/alimtvnetwork/enum-v4/taskcategory"
	"github.com/alimtvnetwork/enum-v4/taskpriority"
	"github.com/alimtvnetwork/enum-v4/timeunit"
	"github.com/alimtvnetwork/enum-v4/verifiertriggertype"
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

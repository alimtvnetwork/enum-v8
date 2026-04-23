package creationtests

import (
	"github.com/alimtvnetwork/core-v8/coreinterface/enuminf"
	"github.com/alimtvnetwork/core-v8/enums/stringcompareas"
	"github.com/alimtvnetwork/core-v8/reqtype"
	"github.com/alimtvnetwork/enum-v1/accesstype"
	"github.com/alimtvnetwork/enum-v1/brackets"
	"github.com/alimtvnetwork/enum-v1/certaction"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/compresscmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/configcmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/crontabscmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/decompresscmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/dnscmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/dockercmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/downloadcmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/envpathcmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/envvarscmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/ethernetcmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/fail2bancmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/firewallcmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/ftpcmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/hostingplancmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/macrocmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/operatingsystemcmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/packagecmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/rootcmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/servicescmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/snapshotcmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/sshcmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/sslcmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/sysgroupcmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/toolingcmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/usercmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/userrolecmdnames"
	"github.com/alimtvnetwork/enum-v1/cmdenumtypes/webservercmdnames"
	"github.com/alimtvnetwork/enum-v1/completionstate"
	"github.com/alimtvnetwork/enum-v1/compressformats"
	"github.com/alimtvnetwork/enum-v1/compresslevels"
	"github.com/alimtvnetwork/enum-v1/configfilestate"
	"github.com/alimtvnetwork/enum-v1/conntrackstate"
	"github.com/alimtvnetwork/enum-v1/dbaction"
	"github.com/alimtvnetwork/enum-v1/dbdrivertype"
	"github.com/alimtvnetwork/enum-v1/dbexposetype"
	dbuserprivilegetype "github.com/alimtvnetwork/enum-v1/dbuserprivillegetype"
	"github.com/alimtvnetwork/enum-v1/eventtype"
	"github.com/alimtvnetwork/enum-v1/instructiontype"
	"github.com/alimtvnetwork/enum-v1/inttype"
	"github.com/alimtvnetwork/enum-v1/iptype"
	"github.com/alimtvnetwork/enum-v1/leveltype"
	"github.com/alimtvnetwork/enum-v1/licensetype"
	"github.com/alimtvnetwork/enum-v1/linescomparetype"
	"github.com/alimtvnetwork/enum-v1/linuxservicestate"
	"github.com/alimtvnetwork/enum-v1/linuxtype"
	"github.com/alimtvnetwork/enum-v1/linuxvendortype"
	"github.com/alimtvnetwork/enum-v1/logtype"
	"github.com/alimtvnetwork/enum-v1/nginxlogtype"
	"github.com/alimtvnetwork/enum-v1/onofftype"
	"github.com/alimtvnetwork/enum-v1/osarchs"
	"github.com/alimtvnetwork/enum-v1/osdetect"
	"github.com/alimtvnetwork/enum-v1/osgroupexecution"
	"github.com/alimtvnetwork/enum-v1/overwritetype"
	"github.com/alimtvnetwork/enum-v1/packageinstallmethod"
	"github.com/alimtvnetwork/enum-v1/pathpatterntype"
	"github.com/alimtvnetwork/enum-v1/promptclitype"
	"github.com/alimtvnetwork/enum-v1/protocoltype"
	"github.com/alimtvnetwork/enum-v1/querymethodtype"
	"github.com/alimtvnetwork/enum-v1/quotes"
	"github.com/alimtvnetwork/enum-v1/resauthtype"
	"github.com/alimtvnetwork/enum-v1/revokereason"
	"github.com/alimtvnetwork/enum-v1/runtype"
	"github.com/alimtvnetwork/enum-v1/scripttype"
	"github.com/alimtvnetwork/enum-v1/servicestate"
	"github.com/alimtvnetwork/enum-v1/sitestatetype"
	"github.com/alimtvnetwork/enum-v1/sqliteconnpathtype"
	"github.com/alimtvnetwork/enum-v1/sqljointype"
	"github.com/alimtvnetwork/enum-v1/strtype"
	"github.com/alimtvnetwork/enum-v1/taskcategory"
	"github.com/alimtvnetwork/enum-v1/taskpriority"
	"github.com/alimtvnetwork/enum-v1/timeunit"
	"github.com/alimtvnetwork/enum-v1/verifiertriggertype"
)

var allBasicEnumsCollection = [...]enuminf.BasicEnumer{
	&setterInvalid,
	reqtype.Invalid.AsBasicByteEnumContractsBinder(),
	stringcompareas.Invalid.AsBasicByteEnumContractsBinder(),
	accesstype.Invalid.AsBasicByteEnumContractsBinder(),
	brackets.Invalid.AsBasicByteEnumContractsBinder(),
	
	certaction.Invalid.AsBasicEnumContractsBinder(),
	
	completionstate.Invalid.AsBasicEnumContractsBinder(),
	
	compressformats.Invalid.AsBasicEnumContractsBinder(),
	compresslevels.Invalid.AsBasicEnumContractsBinder(),
	
	configfilestate.Invalid.AsBasicByteEnumContractsBinder(),
	conntrackstate.Invalid.AsBasicByteEnumContractsBinder(),
	
	dbaction.Invalid.AsBasicEnumContractsBinder(),
	dbexposetype.Invalid.AsBasicEnumContractsBinder(),
	dbdrivertype.Invalid.AsBasicByteEnumContractsBinder(),
	dbuserprivilegetype.Invalid.AsBasicEnumContractsBinder(),
	eventtype.Invalid.AsBasicByteEnumContractsBinder(),
	
	instructiontype.Invalid.AsBasicByteEnumContractsBinder(),
	inttype.Variant(0).AsBasicEnumer(),
	
	iptype.Invalid.AsBasicEnumContractsBinder(),
	
	leveltype.Invalid.AsBasicEnumContractsBinder(),
	licensetype.Invalid.AsBasicEnumContractsBinder(),
	linescomparetype.Invalid.AsBasicByteEnumContractsBinder(),
	linuxservicestate.Invalid.AsBasicByteEnumContractsBinder(),
	linuxtype.Invalid.AsBasicByteEnumContractsBinder(),
	linuxvendortype.Invalid.AsBasicByteEnumContractsBinder(),
	logtype.Invalid.AsBasicByteEnumContractsBinder(),
	
	nginxlogtype.Invalid.AsBasicEnumContractsBinder(),
	
	onofftype.Invalid.AsBasicEnumContractsBinder(),
	osarchs.Invalid.AsBasicByteEnumContractsBinder(),
	osgroupexecution.Invalid.AsBasicByteEnumContractsBinder(),
	osdetect.Invalid.AsBasicByteEnumContractsBinder(),
	overwritetype.Invalid.AsBasicByteEnumContractsBinder(),
	packageinstallmethod.Invalid.AsBasicEnumContractsBinder(),
	pathpatterntype.Invalid.AsBasicByteEnumContractsBinder(),
	promptclitype.Invalid.AsBasicByteEnumContractsBinder(),
	protocoltype.Invalid.AsBasicEnumContractsBinder(),
	querymethodtype.Invalid.AsBasicByteEnumContractsBinder(),
	quotes.Invalid.AsBasicByteEnumContractsBinder(),
	resauthtype.Invalid.AsBasicEnumContractsBinder(),
	revokereason.Unspecified.AsBasicEnumContractsBinder(),
	runtype.Invalid.AsBasicEnumContractsBinder(),
	scripttype.Invalid.AsBasicByteEnumContractsBinder(),
	
	servicestate.Invalid.AsBasicByteEnumContractsBinder(),
	
	sitestatetype.Invalid.AsBasicByteEnumContractsBinder(),
	
	sqliteconnpathtype.Invalid.AsBasicEnumContractsBinder(),
	sqljointype.Invalid.AsBasicEnumContractsBinder(),
	strtype.Variant("Invalid").AsBasicEnumer(),
	taskcategory.Invalid.AsBasicEnumContractsBinder(),
	taskpriority.Invalid.AsBasicEnumContractsBinder(),
	timeunit.Invalid.AsBasicEnumContractsBinder(),
	verifiertriggertype.Invalid.AsBasicEnumContractsBinder(),
	
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
	sysgroupcmdnames.Invalid.AsBasicEnumContractsBinder(),
	toolingcmdnames.Invalid.AsBasicEnumContractsBinder(),
	usercmdnames.Invalid.AsBasicEnumContractsBinder(),
	userrolecmdnames.Invalid.AsBasicEnumContractsBinder(),
	webservercmdnames.Invalid.AsBasicEnumContractsBinder(),
}

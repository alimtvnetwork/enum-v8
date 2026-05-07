package creationtests

import (
	"github.com/alimtvnetwork/core-v9/coreinterface/enuminf"
	"github.com/alimtvnetwork/core-v9/enums/stringcompareas"
	"github.com/alimtvnetwork/core-v9/reqtype"
	"github.com/alimtvnetwork/enum-v7/accesstype"
	"github.com/alimtvnetwork/enum-v7/brackets"
	"github.com/alimtvnetwork/enum-v7/certaction"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/compresscmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/configcmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/crontabscmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/decompresscmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/dnscmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/dockercmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/downloadcmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/envpathcmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/envvarscmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/ethernetcmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/fail2bancmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/firewallcmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/ftpcmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/hostingplancmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/macrocmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/operatingsystemcmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/packagecmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/rootcmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/servicescmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/snapshotcmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/sshcmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/sslcmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/sysgroupcmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/toolingcmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/usercmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/userrolecmdnames"
	"github.com/alimtvnetwork/enum-v7/cmdenumtypes/webservercmdnames"
	"github.com/alimtvnetwork/enum-v7/completionstate"
	"github.com/alimtvnetwork/enum-v7/compressformats"
	"github.com/alimtvnetwork/enum-v7/compresslevels"
	"github.com/alimtvnetwork/enum-v7/configfilestate"
	"github.com/alimtvnetwork/enum-v7/conntrackstate"
	"github.com/alimtvnetwork/enum-v7/dbaction"
	"github.com/alimtvnetwork/enum-v7/dbdrivertype"
	"github.com/alimtvnetwork/enum-v7/dbexposetype"
	dbuserprivilegetype "github.com/alimtvnetwork/enum-v7/dbuserprivillegetype"
	"github.com/alimtvnetwork/enum-v7/eventtype"
	"github.com/alimtvnetwork/enum-v7/instructiontype"
	"github.com/alimtvnetwork/enum-v7/inttype"
	"github.com/alimtvnetwork/enum-v7/iptype"
	"github.com/alimtvnetwork/enum-v7/leveltype"
	"github.com/alimtvnetwork/enum-v7/licensetype"
	"github.com/alimtvnetwork/enum-v7/linescomparetype"
	"github.com/alimtvnetwork/enum-v7/linuxservicestate"
	"github.com/alimtvnetwork/enum-v7/linuxtype"
	"github.com/alimtvnetwork/enum-v7/linuxvendortype"
	"github.com/alimtvnetwork/enum-v7/logtype"
	"github.com/alimtvnetwork/enum-v7/nginxlogtype"
	"github.com/alimtvnetwork/enum-v7/onofftype"
	"github.com/alimtvnetwork/enum-v7/osarchs"
	"github.com/alimtvnetwork/enum-v7/osdetect"
	"github.com/alimtvnetwork/enum-v7/osgroupexecution"
	"github.com/alimtvnetwork/enum-v7/overwritetype"
	"github.com/alimtvnetwork/enum-v7/packageinstallmethod"
	"github.com/alimtvnetwork/enum-v7/pathpatterntype"
	"github.com/alimtvnetwork/enum-v7/promptclitype"
	"github.com/alimtvnetwork/enum-v7/protocoltype"
	"github.com/alimtvnetwork/enum-v7/querymethodtype"
	"github.com/alimtvnetwork/enum-v7/quotes"
	"github.com/alimtvnetwork/enum-v7/resauthtype"
	"github.com/alimtvnetwork/enum-v7/revokereason"
	"github.com/alimtvnetwork/enum-v7/runtype"
	"github.com/alimtvnetwork/enum-v7/scripttype"
	"github.com/alimtvnetwork/enum-v7/servicestate"
	"github.com/alimtvnetwork/enum-v7/sitestatetype"
	"github.com/alimtvnetwork/enum-v7/sqliteconnpathtype"
	"github.com/alimtvnetwork/enum-v7/sqljointype"
	"github.com/alimtvnetwork/enum-v7/strtype"
	"github.com/alimtvnetwork/enum-v7/taskcategory"
	"github.com/alimtvnetwork/enum-v7/taskpriority"
	"github.com/alimtvnetwork/enum-v7/timeunit"
	"github.com/alimtvnetwork/enum-v7/verifiertriggertype"
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

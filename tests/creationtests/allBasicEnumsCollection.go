package creationtests

import (
	"github.com/alimtvnetwork/core-v9/coreinterface/enuminf"
	"github.com/alimtvnetwork/core-v9/enums/stringcompareas"
	"github.com/alimtvnetwork/core-v9/reqtype"
	"github.com/alimtvnetwork/enum-v5/accesstype"
	"github.com/alimtvnetwork/enum-v5/brackets"
	"github.com/alimtvnetwork/enum-v5/certaction"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/compresscmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/configcmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/crontabscmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/decompresscmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/dnscmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/dockercmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/downloadcmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/envpathcmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/envvarscmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/ethernetcmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/fail2bancmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/firewallcmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/ftpcmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/hostingplancmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/macrocmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/operatingsystemcmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/packagecmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/rootcmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/servicescmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/snapshotcmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/sshcmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/sslcmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/sysgroupcmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/toolingcmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/usercmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/userrolecmdnames"
	"github.com/alimtvnetwork/enum-v5/cmdenumtypes/webservercmdnames"
	"github.com/alimtvnetwork/enum-v5/completionstate"
	"github.com/alimtvnetwork/enum-v5/compressformats"
	"github.com/alimtvnetwork/enum-v5/compresslevels"
	"github.com/alimtvnetwork/enum-v5/configfilestate"
	"github.com/alimtvnetwork/enum-v5/conntrackstate"
	"github.com/alimtvnetwork/enum-v5/dbaction"
	"github.com/alimtvnetwork/enum-v5/dbdrivertype"
	"github.com/alimtvnetwork/enum-v5/dbexposetype"
	dbuserprivilegetype "github.com/alimtvnetwork/enum-v5/dbuserprivillegetype"
	"github.com/alimtvnetwork/enum-v5/eventtype"
	"github.com/alimtvnetwork/enum-v5/instructiontype"
	"github.com/alimtvnetwork/enum-v5/inttype"
	"github.com/alimtvnetwork/enum-v5/iptype"
	"github.com/alimtvnetwork/enum-v5/leveltype"
	"github.com/alimtvnetwork/enum-v5/licensetype"
	"github.com/alimtvnetwork/enum-v5/linescomparetype"
	"github.com/alimtvnetwork/enum-v5/linuxservicestate"
	"github.com/alimtvnetwork/enum-v5/linuxtype"
	"github.com/alimtvnetwork/enum-v5/linuxvendortype"
	"github.com/alimtvnetwork/enum-v5/logtype"
	"github.com/alimtvnetwork/enum-v5/nginxlogtype"
	"github.com/alimtvnetwork/enum-v5/onofftype"
	"github.com/alimtvnetwork/enum-v5/osarchs"
	"github.com/alimtvnetwork/enum-v5/osdetect"
	"github.com/alimtvnetwork/enum-v5/osgroupexecution"
	"github.com/alimtvnetwork/enum-v5/overwritetype"
	"github.com/alimtvnetwork/enum-v5/packageinstallmethod"
	"github.com/alimtvnetwork/enum-v5/pathpatterntype"
	"github.com/alimtvnetwork/enum-v5/promptclitype"
	"github.com/alimtvnetwork/enum-v5/protocoltype"
	"github.com/alimtvnetwork/enum-v5/querymethodtype"
	"github.com/alimtvnetwork/enum-v5/quotes"
	"github.com/alimtvnetwork/enum-v5/resauthtype"
	"github.com/alimtvnetwork/enum-v5/revokereason"
	"github.com/alimtvnetwork/enum-v5/runtype"
	"github.com/alimtvnetwork/enum-v5/scripttype"
	"github.com/alimtvnetwork/enum-v5/servicestate"
	"github.com/alimtvnetwork/enum-v5/sitestatetype"
	"github.com/alimtvnetwork/enum-v5/sqliteconnpathtype"
	"github.com/alimtvnetwork/enum-v5/sqljointype"
	"github.com/alimtvnetwork/enum-v5/strtype"
	"github.com/alimtvnetwork/enum-v5/taskcategory"
	"github.com/alimtvnetwork/enum-v5/taskpriority"
	"github.com/alimtvnetwork/enum-v5/timeunit"
	"github.com/alimtvnetwork/enum-v5/verifiertriggertype"
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

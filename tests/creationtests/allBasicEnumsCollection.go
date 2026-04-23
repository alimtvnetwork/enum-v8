package creationtests

import (
	"github.com/alimtvnetwork/core-v8/coreinterface/enuminf"
	"github.com/alimtvnetwork/core-v8/enums/stringcompareas"
	"github.com/alimtvnetwork/core-v8/reqtype"
	"https://github.com/alimtvnetwork/enum-v1/accesstype"
	"https://github.com/alimtvnetwork/enum-v1/brackets"
	"https://github.com/alimtvnetwork/enum-v1/certaction"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/compresscmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/configcmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/crontabscmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/decompresscmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/dnscmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/dockercmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/downloadcmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/envpathcmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/envvarscmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/ethernetcmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/fail2bancmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/firewallcmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/ftpcmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/hostingplancmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/macrocmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/operatingsystemcmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/packagecmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/rootcmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/servicescmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/snapshotcmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/sshcmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/sslcmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/sysgroupcmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/toolingcmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/usercmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/userrolecmdnames"
	"https://github.com/alimtvnetwork/enum-v1/cmdenumtypes/webservercmdnames"
	"https://github.com/alimtvnetwork/enum-v1/completionstate"
	"https://github.com/alimtvnetwork/enum-v1/compressformats"
	"https://github.com/alimtvnetwork/enum-v1/compresslevels"
	"https://github.com/alimtvnetwork/enum-v1/configfilestate"
	"https://github.com/alimtvnetwork/enum-v1/conntrackstate"
	"https://github.com/alimtvnetwork/enum-v1/dbaction"
	"https://github.com/alimtvnetwork/enum-v1/dbdrivertype"
	"https://github.com/alimtvnetwork/enum-v1/dbexposetype"
	dbuserprivilegetype "https://github.com/alimtvnetwork/enum-v1/dbuserprivillegetype"
	"https://github.com/alimtvnetwork/enum-v1/eventtype"
	"https://github.com/alimtvnetwork/enum-v1/instructiontype"
	"https://github.com/alimtvnetwork/enum-v1/inttype"
	"https://github.com/alimtvnetwork/enum-v1/iptype"
	"https://github.com/alimtvnetwork/enum-v1/leveltype"
	"https://github.com/alimtvnetwork/enum-v1/licensetype"
	"https://github.com/alimtvnetwork/enum-v1/linescomparetype"
	"https://github.com/alimtvnetwork/enum-v1/linuxservicestate"
	"https://github.com/alimtvnetwork/enum-v1/linuxtype"
	"https://github.com/alimtvnetwork/enum-v1/linuxvendortype"
	"https://github.com/alimtvnetwork/enum-v1/logtype"
	"https://github.com/alimtvnetwork/enum-v1/nginxlogtype"
	"https://github.com/alimtvnetwork/enum-v1/onofftype"
	"https://github.com/alimtvnetwork/enum-v1/osarchs"
	"https://github.com/alimtvnetwork/enum-v1/osdetect"
	"https://github.com/alimtvnetwork/enum-v1/osgroupexecution"
	"https://github.com/alimtvnetwork/enum-v1/overwritetype"
	"https://github.com/alimtvnetwork/enum-v1/packageinstallmethod"
	"https://github.com/alimtvnetwork/enum-v1/pathpatterntype"
	"https://github.com/alimtvnetwork/enum-v1/promptclitype"
	"https://github.com/alimtvnetwork/enum-v1/protocoltype"
	"https://github.com/alimtvnetwork/enum-v1/querymethodtype"
	"https://github.com/alimtvnetwork/enum-v1/quotes"
	"https://github.com/alimtvnetwork/enum-v1/resauthtype"
	"https://github.com/alimtvnetwork/enum-v1/revokereason"
	"https://github.com/alimtvnetwork/enum-v1/runtype"
	"https://github.com/alimtvnetwork/enum-v1/scripttype"
	"https://github.com/alimtvnetwork/enum-v1/servicestate"
	"https://github.com/alimtvnetwork/enum-v1/sitestatetype"
	"https://github.com/alimtvnetwork/enum-v1/sqliteconnpathtype"
	"https://github.com/alimtvnetwork/enum-v1/sqljointype"
	"https://github.com/alimtvnetwork/enum-v1/strtype"
	"https://github.com/alimtvnetwork/enum-v1/taskcategory"
	"https://github.com/alimtvnetwork/enum-v1/taskpriority"
	"https://github.com/alimtvnetwork/enum-v1/timeunit"
	"https://github.com/alimtvnetwork/enum-v1/verifiertriggertype"
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

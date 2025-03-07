package resources

type Privilege string

func (p Privilege) String() string {
	return string(p)
}

const (
	privilegeSelect                    Privilege = "SELECT"
	privilegeInsert                    Privilege = "INSERT"
	privilegeUpdate                    Privilege = "UPDATE"
	privilegeDelete                    Privilege = "DELETE"
	privilegeTruncate                  Privilege = "TRUNCATE"
	privilegeReferences                Privilege = "REFERENCES"
	privilegeRebuild                   Privilege = "REBUILD"
	privilegeCreateSchema              Privilege = "CREATE SCHEMA"
	privilegeImportedPrivileges        Privilege = "IMPORTED PRIVILEGES"
	privilegeModify                    Privilege = "MODIFY"
	privilegeOperate                   Privilege = "OPERATE"
	privilegeMonitor                   Privilege = "MONITOR"
	privilegeOwnership                 Privilege = "OWNERSHIP"
	privilegeRead                      Privilege = "READ"
	privilegeReferenceUsage            Privilege = "REFERENCE_USAGE"
	privilegeUsage                     Privilege = "USAGE"
	privilegeWrite                     Privilege = "WRITE"
	privilegeCreateTable               Privilege = "CREATE TABLE"
	privilegeCreateTag                 Privilege = "CREATE TAG"
	privilegeCreateView                Privilege = "CREATE VIEW"
	privilegeCreateFileFormat          Privilege = "CREATE FILE FORMAT"
	privilegeCreateStage               Privilege = "CREATE STAGE"
	privilegeCreatePipe                Privilege = "CREATE PIPE"
	privilegeCreateStream              Privilege = "CREATE STREAM"
	privilegeCreateTask                Privilege = "CREATE TASK"
	privilegeCreateSequence            Privilege = "CREATE SEQUENCE"
	privilegeCreateFunction            Privilege = "CREATE FUNCTION"
	privilegeCreateProcedure           Privilege = "CREATE PROCEDURE"
	privilegeCreateExternalTable       Privilege = "CREATE EXTERNAL TABLE"
	privilegeCreateMaterializedView    Privilege = "CREATE MATERIALIZED VIEW"
	privilegeCreateRowAccessPolicy     Privilege = "CREATE ROW ACCESS POLICY"
	privilegeCreateTemporaryTable      Privilege = "CREATE TEMPORARY TABLE"
	privilegeCreateMaskingPolicy       Privilege = "CREATE MASKING POLICY"
	privilegeCreateNetworkPolicy       Privilege = "CREATE NETWORK POLICY"
	privilegeCreateDataExchangeListing Privilege = "CREATE DATA EXCHANGE LISTING"
	privilegeCreateAccount             Privilege = "CREATE ACCOUNT"
	privilegeCreateShare               Privilege = "CREATE SHARE"
	privilegeImportShare               Privilege = "IMPORT SHARE"
	privilegeOverrideShareRestrictions Privilege = "OVERRIDE SHARE RESTRICTIONS"
	privilegeAddSearchOptimization     Privilege = "ADD SEARCH OPTIMIZATION"
	privilegeApplyMaskingPolicy        Privilege = "APPLY MASKING POLICY"
	privilegeApplyPasswordPolicy       Privilege = "APPLY PASSWORD POLICY"
	privilegeApplySessionPolicy        Privilege = "APPLY SESSION POLICY"
	privilegeApplyRowAccessPolicy      Privilege = "APPLY ROW ACCESS POLICY"
	privilegeApplyTag                  Privilege = "APPLY TAG"
	privilegeApply                     Privilege = "APPLY"
	privilegeAttachPolicy              Privilege = "ATTACH POLICY"
	privilegeCreateRole                Privilege = "CREATE ROLE"
	privilegeCreateUser                Privilege = "CREATE USER"
	privilegeCreateWarehouse           Privilege = "CREATE WAREHOUSE"
	privilegeCreateDatabase            Privilege = "CREATE DATABASE"
	privilegeCreateIntegration         Privilege = "CREATE INTEGRATION"
	privilegeManageGrants              Privilege = "MANAGE GRANTS"
	privilegeMonitorUsage              Privilege = "MONITOR USAGE"
	privilegeMonitorExecution          Privilege = "MONITOR EXECUTION"
	privilegeExecuteTask               Privilege = "EXECUTE TASK"
	privilegeExecuteManagedTask        Privilege = "EXECUTE MANAGED TASK"
	privilegeOrganizationSupportCases  Privilege = "MANAGE ORGANIZATION SUPPORT CASES"
	privilegeAccountSupportCases       Privilege = "MANAGE ACCOUNT SUPPORT CASES"
	privilegeUserSupportCases          Privilege = "MANAGE USER SUPPORT CASES"
)

type PrivilegeSet map[Privilege]struct{}

func NewPrivilegeSet(privileges ...Privilege) PrivilegeSet {
	ps := PrivilegeSet{}
	for _, priv := range privileges {
		ps[priv] = struct{}{}
	}
	return ps
}

func (ps PrivilegeSet) ToList() []string {
	privs := []string{}
	for p := range ps {
		privs = append(privs, string(p))
	}
	return privs
}

func (ps PrivilegeSet) addString(s string) {
	ps[Privilege(s)] = struct{}{}
}

func (ps PrivilegeSet) hasString(s string) bool {
	_, ok := ps[Privilege(s)]
	return ok
}

package role

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"postgresql.role.Role",
		reflect.TypeOf((*Role)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "assumeRole", GoGetter: "AssumeRole"},
			_jsii_.MemberProperty{JsiiProperty: "assumeRoleInput", GoGetter: "AssumeRoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "bypassRowLevelSecurity", GoGetter: "BypassRowLevelSecurity"},
			_jsii_.MemberProperty{JsiiProperty: "bypassRowLevelSecurityInput", GoGetter: "BypassRowLevelSecurityInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "connectionLimit", GoGetter: "ConnectionLimit"},
			_jsii_.MemberProperty{JsiiProperty: "connectionLimitInput", GoGetter: "ConnectionLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createDatabase", GoGetter: "CreateDatabase"},
			_jsii_.MemberProperty{JsiiProperty: "createDatabaseInput", GoGetter: "CreateDatabaseInput"},
			_jsii_.MemberProperty{JsiiProperty: "createRole", GoGetter: "CreateRole"},
			_jsii_.MemberProperty{JsiiProperty: "createRoleInput", GoGetter: "CreateRoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "encrypted", GoGetter: "Encrypted"},
			_jsii_.MemberProperty{JsiiProperty: "encryptedInput", GoGetter: "EncryptedInput"},
			_jsii_.MemberProperty{JsiiProperty: "encryptedPassword", GoGetter: "EncryptedPassword"},
			_jsii_.MemberProperty{JsiiProperty: "encryptedPasswordInput", GoGetter: "EncryptedPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "idleInTransactionSessionTimeout", GoGetter: "IdleInTransactionSessionTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "idleInTransactionSessionTimeoutInput", GoGetter: "IdleInTransactionSessionTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "inherit", GoGetter: "Inherit"},
			_jsii_.MemberProperty{JsiiProperty: "inheritInput", GoGetter: "InheritInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "login", GoGetter: "Login"},
			_jsii_.MemberProperty{JsiiProperty: "loginInput", GoGetter: "LoginInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "passwordInput", GoGetter: "PasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "replication", GoGetter: "Replication"},
			_jsii_.MemberProperty{JsiiProperty: "replicationInput", GoGetter: "ReplicationInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAssumeRole", GoMethod: "ResetAssumeRole"},
			_jsii_.MemberMethod{JsiiMethod: "resetBypassRowLevelSecurity", GoMethod: "ResetBypassRowLevelSecurity"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectionLimit", GoMethod: "ResetConnectionLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreateDatabase", GoMethod: "ResetCreateDatabase"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreateRole", GoMethod: "ResetCreateRole"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncrypted", GoMethod: "ResetEncrypted"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptedPassword", GoMethod: "ResetEncryptedPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdleInTransactionSessionTimeout", GoMethod: "ResetIdleInTransactionSessionTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetInherit", GoMethod: "ResetInherit"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogin", GoMethod: "ResetLogin"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPassword", GoMethod: "ResetPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplication", GoMethod: "ResetReplication"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoles", GoMethod: "ResetRoles"},
			_jsii_.MemberMethod{JsiiMethod: "resetSearchPath", GoMethod: "ResetSearchPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipDropRole", GoMethod: "ResetSkipDropRole"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipReassignOwned", GoMethod: "ResetSkipReassignOwned"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatementTimeout", GoMethod: "ResetStatementTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuperuser", GoMethod: "ResetSuperuser"},
			_jsii_.MemberMethod{JsiiMethod: "resetValidUntil", GoMethod: "ResetValidUntil"},
			_jsii_.MemberProperty{JsiiProperty: "roles", GoGetter: "Roles"},
			_jsii_.MemberProperty{JsiiProperty: "rolesInput", GoGetter: "RolesInput"},
			_jsii_.MemberProperty{JsiiProperty: "searchPath", GoGetter: "SearchPath"},
			_jsii_.MemberProperty{JsiiProperty: "searchPathInput", GoGetter: "SearchPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipDropRole", GoGetter: "SkipDropRole"},
			_jsii_.MemberProperty{JsiiProperty: "skipDropRoleInput", GoGetter: "SkipDropRoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipReassignOwned", GoGetter: "SkipReassignOwned"},
			_jsii_.MemberProperty{JsiiProperty: "skipReassignOwnedInput", GoGetter: "SkipReassignOwnedInput"},
			_jsii_.MemberProperty{JsiiProperty: "statementTimeout", GoGetter: "StatementTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "statementTimeoutInput", GoGetter: "StatementTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "superuser", GoGetter: "Superuser"},
			_jsii_.MemberProperty{JsiiProperty: "superuserInput", GoGetter: "SuperuserInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "validUntil", GoGetter: "ValidUntil"},
			_jsii_.MemberProperty{JsiiProperty: "validUntilInput", GoGetter: "ValidUntilInput"},
		},
		func() interface{} {
			j := jsiiProxy_Role{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"postgresql.role.RoleConfig",
		reflect.TypeOf((*RoleConfig)(nil)).Elem(),
	)
}
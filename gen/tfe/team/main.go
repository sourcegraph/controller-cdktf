package team

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-tfe.team.Team",
		reflect.TypeOf((*Team)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
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
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "organization", GoGetter: "Organization"},
			_jsii_.MemberProperty{JsiiProperty: "organizationAccess", GoGetter: "OrganizationAccess"},
			_jsii_.MemberProperty{JsiiProperty: "organizationAccessInput", GoGetter: "OrganizationAccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "organizationInput", GoGetter: "OrganizationInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putOrganizationAccess", GoMethod: "PutOrganizationAccess"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrganization", GoMethod: "ResetOrganization"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrganizationAccess", GoMethod: "ResetOrganizationAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSsoTeamId", GoMethod: "ResetSsoTeamId"},
			_jsii_.MemberMethod{JsiiMethod: "resetVisibility", GoMethod: "ResetVisibility"},
			_jsii_.MemberProperty{JsiiProperty: "ssoTeamId", GoGetter: "SsoTeamId"},
			_jsii_.MemberProperty{JsiiProperty: "ssoTeamIdInput", GoGetter: "SsoTeamIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "visibility", GoGetter: "Visibility"},
			_jsii_.MemberProperty{JsiiProperty: "visibilityInput", GoGetter: "VisibilityInput"},
		},
		func() interface{} {
			j := jsiiProxy_Team{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-tfe.team.TeamConfig",
		reflect.TypeOf((*TeamConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-tfe.team.TeamOrganizationAccess",
		reflect.TypeOf((*TeamOrganizationAccess)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-tfe.team.TeamOrganizationAccessOutputReference",
		reflect.TypeOf((*TeamOrganizationAccessOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "manageModules", GoGetter: "ManageModules"},
			_jsii_.MemberProperty{JsiiProperty: "manageModulesInput", GoGetter: "ManageModulesInput"},
			_jsii_.MemberProperty{JsiiProperty: "managePolicies", GoGetter: "ManagePolicies"},
			_jsii_.MemberProperty{JsiiProperty: "managePoliciesInput", GoGetter: "ManagePoliciesInput"},
			_jsii_.MemberProperty{JsiiProperty: "managePolicyOverrides", GoGetter: "ManagePolicyOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "managePolicyOverridesInput", GoGetter: "ManagePolicyOverridesInput"},
			_jsii_.MemberProperty{JsiiProperty: "manageProjects", GoGetter: "ManageProjects"},
			_jsii_.MemberProperty{JsiiProperty: "manageProjectsInput", GoGetter: "ManageProjectsInput"},
			_jsii_.MemberProperty{JsiiProperty: "manageProviders", GoGetter: "ManageProviders"},
			_jsii_.MemberProperty{JsiiProperty: "manageProvidersInput", GoGetter: "ManageProvidersInput"},
			_jsii_.MemberProperty{JsiiProperty: "manageRunTasks", GoGetter: "ManageRunTasks"},
			_jsii_.MemberProperty{JsiiProperty: "manageRunTasksInput", GoGetter: "ManageRunTasksInput"},
			_jsii_.MemberProperty{JsiiProperty: "manageVcsSettings", GoGetter: "ManageVcsSettings"},
			_jsii_.MemberProperty{JsiiProperty: "manageVcsSettingsInput", GoGetter: "ManageVcsSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "manageWorkspaces", GoGetter: "ManageWorkspaces"},
			_jsii_.MemberProperty{JsiiProperty: "manageWorkspacesInput", GoGetter: "ManageWorkspacesInput"},
			_jsii_.MemberProperty{JsiiProperty: "readProjects", GoGetter: "ReadProjects"},
			_jsii_.MemberProperty{JsiiProperty: "readProjectsInput", GoGetter: "ReadProjectsInput"},
			_jsii_.MemberProperty{JsiiProperty: "readWorkspaces", GoGetter: "ReadWorkspaces"},
			_jsii_.MemberProperty{JsiiProperty: "readWorkspacesInput", GoGetter: "ReadWorkspacesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetManageModules", GoMethod: "ResetManageModules"},
			_jsii_.MemberMethod{JsiiMethod: "resetManagePolicies", GoMethod: "ResetManagePolicies"},
			_jsii_.MemberMethod{JsiiMethod: "resetManagePolicyOverrides", GoMethod: "ResetManagePolicyOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetManageProjects", GoMethod: "ResetManageProjects"},
			_jsii_.MemberMethod{JsiiMethod: "resetManageProviders", GoMethod: "ResetManageProviders"},
			_jsii_.MemberMethod{JsiiMethod: "resetManageRunTasks", GoMethod: "ResetManageRunTasks"},
			_jsii_.MemberMethod{JsiiMethod: "resetManageVcsSettings", GoMethod: "ResetManageVcsSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetManageWorkspaces", GoMethod: "ResetManageWorkspaces"},
			_jsii_.MemberMethod{JsiiMethod: "resetReadProjects", GoMethod: "ResetReadProjects"},
			_jsii_.MemberMethod{JsiiMethod: "resetReadWorkspaces", GoMethod: "ResetReadWorkspaces"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TeamOrganizationAccessOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}

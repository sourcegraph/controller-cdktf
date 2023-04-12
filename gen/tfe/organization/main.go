package organization

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-tfe.organization.Organization",
		reflect.TypeOf((*Organization)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowForceDeleteWorkspaces", GoGetter: "AllowForceDeleteWorkspaces"},
			_jsii_.MemberProperty{JsiiProperty: "allowForceDeleteWorkspacesInput", GoGetter: "AllowForceDeleteWorkspacesInput"},
			_jsii_.MemberProperty{JsiiProperty: "assessmentsEnforced", GoGetter: "AssessmentsEnforced"},
			_jsii_.MemberProperty{JsiiProperty: "assessmentsEnforcedInput", GoGetter: "AssessmentsEnforcedInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "collaboratorAuthPolicy", GoGetter: "CollaboratorAuthPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "collaboratorAuthPolicyInput", GoGetter: "CollaboratorAuthPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "costEstimationEnabled", GoGetter: "CostEstimationEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "costEstimationEnabledInput", GoGetter: "CostEstimationEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultProjectId", GoGetter: "DefaultProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "email", GoGetter: "Email"},
			_jsii_.MemberProperty{JsiiProperty: "emailInput", GoGetter: "EmailInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ownersTeamSamlRoleId", GoGetter: "OwnersTeamSamlRoleId"},
			_jsii_.MemberProperty{JsiiProperty: "ownersTeamSamlRoleIdInput", GoGetter: "OwnersTeamSamlRoleIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowForceDeleteWorkspaces", GoMethod: "ResetAllowForceDeleteWorkspaces"},
			_jsii_.MemberMethod{JsiiMethod: "resetAssessmentsEnforced", GoMethod: "ResetAssessmentsEnforced"},
			_jsii_.MemberMethod{JsiiMethod: "resetCollaboratorAuthPolicy", GoMethod: "ResetCollaboratorAuthPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetCostEstimationEnabled", GoMethod: "ResetCostEstimationEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOwnersTeamSamlRoleId", GoMethod: "ResetOwnersTeamSamlRoleId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSendPassingStatusesForUntriggeredSpeculativePlans", GoMethod: "ResetSendPassingStatusesForUntriggeredSpeculativePlans"},
			_jsii_.MemberMethod{JsiiMethod: "resetSessionRememberMinutes", GoMethod: "ResetSessionRememberMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetSessionTimeoutMinutes", GoMethod: "ResetSessionTimeoutMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "sendPassingStatusesForUntriggeredSpeculativePlans", GoGetter: "SendPassingStatusesForUntriggeredSpeculativePlans"},
			_jsii_.MemberProperty{JsiiProperty: "sendPassingStatusesForUntriggeredSpeculativePlansInput", GoGetter: "SendPassingStatusesForUntriggeredSpeculativePlansInput"},
			_jsii_.MemberProperty{JsiiProperty: "sessionRememberMinutes", GoGetter: "SessionRememberMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "sessionRememberMinutesInput", GoGetter: "SessionRememberMinutesInput"},
			_jsii_.MemberProperty{JsiiProperty: "sessionTimeoutMinutes", GoGetter: "SessionTimeoutMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "sessionTimeoutMinutesInput", GoGetter: "SessionTimeoutMinutesInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_Organization{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-tfe.organization.OrganizationConfig",
		reflect.TypeOf((*OrganizationConfig)(nil)).Elem(),
	)
}

package policyrulepassword

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.policyRulePassword.PolicyRulePassword",
		reflect.TypeOf((*PolicyRulePassword)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "networkConnection", GoGetter: "NetworkConnection"},
			_jsii_.MemberProperty{JsiiProperty: "networkConnectionInput", GoGetter: "NetworkConnectionInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkExcludes", GoGetter: "NetworkExcludes"},
			_jsii_.MemberProperty{JsiiProperty: "networkExcludesInput", GoGetter: "NetworkExcludesInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkIncludes", GoGetter: "NetworkIncludes"},
			_jsii_.MemberProperty{JsiiProperty: "networkIncludesInput", GoGetter: "NetworkIncludesInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "passwordChange", GoGetter: "PasswordChange"},
			_jsii_.MemberProperty{JsiiProperty: "passwordChangeInput", GoGetter: "PasswordChangeInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordReset", GoGetter: "PasswordReset"},
			_jsii_.MemberProperty{JsiiProperty: "passwordResetInput", GoGetter: "PasswordResetInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordUnlock", GoGetter: "PasswordUnlock"},
			_jsii_.MemberProperty{JsiiProperty: "passwordUnlockInput", GoGetter: "PasswordUnlockInput"},
			_jsii_.MemberProperty{JsiiProperty: "policyid", GoGetter: "Policyid"},
			_jsii_.MemberProperty{JsiiProperty: "policyId", GoGetter: "PolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "policyidInput", GoGetter: "PolicyidInput"},
			_jsii_.MemberProperty{JsiiProperty: "policyIdInput", GoGetter: "PolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "priority", GoGetter: "Priority"},
			_jsii_.MemberProperty{JsiiProperty: "priorityInput", GoGetter: "PriorityInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkConnection", GoMethod: "ResetNetworkConnection"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkExcludes", GoMethod: "ResetNetworkExcludes"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkIncludes", GoMethod: "ResetNetworkIncludes"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordChange", GoMethod: "ResetPasswordChange"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordReset", GoMethod: "ResetPasswordReset"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordUnlock", GoMethod: "ResetPasswordUnlock"},
			_jsii_.MemberMethod{JsiiMethod: "resetPolicyid", GoMethod: "ResetPolicyid"},
			_jsii_.MemberMethod{JsiiMethod: "resetPolicyId", GoMethod: "ResetPolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPriority", GoMethod: "ResetPriority"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatus", GoMethod: "ResetStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsersExcluded", GoMethod: "ResetUsersExcluded"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusInput", GoGetter: "StatusInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "usersExcluded", GoGetter: "UsersExcluded"},
			_jsii_.MemberProperty{JsiiProperty: "usersExcludedInput", GoGetter: "UsersExcludedInput"},
		},
		func() interface{} {
			j := jsiiProxy_PolicyRulePassword{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.policyRulePassword.PolicyRulePasswordConfig",
		reflect.TypeOf((*PolicyRulePasswordConfig)(nil)).Elem(),
	)
}

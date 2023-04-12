package policyrulesignon

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.policyRuleSignon.PolicyRuleSignon",
		reflect.TypeOf((*PolicyRuleSignon)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "access", GoGetter: "Access"},
			_jsii_.MemberProperty{JsiiProperty: "accessInput", GoGetter: "AccessInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "authtype", GoGetter: "Authtype"},
			_jsii_.MemberProperty{JsiiProperty: "authtypeInput", GoGetter: "AuthtypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "behaviors", GoGetter: "Behaviors"},
			_jsii_.MemberProperty{JsiiProperty: "behaviorsInput", GoGetter: "BehaviorsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "factorSequence", GoGetter: "FactorSequence"},
			_jsii_.MemberProperty{JsiiProperty: "factorSequenceInput", GoGetter: "FactorSequenceInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "identityProvider", GoGetter: "IdentityProvider"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIds", GoGetter: "IdentityProviderIds"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdsInput", GoGetter: "IdentityProviderIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderInput", GoGetter: "IdentityProviderInput"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "mfaLifetime", GoGetter: "MfaLifetime"},
			_jsii_.MemberProperty{JsiiProperty: "mfaLifetimeInput", GoGetter: "MfaLifetimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "mfaPrompt", GoGetter: "MfaPrompt"},
			_jsii_.MemberProperty{JsiiProperty: "mfaPromptInput", GoGetter: "MfaPromptInput"},
			_jsii_.MemberProperty{JsiiProperty: "mfaRememberDevice", GoGetter: "MfaRememberDevice"},
			_jsii_.MemberProperty{JsiiProperty: "mfaRememberDeviceInput", GoGetter: "MfaRememberDeviceInput"},
			_jsii_.MemberProperty{JsiiProperty: "mfaRequired", GoGetter: "MfaRequired"},
			_jsii_.MemberProperty{JsiiProperty: "mfaRequiredInput", GoGetter: "MfaRequiredInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "policyid", GoGetter: "Policyid"},
			_jsii_.MemberProperty{JsiiProperty: "policyId", GoGetter: "PolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "policyidInput", GoGetter: "PolicyidInput"},
			_jsii_.MemberProperty{JsiiProperty: "policyIdInput", GoGetter: "PolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "primaryFactor", GoGetter: "PrimaryFactor"},
			_jsii_.MemberProperty{JsiiProperty: "primaryFactorInput", GoGetter: "PrimaryFactorInput"},
			_jsii_.MemberProperty{JsiiProperty: "priority", GoGetter: "Priority"},
			_jsii_.MemberProperty{JsiiProperty: "priorityInput", GoGetter: "PriorityInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putFactorSequence", GoMethod: "PutFactorSequence"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccess", GoMethod: "ResetAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthtype", GoMethod: "ResetAuthtype"},
			_jsii_.MemberMethod{JsiiMethod: "resetBehaviors", GoMethod: "ResetBehaviors"},
			_jsii_.MemberMethod{JsiiMethod: "resetFactorSequence", GoMethod: "ResetFactorSequence"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentityProvider", GoMethod: "ResetIdentityProvider"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentityProviderIds", GoMethod: "ResetIdentityProviderIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetMfaLifetime", GoMethod: "ResetMfaLifetime"},
			_jsii_.MemberMethod{JsiiMethod: "resetMfaPrompt", GoMethod: "ResetMfaPrompt"},
			_jsii_.MemberMethod{JsiiMethod: "resetMfaRememberDevice", GoMethod: "ResetMfaRememberDevice"},
			_jsii_.MemberMethod{JsiiMethod: "resetMfaRequired", GoMethod: "ResetMfaRequired"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkConnection", GoMethod: "ResetNetworkConnection"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkExcludes", GoMethod: "ResetNetworkExcludes"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkIncludes", GoMethod: "ResetNetworkIncludes"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPolicyid", GoMethod: "ResetPolicyid"},
			_jsii_.MemberMethod{JsiiMethod: "resetPolicyId", GoMethod: "ResetPolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrimaryFactor", GoMethod: "ResetPrimaryFactor"},
			_jsii_.MemberMethod{JsiiMethod: "resetPriority", GoMethod: "ResetPriority"},
			_jsii_.MemberMethod{JsiiMethod: "resetRiscLevel", GoMethod: "ResetRiscLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetSessionIdle", GoMethod: "ResetSessionIdle"},
			_jsii_.MemberMethod{JsiiMethod: "resetSessionLifetime", GoMethod: "ResetSessionLifetime"},
			_jsii_.MemberMethod{JsiiMethod: "resetSessionPersistent", GoMethod: "ResetSessionPersistent"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatus", GoMethod: "ResetStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsersExcluded", GoMethod: "ResetUsersExcluded"},
			_jsii_.MemberProperty{JsiiProperty: "riscLevel", GoGetter: "RiscLevel"},
			_jsii_.MemberProperty{JsiiProperty: "riscLevelInput", GoGetter: "RiscLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "sessionIdle", GoGetter: "SessionIdle"},
			_jsii_.MemberProperty{JsiiProperty: "sessionIdleInput", GoGetter: "SessionIdleInput"},
			_jsii_.MemberProperty{JsiiProperty: "sessionLifetime", GoGetter: "SessionLifetime"},
			_jsii_.MemberProperty{JsiiProperty: "sessionLifetimeInput", GoGetter: "SessionLifetimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "sessionPersistent", GoGetter: "SessionPersistent"},
			_jsii_.MemberProperty{JsiiProperty: "sessionPersistentInput", GoGetter: "SessionPersistentInput"},
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
			j := jsiiProxy_PolicyRuleSignon{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.policyRuleSignon.PolicyRuleSignonConfig",
		reflect.TypeOf((*PolicyRuleSignonConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.policyRuleSignon.PolicyRuleSignonFactorSequence",
		reflect.TypeOf((*PolicyRuleSignonFactorSequence)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.policyRuleSignon.PolicyRuleSignonFactorSequenceList",
		reflect.TypeOf((*PolicyRuleSignonFactorSequenceList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_PolicyRuleSignonFactorSequenceList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.policyRuleSignon.PolicyRuleSignonFactorSequenceOutputReference",
		reflect.TypeOf((*PolicyRuleSignonFactorSequenceOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "primaryCriteriaFactorType", GoGetter: "PrimaryCriteriaFactorType"},
			_jsii_.MemberProperty{JsiiProperty: "primaryCriteriaFactorTypeInput", GoGetter: "PrimaryCriteriaFactorTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "primaryCriteriaProvider", GoGetter: "PrimaryCriteriaProvider"},
			_jsii_.MemberProperty{JsiiProperty: "primaryCriteriaProviderInput", GoGetter: "PrimaryCriteriaProviderInput"},
			_jsii_.MemberMethod{JsiiMethod: "putSecondaryCriteria", GoMethod: "PutSecondaryCriteria"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecondaryCriteria", GoMethod: "ResetSecondaryCriteria"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryCriteria", GoGetter: "SecondaryCriteria"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryCriteriaInput", GoGetter: "SecondaryCriteriaInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PolicyRuleSignonFactorSequenceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.policyRuleSignon.PolicyRuleSignonFactorSequenceSecondaryCriteria",
		reflect.TypeOf((*PolicyRuleSignonFactorSequenceSecondaryCriteria)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.policyRuleSignon.PolicyRuleSignonFactorSequenceSecondaryCriteriaList",
		reflect.TypeOf((*PolicyRuleSignonFactorSequenceSecondaryCriteriaList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_PolicyRuleSignonFactorSequenceSecondaryCriteriaList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.policyRuleSignon.PolicyRuleSignonFactorSequenceSecondaryCriteriaOutputReference",
		reflect.TypeOf((*PolicyRuleSignonFactorSequenceSecondaryCriteriaOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "factorType", GoGetter: "FactorType"},
			_jsii_.MemberProperty{JsiiProperty: "factorTypeInput", GoGetter: "FactorTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "providerInput", GoGetter: "ProviderInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PolicyRuleSignonFactorSequenceSecondaryCriteriaOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}

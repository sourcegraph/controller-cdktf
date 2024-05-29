package googlenetworkserviceshttproute

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRoute",
		reflect.TypeOf((*GoogleNetworkServicesHttpRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createTime", GoGetter: "CreateTime"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "effectiveLabels", GoGetter: "EffectiveLabels"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "gateways", GoGetter: "Gateways"},
			_jsii_.MemberProperty{JsiiProperty: "gatewaysInput", GoGetter: "GatewaysInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "hostnames", GoGetter: "Hostnames"},
			_jsii_.MemberProperty{JsiiProperty: "hostnamesInput", GoGetter: "HostnamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "labelsInput", GoGetter: "LabelsInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "meshes", GoGetter: "Meshes"},
			_jsii_.MemberProperty{JsiiProperty: "meshesInput", GoGetter: "MeshesInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projectInput", GoGetter: "ProjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putRules", GoMethod: "PutRules"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetGateways", GoMethod: "ResetGateways"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabels", GoMethod: "ResetLabels"},
			_jsii_.MemberMethod{JsiiMethod: "resetMeshes", GoMethod: "ResetMeshes"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProject", GoMethod: "ResetProject"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rules", GoGetter: "Rules"},
			_jsii_.MemberProperty{JsiiProperty: "rulesInput", GoGetter: "RulesInput"},
			_jsii_.MemberProperty{JsiiProperty: "selfLink", GoGetter: "SelfLink"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformLabels", GoGetter: "TerraformLabels"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "updateTime", GoGetter: "UpdateTime"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleNetworkServicesHttpRoute{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteConfig",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRules",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRules)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesAction",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesAction)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionCorsPolicy",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionCorsPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionCorsPolicyOutputReference",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionCorsPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowCredentials", GoGetter: "AllowCredentials"},
			_jsii_.MemberProperty{JsiiProperty: "allowCredentialsInput", GoGetter: "AllowCredentialsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowHeaders", GoGetter: "AllowHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "allowHeadersInput", GoGetter: "AllowHeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowMethods", GoGetter: "AllowMethods"},
			_jsii_.MemberProperty{JsiiProperty: "allowMethodsInput", GoGetter: "AllowMethodsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowOriginRegexes", GoGetter: "AllowOriginRegexes"},
			_jsii_.MemberProperty{JsiiProperty: "allowOriginRegexesInput", GoGetter: "AllowOriginRegexesInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowOrigins", GoGetter: "AllowOrigins"},
			_jsii_.MemberProperty{JsiiProperty: "allowOriginsInput", GoGetter: "AllowOriginsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "disabled", GoGetter: "Disabled"},
			_jsii_.MemberProperty{JsiiProperty: "disabledInput", GoGetter: "DisabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "exposeHeaders", GoGetter: "ExposeHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "exposeHeadersInput", GoGetter: "ExposeHeadersInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "maxAge", GoGetter: "MaxAge"},
			_jsii_.MemberProperty{JsiiProperty: "maxAgeInput", GoGetter: "MaxAgeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowCredentials", GoMethod: "ResetAllowCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowHeaders", GoMethod: "ResetAllowHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowMethods", GoMethod: "ResetAllowMethods"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowOriginRegexes", GoMethod: "ResetAllowOriginRegexes"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowOrigins", GoMethod: "ResetAllowOrigins"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisabled", GoMethod: "ResetDisabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetExposeHeaders", GoMethod: "ResetExposeHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxAge", GoMethod: "ResetMaxAge"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionCorsPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionDestinations",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionDestinations)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionDestinationsList",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionDestinationsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionDestinationsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionDestinationsOutputReference",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionDestinationsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetServiceName", GoMethod: "ResetServiceName"},
			_jsii_.MemberMethod{JsiiMethod: "resetWeight", GoMethod: "ResetWeight"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "serviceName", GoGetter: "ServiceName"},
			_jsii_.MemberProperty{JsiiProperty: "serviceNameInput", GoGetter: "ServiceNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "weight", GoGetter: "Weight"},
			_jsii_.MemberProperty{JsiiProperty: "weightInput", GoGetter: "WeightInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionDestinationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicy",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicyAbort",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicyAbort)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicyAbortOutputReference",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicyAbortOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "httpStatus", GoGetter: "HttpStatus"},
			_jsii_.MemberProperty{JsiiProperty: "httpStatusInput", GoGetter: "HttpStatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "percentage", GoGetter: "Percentage"},
			_jsii_.MemberProperty{JsiiProperty: "percentageInput", GoGetter: "PercentageInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpStatus", GoMethod: "ResetHttpStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetPercentage", GoMethod: "ResetPercentage"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicyAbortOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicyDelay",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicyDelay)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicyDelayOutputReference",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicyDelayOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fixedDelay", GoGetter: "FixedDelay"},
			_jsii_.MemberProperty{JsiiProperty: "fixedDelayInput", GoGetter: "FixedDelayInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "percentage", GoGetter: "Percentage"},
			_jsii_.MemberProperty{JsiiProperty: "percentageInput", GoGetter: "PercentageInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetFixedDelay", GoMethod: "ResetFixedDelay"},
			_jsii_.MemberMethod{JsiiMethod: "resetPercentage", GoMethod: "ResetPercentage"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicyDelayOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicyOutputReference",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "abort", GoGetter: "Abort"},
			_jsii_.MemberProperty{JsiiProperty: "abortInput", GoGetter: "AbortInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delay", GoGetter: "Delay"},
			_jsii_.MemberProperty{JsiiProperty: "delayInput", GoGetter: "DelayInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putAbort", GoMethod: "PutAbort"},
			_jsii_.MemberMethod{JsiiMethod: "putDelay", GoMethod: "PutDelay"},
			_jsii_.MemberMethod{JsiiMethod: "resetAbort", GoMethod: "ResetAbort"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelay", GoMethod: "ResetDelay"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionOutputReference",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "corsPolicy", GoGetter: "CorsPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "corsPolicyInput", GoGetter: "CorsPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destinations", GoGetter: "Destinations"},
			_jsii_.MemberProperty{JsiiProperty: "destinationsInput", GoGetter: "DestinationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "faultInjectionPolicy", GoGetter: "FaultInjectionPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "faultInjectionPolicyInput", GoGetter: "FaultInjectionPolicyInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putCorsPolicy", GoMethod: "PutCorsPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putDestinations", GoMethod: "PutDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "putFaultInjectionPolicy", GoMethod: "PutFaultInjectionPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putRedirect", GoMethod: "PutRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "putRequestHeaderModifier", GoMethod: "PutRequestHeaderModifier"},
			_jsii_.MemberMethod{JsiiMethod: "putRequestMirrorPolicy", GoMethod: "PutRequestMirrorPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putResponseHeaderModifier", GoMethod: "PutResponseHeaderModifier"},
			_jsii_.MemberMethod{JsiiMethod: "putRetryPolicy", GoMethod: "PutRetryPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putUrlRewrite", GoMethod: "PutUrlRewrite"},
			_jsii_.MemberProperty{JsiiProperty: "redirect", GoGetter: "Redirect"},
			_jsii_.MemberProperty{JsiiProperty: "redirectInput", GoGetter: "RedirectInput"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeaderModifier", GoGetter: "RequestHeaderModifier"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeaderModifierInput", GoGetter: "RequestHeaderModifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "requestMirrorPolicy", GoGetter: "RequestMirrorPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "requestMirrorPolicyInput", GoGetter: "RequestMirrorPolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCorsPolicy", GoMethod: "ResetCorsPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestinations", GoMethod: "ResetDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resetFaultInjectionPolicy", GoMethod: "ResetFaultInjectionPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedirect", GoMethod: "ResetRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestHeaderModifier", GoMethod: "ResetRequestHeaderModifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestMirrorPolicy", GoMethod: "ResetRequestMirrorPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseHeaderModifier", GoMethod: "ResetResponseHeaderModifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryPolicy", GoMethod: "ResetRetryPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeout", GoMethod: "ResetTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetUrlRewrite", GoMethod: "ResetUrlRewrite"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeaderModifier", GoGetter: "ResponseHeaderModifier"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeaderModifierInput", GoGetter: "ResponseHeaderModifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "retryPolicy", GoGetter: "RetryPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "retryPolicyInput", GoGetter: "RetryPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutInput", GoGetter: "TimeoutInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "urlRewrite", GoGetter: "UrlRewrite"},
			_jsii_.MemberProperty{JsiiProperty: "urlRewriteInput", GoGetter: "UrlRewriteInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionRedirect",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionRedirect)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionRedirectOutputReference",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionRedirectOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "hostRedirect", GoGetter: "HostRedirect"},
			_jsii_.MemberProperty{JsiiProperty: "hostRedirectInput", GoGetter: "HostRedirectInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpsRedirect", GoGetter: "HttpsRedirect"},
			_jsii_.MemberProperty{JsiiProperty: "httpsRedirectInput", GoGetter: "HttpsRedirectInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "pathRedirect", GoGetter: "PathRedirect"},
			_jsii_.MemberProperty{JsiiProperty: "pathRedirectInput", GoGetter: "PathRedirectInput"},
			_jsii_.MemberProperty{JsiiProperty: "portRedirect", GoGetter: "PortRedirect"},
			_jsii_.MemberProperty{JsiiProperty: "portRedirectInput", GoGetter: "PortRedirectInput"},
			_jsii_.MemberProperty{JsiiProperty: "prefixRewrite", GoGetter: "PrefixRewrite"},
			_jsii_.MemberProperty{JsiiProperty: "prefixRewriteInput", GoGetter: "PrefixRewriteInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostRedirect", GoMethod: "ResetHostRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpsRedirect", GoMethod: "ResetHttpsRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetPathRedirect", GoMethod: "ResetPathRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetPortRedirect", GoMethod: "ResetPortRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefixRewrite", GoMethod: "ResetPrefixRewrite"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseCode", GoMethod: "ResetResponseCode"},
			_jsii_.MemberMethod{JsiiMethod: "resetStripQuery", GoMethod: "ResetStripQuery"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "responseCode", GoGetter: "ResponseCode"},
			_jsii_.MemberProperty{JsiiProperty: "responseCodeInput", GoGetter: "ResponseCodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "stripQuery", GoGetter: "StripQuery"},
			_jsii_.MemberProperty{JsiiProperty: "stripQueryInput", GoGetter: "StripQueryInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionRedirectOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionRequestHeaderModifier",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionRequestHeaderModifier)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionRequestHeaderModifierOutputReference",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionRequestHeaderModifierOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "add", GoGetter: "Add"},
			_jsii_.MemberProperty{JsiiProperty: "addInput", GoGetter: "AddInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "remove", GoGetter: "Remove"},
			_jsii_.MemberProperty{JsiiProperty: "removeInput", GoGetter: "RemoveInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdd", GoMethod: "ResetAdd"},
			_jsii_.MemberMethod{JsiiMethod: "resetRemove", GoMethod: "ResetRemove"},
			_jsii_.MemberMethod{JsiiMethod: "resetSet", GoMethod: "ResetSet"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "set", GoGetter: "Set"},
			_jsii_.MemberProperty{JsiiProperty: "setInput", GoGetter: "SetInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionRequestHeaderModifierOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionRequestMirrorPolicy",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionRequestMirrorPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionRequestMirrorPolicyDestination",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionRequestMirrorPolicyDestination)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionRequestMirrorPolicyDestinationOutputReference",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionRequestMirrorPolicyDestinationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetServiceName", GoMethod: "ResetServiceName"},
			_jsii_.MemberMethod{JsiiMethod: "resetWeight", GoMethod: "ResetWeight"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "serviceName", GoGetter: "ServiceName"},
			_jsii_.MemberProperty{JsiiProperty: "serviceNameInput", GoGetter: "ServiceNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "weight", GoGetter: "Weight"},
			_jsii_.MemberProperty{JsiiProperty: "weightInput", GoGetter: "WeightInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionRequestMirrorPolicyDestinationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionRequestMirrorPolicyOutputReference",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionRequestMirrorPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "destinationInput", GoGetter: "DestinationInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putDestination", GoMethod: "PutDestination"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestination", GoMethod: "ResetDestination"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionRequestMirrorPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionResponseHeaderModifier",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionResponseHeaderModifier)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionResponseHeaderModifierOutputReference",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionResponseHeaderModifierOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "add", GoGetter: "Add"},
			_jsii_.MemberProperty{JsiiProperty: "addInput", GoGetter: "AddInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "remove", GoGetter: "Remove"},
			_jsii_.MemberProperty{JsiiProperty: "removeInput", GoGetter: "RemoveInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdd", GoMethod: "ResetAdd"},
			_jsii_.MemberMethod{JsiiMethod: "resetRemove", GoMethod: "ResetRemove"},
			_jsii_.MemberMethod{JsiiMethod: "resetSet", GoMethod: "ResetSet"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "set", GoGetter: "Set"},
			_jsii_.MemberProperty{JsiiProperty: "setInput", GoGetter: "SetInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionResponseHeaderModifierOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionRetryPolicy",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionRetryPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionRetryPolicyOutputReference",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionRetryPolicyOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "numRetries", GoGetter: "NumRetries"},
			_jsii_.MemberProperty{JsiiProperty: "numRetriesInput", GoGetter: "NumRetriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "perTryTimeout", GoGetter: "PerTryTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "perTryTimeoutInput", GoGetter: "PerTryTimeoutInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNumRetries", GoMethod: "ResetNumRetries"},
			_jsii_.MemberMethod{JsiiMethod: "resetPerTryTimeout", GoMethod: "ResetPerTryTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryConditions", GoMethod: "ResetRetryConditions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retryConditions", GoGetter: "RetryConditions"},
			_jsii_.MemberProperty{JsiiProperty: "retryConditionsInput", GoGetter: "RetryConditionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionRetryPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionUrlRewrite",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionUrlRewrite)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionUrlRewriteOutputReference",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesActionUrlRewriteOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "hostRewrite", GoGetter: "HostRewrite"},
			_jsii_.MemberProperty{JsiiProperty: "hostRewriteInput", GoGetter: "HostRewriteInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "pathPrefixRewrite", GoGetter: "PathPrefixRewrite"},
			_jsii_.MemberProperty{JsiiProperty: "pathPrefixRewriteInput", GoGetter: "PathPrefixRewriteInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostRewrite", GoMethod: "ResetHostRewrite"},
			_jsii_.MemberMethod{JsiiMethod: "resetPathPrefixRewrite", GoMethod: "ResetPathPrefixRewrite"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionUrlRewriteOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesList",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesMatches",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesMatches)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesMatchesHeaders",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesMatchesHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesMatchesHeadersList",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesMatchesHeadersList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesHeadersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesMatchesHeadersOutputReference",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesMatchesHeadersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "exactMatch", GoGetter: "ExactMatch"},
			_jsii_.MemberProperty{JsiiProperty: "exactMatchInput", GoGetter: "ExactMatchInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "header", GoGetter: "Header"},
			_jsii_.MemberProperty{JsiiProperty: "headerInput", GoGetter: "HeaderInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "invertMatch", GoGetter: "InvertMatch"},
			_jsii_.MemberProperty{JsiiProperty: "invertMatchInput", GoGetter: "InvertMatchInput"},
			_jsii_.MemberProperty{JsiiProperty: "prefixMatch", GoGetter: "PrefixMatch"},
			_jsii_.MemberProperty{JsiiProperty: "prefixMatchInput", GoGetter: "PrefixMatchInput"},
			_jsii_.MemberProperty{JsiiProperty: "presentMatch", GoGetter: "PresentMatch"},
			_jsii_.MemberProperty{JsiiProperty: "presentMatchInput", GoGetter: "PresentMatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "putRangeMatch", GoMethod: "PutRangeMatch"},
			_jsii_.MemberProperty{JsiiProperty: "rangeMatch", GoGetter: "RangeMatch"},
			_jsii_.MemberProperty{JsiiProperty: "rangeMatchInput", GoGetter: "RangeMatchInput"},
			_jsii_.MemberProperty{JsiiProperty: "regexMatch", GoGetter: "RegexMatch"},
			_jsii_.MemberProperty{JsiiProperty: "regexMatchInput", GoGetter: "RegexMatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetExactMatch", GoMethod: "ResetExactMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeader", GoMethod: "ResetHeader"},
			_jsii_.MemberMethod{JsiiMethod: "resetInvertMatch", GoMethod: "ResetInvertMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefixMatch", GoMethod: "ResetPrefixMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetPresentMatch", GoMethod: "ResetPresentMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetRangeMatch", GoMethod: "ResetRangeMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegexMatch", GoMethod: "ResetRegexMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuffixMatch", GoMethod: "ResetSuffixMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "suffixMatch", GoGetter: "SuffixMatch"},
			_jsii_.MemberProperty{JsiiProperty: "suffixMatchInput", GoGetter: "SuffixMatchInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesHeadersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesMatchesHeadersRangeMatch",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesMatchesHeadersRangeMatch)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesMatchesHeadersRangeMatchOutputReference",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesMatchesHeadersRangeMatchOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "end", GoGetter: "End"},
			_jsii_.MemberProperty{JsiiProperty: "endInput", GoGetter: "EndInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "start", GoGetter: "Start"},
			_jsii_.MemberProperty{JsiiProperty: "startInput", GoGetter: "StartInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesHeadersRangeMatchOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesMatchesList",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesMatchesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesMatchesOutputReference",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesMatchesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "fullPathMatch", GoGetter: "FullPathMatch"},
			_jsii_.MemberProperty{JsiiProperty: "fullPathMatchInput", GoGetter: "FullPathMatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "headers", GoGetter: "Headers"},
			_jsii_.MemberProperty{JsiiProperty: "headersInput", GoGetter: "HeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreCase", GoGetter: "IgnoreCase"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreCaseInput", GoGetter: "IgnoreCaseInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "prefixMatch", GoGetter: "PrefixMatch"},
			_jsii_.MemberProperty{JsiiProperty: "prefixMatchInput", GoGetter: "PrefixMatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "putHeaders", GoMethod: "PutHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "putQueryParameters", GoMethod: "PutQueryParameters"},
			_jsii_.MemberProperty{JsiiProperty: "queryParameters", GoGetter: "QueryParameters"},
			_jsii_.MemberProperty{JsiiProperty: "queryParametersInput", GoGetter: "QueryParametersInput"},
			_jsii_.MemberProperty{JsiiProperty: "regexMatch", GoGetter: "RegexMatch"},
			_jsii_.MemberProperty{JsiiProperty: "regexMatchInput", GoGetter: "RegexMatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetFullPathMatch", GoMethod: "ResetFullPathMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaders", GoMethod: "ResetHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetIgnoreCase", GoMethod: "ResetIgnoreCase"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefixMatch", GoMethod: "ResetPrefixMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryParameters", GoMethod: "ResetQueryParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegexMatch", GoMethod: "ResetRegexMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesMatchesQueryParameters",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesMatchesQueryParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersList",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "exactMatch", GoGetter: "ExactMatch"},
			_jsii_.MemberProperty{JsiiProperty: "exactMatchInput", GoGetter: "ExactMatchInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "presentMatch", GoGetter: "PresentMatch"},
			_jsii_.MemberProperty{JsiiProperty: "presentMatchInput", GoGetter: "PresentMatchInput"},
			_jsii_.MemberProperty{JsiiProperty: "queryParameter", GoGetter: "QueryParameter"},
			_jsii_.MemberProperty{JsiiProperty: "queryParameterInput", GoGetter: "QueryParameterInput"},
			_jsii_.MemberProperty{JsiiProperty: "regexMatch", GoGetter: "RegexMatch"},
			_jsii_.MemberProperty{JsiiProperty: "regexMatchInput", GoGetter: "RegexMatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetExactMatch", GoMethod: "ResetExactMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetPresentMatch", GoMethod: "ResetPresentMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryParameter", GoMethod: "ResetQueryParameter"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegexMatch", GoMethod: "ResetRegexMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesOutputReference",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteRulesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "action", GoGetter: "Action"},
			_jsii_.MemberProperty{JsiiProperty: "actionInput", GoGetter: "ActionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "matches", GoGetter: "Matches"},
			_jsii_.MemberProperty{JsiiProperty: "matchesInput", GoGetter: "MatchesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAction", GoMethod: "PutAction"},
			_jsii_.MemberMethod{JsiiMethod: "putMatches", GoMethod: "PutMatches"},
			_jsii_.MemberMethod{JsiiMethod: "resetAction", GoMethod: "ResetAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetMatches", GoMethod: "ResetMatches"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteTimeouts",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteTimeoutsOutputReference",
		reflect.TypeOf((*GoogleNetworkServicesHttpRouteTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleNetworkServicesHttpRouteTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}

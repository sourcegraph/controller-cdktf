package computeregionurlmap

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMap",
		reflect.TypeFor[ComputeRegionUrlMap](),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "creationTimestamp", GoGetter: "CreationTimestamp"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRouteAction", GoGetter: "DefaultRouteAction"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRouteActionInput", GoGetter: "DefaultRouteActionInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultService", GoGetter: "DefaultService"},
			_jsii_.MemberProperty{JsiiProperty: "defaultServiceInput", GoGetter: "DefaultServiceInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultUrlRedirect", GoGetter: "DefaultUrlRedirect"},
			_jsii_.MemberProperty{JsiiProperty: "defaultUrlRedirectInput", GoGetter: "DefaultUrlRedirectInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "fingerprint", GoGetter: "Fingerprint"},
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
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "headerAction", GoGetter: "HeaderAction"},
			_jsii_.MemberProperty{JsiiProperty: "headerActionInput", GoGetter: "HeaderActionInput"},
			_jsii_.MemberProperty{JsiiProperty: "hostRule", GoGetter: "HostRule"},
			_jsii_.MemberProperty{JsiiProperty: "hostRuleInput", GoGetter: "HostRuleInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "mapId", GoGetter: "MapId"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pathMatcher", GoGetter: "PathMatcher"},
			_jsii_.MemberProperty{JsiiProperty: "pathMatcherInput", GoGetter: "PathMatcherInput"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projectInput", GoGetter: "ProjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultRouteAction", GoMethod: "PutDefaultRouteAction"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultUrlRedirect", GoMethod: "PutDefaultUrlRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "putHeaderAction", GoMethod: "PutHeaderAction"},
			_jsii_.MemberMethod{JsiiMethod: "putHostRule", GoMethod: "PutHostRule"},
			_jsii_.MemberMethod{JsiiMethod: "putPathMatcher", GoMethod: "PutPathMatcher"},
			_jsii_.MemberMethod{JsiiMethod: "putTest", GoMethod: "PutTest"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultRouteAction", GoMethod: "ResetDefaultRouteAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultService", GoMethod: "ResetDefaultService"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultUrlRedirect", GoMethod: "ResetDefaultUrlRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderAction", GoMethod: "ResetHeaderAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostRule", GoMethod: "ResetHostRule"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPathMatcher", GoMethod: "ResetPathMatcher"},
			_jsii_.MemberMethod{JsiiMethod: "resetProject", GoMethod: "ResetProject"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetTest", GoMethod: "ResetTest"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "selfLink", GoGetter: "SelfLink"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "test", GoGetter: "Test"},
			_jsii_.MemberProperty{JsiiProperty: "testInput", GoGetter: "TestInput"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMap{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapConfig",
		reflect.TypeFor[ComputeRegionUrlMapConfig](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteAction",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteAction](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionCorsPolicy",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionCorsPolicy](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionCorsPolicyOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionCorsPolicyOutputReference](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapDefaultRouteActionCorsPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicy",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicy](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyAbort",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyAbort](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyAbortOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyAbortOutputReference](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyAbortOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyDelay",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyDelay](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelayOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelayOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "nanos", GoGetter: "Nanos"},
			_jsii_.MemberProperty{JsiiProperty: "nanosInput", GoGetter: "NanosInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNanos", GoMethod: "ResetNanos"},
			_jsii_.MemberMethod{JsiiMethod: "resetSeconds", GoMethod: "ResetSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "seconds", GoGetter: "Seconds"},
			_jsii_.MemberProperty{JsiiProperty: "secondsInput", GoGetter: "SecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelayOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyDelayOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyDelayOutputReference](),
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
			_jsii_.MemberMethod{JsiiMethod: "putFixedDelay", GoMethod: "PutFixedDelay"},
			_jsii_.MemberMethod{JsiiMethod: "resetFixedDelay", GoMethod: "ResetFixedDelay"},
			_jsii_.MemberMethod{JsiiMethod: "resetPercentage", GoMethod: "ResetPercentage"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyDelayOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyOutputReference](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "corsPolicy", GoGetter: "CorsPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "corsPolicyInput", GoGetter: "CorsPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberMethod{JsiiMethod: "putFaultInjectionPolicy", GoMethod: "PutFaultInjectionPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putRequestMirrorPolicy", GoMethod: "PutRequestMirrorPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putRetryPolicy", GoMethod: "PutRetryPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeout", GoMethod: "PutTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "putUrlRewrite", GoMethod: "PutUrlRewrite"},
			_jsii_.MemberMethod{JsiiMethod: "putWeightedBackendServices", GoMethod: "PutWeightedBackendServices"},
			_jsii_.MemberProperty{JsiiProperty: "requestMirrorPolicy", GoGetter: "RequestMirrorPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "requestMirrorPolicyInput", GoGetter: "RequestMirrorPolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCorsPolicy", GoMethod: "ResetCorsPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetFaultInjectionPolicy", GoMethod: "ResetFaultInjectionPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestMirrorPolicy", GoMethod: "ResetRequestMirrorPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryPolicy", GoMethod: "ResetRetryPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeout", GoMethod: "ResetTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetUrlRewrite", GoMethod: "ResetUrlRewrite"},
			_jsii_.MemberMethod{JsiiMethod: "resetWeightedBackendServices", GoMethod: "ResetWeightedBackendServices"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retryPolicy", GoGetter: "RetryPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "retryPolicyInput", GoGetter: "RetryPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutInput", GoGetter: "TimeoutInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "urlRewrite", GoGetter: "UrlRewrite"},
			_jsii_.MemberProperty{JsiiProperty: "urlRewriteInput", GoGetter: "UrlRewriteInput"},
			_jsii_.MemberProperty{JsiiProperty: "weightedBackendServices", GoGetter: "WeightedBackendServices"},
			_jsii_.MemberProperty{JsiiProperty: "weightedBackendServicesInput", GoGetter: "WeightedBackendServicesInput"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicy",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicy](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicyOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicyOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "backendService", GoGetter: "BackendService"},
			_jsii_.MemberProperty{JsiiProperty: "backendServiceInput", GoGetter: "BackendServiceInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetBackendService", GoMethod: "ResetBackendService"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionRetryPolicy",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionRetryPolicy](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionRetryPolicyOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionRetryPolicyOutputReference](),
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
			_jsii_.MemberMethod{JsiiMethod: "putPerTryTimeout", GoMethod: "PutPerTryTimeout"},
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapDefaultRouteActionRetryPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionRetryPolicyPerTryTimeout",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionRetryPolicyPerTryTimeout](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionRetryPolicyPerTryTimeoutOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionRetryPolicyPerTryTimeoutOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "nanos", GoGetter: "Nanos"},
			_jsii_.MemberProperty{JsiiProperty: "nanosInput", GoGetter: "NanosInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNanos", GoMethod: "ResetNanos"},
			_jsii_.MemberMethod{JsiiMethod: "resetSeconds", GoMethod: "ResetSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "seconds", GoGetter: "Seconds"},
			_jsii_.MemberProperty{JsiiProperty: "secondsInput", GoGetter: "SecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapDefaultRouteActionRetryPolicyPerTryTimeoutOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionTimeout",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionTimeout](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionTimeoutOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionTimeoutOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "nanos", GoGetter: "Nanos"},
			_jsii_.MemberProperty{JsiiProperty: "nanosInput", GoGetter: "NanosInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNanos", GoMethod: "ResetNanos"},
			_jsii_.MemberMethod{JsiiMethod: "resetSeconds", GoMethod: "ResetSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "seconds", GoGetter: "Seconds"},
			_jsii_.MemberProperty{JsiiProperty: "secondsInput", GoGetter: "SecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapDefaultRouteActionTimeoutOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionUrlRewrite",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionUrlRewrite](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionUrlRewriteOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionUrlRewriteOutputReference](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapDefaultRouteActionUrlRewriteOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionWeightedBackendServices",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionWeightedBackendServices](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderAction",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderAction](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference](),
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
			_jsii_.MemberMethod{JsiiMethod: "putRequestHeadersToAdd", GoMethod: "PutRequestHeadersToAdd"},
			_jsii_.MemberMethod{JsiiMethod: "putResponseHeadersToAdd", GoMethod: "PutResponseHeadersToAdd"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToAdd", GoGetter: "RequestHeadersToAdd"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToAddInput", GoGetter: "RequestHeadersToAddInput"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToRemove", GoGetter: "RequestHeadersToRemove"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToRemoveInput", GoGetter: "RequestHeadersToRemoveInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestHeadersToAdd", GoMethod: "ResetRequestHeadersToAdd"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestHeadersToRemove", GoMethod: "ResetRequestHeadersToRemove"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseHeadersToAdd", GoMethod: "ResetResponseHeadersToAdd"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseHeadersToRemove", GoMethod: "ResetResponseHeadersToRemove"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToAdd", GoGetter: "ResponseHeadersToAdd"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToAddInput", GoGetter: "ResponseHeadersToAddInput"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToRemove", GoGetter: "ResponseHeadersToRemove"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToRemoveInput", GoGetter: "ResponseHeadersToRemoveInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAdd",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAdd](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "headerName", GoGetter: "HeaderName"},
			_jsii_.MemberProperty{JsiiProperty: "headerNameInput", GoGetter: "HeaderNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "headerValue", GoGetter: "HeaderValue"},
			_jsii_.MemberProperty{JsiiProperty: "headerValueInput", GoGetter: "HeaderValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "replace", GoGetter: "Replace"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInput", GoGetter: "ReplaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderName", GoMethod: "ResetHeaderName"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderValue", GoMethod: "ResetHeaderValue"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplace", GoMethod: "ResetReplace"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAdd",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAdd](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddList",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "headerName", GoGetter: "HeaderName"},
			_jsii_.MemberProperty{JsiiProperty: "headerNameInput", GoGetter: "HeaderNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "headerValue", GoGetter: "HeaderValue"},
			_jsii_.MemberProperty{JsiiProperty: "headerValueInput", GoGetter: "HeaderValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "replace", GoGetter: "Replace"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInput", GoGetter: "ReplaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderName", GoMethod: "ResetHeaderName"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderValue", GoMethod: "ResetHeaderValue"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplace", GoMethod: "ResetReplace"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesList",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "backendService", GoGetter: "BackendService"},
			_jsii_.MemberProperty{JsiiProperty: "backendServiceInput", GoGetter: "BackendServiceInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "headerAction", GoGetter: "HeaderAction"},
			_jsii_.MemberProperty{JsiiProperty: "headerActionInput", GoGetter: "HeaderActionInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putHeaderAction", GoMethod: "PutHeaderAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackendService", GoMethod: "ResetBackendService"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderAction", GoMethod: "ResetHeaderAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetWeight", GoMethod: "ResetWeight"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "weight", GoGetter: "Weight"},
			_jsii_.MemberProperty{JsiiProperty: "weightInput", GoGetter: "WeightInput"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultUrlRedirect",
		reflect.TypeFor[ComputeRegionUrlMapDefaultUrlRedirect](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultUrlRedirectOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapDefaultUrlRedirectOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "prefixRedirect", GoGetter: "PrefixRedirect"},
			_jsii_.MemberProperty{JsiiProperty: "prefixRedirectInput", GoGetter: "PrefixRedirectInput"},
			_jsii_.MemberProperty{JsiiProperty: "redirectResponseCode", GoGetter: "RedirectResponseCode"},
			_jsii_.MemberProperty{JsiiProperty: "redirectResponseCodeInput", GoGetter: "RedirectResponseCodeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostRedirect", GoMethod: "ResetHostRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpsRedirect", GoMethod: "ResetHttpsRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetPathRedirect", GoMethod: "ResetPathRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefixRedirect", GoMethod: "ResetPrefixRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedirectResponseCode", GoMethod: "ResetRedirectResponseCode"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "stripQuery", GoGetter: "StripQuery"},
			_jsii_.MemberProperty{JsiiProperty: "stripQueryInput", GoGetter: "StripQueryInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapDefaultUrlRedirectOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapHeaderAction",
		reflect.TypeFor[ComputeRegionUrlMapHeaderAction](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapHeaderActionOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapHeaderActionOutputReference](),
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
			_jsii_.MemberMethod{JsiiMethod: "putRequestHeadersToAdd", GoMethod: "PutRequestHeadersToAdd"},
			_jsii_.MemberMethod{JsiiMethod: "putResponseHeadersToAdd", GoMethod: "PutResponseHeadersToAdd"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToAdd", GoGetter: "RequestHeadersToAdd"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToAddInput", GoGetter: "RequestHeadersToAddInput"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToRemove", GoGetter: "RequestHeadersToRemove"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToRemoveInput", GoGetter: "RequestHeadersToRemoveInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestHeadersToAdd", GoMethod: "ResetRequestHeadersToAdd"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestHeadersToRemove", GoMethod: "ResetRequestHeadersToRemove"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseHeadersToAdd", GoMethod: "ResetResponseHeadersToAdd"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseHeadersToRemove", GoMethod: "ResetResponseHeadersToRemove"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToAdd", GoGetter: "ResponseHeadersToAdd"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToAddInput", GoGetter: "ResponseHeadersToAddInput"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToRemove", GoGetter: "ResponseHeadersToRemove"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToRemoveInput", GoGetter: "ResponseHeadersToRemoveInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapHeaderActionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapHeaderActionRequestHeadersToAdd",
		reflect.TypeFor[ComputeRegionUrlMapHeaderActionRequestHeadersToAdd](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapHeaderActionRequestHeadersToAddList",
		reflect.TypeFor[ComputeRegionUrlMapHeaderActionRequestHeadersToAddList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapHeaderActionRequestHeadersToAddList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapHeaderActionRequestHeadersToAddOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapHeaderActionRequestHeadersToAddOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "headerName", GoGetter: "HeaderName"},
			_jsii_.MemberProperty{JsiiProperty: "headerNameInput", GoGetter: "HeaderNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "headerValue", GoGetter: "HeaderValue"},
			_jsii_.MemberProperty{JsiiProperty: "headerValueInput", GoGetter: "HeaderValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "replace", GoGetter: "Replace"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInput", GoGetter: "ReplaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderName", GoMethod: "ResetHeaderName"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderValue", GoMethod: "ResetHeaderValue"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplace", GoMethod: "ResetReplace"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapHeaderActionRequestHeadersToAddOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapHeaderActionResponseHeadersToAdd",
		reflect.TypeFor[ComputeRegionUrlMapHeaderActionResponseHeadersToAdd](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapHeaderActionResponseHeadersToAddList",
		reflect.TypeFor[ComputeRegionUrlMapHeaderActionResponseHeadersToAddList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapHeaderActionResponseHeadersToAddList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapHeaderActionResponseHeadersToAddOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapHeaderActionResponseHeadersToAddOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "headerName", GoGetter: "HeaderName"},
			_jsii_.MemberProperty{JsiiProperty: "headerNameInput", GoGetter: "HeaderNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "headerValue", GoGetter: "HeaderValue"},
			_jsii_.MemberProperty{JsiiProperty: "headerValueInput", GoGetter: "HeaderValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "replace", GoGetter: "Replace"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInput", GoGetter: "ReplaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderName", GoMethod: "ResetHeaderName"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderValue", GoMethod: "ResetHeaderValue"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplace", GoMethod: "ResetReplace"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapHeaderActionResponseHeadersToAddOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapHostRule",
		reflect.TypeFor[ComputeRegionUrlMapHostRule](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapHostRuleList",
		reflect.TypeFor[ComputeRegionUrlMapHostRuleList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapHostRuleList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapHostRuleOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapHostRuleOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "hosts", GoGetter: "Hosts"},
			_jsii_.MemberProperty{JsiiProperty: "hostsInput", GoGetter: "HostsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "pathMatcher", GoGetter: "PathMatcher"},
			_jsii_.MemberProperty{JsiiProperty: "pathMatcherInput", GoGetter: "PathMatcherInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapHostRuleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcher",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcher](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteAction",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteAction](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicy",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicy](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicy",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicy](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyAbort",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyAbort](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyAbortOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyAbortOutputReference](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyAbortOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyDelay",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyDelay](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyDelayFixedDelay",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyDelayFixedDelay](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyDelayFixedDelayOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyDelayFixedDelayOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "nanos", GoGetter: "Nanos"},
			_jsii_.MemberProperty{JsiiProperty: "nanosInput", GoGetter: "NanosInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNanos", GoMethod: "ResetNanos"},
			_jsii_.MemberMethod{JsiiMethod: "resetSeconds", GoMethod: "ResetSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "seconds", GoGetter: "Seconds"},
			_jsii_.MemberProperty{JsiiProperty: "secondsInput", GoGetter: "SecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyDelayFixedDelayOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyDelayOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyDelayOutputReference](),
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
			_jsii_.MemberMethod{JsiiMethod: "putFixedDelay", GoMethod: "PutFixedDelay"},
			_jsii_.MemberMethod{JsiiMethod: "resetFixedDelay", GoMethod: "ResetFixedDelay"},
			_jsii_.MemberMethod{JsiiMethod: "resetPercentage", GoMethod: "ResetPercentage"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyDelayOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyOutputReference](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionMaxStreamDuration",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionMaxStreamDuration](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionMaxStreamDurationOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionMaxStreamDurationOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "nanos", GoGetter: "Nanos"},
			_jsii_.MemberProperty{JsiiProperty: "nanosInput", GoGetter: "NanosInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNanos", GoMethod: "ResetNanos"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "seconds", GoGetter: "Seconds"},
			_jsii_.MemberProperty{JsiiProperty: "secondsInput", GoGetter: "SecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionMaxStreamDurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "corsPolicy", GoGetter: "CorsPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "corsPolicyInput", GoGetter: "CorsPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "maxStreamDuration", GoGetter: "MaxStreamDuration"},
			_jsii_.MemberProperty{JsiiProperty: "maxStreamDurationInput", GoGetter: "MaxStreamDurationInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCorsPolicy", GoMethod: "PutCorsPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putFaultInjectionPolicy", GoMethod: "PutFaultInjectionPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putMaxStreamDuration", GoMethod: "PutMaxStreamDuration"},
			_jsii_.MemberMethod{JsiiMethod: "putRequestMirrorPolicy", GoMethod: "PutRequestMirrorPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putRetryPolicy", GoMethod: "PutRetryPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeout", GoMethod: "PutTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "putUrlRewrite", GoMethod: "PutUrlRewrite"},
			_jsii_.MemberMethod{JsiiMethod: "putWeightedBackendServices", GoMethod: "PutWeightedBackendServices"},
			_jsii_.MemberProperty{JsiiProperty: "requestMirrorPolicy", GoGetter: "RequestMirrorPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "requestMirrorPolicyInput", GoGetter: "RequestMirrorPolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCorsPolicy", GoMethod: "ResetCorsPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetFaultInjectionPolicy", GoMethod: "ResetFaultInjectionPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxStreamDuration", GoMethod: "ResetMaxStreamDuration"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestMirrorPolicy", GoMethod: "ResetRequestMirrorPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryPolicy", GoMethod: "ResetRetryPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeout", GoMethod: "ResetTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetUrlRewrite", GoMethod: "ResetUrlRewrite"},
			_jsii_.MemberMethod{JsiiMethod: "resetWeightedBackendServices", GoMethod: "ResetWeightedBackendServices"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retryPolicy", GoGetter: "RetryPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "retryPolicyInput", GoGetter: "RetryPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutInput", GoGetter: "TimeoutInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "urlRewrite", GoGetter: "UrlRewrite"},
			_jsii_.MemberProperty{JsiiProperty: "urlRewriteInput", GoGetter: "UrlRewriteInput"},
			_jsii_.MemberProperty{JsiiProperty: "weightedBackendServices", GoGetter: "WeightedBackendServices"},
			_jsii_.MemberProperty{JsiiProperty: "weightedBackendServicesInput", GoGetter: "WeightedBackendServicesInput"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionRequestMirrorPolicy",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionRequestMirrorPolicy](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionRequestMirrorPolicyOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionRequestMirrorPolicyOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "backendService", GoGetter: "BackendService"},
			_jsii_.MemberProperty{JsiiProperty: "backendServiceInput", GoGetter: "BackendServiceInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionRequestMirrorPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionRetryPolicy",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionRetryPolicy](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionRetryPolicyOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionRetryPolicyOutputReference](),
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
			_jsii_.MemberMethod{JsiiMethod: "putPerTryTimeout", GoMethod: "PutPerTryTimeout"},
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionRetryPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionRetryPolicyPerTryTimeout",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionRetryPolicyPerTryTimeout](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionRetryPolicyPerTryTimeoutOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionRetryPolicyPerTryTimeoutOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "nanos", GoGetter: "Nanos"},
			_jsii_.MemberProperty{JsiiProperty: "nanosInput", GoGetter: "NanosInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNanos", GoMethod: "ResetNanos"},
			_jsii_.MemberMethod{JsiiMethod: "resetSeconds", GoMethod: "ResetSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "seconds", GoGetter: "Seconds"},
			_jsii_.MemberProperty{JsiiProperty: "secondsInput", GoGetter: "SecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionRetryPolicyPerTryTimeoutOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionTimeout",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionTimeout](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionTimeoutOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionTimeoutOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "nanos", GoGetter: "Nanos"},
			_jsii_.MemberProperty{JsiiProperty: "nanosInput", GoGetter: "NanosInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNanos", GoMethod: "ResetNanos"},
			_jsii_.MemberMethod{JsiiMethod: "resetSeconds", GoMethod: "ResetSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "seconds", GoGetter: "Seconds"},
			_jsii_.MemberProperty{JsiiProperty: "secondsInput", GoGetter: "SecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionTimeoutOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionUrlRewrite",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionUrlRewrite](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionUrlRewriteOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionUrlRewriteOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "pathTemplateRewrite", GoGetter: "PathTemplateRewrite"},
			_jsii_.MemberProperty{JsiiProperty: "pathTemplateRewriteInput", GoGetter: "PathTemplateRewriteInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostRewrite", GoMethod: "ResetHostRewrite"},
			_jsii_.MemberMethod{JsiiMethod: "resetPathPrefixRewrite", GoMethod: "ResetPathPrefixRewrite"},
			_jsii_.MemberMethod{JsiiMethod: "resetPathTemplateRewrite", GoMethod: "ResetPathTemplateRewrite"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionUrlRewriteOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServices",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServices](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderAction",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderAction](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference](),
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
			_jsii_.MemberMethod{JsiiMethod: "putRequestHeadersToAdd", GoMethod: "PutRequestHeadersToAdd"},
			_jsii_.MemberMethod{JsiiMethod: "putResponseHeadersToAdd", GoMethod: "PutResponseHeadersToAdd"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToAdd", GoGetter: "RequestHeadersToAdd"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToAddInput", GoGetter: "RequestHeadersToAddInput"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToRemove", GoGetter: "RequestHeadersToRemove"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToRemoveInput", GoGetter: "RequestHeadersToRemoveInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestHeadersToAdd", GoMethod: "ResetRequestHeadersToAdd"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestHeadersToRemove", GoMethod: "ResetRequestHeadersToRemove"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseHeadersToAdd", GoMethod: "ResetResponseHeadersToAdd"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseHeadersToRemove", GoMethod: "ResetResponseHeadersToRemove"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToAdd", GoGetter: "ResponseHeadersToAdd"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToAddInput", GoGetter: "ResponseHeadersToAddInput"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToRemove", GoGetter: "ResponseHeadersToRemove"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToRemoveInput", GoGetter: "ResponseHeadersToRemoveInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAdd",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAdd](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "headerName", GoGetter: "HeaderName"},
			_jsii_.MemberProperty{JsiiProperty: "headerNameInput", GoGetter: "HeaderNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "headerValue", GoGetter: "HeaderValue"},
			_jsii_.MemberProperty{JsiiProperty: "headerValueInput", GoGetter: "HeaderValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "replace", GoGetter: "Replace"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInput", GoGetter: "ReplaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderName", GoMethod: "ResetHeaderName"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderValue", GoMethod: "ResetHeaderValue"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplace", GoMethod: "ResetReplace"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAdd",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAdd](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddList",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "headerName", GoGetter: "HeaderName"},
			_jsii_.MemberProperty{JsiiProperty: "headerNameInput", GoGetter: "HeaderNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "headerValue", GoGetter: "HeaderValue"},
			_jsii_.MemberProperty{JsiiProperty: "headerValueInput", GoGetter: "HeaderValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "replace", GoGetter: "Replace"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInput", GoGetter: "ReplaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderName", GoMethod: "ResetHeaderName"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderValue", GoMethod: "ResetHeaderValue"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplace", GoMethod: "ResetReplace"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesList",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "backendService", GoGetter: "BackendService"},
			_jsii_.MemberProperty{JsiiProperty: "backendServiceInput", GoGetter: "BackendServiceInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "headerAction", GoGetter: "HeaderAction"},
			_jsii_.MemberProperty{JsiiProperty: "headerActionInput", GoGetter: "HeaderActionInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putHeaderAction", GoMethod: "PutHeaderAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackendService", GoMethod: "ResetBackendService"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderAction", GoMethod: "ResetHeaderAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetWeight", GoMethod: "ResetWeight"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "weight", GoGetter: "Weight"},
			_jsii_.MemberProperty{JsiiProperty: "weightInput", GoGetter: "WeightInput"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultUrlRedirect",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultUrlRedirect](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultUrlRedirectOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherDefaultUrlRedirectOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "prefixRedirect", GoGetter: "PrefixRedirect"},
			_jsii_.MemberProperty{JsiiProperty: "prefixRedirectInput", GoGetter: "PrefixRedirectInput"},
			_jsii_.MemberProperty{JsiiProperty: "redirectResponseCode", GoGetter: "RedirectResponseCode"},
			_jsii_.MemberProperty{JsiiProperty: "redirectResponseCodeInput", GoGetter: "RedirectResponseCodeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostRedirect", GoMethod: "ResetHostRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpsRedirect", GoMethod: "ResetHttpsRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetPathRedirect", GoMethod: "ResetPathRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefixRedirect", GoMethod: "ResetPrefixRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedirectResponseCode", GoMethod: "ResetRedirectResponseCode"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "stripQuery", GoGetter: "StripQuery"},
			_jsii_.MemberProperty{JsiiProperty: "stripQueryInput", GoGetter: "StripQueryInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultUrlRedirectOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherHeaderAction",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherHeaderAction](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherHeaderActionOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherHeaderActionOutputReference](),
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
			_jsii_.MemberMethod{JsiiMethod: "putRequestHeadersToAdd", GoMethod: "PutRequestHeadersToAdd"},
			_jsii_.MemberMethod{JsiiMethod: "putResponseHeadersToAdd", GoMethod: "PutResponseHeadersToAdd"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToAdd", GoGetter: "RequestHeadersToAdd"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToAddInput", GoGetter: "RequestHeadersToAddInput"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToRemove", GoGetter: "RequestHeadersToRemove"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToRemoveInput", GoGetter: "RequestHeadersToRemoveInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestHeadersToAdd", GoMethod: "ResetRequestHeadersToAdd"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestHeadersToRemove", GoMethod: "ResetRequestHeadersToRemove"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseHeadersToAdd", GoMethod: "ResetResponseHeadersToAdd"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseHeadersToRemove", GoMethod: "ResetResponseHeadersToRemove"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToAdd", GoGetter: "ResponseHeadersToAdd"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToAddInput", GoGetter: "ResponseHeadersToAddInput"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToRemove", GoGetter: "ResponseHeadersToRemove"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToRemoveInput", GoGetter: "ResponseHeadersToRemoveInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherHeaderActionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherHeaderActionRequestHeadersToAdd",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherHeaderActionRequestHeadersToAdd](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherHeaderActionRequestHeadersToAddList",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherHeaderActionRequestHeadersToAddList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherHeaderActionRequestHeadersToAddList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherHeaderActionRequestHeadersToAddOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherHeaderActionRequestHeadersToAddOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "headerName", GoGetter: "HeaderName"},
			_jsii_.MemberProperty{JsiiProperty: "headerNameInput", GoGetter: "HeaderNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "headerValue", GoGetter: "HeaderValue"},
			_jsii_.MemberProperty{JsiiProperty: "headerValueInput", GoGetter: "HeaderValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "replace", GoGetter: "Replace"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInput", GoGetter: "ReplaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderName", GoMethod: "ResetHeaderName"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderValue", GoMethod: "ResetHeaderValue"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplace", GoMethod: "ResetReplace"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherHeaderActionRequestHeadersToAddOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherHeaderActionResponseHeadersToAdd",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherHeaderActionResponseHeadersToAdd](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherHeaderActionResponseHeadersToAddList",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherHeaderActionResponseHeadersToAddList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherHeaderActionResponseHeadersToAddList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherHeaderActionResponseHeadersToAddOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherHeaderActionResponseHeadersToAddOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "headerName", GoGetter: "HeaderName"},
			_jsii_.MemberProperty{JsiiProperty: "headerNameInput", GoGetter: "HeaderNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "headerValue", GoGetter: "HeaderValue"},
			_jsii_.MemberProperty{JsiiProperty: "headerValueInput", GoGetter: "HeaderValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "replace", GoGetter: "Replace"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInput", GoGetter: "ReplaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderName", GoMethod: "ResetHeaderName"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderValue", GoMethod: "ResetHeaderValue"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplace", GoMethod: "ResetReplace"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherHeaderActionResponseHeadersToAddOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherList",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRouteAction", GoGetter: "DefaultRouteAction"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRouteActionInput", GoGetter: "DefaultRouteActionInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultService", GoGetter: "DefaultService"},
			_jsii_.MemberProperty{JsiiProperty: "defaultServiceInput", GoGetter: "DefaultServiceInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultUrlRedirect", GoGetter: "DefaultUrlRedirect"},
			_jsii_.MemberProperty{JsiiProperty: "defaultUrlRedirectInput", GoGetter: "DefaultUrlRedirectInput"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "headerAction", GoGetter: "HeaderAction"},
			_jsii_.MemberProperty{JsiiProperty: "headerActionInput", GoGetter: "HeaderActionInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "pathRule", GoGetter: "PathRule"},
			_jsii_.MemberProperty{JsiiProperty: "pathRuleInput", GoGetter: "PathRuleInput"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultRouteAction", GoMethod: "PutDefaultRouteAction"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultUrlRedirect", GoMethod: "PutDefaultUrlRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "putHeaderAction", GoMethod: "PutHeaderAction"},
			_jsii_.MemberMethod{JsiiMethod: "putPathRule", GoMethod: "PutPathRule"},
			_jsii_.MemberMethod{JsiiMethod: "putRouteRules", GoMethod: "PutRouteRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultRouteAction", GoMethod: "ResetDefaultRouteAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultService", GoMethod: "ResetDefaultService"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultUrlRedirect", GoMethod: "ResetDefaultUrlRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderAction", GoMethod: "ResetHeaderAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetPathRule", GoMethod: "ResetPathRule"},
			_jsii_.MemberMethod{JsiiMethod: "resetRouteRules", GoMethod: "ResetRouteRules"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "routeRules", GoGetter: "RouteRules"},
			_jsii_.MemberProperty{JsiiProperty: "routeRulesInput", GoGetter: "RouteRulesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRule",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRule](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleList",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "paths", GoGetter: "Paths"},
			_jsii_.MemberProperty{JsiiProperty: "pathsInput", GoGetter: "PathsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putRouteAction", GoMethod: "PutRouteAction"},
			_jsii_.MemberMethod{JsiiMethod: "putUrlRedirect", GoMethod: "PutUrlRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetRouteAction", GoMethod: "ResetRouteAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetService", GoMethod: "ResetService"},
			_jsii_.MemberMethod{JsiiMethod: "resetUrlRedirect", GoMethod: "ResetUrlRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "routeAction", GoGetter: "RouteAction"},
			_jsii_.MemberProperty{JsiiProperty: "routeActionInput", GoGetter: "RouteActionInput"},
			_jsii_.MemberProperty{JsiiProperty: "service", GoGetter: "Service"},
			_jsii_.MemberProperty{JsiiProperty: "serviceInput", GoGetter: "ServiceInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "urlRedirect", GoGetter: "UrlRedirect"},
			_jsii_.MemberProperty{JsiiProperty: "urlRedirectInput", GoGetter: "UrlRedirectInput"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteAction",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteAction](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutputReference](),
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
			_jsii_.MemberMethod{JsiiMethod: "resetExposeHeaders", GoMethod: "ResetExposeHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxAge", GoMethod: "ResetMaxAge"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortOutputReference](),
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelayOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelayOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "nanos", GoGetter: "Nanos"},
			_jsii_.MemberProperty{JsiiProperty: "nanosInput", GoGetter: "NanosInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNanos", GoMethod: "ResetNanos"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "seconds", GoGetter: "Seconds"},
			_jsii_.MemberProperty{JsiiProperty: "secondsInput", GoGetter: "SecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelayOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayOutputReference](),
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
			_jsii_.MemberMethod{JsiiMethod: "putFixedDelay", GoMethod: "PutFixedDelay"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyOutputReference](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "corsPolicy", GoGetter: "CorsPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "corsPolicyInput", GoGetter: "CorsPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberMethod{JsiiMethod: "putFaultInjectionPolicy", GoMethod: "PutFaultInjectionPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putRequestMirrorPolicy", GoMethod: "PutRequestMirrorPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putRetryPolicy", GoMethod: "PutRetryPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeout", GoMethod: "PutTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "putUrlRewrite", GoMethod: "PutUrlRewrite"},
			_jsii_.MemberMethod{JsiiMethod: "putWeightedBackendServices", GoMethod: "PutWeightedBackendServices"},
			_jsii_.MemberProperty{JsiiProperty: "requestMirrorPolicy", GoGetter: "RequestMirrorPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "requestMirrorPolicyInput", GoGetter: "RequestMirrorPolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCorsPolicy", GoMethod: "ResetCorsPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetFaultInjectionPolicy", GoMethod: "ResetFaultInjectionPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestMirrorPolicy", GoMethod: "ResetRequestMirrorPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryPolicy", GoMethod: "ResetRetryPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeout", GoMethod: "ResetTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetUrlRewrite", GoMethod: "ResetUrlRewrite"},
			_jsii_.MemberMethod{JsiiMethod: "resetWeightedBackendServices", GoMethod: "ResetWeightedBackendServices"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retryPolicy", GoGetter: "RetryPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "retryPolicyInput", GoGetter: "RetryPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutInput", GoGetter: "TimeoutInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "urlRewrite", GoGetter: "UrlRewrite"},
			_jsii_.MemberProperty{JsiiProperty: "urlRewriteInput", GoGetter: "UrlRewriteInput"},
			_jsii_.MemberProperty{JsiiProperty: "weightedBackendServices", GoGetter: "WeightedBackendServices"},
			_jsii_.MemberProperty{JsiiProperty: "weightedBackendServicesInput", GoGetter: "WeightedBackendServicesInput"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleRouteActionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "backendService", GoGetter: "BackendService"},
			_jsii_.MemberProperty{JsiiProperty: "backendServiceInput", GoGetter: "BackendServiceInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionRetryPolicy",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionRetryPolicy](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionRetryPolicyOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionRetryPolicyOutputReference](),
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
			_jsii_.MemberMethod{JsiiMethod: "putPerTryTimeout", GoMethod: "PutPerTryTimeout"},
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleRouteActionRetryPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "nanos", GoGetter: "Nanos"},
			_jsii_.MemberProperty{JsiiProperty: "nanosInput", GoGetter: "NanosInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNanos", GoMethod: "ResetNanos"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "seconds", GoGetter: "Seconds"},
			_jsii_.MemberProperty{JsiiProperty: "secondsInput", GoGetter: "SecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionTimeout",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionTimeout](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionTimeoutOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionTimeoutOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "nanos", GoGetter: "Nanos"},
			_jsii_.MemberProperty{JsiiProperty: "nanosInput", GoGetter: "NanosInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNanos", GoMethod: "ResetNanos"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "seconds", GoGetter: "Seconds"},
			_jsii_.MemberProperty{JsiiProperty: "secondsInput", GoGetter: "SecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleRouteActionTimeoutOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionUrlRewrite",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionUrlRewrite](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionUrlRewriteOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionUrlRewriteOutputReference](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleRouteActionUrlRewriteOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServices",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServices](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderAction",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderAction](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderActionOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderActionOutputReference](),
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
			_jsii_.MemberMethod{JsiiMethod: "putRequestHeadersToAdd", GoMethod: "PutRequestHeadersToAdd"},
			_jsii_.MemberMethod{JsiiMethod: "putResponseHeadersToAdd", GoMethod: "PutResponseHeadersToAdd"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToAdd", GoGetter: "RequestHeadersToAdd"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToAddInput", GoGetter: "RequestHeadersToAddInput"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToRemove", GoGetter: "RequestHeadersToRemove"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToRemoveInput", GoGetter: "RequestHeadersToRemoveInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestHeadersToAdd", GoMethod: "ResetRequestHeadersToAdd"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestHeadersToRemove", GoMethod: "ResetRequestHeadersToRemove"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseHeadersToAdd", GoMethod: "ResetResponseHeadersToAdd"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseHeadersToRemove", GoMethod: "ResetResponseHeadersToRemove"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToAdd", GoGetter: "ResponseHeadersToAdd"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToAddInput", GoGetter: "ResponseHeadersToAddInput"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToRemove", GoGetter: "ResponseHeadersToRemove"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToRemoveInput", GoGetter: "ResponseHeadersToRemoveInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderActionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAdd",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAdd](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "headerName", GoGetter: "HeaderName"},
			_jsii_.MemberProperty{JsiiProperty: "headerNameInput", GoGetter: "HeaderNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "headerValue", GoGetter: "HeaderValue"},
			_jsii_.MemberProperty{JsiiProperty: "headerValueInput", GoGetter: "HeaderValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "replace", GoGetter: "Replace"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInput", GoGetter: "ReplaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAdd",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAdd](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddList",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "headerName", GoGetter: "HeaderName"},
			_jsii_.MemberProperty{JsiiProperty: "headerNameInput", GoGetter: "HeaderNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "headerValue", GoGetter: "HeaderValue"},
			_jsii_.MemberProperty{JsiiProperty: "headerValueInput", GoGetter: "HeaderValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "replace", GoGetter: "Replace"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInput", GoGetter: "ReplaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesList",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "backendService", GoGetter: "BackendService"},
			_jsii_.MemberProperty{JsiiProperty: "backendServiceInput", GoGetter: "BackendServiceInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "headerAction", GoGetter: "HeaderAction"},
			_jsii_.MemberProperty{JsiiProperty: "headerActionInput", GoGetter: "HeaderActionInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putHeaderAction", GoMethod: "PutHeaderAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderAction", GoMethod: "ResetHeaderAction"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "weight", GoGetter: "Weight"},
			_jsii_.MemberProperty{JsiiProperty: "weightInput", GoGetter: "WeightInput"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleUrlRedirect",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleUrlRedirect](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleUrlRedirectOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherPathRuleUrlRedirectOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "prefixRedirect", GoGetter: "PrefixRedirect"},
			_jsii_.MemberProperty{JsiiProperty: "prefixRedirectInput", GoGetter: "PrefixRedirectInput"},
			_jsii_.MemberProperty{JsiiProperty: "redirectResponseCode", GoGetter: "RedirectResponseCode"},
			_jsii_.MemberProperty{JsiiProperty: "redirectResponseCodeInput", GoGetter: "RedirectResponseCodeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostRedirect", GoMethod: "ResetHostRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpsRedirect", GoMethod: "ResetHttpsRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetPathRedirect", GoMethod: "ResetPathRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefixRedirect", GoMethod: "ResetPrefixRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedirectResponseCode", GoMethod: "ResetRedirectResponseCode"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "stripQuery", GoGetter: "StripQuery"},
			_jsii_.MemberProperty{JsiiProperty: "stripQueryInput", GoGetter: "StripQueryInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleUrlRedirectOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRules",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRules](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesHeaderAction",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesHeaderAction](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesHeaderActionOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesHeaderActionOutputReference](),
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
			_jsii_.MemberMethod{JsiiMethod: "putRequestHeadersToAdd", GoMethod: "PutRequestHeadersToAdd"},
			_jsii_.MemberMethod{JsiiMethod: "putResponseHeadersToAdd", GoMethod: "PutResponseHeadersToAdd"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToAdd", GoGetter: "RequestHeadersToAdd"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToAddInput", GoGetter: "RequestHeadersToAddInput"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToRemove", GoGetter: "RequestHeadersToRemove"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToRemoveInput", GoGetter: "RequestHeadersToRemoveInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestHeadersToAdd", GoMethod: "ResetRequestHeadersToAdd"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestHeadersToRemove", GoMethod: "ResetRequestHeadersToRemove"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseHeadersToAdd", GoMethod: "ResetResponseHeadersToAdd"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseHeadersToRemove", GoMethod: "ResetResponseHeadersToRemove"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToAdd", GoGetter: "ResponseHeadersToAdd"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToAddInput", GoGetter: "ResponseHeadersToAddInput"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToRemove", GoGetter: "ResponseHeadersToRemove"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToRemoveInput", GoGetter: "ResponseHeadersToRemoveInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesHeaderActionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesHeaderActionRequestHeadersToAdd",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesHeaderActionRequestHeadersToAdd](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesHeaderActionRequestHeadersToAddList",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesHeaderActionRequestHeadersToAddList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesHeaderActionRequestHeadersToAddList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesHeaderActionRequestHeadersToAddOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesHeaderActionRequestHeadersToAddOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "headerName", GoGetter: "HeaderName"},
			_jsii_.MemberProperty{JsiiProperty: "headerNameInput", GoGetter: "HeaderNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "headerValue", GoGetter: "HeaderValue"},
			_jsii_.MemberProperty{JsiiProperty: "headerValueInput", GoGetter: "HeaderValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "replace", GoGetter: "Replace"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInput", GoGetter: "ReplaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesHeaderActionRequestHeadersToAddOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesHeaderActionResponseHeadersToAdd",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesHeaderActionResponseHeadersToAdd](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesHeaderActionResponseHeadersToAddList",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesHeaderActionResponseHeadersToAddList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesHeaderActionResponseHeadersToAddList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesHeaderActionResponseHeadersToAddOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesHeaderActionResponseHeadersToAddOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "headerName", GoGetter: "HeaderName"},
			_jsii_.MemberProperty{JsiiProperty: "headerNameInput", GoGetter: "HeaderNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "headerValue", GoGetter: "HeaderValue"},
			_jsii_.MemberProperty{JsiiProperty: "headerValueInput", GoGetter: "HeaderValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "replace", GoGetter: "Replace"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInput", GoGetter: "ReplaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesHeaderActionResponseHeadersToAddOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesList",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesMatchRules",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesMatchRules](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatches",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatches](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesList",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "headerName", GoGetter: "HeaderName"},
			_jsii_.MemberProperty{JsiiProperty: "headerNameInput", GoGetter: "HeaderNameInput"},
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatch",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatch](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "rangeEnd", GoGetter: "RangeEnd"},
			_jsii_.MemberProperty{JsiiProperty: "rangeEndInput", GoGetter: "RangeEndInput"},
			_jsii_.MemberProperty{JsiiProperty: "rangeStart", GoGetter: "RangeStart"},
			_jsii_.MemberProperty{JsiiProperty: "rangeStartInput", GoGetter: "RangeStartInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesList",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesMetadataFilters",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesMetadataFilters](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesMetadataFiltersFilterLabels",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesMetadataFiltersFilterLabels](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesMetadataFiltersFilterLabelsList",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesMetadataFiltersFilterLabelsList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesMetadataFiltersFilterLabelsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesMetadataFiltersFilterLabelsOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesMetadataFiltersFilterLabelsOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesMetadataFiltersFilterLabelsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesMetadataFiltersList",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesMetadataFiltersList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesMetadataFiltersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesMetadataFiltersOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesMetadataFiltersOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "filterLabels", GoGetter: "FilterLabels"},
			_jsii_.MemberProperty{JsiiProperty: "filterLabelsInput", GoGetter: "FilterLabelsInput"},
			_jsii_.MemberProperty{JsiiProperty: "filterMatchCriteria", GoGetter: "FilterMatchCriteria"},
			_jsii_.MemberProperty{JsiiProperty: "filterMatchCriteriaInput", GoGetter: "FilterMatchCriteriaInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putFilterLabels", GoMethod: "PutFilterLabels"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesMetadataFiltersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "headerMatches", GoGetter: "HeaderMatches"},
			_jsii_.MemberProperty{JsiiProperty: "headerMatchesInput", GoGetter: "HeaderMatchesInput"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreCase", GoGetter: "IgnoreCase"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreCaseInput", GoGetter: "IgnoreCaseInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "metadataFilters", GoGetter: "MetadataFilters"},
			_jsii_.MemberProperty{JsiiProperty: "metadataFiltersInput", GoGetter: "MetadataFiltersInput"},
			_jsii_.MemberProperty{JsiiProperty: "pathTemplateMatch", GoGetter: "PathTemplateMatch"},
			_jsii_.MemberProperty{JsiiProperty: "pathTemplateMatchInput", GoGetter: "PathTemplateMatchInput"},
			_jsii_.MemberProperty{JsiiProperty: "prefixMatch", GoGetter: "PrefixMatch"},
			_jsii_.MemberProperty{JsiiProperty: "prefixMatchInput", GoGetter: "PrefixMatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "putHeaderMatches", GoMethod: "PutHeaderMatches"},
			_jsii_.MemberMethod{JsiiMethod: "putMetadataFilters", GoMethod: "PutMetadataFilters"},
			_jsii_.MemberMethod{JsiiMethod: "putQueryParameterMatches", GoMethod: "PutQueryParameterMatches"},
			_jsii_.MemberProperty{JsiiProperty: "queryParameterMatches", GoGetter: "QueryParameterMatches"},
			_jsii_.MemberProperty{JsiiProperty: "queryParameterMatchesInput", GoGetter: "QueryParameterMatchesInput"},
			_jsii_.MemberProperty{JsiiProperty: "regexMatch", GoGetter: "RegexMatch"},
			_jsii_.MemberProperty{JsiiProperty: "regexMatchInput", GoGetter: "RegexMatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetFullPathMatch", GoMethod: "ResetFullPathMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderMatches", GoMethod: "ResetHeaderMatches"},
			_jsii_.MemberMethod{JsiiMethod: "resetIgnoreCase", GoMethod: "ResetIgnoreCase"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadataFilters", GoMethod: "ResetMetadataFilters"},
			_jsii_.MemberMethod{JsiiMethod: "resetPathTemplateMatch", GoMethod: "ResetPathTemplateMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefixMatch", GoMethod: "ResetPrefixMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryParameterMatches", GoMethod: "ResetQueryParameterMatches"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegexMatch", GoMethod: "ResetRegexMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesQueryParameterMatches",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesQueryParameterMatches](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesQueryParameterMatchesList",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesQueryParameterMatchesList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesQueryParameterMatchesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesQueryParameterMatchesOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesQueryParameterMatchesOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "presentMatch", GoGetter: "PresentMatch"},
			_jsii_.MemberProperty{JsiiProperty: "presentMatchInput", GoGetter: "PresentMatchInput"},
			_jsii_.MemberProperty{JsiiProperty: "regexMatch", GoGetter: "RegexMatch"},
			_jsii_.MemberProperty{JsiiProperty: "regexMatchInput", GoGetter: "RegexMatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetExactMatch", GoMethod: "ResetExactMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetPresentMatch", GoMethod: "ResetPresentMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegexMatch", GoMethod: "ResetRegexMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesQueryParameterMatchesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "headerAction", GoGetter: "HeaderAction"},
			_jsii_.MemberProperty{JsiiProperty: "headerActionInput", GoGetter: "HeaderActionInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "matchRules", GoGetter: "MatchRules"},
			_jsii_.MemberProperty{JsiiProperty: "matchRulesInput", GoGetter: "MatchRulesInput"},
			_jsii_.MemberProperty{JsiiProperty: "priority", GoGetter: "Priority"},
			_jsii_.MemberProperty{JsiiProperty: "priorityInput", GoGetter: "PriorityInput"},
			_jsii_.MemberMethod{JsiiMethod: "putHeaderAction", GoMethod: "PutHeaderAction"},
			_jsii_.MemberMethod{JsiiMethod: "putMatchRules", GoMethod: "PutMatchRules"},
			_jsii_.MemberMethod{JsiiMethod: "putRouteAction", GoMethod: "PutRouteAction"},
			_jsii_.MemberMethod{JsiiMethod: "putUrlRedirect", GoMethod: "PutUrlRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderAction", GoMethod: "ResetHeaderAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetMatchRules", GoMethod: "ResetMatchRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetRouteAction", GoMethod: "ResetRouteAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetService", GoMethod: "ResetService"},
			_jsii_.MemberMethod{JsiiMethod: "resetUrlRedirect", GoMethod: "ResetUrlRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "routeAction", GoGetter: "RouteAction"},
			_jsii_.MemberProperty{JsiiProperty: "routeActionInput", GoGetter: "RouteActionInput"},
			_jsii_.MemberProperty{JsiiProperty: "service", GoGetter: "Service"},
			_jsii_.MemberProperty{JsiiProperty: "serviceInput", GoGetter: "ServiceInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "urlRedirect", GoGetter: "UrlRedirect"},
			_jsii_.MemberProperty{JsiiProperty: "urlRedirectInput", GoGetter: "UrlRedirectInput"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteAction",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteAction](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionCorsPolicy",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionCorsPolicy](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionCorsPolicyOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionCorsPolicyOutputReference](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesRouteActionCorsPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicy",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicy](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyAbort",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyAbort](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyAbortOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyAbortOutputReference](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyAbortOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyDelay",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyDelay](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyDelayFixedDelay",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyDelayFixedDelay](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyDelayFixedDelayOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyDelayFixedDelayOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "nanos", GoGetter: "Nanos"},
			_jsii_.MemberProperty{JsiiProperty: "nanosInput", GoGetter: "NanosInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNanos", GoMethod: "ResetNanos"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "seconds", GoGetter: "Seconds"},
			_jsii_.MemberProperty{JsiiProperty: "secondsInput", GoGetter: "SecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyDelayFixedDelayOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyDelayOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyDelayOutputReference](),
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
			_jsii_.MemberMethod{JsiiMethod: "putFixedDelay", GoMethod: "PutFixedDelay"},
			_jsii_.MemberMethod{JsiiMethod: "resetFixedDelay", GoMethod: "ResetFixedDelay"},
			_jsii_.MemberMethod{JsiiMethod: "resetPercentage", GoMethod: "ResetPercentage"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyDelayOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyOutputReference](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "corsPolicy", GoGetter: "CorsPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "corsPolicyInput", GoGetter: "CorsPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberMethod{JsiiMethod: "putFaultInjectionPolicy", GoMethod: "PutFaultInjectionPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putRequestMirrorPolicy", GoMethod: "PutRequestMirrorPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putRetryPolicy", GoMethod: "PutRetryPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeout", GoMethod: "PutTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "putUrlRewrite", GoMethod: "PutUrlRewrite"},
			_jsii_.MemberMethod{JsiiMethod: "putWeightedBackendServices", GoMethod: "PutWeightedBackendServices"},
			_jsii_.MemberProperty{JsiiProperty: "requestMirrorPolicy", GoGetter: "RequestMirrorPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "requestMirrorPolicyInput", GoGetter: "RequestMirrorPolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCorsPolicy", GoMethod: "ResetCorsPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetFaultInjectionPolicy", GoMethod: "ResetFaultInjectionPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestMirrorPolicy", GoMethod: "ResetRequestMirrorPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryPolicy", GoMethod: "ResetRetryPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeout", GoMethod: "ResetTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetUrlRewrite", GoMethod: "ResetUrlRewrite"},
			_jsii_.MemberMethod{JsiiMethod: "resetWeightedBackendServices", GoMethod: "ResetWeightedBackendServices"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retryPolicy", GoGetter: "RetryPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "retryPolicyInput", GoGetter: "RetryPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutInput", GoGetter: "TimeoutInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "urlRewrite", GoGetter: "UrlRewrite"},
			_jsii_.MemberProperty{JsiiProperty: "urlRewriteInput", GoGetter: "UrlRewriteInput"},
			_jsii_.MemberProperty{JsiiProperty: "weightedBackendServices", GoGetter: "WeightedBackendServices"},
			_jsii_.MemberProperty{JsiiProperty: "weightedBackendServicesInput", GoGetter: "WeightedBackendServicesInput"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicy",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicy](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicyOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicyOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "backendService", GoGetter: "BackendService"},
			_jsii_.MemberProperty{JsiiProperty: "backendServiceInput", GoGetter: "BackendServiceInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionRetryPolicy",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionRetryPolicy](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionRetryPolicyOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionRetryPolicyOutputReference](),
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
			_jsii_.MemberMethod{JsiiMethod: "putPerTryTimeout", GoMethod: "PutPerTryTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetPerTryTimeout", GoMethod: "ResetPerTryTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryConditions", GoMethod: "ResetRetryConditions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retryConditions", GoGetter: "RetryConditions"},
			_jsii_.MemberProperty{JsiiProperty: "retryConditionsInput", GoGetter: "RetryConditionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesRouteActionRetryPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionRetryPolicyPerTryTimeout",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionRetryPolicyPerTryTimeout](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionRetryPolicyPerTryTimeoutOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionRetryPolicyPerTryTimeoutOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "nanos", GoGetter: "Nanos"},
			_jsii_.MemberProperty{JsiiProperty: "nanosInput", GoGetter: "NanosInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNanos", GoMethod: "ResetNanos"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "seconds", GoGetter: "Seconds"},
			_jsii_.MemberProperty{JsiiProperty: "secondsInput", GoGetter: "SecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesRouteActionRetryPolicyPerTryTimeoutOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionTimeout",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionTimeout](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionTimeoutOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionTimeoutOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "nanos", GoGetter: "Nanos"},
			_jsii_.MemberProperty{JsiiProperty: "nanosInput", GoGetter: "NanosInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNanos", GoMethod: "ResetNanos"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "seconds", GoGetter: "Seconds"},
			_jsii_.MemberProperty{JsiiProperty: "secondsInput", GoGetter: "SecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesRouteActionTimeoutOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionUrlRewrite",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionUrlRewrite](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionUrlRewriteOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionUrlRewriteOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "pathTemplateRewrite", GoGetter: "PathTemplateRewrite"},
			_jsii_.MemberProperty{JsiiProperty: "pathTemplateRewriteInput", GoGetter: "PathTemplateRewriteInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostRewrite", GoMethod: "ResetHostRewrite"},
			_jsii_.MemberMethod{JsiiMethod: "resetPathPrefixRewrite", GoMethod: "ResetPathPrefixRewrite"},
			_jsii_.MemberMethod{JsiiMethod: "resetPathTemplateRewrite", GoMethod: "ResetPathTemplateRewrite"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesRouteActionUrlRewriteOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServices",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServices](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesHeaderAction",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesHeaderAction](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesHeaderActionOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesHeaderActionOutputReference](),
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
			_jsii_.MemberMethod{JsiiMethod: "putRequestHeadersToAdd", GoMethod: "PutRequestHeadersToAdd"},
			_jsii_.MemberMethod{JsiiMethod: "putResponseHeadersToAdd", GoMethod: "PutResponseHeadersToAdd"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToAdd", GoGetter: "RequestHeadersToAdd"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToAddInput", GoGetter: "RequestHeadersToAddInput"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToRemove", GoGetter: "RequestHeadersToRemove"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersToRemoveInput", GoGetter: "RequestHeadersToRemoveInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestHeadersToAdd", GoMethod: "ResetRequestHeadersToAdd"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestHeadersToRemove", GoMethod: "ResetRequestHeadersToRemove"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseHeadersToAdd", GoMethod: "ResetResponseHeadersToAdd"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseHeadersToRemove", GoMethod: "ResetResponseHeadersToRemove"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToAdd", GoGetter: "ResponseHeadersToAdd"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToAddInput", GoGetter: "ResponseHeadersToAddInput"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToRemove", GoGetter: "ResponseHeadersToRemove"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersToRemoveInput", GoGetter: "ResponseHeadersToRemoveInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesHeaderActionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAdd",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAdd](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "headerName", GoGetter: "HeaderName"},
			_jsii_.MemberProperty{JsiiProperty: "headerNameInput", GoGetter: "HeaderNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "headerValue", GoGetter: "HeaderValue"},
			_jsii_.MemberProperty{JsiiProperty: "headerValueInput", GoGetter: "HeaderValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "replace", GoGetter: "Replace"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInput", GoGetter: "ReplaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAdd",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAdd](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddList",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "headerName", GoGetter: "HeaderName"},
			_jsii_.MemberProperty{JsiiProperty: "headerNameInput", GoGetter: "HeaderNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "headerValue", GoGetter: "HeaderValue"},
			_jsii_.MemberProperty{JsiiProperty: "headerValueInput", GoGetter: "HeaderValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "replace", GoGetter: "Replace"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInput", GoGetter: "ReplaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesList",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "backendService", GoGetter: "BackendService"},
			_jsii_.MemberProperty{JsiiProperty: "backendServiceInput", GoGetter: "BackendServiceInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "headerAction", GoGetter: "HeaderAction"},
			_jsii_.MemberProperty{JsiiProperty: "headerActionInput", GoGetter: "HeaderActionInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putHeaderAction", GoMethod: "PutHeaderAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderAction", GoMethod: "ResetHeaderAction"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "weight", GoGetter: "Weight"},
			_jsii_.MemberProperty{JsiiProperty: "weightInput", GoGetter: "WeightInput"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesUrlRedirect",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesUrlRedirect](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesUrlRedirectOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapPathMatcherRouteRulesUrlRedirectOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "prefixRedirect", GoGetter: "PrefixRedirect"},
			_jsii_.MemberProperty{JsiiProperty: "prefixRedirectInput", GoGetter: "PrefixRedirectInput"},
			_jsii_.MemberProperty{JsiiProperty: "redirectResponseCode", GoGetter: "RedirectResponseCode"},
			_jsii_.MemberProperty{JsiiProperty: "redirectResponseCodeInput", GoGetter: "RedirectResponseCodeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostRedirect", GoMethod: "ResetHostRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpsRedirect", GoMethod: "ResetHttpsRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetPathRedirect", GoMethod: "ResetPathRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefixRedirect", GoMethod: "ResetPrefixRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedirectResponseCode", GoMethod: "ResetRedirectResponseCode"},
			_jsii_.MemberMethod{JsiiMethod: "resetStripQuery", GoMethod: "ResetStripQuery"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "stripQuery", GoGetter: "StripQuery"},
			_jsii_.MemberProperty{JsiiProperty: "stripQueryInput", GoGetter: "StripQueryInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesUrlRedirectOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapTest",
		reflect.TypeFor[ComputeRegionUrlMapTest](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapTestList",
		reflect.TypeFor[ComputeRegionUrlMapTestList](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapTestList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapTestOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapTestOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "host", GoGetter: "Host"},
			_jsii_.MemberProperty{JsiiProperty: "hostInput", GoGetter: "HostInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathInput", GoGetter: "PathInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "service", GoGetter: "Service"},
			_jsii_.MemberProperty{JsiiProperty: "serviceInput", GoGetter: "ServiceInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapTestOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapTimeouts",
		reflect.TypeFor[ComputeRegionUrlMapTimeouts](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapTimeoutsOutputReference",
		reflect.TypeFor[ComputeRegionUrlMapTimeoutsOutputReference](),
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
		func() any {
			j := jsiiProxy_ComputeRegionUrlMapTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}

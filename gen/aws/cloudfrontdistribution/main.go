package cloudfrontdistribution

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistribution",
		reflect.TypeFor[CloudfrontDistribution](),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "aliases", GoGetter: "Aliases"},
			_jsii_.MemberProperty{JsiiProperty: "aliasesInput", GoGetter: "AliasesInput"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "callerReference", GoGetter: "CallerReference"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "commentInput", GoGetter: "CommentInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customErrorResponse", GoGetter: "CustomErrorResponse"},
			_jsii_.MemberProperty{JsiiProperty: "customErrorResponseInput", GoGetter: "CustomErrorResponseInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultCacheBehavior", GoGetter: "DefaultCacheBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "defaultCacheBehaviorInput", GoGetter: "DefaultCacheBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRootObject", GoGetter: "DefaultRootObject"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRootObjectInput", GoGetter: "DefaultRootObjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "etag", GoGetter: "Etag"},
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
			_jsii_.MemberProperty{JsiiProperty: "hostedZoneId", GoGetter: "HostedZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "httpVersion", GoGetter: "HttpVersion"},
			_jsii_.MemberProperty{JsiiProperty: "httpVersionInput", GoGetter: "HttpVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberProperty{JsiiProperty: "inProgressValidationBatches", GoGetter: "InProgressValidationBatches"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isIpv6Enabled", GoGetter: "IsIpv6Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "isIpv6EnabledInput", GoGetter: "IsIpv6EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "lastModifiedTime", GoGetter: "LastModifiedTime"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfig", GoGetter: "LoggingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigInput", GoGetter: "LoggingConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "orderedCacheBehavior", GoGetter: "OrderedCacheBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "orderedCacheBehaviorInput", GoGetter: "OrderedCacheBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "origin", GoGetter: "Origin"},
			_jsii_.MemberProperty{JsiiProperty: "originGroup", GoGetter: "OriginGroup"},
			_jsii_.MemberProperty{JsiiProperty: "originGroupInput", GoGetter: "OriginGroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "originInput", GoGetter: "OriginInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "priceClass", GoGetter: "PriceClass"},
			_jsii_.MemberProperty{JsiiProperty: "priceClassInput", GoGetter: "PriceClassInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomErrorResponse", GoMethod: "PutCustomErrorResponse"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultCacheBehavior", GoMethod: "PutDefaultCacheBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "putLoggingConfig", GoMethod: "PutLoggingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putOrderedCacheBehavior", GoMethod: "PutOrderedCacheBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "putOrigin", GoMethod: "PutOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "putOriginGroup", GoMethod: "PutOriginGroup"},
			_jsii_.MemberMethod{JsiiMethod: "putRestrictions", GoMethod: "PutRestrictions"},
			_jsii_.MemberMethod{JsiiMethod: "putViewerCertificate", GoMethod: "PutViewerCertificate"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAliases", GoMethod: "ResetAliases"},
			_jsii_.MemberMethod{JsiiMethod: "resetComment", GoMethod: "ResetComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomErrorResponse", GoMethod: "ResetCustomErrorResponse"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultRootObject", GoMethod: "ResetDefaultRootObject"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpVersion", GoMethod: "ResetHttpVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsIpv6Enabled", GoMethod: "ResetIsIpv6Enabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoggingConfig", GoMethod: "ResetLoggingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrderedCacheBehavior", GoMethod: "ResetOrderedCacheBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginGroup", GoMethod: "ResetOriginGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPriceClass", GoMethod: "ResetPriceClass"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetainOnDelete", GoMethod: "ResetRetainOnDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetWaitForDeployment", GoMethod: "ResetWaitForDeployment"},
			_jsii_.MemberMethod{JsiiMethod: "resetWebAclId", GoMethod: "ResetWebAclId"},
			_jsii_.MemberProperty{JsiiProperty: "restrictions", GoGetter: "Restrictions"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionsInput", GoGetter: "RestrictionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "retainOnDelete", GoGetter: "RetainOnDelete"},
			_jsii_.MemberProperty{JsiiProperty: "retainOnDeleteInput", GoGetter: "RetainOnDeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "trustedKeyGroups", GoGetter: "TrustedKeyGroups"},
			_jsii_.MemberProperty{JsiiProperty: "trustedSigners", GoGetter: "TrustedSigners"},
			_jsii_.MemberProperty{JsiiProperty: "viewerCertificate", GoGetter: "ViewerCertificate"},
			_jsii_.MemberProperty{JsiiProperty: "viewerCertificateInput", GoGetter: "ViewerCertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "waitForDeployment", GoGetter: "WaitForDeployment"},
			_jsii_.MemberProperty{JsiiProperty: "waitForDeploymentInput", GoGetter: "WaitForDeploymentInput"},
			_jsii_.MemberProperty{JsiiProperty: "webAclId", GoGetter: "WebAclId"},
			_jsii_.MemberProperty{JsiiProperty: "webAclIdInput", GoGetter: "WebAclIdInput"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistribution{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionConfig",
		reflect.TypeFor[CloudfrontDistributionConfig](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionCustomErrorResponse",
		reflect.TypeFor[CloudfrontDistributionCustomErrorResponse](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionCustomErrorResponseList",
		reflect.TypeFor[CloudfrontDistributionCustomErrorResponseList](),
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
			j := jsiiProxy_CloudfrontDistributionCustomErrorResponseList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionCustomErrorResponseOutputReference",
		reflect.TypeFor[CloudfrontDistributionCustomErrorResponseOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "errorCachingMinTtl", GoGetter: "ErrorCachingMinTtl"},
			_jsii_.MemberProperty{JsiiProperty: "errorCachingMinTtlInput", GoGetter: "ErrorCachingMinTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "errorCode", GoGetter: "ErrorCode"},
			_jsii_.MemberProperty{JsiiProperty: "errorCodeInput", GoGetter: "ErrorCodeInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetErrorCachingMinTtl", GoMethod: "ResetErrorCachingMinTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseCode", GoMethod: "ResetResponseCode"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponsePagePath", GoMethod: "ResetResponsePagePath"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "responseCode", GoGetter: "ResponseCode"},
			_jsii_.MemberProperty{JsiiProperty: "responseCodeInput", GoGetter: "ResponseCodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "responsePagePath", GoGetter: "ResponsePagePath"},
			_jsii_.MemberProperty{JsiiProperty: "responsePagePathInput", GoGetter: "ResponsePagePathInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionCustomErrorResponseOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionDefaultCacheBehavior",
		reflect.TypeFor[CloudfrontDistributionDefaultCacheBehavior](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionDefaultCacheBehaviorForwardedValues",
		reflect.TypeFor[CloudfrontDistributionDefaultCacheBehaviorForwardedValues](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookies",
		reflect.TypeFor[CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookies](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference",
		reflect.TypeFor[CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "forward", GoGetter: "Forward"},
			_jsii_.MemberProperty{JsiiProperty: "forwardInput", GoGetter: "ForwardInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetWhitelistedNames", GoMethod: "ResetWhitelistedNames"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "whitelistedNames", GoGetter: "WhitelistedNames"},
			_jsii_.MemberProperty{JsiiProperty: "whitelistedNamesInput", GoGetter: "WhitelistedNamesInput"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference",
		reflect.TypeFor[CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "cookies", GoGetter: "Cookies"},
			_jsii_.MemberProperty{JsiiProperty: "cookiesInput", GoGetter: "CookiesInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "headers", GoGetter: "Headers"},
			_jsii_.MemberProperty{JsiiProperty: "headersInput", GoGetter: "HeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putCookies", GoMethod: "PutCookies"},
			_jsii_.MemberProperty{JsiiProperty: "queryString", GoGetter: "QueryString"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringCacheKeys", GoGetter: "QueryStringCacheKeys"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringCacheKeysInput", GoGetter: "QueryStringCacheKeysInput"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringInput", GoGetter: "QueryStringInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaders", GoMethod: "ResetHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryStringCacheKeys", GoMethod: "ResetQueryStringCacheKeys"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionDefaultCacheBehaviorFunctionAssociation",
		reflect.TypeFor[CloudfrontDistributionDefaultCacheBehaviorFunctionAssociation](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionDefaultCacheBehaviorFunctionAssociationList",
		reflect.TypeFor[CloudfrontDistributionDefaultCacheBehaviorFunctionAssociationList](),
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
			j := jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorFunctionAssociationList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionDefaultCacheBehaviorFunctionAssociationOutputReference",
		reflect.TypeFor[CloudfrontDistributionDefaultCacheBehaviorFunctionAssociationOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "eventType", GoGetter: "EventType"},
			_jsii_.MemberProperty{JsiiProperty: "eventTypeInput", GoGetter: "EventTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "functionArn", GoGetter: "FunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "functionArnInput", GoGetter: "FunctionArnInput"},
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
			j := jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorFunctionAssociationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociation",
		reflect.TypeFor[CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociation](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociationList",
		reflect.TypeFor[CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociationList](),
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
			j := jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociationList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociationOutputReference",
		reflect.TypeFor[CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociationOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "eventType", GoGetter: "EventType"},
			_jsii_.MemberProperty{JsiiProperty: "eventTypeInput", GoGetter: "EventTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "includeBody", GoGetter: "IncludeBody"},
			_jsii_.MemberProperty{JsiiProperty: "includeBodyInput", GoGetter: "IncludeBodyInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaArn", GoGetter: "LambdaArn"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaArnInput", GoGetter: "LambdaArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeBody", GoMethod: "ResetIncludeBody"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionDefaultCacheBehaviorOutputReference",
		reflect.TypeFor[CloudfrontDistributionDefaultCacheBehaviorOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowedMethods", GoGetter: "AllowedMethods"},
			_jsii_.MemberProperty{JsiiProperty: "allowedMethodsInput", GoGetter: "AllowedMethodsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cachedMethods", GoGetter: "CachedMethods"},
			_jsii_.MemberProperty{JsiiProperty: "cachedMethodsInput", GoGetter: "CachedMethodsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cachePolicyId", GoGetter: "CachePolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "cachePolicyIdInput", GoGetter: "CachePolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compress", GoGetter: "Compress"},
			_jsii_.MemberProperty{JsiiProperty: "compressInput", GoGetter: "CompressInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTtl", GoGetter: "DefaultTtl"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTtlInput", GoGetter: "DefaultTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "fieldLevelEncryptionId", GoGetter: "FieldLevelEncryptionId"},
			_jsii_.MemberProperty{JsiiProperty: "fieldLevelEncryptionIdInput", GoGetter: "FieldLevelEncryptionIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "forwardedValues", GoGetter: "ForwardedValues"},
			_jsii_.MemberProperty{JsiiProperty: "forwardedValuesInput", GoGetter: "ForwardedValuesInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "functionAssociation", GoGetter: "FunctionAssociation"},
			_jsii_.MemberProperty{JsiiProperty: "functionAssociationInput", GoGetter: "FunctionAssociationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunctionAssociation", GoGetter: "LambdaFunctionAssociation"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunctionAssociationInput", GoGetter: "LambdaFunctionAssociationInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtl", GoGetter: "MaxTtl"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtlInput", GoGetter: "MaxTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "minTtl", GoGetter: "MinTtl"},
			_jsii_.MemberProperty{JsiiProperty: "minTtlInput", GoGetter: "MinTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "originRequestPolicyId", GoGetter: "OriginRequestPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "originRequestPolicyIdInput", GoGetter: "OriginRequestPolicyIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "putForwardedValues", GoMethod: "PutForwardedValues"},
			_jsii_.MemberMethod{JsiiMethod: "putFunctionAssociation", GoMethod: "PutFunctionAssociation"},
			_jsii_.MemberMethod{JsiiMethod: "putLambdaFunctionAssociation", GoMethod: "PutLambdaFunctionAssociation"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeLogConfigArn", GoGetter: "RealtimeLogConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeLogConfigArnInput", GoGetter: "RealtimeLogConfigArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCachePolicyId", GoMethod: "ResetCachePolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompress", GoMethod: "ResetCompress"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultTtl", GoMethod: "ResetDefaultTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetFieldLevelEncryptionId", GoMethod: "ResetFieldLevelEncryptionId"},
			_jsii_.MemberMethod{JsiiMethod: "resetForwardedValues", GoMethod: "ResetForwardedValues"},
			_jsii_.MemberMethod{JsiiMethod: "resetFunctionAssociation", GoMethod: "ResetFunctionAssociation"},
			_jsii_.MemberMethod{JsiiMethod: "resetLambdaFunctionAssociation", GoMethod: "ResetLambdaFunctionAssociation"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxTtl", GoMethod: "ResetMaxTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinTtl", GoMethod: "ResetMinTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginRequestPolicyId", GoMethod: "ResetOriginRequestPolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRealtimeLogConfigArn", GoMethod: "ResetRealtimeLogConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseHeadersPolicyId", GoMethod: "ResetResponseHeadersPolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSmoothStreaming", GoMethod: "ResetSmoothStreaming"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrustedKeyGroups", GoMethod: "ResetTrustedKeyGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrustedSigners", GoMethod: "ResetTrustedSigners"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersPolicyId", GoGetter: "ResponseHeadersPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersPolicyIdInput", GoGetter: "ResponseHeadersPolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "smoothStreaming", GoGetter: "SmoothStreaming"},
			_jsii_.MemberProperty{JsiiProperty: "smoothStreamingInput", GoGetter: "SmoothStreamingInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetOriginId", GoGetter: "TargetOriginId"},
			_jsii_.MemberProperty{JsiiProperty: "targetOriginIdInput", GoGetter: "TargetOriginIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trustedKeyGroups", GoGetter: "TrustedKeyGroups"},
			_jsii_.MemberProperty{JsiiProperty: "trustedKeyGroupsInput", GoGetter: "TrustedKeyGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "trustedSigners", GoGetter: "TrustedSigners"},
			_jsii_.MemberProperty{JsiiProperty: "trustedSignersInput", GoGetter: "TrustedSignersInput"},
			_jsii_.MemberProperty{JsiiProperty: "viewerProtocolPolicy", GoGetter: "ViewerProtocolPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "viewerProtocolPolicyInput", GoGetter: "ViewerProtocolPolicyInput"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionLoggingConfig",
		reflect.TypeFor[CloudfrontDistributionLoggingConfig](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionLoggingConfigOutputReference",
		reflect.TypeFor[CloudfrontDistributionLoggingConfigOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "bucketInput", GoGetter: "BucketInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "includeCookies", GoGetter: "IncludeCookies"},
			_jsii_.MemberProperty{JsiiProperty: "includeCookiesInput", GoGetter: "IncludeCookiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeCookies", GoMethod: "ResetIncludeCookies"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefix", GoMethod: "ResetPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOrderedCacheBehavior",
		reflect.TypeFor[CloudfrontDistributionOrderedCacheBehavior](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOrderedCacheBehaviorForwardedValues",
		reflect.TypeFor[CloudfrontDistributionOrderedCacheBehaviorForwardedValues](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookies",
		reflect.TypeFor[CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookies](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference",
		reflect.TypeFor[CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "forward", GoGetter: "Forward"},
			_jsii_.MemberProperty{JsiiProperty: "forwardInput", GoGetter: "ForwardInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetWhitelistedNames", GoMethod: "ResetWhitelistedNames"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "whitelistedNames", GoGetter: "WhitelistedNames"},
			_jsii_.MemberProperty{JsiiProperty: "whitelistedNamesInput", GoGetter: "WhitelistedNamesInput"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference",
		reflect.TypeFor[CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "cookies", GoGetter: "Cookies"},
			_jsii_.MemberProperty{JsiiProperty: "cookiesInput", GoGetter: "CookiesInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "headers", GoGetter: "Headers"},
			_jsii_.MemberProperty{JsiiProperty: "headersInput", GoGetter: "HeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putCookies", GoMethod: "PutCookies"},
			_jsii_.MemberProperty{JsiiProperty: "queryString", GoGetter: "QueryString"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringCacheKeys", GoGetter: "QueryStringCacheKeys"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringCacheKeysInput", GoGetter: "QueryStringCacheKeysInput"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringInput", GoGetter: "QueryStringInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaders", GoMethod: "ResetHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryStringCacheKeys", GoMethod: "ResetQueryStringCacheKeys"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOrderedCacheBehaviorFunctionAssociation",
		reflect.TypeFor[CloudfrontDistributionOrderedCacheBehaviorFunctionAssociation](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOrderedCacheBehaviorFunctionAssociationList",
		reflect.TypeFor[CloudfrontDistributionOrderedCacheBehaviorFunctionAssociationList](),
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
			j := jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorFunctionAssociationList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOrderedCacheBehaviorFunctionAssociationOutputReference",
		reflect.TypeFor[CloudfrontDistributionOrderedCacheBehaviorFunctionAssociationOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "eventType", GoGetter: "EventType"},
			_jsii_.MemberProperty{JsiiProperty: "eventTypeInput", GoGetter: "EventTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "functionArn", GoGetter: "FunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "functionArnInput", GoGetter: "FunctionArnInput"},
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
			j := jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorFunctionAssociationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOrderedCacheBehaviorLambdaFunctionAssociation",
		reflect.TypeFor[CloudfrontDistributionOrderedCacheBehaviorLambdaFunctionAssociation](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOrderedCacheBehaviorLambdaFunctionAssociationList",
		reflect.TypeFor[CloudfrontDistributionOrderedCacheBehaviorLambdaFunctionAssociationList](),
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
			j := jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorLambdaFunctionAssociationList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOrderedCacheBehaviorLambdaFunctionAssociationOutputReference",
		reflect.TypeFor[CloudfrontDistributionOrderedCacheBehaviorLambdaFunctionAssociationOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "eventType", GoGetter: "EventType"},
			_jsii_.MemberProperty{JsiiProperty: "eventTypeInput", GoGetter: "EventTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "includeBody", GoGetter: "IncludeBody"},
			_jsii_.MemberProperty{JsiiProperty: "includeBodyInput", GoGetter: "IncludeBodyInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaArn", GoGetter: "LambdaArn"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaArnInput", GoGetter: "LambdaArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeBody", GoMethod: "ResetIncludeBody"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorLambdaFunctionAssociationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOrderedCacheBehaviorList",
		reflect.TypeFor[CloudfrontDistributionOrderedCacheBehaviorList](),
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
			j := jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOrderedCacheBehaviorOutputReference",
		reflect.TypeFor[CloudfrontDistributionOrderedCacheBehaviorOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowedMethods", GoGetter: "AllowedMethods"},
			_jsii_.MemberProperty{JsiiProperty: "allowedMethodsInput", GoGetter: "AllowedMethodsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cachedMethods", GoGetter: "CachedMethods"},
			_jsii_.MemberProperty{JsiiProperty: "cachedMethodsInput", GoGetter: "CachedMethodsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cachePolicyId", GoGetter: "CachePolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "cachePolicyIdInput", GoGetter: "CachePolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compress", GoGetter: "Compress"},
			_jsii_.MemberProperty{JsiiProperty: "compressInput", GoGetter: "CompressInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTtl", GoGetter: "DefaultTtl"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTtlInput", GoGetter: "DefaultTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "fieldLevelEncryptionId", GoGetter: "FieldLevelEncryptionId"},
			_jsii_.MemberProperty{JsiiProperty: "fieldLevelEncryptionIdInput", GoGetter: "FieldLevelEncryptionIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "forwardedValues", GoGetter: "ForwardedValues"},
			_jsii_.MemberProperty{JsiiProperty: "forwardedValuesInput", GoGetter: "ForwardedValuesInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "functionAssociation", GoGetter: "FunctionAssociation"},
			_jsii_.MemberProperty{JsiiProperty: "functionAssociationInput", GoGetter: "FunctionAssociationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunctionAssociation", GoGetter: "LambdaFunctionAssociation"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunctionAssociationInput", GoGetter: "LambdaFunctionAssociationInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtl", GoGetter: "MaxTtl"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtlInput", GoGetter: "MaxTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "minTtl", GoGetter: "MinTtl"},
			_jsii_.MemberProperty{JsiiProperty: "minTtlInput", GoGetter: "MinTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "originRequestPolicyId", GoGetter: "OriginRequestPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "originRequestPolicyIdInput", GoGetter: "OriginRequestPolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "pathPattern", GoGetter: "PathPattern"},
			_jsii_.MemberProperty{JsiiProperty: "pathPatternInput", GoGetter: "PathPatternInput"},
			_jsii_.MemberMethod{JsiiMethod: "putForwardedValues", GoMethod: "PutForwardedValues"},
			_jsii_.MemberMethod{JsiiMethod: "putFunctionAssociation", GoMethod: "PutFunctionAssociation"},
			_jsii_.MemberMethod{JsiiMethod: "putLambdaFunctionAssociation", GoMethod: "PutLambdaFunctionAssociation"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeLogConfigArn", GoGetter: "RealtimeLogConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeLogConfigArnInput", GoGetter: "RealtimeLogConfigArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCachePolicyId", GoMethod: "ResetCachePolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompress", GoMethod: "ResetCompress"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultTtl", GoMethod: "ResetDefaultTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetFieldLevelEncryptionId", GoMethod: "ResetFieldLevelEncryptionId"},
			_jsii_.MemberMethod{JsiiMethod: "resetForwardedValues", GoMethod: "ResetForwardedValues"},
			_jsii_.MemberMethod{JsiiMethod: "resetFunctionAssociation", GoMethod: "ResetFunctionAssociation"},
			_jsii_.MemberMethod{JsiiMethod: "resetLambdaFunctionAssociation", GoMethod: "ResetLambdaFunctionAssociation"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxTtl", GoMethod: "ResetMaxTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinTtl", GoMethod: "ResetMinTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginRequestPolicyId", GoMethod: "ResetOriginRequestPolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRealtimeLogConfigArn", GoMethod: "ResetRealtimeLogConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseHeadersPolicyId", GoMethod: "ResetResponseHeadersPolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSmoothStreaming", GoMethod: "ResetSmoothStreaming"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrustedKeyGroups", GoMethod: "ResetTrustedKeyGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrustedSigners", GoMethod: "ResetTrustedSigners"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersPolicyId", GoGetter: "ResponseHeadersPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersPolicyIdInput", GoGetter: "ResponseHeadersPolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "smoothStreaming", GoGetter: "SmoothStreaming"},
			_jsii_.MemberProperty{JsiiProperty: "smoothStreamingInput", GoGetter: "SmoothStreamingInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetOriginId", GoGetter: "TargetOriginId"},
			_jsii_.MemberProperty{JsiiProperty: "targetOriginIdInput", GoGetter: "TargetOriginIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trustedKeyGroups", GoGetter: "TrustedKeyGroups"},
			_jsii_.MemberProperty{JsiiProperty: "trustedKeyGroupsInput", GoGetter: "TrustedKeyGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "trustedSigners", GoGetter: "TrustedSigners"},
			_jsii_.MemberProperty{JsiiProperty: "trustedSignersInput", GoGetter: "TrustedSignersInput"},
			_jsii_.MemberProperty{JsiiProperty: "viewerProtocolPolicy", GoGetter: "ViewerProtocolPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "viewerProtocolPolicyInput", GoGetter: "ViewerProtocolPolicyInput"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOrigin",
		reflect.TypeFor[CloudfrontDistributionOrigin](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOriginCustomHeader",
		reflect.TypeFor[CloudfrontDistributionOriginCustomHeader](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOriginCustomHeaderList",
		reflect.TypeFor[CloudfrontDistributionOriginCustomHeaderList](),
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
			j := jsiiProxy_CloudfrontDistributionOriginCustomHeaderList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOriginCustomHeaderOutputReference",
		reflect.TypeFor[CloudfrontDistributionOriginCustomHeaderOutputReference](),
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
			j := jsiiProxy_CloudfrontDistributionOriginCustomHeaderOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOriginCustomOriginConfig",
		reflect.TypeFor[CloudfrontDistributionOriginCustomOriginConfig](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOriginCustomOriginConfigOutputReference",
		reflect.TypeFor[CloudfrontDistributionOriginCustomOriginConfigOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "httpPort", GoGetter: "HttpPort"},
			_jsii_.MemberProperty{JsiiProperty: "httpPortInput", GoGetter: "HttpPortInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpsPort", GoGetter: "HttpsPort"},
			_jsii_.MemberProperty{JsiiProperty: "httpsPortInput", GoGetter: "HttpsPortInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "originKeepaliveTimeout", GoGetter: "OriginKeepaliveTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "originKeepaliveTimeoutInput", GoGetter: "OriginKeepaliveTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "originProtocolPolicy", GoGetter: "OriginProtocolPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "originProtocolPolicyInput", GoGetter: "OriginProtocolPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "originReadTimeout", GoGetter: "OriginReadTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "originReadTimeoutInput", GoGetter: "OriginReadTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "originSslProtocols", GoGetter: "OriginSslProtocols"},
			_jsii_.MemberProperty{JsiiProperty: "originSslProtocolsInput", GoGetter: "OriginSslProtocolsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginKeepaliveTimeout", GoMethod: "ResetOriginKeepaliveTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginReadTimeout", GoMethod: "ResetOriginReadTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOriginGroup",
		reflect.TypeFor[CloudfrontDistributionOriginGroup](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOriginGroupFailoverCriteria",
		reflect.TypeFor[CloudfrontDistributionOriginGroupFailoverCriteria](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference",
		reflect.TypeFor[CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference](),
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "statusCodes", GoGetter: "StatusCodes"},
			_jsii_.MemberProperty{JsiiProperty: "statusCodesInput", GoGetter: "StatusCodesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOriginGroupList",
		reflect.TypeFor[CloudfrontDistributionOriginGroupList](),
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
			j := jsiiProxy_CloudfrontDistributionOriginGroupList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOriginGroupMember",
		reflect.TypeFor[CloudfrontDistributionOriginGroupMember](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOriginGroupMemberList",
		reflect.TypeFor[CloudfrontDistributionOriginGroupMemberList](),
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
			j := jsiiProxy_CloudfrontDistributionOriginGroupMemberList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOriginGroupMemberOutputReference",
		reflect.TypeFor[CloudfrontDistributionOriginGroupMemberOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "originId", GoGetter: "OriginId"},
			_jsii_.MemberProperty{JsiiProperty: "originIdInput", GoGetter: "OriginIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionOriginGroupMemberOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOriginGroupOutputReference",
		reflect.TypeFor[CloudfrontDistributionOriginGroupOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "failoverCriteria", GoGetter: "FailoverCriteria"},
			_jsii_.MemberProperty{JsiiProperty: "failoverCriteriaInput", GoGetter: "FailoverCriteriaInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "member", GoGetter: "Member"},
			_jsii_.MemberProperty{JsiiProperty: "memberInput", GoGetter: "MemberInput"},
			_jsii_.MemberProperty{JsiiProperty: "originId", GoGetter: "OriginId"},
			_jsii_.MemberProperty{JsiiProperty: "originIdInput", GoGetter: "OriginIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "putFailoverCriteria", GoMethod: "PutFailoverCriteria"},
			_jsii_.MemberMethod{JsiiMethod: "putMember", GoMethod: "PutMember"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionOriginGroupOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOriginList",
		reflect.TypeFor[CloudfrontDistributionOriginList](),
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
			j := jsiiProxy_CloudfrontDistributionOriginList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOriginOriginShield",
		reflect.TypeFor[CloudfrontDistributionOriginOriginShield](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOriginOriginShieldOutputReference",
		reflect.TypeFor[CloudfrontDistributionOriginOriginShieldOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "originShieldRegion", GoGetter: "OriginShieldRegion"},
			_jsii_.MemberProperty{JsiiProperty: "originShieldRegionInput", GoGetter: "OriginShieldRegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOriginOutputReference",
		reflect.TypeFor[CloudfrontDistributionOriginOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "connectionAttempts", GoGetter: "ConnectionAttempts"},
			_jsii_.MemberProperty{JsiiProperty: "connectionAttemptsInput", GoGetter: "ConnectionAttemptsInput"},
			_jsii_.MemberProperty{JsiiProperty: "connectionTimeout", GoGetter: "ConnectionTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "connectionTimeoutInput", GoGetter: "ConnectionTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customHeader", GoGetter: "CustomHeader"},
			_jsii_.MemberProperty{JsiiProperty: "customHeaderInput", GoGetter: "CustomHeaderInput"},
			_jsii_.MemberProperty{JsiiProperty: "customOriginConfig", GoGetter: "CustomOriginConfig"},
			_jsii_.MemberProperty{JsiiProperty: "customOriginConfigInput", GoGetter: "CustomOriginConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "domainNameInput", GoGetter: "DomainNameInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "originAccessControlId", GoGetter: "OriginAccessControlId"},
			_jsii_.MemberProperty{JsiiProperty: "originAccessControlIdInput", GoGetter: "OriginAccessControlIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "originId", GoGetter: "OriginId"},
			_jsii_.MemberProperty{JsiiProperty: "originIdInput", GoGetter: "OriginIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "originPath", GoGetter: "OriginPath"},
			_jsii_.MemberProperty{JsiiProperty: "originPathInput", GoGetter: "OriginPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "originShield", GoGetter: "OriginShield"},
			_jsii_.MemberProperty{JsiiProperty: "originShieldInput", GoGetter: "OriginShieldInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomHeader", GoMethod: "PutCustomHeader"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomOriginConfig", GoMethod: "PutCustomOriginConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putOriginShield", GoMethod: "PutOriginShield"},
			_jsii_.MemberMethod{JsiiMethod: "putS3OriginConfig", GoMethod: "PutS3OriginConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectionAttempts", GoMethod: "ResetConnectionAttempts"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectionTimeout", GoMethod: "ResetConnectionTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomHeader", GoMethod: "ResetCustomHeader"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomOriginConfig", GoMethod: "ResetCustomOriginConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginAccessControlId", GoMethod: "ResetOriginAccessControlId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginPath", GoMethod: "ResetOriginPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginShield", GoMethod: "ResetOriginShield"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3OriginConfig", GoMethod: "ResetS3OriginConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "s3OriginConfig", GoGetter: "S3OriginConfig"},
			_jsii_.MemberProperty{JsiiProperty: "s3OriginConfigInput", GoGetter: "S3OriginConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionOriginOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOriginS3OriginConfig",
		reflect.TypeFor[CloudfrontDistributionOriginS3OriginConfig](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionOriginS3OriginConfigOutputReference",
		reflect.TypeFor[CloudfrontDistributionOriginS3OriginConfigOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "originAccessIdentity", GoGetter: "OriginAccessIdentity"},
			_jsii_.MemberProperty{JsiiProperty: "originAccessIdentityInput", GoGetter: "OriginAccessIdentityInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionOriginS3OriginConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionRestrictions",
		reflect.TypeFor[CloudfrontDistributionRestrictions](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionRestrictionsGeoRestriction",
		reflect.TypeFor[CloudfrontDistributionRestrictionsGeoRestriction](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionRestrictionsGeoRestrictionOutputReference",
		reflect.TypeFor[CloudfrontDistributionRestrictionsGeoRestrictionOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "locations", GoGetter: "Locations"},
			_jsii_.MemberProperty{JsiiProperty: "locationsInput", GoGetter: "LocationsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocations", GoMethod: "ResetLocations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionType", GoGetter: "RestrictionType"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionTypeInput", GoGetter: "RestrictionTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionRestrictionsOutputReference",
		reflect.TypeFor[CloudfrontDistributionRestrictionsOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "geoRestriction", GoGetter: "GeoRestriction"},
			_jsii_.MemberProperty{JsiiProperty: "geoRestrictionInput", GoGetter: "GeoRestrictionInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putGeoRestriction", GoMethod: "PutGeoRestriction"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionRestrictionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionTrustedKeyGroups",
		reflect.TypeFor[CloudfrontDistributionTrustedKeyGroups](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionTrustedKeyGroupsItems",
		reflect.TypeFor[CloudfrontDistributionTrustedKeyGroupsItems](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionTrustedKeyGroupsItemsList",
		reflect.TypeFor[CloudfrontDistributionTrustedKeyGroupsItemsList](),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionTrustedKeyGroupsItemsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionTrustedKeyGroupsItemsOutputReference",
		reflect.TypeFor[CloudfrontDistributionTrustedKeyGroupsItemsOutputReference](),
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
			_jsii_.MemberProperty{JsiiProperty: "keyGroupId", GoGetter: "KeyGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "keyPairIds", GoGetter: "KeyPairIds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionTrustedKeyGroupsItemsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionTrustedKeyGroupsList",
		reflect.TypeFor[CloudfrontDistributionTrustedKeyGroupsList](),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionTrustedKeyGroupsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionTrustedKeyGroupsOutputReference",
		reflect.TypeFor[CloudfrontDistributionTrustedKeyGroupsOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
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
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionTrustedKeyGroupsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionTrustedSigners",
		reflect.TypeFor[CloudfrontDistributionTrustedSigners](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionTrustedSignersItems",
		reflect.TypeFor[CloudfrontDistributionTrustedSignersItems](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionTrustedSignersItemsList",
		reflect.TypeFor[CloudfrontDistributionTrustedSignersItemsList](),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionTrustedSignersItemsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionTrustedSignersItemsOutputReference",
		reflect.TypeFor[CloudfrontDistributionTrustedSignersItemsOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awsAccountNumber", GoGetter: "AwsAccountNumber"},
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
			_jsii_.MemberProperty{JsiiProperty: "keyPairIds", GoGetter: "KeyPairIds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionTrustedSignersItemsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionTrustedSignersList",
		reflect.TypeFor[CloudfrontDistributionTrustedSignersList](),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionTrustedSignersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionTrustedSignersOutputReference",
		reflect.TypeFor[CloudfrontDistributionTrustedSignersOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
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
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionTrustedSignersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionViewerCertificate",
		reflect.TypeFor[CloudfrontDistributionViewerCertificate](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cloudfrontDistribution.CloudfrontDistributionViewerCertificateOutputReference",
		reflect.TypeFor[CloudfrontDistributionViewerCertificateOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acmCertificateArn", GoGetter: "AcmCertificateArn"},
			_jsii_.MemberProperty{JsiiProperty: "acmCertificateArnInput", GoGetter: "AcmCertificateArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudfrontDefaultCertificate", GoGetter: "CloudfrontDefaultCertificate"},
			_jsii_.MemberProperty{JsiiProperty: "cloudfrontDefaultCertificateInput", GoGetter: "CloudfrontDefaultCertificateInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "iamCertificateId", GoGetter: "IamCertificateId"},
			_jsii_.MemberProperty{JsiiProperty: "iamCertificateIdInput", GoGetter: "IamCertificateIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "minimumProtocolVersion", GoGetter: "MinimumProtocolVersion"},
			_jsii_.MemberProperty{JsiiProperty: "minimumProtocolVersionInput", GoGetter: "MinimumProtocolVersionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAcmCertificateArn", GoMethod: "ResetAcmCertificateArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudfrontDefaultCertificate", GoMethod: "ResetCloudfrontDefaultCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetIamCertificateId", GoMethod: "ResetIamCertificateId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinimumProtocolVersion", GoMethod: "ResetMinimumProtocolVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetSslSupportMethod", GoMethod: "ResetSslSupportMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sslSupportMethod", GoGetter: "SslSupportMethod"},
			_jsii_.MemberProperty{JsiiProperty: "sslSupportMethodInput", GoGetter: "SslSupportMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}

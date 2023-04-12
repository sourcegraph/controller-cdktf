package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.provider.OktaProvider",
		reflect.TypeOf((*OktaProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessToken", GoGetter: "AccessToken"},
			_jsii_.MemberProperty{JsiiProperty: "accessTokenInput", GoGetter: "AccessTokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "apiToken", GoGetter: "ApiToken"},
			_jsii_.MemberProperty{JsiiProperty: "apiTokenInput", GoGetter: "ApiTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "backoff", GoGetter: "Backoff"},
			_jsii_.MemberProperty{JsiiProperty: "backoffInput", GoGetter: "BackoffInput"},
			_jsii_.MemberProperty{JsiiProperty: "baseUrl", GoGetter: "BaseUrl"},
			_jsii_.MemberProperty{JsiiProperty: "baseUrlInput", GoGetter: "BaseUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientId", GoGetter: "ClientId"},
			_jsii_.MemberProperty{JsiiProperty: "clientIdInput", GoGetter: "ClientIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "httpProxy", GoGetter: "HttpProxy"},
			_jsii_.MemberProperty{JsiiProperty: "httpProxyInput", GoGetter: "HttpProxyInput"},
			_jsii_.MemberProperty{JsiiProperty: "logLevel", GoGetter: "LogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "logLevelInput", GoGetter: "LogLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxApiCapacity", GoGetter: "MaxApiCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "maxApiCapacityInput", GoGetter: "MaxApiCapacityInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetries", GoGetter: "MaxRetries"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetriesInput", GoGetter: "MaxRetriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxWaitSeconds", GoGetter: "MaxWaitSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "maxWaitSecondsInput", GoGetter: "MaxWaitSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "minWaitSeconds", GoGetter: "MinWaitSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "minWaitSecondsInput", GoGetter: "MinWaitSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "orgName", GoGetter: "OrgName"},
			_jsii_.MemberProperty{JsiiProperty: "orgNameInput", GoGetter: "OrgNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "parallelism", GoGetter: "Parallelism"},
			_jsii_.MemberProperty{JsiiProperty: "parallelismInput", GoGetter: "ParallelismInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateKey", GoGetter: "PrivateKey"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyId", GoGetter: "PrivateKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyIdInput", GoGetter: "PrivateKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyInput", GoGetter: "PrivateKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "requestTimeout", GoGetter: "RequestTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "requestTimeoutInput", GoGetter: "RequestTimeoutInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessToken", GoMethod: "ResetAccessToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiToken", GoMethod: "ResetApiToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackoff", GoMethod: "ResetBackoff"},
			_jsii_.MemberMethod{JsiiMethod: "resetBaseUrl", GoMethod: "ResetBaseUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientId", GoMethod: "ResetClientId"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpProxy", GoMethod: "ResetHttpProxy"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogLevel", GoMethod: "ResetLogLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxApiCapacity", GoMethod: "ResetMaxApiCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxRetries", GoMethod: "ResetMaxRetries"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxWaitSeconds", GoMethod: "ResetMaxWaitSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinWaitSeconds", GoMethod: "ResetMinWaitSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrgName", GoMethod: "ResetOrgName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetParallelism", GoMethod: "ResetParallelism"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateKey", GoMethod: "ResetPrivateKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateKeyId", GoMethod: "ResetPrivateKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestTimeout", GoMethod: "ResetRequestTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_OktaProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.provider.OktaProviderConfig",
		reflect.TypeOf((*OktaProviderConfig)(nil)).Elem(),
	)
}

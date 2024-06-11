package tunnelconfig

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.tunnelConfig.TunnelConfigA",
		reflect.TypeOf((*TunnelConfigA)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "accountIdInput", GoGetter: "AccountIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "configInput", GoGetter: "ConfigInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putConfig", GoMethod: "PutConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelId", GoGetter: "TunnelId"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelIdInput", GoGetter: "TunnelIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_TunnelConfigA{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.tunnelConfig.TunnelConfigAConfig",
		reflect.TypeOf((*TunnelConfigAConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.tunnelConfig.TunnelConfigConfig",
		reflect.TypeOf((*TunnelConfigConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.tunnelConfig.TunnelConfigConfigIngressRule",
		reflect.TypeOf((*TunnelConfigConfigIngressRule)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.tunnelConfig.TunnelConfigConfigIngressRuleList",
		reflect.TypeOf((*TunnelConfigConfigIngressRuleList)(nil)).Elem(),
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
			j := jsiiProxy_TunnelConfigConfigIngressRuleList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.tunnelConfig.TunnelConfigConfigIngressRuleOutputReference",
		reflect.TypeOf((*TunnelConfigConfigIngressRuleOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "hostname", GoGetter: "Hostname"},
			_jsii_.MemberProperty{JsiiProperty: "hostnameInput", GoGetter: "HostnameInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathInput", GoGetter: "PathInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostname", GoMethod: "ResetHostname"},
			_jsii_.MemberMethod{JsiiMethod: "resetPath", GoMethod: "ResetPath"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "service", GoGetter: "Service"},
			_jsii_.MemberProperty{JsiiProperty: "serviceInput", GoGetter: "ServiceInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TunnelConfigConfigIngressRuleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.tunnelConfig.TunnelConfigConfigOriginRequest",
		reflect.TypeOf((*TunnelConfigConfigOriginRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.tunnelConfig.TunnelConfigConfigOriginRequestIpRules",
		reflect.TypeOf((*TunnelConfigConfigOriginRequestIpRules)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.tunnelConfig.TunnelConfigConfigOriginRequestIpRulesList",
		reflect.TypeOf((*TunnelConfigConfigOriginRequestIpRulesList)(nil)).Elem(),
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
			j := jsiiProxy_TunnelConfigConfigOriginRequestIpRulesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.tunnelConfig.TunnelConfigConfigOriginRequestIpRulesOutputReference",
		reflect.TypeOf((*TunnelConfigConfigOriginRequestIpRulesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allow", GoGetter: "Allow"},
			_jsii_.MemberProperty{JsiiProperty: "allowInput", GoGetter: "AllowInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "ports", GoGetter: "Ports"},
			_jsii_.MemberProperty{JsiiProperty: "portsInput", GoGetter: "PortsInput"},
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllow", GoMethod: "ResetAllow"},
			_jsii_.MemberMethod{JsiiMethod: "resetPorts", GoMethod: "ResetPorts"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefix", GoMethod: "ResetPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TunnelConfigConfigOriginRequestIpRulesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.tunnelConfig.TunnelConfigConfigOriginRequestOutputReference",
		reflect.TypeOf((*TunnelConfigConfigOriginRequestOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bastionMode", GoGetter: "BastionMode"},
			_jsii_.MemberProperty{JsiiProperty: "bastionModeInput", GoGetter: "BastionModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "caPool", GoGetter: "CaPool"},
			_jsii_.MemberProperty{JsiiProperty: "caPoolInput", GoGetter: "CaPoolInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "connectTimeout", GoGetter: "ConnectTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "connectTimeoutInput", GoGetter: "ConnectTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "disableChunkedEncoding", GoGetter: "DisableChunkedEncoding"},
			_jsii_.MemberProperty{JsiiProperty: "disableChunkedEncodingInput", GoGetter: "DisableChunkedEncodingInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "httpHostHeader", GoGetter: "HttpHostHeader"},
			_jsii_.MemberProperty{JsiiProperty: "httpHostHeaderInput", GoGetter: "HttpHostHeaderInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipRules", GoGetter: "IpRules"},
			_jsii_.MemberProperty{JsiiProperty: "ipRulesInput", GoGetter: "IpRulesInput"},
			_jsii_.MemberProperty{JsiiProperty: "keepAliveConnections", GoGetter: "KeepAliveConnections"},
			_jsii_.MemberProperty{JsiiProperty: "keepAliveConnectionsInput", GoGetter: "KeepAliveConnectionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "keepAliveTimeout", GoGetter: "KeepAliveTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "keepAliveTimeoutInput", GoGetter: "KeepAliveTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "noHappyEyeballs", GoGetter: "NoHappyEyeballs"},
			_jsii_.MemberProperty{JsiiProperty: "noHappyEyeballsInput", GoGetter: "NoHappyEyeballsInput"},
			_jsii_.MemberProperty{JsiiProperty: "noTlsVerify", GoGetter: "NoTlsVerify"},
			_jsii_.MemberProperty{JsiiProperty: "noTlsVerifyInput", GoGetter: "NoTlsVerifyInput"},
			_jsii_.MemberProperty{JsiiProperty: "originServerName", GoGetter: "OriginServerName"},
			_jsii_.MemberProperty{JsiiProperty: "originServerNameInput", GoGetter: "OriginServerNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "proxyAddress", GoGetter: "ProxyAddress"},
			_jsii_.MemberProperty{JsiiProperty: "proxyAddressInput", GoGetter: "ProxyAddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "proxyPort", GoGetter: "ProxyPort"},
			_jsii_.MemberProperty{JsiiProperty: "proxyPortInput", GoGetter: "ProxyPortInput"},
			_jsii_.MemberProperty{JsiiProperty: "proxyType", GoGetter: "ProxyType"},
			_jsii_.MemberProperty{JsiiProperty: "proxyTypeInput", GoGetter: "ProxyTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putIpRules", GoMethod: "PutIpRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetBastionMode", GoMethod: "ResetBastionMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetCaPool", GoMethod: "ResetCaPool"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectTimeout", GoMethod: "ResetConnectTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableChunkedEncoding", GoMethod: "ResetDisableChunkedEncoding"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpHostHeader", GoMethod: "ResetHttpHostHeader"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpRules", GoMethod: "ResetIpRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeepAliveConnections", GoMethod: "ResetKeepAliveConnections"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeepAliveTimeout", GoMethod: "ResetKeepAliveTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoHappyEyeballs", GoMethod: "ResetNoHappyEyeballs"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoTlsVerify", GoMethod: "ResetNoTlsVerify"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginServerName", GoMethod: "ResetOriginServerName"},
			_jsii_.MemberMethod{JsiiMethod: "resetProxyAddress", GoMethod: "ResetProxyAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resetProxyPort", GoMethod: "ResetProxyPort"},
			_jsii_.MemberMethod{JsiiMethod: "resetProxyType", GoMethod: "ResetProxyType"},
			_jsii_.MemberMethod{JsiiMethod: "resetTcpKeepAlive", GoMethod: "ResetTcpKeepAlive"},
			_jsii_.MemberMethod{JsiiMethod: "resetTlsTimeout", GoMethod: "ResetTlsTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "tcpKeepAlive", GoGetter: "TcpKeepAlive"},
			_jsii_.MemberProperty{JsiiProperty: "tcpKeepAliveInput", GoGetter: "TcpKeepAliveInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tlsTimeout", GoGetter: "TlsTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "tlsTimeoutInput", GoGetter: "TlsTimeoutInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TunnelConfigConfigOriginRequestOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.tunnelConfig.TunnelConfigConfigOutputReference",
		reflect.TypeOf((*TunnelConfigConfigOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "ingressRule", GoGetter: "IngressRule"},
			_jsii_.MemberProperty{JsiiProperty: "ingressRuleInput", GoGetter: "IngressRuleInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "originRequest", GoGetter: "OriginRequest"},
			_jsii_.MemberProperty{JsiiProperty: "originRequestInput", GoGetter: "OriginRequestInput"},
			_jsii_.MemberMethod{JsiiMethod: "putIngressRule", GoMethod: "PutIngressRule"},
			_jsii_.MemberMethod{JsiiMethod: "putOriginRequest", GoMethod: "PutOriginRequest"},
			_jsii_.MemberMethod{JsiiMethod: "putWarpRouting", GoMethod: "PutWarpRouting"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginRequest", GoMethod: "ResetOriginRequest"},
			_jsii_.MemberMethod{JsiiMethod: "resetWarpRouting", GoMethod: "ResetWarpRouting"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "warpRouting", GoGetter: "WarpRouting"},
			_jsii_.MemberProperty{JsiiProperty: "warpRoutingInput", GoGetter: "WarpRoutingInput"},
		},
		func() interface{} {
			j := jsiiProxy_TunnelConfigConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.tunnelConfig.TunnelConfigConfigWarpRouting",
		reflect.TypeOf((*TunnelConfigConfigWarpRouting)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.tunnelConfig.TunnelConfigConfigWarpRoutingOutputReference",
		reflect.TypeOf((*TunnelConfigConfigWarpRoutingOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TunnelConfigConfigWarpRoutingOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}

package tunnelconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/tunnelconfig/internal"
)

type TunnelConfigConfigIngressRuleOriginRequestOutputReference interface {
	cdktf.ComplexObject
	Access() TunnelConfigConfigIngressRuleOriginRequestAccessOutputReference
	AccessInput() *TunnelConfigConfigIngressRuleOriginRequestAccess
	BastionMode() interface{}
	SetBastionMode(val interface{})
	BastionModeInput() interface{}
	CaPool() *string
	SetCaPool(val *string)
	CaPoolInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ConnectTimeout() *string
	SetConnectTimeout(val *string)
	ConnectTimeoutInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisableChunkedEncoding() interface{}
	SetDisableChunkedEncoding(val interface{})
	DisableChunkedEncodingInput() interface{}
	// Experimental.
	Fqn() *string
	Http2Origin() interface{}
	SetHttp2Origin(val interface{})
	Http2OriginInput() interface{}
	HttpHostHeader() *string
	SetHttpHostHeader(val *string)
	HttpHostHeaderInput() *string
	InternalValue() *TunnelConfigConfigIngressRuleOriginRequest
	SetInternalValue(val *TunnelConfigConfigIngressRuleOriginRequest)
	IpRules() TunnelConfigConfigIngressRuleOriginRequestIpRulesList
	IpRulesInput() interface{}
	KeepAliveConnections() *float64
	SetKeepAliveConnections(val *float64)
	KeepAliveConnectionsInput() *float64
	KeepAliveTimeout() *string
	SetKeepAliveTimeout(val *string)
	KeepAliveTimeoutInput() *string
	NoHappyEyeballs() interface{}
	SetNoHappyEyeballs(val interface{})
	NoHappyEyeballsInput() interface{}
	NoTlsVerify() interface{}
	SetNoTlsVerify(val interface{})
	NoTlsVerifyInput() interface{}
	OriginServerName() *string
	SetOriginServerName(val *string)
	OriginServerNameInput() *string
	ProxyAddress() *string
	SetProxyAddress(val *string)
	ProxyAddressInput() *string
	ProxyPort() *float64
	SetProxyPort(val *float64)
	ProxyPortInput() *float64
	ProxyType() *string
	SetProxyType(val *string)
	ProxyTypeInput() *string
	TcpKeepAlive() *string
	SetTcpKeepAlive(val *string)
	TcpKeepAliveInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsTimeout() *string
	SetTlsTimeout(val *string)
	TlsTimeoutInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutAccess(value *TunnelConfigConfigIngressRuleOriginRequestAccess)
	PutIpRules(value interface{})
	ResetAccess()
	ResetBastionMode()
	ResetCaPool()
	ResetConnectTimeout()
	ResetDisableChunkedEncoding()
	ResetHttp2Origin()
	ResetHttpHostHeader()
	ResetIpRules()
	ResetKeepAliveConnections()
	ResetKeepAliveTimeout()
	ResetNoHappyEyeballs()
	ResetNoTlsVerify()
	ResetOriginServerName()
	ResetProxyAddress()
	ResetProxyPort()
	ResetProxyType()
	ResetTcpKeepAlive()
	ResetTlsTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TunnelConfigConfigIngressRuleOriginRequestOutputReference
type jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) Access() TunnelConfigConfigIngressRuleOriginRequestAccessOutputReference {
	var returns TunnelConfigConfigIngressRuleOriginRequestAccessOutputReference
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) AccessInput() *TunnelConfigConfigIngressRuleOriginRequestAccess {
	var returns *TunnelConfigConfigIngressRuleOriginRequestAccess
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) BastionMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bastionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) BastionModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bastionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) CaPool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) CaPoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ConnectTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ConnectTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) DisableChunkedEncoding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableChunkedEncoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) DisableChunkedEncodingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableChunkedEncodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) Http2Origin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2Origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) Http2OriginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2OriginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) HttpHostHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpHostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) HttpHostHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpHostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) InternalValue() *TunnelConfigConfigIngressRuleOriginRequest {
	var returns *TunnelConfigConfigIngressRuleOriginRequest
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) IpRules() TunnelConfigConfigIngressRuleOriginRequestIpRulesList {
	var returns TunnelConfigConfigIngressRuleOriginRequestIpRulesList
	_jsii_.Get(
		j,
		"ipRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) IpRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) KeepAliveConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) KeepAliveConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) KeepAliveTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keepAliveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) KeepAliveTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keepAliveTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) NoHappyEyeballs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noHappyEyeballs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) NoHappyEyeballsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noHappyEyeballsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) NoTlsVerify() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noTlsVerify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) NoTlsVerifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noTlsVerifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) OriginServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originServerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) OriginServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originServerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ProxyAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ProxyAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ProxyPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"proxyPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ProxyPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"proxyPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ProxyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ProxyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) TcpKeepAlive() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tcpKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) TcpKeepAliveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tcpKeepAliveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) TlsTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) TlsTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsTimeoutInput",
		&returns,
	)
	return returns
}


func NewTunnelConfigConfigIngressRuleOriginRequestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TunnelConfigConfigIngressRuleOriginRequestOutputReference {
	_init_.Initialize()

	if err := validateNewTunnelConfigConfigIngressRuleOriginRequestOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.tunnelConfig.TunnelConfigConfigIngressRuleOriginRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTunnelConfigConfigIngressRuleOriginRequestOutputReference_Override(t TunnelConfigConfigIngressRuleOriginRequestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.tunnelConfig.TunnelConfigConfigIngressRuleOriginRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference)SetBastionMode(val interface{}) {
	if err := j.validateSetBastionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bastionMode",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference)SetCaPool(val *string) {
	if err := j.validateSetCaPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caPool",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference)SetConnectTimeout(val *string) {
	if err := j.validateSetConnectTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectTimeout",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference)SetDisableChunkedEncoding(val interface{}) {
	if err := j.validateSetDisableChunkedEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableChunkedEncoding",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference)SetHttp2Origin(val interface{}) {
	if err := j.validateSetHttp2OriginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"http2Origin",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference)SetHttpHostHeader(val *string) {
	if err := j.validateSetHttpHostHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpHostHeader",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference)SetInternalValue(val *TunnelConfigConfigIngressRuleOriginRequest) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference)SetKeepAliveConnections(val *float64) {
	if err := j.validateSetKeepAliveConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepAliveConnections",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference)SetKeepAliveTimeout(val *string) {
	if err := j.validateSetKeepAliveTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepAliveTimeout",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference)SetNoHappyEyeballs(val interface{}) {
	if err := j.validateSetNoHappyEyeballsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noHappyEyeballs",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference)SetNoTlsVerify(val interface{}) {
	if err := j.validateSetNoTlsVerifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noTlsVerify",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference)SetOriginServerName(val *string) {
	if err := j.validateSetOriginServerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originServerName",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference)SetProxyAddress(val *string) {
	if err := j.validateSetProxyAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyAddress",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference)SetProxyPort(val *float64) {
	if err := j.validateSetProxyPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyPort",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference)SetProxyType(val *string) {
	if err := j.validateSetProxyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyType",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference)SetTcpKeepAlive(val *string) {
	if err := j.validateSetTcpKeepAliveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpKeepAlive",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference)SetTlsTimeout(val *string) {
	if err := j.validateSetTlsTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsTimeout",
		val,
	)
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) PutAccess(value *TunnelConfigConfigIngressRuleOriginRequestAccess) {
	if err := t.validatePutAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAccess",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) PutIpRules(value interface{}) {
	if err := t.validatePutIpRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIpRules",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ResetAccess() {
	_jsii_.InvokeVoid(
		t,
		"resetAccess",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ResetBastionMode() {
	_jsii_.InvokeVoid(
		t,
		"resetBastionMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ResetCaPool() {
	_jsii_.InvokeVoid(
		t,
		"resetCaPool",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ResetConnectTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetConnectTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ResetDisableChunkedEncoding() {
	_jsii_.InvokeVoid(
		t,
		"resetDisableChunkedEncoding",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ResetHttp2Origin() {
	_jsii_.InvokeVoid(
		t,
		"resetHttp2Origin",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ResetHttpHostHeader() {
	_jsii_.InvokeVoid(
		t,
		"resetHttpHostHeader",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ResetIpRules() {
	_jsii_.InvokeVoid(
		t,
		"resetIpRules",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ResetKeepAliveConnections() {
	_jsii_.InvokeVoid(
		t,
		"resetKeepAliveConnections",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ResetKeepAliveTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetKeepAliveTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ResetNoHappyEyeballs() {
	_jsii_.InvokeVoid(
		t,
		"resetNoHappyEyeballs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ResetNoTlsVerify() {
	_jsii_.InvokeVoid(
		t,
		"resetNoTlsVerify",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ResetOriginServerName() {
	_jsii_.InvokeVoid(
		t,
		"resetOriginServerName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ResetProxyAddress() {
	_jsii_.InvokeVoid(
		t,
		"resetProxyAddress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ResetProxyPort() {
	_jsii_.InvokeVoid(
		t,
		"resetProxyPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ResetProxyType() {
	_jsii_.InvokeVoid(
		t,
		"resetProxyType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ResetTcpKeepAlive() {
	_jsii_.InvokeVoid(
		t,
		"resetTcpKeepAlive",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ResetTlsTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetTlsTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigIngressRuleOriginRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


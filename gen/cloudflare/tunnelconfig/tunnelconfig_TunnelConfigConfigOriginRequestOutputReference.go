package tunnelconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/tunnelconfig/internal"
)

type TunnelConfigConfigOriginRequestOutputReference interface {
	cdktf.ComplexObject
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
	HttpHostHeader() *string
	SetHttpHostHeader(val *string)
	HttpHostHeaderInput() *string
	InternalValue() *TunnelConfigConfigOriginRequest
	SetInternalValue(val *TunnelConfigConfigOriginRequest)
	IpRules() TunnelConfigConfigOriginRequestIpRulesList
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
	PutIpRules(value interface{})
	ResetBastionMode()
	ResetCaPool()
	ResetConnectTimeout()
	ResetDisableChunkedEncoding()
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

// The jsii proxy struct for TunnelConfigConfigOriginRequestOutputReference
type jsiiProxy_TunnelConfigConfigOriginRequestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) BastionMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bastionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) BastionModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bastionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) CaPool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) CaPoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ConnectTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ConnectTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) DisableChunkedEncoding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableChunkedEncoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) DisableChunkedEncodingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableChunkedEncodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) HttpHostHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpHostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) HttpHostHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpHostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) InternalValue() *TunnelConfigConfigOriginRequest {
	var returns *TunnelConfigConfigOriginRequest
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) IpRules() TunnelConfigConfigOriginRequestIpRulesList {
	var returns TunnelConfigConfigOriginRequestIpRulesList
	_jsii_.Get(
		j,
		"ipRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) IpRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) KeepAliveConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) KeepAliveConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) KeepAliveTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keepAliveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) KeepAliveTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keepAliveTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) NoHappyEyeballs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noHappyEyeballs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) NoHappyEyeballsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noHappyEyeballsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) NoTlsVerify() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noTlsVerify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) NoTlsVerifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noTlsVerifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) OriginServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originServerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) OriginServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originServerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ProxyAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ProxyAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ProxyPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"proxyPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ProxyPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"proxyPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ProxyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ProxyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) TcpKeepAlive() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tcpKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) TcpKeepAliveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tcpKeepAliveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) TlsTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) TlsTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsTimeoutInput",
		&returns,
	)
	return returns
}


func NewTunnelConfigConfigOriginRequestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TunnelConfigConfigOriginRequestOutputReference {
	_init_.Initialize()

	if err := validateNewTunnelConfigConfigOriginRequestOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TunnelConfigConfigOriginRequestOutputReference{}

	_jsii_.Create(
		"cloudflare.tunnelConfig.TunnelConfigConfigOriginRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTunnelConfigConfigOriginRequestOutputReference_Override(t TunnelConfigConfigOriginRequestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.tunnelConfig.TunnelConfigConfigOriginRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference)SetBastionMode(val interface{}) {
	if err := j.validateSetBastionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bastionMode",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference)SetCaPool(val *string) {
	if err := j.validateSetCaPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caPool",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference)SetConnectTimeout(val *string) {
	if err := j.validateSetConnectTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectTimeout",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference)SetDisableChunkedEncoding(val interface{}) {
	if err := j.validateSetDisableChunkedEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableChunkedEncoding",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference)SetHttpHostHeader(val *string) {
	if err := j.validateSetHttpHostHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpHostHeader",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference)SetInternalValue(val *TunnelConfigConfigOriginRequest) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference)SetKeepAliveConnections(val *float64) {
	if err := j.validateSetKeepAliveConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepAliveConnections",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference)SetKeepAliveTimeout(val *string) {
	if err := j.validateSetKeepAliveTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepAliveTimeout",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference)SetNoHappyEyeballs(val interface{}) {
	if err := j.validateSetNoHappyEyeballsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noHappyEyeballs",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference)SetNoTlsVerify(val interface{}) {
	if err := j.validateSetNoTlsVerifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noTlsVerify",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference)SetOriginServerName(val *string) {
	if err := j.validateSetOriginServerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originServerName",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference)SetProxyAddress(val *string) {
	if err := j.validateSetProxyAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyAddress",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference)SetProxyPort(val *float64) {
	if err := j.validateSetProxyPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyPort",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference)SetProxyType(val *string) {
	if err := j.validateSetProxyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyType",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference)SetTcpKeepAlive(val *string) {
	if err := j.validateSetTcpKeepAliveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpKeepAlive",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference)SetTlsTimeout(val *string) {
	if err := j.validateSetTlsTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsTimeout",
		val,
	)
}

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) PutIpRules(value interface{}) {
	if err := t.validatePutIpRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIpRules",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ResetBastionMode() {
	_jsii_.InvokeVoid(
		t,
		"resetBastionMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ResetCaPool() {
	_jsii_.InvokeVoid(
		t,
		"resetCaPool",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ResetConnectTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetConnectTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ResetDisableChunkedEncoding() {
	_jsii_.InvokeVoid(
		t,
		"resetDisableChunkedEncoding",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ResetHttpHostHeader() {
	_jsii_.InvokeVoid(
		t,
		"resetHttpHostHeader",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ResetIpRules() {
	_jsii_.InvokeVoid(
		t,
		"resetIpRules",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ResetKeepAliveConnections() {
	_jsii_.InvokeVoid(
		t,
		"resetKeepAliveConnections",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ResetKeepAliveTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetKeepAliveTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ResetNoHappyEyeballs() {
	_jsii_.InvokeVoid(
		t,
		"resetNoHappyEyeballs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ResetNoTlsVerify() {
	_jsii_.InvokeVoid(
		t,
		"resetNoTlsVerify",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ResetOriginServerName() {
	_jsii_.InvokeVoid(
		t,
		"resetOriginServerName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ResetProxyAddress() {
	_jsii_.InvokeVoid(
		t,
		"resetProxyAddress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ResetProxyPort() {
	_jsii_.InvokeVoid(
		t,
		"resetProxyPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ResetProxyType() {
	_jsii_.InvokeVoid(
		t,
		"resetProxyType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ResetTcpKeepAlive() {
	_jsii_.InvokeVoid(
		t,
		"resetTcpKeepAlive",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ResetTlsTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetTlsTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (t *jsiiProxy_TunnelConfigConfigOriginRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


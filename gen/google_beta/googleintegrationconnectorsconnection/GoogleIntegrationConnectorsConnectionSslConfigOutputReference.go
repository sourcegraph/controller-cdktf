package googleintegrationconnectorsconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleintegrationconnectorsconnection/internal"
)

type GoogleIntegrationConnectorsConnectionSslConfigOutputReference interface {
	cdktf.ComplexObject
	AdditionalVariable() GoogleIntegrationConnectorsConnectionSslConfigAdditionalVariableList
	AdditionalVariableInput() interface{}
	ClientCertificate() GoogleIntegrationConnectorsConnectionSslConfigClientCertificateOutputReference
	ClientCertificateInput() *GoogleIntegrationConnectorsConnectionSslConfigClientCertificate
	ClientCertType() *string
	SetClientCertType(val *string)
	ClientCertTypeInput() *string
	ClientPrivateKey() GoogleIntegrationConnectorsConnectionSslConfigClientPrivateKeyOutputReference
	ClientPrivateKeyInput() *GoogleIntegrationConnectorsConnectionSslConfigClientPrivateKey
	ClientPrivateKeyPass() GoogleIntegrationConnectorsConnectionSslConfigClientPrivateKeyPassOutputReference
	ClientPrivateKeyPassInput() *GoogleIntegrationConnectorsConnectionSslConfigClientPrivateKeyPass
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleIntegrationConnectorsConnectionSslConfig
	SetInternalValue(val *GoogleIntegrationConnectorsConnectionSslConfig)
	PrivateServerCertificate() GoogleIntegrationConnectorsConnectionSslConfigPrivateServerCertificateOutputReference
	PrivateServerCertificateInput() *GoogleIntegrationConnectorsConnectionSslConfigPrivateServerCertificate
	ServerCertType() *string
	SetServerCertType(val *string)
	ServerCertTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrustModel() *string
	SetTrustModel(val *string)
	TrustModelInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UseSsl() interface{}
	SetUseSsl(val interface{})
	UseSslInput() interface{}
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
	PutAdditionalVariable(value interface{})
	PutClientCertificate(value *GoogleIntegrationConnectorsConnectionSslConfigClientCertificate)
	PutClientPrivateKey(value *GoogleIntegrationConnectorsConnectionSslConfigClientPrivateKey)
	PutClientPrivateKeyPass(value *GoogleIntegrationConnectorsConnectionSslConfigClientPrivateKeyPass)
	PutPrivateServerCertificate(value *GoogleIntegrationConnectorsConnectionSslConfigPrivateServerCertificate)
	ResetAdditionalVariable()
	ResetClientCertificate()
	ResetClientCertType()
	ResetClientPrivateKey()
	ResetClientPrivateKeyPass()
	ResetPrivateServerCertificate()
	ResetServerCertType()
	ResetTrustModel()
	ResetUseSsl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIntegrationConnectorsConnectionSslConfigOutputReference
type jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) AdditionalVariable() GoogleIntegrationConnectorsConnectionSslConfigAdditionalVariableList {
	var returns GoogleIntegrationConnectorsConnectionSslConfigAdditionalVariableList
	_jsii_.Get(
		j,
		"additionalVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) AdditionalVariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ClientCertificate() GoogleIntegrationConnectorsConnectionSslConfigClientCertificateOutputReference {
	var returns GoogleIntegrationConnectorsConnectionSslConfigClientCertificateOutputReference
	_jsii_.Get(
		j,
		"clientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ClientCertificateInput() *GoogleIntegrationConnectorsConnectionSslConfigClientCertificate {
	var returns *GoogleIntegrationConnectorsConnectionSslConfigClientCertificate
	_jsii_.Get(
		j,
		"clientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ClientCertType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ClientCertTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ClientPrivateKey() GoogleIntegrationConnectorsConnectionSslConfigClientPrivateKeyOutputReference {
	var returns GoogleIntegrationConnectorsConnectionSslConfigClientPrivateKeyOutputReference
	_jsii_.Get(
		j,
		"clientPrivateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ClientPrivateKeyInput() *GoogleIntegrationConnectorsConnectionSslConfigClientPrivateKey {
	var returns *GoogleIntegrationConnectorsConnectionSslConfigClientPrivateKey
	_jsii_.Get(
		j,
		"clientPrivateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ClientPrivateKeyPass() GoogleIntegrationConnectorsConnectionSslConfigClientPrivateKeyPassOutputReference {
	var returns GoogleIntegrationConnectorsConnectionSslConfigClientPrivateKeyPassOutputReference
	_jsii_.Get(
		j,
		"clientPrivateKeyPass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ClientPrivateKeyPassInput() *GoogleIntegrationConnectorsConnectionSslConfigClientPrivateKeyPass {
	var returns *GoogleIntegrationConnectorsConnectionSslConfigClientPrivateKeyPass
	_jsii_.Get(
		j,
		"clientPrivateKeyPassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) InternalValue() *GoogleIntegrationConnectorsConnectionSslConfig {
	var returns *GoogleIntegrationConnectorsConnectionSslConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) PrivateServerCertificate() GoogleIntegrationConnectorsConnectionSslConfigPrivateServerCertificateOutputReference {
	var returns GoogleIntegrationConnectorsConnectionSslConfigPrivateServerCertificateOutputReference
	_jsii_.Get(
		j,
		"privateServerCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) PrivateServerCertificateInput() *GoogleIntegrationConnectorsConnectionSslConfigPrivateServerCertificate {
	var returns *GoogleIntegrationConnectorsConnectionSslConfigPrivateServerCertificate
	_jsii_.Get(
		j,
		"privateServerCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ServerCertType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCertType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ServerCertTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCertTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) TrustModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) TrustModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) UseSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) UseSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useSslInput",
		&returns,
	)
	return returns
}


func NewGoogleIntegrationConnectorsConnectionSslConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleIntegrationConnectorsConnectionSslConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIntegrationConnectorsConnectionSslConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIntegrationConnectorsConnection.GoogleIntegrationConnectorsConnectionSslConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleIntegrationConnectorsConnectionSslConfigOutputReference_Override(g GoogleIntegrationConnectorsConnectionSslConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIntegrationConnectorsConnection.GoogleIntegrationConnectorsConnectionSslConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference)SetClientCertType(val *string) {
	if err := j.validateSetClientCertTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertType",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference)SetInternalValue(val *GoogleIntegrationConnectorsConnectionSslConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference)SetServerCertType(val *string) {
	if err := j.validateSetServerCertTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverCertType",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference)SetTrustModel(val *string) {
	if err := j.validateSetTrustModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustModel",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference)SetUseSsl(val interface{}) {
	if err := j.validateSetUseSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSsl",
		val,
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) PutAdditionalVariable(value interface{}) {
	if err := g.validatePutAdditionalVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdditionalVariable",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) PutClientCertificate(value *GoogleIntegrationConnectorsConnectionSslConfigClientCertificate) {
	if err := g.validatePutClientCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClientCertificate",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) PutClientPrivateKey(value *GoogleIntegrationConnectorsConnectionSslConfigClientPrivateKey) {
	if err := g.validatePutClientPrivateKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClientPrivateKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) PutClientPrivateKeyPass(value *GoogleIntegrationConnectorsConnectionSslConfigClientPrivateKeyPass) {
	if err := g.validatePutClientPrivateKeyPassParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClientPrivateKeyPass",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) PutPrivateServerCertificate(value *GoogleIntegrationConnectorsConnectionSslConfigPrivateServerCertificate) {
	if err := g.validatePutPrivateServerCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPrivateServerCertificate",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ResetAdditionalVariable() {
	_jsii_.InvokeVoid(
		g,
		"resetAdditionalVariable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ResetClientCertificate() {
	_jsii_.InvokeVoid(
		g,
		"resetClientCertificate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ResetClientCertType() {
	_jsii_.InvokeVoid(
		g,
		"resetClientCertType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ResetClientPrivateKey() {
	_jsii_.InvokeVoid(
		g,
		"resetClientPrivateKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ResetClientPrivateKeyPass() {
	_jsii_.InvokeVoid(
		g,
		"resetClientPrivateKeyPass",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ResetPrivateServerCertificate() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateServerCertificate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ResetServerCertType() {
	_jsii_.InvokeVoid(
		g,
		"resetServerCertType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ResetTrustModel() {
	_jsii_.InvokeVoid(
		g,
		"resetTrustModel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ResetUseSsl() {
	_jsii_.InvokeVoid(
		g,
		"resetUseSsl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionSslConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package googledatastreamconnectionprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatastreamconnectionprofile/internal"
)

type GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference interface {
	cdktf.ComplexObject
	CaCertificate() *string
	SetCaCertificate(val *string)
	CaCertificateInput() *string
	CaCertificateSet() cdktf.IResolvable
	ClientCertificate() *string
	SetClientCertificate(val *string)
	ClientCertificateInput() *string
	ClientCertificateSet() cdktf.IResolvable
	ClientKey() *string
	SetClientKey(val *string)
	ClientKeyInput() *string
	ClientKeySet() cdktf.IResolvable
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
	InternalValue() *GoogleDatastreamConnectionProfileMongodbProfileSslConfig
	SetInternalValue(val *GoogleDatastreamConnectionProfileMongodbProfileSslConfig)
	SecretManagerStoredClientKey() *string
	SetSecretManagerStoredClientKey(val *string)
	SecretManagerStoredClientKeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetCaCertificate()
	ResetClientCertificate()
	ResetClientKey()
	ResetSecretManagerStoredClientKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference
type jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) CaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) CaCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) CaCertificateSet() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"caCertificateSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) ClientCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) ClientCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) ClientCertificateSet() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"clientCertificateSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) ClientKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) ClientKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) ClientKeySet() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"clientKeySet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) InternalValue() *GoogleDatastreamConnectionProfileMongodbProfileSslConfig {
	var returns *GoogleDatastreamConnectionProfileMongodbProfileSslConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) SecretManagerStoredClientKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretManagerStoredClientKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) SecretManagerStoredClientKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretManagerStoredClientKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatastreamConnectionProfile.GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference_Override(g GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatastreamConnectionProfile.GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference)SetCaCertificate(val *string) {
	if err := j.validateSetCaCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificate",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference)SetClientCertificate(val *string) {
	if err := j.validateSetClientCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificate",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference)SetClientKey(val *string) {
	if err := j.validateSetClientKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientKey",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference)SetInternalValue(val *GoogleDatastreamConnectionProfileMongodbProfileSslConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference)SetSecretManagerStoredClientKey(val *string) {
	if err := j.validateSetSecretManagerStoredClientKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretManagerStoredClientKey",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) ResetCaCertificate() {
	_jsii_.InvokeVoid(
		g,
		"resetCaCertificate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) ResetClientCertificate() {
	_jsii_.InvokeVoid(
		g,
		"resetClientCertificate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) ResetClientKey() {
	_jsii_.InvokeVoid(
		g,
		"resetClientKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) ResetSecretManagerStoredClientKey() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretManagerStoredClientKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


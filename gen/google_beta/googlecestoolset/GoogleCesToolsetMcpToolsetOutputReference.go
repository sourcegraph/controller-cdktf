package googlecestoolset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecestoolset/internal"
)

type GoogleCesToolsetMcpToolsetOutputReference interface {
	cdktf.ComplexObject
	ApiAuthentication() GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference
	ApiAuthenticationInput() *GoogleCesToolsetMcpToolsetApiAuthentication
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
	CustomHeaders() *map[string]*string
	SetCustomHeaders(val *map[string]*string)
	CustomHeadersInput() *map[string]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCesToolsetMcpToolset
	SetInternalValue(val *GoogleCesToolsetMcpToolset)
	ServerAddress() *string
	SetServerAddress(val *string)
	ServerAddressInput() *string
	ServiceDirectoryConfig() GoogleCesToolsetMcpToolsetServiceDirectoryConfigOutputReference
	ServiceDirectoryConfigInput() *GoogleCesToolsetMcpToolsetServiceDirectoryConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsConfig() GoogleCesToolsetMcpToolsetTlsConfigOutputReference
	TlsConfigInput() *GoogleCesToolsetMcpToolsetTlsConfig
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
	PutApiAuthentication(value *GoogleCesToolsetMcpToolsetApiAuthentication)
	PutServiceDirectoryConfig(value *GoogleCesToolsetMcpToolsetServiceDirectoryConfig)
	PutTlsConfig(value *GoogleCesToolsetMcpToolsetTlsConfig)
	ResetApiAuthentication()
	ResetCustomHeaders()
	ResetServiceDirectoryConfig()
	ResetTlsConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesToolsetMcpToolsetOutputReference
type jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) ApiAuthentication() GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference {
	var returns GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference
	_jsii_.Get(
		j,
		"apiAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) ApiAuthenticationInput() *GoogleCesToolsetMcpToolsetApiAuthentication {
	var returns *GoogleCesToolsetMcpToolsetApiAuthentication
	_jsii_.Get(
		j,
		"apiAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) CustomHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) CustomHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) InternalValue() *GoogleCesToolsetMcpToolset {
	var returns *GoogleCesToolsetMcpToolset
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) ServerAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) ServerAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) ServiceDirectoryConfig() GoogleCesToolsetMcpToolsetServiceDirectoryConfigOutputReference {
	var returns GoogleCesToolsetMcpToolsetServiceDirectoryConfigOutputReference
	_jsii_.Get(
		j,
		"serviceDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) ServiceDirectoryConfigInput() *GoogleCesToolsetMcpToolsetServiceDirectoryConfig {
	var returns *GoogleCesToolsetMcpToolsetServiceDirectoryConfig
	_jsii_.Get(
		j,
		"serviceDirectoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) TlsConfig() GoogleCesToolsetMcpToolsetTlsConfigOutputReference {
	var returns GoogleCesToolsetMcpToolsetTlsConfigOutputReference
	_jsii_.Get(
		j,
		"tlsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) TlsConfigInput() *GoogleCesToolsetMcpToolsetTlsConfig {
	var returns *GoogleCesToolsetMcpToolsetTlsConfig
	_jsii_.Get(
		j,
		"tlsConfigInput",
		&returns,
	)
	return returns
}


func NewGoogleCesToolsetMcpToolsetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCesToolsetMcpToolsetOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesToolsetMcpToolsetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesToolset.GoogleCesToolsetMcpToolsetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesToolsetMcpToolsetOutputReference_Override(g GoogleCesToolsetMcpToolsetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesToolset.GoogleCesToolsetMcpToolsetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference)SetCustomHeaders(val *map[string]*string) {
	if err := j.validateSetCustomHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customHeaders",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference)SetInternalValue(val *GoogleCesToolsetMcpToolset) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference)SetServerAddress(val *string) {
	if err := j.validateSetServerAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverAddress",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) PutApiAuthentication(value *GoogleCesToolsetMcpToolsetApiAuthentication) {
	if err := g.validatePutApiAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApiAuthentication",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) PutServiceDirectoryConfig(value *GoogleCesToolsetMcpToolsetServiceDirectoryConfig) {
	if err := g.validatePutServiceDirectoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceDirectoryConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) PutTlsConfig(value *GoogleCesToolsetMcpToolsetTlsConfig) {
	if err := g.validatePutTlsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTlsConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) ResetApiAuthentication() {
	_jsii_.InvokeVoid(
		g,
		"resetApiAuthentication",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) ResetCustomHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) ResetServiceDirectoryConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceDirectoryConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) ResetTlsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetTlsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


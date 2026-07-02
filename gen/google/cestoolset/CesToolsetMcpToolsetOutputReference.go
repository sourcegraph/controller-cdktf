package cestoolset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cestoolset/internal"
)

type CesToolsetMcpToolsetOutputReference interface {
	cdktf.ComplexObject
	ApiAuthentication() CesToolsetMcpToolsetApiAuthenticationOutputReference
	ApiAuthenticationInput() *CesToolsetMcpToolsetApiAuthentication
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	InternalValue() *CesToolsetMcpToolset
	SetInternalValue(val *CesToolsetMcpToolset)
	ServerAddress() *string
	SetServerAddress(val *string)
	ServerAddressInput() *string
	ServiceDirectoryConfig() CesToolsetMcpToolsetServiceDirectoryConfigOutputReference
	ServiceDirectoryConfigInput() *CesToolsetMcpToolsetServiceDirectoryConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsConfig() CesToolsetMcpToolsetTlsConfigOutputReference
	TlsConfigInput() *CesToolsetMcpToolsetTlsConfig
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutApiAuthentication(value *CesToolsetMcpToolsetApiAuthentication)
	PutServiceDirectoryConfig(value *CesToolsetMcpToolsetServiceDirectoryConfig)
	PutTlsConfig(value *CesToolsetMcpToolsetTlsConfig)
	ResetApiAuthentication()
	ResetCustomHeaders()
	ResetServiceDirectoryConfig()
	ResetTlsConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesToolsetMcpToolsetOutputReference
type jsiiProxy_CesToolsetMcpToolsetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) ApiAuthentication() CesToolsetMcpToolsetApiAuthenticationOutputReference {
	var returns CesToolsetMcpToolsetApiAuthenticationOutputReference
	_jsii_.Get(
		j,
		"apiAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) ApiAuthenticationInput() *CesToolsetMcpToolsetApiAuthentication {
	var returns *CesToolsetMcpToolsetApiAuthentication
	_jsii_.Get(
		j,
		"apiAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) CustomHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) CustomHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) InternalValue() *CesToolsetMcpToolset {
	var returns *CesToolsetMcpToolset
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) ServerAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) ServerAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) ServiceDirectoryConfig() CesToolsetMcpToolsetServiceDirectoryConfigOutputReference {
	var returns CesToolsetMcpToolsetServiceDirectoryConfigOutputReference
	_jsii_.Get(
		j,
		"serviceDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) ServiceDirectoryConfigInput() *CesToolsetMcpToolsetServiceDirectoryConfig {
	var returns *CesToolsetMcpToolsetServiceDirectoryConfig
	_jsii_.Get(
		j,
		"serviceDirectoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) TlsConfig() CesToolsetMcpToolsetTlsConfigOutputReference {
	var returns CesToolsetMcpToolsetTlsConfigOutputReference
	_jsii_.Get(
		j,
		"tlsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) TlsConfigInput() *CesToolsetMcpToolsetTlsConfig {
	var returns *CesToolsetMcpToolsetTlsConfig
	_jsii_.Get(
		j,
		"tlsConfigInput",
		&returns,
	)
	return returns
}

func NewCesToolsetMcpToolsetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CesToolsetMcpToolsetOutputReference {
	_init_.Initialize()

	if err := validateNewCesToolsetMcpToolsetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesToolsetMcpToolsetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cesToolset.CesToolsetMcpToolsetOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesToolsetMcpToolsetOutputReference_Override(c CesToolsetMcpToolsetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesToolset.CesToolsetMcpToolsetOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) SetCustomHeaders(val *map[string]*string) {
	if err := j.validateSetCustomHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customHeaders",
		val,
	)
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) SetInternalValue(val *CesToolsetMcpToolset) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) SetServerAddress(val *string) {
	if err := j.validateSetServerAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverAddress",
		val,
	)
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesToolsetMcpToolsetOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesToolsetMcpToolsetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetMcpToolsetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetMcpToolsetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetMcpToolsetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetMcpToolsetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetMcpToolsetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetMcpToolsetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetMcpToolsetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetMcpToolsetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetMcpToolsetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetMcpToolsetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetMcpToolsetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetMcpToolsetOutputReference) PutApiAuthentication(value *CesToolsetMcpToolsetApiAuthentication) {
	if err := c.validatePutApiAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putApiAuthentication",
		[]any{value},
	)
}

func (c *jsiiProxy_CesToolsetMcpToolsetOutputReference) PutServiceDirectoryConfig(value *CesToolsetMcpToolsetServiceDirectoryConfig) {
	if err := c.validatePutServiceDirectoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServiceDirectoryConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_CesToolsetMcpToolsetOutputReference) PutTlsConfig(value *CesToolsetMcpToolsetTlsConfig) {
	if err := c.validatePutTlsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTlsConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_CesToolsetMcpToolsetOutputReference) ResetApiAuthentication() {
	_jsii_.InvokeVoid(
		c,
		"resetApiAuthentication",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetMcpToolsetOutputReference) ResetCustomHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetMcpToolsetOutputReference) ResetServiceDirectoryConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceDirectoryConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetMcpToolsetOutputReference) ResetTlsConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetTlsConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetMcpToolsetOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		c,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetMcpToolsetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

package cestoolset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cestoolset/internal"
)

type CesToolsetOpenApiToolsetOutputReference interface {
	cdktf.ComplexObject
	ApiAuthentication() CesToolsetOpenApiToolsetApiAuthenticationOutputReference
	ApiAuthenticationInput() *CesToolsetOpenApiToolsetApiAuthentication
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
	IgnoreUnknownFields() interface{}
	SetIgnoreUnknownFields(val interface{})
	IgnoreUnknownFieldsInput() interface{}
	InternalValue() *CesToolsetOpenApiToolset
	SetInternalValue(val *CesToolsetOpenApiToolset)
	OpenApiSchema() *string
	SetOpenApiSchema(val *string)
	OpenApiSchemaInput() *string
	ServiceDirectoryConfig() CesToolsetOpenApiToolsetServiceDirectoryConfigOutputReference
	ServiceDirectoryConfigInput() *CesToolsetOpenApiToolsetServiceDirectoryConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsConfig() CesToolsetOpenApiToolsetTlsConfigOutputReference
	TlsConfigInput() *CesToolsetOpenApiToolsetTlsConfig
	Url() *string
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
	PutApiAuthentication(value *CesToolsetOpenApiToolsetApiAuthentication)
	PutServiceDirectoryConfig(value *CesToolsetOpenApiToolsetServiceDirectoryConfig)
	PutTlsConfig(value *CesToolsetOpenApiToolsetTlsConfig)
	ResetApiAuthentication()
	ResetIgnoreUnknownFields()
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

// The jsii proxy struct for CesToolsetOpenApiToolsetOutputReference
type jsiiProxy_CesToolsetOpenApiToolsetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) ApiAuthentication() CesToolsetOpenApiToolsetApiAuthenticationOutputReference {
	var returns CesToolsetOpenApiToolsetApiAuthenticationOutputReference
	_jsii_.Get(
		j,
		"apiAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) ApiAuthenticationInput() *CesToolsetOpenApiToolsetApiAuthentication {
	var returns *CesToolsetOpenApiToolsetApiAuthentication
	_jsii_.Get(
		j,
		"apiAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) IgnoreUnknownFields() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreUnknownFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) IgnoreUnknownFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreUnknownFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) InternalValue() *CesToolsetOpenApiToolset {
	var returns *CesToolsetOpenApiToolset
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) OpenApiSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openApiSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) OpenApiSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openApiSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) ServiceDirectoryConfig() CesToolsetOpenApiToolsetServiceDirectoryConfigOutputReference {
	var returns CesToolsetOpenApiToolsetServiceDirectoryConfigOutputReference
	_jsii_.Get(
		j,
		"serviceDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) ServiceDirectoryConfigInput() *CesToolsetOpenApiToolsetServiceDirectoryConfig {
	var returns *CesToolsetOpenApiToolsetServiceDirectoryConfig
	_jsii_.Get(
		j,
		"serviceDirectoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) TlsConfig() CesToolsetOpenApiToolsetTlsConfigOutputReference {
	var returns CesToolsetOpenApiToolsetTlsConfigOutputReference
	_jsii_.Get(
		j,
		"tlsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) TlsConfigInput() *CesToolsetOpenApiToolsetTlsConfig {
	var returns *CesToolsetOpenApiToolsetTlsConfig
	_jsii_.Get(
		j,
		"tlsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


func NewCesToolsetOpenApiToolsetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CesToolsetOpenApiToolsetOutputReference {
	_init_.Initialize()

	if err := validateNewCesToolsetOpenApiToolsetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesToolsetOpenApiToolsetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cesToolset.CesToolsetOpenApiToolsetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesToolsetOpenApiToolsetOutputReference_Override(c CesToolsetOpenApiToolsetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesToolset.CesToolsetOpenApiToolsetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference)SetIgnoreUnknownFields(val interface{}) {
	if err := j.validateSetIgnoreUnknownFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreUnknownFields",
		val,
	)
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference)SetInternalValue(val *CesToolsetOpenApiToolset) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference)SetOpenApiSchema(val *string) {
	if err := j.validateSetOpenApiSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openApiSchema",
		val,
	)
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) PutApiAuthentication(value *CesToolsetOpenApiToolsetApiAuthentication) {
	if err := c.validatePutApiAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putApiAuthentication",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) PutServiceDirectoryConfig(value *CesToolsetOpenApiToolsetServiceDirectoryConfig) {
	if err := c.validatePutServiceDirectoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServiceDirectoryConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) PutTlsConfig(value *CesToolsetOpenApiToolsetTlsConfig) {
	if err := c.validatePutTlsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTlsConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) ResetApiAuthentication() {
	_jsii_.InvokeVoid(
		c,
		"resetApiAuthentication",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) ResetIgnoreUnknownFields() {
	_jsii_.InvokeVoid(
		c,
		"resetIgnoreUnknownFields",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) ResetServiceDirectoryConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceDirectoryConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) ResetTlsConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetTlsConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package cestoolset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cestoolset/internal"
)

type CesToolsetMcpToolsetApiAuthenticationOutputReference interface {
	cdktf.ComplexObject
	ApiKeyConfig() CesToolsetMcpToolsetApiAuthenticationApiKeyConfigOutputReference
	ApiKeyConfigInput() *CesToolsetMcpToolsetApiAuthenticationApiKeyConfig
	BearerTokenConfig() CesToolsetMcpToolsetApiAuthenticationBearerTokenConfigOutputReference
	BearerTokenConfigInput() *CesToolsetMcpToolsetApiAuthenticationBearerTokenConfig
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
	InternalValue() *CesToolsetMcpToolsetApiAuthentication
	SetInternalValue(val *CesToolsetMcpToolsetApiAuthentication)
	OauthConfig() CesToolsetMcpToolsetApiAuthenticationOauthConfigOutputReference
	OauthConfigInput() *CesToolsetMcpToolsetApiAuthenticationOauthConfig
	ServiceAccountAuthConfig() CesToolsetMcpToolsetApiAuthenticationServiceAccountAuthConfigOutputReference
	ServiceAccountAuthConfigInput() *CesToolsetMcpToolsetApiAuthenticationServiceAccountAuthConfig
	ServiceAgentIdTokenAuthConfig() CesToolsetMcpToolsetApiAuthenticationServiceAgentIdTokenAuthConfigOutputReference
	ServiceAgentIdTokenAuthConfigInput() *CesToolsetMcpToolsetApiAuthenticationServiceAgentIdTokenAuthConfig
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
	PutApiKeyConfig(value *CesToolsetMcpToolsetApiAuthenticationApiKeyConfig)
	PutBearerTokenConfig(value *CesToolsetMcpToolsetApiAuthenticationBearerTokenConfig)
	PutOauthConfig(value *CesToolsetMcpToolsetApiAuthenticationOauthConfig)
	PutServiceAccountAuthConfig(value *CesToolsetMcpToolsetApiAuthenticationServiceAccountAuthConfig)
	PutServiceAgentIdTokenAuthConfig(value *CesToolsetMcpToolsetApiAuthenticationServiceAgentIdTokenAuthConfig)
	ResetApiKeyConfig()
	ResetBearerTokenConfig()
	ResetOauthConfig()
	ResetServiceAccountAuthConfig()
	ResetServiceAgentIdTokenAuthConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesToolsetMcpToolsetApiAuthenticationOutputReference
type jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) ApiKeyConfig() CesToolsetMcpToolsetApiAuthenticationApiKeyConfigOutputReference {
	var returns CesToolsetMcpToolsetApiAuthenticationApiKeyConfigOutputReference
	_jsii_.Get(
		j,
		"apiKeyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) ApiKeyConfigInput() *CesToolsetMcpToolsetApiAuthenticationApiKeyConfig {
	var returns *CesToolsetMcpToolsetApiAuthenticationApiKeyConfig
	_jsii_.Get(
		j,
		"apiKeyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) BearerTokenConfig() CesToolsetMcpToolsetApiAuthenticationBearerTokenConfigOutputReference {
	var returns CesToolsetMcpToolsetApiAuthenticationBearerTokenConfigOutputReference
	_jsii_.Get(
		j,
		"bearerTokenConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) BearerTokenConfigInput() *CesToolsetMcpToolsetApiAuthenticationBearerTokenConfig {
	var returns *CesToolsetMcpToolsetApiAuthenticationBearerTokenConfig
	_jsii_.Get(
		j,
		"bearerTokenConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) InternalValue() *CesToolsetMcpToolsetApiAuthentication {
	var returns *CesToolsetMcpToolsetApiAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) OauthConfig() CesToolsetMcpToolsetApiAuthenticationOauthConfigOutputReference {
	var returns CesToolsetMcpToolsetApiAuthenticationOauthConfigOutputReference
	_jsii_.Get(
		j,
		"oauthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) OauthConfigInput() *CesToolsetMcpToolsetApiAuthenticationOauthConfig {
	var returns *CesToolsetMcpToolsetApiAuthenticationOauthConfig
	_jsii_.Get(
		j,
		"oauthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) ServiceAccountAuthConfig() CesToolsetMcpToolsetApiAuthenticationServiceAccountAuthConfigOutputReference {
	var returns CesToolsetMcpToolsetApiAuthenticationServiceAccountAuthConfigOutputReference
	_jsii_.Get(
		j,
		"serviceAccountAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) ServiceAccountAuthConfigInput() *CesToolsetMcpToolsetApiAuthenticationServiceAccountAuthConfig {
	var returns *CesToolsetMcpToolsetApiAuthenticationServiceAccountAuthConfig
	_jsii_.Get(
		j,
		"serviceAccountAuthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) ServiceAgentIdTokenAuthConfig() CesToolsetMcpToolsetApiAuthenticationServiceAgentIdTokenAuthConfigOutputReference {
	var returns CesToolsetMcpToolsetApiAuthenticationServiceAgentIdTokenAuthConfigOutputReference
	_jsii_.Get(
		j,
		"serviceAgentIdTokenAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) ServiceAgentIdTokenAuthConfigInput() *CesToolsetMcpToolsetApiAuthenticationServiceAgentIdTokenAuthConfig {
	var returns *CesToolsetMcpToolsetApiAuthenticationServiceAgentIdTokenAuthConfig
	_jsii_.Get(
		j,
		"serviceAgentIdTokenAuthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCesToolsetMcpToolsetApiAuthenticationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CesToolsetMcpToolsetApiAuthenticationOutputReference {
	_init_.Initialize()

	if err := validateNewCesToolsetMcpToolsetApiAuthenticationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cesToolset.CesToolsetMcpToolsetApiAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesToolsetMcpToolsetApiAuthenticationOutputReference_Override(c CesToolsetMcpToolsetApiAuthenticationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesToolset.CesToolsetMcpToolsetApiAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference)SetInternalValue(val *CesToolsetMcpToolsetApiAuthentication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) PutApiKeyConfig(value *CesToolsetMcpToolsetApiAuthenticationApiKeyConfig) {
	if err := c.validatePutApiKeyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putApiKeyConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) PutBearerTokenConfig(value *CesToolsetMcpToolsetApiAuthenticationBearerTokenConfig) {
	if err := c.validatePutBearerTokenConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBearerTokenConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) PutOauthConfig(value *CesToolsetMcpToolsetApiAuthenticationOauthConfig) {
	if err := c.validatePutOauthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOauthConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) PutServiceAccountAuthConfig(value *CesToolsetMcpToolsetApiAuthenticationServiceAccountAuthConfig) {
	if err := c.validatePutServiceAccountAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServiceAccountAuthConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) PutServiceAgentIdTokenAuthConfig(value *CesToolsetMcpToolsetApiAuthenticationServiceAgentIdTokenAuthConfig) {
	if err := c.validatePutServiceAgentIdTokenAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServiceAgentIdTokenAuthConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) ResetApiKeyConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetApiKeyConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) ResetBearerTokenConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetBearerTokenConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) ResetOauthConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetOauthConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) ResetServiceAccountAuthConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAccountAuthConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) ResetServiceAgentIdTokenAuthConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAgentIdTokenAuthConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CesToolsetMcpToolsetApiAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


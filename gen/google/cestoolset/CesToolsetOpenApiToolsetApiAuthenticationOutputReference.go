package cestoolset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/cestoolset/internal"
)

type CesToolsetOpenApiToolsetApiAuthenticationOutputReference interface {
	cdktn.ComplexObject
	ApiKeyConfig() CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference
	ApiKeyConfigInput() *CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfig
	BearerTokenConfig() CesToolsetOpenApiToolsetApiAuthenticationBearerTokenConfigOutputReference
	BearerTokenConfigInput() *CesToolsetOpenApiToolsetApiAuthenticationBearerTokenConfig
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
	InternalValue() *CesToolsetOpenApiToolsetApiAuthentication
	SetInternalValue(val *CesToolsetOpenApiToolsetApiAuthentication)
	OauthConfig() CesToolsetOpenApiToolsetApiAuthenticationOauthConfigOutputReference
	OauthConfigInput() *CesToolsetOpenApiToolsetApiAuthenticationOauthConfig
	ServiceAccountAuthConfig() CesToolsetOpenApiToolsetApiAuthenticationServiceAccountAuthConfigOutputReference
	ServiceAccountAuthConfigInput() *CesToolsetOpenApiToolsetApiAuthenticationServiceAccountAuthConfig
	ServiceAgentIdTokenAuthConfig() CesToolsetOpenApiToolsetApiAuthenticationServiceAgentIdTokenAuthConfigOutputReference
	ServiceAgentIdTokenAuthConfigInput() *CesToolsetOpenApiToolsetApiAuthenticationServiceAgentIdTokenAuthConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutApiKeyConfig(value *CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfig)
	PutBearerTokenConfig(value *CesToolsetOpenApiToolsetApiAuthenticationBearerTokenConfig)
	PutOauthConfig(value *CesToolsetOpenApiToolsetApiAuthenticationOauthConfig)
	PutServiceAccountAuthConfig(value *CesToolsetOpenApiToolsetApiAuthenticationServiceAccountAuthConfig)
	PutServiceAgentIdTokenAuthConfig(value *CesToolsetOpenApiToolsetApiAuthenticationServiceAgentIdTokenAuthConfig)
	ResetApiKeyConfig()
	ResetBearerTokenConfig()
	ResetOauthConfig()
	ResetServiceAccountAuthConfig()
	ResetServiceAgentIdTokenAuthConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesToolsetOpenApiToolsetApiAuthenticationOutputReference
type jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) ApiKeyConfig() CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference {
	var returns CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference
	_jsii_.Get(
		j,
		"apiKeyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) ApiKeyConfigInput() *CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfig {
	var returns *CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfig
	_jsii_.Get(
		j,
		"apiKeyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) BearerTokenConfig() CesToolsetOpenApiToolsetApiAuthenticationBearerTokenConfigOutputReference {
	var returns CesToolsetOpenApiToolsetApiAuthenticationBearerTokenConfigOutputReference
	_jsii_.Get(
		j,
		"bearerTokenConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) BearerTokenConfigInput() *CesToolsetOpenApiToolsetApiAuthenticationBearerTokenConfig {
	var returns *CesToolsetOpenApiToolsetApiAuthenticationBearerTokenConfig
	_jsii_.Get(
		j,
		"bearerTokenConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) InternalValue() *CesToolsetOpenApiToolsetApiAuthentication {
	var returns *CesToolsetOpenApiToolsetApiAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) OauthConfig() CesToolsetOpenApiToolsetApiAuthenticationOauthConfigOutputReference {
	var returns CesToolsetOpenApiToolsetApiAuthenticationOauthConfigOutputReference
	_jsii_.Get(
		j,
		"oauthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) OauthConfigInput() *CesToolsetOpenApiToolsetApiAuthenticationOauthConfig {
	var returns *CesToolsetOpenApiToolsetApiAuthenticationOauthConfig
	_jsii_.Get(
		j,
		"oauthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) ServiceAccountAuthConfig() CesToolsetOpenApiToolsetApiAuthenticationServiceAccountAuthConfigOutputReference {
	var returns CesToolsetOpenApiToolsetApiAuthenticationServiceAccountAuthConfigOutputReference
	_jsii_.Get(
		j,
		"serviceAccountAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) ServiceAccountAuthConfigInput() *CesToolsetOpenApiToolsetApiAuthenticationServiceAccountAuthConfig {
	var returns *CesToolsetOpenApiToolsetApiAuthenticationServiceAccountAuthConfig
	_jsii_.Get(
		j,
		"serviceAccountAuthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) ServiceAgentIdTokenAuthConfig() CesToolsetOpenApiToolsetApiAuthenticationServiceAgentIdTokenAuthConfigOutputReference {
	var returns CesToolsetOpenApiToolsetApiAuthenticationServiceAgentIdTokenAuthConfigOutputReference
	_jsii_.Get(
		j,
		"serviceAgentIdTokenAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) ServiceAgentIdTokenAuthConfigInput() *CesToolsetOpenApiToolsetApiAuthenticationServiceAgentIdTokenAuthConfig {
	var returns *CesToolsetOpenApiToolsetApiAuthenticationServiceAgentIdTokenAuthConfig
	_jsii_.Get(
		j,
		"serviceAgentIdTokenAuthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCesToolsetOpenApiToolsetApiAuthenticationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CesToolsetOpenApiToolsetApiAuthenticationOutputReference {
	_init_.Initialize()

	if err := validateNewCesToolsetOpenApiToolsetApiAuthenticationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cesToolset.CesToolsetOpenApiToolsetApiAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesToolsetOpenApiToolsetApiAuthenticationOutputReference_Override(c CesToolsetOpenApiToolsetApiAuthenticationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cesToolset.CesToolsetOpenApiToolsetApiAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference)SetInternalValue(val *CesToolsetOpenApiToolsetApiAuthentication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) PutApiKeyConfig(value *CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfig) {
	if err := c.validatePutApiKeyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putApiKeyConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) PutBearerTokenConfig(value *CesToolsetOpenApiToolsetApiAuthenticationBearerTokenConfig) {
	if err := c.validatePutBearerTokenConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBearerTokenConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) PutOauthConfig(value *CesToolsetOpenApiToolsetApiAuthenticationOauthConfig) {
	if err := c.validatePutOauthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOauthConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) PutServiceAccountAuthConfig(value *CesToolsetOpenApiToolsetApiAuthenticationServiceAccountAuthConfig) {
	if err := c.validatePutServiceAccountAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServiceAccountAuthConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) PutServiceAgentIdTokenAuthConfig(value *CesToolsetOpenApiToolsetApiAuthenticationServiceAgentIdTokenAuthConfig) {
	if err := c.validatePutServiceAgentIdTokenAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServiceAgentIdTokenAuthConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) ResetApiKeyConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetApiKeyConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) ResetBearerTokenConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetBearerTokenConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) ResetOauthConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetOauthConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) ResetServiceAccountAuthConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAccountAuthConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) ResetServiceAgentIdTokenAuthConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAgentIdTokenAuthConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetOpenApiToolsetApiAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


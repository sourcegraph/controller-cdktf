package googlecestoolset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecestoolset/internal"
)

type GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference interface {
	cdktn.ComplexObject
	ApiKeyConfig() GoogleCesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference
	ApiKeyConfigInput() *GoogleCesToolsetOpenApiToolsetApiAuthenticationApiKeyConfig
	BearerTokenConfig() GoogleCesToolsetOpenApiToolsetApiAuthenticationBearerTokenConfigOutputReference
	BearerTokenConfigInput() *GoogleCesToolsetOpenApiToolsetApiAuthenticationBearerTokenConfig
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
	InternalValue() *GoogleCesToolsetOpenApiToolsetApiAuthentication
	SetInternalValue(val *GoogleCesToolsetOpenApiToolsetApiAuthentication)
	OauthConfig() GoogleCesToolsetOpenApiToolsetApiAuthenticationOauthConfigOutputReference
	OauthConfigInput() *GoogleCesToolsetOpenApiToolsetApiAuthenticationOauthConfig
	ServiceAccountAuthConfig() GoogleCesToolsetOpenApiToolsetApiAuthenticationServiceAccountAuthConfigOutputReference
	ServiceAccountAuthConfigInput() *GoogleCesToolsetOpenApiToolsetApiAuthenticationServiceAccountAuthConfig
	ServiceAgentIdTokenAuthConfig() GoogleCesToolsetOpenApiToolsetApiAuthenticationServiceAgentIdTokenAuthConfigOutputReference
	ServiceAgentIdTokenAuthConfigInput() *GoogleCesToolsetOpenApiToolsetApiAuthenticationServiceAgentIdTokenAuthConfig
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
	PutApiKeyConfig(value *GoogleCesToolsetOpenApiToolsetApiAuthenticationApiKeyConfig)
	PutBearerTokenConfig(value *GoogleCesToolsetOpenApiToolsetApiAuthenticationBearerTokenConfig)
	PutOauthConfig(value *GoogleCesToolsetOpenApiToolsetApiAuthenticationOauthConfig)
	PutServiceAccountAuthConfig(value *GoogleCesToolsetOpenApiToolsetApiAuthenticationServiceAccountAuthConfig)
	PutServiceAgentIdTokenAuthConfig(value *GoogleCesToolsetOpenApiToolsetApiAuthenticationServiceAgentIdTokenAuthConfig)
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

// The jsii proxy struct for GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference
type jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) ApiKeyConfig() GoogleCesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference {
	var returns GoogleCesToolsetOpenApiToolsetApiAuthenticationApiKeyConfigOutputReference
	_jsii_.Get(
		j,
		"apiKeyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) ApiKeyConfigInput() *GoogleCesToolsetOpenApiToolsetApiAuthenticationApiKeyConfig {
	var returns *GoogleCesToolsetOpenApiToolsetApiAuthenticationApiKeyConfig
	_jsii_.Get(
		j,
		"apiKeyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) BearerTokenConfig() GoogleCesToolsetOpenApiToolsetApiAuthenticationBearerTokenConfigOutputReference {
	var returns GoogleCesToolsetOpenApiToolsetApiAuthenticationBearerTokenConfigOutputReference
	_jsii_.Get(
		j,
		"bearerTokenConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) BearerTokenConfigInput() *GoogleCesToolsetOpenApiToolsetApiAuthenticationBearerTokenConfig {
	var returns *GoogleCesToolsetOpenApiToolsetApiAuthenticationBearerTokenConfig
	_jsii_.Get(
		j,
		"bearerTokenConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) InternalValue() *GoogleCesToolsetOpenApiToolsetApiAuthentication {
	var returns *GoogleCesToolsetOpenApiToolsetApiAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) OauthConfig() GoogleCesToolsetOpenApiToolsetApiAuthenticationOauthConfigOutputReference {
	var returns GoogleCesToolsetOpenApiToolsetApiAuthenticationOauthConfigOutputReference
	_jsii_.Get(
		j,
		"oauthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) OauthConfigInput() *GoogleCesToolsetOpenApiToolsetApiAuthenticationOauthConfig {
	var returns *GoogleCesToolsetOpenApiToolsetApiAuthenticationOauthConfig
	_jsii_.Get(
		j,
		"oauthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) ServiceAccountAuthConfig() GoogleCesToolsetOpenApiToolsetApiAuthenticationServiceAccountAuthConfigOutputReference {
	var returns GoogleCesToolsetOpenApiToolsetApiAuthenticationServiceAccountAuthConfigOutputReference
	_jsii_.Get(
		j,
		"serviceAccountAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) ServiceAccountAuthConfigInput() *GoogleCesToolsetOpenApiToolsetApiAuthenticationServiceAccountAuthConfig {
	var returns *GoogleCesToolsetOpenApiToolsetApiAuthenticationServiceAccountAuthConfig
	_jsii_.Get(
		j,
		"serviceAccountAuthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) ServiceAgentIdTokenAuthConfig() GoogleCesToolsetOpenApiToolsetApiAuthenticationServiceAgentIdTokenAuthConfigOutputReference {
	var returns GoogleCesToolsetOpenApiToolsetApiAuthenticationServiceAgentIdTokenAuthConfigOutputReference
	_jsii_.Get(
		j,
		"serviceAgentIdTokenAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) ServiceAgentIdTokenAuthConfigInput() *GoogleCesToolsetOpenApiToolsetApiAuthenticationServiceAgentIdTokenAuthConfig {
	var returns *GoogleCesToolsetOpenApiToolsetApiAuthenticationServiceAgentIdTokenAuthConfig
	_jsii_.Get(
		j,
		"serviceAgentIdTokenAuthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesToolset.GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference_Override(g GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesToolset.GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference)SetInternalValue(val *GoogleCesToolsetOpenApiToolsetApiAuthentication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) PutApiKeyConfig(value *GoogleCesToolsetOpenApiToolsetApiAuthenticationApiKeyConfig) {
	if err := g.validatePutApiKeyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApiKeyConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) PutBearerTokenConfig(value *GoogleCesToolsetOpenApiToolsetApiAuthenticationBearerTokenConfig) {
	if err := g.validatePutBearerTokenConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBearerTokenConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) PutOauthConfig(value *GoogleCesToolsetOpenApiToolsetApiAuthenticationOauthConfig) {
	if err := g.validatePutOauthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauthConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) PutServiceAccountAuthConfig(value *GoogleCesToolsetOpenApiToolsetApiAuthenticationServiceAccountAuthConfig) {
	if err := g.validatePutServiceAccountAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceAccountAuthConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) PutServiceAgentIdTokenAuthConfig(value *GoogleCesToolsetOpenApiToolsetApiAuthenticationServiceAgentIdTokenAuthConfig) {
	if err := g.validatePutServiceAgentIdTokenAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceAgentIdTokenAuthConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) ResetApiKeyConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetApiKeyConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) ResetBearerTokenConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetBearerTokenConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) ResetOauthConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetOauthConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) ResetServiceAccountAuthConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccountAuthConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) ResetServiceAgentIdTokenAuthConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAgentIdTokenAuthConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolsetOpenApiToolsetApiAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package googlecestoolset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecestoolset/internal"
)

type GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference interface {
	cdktn.ComplexObject
	ApiKeyConfig() GoogleCesToolsetMcpToolsetApiAuthenticationApiKeyConfigOutputReference
	ApiKeyConfigInput() *GoogleCesToolsetMcpToolsetApiAuthenticationApiKeyConfig
	BearerTokenConfig() GoogleCesToolsetMcpToolsetApiAuthenticationBearerTokenConfigOutputReference
	BearerTokenConfigInput() *GoogleCesToolsetMcpToolsetApiAuthenticationBearerTokenConfig
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
	InternalValue() *GoogleCesToolsetMcpToolsetApiAuthentication
	SetInternalValue(val *GoogleCesToolsetMcpToolsetApiAuthentication)
	OauthConfig() GoogleCesToolsetMcpToolsetApiAuthenticationOauthConfigOutputReference
	OauthConfigInput() *GoogleCesToolsetMcpToolsetApiAuthenticationOauthConfig
	ServiceAccountAuthConfig() GoogleCesToolsetMcpToolsetApiAuthenticationServiceAccountAuthConfigOutputReference
	ServiceAccountAuthConfigInput() *GoogleCesToolsetMcpToolsetApiAuthenticationServiceAccountAuthConfig
	ServiceAgentIdTokenAuthConfig() GoogleCesToolsetMcpToolsetApiAuthenticationServiceAgentIdTokenAuthConfigOutputReference
	ServiceAgentIdTokenAuthConfigInput() *GoogleCesToolsetMcpToolsetApiAuthenticationServiceAgentIdTokenAuthConfig
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
	PutApiKeyConfig(value *GoogleCesToolsetMcpToolsetApiAuthenticationApiKeyConfig)
	PutBearerTokenConfig(value *GoogleCesToolsetMcpToolsetApiAuthenticationBearerTokenConfig)
	PutOauthConfig(value *GoogleCesToolsetMcpToolsetApiAuthenticationOauthConfig)
	PutServiceAccountAuthConfig(value *GoogleCesToolsetMcpToolsetApiAuthenticationServiceAccountAuthConfig)
	PutServiceAgentIdTokenAuthConfig(value *GoogleCesToolsetMcpToolsetApiAuthenticationServiceAgentIdTokenAuthConfig)
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

// The jsii proxy struct for GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference
type jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) ApiKeyConfig() GoogleCesToolsetMcpToolsetApiAuthenticationApiKeyConfigOutputReference {
	var returns GoogleCesToolsetMcpToolsetApiAuthenticationApiKeyConfigOutputReference
	_jsii_.Get(
		j,
		"apiKeyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) ApiKeyConfigInput() *GoogleCesToolsetMcpToolsetApiAuthenticationApiKeyConfig {
	var returns *GoogleCesToolsetMcpToolsetApiAuthenticationApiKeyConfig
	_jsii_.Get(
		j,
		"apiKeyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) BearerTokenConfig() GoogleCesToolsetMcpToolsetApiAuthenticationBearerTokenConfigOutputReference {
	var returns GoogleCesToolsetMcpToolsetApiAuthenticationBearerTokenConfigOutputReference
	_jsii_.Get(
		j,
		"bearerTokenConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) BearerTokenConfigInput() *GoogleCesToolsetMcpToolsetApiAuthenticationBearerTokenConfig {
	var returns *GoogleCesToolsetMcpToolsetApiAuthenticationBearerTokenConfig
	_jsii_.Get(
		j,
		"bearerTokenConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) InternalValue() *GoogleCesToolsetMcpToolsetApiAuthentication {
	var returns *GoogleCesToolsetMcpToolsetApiAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) OauthConfig() GoogleCesToolsetMcpToolsetApiAuthenticationOauthConfigOutputReference {
	var returns GoogleCesToolsetMcpToolsetApiAuthenticationOauthConfigOutputReference
	_jsii_.Get(
		j,
		"oauthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) OauthConfigInput() *GoogleCesToolsetMcpToolsetApiAuthenticationOauthConfig {
	var returns *GoogleCesToolsetMcpToolsetApiAuthenticationOauthConfig
	_jsii_.Get(
		j,
		"oauthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) ServiceAccountAuthConfig() GoogleCesToolsetMcpToolsetApiAuthenticationServiceAccountAuthConfigOutputReference {
	var returns GoogleCesToolsetMcpToolsetApiAuthenticationServiceAccountAuthConfigOutputReference
	_jsii_.Get(
		j,
		"serviceAccountAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) ServiceAccountAuthConfigInput() *GoogleCesToolsetMcpToolsetApiAuthenticationServiceAccountAuthConfig {
	var returns *GoogleCesToolsetMcpToolsetApiAuthenticationServiceAccountAuthConfig
	_jsii_.Get(
		j,
		"serviceAccountAuthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) ServiceAgentIdTokenAuthConfig() GoogleCesToolsetMcpToolsetApiAuthenticationServiceAgentIdTokenAuthConfigOutputReference {
	var returns GoogleCesToolsetMcpToolsetApiAuthenticationServiceAgentIdTokenAuthConfigOutputReference
	_jsii_.Get(
		j,
		"serviceAgentIdTokenAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) ServiceAgentIdTokenAuthConfigInput() *GoogleCesToolsetMcpToolsetApiAuthenticationServiceAgentIdTokenAuthConfig {
	var returns *GoogleCesToolsetMcpToolsetApiAuthenticationServiceAgentIdTokenAuthConfig
	_jsii_.Get(
		j,
		"serviceAgentIdTokenAuthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCesToolsetMcpToolsetApiAuthenticationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesToolsetMcpToolsetApiAuthenticationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesToolset.GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesToolsetMcpToolsetApiAuthenticationOutputReference_Override(g GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesToolset.GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference)SetInternalValue(val *GoogleCesToolsetMcpToolsetApiAuthentication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) PutApiKeyConfig(value *GoogleCesToolsetMcpToolsetApiAuthenticationApiKeyConfig) {
	if err := g.validatePutApiKeyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApiKeyConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) PutBearerTokenConfig(value *GoogleCesToolsetMcpToolsetApiAuthenticationBearerTokenConfig) {
	if err := g.validatePutBearerTokenConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBearerTokenConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) PutOauthConfig(value *GoogleCesToolsetMcpToolsetApiAuthenticationOauthConfig) {
	if err := g.validatePutOauthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauthConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) PutServiceAccountAuthConfig(value *GoogleCesToolsetMcpToolsetApiAuthenticationServiceAccountAuthConfig) {
	if err := g.validatePutServiceAccountAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceAccountAuthConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) PutServiceAgentIdTokenAuthConfig(value *GoogleCesToolsetMcpToolsetApiAuthenticationServiceAgentIdTokenAuthConfig) {
	if err := g.validatePutServiceAgentIdTokenAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceAgentIdTokenAuthConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) ResetApiKeyConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetApiKeyConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) ResetBearerTokenConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetBearerTokenConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) ResetOauthConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetOauthConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) ResetServiceAccountAuthConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccountAuthConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) ResetServiceAgentIdTokenAuthConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAgentIdTokenAuthConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesToolsetMcpToolsetApiAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


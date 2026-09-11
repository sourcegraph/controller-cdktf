package googledialogflowcxtoolversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowcxtoolversion/internal"
)

type GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference interface {
	cdktn.ComplexObject
	ApiKeyConfig() GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationApiKeyConfigOutputReference
	ApiKeyConfigInput() *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationApiKeyConfig
	BearerTokenConfig() GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationBearerTokenConfigOutputReference
	BearerTokenConfigInput() *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationBearerTokenConfig
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
	InternalValue() *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthentication
	SetInternalValue(val *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthentication)
	OauthConfig() GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOauthConfigOutputReference
	OauthConfigInput() *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOauthConfig
	ServiceAgentAuthConfig() GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationServiceAgentAuthConfigOutputReference
	ServiceAgentAuthConfigInput() *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationServiceAgentAuthConfig
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
	PutApiKeyConfig(value *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationApiKeyConfig)
	PutBearerTokenConfig(value *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationBearerTokenConfig)
	PutOauthConfig(value *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOauthConfig)
	PutServiceAgentAuthConfig(value *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationServiceAgentAuthConfig)
	ResetApiKeyConfig()
	ResetBearerTokenConfig()
	ResetOauthConfig()
	ResetServiceAgentAuthConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference
type jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ApiKeyConfig() GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationApiKeyConfigOutputReference {
	var returns GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationApiKeyConfigOutputReference
	_jsii_.Get(
		j,
		"apiKeyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ApiKeyConfigInput() *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationApiKeyConfig {
	var returns *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationApiKeyConfig
	_jsii_.Get(
		j,
		"apiKeyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) BearerTokenConfig() GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationBearerTokenConfigOutputReference {
	var returns GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationBearerTokenConfigOutputReference
	_jsii_.Get(
		j,
		"bearerTokenConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) BearerTokenConfigInput() *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationBearerTokenConfig {
	var returns *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationBearerTokenConfig
	_jsii_.Get(
		j,
		"bearerTokenConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) InternalValue() *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthentication {
	var returns *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) OauthConfig() GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOauthConfigOutputReference {
	var returns GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOauthConfigOutputReference
	_jsii_.Get(
		j,
		"oauthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) OauthConfigInput() *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOauthConfig {
	var returns *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOauthConfig
	_jsii_.Get(
		j,
		"oauthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ServiceAgentAuthConfig() GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationServiceAgentAuthConfigOutputReference {
	var returns GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationServiceAgentAuthConfigOutputReference
	_jsii_.Get(
		j,
		"serviceAgentAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ServiceAgentAuthConfigInput() *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationServiceAgentAuthConfig {
	var returns *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationServiceAgentAuthConfig
	_jsii_.Get(
		j,
		"serviceAgentAuthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowCxToolVersion.GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference_Override(g GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowCxToolVersion.GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference)SetInternalValue(val *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthentication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) PutApiKeyConfig(value *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationApiKeyConfig) {
	if err := g.validatePutApiKeyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApiKeyConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) PutBearerTokenConfig(value *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationBearerTokenConfig) {
	if err := g.validatePutBearerTokenConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBearerTokenConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) PutOauthConfig(value *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOauthConfig) {
	if err := g.validatePutOauthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauthConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) PutServiceAgentAuthConfig(value *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationServiceAgentAuthConfig) {
	if err := g.validatePutServiceAgentAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceAgentAuthConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ResetApiKeyConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetApiKeyConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ResetBearerTokenConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetBearerTokenConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ResetOauthConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetOauthConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ResetServiceAgentAuthConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAgentAuthConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package googleapihubplugininstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleapihubplugininstance/internal"
)

type GoogleApihubPluginInstanceAuthConfigOutputReference interface {
	cdktf.ComplexObject
	ApiKeyConfig() GoogleApihubPluginInstanceAuthConfigApiKeyConfigOutputReference
	ApiKeyConfigInput() *GoogleApihubPluginInstanceAuthConfigApiKeyConfig
	AuthType() *string
	SetAuthType(val *string)
	AuthTypeInput() *string
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
	GoogleServiceAccountConfig() GoogleApihubPluginInstanceAuthConfigGoogleServiceAccountConfigOutputReference
	GoogleServiceAccountConfigInput() *GoogleApihubPluginInstanceAuthConfigGoogleServiceAccountConfig
	InternalValue() *GoogleApihubPluginInstanceAuthConfig
	SetInternalValue(val *GoogleApihubPluginInstanceAuthConfig)
	Oauth2ClientCredentialsConfig() GoogleApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigOutputReference
	Oauth2ClientCredentialsConfigInput() *GoogleApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserPasswordConfig() GoogleApihubPluginInstanceAuthConfigUserPasswordConfigOutputReference
	UserPasswordConfigInput() *GoogleApihubPluginInstanceAuthConfigUserPasswordConfig
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
	PutApiKeyConfig(value *GoogleApihubPluginInstanceAuthConfigApiKeyConfig)
	PutGoogleServiceAccountConfig(value *GoogleApihubPluginInstanceAuthConfigGoogleServiceAccountConfig)
	PutOauth2ClientCredentialsConfig(value *GoogleApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfig)
	PutUserPasswordConfig(value *GoogleApihubPluginInstanceAuthConfigUserPasswordConfig)
	ResetApiKeyConfig()
	ResetGoogleServiceAccountConfig()
	ResetOauth2ClientCredentialsConfig()
	ResetUserPasswordConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleApihubPluginInstanceAuthConfigOutputReference
type jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) ApiKeyConfig() GoogleApihubPluginInstanceAuthConfigApiKeyConfigOutputReference {
	var returns GoogleApihubPluginInstanceAuthConfigApiKeyConfigOutputReference
	_jsii_.Get(
		j,
		"apiKeyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) ApiKeyConfigInput() *GoogleApihubPluginInstanceAuthConfigApiKeyConfig {
	var returns *GoogleApihubPluginInstanceAuthConfigApiKeyConfig
	_jsii_.Get(
		j,
		"apiKeyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) GoogleServiceAccountConfig() GoogleApihubPluginInstanceAuthConfigGoogleServiceAccountConfigOutputReference {
	var returns GoogleApihubPluginInstanceAuthConfigGoogleServiceAccountConfigOutputReference
	_jsii_.Get(
		j,
		"googleServiceAccountConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) GoogleServiceAccountConfigInput() *GoogleApihubPluginInstanceAuthConfigGoogleServiceAccountConfig {
	var returns *GoogleApihubPluginInstanceAuthConfigGoogleServiceAccountConfig
	_jsii_.Get(
		j,
		"googleServiceAccountConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) InternalValue() *GoogleApihubPluginInstanceAuthConfig {
	var returns *GoogleApihubPluginInstanceAuthConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) Oauth2ClientCredentialsConfig() GoogleApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigOutputReference {
	var returns GoogleApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigOutputReference
	_jsii_.Get(
		j,
		"oauth2ClientCredentialsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) Oauth2ClientCredentialsConfigInput() *GoogleApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfig {
	var returns *GoogleApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfig
	_jsii_.Get(
		j,
		"oauth2ClientCredentialsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) UserPasswordConfig() GoogleApihubPluginInstanceAuthConfigUserPasswordConfigOutputReference {
	var returns GoogleApihubPluginInstanceAuthConfigUserPasswordConfigOutputReference
	_jsii_.Get(
		j,
		"userPasswordConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) UserPasswordConfigInput() *GoogleApihubPluginInstanceAuthConfigUserPasswordConfig {
	var returns *GoogleApihubPluginInstanceAuthConfigUserPasswordConfig
	_jsii_.Get(
		j,
		"userPasswordConfigInput",
		&returns,
	)
	return returns
}


func NewGoogleApihubPluginInstanceAuthConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleApihubPluginInstanceAuthConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleApihubPluginInstanceAuthConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleApihubPluginInstance.GoogleApihubPluginInstanceAuthConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleApihubPluginInstanceAuthConfigOutputReference_Override(g GoogleApihubPluginInstanceAuthConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleApihubPluginInstance.GoogleApihubPluginInstanceAuthConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference)SetAuthType(val *string) {
	if err := j.validateSetAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference)SetInternalValue(val *GoogleApihubPluginInstanceAuthConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) PutApiKeyConfig(value *GoogleApihubPluginInstanceAuthConfigApiKeyConfig) {
	if err := g.validatePutApiKeyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApiKeyConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) PutGoogleServiceAccountConfig(value *GoogleApihubPluginInstanceAuthConfigGoogleServiceAccountConfig) {
	if err := g.validatePutGoogleServiceAccountConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGoogleServiceAccountConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) PutOauth2ClientCredentialsConfig(value *GoogleApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfig) {
	if err := g.validatePutOauth2ClientCredentialsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauth2ClientCredentialsConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) PutUserPasswordConfig(value *GoogleApihubPluginInstanceAuthConfigUserPasswordConfig) {
	if err := g.validatePutUserPasswordConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUserPasswordConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) ResetApiKeyConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetApiKeyConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) ResetGoogleServiceAccountConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGoogleServiceAccountConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) ResetOauth2ClientCredentialsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetOauth2ClientCredentialsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) ResetUserPasswordConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetUserPasswordConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceAuthConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


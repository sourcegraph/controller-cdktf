package googleintegrationconnectorsconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleintegrationconnectorsconnection/internal"
)

type GoogleIntegrationConnectorsConnectionAuthConfigOutputReference interface {
	cdktf.ComplexObject
	AdditionalVariable() GoogleIntegrationConnectorsConnectionAuthConfigAdditionalVariableList
	AdditionalVariableInput() interface{}
	AuthKey() *string
	SetAuthKey(val *string)
	AuthKeyInput() *string
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
	InternalValue() *GoogleIntegrationConnectorsConnectionAuthConfig
	SetInternalValue(val *GoogleIntegrationConnectorsConnectionAuthConfig)
	Oauth2AuthCodeFlow() GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference
	Oauth2AuthCodeFlowInput() *GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow
	Oauth2ClientCredentials() GoogleIntegrationConnectorsConnectionAuthConfigOauth2ClientCredentialsOutputReference
	Oauth2ClientCredentialsInput() *GoogleIntegrationConnectorsConnectionAuthConfigOauth2ClientCredentials
	Oauth2JwtBearer() GoogleIntegrationConnectorsConnectionAuthConfigOauth2JwtBearerOutputReference
	Oauth2JwtBearerInput() *GoogleIntegrationConnectorsConnectionAuthConfigOauth2JwtBearer
	SshPublicKey() GoogleIntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference
	SshPublicKeyInput() *GoogleIntegrationConnectorsConnectionAuthConfigSshPublicKey
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserPassword() GoogleIntegrationConnectorsConnectionAuthConfigUserPasswordOutputReference
	UserPasswordInput() *GoogleIntegrationConnectorsConnectionAuthConfigUserPassword
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
	PutAdditionalVariable(value interface{})
	PutOauth2AuthCodeFlow(value *GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow)
	PutOauth2ClientCredentials(value *GoogleIntegrationConnectorsConnectionAuthConfigOauth2ClientCredentials)
	PutOauth2JwtBearer(value *GoogleIntegrationConnectorsConnectionAuthConfigOauth2JwtBearer)
	PutSshPublicKey(value *GoogleIntegrationConnectorsConnectionAuthConfigSshPublicKey)
	PutUserPassword(value *GoogleIntegrationConnectorsConnectionAuthConfigUserPassword)
	ResetAdditionalVariable()
	ResetAuthKey()
	ResetOauth2AuthCodeFlow()
	ResetOauth2ClientCredentials()
	ResetOauth2JwtBearer()
	ResetSshPublicKey()
	ResetUserPassword()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIntegrationConnectorsConnectionAuthConfigOutputReference
type jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) AdditionalVariable() GoogleIntegrationConnectorsConnectionAuthConfigAdditionalVariableList {
	var returns GoogleIntegrationConnectorsConnectionAuthConfigAdditionalVariableList
	_jsii_.Get(
		j,
		"additionalVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) AdditionalVariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) AuthKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) AuthKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) InternalValue() *GoogleIntegrationConnectorsConnectionAuthConfig {
	var returns *GoogleIntegrationConnectorsConnectionAuthConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) Oauth2AuthCodeFlow() GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference {
	var returns GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference
	_jsii_.Get(
		j,
		"oauth2AuthCodeFlow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) Oauth2AuthCodeFlowInput() *GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow {
	var returns *GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow
	_jsii_.Get(
		j,
		"oauth2AuthCodeFlowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) Oauth2ClientCredentials() GoogleIntegrationConnectorsConnectionAuthConfigOauth2ClientCredentialsOutputReference {
	var returns GoogleIntegrationConnectorsConnectionAuthConfigOauth2ClientCredentialsOutputReference
	_jsii_.Get(
		j,
		"oauth2ClientCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) Oauth2ClientCredentialsInput() *GoogleIntegrationConnectorsConnectionAuthConfigOauth2ClientCredentials {
	var returns *GoogleIntegrationConnectorsConnectionAuthConfigOauth2ClientCredentials
	_jsii_.Get(
		j,
		"oauth2ClientCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) Oauth2JwtBearer() GoogleIntegrationConnectorsConnectionAuthConfigOauth2JwtBearerOutputReference {
	var returns GoogleIntegrationConnectorsConnectionAuthConfigOauth2JwtBearerOutputReference
	_jsii_.Get(
		j,
		"oauth2JwtBearer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) Oauth2JwtBearerInput() *GoogleIntegrationConnectorsConnectionAuthConfigOauth2JwtBearer {
	var returns *GoogleIntegrationConnectorsConnectionAuthConfigOauth2JwtBearer
	_jsii_.Get(
		j,
		"oauth2JwtBearerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) SshPublicKey() GoogleIntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference {
	var returns GoogleIntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference
	_jsii_.Get(
		j,
		"sshPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) SshPublicKeyInput() *GoogleIntegrationConnectorsConnectionAuthConfigSshPublicKey {
	var returns *GoogleIntegrationConnectorsConnectionAuthConfigSshPublicKey
	_jsii_.Get(
		j,
		"sshPublicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) UserPassword() GoogleIntegrationConnectorsConnectionAuthConfigUserPasswordOutputReference {
	var returns GoogleIntegrationConnectorsConnectionAuthConfigUserPasswordOutputReference
	_jsii_.Get(
		j,
		"userPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) UserPasswordInput() *GoogleIntegrationConnectorsConnectionAuthConfigUserPassword {
	var returns *GoogleIntegrationConnectorsConnectionAuthConfigUserPassword
	_jsii_.Get(
		j,
		"userPasswordInput",
		&returns,
	)
	return returns
}


func NewGoogleIntegrationConnectorsConnectionAuthConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleIntegrationConnectorsConnectionAuthConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIntegrationConnectorsConnectionAuthConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIntegrationConnectorsConnection.GoogleIntegrationConnectorsConnectionAuthConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleIntegrationConnectorsConnectionAuthConfigOutputReference_Override(g GoogleIntegrationConnectorsConnectionAuthConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIntegrationConnectorsConnection.GoogleIntegrationConnectorsConnectionAuthConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference)SetAuthKey(val *string) {
	if err := j.validateSetAuthKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authKey",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference)SetAuthType(val *string) {
	if err := j.validateSetAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference)SetInternalValue(val *GoogleIntegrationConnectorsConnectionAuthConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) PutAdditionalVariable(value interface{}) {
	if err := g.validatePutAdditionalVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdditionalVariable",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) PutOauth2AuthCodeFlow(value *GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow) {
	if err := g.validatePutOauth2AuthCodeFlowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauth2AuthCodeFlow",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) PutOauth2ClientCredentials(value *GoogleIntegrationConnectorsConnectionAuthConfigOauth2ClientCredentials) {
	if err := g.validatePutOauth2ClientCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauth2ClientCredentials",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) PutOauth2JwtBearer(value *GoogleIntegrationConnectorsConnectionAuthConfigOauth2JwtBearer) {
	if err := g.validatePutOauth2JwtBearerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauth2JwtBearer",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) PutSshPublicKey(value *GoogleIntegrationConnectorsConnectionAuthConfigSshPublicKey) {
	if err := g.validatePutSshPublicKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSshPublicKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) PutUserPassword(value *GoogleIntegrationConnectorsConnectionAuthConfigUserPassword) {
	if err := g.validatePutUserPasswordParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUserPassword",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) ResetAdditionalVariable() {
	_jsii_.InvokeVoid(
		g,
		"resetAdditionalVariable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) ResetAuthKey() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) ResetOauth2AuthCodeFlow() {
	_jsii_.InvokeVoid(
		g,
		"resetOauth2AuthCodeFlow",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) ResetOauth2ClientCredentials() {
	_jsii_.InvokeVoid(
		g,
		"resetOauth2ClientCredentials",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) ResetOauth2JwtBearer() {
	_jsii_.InvokeVoid(
		g,
		"resetOauth2JwtBearer",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) ResetSshPublicKey() {
	_jsii_.InvokeVoid(
		g,
		"resetSshPublicKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) ResetUserPassword() {
	_jsii_.InvokeVoid(
		g,
		"resetUserPassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


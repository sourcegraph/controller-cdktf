package googleintegrationsauthconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleintegrationsauthconfig/internal"
)

type GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference interface {
	cdktf.ComplexObject
	AuthToken() GoogleIntegrationsAuthConfigDecryptedCredentialAuthTokenOutputReference
	AuthTokenInput() *GoogleIntegrationsAuthConfigDecryptedCredentialAuthToken
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
	CredentialType() *string
	SetCredentialType(val *string)
	CredentialTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleIntegrationsAuthConfigDecryptedCredential
	SetInternalValue(val *GoogleIntegrationsAuthConfigDecryptedCredential)
	Jwt() GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference
	JwtInput() *GoogleIntegrationsAuthConfigDecryptedCredentialJwt
	Oauth2AuthorizationCode() GoogleIntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeOutputReference
	Oauth2AuthorizationCodeInput() *GoogleIntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCode
	Oauth2ClientCredentials() GoogleIntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsOutputReference
	Oauth2ClientCredentialsInput() *GoogleIntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentials
	OidcToken() GoogleIntegrationsAuthConfigDecryptedCredentialOidcTokenOutputReference
	OidcTokenInput() *GoogleIntegrationsAuthConfigDecryptedCredentialOidcToken
	ServiceAccountCredentials() GoogleIntegrationsAuthConfigDecryptedCredentialServiceAccountCredentialsOutputReference
	ServiceAccountCredentialsInput() *GoogleIntegrationsAuthConfigDecryptedCredentialServiceAccountCredentials
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UsernameAndPassword() GoogleIntegrationsAuthConfigDecryptedCredentialUsernameAndPasswordOutputReference
	UsernameAndPasswordInput() *GoogleIntegrationsAuthConfigDecryptedCredentialUsernameAndPassword
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
	PutAuthToken(value *GoogleIntegrationsAuthConfigDecryptedCredentialAuthToken)
	PutJwt(value *GoogleIntegrationsAuthConfigDecryptedCredentialJwt)
	PutOauth2AuthorizationCode(value *GoogleIntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCode)
	PutOauth2ClientCredentials(value *GoogleIntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentials)
	PutOidcToken(value *GoogleIntegrationsAuthConfigDecryptedCredentialOidcToken)
	PutServiceAccountCredentials(value *GoogleIntegrationsAuthConfigDecryptedCredentialServiceAccountCredentials)
	PutUsernameAndPassword(value *GoogleIntegrationsAuthConfigDecryptedCredentialUsernameAndPassword)
	ResetAuthToken()
	ResetJwt()
	ResetOauth2AuthorizationCode()
	ResetOauth2ClientCredentials()
	ResetOidcToken()
	ResetServiceAccountCredentials()
	ResetUsernameAndPassword()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference
type jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) AuthToken() GoogleIntegrationsAuthConfigDecryptedCredentialAuthTokenOutputReference {
	var returns GoogleIntegrationsAuthConfigDecryptedCredentialAuthTokenOutputReference
	_jsii_.Get(
		j,
		"authToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) AuthTokenInput() *GoogleIntegrationsAuthConfigDecryptedCredentialAuthToken {
	var returns *GoogleIntegrationsAuthConfigDecryptedCredentialAuthToken
	_jsii_.Get(
		j,
		"authTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) CredentialType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) CredentialTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) InternalValue() *GoogleIntegrationsAuthConfigDecryptedCredential {
	var returns *GoogleIntegrationsAuthConfigDecryptedCredential
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) Jwt() GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference {
	var returns GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference
	_jsii_.Get(
		j,
		"jwt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) JwtInput() *GoogleIntegrationsAuthConfigDecryptedCredentialJwt {
	var returns *GoogleIntegrationsAuthConfigDecryptedCredentialJwt
	_jsii_.Get(
		j,
		"jwtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) Oauth2AuthorizationCode() GoogleIntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeOutputReference {
	var returns GoogleIntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeOutputReference
	_jsii_.Get(
		j,
		"oauth2AuthorizationCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) Oauth2AuthorizationCodeInput() *GoogleIntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCode {
	var returns *GoogleIntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCode
	_jsii_.Get(
		j,
		"oauth2AuthorizationCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) Oauth2ClientCredentials() GoogleIntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsOutputReference {
	var returns GoogleIntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsOutputReference
	_jsii_.Get(
		j,
		"oauth2ClientCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) Oauth2ClientCredentialsInput() *GoogleIntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentials {
	var returns *GoogleIntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentials
	_jsii_.Get(
		j,
		"oauth2ClientCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) OidcToken() GoogleIntegrationsAuthConfigDecryptedCredentialOidcTokenOutputReference {
	var returns GoogleIntegrationsAuthConfigDecryptedCredentialOidcTokenOutputReference
	_jsii_.Get(
		j,
		"oidcToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) OidcTokenInput() *GoogleIntegrationsAuthConfigDecryptedCredentialOidcToken {
	var returns *GoogleIntegrationsAuthConfigDecryptedCredentialOidcToken
	_jsii_.Get(
		j,
		"oidcTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) ServiceAccountCredentials() GoogleIntegrationsAuthConfigDecryptedCredentialServiceAccountCredentialsOutputReference {
	var returns GoogleIntegrationsAuthConfigDecryptedCredentialServiceAccountCredentialsOutputReference
	_jsii_.Get(
		j,
		"serviceAccountCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) ServiceAccountCredentialsInput() *GoogleIntegrationsAuthConfigDecryptedCredentialServiceAccountCredentials {
	var returns *GoogleIntegrationsAuthConfigDecryptedCredentialServiceAccountCredentials
	_jsii_.Get(
		j,
		"serviceAccountCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) UsernameAndPassword() GoogleIntegrationsAuthConfigDecryptedCredentialUsernameAndPasswordOutputReference {
	var returns GoogleIntegrationsAuthConfigDecryptedCredentialUsernameAndPasswordOutputReference
	_jsii_.Get(
		j,
		"usernameAndPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) UsernameAndPasswordInput() *GoogleIntegrationsAuthConfigDecryptedCredentialUsernameAndPassword {
	var returns *GoogleIntegrationsAuthConfigDecryptedCredentialUsernameAndPassword
	_jsii_.Get(
		j,
		"usernameAndPasswordInput",
		&returns,
	)
	return returns
}


func NewGoogleIntegrationsAuthConfigDecryptedCredentialOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIntegrationsAuthConfigDecryptedCredentialOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIntegrationsAuthConfig.GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleIntegrationsAuthConfigDecryptedCredentialOutputReference_Override(g GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIntegrationsAuthConfig.GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference)SetCredentialType(val *string) {
	if err := j.validateSetCredentialTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialType",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference)SetInternalValue(val *GoogleIntegrationsAuthConfigDecryptedCredential) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) PutAuthToken(value *GoogleIntegrationsAuthConfigDecryptedCredentialAuthToken) {
	if err := g.validatePutAuthTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthToken",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) PutJwt(value *GoogleIntegrationsAuthConfigDecryptedCredentialJwt) {
	if err := g.validatePutJwtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putJwt",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) PutOauth2AuthorizationCode(value *GoogleIntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCode) {
	if err := g.validatePutOauth2AuthorizationCodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauth2AuthorizationCode",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) PutOauth2ClientCredentials(value *GoogleIntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentials) {
	if err := g.validatePutOauth2ClientCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauth2ClientCredentials",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) PutOidcToken(value *GoogleIntegrationsAuthConfigDecryptedCredentialOidcToken) {
	if err := g.validatePutOidcTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOidcToken",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) PutServiceAccountCredentials(value *GoogleIntegrationsAuthConfigDecryptedCredentialServiceAccountCredentials) {
	if err := g.validatePutServiceAccountCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceAccountCredentials",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) PutUsernameAndPassword(value *GoogleIntegrationsAuthConfigDecryptedCredentialUsernameAndPassword) {
	if err := g.validatePutUsernameAndPasswordParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUsernameAndPassword",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) ResetAuthToken() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) ResetJwt() {
	_jsii_.InvokeVoid(
		g,
		"resetJwt",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) ResetOauth2AuthorizationCode() {
	_jsii_.InvokeVoid(
		g,
		"resetOauth2AuthorizationCode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) ResetOauth2ClientCredentials() {
	_jsii_.InvokeVoid(
		g,
		"resetOauth2ClientCredentials",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) ResetOidcToken() {
	_jsii_.InvokeVoid(
		g,
		"resetOidcToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) ResetServiceAccountCredentials() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccountCredentials",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) ResetUsernameAndPassword() {
	_jsii_.InvokeVoid(
		g,
		"resetUsernameAndPassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


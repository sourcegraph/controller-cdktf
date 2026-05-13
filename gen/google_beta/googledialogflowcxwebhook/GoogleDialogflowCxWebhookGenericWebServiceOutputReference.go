package googledialogflowcxwebhook

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowcxwebhook/internal"
)

type GoogleDialogflowCxWebhookGenericWebServiceOutputReference interface {
	cdktf.ComplexObject
	AllowedCaCerts() *[]*string
	SetAllowedCaCerts(val *[]*string)
	AllowedCaCertsInput() *[]*string
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
	HttpMethod() *string
	SetHttpMethod(val *string)
	HttpMethodInput() *string
	InternalValue() *GoogleDialogflowCxWebhookGenericWebService
	SetInternalValue(val *GoogleDialogflowCxWebhookGenericWebService)
	OauthConfig() GoogleDialogflowCxWebhookGenericWebServiceOauthConfigOutputReference
	OauthConfigInput() *GoogleDialogflowCxWebhookGenericWebServiceOauthConfig
	ParameterMapping() *map[string]*string
	SetParameterMapping(val *map[string]*string)
	ParameterMappingInput() *map[string]*string
	RequestBody() *string
	SetRequestBody(val *string)
	RequestBodyInput() *string
	RequestHeaders() *map[string]*string
	SetRequestHeaders(val *map[string]*string)
	RequestHeadersInput() *map[string]*string
	SecretVersionForUsernamePassword() *string
	SetSecretVersionForUsernamePassword(val *string)
	SecretVersionForUsernamePasswordInput() *string
	SecretVersionsForRequestHeaders() GoogleDialogflowCxWebhookGenericWebServiceSecretVersionsForRequestHeadersList
	SecretVersionsForRequestHeadersInput() interface{}
	ServiceAccountAuthConfig() GoogleDialogflowCxWebhookGenericWebServiceServiceAccountAuthConfigOutputReference
	ServiceAccountAuthConfigInput() *GoogleDialogflowCxWebhookGenericWebServiceServiceAccountAuthConfig
	ServiceAgentAuth() *string
	SetServiceAgentAuth(val *string)
	ServiceAgentAuthInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uri() *string
	SetUri(val *string)
	UriInput() *string
	WebhookType() *string
	SetWebhookType(val *string)
	WebhookTypeInput() *string
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
	PutOauthConfig(value *GoogleDialogflowCxWebhookGenericWebServiceOauthConfig)
	PutSecretVersionsForRequestHeaders(value interface{})
	PutServiceAccountAuthConfig(value *GoogleDialogflowCxWebhookGenericWebServiceServiceAccountAuthConfig)
	ResetAllowedCaCerts()
	ResetHttpMethod()
	ResetOauthConfig()
	ResetParameterMapping()
	ResetRequestBody()
	ResetRequestHeaders()
	ResetSecretVersionForUsernamePassword()
	ResetSecretVersionsForRequestHeaders()
	ResetServiceAccountAuthConfig()
	ResetServiceAgentAuth()
	ResetWebhookType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxWebhookGenericWebServiceOutputReference
type jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) AllowedCaCerts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedCaCerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) AllowedCaCertsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedCaCertsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) InternalValue() *GoogleDialogflowCxWebhookGenericWebService {
	var returns *GoogleDialogflowCxWebhookGenericWebService
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) OauthConfig() GoogleDialogflowCxWebhookGenericWebServiceOauthConfigOutputReference {
	var returns GoogleDialogflowCxWebhookGenericWebServiceOauthConfigOutputReference
	_jsii_.Get(
		j,
		"oauthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) OauthConfigInput() *GoogleDialogflowCxWebhookGenericWebServiceOauthConfig {
	var returns *GoogleDialogflowCxWebhookGenericWebServiceOauthConfig
	_jsii_.Get(
		j,
		"oauthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) ParameterMapping() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameterMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) ParameterMappingInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameterMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) RequestBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) RequestBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) RequestHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) RequestHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) SecretVersionForUsernamePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretVersionForUsernamePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) SecretVersionForUsernamePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretVersionForUsernamePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) SecretVersionsForRequestHeaders() GoogleDialogflowCxWebhookGenericWebServiceSecretVersionsForRequestHeadersList {
	var returns GoogleDialogflowCxWebhookGenericWebServiceSecretVersionsForRequestHeadersList
	_jsii_.Get(
		j,
		"secretVersionsForRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) SecretVersionsForRequestHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretVersionsForRequestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) ServiceAccountAuthConfig() GoogleDialogflowCxWebhookGenericWebServiceServiceAccountAuthConfigOutputReference {
	var returns GoogleDialogflowCxWebhookGenericWebServiceServiceAccountAuthConfigOutputReference
	_jsii_.Get(
		j,
		"serviceAccountAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) ServiceAccountAuthConfigInput() *GoogleDialogflowCxWebhookGenericWebServiceServiceAccountAuthConfig {
	var returns *GoogleDialogflowCxWebhookGenericWebServiceServiceAccountAuthConfig
	_jsii_.Get(
		j,
		"serviceAccountAuthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) ServiceAgentAuth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAgentAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) ServiceAgentAuthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAgentAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) WebhookType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) WebhookTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookTypeInput",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowCxWebhookGenericWebServiceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowCxWebhookGenericWebServiceOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxWebhookGenericWebServiceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxWebhook.GoogleDialogflowCxWebhookGenericWebServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxWebhookGenericWebServiceOutputReference_Override(g GoogleDialogflowCxWebhookGenericWebServiceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxWebhook.GoogleDialogflowCxWebhookGenericWebServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference)SetAllowedCaCerts(val *[]*string) {
	if err := j.validateSetAllowedCaCertsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedCaCerts",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference)SetHttpMethod(val *string) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference)SetInternalValue(val *GoogleDialogflowCxWebhookGenericWebService) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference)SetParameterMapping(val *map[string]*string) {
	if err := j.validateSetParameterMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterMapping",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference)SetRequestBody(val *string) {
	if err := j.validateSetRequestBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestBody",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference)SetRequestHeaders(val *map[string]*string) {
	if err := j.validateSetRequestHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestHeaders",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference)SetSecretVersionForUsernamePassword(val *string) {
	if err := j.validateSetSecretVersionForUsernamePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretVersionForUsernamePassword",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference)SetServiceAgentAuth(val *string) {
	if err := j.validateSetServiceAgentAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAgentAuth",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference)SetUri(val *string) {
	if err := j.validateSetUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference)SetWebhookType(val *string) {
	if err := j.validateSetWebhookTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhookType",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) PutOauthConfig(value *GoogleDialogflowCxWebhookGenericWebServiceOauthConfig) {
	if err := g.validatePutOauthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauthConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) PutSecretVersionsForRequestHeaders(value interface{}) {
	if err := g.validatePutSecretVersionsForRequestHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecretVersionsForRequestHeaders",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) PutServiceAccountAuthConfig(value *GoogleDialogflowCxWebhookGenericWebServiceServiceAccountAuthConfig) {
	if err := g.validatePutServiceAccountAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceAccountAuthConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) ResetAllowedCaCerts() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedCaCerts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) ResetHttpMethod() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpMethod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) ResetOauthConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetOauthConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) ResetParameterMapping() {
	_jsii_.InvokeVoid(
		g,
		"resetParameterMapping",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) ResetRequestBody() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestBody",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) ResetRequestHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) ResetSecretVersionForUsernamePassword() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretVersionForUsernamePassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) ResetSecretVersionsForRequestHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretVersionsForRequestHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) ResetServiceAccountAuthConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccountAuthConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) ResetServiceAgentAuth() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAgentAuth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) ResetWebhookType() {
	_jsii_.InvokeVoid(
		g,
		"resetWebhookType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookGenericWebServiceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


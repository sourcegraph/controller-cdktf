package googledialogflowcxwebhook

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowcxwebhook/internal"
)

type GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() *GoogleDialogflowCxWebhookServiceDirectoryGenericWebService
	SetInternalValue(val *GoogleDialogflowCxWebhookServiceDirectoryGenericWebService)
	OauthConfig() GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOauthConfigOutputReference
	OauthConfigInput() *GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOauthConfig
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
	SecretVersionsForRequestHeaders() GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList
	SecretVersionsForRequestHeadersInput() interface{}
	ServiceAccountAuthConfig() GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceServiceAccountAuthConfigOutputReference
	ServiceAccountAuthConfigInput() *GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceServiceAccountAuthConfig
	ServiceAgentAuth() *string
	SetServiceAgentAuth(val *string)
	ServiceAgentAuthInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutOauthConfig(value *GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOauthConfig)
	PutSecretVersionsForRequestHeaders(value interface{})
	PutServiceAccountAuthConfig(value *GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceServiceAccountAuthConfig)
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
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference
type jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) AllowedCaCerts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedCaCerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) AllowedCaCertsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedCaCertsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) InternalValue() *GoogleDialogflowCxWebhookServiceDirectoryGenericWebService {
	var returns *GoogleDialogflowCxWebhookServiceDirectoryGenericWebService
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) OauthConfig() GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOauthConfigOutputReference {
	var returns GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOauthConfigOutputReference
	_jsii_.Get(
		j,
		"oauthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) OauthConfigInput() *GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOauthConfig {
	var returns *GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOauthConfig
	_jsii_.Get(
		j,
		"oauthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ParameterMapping() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameterMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ParameterMappingInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameterMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) RequestBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) RequestBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) RequestHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) RequestHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) SecretVersionForUsernamePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretVersionForUsernamePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) SecretVersionForUsernamePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretVersionForUsernamePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) SecretVersionsForRequestHeaders() GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList {
	var returns GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList
	_jsii_.Get(
		j,
		"secretVersionsForRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) SecretVersionsForRequestHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretVersionsForRequestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ServiceAccountAuthConfig() GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceServiceAccountAuthConfigOutputReference {
	var returns GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceServiceAccountAuthConfigOutputReference
	_jsii_.Get(
		j,
		"serviceAccountAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ServiceAccountAuthConfigInput() *GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceServiceAccountAuthConfig {
	var returns *GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceServiceAccountAuthConfig
	_jsii_.Get(
		j,
		"serviceAccountAuthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ServiceAgentAuth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAgentAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ServiceAgentAuthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAgentAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) WebhookType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) WebhookTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookTypeInput",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowCxWebhook.GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference_Override(g GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowCxWebhook.GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetAllowedCaCerts(val *[]*string) {
	if err := j.validateSetAllowedCaCertsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedCaCerts",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetHttpMethod(val *string) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetInternalValue(val *GoogleDialogflowCxWebhookServiceDirectoryGenericWebService) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetParameterMapping(val *map[string]*string) {
	if err := j.validateSetParameterMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterMapping",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetRequestBody(val *string) {
	if err := j.validateSetRequestBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestBody",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetRequestHeaders(val *map[string]*string) {
	if err := j.validateSetRequestHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestHeaders",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetSecretVersionForUsernamePassword(val *string) {
	if err := j.validateSetSecretVersionForUsernamePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretVersionForUsernamePassword",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetServiceAgentAuth(val *string) {
	if err := j.validateSetServiceAgentAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAgentAuth",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetUri(val *string) {
	if err := j.validateSetUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetWebhookType(val *string) {
	if err := j.validateSetWebhookTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhookType",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) PutOauthConfig(value *GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOauthConfig) {
	if err := g.validatePutOauthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauthConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) PutSecretVersionsForRequestHeaders(value interface{}) {
	if err := g.validatePutSecretVersionsForRequestHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecretVersionsForRequestHeaders",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) PutServiceAccountAuthConfig(value *GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceServiceAccountAuthConfig) {
	if err := g.validatePutServiceAccountAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceAccountAuthConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ResetAllowedCaCerts() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedCaCerts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ResetHttpMethod() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpMethod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ResetOauthConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetOauthConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ResetParameterMapping() {
	_jsii_.InvokeVoid(
		g,
		"resetParameterMapping",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ResetRequestBody() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestBody",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ResetRequestHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ResetSecretVersionForUsernamePassword() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretVersionForUsernamePassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ResetSecretVersionsForRequestHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretVersionsForRequestHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ResetServiceAccountAuthConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccountAuthConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ResetServiceAgentAuth() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAgentAuth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ResetWebhookType() {
	_jsii_.InvokeVoid(
		g,
		"resetWebhookType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


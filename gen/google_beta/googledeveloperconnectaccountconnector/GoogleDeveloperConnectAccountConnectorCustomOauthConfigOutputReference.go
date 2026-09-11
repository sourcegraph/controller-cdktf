package googledeveloperconnectaccountconnector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledeveloperconnectaccountconnector/internal"
)

type GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference interface {
	cdktn.ComplexObject
	AuthUri() *string
	SetAuthUri(val *string)
	AuthUriInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
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
	HostUri() *string
	SetHostUri(val *string)
	HostUriInput() *string
	InternalValue() *GoogleDeveloperConnectAccountConnectorCustomOauthConfig
	SetInternalValue(val *GoogleDeveloperConnectAccountConnectorCustomOauthConfig)
	PkceDisabled() interface{}
	SetPkceDisabled(val interface{})
	PkceDisabledInput() interface{}
	ScmProvider() *string
	SetScmProvider(val *string)
	ScmProviderInput() *string
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	ServerVersion() *string
	ServiceDirectoryConfig() GoogleDeveloperConnectAccountConnectorCustomOauthConfigServiceDirectoryConfigOutputReference
	ServiceDirectoryConfigInput() *GoogleDeveloperConnectAccountConnectorCustomOauthConfigServiceDirectoryConfig
	SslCaCertificate() *string
	SetSslCaCertificate(val *string)
	SslCaCertificateInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TokenUri() *string
	SetTokenUri(val *string)
	TokenUriInput() *string
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
	PutServiceDirectoryConfig(value *GoogleDeveloperConnectAccountConnectorCustomOauthConfigServiceDirectoryConfig)
	ResetPkceDisabled()
	ResetServiceDirectoryConfig()
	ResetSslCaCertificate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference
type jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) AuthUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) AuthUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) HostUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) HostUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) InternalValue() *GoogleDeveloperConnectAccountConnectorCustomOauthConfig {
	var returns *GoogleDeveloperConnectAccountConnectorCustomOauthConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) PkceDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pkceDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) PkceDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pkceDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ScmProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ScmProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ServiceDirectoryConfig() GoogleDeveloperConnectAccountConnectorCustomOauthConfigServiceDirectoryConfigOutputReference {
	var returns GoogleDeveloperConnectAccountConnectorCustomOauthConfigServiceDirectoryConfigOutputReference
	_jsii_.Get(
		j,
		"serviceDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ServiceDirectoryConfigInput() *GoogleDeveloperConnectAccountConnectorCustomOauthConfigServiceDirectoryConfig {
	var returns *GoogleDeveloperConnectAccountConnectorCustomOauthConfigServiceDirectoryConfig
	_jsii_.Get(
		j,
		"serviceDirectoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) SslCaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) SslCaCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) TokenUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) TokenUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUriInput",
		&returns,
	)
	return returns
}


func NewGoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDeveloperConnectAccountConnector.GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference_Override(g GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDeveloperConnectAccountConnector.GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetAuthUri(val *string) {
	if err := j.validateSetAuthUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetHostUri(val *string) {
	if err := j.validateSetHostUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetInternalValue(val *GoogleDeveloperConnectAccountConnectorCustomOauthConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetPkceDisabled(val interface{}) {
	if err := j.validateSetPkceDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pkceDisabled",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetScmProvider(val *string) {
	if err := j.validateSetScmProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scmProvider",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetSslCaCertificate(val *string) {
	if err := j.validateSetSslCaCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCaCertificate",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference)SetTokenUri(val *string) {
	if err := j.validateSetTokenUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenUri",
		val,
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) PutServiceDirectoryConfig(value *GoogleDeveloperConnectAccountConnectorCustomOauthConfigServiceDirectoryConfig) {
	if err := g.validatePutServiceDirectoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceDirectoryConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ResetPkceDisabled() {
	_jsii_.InvokeVoid(
		g,
		"resetPkceDisabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ResetServiceDirectoryConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceDirectoryConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ResetSslCaCertificate() {
	_jsii_.InvokeVoid(
		g,
		"resetSslCaCertificate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDeveloperConnectAccountConnectorCustomOauthConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


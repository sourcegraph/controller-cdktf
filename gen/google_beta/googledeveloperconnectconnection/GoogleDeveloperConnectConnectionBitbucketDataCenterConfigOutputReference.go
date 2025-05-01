package googledeveloperconnectconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledeveloperconnectconnection/internal"
)

type GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference interface {
	cdktf.ComplexObject
	AuthorizerCredential() GoogleDeveloperConnectConnectionBitbucketDataCenterConfigAuthorizerCredentialOutputReference
	AuthorizerCredentialInput() *GoogleDeveloperConnectConnectionBitbucketDataCenterConfigAuthorizerCredential
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
	InternalValue() *GoogleDeveloperConnectConnectionBitbucketDataCenterConfig
	SetInternalValue(val *GoogleDeveloperConnectConnectionBitbucketDataCenterConfig)
	ReadAuthorizerCredential() GoogleDeveloperConnectConnectionBitbucketDataCenterConfigReadAuthorizerCredentialOutputReference
	ReadAuthorizerCredentialInput() *GoogleDeveloperConnectConnectionBitbucketDataCenterConfigReadAuthorizerCredential
	ServerVersion() *string
	ServiceDirectoryConfig() GoogleDeveloperConnectConnectionBitbucketDataCenterConfigServiceDirectoryConfigOutputReference
	ServiceDirectoryConfigInput() *GoogleDeveloperConnectConnectionBitbucketDataCenterConfigServiceDirectoryConfig
	SslCaCertificate() *string
	SetSslCaCertificate(val *string)
	SslCaCertificateInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WebhookSecretSecretVersion() *string
	SetWebhookSecretSecretVersion(val *string)
	WebhookSecretSecretVersionInput() *string
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
	PutAuthorizerCredential(value *GoogleDeveloperConnectConnectionBitbucketDataCenterConfigAuthorizerCredential)
	PutReadAuthorizerCredential(value *GoogleDeveloperConnectConnectionBitbucketDataCenterConfigReadAuthorizerCredential)
	PutServiceDirectoryConfig(value *GoogleDeveloperConnectConnectionBitbucketDataCenterConfigServiceDirectoryConfig)
	ResetServiceDirectoryConfig()
	ResetSslCaCertificate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference
type jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) AuthorizerCredential() GoogleDeveloperConnectConnectionBitbucketDataCenterConfigAuthorizerCredentialOutputReference {
	var returns GoogleDeveloperConnectConnectionBitbucketDataCenterConfigAuthorizerCredentialOutputReference
	_jsii_.Get(
		j,
		"authorizerCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) AuthorizerCredentialInput() *GoogleDeveloperConnectConnectionBitbucketDataCenterConfigAuthorizerCredential {
	var returns *GoogleDeveloperConnectConnectionBitbucketDataCenterConfigAuthorizerCredential
	_jsii_.Get(
		j,
		"authorizerCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) HostUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) HostUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) InternalValue() *GoogleDeveloperConnectConnectionBitbucketDataCenterConfig {
	var returns *GoogleDeveloperConnectConnectionBitbucketDataCenterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ReadAuthorizerCredential() GoogleDeveloperConnectConnectionBitbucketDataCenterConfigReadAuthorizerCredentialOutputReference {
	var returns GoogleDeveloperConnectConnectionBitbucketDataCenterConfigReadAuthorizerCredentialOutputReference
	_jsii_.Get(
		j,
		"readAuthorizerCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ReadAuthorizerCredentialInput() *GoogleDeveloperConnectConnectionBitbucketDataCenterConfigReadAuthorizerCredential {
	var returns *GoogleDeveloperConnectConnectionBitbucketDataCenterConfigReadAuthorizerCredential
	_jsii_.Get(
		j,
		"readAuthorizerCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ServiceDirectoryConfig() GoogleDeveloperConnectConnectionBitbucketDataCenterConfigServiceDirectoryConfigOutputReference {
	var returns GoogleDeveloperConnectConnectionBitbucketDataCenterConfigServiceDirectoryConfigOutputReference
	_jsii_.Get(
		j,
		"serviceDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ServiceDirectoryConfigInput() *GoogleDeveloperConnectConnectionBitbucketDataCenterConfigServiceDirectoryConfig {
	var returns *GoogleDeveloperConnectConnectionBitbucketDataCenterConfigServiceDirectoryConfig
	_jsii_.Get(
		j,
		"serviceDirectoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) SslCaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) SslCaCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) WebhookSecretSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookSecretSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) WebhookSecretSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookSecretSecretVersionInput",
		&returns,
	)
	return returns
}


func NewGoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDeveloperConnectConnection.GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference_Override(g GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDeveloperConnectConnection.GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference)SetHostUri(val *string) {
	if err := j.validateSetHostUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference)SetInternalValue(val *GoogleDeveloperConnectConnectionBitbucketDataCenterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference)SetSslCaCertificate(val *string) {
	if err := j.validateSetSslCaCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCaCertificate",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference)SetWebhookSecretSecretVersion(val *string) {
	if err := j.validateSetWebhookSecretSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhookSecretSecretVersion",
		val,
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) PutAuthorizerCredential(value *GoogleDeveloperConnectConnectionBitbucketDataCenterConfigAuthorizerCredential) {
	if err := g.validatePutAuthorizerCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthorizerCredential",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) PutReadAuthorizerCredential(value *GoogleDeveloperConnectConnectionBitbucketDataCenterConfigReadAuthorizerCredential) {
	if err := g.validatePutReadAuthorizerCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putReadAuthorizerCredential",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) PutServiceDirectoryConfig(value *GoogleDeveloperConnectConnectionBitbucketDataCenterConfigServiceDirectoryConfig) {
	if err := g.validatePutServiceDirectoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceDirectoryConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ResetServiceDirectoryConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceDirectoryConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ResetSslCaCertificate() {
	_jsii_.InvokeVoid(
		g,
		"resetSslCaCertificate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package googledeveloperconnectconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledeveloperconnectconnection/internal"
)

type GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference interface {
	cdktf.ComplexObject
	AppId() *string
	SetAppId(val *string)
	AppIdInput() *string
	AppInstallationId() *string
	SetAppInstallationId(val *string)
	AppInstallationIdInput() *string
	AppSlug() *string
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
	InstallationUri() *string
	InternalValue() *GoogleDeveloperConnectConnectionGithubEnterpriseConfig
	SetInternalValue(val *GoogleDeveloperConnectConnectionGithubEnterpriseConfig)
	PrivateKeySecretVersion() *string
	SetPrivateKeySecretVersion(val *string)
	PrivateKeySecretVersionInput() *string
	ServerVersion() *string
	ServiceDirectoryConfig() GoogleDeveloperConnectConnectionGithubEnterpriseConfigServiceDirectoryConfigOutputReference
	ServiceDirectoryConfigInput() *GoogleDeveloperConnectConnectionGithubEnterpriseConfigServiceDirectoryConfig
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
	PutServiceDirectoryConfig(value *GoogleDeveloperConnectConnectionGithubEnterpriseConfigServiceDirectoryConfig)
	ResetAppId()
	ResetAppInstallationId()
	ResetPrivateKeySecretVersion()
	ResetServiceDirectoryConfig()
	ResetSslCaCertificate()
	ResetWebhookSecretSecretVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference
type jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) AppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) AppInstallationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appInstallationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) AppInstallationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appInstallationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) AppSlug() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) HostUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) HostUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) InstallationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"installationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) InternalValue() *GoogleDeveloperConnectConnectionGithubEnterpriseConfig {
	var returns *GoogleDeveloperConnectConnectionGithubEnterpriseConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) PrivateKeySecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeySecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) PrivateKeySecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeySecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ServiceDirectoryConfig() GoogleDeveloperConnectConnectionGithubEnterpriseConfigServiceDirectoryConfigOutputReference {
	var returns GoogleDeveloperConnectConnectionGithubEnterpriseConfigServiceDirectoryConfigOutputReference
	_jsii_.Get(
		j,
		"serviceDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ServiceDirectoryConfigInput() *GoogleDeveloperConnectConnectionGithubEnterpriseConfigServiceDirectoryConfig {
	var returns *GoogleDeveloperConnectConnectionGithubEnterpriseConfigServiceDirectoryConfig
	_jsii_.Get(
		j,
		"serviceDirectoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) SslCaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) SslCaCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) WebhookSecretSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookSecretSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) WebhookSecretSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookSecretSecretVersionInput",
		&returns,
	)
	return returns
}


func NewGoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDeveloperConnectConnection.GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference_Override(g GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDeveloperConnectConnection.GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetAppId(val *string) {
	if err := j.validateSetAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetAppInstallationId(val *string) {
	if err := j.validateSetAppInstallationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appInstallationId",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetHostUri(val *string) {
	if err := j.validateSetHostUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetInternalValue(val *GoogleDeveloperConnectConnectionGithubEnterpriseConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetPrivateKeySecretVersion(val *string) {
	if err := j.validateSetPrivateKeySecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKeySecretVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetSslCaCertificate(val *string) {
	if err := j.validateSetSslCaCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCaCertificate",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference)SetWebhookSecretSecretVersion(val *string) {
	if err := j.validateSetWebhookSecretSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhookSecretSecretVersion",
		val,
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) PutServiceDirectoryConfig(value *GoogleDeveloperConnectConnectionGithubEnterpriseConfigServiceDirectoryConfig) {
	if err := g.validatePutServiceDirectoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceDirectoryConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ResetAppId() {
	_jsii_.InvokeVoid(
		g,
		"resetAppId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ResetAppInstallationId() {
	_jsii_.InvokeVoid(
		g,
		"resetAppInstallationId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ResetPrivateKeySecretVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateKeySecretVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ResetServiceDirectoryConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceDirectoryConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ResetSslCaCertificate() {
	_jsii_.InvokeVoid(
		g,
		"resetSslCaCertificate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ResetWebhookSecretSecretVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetWebhookSecretSecretVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


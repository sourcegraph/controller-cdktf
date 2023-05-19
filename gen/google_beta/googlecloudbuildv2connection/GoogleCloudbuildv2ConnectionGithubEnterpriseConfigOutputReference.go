package googlecloudbuildv2connection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudbuildv2connection/internal"
)

type GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference interface {
	cdktf.ComplexObject
	AppId() *float64
	SetAppId(val *float64)
	AppIdInput() *float64
	AppInstallationId() *float64
	SetAppInstallationId(val *float64)
	AppInstallationIdInput() *float64
	AppSlug() *string
	SetAppSlug(val *string)
	AppSlugInput() *string
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
	InternalValue() *GoogleCloudbuildv2ConnectionGithubEnterpriseConfig
	SetInternalValue(val *GoogleCloudbuildv2ConnectionGithubEnterpriseConfig)
	PrivateKeySecretVersion() *string
	SetPrivateKeySecretVersion(val *string)
	PrivateKeySecretVersionInput() *string
	ServiceDirectoryConfig() GoogleCloudbuildv2ConnectionGithubEnterpriseConfigServiceDirectoryConfigOutputReference
	ServiceDirectoryConfigInput() *GoogleCloudbuildv2ConnectionGithubEnterpriseConfigServiceDirectoryConfig
	SslCa() *string
	SetSslCa(val *string)
	SslCaInput() *string
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
	PutServiceDirectoryConfig(value *GoogleCloudbuildv2ConnectionGithubEnterpriseConfigServiceDirectoryConfig)
	ResetAppId()
	ResetAppInstallationId()
	ResetAppSlug()
	ResetPrivateKeySecretVersion()
	ResetServiceDirectoryConfig()
	ResetSslCa()
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

// The jsii proxy struct for GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference
type jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) AppId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) AppIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"appIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) AppInstallationId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"appInstallationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) AppInstallationIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"appInstallationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) AppSlug() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) AppSlugInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSlugInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) HostUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) HostUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) InternalValue() *GoogleCloudbuildv2ConnectionGithubEnterpriseConfig {
	var returns *GoogleCloudbuildv2ConnectionGithubEnterpriseConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) PrivateKeySecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeySecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) PrivateKeySecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeySecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) ServiceDirectoryConfig() GoogleCloudbuildv2ConnectionGithubEnterpriseConfigServiceDirectoryConfigOutputReference {
	var returns GoogleCloudbuildv2ConnectionGithubEnterpriseConfigServiceDirectoryConfigOutputReference
	_jsii_.Get(
		j,
		"serviceDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) ServiceDirectoryConfigInput() *GoogleCloudbuildv2ConnectionGithubEnterpriseConfigServiceDirectoryConfig {
	var returns *GoogleCloudbuildv2ConnectionGithubEnterpriseConfigServiceDirectoryConfig
	_jsii_.Get(
		j,
		"serviceDirectoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) SslCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) SslCaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) WebhookSecretSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookSecretSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) WebhookSecretSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookSecretSecretVersionInput",
		&returns,
	)
	return returns
}


func NewGoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleCloudbuildv2Connection.GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference_Override(g GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleCloudbuildv2Connection.GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference)SetAppId(val *float64) {
	if err := j.validateSetAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference)SetAppInstallationId(val *float64) {
	if err := j.validateSetAppInstallationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appInstallationId",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference)SetAppSlug(val *string) {
	if err := j.validateSetAppSlugParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appSlug",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference)SetHostUri(val *string) {
	if err := j.validateSetHostUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostUri",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference)SetInternalValue(val *GoogleCloudbuildv2ConnectionGithubEnterpriseConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference)SetPrivateKeySecretVersion(val *string) {
	if err := j.validateSetPrivateKeySecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKeySecretVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference)SetSslCa(val *string) {
	if err := j.validateSetSslCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCa",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference)SetWebhookSecretSecretVersion(val *string) {
	if err := j.validateSetWebhookSecretSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhookSecretSecretVersion",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) PutServiceDirectoryConfig(value *GoogleCloudbuildv2ConnectionGithubEnterpriseConfigServiceDirectoryConfig) {
	if err := g.validatePutServiceDirectoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceDirectoryConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) ResetAppId() {
	_jsii_.InvokeVoid(
		g,
		"resetAppId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) ResetAppInstallationId() {
	_jsii_.InvokeVoid(
		g,
		"resetAppInstallationId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) ResetAppSlug() {
	_jsii_.InvokeVoid(
		g,
		"resetAppSlug",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) ResetPrivateKeySecretVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateKeySecretVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) ResetServiceDirectoryConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceDirectoryConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) ResetSslCa() {
	_jsii_.InvokeVoid(
		g,
		"resetSslCa",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) ResetWebhookSecretSecretVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetWebhookSecretSecretVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


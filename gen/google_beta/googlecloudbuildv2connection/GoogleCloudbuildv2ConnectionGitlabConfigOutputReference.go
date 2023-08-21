package googlecloudbuildv2connection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudbuildv2connection/internal"
)

type GoogleCloudbuildv2ConnectionGitlabConfigOutputReference interface {
	cdktf.ComplexObject
	AuthorizerCredential() GoogleCloudbuildv2ConnectionGitlabConfigAuthorizerCredentialOutputReference
	AuthorizerCredentialInput() *GoogleCloudbuildv2ConnectionGitlabConfigAuthorizerCredential
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
	InternalValue() *GoogleCloudbuildv2ConnectionGitlabConfig
	SetInternalValue(val *GoogleCloudbuildv2ConnectionGitlabConfig)
	ReadAuthorizerCredential() GoogleCloudbuildv2ConnectionGitlabConfigReadAuthorizerCredentialOutputReference
	ReadAuthorizerCredentialInput() *GoogleCloudbuildv2ConnectionGitlabConfigReadAuthorizerCredential
	ServerVersion() *string
	ServiceDirectoryConfig() GoogleCloudbuildv2ConnectionGitlabConfigServiceDirectoryConfigOutputReference
	ServiceDirectoryConfigInput() *GoogleCloudbuildv2ConnectionGitlabConfigServiceDirectoryConfig
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
	PutAuthorizerCredential(value *GoogleCloudbuildv2ConnectionGitlabConfigAuthorizerCredential)
	PutReadAuthorizerCredential(value *GoogleCloudbuildv2ConnectionGitlabConfigReadAuthorizerCredential)
	PutServiceDirectoryConfig(value *GoogleCloudbuildv2ConnectionGitlabConfigServiceDirectoryConfig)
	ResetHostUri()
	ResetServiceDirectoryConfig()
	ResetSslCa()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudbuildv2ConnectionGitlabConfigOutputReference
type jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) AuthorizerCredential() GoogleCloudbuildv2ConnectionGitlabConfigAuthorizerCredentialOutputReference {
	var returns GoogleCloudbuildv2ConnectionGitlabConfigAuthorizerCredentialOutputReference
	_jsii_.Get(
		j,
		"authorizerCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) AuthorizerCredentialInput() *GoogleCloudbuildv2ConnectionGitlabConfigAuthorizerCredential {
	var returns *GoogleCloudbuildv2ConnectionGitlabConfigAuthorizerCredential
	_jsii_.Get(
		j,
		"authorizerCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) HostUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) HostUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) InternalValue() *GoogleCloudbuildv2ConnectionGitlabConfig {
	var returns *GoogleCloudbuildv2ConnectionGitlabConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) ReadAuthorizerCredential() GoogleCloudbuildv2ConnectionGitlabConfigReadAuthorizerCredentialOutputReference {
	var returns GoogleCloudbuildv2ConnectionGitlabConfigReadAuthorizerCredentialOutputReference
	_jsii_.Get(
		j,
		"readAuthorizerCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) ReadAuthorizerCredentialInput() *GoogleCloudbuildv2ConnectionGitlabConfigReadAuthorizerCredential {
	var returns *GoogleCloudbuildv2ConnectionGitlabConfigReadAuthorizerCredential
	_jsii_.Get(
		j,
		"readAuthorizerCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) ServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) ServiceDirectoryConfig() GoogleCloudbuildv2ConnectionGitlabConfigServiceDirectoryConfigOutputReference {
	var returns GoogleCloudbuildv2ConnectionGitlabConfigServiceDirectoryConfigOutputReference
	_jsii_.Get(
		j,
		"serviceDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) ServiceDirectoryConfigInput() *GoogleCloudbuildv2ConnectionGitlabConfigServiceDirectoryConfig {
	var returns *GoogleCloudbuildv2ConnectionGitlabConfigServiceDirectoryConfig
	_jsii_.Get(
		j,
		"serviceDirectoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) SslCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) SslCaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) WebhookSecretSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookSecretSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) WebhookSecretSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookSecretSecretVersionInput",
		&returns,
	)
	return returns
}


func NewGoogleCloudbuildv2ConnectionGitlabConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudbuildv2ConnectionGitlabConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudbuildv2ConnectionGitlabConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudbuildv2Connection.GoogleCloudbuildv2ConnectionGitlabConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudbuildv2ConnectionGitlabConfigOutputReference_Override(g GoogleCloudbuildv2ConnectionGitlabConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudbuildv2Connection.GoogleCloudbuildv2ConnectionGitlabConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference)SetHostUri(val *string) {
	if err := j.validateSetHostUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostUri",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference)SetInternalValue(val *GoogleCloudbuildv2ConnectionGitlabConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference)SetSslCa(val *string) {
	if err := j.validateSetSslCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCa",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference)SetWebhookSecretSecretVersion(val *string) {
	if err := j.validateSetWebhookSecretSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhookSecretSecretVersion",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) PutAuthorizerCredential(value *GoogleCloudbuildv2ConnectionGitlabConfigAuthorizerCredential) {
	if err := g.validatePutAuthorizerCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthorizerCredential",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) PutReadAuthorizerCredential(value *GoogleCloudbuildv2ConnectionGitlabConfigReadAuthorizerCredential) {
	if err := g.validatePutReadAuthorizerCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putReadAuthorizerCredential",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) PutServiceDirectoryConfig(value *GoogleCloudbuildv2ConnectionGitlabConfigServiceDirectoryConfig) {
	if err := g.validatePutServiceDirectoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceDirectoryConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) ResetHostUri() {
	_jsii_.InvokeVoid(
		g,
		"resetHostUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) ResetServiceDirectoryConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceDirectoryConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) ResetSslCa() {
	_jsii_.InvokeVoid(
		g,
		"resetSslCa",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudbuildv2ConnectionGitlabConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package googledeveloperconnectconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledeveloperconnectconnection/internal"
)

type GoogleDeveloperConnectConnectionHttpConfigOutputReference interface {
	cdktf.ComplexObject
	BasicAuthentication() GoogleDeveloperConnectConnectionHttpConfigBasicAuthenticationOutputReference
	BasicAuthenticationInput() *GoogleDeveloperConnectConnectionHttpConfigBasicAuthentication
	BearerTokenAuthentication() GoogleDeveloperConnectConnectionHttpConfigBearerTokenAuthenticationOutputReference
	BearerTokenAuthenticationInput() *GoogleDeveloperConnectConnectionHttpConfigBearerTokenAuthentication
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
	InternalValue() *GoogleDeveloperConnectConnectionHttpConfig
	SetInternalValue(val *GoogleDeveloperConnectConnectionHttpConfig)
	ServiceDirectoryConfig() GoogleDeveloperConnectConnectionHttpConfigServiceDirectoryConfigOutputReference
	ServiceDirectoryConfigInput() *GoogleDeveloperConnectConnectionHttpConfigServiceDirectoryConfig
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
	PutBasicAuthentication(value *GoogleDeveloperConnectConnectionHttpConfigBasicAuthentication)
	PutBearerTokenAuthentication(value *GoogleDeveloperConnectConnectionHttpConfigBearerTokenAuthentication)
	PutServiceDirectoryConfig(value *GoogleDeveloperConnectConnectionHttpConfigServiceDirectoryConfig)
	ResetBasicAuthentication()
	ResetBearerTokenAuthentication()
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

// The jsii proxy struct for GoogleDeveloperConnectConnectionHttpConfigOutputReference
type jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) BasicAuthentication() GoogleDeveloperConnectConnectionHttpConfigBasicAuthenticationOutputReference {
	var returns GoogleDeveloperConnectConnectionHttpConfigBasicAuthenticationOutputReference
	_jsii_.Get(
		j,
		"basicAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) BasicAuthenticationInput() *GoogleDeveloperConnectConnectionHttpConfigBasicAuthentication {
	var returns *GoogleDeveloperConnectConnectionHttpConfigBasicAuthentication
	_jsii_.Get(
		j,
		"basicAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) BearerTokenAuthentication() GoogleDeveloperConnectConnectionHttpConfigBearerTokenAuthenticationOutputReference {
	var returns GoogleDeveloperConnectConnectionHttpConfigBearerTokenAuthenticationOutputReference
	_jsii_.Get(
		j,
		"bearerTokenAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) BearerTokenAuthenticationInput() *GoogleDeveloperConnectConnectionHttpConfigBearerTokenAuthentication {
	var returns *GoogleDeveloperConnectConnectionHttpConfigBearerTokenAuthentication
	_jsii_.Get(
		j,
		"bearerTokenAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) HostUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) HostUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) InternalValue() *GoogleDeveloperConnectConnectionHttpConfig {
	var returns *GoogleDeveloperConnectConnectionHttpConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) ServiceDirectoryConfig() GoogleDeveloperConnectConnectionHttpConfigServiceDirectoryConfigOutputReference {
	var returns GoogleDeveloperConnectConnectionHttpConfigServiceDirectoryConfigOutputReference
	_jsii_.Get(
		j,
		"serviceDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) ServiceDirectoryConfigInput() *GoogleDeveloperConnectConnectionHttpConfigServiceDirectoryConfig {
	var returns *GoogleDeveloperConnectConnectionHttpConfigServiceDirectoryConfig
	_jsii_.Get(
		j,
		"serviceDirectoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) SslCaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) SslCaCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDeveloperConnectConnectionHttpConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDeveloperConnectConnectionHttpConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDeveloperConnectConnectionHttpConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDeveloperConnectConnection.GoogleDeveloperConnectConnectionHttpConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDeveloperConnectConnectionHttpConfigOutputReference_Override(g GoogleDeveloperConnectConnectionHttpConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDeveloperConnectConnection.GoogleDeveloperConnectConnectionHttpConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference)SetHostUri(val *string) {
	if err := j.validateSetHostUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference)SetInternalValue(val *GoogleDeveloperConnectConnectionHttpConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference)SetSslCaCertificate(val *string) {
	if err := j.validateSetSslCaCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCaCertificate",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) PutBasicAuthentication(value *GoogleDeveloperConnectConnectionHttpConfigBasicAuthentication) {
	if err := g.validatePutBasicAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBasicAuthentication",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) PutBearerTokenAuthentication(value *GoogleDeveloperConnectConnectionHttpConfigBearerTokenAuthentication) {
	if err := g.validatePutBearerTokenAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBearerTokenAuthentication",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) PutServiceDirectoryConfig(value *GoogleDeveloperConnectConnectionHttpConfigServiceDirectoryConfig) {
	if err := g.validatePutServiceDirectoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceDirectoryConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) ResetBasicAuthentication() {
	_jsii_.InvokeVoid(
		g,
		"resetBasicAuthentication",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) ResetBearerTokenAuthentication() {
	_jsii_.InvokeVoid(
		g,
		"resetBearerTokenAuthentication",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) ResetServiceDirectoryConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceDirectoryConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) ResetSslCaCertificate() {
	_jsii_.InvokeVoid(
		g,
		"resetSslCaCertificate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionHttpConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


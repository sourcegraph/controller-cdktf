package developerconnectconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/developerconnectconnection/internal"
)

type DeveloperConnectConnectionHttpConfigOutputReference interface {
	cdktf.ComplexObject
	BasicAuthentication() DeveloperConnectConnectionHttpConfigBasicAuthenticationOutputReference
	BasicAuthenticationInput() *DeveloperConnectConnectionHttpConfigBasicAuthentication
	BearerTokenAuthentication() DeveloperConnectConnectionHttpConfigBearerTokenAuthenticationOutputReference
	BearerTokenAuthenticationInput() *DeveloperConnectConnectionHttpConfigBearerTokenAuthentication
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
	InternalValue() *DeveloperConnectConnectionHttpConfig
	SetInternalValue(val *DeveloperConnectConnectionHttpConfig)
	ServiceDirectoryConfig() DeveloperConnectConnectionHttpConfigServiceDirectoryConfigOutputReference
	ServiceDirectoryConfigInput() *DeveloperConnectConnectionHttpConfigServiceDirectoryConfig
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
	PutBasicAuthentication(value *DeveloperConnectConnectionHttpConfigBasicAuthentication)
	PutBearerTokenAuthentication(value *DeveloperConnectConnectionHttpConfigBearerTokenAuthentication)
	PutServiceDirectoryConfig(value *DeveloperConnectConnectionHttpConfigServiceDirectoryConfig)
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

// The jsii proxy struct for DeveloperConnectConnectionHttpConfigOutputReference
type jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) BasicAuthentication() DeveloperConnectConnectionHttpConfigBasicAuthenticationOutputReference {
	var returns DeveloperConnectConnectionHttpConfigBasicAuthenticationOutputReference
	_jsii_.Get(
		j,
		"basicAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) BasicAuthenticationInput() *DeveloperConnectConnectionHttpConfigBasicAuthentication {
	var returns *DeveloperConnectConnectionHttpConfigBasicAuthentication
	_jsii_.Get(
		j,
		"basicAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) BearerTokenAuthentication() DeveloperConnectConnectionHttpConfigBearerTokenAuthenticationOutputReference {
	var returns DeveloperConnectConnectionHttpConfigBearerTokenAuthenticationOutputReference
	_jsii_.Get(
		j,
		"bearerTokenAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) BearerTokenAuthenticationInput() *DeveloperConnectConnectionHttpConfigBearerTokenAuthentication {
	var returns *DeveloperConnectConnectionHttpConfigBearerTokenAuthentication
	_jsii_.Get(
		j,
		"bearerTokenAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) HostUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) HostUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) InternalValue() *DeveloperConnectConnectionHttpConfig {
	var returns *DeveloperConnectConnectionHttpConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) ServiceDirectoryConfig() DeveloperConnectConnectionHttpConfigServiceDirectoryConfigOutputReference {
	var returns DeveloperConnectConnectionHttpConfigServiceDirectoryConfigOutputReference
	_jsii_.Get(
		j,
		"serviceDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) ServiceDirectoryConfigInput() *DeveloperConnectConnectionHttpConfigServiceDirectoryConfig {
	var returns *DeveloperConnectConnectionHttpConfigServiceDirectoryConfig
	_jsii_.Get(
		j,
		"serviceDirectoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) SslCaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) SslCaCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDeveloperConnectConnectionHttpConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DeveloperConnectConnectionHttpConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDeveloperConnectConnectionHttpConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.developerConnectConnection.DeveloperConnectConnectionHttpConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDeveloperConnectConnectionHttpConfigOutputReference_Override(d DeveloperConnectConnectionHttpConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.developerConnectConnection.DeveloperConnectConnectionHttpConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference)SetHostUri(val *string) {
	if err := j.validateSetHostUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostUri",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference)SetInternalValue(val *DeveloperConnectConnectionHttpConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference)SetSslCaCertificate(val *string) {
	if err := j.validateSetSslCaCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCaCertificate",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) PutBasicAuthentication(value *DeveloperConnectConnectionHttpConfigBasicAuthentication) {
	if err := d.validatePutBasicAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBasicAuthentication",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) PutBearerTokenAuthentication(value *DeveloperConnectConnectionHttpConfigBearerTokenAuthentication) {
	if err := d.validatePutBearerTokenAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBearerTokenAuthentication",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) PutServiceDirectoryConfig(value *DeveloperConnectConnectionHttpConfigServiceDirectoryConfig) {
	if err := d.validatePutServiceDirectoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServiceDirectoryConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) ResetBasicAuthentication() {
	_jsii_.InvokeVoid(
		d,
		"resetBasicAuthentication",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) ResetBearerTokenAuthentication() {
	_jsii_.InvokeVoid(
		d,
		"resetBearerTokenAuthentication",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) ResetServiceDirectoryConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceDirectoryConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) ResetSslCaCertificate() {
	_jsii_.InvokeVoid(
		d,
		"resetSslCaCertificate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnectionHttpConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


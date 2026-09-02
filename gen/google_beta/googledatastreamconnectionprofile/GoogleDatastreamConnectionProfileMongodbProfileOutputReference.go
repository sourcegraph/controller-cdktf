package googledatastreamconnectionprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatastreamconnectionprofile/internal"
)

type GoogleDatastreamConnectionProfileMongodbProfileOutputReference interface {
	cdktn.ComplexObject
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
	HostAddresses() GoogleDatastreamConnectionProfileMongodbProfileHostAddressesList
	HostAddressesInput() interface{}
	InternalValue() *GoogleDatastreamConnectionProfileMongodbProfile
	SetInternalValue(val *GoogleDatastreamConnectionProfileMongodbProfile)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	ReplicaSet() *string
	SetReplicaSet(val *string)
	ReplicaSetInput() *string
	SecretManagerStoredPassword() *string
	SetSecretManagerStoredPassword(val *string)
	SecretManagerStoredPasswordInput() *string
	SrvConnectionFormat() GoogleDatastreamConnectionProfileMongodbProfileSrvConnectionFormatOutputReference
	SrvConnectionFormatInput() *GoogleDatastreamConnectionProfileMongodbProfileSrvConnectionFormat
	SslConfig() GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference
	SslConfigInput() *GoogleDatastreamConnectionProfileMongodbProfileSslConfig
	StandardConnectionFormat() GoogleDatastreamConnectionProfileMongodbProfileStandardConnectionFormatOutputReference
	StandardConnectionFormatInput() *GoogleDatastreamConnectionProfileMongodbProfileStandardConnectionFormat
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	PutHostAddresses(value interface{})
	PutSrvConnectionFormat(value *GoogleDatastreamConnectionProfileMongodbProfileSrvConnectionFormat)
	PutSslConfig(value *GoogleDatastreamConnectionProfileMongodbProfileSslConfig)
	PutStandardConnectionFormat(value *GoogleDatastreamConnectionProfileMongodbProfileStandardConnectionFormat)
	ResetPassword()
	ResetReplicaSet()
	ResetSecretManagerStoredPassword()
	ResetSrvConnectionFormat()
	ResetSslConfig()
	ResetStandardConnectionFormat()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDatastreamConnectionProfileMongodbProfileOutputReference
type jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) HostAddresses() GoogleDatastreamConnectionProfileMongodbProfileHostAddressesList {
	var returns GoogleDatastreamConnectionProfileMongodbProfileHostAddressesList
	_jsii_.Get(
		j,
		"hostAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) HostAddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) InternalValue() *GoogleDatastreamConnectionProfileMongodbProfile {
	var returns *GoogleDatastreamConnectionProfileMongodbProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) ReplicaSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) ReplicaSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) SecretManagerStoredPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretManagerStoredPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) SecretManagerStoredPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretManagerStoredPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) SrvConnectionFormat() GoogleDatastreamConnectionProfileMongodbProfileSrvConnectionFormatOutputReference {
	var returns GoogleDatastreamConnectionProfileMongodbProfileSrvConnectionFormatOutputReference
	_jsii_.Get(
		j,
		"srvConnectionFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) SrvConnectionFormatInput() *GoogleDatastreamConnectionProfileMongodbProfileSrvConnectionFormat {
	var returns *GoogleDatastreamConnectionProfileMongodbProfileSrvConnectionFormat
	_jsii_.Get(
		j,
		"srvConnectionFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) SslConfig() GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference {
	var returns GoogleDatastreamConnectionProfileMongodbProfileSslConfigOutputReference
	_jsii_.Get(
		j,
		"sslConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) SslConfigInput() *GoogleDatastreamConnectionProfileMongodbProfileSslConfig {
	var returns *GoogleDatastreamConnectionProfileMongodbProfileSslConfig
	_jsii_.Get(
		j,
		"sslConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) StandardConnectionFormat() GoogleDatastreamConnectionProfileMongodbProfileStandardConnectionFormatOutputReference {
	var returns GoogleDatastreamConnectionProfileMongodbProfileStandardConnectionFormatOutputReference
	_jsii_.Get(
		j,
		"standardConnectionFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) StandardConnectionFormatInput() *GoogleDatastreamConnectionProfileMongodbProfileStandardConnectionFormat {
	var returns *GoogleDatastreamConnectionProfileMongodbProfileStandardConnectionFormat
	_jsii_.Get(
		j,
		"standardConnectionFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewGoogleDatastreamConnectionProfileMongodbProfileOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDatastreamConnectionProfileMongodbProfileOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDatastreamConnectionProfileMongodbProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDatastreamConnectionProfile.GoogleDatastreamConnectionProfileMongodbProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDatastreamConnectionProfileMongodbProfileOutputReference_Override(g GoogleDatastreamConnectionProfileMongodbProfileOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDatastreamConnectionProfile.GoogleDatastreamConnectionProfileMongodbProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference)SetInternalValue(val *GoogleDatastreamConnectionProfileMongodbProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference)SetReplicaSet(val *string) {
	if err := j.validateSetReplicaSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference)SetSecretManagerStoredPassword(val *string) {
	if err := j.validateSetSecretManagerStoredPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretManagerStoredPassword",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) PutHostAddresses(value interface{}) {
	if err := g.validatePutHostAddressesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHostAddresses",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) PutSrvConnectionFormat(value *GoogleDatastreamConnectionProfileMongodbProfileSrvConnectionFormat) {
	if err := g.validatePutSrvConnectionFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSrvConnectionFormat",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) PutSslConfig(value *GoogleDatastreamConnectionProfileMongodbProfileSslConfig) {
	if err := g.validatePutSslConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSslConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) PutStandardConnectionFormat(value *GoogleDatastreamConnectionProfileMongodbProfileStandardConnectionFormat) {
	if err := g.validatePutStandardConnectionFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStandardConnectionFormat",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		g,
		"resetPassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) ResetReplicaSet() {
	_jsii_.InvokeVoid(
		g,
		"resetReplicaSet",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) ResetSecretManagerStoredPassword() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretManagerStoredPassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) ResetSrvConnectionFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetSrvConnectionFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) ResetSslConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSslConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) ResetStandardConnectionFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetStandardConnectionFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileMongodbProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


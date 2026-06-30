package datastreamconnectionprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/datastreamconnectionprofile/internal"
)

type DatastreamConnectionProfileMongodbProfileOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	HostAddresses() DatastreamConnectionProfileMongodbProfileHostAddressesList
	HostAddressesInput() any
	InternalValue() *DatastreamConnectionProfileMongodbProfile
	SetInternalValue(val *DatastreamConnectionProfileMongodbProfile)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	ReplicaSet() *string
	SetReplicaSet(val *string)
	ReplicaSetInput() *string
	SecretManagerStoredPassword() *string
	SetSecretManagerStoredPassword(val *string)
	SecretManagerStoredPasswordInput() *string
	SrvConnectionFormat() DatastreamConnectionProfileMongodbProfileSrvConnectionFormatOutputReference
	SrvConnectionFormatInput() *DatastreamConnectionProfileMongodbProfileSrvConnectionFormat
	SslConfig() DatastreamConnectionProfileMongodbProfileSslConfigOutputReference
	SslConfigInput() *DatastreamConnectionProfileMongodbProfileSslConfig
	StandardConnectionFormat() DatastreamConnectionProfileMongodbProfileStandardConnectionFormatOutputReference
	StandardConnectionFormatInput() *DatastreamConnectionProfileMongodbProfileStandardConnectionFormat
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutHostAddresses(value any)
	PutSrvConnectionFormat(value *DatastreamConnectionProfileMongodbProfileSrvConnectionFormat)
	PutSslConfig(value *DatastreamConnectionProfileMongodbProfileSslConfig)
	PutStandardConnectionFormat(value *DatastreamConnectionProfileMongodbProfileStandardConnectionFormat)
	ResetPassword()
	ResetReplicaSet()
	ResetSecretManagerStoredPassword()
	ResetSrvConnectionFormat()
	ResetSslConfig()
	ResetStandardConnectionFormat()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatastreamConnectionProfileMongodbProfileOutputReference
type jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) HostAddresses() DatastreamConnectionProfileMongodbProfileHostAddressesList {
	var returns DatastreamConnectionProfileMongodbProfileHostAddressesList
	_jsii_.Get(
		j,
		"hostAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) HostAddressesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) InternalValue() *DatastreamConnectionProfileMongodbProfile {
	var returns *DatastreamConnectionProfileMongodbProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) ReplicaSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) ReplicaSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) SecretManagerStoredPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretManagerStoredPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) SecretManagerStoredPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretManagerStoredPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) SrvConnectionFormat() DatastreamConnectionProfileMongodbProfileSrvConnectionFormatOutputReference {
	var returns DatastreamConnectionProfileMongodbProfileSrvConnectionFormatOutputReference
	_jsii_.Get(
		j,
		"srvConnectionFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) SrvConnectionFormatInput() *DatastreamConnectionProfileMongodbProfileSrvConnectionFormat {
	var returns *DatastreamConnectionProfileMongodbProfileSrvConnectionFormat
	_jsii_.Get(
		j,
		"srvConnectionFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) SslConfig() DatastreamConnectionProfileMongodbProfileSslConfigOutputReference {
	var returns DatastreamConnectionProfileMongodbProfileSslConfigOutputReference
	_jsii_.Get(
		j,
		"sslConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) SslConfigInput() *DatastreamConnectionProfileMongodbProfileSslConfig {
	var returns *DatastreamConnectionProfileMongodbProfileSslConfig
	_jsii_.Get(
		j,
		"sslConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) StandardConnectionFormat() DatastreamConnectionProfileMongodbProfileStandardConnectionFormatOutputReference {
	var returns DatastreamConnectionProfileMongodbProfileStandardConnectionFormatOutputReference
	_jsii_.Get(
		j,
		"standardConnectionFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) StandardConnectionFormatInput() *DatastreamConnectionProfileMongodbProfileStandardConnectionFormat {
	var returns *DatastreamConnectionProfileMongodbProfileStandardConnectionFormat
	_jsii_.Get(
		j,
		"standardConnectionFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func NewDatastreamConnectionProfileMongodbProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatastreamConnectionProfileMongodbProfileOutputReference {
	_init_.Initialize()

	if err := validateNewDatastreamConnectionProfileMongodbProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.datastreamConnectionProfile.DatastreamConnectionProfileMongodbProfileOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatastreamConnectionProfileMongodbProfileOutputReference_Override(d DatastreamConnectionProfileMongodbProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.datastreamConnectionProfile.DatastreamConnectionProfileMongodbProfileOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) SetInternalValue(val *DatastreamConnectionProfileMongodbProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) SetReplicaSet(val *string) {
	if err := j.validateSetReplicaSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaSet",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) SetSecretManagerStoredPassword(val *string) {
	if err := j.validateSetSecretManagerStoredPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretManagerStoredPassword",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) PutHostAddresses(value any) {
	if err := d.validatePutHostAddressesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHostAddresses",
		[]any{value},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) PutSrvConnectionFormat(value *DatastreamConnectionProfileMongodbProfileSrvConnectionFormat) {
	if err := d.validatePutSrvConnectionFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSrvConnectionFormat",
		[]any{value},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) PutSslConfig(value *DatastreamConnectionProfileMongodbProfileSslConfig) {
	if err := d.validatePutSslConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSslConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) PutStandardConnectionFormat(value *DatastreamConnectionProfileMongodbProfileStandardConnectionFormat) {
	if err := d.validatePutStandardConnectionFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStandardConnectionFormat",
		[]any{value},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) ResetReplicaSet() {
	_jsii_.InvokeVoid(
		d,
		"resetReplicaSet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) ResetSecretManagerStoredPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretManagerStoredPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) ResetSrvConnectionFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetSrvConnectionFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) ResetSslConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetSslConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) ResetStandardConnectionFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetStandardConnectionFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		d,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfileMongodbProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

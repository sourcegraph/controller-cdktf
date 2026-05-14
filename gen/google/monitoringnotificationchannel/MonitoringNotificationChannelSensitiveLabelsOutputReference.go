package monitoringnotificationchannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/monitoringnotificationchannel/internal"
)

type MonitoringNotificationChannelSensitiveLabelsOutputReference interface {
	cdktf.ComplexObject
	AuthToken() *string
	SetAuthToken(val *string)
	AuthTokenInput() *string
	AuthTokenWo() *string
	SetAuthTokenWo(val *string)
	AuthTokenWoInput() *string
	AuthTokenWoVersion() *string
	SetAuthTokenWoVersion(val *string)
	AuthTokenWoVersionInput() *string
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
	InternalValue() *MonitoringNotificationChannelSensitiveLabels
	SetInternalValue(val *MonitoringNotificationChannelSensitiveLabels)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	PasswordWo() *string
	SetPasswordWo(val *string)
	PasswordWoInput() *string
	PasswordWoVersion() *string
	SetPasswordWoVersion(val *string)
	PasswordWoVersionInput() *string
	ServiceKey() *string
	SetServiceKey(val *string)
	ServiceKeyInput() *string
	ServiceKeyWo() *string
	SetServiceKeyWo(val *string)
	ServiceKeyWoInput() *string
	ServiceKeyWoVersion() *string
	SetServiceKeyWoVersion(val *string)
	ServiceKeyWoVersionInput() *string
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
	ResetAuthToken()
	ResetAuthTokenWo()
	ResetAuthTokenWoVersion()
	ResetPassword()
	ResetPasswordWo()
	ResetPasswordWoVersion()
	ResetServiceKey()
	ResetServiceKeyWo()
	ResetServiceKeyWoVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitoringNotificationChannelSensitiveLabelsOutputReference
type jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) AuthToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) AuthTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) AuthTokenWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTokenWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) AuthTokenWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTokenWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) AuthTokenWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTokenWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) AuthTokenWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTokenWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) InternalValue() *MonitoringNotificationChannelSensitiveLabels {
	var returns *MonitoringNotificationChannelSensitiveLabels
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) PasswordWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) PasswordWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) PasswordWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) PasswordWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) ServiceKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) ServiceKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) ServiceKeyWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKeyWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) ServiceKeyWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKeyWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) ServiceKeyWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKeyWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) ServiceKeyWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKeyWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitoringNotificationChannelSensitiveLabelsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitoringNotificationChannelSensitiveLabelsOutputReference {
	_init_.Initialize()

	if err := validateNewMonitoringNotificationChannelSensitiveLabelsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.monitoringNotificationChannel.MonitoringNotificationChannelSensitiveLabelsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitoringNotificationChannelSensitiveLabelsOutputReference_Override(m MonitoringNotificationChannelSensitiveLabelsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.monitoringNotificationChannel.MonitoringNotificationChannelSensitiveLabelsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference)SetAuthToken(val *string) {
	if err := j.validateSetAuthTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authToken",
		val,
	)
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference)SetAuthTokenWo(val *string) {
	if err := j.validateSetAuthTokenWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authTokenWo",
		val,
	)
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference)SetAuthTokenWoVersion(val *string) {
	if err := j.validateSetAuthTokenWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authTokenWoVersion",
		val,
	)
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference)SetInternalValue(val *MonitoringNotificationChannelSensitiveLabels) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference)SetPasswordWo(val *string) {
	if err := j.validateSetPasswordWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordWo",
		val,
	)
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference)SetPasswordWoVersion(val *string) {
	if err := j.validateSetPasswordWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordWoVersion",
		val,
	)
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference)SetServiceKey(val *string) {
	if err := j.validateSetServiceKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceKey",
		val,
	)
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference)SetServiceKeyWo(val *string) {
	if err := j.validateSetServiceKeyWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceKeyWo",
		val,
	)
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference)SetServiceKeyWoVersion(val *string) {
	if err := j.validateSetServiceKeyWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceKeyWoVersion",
		val,
	)
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) ResetAuthToken() {
	_jsii_.InvokeVoid(
		m,
		"resetAuthToken",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) ResetAuthTokenWo() {
	_jsii_.InvokeVoid(
		m,
		"resetAuthTokenWo",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) ResetAuthTokenWoVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetAuthTokenWoVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		m,
		"resetPassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) ResetPasswordWo() {
	_jsii_.InvokeVoid(
		m,
		"resetPasswordWo",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) ResetPasswordWoVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetPasswordWoVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) ResetServiceKey() {
	_jsii_.InvokeVoid(
		m,
		"resetServiceKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) ResetServiceKeyWo() {
	_jsii_.InvokeVoid(
		m,
		"resetServiceKeyWo",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) ResetServiceKeyWoVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetServiceKeyWoVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringNotificationChannelSensitiveLabelsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package googlemonitoringnotificationchannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlemonitoringnotificationchannel/internal"
)

type GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() *GoogleMonitoringNotificationChannelSensitiveLabels
	SetInternalValue(val *GoogleMonitoringNotificationChannelSensitiveLabels)
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
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference
type jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) AuthToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) AuthTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) AuthTokenWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTokenWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) AuthTokenWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTokenWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) AuthTokenWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTokenWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) AuthTokenWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTokenWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) InternalValue() *GoogleMonitoringNotificationChannelSensitiveLabels {
	var returns *GoogleMonitoringNotificationChannelSensitiveLabels
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) PasswordWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) PasswordWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) PasswordWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) PasswordWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) ServiceKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) ServiceKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) ServiceKeyWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKeyWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) ServiceKeyWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKeyWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) ServiceKeyWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKeyWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) ServiceKeyWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKeyWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleMonitoringNotificationChannelSensitiveLabelsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleMonitoringNotificationChannelSensitiveLabelsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleMonitoringNotificationChannel.GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleMonitoringNotificationChannelSensitiveLabelsOutputReference_Override(g GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleMonitoringNotificationChannel.GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference)SetAuthToken(val *string) {
	if err := j.validateSetAuthTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authToken",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference)SetAuthTokenWo(val *string) {
	if err := j.validateSetAuthTokenWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authTokenWo",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference)SetAuthTokenWoVersion(val *string) {
	if err := j.validateSetAuthTokenWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authTokenWoVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference)SetInternalValue(val *GoogleMonitoringNotificationChannelSensitiveLabels) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference)SetPasswordWo(val *string) {
	if err := j.validateSetPasswordWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordWo",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference)SetPasswordWoVersion(val *string) {
	if err := j.validateSetPasswordWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordWoVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference)SetServiceKey(val *string) {
	if err := j.validateSetServiceKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceKey",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference)SetServiceKeyWo(val *string) {
	if err := j.validateSetServiceKeyWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceKeyWo",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference)SetServiceKeyWoVersion(val *string) {
	if err := j.validateSetServiceKeyWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceKeyWoVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) ResetAuthToken() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) ResetAuthTokenWo() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthTokenWo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) ResetAuthTokenWoVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthTokenWoVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		g,
		"resetPassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) ResetPasswordWo() {
	_jsii_.InvokeVoid(
		g,
		"resetPasswordWo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) ResetPasswordWoVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetPasswordWoVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) ResetServiceKey() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) ResetServiceKeyWo() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceKeyWo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) ResetServiceKeyWoVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceKeyWoVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleMonitoringNotificationChannelSensitiveLabelsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package dataoktadefaultsigninpage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/dataoktadefaultsigninpage/internal"
)

type DataOktaDefaultSigninPageWidgetCustomizationsOutputReference interface {
	cdktf.ComplexObject
	AuthenticatorPageCustomLinkLabel() *string
	AuthenticatorPageCustomLinkUrl() *string
	ClassicRecoveryFlowEmailOrUsernameLabel() *string
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
	CustomLink1Label() *string
	CustomLink1Url() *string
	CustomLink2Label() *string
	CustomLink2Url() *string
	ForgotPasswordLabel() *string
	ForgotPasswordUrl() *string
	// Experimental.
	Fqn() *string
	HelpLabel() *string
	HelpUrl() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PasswordInfoTip() *string
	PasswordLabel() *string
	ShowPasswordVisibilityToggle() cdktf.IResolvable
	ShowUserIdentifier() cdktf.IResolvable
	SignInLabel() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UnlockAccountLabel() *string
	UnlockAccountUrl() *string
	UsernameInfoTip() *string
	UsernameLabel() *string
	WidgetGeneration() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataOktaDefaultSigninPageWidgetCustomizationsOutputReference
type jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) AuthenticatorPageCustomLinkLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticatorPageCustomLinkLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) AuthenticatorPageCustomLinkUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticatorPageCustomLinkUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) ClassicRecoveryFlowEmailOrUsernameLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classicRecoveryFlowEmailOrUsernameLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) CustomLink1Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLink1Label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) CustomLink1Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLink1Url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) CustomLink2Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLink2Label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) CustomLink2Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLink2Url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) ForgotPasswordLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forgotPasswordLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) ForgotPasswordUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forgotPasswordUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) HelpLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"helpLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) HelpUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"helpUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) PasswordInfoTip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInfoTip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) PasswordLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) ShowPasswordVisibilityToggle() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"showPasswordVisibilityToggle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) ShowUserIdentifier() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"showUserIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) SignInLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signInLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) UnlockAccountLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unlockAccountLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) UnlockAccountUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unlockAccountUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) UsernameInfoTip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInfoTip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) UsernameLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) WidgetGeneration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widgetGeneration",
		&returns,
	)
	return returns
}


func NewDataOktaDefaultSigninPageWidgetCustomizationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataOktaDefaultSigninPageWidgetCustomizationsOutputReference {
	_init_.Initialize()

	if err := validateNewDataOktaDefaultSigninPageWidgetCustomizationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-okta.dataOktaDefaultSigninPage.DataOktaDefaultSigninPageWidgetCustomizationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataOktaDefaultSigninPageWidgetCustomizationsOutputReference_Override(d DataOktaDefaultSigninPageWidgetCustomizationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.dataOktaDefaultSigninPage.DataOktaDefaultSigninPageWidgetCustomizationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataOktaDefaultSigninPageWidgetCustomizationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


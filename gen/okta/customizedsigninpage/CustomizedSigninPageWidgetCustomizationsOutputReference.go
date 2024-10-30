package customizedsigninpage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/customizedsigninpage/internal"
)

type CustomizedSigninPageWidgetCustomizationsOutputReference interface {
	cdktf.ComplexObject
	AuthenticatorPageCustomLinkLabel() *string
	SetAuthenticatorPageCustomLinkLabel(val *string)
	AuthenticatorPageCustomLinkLabelInput() *string
	AuthenticatorPageCustomLinkUrl() *string
	SetAuthenticatorPageCustomLinkUrl(val *string)
	AuthenticatorPageCustomLinkUrlInput() *string
	ClassicRecoveryFlowEmailOrUsernameLabel() *string
	SetClassicRecoveryFlowEmailOrUsernameLabel(val *string)
	ClassicRecoveryFlowEmailOrUsernameLabelInput() *string
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
	SetCustomLink1Label(val *string)
	CustomLink1LabelInput() *string
	CustomLink1Url() *string
	SetCustomLink1Url(val *string)
	CustomLink1UrlInput() *string
	CustomLink2Label() *string
	SetCustomLink2Label(val *string)
	CustomLink2LabelInput() *string
	CustomLink2Url() *string
	SetCustomLink2Url(val *string)
	CustomLink2UrlInput() *string
	ForgotPasswordLabel() *string
	SetForgotPasswordLabel(val *string)
	ForgotPasswordLabelInput() *string
	ForgotPasswordUrl() *string
	SetForgotPasswordUrl(val *string)
	ForgotPasswordUrlInput() *string
	// Experimental.
	Fqn() *string
	HelpLabel() *string
	SetHelpLabel(val *string)
	HelpLabelInput() *string
	HelpUrl() *string
	SetHelpUrl(val *string)
	HelpUrlInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PasswordInfoTip() *string
	SetPasswordInfoTip(val *string)
	PasswordInfoTipInput() *string
	PasswordLabel() *string
	SetPasswordLabel(val *string)
	PasswordLabelInput() *string
	ShowPasswordVisibilityToggle() interface{}
	SetShowPasswordVisibilityToggle(val interface{})
	ShowPasswordVisibilityToggleInput() interface{}
	ShowUserIdentifier() interface{}
	SetShowUserIdentifier(val interface{})
	ShowUserIdentifierInput() interface{}
	SignInLabel() *string
	SetSignInLabel(val *string)
	SignInLabelInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UnlockAccountLabel() *string
	SetUnlockAccountLabel(val *string)
	UnlockAccountLabelInput() *string
	UnlockAccountUrl() *string
	SetUnlockAccountUrl(val *string)
	UnlockAccountUrlInput() *string
	UsernameInfoTip() *string
	SetUsernameInfoTip(val *string)
	UsernameInfoTipInput() *string
	UsernameLabel() *string
	SetUsernameLabel(val *string)
	UsernameLabelInput() *string
	WidgetGeneration() *string
	SetWidgetGeneration(val *string)
	WidgetGenerationInput() *string
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
	ResetAuthenticatorPageCustomLinkLabel()
	ResetAuthenticatorPageCustomLinkUrl()
	ResetClassicRecoveryFlowEmailOrUsernameLabel()
	ResetCustomLink1Label()
	ResetCustomLink1Url()
	ResetCustomLink2Label()
	ResetCustomLink2Url()
	ResetForgotPasswordLabel()
	ResetForgotPasswordUrl()
	ResetHelpLabel()
	ResetHelpUrl()
	ResetPasswordInfoTip()
	ResetPasswordLabel()
	ResetShowPasswordVisibilityToggle()
	ResetShowUserIdentifier()
	ResetSignInLabel()
	ResetUnlockAccountLabel()
	ResetUnlockAccountUrl()
	ResetUsernameInfoTip()
	ResetUsernameLabel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CustomizedSigninPageWidgetCustomizationsOutputReference
type jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) AuthenticatorPageCustomLinkLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticatorPageCustomLinkLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) AuthenticatorPageCustomLinkLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticatorPageCustomLinkLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) AuthenticatorPageCustomLinkUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticatorPageCustomLinkUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) AuthenticatorPageCustomLinkUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticatorPageCustomLinkUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ClassicRecoveryFlowEmailOrUsernameLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classicRecoveryFlowEmailOrUsernameLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ClassicRecoveryFlowEmailOrUsernameLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classicRecoveryFlowEmailOrUsernameLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) CustomLink1Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLink1Label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) CustomLink1LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLink1LabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) CustomLink1Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLink1Url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) CustomLink1UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLink1UrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) CustomLink2Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLink2Label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) CustomLink2LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLink2LabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) CustomLink2Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLink2Url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) CustomLink2UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLink2UrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ForgotPasswordLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forgotPasswordLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ForgotPasswordLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forgotPasswordLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ForgotPasswordUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forgotPasswordUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ForgotPasswordUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forgotPasswordUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) HelpLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"helpLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) HelpLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"helpLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) HelpUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"helpUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) HelpUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"helpUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) PasswordInfoTip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInfoTip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) PasswordInfoTipInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInfoTipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) PasswordLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) PasswordLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ShowPasswordVisibilityToggle() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showPasswordVisibilityToggle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ShowPasswordVisibilityToggleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showPasswordVisibilityToggleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ShowUserIdentifier() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showUserIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ShowUserIdentifierInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showUserIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) SignInLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signInLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) SignInLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signInLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) UnlockAccountLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unlockAccountLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) UnlockAccountLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unlockAccountLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) UnlockAccountUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unlockAccountUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) UnlockAccountUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unlockAccountUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) UsernameInfoTip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInfoTip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) UsernameInfoTipInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInfoTipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) UsernameLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) UsernameLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) WidgetGeneration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widgetGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) WidgetGenerationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widgetGenerationInput",
		&returns,
	)
	return returns
}


func NewCustomizedSigninPageWidgetCustomizationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CustomizedSigninPageWidgetCustomizationsOutputReference {
	_init_.Initialize()

	if err := validateNewCustomizedSigninPageWidgetCustomizationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-okta.customizedSigninPage.CustomizedSigninPageWidgetCustomizationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCustomizedSigninPageWidgetCustomizationsOutputReference_Override(c CustomizedSigninPageWidgetCustomizationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.customizedSigninPage.CustomizedSigninPageWidgetCustomizationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetAuthenticatorPageCustomLinkLabel(val *string) {
	if err := j.validateSetAuthenticatorPageCustomLinkLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticatorPageCustomLinkLabel",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetAuthenticatorPageCustomLinkUrl(val *string) {
	if err := j.validateSetAuthenticatorPageCustomLinkUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticatorPageCustomLinkUrl",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetClassicRecoveryFlowEmailOrUsernameLabel(val *string) {
	if err := j.validateSetClassicRecoveryFlowEmailOrUsernameLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"classicRecoveryFlowEmailOrUsernameLabel",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetCustomLink1Label(val *string) {
	if err := j.validateSetCustomLink1LabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customLink1Label",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetCustomLink1Url(val *string) {
	if err := j.validateSetCustomLink1UrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customLink1Url",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetCustomLink2Label(val *string) {
	if err := j.validateSetCustomLink2LabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customLink2Label",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetCustomLink2Url(val *string) {
	if err := j.validateSetCustomLink2UrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customLink2Url",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetForgotPasswordLabel(val *string) {
	if err := j.validateSetForgotPasswordLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forgotPasswordLabel",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetForgotPasswordUrl(val *string) {
	if err := j.validateSetForgotPasswordUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forgotPasswordUrl",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetHelpLabel(val *string) {
	if err := j.validateSetHelpLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"helpLabel",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetHelpUrl(val *string) {
	if err := j.validateSetHelpUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"helpUrl",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetPasswordInfoTip(val *string) {
	if err := j.validateSetPasswordInfoTipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordInfoTip",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetPasswordLabel(val *string) {
	if err := j.validateSetPasswordLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordLabel",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetShowPasswordVisibilityToggle(val interface{}) {
	if err := j.validateSetShowPasswordVisibilityToggleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showPasswordVisibilityToggle",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetShowUserIdentifier(val interface{}) {
	if err := j.validateSetShowUserIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showUserIdentifier",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetSignInLabel(val *string) {
	if err := j.validateSetSignInLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signInLabel",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetUnlockAccountLabel(val *string) {
	if err := j.validateSetUnlockAccountLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unlockAccountLabel",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetUnlockAccountUrl(val *string) {
	if err := j.validateSetUnlockAccountUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unlockAccountUrl",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetUsernameInfoTip(val *string) {
	if err := j.validateSetUsernameInfoTipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameInfoTip",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetUsernameLabel(val *string) {
	if err := j.validateSetUsernameLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameLabel",
		val,
	)
}

func (j *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference)SetWidgetGeneration(val *string) {
	if err := j.validateSetWidgetGenerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"widgetGeneration",
		val,
	)
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ResetAuthenticatorPageCustomLinkLabel() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthenticatorPageCustomLinkLabel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ResetAuthenticatorPageCustomLinkUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthenticatorPageCustomLinkUrl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ResetClassicRecoveryFlowEmailOrUsernameLabel() {
	_jsii_.InvokeVoid(
		c,
		"resetClassicRecoveryFlowEmailOrUsernameLabel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ResetCustomLink1Label() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomLink1Label",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ResetCustomLink1Url() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomLink1Url",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ResetCustomLink2Label() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomLink2Label",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ResetCustomLink2Url() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomLink2Url",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ResetForgotPasswordLabel() {
	_jsii_.InvokeVoid(
		c,
		"resetForgotPasswordLabel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ResetForgotPasswordUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetForgotPasswordUrl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ResetHelpLabel() {
	_jsii_.InvokeVoid(
		c,
		"resetHelpLabel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ResetHelpUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetHelpUrl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ResetPasswordInfoTip() {
	_jsii_.InvokeVoid(
		c,
		"resetPasswordInfoTip",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ResetPasswordLabel() {
	_jsii_.InvokeVoid(
		c,
		"resetPasswordLabel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ResetShowPasswordVisibilityToggle() {
	_jsii_.InvokeVoid(
		c,
		"resetShowPasswordVisibilityToggle",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ResetShowUserIdentifier() {
	_jsii_.InvokeVoid(
		c,
		"resetShowUserIdentifier",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ResetSignInLabel() {
	_jsii_.InvokeVoid(
		c,
		"resetSignInLabel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ResetUnlockAccountLabel() {
	_jsii_.InvokeVoid(
		c,
		"resetUnlockAccountLabel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ResetUnlockAccountUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetUnlockAccountUrl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ResetUsernameInfoTip() {
	_jsii_.InvokeVoid(
		c,
		"resetUsernameInfoTip",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ResetUsernameLabel() {
	_jsii_.InvokeVoid(
		c,
		"resetUsernameLabel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomizedSigninPageWidgetCustomizationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


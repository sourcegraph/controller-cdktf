package previewsigninpage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/previewsigninpage/internal"
)

type PreviewSigninPageWidgetCustomizationsOutputReference interface {
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

// The jsii proxy struct for PreviewSigninPageWidgetCustomizationsOutputReference
type jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) AuthenticatorPageCustomLinkLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticatorPageCustomLinkLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) AuthenticatorPageCustomLinkLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticatorPageCustomLinkLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) AuthenticatorPageCustomLinkUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticatorPageCustomLinkUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) AuthenticatorPageCustomLinkUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticatorPageCustomLinkUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ClassicRecoveryFlowEmailOrUsernameLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classicRecoveryFlowEmailOrUsernameLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ClassicRecoveryFlowEmailOrUsernameLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classicRecoveryFlowEmailOrUsernameLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) CustomLink1Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLink1Label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) CustomLink1LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLink1LabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) CustomLink1Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLink1Url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) CustomLink1UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLink1UrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) CustomLink2Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLink2Label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) CustomLink2LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLink2LabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) CustomLink2Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLink2Url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) CustomLink2UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLink2UrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ForgotPasswordLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forgotPasswordLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ForgotPasswordLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forgotPasswordLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ForgotPasswordUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forgotPasswordUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ForgotPasswordUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forgotPasswordUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) HelpLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"helpLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) HelpLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"helpLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) HelpUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"helpUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) HelpUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"helpUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) PasswordInfoTip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInfoTip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) PasswordInfoTipInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInfoTipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) PasswordLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) PasswordLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ShowPasswordVisibilityToggle() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showPasswordVisibilityToggle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ShowPasswordVisibilityToggleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showPasswordVisibilityToggleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ShowUserIdentifier() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showUserIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ShowUserIdentifierInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showUserIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) SignInLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signInLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) SignInLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signInLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) UnlockAccountLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unlockAccountLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) UnlockAccountLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unlockAccountLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) UnlockAccountUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unlockAccountUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) UnlockAccountUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unlockAccountUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) UsernameInfoTip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInfoTip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) UsernameInfoTipInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInfoTipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) UsernameLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) UsernameLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) WidgetGeneration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widgetGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) WidgetGenerationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widgetGenerationInput",
		&returns,
	)
	return returns
}


func NewPreviewSigninPageWidgetCustomizationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PreviewSigninPageWidgetCustomizationsOutputReference {
	_init_.Initialize()

	if err := validateNewPreviewSigninPageWidgetCustomizationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-okta.previewSigninPage.PreviewSigninPageWidgetCustomizationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPreviewSigninPageWidgetCustomizationsOutputReference_Override(p PreviewSigninPageWidgetCustomizationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.previewSigninPage.PreviewSigninPageWidgetCustomizationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetAuthenticatorPageCustomLinkLabel(val *string) {
	if err := j.validateSetAuthenticatorPageCustomLinkLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticatorPageCustomLinkLabel",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetAuthenticatorPageCustomLinkUrl(val *string) {
	if err := j.validateSetAuthenticatorPageCustomLinkUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticatorPageCustomLinkUrl",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetClassicRecoveryFlowEmailOrUsernameLabel(val *string) {
	if err := j.validateSetClassicRecoveryFlowEmailOrUsernameLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"classicRecoveryFlowEmailOrUsernameLabel",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetCustomLink1Label(val *string) {
	if err := j.validateSetCustomLink1LabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customLink1Label",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetCustomLink1Url(val *string) {
	if err := j.validateSetCustomLink1UrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customLink1Url",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetCustomLink2Label(val *string) {
	if err := j.validateSetCustomLink2LabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customLink2Label",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetCustomLink2Url(val *string) {
	if err := j.validateSetCustomLink2UrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customLink2Url",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetForgotPasswordLabel(val *string) {
	if err := j.validateSetForgotPasswordLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forgotPasswordLabel",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetForgotPasswordUrl(val *string) {
	if err := j.validateSetForgotPasswordUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forgotPasswordUrl",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetHelpLabel(val *string) {
	if err := j.validateSetHelpLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"helpLabel",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetHelpUrl(val *string) {
	if err := j.validateSetHelpUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"helpUrl",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetPasswordInfoTip(val *string) {
	if err := j.validateSetPasswordInfoTipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordInfoTip",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetPasswordLabel(val *string) {
	if err := j.validateSetPasswordLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordLabel",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetShowPasswordVisibilityToggle(val interface{}) {
	if err := j.validateSetShowPasswordVisibilityToggleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showPasswordVisibilityToggle",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetShowUserIdentifier(val interface{}) {
	if err := j.validateSetShowUserIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showUserIdentifier",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetSignInLabel(val *string) {
	if err := j.validateSetSignInLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signInLabel",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetUnlockAccountLabel(val *string) {
	if err := j.validateSetUnlockAccountLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unlockAccountLabel",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetUnlockAccountUrl(val *string) {
	if err := j.validateSetUnlockAccountUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unlockAccountUrl",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetUsernameInfoTip(val *string) {
	if err := j.validateSetUsernameInfoTipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameInfoTip",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetUsernameLabel(val *string) {
	if err := j.validateSetUsernameLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameLabel",
		val,
	)
}

func (j *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference)SetWidgetGeneration(val *string) {
	if err := j.validateSetWidgetGenerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"widgetGeneration",
		val,
	)
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ResetAuthenticatorPageCustomLinkLabel() {
	_jsii_.InvokeVoid(
		p,
		"resetAuthenticatorPageCustomLinkLabel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ResetAuthenticatorPageCustomLinkUrl() {
	_jsii_.InvokeVoid(
		p,
		"resetAuthenticatorPageCustomLinkUrl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ResetClassicRecoveryFlowEmailOrUsernameLabel() {
	_jsii_.InvokeVoid(
		p,
		"resetClassicRecoveryFlowEmailOrUsernameLabel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ResetCustomLink1Label() {
	_jsii_.InvokeVoid(
		p,
		"resetCustomLink1Label",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ResetCustomLink1Url() {
	_jsii_.InvokeVoid(
		p,
		"resetCustomLink1Url",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ResetCustomLink2Label() {
	_jsii_.InvokeVoid(
		p,
		"resetCustomLink2Label",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ResetCustomLink2Url() {
	_jsii_.InvokeVoid(
		p,
		"resetCustomLink2Url",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ResetForgotPasswordLabel() {
	_jsii_.InvokeVoid(
		p,
		"resetForgotPasswordLabel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ResetForgotPasswordUrl() {
	_jsii_.InvokeVoid(
		p,
		"resetForgotPasswordUrl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ResetHelpLabel() {
	_jsii_.InvokeVoid(
		p,
		"resetHelpLabel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ResetHelpUrl() {
	_jsii_.InvokeVoid(
		p,
		"resetHelpUrl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ResetPasswordInfoTip() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordInfoTip",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ResetPasswordLabel() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordLabel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ResetShowPasswordVisibilityToggle() {
	_jsii_.InvokeVoid(
		p,
		"resetShowPasswordVisibilityToggle",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ResetShowUserIdentifier() {
	_jsii_.InvokeVoid(
		p,
		"resetShowUserIdentifier",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ResetSignInLabel() {
	_jsii_.InvokeVoid(
		p,
		"resetSignInLabel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ResetUnlockAccountLabel() {
	_jsii_.InvokeVoid(
		p,
		"resetUnlockAccountLabel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ResetUnlockAccountUrl() {
	_jsii_.InvokeVoid(
		p,
		"resetUnlockAccountUrl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ResetUsernameInfoTip() {
	_jsii_.InvokeVoid(
		p,
		"resetUsernameInfoTip",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ResetUsernameLabel() {
	_jsii_.InvokeVoid(
		p,
		"resetUsernameLabel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PreviewSigninPageWidgetCustomizationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


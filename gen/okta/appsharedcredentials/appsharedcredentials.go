package appsharedcredentials

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/appsharedcredentials/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_shared_credentials okta_app_shared_credentials}.
type AppSharedCredentials interface {
	cdktf.TerraformResource
	AccessibilityErrorRedirectUrl() *string
	SetAccessibilityErrorRedirectUrl(val *string)
	AccessibilityErrorRedirectUrlInput() *string
	AccessibilityLoginRedirectUrl() *string
	SetAccessibilityLoginRedirectUrl(val *string)
	AccessibilityLoginRedirectUrlInput() *string
	AccessibilitySelfService() interface{}
	SetAccessibilitySelfService(val interface{})
	AccessibilitySelfServiceInput() interface{}
	AdminNote() *string
	SetAdminNote(val *string)
	AdminNoteInput() *string
	AppLinksJson() *string
	SetAppLinksJson(val *string)
	AppLinksJsonInput() *string
	AutoSubmitToolbar() interface{}
	SetAutoSubmitToolbar(val interface{})
	AutoSubmitToolbarInput() interface{}
	ButtonField() *string
	SetButtonField(val *string)
	ButtonFieldInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Checkbox() *string
	SetCheckbox(val *string)
	CheckboxInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnduserNote() *string
	SetEnduserNote(val *string)
	EnduserNoteInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HideIos() interface{}
	SetHideIos(val interface{})
	HideIosInput() interface{}
	HideWeb() interface{}
	SetHideWeb(val interface{})
	HideWebInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Logo() *string
	SetLogo(val *string)
	LogoInput() *string
	LogoUrl() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	PasswordField() *string
	SetPasswordField(val *string)
	PasswordFieldInput() *string
	PreconfiguredApp() *string
	SetPreconfiguredApp(val *string)
	PreconfiguredAppInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RedirectUrl() *string
	SetRedirectUrl(val *string)
	RedirectUrlInput() *string
	SharedPassword() *string
	SetSharedPassword(val *string)
	SharedPasswordInput() *string
	SharedUsername() *string
	SetSharedUsername(val *string)
	SharedUsernameInput() *string
	SignOnMode() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() AppSharedCredentialsTimeoutsOutputReference
	TimeoutsInput() interface{}
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
	UrlRegex() *string
	SetUrlRegex(val *string)
	UrlRegexInput() *string
	UsernameField() *string
	SetUsernameField(val *string)
	UsernameFieldInput() *string
	UserNameTemplate() *string
	SetUserNameTemplate(val *string)
	UserNameTemplateInput() *string
	UserNameTemplatePushStatus() *string
	SetUserNameTemplatePushStatus(val *string)
	UserNameTemplatePushStatusInput() *string
	UserNameTemplateSuffix() *string
	SetUserNameTemplateSuffix(val *string)
	UserNameTemplateSuffixInput() *string
	UserNameTemplateType() *string
	SetUserNameTemplateType(val *string)
	UserNameTemplateTypeInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *AppSharedCredentialsTimeouts)
	ResetAccessibilityErrorRedirectUrl()
	ResetAccessibilityLoginRedirectUrl()
	ResetAccessibilitySelfService()
	ResetAdminNote()
	ResetAppLinksJson()
	ResetAutoSubmitToolbar()
	ResetButtonField()
	ResetCheckbox()
	ResetEnduserNote()
	ResetHideIos()
	ResetHideWeb()
	ResetId()
	ResetLogo()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPasswordField()
	ResetPreconfiguredApp()
	ResetRedirectUrl()
	ResetSharedPassword()
	ResetSharedUsername()
	ResetStatus()
	ResetTimeouts()
	ResetUrl()
	ResetUrlRegex()
	ResetUsernameField()
	ResetUserNameTemplate()
	ResetUserNameTemplatePushStatus()
	ResetUserNameTemplateSuffix()
	ResetUserNameTemplateType()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AppSharedCredentials
type jsiiProxy_AppSharedCredentials struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppSharedCredentials) AccessibilityErrorRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityErrorRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) AccessibilityErrorRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityErrorRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) AccessibilityLoginRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityLoginRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) AccessibilityLoginRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityLoginRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) AccessibilitySelfService() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessibilitySelfService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) AccessibilitySelfServiceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessibilitySelfServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) AdminNote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminNote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) AdminNoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminNoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) AppLinksJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLinksJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) AppLinksJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLinksJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) AutoSubmitToolbar() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSubmitToolbar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) AutoSubmitToolbarInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSubmitToolbarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) ButtonField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buttonField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) ButtonFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buttonFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) Checkbox() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkbox",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) CheckboxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkboxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) EnduserNote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enduserNote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) EnduserNoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enduserNoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) HideIos() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) HideIosInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) HideWeb() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideWeb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) HideWebInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideWebInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) Logo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) LogoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) LogoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) PasswordField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) PasswordFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) PreconfiguredApp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconfiguredApp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) PreconfiguredAppInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconfiguredAppInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) RedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) RedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) SharedPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) SharedPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) SharedUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) SharedUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) SignOnMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signOnMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) Timeouts() AppSharedCredentialsTimeoutsOutputReference {
	var returns AppSharedCredentialsTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) UrlRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) UrlRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) UsernameField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) UsernameFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) UserNameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) UserNameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) UserNameTemplatePushStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplatePushStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) UserNameTemplatePushStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplatePushStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) UserNameTemplateSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) UserNameTemplateSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) UserNameTemplateType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSharedCredentials) UserNameTemplateTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_shared_credentials okta_app_shared_credentials} Resource.
func NewAppSharedCredentials(scope constructs.Construct, id *string, config *AppSharedCredentialsConfig) AppSharedCredentials {
	_init_.Initialize()

	if err := validateNewAppSharedCredentialsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppSharedCredentials{}

	_jsii_.Create(
		"@cdktf/provider-okta.appSharedCredentials.AppSharedCredentials",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_shared_credentials okta_app_shared_credentials} Resource.
func NewAppSharedCredentials_Override(a AppSharedCredentials, scope constructs.Construct, id *string, config *AppSharedCredentialsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.appSharedCredentials.AppSharedCredentials",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetAccessibilityErrorRedirectUrl(val *string) {
	if err := j.validateSetAccessibilityErrorRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilityErrorRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetAccessibilityLoginRedirectUrl(val *string) {
	if err := j.validateSetAccessibilityLoginRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilityLoginRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetAccessibilitySelfService(val interface{}) {
	if err := j.validateSetAccessibilitySelfServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilitySelfService",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetAdminNote(val *string) {
	if err := j.validateSetAdminNoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminNote",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetAppLinksJson(val *string) {
	if err := j.validateSetAppLinksJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appLinksJson",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetAutoSubmitToolbar(val interface{}) {
	if err := j.validateSetAutoSubmitToolbarParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoSubmitToolbar",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetButtonField(val *string) {
	if err := j.validateSetButtonFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buttonField",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetCheckbox(val *string) {
	if err := j.validateSetCheckboxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkbox",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetEnduserNote(val *string) {
	if err := j.validateSetEnduserNoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enduserNote",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetHideIos(val interface{}) {
	if err := j.validateSetHideIosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideIos",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetHideWeb(val interface{}) {
	if err := j.validateSetHideWebParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideWeb",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetLogo(val *string) {
	if err := j.validateSetLogoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logo",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetPasswordField(val *string) {
	if err := j.validateSetPasswordFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordField",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetPreconfiguredApp(val *string) {
	if err := j.validateSetPreconfiguredAppParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preconfiguredApp",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetRedirectUrl(val *string) {
	if err := j.validateSetRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectUrl",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetSharedPassword(val *string) {
	if err := j.validateSetSharedPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedPassword",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetSharedUsername(val *string) {
	if err := j.validateSetSharedUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedUsername",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetUrlRegex(val *string) {
	if err := j.validateSetUrlRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlRegex",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetUsernameField(val *string) {
	if err := j.validateSetUsernameFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameField",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetUserNameTemplate(val *string) {
	if err := j.validateSetUserNameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplate",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetUserNameTemplatePushStatus(val *string) {
	if err := j.validateSetUserNameTemplatePushStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplatePushStatus",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetUserNameTemplateSuffix(val *string) {
	if err := j.validateSetUserNameTemplateSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplateSuffix",
		val,
	)
}

func (j *jsiiProxy_AppSharedCredentials)SetUserNameTemplateType(val *string) {
	if err := j.validateSetUserNameTemplateTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplateType",
		val,
	)
}

// Generates CDKTF code for importing a AppSharedCredentials resource upon running "cdktf plan <stack-name>".
func AppSharedCredentials_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAppSharedCredentials_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.appSharedCredentials.AppSharedCredentials",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func AppSharedCredentials_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppSharedCredentials_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.appSharedCredentials.AppSharedCredentials",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppSharedCredentials_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppSharedCredentials_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.appSharedCredentials.AppSharedCredentials",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppSharedCredentials_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppSharedCredentials_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.appSharedCredentials.AppSharedCredentials",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppSharedCredentials_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.appSharedCredentials.AppSharedCredentials",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppSharedCredentials) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AppSharedCredentials) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppSharedCredentials) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSharedCredentials) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSharedCredentials) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSharedCredentials) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSharedCredentials) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSharedCredentials) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSharedCredentials) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSharedCredentials) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSharedCredentials) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSharedCredentials) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSharedCredentials) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AppSharedCredentials) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSharedCredentials) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppSharedCredentials) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AppSharedCredentials) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppSharedCredentials) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppSharedCredentials) PutTimeouts(value *AppSharedCredentialsTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetAccessibilityErrorRedirectUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessibilityErrorRedirectUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetAccessibilityLoginRedirectUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessibilityLoginRedirectUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetAccessibilitySelfService() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessibilitySelfService",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetAdminNote() {
	_jsii_.InvokeVoid(
		a,
		"resetAdminNote",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetAppLinksJson() {
	_jsii_.InvokeVoid(
		a,
		"resetAppLinksJson",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetAutoSubmitToolbar() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoSubmitToolbar",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetButtonField() {
	_jsii_.InvokeVoid(
		a,
		"resetButtonField",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetCheckbox() {
	_jsii_.InvokeVoid(
		a,
		"resetCheckbox",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetEnduserNote() {
	_jsii_.InvokeVoid(
		a,
		"resetEnduserNote",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetHideIos() {
	_jsii_.InvokeVoid(
		a,
		"resetHideIos",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetHideWeb() {
	_jsii_.InvokeVoid(
		a,
		"resetHideWeb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetLogo() {
	_jsii_.InvokeVoid(
		a,
		"resetLogo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetPasswordField() {
	_jsii_.InvokeVoid(
		a,
		"resetPasswordField",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetPreconfiguredApp() {
	_jsii_.InvokeVoid(
		a,
		"resetPreconfiguredApp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetRedirectUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetRedirectUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetSharedPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetSharedPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetSharedUsername() {
	_jsii_.InvokeVoid(
		a,
		"resetSharedUsername",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetUrlRegex() {
	_jsii_.InvokeVoid(
		a,
		"resetUrlRegex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetUsernameField() {
	_jsii_.InvokeVoid(
		a,
		"resetUsernameField",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetUserNameTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetUserNameTemplatePushStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplatePushStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetUserNameTemplateSuffix() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplateSuffix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) ResetUserNameTemplateType() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplateType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSharedCredentials) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSharedCredentials) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSharedCredentials) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSharedCredentials) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSharedCredentials) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSharedCredentials) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


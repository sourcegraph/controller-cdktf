package securepasswordstoreapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/securepasswordstoreapp/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/secure_password_store_app okta_secure_password_store_app}.
type SecurePasswordStoreApp interface {
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
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CredentialsScheme() *string
	SetCredentialsScheme(val *string)
	CredentialsSchemeInput() *string
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
	Groups() *[]*string
	SetGroups(val *[]*string)
	GroupsInput() *[]*string
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
	OptionalField1() *string
	SetOptionalField1(val *string)
	OptionalField1Input() *string
	OptionalField1Value() *string
	SetOptionalField1Value(val *string)
	OptionalField1ValueInput() *string
	OptionalField2() *string
	SetOptionalField2(val *string)
	OptionalField2Input() *string
	OptionalField2Value() *string
	SetOptionalField2Value(val *string)
	OptionalField2ValueInput() *string
	OptionalField3() *string
	SetOptionalField3(val *string)
	OptionalField3Input() *string
	OptionalField3Value() *string
	SetOptionalField3Value(val *string)
	OptionalField3ValueInput() *string
	PasswordField() *string
	SetPasswordField(val *string)
	PasswordFieldInput() *string
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
	RevealPassword() interface{}
	SetRevealPassword(val interface{})
	RevealPasswordInput() interface{}
	SharedPassword() *string
	SetSharedPassword(val *string)
	SharedPasswordInput() *string
	SharedUsername() *string
	SetSharedUsername(val *string)
	SharedUsernameInput() *string
	SignOnMode() *string
	SkipGroups() interface{}
	SetSkipGroups(val interface{})
	SkipGroupsInput() interface{}
	SkipUsers() interface{}
	SetSkipUsers(val interface{})
	SkipUsersInput() interface{}
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SecurePasswordStoreAppTimeoutsOutputReference
	TimeoutsInput() interface{}
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
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
	Users() SecurePasswordStoreAppUsersList
	UsersInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *SecurePasswordStoreAppTimeouts)
	PutUsers(value interface{})
	ResetAccessibilityErrorRedirectUrl()
	ResetAccessibilityLoginRedirectUrl()
	ResetAccessibilitySelfService()
	ResetAdminNote()
	ResetAppLinksJson()
	ResetAutoSubmitToolbar()
	ResetCredentialsScheme()
	ResetEnduserNote()
	ResetGroups()
	ResetHideIos()
	ResetHideWeb()
	ResetId()
	ResetLogo()
	ResetOptionalField1()
	ResetOptionalField1Value()
	ResetOptionalField2()
	ResetOptionalField2Value()
	ResetOptionalField3()
	ResetOptionalField3Value()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRevealPassword()
	ResetSharedPassword()
	ResetSharedUsername()
	ResetSkipGroups()
	ResetSkipUsers()
	ResetStatus()
	ResetTimeouts()
	ResetUserNameTemplate()
	ResetUserNameTemplatePushStatus()
	ResetUserNameTemplateSuffix()
	ResetUserNameTemplateType()
	ResetUsers()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SecurePasswordStoreApp
type jsiiProxy_SecurePasswordStoreApp struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SecurePasswordStoreApp) AccessibilityErrorRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityErrorRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) AccessibilityErrorRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityErrorRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) AccessibilityLoginRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityLoginRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) AccessibilityLoginRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityLoginRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) AccessibilitySelfService() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessibilitySelfService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) AccessibilitySelfServiceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessibilitySelfServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) AdminNote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminNote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) AdminNoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminNoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) AppLinksJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLinksJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) AppLinksJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLinksJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) AutoSubmitToolbar() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSubmitToolbar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) AutoSubmitToolbarInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSubmitToolbarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) CredentialsScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) CredentialsSchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) EnduserNote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enduserNote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) EnduserNoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enduserNoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) Groups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) GroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) HideIos() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) HideIosInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) HideWeb() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideWeb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) HideWebInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideWebInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) Logo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) LogoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) LogoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) OptionalField1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) OptionalField1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) OptionalField1Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField1Value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) OptionalField1ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField1ValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) OptionalField2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) OptionalField2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) OptionalField2Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField2Value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) OptionalField2ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField2ValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) OptionalField3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) OptionalField3Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) OptionalField3Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField3Value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) OptionalField3ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField3ValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) PasswordField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) PasswordFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) RevealPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revealPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) RevealPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revealPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) SharedPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) SharedPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) SharedUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) SharedUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) SignOnMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signOnMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) SkipGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) SkipGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) SkipUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) SkipUsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) Timeouts() SecurePasswordStoreAppTimeoutsOutputReference {
	var returns SecurePasswordStoreAppTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) UsernameField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) UsernameFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) UserNameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) UserNameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) UserNameTemplatePushStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplatePushStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) UserNameTemplatePushStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplatePushStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) UserNameTemplateSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) UserNameTemplateSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) UserNameTemplateType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) UserNameTemplateTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) Users() SecurePasswordStoreAppUsersList {
	var returns SecurePasswordStoreAppUsersList
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurePasswordStoreApp) UsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usersInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/secure_password_store_app okta_secure_password_store_app} Resource.
func NewSecurePasswordStoreApp(scope constructs.Construct, id *string, config *SecurePasswordStoreAppConfig) SecurePasswordStoreApp {
	_init_.Initialize()

	if err := validateNewSecurePasswordStoreAppParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurePasswordStoreApp{}

	_jsii_.Create(
		"@cdktf/provider-okta.securePasswordStoreApp.SecurePasswordStoreApp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/secure_password_store_app okta_secure_password_store_app} Resource.
func NewSecurePasswordStoreApp_Override(s SecurePasswordStoreApp, scope constructs.Construct, id *string, config *SecurePasswordStoreAppConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.securePasswordStoreApp.SecurePasswordStoreApp",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetAccessibilityErrorRedirectUrl(val *string) {
	if err := j.validateSetAccessibilityErrorRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilityErrorRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetAccessibilityLoginRedirectUrl(val *string) {
	if err := j.validateSetAccessibilityLoginRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilityLoginRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetAccessibilitySelfService(val interface{}) {
	if err := j.validateSetAccessibilitySelfServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilitySelfService",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetAdminNote(val *string) {
	if err := j.validateSetAdminNoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminNote",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetAppLinksJson(val *string) {
	if err := j.validateSetAppLinksJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appLinksJson",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetAutoSubmitToolbar(val interface{}) {
	if err := j.validateSetAutoSubmitToolbarParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoSubmitToolbar",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetCredentialsScheme(val *string) {
	if err := j.validateSetCredentialsSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialsScheme",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetEnduserNote(val *string) {
	if err := j.validateSetEnduserNoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enduserNote",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetGroups(val *[]*string) {
	if err := j.validateSetGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groups",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetHideIos(val interface{}) {
	if err := j.validateSetHideIosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideIos",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetHideWeb(val interface{}) {
	if err := j.validateSetHideWebParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideWeb",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetLogo(val *string) {
	if err := j.validateSetLogoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logo",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetOptionalField1(val *string) {
	if err := j.validateSetOptionalField1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionalField1",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetOptionalField1Value(val *string) {
	if err := j.validateSetOptionalField1ValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionalField1Value",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetOptionalField2(val *string) {
	if err := j.validateSetOptionalField2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionalField2",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetOptionalField2Value(val *string) {
	if err := j.validateSetOptionalField2ValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionalField2Value",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetOptionalField3(val *string) {
	if err := j.validateSetOptionalField3Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionalField3",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetOptionalField3Value(val *string) {
	if err := j.validateSetOptionalField3ValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionalField3Value",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetPasswordField(val *string) {
	if err := j.validateSetPasswordFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordField",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetRevealPassword(val interface{}) {
	if err := j.validateSetRevealPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revealPassword",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetSharedPassword(val *string) {
	if err := j.validateSetSharedPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedPassword",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetSharedUsername(val *string) {
	if err := j.validateSetSharedUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedUsername",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetSkipGroups(val interface{}) {
	if err := j.validateSetSkipGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipGroups",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetSkipUsers(val interface{}) {
	if err := j.validateSetSkipUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipUsers",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetUsernameField(val *string) {
	if err := j.validateSetUsernameFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameField",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetUserNameTemplate(val *string) {
	if err := j.validateSetUserNameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplate",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetUserNameTemplatePushStatus(val *string) {
	if err := j.validateSetUserNameTemplatePushStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplatePushStatus",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetUserNameTemplateSuffix(val *string) {
	if err := j.validateSetUserNameTemplateSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplateSuffix",
		val,
	)
}

func (j *jsiiProxy_SecurePasswordStoreApp)SetUserNameTemplateType(val *string) {
	if err := j.validateSetUserNameTemplateTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplateType",
		val,
	)
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
func SecurePasswordStoreApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecurePasswordStoreApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.securePasswordStoreApp.SecurePasswordStoreApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecurePasswordStoreApp_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecurePasswordStoreApp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.securePasswordStoreApp.SecurePasswordStoreApp",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecurePasswordStoreApp_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecurePasswordStoreApp_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.securePasswordStoreApp.SecurePasswordStoreApp",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecurePasswordStoreApp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.securePasswordStoreApp.SecurePasswordStoreApp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SecurePasswordStoreApp) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurePasswordStoreApp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurePasswordStoreApp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurePasswordStoreApp) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurePasswordStoreApp) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurePasswordStoreApp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurePasswordStoreApp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurePasswordStoreApp) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurePasswordStoreApp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurePasswordStoreApp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurePasswordStoreApp) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) PutTimeouts(value *SecurePasswordStoreAppTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) PutUsers(value interface{}) {
	if err := s.validatePutUsersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putUsers",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetAccessibilityErrorRedirectUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetAccessibilityErrorRedirectUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetAccessibilityLoginRedirectUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetAccessibilityLoginRedirectUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetAccessibilitySelfService() {
	_jsii_.InvokeVoid(
		s,
		"resetAccessibilitySelfService",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetAdminNote() {
	_jsii_.InvokeVoid(
		s,
		"resetAdminNote",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetAppLinksJson() {
	_jsii_.InvokeVoid(
		s,
		"resetAppLinksJson",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetAutoSubmitToolbar() {
	_jsii_.InvokeVoid(
		s,
		"resetAutoSubmitToolbar",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetCredentialsScheme() {
	_jsii_.InvokeVoid(
		s,
		"resetCredentialsScheme",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetEnduserNote() {
	_jsii_.InvokeVoid(
		s,
		"resetEnduserNote",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetGroups() {
	_jsii_.InvokeVoid(
		s,
		"resetGroups",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetHideIos() {
	_jsii_.InvokeVoid(
		s,
		"resetHideIos",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetHideWeb() {
	_jsii_.InvokeVoid(
		s,
		"resetHideWeb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetLogo() {
	_jsii_.InvokeVoid(
		s,
		"resetLogo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetOptionalField1() {
	_jsii_.InvokeVoid(
		s,
		"resetOptionalField1",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetOptionalField1Value() {
	_jsii_.InvokeVoid(
		s,
		"resetOptionalField1Value",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetOptionalField2() {
	_jsii_.InvokeVoid(
		s,
		"resetOptionalField2",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetOptionalField2Value() {
	_jsii_.InvokeVoid(
		s,
		"resetOptionalField2Value",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetOptionalField3() {
	_jsii_.InvokeVoid(
		s,
		"resetOptionalField3",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetOptionalField3Value() {
	_jsii_.InvokeVoid(
		s,
		"resetOptionalField3Value",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetRevealPassword() {
	_jsii_.InvokeVoid(
		s,
		"resetRevealPassword",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetSharedPassword() {
	_jsii_.InvokeVoid(
		s,
		"resetSharedPassword",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetSharedUsername() {
	_jsii_.InvokeVoid(
		s,
		"resetSharedUsername",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetSkipGroups() {
	_jsii_.InvokeVoid(
		s,
		"resetSkipGroups",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetSkipUsers() {
	_jsii_.InvokeVoid(
		s,
		"resetSkipUsers",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetUserNameTemplate() {
	_jsii_.InvokeVoid(
		s,
		"resetUserNameTemplate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetUserNameTemplatePushStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetUserNameTemplatePushStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetUserNameTemplateSuffix() {
	_jsii_.InvokeVoid(
		s,
		"resetUserNameTemplateSuffix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetUserNameTemplateType() {
	_jsii_.InvokeVoid(
		s,
		"resetUserNameTemplateType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) ResetUsers() {
	_jsii_.InvokeVoid(
		s,
		"resetUsers",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurePasswordStoreApp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurePasswordStoreApp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurePasswordStoreApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurePasswordStoreApp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


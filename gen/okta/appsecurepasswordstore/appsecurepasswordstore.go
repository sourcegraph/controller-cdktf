package appsecurepasswordstore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/appsecurepasswordstore/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/app_secure_password_store okta_app_secure_password_store}.
type AppSecurePasswordStore interface {
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
	Timeouts() AppSecurePasswordStoreTimeoutsOutputReference
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
	Users() AppSecurePasswordStoreUsersList
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
	PutTimeouts(value *AppSecurePasswordStoreTimeouts)
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

// The jsii proxy struct for AppSecurePasswordStore
type jsiiProxy_AppSecurePasswordStore struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppSecurePasswordStore) AccessibilityErrorRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityErrorRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) AccessibilityErrorRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityErrorRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) AccessibilityLoginRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityLoginRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) AccessibilityLoginRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityLoginRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) AccessibilitySelfService() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessibilitySelfService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) AccessibilitySelfServiceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessibilitySelfServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) AdminNote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminNote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) AdminNoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminNoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) AppLinksJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLinksJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) AppLinksJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLinksJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) AutoSubmitToolbar() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSubmitToolbar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) AutoSubmitToolbarInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSubmitToolbarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) CredentialsScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) CredentialsSchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) EnduserNote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enduserNote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) EnduserNoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enduserNoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) Groups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) GroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) HideIos() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) HideIosInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) HideWeb() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideWeb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) HideWebInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideWebInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) Logo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) LogoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) LogoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) OptionalField1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) OptionalField1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) OptionalField1Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField1Value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) OptionalField1ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField1ValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) OptionalField2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) OptionalField2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) OptionalField2Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField2Value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) OptionalField2ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField2ValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) OptionalField3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) OptionalField3Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) OptionalField3Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField3Value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) OptionalField3ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionalField3ValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) PasswordField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) PasswordFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) RevealPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revealPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) RevealPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revealPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) SharedPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) SharedPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) SharedUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) SharedUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) SignOnMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signOnMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) SkipGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) SkipGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) SkipUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) SkipUsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) Timeouts() AppSecurePasswordStoreTimeoutsOutputReference {
	var returns AppSecurePasswordStoreTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) UsernameField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) UsernameFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) UserNameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) UserNameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) UserNameTemplatePushStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplatePushStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) UserNameTemplatePushStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplatePushStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) UserNameTemplateSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) UserNameTemplateSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) UserNameTemplateType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) UserNameTemplateTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) Users() AppSecurePasswordStoreUsersList {
	var returns AppSecurePasswordStoreUsersList
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSecurePasswordStore) UsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usersInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/app_secure_password_store okta_app_secure_password_store} Resource.
func NewAppSecurePasswordStore(scope constructs.Construct, id *string, config *AppSecurePasswordStoreConfig) AppSecurePasswordStore {
	_init_.Initialize()

	if err := validateNewAppSecurePasswordStoreParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppSecurePasswordStore{}

	_jsii_.Create(
		"@cdktf/provider-okta.appSecurePasswordStore.AppSecurePasswordStore",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/app_secure_password_store okta_app_secure_password_store} Resource.
func NewAppSecurePasswordStore_Override(a AppSecurePasswordStore, scope constructs.Construct, id *string, config *AppSecurePasswordStoreConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.appSecurePasswordStore.AppSecurePasswordStore",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetAccessibilityErrorRedirectUrl(val *string) {
	if err := j.validateSetAccessibilityErrorRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilityErrorRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetAccessibilityLoginRedirectUrl(val *string) {
	if err := j.validateSetAccessibilityLoginRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilityLoginRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetAccessibilitySelfService(val interface{}) {
	if err := j.validateSetAccessibilitySelfServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilitySelfService",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetAdminNote(val *string) {
	if err := j.validateSetAdminNoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminNote",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetAppLinksJson(val *string) {
	if err := j.validateSetAppLinksJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appLinksJson",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetAutoSubmitToolbar(val interface{}) {
	if err := j.validateSetAutoSubmitToolbarParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoSubmitToolbar",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetCredentialsScheme(val *string) {
	if err := j.validateSetCredentialsSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialsScheme",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetEnduserNote(val *string) {
	if err := j.validateSetEnduserNoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enduserNote",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetGroups(val *[]*string) {
	if err := j.validateSetGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groups",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetHideIos(val interface{}) {
	if err := j.validateSetHideIosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideIos",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetHideWeb(val interface{}) {
	if err := j.validateSetHideWebParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideWeb",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetLogo(val *string) {
	if err := j.validateSetLogoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logo",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetOptionalField1(val *string) {
	if err := j.validateSetOptionalField1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionalField1",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetOptionalField1Value(val *string) {
	if err := j.validateSetOptionalField1ValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionalField1Value",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetOptionalField2(val *string) {
	if err := j.validateSetOptionalField2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionalField2",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetOptionalField2Value(val *string) {
	if err := j.validateSetOptionalField2ValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionalField2Value",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetOptionalField3(val *string) {
	if err := j.validateSetOptionalField3Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionalField3",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetOptionalField3Value(val *string) {
	if err := j.validateSetOptionalField3ValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionalField3Value",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetPasswordField(val *string) {
	if err := j.validateSetPasswordFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordField",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetRevealPassword(val interface{}) {
	if err := j.validateSetRevealPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revealPassword",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetSharedPassword(val *string) {
	if err := j.validateSetSharedPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedPassword",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetSharedUsername(val *string) {
	if err := j.validateSetSharedUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedUsername",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetSkipGroups(val interface{}) {
	if err := j.validateSetSkipGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipGroups",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetSkipUsers(val interface{}) {
	if err := j.validateSetSkipUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipUsers",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetUsernameField(val *string) {
	if err := j.validateSetUsernameFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameField",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetUserNameTemplate(val *string) {
	if err := j.validateSetUserNameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplate",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetUserNameTemplatePushStatus(val *string) {
	if err := j.validateSetUserNameTemplatePushStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplatePushStatus",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetUserNameTemplateSuffix(val *string) {
	if err := j.validateSetUserNameTemplateSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplateSuffix",
		val,
	)
}

func (j *jsiiProxy_AppSecurePasswordStore)SetUserNameTemplateType(val *string) {
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
func AppSecurePasswordStore_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppSecurePasswordStore_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.appSecurePasswordStore.AppSecurePasswordStore",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppSecurePasswordStore_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppSecurePasswordStore_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.appSecurePasswordStore.AppSecurePasswordStore",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppSecurePasswordStore_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppSecurePasswordStore_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.appSecurePasswordStore.AppSecurePasswordStore",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppSecurePasswordStore_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.appSecurePasswordStore.AppSecurePasswordStore",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppSecurePasswordStore) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppSecurePasswordStore) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppSecurePasswordStore) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppSecurePasswordStore) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppSecurePasswordStore) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppSecurePasswordStore) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppSecurePasswordStore) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppSecurePasswordStore) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppSecurePasswordStore) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppSecurePasswordStore) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppSecurePasswordStore) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) PutTimeouts(value *AppSecurePasswordStoreTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) PutUsers(value interface{}) {
	if err := a.validatePutUsersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUsers",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetAccessibilityErrorRedirectUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessibilityErrorRedirectUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetAccessibilityLoginRedirectUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessibilityLoginRedirectUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetAccessibilitySelfService() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessibilitySelfService",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetAdminNote() {
	_jsii_.InvokeVoid(
		a,
		"resetAdminNote",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetAppLinksJson() {
	_jsii_.InvokeVoid(
		a,
		"resetAppLinksJson",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetAutoSubmitToolbar() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoSubmitToolbar",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetCredentialsScheme() {
	_jsii_.InvokeVoid(
		a,
		"resetCredentialsScheme",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetEnduserNote() {
	_jsii_.InvokeVoid(
		a,
		"resetEnduserNote",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetHideIos() {
	_jsii_.InvokeVoid(
		a,
		"resetHideIos",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetHideWeb() {
	_jsii_.InvokeVoid(
		a,
		"resetHideWeb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetLogo() {
	_jsii_.InvokeVoid(
		a,
		"resetLogo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetOptionalField1() {
	_jsii_.InvokeVoid(
		a,
		"resetOptionalField1",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetOptionalField1Value() {
	_jsii_.InvokeVoid(
		a,
		"resetOptionalField1Value",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetOptionalField2() {
	_jsii_.InvokeVoid(
		a,
		"resetOptionalField2",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetOptionalField2Value() {
	_jsii_.InvokeVoid(
		a,
		"resetOptionalField2Value",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetOptionalField3() {
	_jsii_.InvokeVoid(
		a,
		"resetOptionalField3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetOptionalField3Value() {
	_jsii_.InvokeVoid(
		a,
		"resetOptionalField3Value",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetRevealPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetRevealPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetSharedPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetSharedPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetSharedUsername() {
	_jsii_.InvokeVoid(
		a,
		"resetSharedUsername",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetSkipGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetSkipUsers() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipUsers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetUserNameTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetUserNameTemplatePushStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplatePushStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetUserNameTemplateSuffix() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplateSuffix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetUserNameTemplateType() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplateType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) ResetUsers() {
	_jsii_.InvokeVoid(
		a,
		"resetUsers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSecurePasswordStore) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSecurePasswordStore) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSecurePasswordStore) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSecurePasswordStore) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


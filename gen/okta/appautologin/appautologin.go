package appautologin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/appautologin/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/app_auto_login okta_app_auto_login}.
type AppAutoLogin interface {
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
	AppSettingsJson() *string
	SetAppSettingsJson(val *string)
	AppSettingsJsonInput() *string
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
	SignOnRedirectUrl() *string
	SetSignOnRedirectUrl(val *string)
	SignOnRedirectUrlInput() *string
	SignOnUrl() *string
	SetSignOnUrl(val *string)
	SignOnUrlInput() *string
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
	Timeouts() AppAutoLoginTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	Users() AppAutoLoginUsersList
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
	PutTimeouts(value *AppAutoLoginTimeouts)
	PutUsers(value interface{})
	ResetAccessibilityErrorRedirectUrl()
	ResetAccessibilityLoginRedirectUrl()
	ResetAccessibilitySelfService()
	ResetAdminNote()
	ResetAppLinksJson()
	ResetAppSettingsJson()
	ResetAutoSubmitToolbar()
	ResetCredentialsScheme()
	ResetEnduserNote()
	ResetGroups()
	ResetHideIos()
	ResetHideWeb()
	ResetId()
	ResetLogo()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreconfiguredApp()
	ResetRevealPassword()
	ResetSharedPassword()
	ResetSharedUsername()
	ResetSignOnRedirectUrl()
	ResetSignOnUrl()
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

// The jsii proxy struct for AppAutoLogin
type jsiiProxy_AppAutoLogin struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppAutoLogin) AccessibilityErrorRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityErrorRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) AccessibilityErrorRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityErrorRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) AccessibilityLoginRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityLoginRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) AccessibilityLoginRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityLoginRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) AccessibilitySelfService() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessibilitySelfService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) AccessibilitySelfServiceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessibilitySelfServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) AdminNote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminNote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) AdminNoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminNoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) AppLinksJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLinksJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) AppLinksJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLinksJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) AppSettingsJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSettingsJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) AppSettingsJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSettingsJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) AutoSubmitToolbar() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSubmitToolbar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) AutoSubmitToolbarInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSubmitToolbarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) CredentialsScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) CredentialsSchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) EnduserNote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enduserNote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) EnduserNoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enduserNoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) Groups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) GroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) HideIos() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) HideIosInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) HideWeb() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideWeb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) HideWebInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideWebInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) Logo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) LogoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) LogoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) PreconfiguredApp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconfiguredApp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) PreconfiguredAppInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconfiguredAppInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) RevealPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revealPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) RevealPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revealPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) SharedPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) SharedPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) SharedUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) SharedUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) SignOnMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signOnMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) SignOnRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signOnRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) SignOnRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signOnRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) SignOnUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signOnUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) SignOnUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signOnUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) SkipGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) SkipGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) SkipUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) SkipUsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) Timeouts() AppAutoLoginTimeoutsOutputReference {
	var returns AppAutoLoginTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) UserNameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) UserNameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) UserNameTemplatePushStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplatePushStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) UserNameTemplatePushStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplatePushStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) UserNameTemplateSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) UserNameTemplateSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) UserNameTemplateType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) UserNameTemplateTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) Users() AppAutoLoginUsersList {
	var returns AppAutoLoginUsersList
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppAutoLogin) UsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usersInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/app_auto_login okta_app_auto_login} Resource.
func NewAppAutoLogin(scope constructs.Construct, id *string, config *AppAutoLoginConfig) AppAutoLogin {
	_init_.Initialize()

	if err := validateNewAppAutoLoginParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppAutoLogin{}

	_jsii_.Create(
		"okta.appAutoLogin.AppAutoLogin",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/app_auto_login okta_app_auto_login} Resource.
func NewAppAutoLogin_Override(a AppAutoLogin, scope constructs.Construct, id *string, config *AppAutoLoginConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"okta.appAutoLogin.AppAutoLogin",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetAccessibilityErrorRedirectUrl(val *string) {
	if err := j.validateSetAccessibilityErrorRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilityErrorRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetAccessibilityLoginRedirectUrl(val *string) {
	if err := j.validateSetAccessibilityLoginRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilityLoginRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetAccessibilitySelfService(val interface{}) {
	if err := j.validateSetAccessibilitySelfServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilitySelfService",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetAdminNote(val *string) {
	if err := j.validateSetAdminNoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminNote",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetAppLinksJson(val *string) {
	if err := j.validateSetAppLinksJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appLinksJson",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetAppSettingsJson(val *string) {
	if err := j.validateSetAppSettingsJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appSettingsJson",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetAutoSubmitToolbar(val interface{}) {
	if err := j.validateSetAutoSubmitToolbarParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoSubmitToolbar",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetCredentialsScheme(val *string) {
	if err := j.validateSetCredentialsSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialsScheme",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetEnduserNote(val *string) {
	if err := j.validateSetEnduserNoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enduserNote",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetGroups(val *[]*string) {
	if err := j.validateSetGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groups",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetHideIos(val interface{}) {
	if err := j.validateSetHideIosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideIos",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetHideWeb(val interface{}) {
	if err := j.validateSetHideWebParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideWeb",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetLogo(val *string) {
	if err := j.validateSetLogoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logo",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetPreconfiguredApp(val *string) {
	if err := j.validateSetPreconfiguredAppParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preconfiguredApp",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetRevealPassword(val interface{}) {
	if err := j.validateSetRevealPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revealPassword",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetSharedPassword(val *string) {
	if err := j.validateSetSharedPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedPassword",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetSharedUsername(val *string) {
	if err := j.validateSetSharedUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedUsername",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetSignOnRedirectUrl(val *string) {
	if err := j.validateSetSignOnRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signOnRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetSignOnUrl(val *string) {
	if err := j.validateSetSignOnUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signOnUrl",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetSkipGroups(val interface{}) {
	if err := j.validateSetSkipGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipGroups",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetSkipUsers(val interface{}) {
	if err := j.validateSetSkipUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipUsers",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetUserNameTemplate(val *string) {
	if err := j.validateSetUserNameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplate",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetUserNameTemplatePushStatus(val *string) {
	if err := j.validateSetUserNameTemplatePushStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplatePushStatus",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetUserNameTemplateSuffix(val *string) {
	if err := j.validateSetUserNameTemplateSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplateSuffix",
		val,
	)
}

func (j *jsiiProxy_AppAutoLogin)SetUserNameTemplateType(val *string) {
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
func AppAutoLogin_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppAutoLogin_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"okta.appAutoLogin.AppAutoLogin",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppAutoLogin_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppAutoLogin_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"okta.appAutoLogin.AppAutoLogin",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppAutoLogin_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppAutoLogin_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"okta.appAutoLogin.AppAutoLogin",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppAutoLogin_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"okta.appAutoLogin.AppAutoLogin",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppAutoLogin) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppAutoLogin) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppAutoLogin) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppAutoLogin) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppAutoLogin) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppAutoLogin) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppAutoLogin) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppAutoLogin) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppAutoLogin) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppAutoLogin) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppAutoLogin) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppAutoLogin) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppAutoLogin) PutTimeouts(value *AppAutoLoginTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppAutoLogin) PutUsers(value interface{}) {
	if err := a.validatePutUsersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUsers",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetAccessibilityErrorRedirectUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessibilityErrorRedirectUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetAccessibilityLoginRedirectUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessibilityLoginRedirectUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetAccessibilitySelfService() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessibilitySelfService",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetAdminNote() {
	_jsii_.InvokeVoid(
		a,
		"resetAdminNote",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetAppLinksJson() {
	_jsii_.InvokeVoid(
		a,
		"resetAppLinksJson",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetAppSettingsJson() {
	_jsii_.InvokeVoid(
		a,
		"resetAppSettingsJson",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetAutoSubmitToolbar() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoSubmitToolbar",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetCredentialsScheme() {
	_jsii_.InvokeVoid(
		a,
		"resetCredentialsScheme",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetEnduserNote() {
	_jsii_.InvokeVoid(
		a,
		"resetEnduserNote",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetHideIos() {
	_jsii_.InvokeVoid(
		a,
		"resetHideIos",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetHideWeb() {
	_jsii_.InvokeVoid(
		a,
		"resetHideWeb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetLogo() {
	_jsii_.InvokeVoid(
		a,
		"resetLogo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetPreconfiguredApp() {
	_jsii_.InvokeVoid(
		a,
		"resetPreconfiguredApp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetRevealPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetRevealPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetSharedPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetSharedPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetSharedUsername() {
	_jsii_.InvokeVoid(
		a,
		"resetSharedUsername",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetSignOnRedirectUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetSignOnRedirectUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetSignOnUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetSignOnUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetSkipGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetSkipUsers() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipUsers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetUserNameTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetUserNameTemplatePushStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplatePushStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetUserNameTemplateSuffix() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplateSuffix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetUserNameTemplateType() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplateType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) ResetUsers() {
	_jsii_.InvokeVoid(
		a,
		"resetUsers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppAutoLogin) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppAutoLogin) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppAutoLogin) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppAutoLogin) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


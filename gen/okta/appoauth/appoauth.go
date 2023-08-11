package appoauth

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/appoauth/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/app_oauth okta_app_oauth}.
type AppOauth interface {
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
	AuthenticationPolicy() *string
	SetAuthenticationPolicy(val *string)
	AuthenticationPolicyInput() *string
	AutoKeyRotation() interface{}
	SetAutoKeyRotation(val interface{})
	AutoKeyRotationInput() interface{}
	AutoSubmitToolbar() interface{}
	SetAutoSubmitToolbar(val interface{})
	AutoSubmitToolbarInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientBasicSecret() *string
	SetClientBasicSecret(val *string)
	ClientBasicSecretInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	ClientUri() *string
	SetClientUri(val *string)
	ClientUriInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConsentMethod() *string
	SetConsentMethod(val *string)
	ConsentMethodInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomClientId() *string
	SetCustomClientId(val *string)
	CustomClientIdInput() *string
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
	GrantTypes() *[]*string
	SetGrantTypes(val *[]*string)
	GrantTypesInput() *[]*string
	Groups() *[]*string
	SetGroups(val *[]*string)
	GroupsClaim() AppOauthGroupsClaimOutputReference
	GroupsClaimInput() *AppOauthGroupsClaim
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
	ImplicitAssignment() interface{}
	SetImplicitAssignment(val interface{})
	ImplicitAssignmentInput() interface{}
	IssuerMode() *string
	SetIssuerMode(val *string)
	IssuerModeInput() *string
	Jwks() AppOauthJwksList
	JwksInput() interface{}
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoginMode() *string
	SetLoginMode(val *string)
	LoginModeInput() *string
	LoginScopes() *[]*string
	SetLoginScopes(val *[]*string)
	LoginScopesInput() *[]*string
	LoginUri() *string
	SetLoginUri(val *string)
	LoginUriInput() *string
	Logo() *string
	SetLogo(val *string)
	LogoInput() *string
	LogoUri() *string
	SetLogoUri(val *string)
	LogoUriInput() *string
	LogoUrl() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	OmitSecret() interface{}
	SetOmitSecret(val interface{})
	OmitSecretInput() interface{}
	PkceRequired() interface{}
	SetPkceRequired(val interface{})
	PkceRequiredInput() interface{}
	PolicyUri() *string
	SetPolicyUri(val *string)
	PolicyUriInput() *string
	PostLogoutRedirectUris() *[]*string
	SetPostLogoutRedirectUris(val *[]*string)
	PostLogoutRedirectUrisInput() *[]*string
	Profile() *string
	SetProfile(val *string)
	ProfileInput() *string
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
	RedirectUris() *[]*string
	SetRedirectUris(val *[]*string)
	RedirectUrisInput() *[]*string
	RefreshTokenLeeway() *float64
	SetRefreshTokenLeeway(val *float64)
	RefreshTokenLeewayInput() *float64
	RefreshTokenRotation() *string
	SetRefreshTokenRotation(val *string)
	RefreshTokenRotationInput() *string
	ResponseTypes() *[]*string
	SetResponseTypes(val *[]*string)
	ResponseTypesInput() *[]*string
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
	Timeouts() AppOauthTimeoutsOutputReference
	TimeoutsInput() interface{}
	TokenEndpointAuthMethod() *string
	SetTokenEndpointAuthMethod(val *string)
	TokenEndpointAuthMethodInput() *string
	TosUri() *string
	SetTosUri(val *string)
	TosUriInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	Users() AppOauthUsersList
	UsersInput() interface{}
	WildcardRedirect() *string
	SetWildcardRedirect(val *string)
	WildcardRedirectInput() *string
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
	PutGroupsClaim(value *AppOauthGroupsClaim)
	PutJwks(value interface{})
	PutTimeouts(value *AppOauthTimeouts)
	PutUsers(value interface{})
	ResetAccessibilityErrorRedirectUrl()
	ResetAccessibilityLoginRedirectUrl()
	ResetAccessibilitySelfService()
	ResetAdminNote()
	ResetAppLinksJson()
	ResetAppSettingsJson()
	ResetAuthenticationPolicy()
	ResetAutoKeyRotation()
	ResetAutoSubmitToolbar()
	ResetClientBasicSecret()
	ResetClientId()
	ResetClientUri()
	ResetConsentMethod()
	ResetCustomClientId()
	ResetEnduserNote()
	ResetGrantTypes()
	ResetGroups()
	ResetGroupsClaim()
	ResetHideIos()
	ResetHideWeb()
	ResetId()
	ResetImplicitAssignment()
	ResetIssuerMode()
	ResetJwks()
	ResetLoginMode()
	ResetLoginScopes()
	ResetLoginUri()
	ResetLogo()
	ResetLogoUri()
	ResetOmitSecret()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPkceRequired()
	ResetPolicyUri()
	ResetPostLogoutRedirectUris()
	ResetProfile()
	ResetRedirectUris()
	ResetRefreshTokenLeeway()
	ResetRefreshTokenRotation()
	ResetResponseTypes()
	ResetSkipGroups()
	ResetSkipUsers()
	ResetStatus()
	ResetTimeouts()
	ResetTokenEndpointAuthMethod()
	ResetTosUri()
	ResetUserNameTemplate()
	ResetUserNameTemplatePushStatus()
	ResetUserNameTemplateSuffix()
	ResetUserNameTemplateType()
	ResetUsers()
	ResetWildcardRedirect()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AppOauth
type jsiiProxy_AppOauth struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppOauth) AccessibilityErrorRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityErrorRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) AccessibilityErrorRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityErrorRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) AccessibilityLoginRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityLoginRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) AccessibilityLoginRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityLoginRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) AccessibilitySelfService() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessibilitySelfService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) AccessibilitySelfServiceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessibilitySelfServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) AdminNote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminNote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) AdminNoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminNoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) AppLinksJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLinksJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) AppLinksJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLinksJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) AppSettingsJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSettingsJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) AppSettingsJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSettingsJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) AuthenticationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) AuthenticationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) AutoKeyRotation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoKeyRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) AutoKeyRotationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoKeyRotationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) AutoSubmitToolbar() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSubmitToolbar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) AutoSubmitToolbarInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSubmitToolbarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) ClientBasicSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientBasicSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) ClientBasicSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientBasicSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) ClientUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) ClientUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) ConsentMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consentMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) ConsentMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consentMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) CustomClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) CustomClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) EnduserNote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enduserNote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) EnduserNoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enduserNoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) GrantTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grantTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) GrantTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grantTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) Groups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) GroupsClaim() AppOauthGroupsClaimOutputReference {
	var returns AppOauthGroupsClaimOutputReference
	_jsii_.Get(
		j,
		"groupsClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) GroupsClaimInput() *AppOauthGroupsClaim {
	var returns *AppOauthGroupsClaim
	_jsii_.Get(
		j,
		"groupsClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) GroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) HideIos() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) HideIosInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) HideWeb() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideWeb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) HideWebInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideWebInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) ImplicitAssignment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"implicitAssignment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) ImplicitAssignmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"implicitAssignmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) IssuerMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) IssuerModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) Jwks() AppOauthJwksList {
	var returns AppOauthJwksList
	_jsii_.Get(
		j,
		"jwks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) JwksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jwksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) LoginMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) LoginModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) LoginScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loginScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) LoginScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loginScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) LoginUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) LoginUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) Logo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) LogoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) LogoUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) LogoUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) LogoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) OmitSecret() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"omitSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) OmitSecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"omitSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) PkceRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pkceRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) PkceRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pkceRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) PolicyUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) PolicyUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) PostLogoutRedirectUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"postLogoutRedirectUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) PostLogoutRedirectUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"postLogoutRedirectUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) RedirectUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"redirectUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) RedirectUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"redirectUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) RefreshTokenLeeway() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTokenLeeway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) RefreshTokenLeewayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTokenLeewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) RefreshTokenRotation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshTokenRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) RefreshTokenRotationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshTokenRotationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) ResponseTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"responseTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) ResponseTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"responseTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) SignOnMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signOnMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) SkipGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) SkipGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) SkipUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) SkipUsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) Timeouts() AppOauthTimeoutsOutputReference {
	var returns AppOauthTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) TokenEndpointAuthMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpointAuthMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) TokenEndpointAuthMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpointAuthMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) TosUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tosUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) TosUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tosUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) UserNameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) UserNameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) UserNameTemplatePushStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplatePushStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) UserNameTemplatePushStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplatePushStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) UserNameTemplateSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) UserNameTemplateSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) UserNameTemplateType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) UserNameTemplateTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) Users() AppOauthUsersList {
	var returns AppOauthUsersList
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) UsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) WildcardRedirect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wildcardRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppOauth) WildcardRedirectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wildcardRedirectInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/app_oauth okta_app_oauth} Resource.
func NewAppOauth(scope constructs.Construct, id *string, config *AppOauthConfig) AppOauth {
	_init_.Initialize()

	if err := validateNewAppOauthParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppOauth{}

	_jsii_.Create(
		"@cdktf/provider-okta.appOauth.AppOauth",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/app_oauth okta_app_oauth} Resource.
func NewAppOauth_Override(a AppOauth, scope constructs.Construct, id *string, config *AppOauthConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.appOauth.AppOauth",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppOauth)SetAccessibilityErrorRedirectUrl(val *string) {
	if err := j.validateSetAccessibilityErrorRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilityErrorRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetAccessibilityLoginRedirectUrl(val *string) {
	if err := j.validateSetAccessibilityLoginRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilityLoginRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetAccessibilitySelfService(val interface{}) {
	if err := j.validateSetAccessibilitySelfServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilitySelfService",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetAdminNote(val *string) {
	if err := j.validateSetAdminNoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminNote",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetAppLinksJson(val *string) {
	if err := j.validateSetAppLinksJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appLinksJson",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetAppSettingsJson(val *string) {
	if err := j.validateSetAppSettingsJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appSettingsJson",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetAuthenticationPolicy(val *string) {
	if err := j.validateSetAuthenticationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationPolicy",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetAutoKeyRotation(val interface{}) {
	if err := j.validateSetAutoKeyRotationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoKeyRotation",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetAutoSubmitToolbar(val interface{}) {
	if err := j.validateSetAutoSubmitToolbarParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoSubmitToolbar",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetClientBasicSecret(val *string) {
	if err := j.validateSetClientBasicSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientBasicSecret",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetClientUri(val *string) {
	if err := j.validateSetClientUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientUri",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetConsentMethod(val *string) {
	if err := j.validateSetConsentMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consentMethod",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetCustomClientId(val *string) {
	if err := j.validateSetCustomClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customClientId",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetEnduserNote(val *string) {
	if err := j.validateSetEnduserNoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enduserNote",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetGrantTypes(val *[]*string) {
	if err := j.validateSetGrantTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grantTypes",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetGroups(val *[]*string) {
	if err := j.validateSetGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groups",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetHideIos(val interface{}) {
	if err := j.validateSetHideIosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideIos",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetHideWeb(val interface{}) {
	if err := j.validateSetHideWebParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideWeb",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetImplicitAssignment(val interface{}) {
	if err := j.validateSetImplicitAssignmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"implicitAssignment",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetIssuerMode(val *string) {
	if err := j.validateSetIssuerModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerMode",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetLoginMode(val *string) {
	if err := j.validateSetLoginModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginMode",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetLoginScopes(val *[]*string) {
	if err := j.validateSetLoginScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginScopes",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetLoginUri(val *string) {
	if err := j.validateSetLoginUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginUri",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetLogo(val *string) {
	if err := j.validateSetLogoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logo",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetLogoUri(val *string) {
	if err := j.validateSetLogoUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logoUri",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetOmitSecret(val interface{}) {
	if err := j.validateSetOmitSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"omitSecret",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetPkceRequired(val interface{}) {
	if err := j.validateSetPkceRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pkceRequired",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetPolicyUri(val *string) {
	if err := j.validateSetPolicyUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyUri",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetPostLogoutRedirectUris(val *[]*string) {
	if err := j.validateSetPostLogoutRedirectUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postLogoutRedirectUris",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetProfile(val *string) {
	if err := j.validateSetProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetRedirectUris(val *[]*string) {
	if err := j.validateSetRedirectUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectUris",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetRefreshTokenLeeway(val *float64) {
	if err := j.validateSetRefreshTokenLeewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshTokenLeeway",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetRefreshTokenRotation(val *string) {
	if err := j.validateSetRefreshTokenRotationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshTokenRotation",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetResponseTypes(val *[]*string) {
	if err := j.validateSetResponseTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseTypes",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetSkipGroups(val interface{}) {
	if err := j.validateSetSkipGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipGroups",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetSkipUsers(val interface{}) {
	if err := j.validateSetSkipUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipUsers",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetTokenEndpointAuthMethod(val *string) {
	if err := j.validateSetTokenEndpointAuthMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenEndpointAuthMethod",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetTosUri(val *string) {
	if err := j.validateSetTosUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tosUri",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetUserNameTemplate(val *string) {
	if err := j.validateSetUserNameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplate",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetUserNameTemplatePushStatus(val *string) {
	if err := j.validateSetUserNameTemplatePushStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplatePushStatus",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetUserNameTemplateSuffix(val *string) {
	if err := j.validateSetUserNameTemplateSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplateSuffix",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetUserNameTemplateType(val *string) {
	if err := j.validateSetUserNameTemplateTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplateType",
		val,
	)
}

func (j *jsiiProxy_AppOauth)SetWildcardRedirect(val *string) {
	if err := j.validateSetWildcardRedirectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wildcardRedirect",
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
func AppOauth_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppOauth_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.appOauth.AppOauth",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppOauth_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppOauth_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.appOauth.AppOauth",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppOauth_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppOauth_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.appOauth.AppOauth",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppOauth_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.appOauth.AppOauth",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppOauth) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppOauth) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppOauth) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppOauth) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppOauth) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppOauth) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppOauth) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppOauth) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppOauth) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppOauth) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppOauth) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppOauth) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppOauth) PutGroupsClaim(value *AppOauthGroupsClaim) {
	if err := a.validatePutGroupsClaimParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGroupsClaim",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppOauth) PutJwks(value interface{}) {
	if err := a.validatePutJwksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJwks",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppOauth) PutTimeouts(value *AppOauthTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppOauth) PutUsers(value interface{}) {
	if err := a.validatePutUsersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUsers",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppOauth) ResetAccessibilityErrorRedirectUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessibilityErrorRedirectUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetAccessibilityLoginRedirectUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessibilityLoginRedirectUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetAccessibilitySelfService() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessibilitySelfService",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetAdminNote() {
	_jsii_.InvokeVoid(
		a,
		"resetAdminNote",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetAppLinksJson() {
	_jsii_.InvokeVoid(
		a,
		"resetAppLinksJson",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetAppSettingsJson() {
	_jsii_.InvokeVoid(
		a,
		"resetAppSettingsJson",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetAuthenticationPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthenticationPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetAutoKeyRotation() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoKeyRotation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetAutoSubmitToolbar() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoSubmitToolbar",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetClientBasicSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetClientBasicSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetClientUri() {
	_jsii_.InvokeVoid(
		a,
		"resetClientUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetConsentMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetConsentMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetCustomClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetEnduserNote() {
	_jsii_.InvokeVoid(
		a,
		"resetEnduserNote",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetGrantTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetGrantTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetGroupsClaim() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupsClaim",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetHideIos() {
	_jsii_.InvokeVoid(
		a,
		"resetHideIos",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetHideWeb() {
	_jsii_.InvokeVoid(
		a,
		"resetHideWeb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetImplicitAssignment() {
	_jsii_.InvokeVoid(
		a,
		"resetImplicitAssignment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetIssuerMode() {
	_jsii_.InvokeVoid(
		a,
		"resetIssuerMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetJwks() {
	_jsii_.InvokeVoid(
		a,
		"resetJwks",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetLoginMode() {
	_jsii_.InvokeVoid(
		a,
		"resetLoginMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetLoginScopes() {
	_jsii_.InvokeVoid(
		a,
		"resetLoginScopes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetLoginUri() {
	_jsii_.InvokeVoid(
		a,
		"resetLoginUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetLogo() {
	_jsii_.InvokeVoid(
		a,
		"resetLogo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetLogoUri() {
	_jsii_.InvokeVoid(
		a,
		"resetLogoUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetOmitSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetOmitSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetPkceRequired() {
	_jsii_.InvokeVoid(
		a,
		"resetPkceRequired",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetPolicyUri() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicyUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetPostLogoutRedirectUris() {
	_jsii_.InvokeVoid(
		a,
		"resetPostLogoutRedirectUris",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetProfile() {
	_jsii_.InvokeVoid(
		a,
		"resetProfile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetRedirectUris() {
	_jsii_.InvokeVoid(
		a,
		"resetRedirectUris",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetRefreshTokenLeeway() {
	_jsii_.InvokeVoid(
		a,
		"resetRefreshTokenLeeway",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetRefreshTokenRotation() {
	_jsii_.InvokeVoid(
		a,
		"resetRefreshTokenRotation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetResponseTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetSkipGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetSkipUsers() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipUsers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetTokenEndpointAuthMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenEndpointAuthMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetTosUri() {
	_jsii_.InvokeVoid(
		a,
		"resetTosUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetUserNameTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetUserNameTemplatePushStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplatePushStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetUserNameTemplateSuffix() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplateSuffix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetUserNameTemplateType() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplateType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetUsers() {
	_jsii_.InvokeVoid(
		a,
		"resetUsers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) ResetWildcardRedirect() {
	_jsii_.InvokeVoid(
		a,
		"resetWildcardRedirect",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppOauth) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppOauth) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppOauth) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppOauth) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


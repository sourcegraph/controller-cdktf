package oauthapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/oauthapp/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/okta/r/oauth_app okta_oauth_app}.
type OauthApp interface {
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	GroupsClaim() OauthAppGroupsClaimOutputReference
	GroupsClaimInput() *OauthAppGroupsClaim
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
	Jwks() OauthAppJwksList
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
	Timeouts() OauthAppTimeoutsOutputReference
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
	Users() OauthAppUsersList
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
	PutGroupsClaim(value *OauthAppGroupsClaim)
	PutJwks(value interface{})
	PutTimeouts(value *OauthAppTimeouts)
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

// The jsii proxy struct for OauthApp
type jsiiProxy_OauthApp struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OauthApp) AccessibilityErrorRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityErrorRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) AccessibilityErrorRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityErrorRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) AccessibilityLoginRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityLoginRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) AccessibilityLoginRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityLoginRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) AccessibilitySelfService() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessibilitySelfService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) AccessibilitySelfServiceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessibilitySelfServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) AdminNote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminNote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) AdminNoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminNoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) AppLinksJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLinksJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) AppLinksJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLinksJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) AppSettingsJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSettingsJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) AppSettingsJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSettingsJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) AuthenticationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) AuthenticationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) AutoKeyRotation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoKeyRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) AutoKeyRotationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoKeyRotationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) AutoSubmitToolbar() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSubmitToolbar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) AutoSubmitToolbarInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSubmitToolbarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) ClientBasicSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientBasicSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) ClientBasicSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientBasicSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) ClientUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) ClientUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) ConsentMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consentMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) ConsentMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consentMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) CustomClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) CustomClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) EnduserNote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enduserNote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) EnduserNoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enduserNoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) GrantTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grantTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) GrantTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grantTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) Groups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) GroupsClaim() OauthAppGroupsClaimOutputReference {
	var returns OauthAppGroupsClaimOutputReference
	_jsii_.Get(
		j,
		"groupsClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) GroupsClaimInput() *OauthAppGroupsClaim {
	var returns *OauthAppGroupsClaim
	_jsii_.Get(
		j,
		"groupsClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) GroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) HideIos() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) HideIosInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) HideWeb() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideWeb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) HideWebInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideWebInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) ImplicitAssignment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"implicitAssignment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) ImplicitAssignmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"implicitAssignmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) IssuerMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) IssuerModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) Jwks() OauthAppJwksList {
	var returns OauthAppJwksList
	_jsii_.Get(
		j,
		"jwks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) JwksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jwksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) LoginMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) LoginModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) LoginScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loginScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) LoginScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loginScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) LoginUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) LoginUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) Logo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) LogoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) LogoUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) LogoUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) LogoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) OmitSecret() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"omitSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) OmitSecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"omitSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) PkceRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pkceRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) PkceRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pkceRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) PolicyUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) PolicyUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) PostLogoutRedirectUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"postLogoutRedirectUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) PostLogoutRedirectUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"postLogoutRedirectUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) RedirectUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"redirectUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) RedirectUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"redirectUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) RefreshTokenLeeway() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTokenLeeway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) RefreshTokenLeewayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTokenLeewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) RefreshTokenRotation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshTokenRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) RefreshTokenRotationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshTokenRotationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) ResponseTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"responseTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) ResponseTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"responseTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) SignOnMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signOnMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) SkipGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) SkipGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) SkipUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) SkipUsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) Timeouts() OauthAppTimeoutsOutputReference {
	var returns OauthAppTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) TokenEndpointAuthMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpointAuthMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) TokenEndpointAuthMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpointAuthMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) TosUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tosUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) TosUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tosUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) UserNameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) UserNameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) UserNameTemplatePushStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplatePushStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) UserNameTemplatePushStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplatePushStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) UserNameTemplateSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) UserNameTemplateSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) UserNameTemplateType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) UserNameTemplateTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) Users() OauthAppUsersList {
	var returns OauthAppUsersList
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) UsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) WildcardRedirect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wildcardRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthApp) WildcardRedirectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wildcardRedirectInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/okta/r/oauth_app okta_oauth_app} Resource.
func NewOauthApp(scope constructs.Construct, id *string, config *OauthAppConfig) OauthApp {
	_init_.Initialize()

	if err := validateNewOauthAppParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OauthApp{}

	_jsii_.Create(
		"okta.oauthApp.OauthApp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/okta/r/oauth_app okta_oauth_app} Resource.
func NewOauthApp_Override(o OauthApp, scope constructs.Construct, id *string, config *OauthAppConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"okta.oauthApp.OauthApp",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OauthApp)SetAccessibilityErrorRedirectUrl(val *string) {
	if err := j.validateSetAccessibilityErrorRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilityErrorRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetAccessibilityLoginRedirectUrl(val *string) {
	if err := j.validateSetAccessibilityLoginRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilityLoginRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetAccessibilitySelfService(val interface{}) {
	if err := j.validateSetAccessibilitySelfServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilitySelfService",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetAdminNote(val *string) {
	if err := j.validateSetAdminNoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminNote",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetAppLinksJson(val *string) {
	if err := j.validateSetAppLinksJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appLinksJson",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetAppSettingsJson(val *string) {
	if err := j.validateSetAppSettingsJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appSettingsJson",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetAuthenticationPolicy(val *string) {
	if err := j.validateSetAuthenticationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationPolicy",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetAutoKeyRotation(val interface{}) {
	if err := j.validateSetAutoKeyRotationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoKeyRotation",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetAutoSubmitToolbar(val interface{}) {
	if err := j.validateSetAutoSubmitToolbarParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoSubmitToolbar",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetClientBasicSecret(val *string) {
	if err := j.validateSetClientBasicSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientBasicSecret",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetClientUri(val *string) {
	if err := j.validateSetClientUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientUri",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetConsentMethod(val *string) {
	if err := j.validateSetConsentMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consentMethod",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetCustomClientId(val *string) {
	if err := j.validateSetCustomClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customClientId",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetEnduserNote(val *string) {
	if err := j.validateSetEnduserNoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enduserNote",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetGrantTypes(val *[]*string) {
	if err := j.validateSetGrantTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grantTypes",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetGroups(val *[]*string) {
	if err := j.validateSetGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groups",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetHideIos(val interface{}) {
	if err := j.validateSetHideIosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideIos",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetHideWeb(val interface{}) {
	if err := j.validateSetHideWebParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideWeb",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetImplicitAssignment(val interface{}) {
	if err := j.validateSetImplicitAssignmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"implicitAssignment",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetIssuerMode(val *string) {
	if err := j.validateSetIssuerModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerMode",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetLoginMode(val *string) {
	if err := j.validateSetLoginModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginMode",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetLoginScopes(val *[]*string) {
	if err := j.validateSetLoginScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginScopes",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetLoginUri(val *string) {
	if err := j.validateSetLoginUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginUri",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetLogo(val *string) {
	if err := j.validateSetLogoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logo",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetLogoUri(val *string) {
	if err := j.validateSetLogoUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logoUri",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetOmitSecret(val interface{}) {
	if err := j.validateSetOmitSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"omitSecret",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetPkceRequired(val interface{}) {
	if err := j.validateSetPkceRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pkceRequired",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetPolicyUri(val *string) {
	if err := j.validateSetPolicyUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyUri",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetPostLogoutRedirectUris(val *[]*string) {
	if err := j.validateSetPostLogoutRedirectUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postLogoutRedirectUris",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetProfile(val *string) {
	if err := j.validateSetProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetRedirectUris(val *[]*string) {
	if err := j.validateSetRedirectUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectUris",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetRefreshTokenLeeway(val *float64) {
	if err := j.validateSetRefreshTokenLeewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshTokenLeeway",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetRefreshTokenRotation(val *string) {
	if err := j.validateSetRefreshTokenRotationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshTokenRotation",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetResponseTypes(val *[]*string) {
	if err := j.validateSetResponseTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseTypes",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetSkipGroups(val interface{}) {
	if err := j.validateSetSkipGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipGroups",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetSkipUsers(val interface{}) {
	if err := j.validateSetSkipUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipUsers",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetTokenEndpointAuthMethod(val *string) {
	if err := j.validateSetTokenEndpointAuthMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenEndpointAuthMethod",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetTosUri(val *string) {
	if err := j.validateSetTosUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tosUri",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetUserNameTemplate(val *string) {
	if err := j.validateSetUserNameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplate",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetUserNameTemplatePushStatus(val *string) {
	if err := j.validateSetUserNameTemplatePushStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplatePushStatus",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetUserNameTemplateSuffix(val *string) {
	if err := j.validateSetUserNameTemplateSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplateSuffix",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetUserNameTemplateType(val *string) {
	if err := j.validateSetUserNameTemplateTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplateType",
		val,
	)
}

func (j *jsiiProxy_OauthApp)SetWildcardRedirect(val *string) {
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
func OauthApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOauthApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"okta.oauthApp.OauthApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OauthApp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"okta.oauthApp.OauthApp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OauthApp) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OauthApp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthApp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthApp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthApp) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthApp) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthApp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthApp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthApp) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthApp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthApp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthApp) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OauthApp) PutGroupsClaim(value *OauthAppGroupsClaim) {
	if err := o.validatePutGroupsClaimParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGroupsClaim",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OauthApp) PutJwks(value interface{}) {
	if err := o.validatePutJwksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putJwks",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OauthApp) PutTimeouts(value *OauthAppTimeouts) {
	if err := o.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OauthApp) PutUsers(value interface{}) {
	if err := o.validatePutUsersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putUsers",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OauthApp) ResetAccessibilityErrorRedirectUrl() {
	_jsii_.InvokeVoid(
		o,
		"resetAccessibilityErrorRedirectUrl",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetAccessibilityLoginRedirectUrl() {
	_jsii_.InvokeVoid(
		o,
		"resetAccessibilityLoginRedirectUrl",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetAccessibilitySelfService() {
	_jsii_.InvokeVoid(
		o,
		"resetAccessibilitySelfService",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetAdminNote() {
	_jsii_.InvokeVoid(
		o,
		"resetAdminNote",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetAppLinksJson() {
	_jsii_.InvokeVoid(
		o,
		"resetAppLinksJson",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetAppSettingsJson() {
	_jsii_.InvokeVoid(
		o,
		"resetAppSettingsJson",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetAuthenticationPolicy() {
	_jsii_.InvokeVoid(
		o,
		"resetAuthenticationPolicy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetAutoKeyRotation() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoKeyRotation",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetAutoSubmitToolbar() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoSubmitToolbar",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetClientBasicSecret() {
	_jsii_.InvokeVoid(
		o,
		"resetClientBasicSecret",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetClientId() {
	_jsii_.InvokeVoid(
		o,
		"resetClientId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetClientUri() {
	_jsii_.InvokeVoid(
		o,
		"resetClientUri",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetConsentMethod() {
	_jsii_.InvokeVoid(
		o,
		"resetConsentMethod",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetCustomClientId() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomClientId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetEnduserNote() {
	_jsii_.InvokeVoid(
		o,
		"resetEnduserNote",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetGrantTypes() {
	_jsii_.InvokeVoid(
		o,
		"resetGrantTypes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetGroups() {
	_jsii_.InvokeVoid(
		o,
		"resetGroups",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetGroupsClaim() {
	_jsii_.InvokeVoid(
		o,
		"resetGroupsClaim",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetHideIos() {
	_jsii_.InvokeVoid(
		o,
		"resetHideIos",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetHideWeb() {
	_jsii_.InvokeVoid(
		o,
		"resetHideWeb",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetImplicitAssignment() {
	_jsii_.InvokeVoid(
		o,
		"resetImplicitAssignment",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetIssuerMode() {
	_jsii_.InvokeVoid(
		o,
		"resetIssuerMode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetJwks() {
	_jsii_.InvokeVoid(
		o,
		"resetJwks",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetLoginMode() {
	_jsii_.InvokeVoid(
		o,
		"resetLoginMode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetLoginScopes() {
	_jsii_.InvokeVoid(
		o,
		"resetLoginScopes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetLoginUri() {
	_jsii_.InvokeVoid(
		o,
		"resetLoginUri",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetLogo() {
	_jsii_.InvokeVoid(
		o,
		"resetLogo",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetLogoUri() {
	_jsii_.InvokeVoid(
		o,
		"resetLogoUri",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetOmitSecret() {
	_jsii_.InvokeVoid(
		o,
		"resetOmitSecret",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetPkceRequired() {
	_jsii_.InvokeVoid(
		o,
		"resetPkceRequired",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetPolicyUri() {
	_jsii_.InvokeVoid(
		o,
		"resetPolicyUri",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetPostLogoutRedirectUris() {
	_jsii_.InvokeVoid(
		o,
		"resetPostLogoutRedirectUris",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetProfile() {
	_jsii_.InvokeVoid(
		o,
		"resetProfile",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetRedirectUris() {
	_jsii_.InvokeVoid(
		o,
		"resetRedirectUris",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetRefreshTokenLeeway() {
	_jsii_.InvokeVoid(
		o,
		"resetRefreshTokenLeeway",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetRefreshTokenRotation() {
	_jsii_.InvokeVoid(
		o,
		"resetRefreshTokenRotation",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetResponseTypes() {
	_jsii_.InvokeVoid(
		o,
		"resetResponseTypes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetSkipGroups() {
	_jsii_.InvokeVoid(
		o,
		"resetSkipGroups",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetSkipUsers() {
	_jsii_.InvokeVoid(
		o,
		"resetSkipUsers",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetStatus() {
	_jsii_.InvokeVoid(
		o,
		"resetStatus",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetTimeouts() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetTokenEndpointAuthMethod() {
	_jsii_.InvokeVoid(
		o,
		"resetTokenEndpointAuthMethod",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetTosUri() {
	_jsii_.InvokeVoid(
		o,
		"resetTosUri",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetUserNameTemplate() {
	_jsii_.InvokeVoid(
		o,
		"resetUserNameTemplate",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetUserNameTemplatePushStatus() {
	_jsii_.InvokeVoid(
		o,
		"resetUserNameTemplatePushStatus",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetUserNameTemplateSuffix() {
	_jsii_.InvokeVoid(
		o,
		"resetUserNameTemplateSuffix",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetUserNameTemplateType() {
	_jsii_.InvokeVoid(
		o,
		"resetUserNameTemplateType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetUsers() {
	_jsii_.InvokeVoid(
		o,
		"resetUsers",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) ResetWildcardRedirect() {
	_jsii_.InvokeVoid(
		o,
		"resetWildcardRedirect",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthApp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthApp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthApp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


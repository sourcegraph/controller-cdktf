package appsaml

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/appsaml/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/app_saml okta_app_saml}.
type AppSaml interface {
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
	AcsEndpoints() *[]*string
	SetAcsEndpoints(val *[]*string)
	AcsEndpointsInput() *[]*string
	AdminNote() *string
	SetAdminNote(val *string)
	AdminNoteInput() *string
	AppLinksJson() *string
	SetAppLinksJson(val *string)
	AppLinksJsonInput() *string
	AppSettingsJson() *string
	SetAppSettingsJson(val *string)
	AppSettingsJsonInput() *string
	AssertionSigned() interface{}
	SetAssertionSigned(val interface{})
	AssertionSignedInput() interface{}
	AttributeStatements() AppSamlAttributeStatementsList
	AttributeStatementsInput() interface{}
	Audience() *string
	SetAudience(val *string)
	AudienceInput() *string
	AuthenticationPolicy() *string
	SetAuthenticationPolicy(val *string)
	AuthenticationPolicyInput() *string
	AuthnContextClassRef() *string
	SetAuthnContextClassRef(val *string)
	AuthnContextClassRefInput() *string
	AutoSubmitToolbar() interface{}
	SetAutoSubmitToolbar(val interface{})
	AutoSubmitToolbarInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Certificate() *string
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
	DefaultRelayState() *string
	SetDefaultRelayState(val *string)
	DefaultRelayStateInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Destination() *string
	SetDestination(val *string)
	DestinationInput() *string
	DigestAlgorithm() *string
	SetDigestAlgorithm(val *string)
	DigestAlgorithmInput() *string
	EmbedUrl() *string
	EnduserNote() *string
	SetEnduserNote(val *string)
	EnduserNoteInput() *string
	EntityKey() *string
	EntityUrl() *string
	Features() *[]*string
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
	HonorForceAuthn() interface{}
	SetHonorForceAuthn(val interface{})
	HonorForceAuthnInput() interface{}
	HttpPostBinding() *string
	HttpRedirectBinding() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IdpIssuer() *string
	SetIdpIssuer(val *string)
	IdpIssuerInput() *string
	ImplicitAssignment() interface{}
	SetImplicitAssignment(val interface{})
	ImplicitAssignmentInput() interface{}
	InlineHookId() *string
	SetInlineHookId(val *string)
	InlineHookIdInput() *string
	KeyId() *string
	KeyName() *string
	SetKeyName(val *string)
	KeyNameInput() *string
	Keys() AppSamlKeysList
	KeyYearsValid() *float64
	SetKeyYearsValid(val *float64)
	KeyYearsValidInput() *float64
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
	Metadata() *string
	MetadataUrl() *string
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
	Recipient() *string
	SetRecipient(val *string)
	RecipientInput() *string
	RequestCompressed() interface{}
	SetRequestCompressed(val interface{})
	RequestCompressedInput() interface{}
	ResponseSigned() interface{}
	SetResponseSigned(val interface{})
	ResponseSignedInput() interface{}
	SamlVersion() *string
	SetSamlVersion(val *string)
	SamlVersionInput() *string
	SignatureAlgorithm() *string
	SetSignatureAlgorithm(val *string)
	SignatureAlgorithmInput() *string
	SignOnMode() *string
	SingleLogoutCertificate() *string
	SetSingleLogoutCertificate(val *string)
	SingleLogoutCertificateInput() *string
	SingleLogoutIssuer() *string
	SetSingleLogoutIssuer(val *string)
	SingleLogoutIssuerInput() *string
	SingleLogoutUrl() *string
	SetSingleLogoutUrl(val *string)
	SingleLogoutUrlInput() *string
	SkipGroups() interface{}
	SetSkipGroups(val interface{})
	SkipGroupsInput() interface{}
	SkipUsers() interface{}
	SetSkipUsers(val interface{})
	SkipUsersInput() interface{}
	SpIssuer() *string
	SetSpIssuer(val *string)
	SpIssuerInput() *string
	SsoUrl() *string
	SetSsoUrl(val *string)
	SsoUrlInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	SubjectNameIdFormat() *string
	SetSubjectNameIdFormat(val *string)
	SubjectNameIdFormatInput() *string
	SubjectNameIdTemplate() *string
	SetSubjectNameIdTemplate(val *string)
	SubjectNameIdTemplateInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() AppSamlTimeoutsOutputReference
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
	Users() AppSamlUsersList
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
	PutAttributeStatements(value interface{})
	PutTimeouts(value *AppSamlTimeouts)
	PutUsers(value interface{})
	ResetAccessibilityErrorRedirectUrl()
	ResetAccessibilityLoginRedirectUrl()
	ResetAccessibilitySelfService()
	ResetAcsEndpoints()
	ResetAdminNote()
	ResetAppLinksJson()
	ResetAppSettingsJson()
	ResetAssertionSigned()
	ResetAttributeStatements()
	ResetAudience()
	ResetAuthenticationPolicy()
	ResetAuthnContextClassRef()
	ResetAutoSubmitToolbar()
	ResetDefaultRelayState()
	ResetDestination()
	ResetDigestAlgorithm()
	ResetEnduserNote()
	ResetGroups()
	ResetHideIos()
	ResetHideWeb()
	ResetHonorForceAuthn()
	ResetId()
	ResetIdpIssuer()
	ResetImplicitAssignment()
	ResetInlineHookId()
	ResetKeyName()
	ResetKeyYearsValid()
	ResetLogo()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreconfiguredApp()
	ResetRecipient()
	ResetRequestCompressed()
	ResetResponseSigned()
	ResetSamlVersion()
	ResetSignatureAlgorithm()
	ResetSingleLogoutCertificate()
	ResetSingleLogoutIssuer()
	ResetSingleLogoutUrl()
	ResetSkipGroups()
	ResetSkipUsers()
	ResetSpIssuer()
	ResetSsoUrl()
	ResetStatus()
	ResetSubjectNameIdFormat()
	ResetSubjectNameIdTemplate()
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

// The jsii proxy struct for AppSaml
type jsiiProxy_AppSaml struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppSaml) AccessibilityErrorRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityErrorRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AccessibilityErrorRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityErrorRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AccessibilityLoginRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityLoginRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AccessibilityLoginRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityLoginRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AccessibilitySelfService() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessibilitySelfService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AccessibilitySelfServiceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessibilitySelfServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AcsEndpoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acsEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AcsEndpointsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acsEndpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AdminNote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminNote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AdminNoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminNoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AppLinksJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLinksJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AppLinksJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLinksJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AppSettingsJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSettingsJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AppSettingsJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSettingsJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AssertionSigned() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assertionSigned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AssertionSignedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assertionSignedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AttributeStatements() AppSamlAttributeStatementsList {
	var returns AppSamlAttributeStatementsList
	_jsii_.Get(
		j,
		"attributeStatements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AttributeStatementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributeStatementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Audience() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AudienceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AuthenticationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AuthenticationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AuthnContextClassRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authnContextClassRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AuthnContextClassRefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authnContextClassRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AutoSubmitToolbar() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSubmitToolbar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) AutoSubmitToolbarInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSubmitToolbarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) DefaultRelayState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRelayState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) DefaultRelayStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRelayStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) DestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) DigestAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"digestAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) DigestAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"digestAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) EmbedUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embedUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) EnduserNote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enduserNote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) EnduserNoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enduserNoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) EntityKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) EntityUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Features() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"features",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Groups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) GroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) HideIos() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) HideIosInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) HideWeb() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideWeb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) HideWebInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideWebInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) HonorForceAuthn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"honorForceAuthn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) HonorForceAuthnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"honorForceAuthnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) HttpPostBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpPostBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) HttpRedirectBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRedirectBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) IdpIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) IdpIssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpIssuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) ImplicitAssignment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"implicitAssignment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) ImplicitAssignmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"implicitAssignmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) InlineHookId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inlineHookId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) InlineHookIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inlineHookIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) KeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Keys() AppSamlKeysList {
	var returns AppSamlKeysList
	_jsii_.Get(
		j,
		"keys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) KeyYearsValid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyYearsValid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) KeyYearsValidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyYearsValidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Logo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) LogoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) LogoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Metadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) MetadataUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) PreconfiguredApp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconfiguredApp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) PreconfiguredAppInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconfiguredAppInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Recipient() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recipient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) RecipientInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recipientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) RequestCompressed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestCompressed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) RequestCompressedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestCompressedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) ResponseSigned() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseSigned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) ResponseSignedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseSignedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SamlVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SamlVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SignatureAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SignatureAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SignOnMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signOnMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SingleLogoutCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleLogoutCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SingleLogoutCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleLogoutCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SingleLogoutIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleLogoutIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SingleLogoutIssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleLogoutIssuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SingleLogoutUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleLogoutUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SingleLogoutUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleLogoutUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SkipGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SkipGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SkipUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SkipUsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SpIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SpIssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spIssuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SsoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SsoUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SubjectNameIdFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectNameIdFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SubjectNameIdFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectNameIdFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SubjectNameIdTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectNameIdTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) SubjectNameIdTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectNameIdTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Timeouts() AppSamlTimeoutsOutputReference {
	var returns AppSamlTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) UserNameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) UserNameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) UserNameTemplatePushStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplatePushStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) UserNameTemplatePushStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplatePushStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) UserNameTemplateSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) UserNameTemplateSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) UserNameTemplateType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) UserNameTemplateTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) Users() AppSamlUsersList {
	var returns AppSamlUsersList
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSaml) UsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usersInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/app_saml okta_app_saml} Resource.
func NewAppSaml(scope constructs.Construct, id *string, config *AppSamlConfig) AppSaml {
	_init_.Initialize()

	if err := validateNewAppSamlParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppSaml{}

	_jsii_.Create(
		"okta.appSaml.AppSaml",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/app_saml okta_app_saml} Resource.
func NewAppSaml_Override(a AppSaml, scope constructs.Construct, id *string, config *AppSamlConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"okta.appSaml.AppSaml",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppSaml)SetAccessibilityErrorRedirectUrl(val *string) {
	if err := j.validateSetAccessibilityErrorRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilityErrorRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetAccessibilityLoginRedirectUrl(val *string) {
	if err := j.validateSetAccessibilityLoginRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilityLoginRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetAccessibilitySelfService(val interface{}) {
	if err := j.validateSetAccessibilitySelfServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilitySelfService",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetAcsEndpoints(val *[]*string) {
	if err := j.validateSetAcsEndpointsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acsEndpoints",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetAdminNote(val *string) {
	if err := j.validateSetAdminNoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminNote",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetAppLinksJson(val *string) {
	if err := j.validateSetAppLinksJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appLinksJson",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetAppSettingsJson(val *string) {
	if err := j.validateSetAppSettingsJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appSettingsJson",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetAssertionSigned(val interface{}) {
	if err := j.validateSetAssertionSignedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assertionSigned",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetAudience(val *string) {
	if err := j.validateSetAudienceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audience",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetAuthenticationPolicy(val *string) {
	if err := j.validateSetAuthenticationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationPolicy",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetAuthnContextClassRef(val *string) {
	if err := j.validateSetAuthnContextClassRefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authnContextClassRef",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetAutoSubmitToolbar(val interface{}) {
	if err := j.validateSetAutoSubmitToolbarParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoSubmitToolbar",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetDefaultRelayState(val *string) {
	if err := j.validateSetDefaultRelayStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRelayState",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetDestination(val *string) {
	if err := j.validateSetDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destination",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetDigestAlgorithm(val *string) {
	if err := j.validateSetDigestAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"digestAlgorithm",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetEnduserNote(val *string) {
	if err := j.validateSetEnduserNoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enduserNote",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetGroups(val *[]*string) {
	if err := j.validateSetGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groups",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetHideIos(val interface{}) {
	if err := j.validateSetHideIosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideIos",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetHideWeb(val interface{}) {
	if err := j.validateSetHideWebParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideWeb",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetHonorForceAuthn(val interface{}) {
	if err := j.validateSetHonorForceAuthnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"honorForceAuthn",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetIdpIssuer(val *string) {
	if err := j.validateSetIdpIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpIssuer",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetImplicitAssignment(val interface{}) {
	if err := j.validateSetImplicitAssignmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"implicitAssignment",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetInlineHookId(val *string) {
	if err := j.validateSetInlineHookIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inlineHookId",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetKeyYearsValid(val *float64) {
	if err := j.validateSetKeyYearsValidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyYearsValid",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetLogo(val *string) {
	if err := j.validateSetLogoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logo",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetPreconfiguredApp(val *string) {
	if err := j.validateSetPreconfiguredAppParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preconfiguredApp",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetRecipient(val *string) {
	if err := j.validateSetRecipientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recipient",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetRequestCompressed(val interface{}) {
	if err := j.validateSetRequestCompressedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestCompressed",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetResponseSigned(val interface{}) {
	if err := j.validateSetResponseSignedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseSigned",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetSamlVersion(val *string) {
	if err := j.validateSetSamlVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samlVersion",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetSignatureAlgorithm(val *string) {
	if err := j.validateSetSignatureAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signatureAlgorithm",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetSingleLogoutCertificate(val *string) {
	if err := j.validateSetSingleLogoutCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleLogoutCertificate",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetSingleLogoutIssuer(val *string) {
	if err := j.validateSetSingleLogoutIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleLogoutIssuer",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetSingleLogoutUrl(val *string) {
	if err := j.validateSetSingleLogoutUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleLogoutUrl",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetSkipGroups(val interface{}) {
	if err := j.validateSetSkipGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipGroups",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetSkipUsers(val interface{}) {
	if err := j.validateSetSkipUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipUsers",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetSpIssuer(val *string) {
	if err := j.validateSetSpIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spIssuer",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetSsoUrl(val *string) {
	if err := j.validateSetSsoUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssoUrl",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetSubjectNameIdFormat(val *string) {
	if err := j.validateSetSubjectNameIdFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectNameIdFormat",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetSubjectNameIdTemplate(val *string) {
	if err := j.validateSetSubjectNameIdTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectNameIdTemplate",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetUserNameTemplate(val *string) {
	if err := j.validateSetUserNameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplate",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetUserNameTemplatePushStatus(val *string) {
	if err := j.validateSetUserNameTemplatePushStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplatePushStatus",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetUserNameTemplateSuffix(val *string) {
	if err := j.validateSetUserNameTemplateSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplateSuffix",
		val,
	)
}

func (j *jsiiProxy_AppSaml)SetUserNameTemplateType(val *string) {
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
func AppSaml_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppSaml_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"okta.appSaml.AppSaml",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppSaml_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppSaml_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"okta.appSaml.AppSaml",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppSaml_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppSaml_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"okta.appSaml.AppSaml",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppSaml_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"okta.appSaml.AppSaml",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppSaml) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppSaml) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppSaml) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppSaml) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppSaml) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppSaml) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppSaml) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppSaml) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppSaml) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppSaml) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppSaml) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppSaml) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppSaml) PutAttributeStatements(value interface{}) {
	if err := a.validatePutAttributeStatementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAttributeStatements",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSaml) PutTimeouts(value *AppSamlTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSaml) PutUsers(value interface{}) {
	if err := a.validatePutUsersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUsers",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSaml) ResetAccessibilityErrorRedirectUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessibilityErrorRedirectUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetAccessibilityLoginRedirectUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessibilityLoginRedirectUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetAccessibilitySelfService() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessibilitySelfService",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetAcsEndpoints() {
	_jsii_.InvokeVoid(
		a,
		"resetAcsEndpoints",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetAdminNote() {
	_jsii_.InvokeVoid(
		a,
		"resetAdminNote",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetAppLinksJson() {
	_jsii_.InvokeVoid(
		a,
		"resetAppLinksJson",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetAppSettingsJson() {
	_jsii_.InvokeVoid(
		a,
		"resetAppSettingsJson",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetAssertionSigned() {
	_jsii_.InvokeVoid(
		a,
		"resetAssertionSigned",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetAttributeStatements() {
	_jsii_.InvokeVoid(
		a,
		"resetAttributeStatements",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetAudience() {
	_jsii_.InvokeVoid(
		a,
		"resetAudience",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetAuthenticationPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthenticationPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetAuthnContextClassRef() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthnContextClassRef",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetAutoSubmitToolbar() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoSubmitToolbar",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetDefaultRelayState() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultRelayState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetDestination() {
	_jsii_.InvokeVoid(
		a,
		"resetDestination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetDigestAlgorithm() {
	_jsii_.InvokeVoid(
		a,
		"resetDigestAlgorithm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetEnduserNote() {
	_jsii_.InvokeVoid(
		a,
		"resetEnduserNote",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetHideIos() {
	_jsii_.InvokeVoid(
		a,
		"resetHideIos",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetHideWeb() {
	_jsii_.InvokeVoid(
		a,
		"resetHideWeb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetHonorForceAuthn() {
	_jsii_.InvokeVoid(
		a,
		"resetHonorForceAuthn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetIdpIssuer() {
	_jsii_.InvokeVoid(
		a,
		"resetIdpIssuer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetImplicitAssignment() {
	_jsii_.InvokeVoid(
		a,
		"resetImplicitAssignment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetInlineHookId() {
	_jsii_.InvokeVoid(
		a,
		"resetInlineHookId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetKeyName() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetKeyYearsValid() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyYearsValid",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetLogo() {
	_jsii_.InvokeVoid(
		a,
		"resetLogo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetPreconfiguredApp() {
	_jsii_.InvokeVoid(
		a,
		"resetPreconfiguredApp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetRecipient() {
	_jsii_.InvokeVoid(
		a,
		"resetRecipient",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetRequestCompressed() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestCompressed",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetResponseSigned() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseSigned",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetSamlVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetSamlVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetSignatureAlgorithm() {
	_jsii_.InvokeVoid(
		a,
		"resetSignatureAlgorithm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetSingleLogoutCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetSingleLogoutCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetSingleLogoutIssuer() {
	_jsii_.InvokeVoid(
		a,
		"resetSingleLogoutIssuer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetSingleLogoutUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetSingleLogoutUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetSkipGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetSkipUsers() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipUsers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetSpIssuer() {
	_jsii_.InvokeVoid(
		a,
		"resetSpIssuer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetSsoUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetSsoUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetSubjectNameIdFormat() {
	_jsii_.InvokeVoid(
		a,
		"resetSubjectNameIdFormat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetSubjectNameIdTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetSubjectNameIdTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetUserNameTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetUserNameTemplatePushStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplatePushStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetUserNameTemplateSuffix() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplateSuffix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetUserNameTemplateType() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameTemplateType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) ResetUsers() {
	_jsii_.InvokeVoid(
		a,
		"resetUsers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSaml) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSaml) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSaml) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSaml) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


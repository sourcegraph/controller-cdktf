package samlapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/samlapp/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/saml_app okta_saml_app}.
type SamlApp interface {
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
	AttributeStatements() SamlAppAttributeStatementsList
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
	Keys() SamlAppKeysList
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
	Timeouts() SamlAppTimeoutsOutputReference
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
	Users() SamlAppUsersList
	UsersInput() interface{}
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
	PutAttributeStatements(value interface{})
	PutTimeouts(value *SamlAppTimeouts)
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

// The jsii proxy struct for SamlApp
type jsiiProxy_SamlApp struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SamlApp) AccessibilityErrorRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityErrorRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AccessibilityErrorRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityErrorRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AccessibilityLoginRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityLoginRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AccessibilityLoginRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityLoginRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AccessibilitySelfService() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessibilitySelfService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AccessibilitySelfServiceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessibilitySelfServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AcsEndpoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acsEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AcsEndpointsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acsEndpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AdminNote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminNote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AdminNoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminNoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AppLinksJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLinksJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AppLinksJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLinksJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AppSettingsJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSettingsJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AppSettingsJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSettingsJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AssertionSigned() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assertionSigned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AssertionSignedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assertionSignedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AttributeStatements() SamlAppAttributeStatementsList {
	var returns SamlAppAttributeStatementsList
	_jsii_.Get(
		j,
		"attributeStatements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AttributeStatementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributeStatementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Audience() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AudienceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AuthenticationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AuthenticationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AuthnContextClassRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authnContextClassRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AuthnContextClassRefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authnContextClassRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AutoSubmitToolbar() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSubmitToolbar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) AutoSubmitToolbarInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSubmitToolbarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) DefaultRelayState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRelayState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) DefaultRelayStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRelayStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) DestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) DigestAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"digestAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) DigestAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"digestAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) EmbedUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embedUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) EnduserNote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enduserNote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) EnduserNoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enduserNoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) EntityKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) EntityUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Features() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"features",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Groups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) GroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) HideIos() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) HideIosInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) HideWeb() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideWeb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) HideWebInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideWebInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) HonorForceAuthn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"honorForceAuthn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) HonorForceAuthnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"honorForceAuthnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) HttpPostBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpPostBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) HttpRedirectBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRedirectBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) IdpIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) IdpIssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpIssuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) ImplicitAssignment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"implicitAssignment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) ImplicitAssignmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"implicitAssignmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) InlineHookId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inlineHookId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) InlineHookIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inlineHookIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) KeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Keys() SamlAppKeysList {
	var returns SamlAppKeysList
	_jsii_.Get(
		j,
		"keys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) KeyYearsValid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyYearsValid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) KeyYearsValidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyYearsValidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Logo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) LogoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) LogoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Metadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) MetadataUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) PreconfiguredApp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconfiguredApp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) PreconfiguredAppInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconfiguredAppInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Recipient() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recipient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) RecipientInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recipientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) RequestCompressed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestCompressed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) RequestCompressedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestCompressedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) ResponseSigned() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseSigned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) ResponseSignedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseSignedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SamlVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SamlVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SignatureAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SignatureAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SignOnMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signOnMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SingleLogoutCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleLogoutCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SingleLogoutCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleLogoutCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SingleLogoutIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleLogoutIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SingleLogoutIssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleLogoutIssuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SingleLogoutUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleLogoutUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SingleLogoutUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleLogoutUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SkipGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SkipGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SkipUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SkipUsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SpIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SpIssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spIssuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SsoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SsoUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SubjectNameIdFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectNameIdFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SubjectNameIdFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectNameIdFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SubjectNameIdTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectNameIdTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) SubjectNameIdTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectNameIdTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Timeouts() SamlAppTimeoutsOutputReference {
	var returns SamlAppTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) UserNameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) UserNameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) UserNameTemplatePushStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplatePushStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) UserNameTemplatePushStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplatePushStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) UserNameTemplateSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) UserNameTemplateSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) UserNameTemplateType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) UserNameTemplateTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) Users() SamlAppUsersList {
	var returns SamlAppUsersList
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlApp) UsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usersInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/saml_app okta_saml_app} Resource.
func NewSamlApp(scope constructs.Construct, id *string, config *SamlAppConfig) SamlApp {
	_init_.Initialize()

	if err := validateNewSamlAppParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SamlApp{}

	_jsii_.Create(
		"@cdktf/provider-okta.samlApp.SamlApp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/saml_app okta_saml_app} Resource.
func NewSamlApp_Override(s SamlApp, scope constructs.Construct, id *string, config *SamlAppConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.samlApp.SamlApp",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SamlApp)SetAccessibilityErrorRedirectUrl(val *string) {
	if err := j.validateSetAccessibilityErrorRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilityErrorRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetAccessibilityLoginRedirectUrl(val *string) {
	if err := j.validateSetAccessibilityLoginRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilityLoginRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetAccessibilitySelfService(val interface{}) {
	if err := j.validateSetAccessibilitySelfServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessibilitySelfService",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetAcsEndpoints(val *[]*string) {
	if err := j.validateSetAcsEndpointsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acsEndpoints",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetAdminNote(val *string) {
	if err := j.validateSetAdminNoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminNote",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetAppLinksJson(val *string) {
	if err := j.validateSetAppLinksJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appLinksJson",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetAppSettingsJson(val *string) {
	if err := j.validateSetAppSettingsJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appSettingsJson",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetAssertionSigned(val interface{}) {
	if err := j.validateSetAssertionSignedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assertionSigned",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetAudience(val *string) {
	if err := j.validateSetAudienceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audience",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetAuthenticationPolicy(val *string) {
	if err := j.validateSetAuthenticationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationPolicy",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetAuthnContextClassRef(val *string) {
	if err := j.validateSetAuthnContextClassRefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authnContextClassRef",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetAutoSubmitToolbar(val interface{}) {
	if err := j.validateSetAutoSubmitToolbarParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoSubmitToolbar",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetDefaultRelayState(val *string) {
	if err := j.validateSetDefaultRelayStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRelayState",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetDestination(val *string) {
	if err := j.validateSetDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destination",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetDigestAlgorithm(val *string) {
	if err := j.validateSetDigestAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"digestAlgorithm",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetEnduserNote(val *string) {
	if err := j.validateSetEnduserNoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enduserNote",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetGroups(val *[]*string) {
	if err := j.validateSetGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groups",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetHideIos(val interface{}) {
	if err := j.validateSetHideIosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideIos",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetHideWeb(val interface{}) {
	if err := j.validateSetHideWebParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideWeb",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetHonorForceAuthn(val interface{}) {
	if err := j.validateSetHonorForceAuthnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"honorForceAuthn",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetIdpIssuer(val *string) {
	if err := j.validateSetIdpIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpIssuer",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetImplicitAssignment(val interface{}) {
	if err := j.validateSetImplicitAssignmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"implicitAssignment",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetInlineHookId(val *string) {
	if err := j.validateSetInlineHookIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inlineHookId",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetKeyYearsValid(val *float64) {
	if err := j.validateSetKeyYearsValidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyYearsValid",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetLogo(val *string) {
	if err := j.validateSetLogoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logo",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetPreconfiguredApp(val *string) {
	if err := j.validateSetPreconfiguredAppParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preconfiguredApp",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetRecipient(val *string) {
	if err := j.validateSetRecipientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recipient",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetRequestCompressed(val interface{}) {
	if err := j.validateSetRequestCompressedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestCompressed",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetResponseSigned(val interface{}) {
	if err := j.validateSetResponseSignedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseSigned",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetSamlVersion(val *string) {
	if err := j.validateSetSamlVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samlVersion",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetSignatureAlgorithm(val *string) {
	if err := j.validateSetSignatureAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signatureAlgorithm",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetSingleLogoutCertificate(val *string) {
	if err := j.validateSetSingleLogoutCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleLogoutCertificate",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetSingleLogoutIssuer(val *string) {
	if err := j.validateSetSingleLogoutIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleLogoutIssuer",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetSingleLogoutUrl(val *string) {
	if err := j.validateSetSingleLogoutUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleLogoutUrl",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetSkipGroups(val interface{}) {
	if err := j.validateSetSkipGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipGroups",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetSkipUsers(val interface{}) {
	if err := j.validateSetSkipUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipUsers",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetSpIssuer(val *string) {
	if err := j.validateSetSpIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spIssuer",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetSsoUrl(val *string) {
	if err := j.validateSetSsoUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssoUrl",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetSubjectNameIdFormat(val *string) {
	if err := j.validateSetSubjectNameIdFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectNameIdFormat",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetSubjectNameIdTemplate(val *string) {
	if err := j.validateSetSubjectNameIdTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectNameIdTemplate",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetUserNameTemplate(val *string) {
	if err := j.validateSetUserNameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplate",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetUserNameTemplatePushStatus(val *string) {
	if err := j.validateSetUserNameTemplatePushStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplatePushStatus",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetUserNameTemplateSuffix(val *string) {
	if err := j.validateSetUserNameTemplateSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplateSuffix",
		val,
	)
}

func (j *jsiiProxy_SamlApp)SetUserNameTemplateType(val *string) {
	if err := j.validateSetUserNameTemplateTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameTemplateType",
		val,
	)
}

// Generates CDKTF code for importing a SamlApp resource upon running "cdktf plan <stack-name>".
func SamlApp_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSamlApp_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.samlApp.SamlApp",
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
func SamlApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSamlApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.samlApp.SamlApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SamlApp_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSamlApp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.samlApp.SamlApp",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SamlApp_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSamlApp_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.samlApp.SamlApp",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SamlApp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.samlApp.SamlApp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SamlApp) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SamlApp) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SamlApp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SamlApp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SamlApp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SamlApp) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SamlApp) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SamlApp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SamlApp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SamlApp) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SamlApp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SamlApp) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlApp) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SamlApp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SamlApp) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SamlApp) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SamlApp) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SamlApp) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SamlApp) PutAttributeStatements(value interface{}) {
	if err := s.validatePutAttributeStatementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAttributeStatements",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SamlApp) PutTimeouts(value *SamlAppTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SamlApp) PutUsers(value interface{}) {
	if err := s.validatePutUsersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putUsers",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SamlApp) ResetAccessibilityErrorRedirectUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetAccessibilityErrorRedirectUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetAccessibilityLoginRedirectUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetAccessibilityLoginRedirectUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetAccessibilitySelfService() {
	_jsii_.InvokeVoid(
		s,
		"resetAccessibilitySelfService",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetAcsEndpoints() {
	_jsii_.InvokeVoid(
		s,
		"resetAcsEndpoints",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetAdminNote() {
	_jsii_.InvokeVoid(
		s,
		"resetAdminNote",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetAppLinksJson() {
	_jsii_.InvokeVoid(
		s,
		"resetAppLinksJson",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetAppSettingsJson() {
	_jsii_.InvokeVoid(
		s,
		"resetAppSettingsJson",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetAssertionSigned() {
	_jsii_.InvokeVoid(
		s,
		"resetAssertionSigned",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetAttributeStatements() {
	_jsii_.InvokeVoid(
		s,
		"resetAttributeStatements",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetAudience() {
	_jsii_.InvokeVoid(
		s,
		"resetAudience",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetAuthenticationPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetAuthenticationPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetAuthnContextClassRef() {
	_jsii_.InvokeVoid(
		s,
		"resetAuthnContextClassRef",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetAutoSubmitToolbar() {
	_jsii_.InvokeVoid(
		s,
		"resetAutoSubmitToolbar",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetDefaultRelayState() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultRelayState",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetDestination() {
	_jsii_.InvokeVoid(
		s,
		"resetDestination",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetDigestAlgorithm() {
	_jsii_.InvokeVoid(
		s,
		"resetDigestAlgorithm",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetEnduserNote() {
	_jsii_.InvokeVoid(
		s,
		"resetEnduserNote",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetGroups() {
	_jsii_.InvokeVoid(
		s,
		"resetGroups",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetHideIos() {
	_jsii_.InvokeVoid(
		s,
		"resetHideIos",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetHideWeb() {
	_jsii_.InvokeVoid(
		s,
		"resetHideWeb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetHonorForceAuthn() {
	_jsii_.InvokeVoid(
		s,
		"resetHonorForceAuthn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetIdpIssuer() {
	_jsii_.InvokeVoid(
		s,
		"resetIdpIssuer",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetImplicitAssignment() {
	_jsii_.InvokeVoid(
		s,
		"resetImplicitAssignment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetInlineHookId() {
	_jsii_.InvokeVoid(
		s,
		"resetInlineHookId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetKeyName() {
	_jsii_.InvokeVoid(
		s,
		"resetKeyName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetKeyYearsValid() {
	_jsii_.InvokeVoid(
		s,
		"resetKeyYearsValid",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetLogo() {
	_jsii_.InvokeVoid(
		s,
		"resetLogo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetPreconfiguredApp() {
	_jsii_.InvokeVoid(
		s,
		"resetPreconfiguredApp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetRecipient() {
	_jsii_.InvokeVoid(
		s,
		"resetRecipient",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetRequestCompressed() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestCompressed",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetResponseSigned() {
	_jsii_.InvokeVoid(
		s,
		"resetResponseSigned",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetSamlVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetSamlVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetSignatureAlgorithm() {
	_jsii_.InvokeVoid(
		s,
		"resetSignatureAlgorithm",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetSingleLogoutCertificate() {
	_jsii_.InvokeVoid(
		s,
		"resetSingleLogoutCertificate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetSingleLogoutIssuer() {
	_jsii_.InvokeVoid(
		s,
		"resetSingleLogoutIssuer",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetSingleLogoutUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetSingleLogoutUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetSkipGroups() {
	_jsii_.InvokeVoid(
		s,
		"resetSkipGroups",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetSkipUsers() {
	_jsii_.InvokeVoid(
		s,
		"resetSkipUsers",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetSpIssuer() {
	_jsii_.InvokeVoid(
		s,
		"resetSpIssuer",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetSsoUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetSsoUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetSubjectNameIdFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetSubjectNameIdFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetSubjectNameIdTemplate() {
	_jsii_.InvokeVoid(
		s,
		"resetSubjectNameIdTemplate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetUserNameTemplate() {
	_jsii_.InvokeVoid(
		s,
		"resetUserNameTemplate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetUserNameTemplatePushStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetUserNameTemplatePushStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetUserNameTemplateSuffix() {
	_jsii_.InvokeVoid(
		s,
		"resetUserNameTemplateSuffix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetUserNameTemplateType() {
	_jsii_.InvokeVoid(
		s,
		"resetUserNameTemplateType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) ResetUsers() {
	_jsii_.InvokeVoid(
		s,
		"resetUsers",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlApp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlApp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlApp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


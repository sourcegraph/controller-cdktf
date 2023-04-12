package socialidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/socialidp/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/okta/r/social_idp okta_social_idp}.
type SocialIdp interface {
	cdktf.TerraformResource
	AccountLinkAction() *string
	SetAccountLinkAction(val *string)
	AccountLinkActionInput() *string
	AccountLinkGroupInclude() *[]*string
	SetAccountLinkGroupInclude(val *[]*string)
	AccountLinkGroupIncludeInput() *[]*string
	AppleKid() *string
	SetAppleKid(val *string)
	AppleKidInput() *string
	ApplePrivateKey() *string
	SetApplePrivateKey(val *string)
	ApplePrivateKeyInput() *string
	AppleTeamId() *string
	SetAppleTeamId(val *string)
	AppleTeamIdInput() *string
	AuthorizationBinding() *string
	AuthorizationUrl() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeprovisionedAction() *string
	SetDeprovisionedAction(val *string)
	DeprovisionedActionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupsAction() *string
	SetGroupsAction(val *string)
	GroupsActionInput() *string
	GroupsAssignment() *[]*string
	SetGroupsAssignment(val *[]*string)
	GroupsAssignmentInput() *[]*string
	GroupsAttribute() *string
	SetGroupsAttribute(val *string)
	GroupsAttributeInput() *string
	GroupsFilter() *[]*string
	SetGroupsFilter(val *[]*string)
	GroupsFilterInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IssuerMode() *string
	SetIssuerMode(val *string)
	IssuerModeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MatchAttribute() *string
	SetMatchAttribute(val *string)
	MatchAttributeInput() *string
	MatchType() *string
	SetMatchType(val *string)
	MatchTypeInput() *string
	MaxClockSkew() *float64
	SetMaxClockSkew(val *float64)
	MaxClockSkewInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ProfileMaster() interface{}
	SetProfileMaster(val interface{})
	ProfileMasterInput() interface{}
	ProtocolType() *string
	SetProtocolType(val *string)
	ProtocolTypeInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProvisioningAction() *string
	SetProvisioningAction(val *string)
	ProvisioningActionInput() *string
	// Experimental.
	RawOverrides() interface{}
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	SubjectMatchAttribute() *string
	SetSubjectMatchAttribute(val *string)
	SubjectMatchAttributeInput() *string
	SubjectMatchType() *string
	SetSubjectMatchType(val *string)
	SubjectMatchTypeInput() *string
	SuspendedAction() *string
	SetSuspendedAction(val *string)
	SuspendedActionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TokenBinding() *string
	TokenUrl() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UsernameTemplate() *string
	SetUsernameTemplate(val *string)
	UsernameTemplateInput() *string
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
	ResetAccountLinkAction()
	ResetAccountLinkGroupInclude()
	ResetAppleKid()
	ResetApplePrivateKey()
	ResetAppleTeamId()
	ResetClientId()
	ResetClientSecret()
	ResetDeprovisionedAction()
	ResetGroupsAction()
	ResetGroupsAssignment()
	ResetGroupsAttribute()
	ResetGroupsFilter()
	ResetId()
	ResetIssuerMode()
	ResetMatchAttribute()
	ResetMatchType()
	ResetMaxClockSkew()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProfileMaster()
	ResetProtocolType()
	ResetProvisioningAction()
	ResetStatus()
	ResetSubjectMatchAttribute()
	ResetSubjectMatchType()
	ResetSuspendedAction()
	ResetUsernameTemplate()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SocialIdp
type jsiiProxy_SocialIdp struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SocialIdp) AccountLinkAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountLinkAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) AccountLinkActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountLinkActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) AccountLinkGroupInclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountLinkGroupInclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) AccountLinkGroupIncludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountLinkGroupIncludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) AppleKid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appleKid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) AppleKidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appleKidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) ApplePrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applePrivateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) ApplePrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applePrivateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) AppleTeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appleTeamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) AppleTeamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appleTeamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) AuthorizationBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) AuthorizationUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) DeprovisionedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deprovisionedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) DeprovisionedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deprovisionedActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) GroupsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) GroupsActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) GroupsAssignment() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsAssignment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) GroupsAssignmentInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsAssignmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) GroupsAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) GroupsAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) GroupsFilter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) GroupsFilterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) IssuerMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) IssuerModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) MatchAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) MatchAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) MatchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) MatchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) MaxClockSkew() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxClockSkew",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) MaxClockSkewInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxClockSkewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) ProfileMaster() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profileMaster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) ProfileMasterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profileMasterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) ProtocolType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) ProtocolTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) ProvisioningAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) ProvisioningActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) SubjectMatchAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) SubjectMatchAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) SubjectMatchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) SubjectMatchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) SuspendedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspendedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) SuspendedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspendedActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) TokenBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) TokenUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) UsernameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SocialIdp) UsernameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplateInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/okta/r/social_idp okta_social_idp} Resource.
func NewSocialIdp(scope constructs.Construct, id *string, config *SocialIdpConfig) SocialIdp {
	_init_.Initialize()

	if err := validateNewSocialIdpParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SocialIdp{}

	_jsii_.Create(
		"@cdktf/provider-okta.socialIdp.SocialIdp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/okta/r/social_idp okta_social_idp} Resource.
func NewSocialIdp_Override(s SocialIdp, scope constructs.Construct, id *string, config *SocialIdpConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.socialIdp.SocialIdp",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SocialIdp)SetAccountLinkAction(val *string) {
	if err := j.validateSetAccountLinkActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountLinkAction",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetAccountLinkGroupInclude(val *[]*string) {
	if err := j.validateSetAccountLinkGroupIncludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountLinkGroupInclude",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetAppleKid(val *string) {
	if err := j.validateSetAppleKidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appleKid",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetApplePrivateKey(val *string) {
	if err := j.validateSetApplePrivateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applePrivateKey",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetAppleTeamId(val *string) {
	if err := j.validateSetAppleTeamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appleTeamId",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetDeprovisionedAction(val *string) {
	if err := j.validateSetDeprovisionedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deprovisionedAction",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetGroupsAction(val *string) {
	if err := j.validateSetGroupsActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsAction",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetGroupsAssignment(val *[]*string) {
	if err := j.validateSetGroupsAssignmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsAssignment",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetGroupsAttribute(val *string) {
	if err := j.validateSetGroupsAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsAttribute",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetGroupsFilter(val *[]*string) {
	if err := j.validateSetGroupsFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsFilter",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetIssuerMode(val *string) {
	if err := j.validateSetIssuerModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerMode",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetMatchAttribute(val *string) {
	if err := j.validateSetMatchAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchAttribute",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetMatchType(val *string) {
	if err := j.validateSetMatchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchType",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetMaxClockSkew(val *float64) {
	if err := j.validateSetMaxClockSkewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxClockSkew",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetProfileMaster(val interface{}) {
	if err := j.validateSetProfileMasterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profileMaster",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetProtocolType(val *string) {
	if err := j.validateSetProtocolTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocolType",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetProvisioningAction(val *string) {
	if err := j.validateSetProvisioningActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningAction",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetSubjectMatchAttribute(val *string) {
	if err := j.validateSetSubjectMatchAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectMatchAttribute",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetSubjectMatchType(val *string) {
	if err := j.validateSetSubjectMatchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectMatchType",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetSuspendedAction(val *string) {
	if err := j.validateSetSuspendedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspendedAction",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_SocialIdp)SetUsernameTemplate(val *string) {
	if err := j.validateSetUsernameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameTemplate",
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
func SocialIdp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSocialIdp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.socialIdp.SocialIdp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SocialIdp_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSocialIdp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.socialIdp.SocialIdp",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SocialIdp_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSocialIdp_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.socialIdp.SocialIdp",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SocialIdp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.socialIdp.SocialIdp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SocialIdp) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SocialIdp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SocialIdp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SocialIdp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SocialIdp) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SocialIdp) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SocialIdp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SocialIdp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SocialIdp) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SocialIdp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SocialIdp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SocialIdp) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SocialIdp) ResetAccountLinkAction() {
	_jsii_.InvokeVoid(
		s,
		"resetAccountLinkAction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetAccountLinkGroupInclude() {
	_jsii_.InvokeVoid(
		s,
		"resetAccountLinkGroupInclude",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetAppleKid() {
	_jsii_.InvokeVoid(
		s,
		"resetAppleKid",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetApplePrivateKey() {
	_jsii_.InvokeVoid(
		s,
		"resetApplePrivateKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetAppleTeamId() {
	_jsii_.InvokeVoid(
		s,
		"resetAppleTeamId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetClientId() {
	_jsii_.InvokeVoid(
		s,
		"resetClientId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetClientSecret() {
	_jsii_.InvokeVoid(
		s,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetDeprovisionedAction() {
	_jsii_.InvokeVoid(
		s,
		"resetDeprovisionedAction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetGroupsAction() {
	_jsii_.InvokeVoid(
		s,
		"resetGroupsAction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetGroupsAssignment() {
	_jsii_.InvokeVoid(
		s,
		"resetGroupsAssignment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetGroupsAttribute() {
	_jsii_.InvokeVoid(
		s,
		"resetGroupsAttribute",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetGroupsFilter() {
	_jsii_.InvokeVoid(
		s,
		"resetGroupsFilter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetIssuerMode() {
	_jsii_.InvokeVoid(
		s,
		"resetIssuerMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetMatchAttribute() {
	_jsii_.InvokeVoid(
		s,
		"resetMatchAttribute",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetMatchType() {
	_jsii_.InvokeVoid(
		s,
		"resetMatchType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetMaxClockSkew() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxClockSkew",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetProfileMaster() {
	_jsii_.InvokeVoid(
		s,
		"resetProfileMaster",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetProtocolType() {
	_jsii_.InvokeVoid(
		s,
		"resetProtocolType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetProvisioningAction() {
	_jsii_.InvokeVoid(
		s,
		"resetProvisioningAction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetSubjectMatchAttribute() {
	_jsii_.InvokeVoid(
		s,
		"resetSubjectMatchAttribute",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetSubjectMatchType() {
	_jsii_.InvokeVoid(
		s,
		"resetSubjectMatchType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetSuspendedAction() {
	_jsii_.InvokeVoid(
		s,
		"resetSuspendedAction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) ResetUsernameTemplate() {
	_jsii_.InvokeVoid(
		s,
		"resetUsernameTemplate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SocialIdp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SocialIdp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SocialIdp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SocialIdp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


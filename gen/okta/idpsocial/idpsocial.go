package idpsocial

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/idpsocial/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/idp_social okta_idp_social}.
type IdpSocial interface {
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
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
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

// The jsii proxy struct for IdpSocial
type jsiiProxy_IdpSocial struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IdpSocial) AccountLinkAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountLinkAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) AccountLinkActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountLinkActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) AccountLinkGroupInclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountLinkGroupInclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) AccountLinkGroupIncludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountLinkGroupIncludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) AppleKid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appleKid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) AppleKidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appleKidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) ApplePrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applePrivateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) ApplePrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applePrivateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) AppleTeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appleTeamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) AppleTeamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appleTeamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) AuthorizationBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) AuthorizationUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) DeprovisionedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deprovisionedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) DeprovisionedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deprovisionedActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) GroupsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) GroupsActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) GroupsAssignment() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsAssignment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) GroupsAssignmentInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsAssignmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) GroupsAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) GroupsAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) GroupsFilter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) GroupsFilterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) IssuerMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) IssuerModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) MatchAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) MatchAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) MatchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) MatchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) MaxClockSkew() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxClockSkew",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) MaxClockSkewInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxClockSkewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) ProfileMaster() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profileMaster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) ProfileMasterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profileMasterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) ProtocolType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) ProtocolTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) ProvisioningAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) ProvisioningActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) SubjectMatchAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) SubjectMatchAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) SubjectMatchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) SubjectMatchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) SuspendedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspendedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) SuspendedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspendedActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) TokenBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) TokenUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) UsernameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSocial) UsernameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplateInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/idp_social okta_idp_social} Resource.
func NewIdpSocial(scope constructs.Construct, id *string, config *IdpSocialConfig) IdpSocial {
	_init_.Initialize()

	if err := validateNewIdpSocialParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IdpSocial{}

	_jsii_.Create(
		"@cdktf/provider-okta.idpSocial.IdpSocial",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/idp_social okta_idp_social} Resource.
func NewIdpSocial_Override(i IdpSocial, scope constructs.Construct, id *string, config *IdpSocialConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.idpSocial.IdpSocial",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IdpSocial)SetAccountLinkAction(val *string) {
	if err := j.validateSetAccountLinkActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountLinkAction",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetAccountLinkGroupInclude(val *[]*string) {
	if err := j.validateSetAccountLinkGroupIncludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountLinkGroupInclude",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetAppleKid(val *string) {
	if err := j.validateSetAppleKidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appleKid",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetApplePrivateKey(val *string) {
	if err := j.validateSetApplePrivateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applePrivateKey",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetAppleTeamId(val *string) {
	if err := j.validateSetAppleTeamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appleTeamId",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetDeprovisionedAction(val *string) {
	if err := j.validateSetDeprovisionedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deprovisionedAction",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetGroupsAction(val *string) {
	if err := j.validateSetGroupsActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsAction",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetGroupsAssignment(val *[]*string) {
	if err := j.validateSetGroupsAssignmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsAssignment",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetGroupsAttribute(val *string) {
	if err := j.validateSetGroupsAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsAttribute",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetGroupsFilter(val *[]*string) {
	if err := j.validateSetGroupsFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsFilter",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetIssuerMode(val *string) {
	if err := j.validateSetIssuerModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerMode",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetMatchAttribute(val *string) {
	if err := j.validateSetMatchAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchAttribute",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetMatchType(val *string) {
	if err := j.validateSetMatchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchType",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetMaxClockSkew(val *float64) {
	if err := j.validateSetMaxClockSkewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxClockSkew",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetProfileMaster(val interface{}) {
	if err := j.validateSetProfileMasterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profileMaster",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetProtocolType(val *string) {
	if err := j.validateSetProtocolTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocolType",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetProvisioningAction(val *string) {
	if err := j.validateSetProvisioningActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningAction",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetSubjectMatchAttribute(val *string) {
	if err := j.validateSetSubjectMatchAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectMatchAttribute",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetSubjectMatchType(val *string) {
	if err := j.validateSetSubjectMatchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectMatchType",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetSuspendedAction(val *string) {
	if err := j.validateSetSuspendedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspendedAction",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_IdpSocial)SetUsernameTemplate(val *string) {
	if err := j.validateSetUsernameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameTemplate",
		val,
	)
}

// Generates CDKTF code for importing a IdpSocial resource upon running "cdktf plan <stack-name>".
func IdpSocial_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIdpSocial_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.idpSocial.IdpSocial",
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
func IdpSocial_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdpSocial_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.idpSocial.IdpSocial",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdpSocial_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdpSocial_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.idpSocial.IdpSocial",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdpSocial_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdpSocial_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.idpSocial.IdpSocial",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IdpSocial_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.idpSocial.IdpSocial",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IdpSocial) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IdpSocial) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IdpSocial) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpSocial) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpSocial) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpSocial) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpSocial) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpSocial) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpSocial) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpSocial) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpSocial) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpSocial) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpSocial) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IdpSocial) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpSocial) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IdpSocial) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IdpSocial) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IdpSocial) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IdpSocial) ResetAccountLinkAction() {
	_jsii_.InvokeVoid(
		i,
		"resetAccountLinkAction",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetAccountLinkGroupInclude() {
	_jsii_.InvokeVoid(
		i,
		"resetAccountLinkGroupInclude",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetAppleKid() {
	_jsii_.InvokeVoid(
		i,
		"resetAppleKid",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetApplePrivateKey() {
	_jsii_.InvokeVoid(
		i,
		"resetApplePrivateKey",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetAppleTeamId() {
	_jsii_.InvokeVoid(
		i,
		"resetAppleTeamId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetClientId() {
	_jsii_.InvokeVoid(
		i,
		"resetClientId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetClientSecret() {
	_jsii_.InvokeVoid(
		i,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetDeprovisionedAction() {
	_jsii_.InvokeVoid(
		i,
		"resetDeprovisionedAction",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetGroupsAction() {
	_jsii_.InvokeVoid(
		i,
		"resetGroupsAction",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetGroupsAssignment() {
	_jsii_.InvokeVoid(
		i,
		"resetGroupsAssignment",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetGroupsAttribute() {
	_jsii_.InvokeVoid(
		i,
		"resetGroupsAttribute",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetGroupsFilter() {
	_jsii_.InvokeVoid(
		i,
		"resetGroupsFilter",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetIssuerMode() {
	_jsii_.InvokeVoid(
		i,
		"resetIssuerMode",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetMatchAttribute() {
	_jsii_.InvokeVoid(
		i,
		"resetMatchAttribute",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetMatchType() {
	_jsii_.InvokeVoid(
		i,
		"resetMatchType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetMaxClockSkew() {
	_jsii_.InvokeVoid(
		i,
		"resetMaxClockSkew",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetProfileMaster() {
	_jsii_.InvokeVoid(
		i,
		"resetProfileMaster",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetProtocolType() {
	_jsii_.InvokeVoid(
		i,
		"resetProtocolType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetProvisioningAction() {
	_jsii_.InvokeVoid(
		i,
		"resetProvisioningAction",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetStatus() {
	_jsii_.InvokeVoid(
		i,
		"resetStatus",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetSubjectMatchAttribute() {
	_jsii_.InvokeVoid(
		i,
		"resetSubjectMatchAttribute",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetSubjectMatchType() {
	_jsii_.InvokeVoid(
		i,
		"resetSubjectMatchType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetSuspendedAction() {
	_jsii_.InvokeVoid(
		i,
		"resetSuspendedAction",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) ResetUsernameTemplate() {
	_jsii_.InvokeVoid(
		i,
		"resetUsernameTemplate",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSocial) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpSocial) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpSocial) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpSocial) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


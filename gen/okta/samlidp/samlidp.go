package samlidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/samlidp/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/saml_idp okta_saml_idp}.
type SamlIdp interface {
	cdktf.TerraformResource
	AccountLinkAction() *string
	SetAccountLinkAction(val *string)
	AccountLinkActionInput() *string
	AccountLinkGroupInclude() *[]*string
	SetAccountLinkGroupInclude(val *[]*string)
	AccountLinkGroupIncludeInput() *[]*string
	AcsBinding() *string
	SetAcsBinding(val *string)
	AcsBindingInput() *string
	AcsType() *string
	SetAcsType(val *string)
	AcsTypeInput() *string
	Audience() *string
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
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
	IssuerMode() *string
	SetIssuerMode(val *string)
	IssuerModeInput() *string
	Kid() *string
	SetKid(val *string)
	KidInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxClockSkew() *float64
	SetMaxClockSkew(val *float64)
	MaxClockSkewInput() *float64
	Name() *string
	SetName(val *string)
	NameFormat() *string
	SetNameFormat(val *string)
	NameFormatInput() *string
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ProfileMaster() interface{}
	SetProfileMaster(val interface{})
	ProfileMasterInput() interface{}
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
	RequestSignatureAlgorithm() *string
	SetRequestSignatureAlgorithm(val *string)
	RequestSignatureAlgorithmInput() *string
	RequestSignatureScope() *string
	SetRequestSignatureScope(val *string)
	RequestSignatureScopeInput() *string
	ResponseSignatureAlgorithm() *string
	SetResponseSignatureAlgorithm(val *string)
	ResponseSignatureAlgorithmInput() *string
	ResponseSignatureScope() *string
	SetResponseSignatureScope(val *string)
	ResponseSignatureScopeInput() *string
	SsoBinding() *string
	SetSsoBinding(val *string)
	SsoBindingInput() *string
	SsoDestination() *string
	SetSsoDestination(val *string)
	SsoDestinationInput() *string
	SsoUrl() *string
	SetSsoUrl(val *string)
	SsoUrlInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	SubjectFilter() *string
	SetSubjectFilter(val *string)
	SubjectFilterInput() *string
	SubjectFormat() *[]*string
	SetSubjectFormat(val *[]*string)
	SubjectFormatInput() *[]*string
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
	Type() *string
	UsernameTemplate() *string
	SetUsernameTemplate(val *string)
	UsernameTemplateInput() *string
	UserTypeId() *string
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
	ResetAcsBinding()
	ResetAcsType()
	ResetDeprovisionedAction()
	ResetGroupsAction()
	ResetGroupsAssignment()
	ResetGroupsAttribute()
	ResetGroupsFilter()
	ResetId()
	ResetIssuerMode()
	ResetMaxClockSkew()
	ResetNameFormat()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProfileMaster()
	ResetProvisioningAction()
	ResetRequestSignatureAlgorithm()
	ResetRequestSignatureScope()
	ResetResponseSignatureAlgorithm()
	ResetResponseSignatureScope()
	ResetSsoBinding()
	ResetSsoDestination()
	ResetStatus()
	ResetSubjectFilter()
	ResetSubjectFormat()
	ResetSubjectMatchAttribute()
	ResetSubjectMatchType()
	ResetSuspendedAction()
	ResetUsernameTemplate()
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

// The jsii proxy struct for SamlIdp
type jsiiProxy_SamlIdp struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SamlIdp) AccountLinkAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountLinkAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) AccountLinkActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountLinkActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) AccountLinkGroupInclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountLinkGroupInclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) AccountLinkGroupIncludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountLinkGroupIncludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) AcsBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acsBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) AcsBindingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acsBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) AcsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) AcsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acsTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) Audience() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) DeprovisionedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deprovisionedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) DeprovisionedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deprovisionedActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) GroupsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) GroupsActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) GroupsAssignment() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsAssignment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) GroupsAssignmentInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsAssignmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) GroupsAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) GroupsAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) GroupsFilter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) GroupsFilterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) IssuerMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) IssuerModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) Kid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) KidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) MaxClockSkew() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxClockSkew",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) MaxClockSkewInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxClockSkewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) NameFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) NameFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) ProfileMaster() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profileMaster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) ProfileMasterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profileMasterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) ProvisioningAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) ProvisioningActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) RequestSignatureAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestSignatureAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) RequestSignatureAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestSignatureAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) RequestSignatureScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestSignatureScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) RequestSignatureScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestSignatureScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) ResponseSignatureAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseSignatureAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) ResponseSignatureAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseSignatureAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) ResponseSignatureScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseSignatureScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) ResponseSignatureScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseSignatureScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) SsoBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) SsoBindingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) SsoDestination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) SsoDestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) SsoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) SsoUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) SubjectFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) SubjectFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) SubjectFormat() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subjectFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) SubjectFormatInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subjectFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) SubjectMatchAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) SubjectMatchAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) SubjectMatchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) SubjectMatchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) SuspendedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspendedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) SuspendedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspendedActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) UsernameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) UsernameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdp) UserTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTypeId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/saml_idp okta_saml_idp} Resource.
func NewSamlIdp(scope constructs.Construct, id *string, config *SamlIdpConfig) SamlIdp {
	_init_.Initialize()

	if err := validateNewSamlIdpParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SamlIdp{}

	_jsii_.Create(
		"@cdktf/provider-okta.samlIdp.SamlIdp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/saml_idp okta_saml_idp} Resource.
func NewSamlIdp_Override(s SamlIdp, scope constructs.Construct, id *string, config *SamlIdpConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.samlIdp.SamlIdp",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SamlIdp)SetAccountLinkAction(val *string) {
	if err := j.validateSetAccountLinkActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountLinkAction",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetAccountLinkGroupInclude(val *[]*string) {
	if err := j.validateSetAccountLinkGroupIncludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountLinkGroupInclude",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetAcsBinding(val *string) {
	if err := j.validateSetAcsBindingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acsBinding",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetAcsType(val *string) {
	if err := j.validateSetAcsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acsType",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetDeprovisionedAction(val *string) {
	if err := j.validateSetDeprovisionedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deprovisionedAction",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetGroupsAction(val *string) {
	if err := j.validateSetGroupsActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsAction",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetGroupsAssignment(val *[]*string) {
	if err := j.validateSetGroupsAssignmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsAssignment",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetGroupsAttribute(val *string) {
	if err := j.validateSetGroupsAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsAttribute",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetGroupsFilter(val *[]*string) {
	if err := j.validateSetGroupsFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsFilter",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetIssuerMode(val *string) {
	if err := j.validateSetIssuerModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerMode",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetKid(val *string) {
	if err := j.validateSetKidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kid",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetMaxClockSkew(val *float64) {
	if err := j.validateSetMaxClockSkewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxClockSkew",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetNameFormat(val *string) {
	if err := j.validateSetNameFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameFormat",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetProfileMaster(val interface{}) {
	if err := j.validateSetProfileMasterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profileMaster",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetProvisioningAction(val *string) {
	if err := j.validateSetProvisioningActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningAction",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetRequestSignatureAlgorithm(val *string) {
	if err := j.validateSetRequestSignatureAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestSignatureAlgorithm",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetRequestSignatureScope(val *string) {
	if err := j.validateSetRequestSignatureScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestSignatureScope",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetResponseSignatureAlgorithm(val *string) {
	if err := j.validateSetResponseSignatureAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseSignatureAlgorithm",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetResponseSignatureScope(val *string) {
	if err := j.validateSetResponseSignatureScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseSignatureScope",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetSsoBinding(val *string) {
	if err := j.validateSetSsoBindingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssoBinding",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetSsoDestination(val *string) {
	if err := j.validateSetSsoDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssoDestination",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetSsoUrl(val *string) {
	if err := j.validateSetSsoUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssoUrl",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetSubjectFilter(val *string) {
	if err := j.validateSetSubjectFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectFilter",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetSubjectFormat(val *[]*string) {
	if err := j.validateSetSubjectFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectFormat",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetSubjectMatchAttribute(val *string) {
	if err := j.validateSetSubjectMatchAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectMatchAttribute",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetSubjectMatchType(val *string) {
	if err := j.validateSetSubjectMatchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectMatchType",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetSuspendedAction(val *string) {
	if err := j.validateSetSuspendedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspendedAction",
		val,
	)
}

func (j *jsiiProxy_SamlIdp)SetUsernameTemplate(val *string) {
	if err := j.validateSetUsernameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameTemplate",
		val,
	)
}

// Generates CDKTF code for importing a SamlIdp resource upon running "cdktf plan <stack-name>".
func SamlIdp_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSamlIdp_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.samlIdp.SamlIdp",
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
func SamlIdp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSamlIdp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.samlIdp.SamlIdp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SamlIdp_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSamlIdp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.samlIdp.SamlIdp",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SamlIdp_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSamlIdp_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.samlIdp.SamlIdp",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SamlIdp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.samlIdp.SamlIdp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SamlIdp) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SamlIdp) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SamlIdp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SamlIdp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SamlIdp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SamlIdp) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SamlIdp) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SamlIdp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SamlIdp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SamlIdp) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SamlIdp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SamlIdp) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlIdp) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SamlIdp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SamlIdp) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SamlIdp) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SamlIdp) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SamlIdp) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SamlIdp) ResetAccountLinkAction() {
	_jsii_.InvokeVoid(
		s,
		"resetAccountLinkAction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetAccountLinkGroupInclude() {
	_jsii_.InvokeVoid(
		s,
		"resetAccountLinkGroupInclude",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetAcsBinding() {
	_jsii_.InvokeVoid(
		s,
		"resetAcsBinding",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetAcsType() {
	_jsii_.InvokeVoid(
		s,
		"resetAcsType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetDeprovisionedAction() {
	_jsii_.InvokeVoid(
		s,
		"resetDeprovisionedAction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetGroupsAction() {
	_jsii_.InvokeVoid(
		s,
		"resetGroupsAction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetGroupsAssignment() {
	_jsii_.InvokeVoid(
		s,
		"resetGroupsAssignment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetGroupsAttribute() {
	_jsii_.InvokeVoid(
		s,
		"resetGroupsAttribute",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetGroupsFilter() {
	_jsii_.InvokeVoid(
		s,
		"resetGroupsFilter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetIssuerMode() {
	_jsii_.InvokeVoid(
		s,
		"resetIssuerMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetMaxClockSkew() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxClockSkew",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetNameFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetNameFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetProfileMaster() {
	_jsii_.InvokeVoid(
		s,
		"resetProfileMaster",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetProvisioningAction() {
	_jsii_.InvokeVoid(
		s,
		"resetProvisioningAction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetRequestSignatureAlgorithm() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestSignatureAlgorithm",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetRequestSignatureScope() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestSignatureScope",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetResponseSignatureAlgorithm() {
	_jsii_.InvokeVoid(
		s,
		"resetResponseSignatureAlgorithm",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetResponseSignatureScope() {
	_jsii_.InvokeVoid(
		s,
		"resetResponseSignatureScope",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetSsoBinding() {
	_jsii_.InvokeVoid(
		s,
		"resetSsoBinding",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetSsoDestination() {
	_jsii_.InvokeVoid(
		s,
		"resetSsoDestination",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetSubjectFilter() {
	_jsii_.InvokeVoid(
		s,
		"resetSubjectFilter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetSubjectFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetSubjectFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetSubjectMatchAttribute() {
	_jsii_.InvokeVoid(
		s,
		"resetSubjectMatchAttribute",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetSubjectMatchType() {
	_jsii_.InvokeVoid(
		s,
		"resetSubjectMatchType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetSuspendedAction() {
	_jsii_.InvokeVoid(
		s,
		"resetSuspendedAction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) ResetUsernameTemplate() {
	_jsii_.InvokeVoid(
		s,
		"resetUsernameTemplate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIdp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlIdp) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlIdp) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlIdp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlIdp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlIdp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


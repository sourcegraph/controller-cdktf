package idpsaml

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/idpsaml/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/idp_saml okta_idp_saml}.
type IdpSaml interface {
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
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for IdpSaml
type jsiiProxy_IdpSaml struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IdpSaml) AccountLinkAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountLinkAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) AccountLinkActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountLinkActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) AccountLinkGroupInclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountLinkGroupInclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) AccountLinkGroupIncludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountLinkGroupIncludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) AcsBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acsBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) AcsBindingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acsBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) AcsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) AcsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acsTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) Audience() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) DeprovisionedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deprovisionedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) DeprovisionedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deprovisionedActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) GroupsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) GroupsActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) GroupsAssignment() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsAssignment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) GroupsAssignmentInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsAssignmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) GroupsAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) GroupsAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) GroupsFilter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) GroupsFilterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) IssuerMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) IssuerModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) Kid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) KidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) MaxClockSkew() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxClockSkew",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) MaxClockSkewInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxClockSkewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) NameFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) NameFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) ProfileMaster() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profileMaster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) ProfileMasterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profileMasterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) ProvisioningAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) ProvisioningActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) RequestSignatureAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestSignatureAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) RequestSignatureAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestSignatureAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) RequestSignatureScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestSignatureScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) RequestSignatureScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestSignatureScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) ResponseSignatureAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseSignatureAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) ResponseSignatureAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseSignatureAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) ResponseSignatureScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseSignatureScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) ResponseSignatureScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseSignatureScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) SsoBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) SsoBindingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) SsoDestination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) SsoDestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) SsoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) SsoUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) SubjectFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) SubjectFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) SubjectFormat() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subjectFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) SubjectFormatInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subjectFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) SubjectMatchAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) SubjectMatchAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) SubjectMatchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) SubjectMatchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) SuspendedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspendedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) SuspendedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspendedActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) UsernameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) UsernameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdpSaml) UserTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTypeId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/idp_saml okta_idp_saml} Resource.
func NewIdpSaml(scope constructs.Construct, id *string, config *IdpSamlConfig) IdpSaml {
	_init_.Initialize()

	if err := validateNewIdpSamlParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IdpSaml{}

	_jsii_.Create(
		"okta.idpSaml.IdpSaml",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/idp_saml okta_idp_saml} Resource.
func NewIdpSaml_Override(i IdpSaml, scope constructs.Construct, id *string, config *IdpSamlConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"okta.idpSaml.IdpSaml",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IdpSaml)SetAccountLinkAction(val *string) {
	if err := j.validateSetAccountLinkActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountLinkAction",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetAccountLinkGroupInclude(val *[]*string) {
	if err := j.validateSetAccountLinkGroupIncludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountLinkGroupInclude",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetAcsBinding(val *string) {
	if err := j.validateSetAcsBindingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acsBinding",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetAcsType(val *string) {
	if err := j.validateSetAcsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acsType",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetDeprovisionedAction(val *string) {
	if err := j.validateSetDeprovisionedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deprovisionedAction",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetGroupsAction(val *string) {
	if err := j.validateSetGroupsActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsAction",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetGroupsAssignment(val *[]*string) {
	if err := j.validateSetGroupsAssignmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsAssignment",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetGroupsAttribute(val *string) {
	if err := j.validateSetGroupsAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsAttribute",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetGroupsFilter(val *[]*string) {
	if err := j.validateSetGroupsFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsFilter",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetIssuerMode(val *string) {
	if err := j.validateSetIssuerModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerMode",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetKid(val *string) {
	if err := j.validateSetKidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kid",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetMaxClockSkew(val *float64) {
	if err := j.validateSetMaxClockSkewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxClockSkew",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetNameFormat(val *string) {
	if err := j.validateSetNameFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameFormat",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetProfileMaster(val interface{}) {
	if err := j.validateSetProfileMasterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profileMaster",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetProvisioningAction(val *string) {
	if err := j.validateSetProvisioningActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningAction",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetRequestSignatureAlgorithm(val *string) {
	if err := j.validateSetRequestSignatureAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestSignatureAlgorithm",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetRequestSignatureScope(val *string) {
	if err := j.validateSetRequestSignatureScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestSignatureScope",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetResponseSignatureAlgorithm(val *string) {
	if err := j.validateSetResponseSignatureAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseSignatureAlgorithm",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetResponseSignatureScope(val *string) {
	if err := j.validateSetResponseSignatureScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseSignatureScope",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetSsoBinding(val *string) {
	if err := j.validateSetSsoBindingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssoBinding",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetSsoDestination(val *string) {
	if err := j.validateSetSsoDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssoDestination",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetSsoUrl(val *string) {
	if err := j.validateSetSsoUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssoUrl",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetSubjectFilter(val *string) {
	if err := j.validateSetSubjectFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectFilter",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetSubjectFormat(val *[]*string) {
	if err := j.validateSetSubjectFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectFormat",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetSubjectMatchAttribute(val *string) {
	if err := j.validateSetSubjectMatchAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectMatchAttribute",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetSubjectMatchType(val *string) {
	if err := j.validateSetSubjectMatchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectMatchType",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetSuspendedAction(val *string) {
	if err := j.validateSetSuspendedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspendedAction",
		val,
	)
}

func (j *jsiiProxy_IdpSaml)SetUsernameTemplate(val *string) {
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
func IdpSaml_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdpSaml_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"okta.idpSaml.IdpSaml",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdpSaml_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdpSaml_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"okta.idpSaml.IdpSaml",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdpSaml_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdpSaml_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"okta.idpSaml.IdpSaml",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IdpSaml_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"okta.idpSaml.IdpSaml",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IdpSaml) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IdpSaml) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IdpSaml) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IdpSaml) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IdpSaml) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IdpSaml) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IdpSaml) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IdpSaml) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IdpSaml) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IdpSaml) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IdpSaml) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IdpSaml) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IdpSaml) ResetAccountLinkAction() {
	_jsii_.InvokeVoid(
		i,
		"resetAccountLinkAction",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetAccountLinkGroupInclude() {
	_jsii_.InvokeVoid(
		i,
		"resetAccountLinkGroupInclude",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetAcsBinding() {
	_jsii_.InvokeVoid(
		i,
		"resetAcsBinding",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetAcsType() {
	_jsii_.InvokeVoid(
		i,
		"resetAcsType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetDeprovisionedAction() {
	_jsii_.InvokeVoid(
		i,
		"resetDeprovisionedAction",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetGroupsAction() {
	_jsii_.InvokeVoid(
		i,
		"resetGroupsAction",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetGroupsAssignment() {
	_jsii_.InvokeVoid(
		i,
		"resetGroupsAssignment",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetGroupsAttribute() {
	_jsii_.InvokeVoid(
		i,
		"resetGroupsAttribute",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetGroupsFilter() {
	_jsii_.InvokeVoid(
		i,
		"resetGroupsFilter",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetIssuerMode() {
	_jsii_.InvokeVoid(
		i,
		"resetIssuerMode",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetMaxClockSkew() {
	_jsii_.InvokeVoid(
		i,
		"resetMaxClockSkew",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetNameFormat() {
	_jsii_.InvokeVoid(
		i,
		"resetNameFormat",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetProfileMaster() {
	_jsii_.InvokeVoid(
		i,
		"resetProfileMaster",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetProvisioningAction() {
	_jsii_.InvokeVoid(
		i,
		"resetProvisioningAction",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetRequestSignatureAlgorithm() {
	_jsii_.InvokeVoid(
		i,
		"resetRequestSignatureAlgorithm",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetRequestSignatureScope() {
	_jsii_.InvokeVoid(
		i,
		"resetRequestSignatureScope",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetResponseSignatureAlgorithm() {
	_jsii_.InvokeVoid(
		i,
		"resetResponseSignatureAlgorithm",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetResponseSignatureScope() {
	_jsii_.InvokeVoid(
		i,
		"resetResponseSignatureScope",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetSsoBinding() {
	_jsii_.InvokeVoid(
		i,
		"resetSsoBinding",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetSsoDestination() {
	_jsii_.InvokeVoid(
		i,
		"resetSsoDestination",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetStatus() {
	_jsii_.InvokeVoid(
		i,
		"resetStatus",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetSubjectFilter() {
	_jsii_.InvokeVoid(
		i,
		"resetSubjectFilter",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetSubjectFormat() {
	_jsii_.InvokeVoid(
		i,
		"resetSubjectFormat",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetSubjectMatchAttribute() {
	_jsii_.InvokeVoid(
		i,
		"resetSubjectMatchAttribute",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetSubjectMatchType() {
	_jsii_.InvokeVoid(
		i,
		"resetSubjectMatchType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetSuspendedAction() {
	_jsii_.InvokeVoid(
		i,
		"resetSuspendedAction",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) ResetUsernameTemplate() {
	_jsii_.InvokeVoid(
		i,
		"resetUsernameTemplate",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdpSaml) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpSaml) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpSaml) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdpSaml) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


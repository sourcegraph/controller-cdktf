package organization

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/tfe/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/tfe/organization/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/tfe/0.42.0/docs/resources/organization tfe_organization}.
type Organization interface {
	cdktf.TerraformResource
	AllowForceDeleteWorkspaces() interface{}
	SetAllowForceDeleteWorkspaces(val interface{})
	AllowForceDeleteWorkspacesInput() interface{}
	AssessmentsEnforced() interface{}
	SetAssessmentsEnforced(val interface{})
	AssessmentsEnforcedInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CollaboratorAuthPolicy() *string
	SetCollaboratorAuthPolicy(val *string)
	CollaboratorAuthPolicyInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CostEstimationEnabled() interface{}
	SetCostEstimationEnabled(val interface{})
	CostEstimationEnabledInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DefaultProjectId() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OwnersTeamSamlRoleId() *string
	SetOwnersTeamSamlRoleId(val *string)
	OwnersTeamSamlRoleIdInput() *string
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
	SendPassingStatusesForUntriggeredSpeculativePlans() interface{}
	SetSendPassingStatusesForUntriggeredSpeculativePlans(val interface{})
	SendPassingStatusesForUntriggeredSpeculativePlansInput() interface{}
	SessionRememberMinutes() *float64
	SetSessionRememberMinutes(val *float64)
	SessionRememberMinutesInput() *float64
	SessionTimeoutMinutes() *float64
	SetSessionTimeoutMinutes(val *float64)
	SessionTimeoutMinutesInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetAllowForceDeleteWorkspaces()
	ResetAssessmentsEnforced()
	ResetCollaboratorAuthPolicy()
	ResetCostEstimationEnabled()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwnersTeamSamlRoleId()
	ResetSendPassingStatusesForUntriggeredSpeculativePlans()
	ResetSessionRememberMinutes()
	ResetSessionTimeoutMinutes()
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

// The jsii proxy struct for Organization
type jsiiProxy_Organization struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Organization) AllowForceDeleteWorkspaces() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowForceDeleteWorkspaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) AllowForceDeleteWorkspacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowForceDeleteWorkspacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) AssessmentsEnforced() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assessmentsEnforced",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) AssessmentsEnforcedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assessmentsEnforcedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) CollaboratorAuthPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collaboratorAuthPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) CollaboratorAuthPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collaboratorAuthPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) CostEstimationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"costEstimationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) CostEstimationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"costEstimationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) DefaultProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) OwnersTeamSamlRoleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownersTeamSamlRoleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) OwnersTeamSamlRoleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownersTeamSamlRoleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) SendPassingStatusesForUntriggeredSpeculativePlans() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendPassingStatusesForUntriggeredSpeculativePlans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) SendPassingStatusesForUntriggeredSpeculativePlansInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendPassingStatusesForUntriggeredSpeculativePlansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) SessionRememberMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionRememberMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) SessionRememberMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionRememberMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) SessionTimeoutMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) SessionTimeoutMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Organization) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/tfe/0.42.0/docs/resources/organization tfe_organization} Resource.
func NewOrganization(scope constructs.Construct, id *string, config *OrganizationConfig) Organization {
	_init_.Initialize()

	if err := validateNewOrganizationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Organization{}

	_jsii_.Create(
		"@cdktf/provider-tfe.organization.Organization",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/tfe/0.42.0/docs/resources/organization tfe_organization} Resource.
func NewOrganization_Override(o Organization, scope constructs.Construct, id *string, config *OrganizationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-tfe.organization.Organization",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_Organization)SetAllowForceDeleteWorkspaces(val interface{}) {
	if err := j.validateSetAllowForceDeleteWorkspacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowForceDeleteWorkspaces",
		val,
	)
}

func (j *jsiiProxy_Organization)SetAssessmentsEnforced(val interface{}) {
	if err := j.validateSetAssessmentsEnforcedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assessmentsEnforced",
		val,
	)
}

func (j *jsiiProxy_Organization)SetCollaboratorAuthPolicy(val *string) {
	if err := j.validateSetCollaboratorAuthPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collaboratorAuthPolicy",
		val,
	)
}

func (j *jsiiProxy_Organization)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Organization)SetCostEstimationEnabled(val interface{}) {
	if err := j.validateSetCostEstimationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"costEstimationEnabled",
		val,
	)
}

func (j *jsiiProxy_Organization)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Organization)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Organization)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_Organization)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Organization)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Organization)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Organization)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Organization)SetOwnersTeamSamlRoleId(val *string) {
	if err := j.validateSetOwnersTeamSamlRoleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ownersTeamSamlRoleId",
		val,
	)
}

func (j *jsiiProxy_Organization)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Organization)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Organization)SetSendPassingStatusesForUntriggeredSpeculativePlans(val interface{}) {
	if err := j.validateSetSendPassingStatusesForUntriggeredSpeculativePlansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendPassingStatusesForUntriggeredSpeculativePlans",
		val,
	)
}

func (j *jsiiProxy_Organization)SetSessionRememberMinutes(val *float64) {
	if err := j.validateSetSessionRememberMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionRememberMinutes",
		val,
	)
}

func (j *jsiiProxy_Organization)SetSessionTimeoutMinutes(val *float64) {
	if err := j.validateSetSessionTimeoutMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionTimeoutMinutes",
		val,
	)
}

// Generates CDKTF code for importing a Organization resource upon running "cdktf plan <stack-name>".
func Organization_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOrganization_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-tfe.organization.Organization",
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
func Organization_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOrganization_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-tfe.organization.Organization",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Organization_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOrganization_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-tfe.organization.Organization",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Organization_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOrganization_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-tfe.organization.Organization",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Organization_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-tfe.organization.Organization",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_Organization) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_Organization) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_Organization) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_Organization) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_Organization) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_Organization) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_Organization) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_Organization) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_Organization) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_Organization) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_Organization) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_Organization) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Organization) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_Organization) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_Organization) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_Organization) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_Organization) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_Organization) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_Organization) ResetAllowForceDeleteWorkspaces() {
	_jsii_.InvokeVoid(
		o,
		"resetAllowForceDeleteWorkspaces",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Organization) ResetAssessmentsEnforced() {
	_jsii_.InvokeVoid(
		o,
		"resetAssessmentsEnforced",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Organization) ResetCollaboratorAuthPolicy() {
	_jsii_.InvokeVoid(
		o,
		"resetCollaboratorAuthPolicy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Organization) ResetCostEstimationEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetCostEstimationEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Organization) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Organization) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Organization) ResetOwnersTeamSamlRoleId() {
	_jsii_.InvokeVoid(
		o,
		"resetOwnersTeamSamlRoleId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Organization) ResetSendPassingStatusesForUntriggeredSpeculativePlans() {
	_jsii_.InvokeVoid(
		o,
		"resetSendPassingStatusesForUntriggeredSpeculativePlans",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Organization) ResetSessionRememberMinutes() {
	_jsii_.InvokeVoid(
		o,
		"resetSessionRememberMinutes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Organization) ResetSessionTimeoutMinutes() {
	_jsii_.InvokeVoid(
		o,
		"resetSessionTimeoutMinutes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Organization) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Organization) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Organization) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Organization) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Organization) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Organization) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


package googlebigquerycapacitycommitment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebigquerycapacitycommitment/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_bigquery_capacity_commitment google_bigquery_capacity_commitment}.
type GoogleBigqueryCapacityCommitment interface {
	cdktf.TerraformResource
	CapacityCommitmentId() *string
	SetCapacityCommitmentId(val *string)
	CapacityCommitmentIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CommitmentEndTime() *string
	CommitmentStartTime() *string
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
	Edition() *string
	SetEdition(val *string)
	EditionInput() *string
	EnforceSingleAdminProjectPerOrg() *string
	SetEnforceSingleAdminProjectPerOrg(val *string)
	EnforceSingleAdminProjectPerOrgInput() *string
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
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	Plan() *string
	SetPlan(val *string)
	PlanInput() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
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
	RenewalPlan() *string
	SetRenewalPlan(val *string)
	RenewalPlanInput() *string
	SlotCount() *float64
	SetSlotCount(val *float64)
	SlotCountInput() *float64
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleBigqueryCapacityCommitmentTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutTimeouts(value *GoogleBigqueryCapacityCommitmentTimeouts)
	ResetCapacityCommitmentId()
	ResetEdition()
	ResetEnforceSingleAdminProjectPerOrg()
	ResetId()
	ResetLocation()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRenewalPlan()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleBigqueryCapacityCommitment
type jsiiProxy_GoogleBigqueryCapacityCommitment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) CapacityCommitmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityCommitmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) CapacityCommitmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityCommitmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) CommitmentEndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitmentEndTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) CommitmentStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitmentStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) EditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"editionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) EnforceSingleAdminProjectPerOrg() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceSingleAdminProjectPerOrg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) EnforceSingleAdminProjectPerOrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceSingleAdminProjectPerOrgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) Plan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"plan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) PlanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"planInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) RenewalPlan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renewalPlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) RenewalPlanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renewalPlanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) SlotCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slotCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) SlotCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slotCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) Timeouts() GoogleBigqueryCapacityCommitmentTimeoutsOutputReference {
	var returns GoogleBigqueryCapacityCommitmentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_bigquery_capacity_commitment google_bigquery_capacity_commitment} Resource.
func NewGoogleBigqueryCapacityCommitment(scope constructs.Construct, id *string, config *GoogleBigqueryCapacityCommitmentConfig) GoogleBigqueryCapacityCommitment {
	_init_.Initialize()

	if err := validateNewGoogleBigqueryCapacityCommitmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigqueryCapacityCommitment{}

	_jsii_.Create(
		"google-beta.googleBigqueryCapacityCommitment.GoogleBigqueryCapacityCommitment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_bigquery_capacity_commitment google_bigquery_capacity_commitment} Resource.
func NewGoogleBigqueryCapacityCommitment_Override(g GoogleBigqueryCapacityCommitment, scope constructs.Construct, id *string, config *GoogleBigqueryCapacityCommitmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleBigqueryCapacityCommitment.GoogleBigqueryCapacityCommitment",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment)SetCapacityCommitmentId(val *string) {
	if err := j.validateSetCapacityCommitmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityCommitmentId",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment)SetEdition(val *string) {
	if err := j.validateSetEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edition",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment)SetEnforceSingleAdminProjectPerOrg(val *string) {
	if err := j.validateSetEnforceSingleAdminProjectPerOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceSingleAdminProjectPerOrg",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment)SetPlan(val *string) {
	if err := j.validateSetPlanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"plan",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment)SetRenewalPlan(val *string) {
	if err := j.validateSetRenewalPlanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renewalPlan",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryCapacityCommitment)SetSlotCount(val *float64) {
	if err := j.validateSetSlotCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slotCount",
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
func GoogleBigqueryCapacityCommitment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBigqueryCapacityCommitment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleBigqueryCapacityCommitment.GoogleBigqueryCapacityCommitment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleBigqueryCapacityCommitment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBigqueryCapacityCommitment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleBigqueryCapacityCommitment.GoogleBigqueryCapacityCommitment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleBigqueryCapacityCommitment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBigqueryCapacityCommitment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleBigqueryCapacityCommitment.GoogleBigqueryCapacityCommitment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleBigqueryCapacityCommitment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.googleBigqueryCapacityCommitment.GoogleBigqueryCapacityCommitment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) PutTimeouts(value *GoogleBigqueryCapacityCommitmentTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) ResetCapacityCommitmentId() {
	_jsii_.InvokeVoid(
		g,
		"resetCapacityCommitmentId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) ResetEdition() {
	_jsii_.InvokeVoid(
		g,
		"resetEdition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) ResetEnforceSingleAdminProjectPerOrg() {
	_jsii_.InvokeVoid(
		g,
		"resetEnforceSingleAdminProjectPerOrg",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) ResetLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) ResetRenewalPlan() {
	_jsii_.InvokeVoid(
		g,
		"resetRenewalPlan",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryCapacityCommitment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


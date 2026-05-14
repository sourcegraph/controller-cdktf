package discoveryenginecontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/discoveryenginecontrol/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/discovery_engine_control google_discovery_engine_control}.
type DiscoveryEngineControl interface {
	cdktf.TerraformResource
	BoostAction() DiscoveryEngineControlBoostActionOutputReference
	BoostActionInput() *DiscoveryEngineControlBoostAction
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CollectionId() *string
	SetCollectionId(val *string)
	CollectionIdInput() *string
	Conditions() DiscoveryEngineControlConditionsList
	ConditionsInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ControlId() *string
	SetControlId(val *string)
	ControlIdInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EngineId() *string
	SetEngineId(val *string)
	EngineIdInput() *string
	FilterAction() DiscoveryEngineControlFilterActionOutputReference
	FilterActionInput() *DiscoveryEngineControlFilterAction
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
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	PromoteAction() DiscoveryEngineControlPromoteActionOutputReference
	PromoteActionInput() *DiscoveryEngineControlPromoteAction
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
	RedirectAction() DiscoveryEngineControlRedirectActionOutputReference
	RedirectActionInput() *DiscoveryEngineControlRedirectAction
	SolutionType() *string
	SetSolutionType(val *string)
	SolutionTypeInput() *string
	SynonymsAction() DiscoveryEngineControlSynonymsActionOutputReference
	SynonymsActionInput() *DiscoveryEngineControlSynonymsAction
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DiscoveryEngineControlTimeoutsOutputReference
	TimeoutsInput() interface{}
	UseCases() *[]*string
	SetUseCases(val *[]*string)
	UseCasesInput() *[]*string
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
	PutBoostAction(value *DiscoveryEngineControlBoostAction)
	PutConditions(value interface{})
	PutFilterAction(value *DiscoveryEngineControlFilterAction)
	PutPromoteAction(value *DiscoveryEngineControlPromoteAction)
	PutRedirectAction(value *DiscoveryEngineControlRedirectAction)
	PutSynonymsAction(value *DiscoveryEngineControlSynonymsAction)
	PutTimeouts(value *DiscoveryEngineControlTimeouts)
	ResetBoostAction()
	ResetCollectionId()
	ResetConditions()
	ResetFilterAction()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetPromoteAction()
	ResetRedirectAction()
	ResetSynonymsAction()
	ResetTimeouts()
	ResetUseCases()
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

// The jsii proxy struct for DiscoveryEngineControl
type jsiiProxy_DiscoveryEngineControl struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DiscoveryEngineControl) BoostAction() DiscoveryEngineControlBoostActionOutputReference {
	var returns DiscoveryEngineControlBoostActionOutputReference
	_jsii_.Get(
		j,
		"boostAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) BoostActionInput() *DiscoveryEngineControlBoostAction {
	var returns *DiscoveryEngineControlBoostAction
	_jsii_.Get(
		j,
		"boostActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) CollectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) CollectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) Conditions() DiscoveryEngineControlConditionsList {
	var returns DiscoveryEngineControlConditionsList
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) ConditionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) ControlId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) ControlIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) EngineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) EngineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) FilterAction() DiscoveryEngineControlFilterActionOutputReference {
	var returns DiscoveryEngineControlFilterActionOutputReference
	_jsii_.Get(
		j,
		"filterAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) FilterActionInput() *DiscoveryEngineControlFilterAction {
	var returns *DiscoveryEngineControlFilterAction
	_jsii_.Get(
		j,
		"filterActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) PromoteAction() DiscoveryEngineControlPromoteActionOutputReference {
	var returns DiscoveryEngineControlPromoteActionOutputReference
	_jsii_.Get(
		j,
		"promoteAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) PromoteActionInput() *DiscoveryEngineControlPromoteAction {
	var returns *DiscoveryEngineControlPromoteAction
	_jsii_.Get(
		j,
		"promoteActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) RedirectAction() DiscoveryEngineControlRedirectActionOutputReference {
	var returns DiscoveryEngineControlRedirectActionOutputReference
	_jsii_.Get(
		j,
		"redirectAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) RedirectActionInput() *DiscoveryEngineControlRedirectAction {
	var returns *DiscoveryEngineControlRedirectAction
	_jsii_.Get(
		j,
		"redirectActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) SolutionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"solutionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) SolutionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"solutionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) SynonymsAction() DiscoveryEngineControlSynonymsActionOutputReference {
	var returns DiscoveryEngineControlSynonymsActionOutputReference
	_jsii_.Get(
		j,
		"synonymsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) SynonymsActionInput() *DiscoveryEngineControlSynonymsAction {
	var returns *DiscoveryEngineControlSynonymsAction
	_jsii_.Get(
		j,
		"synonymsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) Timeouts() DiscoveryEngineControlTimeoutsOutputReference {
	var returns DiscoveryEngineControlTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) UseCases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"useCases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControl) UseCasesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"useCasesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/discovery_engine_control google_discovery_engine_control} Resource.
func NewDiscoveryEngineControl(scope constructs.Construct, id *string, config *DiscoveryEngineControlConfig) DiscoveryEngineControl {
	_init_.Initialize()

	if err := validateNewDiscoveryEngineControlParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DiscoveryEngineControl{}

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineControl.DiscoveryEngineControl",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/discovery_engine_control google_discovery_engine_control} Resource.
func NewDiscoveryEngineControl_Override(d DiscoveryEngineControl, scope constructs.Construct, id *string, config *DiscoveryEngineControlConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineControl.DiscoveryEngineControl",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DiscoveryEngineControl)SetCollectionId(val *string) {
	if err := j.validateSetCollectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectionId",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControl)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControl)SetControlId(val *string) {
	if err := j.validateSetControlIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controlId",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControl)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControl)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControl)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControl)SetEngineId(val *string) {
	if err := j.validateSetEngineIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineId",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControl)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControl)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControl)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControl)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControl)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControl)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControl)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControl)SetSolutionType(val *string) {
	if err := j.validateSetSolutionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"solutionType",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControl)SetUseCases(val *[]*string) {
	if err := j.validateSetUseCasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCases",
		val,
	)
}

// Generates CDKTF code for importing a DiscoveryEngineControl resource upon running "cdktf plan <stack-name>".
func DiscoveryEngineControl_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDiscoveryEngineControl_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.discoveryEngineControl.DiscoveryEngineControl",
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
func DiscoveryEngineControl_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDiscoveryEngineControl_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.discoveryEngineControl.DiscoveryEngineControl",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DiscoveryEngineControl_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDiscoveryEngineControl_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.discoveryEngineControl.DiscoveryEngineControl",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DiscoveryEngineControl_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDiscoveryEngineControl_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.discoveryEngineControl.DiscoveryEngineControl",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DiscoveryEngineControl_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.discoveryEngineControl.DiscoveryEngineControl",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DiscoveryEngineControl) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineControl) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineControl) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineControl) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineControl) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineControl) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineControl) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineControl) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineControl) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineControl) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineControl) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineControl) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) PutBoostAction(value *DiscoveryEngineControlBoostAction) {
	if err := d.validatePutBoostActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBoostAction",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) PutConditions(value interface{}) {
	if err := d.validatePutConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConditions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) PutFilterAction(value *DiscoveryEngineControlFilterAction) {
	if err := d.validatePutFilterActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFilterAction",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) PutPromoteAction(value *DiscoveryEngineControlPromoteAction) {
	if err := d.validatePutPromoteActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPromoteAction",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) PutRedirectAction(value *DiscoveryEngineControlRedirectAction) {
	if err := d.validatePutRedirectActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRedirectAction",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) PutSynonymsAction(value *DiscoveryEngineControlSynonymsAction) {
	if err := d.validatePutSynonymsActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSynonymsAction",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) PutTimeouts(value *DiscoveryEngineControlTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) ResetBoostAction() {
	_jsii_.InvokeVoid(
		d,
		"resetBoostAction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) ResetCollectionId() {
	_jsii_.InvokeVoid(
		d,
		"resetCollectionId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) ResetConditions() {
	_jsii_.InvokeVoid(
		d,
		"resetConditions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) ResetFilterAction() {
	_jsii_.InvokeVoid(
		d,
		"resetFilterAction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) ResetPromoteAction() {
	_jsii_.InvokeVoid(
		d,
		"resetPromoteAction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) ResetRedirectAction() {
	_jsii_.InvokeVoid(
		d,
		"resetRedirectAction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) ResetSynonymsAction() {
	_jsii_.InvokeVoid(
		d,
		"resetSynonymsAction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) ResetUseCases() {
	_jsii_.InvokeVoid(
		d,
		"resetUseCases",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineControl) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineControl) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineControl) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineControl) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineControl) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineControl) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


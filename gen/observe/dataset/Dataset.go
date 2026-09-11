package dataset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/observe/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/observe/dataset/internal"
)

// Represents a {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset observe_dataset}.
type Dataset interface {
	cdktn.TerraformResource
	AccelerationDisabled() interface{}
	SetAccelerationDisabled(val interface{})
	AccelerationDisabledInput() interface{}
	AccelerationDisabledSource() *string
	SetAccelerationDisabledSource(val *string)
	AccelerationDisabledSourceInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	DataTableViewState() *string
	SetDataTableViewState(val *string)
	DataTableViewStateInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	Freshness() *string
	SetFreshness(val *string)
	FreshnessInput() *string
	// Experimental.
	FriendlyUniqueId() *string
	IconUrl() *string
	SetIconUrl(val *string)
	IconUrlInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Inputs() *map[string]*string
	SetInputs(val *map[string]*string)
	InputsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Oid() *string
	OnDemandMaterializationLength() *string
	SetOnDemandMaterializationLength(val *string)
	OnDemandMaterializationLengthInput() *string
	PathCost() *float64
	SetPathCost(val *float64)
	PathCostInput() *float64
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RematerializationMode() *string
	SetRematerializationMode(val *string)
	RematerializationModeInput() *string
	Stage() DatasetStageList
	StageInput() interface{}
	StorageIntegration() *string
	SetStorageIntegration(val *string)
	StorageIntegrationInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Workspace() *string
	SetWorkspace(val *string)
	WorkspaceInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
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
	PutStage(value interface{})
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetAccelerationDisabled()
	ResetAccelerationDisabledSource()
	ResetDataTableViewState()
	ResetDescription()
	ResetFreshness()
	ResetIconUrl()
	ResetId()
	ResetOnDemandMaterializationLength()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPathCost()
	ResetRematerializationMode()
	ResetStorageIntegration()
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for Dataset
type jsiiProxy_Dataset struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_Dataset) AccelerationDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accelerationDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) AccelerationDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accelerationDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) AccelerationDisabledSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accelerationDisabledSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) AccelerationDisabledSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accelerationDisabledSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) DataTableViewState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTableViewState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) DataTableViewStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTableViewStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) Freshness() *string {
	var returns *string
	_jsii_.Get(
		j,
		"freshness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) FreshnessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"freshnessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) IconUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iconUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) IconUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iconUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) Inputs() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"inputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) InputsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"inputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) Oid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) OnDemandMaterializationLength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDemandMaterializationLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) OnDemandMaterializationLengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDemandMaterializationLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) PathCost() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pathCost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) PathCostInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pathCostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) RematerializationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rematerializationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) RematerializationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rematerializationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) Stage() DatasetStageList {
	var returns DatasetStageList
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) StageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) StorageIntegration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) StorageIntegrationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageIntegrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) Workspace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataset) WorkspaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset observe_dataset} Resource.
func NewDataset(scope constructs.Construct, id *string, config *DatasetConfig) Dataset {
	_init_.Initialize()

	if err := validateNewDatasetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Dataset{}

	_jsii_.Create(
		"@cdktn/provider-observe.dataset.Dataset",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset observe_dataset} Resource.
func NewDataset_Override(d Dataset, scope constructs.Construct, id *string, config *DatasetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-observe.dataset.Dataset",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_Dataset)SetAccelerationDisabled(val interface{}) {
	if err := j.validateSetAccelerationDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accelerationDisabled",
		val,
	)
}

func (j *jsiiProxy_Dataset)SetAccelerationDisabledSource(val *string) {
	if err := j.validateSetAccelerationDisabledSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accelerationDisabledSource",
		val,
	)
}

func (j *jsiiProxy_Dataset)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Dataset)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Dataset)SetDataTableViewState(val *string) {
	if err := j.validateSetDataTableViewStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataTableViewState",
		val,
	)
}

func (j *jsiiProxy_Dataset)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Dataset)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Dataset)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Dataset)SetFreshness(val *string) {
	if err := j.validateSetFreshnessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"freshness",
		val,
	)
}

func (j *jsiiProxy_Dataset)SetIconUrl(val *string) {
	if err := j.validateSetIconUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iconUrl",
		val,
	)
}

func (j *jsiiProxy_Dataset)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Dataset)SetInputs(val *map[string]*string) {
	if err := j.validateSetInputsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputs",
		val,
	)
}

func (j *jsiiProxy_Dataset)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Dataset)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Dataset)SetOnDemandMaterializationLength(val *string) {
	if err := j.validateSetOnDemandMaterializationLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandMaterializationLength",
		val,
	)
}

func (j *jsiiProxy_Dataset)SetPathCost(val *float64) {
	if err := j.validateSetPathCostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathCost",
		val,
	)
}

func (j *jsiiProxy_Dataset)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Dataset)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Dataset)SetRematerializationMode(val *string) {
	if err := j.validateSetRematerializationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rematerializationMode",
		val,
	)
}

func (j *jsiiProxy_Dataset)SetStorageIntegration(val *string) {
	if err := j.validateSetStorageIntegrationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageIntegration",
		val,
	)
}

func (j *jsiiProxy_Dataset)SetWorkspace(val *string) {
	if err := j.validateSetWorkspaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspace",
		val,
	)
}

// Generates CDKTN code for importing a Dataset resource upon running "cdktn plan <stack-name>".
func Dataset_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataset_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-observe.dataset.Dataset",
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
func Dataset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataset_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-observe.dataset.Dataset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Dataset_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataset_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-observe.dataset.Dataset",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Dataset_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataset_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-observe.dataset.Dataset",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Dataset_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-observe.dataset.Dataset",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_Dataset) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_Dataset) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_Dataset) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_Dataset) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dataset) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_Dataset) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_Dataset) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_Dataset) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_Dataset) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_Dataset) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_Dataset) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_Dataset) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dataset) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_Dataset) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dataset) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := d.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dataset) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_Dataset) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_Dataset) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_Dataset) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_Dataset) PutStage(value interface{}) {
	if err := d.validatePutStageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStage",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_Dataset) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_Dataset) ResetAccelerationDisabled() {
	_jsii_.InvokeVoid(
		d,
		"resetAccelerationDisabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dataset) ResetAccelerationDisabledSource() {
	_jsii_.InvokeVoid(
		d,
		"resetAccelerationDisabledSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dataset) ResetDataTableViewState() {
	_jsii_.InvokeVoid(
		d,
		"resetDataTableViewState",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dataset) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dataset) ResetFreshness() {
	_jsii_.InvokeVoid(
		d,
		"resetFreshness",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dataset) ResetIconUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetIconUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dataset) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dataset) ResetOnDemandMaterializationLength() {
	_jsii_.InvokeVoid(
		d,
		"resetOnDemandMaterializationLength",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dataset) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dataset) ResetPathCost() {
	_jsii_.InvokeVoid(
		d,
		"resetPathCost",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dataset) ResetRematerializationMode() {
	_jsii_.InvokeVoid(
		d,
		"resetRematerializationMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dataset) ResetStorageIntegration() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageIntegration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dataset) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dataset) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dataset) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dataset) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dataset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dataset) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dataset) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}


package workflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/incident/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/incident/workflow/internal"
)

// Represents a {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow incident_workflow}.
type Workflow interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ConditionGroups() WorkflowConditionGroupsList
	ConditionGroupsInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContinueOnStepError() interface{}
	SetContinueOnStepError(val interface{})
	ContinueOnStepErrorInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Delay() WorkflowDelayOutputReference
	DelayInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Expressions() WorkflowExpressionsList
	ExpressionsInput() interface{}
	Folder() *string
	SetFolder(val *string)
	FolderInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IncludePrivateIncidents() interface{}
	SetIncludePrivateIncidents(val interface{})
	IncludePrivateIncidentsInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OnceFor() *[]*string
	SetOnceFor(val *[]*string)
	OnceForInput() *[]*string
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
	RunsOnIncidentModes() *[]*string
	SetRunsOnIncidentModes(val *[]*string)
	RunsOnIncidentModesInput() *[]*string
	RunsOnIncidents() *string
	SetRunsOnIncidents(val *string)
	RunsOnIncidentsInput() *string
	Shortform() *string
	SetShortform(val *string)
	ShortformInput() *string
	State() *string
	SetState(val *string)
	StateInput() *string
	Steps() WorkflowStepsList
	StepsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Trigger() *string
	SetTrigger(val *string)
	TriggerInput() *string
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
	PutConditionGroups(value interface{})
	PutDelay(value *WorkflowDelay)
	PutExpressions(value interface{})
	PutSteps(value interface{})
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
	ResetDelay()
	ResetFolder()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetShortform()
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

// The jsii proxy struct for Workflow
type jsiiProxy_Workflow struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_Workflow) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) ConditionGroups() WorkflowConditionGroupsList {
	var returns WorkflowConditionGroupsList
	_jsii_.Get(
		j,
		"conditionGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) ConditionGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) ContinueOnStepError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continueOnStepError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) ContinueOnStepErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continueOnStepErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Delay() WorkflowDelayOutputReference {
	var returns WorkflowDelayOutputReference
	_jsii_.Get(
		j,
		"delay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) DelayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"delayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Expressions() WorkflowExpressionsList {
	var returns WorkflowExpressionsList
	_jsii_.Get(
		j,
		"expressions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) ExpressionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expressionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Folder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) FolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) IncludePrivateIncidents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includePrivateIncidents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) IncludePrivateIncidentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includePrivateIncidentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) OnceFor() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onceFor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) OnceForInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onceForInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) RunsOnIncidentModes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"runsOnIncidentModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) RunsOnIncidentModesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"runsOnIncidentModesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) RunsOnIncidents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runsOnIncidents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) RunsOnIncidentsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runsOnIncidentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Shortform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) ShortformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Steps() WorkflowStepsList {
	var returns WorkflowStepsList
	_jsii_.Get(
		j,
		"steps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) StepsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Trigger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) TriggerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow incident_workflow} Resource.
func NewWorkflow(scope constructs.Construct, id *string, config *WorkflowConfig) Workflow {
	_init_.Initialize()

	if err := validateNewWorkflowParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Workflow{}

	_jsii_.Create(
		"@cdktn/provider-incident.workflow.Workflow",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow incident_workflow} Resource.
func NewWorkflow_Override(w Workflow, scope constructs.Construct, id *string, config *WorkflowConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-incident.workflow.Workflow",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_Workflow)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Workflow)SetContinueOnStepError(val interface{}) {
	if err := j.validateSetContinueOnStepErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"continueOnStepError",
		val,
	)
}

func (j *jsiiProxy_Workflow)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Workflow)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Workflow)SetFolder(val *string) {
	if err := j.validateSetFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"folder",
		val,
	)
}

func (j *jsiiProxy_Workflow)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Workflow)SetIncludePrivateIncidents(val interface{}) {
	if err := j.validateSetIncludePrivateIncidentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includePrivateIncidents",
		val,
	)
}

func (j *jsiiProxy_Workflow)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Workflow)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Workflow)SetOnceFor(val *[]*string) {
	if err := j.validateSetOnceForParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onceFor",
		val,
	)
}

func (j *jsiiProxy_Workflow)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Workflow)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Workflow)SetRunsOnIncidentModes(val *[]*string) {
	if err := j.validateSetRunsOnIncidentModesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runsOnIncidentModes",
		val,
	)
}

func (j *jsiiProxy_Workflow)SetRunsOnIncidents(val *string) {
	if err := j.validateSetRunsOnIncidentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runsOnIncidents",
		val,
	)
}

func (j *jsiiProxy_Workflow)SetShortform(val *string) {
	if err := j.validateSetShortformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shortform",
		val,
	)
}

func (j *jsiiProxy_Workflow)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_Workflow)SetTrigger(val *string) {
	if err := j.validateSetTriggerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trigger",
		val,
	)
}

// Generates CDKTN code for importing a Workflow resource upon running "cdktn plan <stack-name>".
func Workflow_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateWorkflow_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-incident.workflow.Workflow",
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
func Workflow_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkflow_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-incident.workflow.Workflow",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Workflow_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkflow_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-incident.workflow.Workflow",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Workflow_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkflow_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-incident.workflow.Workflow",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Workflow_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-incident.workflow.Workflow",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_Workflow) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_Workflow) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_Workflow) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_Workflow) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := w.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_Workflow) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_Workflow) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_Workflow) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_Workflow) PutConditionGroups(value interface{}) {
	if err := w.validatePutConditionGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putConditionGroups",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Workflow) PutDelay(value *WorkflowDelay) {
	if err := w.validatePutDelayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putDelay",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Workflow) PutExpressions(value interface{}) {
	if err := w.validatePutExpressionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putExpressions",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Workflow) PutSteps(value interface{}) {
	if err := w.validatePutStepsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSteps",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Workflow) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := w.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (w *jsiiProxy_Workflow) ResetDelay() {
	_jsii_.InvokeVoid(
		w,
		"resetDelay",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workflow) ResetFolder() {
	_jsii_.InvokeVoid(
		w,
		"resetFolder",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workflow) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workflow) ResetShortform() {
	_jsii_.InvokeVoid(
		w,
		"resetShortform",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workflow) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		w,
		"with",
		args,
		&returns,
	)

	return returns
}


package workflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/incident/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/incident/workflow/internal"
)

// Represents a {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow incident_workflow}.
type Workflow interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConditionGroups() WorkflowConditionGroupsList
	ConditionGroupsInput() any
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	ContinueOnStepError() any
	SetContinueOnStepError(val any)
	ContinueOnStepErrorInput() any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	Delay() WorkflowDelayOutputReference
	DelayInput() any
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Expressions() WorkflowExpressionsList
	ExpressionsInput() any
	Folder() *string
	SetFolder(val *string)
	FolderInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IncludePrivateIncidents() any
	SetIncludePrivateIncidents(val any)
	IncludePrivateIncidentsInput() any
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OnceFor() *[]*string
	SetOnceFor(val *[]*string)
	OnceForInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	// Experimental.
	RawOverrides() any
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
	StepsInput() any
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Trigger() *string
	SetTrigger(val *string)
	TriggerInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value any)
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	HasResourceMove() any
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
	MoveTo(moveTarget *string, index any)
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutConditionGroups(value any)
	PutDelay(value *WorkflowDelay)
	PutExpressions(value any)
	PutSteps(value any)
	ResetDelay()
	ResetFolder()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetShortform()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for Workflow
type jsiiProxy_Workflow struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Workflow) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
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

func (j *jsiiProxy_Workflow) ConditionGroupsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"conditionGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) ContinueOnStepError() any {
	var returns any
	_jsii_.Get(
		j,
		"continueOnStepError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) ContinueOnStepErrorInput() any {
	var returns any
	_jsii_.Get(
		j,
		"continueOnStepErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Count() any {
	var returns any
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

func (j *jsiiProxy_Workflow) DelayInput() any {
	var returns any
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

func (j *jsiiProxy_Workflow) ExpressionsInput() any {
	var returns any
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

func (j *jsiiProxy_Workflow) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
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

func (j *jsiiProxy_Workflow) IncludePrivateIncidents() any {
	var returns any
	_jsii_.Get(
		j,
		"includePrivateIncidents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) IncludePrivateIncidentsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includePrivateIncidentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
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

func (j *jsiiProxy_Workflow) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) RawOverrides() any {
	var returns any
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

func (j *jsiiProxy_Workflow) StepsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"stepsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
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
		"@cdktf/provider-incident.workflow.Workflow",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow incident_workflow} Resource.
func NewWorkflow_Override(w Workflow, scope constructs.Construct, id *string, config *WorkflowConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-incident.workflow.Workflow",
		[]any{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_Workflow) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetContinueOnStepError(val any) {
	if err := j.validateSetContinueOnStepErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"continueOnStepError",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetFolder(val *string) {
	if err := j.validateSetFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"folder",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetIncludePrivateIncidents(val any) {
	if err := j.validateSetIncludePrivateIncidentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includePrivateIncidents",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetOnceFor(val *[]*string) {
	if err := j.validateSetOnceForParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onceFor",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetRunsOnIncidentModes(val *[]*string) {
	if err := j.validateSetRunsOnIncidentModesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runsOnIncidentModes",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetRunsOnIncidents(val *string) {
	if err := j.validateSetRunsOnIncidentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runsOnIncidents",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetShortform(val *string) {
	if err := j.validateSetShortformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shortform",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetTrigger(val *string) {
	if err := j.validateSetTriggerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trigger",
		val,
	)
}

// Generates CDKTF code for importing a Workflow resource upon running "cdktf plan <stack-name>".
func Workflow_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateWorkflow_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-incident.workflow.Workflow",
		"generateConfigForImport",
		[]any{scope, importToId, importFromId, provider},
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
func Workflow_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateWorkflow_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-incident.workflow.Workflow",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func Workflow_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateWorkflow_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-incident.workflow.Workflow",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func Workflow_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateWorkflow_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-incident.workflow.Workflow",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func Workflow_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-incident.workflow.Workflow",
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
		[]any{moveTarget},
	)
}

func (w *jsiiProxy_Workflow) AddOverride(path *string, value any) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]any{path, value},
	)
}

func (w *jsiiProxy_Workflow) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]any{id, provider},
	)
}

func (w *jsiiProxy_Workflow) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]any{terraformAttribute},
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
		[]any{id},
	)
}

func (w *jsiiProxy_Workflow) MoveTo(moveTarget *string, index any) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (w *jsiiProxy_Workflow) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]any{id},
	)
}

func (w *jsiiProxy_Workflow) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (w *jsiiProxy_Workflow) PutConditionGroups(value any) {
	if err := w.validatePutConditionGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putConditionGroups",
		[]any{value},
	)
}

func (w *jsiiProxy_Workflow) PutDelay(value *WorkflowDelay) {
	if err := w.validatePutDelayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putDelay",
		[]any{value},
	)
}

func (w *jsiiProxy_Workflow) PutExpressions(value any) {
	if err := w.validatePutExpressionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putExpressions",
		[]any{value},
	)
}

func (w *jsiiProxy_Workflow) PutSteps(value any) {
	if err := w.validatePutStepsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSteps",
		[]any{value},
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

func (w *jsiiProxy_Workflow) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) ToMetadata() any {
	var returns any

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

func (w *jsiiProxy_Workflow) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

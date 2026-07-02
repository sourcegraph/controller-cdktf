package cesagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cesagent/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/ces_agent google_ces_agent}.
type CesAgent interface {
	cdktf.TerraformResource
	AfterAgentCallbacks() CesAgentAfterAgentCallbacksList
	AfterAgentCallbacksInput() any
	AfterModelCallbacks() CesAgentAfterModelCallbacksList
	AfterModelCallbacksInput() any
	AfterToolCallbacks() CesAgentAfterToolCallbacksList
	AfterToolCallbacksInput() any
	AgentId() *string
	SetAgentId(val *string)
	AgentIdInput() *string
	App() *string
	SetApp(val *string)
	AppInput() *string
	BeforeAgentCallbacks() CesAgentBeforeAgentCallbacksList
	BeforeAgentCallbacksInput() any
	BeforeModelCallbacks() CesAgentBeforeModelCallbacksList
	BeforeModelCallbacksInput() any
	BeforeToolCallbacks() CesAgentBeforeToolCallbacksList
	BeforeToolCallbacksInput() any
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChildAgents() *[]*string
	SetChildAgents(val *[]*string)
	ChildAgentsInput() *[]*string
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Etag() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GeneratedSummary() *string
	Guardrails() *[]*string
	SetGuardrails(val *[]*string)
	GuardrailsInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Instruction() *string
	SetInstruction(val *string)
	InstructionInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LlmAgent() CesAgentLlmAgentOutputReference
	LlmAgentInput() *CesAgentLlmAgent
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ModelSettings() CesAgentModelSettingsOutputReference
	ModelSettingsInput() *CesAgentModelSettings
	Name() *string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
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
	RemoteDialogflowAgent() CesAgentRemoteDialogflowAgentOutputReference
	RemoteDialogflowAgentInput() *CesAgentRemoteDialogflowAgent
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CesAgentTimeoutsOutputReference
	TimeoutsInput() any
	Tools() *[]*string
	SetTools(val *[]*string)
	Toolsets() CesAgentToolsetsList
	ToolsetsInput() any
	ToolsInput() *[]*string
	UpdateTime() *string
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
	PutAfterAgentCallbacks(value any)
	PutAfterModelCallbacks(value any)
	PutAfterToolCallbacks(value any)
	PutBeforeAgentCallbacks(value any)
	PutBeforeModelCallbacks(value any)
	PutBeforeToolCallbacks(value any)
	PutLlmAgent(value *CesAgentLlmAgent)
	PutModelSettings(value *CesAgentModelSettings)
	PutRemoteDialogflowAgent(value *CesAgentRemoteDialogflowAgent)
	PutTimeouts(value *CesAgentTimeouts)
	PutToolsets(value any)
	ResetAfterAgentCallbacks()
	ResetAfterModelCallbacks()
	ResetAfterToolCallbacks()
	ResetAgentId()
	ResetBeforeAgentCallbacks()
	ResetBeforeModelCallbacks()
	ResetBeforeToolCallbacks()
	ResetChildAgents()
	ResetDescription()
	ResetGuardrails()
	ResetId()
	ResetInstruction()
	ResetLlmAgent()
	ResetModelSettings()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRemoteDialogflowAgent()
	ResetTimeouts()
	ResetTools()
	ResetToolsets()
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

// The jsii proxy struct for CesAgent
type jsiiProxy_CesAgent struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CesAgent) AfterAgentCallbacks() CesAgentAfterAgentCallbacksList {
	var returns CesAgentAfterAgentCallbacksList
	_jsii_.Get(
		j,
		"afterAgentCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) AfterAgentCallbacksInput() any {
	var returns any
	_jsii_.Get(
		j,
		"afterAgentCallbacksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) AfterModelCallbacks() CesAgentAfterModelCallbacksList {
	var returns CesAgentAfterModelCallbacksList
	_jsii_.Get(
		j,
		"afterModelCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) AfterModelCallbacksInput() any {
	var returns any
	_jsii_.Get(
		j,
		"afterModelCallbacksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) AfterToolCallbacks() CesAgentAfterToolCallbacksList {
	var returns CesAgentAfterToolCallbacksList
	_jsii_.Get(
		j,
		"afterToolCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) AfterToolCallbacksInput() any {
	var returns any
	_jsii_.Get(
		j,
		"afterToolCallbacksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) AgentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) AgentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) App() *string {
	var returns *string
	_jsii_.Get(
		j,
		"app",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) AppInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) BeforeAgentCallbacks() CesAgentBeforeAgentCallbacksList {
	var returns CesAgentBeforeAgentCallbacksList
	_jsii_.Get(
		j,
		"beforeAgentCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) BeforeAgentCallbacksInput() any {
	var returns any
	_jsii_.Get(
		j,
		"beforeAgentCallbacksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) BeforeModelCallbacks() CesAgentBeforeModelCallbacksList {
	var returns CesAgentBeforeModelCallbacksList
	_jsii_.Get(
		j,
		"beforeModelCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) BeforeModelCallbacksInput() any {
	var returns any
	_jsii_.Get(
		j,
		"beforeModelCallbacksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) BeforeToolCallbacks() CesAgentBeforeToolCallbacksList {
	var returns CesAgentBeforeToolCallbacksList
	_jsii_.Get(
		j,
		"beforeToolCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) BeforeToolCallbacksInput() any {
	var returns any
	_jsii_.Get(
		j,
		"beforeToolCallbacksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) ChildAgents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"childAgents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) ChildAgentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"childAgentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) GeneratedSummary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generatedSummary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) Guardrails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guardrails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) GuardrailsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guardrailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) Instruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) InstructionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instructionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) LlmAgent() CesAgentLlmAgentOutputReference {
	var returns CesAgentLlmAgentOutputReference
	_jsii_.Get(
		j,
		"llmAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) LlmAgentInput() *CesAgentLlmAgent {
	var returns *CesAgentLlmAgent
	_jsii_.Get(
		j,
		"llmAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) ModelSettings() CesAgentModelSettingsOutputReference {
	var returns CesAgentModelSettingsOutputReference
	_jsii_.Get(
		j,
		"modelSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) ModelSettingsInput() *CesAgentModelSettings {
	var returns *CesAgentModelSettings
	_jsii_.Get(
		j,
		"modelSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) RemoteDialogflowAgent() CesAgentRemoteDialogflowAgentOutputReference {
	var returns CesAgentRemoteDialogflowAgentOutputReference
	_jsii_.Get(
		j,
		"remoteDialogflowAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) RemoteDialogflowAgentInput() *CesAgentRemoteDialogflowAgent {
	var returns *CesAgentRemoteDialogflowAgent
	_jsii_.Get(
		j,
		"remoteDialogflowAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) Timeouts() CesAgentTimeoutsOutputReference {
	var returns CesAgentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) Tools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) Toolsets() CesAgentToolsetsList {
	var returns CesAgentToolsetsList
	_jsii_.Get(
		j,
		"toolsets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) ToolsetsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"toolsetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) ToolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"toolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgent) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/ces_agent google_ces_agent} Resource.
func NewCesAgent(scope constructs.Construct, id *string, config *CesAgentConfig) CesAgent {
	_init_.Initialize()

	if err := validateNewCesAgentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesAgent{}

	_jsii_.Create(
		"@cdktf/provider-google.cesAgent.CesAgent",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/ces_agent google_ces_agent} Resource.
func NewCesAgent_Override(c CesAgent, scope constructs.Construct, id *string, config *CesAgentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesAgent.CesAgent",
		[]any{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CesAgent) SetAgentId(val *string) {
	if err := j.validateSetAgentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentId",
		val,
	)
}

func (j *jsiiProxy_CesAgent) SetApp(val *string) {
	if err := j.validateSetAppParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"app",
		val,
	)
}

func (j *jsiiProxy_CesAgent) SetChildAgents(val *[]*string) {
	if err := j.validateSetChildAgentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"childAgents",
		val,
	)
}

func (j *jsiiProxy_CesAgent) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CesAgent) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CesAgent) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CesAgent) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CesAgent) SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_CesAgent) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CesAgent) SetGuardrails(val *[]*string) {
	if err := j.validateSetGuardrailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guardrails",
		val,
	)
}

func (j *jsiiProxy_CesAgent) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CesAgent) SetInstruction(val *string) {
	if err := j.validateSetInstructionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instruction",
		val,
	)
}

func (j *jsiiProxy_CesAgent) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CesAgent) SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CesAgent) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_CesAgent) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CesAgent) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CesAgent) SetTools(val *[]*string) {
	if err := j.validateSetToolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tools",
		val,
	)
}

// Generates CDKTF code for importing a CesAgent resource upon running "cdktf plan <stack-name>".
func CesAgent_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCesAgent_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cesAgent.CesAgent",
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
func CesAgent_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateCesAgent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cesAgent.CesAgent",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func CesAgent_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateCesAgent_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cesAgent.CesAgent",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func CesAgent_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateCesAgent_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cesAgent.CesAgent",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func CesAgent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.cesAgent.CesAgent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CesAgent) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (c *jsiiProxy_CesAgent) AddOverride(path *string, value any) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]any{path, value},
	)
}

func (c *jsiiProxy_CesAgent) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgent) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgent) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgent) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgent) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgent) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgent) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]any{id, provider},
	)
}

func (c *jsiiProxy_CesAgent) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgent) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]any{id},
	)
}

func (c *jsiiProxy_CesAgent) MoveTo(moveTarget *string, index any) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (c *jsiiProxy_CesAgent) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]any{id},
	)
}

func (c *jsiiProxy_CesAgent) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (c *jsiiProxy_CesAgent) PutAfterAgentCallbacks(value any) {
	if err := c.validatePutAfterAgentCallbacksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAfterAgentCallbacks",
		[]any{value},
	)
}

func (c *jsiiProxy_CesAgent) PutAfterModelCallbacks(value any) {
	if err := c.validatePutAfterModelCallbacksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAfterModelCallbacks",
		[]any{value},
	)
}

func (c *jsiiProxy_CesAgent) PutAfterToolCallbacks(value any) {
	if err := c.validatePutAfterToolCallbacksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAfterToolCallbacks",
		[]any{value},
	)
}

func (c *jsiiProxy_CesAgent) PutBeforeAgentCallbacks(value any) {
	if err := c.validatePutBeforeAgentCallbacksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBeforeAgentCallbacks",
		[]any{value},
	)
}

func (c *jsiiProxy_CesAgent) PutBeforeModelCallbacks(value any) {
	if err := c.validatePutBeforeModelCallbacksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBeforeModelCallbacks",
		[]any{value},
	)
}

func (c *jsiiProxy_CesAgent) PutBeforeToolCallbacks(value any) {
	if err := c.validatePutBeforeToolCallbacksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBeforeToolCallbacks",
		[]any{value},
	)
}

func (c *jsiiProxy_CesAgent) PutLlmAgent(value *CesAgentLlmAgent) {
	if err := c.validatePutLlmAgentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLlmAgent",
		[]any{value},
	)
}

func (c *jsiiProxy_CesAgent) PutModelSettings(value *CesAgentModelSettings) {
	if err := c.validatePutModelSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putModelSettings",
		[]any{value},
	)
}

func (c *jsiiProxy_CesAgent) PutRemoteDialogflowAgent(value *CesAgentRemoteDialogflowAgent) {
	if err := c.validatePutRemoteDialogflowAgentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRemoteDialogflowAgent",
		[]any{value},
	)
}

func (c *jsiiProxy_CesAgent) PutTimeouts(value *CesAgentTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]any{value},
	)
}

func (c *jsiiProxy_CesAgent) PutToolsets(value any) {
	if err := c.validatePutToolsetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putToolsets",
		[]any{value},
	)
}

func (c *jsiiProxy_CesAgent) ResetAfterAgentCallbacks() {
	_jsii_.InvokeVoid(
		c,
		"resetAfterAgentCallbacks",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgent) ResetAfterModelCallbacks() {
	_jsii_.InvokeVoid(
		c,
		"resetAfterModelCallbacks",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgent) ResetAfterToolCallbacks() {
	_jsii_.InvokeVoid(
		c,
		"resetAfterToolCallbacks",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgent) ResetAgentId() {
	_jsii_.InvokeVoid(
		c,
		"resetAgentId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgent) ResetBeforeAgentCallbacks() {
	_jsii_.InvokeVoid(
		c,
		"resetBeforeAgentCallbacks",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgent) ResetBeforeModelCallbacks() {
	_jsii_.InvokeVoid(
		c,
		"resetBeforeModelCallbacks",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgent) ResetBeforeToolCallbacks() {
	_jsii_.InvokeVoid(
		c,
		"resetBeforeToolCallbacks",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgent) ResetChildAgents() {
	_jsii_.InvokeVoid(
		c,
		"resetChildAgents",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgent) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgent) ResetGuardrails() {
	_jsii_.InvokeVoid(
		c,
		"resetGuardrails",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgent) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgent) ResetInstruction() {
	_jsii_.InvokeVoid(
		c,
		"resetInstruction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgent) ResetLlmAgent() {
	_jsii_.InvokeVoid(
		c,
		"resetLlmAgent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgent) ResetModelSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetModelSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgent) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgent) ResetRemoteDialogflowAgent() {
	_jsii_.InvokeVoid(
		c,
		"resetRemoteDialogflowAgent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgent) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgent) ResetTools() {
	_jsii_.InvokeVoid(
		c,
		"resetTools",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgent) ResetToolsets() {
	_jsii_.InvokeVoid(
		c,
		"resetToolsets",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgent) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgent) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgent) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgent) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgent) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

package googlecesagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesagent/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_agent google_ces_agent}.
type GoogleCesAgent interface {
	cdktn.TerraformResource
	AfterAgentCallbacks() GoogleCesAgentAfterAgentCallbacksList
	AfterAgentCallbacksInput() interface{}
	AfterModelCallbacks() GoogleCesAgentAfterModelCallbacksList
	AfterModelCallbacksInput() interface{}
	AfterToolCallbacks() GoogleCesAgentAfterToolCallbacksList
	AfterToolCallbacksInput() interface{}
	AgentId() *string
	SetAgentId(val *string)
	AgentIdInput() *string
	App() *string
	SetApp(val *string)
	AppInput() *string
	BeforeAgentCallbacks() GoogleCesAgentBeforeAgentCallbacksList
	BeforeAgentCallbacksInput() interface{}
	BeforeModelCallbacks() GoogleCesAgentBeforeModelCallbacksList
	BeforeModelCallbacksInput() interface{}
	BeforeToolCallbacks() GoogleCesAgentBeforeToolCallbacksList
	BeforeToolCallbacksInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ChildAgents() *[]*string
	SetChildAgents(val *[]*string)
	ChildAgentsInput() *[]*string
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
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
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
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LlmAgent() GoogleCesAgentLlmAgentOutputReference
	LlmAgentInput() *GoogleCesAgentLlmAgent
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ModelSettings() GoogleCesAgentModelSettingsOutputReference
	ModelSettingsInput() *GoogleCesAgentModelSettings
	Name() *string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
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
	RemoteDialogflowAgent() GoogleCesAgentRemoteDialogflowAgentOutputReference
	RemoteDialogflowAgentInput() *GoogleCesAgentRemoteDialogflowAgent
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleCesAgentTimeoutsOutputReference
	TimeoutsInput() interface{}
	Tools() *[]*string
	SetTools(val *[]*string)
	Toolsets() GoogleCesAgentToolsetsList
	ToolsetsInput() interface{}
	ToolsInput() *[]*string
	UpdateTime() *string
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
	PutAfterAgentCallbacks(value interface{})
	PutAfterModelCallbacks(value interface{})
	PutAfterToolCallbacks(value interface{})
	PutBeforeAgentCallbacks(value interface{})
	PutBeforeModelCallbacks(value interface{})
	PutBeforeToolCallbacks(value interface{})
	PutLlmAgent(value *GoogleCesAgentLlmAgent)
	PutModelSettings(value *GoogleCesAgentModelSettings)
	PutRemoteDialogflowAgent(value *GoogleCesAgentRemoteDialogflowAgent)
	PutTimeouts(value *GoogleCesAgentTimeouts)
	PutToolsets(value interface{})
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

// The jsii proxy struct for GoogleCesAgent
type jsiiProxy_GoogleCesAgent struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleCesAgent) AfterAgentCallbacks() GoogleCesAgentAfterAgentCallbacksList {
	var returns GoogleCesAgentAfterAgentCallbacksList
	_jsii_.Get(
		j,
		"afterAgentCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) AfterAgentCallbacksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"afterAgentCallbacksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) AfterModelCallbacks() GoogleCesAgentAfterModelCallbacksList {
	var returns GoogleCesAgentAfterModelCallbacksList
	_jsii_.Get(
		j,
		"afterModelCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) AfterModelCallbacksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"afterModelCallbacksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) AfterToolCallbacks() GoogleCesAgentAfterToolCallbacksList {
	var returns GoogleCesAgentAfterToolCallbacksList
	_jsii_.Get(
		j,
		"afterToolCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) AfterToolCallbacksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"afterToolCallbacksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) AgentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) AgentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) App() *string {
	var returns *string
	_jsii_.Get(
		j,
		"app",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) AppInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) BeforeAgentCallbacks() GoogleCesAgentBeforeAgentCallbacksList {
	var returns GoogleCesAgentBeforeAgentCallbacksList
	_jsii_.Get(
		j,
		"beforeAgentCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) BeforeAgentCallbacksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"beforeAgentCallbacksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) BeforeModelCallbacks() GoogleCesAgentBeforeModelCallbacksList {
	var returns GoogleCesAgentBeforeModelCallbacksList
	_jsii_.Get(
		j,
		"beforeModelCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) BeforeModelCallbacksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"beforeModelCallbacksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) BeforeToolCallbacks() GoogleCesAgentBeforeToolCallbacksList {
	var returns GoogleCesAgentBeforeToolCallbacksList
	_jsii_.Get(
		j,
		"beforeToolCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) BeforeToolCallbacksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"beforeToolCallbacksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) ChildAgents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"childAgents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) ChildAgentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"childAgentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) GeneratedSummary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generatedSummary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) Guardrails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guardrails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) GuardrailsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guardrailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) Instruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) InstructionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instructionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) LlmAgent() GoogleCesAgentLlmAgentOutputReference {
	var returns GoogleCesAgentLlmAgentOutputReference
	_jsii_.Get(
		j,
		"llmAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) LlmAgentInput() *GoogleCesAgentLlmAgent {
	var returns *GoogleCesAgentLlmAgent
	_jsii_.Get(
		j,
		"llmAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) ModelSettings() GoogleCesAgentModelSettingsOutputReference {
	var returns GoogleCesAgentModelSettingsOutputReference
	_jsii_.Get(
		j,
		"modelSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) ModelSettingsInput() *GoogleCesAgentModelSettings {
	var returns *GoogleCesAgentModelSettings
	_jsii_.Get(
		j,
		"modelSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) RemoteDialogflowAgent() GoogleCesAgentRemoteDialogflowAgentOutputReference {
	var returns GoogleCesAgentRemoteDialogflowAgentOutputReference
	_jsii_.Get(
		j,
		"remoteDialogflowAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) RemoteDialogflowAgentInput() *GoogleCesAgentRemoteDialogflowAgent {
	var returns *GoogleCesAgentRemoteDialogflowAgent
	_jsii_.Get(
		j,
		"remoteDialogflowAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) Timeouts() GoogleCesAgentTimeoutsOutputReference {
	var returns GoogleCesAgentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) Tools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) Toolsets() GoogleCesAgentToolsetsList {
	var returns GoogleCesAgentToolsetsList
	_jsii_.Get(
		j,
		"toolsets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) ToolsetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"toolsetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) ToolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"toolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgent) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_agent google_ces_agent} Resource.
func NewGoogleCesAgent(scope constructs.Construct, id *string, config *GoogleCesAgentConfig) GoogleCesAgent {
	_init_.Initialize()

	if err := validateNewGoogleCesAgentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesAgent{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesAgent.GoogleCesAgent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_agent google_ces_agent} Resource.
func NewGoogleCesAgent_Override(g GoogleCesAgent, scope constructs.Construct, id *string, config *GoogleCesAgentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesAgent.GoogleCesAgent",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleCesAgent)SetAgentId(val *string) {
	if err := j.validateSetAgentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentId",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgent)SetApp(val *string) {
	if err := j.validateSetAppParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"app",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgent)SetChildAgents(val *[]*string) {
	if err := j.validateSetChildAgentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"childAgents",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgent)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgent)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgent)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgent)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgent)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgent)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgent)SetGuardrails(val *[]*string) {
	if err := j.validateSetGuardrailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guardrails",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgent)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgent)SetInstruction(val *string) {
	if err := j.validateSetInstructionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instruction",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgent)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgent)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgent)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgent)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgent)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgent)SetTools(val *[]*string) {
	if err := j.validateSetToolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tools",
		val,
	)
}

// Generates CDKTN code for importing a GoogleCesAgent resource upon running "cdktn plan <stack-name>".
func GoogleCesAgent_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleCesAgent_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleCesAgent.GoogleCesAgent",
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
func GoogleCesAgent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCesAgent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleCesAgent.GoogleCesAgent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCesAgent_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCesAgent_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleCesAgent.GoogleCesAgent",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCesAgent_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCesAgent_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleCesAgent.GoogleCesAgent",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleCesAgent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleCesAgent.GoogleCesAgent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleCesAgent) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleCesAgent) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleCesAgent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesAgent) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAgent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesAgent) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesAgent) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesAgent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesAgent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesAgent) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesAgent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesAgent) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAgent) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleCesAgent) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAgent) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := g.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAgent) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleCesAgent) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleCesAgent) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleCesAgent) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleCesAgent) PutAfterAgentCallbacks(value interface{}) {
	if err := g.validatePutAfterAgentCallbacksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAfterAgentCallbacks",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAgent) PutAfterModelCallbacks(value interface{}) {
	if err := g.validatePutAfterModelCallbacksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAfterModelCallbacks",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAgent) PutAfterToolCallbacks(value interface{}) {
	if err := g.validatePutAfterToolCallbacksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAfterToolCallbacks",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAgent) PutBeforeAgentCallbacks(value interface{}) {
	if err := g.validatePutBeforeAgentCallbacksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBeforeAgentCallbacks",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAgent) PutBeforeModelCallbacks(value interface{}) {
	if err := g.validatePutBeforeModelCallbacksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBeforeModelCallbacks",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAgent) PutBeforeToolCallbacks(value interface{}) {
	if err := g.validatePutBeforeToolCallbacksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBeforeToolCallbacks",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAgent) PutLlmAgent(value *GoogleCesAgentLlmAgent) {
	if err := g.validatePutLlmAgentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLlmAgent",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAgent) PutModelSettings(value *GoogleCesAgentModelSettings) {
	if err := g.validatePutModelSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putModelSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAgent) PutRemoteDialogflowAgent(value *GoogleCesAgentRemoteDialogflowAgent) {
	if err := g.validatePutRemoteDialogflowAgentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRemoteDialogflowAgent",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAgent) PutTimeouts(value *GoogleCesAgentTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAgent) PutToolsets(value interface{}) {
	if err := g.validatePutToolsetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putToolsets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAgent) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleCesAgent) ResetAfterAgentCallbacks() {
	_jsii_.InvokeVoid(
		g,
		"resetAfterAgentCallbacks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgent) ResetAfterModelCallbacks() {
	_jsii_.InvokeVoid(
		g,
		"resetAfterModelCallbacks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgent) ResetAfterToolCallbacks() {
	_jsii_.InvokeVoid(
		g,
		"resetAfterToolCallbacks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgent) ResetAgentId() {
	_jsii_.InvokeVoid(
		g,
		"resetAgentId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgent) ResetBeforeAgentCallbacks() {
	_jsii_.InvokeVoid(
		g,
		"resetBeforeAgentCallbacks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgent) ResetBeforeModelCallbacks() {
	_jsii_.InvokeVoid(
		g,
		"resetBeforeModelCallbacks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgent) ResetBeforeToolCallbacks() {
	_jsii_.InvokeVoid(
		g,
		"resetBeforeToolCallbacks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgent) ResetChildAgents() {
	_jsii_.InvokeVoid(
		g,
		"resetChildAgents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgent) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgent) ResetGuardrails() {
	_jsii_.InvokeVoid(
		g,
		"resetGuardrails",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgent) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgent) ResetInstruction() {
	_jsii_.InvokeVoid(
		g,
		"resetInstruction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgent) ResetLlmAgent() {
	_jsii_.InvokeVoid(
		g,
		"resetLlmAgent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgent) ResetModelSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetModelSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgent) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgent) ResetRemoteDialogflowAgent() {
	_jsii_.InvokeVoid(
		g,
		"resetRemoteDialogflowAgent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgent) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgent) ResetTools() {
	_jsii_.InvokeVoid(
		g,
		"resetTools",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgent) ResetToolsets() {
	_jsii_.InvokeVoid(
		g,
		"resetToolsets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAgent) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAgent) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAgent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAgent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAgent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAgent) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		g,
		"with",
		args,
		&returns,
	)

	return returns
}


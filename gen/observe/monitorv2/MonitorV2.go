package monitorv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/observe/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/observe/monitorv2/internal"
)

// Represents a {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2 observe_monitor_v2}.
type MonitorV2 interface {
	cdktn.TerraformResource
	Actions() MonitorV2ActionsList
	ActionsInput() interface{}
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
	CustomVariables() *string
	SetCustomVariables(val *string)
	CustomVariablesInput() *string
	DataStabilizationDelay() *string
	SetDataStabilizationDelay(val *string)
	DataStabilizationDelayInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Groupings() MonitorV2GroupingsList
	GroupingsInput() interface{}
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
	LookbackTime() *string
	SetLookbackTime(val *string)
	LookbackTimeInput() *string
	MaxAlertsPerHour() *float64
	SetMaxAlertsPerHour(val *float64)
	MaxAlertsPerHourInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NoDataRules() MonitorV2NoDataRulesOutputReference
	NoDataRulesInput() *MonitorV2NoDataRules
	// The tree node.
	Node() constructs.Node
	Oid() *string
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
	RuleKind() *string
	SetRuleKind(val *string)
	RuleKindInput() *string
	Rules() MonitorV2RulesList
	RulesInput() interface{}
	Scheduling() MonitorV2SchedulingOutputReference
	SchedulingInput() *MonitorV2Scheduling
	Stage() MonitorV2StageList
	StageInput() interface{}
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
	PutActions(value interface{})
	PutGroupings(value interface{})
	PutNoDataRules(value *MonitorV2NoDataRules)
	PutRules(value interface{})
	PutScheduling(value *MonitorV2Scheduling)
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
	ResetActions()
	ResetCustomVariables()
	ResetDataStabilizationDelay()
	ResetDescription()
	ResetDisabled()
	ResetGroupings()
	ResetIconUrl()
	ResetId()
	ResetLookbackTime()
	ResetMaxAlertsPerHour()
	ResetNoDataRules()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetScheduling()
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

// The jsii proxy struct for MonitorV2
type jsiiProxy_MonitorV2 struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_MonitorV2) Actions() MonitorV2ActionsList {
	var returns MonitorV2ActionsList
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) ActionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) CustomVariables() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) CustomVariablesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) DataStabilizationDelay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataStabilizationDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) DataStabilizationDelayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataStabilizationDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) Groupings() MonitorV2GroupingsList {
	var returns MonitorV2GroupingsList
	_jsii_.Get(
		j,
		"groupings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) GroupingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) IconUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iconUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) IconUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iconUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) Inputs() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"inputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) InputsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"inputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) LookbackTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookbackTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) LookbackTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookbackTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) MaxAlertsPerHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAlertsPerHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) MaxAlertsPerHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAlertsPerHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) NoDataRules() MonitorV2NoDataRulesOutputReference {
	var returns MonitorV2NoDataRulesOutputReference
	_jsii_.Get(
		j,
		"noDataRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) NoDataRulesInput() *MonitorV2NoDataRules {
	var returns *MonitorV2NoDataRules
	_jsii_.Get(
		j,
		"noDataRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) Oid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) RuleKind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleKind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) RuleKindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleKindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) Rules() MonitorV2RulesList {
	var returns MonitorV2RulesList
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) RulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) Scheduling() MonitorV2SchedulingOutputReference {
	var returns MonitorV2SchedulingOutputReference
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) SchedulingInput() *MonitorV2Scheduling {
	var returns *MonitorV2Scheduling
	_jsii_.Get(
		j,
		"schedulingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) Stage() MonitorV2StageList {
	var returns MonitorV2StageList
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) StageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) Workspace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2) WorkspaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2 observe_monitor_v2} Resource.
func NewMonitorV2(scope constructs.Construct, id *string, config *MonitorV2Config) MonitorV2 {
	_init_.Initialize()

	if err := validateNewMonitorV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorV2{}

	_jsii_.Create(
		"@cdktn/provider-observe.monitorV2.MonitorV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2 observe_monitor_v2} Resource.
func NewMonitorV2_Override(m MonitorV2, scope constructs.Construct, id *string, config *MonitorV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-observe.monitorV2.MonitorV2",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MonitorV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MonitorV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MonitorV2)SetCustomVariables(val *string) {
	if err := j.validateSetCustomVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customVariables",
		val,
	)
}

func (j *jsiiProxy_MonitorV2)SetDataStabilizationDelay(val *string) {
	if err := j.validateSetDataStabilizationDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataStabilizationDelay",
		val,
	)
}

func (j *jsiiProxy_MonitorV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MonitorV2)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MonitorV2)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_MonitorV2)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MonitorV2)SetIconUrl(val *string) {
	if err := j.validateSetIconUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iconUrl",
		val,
	)
}

func (j *jsiiProxy_MonitorV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MonitorV2)SetInputs(val *map[string]*string) {
	if err := j.validateSetInputsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputs",
		val,
	)
}

func (j *jsiiProxy_MonitorV2)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MonitorV2)SetLookbackTime(val *string) {
	if err := j.validateSetLookbackTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lookbackTime",
		val,
	)
}

func (j *jsiiProxy_MonitorV2)SetMaxAlertsPerHour(val *float64) {
	if err := j.validateSetMaxAlertsPerHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAlertsPerHour",
		val,
	)
}

func (j *jsiiProxy_MonitorV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MonitorV2)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MonitorV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MonitorV2)SetRuleKind(val *string) {
	if err := j.validateSetRuleKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleKind",
		val,
	)
}

func (j *jsiiProxy_MonitorV2)SetWorkspace(val *string) {
	if err := j.validateSetWorkspaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspace",
		val,
	)
}

// Generates CDKTN code for importing a MonitorV2 resource upon running "cdktn plan <stack-name>".
func MonitorV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateMonitorV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-observe.monitorV2.MonitorV2",
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
func MonitorV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitorV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-observe.monitorV2.MonitorV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MonitorV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitorV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-observe.monitorV2.MonitorV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MonitorV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitorV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-observe.monitorV2.MonitorV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MonitorV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-observe.monitorV2.MonitorV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MonitorV2) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MonitorV2) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MonitorV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MonitorV2) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := m.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MonitorV2) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MonitorV2) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MonitorV2) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MonitorV2) PutActions(value interface{}) {
	if err := m.validatePutActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putActions",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorV2) PutGroupings(value interface{}) {
	if err := m.validatePutGroupingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putGroupings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorV2) PutNoDataRules(value *MonitorV2NoDataRules) {
	if err := m.validatePutNoDataRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putNoDataRules",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorV2) PutRules(value interface{}) {
	if err := m.validatePutRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRules",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorV2) PutScheduling(value *MonitorV2Scheduling) {
	if err := m.validatePutSchedulingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putScheduling",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorV2) PutStage(value interface{}) {
	if err := m.validatePutStageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putStage",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorV2) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := m.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (m *jsiiProxy_MonitorV2) ResetActions() {
	_jsii_.InvokeVoid(
		m,
		"resetActions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2) ResetCustomVariables() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomVariables",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2) ResetDataStabilizationDelay() {
	_jsii_.InvokeVoid(
		m,
		"resetDataStabilizationDelay",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2) ResetDisabled() {
	_jsii_.InvokeVoid(
		m,
		"resetDisabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2) ResetGroupings() {
	_jsii_.InvokeVoid(
		m,
		"resetGroupings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2) ResetIconUrl() {
	_jsii_.InvokeVoid(
		m,
		"resetIconUrl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2) ResetLookbackTime() {
	_jsii_.InvokeVoid(
		m,
		"resetLookbackTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2) ResetMaxAlertsPerHour() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxAlertsPerHour",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2) ResetNoDataRules() {
	_jsii_.InvokeVoid(
		m,
		"resetNoDataRules",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2) ResetScheduling() {
	_jsii_.InvokeVoid(
		m,
		"resetScheduling",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		m,
		"with",
		args,
		&returns,
	)

	return returns
}


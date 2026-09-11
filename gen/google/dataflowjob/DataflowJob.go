package dataflowjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/dataflowjob/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/dataflow_job google_dataflow_job}.
type DataflowJob interface {
	cdktn.TerraformResource
	AdditionalExperiments() *[]*string
	SetAdditionalExperiments(val *[]*string)
	AdditionalExperimentsInput() *[]*string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EffectiveLabels() cdktn.StringMap
	EnableStreamingEngine() interface{}
	SetEnableStreamingEngine(val interface{})
	EnableStreamingEngineInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IpConfiguration() *string
	SetIpConfiguration(val *string)
	IpConfigurationInput() *string
	JobId() *string
	KmsKeyName() *string
	SetKmsKeyName(val *string)
	KmsKeyNameInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	MaxWorkers() *float64
	SetMaxWorkers(val *float64)
	MaxWorkersInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	// The tree node.
	Node() constructs.Node
	OnDelete() *string
	SetOnDelete(val *string)
	OnDeleteInput() *string
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ServiceAccountEmail() *string
	SetServiceAccountEmail(val *string)
	ServiceAccountEmailInput() *string
	SkipWaitOnJobTermination() interface{}
	SetSkipWaitOnJobTermination(val interface{})
	SkipWaitOnJobTerminationInput() interface{}
	State() *string
	Subnetwork() *string
	SetSubnetwork(val *string)
	SubnetworkInput() *string
	TempGcsLocation() *string
	SetTempGcsLocation(val *string)
	TempGcsLocationInput() *string
	TemplateGcsPath() *string
	SetTemplateGcsPath(val *string)
	TemplateGcsPathInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktn.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataflowJobTimeoutsOutputReference
	TimeoutsInput() interface{}
	TransformNameMapping() *map[string]*string
	SetTransformNameMapping(val *map[string]*string)
	TransformNameMappingInput() *map[string]*string
	Type() *string
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
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
	PutTimeouts(value *DataflowJobTimeouts)
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
	ResetAdditionalExperiments()
	ResetEnableStreamingEngine()
	ResetId()
	ResetIpConfiguration()
	ResetKmsKeyName()
	ResetLabels()
	ResetMachineType()
	ResetMaxWorkers()
	ResetNetwork()
	ResetOnDelete()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
	ResetProject()
	ResetRegion()
	ResetServiceAccountEmail()
	ResetSkipWaitOnJobTermination()
	ResetSubnetwork()
	ResetTimeouts()
	ResetTransformNameMapping()
	ResetZone()
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

// The jsii proxy struct for DataflowJob
type jsiiProxy_DataflowJob struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_DataflowJob) AdditionalExperiments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalExperiments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) AdditionalExperimentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalExperimentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) EnableStreamingEngine() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStreamingEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) EnableStreamingEngineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStreamingEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) IpConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) IpConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) JobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) KmsKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) KmsKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) MaxWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) MaxWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) OnDelete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) OnDeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) ServiceAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) ServiceAccountEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) SkipWaitOnJobTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipWaitOnJobTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) SkipWaitOnJobTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipWaitOnJobTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) SubnetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) TempGcsLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tempGcsLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) TempGcsLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tempGcsLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) TemplateGcsPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateGcsPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) TemplateGcsPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateGcsPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) Timeouts() DataflowJobTimeoutsOutputReference {
	var returns DataflowJobTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) TransformNameMapping() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"transformNameMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) TransformNameMappingInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"transformNameMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowJob) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/dataflow_job google_dataflow_job} Resource.
func NewDataflowJob(scope constructs.Construct, id *string, config *DataflowJobConfig) DataflowJob {
	_init_.Initialize()

	if err := validateNewDataflowJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataflowJob{}

	_jsii_.Create(
		"@cdktn/provider-google.dataflowJob.DataflowJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/dataflow_job google_dataflow_job} Resource.
func NewDataflowJob_Override(d DataflowJob, scope constructs.Construct, id *string, config *DataflowJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataflowJob.DataflowJob",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataflowJob)SetAdditionalExperiments(val *[]*string) {
	if err := j.validateSetAdditionalExperimentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalExperiments",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetEnableStreamingEngine(val interface{}) {
	if err := j.validateSetEnableStreamingEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableStreamingEngine",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetIpConfiguration(val *string) {
	if err := j.validateSetIpConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipConfiguration",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetKmsKeyName(val *string) {
	if err := j.validateSetKmsKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyName",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetMaxWorkers(val *float64) {
	if err := j.validateSetMaxWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWorkers",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetOnDelete(val *string) {
	if err := j.validateSetOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDelete",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetServiceAccountEmail(val *string) {
	if err := j.validateSetServiceAccountEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountEmail",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetSkipWaitOnJobTermination(val interface{}) {
	if err := j.validateSetSkipWaitOnJobTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipWaitOnJobTermination",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetSubnetwork(val *string) {
	if err := j.validateSetSubnetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetwork",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetTempGcsLocation(val *string) {
	if err := j.validateSetTempGcsLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tempGcsLocation",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetTemplateGcsPath(val *string) {
	if err := j.validateSetTemplateGcsPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateGcsPath",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetTransformNameMapping(val *map[string]*string) {
	if err := j.validateSetTransformNameMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transformNameMapping",
		val,
	)
}

func (j *jsiiProxy_DataflowJob)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

// Generates CDKTN code for importing a DataflowJob resource upon running "cdktn plan <stack-name>".
func DataflowJob_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataflowJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dataflowJob.DataflowJob",
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
func DataflowJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataflowJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dataflowJob.DataflowJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataflowJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataflowJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dataflowJob.DataflowJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataflowJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataflowJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dataflowJob.DataflowJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataflowJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.dataflowJob.DataflowJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataflowJob) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DataflowJob) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataflowJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataflowJob) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataflowJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataflowJob) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataflowJob) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataflowJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataflowJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataflowJob) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataflowJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataflowJob) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowJob) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DataflowJob) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataflowJob) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (d *jsiiProxy_DataflowJob) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DataflowJob) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DataflowJob) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DataflowJob) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataflowJob) PutTimeouts(value *DataflowJobTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataflowJob) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataflowJob) ResetAdditionalExperiments() {
	_jsii_.InvokeVoid(
		d,
		"resetAdditionalExperiments",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowJob) ResetEnableStreamingEngine() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableStreamingEngine",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowJob) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowJob) ResetIpConfiguration() {
	_jsii_.InvokeVoid(
		d,
		"resetIpConfiguration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowJob) ResetKmsKeyName() {
	_jsii_.InvokeVoid(
		d,
		"resetKmsKeyName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowJob) ResetLabels() {
	_jsii_.InvokeVoid(
		d,
		"resetLabels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowJob) ResetMachineType() {
	_jsii_.InvokeVoid(
		d,
		"resetMachineType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowJob) ResetMaxWorkers() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxWorkers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowJob) ResetNetwork() {
	_jsii_.InvokeVoid(
		d,
		"resetNetwork",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowJob) ResetOnDelete() {
	_jsii_.InvokeVoid(
		d,
		"resetOnDelete",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowJob) ResetParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowJob) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowJob) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowJob) ResetServiceAccountEmail() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAccountEmail",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowJob) ResetSkipWaitOnJobTermination() {
	_jsii_.InvokeVoid(
		d,
		"resetSkipWaitOnJobTermination",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowJob) ResetSubnetwork() {
	_jsii_.InvokeVoid(
		d,
		"resetSubnetwork",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowJob) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowJob) ResetTransformNameMapping() {
	_jsii_.InvokeVoid(
		d,
		"resetTransformNameMapping",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowJob) ResetZone() {
	_jsii_.InvokeVoid(
		d,
		"resetZone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowJob) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowJob) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowJob) With(mixins ...constructs.IMixin) constructs.IConstruct {
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


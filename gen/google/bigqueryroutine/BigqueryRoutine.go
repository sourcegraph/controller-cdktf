package bigqueryroutine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/bigqueryroutine/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/bigquery_routine google_bigquery_routine}.
type BigqueryRoutine interface {
	cdktn.TerraformResource
	Arguments() BigqueryRoutineArgumentsList
	ArgumentsInput() interface{}
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
	CreationTime() *float64
	DataGovernanceType() *string
	SetDataGovernanceType(val *string)
	DataGovernanceTypeInput() *string
	DatasetId() *string
	SetDatasetId(val *string)
	DatasetIdInput() *string
	DefinitionBody() *string
	SetDefinitionBody(val *string)
	DefinitionBodyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DeterminismLevel() *string
	SetDeterminismLevel(val *string)
	DeterminismLevelInput() *string
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
	ImportedLibraries() *[]*string
	SetImportedLibraries(val *[]*string)
	ImportedLibrariesInput() *[]*string
	Language() *string
	SetLanguage(val *string)
	LanguageInput() *string
	LastModifiedTime() *float64
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
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
	RemoteFunctionOptions() BigqueryRoutineRemoteFunctionOptionsOutputReference
	RemoteFunctionOptionsInput() *BigqueryRoutineRemoteFunctionOptions
	ReturnTableType() *string
	SetReturnTableType(val *string)
	ReturnTableTypeInput() *string
	ReturnType() *string
	SetReturnType(val *string)
	ReturnTypeInput() *string
	RoutineId() *string
	SetRoutineId(val *string)
	RoutineIdInput() *string
	RoutineType() *string
	SetRoutineType(val *string)
	RoutineTypeInput() *string
	SecurityMode() *string
	SetSecurityMode(val *string)
	SecurityModeInput() *string
	SparkOptions() BigqueryRoutineSparkOptionsOutputReference
	SparkOptionsInput() *BigqueryRoutineSparkOptions
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() BigqueryRoutineTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutArguments(value interface{})
	PutRemoteFunctionOptions(value *BigqueryRoutineRemoteFunctionOptions)
	PutSparkOptions(value *BigqueryRoutineSparkOptions)
	PutTimeouts(value *BigqueryRoutineTimeouts)
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
	ResetArguments()
	ResetDataGovernanceType()
	ResetDescription()
	ResetDeterminismLevel()
	ResetId()
	ResetImportedLibraries()
	ResetLanguage()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRemoteFunctionOptions()
	ResetReturnTableType()
	ResetReturnType()
	ResetSecurityMode()
	ResetSparkOptions()
	ResetTimeouts()
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

// The jsii proxy struct for BigqueryRoutine
type jsiiProxy_BigqueryRoutine struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_BigqueryRoutine) Arguments() BigqueryRoutineArgumentsList {
	var returns BigqueryRoutineArgumentsList
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) ArgumentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"argumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) CreationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) DataGovernanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataGovernanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) DataGovernanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataGovernanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) DatasetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) DatasetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) DefinitionBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definitionBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) DefinitionBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definitionBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) DeterminismLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"determinismLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) DeterminismLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"determinismLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) ImportedLibraries() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"importedLibraries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) ImportedLibrariesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"importedLibrariesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) Language() *string {
	var returns *string
	_jsii_.Get(
		j,
		"language",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) LanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) LastModifiedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) RemoteFunctionOptions() BigqueryRoutineRemoteFunctionOptionsOutputReference {
	var returns BigqueryRoutineRemoteFunctionOptionsOutputReference
	_jsii_.Get(
		j,
		"remoteFunctionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) RemoteFunctionOptionsInput() *BigqueryRoutineRemoteFunctionOptions {
	var returns *BigqueryRoutineRemoteFunctionOptions
	_jsii_.Get(
		j,
		"remoteFunctionOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) ReturnTableType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnTableType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) ReturnTableTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnTableTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) ReturnType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) ReturnTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) RoutineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) RoutineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routineIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) RoutineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) RoutineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) SecurityMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) SecurityModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) SparkOptions() BigqueryRoutineSparkOptionsOutputReference {
	var returns BigqueryRoutineSparkOptionsOutputReference
	_jsii_.Get(
		j,
		"sparkOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) SparkOptionsInput() *BigqueryRoutineSparkOptions {
	var returns *BigqueryRoutineSparkOptions
	_jsii_.Get(
		j,
		"sparkOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) Timeouts() BigqueryRoutineTimeoutsOutputReference {
	var returns BigqueryRoutineTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutine) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/bigquery_routine google_bigquery_routine} Resource.
func NewBigqueryRoutine(scope constructs.Construct, id *string, config *BigqueryRoutineConfig) BigqueryRoutine {
	_init_.Initialize()

	if err := validateNewBigqueryRoutineParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigqueryRoutine{}

	_jsii_.Create(
		"@cdktn/provider-google.bigqueryRoutine.BigqueryRoutine",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/bigquery_routine google_bigquery_routine} Resource.
func NewBigqueryRoutine_Override(b BigqueryRoutine, scope constructs.Construct, id *string, config *BigqueryRoutineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.bigqueryRoutine.BigqueryRoutine",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BigqueryRoutine)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutine)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutine)SetDataGovernanceType(val *string) {
	if err := j.validateSetDataGovernanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataGovernanceType",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutine)SetDatasetId(val *string) {
	if err := j.validateSetDatasetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetId",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutine)SetDefinitionBody(val *string) {
	if err := j.validateSetDefinitionBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"definitionBody",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutine)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutine)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutine)SetDeterminismLevel(val *string) {
	if err := j.validateSetDeterminismLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"determinismLevel",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutine)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutine)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutine)SetImportedLibraries(val *[]*string) {
	if err := j.validateSetImportedLibrariesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importedLibraries",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutine)SetLanguage(val *string) {
	if err := j.validateSetLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"language",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutine)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutine)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutine)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutine)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutine)SetReturnTableType(val *string) {
	if err := j.validateSetReturnTableTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnTableType",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutine)SetReturnType(val *string) {
	if err := j.validateSetReturnTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnType",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutine)SetRoutineId(val *string) {
	if err := j.validateSetRoutineIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routineId",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutine)SetRoutineType(val *string) {
	if err := j.validateSetRoutineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routineType",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutine)SetSecurityMode(val *string) {
	if err := j.validateSetSecurityModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityMode",
		val,
	)
}

// Generates CDKTN code for importing a BigqueryRoutine resource upon running "cdktn plan <stack-name>".
func BigqueryRoutine_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateBigqueryRoutine_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.bigqueryRoutine.BigqueryRoutine",
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
func BigqueryRoutine_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBigqueryRoutine_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.bigqueryRoutine.BigqueryRoutine",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BigqueryRoutine_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBigqueryRoutine_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.bigqueryRoutine.BigqueryRoutine",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BigqueryRoutine_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBigqueryRoutine_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.bigqueryRoutine.BigqueryRoutine",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BigqueryRoutine_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.bigqueryRoutine.BigqueryRoutine",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BigqueryRoutine) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BigqueryRoutine) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BigqueryRoutine) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutine) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutine) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutine) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutine) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutine) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutine) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutine) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutine) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutine) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutine) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BigqueryRoutine) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutine) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := b.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutine) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BigqueryRoutine) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BigqueryRoutine) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BigqueryRoutine) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BigqueryRoutine) PutArguments(value interface{}) {
	if err := b.validatePutArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putArguments",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryRoutine) PutRemoteFunctionOptions(value *BigqueryRoutineRemoteFunctionOptions) {
	if err := b.validatePutRemoteFunctionOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRemoteFunctionOptions",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryRoutine) PutSparkOptions(value *BigqueryRoutineSparkOptions) {
	if err := b.validatePutSparkOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSparkOptions",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryRoutine) PutTimeouts(value *BigqueryRoutineTimeouts) {
	if err := b.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryRoutine) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := b.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (b *jsiiProxy_BigqueryRoutine) ResetArguments() {
	_jsii_.InvokeVoid(
		b,
		"resetArguments",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutine) ResetDataGovernanceType() {
	_jsii_.InvokeVoid(
		b,
		"resetDataGovernanceType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutine) ResetDescription() {
	_jsii_.InvokeVoid(
		b,
		"resetDescription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutine) ResetDeterminismLevel() {
	_jsii_.InvokeVoid(
		b,
		"resetDeterminismLevel",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutine) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutine) ResetImportedLibraries() {
	_jsii_.InvokeVoid(
		b,
		"resetImportedLibraries",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutine) ResetLanguage() {
	_jsii_.InvokeVoid(
		b,
		"resetLanguage",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutine) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutine) ResetProject() {
	_jsii_.InvokeVoid(
		b,
		"resetProject",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutine) ResetRemoteFunctionOptions() {
	_jsii_.InvokeVoid(
		b,
		"resetRemoteFunctionOptions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutine) ResetReturnTableType() {
	_jsii_.InvokeVoid(
		b,
		"resetReturnTableType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutine) ResetReturnType() {
	_jsii_.InvokeVoid(
		b,
		"resetReturnType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutine) ResetSecurityMode() {
	_jsii_.InvokeVoid(
		b,
		"resetSecurityMode",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutine) ResetSparkOptions() {
	_jsii_.InvokeVoid(
		b,
		"resetSparkOptions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutine) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutine) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutine) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutine) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutine) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutine) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutine) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutine) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		b,
		"with",
		args,
		&returns,
	)

	return returns
}


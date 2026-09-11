package dataprocgdcsparkapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/dataprocgdcsparkapplication/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/dataproc_gdc_spark_application google_dataproc_gdc_spark_application}.
type DataprocGdcSparkApplication interface {
	cdktn.TerraformResource
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	ApplicationEnvironment() *string
	SetApplicationEnvironment(val *string)
	ApplicationEnvironmentInput() *string
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
	CreateTime() *string
	DependencyImages() *[]*string
	SetDependencyImages(val *[]*string)
	DependencyImagesInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EffectiveAnnotations() cdktn.StringMap
	EffectiveLabels() cdktn.StringMap
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
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MonitoringEndpoint() *string
	Name() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	OutputUri() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	Properties() *map[string]*string
	SetProperties(val *map[string]*string)
	PropertiesInput() *map[string]*string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PysparkApplicationConfig() DataprocGdcSparkApplicationPysparkApplicationConfigOutputReference
	PysparkApplicationConfigInput() *DataprocGdcSparkApplicationPysparkApplicationConfig
	// Experimental.
	RawOverrides() interface{}
	Reconciling() cdktn.IResolvable
	Serviceinstance() *string
	SetServiceinstance(val *string)
	ServiceinstanceInput() *string
	SparkApplicationConfig() DataprocGdcSparkApplicationSparkApplicationConfigOutputReference
	SparkApplicationConfigInput() *DataprocGdcSparkApplicationSparkApplicationConfig
	SparkApplicationId() *string
	SetSparkApplicationId(val *string)
	SparkApplicationIdInput() *string
	SparkRApplicationConfig() DataprocGdcSparkApplicationSparkRApplicationConfigOutputReference
	SparkRApplicationConfigInput() *DataprocGdcSparkApplicationSparkRApplicationConfig
	SparkSqlApplicationConfig() DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference
	SparkSqlApplicationConfigInput() *DataprocGdcSparkApplicationSparkSqlApplicationConfig
	State() *string
	StateMessage() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktn.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataprocGdcSparkApplicationTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
	UpdateTime() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutPysparkApplicationConfig(value *DataprocGdcSparkApplicationPysparkApplicationConfig)
	PutSparkApplicationConfig(value *DataprocGdcSparkApplicationSparkApplicationConfig)
	PutSparkRApplicationConfig(value *DataprocGdcSparkApplicationSparkRApplicationConfig)
	PutSparkSqlApplicationConfig(value *DataprocGdcSparkApplicationSparkSqlApplicationConfig)
	PutTimeouts(value *DataprocGdcSparkApplicationTimeouts)
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
	ResetAnnotations()
	ResetApplicationEnvironment()
	ResetDependencyImages()
	ResetDisplayName()
	ResetId()
	ResetLabels()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetProperties()
	ResetPysparkApplicationConfig()
	ResetSparkApplicationConfig()
	ResetSparkRApplicationConfig()
	ResetSparkSqlApplicationConfig()
	ResetTimeouts()
	ResetVersion()
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

// The jsii proxy struct for DataprocGdcSparkApplication
type jsiiProxy_DataprocGdcSparkApplication struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_DataprocGdcSparkApplication) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) ApplicationEnvironment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) ApplicationEnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationEnvironmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) DependencyImages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependencyImages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) DependencyImagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependencyImagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) EffectiveAnnotations() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) MonitoringEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) OutputUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) Properties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) PropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) PysparkApplicationConfig() DataprocGdcSparkApplicationPysparkApplicationConfigOutputReference {
	var returns DataprocGdcSparkApplicationPysparkApplicationConfigOutputReference
	_jsii_.Get(
		j,
		"pysparkApplicationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) PysparkApplicationConfigInput() *DataprocGdcSparkApplicationPysparkApplicationConfig {
	var returns *DataprocGdcSparkApplicationPysparkApplicationConfig
	_jsii_.Get(
		j,
		"pysparkApplicationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) Reconciling() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) Serviceinstance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceinstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) ServiceinstanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceinstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) SparkApplicationConfig() DataprocGdcSparkApplicationSparkApplicationConfigOutputReference {
	var returns DataprocGdcSparkApplicationSparkApplicationConfigOutputReference
	_jsii_.Get(
		j,
		"sparkApplicationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) SparkApplicationConfigInput() *DataprocGdcSparkApplicationSparkApplicationConfig {
	var returns *DataprocGdcSparkApplicationSparkApplicationConfig
	_jsii_.Get(
		j,
		"sparkApplicationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) SparkApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkApplicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) SparkApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkApplicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) SparkRApplicationConfig() DataprocGdcSparkApplicationSparkRApplicationConfigOutputReference {
	var returns DataprocGdcSparkApplicationSparkRApplicationConfigOutputReference
	_jsii_.Get(
		j,
		"sparkRApplicationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) SparkRApplicationConfigInput() *DataprocGdcSparkApplicationSparkRApplicationConfig {
	var returns *DataprocGdcSparkApplicationSparkRApplicationConfig
	_jsii_.Get(
		j,
		"sparkRApplicationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) SparkSqlApplicationConfig() DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference {
	var returns DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference
	_jsii_.Get(
		j,
		"sparkSqlApplicationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) SparkSqlApplicationConfigInput() *DataprocGdcSparkApplicationSparkSqlApplicationConfig {
	var returns *DataprocGdcSparkApplicationSparkSqlApplicationConfig
	_jsii_.Get(
		j,
		"sparkSqlApplicationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) StateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) Timeouts() DataprocGdcSparkApplicationTimeoutsOutputReference {
	var returns DataprocGdcSparkApplicationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplication) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/dataproc_gdc_spark_application google_dataproc_gdc_spark_application} Resource.
func NewDataprocGdcSparkApplication(scope constructs.Construct, id *string, config *DataprocGdcSparkApplicationConfig) DataprocGdcSparkApplication {
	_init_.Initialize()

	if err := validateNewDataprocGdcSparkApplicationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocGdcSparkApplication{}

	_jsii_.Create(
		"@cdktn/provider-google.dataprocGdcSparkApplication.DataprocGdcSparkApplication",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/dataproc_gdc_spark_application google_dataproc_gdc_spark_application} Resource.
func NewDataprocGdcSparkApplication_Override(d DataprocGdcSparkApplication, scope constructs.Construct, id *string, config *DataprocGdcSparkApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataprocGdcSparkApplication.DataprocGdcSparkApplication",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplication)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplication)SetApplicationEnvironment(val *string) {
	if err := j.validateSetApplicationEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationEnvironment",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplication)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplication)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplication)SetDependencyImages(val *[]*string) {
	if err := j.validateSetDependencyImagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dependencyImages",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplication)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplication)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplication)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplication)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplication)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplication)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplication)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplication)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplication)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplication)SetProperties(val *map[string]*string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplication)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplication)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplication)SetServiceinstance(val *string) {
	if err := j.validateSetServiceinstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceinstance",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplication)SetSparkApplicationId(val *string) {
	if err := j.validateSetSparkApplicationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkApplicationId",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplication)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTN code for importing a DataprocGdcSparkApplication resource upon running "cdktn plan <stack-name>".
func DataprocGdcSparkApplication_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataprocGdcSparkApplication_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dataprocGdcSparkApplication.DataprocGdcSparkApplication",
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
func DataprocGdcSparkApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataprocGdcSparkApplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dataprocGdcSparkApplication.DataprocGdcSparkApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataprocGdcSparkApplication_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataprocGdcSparkApplication_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dataprocGdcSparkApplication.DataprocGdcSparkApplication",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataprocGdcSparkApplication_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataprocGdcSparkApplication_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dataprocGdcSparkApplication.DataprocGdcSparkApplication",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataprocGdcSparkApplication_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.dataprocGdcSparkApplication.DataprocGdcSparkApplication",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataprocGdcSparkApplication) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocGdcSparkApplication) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataprocGdcSparkApplication) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocGdcSparkApplication) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocGdcSparkApplication) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocGdcSparkApplication) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocGdcSparkApplication) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocGdcSparkApplication) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocGdcSparkApplication) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocGdcSparkApplication) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocGdcSparkApplication) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataprocGdcSparkApplication) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (d *jsiiProxy_DataprocGdcSparkApplication) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) PutPysparkApplicationConfig(value *DataprocGdcSparkApplicationPysparkApplicationConfig) {
	if err := d.validatePutPysparkApplicationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPysparkApplicationConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) PutSparkApplicationConfig(value *DataprocGdcSparkApplicationSparkApplicationConfig) {
	if err := d.validatePutSparkApplicationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSparkApplicationConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) PutSparkRApplicationConfig(value *DataprocGdcSparkApplicationSparkRApplicationConfig) {
	if err := d.validatePutSparkRApplicationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSparkRApplicationConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) PutSparkSqlApplicationConfig(value *DataprocGdcSparkApplicationSparkSqlApplicationConfig) {
	if err := d.validatePutSparkSqlApplicationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSparkSqlApplicationConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) PutTimeouts(value *DataprocGdcSparkApplicationTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) ResetAnnotations() {
	_jsii_.InvokeVoid(
		d,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) ResetApplicationEnvironment() {
	_jsii_.InvokeVoid(
		d,
		"resetApplicationEnvironment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) ResetDependencyImages() {
	_jsii_.InvokeVoid(
		d,
		"resetDependencyImages",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) ResetDisplayName() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) ResetLabels() {
	_jsii_.InvokeVoid(
		d,
		"resetLabels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) ResetNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) ResetProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) ResetPysparkApplicationConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetPysparkApplicationConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) ResetSparkApplicationConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkApplicationConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) ResetSparkRApplicationConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkRApplicationConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) ResetSparkSqlApplicationConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkSqlApplicationConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) ResetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplication) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocGdcSparkApplication) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocGdcSparkApplication) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocGdcSparkApplication) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocGdcSparkApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocGdcSparkApplication) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocGdcSparkApplication) With(mixins ...constructs.IMixin) constructs.IConstruct {
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


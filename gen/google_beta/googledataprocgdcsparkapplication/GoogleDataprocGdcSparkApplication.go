package googledataprocgdcsparkapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataprocgdcsparkapplication/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_dataproc_gdc_spark_application google_dataproc_gdc_spark_application}.
type GoogleDataprocGdcSparkApplication interface {
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
	PysparkApplicationConfig() GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference
	PysparkApplicationConfigInput() *GoogleDataprocGdcSparkApplicationPysparkApplicationConfig
	// Experimental.
	RawOverrides() interface{}
	Reconciling() cdktn.IResolvable
	Serviceinstance() *string
	SetServiceinstance(val *string)
	ServiceinstanceInput() *string
	SparkApplicationConfig() GoogleDataprocGdcSparkApplicationSparkApplicationConfigOutputReference
	SparkApplicationConfigInput() *GoogleDataprocGdcSparkApplicationSparkApplicationConfig
	SparkApplicationId() *string
	SetSparkApplicationId(val *string)
	SparkApplicationIdInput() *string
	SparkRApplicationConfig() GoogleDataprocGdcSparkApplicationSparkRApplicationConfigOutputReference
	SparkRApplicationConfigInput() *GoogleDataprocGdcSparkApplicationSparkRApplicationConfig
	SparkSqlApplicationConfig() GoogleDataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference
	SparkSqlApplicationConfigInput() *GoogleDataprocGdcSparkApplicationSparkSqlApplicationConfig
	State() *string
	StateMessage() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktn.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleDataprocGdcSparkApplicationTimeoutsOutputReference
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
	PutPysparkApplicationConfig(value *GoogleDataprocGdcSparkApplicationPysparkApplicationConfig)
	PutSparkApplicationConfig(value *GoogleDataprocGdcSparkApplicationSparkApplicationConfig)
	PutSparkRApplicationConfig(value *GoogleDataprocGdcSparkApplicationSparkRApplicationConfig)
	PutSparkSqlApplicationConfig(value *GoogleDataprocGdcSparkApplicationSparkSqlApplicationConfig)
	PutTimeouts(value *GoogleDataprocGdcSparkApplicationTimeouts)
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

// The jsii proxy struct for GoogleDataprocGdcSparkApplication
type jsiiProxy_GoogleDataprocGdcSparkApplication struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) ApplicationEnvironment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) ApplicationEnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationEnvironmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) DependencyImages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependencyImages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) DependencyImagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependencyImagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) EffectiveAnnotations() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) MonitoringEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) OutputUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) Properties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) PropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) PysparkApplicationConfig() GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference {
	var returns GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference
	_jsii_.Get(
		j,
		"pysparkApplicationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) PysparkApplicationConfigInput() *GoogleDataprocGdcSparkApplicationPysparkApplicationConfig {
	var returns *GoogleDataprocGdcSparkApplicationPysparkApplicationConfig
	_jsii_.Get(
		j,
		"pysparkApplicationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) Reconciling() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) Serviceinstance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceinstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) ServiceinstanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceinstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) SparkApplicationConfig() GoogleDataprocGdcSparkApplicationSparkApplicationConfigOutputReference {
	var returns GoogleDataprocGdcSparkApplicationSparkApplicationConfigOutputReference
	_jsii_.Get(
		j,
		"sparkApplicationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) SparkApplicationConfigInput() *GoogleDataprocGdcSparkApplicationSparkApplicationConfig {
	var returns *GoogleDataprocGdcSparkApplicationSparkApplicationConfig
	_jsii_.Get(
		j,
		"sparkApplicationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) SparkApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkApplicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) SparkApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkApplicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) SparkRApplicationConfig() GoogleDataprocGdcSparkApplicationSparkRApplicationConfigOutputReference {
	var returns GoogleDataprocGdcSparkApplicationSparkRApplicationConfigOutputReference
	_jsii_.Get(
		j,
		"sparkRApplicationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) SparkRApplicationConfigInput() *GoogleDataprocGdcSparkApplicationSparkRApplicationConfig {
	var returns *GoogleDataprocGdcSparkApplicationSparkRApplicationConfig
	_jsii_.Get(
		j,
		"sparkRApplicationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) SparkSqlApplicationConfig() GoogleDataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference {
	var returns GoogleDataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference
	_jsii_.Get(
		j,
		"sparkSqlApplicationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) SparkSqlApplicationConfigInput() *GoogleDataprocGdcSparkApplicationSparkSqlApplicationConfig {
	var returns *GoogleDataprocGdcSparkApplicationSparkSqlApplicationConfig
	_jsii_.Get(
		j,
		"sparkSqlApplicationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) StateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) Timeouts() GoogleDataprocGdcSparkApplicationTimeoutsOutputReference {
	var returns GoogleDataprocGdcSparkApplicationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_dataproc_gdc_spark_application google_dataproc_gdc_spark_application} Resource.
func NewGoogleDataprocGdcSparkApplication(scope constructs.Construct, id *string, config *GoogleDataprocGdcSparkApplicationConfig) GoogleDataprocGdcSparkApplication {
	_init_.Initialize()

	if err := validateNewGoogleDataprocGdcSparkApplicationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocGdcSparkApplication{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataprocGdcSparkApplication.GoogleDataprocGdcSparkApplication",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_dataproc_gdc_spark_application google_dataproc_gdc_spark_application} Resource.
func NewGoogleDataprocGdcSparkApplication_Override(g GoogleDataprocGdcSparkApplication, scope constructs.Construct, id *string, config *GoogleDataprocGdcSparkApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataprocGdcSparkApplication.GoogleDataprocGdcSparkApplication",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication)SetApplicationEnvironment(val *string) {
	if err := j.validateSetApplicationEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationEnvironment",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication)SetDependencyImages(val *[]*string) {
	if err := j.validateSetDependencyImagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dependencyImages",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication)SetProperties(val *map[string]*string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication)SetServiceinstance(val *string) {
	if err := j.validateSetServiceinstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceinstance",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication)SetSparkApplicationId(val *string) {
	if err := j.validateSetSparkApplicationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkApplicationId",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplication)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTN code for importing a GoogleDataprocGdcSparkApplication resource upon running "cdktn plan <stack-name>".
func GoogleDataprocGdcSparkApplication_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleDataprocGdcSparkApplication_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleDataprocGdcSparkApplication.GoogleDataprocGdcSparkApplication",
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
func GoogleDataprocGdcSparkApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDataprocGdcSparkApplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleDataprocGdcSparkApplication.GoogleDataprocGdcSparkApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDataprocGdcSparkApplication_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDataprocGdcSparkApplication_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleDataprocGdcSparkApplication.GoogleDataprocGdcSparkApplication",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDataprocGdcSparkApplication_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDataprocGdcSparkApplication_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleDataprocGdcSparkApplication.GoogleDataprocGdcSparkApplication",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleDataprocGdcSparkApplication_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleDataprocGdcSparkApplication.GoogleDataprocGdcSparkApplication",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) PutPysparkApplicationConfig(value *GoogleDataprocGdcSparkApplicationPysparkApplicationConfig) {
	if err := g.validatePutPysparkApplicationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPysparkApplicationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) PutSparkApplicationConfig(value *GoogleDataprocGdcSparkApplicationSparkApplicationConfig) {
	if err := g.validatePutSparkApplicationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSparkApplicationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) PutSparkRApplicationConfig(value *GoogleDataprocGdcSparkApplicationSparkRApplicationConfig) {
	if err := g.validatePutSparkRApplicationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSparkRApplicationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) PutSparkSqlApplicationConfig(value *GoogleDataprocGdcSparkApplicationSparkSqlApplicationConfig) {
	if err := g.validatePutSparkSqlApplicationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSparkSqlApplicationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) PutTimeouts(value *GoogleDataprocGdcSparkApplicationTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) ResetAnnotations() {
	_jsii_.InvokeVoid(
		g,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) ResetApplicationEnvironment() {
	_jsii_.InvokeVoid(
		g,
		"resetApplicationEnvironment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) ResetDependencyImages() {
	_jsii_.InvokeVoid(
		g,
		"resetDependencyImages",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) ResetNamespace() {
	_jsii_.InvokeVoid(
		g,
		"resetNamespace",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) ResetProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) ResetPysparkApplicationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPysparkApplicationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) ResetSparkApplicationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSparkApplicationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) ResetSparkRApplicationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSparkRApplicationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) ResetSparkSqlApplicationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSparkSqlApplicationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) ResetVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplication) With(mixins ...constructs.IMixin) constructs.IConstruct {
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


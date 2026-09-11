package googlevertexaifeatureonlinestorefeatureview

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevertexaifeatureonlinestorefeatureview/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_vertex_ai_feature_online_store_featureview google_vertex_ai_feature_online_store_featureview}.
type GoogleVertexAiFeatureOnlineStoreFeatureview interface {
	cdktn.TerraformResource
	BigQuerySource() GoogleVertexAiFeatureOnlineStoreFeatureviewBigQuerySourceOutputReference
	BigQuerySourceInput() *GoogleVertexAiFeatureOnlineStoreFeatureviewBigQuerySource
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EffectiveLabels() cdktn.StringMap
	FeatureOnlineStore() *string
	SetFeatureOnlineStore(val *string)
	FeatureOnlineStoreInput() *string
	FeatureRegistrySource() GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference
	FeatureRegistrySourceInput() *GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySource
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
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SyncConfig() GoogleVertexAiFeatureOnlineStoreFeatureviewSyncConfigOutputReference
	SyncConfigInput() *GoogleVertexAiFeatureOnlineStoreFeatureviewSyncConfig
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktn.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleVertexAiFeatureOnlineStoreFeatureviewTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
	VectorSearchConfig() GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference
	VectorSearchConfigInput() *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfig
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
	PutBigQuerySource(value *GoogleVertexAiFeatureOnlineStoreFeatureviewBigQuerySource)
	PutFeatureRegistrySource(value *GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySource)
	PutSyncConfig(value *GoogleVertexAiFeatureOnlineStoreFeatureviewSyncConfig)
	PutTimeouts(value *GoogleVertexAiFeatureOnlineStoreFeatureviewTimeouts)
	PutVectorSearchConfig(value *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfig)
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
	ResetBigQuerySource()
	ResetFeatureRegistrySource()
	ResetId()
	ResetLabels()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRegion()
	ResetSyncConfig()
	ResetTimeouts()
	ResetVectorSearchConfig()
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

// The jsii proxy struct for GoogleVertexAiFeatureOnlineStoreFeatureview
type jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) BigQuerySource() GoogleVertexAiFeatureOnlineStoreFeatureviewBigQuerySourceOutputReference {
	var returns GoogleVertexAiFeatureOnlineStoreFeatureviewBigQuerySourceOutputReference
	_jsii_.Get(
		j,
		"bigQuerySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) BigQuerySourceInput() *GoogleVertexAiFeatureOnlineStoreFeatureviewBigQuerySource {
	var returns *GoogleVertexAiFeatureOnlineStoreFeatureviewBigQuerySource
	_jsii_.Get(
		j,
		"bigQuerySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) FeatureOnlineStore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureOnlineStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) FeatureOnlineStoreInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureOnlineStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) FeatureRegistrySource() GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference {
	var returns GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference
	_jsii_.Get(
		j,
		"featureRegistrySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) FeatureRegistrySourceInput() *GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySource {
	var returns *GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySource
	_jsii_.Get(
		j,
		"featureRegistrySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) SyncConfig() GoogleVertexAiFeatureOnlineStoreFeatureviewSyncConfigOutputReference {
	var returns GoogleVertexAiFeatureOnlineStoreFeatureviewSyncConfigOutputReference
	_jsii_.Get(
		j,
		"syncConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) SyncConfigInput() *GoogleVertexAiFeatureOnlineStoreFeatureviewSyncConfig {
	var returns *GoogleVertexAiFeatureOnlineStoreFeatureviewSyncConfig
	_jsii_.Get(
		j,
		"syncConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) Timeouts() GoogleVertexAiFeatureOnlineStoreFeatureviewTimeoutsOutputReference {
	var returns GoogleVertexAiFeatureOnlineStoreFeatureviewTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) VectorSearchConfig() GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference {
	var returns GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference
	_jsii_.Get(
		j,
		"vectorSearchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) VectorSearchConfigInput() *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfig {
	var returns *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfig
	_jsii_.Get(
		j,
		"vectorSearchConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_vertex_ai_feature_online_store_featureview google_vertex_ai_feature_online_store_featureview} Resource.
func NewGoogleVertexAiFeatureOnlineStoreFeatureview(scope constructs.Construct, id *string, config *GoogleVertexAiFeatureOnlineStoreFeatureviewConfig) GoogleVertexAiFeatureOnlineStoreFeatureview {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiFeatureOnlineStoreFeatureviewParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiFeatureOnlineStoreFeatureview.GoogleVertexAiFeatureOnlineStoreFeatureview",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_vertex_ai_feature_online_store_featureview google_vertex_ai_feature_online_store_featureview} Resource.
func NewGoogleVertexAiFeatureOnlineStoreFeatureview_Override(g GoogleVertexAiFeatureOnlineStoreFeatureview, scope constructs.Construct, id *string, config *GoogleVertexAiFeatureOnlineStoreFeatureviewConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiFeatureOnlineStoreFeatureview.GoogleVertexAiFeatureOnlineStoreFeatureview",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview)SetFeatureOnlineStore(val *string) {
	if err := j.validateSetFeatureOnlineStoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"featureOnlineStore",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTN code for importing a GoogleVertexAiFeatureOnlineStoreFeatureview resource upon running "cdktn plan <stack-name>".
func GoogleVertexAiFeatureOnlineStoreFeatureview_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleVertexAiFeatureOnlineStoreFeatureview_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleVertexAiFeatureOnlineStoreFeatureview.GoogleVertexAiFeatureOnlineStoreFeatureview",
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
func GoogleVertexAiFeatureOnlineStoreFeatureview_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleVertexAiFeatureOnlineStoreFeatureview_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleVertexAiFeatureOnlineStoreFeatureview.GoogleVertexAiFeatureOnlineStoreFeatureview",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleVertexAiFeatureOnlineStoreFeatureview_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleVertexAiFeatureOnlineStoreFeatureview_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleVertexAiFeatureOnlineStoreFeatureview.GoogleVertexAiFeatureOnlineStoreFeatureview",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleVertexAiFeatureOnlineStoreFeatureview_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleVertexAiFeatureOnlineStoreFeatureview_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleVertexAiFeatureOnlineStoreFeatureview.GoogleVertexAiFeatureOnlineStoreFeatureview",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleVertexAiFeatureOnlineStoreFeatureview_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleVertexAiFeatureOnlineStoreFeatureview.GoogleVertexAiFeatureOnlineStoreFeatureview",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) PutBigQuerySource(value *GoogleVertexAiFeatureOnlineStoreFeatureviewBigQuerySource) {
	if err := g.validatePutBigQuerySourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBigQuerySource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) PutFeatureRegistrySource(value *GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySource) {
	if err := g.validatePutFeatureRegistrySourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFeatureRegistrySource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) PutSyncConfig(value *GoogleVertexAiFeatureOnlineStoreFeatureviewSyncConfig) {
	if err := g.validatePutSyncConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSyncConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) PutTimeouts(value *GoogleVertexAiFeatureOnlineStoreFeatureviewTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) PutVectorSearchConfig(value *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfig) {
	if err := g.validatePutVectorSearchConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVectorSearchConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) ResetBigQuerySource() {
	_jsii_.InvokeVoid(
		g,
		"resetBigQuerySource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) ResetFeatureRegistrySource() {
	_jsii_.InvokeVoid(
		g,
		"resetFeatureRegistrySource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) ResetSyncConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSyncConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) ResetVectorSearchConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetVectorSearchConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureview) With(mixins ...constructs.IMixin) constructs.IConstruct {
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


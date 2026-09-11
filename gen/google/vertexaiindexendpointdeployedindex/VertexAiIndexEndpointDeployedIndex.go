package vertexaiindexendpointdeployedindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/vertexaiindexendpointdeployedindex/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/vertex_ai_index_endpoint_deployed_index google_vertex_ai_index_endpoint_deployed_index}.
type VertexAiIndexEndpointDeployedIndex interface {
	cdktn.TerraformResource
	AutomaticResources() VertexAiIndexEndpointDeployedIndexAutomaticResourcesOutputReference
	AutomaticResourcesInput() *VertexAiIndexEndpointDeployedIndexAutomaticResources
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
	DedicatedResources() VertexAiIndexEndpointDeployedIndexDedicatedResourcesOutputReference
	DedicatedResourcesInput() *VertexAiIndexEndpointDeployedIndexDedicatedResources
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeployedIndexAuthConfig() VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigOutputReference
	DeployedIndexAuthConfigInput() *VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfig
	DeployedIndexId() *string
	SetDeployedIndexId(val *string)
	DeployedIndexIdInput() *string
	DeploymentGroup() *string
	SetDeploymentGroup(val *string)
	DeploymentGroupInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EnableAccessLogging() interface{}
	SetEnableAccessLogging(val interface{})
	EnableAccessLoggingInput() interface{}
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
	Index() *string
	SetIndex(val *string)
	IndexEndpoint() *string
	SetIndexEndpoint(val *string)
	IndexEndpointInput() *string
	IndexInput() *string
	IndexSyncTime() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	// The tree node.
	Node() constructs.Node
	PrivateEndpoints() VertexAiIndexEndpointDeployedIndexPrivateEndpointsList
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
	ReservedIpRanges() *[]*string
	SetReservedIpRanges(val *[]*string)
	ReservedIpRangesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VertexAiIndexEndpointDeployedIndexTimeoutsOutputReference
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
	PutAutomaticResources(value *VertexAiIndexEndpointDeployedIndexAutomaticResources)
	PutDedicatedResources(value *VertexAiIndexEndpointDeployedIndexDedicatedResources)
	PutDeployedIndexAuthConfig(value *VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfig)
	PutTimeouts(value *VertexAiIndexEndpointDeployedIndexTimeouts)
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
	ResetAutomaticResources()
	ResetDedicatedResources()
	ResetDeployedIndexAuthConfig()
	ResetDeploymentGroup()
	ResetDisplayName()
	ResetEnableAccessLogging()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetReservedIpRanges()
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

// The jsii proxy struct for VertexAiIndexEndpointDeployedIndex
type jsiiProxy_VertexAiIndexEndpointDeployedIndex struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) AutomaticResources() VertexAiIndexEndpointDeployedIndexAutomaticResourcesOutputReference {
	var returns VertexAiIndexEndpointDeployedIndexAutomaticResourcesOutputReference
	_jsii_.Get(
		j,
		"automaticResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) AutomaticResourcesInput() *VertexAiIndexEndpointDeployedIndexAutomaticResources {
	var returns *VertexAiIndexEndpointDeployedIndexAutomaticResources
	_jsii_.Get(
		j,
		"automaticResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) DedicatedResources() VertexAiIndexEndpointDeployedIndexDedicatedResourcesOutputReference {
	var returns VertexAiIndexEndpointDeployedIndexDedicatedResourcesOutputReference
	_jsii_.Get(
		j,
		"dedicatedResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) DedicatedResourcesInput() *VertexAiIndexEndpointDeployedIndexDedicatedResources {
	var returns *VertexAiIndexEndpointDeployedIndexDedicatedResources
	_jsii_.Get(
		j,
		"dedicatedResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) DeployedIndexAuthConfig() VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigOutputReference {
	var returns VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigOutputReference
	_jsii_.Get(
		j,
		"deployedIndexAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) DeployedIndexAuthConfigInput() *VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfig {
	var returns *VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfig
	_jsii_.Get(
		j,
		"deployedIndexAuthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) DeployedIndexId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deployedIndexId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) DeployedIndexIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deployedIndexIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) DeploymentGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) DeploymentGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) EnableAccessLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAccessLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) EnableAccessLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAccessLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) Index() *string {
	var returns *string
	_jsii_.Get(
		j,
		"index",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) IndexEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) IndexEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) IndexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) IndexSyncTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexSyncTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) PrivateEndpoints() VertexAiIndexEndpointDeployedIndexPrivateEndpointsList {
	var returns VertexAiIndexEndpointDeployedIndexPrivateEndpointsList
	_jsii_.Get(
		j,
		"privateEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) ReservedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reservedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) ReservedIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reservedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) Timeouts() VertexAiIndexEndpointDeployedIndexTimeoutsOutputReference {
	var returns VertexAiIndexEndpointDeployedIndexTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/vertex_ai_index_endpoint_deployed_index google_vertex_ai_index_endpoint_deployed_index} Resource.
func NewVertexAiIndexEndpointDeployedIndex(scope constructs.Construct, id *string, config *VertexAiIndexEndpointDeployedIndexConfig) VertexAiIndexEndpointDeployedIndex {
	_init_.Initialize()

	if err := validateNewVertexAiIndexEndpointDeployedIndexParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiIndexEndpointDeployedIndex{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiIndexEndpointDeployedIndex.VertexAiIndexEndpointDeployedIndex",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/vertex_ai_index_endpoint_deployed_index google_vertex_ai_index_endpoint_deployed_index} Resource.
func NewVertexAiIndexEndpointDeployedIndex_Override(v VertexAiIndexEndpointDeployedIndex, scope constructs.Construct, id *string, config *VertexAiIndexEndpointDeployedIndexConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiIndexEndpointDeployedIndex.VertexAiIndexEndpointDeployedIndex",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex)SetDeployedIndexId(val *string) {
	if err := j.validateSetDeployedIndexIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deployedIndexId",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex)SetDeploymentGroup(val *string) {
	if err := j.validateSetDeploymentGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentGroup",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex)SetEnableAccessLogging(val interface{}) {
	if err := j.validateSetEnableAccessLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAccessLogging",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex)SetIndex(val *string) {
	if err := j.validateSetIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"index",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex)SetIndexEndpoint(val *string) {
	if err := j.validateSetIndexEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexEndpoint",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndex)SetReservedIpRanges(val *[]*string) {
	if err := j.validateSetReservedIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedIpRanges",
		val,
	)
}

// Generates CDKTN code for importing a VertexAiIndexEndpointDeployedIndex resource upon running "cdktn plan <stack-name>".
func VertexAiIndexEndpointDeployedIndex_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateVertexAiIndexEndpointDeployedIndex_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.vertexAiIndexEndpointDeployedIndex.VertexAiIndexEndpointDeployedIndex",
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
func VertexAiIndexEndpointDeployedIndex_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVertexAiIndexEndpointDeployedIndex_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.vertexAiIndexEndpointDeployedIndex.VertexAiIndexEndpointDeployedIndex",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VertexAiIndexEndpointDeployedIndex_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVertexAiIndexEndpointDeployedIndex_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.vertexAiIndexEndpointDeployedIndex.VertexAiIndexEndpointDeployedIndex",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VertexAiIndexEndpointDeployedIndex_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVertexAiIndexEndpointDeployedIndex_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.vertexAiIndexEndpointDeployedIndex.VertexAiIndexEndpointDeployedIndex",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VertexAiIndexEndpointDeployedIndex_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.vertexAiIndexEndpointDeployedIndex.VertexAiIndexEndpointDeployedIndex",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := v.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) PutAutomaticResources(value *VertexAiIndexEndpointDeployedIndexAutomaticResources) {
	if err := v.validatePutAutomaticResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putAutomaticResources",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) PutDedicatedResources(value *VertexAiIndexEndpointDeployedIndexDedicatedResources) {
	if err := v.validatePutDedicatedResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putDedicatedResources",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) PutDeployedIndexAuthConfig(value *VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfig) {
	if err := v.validatePutDeployedIndexAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putDeployedIndexAuthConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) PutTimeouts(value *VertexAiIndexEndpointDeployedIndexTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := v.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) ResetAutomaticResources() {
	_jsii_.InvokeVoid(
		v,
		"resetAutomaticResources",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) ResetDedicatedResources() {
	_jsii_.InvokeVoid(
		v,
		"resetDedicatedResources",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) ResetDeployedIndexAuthConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetDeployedIndexAuthConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) ResetDeploymentGroup() {
	_jsii_.InvokeVoid(
		v,
		"resetDeploymentGroup",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) ResetDisplayName() {
	_jsii_.InvokeVoid(
		v,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) ResetEnableAccessLogging() {
	_jsii_.InvokeVoid(
		v,
		"resetEnableAccessLogging",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) ResetRegion() {
	_jsii_.InvokeVoid(
		v,
		"resetRegion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) ResetReservedIpRanges() {
	_jsii_.InvokeVoid(
		v,
		"resetReservedIpRanges",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndex) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		v,
		"with",
		args,
		&returns,
	)

	return returns
}


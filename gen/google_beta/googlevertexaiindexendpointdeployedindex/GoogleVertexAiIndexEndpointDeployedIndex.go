package googlevertexaiindexendpointdeployedindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevertexaiindexendpointdeployedindex/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_vertex_ai_index_endpoint_deployed_index google_vertex_ai_index_endpoint_deployed_index}.
type GoogleVertexAiIndexEndpointDeployedIndex interface {
	cdktn.TerraformResource
	AutomaticResources() GoogleVertexAiIndexEndpointDeployedIndexAutomaticResourcesOutputReference
	AutomaticResourcesInput() *GoogleVertexAiIndexEndpointDeployedIndexAutomaticResources
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
	DedicatedResources() GoogleVertexAiIndexEndpointDeployedIndexDedicatedResourcesOutputReference
	DedicatedResourcesInput() *GoogleVertexAiIndexEndpointDeployedIndexDedicatedResources
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeployedIndexAuthConfig() GoogleVertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigOutputReference
	DeployedIndexAuthConfigInput() *GoogleVertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfig
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
	PrivateEndpoints() GoogleVertexAiIndexEndpointDeployedIndexPrivateEndpointsList
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
	Timeouts() GoogleVertexAiIndexEndpointDeployedIndexTimeoutsOutputReference
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
	PutAutomaticResources(value *GoogleVertexAiIndexEndpointDeployedIndexAutomaticResources)
	PutDedicatedResources(value *GoogleVertexAiIndexEndpointDeployedIndexDedicatedResources)
	PutDeployedIndexAuthConfig(value *GoogleVertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfig)
	PutTimeouts(value *GoogleVertexAiIndexEndpointDeployedIndexTimeouts)
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

// The jsii proxy struct for GoogleVertexAiIndexEndpointDeployedIndex
type jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) AutomaticResources() GoogleVertexAiIndexEndpointDeployedIndexAutomaticResourcesOutputReference {
	var returns GoogleVertexAiIndexEndpointDeployedIndexAutomaticResourcesOutputReference
	_jsii_.Get(
		j,
		"automaticResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) AutomaticResourcesInput() *GoogleVertexAiIndexEndpointDeployedIndexAutomaticResources {
	var returns *GoogleVertexAiIndexEndpointDeployedIndexAutomaticResources
	_jsii_.Get(
		j,
		"automaticResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DedicatedResources() GoogleVertexAiIndexEndpointDeployedIndexDedicatedResourcesOutputReference {
	var returns GoogleVertexAiIndexEndpointDeployedIndexDedicatedResourcesOutputReference
	_jsii_.Get(
		j,
		"dedicatedResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DedicatedResourcesInput() *GoogleVertexAiIndexEndpointDeployedIndexDedicatedResources {
	var returns *GoogleVertexAiIndexEndpointDeployedIndexDedicatedResources
	_jsii_.Get(
		j,
		"dedicatedResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DeployedIndexAuthConfig() GoogleVertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigOutputReference {
	var returns GoogleVertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigOutputReference
	_jsii_.Get(
		j,
		"deployedIndexAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DeployedIndexAuthConfigInput() *GoogleVertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfig {
	var returns *GoogleVertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfig
	_jsii_.Get(
		j,
		"deployedIndexAuthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DeployedIndexId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deployedIndexId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DeployedIndexIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deployedIndexIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DeploymentGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DeploymentGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) EnableAccessLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAccessLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) EnableAccessLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAccessLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Index() *string {
	var returns *string
	_jsii_.Get(
		j,
		"index",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) IndexEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) IndexEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) IndexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) IndexSyncTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexSyncTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) PrivateEndpoints() GoogleVertexAiIndexEndpointDeployedIndexPrivateEndpointsList {
	var returns GoogleVertexAiIndexEndpointDeployedIndexPrivateEndpointsList
	_jsii_.Get(
		j,
		"privateEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ReservedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reservedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ReservedIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reservedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Timeouts() GoogleVertexAiIndexEndpointDeployedIndexTimeoutsOutputReference {
	var returns GoogleVertexAiIndexEndpointDeployedIndexTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_vertex_ai_index_endpoint_deployed_index google_vertex_ai_index_endpoint_deployed_index} Resource.
func NewGoogleVertexAiIndexEndpointDeployedIndex(scope constructs.Construct, id *string, config *GoogleVertexAiIndexEndpointDeployedIndexConfig) GoogleVertexAiIndexEndpointDeployedIndex {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiIndexEndpointDeployedIndexParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiIndexEndpointDeployedIndex.GoogleVertexAiIndexEndpointDeployedIndex",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_vertex_ai_index_endpoint_deployed_index google_vertex_ai_index_endpoint_deployed_index} Resource.
func NewGoogleVertexAiIndexEndpointDeployedIndex_Override(g GoogleVertexAiIndexEndpointDeployedIndex, scope constructs.Construct, id *string, config *GoogleVertexAiIndexEndpointDeployedIndexConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiIndexEndpointDeployedIndex.GoogleVertexAiIndexEndpointDeployedIndex",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetDeployedIndexId(val *string) {
	if err := j.validateSetDeployedIndexIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deployedIndexId",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetDeploymentGroup(val *string) {
	if err := j.validateSetDeploymentGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentGroup",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetEnableAccessLogging(val interface{}) {
	if err := j.validateSetEnableAccessLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAccessLogging",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetIndex(val *string) {
	if err := j.validateSetIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"index",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetIndexEndpoint(val *string) {
	if err := j.validateSetIndexEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetReservedIpRanges(val *[]*string) {
	if err := j.validateSetReservedIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedIpRanges",
		val,
	)
}

// Generates CDKTN code for importing a GoogleVertexAiIndexEndpointDeployedIndex resource upon running "cdktn plan <stack-name>".
func GoogleVertexAiIndexEndpointDeployedIndex_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleVertexAiIndexEndpointDeployedIndex_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleVertexAiIndexEndpointDeployedIndex.GoogleVertexAiIndexEndpointDeployedIndex",
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
func GoogleVertexAiIndexEndpointDeployedIndex_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleVertexAiIndexEndpointDeployedIndex_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleVertexAiIndexEndpointDeployedIndex.GoogleVertexAiIndexEndpointDeployedIndex",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleVertexAiIndexEndpointDeployedIndex_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleVertexAiIndexEndpointDeployedIndex_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleVertexAiIndexEndpointDeployedIndex.GoogleVertexAiIndexEndpointDeployedIndex",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleVertexAiIndexEndpointDeployedIndex_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleVertexAiIndexEndpointDeployedIndex_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleVertexAiIndexEndpointDeployedIndex.GoogleVertexAiIndexEndpointDeployedIndex",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleVertexAiIndexEndpointDeployedIndex_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleVertexAiIndexEndpointDeployedIndex.GoogleVertexAiIndexEndpointDeployedIndex",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) PutAutomaticResources(value *GoogleVertexAiIndexEndpointDeployedIndexAutomaticResources) {
	if err := g.validatePutAutomaticResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutomaticResources",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) PutDedicatedResources(value *GoogleVertexAiIndexEndpointDeployedIndexDedicatedResources) {
	if err := g.validatePutDedicatedResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDedicatedResources",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) PutDeployedIndexAuthConfig(value *GoogleVertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfig) {
	if err := g.validatePutDeployedIndexAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeployedIndexAuthConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) PutTimeouts(value *GoogleVertexAiIndexEndpointDeployedIndexTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ResetAutomaticResources() {
	_jsii_.InvokeVoid(
		g,
		"resetAutomaticResources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ResetDedicatedResources() {
	_jsii_.InvokeVoid(
		g,
		"resetDedicatedResources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ResetDeployedIndexAuthConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetDeployedIndexAuthConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ResetDeploymentGroup() {
	_jsii_.InvokeVoid(
		g,
		"resetDeploymentGroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ResetEnableAccessLogging() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableAccessLogging",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ResetReservedIpRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetReservedIpRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) With(mixins ...constructs.IMixin) constructs.IConstruct {
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


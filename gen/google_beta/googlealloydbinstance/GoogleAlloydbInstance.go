package googlealloydbinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlealloydbinstance/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_alloydb_instance google_alloydb_instance}.
type GoogleAlloydbInstance interface {
	cdktn.TerraformResource
	ActivationPolicy() *string
	SetActivationPolicy(val *string)
	ActivationPolicyInput() *string
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	AvailabilityType() *string
	SetAvailabilityType(val *string)
	AvailabilityTypeInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClientConnectionConfig() GoogleAlloydbInstanceClientConnectionConfigOutputReference
	ClientConnectionConfigInput() *GoogleAlloydbInstanceClientConnectionConfig
	Cluster() *string
	SetCluster(val *string)
	ClusterInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionPoolConfig() GoogleAlloydbInstanceConnectionPoolConfigOutputReference
	ConnectionPoolConfigInput() *GoogleAlloydbInstanceConnectionPoolConfig
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	DatabaseFlags() *map[string]*string
	SetDatabaseFlags(val *map[string]*string)
	DatabaseFlagsInput() *map[string]*string
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
	GceZone() *string
	SetGceZone(val *string)
	GceZoneInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceId() *string
	SetInstanceId(val *string)
	InstanceIdInput() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	IpAddress() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MachineConfig() GoogleAlloydbInstanceMachineConfigOutputReference
	MachineConfigInput() *GoogleAlloydbInstanceMachineConfig
	Name() *string
	NetworkConfig() GoogleAlloydbInstanceNetworkConfigOutputReference
	NetworkConfigInput() *GoogleAlloydbInstanceNetworkConfig
	// The tree node.
	Node() constructs.Node
	ObservabilityConfig() GoogleAlloydbInstanceObservabilityConfigOutputReference
	ObservabilityConfigInput() *GoogleAlloydbInstanceObservabilityConfig
	OutboundPublicIpAddresses() *[]*string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PscInstanceConfig() GoogleAlloydbInstancePscInstanceConfigOutputReference
	PscInstanceConfigInput() *GoogleAlloydbInstancePscInstanceConfig
	PublicIpAddress() *string
	QueryInsightsConfig() GoogleAlloydbInstanceQueryInsightsConfigOutputReference
	QueryInsightsConfigInput() *GoogleAlloydbInstanceQueryInsightsConfig
	// Experimental.
	RawOverrides() interface{}
	ReadPoolConfig() GoogleAlloydbInstanceReadPoolConfigOutputReference
	ReadPoolConfigInput() *GoogleAlloydbInstanceReadPoolConfig
	Reconciling() cdktn.IResolvable
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktn.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleAlloydbInstanceTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
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
	PutClientConnectionConfig(value *GoogleAlloydbInstanceClientConnectionConfig)
	PutConnectionPoolConfig(value *GoogleAlloydbInstanceConnectionPoolConfig)
	PutMachineConfig(value *GoogleAlloydbInstanceMachineConfig)
	PutNetworkConfig(value *GoogleAlloydbInstanceNetworkConfig)
	PutObservabilityConfig(value *GoogleAlloydbInstanceObservabilityConfig)
	PutPscInstanceConfig(value *GoogleAlloydbInstancePscInstanceConfig)
	PutQueryInsightsConfig(value *GoogleAlloydbInstanceQueryInsightsConfig)
	PutReadPoolConfig(value *GoogleAlloydbInstanceReadPoolConfig)
	PutTimeouts(value *GoogleAlloydbInstanceTimeouts)
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
	ResetActivationPolicy()
	ResetAnnotations()
	ResetAvailabilityType()
	ResetClientConnectionConfig()
	ResetConnectionPoolConfig()
	ResetDatabaseFlags()
	ResetDisplayName()
	ResetGceZone()
	ResetId()
	ResetLabels()
	ResetMachineConfig()
	ResetNetworkConfig()
	ResetObservabilityConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPscInstanceConfig()
	ResetQueryInsightsConfig()
	ResetReadPoolConfig()
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

// The jsii proxy struct for GoogleAlloydbInstance
type jsiiProxy_GoogleAlloydbInstance struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleAlloydbInstance) ActivationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ActivationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) AvailabilityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) AvailabilityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ClientConnectionConfig() GoogleAlloydbInstanceClientConnectionConfigOutputReference {
	var returns GoogleAlloydbInstanceClientConnectionConfigOutputReference
	_jsii_.Get(
		j,
		"clientConnectionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ClientConnectionConfigInput() *GoogleAlloydbInstanceClientConnectionConfig {
	var returns *GoogleAlloydbInstanceClientConnectionConfig
	_jsii_.Get(
		j,
		"clientConnectionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Cluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ClusterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ConnectionPoolConfig() GoogleAlloydbInstanceConnectionPoolConfigOutputReference {
	var returns GoogleAlloydbInstanceConnectionPoolConfigOutputReference
	_jsii_.Get(
		j,
		"connectionPoolConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ConnectionPoolConfigInput() *GoogleAlloydbInstanceConnectionPoolConfig {
	var returns *GoogleAlloydbInstanceConnectionPoolConfig
	_jsii_.Get(
		j,
		"connectionPoolConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) DatabaseFlags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"databaseFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) DatabaseFlagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"databaseFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) EffectiveAnnotations() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) GceZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gceZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) GceZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gceZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) InstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) MachineConfig() GoogleAlloydbInstanceMachineConfigOutputReference {
	var returns GoogleAlloydbInstanceMachineConfigOutputReference
	_jsii_.Get(
		j,
		"machineConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) MachineConfigInput() *GoogleAlloydbInstanceMachineConfig {
	var returns *GoogleAlloydbInstanceMachineConfig
	_jsii_.Get(
		j,
		"machineConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) NetworkConfig() GoogleAlloydbInstanceNetworkConfigOutputReference {
	var returns GoogleAlloydbInstanceNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) NetworkConfigInput() *GoogleAlloydbInstanceNetworkConfig {
	var returns *GoogleAlloydbInstanceNetworkConfig
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ObservabilityConfig() GoogleAlloydbInstanceObservabilityConfigOutputReference {
	var returns GoogleAlloydbInstanceObservabilityConfigOutputReference
	_jsii_.Get(
		j,
		"observabilityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ObservabilityConfigInput() *GoogleAlloydbInstanceObservabilityConfig {
	var returns *GoogleAlloydbInstanceObservabilityConfig
	_jsii_.Get(
		j,
		"observabilityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) OutboundPublicIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundPublicIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) PscInstanceConfig() GoogleAlloydbInstancePscInstanceConfigOutputReference {
	var returns GoogleAlloydbInstancePscInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"pscInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) PscInstanceConfigInput() *GoogleAlloydbInstancePscInstanceConfig {
	var returns *GoogleAlloydbInstancePscInstanceConfig
	_jsii_.Get(
		j,
		"pscInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) PublicIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) QueryInsightsConfig() GoogleAlloydbInstanceQueryInsightsConfigOutputReference {
	var returns GoogleAlloydbInstanceQueryInsightsConfigOutputReference
	_jsii_.Get(
		j,
		"queryInsightsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) QueryInsightsConfigInput() *GoogleAlloydbInstanceQueryInsightsConfig {
	var returns *GoogleAlloydbInstanceQueryInsightsConfig
	_jsii_.Get(
		j,
		"queryInsightsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ReadPoolConfig() GoogleAlloydbInstanceReadPoolConfigOutputReference {
	var returns GoogleAlloydbInstanceReadPoolConfigOutputReference
	_jsii_.Get(
		j,
		"readPoolConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ReadPoolConfigInput() *GoogleAlloydbInstanceReadPoolConfig {
	var returns *GoogleAlloydbInstanceReadPoolConfig
	_jsii_.Get(
		j,
		"readPoolConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Reconciling() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Timeouts() GoogleAlloydbInstanceTimeoutsOutputReference {
	var returns GoogleAlloydbInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_alloydb_instance google_alloydb_instance} Resource.
func NewGoogleAlloydbInstance(scope constructs.Construct, id *string, config *GoogleAlloydbInstanceConfig) GoogleAlloydbInstance {
	_init_.Initialize()

	if err := validateNewGoogleAlloydbInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAlloydbInstance{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleAlloydbInstance.GoogleAlloydbInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_alloydb_instance google_alloydb_instance} Resource.
func NewGoogleAlloydbInstance_Override(g GoogleAlloydbInstance, scope constructs.Construct, id *string, config *GoogleAlloydbInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleAlloydbInstance.GoogleAlloydbInstance",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetActivationPolicy(val *string) {
	if err := j.validateSetActivationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activationPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetAvailabilityType(val *string) {
	if err := j.validateSetAvailabilityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityType",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetCluster(val *string) {
	if err := j.validateSetClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cluster",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetDatabaseFlags(val *map[string]*string) {
	if err := j.validateSetDatabaseFlagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseFlags",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetGceZone(val *string) {
	if err := j.validateSetGceZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gceZone",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetInstanceId(val *string) {
	if err := j.validateSetInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a GoogleAlloydbInstance resource upon running "cdktn plan <stack-name>".
func GoogleAlloydbInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleAlloydbInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleAlloydbInstance.GoogleAlloydbInstance",
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
func GoogleAlloydbInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAlloydbInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleAlloydbInstance.GoogleAlloydbInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleAlloydbInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAlloydbInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleAlloydbInstance.GoogleAlloydbInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleAlloydbInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAlloydbInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleAlloydbInstance.GoogleAlloydbInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleAlloydbInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleAlloydbInstance.GoogleAlloydbInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleAlloydbInstance) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAlloydbInstance) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleAlloydbInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAlloydbInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAlloydbInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAlloydbInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAlloydbInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAlloydbInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAlloydbInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAlloydbInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbInstance) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleAlloydbInstance) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (g *jsiiProxy_GoogleAlloydbInstance) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) PutClientConnectionConfig(value *GoogleAlloydbInstanceClientConnectionConfig) {
	if err := g.validatePutClientConnectionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClientConnectionConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) PutConnectionPoolConfig(value *GoogleAlloydbInstanceConnectionPoolConfig) {
	if err := g.validatePutConnectionPoolConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConnectionPoolConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) PutMachineConfig(value *GoogleAlloydbInstanceMachineConfig) {
	if err := g.validatePutMachineConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMachineConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) PutNetworkConfig(value *GoogleAlloydbInstanceNetworkConfig) {
	if err := g.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) PutObservabilityConfig(value *GoogleAlloydbInstanceObservabilityConfig) {
	if err := g.validatePutObservabilityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putObservabilityConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) PutPscInstanceConfig(value *GoogleAlloydbInstancePscInstanceConfig) {
	if err := g.validatePutPscInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPscInstanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) PutQueryInsightsConfig(value *GoogleAlloydbInstanceQueryInsightsConfig) {
	if err := g.validatePutQueryInsightsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putQueryInsightsConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) PutReadPoolConfig(value *GoogleAlloydbInstanceReadPoolConfig) {
	if err := g.validatePutReadPoolConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putReadPoolConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) PutTimeouts(value *GoogleAlloydbInstanceTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetActivationPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetActivationPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetAnnotations() {
	_jsii_.InvokeVoid(
		g,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetAvailabilityType() {
	_jsii_.InvokeVoid(
		g,
		"resetAvailabilityType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetClientConnectionConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetClientConnectionConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetConnectionPoolConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetConnectionPoolConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetDatabaseFlags() {
	_jsii_.InvokeVoid(
		g,
		"resetDatabaseFlags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetGceZone() {
	_jsii_.InvokeVoid(
		g,
		"resetGceZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetMachineConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetMachineConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetObservabilityConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetObservabilityConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetPscInstanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPscInstanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetQueryInsightsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetQueryInsightsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetReadPoolConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetReadPoolConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbInstance) With(mixins ...constructs.IMixin) constructs.IConstruct {
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


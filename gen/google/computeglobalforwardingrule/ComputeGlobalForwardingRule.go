package computeglobalforwardingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeglobalforwardingrule/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_global_forwarding_rule google_compute_global_forwarding_rule}.
type ComputeGlobalForwardingRule interface {
	cdktn.TerraformResource
	BaseForwardingRule() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveLabels() cdktn.StringMap
	ExternalManagedBackendBucketMigrationState() *string
	SetExternalManagedBackendBucketMigrationState(val *string)
	ExternalManagedBackendBucketMigrationStateInput() *string
	ExternalManagedBackendBucketMigrationTestingPercentage() *float64
	SetExternalManagedBackendBucketMigrationTestingPercentage(val *float64)
	ExternalManagedBackendBucketMigrationTestingPercentageInput() *float64
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	ForwardingRuleId() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IpAddress() *string
	SetIpAddress(val *string)
	IpAddressInput() *string
	IpProtocol() *string
	SetIpProtocol(val *string)
	IpProtocolInput() *string
	IpVersion() *string
	SetIpVersion(val *string)
	IpVersionInput() *string
	LabelFingerprint() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LoadBalancingScheme() *string
	SetLoadBalancingScheme(val *string)
	LoadBalancingSchemeInput() *string
	MetadataFilters() ComputeGlobalForwardingRuleMetadataFiltersList
	MetadataFiltersInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	NetworkTier() *string
	SetNetworkTier(val *string)
	NetworkTierInput() *string
	NoAutomateDnsZone() interface{}
	SetNoAutomateDnsZone(val interface{})
	NoAutomateDnsZoneInput() interface{}
	// The tree node.
	Node() constructs.Node
	PortRange() *string
	SetPortRange(val *string)
	PortRangeInput() *string
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
	PscConnectionId() *string
	PscConnectionStatus() *string
	// Experimental.
	RawOverrides() interface{}
	SelfLink() *string
	ServiceDirectoryRegistrations() ComputeGlobalForwardingRuleServiceDirectoryRegistrationsOutputReference
	ServiceDirectoryRegistrationsInput() *ComputeGlobalForwardingRuleServiceDirectoryRegistrations
	SourceIpRanges() *[]*string
	SetSourceIpRanges(val *[]*string)
	SourceIpRangesInput() *[]*string
	Subnetwork() *string
	SetSubnetwork(val *string)
	SubnetworkInput() *string
	Target() *string
	SetTarget(val *string)
	TargetInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktn.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeGlobalForwardingRuleTimeoutsOutputReference
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
	PutMetadataFilters(value interface{})
	PutServiceDirectoryRegistrations(value *ComputeGlobalForwardingRuleServiceDirectoryRegistrations)
	PutTimeouts(value *ComputeGlobalForwardingRuleTimeouts)
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
	ResetDescription()
	ResetExternalManagedBackendBucketMigrationState()
	ResetExternalManagedBackendBucketMigrationTestingPercentage()
	ResetId()
	ResetIpAddress()
	ResetIpProtocol()
	ResetIpVersion()
	ResetLabels()
	ResetLoadBalancingScheme()
	ResetMetadataFilters()
	ResetNetwork()
	ResetNetworkTier()
	ResetNoAutomateDnsZone()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPortRange()
	ResetProject()
	ResetServiceDirectoryRegistrations()
	ResetSourceIpRanges()
	ResetSubnetwork()
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

// The jsii proxy struct for ComputeGlobalForwardingRule
type jsiiProxy_ComputeGlobalForwardingRule struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) BaseForwardingRule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseForwardingRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) ExternalManagedBackendBucketMigrationState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalManagedBackendBucketMigrationState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) ExternalManagedBackendBucketMigrationStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalManagedBackendBucketMigrationStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) ExternalManagedBackendBucketMigrationTestingPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"externalManagedBackendBucketMigrationTestingPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) ExternalManagedBackendBucketMigrationTestingPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"externalManagedBackendBucketMigrationTestingPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) ForwardingRuleId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"forwardingRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) IpProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) IpProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) IpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) IpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) LabelFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) LoadBalancingScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) LoadBalancingSchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) MetadataFilters() ComputeGlobalForwardingRuleMetadataFiltersList {
	var returns ComputeGlobalForwardingRuleMetadataFiltersList
	_jsii_.Get(
		j,
		"metadataFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) MetadataFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) NetworkTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) NetworkTierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) NoAutomateDnsZone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noAutomateDnsZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) NoAutomateDnsZoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noAutomateDnsZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) PortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) PortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) PscConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pscConnectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) PscConnectionStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pscConnectionStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) ServiceDirectoryRegistrations() ComputeGlobalForwardingRuleServiceDirectoryRegistrationsOutputReference {
	var returns ComputeGlobalForwardingRuleServiceDirectoryRegistrationsOutputReference
	_jsii_.Get(
		j,
		"serviceDirectoryRegistrations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) ServiceDirectoryRegistrationsInput() *ComputeGlobalForwardingRuleServiceDirectoryRegistrations {
	var returns *ComputeGlobalForwardingRuleServiceDirectoryRegistrations
	_jsii_.Get(
		j,
		"serviceDirectoryRegistrationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) SourceIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) SourceIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) SubnetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) Timeouts() ComputeGlobalForwardingRuleTimeoutsOutputReference {
	var returns ComputeGlobalForwardingRuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalForwardingRule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_global_forwarding_rule google_compute_global_forwarding_rule} Resource.
func NewComputeGlobalForwardingRule(scope constructs.Construct, id *string, config *ComputeGlobalForwardingRuleConfig) ComputeGlobalForwardingRule {
	_init_.Initialize()

	if err := validateNewComputeGlobalForwardingRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeGlobalForwardingRule{}

	_jsii_.Create(
		"@cdktn/provider-google.computeGlobalForwardingRule.ComputeGlobalForwardingRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_global_forwarding_rule google_compute_global_forwarding_rule} Resource.
func NewComputeGlobalForwardingRule_Override(c ComputeGlobalForwardingRule, scope constructs.Construct, id *string, config *ComputeGlobalForwardingRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeGlobalForwardingRule.ComputeGlobalForwardingRule",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetExternalManagedBackendBucketMigrationState(val *string) {
	if err := j.validateSetExternalManagedBackendBucketMigrationStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalManagedBackendBucketMigrationState",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetExternalManagedBackendBucketMigrationTestingPercentage(val *float64) {
	if err := j.validateSetExternalManagedBackendBucketMigrationTestingPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalManagedBackendBucketMigrationTestingPercentage",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetIpAddress(val *string) {
	if err := j.validateSetIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetIpProtocol(val *string) {
	if err := j.validateSetIpProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipProtocol",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetIpVersion(val *string) {
	if err := j.validateSetIpVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipVersion",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetLoadBalancingScheme(val *string) {
	if err := j.validateSetLoadBalancingSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancingScheme",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetNetworkTier(val *string) {
	if err := j.validateSetNetworkTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkTier",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetNoAutomateDnsZone(val interface{}) {
	if err := j.validateSetNoAutomateDnsZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noAutomateDnsZone",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetPortRange(val *string) {
	if err := j.validateSetPortRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portRange",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetSourceIpRanges(val *[]*string) {
	if err := j.validateSetSourceIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIpRanges",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetSubnetwork(val *string) {
	if err := j.validateSetSubnetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetwork",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalForwardingRule)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

// Generates CDKTN code for importing a ComputeGlobalForwardingRule resource upon running "cdktn plan <stack-name>".
func ComputeGlobalForwardingRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateComputeGlobalForwardingRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.computeGlobalForwardingRule.ComputeGlobalForwardingRule",
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
func ComputeGlobalForwardingRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeGlobalForwardingRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.computeGlobalForwardingRule.ComputeGlobalForwardingRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeGlobalForwardingRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeGlobalForwardingRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.computeGlobalForwardingRule.ComputeGlobalForwardingRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeGlobalForwardingRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeGlobalForwardingRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.computeGlobalForwardingRule.ComputeGlobalForwardingRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeGlobalForwardingRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.computeGlobalForwardingRule.ComputeGlobalForwardingRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := c.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) PutMetadataFilters(value interface{}) {
	if err := c.validatePutMetadataFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMetadataFilters",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) PutServiceDirectoryRegistrations(value *ComputeGlobalForwardingRuleServiceDirectoryRegistrations) {
	if err := c.validatePutServiceDirectoryRegistrationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServiceDirectoryRegistrations",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) PutTimeouts(value *ComputeGlobalForwardingRuleTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := c.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ResetExternalManagedBackendBucketMigrationState() {
	_jsii_.InvokeVoid(
		c,
		"resetExternalManagedBackendBucketMigrationState",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ResetExternalManagedBackendBucketMigrationTestingPercentage() {
	_jsii_.InvokeVoid(
		c,
		"resetExternalManagedBackendBucketMigrationTestingPercentage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ResetIpAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ResetIpProtocol() {
	_jsii_.InvokeVoid(
		c,
		"resetIpProtocol",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ResetIpVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetIpVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ResetLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ResetLoadBalancingScheme() {
	_jsii_.InvokeVoid(
		c,
		"resetLoadBalancingScheme",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ResetMetadataFilters() {
	_jsii_.InvokeVoid(
		c,
		"resetMetadataFilters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ResetNetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetNetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ResetNetworkTier() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkTier",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ResetNoAutomateDnsZone() {
	_jsii_.InvokeVoid(
		c,
		"resetNoAutomateDnsZone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ResetPortRange() {
	_jsii_.InvokeVoid(
		c,
		"resetPortRange",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ResetServiceDirectoryRegistrations() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceDirectoryRegistrations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ResetSourceIpRanges() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceIpRanges",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ResetSubnetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetSubnetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeGlobalForwardingRule) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}


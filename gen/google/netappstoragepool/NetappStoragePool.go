package netappstoragepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/netappstoragepool/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/netapp_storage_pool google_netapp_storage_pool}.
type NetappStoragePool interface {
	cdktn.TerraformResource
	ActiveDirectory() *string
	SetActiveDirectory(val *string)
	ActiveDirectoryInput() *string
	AllowAutoTiering() interface{}
	SetAllowAutoTiering(val interface{})
	AllowAutoTieringInput() interface{}
	AvailableThroughputMibps() *float64
	CapacityGib() *string
	SetCapacityGib(val *string)
	CapacityGibInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ColdTierSizeUsedGib() *string
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
	CustomPerformanceEnabled() interface{}
	SetCustomPerformanceEnabled(val interface{})
	CustomPerformanceEnabledInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveLabels() cdktn.StringMap
	EnableHotTierAutoResize() interface{}
	SetEnableHotTierAutoResize(val interface{})
	EnableHotTierAutoResizeInput() interface{}
	EncryptionType() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HotTierSizeGib() *string
	SetHotTierSizeGib(val *string)
	HotTierSizeGibInput() *string
	HotTierSizeUsedGib() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	KmsConfig() *string
	SetKmsConfig(val *string)
	KmsConfigInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LdapEnabled() interface{}
	SetLdapEnabled(val interface{})
	LdapEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
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
	QosType() *string
	SetQosType(val *string)
	QosTypeInput() *string
	// Experimental.
	RawOverrides() interface{}
	ReplicaZone() *string
	SetReplicaZone(val *string)
	ReplicaZoneInput() *string
	ScaleType() *string
	SetScaleType(val *string)
	ScaleTypeInput() *string
	ServiceLevel() *string
	SetServiceLevel(val *string)
	ServiceLevelInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktn.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() NetappStoragePoolTimeoutsOutputReference
	TimeoutsInput() interface{}
	TotalIops() *string
	SetTotalIops(val *string)
	TotalIopsInput() *string
	TotalThroughputMibps() *string
	SetTotalThroughputMibps(val *string)
	TotalThroughputMibpsInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	VolumeCapacityGib() *string
	VolumeCount() *float64
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
	PutTimeouts(value *NetappStoragePoolTimeouts)
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
	ResetActiveDirectory()
	ResetAllowAutoTiering()
	ResetCustomPerformanceEnabled()
	ResetDescription()
	ResetEnableHotTierAutoResize()
	ResetHotTierSizeGib()
	ResetId()
	ResetKmsConfig()
	ResetLabels()
	ResetLdapEnabled()
	ResetMode()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetQosType()
	ResetReplicaZone()
	ResetScaleType()
	ResetTimeouts()
	ResetTotalIops()
	ResetTotalThroughputMibps()
	ResetType()
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

// The jsii proxy struct for NetappStoragePool
type jsiiProxy_NetappStoragePool struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_NetappStoragePool) ActiveDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) ActiveDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) AllowAutoTiering() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAutoTiering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) AllowAutoTieringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAutoTieringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) AvailableThroughputMibps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableThroughputMibps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) CapacityGib() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) CapacityGibInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityGibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) ColdTierSizeUsedGib() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coldTierSizeUsedGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) CustomPerformanceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customPerformanceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) CustomPerformanceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customPerformanceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) EnableHotTierAutoResize() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHotTierAutoResize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) EnableHotTierAutoResizeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHotTierAutoResizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) EncryptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) HotTierSizeGib() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hotTierSizeGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) HotTierSizeGibInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hotTierSizeGibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) HotTierSizeUsedGib() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hotTierSizeUsedGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) KmsConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) KmsConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) LdapEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ldapEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) LdapEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ldapEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) QosType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qosType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) QosTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qosTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) ReplicaZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) ReplicaZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) ScaleType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) ScaleTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) ServiceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) ServiceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) Timeouts() NetappStoragePoolTimeoutsOutputReference {
	var returns NetappStoragePoolTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) TotalIops() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) TotalIopsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalIopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) TotalThroughputMibps() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalThroughputMibps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) TotalThroughputMibpsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalThroughputMibpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) VolumeCapacityGib() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeCapacityGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) VolumeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappStoragePool) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/netapp_storage_pool google_netapp_storage_pool} Resource.
func NewNetappStoragePool(scope constructs.Construct, id *string, config *NetappStoragePoolConfig) NetappStoragePool {
	_init_.Initialize()

	if err := validateNewNetappStoragePoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetappStoragePool{}

	_jsii_.Create(
		"@cdktn/provider-google.netappStoragePool.NetappStoragePool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/netapp_storage_pool google_netapp_storage_pool} Resource.
func NewNetappStoragePool_Override(n NetappStoragePool, scope constructs.Construct, id *string, config *NetappStoragePoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.netappStoragePool.NetappStoragePool",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetActiveDirectory(val *string) {
	if err := j.validateSetActiveDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeDirectory",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetAllowAutoTiering(val interface{}) {
	if err := j.validateSetAllowAutoTieringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAutoTiering",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetCapacityGib(val *string) {
	if err := j.validateSetCapacityGibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityGib",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetCustomPerformanceEnabled(val interface{}) {
	if err := j.validateSetCustomPerformanceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customPerformanceEnabled",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetEnableHotTierAutoResize(val interface{}) {
	if err := j.validateSetEnableHotTierAutoResizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHotTierAutoResize",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetHotTierSizeGib(val *string) {
	if err := j.validateSetHotTierSizeGibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hotTierSizeGib",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetKmsConfig(val *string) {
	if err := j.validateSetKmsConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsConfig",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetLdapEnabled(val interface{}) {
	if err := j.validateSetLdapEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ldapEnabled",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetQosType(val *string) {
	if err := j.validateSetQosTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qosType",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetReplicaZone(val *string) {
	if err := j.validateSetReplicaZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaZone",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetScaleType(val *string) {
	if err := j.validateSetScaleTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleType",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetServiceLevel(val *string) {
	if err := j.validateSetServiceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceLevel",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetTotalIops(val *string) {
	if err := j.validateSetTotalIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalIops",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetTotalThroughputMibps(val *string) {
	if err := j.validateSetTotalThroughputMibpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalThroughputMibps",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_NetappStoragePool)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

// Generates CDKTN code for importing a NetappStoragePool resource upon running "cdktn plan <stack-name>".
func NetappStoragePool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateNetappStoragePool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.netappStoragePool.NetappStoragePool",
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
func NetappStoragePool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetappStoragePool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.netappStoragePool.NetappStoragePool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetappStoragePool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetappStoragePool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.netappStoragePool.NetappStoragePool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetappStoragePool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetappStoragePool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.netappStoragePool.NetappStoragePool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NetappStoragePool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.netappStoragePool.NetappStoragePool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetappStoragePool) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NetappStoragePool) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NetappStoragePool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappStoragePool) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappStoragePool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappStoragePool) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappStoragePool) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappStoragePool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappStoragePool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappStoragePool) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappStoragePool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappStoragePool) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappStoragePool) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NetappStoragePool) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappStoragePool) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := n.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappStoragePool) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetappStoragePool) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NetappStoragePool) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetappStoragePool) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NetappStoragePool) PutTimeouts(value *NetappStoragePoolTimeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetappStoragePool) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := n.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (n *jsiiProxy_NetappStoragePool) ResetActiveDirectory() {
	_jsii_.InvokeVoid(
		n,
		"resetActiveDirectory",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappStoragePool) ResetAllowAutoTiering() {
	_jsii_.InvokeVoid(
		n,
		"resetAllowAutoTiering",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappStoragePool) ResetCustomPerformanceEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetCustomPerformanceEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappStoragePool) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappStoragePool) ResetEnableHotTierAutoResize() {
	_jsii_.InvokeVoid(
		n,
		"resetEnableHotTierAutoResize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappStoragePool) ResetHotTierSizeGib() {
	_jsii_.InvokeVoid(
		n,
		"resetHotTierSizeGib",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappStoragePool) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappStoragePool) ResetKmsConfig() {
	_jsii_.InvokeVoid(
		n,
		"resetKmsConfig",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappStoragePool) ResetLabels() {
	_jsii_.InvokeVoid(
		n,
		"resetLabels",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappStoragePool) ResetLdapEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetLdapEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappStoragePool) ResetMode() {
	_jsii_.InvokeVoid(
		n,
		"resetMode",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappStoragePool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappStoragePool) ResetProject() {
	_jsii_.InvokeVoid(
		n,
		"resetProject",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappStoragePool) ResetQosType() {
	_jsii_.InvokeVoid(
		n,
		"resetQosType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappStoragePool) ResetReplicaZone() {
	_jsii_.InvokeVoid(
		n,
		"resetReplicaZone",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappStoragePool) ResetScaleType() {
	_jsii_.InvokeVoid(
		n,
		"resetScaleType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappStoragePool) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappStoragePool) ResetTotalIops() {
	_jsii_.InvokeVoid(
		n,
		"resetTotalIops",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappStoragePool) ResetTotalThroughputMibps() {
	_jsii_.InvokeVoid(
		n,
		"resetTotalThroughputMibps",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappStoragePool) ResetType() {
	_jsii_.InvokeVoid(
		n,
		"resetType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappStoragePool) ResetZone() {
	_jsii_.InvokeVoid(
		n,
		"resetZone",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappStoragePool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappStoragePool) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappStoragePool) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappStoragePool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappStoragePool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappStoragePool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappStoragePool) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		n,
		"with",
		args,
		&returns,
	)

	return returns
}


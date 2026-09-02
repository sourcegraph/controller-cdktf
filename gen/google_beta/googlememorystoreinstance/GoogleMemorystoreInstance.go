package googlememorystoreinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlememorystoreinstance/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_memorystore_instance google_memorystore_instance}.
type GoogleMemorystoreInstance interface {
	cdktn.TerraformResource
	AuthorizationMode() *string
	SetAuthorizationMode(val *string)
	AuthorizationModeInput() *string
	AutomatedBackupConfig() GoogleMemorystoreInstanceAutomatedBackupConfigOutputReference
	AutomatedBackupConfigInput() *GoogleMemorystoreInstanceAutomatedBackupConfig
	AvailableMaintenanceVersions() *[]*string
	BackupCollection() *string
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
	CrossInstanceReplicationConfig() GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference
	CrossInstanceReplicationConfigInput() *GoogleMemorystoreInstanceCrossInstanceReplicationConfig
	DeletionProtectionEnabled() interface{}
	SetDeletionProtectionEnabled(val interface{})
	DeletionProtectionEnabledInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DesiredAutoCreatedEndpoints() GoogleMemorystoreInstanceDesiredAutoCreatedEndpointsList
	DesiredAutoCreatedEndpointsInput() interface{}
	DesiredPscAutoConnections() GoogleMemorystoreInstanceDesiredPscAutoConnectionsList
	DesiredPscAutoConnectionsInput() interface{}
	DiscoveryEndpoints() GoogleMemorystoreInstanceDiscoveryEndpointsList
	EffectiveLabels() cdktn.StringMap
	EffectiveMaintenanceVersion() *string
	Endpoints() GoogleMemorystoreInstanceEndpointsList
	EngineConfigs() *map[string]*string
	SetEngineConfigs(val *map[string]*string)
	EngineConfigsInput() *map[string]*string
	EngineVersion() *string
	SetEngineVersion(val *string)
	EngineVersionInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GcsSource() GoogleMemorystoreInstanceGcsSourceOutputReference
	GcsSourceInput() *GoogleMemorystoreInstanceGcsSource
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceId() *string
	SetInstanceId(val *string)
	InstanceIdInput() *string
	KmsKey() *string
	SetKmsKey(val *string)
	KmsKeyInput() *string
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
	MaintenancePolicy() GoogleMemorystoreInstanceMaintenancePolicyOutputReference
	MaintenancePolicyInput() *GoogleMemorystoreInstanceMaintenancePolicy
	MaintenanceSchedule() GoogleMemorystoreInstanceMaintenanceScheduleList
	MaintenanceVersion() *string
	SetMaintenanceVersion(val *string)
	MaintenanceVersionInput() *string
	ManagedBackupSource() GoogleMemorystoreInstanceManagedBackupSourceOutputReference
	ManagedBackupSourceInput() *GoogleMemorystoreInstanceManagedBackupSource
	ManagedServerCa() GoogleMemorystoreInstanceManagedServerCaList
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	NodeConfig() GoogleMemorystoreInstanceNodeConfigList
	NodeType() *string
	SetNodeType(val *string)
	NodeTypeInput() *string
	PersistenceConfig() GoogleMemorystoreInstancePersistenceConfigOutputReference
	PersistenceConfigInput() *GoogleMemorystoreInstancePersistenceConfig
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
	PscAttachmentDetails() GoogleMemorystoreInstancePscAttachmentDetailsList
	PscAutoConnections() GoogleMemorystoreInstancePscAutoConnectionsList
	// Experimental.
	RawOverrides() interface{}
	ReplicaCount() *float64
	SetReplicaCount(val *float64)
	ReplicaCountInput() *float64
	ServerCaMode() *string
	SetServerCaMode(val *string)
	ServerCaModeInput() *string
	ServerCaPool() *string
	SetServerCaPool(val *string)
	ServerCaPoolInput() *string
	ShardCount() *float64
	SetShardCount(val *float64)
	ShardCountInput() *float64
	State() *string
	StateInfo() GoogleMemorystoreInstanceStateInfoList
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktn.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleMemorystoreInstanceTimeoutsOutputReference
	TimeoutsInput() interface{}
	TransitEncryptionMode() *string
	SetTransitEncryptionMode(val *string)
	TransitEncryptionModeInput() *string
	Uid() *string
	UpdateTime() *string
	ZoneDistributionConfig() GoogleMemorystoreInstanceZoneDistributionConfigOutputReference
	ZoneDistributionConfigInput() *GoogleMemorystoreInstanceZoneDistributionConfig
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
	PutAutomatedBackupConfig(value *GoogleMemorystoreInstanceAutomatedBackupConfig)
	PutCrossInstanceReplicationConfig(value *GoogleMemorystoreInstanceCrossInstanceReplicationConfig)
	PutDesiredAutoCreatedEndpoints(value interface{})
	PutDesiredPscAutoConnections(value interface{})
	PutGcsSource(value *GoogleMemorystoreInstanceGcsSource)
	PutMaintenancePolicy(value *GoogleMemorystoreInstanceMaintenancePolicy)
	PutManagedBackupSource(value *GoogleMemorystoreInstanceManagedBackupSource)
	PutPersistenceConfig(value *GoogleMemorystoreInstancePersistenceConfig)
	PutTimeouts(value *GoogleMemorystoreInstanceTimeouts)
	PutZoneDistributionConfig(value *GoogleMemorystoreInstanceZoneDistributionConfig)
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
	ResetAuthorizationMode()
	ResetAutomatedBackupConfig()
	ResetCrossInstanceReplicationConfig()
	ResetDeletionProtectionEnabled()
	ResetDesiredAutoCreatedEndpoints()
	ResetDesiredPscAutoConnections()
	ResetEngineConfigs()
	ResetEngineVersion()
	ResetGcsSource()
	ResetId()
	ResetKmsKey()
	ResetLabels()
	ResetMaintenancePolicy()
	ResetMaintenanceVersion()
	ResetManagedBackupSource()
	ResetMode()
	ResetNodeType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPersistenceConfig()
	ResetProject()
	ResetReplicaCount()
	ResetServerCaMode()
	ResetServerCaPool()
	ResetTimeouts()
	ResetTransitEncryptionMode()
	ResetZoneDistributionConfig()
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

// The jsii proxy struct for GoogleMemorystoreInstance
type jsiiProxy_GoogleMemorystoreInstance struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleMemorystoreInstance) AuthorizationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) AuthorizationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) AutomatedBackupConfig() GoogleMemorystoreInstanceAutomatedBackupConfigOutputReference {
	var returns GoogleMemorystoreInstanceAutomatedBackupConfigOutputReference
	_jsii_.Get(
		j,
		"automatedBackupConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) AutomatedBackupConfigInput() *GoogleMemorystoreInstanceAutomatedBackupConfig {
	var returns *GoogleMemorystoreInstanceAutomatedBackupConfig
	_jsii_.Get(
		j,
		"automatedBackupConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) AvailableMaintenanceVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableMaintenanceVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) BackupCollection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupCollection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) CrossInstanceReplicationConfig() GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference {
	var returns GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference
	_jsii_.Get(
		j,
		"crossInstanceReplicationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) CrossInstanceReplicationConfigInput() *GoogleMemorystoreInstanceCrossInstanceReplicationConfig {
	var returns *GoogleMemorystoreInstanceCrossInstanceReplicationConfig
	_jsii_.Get(
		j,
		"crossInstanceReplicationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) DeletionProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) DeletionProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) DesiredAutoCreatedEndpoints() GoogleMemorystoreInstanceDesiredAutoCreatedEndpointsList {
	var returns GoogleMemorystoreInstanceDesiredAutoCreatedEndpointsList
	_jsii_.Get(
		j,
		"desiredAutoCreatedEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) DesiredAutoCreatedEndpointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"desiredAutoCreatedEndpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) DesiredPscAutoConnections() GoogleMemorystoreInstanceDesiredPscAutoConnectionsList {
	var returns GoogleMemorystoreInstanceDesiredPscAutoConnectionsList
	_jsii_.Get(
		j,
		"desiredPscAutoConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) DesiredPscAutoConnectionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"desiredPscAutoConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) DiscoveryEndpoints() GoogleMemorystoreInstanceDiscoveryEndpointsList {
	var returns GoogleMemorystoreInstanceDiscoveryEndpointsList
	_jsii_.Get(
		j,
		"discoveryEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) EffectiveMaintenanceVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveMaintenanceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) Endpoints() GoogleMemorystoreInstanceEndpointsList {
	var returns GoogleMemorystoreInstanceEndpointsList
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) EngineConfigs() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"engineConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) EngineConfigsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"engineConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) GcsSource() GoogleMemorystoreInstanceGcsSourceOutputReference {
	var returns GoogleMemorystoreInstanceGcsSourceOutputReference
	_jsii_.Get(
		j,
		"gcsSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) GcsSourceInput() *GoogleMemorystoreInstanceGcsSource {
	var returns *GoogleMemorystoreInstanceGcsSource
	_jsii_.Get(
		j,
		"gcsSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) InstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) MaintenancePolicy() GoogleMemorystoreInstanceMaintenancePolicyOutputReference {
	var returns GoogleMemorystoreInstanceMaintenancePolicyOutputReference
	_jsii_.Get(
		j,
		"maintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) MaintenancePolicyInput() *GoogleMemorystoreInstanceMaintenancePolicy {
	var returns *GoogleMemorystoreInstanceMaintenancePolicy
	_jsii_.Get(
		j,
		"maintenancePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) MaintenanceSchedule() GoogleMemorystoreInstanceMaintenanceScheduleList {
	var returns GoogleMemorystoreInstanceMaintenanceScheduleList
	_jsii_.Get(
		j,
		"maintenanceSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) MaintenanceVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) MaintenanceVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) ManagedBackupSource() GoogleMemorystoreInstanceManagedBackupSourceOutputReference {
	var returns GoogleMemorystoreInstanceManagedBackupSourceOutputReference
	_jsii_.Get(
		j,
		"managedBackupSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) ManagedBackupSourceInput() *GoogleMemorystoreInstanceManagedBackupSource {
	var returns *GoogleMemorystoreInstanceManagedBackupSource
	_jsii_.Get(
		j,
		"managedBackupSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) ManagedServerCa() GoogleMemorystoreInstanceManagedServerCaList {
	var returns GoogleMemorystoreInstanceManagedServerCaList
	_jsii_.Get(
		j,
		"managedServerCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) NodeConfig() GoogleMemorystoreInstanceNodeConfigList {
	var returns GoogleMemorystoreInstanceNodeConfigList
	_jsii_.Get(
		j,
		"nodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) NodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) NodeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) PersistenceConfig() GoogleMemorystoreInstancePersistenceConfigOutputReference {
	var returns GoogleMemorystoreInstancePersistenceConfigOutputReference
	_jsii_.Get(
		j,
		"persistenceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) PersistenceConfigInput() *GoogleMemorystoreInstancePersistenceConfig {
	var returns *GoogleMemorystoreInstancePersistenceConfig
	_jsii_.Get(
		j,
		"persistenceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) PscAttachmentDetails() GoogleMemorystoreInstancePscAttachmentDetailsList {
	var returns GoogleMemorystoreInstancePscAttachmentDetailsList
	_jsii_.Get(
		j,
		"pscAttachmentDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) PscAutoConnections() GoogleMemorystoreInstancePscAutoConnectionsList {
	var returns GoogleMemorystoreInstancePscAutoConnectionsList
	_jsii_.Get(
		j,
		"pscAutoConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) ReplicaCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) ReplicaCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) ServerCaMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCaMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) ServerCaModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCaModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) ServerCaPool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCaPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) ServerCaPoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCaPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) ShardCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shardCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) ShardCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shardCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) StateInfo() GoogleMemorystoreInstanceStateInfoList {
	var returns GoogleMemorystoreInstanceStateInfoList
	_jsii_.Get(
		j,
		"stateInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) Timeouts() GoogleMemorystoreInstanceTimeoutsOutputReference {
	var returns GoogleMemorystoreInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) TransitEncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitEncryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) TransitEncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitEncryptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) ZoneDistributionConfig() GoogleMemorystoreInstanceZoneDistributionConfigOutputReference {
	var returns GoogleMemorystoreInstanceZoneDistributionConfigOutputReference
	_jsii_.Get(
		j,
		"zoneDistributionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstance) ZoneDistributionConfigInput() *GoogleMemorystoreInstanceZoneDistributionConfig {
	var returns *GoogleMemorystoreInstanceZoneDistributionConfig
	_jsii_.Get(
		j,
		"zoneDistributionConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_memorystore_instance google_memorystore_instance} Resource.
func NewGoogleMemorystoreInstance(scope constructs.Construct, id *string, config *GoogleMemorystoreInstanceConfig) GoogleMemorystoreInstance {
	_init_.Initialize()

	if err := validateNewGoogleMemorystoreInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMemorystoreInstance{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleMemorystoreInstance.GoogleMemorystoreInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_memorystore_instance google_memorystore_instance} Resource.
func NewGoogleMemorystoreInstance_Override(g GoogleMemorystoreInstance, scope constructs.Construct, id *string, config *GoogleMemorystoreInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleMemorystoreInstance.GoogleMemorystoreInstance",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetAuthorizationMode(val *string) {
	if err := j.validateSetAuthorizationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationMode",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetDeletionProtectionEnabled(val interface{}) {
	if err := j.validateSetDeletionProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetEngineConfigs(val *map[string]*string) {
	if err := j.validateSetEngineConfigsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineConfigs",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetInstanceId(val *string) {
	if err := j.validateSetInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetKmsKey(val *string) {
	if err := j.validateSetKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetMaintenanceVersion(val *string) {
	if err := j.validateSetMaintenanceVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetNodeType(val *string) {
	if err := j.validateSetNodeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeType",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetReplicaCount(val *float64) {
	if err := j.validateSetReplicaCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaCount",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetServerCaMode(val *string) {
	if err := j.validateSetServerCaModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverCaMode",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetServerCaPool(val *string) {
	if err := j.validateSetServerCaPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverCaPool",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetShardCount(val *float64) {
	if err := j.validateSetShardCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shardCount",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstance)SetTransitEncryptionMode(val *string) {
	if err := j.validateSetTransitEncryptionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitEncryptionMode",
		val,
	)
}

// Generates CDKTN code for importing a GoogleMemorystoreInstance resource upon running "cdktn plan <stack-name>".
func GoogleMemorystoreInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleMemorystoreInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleMemorystoreInstance.GoogleMemorystoreInstance",
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
func GoogleMemorystoreInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleMemorystoreInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleMemorystoreInstance.GoogleMemorystoreInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleMemorystoreInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleMemorystoreInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleMemorystoreInstance.GoogleMemorystoreInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleMemorystoreInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleMemorystoreInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleMemorystoreInstance.GoogleMemorystoreInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleMemorystoreInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleMemorystoreInstance.GoogleMemorystoreInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleMemorystoreInstance) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleMemorystoreInstance) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleMemorystoreInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleMemorystoreInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleMemorystoreInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleMemorystoreInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleMemorystoreInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleMemorystoreInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleMemorystoreInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleMemorystoreInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleMemorystoreInstance) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (g *jsiiProxy_GoogleMemorystoreInstance) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) PutAutomatedBackupConfig(value *GoogleMemorystoreInstanceAutomatedBackupConfig) {
	if err := g.validatePutAutomatedBackupConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutomatedBackupConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) PutCrossInstanceReplicationConfig(value *GoogleMemorystoreInstanceCrossInstanceReplicationConfig) {
	if err := g.validatePutCrossInstanceReplicationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCrossInstanceReplicationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) PutDesiredAutoCreatedEndpoints(value interface{}) {
	if err := g.validatePutDesiredAutoCreatedEndpointsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDesiredAutoCreatedEndpoints",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) PutDesiredPscAutoConnections(value interface{}) {
	if err := g.validatePutDesiredPscAutoConnectionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDesiredPscAutoConnections",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) PutGcsSource(value *GoogleMemorystoreInstanceGcsSource) {
	if err := g.validatePutGcsSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcsSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) PutMaintenancePolicy(value *GoogleMemorystoreInstanceMaintenancePolicy) {
	if err := g.validatePutMaintenancePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaintenancePolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) PutManagedBackupSource(value *GoogleMemorystoreInstanceManagedBackupSource) {
	if err := g.validatePutManagedBackupSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putManagedBackupSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) PutPersistenceConfig(value *GoogleMemorystoreInstancePersistenceConfig) {
	if err := g.validatePutPersistenceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPersistenceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) PutTimeouts(value *GoogleMemorystoreInstanceTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) PutZoneDistributionConfig(value *GoogleMemorystoreInstanceZoneDistributionConfig) {
	if err := g.validatePutZoneDistributionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putZoneDistributionConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetAuthorizationMode() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthorizationMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetAutomatedBackupConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAutomatedBackupConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetCrossInstanceReplicationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetCrossInstanceReplicationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetDeletionProtectionEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetDeletionProtectionEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetDesiredAutoCreatedEndpoints() {
	_jsii_.InvokeVoid(
		g,
		"resetDesiredAutoCreatedEndpoints",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetDesiredPscAutoConnections() {
	_jsii_.InvokeVoid(
		g,
		"resetDesiredPscAutoConnections",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetEngineConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetEngineConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetGcsSource() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetKmsKey() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetMaintenancePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenancePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetMaintenanceVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenanceVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetManagedBackupSource() {
	_jsii_.InvokeVoid(
		g,
		"resetManagedBackupSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetMode() {
	_jsii_.InvokeVoid(
		g,
		"resetMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetNodeType() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetPersistenceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPersistenceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetReplicaCount() {
	_jsii_.InvokeVoid(
		g,
		"resetReplicaCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetServerCaMode() {
	_jsii_.InvokeVoid(
		g,
		"resetServerCaMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetServerCaPool() {
	_jsii_.InvokeVoid(
		g,
		"resetServerCaPool",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetTransitEncryptionMode() {
	_jsii_.InvokeVoid(
		g,
		"resetTransitEncryptionMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ResetZoneDistributionConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetZoneDistributionConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMemorystoreInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMemorystoreInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMemorystoreInstance) With(mixins ...constructs.IMixin) constructs.IConstruct {
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


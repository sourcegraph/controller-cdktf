package memorystoreinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/memorystoreinstance/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/memorystore_instance google_memorystore_instance}.
type MemorystoreInstance interface {
	cdktf.TerraformResource
	AuthorizationMode() *string
	SetAuthorizationMode(val *string)
	AuthorizationModeInput() *string
	AutomatedBackupConfig() MemorystoreInstanceAutomatedBackupConfigOutputReference
	AutomatedBackupConfigInput() *MemorystoreInstanceAutomatedBackupConfig
	AvailableMaintenanceVersions() *[]*string
	BackupCollection() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	CreateTime() *string
	CrossInstanceReplicationConfig() MemorystoreInstanceCrossInstanceReplicationConfigOutputReference
	CrossInstanceReplicationConfigInput() *MemorystoreInstanceCrossInstanceReplicationConfig
	DeletionProtectionEnabled() any
	SetDeletionProtectionEnabled(val any)
	DeletionProtectionEnabledInput() any
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DesiredAutoCreatedEndpoints() MemorystoreInstanceDesiredAutoCreatedEndpointsList
	DesiredAutoCreatedEndpointsInput() any
	DesiredPscAutoConnections() MemorystoreInstanceDesiredPscAutoConnectionsList
	DesiredPscAutoConnectionsInput() any
	DiscoveryEndpoints() MemorystoreInstanceDiscoveryEndpointsList
	EffectiveLabels() cdktf.StringMap
	EffectiveMaintenanceVersion() *string
	Endpoints() MemorystoreInstanceEndpointsList
	EngineConfigs() *map[string]*string
	SetEngineConfigs(val *map[string]*string)
	EngineConfigsInput() *map[string]*string
	EngineVersion() *string
	SetEngineVersion(val *string)
	EngineVersionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GcsSource() MemorystoreInstanceGcsSourceOutputReference
	GcsSourceInput() *MemorystoreInstanceGcsSource
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
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaintenancePolicy() MemorystoreInstanceMaintenancePolicyOutputReference
	MaintenancePolicyInput() *MemorystoreInstanceMaintenancePolicy
	MaintenanceSchedule() MemorystoreInstanceMaintenanceScheduleList
	MaintenanceVersion() *string
	SetMaintenanceVersion(val *string)
	MaintenanceVersionInput() *string
	ManagedBackupSource() MemorystoreInstanceManagedBackupSourceOutputReference
	ManagedBackupSourceInput() *MemorystoreInstanceManagedBackupSource
	ManagedServerCa() MemorystoreInstanceManagedServerCaList
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	NodeConfig() MemorystoreInstanceNodeConfigList
	NodeType() *string
	SetNodeType(val *string)
	NodeTypeInput() *string
	PersistenceConfig() MemorystoreInstancePersistenceConfigOutputReference
	PersistenceConfigInput() *MemorystoreInstancePersistenceConfig
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	PscAttachmentDetails() MemorystoreInstancePscAttachmentDetailsList
	PscAutoConnections() MemorystoreInstancePscAutoConnectionsList
	// Experimental.
	RawOverrides() any
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
	StateInfo() MemorystoreInstanceStateInfoList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MemorystoreInstanceTimeoutsOutputReference
	TimeoutsInput() any
	TransitEncryptionMode() *string
	SetTransitEncryptionMode(val *string)
	TransitEncryptionModeInput() *string
	Uid() *string
	UpdateTime() *string
	ZoneDistributionConfig() MemorystoreInstanceZoneDistributionConfigOutputReference
	ZoneDistributionConfigInput() *MemorystoreInstanceZoneDistributionConfig
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value any)
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
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
	HasResourceMove() any
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index any)
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAutomatedBackupConfig(value *MemorystoreInstanceAutomatedBackupConfig)
	PutCrossInstanceReplicationConfig(value *MemorystoreInstanceCrossInstanceReplicationConfig)
	PutDesiredAutoCreatedEndpoints(value any)
	PutDesiredPscAutoConnections(value any)
	PutGcsSource(value *MemorystoreInstanceGcsSource)
	PutMaintenancePolicy(value *MemorystoreInstanceMaintenancePolicy)
	PutManagedBackupSource(value *MemorystoreInstanceManagedBackupSource)
	PutPersistenceConfig(value *MemorystoreInstancePersistenceConfig)
	PutTimeouts(value *MemorystoreInstanceTimeouts)
	PutZoneDistributionConfig(value *MemorystoreInstanceZoneDistributionConfig)
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
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for MemorystoreInstance
type jsiiProxy_MemorystoreInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MemorystoreInstance) AuthorizationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) AuthorizationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) AutomatedBackupConfig() MemorystoreInstanceAutomatedBackupConfigOutputReference {
	var returns MemorystoreInstanceAutomatedBackupConfigOutputReference
	_jsii_.Get(
		j,
		"automatedBackupConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) AutomatedBackupConfigInput() *MemorystoreInstanceAutomatedBackupConfig {
	var returns *MemorystoreInstanceAutomatedBackupConfig
	_jsii_.Get(
		j,
		"automatedBackupConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) AvailableMaintenanceVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableMaintenanceVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) BackupCollection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupCollection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) CrossInstanceReplicationConfig() MemorystoreInstanceCrossInstanceReplicationConfigOutputReference {
	var returns MemorystoreInstanceCrossInstanceReplicationConfigOutputReference
	_jsii_.Get(
		j,
		"crossInstanceReplicationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) CrossInstanceReplicationConfigInput() *MemorystoreInstanceCrossInstanceReplicationConfig {
	var returns *MemorystoreInstanceCrossInstanceReplicationConfig
	_jsii_.Get(
		j,
		"crossInstanceReplicationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) DeletionProtectionEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"deletionProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) DeletionProtectionEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"deletionProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) DesiredAutoCreatedEndpoints() MemorystoreInstanceDesiredAutoCreatedEndpointsList {
	var returns MemorystoreInstanceDesiredAutoCreatedEndpointsList
	_jsii_.Get(
		j,
		"desiredAutoCreatedEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) DesiredAutoCreatedEndpointsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"desiredAutoCreatedEndpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) DesiredPscAutoConnections() MemorystoreInstanceDesiredPscAutoConnectionsList {
	var returns MemorystoreInstanceDesiredPscAutoConnectionsList
	_jsii_.Get(
		j,
		"desiredPscAutoConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) DesiredPscAutoConnectionsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"desiredPscAutoConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) DiscoveryEndpoints() MemorystoreInstanceDiscoveryEndpointsList {
	var returns MemorystoreInstanceDiscoveryEndpointsList
	_jsii_.Get(
		j,
		"discoveryEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) EffectiveMaintenanceVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveMaintenanceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) Endpoints() MemorystoreInstanceEndpointsList {
	var returns MemorystoreInstanceEndpointsList
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) EngineConfigs() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"engineConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) EngineConfigsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"engineConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) GcsSource() MemorystoreInstanceGcsSourceOutputReference {
	var returns MemorystoreInstanceGcsSourceOutputReference
	_jsii_.Get(
		j,
		"gcsSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) GcsSourceInput() *MemorystoreInstanceGcsSource {
	var returns *MemorystoreInstanceGcsSource
	_jsii_.Get(
		j,
		"gcsSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) InstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) MaintenancePolicy() MemorystoreInstanceMaintenancePolicyOutputReference {
	var returns MemorystoreInstanceMaintenancePolicyOutputReference
	_jsii_.Get(
		j,
		"maintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) MaintenancePolicyInput() *MemorystoreInstanceMaintenancePolicy {
	var returns *MemorystoreInstanceMaintenancePolicy
	_jsii_.Get(
		j,
		"maintenancePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) MaintenanceSchedule() MemorystoreInstanceMaintenanceScheduleList {
	var returns MemorystoreInstanceMaintenanceScheduleList
	_jsii_.Get(
		j,
		"maintenanceSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) MaintenanceVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) MaintenanceVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) ManagedBackupSource() MemorystoreInstanceManagedBackupSourceOutputReference {
	var returns MemorystoreInstanceManagedBackupSourceOutputReference
	_jsii_.Get(
		j,
		"managedBackupSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) ManagedBackupSourceInput() *MemorystoreInstanceManagedBackupSource {
	var returns *MemorystoreInstanceManagedBackupSource
	_jsii_.Get(
		j,
		"managedBackupSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) ManagedServerCa() MemorystoreInstanceManagedServerCaList {
	var returns MemorystoreInstanceManagedServerCaList
	_jsii_.Get(
		j,
		"managedServerCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) NodeConfig() MemorystoreInstanceNodeConfigList {
	var returns MemorystoreInstanceNodeConfigList
	_jsii_.Get(
		j,
		"nodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) NodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) NodeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) PersistenceConfig() MemorystoreInstancePersistenceConfigOutputReference {
	var returns MemorystoreInstancePersistenceConfigOutputReference
	_jsii_.Get(
		j,
		"persistenceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) PersistenceConfigInput() *MemorystoreInstancePersistenceConfig {
	var returns *MemorystoreInstancePersistenceConfig
	_jsii_.Get(
		j,
		"persistenceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) PscAttachmentDetails() MemorystoreInstancePscAttachmentDetailsList {
	var returns MemorystoreInstancePscAttachmentDetailsList
	_jsii_.Get(
		j,
		"pscAttachmentDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) PscAutoConnections() MemorystoreInstancePscAutoConnectionsList {
	var returns MemorystoreInstancePscAutoConnectionsList
	_jsii_.Get(
		j,
		"pscAutoConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) ReplicaCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) ReplicaCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) ServerCaMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCaMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) ServerCaModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCaModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) ServerCaPool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCaPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) ServerCaPoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCaPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) ShardCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shardCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) ShardCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shardCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) StateInfo() MemorystoreInstanceStateInfoList {
	var returns MemorystoreInstanceStateInfoList
	_jsii_.Get(
		j,
		"stateInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) Timeouts() MemorystoreInstanceTimeoutsOutputReference {
	var returns MemorystoreInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) TransitEncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitEncryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) TransitEncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitEncryptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) ZoneDistributionConfig() MemorystoreInstanceZoneDistributionConfigOutputReference {
	var returns MemorystoreInstanceZoneDistributionConfigOutputReference
	_jsii_.Get(
		j,
		"zoneDistributionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstance) ZoneDistributionConfigInput() *MemorystoreInstanceZoneDistributionConfig {
	var returns *MemorystoreInstanceZoneDistributionConfig
	_jsii_.Get(
		j,
		"zoneDistributionConfigInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/memorystore_instance google_memorystore_instance} Resource.
func NewMemorystoreInstance(scope constructs.Construct, id *string, config *MemorystoreInstanceConfig) MemorystoreInstance {
	_init_.Initialize()

	if err := validateNewMemorystoreInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MemorystoreInstance{}

	_jsii_.Create(
		"@cdktf/provider-google.memorystoreInstance.MemorystoreInstance",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/memorystore_instance google_memorystore_instance} Resource.
func NewMemorystoreInstance_Override(m MemorystoreInstance, scope constructs.Construct, id *string, config *MemorystoreInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.memorystoreInstance.MemorystoreInstance",
		[]any{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetAuthorizationMode(val *string) {
	if err := j.validateSetAuthorizationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationMode",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetDeletionProtectionEnabled(val any) {
	if err := j.validateSetDeletionProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetEngineConfigs(val *map[string]*string) {
	if err := j.validateSetEngineConfigsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineConfigs",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetInstanceId(val *string) {
	if err := j.validateSetInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetKmsKey(val *string) {
	if err := j.validateSetKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetMaintenanceVersion(val *string) {
	if err := j.validateSetMaintenanceVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceVersion",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetNodeType(val *string) {
	if err := j.validateSetNodeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeType",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetReplicaCount(val *float64) {
	if err := j.validateSetReplicaCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaCount",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetServerCaMode(val *string) {
	if err := j.validateSetServerCaModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverCaMode",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetServerCaPool(val *string) {
	if err := j.validateSetServerCaPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverCaPool",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetShardCount(val *float64) {
	if err := j.validateSetShardCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shardCount",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstance) SetTransitEncryptionMode(val *string) {
	if err := j.validateSetTransitEncryptionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitEncryptionMode",
		val,
	)
}

// Generates CDKTF code for importing a MemorystoreInstance resource upon running "cdktf plan <stack-name>".
func MemorystoreInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMemorystoreInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.memorystoreInstance.MemorystoreInstance",
		"generateConfigForImport",
		[]any{scope, importToId, importFromId, provider},
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
func MemorystoreInstance_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateMemorystoreInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.memorystoreInstance.MemorystoreInstance",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func MemorystoreInstance_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateMemorystoreInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.memorystoreInstance.MemorystoreInstance",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func MemorystoreInstance_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateMemorystoreInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.memorystoreInstance.MemorystoreInstance",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func MemorystoreInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.memorystoreInstance.MemorystoreInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MemorystoreInstance) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (m *jsiiProxy_MemorystoreInstance) AddOverride(path *string, value any) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]any{path, value},
	)
}

func (m *jsiiProxy_MemorystoreInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstance) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstance) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstance) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstance) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]any{id, provider},
	)
}

func (m *jsiiProxy_MemorystoreInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstance) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]any{id},
	)
}

func (m *jsiiProxy_MemorystoreInstance) MoveTo(moveTarget *string, index any) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (m *jsiiProxy_MemorystoreInstance) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]any{id},
	)
}

func (m *jsiiProxy_MemorystoreInstance) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (m *jsiiProxy_MemorystoreInstance) PutAutomatedBackupConfig(value *MemorystoreInstanceAutomatedBackupConfig) {
	if err := m.validatePutAutomatedBackupConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAutomatedBackupConfig",
		[]any{value},
	)
}

func (m *jsiiProxy_MemorystoreInstance) PutCrossInstanceReplicationConfig(value *MemorystoreInstanceCrossInstanceReplicationConfig) {
	if err := m.validatePutCrossInstanceReplicationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCrossInstanceReplicationConfig",
		[]any{value},
	)
}

func (m *jsiiProxy_MemorystoreInstance) PutDesiredAutoCreatedEndpoints(value any) {
	if err := m.validatePutDesiredAutoCreatedEndpointsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDesiredAutoCreatedEndpoints",
		[]any{value},
	)
}

func (m *jsiiProxy_MemorystoreInstance) PutDesiredPscAutoConnections(value any) {
	if err := m.validatePutDesiredPscAutoConnectionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDesiredPscAutoConnections",
		[]any{value},
	)
}

func (m *jsiiProxy_MemorystoreInstance) PutGcsSource(value *MemorystoreInstanceGcsSource) {
	if err := m.validatePutGcsSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putGcsSource",
		[]any{value},
	)
}

func (m *jsiiProxy_MemorystoreInstance) PutMaintenancePolicy(value *MemorystoreInstanceMaintenancePolicy) {
	if err := m.validatePutMaintenancePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMaintenancePolicy",
		[]any{value},
	)
}

func (m *jsiiProxy_MemorystoreInstance) PutManagedBackupSource(value *MemorystoreInstanceManagedBackupSource) {
	if err := m.validatePutManagedBackupSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putManagedBackupSource",
		[]any{value},
	)
}

func (m *jsiiProxy_MemorystoreInstance) PutPersistenceConfig(value *MemorystoreInstancePersistenceConfig) {
	if err := m.validatePutPersistenceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPersistenceConfig",
		[]any{value},
	)
}

func (m *jsiiProxy_MemorystoreInstance) PutTimeouts(value *MemorystoreInstanceTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]any{value},
	)
}

func (m *jsiiProxy_MemorystoreInstance) PutZoneDistributionConfig(value *MemorystoreInstanceZoneDistributionConfig) {
	if err := m.validatePutZoneDistributionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putZoneDistributionConfig",
		[]any{value},
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetAuthorizationMode() {
	_jsii_.InvokeVoid(
		m,
		"resetAuthorizationMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetAutomatedBackupConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetAutomatedBackupConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetCrossInstanceReplicationConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetCrossInstanceReplicationConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetDeletionProtectionEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetDeletionProtectionEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetDesiredAutoCreatedEndpoints() {
	_jsii_.InvokeVoid(
		m,
		"resetDesiredAutoCreatedEndpoints",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetDesiredPscAutoConnections() {
	_jsii_.InvokeVoid(
		m,
		"resetDesiredPscAutoConnections",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetEngineConfigs() {
	_jsii_.InvokeVoid(
		m,
		"resetEngineConfigs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetGcsSource() {
	_jsii_.InvokeVoid(
		m,
		"resetGcsSource",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetKmsKey() {
	_jsii_.InvokeVoid(
		m,
		"resetKmsKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetLabels() {
	_jsii_.InvokeVoid(
		m,
		"resetLabels",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetMaintenancePolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetMaintenancePolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetMaintenanceVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetMaintenanceVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetManagedBackupSource() {
	_jsii_.InvokeVoid(
		m,
		"resetManagedBackupSource",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetMode() {
	_jsii_.InvokeVoid(
		m,
		"resetMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetNodeType() {
	_jsii_.InvokeVoid(
		m,
		"resetNodeType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetPersistenceConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetPersistenceConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetProject() {
	_jsii_.InvokeVoid(
		m,
		"resetProject",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetReplicaCount() {
	_jsii_.InvokeVoid(
		m,
		"resetReplicaCount",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetServerCaMode() {
	_jsii_.InvokeVoid(
		m,
		"resetServerCaMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetServerCaPool() {
	_jsii_.InvokeVoid(
		m,
		"resetServerCaPool",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetTransitEncryptionMode() {
	_jsii_.InvokeVoid(
		m,
		"resetTransitEncryptionMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) ResetZoneDistributionConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetZoneDistributionConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstance) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstance) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstance) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstance) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstance) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

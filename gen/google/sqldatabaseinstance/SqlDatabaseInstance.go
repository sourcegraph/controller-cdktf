package sqldatabaseinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/sqldatabaseinstance/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/sql_database_instance google_sql_database_instance}.
type SqlDatabaseInstance interface {
	cdktn.TerraformResource
	AvailableMaintenanceVersions() *[]*string
	BackupdrBackup() *string
	SetBackupdrBackup(val *string)
	BackupdrBackupInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Clone() SqlDatabaseInstanceCloneOutputReference
	CloneInput() *SqlDatabaseInstanceClone
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionName() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DatabaseVersion() *string
	SetDatabaseVersion(val *string)
	DatabaseVersionInput() *string
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	DeletionProtectionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DnsName() *string
	DnsNames() SqlDatabaseInstanceDnsNamesList
	EncryptionKeyName() *string
	SetEncryptionKeyName(val *string)
	EncryptionKeyNameInput() *string
	FinalBackupDescription() *string
	SetFinalBackupDescription(val *string)
	FinalBackupDescriptionInput() *string
	FirstIpAddress() *string
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
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	IpAddress() SqlDatabaseInstanceIpAddressList
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MaintenanceVersion() *string
	SetMaintenanceVersion(val *string)
	MaintenanceVersionInput() *string
	MasterInstanceName() *string
	SetMasterInstanceName(val *string)
	MasterInstanceNameInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NodeCount() *float64
	SetNodeCount(val *float64)
	NodeCountInput() *float64
	PointInTimeRestoreContext() SqlDatabaseInstancePointInTimeRestoreContextOutputReference
	PointInTimeRestoreContextInput() *SqlDatabaseInstancePointInTimeRestoreContext
	PrivateIpAddress() *string
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
	PscServiceAttachmentLink() *string
	PublicIpAddress() *string
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ReplicaConfiguration() SqlDatabaseInstanceReplicaConfigurationOutputReference
	ReplicaConfigurationInput() *SqlDatabaseInstanceReplicaConfiguration
	ReplicaNames() *[]*string
	SetReplicaNames(val *[]*string)
	ReplicaNamesInput() *[]*string
	ReplicationCluster() SqlDatabaseInstanceReplicationClusterOutputReference
	ReplicationClusterInput() *SqlDatabaseInstanceReplicationCluster
	RestoreBackupContext() SqlDatabaseInstanceRestoreBackupContextOutputReference
	RestoreBackupContextInput() *SqlDatabaseInstanceRestoreBackupContext
	RootPassword() *string
	SetRootPassword(val *string)
	RootPasswordInput() *string
	RootPasswordWo() *string
	SetRootPasswordWo(val *string)
	RootPasswordWoInput() *string
	RootPasswordWoVersion() *string
	SetRootPasswordWoVersion(val *string)
	RootPasswordWoVersionInput() *string
	SelfLink() *string
	ServerCaCert() SqlDatabaseInstanceServerCaCertList
	ServiceAccountEmailAddress() *string
	Settings() SqlDatabaseInstanceSettingsOutputReference
	SettingsInput() *SqlDatabaseInstanceSettings
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SqlDatabaseInstanceTimeoutsOutputReference
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
	PutClone(value *SqlDatabaseInstanceClone)
	PutPointInTimeRestoreContext(value *SqlDatabaseInstancePointInTimeRestoreContext)
	PutReplicaConfiguration(value *SqlDatabaseInstanceReplicaConfiguration)
	PutReplicationCluster(value *SqlDatabaseInstanceReplicationCluster)
	PutRestoreBackupContext(value *SqlDatabaseInstanceRestoreBackupContext)
	PutSettings(value *SqlDatabaseInstanceSettings)
	PutTimeouts(value *SqlDatabaseInstanceTimeouts)
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
	ResetBackupdrBackup()
	ResetClone()
	ResetDeletionProtection()
	ResetEncryptionKeyName()
	ResetFinalBackupDescription()
	ResetId()
	ResetInstanceType()
	ResetMaintenanceVersion()
	ResetMasterInstanceName()
	ResetName()
	ResetNodeCount()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPointInTimeRestoreContext()
	ResetProject()
	ResetRegion()
	ResetReplicaConfiguration()
	ResetReplicaNames()
	ResetReplicationCluster()
	ResetRestoreBackupContext()
	ResetRootPassword()
	ResetRootPasswordWo()
	ResetRootPasswordWoVersion()
	ResetSettings()
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

// The jsii proxy struct for SqlDatabaseInstance
type jsiiProxy_SqlDatabaseInstance struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_SqlDatabaseInstance) AvailableMaintenanceVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableMaintenanceVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) BackupdrBackup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupdrBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) BackupdrBackupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupdrBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) Clone() SqlDatabaseInstanceCloneOutputReference {
	var returns SqlDatabaseInstanceCloneOutputReference
	_jsii_.Get(
		j,
		"clone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) CloneInput() *SqlDatabaseInstanceClone {
	var returns *SqlDatabaseInstanceClone
	_jsii_.Get(
		j,
		"cloneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) ConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) DatabaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) DatabaseVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) DnsNames() SqlDatabaseInstanceDnsNamesList {
	var returns SqlDatabaseInstanceDnsNamesList
	_jsii_.Get(
		j,
		"dnsNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) EncryptionKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) EncryptionKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) FinalBackupDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalBackupDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) FinalBackupDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalBackupDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) FirstIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) IpAddress() SqlDatabaseInstanceIpAddressList {
	var returns SqlDatabaseInstanceIpAddressList
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) MaintenanceVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) MaintenanceVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) MasterInstanceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterInstanceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) MasterInstanceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterInstanceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) PointInTimeRestoreContext() SqlDatabaseInstancePointInTimeRestoreContextOutputReference {
	var returns SqlDatabaseInstancePointInTimeRestoreContextOutputReference
	_jsii_.Get(
		j,
		"pointInTimeRestoreContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) PointInTimeRestoreContextInput() *SqlDatabaseInstancePointInTimeRestoreContext {
	var returns *SqlDatabaseInstancePointInTimeRestoreContext
	_jsii_.Get(
		j,
		"pointInTimeRestoreContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) PrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) PscServiceAttachmentLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pscServiceAttachmentLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) PublicIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) ReplicaConfiguration() SqlDatabaseInstanceReplicaConfigurationOutputReference {
	var returns SqlDatabaseInstanceReplicaConfigurationOutputReference
	_jsii_.Get(
		j,
		"replicaConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) ReplicaConfigurationInput() *SqlDatabaseInstanceReplicaConfiguration {
	var returns *SqlDatabaseInstanceReplicaConfiguration
	_jsii_.Get(
		j,
		"replicaConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) ReplicaNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicaNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) ReplicaNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicaNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) ReplicationCluster() SqlDatabaseInstanceReplicationClusterOutputReference {
	var returns SqlDatabaseInstanceReplicationClusterOutputReference
	_jsii_.Get(
		j,
		"replicationCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) ReplicationClusterInput() *SqlDatabaseInstanceReplicationCluster {
	var returns *SqlDatabaseInstanceReplicationCluster
	_jsii_.Get(
		j,
		"replicationClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) RestoreBackupContext() SqlDatabaseInstanceRestoreBackupContextOutputReference {
	var returns SqlDatabaseInstanceRestoreBackupContextOutputReference
	_jsii_.Get(
		j,
		"restoreBackupContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) RestoreBackupContextInput() *SqlDatabaseInstanceRestoreBackupContext {
	var returns *SqlDatabaseInstanceRestoreBackupContext
	_jsii_.Get(
		j,
		"restoreBackupContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) RootPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) RootPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) RootPasswordWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPasswordWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) RootPasswordWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPasswordWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) RootPasswordWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPasswordWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) RootPasswordWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPasswordWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) ServerCaCert() SqlDatabaseInstanceServerCaCertList {
	var returns SqlDatabaseInstanceServerCaCertList
	_jsii_.Get(
		j,
		"serverCaCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) ServiceAccountEmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) Settings() SqlDatabaseInstanceSettingsOutputReference {
	var returns SqlDatabaseInstanceSettingsOutputReference
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) SettingsInput() *SqlDatabaseInstanceSettings {
	var returns *SqlDatabaseInstanceSettings
	_jsii_.Get(
		j,
		"settingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) Timeouts() SqlDatabaseInstanceTimeoutsOutputReference {
	var returns SqlDatabaseInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/sql_database_instance google_sql_database_instance} Resource.
func NewSqlDatabaseInstance(scope constructs.Construct, id *string, config *SqlDatabaseInstanceConfig) SqlDatabaseInstance {
	_init_.Initialize()

	if err := validateNewSqlDatabaseInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqlDatabaseInstance{}

	_jsii_.Create(
		"@cdktn/provider-google.sqlDatabaseInstance.SqlDatabaseInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/sql_database_instance google_sql_database_instance} Resource.
func NewSqlDatabaseInstance_Override(s SqlDatabaseInstance, scope constructs.Construct, id *string, config *SqlDatabaseInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.sqlDatabaseInstance.SqlDatabaseInstance",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetBackupdrBackup(val *string) {
	if err := j.validateSetBackupdrBackupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupdrBackup",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetDatabaseVersion(val *string) {
	if err := j.validateSetDatabaseVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseVersion",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetEncryptionKeyName(val *string) {
	if err := j.validateSetEncryptionKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionKeyName",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetFinalBackupDescription(val *string) {
	if err := j.validateSetFinalBackupDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalBackupDescription",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetMaintenanceVersion(val *string) {
	if err := j.validateSetMaintenanceVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceVersion",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetMasterInstanceName(val *string) {
	if err := j.validateSetMasterInstanceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterInstanceName",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetReplicaNames(val *[]*string) {
	if err := j.validateSetReplicaNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaNames",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetRootPassword(val *string) {
	if err := j.validateSetRootPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootPassword",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetRootPasswordWo(val *string) {
	if err := j.validateSetRootPasswordWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootPasswordWo",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstance)SetRootPasswordWoVersion(val *string) {
	if err := j.validateSetRootPasswordWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootPasswordWoVersion",
		val,
	)
}

// Generates CDKTN code for importing a SqlDatabaseInstance resource upon running "cdktn plan <stack-name>".
func SqlDatabaseInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateSqlDatabaseInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.sqlDatabaseInstance.SqlDatabaseInstance",
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
func SqlDatabaseInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSqlDatabaseInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.sqlDatabaseInstance.SqlDatabaseInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SqlDatabaseInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSqlDatabaseInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.sqlDatabaseInstance.SqlDatabaseInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SqlDatabaseInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSqlDatabaseInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.sqlDatabaseInstance.SqlDatabaseInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SqlDatabaseInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.sqlDatabaseInstance.SqlDatabaseInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SqlDatabaseInstance) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstance) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstance) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstance) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstance) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstance) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := s.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstance) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) PutClone(value *SqlDatabaseInstanceClone) {
	if err := s.validatePutCloneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putClone",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) PutPointInTimeRestoreContext(value *SqlDatabaseInstancePointInTimeRestoreContext) {
	if err := s.validatePutPointInTimeRestoreContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPointInTimeRestoreContext",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) PutReplicaConfiguration(value *SqlDatabaseInstanceReplicaConfiguration) {
	if err := s.validatePutReplicaConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putReplicaConfiguration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) PutReplicationCluster(value *SqlDatabaseInstanceReplicationCluster) {
	if err := s.validatePutReplicationClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putReplicationCluster",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) PutRestoreBackupContext(value *SqlDatabaseInstanceRestoreBackupContext) {
	if err := s.validatePutRestoreBackupContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRestoreBackupContext",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) PutSettings(value *SqlDatabaseInstanceSettings) {
	if err := s.validatePutSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) PutTimeouts(value *SqlDatabaseInstanceTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := s.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetBackupdrBackup() {
	_jsii_.InvokeVoid(
		s,
		"resetBackupdrBackup",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetClone() {
	_jsii_.InvokeVoid(
		s,
		"resetClone",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		s,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetEncryptionKeyName() {
	_jsii_.InvokeVoid(
		s,
		"resetEncryptionKeyName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetFinalBackupDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetFinalBackupDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetInstanceType() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetMaintenanceVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetMaintenanceVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetMasterInstanceName() {
	_jsii_.InvokeVoid(
		s,
		"resetMasterInstanceName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetNodeCount() {
	_jsii_.InvokeVoid(
		s,
		"resetNodeCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetPointInTimeRestoreContext() {
	_jsii_.InvokeVoid(
		s,
		"resetPointInTimeRestoreContext",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetProject() {
	_jsii_.InvokeVoid(
		s,
		"resetProject",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetRegion() {
	_jsii_.InvokeVoid(
		s,
		"resetRegion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetReplicaConfiguration() {
	_jsii_.InvokeVoid(
		s,
		"resetReplicaConfiguration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetReplicaNames() {
	_jsii_.InvokeVoid(
		s,
		"resetReplicaNames",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetReplicationCluster() {
	_jsii_.InvokeVoid(
		s,
		"resetReplicationCluster",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetRestoreBackupContext() {
	_jsii_.InvokeVoid(
		s,
		"resetRestoreBackupContext",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetRootPassword() {
	_jsii_.InvokeVoid(
		s,
		"resetRootPassword",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetRootPasswordWo() {
	_jsii_.InvokeVoid(
		s,
		"resetRootPasswordWo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetRootPasswordWoVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetRootPasswordWoVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstance) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		s,
		"with",
		args,
		&returns,
	)

	return returns
}


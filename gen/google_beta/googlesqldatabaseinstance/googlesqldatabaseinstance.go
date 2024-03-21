package googlesqldatabaseinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlesqldatabaseinstance/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_sql_database_instance google_sql_database_instance}.
type GoogleSqlDatabaseInstance interface {
	cdktf.TerraformResource
	AvailableMaintenanceVersions() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Clone() GoogleSqlDatabaseInstanceCloneOutputReference
	CloneInput() *GoogleSqlDatabaseInstanceClone
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
	EncryptionKeyName() *string
	SetEncryptionKeyName(val *string)
	EncryptionKeyNameInput() *string
	FirstIpAddress() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
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
	IpAddress() GoogleSqlDatabaseInstanceIpAddressList
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	PrivateIpAddress() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicIpAddress() *string
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ReplicaConfiguration() GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference
	ReplicaConfigurationInput() *GoogleSqlDatabaseInstanceReplicaConfiguration
	RestoreBackupContext() GoogleSqlDatabaseInstanceRestoreBackupContextOutputReference
	RestoreBackupContextInput() *GoogleSqlDatabaseInstanceRestoreBackupContext
	RootPassword() *string
	SetRootPassword(val *string)
	RootPasswordInput() *string
	SelfLink() *string
	ServerCaCert() GoogleSqlDatabaseInstanceServerCaCertList
	ServiceAccountEmailAddress() *string
	Settings() GoogleSqlDatabaseInstanceSettingsOutputReference
	SettingsInput() *GoogleSqlDatabaseInstanceSettings
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleSqlDatabaseInstanceTimeoutsOutputReference
	TimeoutsInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
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
	HasResourceMove() interface{}
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
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutClone(value *GoogleSqlDatabaseInstanceClone)
	PutReplicaConfiguration(value *GoogleSqlDatabaseInstanceReplicaConfiguration)
	PutRestoreBackupContext(value *GoogleSqlDatabaseInstanceRestoreBackupContext)
	PutSettings(value *GoogleSqlDatabaseInstanceSettings)
	PutTimeouts(value *GoogleSqlDatabaseInstanceTimeouts)
	ResetClone()
	ResetDeletionProtection()
	ResetEncryptionKeyName()
	ResetId()
	ResetInstanceType()
	ResetMaintenanceVersion()
	ResetMasterInstanceName()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRegion()
	ResetReplicaConfiguration()
	ResetRestoreBackupContext()
	ResetRootPassword()
	ResetSettings()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleSqlDatabaseInstance
type jsiiProxy_GoogleSqlDatabaseInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) AvailableMaintenanceVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableMaintenanceVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) Clone() GoogleSqlDatabaseInstanceCloneOutputReference {
	var returns GoogleSqlDatabaseInstanceCloneOutputReference
	_jsii_.Get(
		j,
		"clone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) CloneInput() *GoogleSqlDatabaseInstanceClone {
	var returns *GoogleSqlDatabaseInstanceClone
	_jsii_.Get(
		j,
		"cloneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) ConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) DatabaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) DatabaseVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) EncryptionKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) EncryptionKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) FirstIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) IpAddress() GoogleSqlDatabaseInstanceIpAddressList {
	var returns GoogleSqlDatabaseInstanceIpAddressList
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) MaintenanceVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) MaintenanceVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) MasterInstanceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterInstanceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) MasterInstanceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterInstanceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) PrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) PublicIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) ReplicaConfiguration() GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference {
	var returns GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference
	_jsii_.Get(
		j,
		"replicaConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) ReplicaConfigurationInput() *GoogleSqlDatabaseInstanceReplicaConfiguration {
	var returns *GoogleSqlDatabaseInstanceReplicaConfiguration
	_jsii_.Get(
		j,
		"replicaConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) RestoreBackupContext() GoogleSqlDatabaseInstanceRestoreBackupContextOutputReference {
	var returns GoogleSqlDatabaseInstanceRestoreBackupContextOutputReference
	_jsii_.Get(
		j,
		"restoreBackupContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) RestoreBackupContextInput() *GoogleSqlDatabaseInstanceRestoreBackupContext {
	var returns *GoogleSqlDatabaseInstanceRestoreBackupContext
	_jsii_.Get(
		j,
		"restoreBackupContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) RootPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) RootPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) ServerCaCert() GoogleSqlDatabaseInstanceServerCaCertList {
	var returns GoogleSqlDatabaseInstanceServerCaCertList
	_jsii_.Get(
		j,
		"serverCaCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) ServiceAccountEmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) Settings() GoogleSqlDatabaseInstanceSettingsOutputReference {
	var returns GoogleSqlDatabaseInstanceSettingsOutputReference
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) SettingsInput() *GoogleSqlDatabaseInstanceSettings {
	var returns *GoogleSqlDatabaseInstanceSettings
	_jsii_.Get(
		j,
		"settingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) Timeouts() GoogleSqlDatabaseInstanceTimeoutsOutputReference {
	var returns GoogleSqlDatabaseInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_sql_database_instance google_sql_database_instance} Resource.
func NewGoogleSqlDatabaseInstance(scope constructs.Construct, id *string, config *GoogleSqlDatabaseInstanceConfig) GoogleSqlDatabaseInstance {
	_init_.Initialize()

	if err := validateNewGoogleSqlDatabaseInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSqlDatabaseInstance{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_sql_database_instance google_sql_database_instance} Resource.
func NewGoogleSqlDatabaseInstance_Override(g GoogleSqlDatabaseInstance, scope constructs.Construct, id *string, config *GoogleSqlDatabaseInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstance",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance)SetDatabaseVersion(val *string) {
	if err := j.validateSetDatabaseVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance)SetEncryptionKeyName(val *string) {
	if err := j.validateSetEncryptionKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionKeyName",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance)SetMaintenanceVersion(val *string) {
	if err := j.validateSetMaintenanceVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance)SetMasterInstanceName(val *string) {
	if err := j.validateSetMasterInstanceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterInstanceName",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstance)SetRootPassword(val *string) {
	if err := j.validateSetRootPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootPassword",
		val,
	)
}

// Generates CDKTF code for importing a GoogleSqlDatabaseInstance resource upon running "cdktf plan <stack-name>".
func GoogleSqlDatabaseInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleSqlDatabaseInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstance",
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
func GoogleSqlDatabaseInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleSqlDatabaseInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleSqlDatabaseInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleSqlDatabaseInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleSqlDatabaseInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleSqlDatabaseInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleSqlDatabaseInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) PutClone(value *GoogleSqlDatabaseInstanceClone) {
	if err := g.validatePutCloneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClone",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) PutReplicaConfiguration(value *GoogleSqlDatabaseInstanceReplicaConfiguration) {
	if err := g.validatePutReplicaConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putReplicaConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) PutRestoreBackupContext(value *GoogleSqlDatabaseInstanceRestoreBackupContext) {
	if err := g.validatePutRestoreBackupContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRestoreBackupContext",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) PutSettings(value *GoogleSqlDatabaseInstanceSettings) {
	if err := g.validatePutSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) PutTimeouts(value *GoogleSqlDatabaseInstanceTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) ResetClone() {
	_jsii_.InvokeVoid(
		g,
		"resetClone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		g,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) ResetEncryptionKeyName() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionKeyName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) ResetInstanceType() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) ResetMaintenanceVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenanceVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) ResetMasterInstanceName() {
	_jsii_.InvokeVoid(
		g,
		"resetMasterInstanceName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) ResetReplicaConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetReplicaConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) ResetRestoreBackupContext() {
	_jsii_.InvokeVoid(
		g,
		"resetRestoreBackupContext",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) ResetRootPassword() {
	_jsii_.InvokeVoid(
		g,
		"resetRootPassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) ResetSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


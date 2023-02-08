// googlesqldb
package googlesqldb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlesqldb/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlesqldb/internal"
)

type Googlesqldb interface {
	cdktf.TerraformModule
	ActivationPolicy() *string
	SetActivationPolicy(val *string)
	AdditionalDatabases() interface{}
	SetAdditionalDatabases(val interface{})
	AdditionalUsers() interface{}
	SetAdditionalUsers(val interface{})
	AdditionalUsersOutput() *string
	AvailabilityType() *string
	SetAvailabilityType(val *string)
	BackupConfiguration() interface{}
	SetBackupConfiguration(val interface{})
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CreateTimeout() *string
	SetCreateTimeout(val *string)
	DatabaseFlags() interface{}
	SetDatabaseFlags(val interface{})
	DatabaseVersion() *string
	SetDatabaseVersion(val *string)
	DbCharset() *string
	SetDbCharset(val *string)
	DbCollation() *string
	SetDbCollation(val *string)
	DbName() *string
	SetDbName(val *string)
	DeleteTimeout() *string
	SetDeleteTimeout(val *string)
	DeletionProtection() *bool
	SetDeletionProtection(val *bool)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiskAutoresize() *bool
	SetDiskAutoresize(val *bool)
	DiskSize() *float64
	SetDiskSize(val *float64)
	DiskType() *string
	SetDiskType(val *string)
	EnableDefaultDb() *bool
	SetEnableDefaultDb(val *bool)
	EnableDefaultUser() *bool
	SetEnableDefaultUser(val *bool)
	EncryptionKeyName() *string
	SetEncryptionKeyName(val *string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GeneratedUserPasswordOutput() *string
	IamUserEmails() *[]*string
	SetIamUserEmails(val *[]*string)
	InsightsConfig() interface{}
	SetInsightsConfig(val interface{})
	InstanceConnectionNameOutput() *string
	InstanceFirstIpAddressOutput() *string
	InstanceIpAddressOutput() *string
	InstanceNameOutput() *string
	InstanceSelfLinkOutput() *string
	InstanceServerCaCertOutput() *string
	InstanceServiceAccountEmailAddressOutput() *string
	InstancesOutput() *string
	IpConfiguration() interface{}
	SetIpConfiguration(val interface{})
	MaintenanceWindowDay() *float64
	SetMaintenanceWindowDay(val *float64)
	MaintenanceWindowHour() *float64
	SetMaintenanceWindowHour(val *float64)
	MaintenanceWindowUpdateTrack() *string
	SetMaintenanceWindowUpdateTrack(val *string)
	ModuleDependsOn() *[]interface{}
	SetModuleDependsOn(val *[]interface{})
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	PricingPlan() *string
	SetPricingPlan(val *string)
	PrimaryOutput() *string
	PrivateIpAddressOutput() *string
	ProjectId() *string
	SetProjectId(val *string)
	// Experimental.
	Providers() *[]interface{}
	PublicIpAddressOutput() *string
	RandomInstanceName() *bool
	SetRandomInstanceName(val *bool)
	// Experimental.
	RawOverrides() interface{}
	ReadReplicaDeletionProtection() *bool
	SetReadReplicaDeletionProtection(val *bool)
	ReadReplicaInstanceNamesOutput() *string
	ReadReplicaNameSuffix() *string
	SetReadReplicaNameSuffix(val *string)
	ReadReplicas() interface{}
	SetReadReplicas(val interface{})
	Region() *string
	SetRegion(val *string)
	ReplicasInstanceConnectionNamesOutput() *string
	ReplicasInstanceFirstIpAddressesOutput() *string
	ReplicasInstanceSelfLinksOutput() *string
	ReplicasInstanceServerCaCertsOutput() *string
	ReplicasInstanceServiceAccountEmailAddressesOutput() *string
	ReplicasOutput() *string
	// Experimental.
	SkipAssetCreationFromLocalModules() *bool
	// Experimental.
	Source() *string
	Tier() *string
	SetTier(val *string)
	UpdateTimeout() *string
	SetUpdateTimeout(val *string)
	UserLabels() *map[string]*string
	SetUserLabels(val *map[string]*string)
	UserName() *string
	SetUserName(val *string)
	UserPassword() *string
	SetUserPassword(val *string)
	// Experimental.
	Version() *string
	Zone() *string
	SetZone(val *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	AddProvider(provider interface{})
	// Experimental.
	GetString(output *string) *string
	// Experimental.
	InterpolationForOutput(moduleOutput *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Googlesqldb
type jsiiProxy_Googlesqldb struct {
	internal.Type__cdktfTerraformModule
}

func (j *jsiiProxy_Googlesqldb) ActivationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) AdditionalDatabases() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) AdditionalUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) AdditionalUsersOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalUsersOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) AvailabilityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) BackupConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) CreateTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) DatabaseFlags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) DatabaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) DbCharset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbCharset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) DbCollation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbCollation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) DbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) DeleteTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) DeletionProtection() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) DiskAutoresize() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"diskAutoresize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) DiskSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) DiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) EnableDefaultDb() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableDefaultDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) EnableDefaultUser() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableDefaultUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) EncryptionKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) GeneratedUserPasswordOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generatedUserPasswordOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) IamUserEmails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamUserEmails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) InsightsConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insightsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) InstanceConnectionNameOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceConnectionNameOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) InstanceFirstIpAddressOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceFirstIpAddressOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) InstanceIpAddressOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIpAddressOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) InstanceNameOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceNameOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) InstanceSelfLinkOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceSelfLinkOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) InstanceServerCaCertOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceServerCaCertOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) InstanceServiceAccountEmailAddressOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceServiceAccountEmailAddressOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) InstancesOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancesOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) IpConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) MaintenanceWindowDay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maintenanceWindowDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) MaintenanceWindowHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maintenanceWindowHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) MaintenanceWindowUpdateTrack() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowUpdateTrack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) ModuleDependsOn() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"moduleDependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) PricingPlan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pricingPlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) PrimaryOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) PrivateIpAddressOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddressOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) Providers() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"providers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) PublicIpAddressOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddressOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) RandomInstanceName() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"randomInstanceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) ReadReplicaDeletionProtection() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readReplicaDeletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) ReadReplicaInstanceNamesOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readReplicaInstanceNamesOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) ReadReplicaNameSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readReplicaNameSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) ReadReplicas() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) ReplicasInstanceConnectionNamesOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicasInstanceConnectionNamesOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) ReplicasInstanceFirstIpAddressesOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicasInstanceFirstIpAddressesOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) ReplicasInstanceSelfLinksOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicasInstanceSelfLinksOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) ReplicasInstanceServerCaCertsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicasInstanceServerCaCertsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) ReplicasInstanceServiceAccountEmailAddressesOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicasInstanceServiceAccountEmailAddressesOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) ReplicasOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicasOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) SkipAssetCreationFromLocalModules() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipAssetCreationFromLocalModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) UpdateTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) UserLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) UserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Googlesqldb) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}


func NewGooglesqldb(scope constructs.Construct, id *string, options *GooglesqldbOptions) Googlesqldb {
	_init_.Initialize()

	if err := validateNewGooglesqldbParameters(scope, id, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Googlesqldb{}

	_jsii_.Create(
		"googlesqldb.Googlesqldb",
		[]interface{}{scope, id, options},
		&j,
	)

	return &j
}

func NewGooglesqldb_Override(g Googlesqldb, scope constructs.Construct, id *string, options *GooglesqldbOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"googlesqldb.Googlesqldb",
		[]interface{}{scope, id, options},
		g,
	)
}

func (j *jsiiProxy_Googlesqldb)SetActivationPolicy(val *string) {
	_jsii_.Set(
		j,
		"activationPolicy",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetAdditionalDatabases(val interface{}) {
	if err := j.validateSetAdditionalDatabasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalDatabases",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetAdditionalUsers(val interface{}) {
	if err := j.validateSetAdditionalUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalUsers",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetAvailabilityType(val *string) {
	_jsii_.Set(
		j,
		"availabilityType",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetBackupConfiguration(val interface{}) {
	if err := j.validateSetBackupConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupConfiguration",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetCreateTimeout(val *string) {
	_jsii_.Set(
		j,
		"createTimeout",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetDatabaseFlags(val interface{}) {
	if err := j.validateSetDatabaseFlagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseFlags",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetDatabaseVersion(val *string) {
	if err := j.validateSetDatabaseVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseVersion",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetDbCharset(val *string) {
	_jsii_.Set(
		j,
		"dbCharset",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetDbCollation(val *string) {
	_jsii_.Set(
		j,
		"dbCollation",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetDbName(val *string) {
	_jsii_.Set(
		j,
		"dbName",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetDeleteTimeout(val *string) {
	_jsii_.Set(
		j,
		"deleteTimeout",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetDeletionProtection(val *bool) {
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetDiskAutoresize(val *bool) {
	_jsii_.Set(
		j,
		"diskAutoresize",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetDiskSize(val *float64) {
	_jsii_.Set(
		j,
		"diskSize",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetDiskType(val *string) {
	_jsii_.Set(
		j,
		"diskType",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetEnableDefaultDb(val *bool) {
	_jsii_.Set(
		j,
		"enableDefaultDb",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetEnableDefaultUser(val *bool) {
	_jsii_.Set(
		j,
		"enableDefaultUser",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetEncryptionKeyName(val *string) {
	_jsii_.Set(
		j,
		"encryptionKeyName",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetIamUserEmails(val *[]*string) {
	_jsii_.Set(
		j,
		"iamUserEmails",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetInsightsConfig(val interface{}) {
	if err := j.validateSetInsightsConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insightsConfig",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetIpConfiguration(val interface{}) {
	if err := j.validateSetIpConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipConfiguration",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetMaintenanceWindowDay(val *float64) {
	_jsii_.Set(
		j,
		"maintenanceWindowDay",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetMaintenanceWindowHour(val *float64) {
	_jsii_.Set(
		j,
		"maintenanceWindowHour",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetMaintenanceWindowUpdateTrack(val *string) {
	_jsii_.Set(
		j,
		"maintenanceWindowUpdateTrack",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetModuleDependsOn(val *[]interface{}) {
	_jsii_.Set(
		j,
		"moduleDependsOn",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetPricingPlan(val *string) {
	_jsii_.Set(
		j,
		"pricingPlan",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetRandomInstanceName(val *bool) {
	_jsii_.Set(
		j,
		"randomInstanceName",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetReadReplicaDeletionProtection(val *bool) {
	_jsii_.Set(
		j,
		"readReplicaDeletionProtection",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetReadReplicaNameSuffix(val *string) {
	_jsii_.Set(
		j,
		"readReplicaNameSuffix",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetReadReplicas(val interface{}) {
	if err := j.validateSetReadReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readReplicas",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetTier(val *string) {
	_jsii_.Set(
		j,
		"tier",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetUpdateTimeout(val *string) {
	_jsii_.Set(
		j,
		"updateTimeout",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetUserLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"userLabels",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetUserName(val *string) {
	_jsii_.Set(
		j,
		"userName",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetUserPassword(val *string) {
	_jsii_.Set(
		j,
		"userPassword",
		val,
	)
}

func (j *jsiiProxy_Googlesqldb)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
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
func Googlesqldb_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGooglesqldb_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"googlesqldb.Googlesqldb",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Googlesqldb) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_Googlesqldb) AddProvider(provider interface{}) {
	if err := g.validateAddProviderParameters(provider); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addProvider",
		[]interface{}{provider},
	)
}

func (g *jsiiProxy_Googlesqldb) GetString(output *string) *string {
	if err := g.validateGetStringParameters(output); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Googlesqldb) InterpolationForOutput(moduleOutput *string) cdktf.IResolvable {
	if err := g.validateInterpolationForOutputParameters(moduleOutput); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForOutput",
		[]interface{}{moduleOutput},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Googlesqldb) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_Googlesqldb) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Googlesqldb) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Googlesqldb) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Googlesqldb) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Googlesqldb) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


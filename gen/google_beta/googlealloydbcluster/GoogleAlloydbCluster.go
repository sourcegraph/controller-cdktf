package googlealloydbcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlealloydbcluster/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_alloydb_cluster google_alloydb_cluster}.
type GoogleAlloydbCluster interface {
	cdktf.TerraformResource
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	AutomatedBackupPolicy() GoogleAlloydbClusterAutomatedBackupPolicyOutputReference
	AutomatedBackupPolicyInput() *GoogleAlloydbClusterAutomatedBackupPolicy
	BackupSource() GoogleAlloydbClusterBackupSourceList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ClusterType() *string
	SetClusterType(val *string)
	ClusterTypeInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContinuousBackupConfig() GoogleAlloydbClusterContinuousBackupConfigOutputReference
	ContinuousBackupConfigInput() *GoogleAlloydbClusterContinuousBackupConfig
	ContinuousBackupInfo() GoogleAlloydbClusterContinuousBackupInfoList
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DatabaseVersion() *string
	SetDatabaseVersion(val *string)
	DatabaseVersionInput() *string
	DeletionPolicy() *string
	SetDeletionPolicy(val *string)
	DeletionPolicyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EffectiveAnnotations() cdktf.StringMap
	EffectiveLabels() cdktf.StringMap
	EncryptionConfig() GoogleAlloydbClusterEncryptionConfigOutputReference
	EncryptionConfigInput() *GoogleAlloydbClusterEncryptionConfig
	EncryptionInfo() GoogleAlloydbClusterEncryptionInfoList
	Etag() *string
	SetEtag(val *string)
	EtagInput() *string
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
	InitialUser() GoogleAlloydbClusterInitialUserOutputReference
	InitialUserInput() *GoogleAlloydbClusterInitialUser
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
	MaintenanceUpdatePolicy() GoogleAlloydbClusterMaintenanceUpdatePolicyOutputReference
	MaintenanceUpdatePolicyInput() *GoogleAlloydbClusterMaintenanceUpdatePolicy
	MigrationSource() GoogleAlloydbClusterMigrationSourceList
	Name() *string
	Network() *string
	SetNetwork(val *string)
	NetworkConfig() GoogleAlloydbClusterNetworkConfigOutputReference
	NetworkConfigInput() *GoogleAlloydbClusterNetworkConfig
	NetworkInput() *string
	// The tree node.
	Node() constructs.Node
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
	// Experimental.
	RawOverrides() interface{}
	Reconciling() cdktf.IResolvable
	RestoreBackupSource() GoogleAlloydbClusterRestoreBackupSourceOutputReference
	RestoreBackupSourceInput() *GoogleAlloydbClusterRestoreBackupSource
	RestoreContinuousBackupSource() GoogleAlloydbClusterRestoreContinuousBackupSourceOutputReference
	RestoreContinuousBackupSourceInput() *GoogleAlloydbClusterRestoreContinuousBackupSource
	SecondaryConfig() GoogleAlloydbClusterSecondaryConfigOutputReference
	SecondaryConfigInput() *GoogleAlloydbClusterSecondaryConfig
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleAlloydbClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAutomatedBackupPolicy(value *GoogleAlloydbClusterAutomatedBackupPolicy)
	PutContinuousBackupConfig(value *GoogleAlloydbClusterContinuousBackupConfig)
	PutEncryptionConfig(value *GoogleAlloydbClusterEncryptionConfig)
	PutInitialUser(value *GoogleAlloydbClusterInitialUser)
	PutMaintenanceUpdatePolicy(value *GoogleAlloydbClusterMaintenanceUpdatePolicy)
	PutNetworkConfig(value *GoogleAlloydbClusterNetworkConfig)
	PutRestoreBackupSource(value *GoogleAlloydbClusterRestoreBackupSource)
	PutRestoreContinuousBackupSource(value *GoogleAlloydbClusterRestoreContinuousBackupSource)
	PutSecondaryConfig(value *GoogleAlloydbClusterSecondaryConfig)
	PutTimeouts(value *GoogleAlloydbClusterTimeouts)
	ResetAnnotations()
	ResetAutomatedBackupPolicy()
	ResetClusterType()
	ResetContinuousBackupConfig()
	ResetDatabaseVersion()
	ResetDeletionPolicy()
	ResetDisplayName()
	ResetEncryptionConfig()
	ResetEtag()
	ResetId()
	ResetInitialUser()
	ResetLabels()
	ResetMaintenanceUpdatePolicy()
	ResetNetwork()
	ResetNetworkConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRestoreBackupSource()
	ResetRestoreContinuousBackupSource()
	ResetSecondaryConfig()
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

// The jsii proxy struct for GoogleAlloydbCluster
type jsiiProxy_GoogleAlloydbCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleAlloydbCluster) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) AutomatedBackupPolicy() GoogleAlloydbClusterAutomatedBackupPolicyOutputReference {
	var returns GoogleAlloydbClusterAutomatedBackupPolicyOutputReference
	_jsii_.Get(
		j,
		"automatedBackupPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) AutomatedBackupPolicyInput() *GoogleAlloydbClusterAutomatedBackupPolicy {
	var returns *GoogleAlloydbClusterAutomatedBackupPolicy
	_jsii_.Get(
		j,
		"automatedBackupPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) BackupSource() GoogleAlloydbClusterBackupSourceList {
	var returns GoogleAlloydbClusterBackupSourceList
	_jsii_.Get(
		j,
		"backupSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) ClusterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) ClusterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) ContinuousBackupConfig() GoogleAlloydbClusterContinuousBackupConfigOutputReference {
	var returns GoogleAlloydbClusterContinuousBackupConfigOutputReference
	_jsii_.Get(
		j,
		"continuousBackupConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) ContinuousBackupConfigInput() *GoogleAlloydbClusterContinuousBackupConfig {
	var returns *GoogleAlloydbClusterContinuousBackupConfig
	_jsii_.Get(
		j,
		"continuousBackupConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) ContinuousBackupInfo() GoogleAlloydbClusterContinuousBackupInfoList {
	var returns GoogleAlloydbClusterContinuousBackupInfoList
	_jsii_.Get(
		j,
		"continuousBackupInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) DatabaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) DatabaseVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) DeletionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) DeletionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) EffectiveAnnotations() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) EncryptionConfig() GoogleAlloydbClusterEncryptionConfigOutputReference {
	var returns GoogleAlloydbClusterEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) EncryptionConfigInput() *GoogleAlloydbClusterEncryptionConfig {
	var returns *GoogleAlloydbClusterEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) EncryptionInfo() GoogleAlloydbClusterEncryptionInfoList {
	var returns GoogleAlloydbClusterEncryptionInfoList
	_jsii_.Get(
		j,
		"encryptionInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) EtagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) InitialUser() GoogleAlloydbClusterInitialUserOutputReference {
	var returns GoogleAlloydbClusterInitialUserOutputReference
	_jsii_.Get(
		j,
		"initialUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) InitialUserInput() *GoogleAlloydbClusterInitialUser {
	var returns *GoogleAlloydbClusterInitialUser
	_jsii_.Get(
		j,
		"initialUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) MaintenanceUpdatePolicy() GoogleAlloydbClusterMaintenanceUpdatePolicyOutputReference {
	var returns GoogleAlloydbClusterMaintenanceUpdatePolicyOutputReference
	_jsii_.Get(
		j,
		"maintenanceUpdatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) MaintenanceUpdatePolicyInput() *GoogleAlloydbClusterMaintenanceUpdatePolicy {
	var returns *GoogleAlloydbClusterMaintenanceUpdatePolicy
	_jsii_.Get(
		j,
		"maintenanceUpdatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) MigrationSource() GoogleAlloydbClusterMigrationSourceList {
	var returns GoogleAlloydbClusterMigrationSourceList
	_jsii_.Get(
		j,
		"migrationSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) NetworkConfig() GoogleAlloydbClusterNetworkConfigOutputReference {
	var returns GoogleAlloydbClusterNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) NetworkConfigInput() *GoogleAlloydbClusterNetworkConfig {
	var returns *GoogleAlloydbClusterNetworkConfig
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) Reconciling() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) RestoreBackupSource() GoogleAlloydbClusterRestoreBackupSourceOutputReference {
	var returns GoogleAlloydbClusterRestoreBackupSourceOutputReference
	_jsii_.Get(
		j,
		"restoreBackupSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) RestoreBackupSourceInput() *GoogleAlloydbClusterRestoreBackupSource {
	var returns *GoogleAlloydbClusterRestoreBackupSource
	_jsii_.Get(
		j,
		"restoreBackupSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) RestoreContinuousBackupSource() GoogleAlloydbClusterRestoreContinuousBackupSourceOutputReference {
	var returns GoogleAlloydbClusterRestoreContinuousBackupSourceOutputReference
	_jsii_.Get(
		j,
		"restoreContinuousBackupSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) RestoreContinuousBackupSourceInput() *GoogleAlloydbClusterRestoreContinuousBackupSource {
	var returns *GoogleAlloydbClusterRestoreContinuousBackupSource
	_jsii_.Get(
		j,
		"restoreContinuousBackupSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) SecondaryConfig() GoogleAlloydbClusterSecondaryConfigOutputReference {
	var returns GoogleAlloydbClusterSecondaryConfigOutputReference
	_jsii_.Get(
		j,
		"secondaryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) SecondaryConfigInput() *GoogleAlloydbClusterSecondaryConfig {
	var returns *GoogleAlloydbClusterSecondaryConfig
	_jsii_.Get(
		j,
		"secondaryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) Timeouts() GoogleAlloydbClusterTimeoutsOutputReference {
	var returns GoogleAlloydbClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbCluster) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_alloydb_cluster google_alloydb_cluster} Resource.
func NewGoogleAlloydbCluster(scope constructs.Construct, id *string, config *GoogleAlloydbClusterConfig) GoogleAlloydbCluster {
	_init_.Initialize()

	if err := validateNewGoogleAlloydbClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAlloydbCluster{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleAlloydbCluster.GoogleAlloydbCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_alloydb_cluster google_alloydb_cluster} Resource.
func NewGoogleAlloydbCluster_Override(g GoogleAlloydbCluster, scope constructs.Construct, id *string, config *GoogleAlloydbClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleAlloydbCluster.GoogleAlloydbCluster",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleAlloydbCluster)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbCluster)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbCluster)SetClusterType(val *string) {
	if err := j.validateSetClusterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterType",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbCluster)SetDatabaseVersion(val *string) {
	if err := j.validateSetDatabaseVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbCluster)SetDeletionPolicy(val *string) {
	if err := j.validateSetDeletionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbCluster)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbCluster)SetEtag(val *string) {
	if err := j.validateSetEtagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etag",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbCluster)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbCluster)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbCluster)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
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
func GoogleAlloydbCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAlloydbCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleAlloydbCluster.GoogleAlloydbCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleAlloydbCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAlloydbCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleAlloydbCluster.GoogleAlloydbCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleAlloydbCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAlloydbCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleAlloydbCluster.GoogleAlloydbCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleAlloydbCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleAlloydbCluster.GoogleAlloydbCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleAlloydbCluster) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAlloydbCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAlloydbCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAlloydbCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAlloydbCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAlloydbCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAlloydbCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAlloydbCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAlloydbCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAlloydbCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAlloydbCluster) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) PutAutomatedBackupPolicy(value *GoogleAlloydbClusterAutomatedBackupPolicy) {
	if err := g.validatePutAutomatedBackupPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutomatedBackupPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) PutContinuousBackupConfig(value *GoogleAlloydbClusterContinuousBackupConfig) {
	if err := g.validatePutContinuousBackupConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putContinuousBackupConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) PutEncryptionConfig(value *GoogleAlloydbClusterEncryptionConfig) {
	if err := g.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) PutInitialUser(value *GoogleAlloydbClusterInitialUser) {
	if err := g.validatePutInitialUserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInitialUser",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) PutMaintenanceUpdatePolicy(value *GoogleAlloydbClusterMaintenanceUpdatePolicy) {
	if err := g.validatePutMaintenanceUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaintenanceUpdatePolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) PutNetworkConfig(value *GoogleAlloydbClusterNetworkConfig) {
	if err := g.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) PutRestoreBackupSource(value *GoogleAlloydbClusterRestoreBackupSource) {
	if err := g.validatePutRestoreBackupSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRestoreBackupSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) PutRestoreContinuousBackupSource(value *GoogleAlloydbClusterRestoreContinuousBackupSource) {
	if err := g.validatePutRestoreContinuousBackupSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRestoreContinuousBackupSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) PutSecondaryConfig(value *GoogleAlloydbClusterSecondaryConfig) {
	if err := g.validatePutSecondaryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecondaryConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) PutTimeouts(value *GoogleAlloydbClusterTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) ResetAnnotations() {
	_jsii_.InvokeVoid(
		g,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) ResetAutomatedBackupPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetAutomatedBackupPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) ResetClusterType() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) ResetContinuousBackupConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetContinuousBackupConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) ResetDatabaseVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetDatabaseVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) ResetDeletionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetDeletionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) ResetEtag() {
	_jsii_.InvokeVoid(
		g,
		"resetEtag",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) ResetInitialUser() {
	_jsii_.InvokeVoid(
		g,
		"resetInitialUser",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) ResetMaintenanceUpdatePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenanceUpdatePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) ResetNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) ResetRestoreBackupSource() {
	_jsii_.InvokeVoid(
		g,
		"resetRestoreBackupSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) ResetRestoreContinuousBackupSource() {
	_jsii_.InvokeVoid(
		g,
		"resetRestoreContinuousBackupSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) ResetSecondaryConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSecondaryConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


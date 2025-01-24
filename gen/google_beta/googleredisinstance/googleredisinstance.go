package googleredisinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleredisinstance/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_redis_instance google_redis_instance}.
type GoogleRedisInstance interface {
	cdktf.TerraformResource
	AlternativeLocationId() *string
	SetAlternativeLocationId(val *string)
	AlternativeLocationIdInput() *string
	AuthEnabled() interface{}
	SetAuthEnabled(val interface{})
	AuthEnabledInput() interface{}
	AuthorizedNetwork() *string
	SetAuthorizedNetwork(val *string)
	AuthorizedNetworkInput() *string
	AuthString() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectMode() *string
	SetConnectMode(val *string)
	ConnectModeInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	CurrentLocationId() *string
	CustomerManagedKey() *string
	SetCustomerManagedKey(val *string)
	CustomerManagedKeyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EffectiveLabels() cdktf.StringMap
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Host() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocationId() *string
	SetLocationId(val *string)
	LocationIdInput() *string
	MaintenancePolicy() GoogleRedisInstanceMaintenancePolicyOutputReference
	MaintenancePolicyInput() *GoogleRedisInstanceMaintenancePolicy
	MaintenanceSchedule() GoogleRedisInstanceMaintenanceScheduleList
	MaintenanceVersion() *string
	SetMaintenanceVersion(val *string)
	MaintenanceVersionInput() *string
	MemorySizeGb() *float64
	SetMemorySizeGb(val *float64)
	MemorySizeGbInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Nodes() GoogleRedisInstanceNodesList
	PersistenceConfig() GoogleRedisInstancePersistenceConfigOutputReference
	PersistenceConfigInput() *GoogleRedisInstancePersistenceConfig
	PersistenceIamIdentity() *string
	Port() *float64
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
	ReadEndpoint() *string
	ReadEndpointPort() *float64
	ReadReplicasMode() *string
	SetReadReplicasMode(val *string)
	ReadReplicasModeInput() *string
	RedisConfigs() *map[string]*string
	SetRedisConfigs(val *map[string]*string)
	RedisConfigsInput() *map[string]*string
	RedisVersion() *string
	SetRedisVersion(val *string)
	RedisVersionInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ReplicaCount() *float64
	SetReplicaCount(val *float64)
	ReplicaCountInput() *float64
	ReservedIpRange() *string
	SetReservedIpRange(val *string)
	ReservedIpRangeInput() *string
	SecondaryIpRange() *string
	SetSecondaryIpRange(val *string)
	SecondaryIpRangeInput() *string
	ServerCaCerts() GoogleRedisInstanceServerCaCertsList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Tier() *string
	SetTier(val *string)
	TierInput() *string
	Timeouts() GoogleRedisInstanceTimeoutsOutputReference
	TimeoutsInput() interface{}
	TransitEncryptionMode() *string
	SetTransitEncryptionMode(val *string)
	TransitEncryptionModeInput() *string
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
	PutMaintenancePolicy(value *GoogleRedisInstanceMaintenancePolicy)
	PutPersistenceConfig(value *GoogleRedisInstancePersistenceConfig)
	PutTimeouts(value *GoogleRedisInstanceTimeouts)
	ResetAlternativeLocationId()
	ResetAuthEnabled()
	ResetAuthorizedNetwork()
	ResetConnectMode()
	ResetCustomerManagedKey()
	ResetDisplayName()
	ResetId()
	ResetLabels()
	ResetLocationId()
	ResetMaintenancePolicy()
	ResetMaintenanceVersion()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPersistenceConfig()
	ResetProject()
	ResetReadReplicasMode()
	ResetRedisConfigs()
	ResetRedisVersion()
	ResetRegion()
	ResetReplicaCount()
	ResetReservedIpRange()
	ResetSecondaryIpRange()
	ResetTier()
	ResetTimeouts()
	ResetTransitEncryptionMode()
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
}

// The jsii proxy struct for GoogleRedisInstance
type jsiiProxy_GoogleRedisInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleRedisInstance) AlternativeLocationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alternativeLocationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) AlternativeLocationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alternativeLocationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) AuthEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) AuthEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) AuthorizedNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizedNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) AuthorizedNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizedNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) AuthString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) ConnectMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) ConnectModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) CurrentLocationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentLocationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) CustomerManagedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerManagedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) CustomerManagedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerManagedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) LocationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) LocationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) MaintenancePolicy() GoogleRedisInstanceMaintenancePolicyOutputReference {
	var returns GoogleRedisInstanceMaintenancePolicyOutputReference
	_jsii_.Get(
		j,
		"maintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) MaintenancePolicyInput() *GoogleRedisInstanceMaintenancePolicy {
	var returns *GoogleRedisInstanceMaintenancePolicy
	_jsii_.Get(
		j,
		"maintenancePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) MaintenanceSchedule() GoogleRedisInstanceMaintenanceScheduleList {
	var returns GoogleRedisInstanceMaintenanceScheduleList
	_jsii_.Get(
		j,
		"maintenanceSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) MaintenanceVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) MaintenanceVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) MemorySizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) MemorySizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) Nodes() GoogleRedisInstanceNodesList {
	var returns GoogleRedisInstanceNodesList
	_jsii_.Get(
		j,
		"nodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) PersistenceConfig() GoogleRedisInstancePersistenceConfigOutputReference {
	var returns GoogleRedisInstancePersistenceConfigOutputReference
	_jsii_.Get(
		j,
		"persistenceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) PersistenceConfigInput() *GoogleRedisInstancePersistenceConfig {
	var returns *GoogleRedisInstancePersistenceConfig
	_jsii_.Get(
		j,
		"persistenceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) PersistenceIamIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"persistenceIamIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) ReadEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) ReadEndpointPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) ReadReplicasMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readReplicasMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) ReadReplicasModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readReplicasModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) RedisConfigs() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"redisConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) RedisConfigsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"redisConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) RedisVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) RedisVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) ReplicaCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) ReplicaCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) ReservedIpRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservedIpRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) ReservedIpRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservedIpRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) SecondaryIpRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryIpRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) SecondaryIpRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryIpRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) ServerCaCerts() GoogleRedisInstanceServerCaCertsList {
	var returns GoogleRedisInstanceServerCaCertsList
	_jsii_.Get(
		j,
		"serverCaCerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) TierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) Timeouts() GoogleRedisInstanceTimeoutsOutputReference {
	var returns GoogleRedisInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) TransitEncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitEncryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisInstance) TransitEncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitEncryptionModeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_redis_instance google_redis_instance} Resource.
func NewGoogleRedisInstance(scope constructs.Construct, id *string, config *GoogleRedisInstanceConfig) GoogleRedisInstance {
	_init_.Initialize()

	if err := validateNewGoogleRedisInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleRedisInstance{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleRedisInstance.GoogleRedisInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_redis_instance google_redis_instance} Resource.
func NewGoogleRedisInstance_Override(g GoogleRedisInstance, scope constructs.Construct, id *string, config *GoogleRedisInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleRedisInstance.GoogleRedisInstance",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetAlternativeLocationId(val *string) {
	if err := j.validateSetAlternativeLocationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alternativeLocationId",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetAuthEnabled(val interface{}) {
	if err := j.validateSetAuthEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetAuthorizedNetwork(val *string) {
	if err := j.validateSetAuthorizedNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizedNetwork",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetConnectMode(val *string) {
	if err := j.validateSetConnectModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectMode",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetCustomerManagedKey(val *string) {
	if err := j.validateSetCustomerManagedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerManagedKey",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetLocationId(val *string) {
	if err := j.validateSetLocationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locationId",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetMaintenanceVersion(val *string) {
	if err := j.validateSetMaintenanceVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetMemorySizeGb(val *float64) {
	if err := j.validateSetMemorySizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memorySizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetReadReplicasMode(val *string) {
	if err := j.validateSetReadReplicasModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readReplicasMode",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetRedisConfigs(val *map[string]*string) {
	if err := j.validateSetRedisConfigsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redisConfigs",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetRedisVersion(val *string) {
	if err := j.validateSetRedisVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redisVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetReplicaCount(val *float64) {
	if err := j.validateSetReplicaCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaCount",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetReservedIpRange(val *string) {
	if err := j.validateSetReservedIpRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedIpRange",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetSecondaryIpRange(val *string) {
	if err := j.validateSetSecondaryIpRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryIpRange",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetTier(val *string) {
	if err := j.validateSetTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tier",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisInstance)SetTransitEncryptionMode(val *string) {
	if err := j.validateSetTransitEncryptionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitEncryptionMode",
		val,
	)
}

// Generates CDKTF code for importing a GoogleRedisInstance resource upon running "cdktf plan <stack-name>".
func GoogleRedisInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleRedisInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleRedisInstance.GoogleRedisInstance",
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
func GoogleRedisInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleRedisInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleRedisInstance.GoogleRedisInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleRedisInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleRedisInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleRedisInstance.GoogleRedisInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleRedisInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleRedisInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleRedisInstance.GoogleRedisInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleRedisInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleRedisInstance.GoogleRedisInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleRedisInstance) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleRedisInstance) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleRedisInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleRedisInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleRedisInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleRedisInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleRedisInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleRedisInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleRedisInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleRedisInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleRedisInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleRedisInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleRedisInstance) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleRedisInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleRedisInstance) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleRedisInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleRedisInstance) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleRedisInstance) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleRedisInstance) PutMaintenancePolicy(value *GoogleRedisInstanceMaintenancePolicy) {
	if err := g.validatePutMaintenancePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaintenancePolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleRedisInstance) PutPersistenceConfig(value *GoogleRedisInstancePersistenceConfig) {
	if err := g.validatePutPersistenceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPersistenceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleRedisInstance) PutTimeouts(value *GoogleRedisInstanceTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetAlternativeLocationId() {
	_jsii_.InvokeVoid(
		g,
		"resetAlternativeLocationId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetAuthEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetAuthorizedNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthorizedNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetConnectMode() {
	_jsii_.InvokeVoid(
		g,
		"resetConnectMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetCustomerManagedKey() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomerManagedKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetLocationId() {
	_jsii_.InvokeVoid(
		g,
		"resetLocationId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetMaintenancePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenancePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetMaintenanceVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenanceVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetPersistenceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPersistenceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetReadReplicasMode() {
	_jsii_.InvokeVoid(
		g,
		"resetReadReplicasMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetRedisConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetRedisConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetRedisVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetRedisVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetReplicaCount() {
	_jsii_.InvokeVoid(
		g,
		"resetReplicaCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetReservedIpRange() {
	_jsii_.InvokeVoid(
		g,
		"resetReservedIpRange",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetSecondaryIpRange() {
	_jsii_.InvokeVoid(
		g,
		"resetSecondaryIpRange",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetTier() {
	_jsii_.InvokeVoid(
		g,
		"resetTier",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) ResetTransitEncryptionMode() {
	_jsii_.InvokeVoid(
		g,
		"resetTransitEncryptionMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleRedisInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleRedisInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleRedisInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleRedisInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleRedisInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


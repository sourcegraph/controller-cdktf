package googlegkeonprembaremetalcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlegkeonprembaremetalcluster/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_gkeonprem_bare_metal_cluster google_gkeonprem_bare_metal_cluster}.
type GoogleGkeonpremBareMetalCluster interface {
	cdktf.TerraformResource
	AdminClusterMembership() *string
	SetAdminClusterMembership(val *string)
	AdminClusterMembershipInput() *string
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	BareMetalVersion() *string
	SetBareMetalVersion(val *string)
	BareMetalVersionInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterOperations() GoogleGkeonpremBareMetalClusterClusterOperationsOutputReference
	ClusterOperationsInput() *GoogleGkeonpremBareMetalClusterClusterOperations
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ControlPlane() GoogleGkeonpremBareMetalClusterControlPlaneOutputReference
	ControlPlaneInput() *GoogleGkeonpremBareMetalClusterControlPlane
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	DeleteTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Endpoint() *string
	Etag() *string
	Fleet() GoogleGkeonpremBareMetalClusterFleetList
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancer() GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference
	LoadBalancerInput() *GoogleGkeonpremBareMetalClusterLoadBalancer
	LocalName() *string
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaintenanceConfig() GoogleGkeonpremBareMetalClusterMaintenanceConfigOutputReference
	MaintenanceConfigInput() *GoogleGkeonpremBareMetalClusterMaintenanceConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkConfig() GoogleGkeonpremBareMetalClusterNetworkConfigOutputReference
	NetworkConfigInput() *GoogleGkeonpremBareMetalClusterNetworkConfig
	// The tree node.
	Node() constructs.Node
	NodeAccessConfig() GoogleGkeonpremBareMetalClusterNodeAccessConfigOutputReference
	NodeAccessConfigInput() *GoogleGkeonpremBareMetalClusterNodeAccessConfig
	NodeConfig() GoogleGkeonpremBareMetalClusterNodeConfigOutputReference
	NodeConfigInput() *GoogleGkeonpremBareMetalClusterNodeConfig
	OsEnvironmentConfig() GoogleGkeonpremBareMetalClusterOsEnvironmentConfigOutputReference
	OsEnvironmentConfigInput() *GoogleGkeonpremBareMetalClusterOsEnvironmentConfig
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
	Proxy() GoogleGkeonpremBareMetalClusterProxyOutputReference
	ProxyInput() *GoogleGkeonpremBareMetalClusterProxy
	// Experimental.
	RawOverrides() interface{}
	Reconciling() cdktf.IResolvable
	SecurityConfig() GoogleGkeonpremBareMetalClusterSecurityConfigOutputReference
	SecurityConfigInput() *GoogleGkeonpremBareMetalClusterSecurityConfig
	State() *string
	Status() GoogleGkeonpremBareMetalClusterStatusList
	Storage() GoogleGkeonpremBareMetalClusterStorageOutputReference
	StorageInput() *GoogleGkeonpremBareMetalClusterStorage
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleGkeonpremBareMetalClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
	UpdateTime() *string
	ValidationCheck() GoogleGkeonpremBareMetalClusterValidationCheckList
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
	PutClusterOperations(value *GoogleGkeonpremBareMetalClusterClusterOperations)
	PutControlPlane(value *GoogleGkeonpremBareMetalClusterControlPlane)
	PutLoadBalancer(value *GoogleGkeonpremBareMetalClusterLoadBalancer)
	PutMaintenanceConfig(value *GoogleGkeonpremBareMetalClusterMaintenanceConfig)
	PutNetworkConfig(value *GoogleGkeonpremBareMetalClusterNetworkConfig)
	PutNodeAccessConfig(value *GoogleGkeonpremBareMetalClusterNodeAccessConfig)
	PutNodeConfig(value *GoogleGkeonpremBareMetalClusterNodeConfig)
	PutOsEnvironmentConfig(value *GoogleGkeonpremBareMetalClusterOsEnvironmentConfig)
	PutProxy(value *GoogleGkeonpremBareMetalClusterProxy)
	PutSecurityConfig(value *GoogleGkeonpremBareMetalClusterSecurityConfig)
	PutStorage(value *GoogleGkeonpremBareMetalClusterStorage)
	PutTimeouts(value *GoogleGkeonpremBareMetalClusterTimeouts)
	ResetAnnotations()
	ResetClusterOperations()
	ResetDescription()
	ResetId()
	ResetMaintenanceConfig()
	ResetNodeAccessConfig()
	ResetNodeConfig()
	ResetOsEnvironmentConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetProxy()
	ResetSecurityConfig()
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

// The jsii proxy struct for GoogleGkeonpremBareMetalCluster
type jsiiProxy_GoogleGkeonpremBareMetalCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) AdminClusterMembership() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminClusterMembership",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) AdminClusterMembershipInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminClusterMembershipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) BareMetalVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bareMetalVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) BareMetalVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bareMetalVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) ClusterOperations() GoogleGkeonpremBareMetalClusterClusterOperationsOutputReference {
	var returns GoogleGkeonpremBareMetalClusterClusterOperationsOutputReference
	_jsii_.Get(
		j,
		"clusterOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) ClusterOperationsInput() *GoogleGkeonpremBareMetalClusterClusterOperations {
	var returns *GoogleGkeonpremBareMetalClusterClusterOperations
	_jsii_.Get(
		j,
		"clusterOperationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) ControlPlane() GoogleGkeonpremBareMetalClusterControlPlaneOutputReference {
	var returns GoogleGkeonpremBareMetalClusterControlPlaneOutputReference
	_jsii_.Get(
		j,
		"controlPlane",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) ControlPlaneInput() *GoogleGkeonpremBareMetalClusterControlPlane {
	var returns *GoogleGkeonpremBareMetalClusterControlPlane
	_jsii_.Get(
		j,
		"controlPlaneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) DeleteTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Fleet() GoogleGkeonpremBareMetalClusterFleetList {
	var returns GoogleGkeonpremBareMetalClusterFleetList
	_jsii_.Get(
		j,
		"fleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) LoadBalancer() GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference {
	var returns GoogleGkeonpremBareMetalClusterLoadBalancerOutputReference
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) LoadBalancerInput() *GoogleGkeonpremBareMetalClusterLoadBalancer {
	var returns *GoogleGkeonpremBareMetalClusterLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) LocalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) MaintenanceConfig() GoogleGkeonpremBareMetalClusterMaintenanceConfigOutputReference {
	var returns GoogleGkeonpremBareMetalClusterMaintenanceConfigOutputReference
	_jsii_.Get(
		j,
		"maintenanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) MaintenanceConfigInput() *GoogleGkeonpremBareMetalClusterMaintenanceConfig {
	var returns *GoogleGkeonpremBareMetalClusterMaintenanceConfig
	_jsii_.Get(
		j,
		"maintenanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) NetworkConfig() GoogleGkeonpremBareMetalClusterNetworkConfigOutputReference {
	var returns GoogleGkeonpremBareMetalClusterNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) NetworkConfigInput() *GoogleGkeonpremBareMetalClusterNetworkConfig {
	var returns *GoogleGkeonpremBareMetalClusterNetworkConfig
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) NodeAccessConfig() GoogleGkeonpremBareMetalClusterNodeAccessConfigOutputReference {
	var returns GoogleGkeonpremBareMetalClusterNodeAccessConfigOutputReference
	_jsii_.Get(
		j,
		"nodeAccessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) NodeAccessConfigInput() *GoogleGkeonpremBareMetalClusterNodeAccessConfig {
	var returns *GoogleGkeonpremBareMetalClusterNodeAccessConfig
	_jsii_.Get(
		j,
		"nodeAccessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) NodeConfig() GoogleGkeonpremBareMetalClusterNodeConfigOutputReference {
	var returns GoogleGkeonpremBareMetalClusterNodeConfigOutputReference
	_jsii_.Get(
		j,
		"nodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) NodeConfigInput() *GoogleGkeonpremBareMetalClusterNodeConfig {
	var returns *GoogleGkeonpremBareMetalClusterNodeConfig
	_jsii_.Get(
		j,
		"nodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) OsEnvironmentConfig() GoogleGkeonpremBareMetalClusterOsEnvironmentConfigOutputReference {
	var returns GoogleGkeonpremBareMetalClusterOsEnvironmentConfigOutputReference
	_jsii_.Get(
		j,
		"osEnvironmentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) OsEnvironmentConfigInput() *GoogleGkeonpremBareMetalClusterOsEnvironmentConfig {
	var returns *GoogleGkeonpremBareMetalClusterOsEnvironmentConfig
	_jsii_.Get(
		j,
		"osEnvironmentConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Proxy() GoogleGkeonpremBareMetalClusterProxyOutputReference {
	var returns GoogleGkeonpremBareMetalClusterProxyOutputReference
	_jsii_.Get(
		j,
		"proxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) ProxyInput() *GoogleGkeonpremBareMetalClusterProxy {
	var returns *GoogleGkeonpremBareMetalClusterProxy
	_jsii_.Get(
		j,
		"proxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Reconciling() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) SecurityConfig() GoogleGkeonpremBareMetalClusterSecurityConfigOutputReference {
	var returns GoogleGkeonpremBareMetalClusterSecurityConfigOutputReference
	_jsii_.Get(
		j,
		"securityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) SecurityConfigInput() *GoogleGkeonpremBareMetalClusterSecurityConfig {
	var returns *GoogleGkeonpremBareMetalClusterSecurityConfig
	_jsii_.Get(
		j,
		"securityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Status() GoogleGkeonpremBareMetalClusterStatusList {
	var returns GoogleGkeonpremBareMetalClusterStatusList
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Storage() GoogleGkeonpremBareMetalClusterStorageOutputReference {
	var returns GoogleGkeonpremBareMetalClusterStorageOutputReference
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) StorageInput() *GoogleGkeonpremBareMetalClusterStorage {
	var returns *GoogleGkeonpremBareMetalClusterStorage
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Timeouts() GoogleGkeonpremBareMetalClusterTimeoutsOutputReference {
	var returns GoogleGkeonpremBareMetalClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster) ValidationCheck() GoogleGkeonpremBareMetalClusterValidationCheckList {
	var returns GoogleGkeonpremBareMetalClusterValidationCheckList
	_jsii_.Get(
		j,
		"validationCheck",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_gkeonprem_bare_metal_cluster google_gkeonprem_bare_metal_cluster} Resource.
func NewGoogleGkeonpremBareMetalCluster(scope constructs.Construct, id *string, config *GoogleGkeonpremBareMetalClusterConfig) GoogleGkeonpremBareMetalCluster {
	_init_.Initialize()

	if err := validateNewGoogleGkeonpremBareMetalClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeonpremBareMetalCluster{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeonpremBareMetalCluster.GoogleGkeonpremBareMetalCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_gkeonprem_bare_metal_cluster google_gkeonprem_bare_metal_cluster} Resource.
func NewGoogleGkeonpremBareMetalCluster_Override(g GoogleGkeonpremBareMetalCluster, scope constructs.Construct, id *string, config *GoogleGkeonpremBareMetalClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeonpremBareMetalCluster.GoogleGkeonpremBareMetalCluster",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster)SetAdminClusterMembership(val *string) {
	if err := j.validateSetAdminClusterMembershipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminClusterMembership",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster)SetBareMetalVersion(val *string) {
	if err := j.validateSetBareMetalVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bareMetalVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalCluster)SetProvisioners(val *[]interface{}) {
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
func GoogleGkeonpremBareMetalCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleGkeonpremBareMetalCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleGkeonpremBareMetalCluster.GoogleGkeonpremBareMetalCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleGkeonpremBareMetalCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleGkeonpremBareMetalCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleGkeonpremBareMetalCluster.GoogleGkeonpremBareMetalCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleGkeonpremBareMetalCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleGkeonpremBareMetalCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleGkeonpremBareMetalCluster.GoogleGkeonpremBareMetalCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleGkeonpremBareMetalCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleGkeonpremBareMetalCluster.GoogleGkeonpremBareMetalCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) PutClusterOperations(value *GoogleGkeonpremBareMetalClusterClusterOperations) {
	if err := g.validatePutClusterOperationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClusterOperations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) PutControlPlane(value *GoogleGkeonpremBareMetalClusterControlPlane) {
	if err := g.validatePutControlPlaneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putControlPlane",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) PutLoadBalancer(value *GoogleGkeonpremBareMetalClusterLoadBalancer) {
	if err := g.validatePutLoadBalancerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLoadBalancer",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) PutMaintenanceConfig(value *GoogleGkeonpremBareMetalClusterMaintenanceConfig) {
	if err := g.validatePutMaintenanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaintenanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) PutNetworkConfig(value *GoogleGkeonpremBareMetalClusterNetworkConfig) {
	if err := g.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) PutNodeAccessConfig(value *GoogleGkeonpremBareMetalClusterNodeAccessConfig) {
	if err := g.validatePutNodeAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodeAccessConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) PutNodeConfig(value *GoogleGkeonpremBareMetalClusterNodeConfig) {
	if err := g.validatePutNodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) PutOsEnvironmentConfig(value *GoogleGkeonpremBareMetalClusterOsEnvironmentConfig) {
	if err := g.validatePutOsEnvironmentConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOsEnvironmentConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) PutProxy(value *GoogleGkeonpremBareMetalClusterProxy) {
	if err := g.validatePutProxyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putProxy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) PutSecurityConfig(value *GoogleGkeonpremBareMetalClusterSecurityConfig) {
	if err := g.validatePutSecurityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecurityConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) PutStorage(value *GoogleGkeonpremBareMetalClusterStorage) {
	if err := g.validatePutStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStorage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) PutTimeouts(value *GoogleGkeonpremBareMetalClusterTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) ResetAnnotations() {
	_jsii_.InvokeVoid(
		g,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) ResetClusterOperations() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterOperations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) ResetMaintenanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) ResetNodeAccessConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeAccessConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) ResetNodeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) ResetOsEnvironmentConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetOsEnvironmentConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) ResetProxy() {
	_jsii_.InvokeVoid(
		g,
		"resetProxy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) ResetSecurityConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


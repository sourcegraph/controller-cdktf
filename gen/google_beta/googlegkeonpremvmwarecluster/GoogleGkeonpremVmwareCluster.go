package googlegkeonpremvmwarecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlegkeonpremvmwarecluster/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_gkeonprem_vmware_cluster google_gkeonprem_vmware_cluster}.
type GoogleGkeonpremVmwareCluster interface {
	cdktf.TerraformResource
	AdminClusterMembership() *string
	SetAdminClusterMembership(val *string)
	AdminClusterMembershipInput() *string
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	AntiAffinityGroups() GoogleGkeonpremVmwareClusterAntiAffinityGroupsOutputReference
	AntiAffinityGroupsInput() *GoogleGkeonpremVmwareClusterAntiAffinityGroups
	Authorization() GoogleGkeonpremVmwareClusterAuthorizationOutputReference
	AuthorizationInput() *GoogleGkeonpremVmwareClusterAuthorization
	AutoRepairConfig() GoogleGkeonpremVmwareClusterAutoRepairConfigOutputReference
	AutoRepairConfigInput() *GoogleGkeonpremVmwareClusterAutoRepairConfig
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ControlPlaneNode() GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference
	ControlPlaneNodeInput() *GoogleGkeonpremVmwareClusterControlPlaneNode
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	DataplaneV2() GoogleGkeonpremVmwareClusterDataplaneV2OutputReference
	DataplaneV2Input() *GoogleGkeonpremVmwareClusterDataplaneV2
	DeleteTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EnableControlPlaneV2() interface{}
	SetEnableControlPlaneV2(val interface{})
	EnableControlPlaneV2Input() interface{}
	Endpoint() *string
	Etag() *string
	Fleet() GoogleGkeonpremVmwareClusterFleetList
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
	LoadBalancer() GoogleGkeonpremVmwareClusterLoadBalancerOutputReference
	LoadBalancerInput() *GoogleGkeonpremVmwareClusterLoadBalancer
	LocalName() *string
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkConfig() GoogleGkeonpremVmwareClusterNetworkConfigOutputReference
	NetworkConfigInput() *GoogleGkeonpremVmwareClusterNetworkConfig
	// The tree node.
	Node() constructs.Node
	OnPremVersion() *string
	SetOnPremVersion(val *string)
	OnPremVersionInput() *string
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
	State() *string
	Status() GoogleGkeonpremVmwareClusterStatusList
	Storage() GoogleGkeonpremVmwareClusterStorageOutputReference
	StorageInput() *GoogleGkeonpremVmwareClusterStorage
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleGkeonpremVmwareClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
	UpdateTime() *string
	ValidationCheck() GoogleGkeonpremVmwareClusterValidationCheckList
	Vcenter() GoogleGkeonpremVmwareClusterVcenterList
	VmTrackingEnabled() interface{}
	SetVmTrackingEnabled(val interface{})
	VmTrackingEnabledInput() interface{}
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
	PutAntiAffinityGroups(value *GoogleGkeonpremVmwareClusterAntiAffinityGroups)
	PutAuthorization(value *GoogleGkeonpremVmwareClusterAuthorization)
	PutAutoRepairConfig(value *GoogleGkeonpremVmwareClusterAutoRepairConfig)
	PutControlPlaneNode(value *GoogleGkeonpremVmwareClusterControlPlaneNode)
	PutDataplaneV2(value *GoogleGkeonpremVmwareClusterDataplaneV2)
	PutLoadBalancer(value *GoogleGkeonpremVmwareClusterLoadBalancer)
	PutNetworkConfig(value *GoogleGkeonpremVmwareClusterNetworkConfig)
	PutStorage(value *GoogleGkeonpremVmwareClusterStorage)
	PutTimeouts(value *GoogleGkeonpremVmwareClusterTimeouts)
	ResetAnnotations()
	ResetAntiAffinityGroups()
	ResetAuthorization()
	ResetAutoRepairConfig()
	ResetDataplaneV2()
	ResetDescription()
	ResetEnableControlPlaneV2()
	ResetId()
	ResetLoadBalancer()
	ResetNetworkConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetStorage()
	ResetTimeouts()
	ResetVmTrackingEnabled()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleGkeonpremVmwareCluster
type jsiiProxy_GoogleGkeonpremVmwareCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) AdminClusterMembership() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminClusterMembership",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) AdminClusterMembershipInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminClusterMembershipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) AntiAffinityGroups() GoogleGkeonpremVmwareClusterAntiAffinityGroupsOutputReference {
	var returns GoogleGkeonpremVmwareClusterAntiAffinityGroupsOutputReference
	_jsii_.Get(
		j,
		"antiAffinityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) AntiAffinityGroupsInput() *GoogleGkeonpremVmwareClusterAntiAffinityGroups {
	var returns *GoogleGkeonpremVmwareClusterAntiAffinityGroups
	_jsii_.Get(
		j,
		"antiAffinityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Authorization() GoogleGkeonpremVmwareClusterAuthorizationOutputReference {
	var returns GoogleGkeonpremVmwareClusterAuthorizationOutputReference
	_jsii_.Get(
		j,
		"authorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) AuthorizationInput() *GoogleGkeonpremVmwareClusterAuthorization {
	var returns *GoogleGkeonpremVmwareClusterAuthorization
	_jsii_.Get(
		j,
		"authorizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) AutoRepairConfig() GoogleGkeonpremVmwareClusterAutoRepairConfigOutputReference {
	var returns GoogleGkeonpremVmwareClusterAutoRepairConfigOutputReference
	_jsii_.Get(
		j,
		"autoRepairConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) AutoRepairConfigInput() *GoogleGkeonpremVmwareClusterAutoRepairConfig {
	var returns *GoogleGkeonpremVmwareClusterAutoRepairConfig
	_jsii_.Get(
		j,
		"autoRepairConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) ControlPlaneNode() GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference {
	var returns GoogleGkeonpremVmwareClusterControlPlaneNodeOutputReference
	_jsii_.Get(
		j,
		"controlPlaneNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) ControlPlaneNodeInput() *GoogleGkeonpremVmwareClusterControlPlaneNode {
	var returns *GoogleGkeonpremVmwareClusterControlPlaneNode
	_jsii_.Get(
		j,
		"controlPlaneNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) DataplaneV2() GoogleGkeonpremVmwareClusterDataplaneV2OutputReference {
	var returns GoogleGkeonpremVmwareClusterDataplaneV2OutputReference
	_jsii_.Get(
		j,
		"dataplaneV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) DataplaneV2Input() *GoogleGkeonpremVmwareClusterDataplaneV2 {
	var returns *GoogleGkeonpremVmwareClusterDataplaneV2
	_jsii_.Get(
		j,
		"dataplaneV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) DeleteTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) EnableControlPlaneV2() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableControlPlaneV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) EnableControlPlaneV2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableControlPlaneV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Fleet() GoogleGkeonpremVmwareClusterFleetList {
	var returns GoogleGkeonpremVmwareClusterFleetList
	_jsii_.Get(
		j,
		"fleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) LoadBalancer() GoogleGkeonpremVmwareClusterLoadBalancerOutputReference {
	var returns GoogleGkeonpremVmwareClusterLoadBalancerOutputReference
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) LoadBalancerInput() *GoogleGkeonpremVmwareClusterLoadBalancer {
	var returns *GoogleGkeonpremVmwareClusterLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) LocalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) NetworkConfig() GoogleGkeonpremVmwareClusterNetworkConfigOutputReference {
	var returns GoogleGkeonpremVmwareClusterNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) NetworkConfigInput() *GoogleGkeonpremVmwareClusterNetworkConfig {
	var returns *GoogleGkeonpremVmwareClusterNetworkConfig
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) OnPremVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onPremVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) OnPremVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onPremVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Reconciling() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Status() GoogleGkeonpremVmwareClusterStatusList {
	var returns GoogleGkeonpremVmwareClusterStatusList
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Storage() GoogleGkeonpremVmwareClusterStorageOutputReference {
	var returns GoogleGkeonpremVmwareClusterStorageOutputReference
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) StorageInput() *GoogleGkeonpremVmwareClusterStorage {
	var returns *GoogleGkeonpremVmwareClusterStorage
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Timeouts() GoogleGkeonpremVmwareClusterTimeoutsOutputReference {
	var returns GoogleGkeonpremVmwareClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) ValidationCheck() GoogleGkeonpremVmwareClusterValidationCheckList {
	var returns GoogleGkeonpremVmwareClusterValidationCheckList
	_jsii_.Get(
		j,
		"validationCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) Vcenter() GoogleGkeonpremVmwareClusterVcenterList {
	var returns GoogleGkeonpremVmwareClusterVcenterList
	_jsii_.Get(
		j,
		"vcenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) VmTrackingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vmTrackingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster) VmTrackingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vmTrackingEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_gkeonprem_vmware_cluster google_gkeonprem_vmware_cluster} Resource.
func NewGoogleGkeonpremVmwareCluster(scope constructs.Construct, id *string, config *GoogleGkeonpremVmwareClusterConfig) GoogleGkeonpremVmwareCluster {
	_init_.Initialize()

	if err := validateNewGoogleGkeonpremVmwareClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeonpremVmwareCluster{}

	_jsii_.Create(
		"google-beta.googleGkeonpremVmwareCluster.GoogleGkeonpremVmwareCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_gkeonprem_vmware_cluster google_gkeonprem_vmware_cluster} Resource.
func NewGoogleGkeonpremVmwareCluster_Override(g GoogleGkeonpremVmwareCluster, scope constructs.Construct, id *string, config *GoogleGkeonpremVmwareClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleGkeonpremVmwareCluster.GoogleGkeonpremVmwareCluster",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster)SetAdminClusterMembership(val *string) {
	if err := j.validateSetAdminClusterMembershipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminClusterMembership",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster)SetEnableControlPlaneV2(val interface{}) {
	if err := j.validateSetEnableControlPlaneV2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableControlPlaneV2",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster)SetOnPremVersion(val *string) {
	if err := j.validateSetOnPremVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onPremVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareCluster)SetVmTrackingEnabled(val interface{}) {
	if err := j.validateSetVmTrackingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmTrackingEnabled",
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
func GoogleGkeonpremVmwareCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleGkeonpremVmwareCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleGkeonpremVmwareCluster.GoogleGkeonpremVmwareCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleGkeonpremVmwareCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleGkeonpremVmwareCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleGkeonpremVmwareCluster.GoogleGkeonpremVmwareCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleGkeonpremVmwareCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleGkeonpremVmwareCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleGkeonpremVmwareCluster.GoogleGkeonpremVmwareCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleGkeonpremVmwareCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.googleGkeonpremVmwareCluster.GoogleGkeonpremVmwareCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) PutAntiAffinityGroups(value *GoogleGkeonpremVmwareClusterAntiAffinityGroups) {
	if err := g.validatePutAntiAffinityGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAntiAffinityGroups",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) PutAuthorization(value *GoogleGkeonpremVmwareClusterAuthorization) {
	if err := g.validatePutAuthorizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthorization",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) PutAutoRepairConfig(value *GoogleGkeonpremVmwareClusterAutoRepairConfig) {
	if err := g.validatePutAutoRepairConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutoRepairConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) PutControlPlaneNode(value *GoogleGkeonpremVmwareClusterControlPlaneNode) {
	if err := g.validatePutControlPlaneNodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putControlPlaneNode",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) PutDataplaneV2(value *GoogleGkeonpremVmwareClusterDataplaneV2) {
	if err := g.validatePutDataplaneV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataplaneV2",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) PutLoadBalancer(value *GoogleGkeonpremVmwareClusterLoadBalancer) {
	if err := g.validatePutLoadBalancerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLoadBalancer",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) PutNetworkConfig(value *GoogleGkeonpremVmwareClusterNetworkConfig) {
	if err := g.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) PutStorage(value *GoogleGkeonpremVmwareClusterStorage) {
	if err := g.validatePutStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStorage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) PutTimeouts(value *GoogleGkeonpremVmwareClusterTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) ResetAnnotations() {
	_jsii_.InvokeVoid(
		g,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) ResetAntiAffinityGroups() {
	_jsii_.InvokeVoid(
		g,
		"resetAntiAffinityGroups",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) ResetAuthorization() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthorization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) ResetAutoRepairConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoRepairConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) ResetDataplaneV2() {
	_jsii_.InvokeVoid(
		g,
		"resetDataplaneV2",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) ResetEnableControlPlaneV2() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableControlPlaneV2",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) ResetLoadBalancer() {
	_jsii_.InvokeVoid(
		g,
		"resetLoadBalancer",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) ResetStorage() {
	_jsii_.InvokeVoid(
		g,
		"resetStorage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) ResetVmTrackingEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetVmTrackingEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


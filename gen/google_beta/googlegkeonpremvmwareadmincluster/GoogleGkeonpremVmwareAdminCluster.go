package googlegkeonpremvmwareadmincluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlegkeonpremvmwareadmincluster/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster google_gkeonprem_vmware_admin_cluster}.
type GoogleGkeonpremVmwareAdminCluster interface {
	cdktf.TerraformResource
	AddonNode() GoogleGkeonpremVmwareAdminClusterAddonNodeOutputReference
	AddonNodeInput() *GoogleGkeonpremVmwareAdminClusterAddonNode
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	AntiAffinityGroups() GoogleGkeonpremVmwareAdminClusterAntiAffinityGroupsOutputReference
	AntiAffinityGroupsInput() *GoogleGkeonpremVmwareAdminClusterAntiAffinityGroups
	Authorization() GoogleGkeonpremVmwareAdminClusterAuthorizationOutputReference
	AuthorizationInput() *GoogleGkeonpremVmwareAdminClusterAuthorization
	AutoRepairConfig() GoogleGkeonpremVmwareAdminClusterAutoRepairConfigOutputReference
	AutoRepairConfigInput() *GoogleGkeonpremVmwareAdminClusterAutoRepairConfig
	BootstrapClusterMembership() *string
	SetBootstrapClusterMembership(val *string)
	BootstrapClusterMembershipInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ControlPlaneNode() GoogleGkeonpremVmwareAdminClusterControlPlaneNodeOutputReference
	ControlPlaneNodeInput() *GoogleGkeonpremVmwareAdminClusterControlPlaneNode
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveAnnotations() cdktf.StringMap
	EnableAdvancedCluster() cdktf.IResolvable
	Endpoint() *string
	Etag() *string
	Fleet() GoogleGkeonpremVmwareAdminClusterFleetList
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
	ImageType() *string
	SetImageType(val *string)
	ImageTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancer() GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference
	LoadBalancerInput() *GoogleGkeonpremVmwareAdminClusterLoadBalancer
	LocalName() *string
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkConfig() GoogleGkeonpremVmwareAdminClusterNetworkConfigOutputReference
	NetworkConfigInput() *GoogleGkeonpremVmwareAdminClusterNetworkConfig
	// The tree node.
	Node() constructs.Node
	OnPremVersion() *string
	SetOnPremVersion(val *string)
	OnPremVersionInput() *string
	PlatformConfig() GoogleGkeonpremVmwareAdminClusterPlatformConfigOutputReference
	PlatformConfigInput() *GoogleGkeonpremVmwareAdminClusterPlatformConfig
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
	Status() GoogleGkeonpremVmwareAdminClusterStatusList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleGkeonpremVmwareAdminClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
	UpdateTime() *string
	Vcenter() GoogleGkeonpremVmwareAdminClusterVcenterOutputReference
	VcenterInput() *GoogleGkeonpremVmwareAdminClusterVcenter
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
	PutAddonNode(value *GoogleGkeonpremVmwareAdminClusterAddonNode)
	PutAntiAffinityGroups(value *GoogleGkeonpremVmwareAdminClusterAntiAffinityGroups)
	PutAuthorization(value *GoogleGkeonpremVmwareAdminClusterAuthorization)
	PutAutoRepairConfig(value *GoogleGkeonpremVmwareAdminClusterAutoRepairConfig)
	PutControlPlaneNode(value *GoogleGkeonpremVmwareAdminClusterControlPlaneNode)
	PutLoadBalancer(value *GoogleGkeonpremVmwareAdminClusterLoadBalancer)
	PutNetworkConfig(value *GoogleGkeonpremVmwareAdminClusterNetworkConfig)
	PutPlatformConfig(value *GoogleGkeonpremVmwareAdminClusterPlatformConfig)
	PutTimeouts(value *GoogleGkeonpremVmwareAdminClusterTimeouts)
	PutVcenter(value *GoogleGkeonpremVmwareAdminClusterVcenter)
	ResetAddonNode()
	ResetAnnotations()
	ResetAntiAffinityGroups()
	ResetAuthorization()
	ResetAutoRepairConfig()
	ResetBootstrapClusterMembership()
	ResetControlPlaneNode()
	ResetDescription()
	ResetId()
	ResetImageType()
	ResetLoadBalancer()
	ResetOnPremVersion()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlatformConfig()
	ResetProject()
	ResetTimeouts()
	ResetVcenter()
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

// The jsii proxy struct for GoogleGkeonpremVmwareAdminCluster
type jsiiProxy_GoogleGkeonpremVmwareAdminCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) AddonNode() GoogleGkeonpremVmwareAdminClusterAddonNodeOutputReference {
	var returns GoogleGkeonpremVmwareAdminClusterAddonNodeOutputReference
	_jsii_.Get(
		j,
		"addonNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) AddonNodeInput() *GoogleGkeonpremVmwareAdminClusterAddonNode {
	var returns *GoogleGkeonpremVmwareAdminClusterAddonNode
	_jsii_.Get(
		j,
		"addonNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) AntiAffinityGroups() GoogleGkeonpremVmwareAdminClusterAntiAffinityGroupsOutputReference {
	var returns GoogleGkeonpremVmwareAdminClusterAntiAffinityGroupsOutputReference
	_jsii_.Get(
		j,
		"antiAffinityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) AntiAffinityGroupsInput() *GoogleGkeonpremVmwareAdminClusterAntiAffinityGroups {
	var returns *GoogleGkeonpremVmwareAdminClusterAntiAffinityGroups
	_jsii_.Get(
		j,
		"antiAffinityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Authorization() GoogleGkeonpremVmwareAdminClusterAuthorizationOutputReference {
	var returns GoogleGkeonpremVmwareAdminClusterAuthorizationOutputReference
	_jsii_.Get(
		j,
		"authorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) AuthorizationInput() *GoogleGkeonpremVmwareAdminClusterAuthorization {
	var returns *GoogleGkeonpremVmwareAdminClusterAuthorization
	_jsii_.Get(
		j,
		"authorizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) AutoRepairConfig() GoogleGkeonpremVmwareAdminClusterAutoRepairConfigOutputReference {
	var returns GoogleGkeonpremVmwareAdminClusterAutoRepairConfigOutputReference
	_jsii_.Get(
		j,
		"autoRepairConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) AutoRepairConfigInput() *GoogleGkeonpremVmwareAdminClusterAutoRepairConfig {
	var returns *GoogleGkeonpremVmwareAdminClusterAutoRepairConfig
	_jsii_.Get(
		j,
		"autoRepairConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) BootstrapClusterMembership() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapClusterMembership",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) BootstrapClusterMembershipInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapClusterMembershipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ControlPlaneNode() GoogleGkeonpremVmwareAdminClusterControlPlaneNodeOutputReference {
	var returns GoogleGkeonpremVmwareAdminClusterControlPlaneNodeOutputReference
	_jsii_.Get(
		j,
		"controlPlaneNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ControlPlaneNodeInput() *GoogleGkeonpremVmwareAdminClusterControlPlaneNode {
	var returns *GoogleGkeonpremVmwareAdminClusterControlPlaneNode
	_jsii_.Get(
		j,
		"controlPlaneNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) EffectiveAnnotations() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) EnableAdvancedCluster() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableAdvancedCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Fleet() GoogleGkeonpremVmwareAdminClusterFleetList {
	var returns GoogleGkeonpremVmwareAdminClusterFleetList
	_jsii_.Get(
		j,
		"fleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ImageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ImageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) LoadBalancer() GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference {
	var returns GoogleGkeonpremVmwareAdminClusterLoadBalancerOutputReference
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) LoadBalancerInput() *GoogleGkeonpremVmwareAdminClusterLoadBalancer {
	var returns *GoogleGkeonpremVmwareAdminClusterLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) LocalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) NetworkConfig() GoogleGkeonpremVmwareAdminClusterNetworkConfigOutputReference {
	var returns GoogleGkeonpremVmwareAdminClusterNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) NetworkConfigInput() *GoogleGkeonpremVmwareAdminClusterNetworkConfig {
	var returns *GoogleGkeonpremVmwareAdminClusterNetworkConfig
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) OnPremVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onPremVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) OnPremVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onPremVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) PlatformConfig() GoogleGkeonpremVmwareAdminClusterPlatformConfigOutputReference {
	var returns GoogleGkeonpremVmwareAdminClusterPlatformConfigOutputReference
	_jsii_.Get(
		j,
		"platformConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) PlatformConfigInput() *GoogleGkeonpremVmwareAdminClusterPlatformConfig {
	var returns *GoogleGkeonpremVmwareAdminClusterPlatformConfig
	_jsii_.Get(
		j,
		"platformConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Reconciling() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Status() GoogleGkeonpremVmwareAdminClusterStatusList {
	var returns GoogleGkeonpremVmwareAdminClusterStatusList
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Timeouts() GoogleGkeonpremVmwareAdminClusterTimeoutsOutputReference {
	var returns GoogleGkeonpremVmwareAdminClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) Vcenter() GoogleGkeonpremVmwareAdminClusterVcenterOutputReference {
	var returns GoogleGkeonpremVmwareAdminClusterVcenterOutputReference
	_jsii_.Get(
		j,
		"vcenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) VcenterInput() *GoogleGkeonpremVmwareAdminClusterVcenter {
	var returns *GoogleGkeonpremVmwareAdminClusterVcenter
	_jsii_.Get(
		j,
		"vcenterInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster google_gkeonprem_vmware_admin_cluster} Resource.
func NewGoogleGkeonpremVmwareAdminCluster(scope constructs.Construct, id *string, config *GoogleGkeonpremVmwareAdminClusterConfig) GoogleGkeonpremVmwareAdminCluster {
	_init_.Initialize()

	if err := validateNewGoogleGkeonpremVmwareAdminClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeonpremVmwareAdminCluster{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeonpremVmwareAdminCluster.GoogleGkeonpremVmwareAdminCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster google_gkeonprem_vmware_admin_cluster} Resource.
func NewGoogleGkeonpremVmwareAdminCluster_Override(g GoogleGkeonpremVmwareAdminCluster, scope constructs.Construct, id *string, config *GoogleGkeonpremVmwareAdminClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeonpremVmwareAdminCluster.GoogleGkeonpremVmwareAdminCluster",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster)SetBootstrapClusterMembership(val *string) {
	if err := j.validateSetBootstrapClusterMembershipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootstrapClusterMembership",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster)SetImageType(val *string) {
	if err := j.validateSetImageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageType",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster)SetOnPremVersion(val *string) {
	if err := j.validateSetOnPremVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onPremVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a GoogleGkeonpremVmwareAdminCluster resource upon running "cdktf plan <stack-name>".
func GoogleGkeonpremVmwareAdminCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleGkeonpremVmwareAdminCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleGkeonpremVmwareAdminCluster.GoogleGkeonpremVmwareAdminCluster",
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
func GoogleGkeonpremVmwareAdminCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleGkeonpremVmwareAdminCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleGkeonpremVmwareAdminCluster.GoogleGkeonpremVmwareAdminCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleGkeonpremVmwareAdminCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleGkeonpremVmwareAdminCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleGkeonpremVmwareAdminCluster.GoogleGkeonpremVmwareAdminCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleGkeonpremVmwareAdminCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleGkeonpremVmwareAdminCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleGkeonpremVmwareAdminCluster.GoogleGkeonpremVmwareAdminCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleGkeonpremVmwareAdminCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleGkeonpremVmwareAdminCluster.GoogleGkeonpremVmwareAdminCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) PutAddonNode(value *GoogleGkeonpremVmwareAdminClusterAddonNode) {
	if err := g.validatePutAddonNodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAddonNode",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) PutAntiAffinityGroups(value *GoogleGkeonpremVmwareAdminClusterAntiAffinityGroups) {
	if err := g.validatePutAntiAffinityGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAntiAffinityGroups",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) PutAuthorization(value *GoogleGkeonpremVmwareAdminClusterAuthorization) {
	if err := g.validatePutAuthorizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthorization",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) PutAutoRepairConfig(value *GoogleGkeonpremVmwareAdminClusterAutoRepairConfig) {
	if err := g.validatePutAutoRepairConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutoRepairConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) PutControlPlaneNode(value *GoogleGkeonpremVmwareAdminClusterControlPlaneNode) {
	if err := g.validatePutControlPlaneNodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putControlPlaneNode",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) PutLoadBalancer(value *GoogleGkeonpremVmwareAdminClusterLoadBalancer) {
	if err := g.validatePutLoadBalancerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLoadBalancer",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) PutNetworkConfig(value *GoogleGkeonpremVmwareAdminClusterNetworkConfig) {
	if err := g.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) PutPlatformConfig(value *GoogleGkeonpremVmwareAdminClusterPlatformConfig) {
	if err := g.validatePutPlatformConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPlatformConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) PutTimeouts(value *GoogleGkeonpremVmwareAdminClusterTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) PutVcenter(value *GoogleGkeonpremVmwareAdminClusterVcenter) {
	if err := g.validatePutVcenterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVcenter",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ResetAddonNode() {
	_jsii_.InvokeVoid(
		g,
		"resetAddonNode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ResetAnnotations() {
	_jsii_.InvokeVoid(
		g,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ResetAntiAffinityGroups() {
	_jsii_.InvokeVoid(
		g,
		"resetAntiAffinityGroups",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ResetAuthorization() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthorization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ResetAutoRepairConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoRepairConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ResetBootstrapClusterMembership() {
	_jsii_.InvokeVoid(
		g,
		"resetBootstrapClusterMembership",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ResetControlPlaneNode() {
	_jsii_.InvokeVoid(
		g,
		"resetControlPlaneNode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ResetImageType() {
	_jsii_.InvokeVoid(
		g,
		"resetImageType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ResetLoadBalancer() {
	_jsii_.InvokeVoid(
		g,
		"resetLoadBalancer",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ResetOnPremVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetOnPremVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ResetPlatformConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPlatformConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ResetVcenter() {
	_jsii_.InvokeVoid(
		g,
		"resetVcenter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


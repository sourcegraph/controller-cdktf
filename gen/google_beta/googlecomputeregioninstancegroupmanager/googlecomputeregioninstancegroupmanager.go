package googlecomputeregioninstancegroupmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeregioninstancegroupmanager/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_region_instance_group_manager google_compute_region_instance_group_manager}.
type GoogleComputeRegionInstanceGroupManager interface {
	cdktf.TerraformResource
	AllInstancesConfig() GoogleComputeRegionInstanceGroupManagerAllInstancesConfigOutputReference
	AllInstancesConfigInput() *GoogleComputeRegionInstanceGroupManagerAllInstancesConfig
	AutoHealingPolicies() GoogleComputeRegionInstanceGroupManagerAutoHealingPoliciesOutputReference
	AutoHealingPoliciesInput() *GoogleComputeRegionInstanceGroupManagerAutoHealingPolicies
	BaseInstanceName() *string
	SetBaseInstanceName(val *string)
	BaseInstanceNameInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DistributionPolicyTargetShape() *string
	SetDistributionPolicyTargetShape(val *string)
	DistributionPolicyTargetShapeInput() *string
	DistributionPolicyZones() *[]*string
	SetDistributionPolicyZones(val *[]*string)
	DistributionPolicyZonesInput() *[]*string
	Fingerprint() *string
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
	InstanceGroup() *string
	InstanceLifecyclePolicy() GoogleComputeRegionInstanceGroupManagerInstanceLifecyclePolicyOutputReference
	InstanceLifecyclePolicyInput() *GoogleComputeRegionInstanceGroupManagerInstanceLifecyclePolicy
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ListManagedInstancesResults() *string
	SetListManagedInstancesResults(val *string)
	ListManagedInstancesResultsInput() *string
	Name() *string
	SetName(val *string)
	NamedPort() GoogleComputeRegionInstanceGroupManagerNamedPortList
	NamedPortInput() interface{}
	NameInput() *string
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SelfLink() *string
	StatefulDisk() GoogleComputeRegionInstanceGroupManagerStatefulDiskList
	StatefulDiskInput() interface{}
	StatefulExternalIp() GoogleComputeRegionInstanceGroupManagerStatefulExternalIpList
	StatefulExternalIpInput() interface{}
	StatefulInternalIp() GoogleComputeRegionInstanceGroupManagerStatefulInternalIpList
	StatefulInternalIpInput() interface{}
	Status() GoogleComputeRegionInstanceGroupManagerStatusList
	TargetPools() *[]*string
	SetTargetPools(val *[]*string)
	TargetPoolsInput() *[]*string
	TargetSize() *float64
	SetTargetSize(val *float64)
	TargetSizeInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeRegionInstanceGroupManagerTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdatePolicy() GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference
	UpdatePolicyInput() *GoogleComputeRegionInstanceGroupManagerUpdatePolicy
	Version() GoogleComputeRegionInstanceGroupManagerVersionList
	VersionInput() interface{}
	WaitForInstances() interface{}
	SetWaitForInstances(val interface{})
	WaitForInstancesInput() interface{}
	WaitForInstancesStatus() *string
	SetWaitForInstancesStatus(val *string)
	WaitForInstancesStatusInput() *string
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
	PutAllInstancesConfig(value *GoogleComputeRegionInstanceGroupManagerAllInstancesConfig)
	PutAutoHealingPolicies(value *GoogleComputeRegionInstanceGroupManagerAutoHealingPolicies)
	PutInstanceLifecyclePolicy(value *GoogleComputeRegionInstanceGroupManagerInstanceLifecyclePolicy)
	PutNamedPort(value interface{})
	PutStatefulDisk(value interface{})
	PutStatefulExternalIp(value interface{})
	PutStatefulInternalIp(value interface{})
	PutTimeouts(value *GoogleComputeRegionInstanceGroupManagerTimeouts)
	PutUpdatePolicy(value *GoogleComputeRegionInstanceGroupManagerUpdatePolicy)
	PutVersion(value interface{})
	ResetAllInstancesConfig()
	ResetAutoHealingPolicies()
	ResetDescription()
	ResetDistributionPolicyTargetShape()
	ResetDistributionPolicyZones()
	ResetId()
	ResetInstanceLifecyclePolicy()
	ResetListManagedInstancesResults()
	ResetNamedPort()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRegion()
	ResetStatefulDisk()
	ResetStatefulExternalIp()
	ResetStatefulInternalIp()
	ResetTargetPools()
	ResetTargetSize()
	ResetTimeouts()
	ResetUpdatePolicy()
	ResetWaitForInstances()
	ResetWaitForInstancesStatus()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleComputeRegionInstanceGroupManager
type jsiiProxy_GoogleComputeRegionInstanceGroupManager struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) AllInstancesConfig() GoogleComputeRegionInstanceGroupManagerAllInstancesConfigOutputReference {
	var returns GoogleComputeRegionInstanceGroupManagerAllInstancesConfigOutputReference
	_jsii_.Get(
		j,
		"allInstancesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) AllInstancesConfigInput() *GoogleComputeRegionInstanceGroupManagerAllInstancesConfig {
	var returns *GoogleComputeRegionInstanceGroupManagerAllInstancesConfig
	_jsii_.Get(
		j,
		"allInstancesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) AutoHealingPolicies() GoogleComputeRegionInstanceGroupManagerAutoHealingPoliciesOutputReference {
	var returns GoogleComputeRegionInstanceGroupManagerAutoHealingPoliciesOutputReference
	_jsii_.Get(
		j,
		"autoHealingPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) AutoHealingPoliciesInput() *GoogleComputeRegionInstanceGroupManagerAutoHealingPolicies {
	var returns *GoogleComputeRegionInstanceGroupManagerAutoHealingPolicies
	_jsii_.Get(
		j,
		"autoHealingPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) BaseInstanceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseInstanceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) BaseInstanceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseInstanceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) DistributionPolicyTargetShape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionPolicyTargetShape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) DistributionPolicyTargetShapeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionPolicyTargetShapeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) DistributionPolicyZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"distributionPolicyZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) DistributionPolicyZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"distributionPolicyZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) Fingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) InstanceGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) InstanceLifecyclePolicy() GoogleComputeRegionInstanceGroupManagerInstanceLifecyclePolicyOutputReference {
	var returns GoogleComputeRegionInstanceGroupManagerInstanceLifecyclePolicyOutputReference
	_jsii_.Get(
		j,
		"instanceLifecyclePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) InstanceLifecyclePolicyInput() *GoogleComputeRegionInstanceGroupManagerInstanceLifecyclePolicy {
	var returns *GoogleComputeRegionInstanceGroupManagerInstanceLifecyclePolicy
	_jsii_.Get(
		j,
		"instanceLifecyclePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ListManagedInstancesResults() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listManagedInstancesResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ListManagedInstancesResultsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listManagedInstancesResultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) NamedPort() GoogleComputeRegionInstanceGroupManagerNamedPortList {
	var returns GoogleComputeRegionInstanceGroupManagerNamedPortList
	_jsii_.Get(
		j,
		"namedPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) NamedPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"namedPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) StatefulDisk() GoogleComputeRegionInstanceGroupManagerStatefulDiskList {
	var returns GoogleComputeRegionInstanceGroupManagerStatefulDiskList
	_jsii_.Get(
		j,
		"statefulDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) StatefulDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statefulDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) StatefulExternalIp() GoogleComputeRegionInstanceGroupManagerStatefulExternalIpList {
	var returns GoogleComputeRegionInstanceGroupManagerStatefulExternalIpList
	_jsii_.Get(
		j,
		"statefulExternalIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) StatefulExternalIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statefulExternalIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) StatefulInternalIp() GoogleComputeRegionInstanceGroupManagerStatefulInternalIpList {
	var returns GoogleComputeRegionInstanceGroupManagerStatefulInternalIpList
	_jsii_.Get(
		j,
		"statefulInternalIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) StatefulInternalIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statefulInternalIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) Status() GoogleComputeRegionInstanceGroupManagerStatusList {
	var returns GoogleComputeRegionInstanceGroupManagerStatusList
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) TargetPools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) TargetPoolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) TargetSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) TargetSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) Timeouts() GoogleComputeRegionInstanceGroupManagerTimeoutsOutputReference {
	var returns GoogleComputeRegionInstanceGroupManagerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) UpdatePolicy() GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference {
	var returns GoogleComputeRegionInstanceGroupManagerUpdatePolicyOutputReference
	_jsii_.Get(
		j,
		"updatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) UpdatePolicyInput() *GoogleComputeRegionInstanceGroupManagerUpdatePolicy {
	var returns *GoogleComputeRegionInstanceGroupManagerUpdatePolicy
	_jsii_.Get(
		j,
		"updatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) Version() GoogleComputeRegionInstanceGroupManagerVersionList {
	var returns GoogleComputeRegionInstanceGroupManagerVersionList
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) VersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) WaitForInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) WaitForInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) WaitForInstancesStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitForInstancesStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager) WaitForInstancesStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitForInstancesStatusInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_region_instance_group_manager google_compute_region_instance_group_manager} Resource.
func NewGoogleComputeRegionInstanceGroupManager(scope constructs.Construct, id *string, config *GoogleComputeRegionInstanceGroupManagerConfig) GoogleComputeRegionInstanceGroupManager {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionInstanceGroupManagerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionInstanceGroupManager{}

	_jsii_.Create(
		"google-beta.googleComputeRegionInstanceGroupManager.GoogleComputeRegionInstanceGroupManager",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_region_instance_group_manager google_compute_region_instance_group_manager} Resource.
func NewGoogleComputeRegionInstanceGroupManager_Override(g GoogleComputeRegionInstanceGroupManager, scope constructs.Construct, id *string, config *GoogleComputeRegionInstanceGroupManagerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeRegionInstanceGroupManager.GoogleComputeRegionInstanceGroupManager",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager)SetBaseInstanceName(val *string) {
	if err := j.validateSetBaseInstanceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseInstanceName",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager)SetDistributionPolicyTargetShape(val *string) {
	if err := j.validateSetDistributionPolicyTargetShapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"distributionPolicyTargetShape",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager)SetDistributionPolicyZones(val *[]*string) {
	if err := j.validateSetDistributionPolicyZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"distributionPolicyZones",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager)SetListManagedInstancesResults(val *string) {
	if err := j.validateSetListManagedInstancesResultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listManagedInstancesResults",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager)SetTargetPools(val *[]*string) {
	if err := j.validateSetTargetPoolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetPools",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager)SetTargetSize(val *float64) {
	if err := j.validateSetTargetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetSize",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager)SetWaitForInstances(val interface{}) {
	if err := j.validateSetWaitForInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForInstances",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManager)SetWaitForInstancesStatus(val *string) {
	if err := j.validateSetWaitForInstancesStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForInstancesStatus",
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
func GoogleComputeRegionInstanceGroupManager_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeRegionInstanceGroupManager_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeRegionInstanceGroupManager.GoogleComputeRegionInstanceGroupManager",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeRegionInstanceGroupManager_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeRegionInstanceGroupManager_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeRegionInstanceGroupManager.GoogleComputeRegionInstanceGroupManager",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeRegionInstanceGroupManager_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeRegionInstanceGroupManager_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeRegionInstanceGroupManager.GoogleComputeRegionInstanceGroupManager",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeRegionInstanceGroupManager_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.googleComputeRegionInstanceGroupManager.GoogleComputeRegionInstanceGroupManager",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) PutAllInstancesConfig(value *GoogleComputeRegionInstanceGroupManagerAllInstancesConfig) {
	if err := g.validatePutAllInstancesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAllInstancesConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) PutAutoHealingPolicies(value *GoogleComputeRegionInstanceGroupManagerAutoHealingPolicies) {
	if err := g.validatePutAutoHealingPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutoHealingPolicies",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) PutInstanceLifecyclePolicy(value *GoogleComputeRegionInstanceGroupManagerInstanceLifecyclePolicy) {
	if err := g.validatePutInstanceLifecyclePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInstanceLifecyclePolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) PutNamedPort(value interface{}) {
	if err := g.validatePutNamedPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNamedPort",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) PutStatefulDisk(value interface{}) {
	if err := g.validatePutStatefulDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStatefulDisk",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) PutStatefulExternalIp(value interface{}) {
	if err := g.validatePutStatefulExternalIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStatefulExternalIp",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) PutStatefulInternalIp(value interface{}) {
	if err := g.validatePutStatefulInternalIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStatefulInternalIp",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) PutTimeouts(value *GoogleComputeRegionInstanceGroupManagerTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) PutUpdatePolicy(value *GoogleComputeRegionInstanceGroupManagerUpdatePolicy) {
	if err := g.validatePutUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUpdatePolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) PutVersion(value interface{}) {
	if err := g.validatePutVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVersion",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ResetAllInstancesConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAllInstancesConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ResetAutoHealingPolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoHealingPolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ResetDistributionPolicyTargetShape() {
	_jsii_.InvokeVoid(
		g,
		"resetDistributionPolicyTargetShape",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ResetDistributionPolicyZones() {
	_jsii_.InvokeVoid(
		g,
		"resetDistributionPolicyZones",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ResetInstanceLifecyclePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceLifecyclePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ResetListManagedInstancesResults() {
	_jsii_.InvokeVoid(
		g,
		"resetListManagedInstancesResults",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ResetNamedPort() {
	_jsii_.InvokeVoid(
		g,
		"resetNamedPort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ResetStatefulDisk() {
	_jsii_.InvokeVoid(
		g,
		"resetStatefulDisk",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ResetStatefulExternalIp() {
	_jsii_.InvokeVoid(
		g,
		"resetStatefulExternalIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ResetStatefulInternalIp() {
	_jsii_.InvokeVoid(
		g,
		"resetStatefulInternalIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ResetTargetPools() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetPools",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ResetTargetSize() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ResetUpdatePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetUpdatePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ResetWaitForInstances() {
	_jsii_.InvokeVoid(
		g,
		"resetWaitForInstances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ResetWaitForInstancesStatus() {
	_jsii_.InvokeVoid(
		g,
		"resetWaitForInstancesStatus",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManager) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


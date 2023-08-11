package googlecomputeinstancegroupmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeinstancegroupmanager/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_instance_group_manager google_compute_instance_group_manager}.
type GoogleComputeInstanceGroupManager interface {
	cdktf.TerraformResource
	AllInstancesConfig() GoogleComputeInstanceGroupManagerAllInstancesConfigOutputReference
	AllInstancesConfigInput() *GoogleComputeInstanceGroupManagerAllInstancesConfig
	AutoHealingPolicies() GoogleComputeInstanceGroupManagerAutoHealingPoliciesOutputReference
	AutoHealingPoliciesInput() *GoogleComputeInstanceGroupManagerAutoHealingPolicies
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
	InstanceLifecyclePolicy() GoogleComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference
	InstanceLifecyclePolicyInput() *GoogleComputeInstanceGroupManagerInstanceLifecyclePolicy
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ListManagedInstancesResults() *string
	SetListManagedInstancesResults(val *string)
	ListManagedInstancesResultsInput() *string
	Name() *string
	SetName(val *string)
	NamedPort() GoogleComputeInstanceGroupManagerNamedPortList
	NamedPortInput() interface{}
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Operation() *string
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
	SelfLink() *string
	StatefulDisk() GoogleComputeInstanceGroupManagerStatefulDiskList
	StatefulDiskInput() interface{}
	StatefulExternalIp() GoogleComputeInstanceGroupManagerStatefulExternalIpList
	StatefulExternalIpInput() interface{}
	StatefulInternalIp() GoogleComputeInstanceGroupManagerStatefulInternalIpList
	StatefulInternalIpInput() interface{}
	Status() GoogleComputeInstanceGroupManagerStatusList
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
	Timeouts() GoogleComputeInstanceGroupManagerTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdatePolicy() GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference
	UpdatePolicyInput() *GoogleComputeInstanceGroupManagerUpdatePolicy
	Version() GoogleComputeInstanceGroupManagerVersionList
	VersionInput() interface{}
	WaitForInstances() interface{}
	SetWaitForInstances(val interface{})
	WaitForInstancesInput() interface{}
	WaitForInstancesStatus() *string
	SetWaitForInstancesStatus(val *string)
	WaitForInstancesStatusInput() *string
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
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
	PutAllInstancesConfig(value *GoogleComputeInstanceGroupManagerAllInstancesConfig)
	PutAutoHealingPolicies(value *GoogleComputeInstanceGroupManagerAutoHealingPolicies)
	PutInstanceLifecyclePolicy(value *GoogleComputeInstanceGroupManagerInstanceLifecyclePolicy)
	PutNamedPort(value interface{})
	PutStatefulDisk(value interface{})
	PutStatefulExternalIp(value interface{})
	PutStatefulInternalIp(value interface{})
	PutTimeouts(value *GoogleComputeInstanceGroupManagerTimeouts)
	PutUpdatePolicy(value *GoogleComputeInstanceGroupManagerUpdatePolicy)
	PutVersion(value interface{})
	ResetAllInstancesConfig()
	ResetAutoHealingPolicies()
	ResetDescription()
	ResetId()
	ResetInstanceLifecyclePolicy()
	ResetListManagedInstancesResults()
	ResetNamedPort()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetStatefulDisk()
	ResetStatefulExternalIp()
	ResetStatefulInternalIp()
	ResetTargetPools()
	ResetTargetSize()
	ResetTimeouts()
	ResetUpdatePolicy()
	ResetWaitForInstances()
	ResetWaitForInstancesStatus()
	ResetZone()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleComputeInstanceGroupManager
type jsiiProxy_GoogleComputeInstanceGroupManager struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) AllInstancesConfig() GoogleComputeInstanceGroupManagerAllInstancesConfigOutputReference {
	var returns GoogleComputeInstanceGroupManagerAllInstancesConfigOutputReference
	_jsii_.Get(
		j,
		"allInstancesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) AllInstancesConfigInput() *GoogleComputeInstanceGroupManagerAllInstancesConfig {
	var returns *GoogleComputeInstanceGroupManagerAllInstancesConfig
	_jsii_.Get(
		j,
		"allInstancesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) AutoHealingPolicies() GoogleComputeInstanceGroupManagerAutoHealingPoliciesOutputReference {
	var returns GoogleComputeInstanceGroupManagerAutoHealingPoliciesOutputReference
	_jsii_.Get(
		j,
		"autoHealingPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) AutoHealingPoliciesInput() *GoogleComputeInstanceGroupManagerAutoHealingPolicies {
	var returns *GoogleComputeInstanceGroupManagerAutoHealingPolicies
	_jsii_.Get(
		j,
		"autoHealingPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) BaseInstanceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseInstanceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) BaseInstanceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseInstanceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) Fingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) InstanceGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) InstanceLifecyclePolicy() GoogleComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference {
	var returns GoogleComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference
	_jsii_.Get(
		j,
		"instanceLifecyclePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) InstanceLifecyclePolicyInput() *GoogleComputeInstanceGroupManagerInstanceLifecyclePolicy {
	var returns *GoogleComputeInstanceGroupManagerInstanceLifecyclePolicy
	_jsii_.Get(
		j,
		"instanceLifecyclePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) ListManagedInstancesResults() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listManagedInstancesResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) ListManagedInstancesResultsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listManagedInstancesResultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) NamedPort() GoogleComputeInstanceGroupManagerNamedPortList {
	var returns GoogleComputeInstanceGroupManagerNamedPortList
	_jsii_.Get(
		j,
		"namedPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) NamedPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"namedPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) Operation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) StatefulDisk() GoogleComputeInstanceGroupManagerStatefulDiskList {
	var returns GoogleComputeInstanceGroupManagerStatefulDiskList
	_jsii_.Get(
		j,
		"statefulDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) StatefulDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statefulDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) StatefulExternalIp() GoogleComputeInstanceGroupManagerStatefulExternalIpList {
	var returns GoogleComputeInstanceGroupManagerStatefulExternalIpList
	_jsii_.Get(
		j,
		"statefulExternalIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) StatefulExternalIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statefulExternalIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) StatefulInternalIp() GoogleComputeInstanceGroupManagerStatefulInternalIpList {
	var returns GoogleComputeInstanceGroupManagerStatefulInternalIpList
	_jsii_.Get(
		j,
		"statefulInternalIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) StatefulInternalIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statefulInternalIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) Status() GoogleComputeInstanceGroupManagerStatusList {
	var returns GoogleComputeInstanceGroupManagerStatusList
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) TargetPools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) TargetPoolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) TargetSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) TargetSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) Timeouts() GoogleComputeInstanceGroupManagerTimeoutsOutputReference {
	var returns GoogleComputeInstanceGroupManagerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) UpdatePolicy() GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference {
	var returns GoogleComputeInstanceGroupManagerUpdatePolicyOutputReference
	_jsii_.Get(
		j,
		"updatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) UpdatePolicyInput() *GoogleComputeInstanceGroupManagerUpdatePolicy {
	var returns *GoogleComputeInstanceGroupManagerUpdatePolicy
	_jsii_.Get(
		j,
		"updatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) Version() GoogleComputeInstanceGroupManagerVersionList {
	var returns GoogleComputeInstanceGroupManagerVersionList
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) VersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) WaitForInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) WaitForInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) WaitForInstancesStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitForInstancesStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) WaitForInstancesStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitForInstancesStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_instance_group_manager google_compute_instance_group_manager} Resource.
func NewGoogleComputeInstanceGroupManager(scope constructs.Construct, id *string, config *GoogleComputeInstanceGroupManagerConfig) GoogleComputeInstanceGroupManager {
	_init_.Initialize()

	if err := validateNewGoogleComputeInstanceGroupManagerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeInstanceGroupManager{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeInstanceGroupManager.GoogleComputeInstanceGroupManager",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_instance_group_manager google_compute_instance_group_manager} Resource.
func NewGoogleComputeInstanceGroupManager_Override(g GoogleComputeInstanceGroupManager, scope constructs.Construct, id *string, config *GoogleComputeInstanceGroupManagerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeInstanceGroupManager.GoogleComputeInstanceGroupManager",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager)SetBaseInstanceName(val *string) {
	if err := j.validateSetBaseInstanceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseInstanceName",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager)SetListManagedInstancesResults(val *string) {
	if err := j.validateSetListManagedInstancesResultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listManagedInstancesResults",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager)SetTargetPools(val *[]*string) {
	if err := j.validateSetTargetPoolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetPools",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager)SetTargetSize(val *float64) {
	if err := j.validateSetTargetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetSize",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager)SetWaitForInstances(val interface{}) {
	if err := j.validateSetWaitForInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForInstances",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager)SetWaitForInstancesStatus(val *string) {
	if err := j.validateSetWaitForInstancesStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForInstancesStatus",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceGroupManager)SetZone(val *string) {
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
func GoogleComputeInstanceGroupManager_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeInstanceGroupManager_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeInstanceGroupManager.GoogleComputeInstanceGroupManager",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeInstanceGroupManager_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeInstanceGroupManager_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeInstanceGroupManager.GoogleComputeInstanceGroupManager",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeInstanceGroupManager_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeInstanceGroupManager_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeInstanceGroupManager.GoogleComputeInstanceGroupManager",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeInstanceGroupManager_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleComputeInstanceGroupManager.GoogleComputeInstanceGroupManager",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) PutAllInstancesConfig(value *GoogleComputeInstanceGroupManagerAllInstancesConfig) {
	if err := g.validatePutAllInstancesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAllInstancesConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) PutAutoHealingPolicies(value *GoogleComputeInstanceGroupManagerAutoHealingPolicies) {
	if err := g.validatePutAutoHealingPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutoHealingPolicies",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) PutInstanceLifecyclePolicy(value *GoogleComputeInstanceGroupManagerInstanceLifecyclePolicy) {
	if err := g.validatePutInstanceLifecyclePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInstanceLifecyclePolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) PutNamedPort(value interface{}) {
	if err := g.validatePutNamedPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNamedPort",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) PutStatefulDisk(value interface{}) {
	if err := g.validatePutStatefulDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStatefulDisk",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) PutStatefulExternalIp(value interface{}) {
	if err := g.validatePutStatefulExternalIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStatefulExternalIp",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) PutStatefulInternalIp(value interface{}) {
	if err := g.validatePutStatefulInternalIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStatefulInternalIp",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) PutTimeouts(value *GoogleComputeInstanceGroupManagerTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) PutUpdatePolicy(value *GoogleComputeInstanceGroupManagerUpdatePolicy) {
	if err := g.validatePutUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUpdatePolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) PutVersion(value interface{}) {
	if err := g.validatePutVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVersion",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ResetAllInstancesConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAllInstancesConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ResetAutoHealingPolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoHealingPolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ResetInstanceLifecyclePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceLifecyclePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ResetListManagedInstancesResults() {
	_jsii_.InvokeVoid(
		g,
		"resetListManagedInstancesResults",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ResetNamedPort() {
	_jsii_.InvokeVoid(
		g,
		"resetNamedPort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ResetStatefulDisk() {
	_jsii_.InvokeVoid(
		g,
		"resetStatefulDisk",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ResetStatefulExternalIp() {
	_jsii_.InvokeVoid(
		g,
		"resetStatefulExternalIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ResetStatefulInternalIp() {
	_jsii_.InvokeVoid(
		g,
		"resetStatefulInternalIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ResetTargetPools() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetPools",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ResetTargetSize() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ResetUpdatePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetUpdatePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ResetWaitForInstances() {
	_jsii_.InvokeVoid(
		g,
		"resetWaitForInstances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ResetWaitForInstancesStatus() {
	_jsii_.InvokeVoid(
		g,
		"resetWaitForInstancesStatus",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ResetZone() {
	_jsii_.InvokeVoid(
		g,
		"resetZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceGroupManager) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


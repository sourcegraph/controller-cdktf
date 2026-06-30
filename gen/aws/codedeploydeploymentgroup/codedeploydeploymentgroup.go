package codedeploydeploymentgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/codedeploydeploymentgroup/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/codedeploy_deployment_group aws_codedeploy_deployment_group}.
type CodedeployDeploymentGroup interface {
	cdktf.TerraformResource
	AlarmConfiguration() CodedeployDeploymentGroupAlarmConfigurationOutputReference
	AlarmConfigurationInput() *CodedeployDeploymentGroupAlarmConfiguration
	AppName() *string
	SetAppName(val *string)
	AppNameInput() *string
	Arn() *string
	AutoRollbackConfiguration() CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference
	AutoRollbackConfigurationInput() *CodedeployDeploymentGroupAutoRollbackConfiguration
	AutoscalingGroups() *[]*string
	SetAutoscalingGroups(val *[]*string)
	AutoscalingGroupsInput() *[]*string
	BlueGreenDeploymentConfig() CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference
	BlueGreenDeploymentConfigInput() *CodedeployDeploymentGroupBlueGreenDeploymentConfig
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComputePlatform() *string
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeploymentConfigName() *string
	SetDeploymentConfigName(val *string)
	DeploymentConfigNameInput() *string
	DeploymentGroupId() *string
	DeploymentGroupName() *string
	SetDeploymentGroupName(val *string)
	DeploymentGroupNameInput() *string
	DeploymentStyle() CodedeployDeploymentGroupDeploymentStyleOutputReference
	DeploymentStyleInput() *CodedeployDeploymentGroupDeploymentStyle
	Ec2TagFilter() CodedeployDeploymentGroupEc2TagFilterList
	Ec2TagFilterInput() any
	Ec2TagSet() CodedeployDeploymentGroupEc2TagSetList
	Ec2TagSetInput() any
	EcsService() CodedeployDeploymentGroupEcsServiceOutputReference
	EcsServiceInput() *CodedeployDeploymentGroupEcsService
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
	LoadBalancerInfo() CodedeployDeploymentGroupLoadBalancerInfoOutputReference
	LoadBalancerInfoInput() *CodedeployDeploymentGroupLoadBalancerInfo
	// The tree node.
	Node() constructs.Node
	OnPremisesInstanceTagFilter() CodedeployDeploymentGroupOnPremisesInstanceTagFilterList
	OnPremisesInstanceTagFilterInput() any
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	// Experimental.
	RawOverrides() any
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	ServiceRoleArnInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	TriggerConfiguration() CodedeployDeploymentGroupTriggerConfigurationList
	TriggerConfigurationInput() any
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value any)
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	HasResourceMove() any
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
	MoveTo(moveTarget *string, index any)
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAlarmConfiguration(value *CodedeployDeploymentGroupAlarmConfiguration)
	PutAutoRollbackConfiguration(value *CodedeployDeploymentGroupAutoRollbackConfiguration)
	PutBlueGreenDeploymentConfig(value *CodedeployDeploymentGroupBlueGreenDeploymentConfig)
	PutDeploymentStyle(value *CodedeployDeploymentGroupDeploymentStyle)
	PutEc2TagFilter(value any)
	PutEc2TagSet(value any)
	PutEcsService(value *CodedeployDeploymentGroupEcsService)
	PutLoadBalancerInfo(value *CodedeployDeploymentGroupLoadBalancerInfo)
	PutOnPremisesInstanceTagFilter(value any)
	PutTriggerConfiguration(value any)
	ResetAlarmConfiguration()
	ResetAutoRollbackConfiguration()
	ResetAutoscalingGroups()
	ResetBlueGreenDeploymentConfig()
	ResetDeploymentConfigName()
	ResetDeploymentStyle()
	ResetEc2TagFilter()
	ResetEc2TagSet()
	ResetEcsService()
	ResetId()
	ResetLoadBalancerInfo()
	ResetOnPremisesInstanceTagFilter()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	ResetTriggerConfiguration()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for CodedeployDeploymentGroup
type jsiiProxy_CodedeployDeploymentGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AlarmConfiguration() CodedeployDeploymentGroupAlarmConfigurationOutputReference {
	var returns CodedeployDeploymentGroupAlarmConfigurationOutputReference
	_jsii_.Get(
		j,
		"alarmConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AlarmConfigurationInput() *CodedeployDeploymentGroupAlarmConfiguration {
	var returns *CodedeployDeploymentGroupAlarmConfiguration
	_jsii_.Get(
		j,
		"alarmConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AppName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AppNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AutoRollbackConfiguration() CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference {
	var returns CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference
	_jsii_.Get(
		j,
		"autoRollbackConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AutoRollbackConfigurationInput() *CodedeployDeploymentGroupAutoRollbackConfiguration {
	var returns *CodedeployDeploymentGroupAutoRollbackConfiguration
	_jsii_.Get(
		j,
		"autoRollbackConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AutoscalingGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoscalingGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AutoscalingGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoscalingGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) BlueGreenDeploymentConfig() CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference {
	var returns CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference
	_jsii_.Get(
		j,
		"blueGreenDeploymentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) BlueGreenDeploymentConfigInput() *CodedeployDeploymentGroupBlueGreenDeploymentConfig {
	var returns *CodedeployDeploymentGroupBlueGreenDeploymentConfig
	_jsii_.Get(
		j,
		"blueGreenDeploymentConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) ComputePlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computePlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DeploymentConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DeploymentConfigNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DeploymentGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DeploymentGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DeploymentGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DeploymentStyle() CodedeployDeploymentGroupDeploymentStyleOutputReference {
	var returns CodedeployDeploymentGroupDeploymentStyleOutputReference
	_jsii_.Get(
		j,
		"deploymentStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DeploymentStyleInput() *CodedeployDeploymentGroupDeploymentStyle {
	var returns *CodedeployDeploymentGroupDeploymentStyle
	_jsii_.Get(
		j,
		"deploymentStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Ec2TagFilter() CodedeployDeploymentGroupEc2TagFilterList {
	var returns CodedeployDeploymentGroupEc2TagFilterList
	_jsii_.Get(
		j,
		"ec2TagFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Ec2TagFilterInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ec2TagFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Ec2TagSet() CodedeployDeploymentGroupEc2TagSetList {
	var returns CodedeployDeploymentGroupEc2TagSetList
	_jsii_.Get(
		j,
		"ec2TagSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Ec2TagSetInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ec2TagSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) EcsService() CodedeployDeploymentGroupEcsServiceOutputReference {
	var returns CodedeployDeploymentGroupEcsServiceOutputReference
	_jsii_.Get(
		j,
		"ecsService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) EcsServiceInput() *CodedeployDeploymentGroupEcsService {
	var returns *CodedeployDeploymentGroupEcsService
	_jsii_.Get(
		j,
		"ecsServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) LoadBalancerInfo() CodedeployDeploymentGroupLoadBalancerInfoOutputReference {
	var returns CodedeployDeploymentGroupLoadBalancerInfoOutputReference
	_jsii_.Get(
		j,
		"loadBalancerInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) LoadBalancerInfoInput() *CodedeployDeploymentGroupLoadBalancerInfo {
	var returns *CodedeployDeploymentGroupLoadBalancerInfo
	_jsii_.Get(
		j,
		"loadBalancerInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) OnPremisesInstanceTagFilter() CodedeployDeploymentGroupOnPremisesInstanceTagFilterList {
	var returns CodedeployDeploymentGroupOnPremisesInstanceTagFilterList
	_jsii_.Get(
		j,
		"onPremisesInstanceTagFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) OnPremisesInstanceTagFilterInput() any {
	var returns any
	_jsii_.Get(
		j,
		"onPremisesInstanceTagFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) ServiceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TriggerConfiguration() CodedeployDeploymentGroupTriggerConfigurationList {
	var returns CodedeployDeploymentGroupTriggerConfigurationList
	_jsii_.Get(
		j,
		"triggerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TriggerConfigurationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"triggerConfigurationInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/codedeploy_deployment_group aws_codedeploy_deployment_group} Resource.
func NewCodedeployDeploymentGroup(scope constructs.Construct, id *string, config *CodedeployDeploymentGroupConfig) CodedeployDeploymentGroup {
	_init_.Initialize()

	if err := validateNewCodedeployDeploymentGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodedeployDeploymentGroup{}

	_jsii_.Create(
		"@cdktf/provider-aws.codedeployDeploymentGroup.CodedeployDeploymentGroup",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/codedeploy_deployment_group aws_codedeploy_deployment_group} Resource.
func NewCodedeployDeploymentGroup_Override(c CodedeployDeploymentGroup, scope constructs.Construct, id *string, config *CodedeployDeploymentGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.codedeployDeploymentGroup.CodedeployDeploymentGroup",
		[]any{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetAppName(val *string) {
	if err := j.validateSetAppNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appName",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetAutoscalingGroups(val *[]*string) {
	if err := j.validateSetAutoscalingGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalingGroups",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetDeploymentConfigName(val *string) {
	if err := j.validateSetDeploymentConfigNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentConfigName",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetDeploymentGroupName(val *string) {
	if err := j.validateSetDeploymentGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentGroupName",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetServiceRoleArn(val *string) {
	if err := j.validateSetServiceRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup) SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a CodedeployDeploymentGroup resource upon running "cdktf plan <stack-name>".
func CodedeployDeploymentGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCodedeployDeploymentGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.codedeployDeploymentGroup.CodedeployDeploymentGroup",
		"generateConfigForImport",
		[]any{scope, importToId, importFromId, provider},
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
func CodedeployDeploymentGroup_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateCodedeployDeploymentGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.codedeployDeploymentGroup.CodedeployDeploymentGroup",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func CodedeployDeploymentGroup_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateCodedeployDeploymentGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.codedeployDeploymentGroup.CodedeployDeploymentGroup",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func CodedeployDeploymentGroup_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateCodedeployDeploymentGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.codedeployDeploymentGroup.CodedeployDeploymentGroup",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func CodedeployDeploymentGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.codedeployDeploymentGroup.CodedeployDeploymentGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) AddOverride(path *string, value any) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]any{path, value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]any{id, provider},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]any{id},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) MoveTo(moveTarget *string, index any) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]any{id},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutAlarmConfiguration(value *CodedeployDeploymentGroupAlarmConfiguration) {
	if err := c.validatePutAlarmConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAlarmConfiguration",
		[]any{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutAutoRollbackConfiguration(value *CodedeployDeploymentGroupAutoRollbackConfiguration) {
	if err := c.validatePutAutoRollbackConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAutoRollbackConfiguration",
		[]any{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutBlueGreenDeploymentConfig(value *CodedeployDeploymentGroupBlueGreenDeploymentConfig) {
	if err := c.validatePutBlueGreenDeploymentConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBlueGreenDeploymentConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutDeploymentStyle(value *CodedeployDeploymentGroupDeploymentStyle) {
	if err := c.validatePutDeploymentStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDeploymentStyle",
		[]any{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutEc2TagFilter(value any) {
	if err := c.validatePutEc2TagFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEc2TagFilter",
		[]any{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutEc2TagSet(value any) {
	if err := c.validatePutEc2TagSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEc2TagSet",
		[]any{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutEcsService(value *CodedeployDeploymentGroupEcsService) {
	if err := c.validatePutEcsServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEcsService",
		[]any{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutLoadBalancerInfo(value *CodedeployDeploymentGroupLoadBalancerInfo) {
	if err := c.validatePutLoadBalancerInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLoadBalancerInfo",
		[]any{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutOnPremisesInstanceTagFilter(value any) {
	if err := c.validatePutOnPremisesInstanceTagFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOnPremisesInstanceTagFilter",
		[]any{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutTriggerConfiguration(value any) {
	if err := c.validatePutTriggerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTriggerConfiguration",
		[]any{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetAlarmConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetAlarmConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetAutoRollbackConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoRollbackConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetAutoscalingGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoscalingGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetBlueGreenDeploymentConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetBlueGreenDeploymentConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetDeploymentConfigName() {
	_jsii_.InvokeVoid(
		c,
		"resetDeploymentConfigName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetDeploymentStyle() {
	_jsii_.InvokeVoid(
		c,
		"resetDeploymentStyle",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetEc2TagFilter() {
	_jsii_.InvokeVoid(
		c,
		"resetEc2TagFilter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetEc2TagSet() {
	_jsii_.InvokeVoid(
		c,
		"resetEc2TagSet",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetEcsService() {
	_jsii_.InvokeVoid(
		c,
		"resetEcsService",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetLoadBalancerInfo() {
	_jsii_.InvokeVoid(
		c,
		"resetLoadBalancerInfo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetOnPremisesInstanceTagFilter() {
	_jsii_.InvokeVoid(
		c,
		"resetOnPremisesInstanceTagFilter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetTagsAll() {
	_jsii_.InvokeVoid(
		c,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetTriggerConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetTriggerConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

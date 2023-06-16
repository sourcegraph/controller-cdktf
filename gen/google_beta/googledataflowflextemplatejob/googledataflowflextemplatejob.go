package googledataflowflextemplatejob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataflowflextemplatejob/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_dataflow_flex_template_job google_dataflow_flex_template_job}.
type GoogleDataflowFlexTemplateJob interface {
	cdktf.TerraformResource
	AdditionalExperiments() *[]*string
	SetAdditionalExperiments(val *[]*string)
	AdditionalExperimentsInput() *[]*string
	AutoscalingAlgorithm() *string
	SetAutoscalingAlgorithm(val *string)
	AutoscalingAlgorithmInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerSpecGcsPath() *string
	SetContainerSpecGcsPath(val *string)
	ContainerSpecGcsPathInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableStreamingEngine() interface{}
	SetEnableStreamingEngine(val interface{})
	EnableStreamingEngineInput() interface{}
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
	IpConfiguration() *string
	SetIpConfiguration(val *string)
	IpConfigurationInput() *string
	JobId() *string
	KmsKeyName() *string
	SetKmsKeyName(val *string)
	KmsKeyNameInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LauncherMachineType() *string
	SetLauncherMachineType(val *string)
	LauncherMachineTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	MaxWorkers() *float64
	SetMaxWorkers(val *float64)
	MaxWorkersInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	// The tree node.
	Node() constructs.Node
	NumWorkers() *float64
	SetNumWorkers(val *float64)
	NumWorkersInput() *float64
	OnDelete() *string
	SetOnDelete(val *string)
	OnDeleteInput() *string
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
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
	SdkContainerImage() *string
	SetSdkContainerImage(val *string)
	SdkContainerImageInput() *string
	ServiceAccountEmail() *string
	SetServiceAccountEmail(val *string)
	ServiceAccountEmailInput() *string
	SkipWaitOnJobTermination() interface{}
	SetSkipWaitOnJobTermination(val interface{})
	SkipWaitOnJobTerminationInput() interface{}
	StagingLocation() *string
	SetStagingLocation(val *string)
	StagingLocationInput() *string
	State() *string
	Subnetwork() *string
	SetSubnetwork(val *string)
	SubnetworkInput() *string
	TempLocation() *string
	SetTempLocation(val *string)
	TempLocationInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TransformNameMapping() *map[string]*string
	SetTransformNameMapping(val *map[string]*string)
	TransformNameMappingInput() *map[string]*string
	Type() *string
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
	ResetAdditionalExperiments()
	ResetAutoscalingAlgorithm()
	ResetEnableStreamingEngine()
	ResetId()
	ResetIpConfiguration()
	ResetKmsKeyName()
	ResetLabels()
	ResetLauncherMachineType()
	ResetMachineType()
	ResetMaxWorkers()
	ResetNetwork()
	ResetNumWorkers()
	ResetOnDelete()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
	ResetProject()
	ResetRegion()
	ResetSdkContainerImage()
	ResetServiceAccountEmail()
	ResetSkipWaitOnJobTermination()
	ResetStagingLocation()
	ResetSubnetwork()
	ResetTempLocation()
	ResetTransformNameMapping()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleDataflowFlexTemplateJob
type jsiiProxy_GoogleDataflowFlexTemplateJob struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) AdditionalExperiments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalExperiments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) AdditionalExperimentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalExperimentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) AutoscalingAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalingAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) AutoscalingAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalingAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) ContainerSpecGcsPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerSpecGcsPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) ContainerSpecGcsPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerSpecGcsPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) EnableStreamingEngine() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStreamingEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) EnableStreamingEngineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStreamingEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) IpConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) IpConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) JobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) KmsKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) KmsKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) LauncherMachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launcherMachineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) LauncherMachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launcherMachineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) MaxWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) MaxWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) NumWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) NumWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) OnDelete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) OnDeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) SdkContainerImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdkContainerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) SdkContainerImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdkContainerImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) ServiceAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) ServiceAccountEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) SkipWaitOnJobTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipWaitOnJobTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) SkipWaitOnJobTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipWaitOnJobTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) StagingLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) StagingLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) SubnetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) TempLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tempLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) TempLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tempLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) TransformNameMapping() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"transformNameMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) TransformNameMappingInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"transformNameMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_dataflow_flex_template_job google_dataflow_flex_template_job} Resource.
func NewGoogleDataflowFlexTemplateJob(scope constructs.Construct, id *string, config *GoogleDataflowFlexTemplateJobConfig) GoogleDataflowFlexTemplateJob {
	_init_.Initialize()

	if err := validateNewGoogleDataflowFlexTemplateJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataflowFlexTemplateJob{}

	_jsii_.Create(
		"google-beta.googleDataflowFlexTemplateJob.GoogleDataflowFlexTemplateJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_dataflow_flex_template_job google_dataflow_flex_template_job} Resource.
func NewGoogleDataflowFlexTemplateJob_Override(g GoogleDataflowFlexTemplateJob, scope constructs.Construct, id *string, config *GoogleDataflowFlexTemplateJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataflowFlexTemplateJob.GoogleDataflowFlexTemplateJob",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetAdditionalExperiments(val *[]*string) {
	if err := j.validateSetAdditionalExperimentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalExperiments",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetAutoscalingAlgorithm(val *string) {
	if err := j.validateSetAutoscalingAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalingAlgorithm",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetContainerSpecGcsPath(val *string) {
	if err := j.validateSetContainerSpecGcsPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerSpecGcsPath",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetEnableStreamingEngine(val interface{}) {
	if err := j.validateSetEnableStreamingEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableStreamingEngine",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetIpConfiguration(val *string) {
	if err := j.validateSetIpConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipConfiguration",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetKmsKeyName(val *string) {
	if err := j.validateSetKmsKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyName",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetLauncherMachineType(val *string) {
	if err := j.validateSetLauncherMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launcherMachineType",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetMaxWorkers(val *float64) {
	if err := j.validateSetMaxWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWorkers",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetNumWorkers(val *float64) {
	if err := j.validateSetNumWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numWorkers",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetOnDelete(val *string) {
	if err := j.validateSetOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDelete",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetSdkContainerImage(val *string) {
	if err := j.validateSetSdkContainerImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdkContainerImage",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetServiceAccountEmail(val *string) {
	if err := j.validateSetServiceAccountEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountEmail",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetSkipWaitOnJobTermination(val interface{}) {
	if err := j.validateSetSkipWaitOnJobTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipWaitOnJobTermination",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetStagingLocation(val *string) {
	if err := j.validateSetStagingLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stagingLocation",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetSubnetwork(val *string) {
	if err := j.validateSetSubnetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetwork",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetTempLocation(val *string) {
	if err := j.validateSetTempLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tempLocation",
		val,
	)
}

func (j *jsiiProxy_GoogleDataflowFlexTemplateJob)SetTransformNameMapping(val *map[string]*string) {
	if err := j.validateSetTransformNameMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transformNameMapping",
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
func GoogleDataflowFlexTemplateJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDataflowFlexTemplateJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleDataflowFlexTemplateJob.GoogleDataflowFlexTemplateJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDataflowFlexTemplateJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDataflowFlexTemplateJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleDataflowFlexTemplateJob.GoogleDataflowFlexTemplateJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDataflowFlexTemplateJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDataflowFlexTemplateJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleDataflowFlexTemplateJob.GoogleDataflowFlexTemplateJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleDataflowFlexTemplateJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.googleDataflowFlexTemplateJob.GoogleDataflowFlexTemplateJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetAdditionalExperiments() {
	_jsii_.InvokeVoid(
		g,
		"resetAdditionalExperiments",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetAutoscalingAlgorithm() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoscalingAlgorithm",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetEnableStreamingEngine() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableStreamingEngine",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetIpConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetIpConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetKmsKeyName() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsKeyName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetLauncherMachineType() {
	_jsii_.InvokeVoid(
		g,
		"resetLauncherMachineType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetMachineType() {
	_jsii_.InvokeVoid(
		g,
		"resetMachineType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetMaxWorkers() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxWorkers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetNumWorkers() {
	_jsii_.InvokeVoid(
		g,
		"resetNumWorkers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetOnDelete() {
	_jsii_.InvokeVoid(
		g,
		"resetOnDelete",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetSdkContainerImage() {
	_jsii_.InvokeVoid(
		g,
		"resetSdkContainerImage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetServiceAccountEmail() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccountEmail",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetSkipWaitOnJobTermination() {
	_jsii_.InvokeVoid(
		g,
		"resetSkipWaitOnJobTermination",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetStagingLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetStagingLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetSubnetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetSubnetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetTempLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetTempLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ResetTransformNameMapping() {
	_jsii_.InvokeVoid(
		g,
		"resetTransformNameMapping",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataflowFlexTemplateJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


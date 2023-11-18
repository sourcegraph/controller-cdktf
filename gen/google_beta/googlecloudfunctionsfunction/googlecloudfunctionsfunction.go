package googlecloudfunctionsfunction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudfunctionsfunction/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.6.0/docs/resources/google_cloudfunctions_function google_cloudfunctions_function}.
type GoogleCloudfunctionsFunction interface {
	cdktf.TerraformResource
	AvailableMemoryMb() *float64
	SetAvailableMemoryMb(val *float64)
	AvailableMemoryMbInput() *float64
	BuildEnvironmentVariables() *map[string]*string
	SetBuildEnvironmentVariables(val *map[string]*string)
	BuildEnvironmentVariablesInput() *map[string]*string
	BuildWorkerPool() *string
	SetBuildWorkerPool(val *string)
	BuildWorkerPoolInput() *string
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
	DockerRegistry() *string
	SetDockerRegistry(val *string)
	DockerRegistryInput() *string
	DockerRepository() *string
	SetDockerRepository(val *string)
	DockerRepositoryInput() *string
	EffectiveLabels() cdktf.StringMap
	EntryPoint() *string
	SetEntryPoint(val *string)
	EntryPointInput() *string
	EnvironmentVariables() *map[string]*string
	SetEnvironmentVariables(val *map[string]*string)
	EnvironmentVariablesInput() *map[string]*string
	EventTrigger() GoogleCloudfunctionsFunctionEventTriggerOutputReference
	EventTriggerInput() *GoogleCloudfunctionsFunctionEventTrigger
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpsTriggerSecurityLevel() *string
	SetHttpsTriggerSecurityLevel(val *string)
	HttpsTriggerSecurityLevelInput() *string
	HttpsTriggerUrl() *string
	SetHttpsTriggerUrl(val *string)
	HttpsTriggerUrlInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IngressSettings() *string
	SetIngressSettings(val *string)
	IngressSettingsInput() *string
	KmsKeyName() *string
	SetKmsKeyName(val *string)
	KmsKeyNameInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxInstances() *float64
	SetMaxInstances(val *float64)
	MaxInstancesInput() *float64
	MinInstances() *float64
	SetMinInstances(val *float64)
	MinInstancesInput() *float64
	Name() *string
	SetName(val *string)
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
	Runtime() *string
	SetRuntime(val *string)
	RuntimeInput() *string
	SecretEnvironmentVariables() GoogleCloudfunctionsFunctionSecretEnvironmentVariablesList
	SecretEnvironmentVariablesInput() interface{}
	SecretVolumes() GoogleCloudfunctionsFunctionSecretVolumesList
	SecretVolumesInput() interface{}
	ServiceAccountEmail() *string
	SetServiceAccountEmail(val *string)
	ServiceAccountEmailInput() *string
	SourceArchiveBucket() *string
	SetSourceArchiveBucket(val *string)
	SourceArchiveBucketInput() *string
	SourceArchiveObject() *string
	SetSourceArchiveObject(val *string)
	SourceArchiveObjectInput() *string
	SourceRepository() GoogleCloudfunctionsFunctionSourceRepositoryOutputReference
	SourceRepositoryInput() *GoogleCloudfunctionsFunctionSourceRepository
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
	Timeouts() GoogleCloudfunctionsFunctionTimeoutsOutputReference
	TimeoutsInput() interface{}
	TriggerHttp() interface{}
	SetTriggerHttp(val interface{})
	TriggerHttpInput() interface{}
	VpcConnector() *string
	SetVpcConnector(val *string)
	VpcConnectorEgressSettings() *string
	SetVpcConnectorEgressSettings(val *string)
	VpcConnectorEgressSettingsInput() *string
	VpcConnectorInput() *string
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
	PutEventTrigger(value *GoogleCloudfunctionsFunctionEventTrigger)
	PutSecretEnvironmentVariables(value interface{})
	PutSecretVolumes(value interface{})
	PutSourceRepository(value *GoogleCloudfunctionsFunctionSourceRepository)
	PutTimeouts(value *GoogleCloudfunctionsFunctionTimeouts)
	ResetAvailableMemoryMb()
	ResetBuildEnvironmentVariables()
	ResetBuildWorkerPool()
	ResetDescription()
	ResetDockerRegistry()
	ResetDockerRepository()
	ResetEntryPoint()
	ResetEnvironmentVariables()
	ResetEventTrigger()
	ResetHttpsTriggerSecurityLevel()
	ResetHttpsTriggerUrl()
	ResetId()
	ResetIngressSettings()
	ResetKmsKeyName()
	ResetLabels()
	ResetMaxInstances()
	ResetMinInstances()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRegion()
	ResetSecretEnvironmentVariables()
	ResetSecretVolumes()
	ResetServiceAccountEmail()
	ResetSourceArchiveBucket()
	ResetSourceArchiveObject()
	ResetSourceRepository()
	ResetTimeout()
	ResetTimeouts()
	ResetTriggerHttp()
	ResetVpcConnector()
	ResetVpcConnectorEgressSettings()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleCloudfunctionsFunction
type jsiiProxy_GoogleCloudfunctionsFunction struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) AvailableMemoryMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableMemoryMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) AvailableMemoryMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableMemoryMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) BuildEnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"buildEnvironmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) BuildEnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"buildEnvironmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) BuildWorkerPool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildWorkerPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) BuildWorkerPoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildWorkerPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) DockerRegistry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRegistry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) DockerRegistryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRegistryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) DockerRepository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) DockerRepositoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) EntryPoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) EntryPointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) EnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) EnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) EventTrigger() GoogleCloudfunctionsFunctionEventTriggerOutputReference {
	var returns GoogleCloudfunctionsFunctionEventTriggerOutputReference
	_jsii_.Get(
		j,
		"eventTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) EventTriggerInput() *GoogleCloudfunctionsFunctionEventTrigger {
	var returns *GoogleCloudfunctionsFunctionEventTrigger
	_jsii_.Get(
		j,
		"eventTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) HttpsTriggerSecurityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsTriggerSecurityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) HttpsTriggerSecurityLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsTriggerSecurityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) HttpsTriggerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsTriggerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) HttpsTriggerUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsTriggerUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) IngressSettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingressSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) IngressSettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingressSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) KmsKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) KmsKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) MaxInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) MaxInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) MinInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) MinInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) SecretEnvironmentVariables() GoogleCloudfunctionsFunctionSecretEnvironmentVariablesList {
	var returns GoogleCloudfunctionsFunctionSecretEnvironmentVariablesList
	_jsii_.Get(
		j,
		"secretEnvironmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) SecretEnvironmentVariablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretEnvironmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) SecretVolumes() GoogleCloudfunctionsFunctionSecretVolumesList {
	var returns GoogleCloudfunctionsFunctionSecretVolumesList
	_jsii_.Get(
		j,
		"secretVolumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) SecretVolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretVolumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) ServiceAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) ServiceAccountEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) SourceArchiveBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArchiveBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) SourceArchiveBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArchiveBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) SourceArchiveObject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArchiveObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) SourceArchiveObjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArchiveObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) SourceRepository() GoogleCloudfunctionsFunctionSourceRepositoryOutputReference {
	var returns GoogleCloudfunctionsFunctionSourceRepositoryOutputReference
	_jsii_.Get(
		j,
		"sourceRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) SourceRepositoryInput() *GoogleCloudfunctionsFunctionSourceRepository {
	var returns *GoogleCloudfunctionsFunctionSourceRepository
	_jsii_.Get(
		j,
		"sourceRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) Timeouts() GoogleCloudfunctionsFunctionTimeoutsOutputReference {
	var returns GoogleCloudfunctionsFunctionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) TriggerHttp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggerHttp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) TriggerHttpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggerHttpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) VpcConnector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) VpcConnectorEgressSettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorEgressSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) VpcConnectorEgressSettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorEgressSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction) VpcConnectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.6.0/docs/resources/google_cloudfunctions_function google_cloudfunctions_function} Resource.
func NewGoogleCloudfunctionsFunction(scope constructs.Construct, id *string, config *GoogleCloudfunctionsFunctionConfig) GoogleCloudfunctionsFunction {
	_init_.Initialize()

	if err := validateNewGoogleCloudfunctionsFunctionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudfunctionsFunction{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudfunctionsFunction.GoogleCloudfunctionsFunction",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.6.0/docs/resources/google_cloudfunctions_function google_cloudfunctions_function} Resource.
func NewGoogleCloudfunctionsFunction_Override(g GoogleCloudfunctionsFunction, scope constructs.Construct, id *string, config *GoogleCloudfunctionsFunctionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudfunctionsFunction.GoogleCloudfunctionsFunction",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetAvailableMemoryMb(val *float64) {
	if err := j.validateSetAvailableMemoryMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availableMemoryMb",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetBuildEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetBuildEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildEnvironmentVariables",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetBuildWorkerPool(val *string) {
	if err := j.validateSetBuildWorkerPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildWorkerPool",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetDockerRegistry(val *string) {
	if err := j.validateSetDockerRegistryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerRegistry",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetDockerRepository(val *string) {
	if err := j.validateSetDockerRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerRepository",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetEntryPoint(val *string) {
	if err := j.validateSetEntryPointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryPoint",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetHttpsTriggerSecurityLevel(val *string) {
	if err := j.validateSetHttpsTriggerSecurityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsTriggerSecurityLevel",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetHttpsTriggerUrl(val *string) {
	if err := j.validateSetHttpsTriggerUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsTriggerUrl",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetIngressSettings(val *string) {
	if err := j.validateSetIngressSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressSettings",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetKmsKeyName(val *string) {
	if err := j.validateSetKmsKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyName",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetMaxInstances(val *float64) {
	if err := j.validateSetMaxInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxInstances",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetMinInstances(val *float64) {
	if err := j.validateSetMinInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minInstances",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetServiceAccountEmail(val *string) {
	if err := j.validateSetServiceAccountEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountEmail",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetSourceArchiveBucket(val *string) {
	if err := j.validateSetSourceArchiveBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceArchiveBucket",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetSourceArchiveObject(val *string) {
	if err := j.validateSetSourceArchiveObjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceArchiveObject",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetTriggerHttp(val interface{}) {
	if err := j.validateSetTriggerHttpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerHttp",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetVpcConnector(val *string) {
	if err := j.validateSetVpcConnectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcConnector",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctionsFunction)SetVpcConnectorEgressSettings(val *string) {
	if err := j.validateSetVpcConnectorEgressSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcConnectorEgressSettings",
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
func GoogleCloudfunctionsFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudfunctionsFunction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleCloudfunctionsFunction.GoogleCloudfunctionsFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCloudfunctionsFunction_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudfunctionsFunction_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleCloudfunctionsFunction.GoogleCloudfunctionsFunction",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCloudfunctionsFunction_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudfunctionsFunction_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleCloudfunctionsFunction.GoogleCloudfunctionsFunction",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleCloudfunctionsFunction_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleCloudfunctionsFunction.GoogleCloudfunctionsFunction",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudfunctionsFunction) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudfunctionsFunction) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudfunctionsFunction) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudfunctionsFunction) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudfunctionsFunction) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudfunctionsFunction) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudfunctionsFunction) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudfunctionsFunction) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudfunctionsFunction) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudfunctionsFunction) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) PutEventTrigger(value *GoogleCloudfunctionsFunctionEventTrigger) {
	if err := g.validatePutEventTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEventTrigger",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) PutSecretEnvironmentVariables(value interface{}) {
	if err := g.validatePutSecretEnvironmentVariablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecretEnvironmentVariables",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) PutSecretVolumes(value interface{}) {
	if err := g.validatePutSecretVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecretVolumes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) PutSourceRepository(value *GoogleCloudfunctionsFunctionSourceRepository) {
	if err := g.validatePutSourceRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSourceRepository",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) PutTimeouts(value *GoogleCloudfunctionsFunctionTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetAvailableMemoryMb() {
	_jsii_.InvokeVoid(
		g,
		"resetAvailableMemoryMb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetBuildEnvironmentVariables() {
	_jsii_.InvokeVoid(
		g,
		"resetBuildEnvironmentVariables",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetBuildWorkerPool() {
	_jsii_.InvokeVoid(
		g,
		"resetBuildWorkerPool",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetDockerRegistry() {
	_jsii_.InvokeVoid(
		g,
		"resetDockerRegistry",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetDockerRepository() {
	_jsii_.InvokeVoid(
		g,
		"resetDockerRepository",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetEntryPoint() {
	_jsii_.InvokeVoid(
		g,
		"resetEntryPoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		g,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetEventTrigger() {
	_jsii_.InvokeVoid(
		g,
		"resetEventTrigger",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetHttpsTriggerSecurityLevel() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpsTriggerSecurityLevel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetHttpsTriggerUrl() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpsTriggerUrl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetIngressSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetIngressSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetKmsKeyName() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsKeyName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetMaxInstances() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxInstances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetMinInstances() {
	_jsii_.InvokeVoid(
		g,
		"resetMinInstances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetSecretEnvironmentVariables() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretEnvironmentVariables",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetSecretVolumes() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretVolumes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetServiceAccountEmail() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccountEmail",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetSourceArchiveBucket() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceArchiveBucket",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetSourceArchiveObject() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceArchiveObject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetSourceRepository() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceRepository",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetTriggerHttp() {
	_jsii_.InvokeVoid(
		g,
		"resetTriggerHttp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetVpcConnector() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcConnector",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ResetVpcConnectorEgressSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcConnectorEgressSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudfunctionsFunction) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


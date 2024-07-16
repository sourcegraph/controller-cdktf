package googleappengineflexibleappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleappengineflexibleappversion/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version google_app_engine_flexible_app_version}.
type GoogleAppEngineFlexibleAppVersion interface {
	cdktf.TerraformResource
	ApiConfig() GoogleAppEngineFlexibleAppVersionApiConfigOutputReference
	ApiConfigInput() *GoogleAppEngineFlexibleAppVersionApiConfig
	AutomaticScaling() GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference
	AutomaticScalingInput() *GoogleAppEngineFlexibleAppVersionAutomaticScaling
	BetaSettings() *map[string]*string
	SetBetaSettings(val *map[string]*string)
	BetaSettingsInput() *map[string]*string
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
	DefaultExpiration() *string
	SetDefaultExpiration(val *string)
	DefaultExpirationInput() *string
	DeleteServiceOnDestroy() interface{}
	SetDeleteServiceOnDestroy(val interface{})
	DeleteServiceOnDestroyInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Deployment() GoogleAppEngineFlexibleAppVersionDeploymentOutputReference
	DeploymentInput() *GoogleAppEngineFlexibleAppVersionDeployment
	EndpointsApiService() GoogleAppEngineFlexibleAppVersionEndpointsApiServiceOutputReference
	EndpointsApiServiceInput() *GoogleAppEngineFlexibleAppVersionEndpointsApiService
	Entrypoint() GoogleAppEngineFlexibleAppVersionEntrypointOutputReference
	EntrypointInput() *GoogleAppEngineFlexibleAppVersionEntrypoint
	EnvVariables() *map[string]*string
	SetEnvVariables(val *map[string]*string)
	EnvVariablesInput() *map[string]*string
	FlexibleRuntimeSettings() GoogleAppEngineFlexibleAppVersionFlexibleRuntimeSettingsOutputReference
	FlexibleRuntimeSettingsInput() *GoogleAppEngineFlexibleAppVersionFlexibleRuntimeSettings
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Handlers() GoogleAppEngineFlexibleAppVersionHandlersList
	HandlersInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	InboundServices() *[]*string
	SetInboundServices(val *[]*string)
	InboundServicesInput() *[]*string
	InstanceClass() *string
	SetInstanceClass(val *string)
	InstanceClassInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LivenessCheck() GoogleAppEngineFlexibleAppVersionLivenessCheckOutputReference
	LivenessCheckInput() *GoogleAppEngineFlexibleAppVersionLivenessCheck
	ManualScaling() GoogleAppEngineFlexibleAppVersionManualScalingOutputReference
	ManualScalingInput() *GoogleAppEngineFlexibleAppVersionManualScaling
	Name() *string
	Network() GoogleAppEngineFlexibleAppVersionNetworkOutputReference
	NetworkInput() *GoogleAppEngineFlexibleAppVersionNetwork
	NobuildFilesRegex() *string
	SetNobuildFilesRegex(val *string)
	NobuildFilesRegexInput() *string
	// The tree node.
	Node() constructs.Node
	NoopOnDestroy() interface{}
	SetNoopOnDestroy(val interface{})
	NoopOnDestroyInput() interface{}
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
	ReadinessCheck() GoogleAppEngineFlexibleAppVersionReadinessCheckOutputReference
	ReadinessCheckInput() *GoogleAppEngineFlexibleAppVersionReadinessCheck
	Resources() GoogleAppEngineFlexibleAppVersionResourcesOutputReference
	ResourcesInput() *GoogleAppEngineFlexibleAppVersionResources
	Runtime() *string
	SetRuntime(val *string)
	RuntimeApiVersion() *string
	SetRuntimeApiVersion(val *string)
	RuntimeApiVersionInput() *string
	RuntimeChannel() *string
	SetRuntimeChannel(val *string)
	RuntimeChannelInput() *string
	RuntimeInput() *string
	RuntimeMainExecutablePath() *string
	SetRuntimeMainExecutablePath(val *string)
	RuntimeMainExecutablePathInput() *string
	Service() *string
	SetService(val *string)
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	ServiceInput() *string
	ServingStatus() *string
	SetServingStatus(val *string)
	ServingStatusInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleAppEngineFlexibleAppVersionTimeoutsOutputReference
	TimeoutsInput() interface{}
	VersionId() *string
	SetVersionId(val *string)
	VersionIdInput() *string
	VpcAccessConnector() GoogleAppEngineFlexibleAppVersionVpcAccessConnectorOutputReference
	VpcAccessConnectorInput() *GoogleAppEngineFlexibleAppVersionVpcAccessConnector
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
	PutApiConfig(value *GoogleAppEngineFlexibleAppVersionApiConfig)
	PutAutomaticScaling(value *GoogleAppEngineFlexibleAppVersionAutomaticScaling)
	PutDeployment(value *GoogleAppEngineFlexibleAppVersionDeployment)
	PutEndpointsApiService(value *GoogleAppEngineFlexibleAppVersionEndpointsApiService)
	PutEntrypoint(value *GoogleAppEngineFlexibleAppVersionEntrypoint)
	PutFlexibleRuntimeSettings(value *GoogleAppEngineFlexibleAppVersionFlexibleRuntimeSettings)
	PutHandlers(value interface{})
	PutLivenessCheck(value *GoogleAppEngineFlexibleAppVersionLivenessCheck)
	PutManualScaling(value *GoogleAppEngineFlexibleAppVersionManualScaling)
	PutNetwork(value *GoogleAppEngineFlexibleAppVersionNetwork)
	PutReadinessCheck(value *GoogleAppEngineFlexibleAppVersionReadinessCheck)
	PutResources(value *GoogleAppEngineFlexibleAppVersionResources)
	PutTimeouts(value *GoogleAppEngineFlexibleAppVersionTimeouts)
	PutVpcAccessConnector(value *GoogleAppEngineFlexibleAppVersionVpcAccessConnector)
	ResetApiConfig()
	ResetAutomaticScaling()
	ResetBetaSettings()
	ResetDefaultExpiration()
	ResetDeleteServiceOnDestroy()
	ResetDeployment()
	ResetEndpointsApiService()
	ResetEntrypoint()
	ResetEnvVariables()
	ResetFlexibleRuntimeSettings()
	ResetHandlers()
	ResetId()
	ResetInboundServices()
	ResetInstanceClass()
	ResetManualScaling()
	ResetNetwork()
	ResetNobuildFilesRegex()
	ResetNoopOnDestroy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetResources()
	ResetRuntimeApiVersion()
	ResetRuntimeChannel()
	ResetRuntimeMainExecutablePath()
	ResetServiceAccount()
	ResetServingStatus()
	ResetTimeouts()
	ResetVersionId()
	ResetVpcAccessConnector()
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

// The jsii proxy struct for GoogleAppEngineFlexibleAppVersion
type jsiiProxy_GoogleAppEngineFlexibleAppVersion struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ApiConfig() GoogleAppEngineFlexibleAppVersionApiConfigOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionApiConfigOutputReference
	_jsii_.Get(
		j,
		"apiConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ApiConfigInput() *GoogleAppEngineFlexibleAppVersionApiConfig {
	var returns *GoogleAppEngineFlexibleAppVersionApiConfig
	_jsii_.Get(
		j,
		"apiConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) AutomaticScaling() GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference
	_jsii_.Get(
		j,
		"automaticScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) AutomaticScalingInput() *GoogleAppEngineFlexibleAppVersionAutomaticScaling {
	var returns *GoogleAppEngineFlexibleAppVersionAutomaticScaling
	_jsii_.Get(
		j,
		"automaticScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) BetaSettings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"betaSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) BetaSettingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"betaSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) DefaultExpiration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) DefaultExpirationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) DeleteServiceOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteServiceOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) DeleteServiceOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteServiceOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) Deployment() GoogleAppEngineFlexibleAppVersionDeploymentOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionDeploymentOutputReference
	_jsii_.Get(
		j,
		"deployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) DeploymentInput() *GoogleAppEngineFlexibleAppVersionDeployment {
	var returns *GoogleAppEngineFlexibleAppVersionDeployment
	_jsii_.Get(
		j,
		"deploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) EndpointsApiService() GoogleAppEngineFlexibleAppVersionEndpointsApiServiceOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionEndpointsApiServiceOutputReference
	_jsii_.Get(
		j,
		"endpointsApiService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) EndpointsApiServiceInput() *GoogleAppEngineFlexibleAppVersionEndpointsApiService {
	var returns *GoogleAppEngineFlexibleAppVersionEndpointsApiService
	_jsii_.Get(
		j,
		"endpointsApiServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) Entrypoint() GoogleAppEngineFlexibleAppVersionEntrypointOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionEntrypointOutputReference
	_jsii_.Get(
		j,
		"entrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) EntrypointInput() *GoogleAppEngineFlexibleAppVersionEntrypoint {
	var returns *GoogleAppEngineFlexibleAppVersionEntrypoint
	_jsii_.Get(
		j,
		"entrypointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) EnvVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"envVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) EnvVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"envVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) FlexibleRuntimeSettings() GoogleAppEngineFlexibleAppVersionFlexibleRuntimeSettingsOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionFlexibleRuntimeSettingsOutputReference
	_jsii_.Get(
		j,
		"flexibleRuntimeSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) FlexibleRuntimeSettingsInput() *GoogleAppEngineFlexibleAppVersionFlexibleRuntimeSettings {
	var returns *GoogleAppEngineFlexibleAppVersionFlexibleRuntimeSettings
	_jsii_.Get(
		j,
		"flexibleRuntimeSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) Handlers() GoogleAppEngineFlexibleAppVersionHandlersList {
	var returns GoogleAppEngineFlexibleAppVersionHandlersList
	_jsii_.Get(
		j,
		"handlers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) HandlersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"handlersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) InboundServices() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inboundServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) InboundServicesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inboundServicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) InstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) InstanceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) LivenessCheck() GoogleAppEngineFlexibleAppVersionLivenessCheckOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionLivenessCheckOutputReference
	_jsii_.Get(
		j,
		"livenessCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) LivenessCheckInput() *GoogleAppEngineFlexibleAppVersionLivenessCheck {
	var returns *GoogleAppEngineFlexibleAppVersionLivenessCheck
	_jsii_.Get(
		j,
		"livenessCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ManualScaling() GoogleAppEngineFlexibleAppVersionManualScalingOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionManualScalingOutputReference
	_jsii_.Get(
		j,
		"manualScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ManualScalingInput() *GoogleAppEngineFlexibleAppVersionManualScaling {
	var returns *GoogleAppEngineFlexibleAppVersionManualScaling
	_jsii_.Get(
		j,
		"manualScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) Network() GoogleAppEngineFlexibleAppVersionNetworkOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionNetworkOutputReference
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) NetworkInput() *GoogleAppEngineFlexibleAppVersionNetwork {
	var returns *GoogleAppEngineFlexibleAppVersionNetwork
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) NobuildFilesRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nobuildFilesRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) NobuildFilesRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nobuildFilesRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) NoopOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noopOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) NoopOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noopOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ReadinessCheck() GoogleAppEngineFlexibleAppVersionReadinessCheckOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionReadinessCheckOutputReference
	_jsii_.Get(
		j,
		"readinessCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ReadinessCheckInput() *GoogleAppEngineFlexibleAppVersionReadinessCheck {
	var returns *GoogleAppEngineFlexibleAppVersionReadinessCheck
	_jsii_.Get(
		j,
		"readinessCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) Resources() GoogleAppEngineFlexibleAppVersionResourcesOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResourcesInput() *GoogleAppEngineFlexibleAppVersionResources {
	var returns *GoogleAppEngineFlexibleAppVersionResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) RuntimeApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeApiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) RuntimeApiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeApiVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) RuntimeChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) RuntimeChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) RuntimeMainExecutablePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeMainExecutablePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) RuntimeMainExecutablePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeMainExecutablePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ServingStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servingStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ServingStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servingStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) Timeouts() GoogleAppEngineFlexibleAppVersionTimeoutsOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) VersionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) VpcAccessConnector() GoogleAppEngineFlexibleAppVersionVpcAccessConnectorOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionVpcAccessConnectorOutputReference
	_jsii_.Get(
		j,
		"vpcAccessConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion) VpcAccessConnectorInput() *GoogleAppEngineFlexibleAppVersionVpcAccessConnector {
	var returns *GoogleAppEngineFlexibleAppVersionVpcAccessConnector
	_jsii_.Get(
		j,
		"vpcAccessConnectorInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version google_app_engine_flexible_app_version} Resource.
func NewGoogleAppEngineFlexibleAppVersion(scope constructs.Construct, id *string, config *GoogleAppEngineFlexibleAppVersionConfig) GoogleAppEngineFlexibleAppVersion {
	_init_.Initialize()

	if err := validateNewGoogleAppEngineFlexibleAppVersionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAppEngineFlexibleAppVersion{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleAppEngineFlexibleAppVersion.GoogleAppEngineFlexibleAppVersion",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version google_app_engine_flexible_app_version} Resource.
func NewGoogleAppEngineFlexibleAppVersion_Override(g GoogleAppEngineFlexibleAppVersion, scope constructs.Construct, id *string, config *GoogleAppEngineFlexibleAppVersionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleAppEngineFlexibleAppVersion.GoogleAppEngineFlexibleAppVersion",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetBetaSettings(val *map[string]*string) {
	if err := j.validateSetBetaSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"betaSettings",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetDefaultExpiration(val *string) {
	if err := j.validateSetDefaultExpirationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultExpiration",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetDeleteServiceOnDestroy(val interface{}) {
	if err := j.validateSetDeleteServiceOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteServiceOnDestroy",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetEnvVariables(val *map[string]*string) {
	if err := j.validateSetEnvVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"envVariables",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetInboundServices(val *[]*string) {
	if err := j.validateSetInboundServicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inboundServices",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetInstanceClass(val *string) {
	if err := j.validateSetInstanceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceClass",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetNobuildFilesRegex(val *string) {
	if err := j.validateSetNobuildFilesRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nobuildFilesRegex",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetNoopOnDestroy(val interface{}) {
	if err := j.validateSetNoopOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noopOnDestroy",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetRuntimeApiVersion(val *string) {
	if err := j.validateSetRuntimeApiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeApiVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetRuntimeChannel(val *string) {
	if err := j.validateSetRuntimeChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeChannel",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetRuntimeMainExecutablePath(val *string) {
	if err := j.validateSetRuntimeMainExecutablePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeMainExecutablePath",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetServingStatus(val *string) {
	if err := j.validateSetServingStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servingStatus",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersion)SetVersionId(val *string) {
	if err := j.validateSetVersionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionId",
		val,
	)
}

// Generates CDKTF code for importing a GoogleAppEngineFlexibleAppVersion resource upon running "cdktf plan <stack-name>".
func GoogleAppEngineFlexibleAppVersion_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleAppEngineFlexibleAppVersion_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleAppEngineFlexibleAppVersion.GoogleAppEngineFlexibleAppVersion",
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
func GoogleAppEngineFlexibleAppVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAppEngineFlexibleAppVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleAppEngineFlexibleAppVersion.GoogleAppEngineFlexibleAppVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleAppEngineFlexibleAppVersion_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAppEngineFlexibleAppVersion_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleAppEngineFlexibleAppVersion.GoogleAppEngineFlexibleAppVersion",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleAppEngineFlexibleAppVersion_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAppEngineFlexibleAppVersion_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleAppEngineFlexibleAppVersion.GoogleAppEngineFlexibleAppVersion",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleAppEngineFlexibleAppVersion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleAppEngineFlexibleAppVersion.GoogleAppEngineFlexibleAppVersion",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) PutApiConfig(value *GoogleAppEngineFlexibleAppVersionApiConfig) {
	if err := g.validatePutApiConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApiConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) PutAutomaticScaling(value *GoogleAppEngineFlexibleAppVersionAutomaticScaling) {
	if err := g.validatePutAutomaticScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutomaticScaling",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) PutDeployment(value *GoogleAppEngineFlexibleAppVersionDeployment) {
	if err := g.validatePutDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeployment",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) PutEndpointsApiService(value *GoogleAppEngineFlexibleAppVersionEndpointsApiService) {
	if err := g.validatePutEndpointsApiServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEndpointsApiService",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) PutEntrypoint(value *GoogleAppEngineFlexibleAppVersionEntrypoint) {
	if err := g.validatePutEntrypointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEntrypoint",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) PutFlexibleRuntimeSettings(value *GoogleAppEngineFlexibleAppVersionFlexibleRuntimeSettings) {
	if err := g.validatePutFlexibleRuntimeSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFlexibleRuntimeSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) PutHandlers(value interface{}) {
	if err := g.validatePutHandlersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHandlers",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) PutLivenessCheck(value *GoogleAppEngineFlexibleAppVersionLivenessCheck) {
	if err := g.validatePutLivenessCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLivenessCheck",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) PutManualScaling(value *GoogleAppEngineFlexibleAppVersionManualScaling) {
	if err := g.validatePutManualScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putManualScaling",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) PutNetwork(value *GoogleAppEngineFlexibleAppVersionNetwork) {
	if err := g.validatePutNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetwork",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) PutReadinessCheck(value *GoogleAppEngineFlexibleAppVersionReadinessCheck) {
	if err := g.validatePutReadinessCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putReadinessCheck",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) PutResources(value *GoogleAppEngineFlexibleAppVersionResources) {
	if err := g.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putResources",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) PutTimeouts(value *GoogleAppEngineFlexibleAppVersionTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) PutVpcAccessConnector(value *GoogleAppEngineFlexibleAppVersionVpcAccessConnector) {
	if err := g.validatePutVpcAccessConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVpcAccessConnector",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetApiConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetApiConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetAutomaticScaling() {
	_jsii_.InvokeVoid(
		g,
		"resetAutomaticScaling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetBetaSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetBetaSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetDefaultExpiration() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultExpiration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetDeleteServiceOnDestroy() {
	_jsii_.InvokeVoid(
		g,
		"resetDeleteServiceOnDestroy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetDeployment() {
	_jsii_.InvokeVoid(
		g,
		"resetDeployment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetEndpointsApiService() {
	_jsii_.InvokeVoid(
		g,
		"resetEndpointsApiService",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetEntrypoint() {
	_jsii_.InvokeVoid(
		g,
		"resetEntrypoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetEnvVariables() {
	_jsii_.InvokeVoid(
		g,
		"resetEnvVariables",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetFlexibleRuntimeSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetFlexibleRuntimeSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetHandlers() {
	_jsii_.InvokeVoid(
		g,
		"resetHandlers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetInboundServices() {
	_jsii_.InvokeVoid(
		g,
		"resetInboundServices",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetInstanceClass() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceClass",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetManualScaling() {
	_jsii_.InvokeVoid(
		g,
		"resetManualScaling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetNobuildFilesRegex() {
	_jsii_.InvokeVoid(
		g,
		"resetNobuildFilesRegex",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetNoopOnDestroy() {
	_jsii_.InvokeVoid(
		g,
		"resetNoopOnDestroy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetResources() {
	_jsii_.InvokeVoid(
		g,
		"resetResources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetRuntimeApiVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetRuntimeApiVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetRuntimeChannel() {
	_jsii_.InvokeVoid(
		g,
		"resetRuntimeChannel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetRuntimeMainExecutablePath() {
	_jsii_.InvokeVoid(
		g,
		"resetRuntimeMainExecutablePath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetServingStatus() {
	_jsii_.InvokeVoid(
		g,
		"resetServingStatus",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetVersionId() {
	_jsii_.InvokeVoid(
		g,
		"resetVersionId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ResetVpcAccessConnector() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcAccessConnector",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersion) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


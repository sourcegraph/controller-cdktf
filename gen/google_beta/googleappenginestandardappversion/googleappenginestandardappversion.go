package googleappenginestandardappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleappenginestandardappversion/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version google_app_engine_standard_app_version}.
type GoogleAppEngineStandardAppVersion interface {
	cdktf.TerraformResource
	AppEngineApis() interface{}
	SetAppEngineApis(val interface{})
	AppEngineApisInput() interface{}
	AutomaticScaling() GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference
	AutomaticScalingInput() *GoogleAppEngineStandardAppVersionAutomaticScaling
	BasicScaling() GoogleAppEngineStandardAppVersionBasicScalingOutputReference
	BasicScalingInput() *GoogleAppEngineStandardAppVersionBasicScaling
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
	DeleteServiceOnDestroy() interface{}
	SetDeleteServiceOnDestroy(val interface{})
	DeleteServiceOnDestroyInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Deployment() GoogleAppEngineStandardAppVersionDeploymentOutputReference
	DeploymentInput() *GoogleAppEngineStandardAppVersionDeployment
	Entrypoint() GoogleAppEngineStandardAppVersionEntrypointOutputReference
	EntrypointInput() *GoogleAppEngineStandardAppVersionEntrypoint
	EnvVariables() *map[string]*string
	SetEnvVariables(val *map[string]*string)
	EnvVariablesInput() *map[string]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Handlers() GoogleAppEngineStandardAppVersionHandlersList
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
	Libraries() GoogleAppEngineStandardAppVersionLibrariesList
	LibrariesInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ManualScaling() GoogleAppEngineStandardAppVersionManualScalingOutputReference
	ManualScalingInput() *GoogleAppEngineStandardAppVersionManualScaling
	Name() *string
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
	Runtime() *string
	SetRuntime(val *string)
	RuntimeApiVersion() *string
	SetRuntimeApiVersion(val *string)
	RuntimeApiVersionInput() *string
	RuntimeInput() *string
	Service() *string
	SetService(val *string)
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	ServiceInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Threadsafe() interface{}
	SetThreadsafe(val interface{})
	ThreadsafeInput() interface{}
	Timeouts() GoogleAppEngineStandardAppVersionTimeoutsOutputReference
	TimeoutsInput() interface{}
	VersionId() *string
	SetVersionId(val *string)
	VersionIdInput() *string
	VpcAccessConnector() GoogleAppEngineStandardAppVersionVpcAccessConnectorOutputReference
	VpcAccessConnectorInput() *GoogleAppEngineStandardAppVersionVpcAccessConnector
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
	PutAutomaticScaling(value *GoogleAppEngineStandardAppVersionAutomaticScaling)
	PutBasicScaling(value *GoogleAppEngineStandardAppVersionBasicScaling)
	PutDeployment(value *GoogleAppEngineStandardAppVersionDeployment)
	PutEntrypoint(value *GoogleAppEngineStandardAppVersionEntrypoint)
	PutHandlers(value interface{})
	PutLibraries(value interface{})
	PutManualScaling(value *GoogleAppEngineStandardAppVersionManualScaling)
	PutTimeouts(value *GoogleAppEngineStandardAppVersionTimeouts)
	PutVpcAccessConnector(value *GoogleAppEngineStandardAppVersionVpcAccessConnector)
	ResetAppEngineApis()
	ResetAutomaticScaling()
	ResetBasicScaling()
	ResetDeleteServiceOnDestroy()
	ResetEnvVariables()
	ResetHandlers()
	ResetId()
	ResetInboundServices()
	ResetInstanceClass()
	ResetLibraries()
	ResetManualScaling()
	ResetNoopOnDestroy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRuntimeApiVersion()
	ResetServiceAccount()
	ResetThreadsafe()
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

// The jsii proxy struct for GoogleAppEngineStandardAppVersion
type jsiiProxy_GoogleAppEngineStandardAppVersion struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) AppEngineApis() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appEngineApis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) AppEngineApisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appEngineApisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) AutomaticScaling() GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference {
	var returns GoogleAppEngineStandardAppVersionAutomaticScalingOutputReference
	_jsii_.Get(
		j,
		"automaticScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) AutomaticScalingInput() *GoogleAppEngineStandardAppVersionAutomaticScaling {
	var returns *GoogleAppEngineStandardAppVersionAutomaticScaling
	_jsii_.Get(
		j,
		"automaticScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) BasicScaling() GoogleAppEngineStandardAppVersionBasicScalingOutputReference {
	var returns GoogleAppEngineStandardAppVersionBasicScalingOutputReference
	_jsii_.Get(
		j,
		"basicScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) BasicScalingInput() *GoogleAppEngineStandardAppVersionBasicScaling {
	var returns *GoogleAppEngineStandardAppVersionBasicScaling
	_jsii_.Get(
		j,
		"basicScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) DeleteServiceOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteServiceOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) DeleteServiceOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteServiceOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) Deployment() GoogleAppEngineStandardAppVersionDeploymentOutputReference {
	var returns GoogleAppEngineStandardAppVersionDeploymentOutputReference
	_jsii_.Get(
		j,
		"deployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) DeploymentInput() *GoogleAppEngineStandardAppVersionDeployment {
	var returns *GoogleAppEngineStandardAppVersionDeployment
	_jsii_.Get(
		j,
		"deploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) Entrypoint() GoogleAppEngineStandardAppVersionEntrypointOutputReference {
	var returns GoogleAppEngineStandardAppVersionEntrypointOutputReference
	_jsii_.Get(
		j,
		"entrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) EntrypointInput() *GoogleAppEngineStandardAppVersionEntrypoint {
	var returns *GoogleAppEngineStandardAppVersionEntrypoint
	_jsii_.Get(
		j,
		"entrypointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) EnvVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"envVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) EnvVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"envVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) Handlers() GoogleAppEngineStandardAppVersionHandlersList {
	var returns GoogleAppEngineStandardAppVersionHandlersList
	_jsii_.Get(
		j,
		"handlers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) HandlersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"handlersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) InboundServices() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inboundServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) InboundServicesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inboundServicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) InstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) InstanceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) Libraries() GoogleAppEngineStandardAppVersionLibrariesList {
	var returns GoogleAppEngineStandardAppVersionLibrariesList
	_jsii_.Get(
		j,
		"libraries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) LibrariesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"librariesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) ManualScaling() GoogleAppEngineStandardAppVersionManualScalingOutputReference {
	var returns GoogleAppEngineStandardAppVersionManualScalingOutputReference
	_jsii_.Get(
		j,
		"manualScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) ManualScalingInput() *GoogleAppEngineStandardAppVersionManualScaling {
	var returns *GoogleAppEngineStandardAppVersionManualScaling
	_jsii_.Get(
		j,
		"manualScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) NoopOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noopOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) NoopOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noopOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) RuntimeApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeApiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) RuntimeApiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeApiVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) Threadsafe() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threadsafe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) ThreadsafeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threadsafeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) Timeouts() GoogleAppEngineStandardAppVersionTimeoutsOutputReference {
	var returns GoogleAppEngineStandardAppVersionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) VersionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) VpcAccessConnector() GoogleAppEngineStandardAppVersionVpcAccessConnectorOutputReference {
	var returns GoogleAppEngineStandardAppVersionVpcAccessConnectorOutputReference
	_jsii_.Get(
		j,
		"vpcAccessConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion) VpcAccessConnectorInput() *GoogleAppEngineStandardAppVersionVpcAccessConnector {
	var returns *GoogleAppEngineStandardAppVersionVpcAccessConnector
	_jsii_.Get(
		j,
		"vpcAccessConnectorInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version google_app_engine_standard_app_version} Resource.
func NewGoogleAppEngineStandardAppVersion(scope constructs.Construct, id *string, config *GoogleAppEngineStandardAppVersionConfig) GoogleAppEngineStandardAppVersion {
	_init_.Initialize()

	if err := validateNewGoogleAppEngineStandardAppVersionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAppEngineStandardAppVersion{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleAppEngineStandardAppVersion.GoogleAppEngineStandardAppVersion",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version google_app_engine_standard_app_version} Resource.
func NewGoogleAppEngineStandardAppVersion_Override(g GoogleAppEngineStandardAppVersion, scope constructs.Construct, id *string, config *GoogleAppEngineStandardAppVersionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleAppEngineStandardAppVersion.GoogleAppEngineStandardAppVersion",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion)SetAppEngineApis(val interface{}) {
	if err := j.validateSetAppEngineApisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appEngineApis",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion)SetDeleteServiceOnDestroy(val interface{}) {
	if err := j.validateSetDeleteServiceOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteServiceOnDestroy",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion)SetEnvVariables(val *map[string]*string) {
	if err := j.validateSetEnvVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"envVariables",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion)SetInboundServices(val *[]*string) {
	if err := j.validateSetInboundServicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inboundServices",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion)SetInstanceClass(val *string) {
	if err := j.validateSetInstanceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceClass",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion)SetNoopOnDestroy(val interface{}) {
	if err := j.validateSetNoopOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noopOnDestroy",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion)SetRuntimeApiVersion(val *string) {
	if err := j.validateSetRuntimeApiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeApiVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion)SetThreadsafe(val interface{}) {
	if err := j.validateSetThreadsafeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadsafe",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineStandardAppVersion)SetVersionId(val *string) {
	if err := j.validateSetVersionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionId",
		val,
	)
}

// Generates CDKTF code for importing a GoogleAppEngineStandardAppVersion resource upon running "cdktf plan <stack-name>".
func GoogleAppEngineStandardAppVersion_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleAppEngineStandardAppVersion_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleAppEngineStandardAppVersion.GoogleAppEngineStandardAppVersion",
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
func GoogleAppEngineStandardAppVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAppEngineStandardAppVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleAppEngineStandardAppVersion.GoogleAppEngineStandardAppVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleAppEngineStandardAppVersion_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAppEngineStandardAppVersion_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleAppEngineStandardAppVersion.GoogleAppEngineStandardAppVersion",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleAppEngineStandardAppVersion_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAppEngineStandardAppVersion_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleAppEngineStandardAppVersion.GoogleAppEngineStandardAppVersion",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleAppEngineStandardAppVersion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleAppEngineStandardAppVersion.GoogleAppEngineStandardAppVersion",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) PutAutomaticScaling(value *GoogleAppEngineStandardAppVersionAutomaticScaling) {
	if err := g.validatePutAutomaticScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutomaticScaling",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) PutBasicScaling(value *GoogleAppEngineStandardAppVersionBasicScaling) {
	if err := g.validatePutBasicScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBasicScaling",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) PutDeployment(value *GoogleAppEngineStandardAppVersionDeployment) {
	if err := g.validatePutDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeployment",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) PutEntrypoint(value *GoogleAppEngineStandardAppVersionEntrypoint) {
	if err := g.validatePutEntrypointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEntrypoint",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) PutHandlers(value interface{}) {
	if err := g.validatePutHandlersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHandlers",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) PutLibraries(value interface{}) {
	if err := g.validatePutLibrariesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLibraries",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) PutManualScaling(value *GoogleAppEngineStandardAppVersionManualScaling) {
	if err := g.validatePutManualScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putManualScaling",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) PutTimeouts(value *GoogleAppEngineStandardAppVersionTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) PutVpcAccessConnector(value *GoogleAppEngineStandardAppVersionVpcAccessConnector) {
	if err := g.validatePutVpcAccessConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVpcAccessConnector",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ResetAppEngineApis() {
	_jsii_.InvokeVoid(
		g,
		"resetAppEngineApis",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ResetAutomaticScaling() {
	_jsii_.InvokeVoid(
		g,
		"resetAutomaticScaling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ResetBasicScaling() {
	_jsii_.InvokeVoid(
		g,
		"resetBasicScaling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ResetDeleteServiceOnDestroy() {
	_jsii_.InvokeVoid(
		g,
		"resetDeleteServiceOnDestroy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ResetEnvVariables() {
	_jsii_.InvokeVoid(
		g,
		"resetEnvVariables",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ResetHandlers() {
	_jsii_.InvokeVoid(
		g,
		"resetHandlers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ResetInboundServices() {
	_jsii_.InvokeVoid(
		g,
		"resetInboundServices",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ResetInstanceClass() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceClass",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ResetLibraries() {
	_jsii_.InvokeVoid(
		g,
		"resetLibraries",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ResetManualScaling() {
	_jsii_.InvokeVoid(
		g,
		"resetManualScaling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ResetNoopOnDestroy() {
	_jsii_.InvokeVoid(
		g,
		"resetNoopOnDestroy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ResetRuntimeApiVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetRuntimeApiVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ResetThreadsafe() {
	_jsii_.InvokeVoid(
		g,
		"resetThreadsafe",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ResetVersionId() {
	_jsii_.InvokeVoid(
		g,
		"resetVersionId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ResetVpcAccessConnector() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcAccessConnector",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineStandardAppVersion) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


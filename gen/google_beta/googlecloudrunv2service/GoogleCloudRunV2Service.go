package googlecloudrunv2service

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudrunv2service/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_cloud_run_v2_service google_cloud_run_v2_service}.
type GoogleCloudRunV2Service interface {
	cdktf.TerraformResource
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	BinaryAuthorization() GoogleCloudRunV2ServiceBinaryAuthorizationOutputReference
	BinaryAuthorizationInput() *GoogleCloudRunV2ServiceBinaryAuthorization
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Client() *string
	SetClient(val *string)
	ClientInput() *string
	ClientVersion() *string
	SetClientVersion(val *string)
	ClientVersionInput() *string
	Conditions() GoogleCloudRunV2ServiceConditionsList
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
	CreateTime() *string
	Creator() *string
	CustomAudiences() *[]*string
	SetCustomAudiences(val *[]*string)
	CustomAudiencesInput() *[]*string
	DefaultUriDisabled() interface{}
	SetDefaultUriDisabled(val interface{})
	DefaultUriDisabledInput() interface{}
	DeleteTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveAnnotations() cdktf.StringMap
	EffectiveLabels() cdktf.StringMap
	Etag() *string
	ExpireTime() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Generation() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Ingress() *string
	SetIngress(val *string)
	IngressInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LastModifier() *string
	LatestCreatedRevision() *string
	LatestReadyRevision() *string
	LaunchStage() *string
	SetLaunchStage(val *string)
	LaunchStageInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ObservedGeneration() *string
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
	Scaling() GoogleCloudRunV2ServiceScalingOutputReference
	ScalingInput() *GoogleCloudRunV2ServiceScaling
	Template() GoogleCloudRunV2ServiceTemplateOutputReference
	TemplateInput() *GoogleCloudRunV2ServiceTemplate
	TerminalCondition() GoogleCloudRunV2ServiceTerminalConditionList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleCloudRunV2ServiceTimeoutsOutputReference
	TimeoutsInput() interface{}
	Traffic() GoogleCloudRunV2ServiceTrafficList
	TrafficInput() interface{}
	TrafficStatuses() GoogleCloudRunV2ServiceTrafficStatusesList
	Uid() *string
	UpdateTime() *string
	Uri() *string
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
	PutBinaryAuthorization(value *GoogleCloudRunV2ServiceBinaryAuthorization)
	PutScaling(value *GoogleCloudRunV2ServiceScaling)
	PutTemplate(value *GoogleCloudRunV2ServiceTemplate)
	PutTimeouts(value *GoogleCloudRunV2ServiceTimeouts)
	PutTraffic(value interface{})
	ResetAnnotations()
	ResetBinaryAuthorization()
	ResetClient()
	ResetClientVersion()
	ResetCustomAudiences()
	ResetDefaultUriDisabled()
	ResetDescription()
	ResetId()
	ResetIngress()
	ResetLabels()
	ResetLaunchStage()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetScaling()
	ResetTimeouts()
	ResetTraffic()
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

// The jsii proxy struct for GoogleCloudRunV2Service
type jsiiProxy_GoogleCloudRunV2Service struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) BinaryAuthorization() GoogleCloudRunV2ServiceBinaryAuthorizationOutputReference {
	var returns GoogleCloudRunV2ServiceBinaryAuthorizationOutputReference
	_jsii_.Get(
		j,
		"binaryAuthorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) BinaryAuthorizationInput() *GoogleCloudRunV2ServiceBinaryAuthorization {
	var returns *GoogleCloudRunV2ServiceBinaryAuthorization
	_jsii_.Get(
		j,
		"binaryAuthorizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Client() *string {
	var returns *string
	_jsii_.Get(
		j,
		"client",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) ClientInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) ClientVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) ClientVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Conditions() GoogleCloudRunV2ServiceConditionsList {
	var returns GoogleCloudRunV2ServiceConditionsList
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Creator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) CustomAudiences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customAudiences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) CustomAudiencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customAudiencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) DefaultUriDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultUriDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) DefaultUriDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultUriDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) DeleteTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) EffectiveAnnotations() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) ExpireTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expireTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Generation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Ingress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) IngressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) LastModifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) LatestCreatedRevision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latestCreatedRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) LatestReadyRevision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latestReadyRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) LaunchStage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) LaunchStageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchStageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) ObservedGeneration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"observedGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Reconciling() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Scaling() GoogleCloudRunV2ServiceScalingOutputReference {
	var returns GoogleCloudRunV2ServiceScalingOutputReference
	_jsii_.Get(
		j,
		"scaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) ScalingInput() *GoogleCloudRunV2ServiceScaling {
	var returns *GoogleCloudRunV2ServiceScaling
	_jsii_.Get(
		j,
		"scalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Template() GoogleCloudRunV2ServiceTemplateOutputReference {
	var returns GoogleCloudRunV2ServiceTemplateOutputReference
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) TemplateInput() *GoogleCloudRunV2ServiceTemplate {
	var returns *GoogleCloudRunV2ServiceTemplate
	_jsii_.Get(
		j,
		"templateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) TerminalCondition() GoogleCloudRunV2ServiceTerminalConditionList {
	var returns GoogleCloudRunV2ServiceTerminalConditionList
	_jsii_.Get(
		j,
		"terminalCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Timeouts() GoogleCloudRunV2ServiceTimeoutsOutputReference {
	var returns GoogleCloudRunV2ServiceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Traffic() GoogleCloudRunV2ServiceTrafficList {
	var returns GoogleCloudRunV2ServiceTrafficList
	_jsii_.Get(
		j,
		"traffic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) TrafficInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) TrafficStatuses() GoogleCloudRunV2ServiceTrafficStatusesList {
	var returns GoogleCloudRunV2ServiceTrafficStatusesList
	_jsii_.Get(
		j,
		"trafficStatuses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Service) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_cloud_run_v2_service google_cloud_run_v2_service} Resource.
func NewGoogleCloudRunV2Service(scope constructs.Construct, id *string, config *GoogleCloudRunV2ServiceConfig) GoogleCloudRunV2Service {
	_init_.Initialize()

	if err := validateNewGoogleCloudRunV2ServiceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudRunV2Service{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudRunV2Service.GoogleCloudRunV2Service",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_cloud_run_v2_service google_cloud_run_v2_service} Resource.
func NewGoogleCloudRunV2Service_Override(g GoogleCloudRunV2Service, scope constructs.Construct, id *string, config *GoogleCloudRunV2ServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudRunV2Service.GoogleCloudRunV2Service",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Service)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Service)SetClient(val *string) {
	if err := j.validateSetClientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"client",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Service)SetClientVersion(val *string) {
	if err := j.validateSetClientVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Service)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Service)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Service)SetCustomAudiences(val *[]*string) {
	if err := j.validateSetCustomAudiencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customAudiences",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Service)SetDefaultUriDisabled(val interface{}) {
	if err := j.validateSetDefaultUriDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultUriDisabled",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Service)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Service)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Service)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Service)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Service)SetIngress(val *string) {
	if err := j.validateSetIngressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingress",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Service)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Service)SetLaunchStage(val *string) {
	if err := j.validateSetLaunchStageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchStage",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Service)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Service)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Service)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Service)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Service)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Service)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a GoogleCloudRunV2Service resource upon running "cdktf plan <stack-name>".
func GoogleCloudRunV2Service_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleCloudRunV2Service_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleCloudRunV2Service.GoogleCloudRunV2Service",
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
func GoogleCloudRunV2Service_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudRunV2Service_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleCloudRunV2Service.GoogleCloudRunV2Service",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCloudRunV2Service_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudRunV2Service_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleCloudRunV2Service.GoogleCloudRunV2Service",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCloudRunV2Service_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudRunV2Service_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleCloudRunV2Service.GoogleCloudRunV2Service",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleCloudRunV2Service_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleCloudRunV2Service.GoogleCloudRunV2Service",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2Service) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudRunV2Service) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2Service) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudRunV2Service) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2Service) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2Service) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2Service) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2Service) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudRunV2Service) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2Service) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2Service) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2Service) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) PutBinaryAuthorization(value *GoogleCloudRunV2ServiceBinaryAuthorization) {
	if err := g.validatePutBinaryAuthorizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBinaryAuthorization",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) PutScaling(value *GoogleCloudRunV2ServiceScaling) {
	if err := g.validatePutScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScaling",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) PutTemplate(value *GoogleCloudRunV2ServiceTemplate) {
	if err := g.validatePutTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTemplate",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) PutTimeouts(value *GoogleCloudRunV2ServiceTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) PutTraffic(value interface{}) {
	if err := g.validatePutTrafficParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTraffic",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) ResetAnnotations() {
	_jsii_.InvokeVoid(
		g,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) ResetBinaryAuthorization() {
	_jsii_.InvokeVoid(
		g,
		"resetBinaryAuthorization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) ResetClient() {
	_jsii_.InvokeVoid(
		g,
		"resetClient",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) ResetClientVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetClientVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) ResetCustomAudiences() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomAudiences",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) ResetDefaultUriDisabled() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultUriDisabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) ResetIngress() {
	_jsii_.InvokeVoid(
		g,
		"resetIngress",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) ResetLaunchStage() {
	_jsii_.InvokeVoid(
		g,
		"resetLaunchStage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) ResetScaling() {
	_jsii_.InvokeVoid(
		g,
		"resetScaling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) ResetTraffic() {
	_jsii_.InvokeVoid(
		g,
		"resetTraffic",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Service) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2Service) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2Service) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2Service) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2Service) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2Service) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


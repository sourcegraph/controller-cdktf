package googlecloudrunv2job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudrunv2job/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloud_run_v2_job google_cloud_run_v2_job}.
type GoogleCloudRunV2Job interface {
	cdktf.TerraformResource
	BinaryAuthorization() GoogleCloudRunV2JobBinaryAuthorizationOutputReference
	BinaryAuthorizationInput() *GoogleCloudRunV2JobBinaryAuthorization
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Client() *string
	SetClient(val *string)
	ClientInput() *string
	ClientVersion() *string
	SetClientVersion(val *string)
	ClientVersionInput() *string
	Conditions() GoogleCloudRunV2JobConditionsList
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
	Etag() *string
	ExecutionCount() *float64
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
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LatestCreatedExecution() GoogleCloudRunV2JobLatestCreatedExecutionList
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
	Template() GoogleCloudRunV2JobTemplateOutputReference
	TemplateInput() *GoogleCloudRunV2JobTemplate
	TerminalCondition() GoogleCloudRunV2JobTerminalConditionList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleCloudRunV2JobTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
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
	PutBinaryAuthorization(value *GoogleCloudRunV2JobBinaryAuthorization)
	PutTemplate(value *GoogleCloudRunV2JobTemplate)
	PutTimeouts(value *GoogleCloudRunV2JobTimeouts)
	ResetBinaryAuthorization()
	ResetClient()
	ResetClientVersion()
	ResetId()
	ResetLabels()
	ResetLaunchStage()
	ResetLocation()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleCloudRunV2Job
type jsiiProxy_GoogleCloudRunV2Job struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleCloudRunV2Job) BinaryAuthorization() GoogleCloudRunV2JobBinaryAuthorizationOutputReference {
	var returns GoogleCloudRunV2JobBinaryAuthorizationOutputReference
	_jsii_.Get(
		j,
		"binaryAuthorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) BinaryAuthorizationInput() *GoogleCloudRunV2JobBinaryAuthorization {
	var returns *GoogleCloudRunV2JobBinaryAuthorization
	_jsii_.Get(
		j,
		"binaryAuthorizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) Client() *string {
	var returns *string
	_jsii_.Get(
		j,
		"client",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) ClientInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) ClientVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) ClientVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) Conditions() GoogleCloudRunV2JobConditionsList {
	var returns GoogleCloudRunV2JobConditionsList
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) ExecutionCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executionCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) Generation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) LatestCreatedExecution() GoogleCloudRunV2JobLatestCreatedExecutionList {
	var returns GoogleCloudRunV2JobLatestCreatedExecutionList
	_jsii_.Get(
		j,
		"latestCreatedExecution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) LaunchStage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) LaunchStageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchStageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) ObservedGeneration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"observedGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) Reconciling() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) Template() GoogleCloudRunV2JobTemplateOutputReference {
	var returns GoogleCloudRunV2JobTemplateOutputReference
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) TemplateInput() *GoogleCloudRunV2JobTemplate {
	var returns *GoogleCloudRunV2JobTemplate
	_jsii_.Get(
		j,
		"templateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) TerminalCondition() GoogleCloudRunV2JobTerminalConditionList {
	var returns GoogleCloudRunV2JobTerminalConditionList
	_jsii_.Get(
		j,
		"terminalCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) Timeouts() GoogleCloudRunV2JobTimeoutsOutputReference {
	var returns GoogleCloudRunV2JobTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2Job) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloud_run_v2_job google_cloud_run_v2_job} Resource.
func NewGoogleCloudRunV2Job(scope constructs.Construct, id *string, config *GoogleCloudRunV2JobConfig) GoogleCloudRunV2Job {
	_init_.Initialize()

	if err := validateNewGoogleCloudRunV2JobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudRunV2Job{}

	_jsii_.Create(
		"google-beta.googleCloudRunV2Job.GoogleCloudRunV2Job",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloud_run_v2_job google_cloud_run_v2_job} Resource.
func NewGoogleCloudRunV2Job_Override(g GoogleCloudRunV2Job, scope constructs.Construct, id *string, config *GoogleCloudRunV2JobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleCloudRunV2Job.GoogleCloudRunV2Job",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Job)SetClient(val *string) {
	if err := j.validateSetClientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"client",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Job)SetClientVersion(val *string) {
	if err := j.validateSetClientVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Job)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Job)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Job)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Job)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Job)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Job)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Job)SetLaunchStage(val *string) {
	if err := j.validateSetLaunchStageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchStage",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Job)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Job)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Job)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Job)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Job)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2Job)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
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
func GoogleCloudRunV2Job_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudRunV2Job_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleCloudRunV2Job.GoogleCloudRunV2Job",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCloudRunV2Job_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudRunV2Job_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleCloudRunV2Job.GoogleCloudRunV2Job",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCloudRunV2Job_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudRunV2Job_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleCloudRunV2Job.GoogleCloudRunV2Job",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleCloudRunV2Job_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.googleCloudRunV2Job.GoogleCloudRunV2Job",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2Job) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Job) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudRunV2Job) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2Job) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudRunV2Job) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2Job) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2Job) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2Job) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2Job) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudRunV2Job) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2Job) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2Job) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Job) PutBinaryAuthorization(value *GoogleCloudRunV2JobBinaryAuthorization) {
	if err := g.validatePutBinaryAuthorizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBinaryAuthorization",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Job) PutTemplate(value *GoogleCloudRunV2JobTemplate) {
	if err := g.validatePutTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTemplate",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Job) PutTimeouts(value *GoogleCloudRunV2JobTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Job) ResetBinaryAuthorization() {
	_jsii_.InvokeVoid(
		g,
		"resetBinaryAuthorization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Job) ResetClient() {
	_jsii_.InvokeVoid(
		g,
		"resetClient",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Job) ResetClientVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetClientVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Job) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Job) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Job) ResetLaunchStage() {
	_jsii_.InvokeVoid(
		g,
		"resetLaunchStage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Job) ResetLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Job) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Job) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Job) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2Job) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2Job) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2Job) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2Job) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


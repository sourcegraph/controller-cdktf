package googleclouddeploytarget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleclouddeploytarget/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_clouddeploy_target google_clouddeploy_target}.
type GoogleClouddeployTarget interface {
	cdktf.TerraformResource
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	AnthosCluster() GoogleClouddeployTargetAnthosClusterOutputReference
	AnthosClusterInput() *GoogleClouddeployTargetAnthosCluster
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
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeployParameters() *map[string]*string
	SetDeployParameters(val *map[string]*string)
	DeployParametersInput() *map[string]*string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveAnnotations() cdktf.StringMap
	EffectiveLabels() cdktf.StringMap
	Etag() *string
	ExecutionConfigs() GoogleClouddeployTargetExecutionConfigsList
	ExecutionConfigsInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Gke() GoogleClouddeployTargetGkeOutputReference
	GkeInput() *GoogleClouddeployTargetGke
	Id() *string
	SetId(val *string)
	IdInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MultiTarget() GoogleClouddeployTargetMultiTargetOutputReference
	MultiTargetInput() *GoogleClouddeployTargetMultiTarget
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
	RequireApproval() interface{}
	SetRequireApproval(val interface{})
	RequireApprovalInput() interface{}
	Run() GoogleClouddeployTargetRunOutputReference
	RunInput() *GoogleClouddeployTargetRun
	TargetId() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleClouddeployTargetTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
	UpdateTime() *string
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
	PutAnthosCluster(value *GoogleClouddeployTargetAnthosCluster)
	PutExecutionConfigs(value interface{})
	PutGke(value *GoogleClouddeployTargetGke)
	PutMultiTarget(value *GoogleClouddeployTargetMultiTarget)
	PutRun(value *GoogleClouddeployTargetRun)
	PutTimeouts(value *GoogleClouddeployTargetTimeouts)
	ResetAnnotations()
	ResetAnthosCluster()
	ResetDeployParameters()
	ResetDescription()
	ResetExecutionConfigs()
	ResetGke()
	ResetId()
	ResetLabels()
	ResetMultiTarget()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRequireApproval()
	ResetRun()
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

// The jsii proxy struct for GoogleClouddeployTarget
type jsiiProxy_GoogleClouddeployTarget struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleClouddeployTarget) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) AnthosCluster() GoogleClouddeployTargetAnthosClusterOutputReference {
	var returns GoogleClouddeployTargetAnthosClusterOutputReference
	_jsii_.Get(
		j,
		"anthosCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) AnthosClusterInput() *GoogleClouddeployTargetAnthosCluster {
	var returns *GoogleClouddeployTargetAnthosCluster
	_jsii_.Get(
		j,
		"anthosClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) DeployParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"deployParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) DeployParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"deployParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) EffectiveAnnotations() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) ExecutionConfigs() GoogleClouddeployTargetExecutionConfigsList {
	var returns GoogleClouddeployTargetExecutionConfigsList
	_jsii_.Get(
		j,
		"executionConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) ExecutionConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executionConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) Gke() GoogleClouddeployTargetGkeOutputReference {
	var returns GoogleClouddeployTargetGkeOutputReference
	_jsii_.Get(
		j,
		"gke",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) GkeInput() *GoogleClouddeployTargetGke {
	var returns *GoogleClouddeployTargetGke
	_jsii_.Get(
		j,
		"gkeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) MultiTarget() GoogleClouddeployTargetMultiTargetOutputReference {
	var returns GoogleClouddeployTargetMultiTargetOutputReference
	_jsii_.Get(
		j,
		"multiTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) MultiTargetInput() *GoogleClouddeployTargetMultiTarget {
	var returns *GoogleClouddeployTargetMultiTarget
	_jsii_.Get(
		j,
		"multiTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) RequireApproval() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireApproval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) RequireApprovalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireApprovalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) Run() GoogleClouddeployTargetRunOutputReference {
	var returns GoogleClouddeployTargetRunOutputReference
	_jsii_.Get(
		j,
		"run",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) RunInput() *GoogleClouddeployTargetRun {
	var returns *GoogleClouddeployTargetRun
	_jsii_.Get(
		j,
		"runInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) TargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) Timeouts() GoogleClouddeployTargetTimeoutsOutputReference {
	var returns GoogleClouddeployTargetTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTarget) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_clouddeploy_target google_clouddeploy_target} Resource.
func NewGoogleClouddeployTarget(scope constructs.Construct, id *string, config *GoogleClouddeployTargetConfig) GoogleClouddeployTarget {
	_init_.Initialize()

	if err := validateNewGoogleClouddeployTargetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleClouddeployTarget{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddeployTarget.GoogleClouddeployTarget",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_clouddeploy_target google_clouddeploy_target} Resource.
func NewGoogleClouddeployTarget_Override(g GoogleClouddeployTarget, scope constructs.Construct, id *string, config *GoogleClouddeployTargetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddeployTarget.GoogleClouddeployTarget",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleClouddeployTarget)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTarget)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTarget)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTarget)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTarget)SetDeployParameters(val *map[string]*string) {
	if err := j.validateSetDeployParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deployParameters",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTarget)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTarget)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTarget)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTarget)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTarget)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTarget)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTarget)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTarget)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTarget)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTarget)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTarget)SetRequireApproval(val interface{}) {
	if err := j.validateSetRequireApprovalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireApproval",
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
func GoogleClouddeployTarget_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleClouddeployTarget_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleClouddeployTarget.GoogleClouddeployTarget",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleClouddeployTarget_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleClouddeployTarget_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleClouddeployTarget.GoogleClouddeployTarget",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleClouddeployTarget_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleClouddeployTarget_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleClouddeployTarget.GoogleClouddeployTarget",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleClouddeployTarget_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleClouddeployTarget.GoogleClouddeployTarget",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleClouddeployTarget) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleClouddeployTarget) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployTarget) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleClouddeployTarget) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleClouddeployTarget) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleClouddeployTarget) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployTarget) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployTarget) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleClouddeployTarget) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleClouddeployTarget) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployTarget) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) PutAnthosCluster(value *GoogleClouddeployTargetAnthosCluster) {
	if err := g.validatePutAnthosClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAnthosCluster",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) PutExecutionConfigs(value interface{}) {
	if err := g.validatePutExecutionConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExecutionConfigs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) PutGke(value *GoogleClouddeployTargetGke) {
	if err := g.validatePutGkeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGke",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) PutMultiTarget(value *GoogleClouddeployTargetMultiTarget) {
	if err := g.validatePutMultiTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMultiTarget",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) PutRun(value *GoogleClouddeployTargetRun) {
	if err := g.validatePutRunParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRun",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) PutTimeouts(value *GoogleClouddeployTargetTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) ResetAnnotations() {
	_jsii_.InvokeVoid(
		g,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) ResetAnthosCluster() {
	_jsii_.InvokeVoid(
		g,
		"resetAnthosCluster",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) ResetDeployParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetDeployParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) ResetExecutionConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetExecutionConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) ResetGke() {
	_jsii_.InvokeVoid(
		g,
		"resetGke",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) ResetMultiTarget() {
	_jsii_.InvokeVoid(
		g,
		"resetMultiTarget",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) ResetRequireApproval() {
	_jsii_.InvokeVoid(
		g,
		"resetRequireApproval",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) ResetRun() {
	_jsii_.InvokeVoid(
		g,
		"resetRun",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployTarget) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployTarget) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployTarget) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployTarget) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


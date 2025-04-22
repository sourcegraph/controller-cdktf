package googledataprocbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataprocbatch/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_batch google_dataproc_batch}.
type GoogleDataprocBatch interface {
	cdktf.TerraformResource
	BatchId() *string
	SetBatchId(val *string)
	BatchIdInput() *string
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
	Creator() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EffectiveLabels() cdktf.StringMap
	EnvironmentConfig() GoogleDataprocBatchEnvironmentConfigOutputReference
	EnvironmentConfigInput() *GoogleDataprocBatchEnvironmentConfig
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
	Name() *string
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
	PysparkBatch() GoogleDataprocBatchPysparkBatchOutputReference
	PysparkBatchInput() *GoogleDataprocBatchPysparkBatch
	// Experimental.
	RawOverrides() interface{}
	RuntimeConfig() GoogleDataprocBatchRuntimeConfigOutputReference
	RuntimeConfigInput() *GoogleDataprocBatchRuntimeConfig
	RuntimeInfo() GoogleDataprocBatchRuntimeInfoList
	SparkBatch() GoogleDataprocBatchSparkBatchOutputReference
	SparkBatchInput() *GoogleDataprocBatchSparkBatch
	SparkRBatch() GoogleDataprocBatchSparkRBatchOutputReference
	SparkRBatchInput() *GoogleDataprocBatchSparkRBatch
	SparkSqlBatch() GoogleDataprocBatchSparkSqlBatchOutputReference
	SparkSqlBatchInput() *GoogleDataprocBatchSparkSqlBatch
	State() *string
	StateHistory() GoogleDataprocBatchStateHistoryList
	StateMessage() *string
	StateTime() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleDataprocBatchTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uuid() *string
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
	PutEnvironmentConfig(value *GoogleDataprocBatchEnvironmentConfig)
	PutPysparkBatch(value *GoogleDataprocBatchPysparkBatch)
	PutRuntimeConfig(value *GoogleDataprocBatchRuntimeConfig)
	PutSparkBatch(value *GoogleDataprocBatchSparkBatch)
	PutSparkRBatch(value *GoogleDataprocBatchSparkRBatch)
	PutSparkSqlBatch(value *GoogleDataprocBatchSparkSqlBatch)
	PutTimeouts(value *GoogleDataprocBatchTimeouts)
	ResetBatchId()
	ResetEnvironmentConfig()
	ResetId()
	ResetLabels()
	ResetLocation()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetPysparkBatch()
	ResetRuntimeConfig()
	ResetSparkBatch()
	ResetSparkRBatch()
	ResetSparkSqlBatch()
	ResetTimeouts()
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

// The jsii proxy struct for GoogleDataprocBatch
type jsiiProxy_GoogleDataprocBatch struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleDataprocBatch) BatchId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) BatchIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) Creator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) EnvironmentConfig() GoogleDataprocBatchEnvironmentConfigOutputReference {
	var returns GoogleDataprocBatchEnvironmentConfigOutputReference
	_jsii_.Get(
		j,
		"environmentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) EnvironmentConfigInput() *GoogleDataprocBatchEnvironmentConfig {
	var returns *GoogleDataprocBatchEnvironmentConfig
	_jsii_.Get(
		j,
		"environmentConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) Operation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) PysparkBatch() GoogleDataprocBatchPysparkBatchOutputReference {
	var returns GoogleDataprocBatchPysparkBatchOutputReference
	_jsii_.Get(
		j,
		"pysparkBatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) PysparkBatchInput() *GoogleDataprocBatchPysparkBatch {
	var returns *GoogleDataprocBatchPysparkBatch
	_jsii_.Get(
		j,
		"pysparkBatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) RuntimeConfig() GoogleDataprocBatchRuntimeConfigOutputReference {
	var returns GoogleDataprocBatchRuntimeConfigOutputReference
	_jsii_.Get(
		j,
		"runtimeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) RuntimeConfigInput() *GoogleDataprocBatchRuntimeConfig {
	var returns *GoogleDataprocBatchRuntimeConfig
	_jsii_.Get(
		j,
		"runtimeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) RuntimeInfo() GoogleDataprocBatchRuntimeInfoList {
	var returns GoogleDataprocBatchRuntimeInfoList
	_jsii_.Get(
		j,
		"runtimeInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) SparkBatch() GoogleDataprocBatchSparkBatchOutputReference {
	var returns GoogleDataprocBatchSparkBatchOutputReference
	_jsii_.Get(
		j,
		"sparkBatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) SparkBatchInput() *GoogleDataprocBatchSparkBatch {
	var returns *GoogleDataprocBatchSparkBatch
	_jsii_.Get(
		j,
		"sparkBatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) SparkRBatch() GoogleDataprocBatchSparkRBatchOutputReference {
	var returns GoogleDataprocBatchSparkRBatchOutputReference
	_jsii_.Get(
		j,
		"sparkRBatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) SparkRBatchInput() *GoogleDataprocBatchSparkRBatch {
	var returns *GoogleDataprocBatchSparkRBatch
	_jsii_.Get(
		j,
		"sparkRBatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) SparkSqlBatch() GoogleDataprocBatchSparkSqlBatchOutputReference {
	var returns GoogleDataprocBatchSparkSqlBatchOutputReference
	_jsii_.Get(
		j,
		"sparkSqlBatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) SparkSqlBatchInput() *GoogleDataprocBatchSparkSqlBatch {
	var returns *GoogleDataprocBatchSparkSqlBatch
	_jsii_.Get(
		j,
		"sparkSqlBatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) StateHistory() GoogleDataprocBatchStateHistoryList {
	var returns GoogleDataprocBatchStateHistoryList
	_jsii_.Get(
		j,
		"stateHistory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) StateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) StateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) Timeouts() GoogleDataprocBatchTimeoutsOutputReference {
	var returns GoogleDataprocBatchTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocBatch) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_batch google_dataproc_batch} Resource.
func NewGoogleDataprocBatch(scope constructs.Construct, id *string, config *GoogleDataprocBatchConfig) GoogleDataprocBatch {
	_init_.Initialize()

	if err := validateNewGoogleDataprocBatchParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocBatch{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocBatch.GoogleDataprocBatch",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_batch google_dataproc_batch} Resource.
func NewGoogleDataprocBatch_Override(g GoogleDataprocBatch, scope constructs.Construct, id *string, config *GoogleDataprocBatchConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocBatch.GoogleDataprocBatch",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocBatch)SetBatchId(val *string) {
	if err := j.validateSetBatchIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchId",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocBatch)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocBatch)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocBatch)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocBatch)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocBatch)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocBatch)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocBatch)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocBatch)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocBatch)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocBatch)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocBatch)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a GoogleDataprocBatch resource upon running "cdktf plan <stack-name>".
func GoogleDataprocBatch_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleDataprocBatch_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDataprocBatch.GoogleDataprocBatch",
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
func GoogleDataprocBatch_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDataprocBatch_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDataprocBatch.GoogleDataprocBatch",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDataprocBatch_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDataprocBatch_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDataprocBatch.GoogleDataprocBatch",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDataprocBatch_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDataprocBatch_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDataprocBatch.GoogleDataprocBatch",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleDataprocBatch_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleDataprocBatch.GoogleDataprocBatch",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleDataprocBatch) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocBatch) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocBatch) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocBatch) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocBatch) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocBatch) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocBatch) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocBatch) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocBatch) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocBatch) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocBatch) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocBatch) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) PutEnvironmentConfig(value *GoogleDataprocBatchEnvironmentConfig) {
	if err := g.validatePutEnvironmentConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEnvironmentConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) PutPysparkBatch(value *GoogleDataprocBatchPysparkBatch) {
	if err := g.validatePutPysparkBatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPysparkBatch",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) PutRuntimeConfig(value *GoogleDataprocBatchRuntimeConfig) {
	if err := g.validatePutRuntimeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRuntimeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) PutSparkBatch(value *GoogleDataprocBatchSparkBatch) {
	if err := g.validatePutSparkBatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSparkBatch",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) PutSparkRBatch(value *GoogleDataprocBatchSparkRBatch) {
	if err := g.validatePutSparkRBatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSparkRBatch",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) PutSparkSqlBatch(value *GoogleDataprocBatchSparkSqlBatch) {
	if err := g.validatePutSparkSqlBatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSparkSqlBatch",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) PutTimeouts(value *GoogleDataprocBatchTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) ResetBatchId() {
	_jsii_.InvokeVoid(
		g,
		"resetBatchId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) ResetEnvironmentConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEnvironmentConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) ResetLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) ResetPysparkBatch() {
	_jsii_.InvokeVoid(
		g,
		"resetPysparkBatch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) ResetRuntimeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetRuntimeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) ResetSparkBatch() {
	_jsii_.InvokeVoid(
		g,
		"resetSparkBatch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) ResetSparkRBatch() {
	_jsii_.InvokeVoid(
		g,
		"resetSparkRBatch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) ResetSparkSqlBatch() {
	_jsii_.InvokeVoid(
		g,
		"resetSparkSqlBatch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocBatch) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocBatch) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocBatch) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocBatch) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocBatch) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocBatch) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


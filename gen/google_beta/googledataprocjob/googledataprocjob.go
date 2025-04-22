package googledataprocjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataprocjob/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_job google_dataproc_job}.
type GoogleDataprocJob interface {
	cdktf.TerraformResource
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
	DriverControlsFilesUri() *string
	DriverOutputResourceUri() *string
	EffectiveLabels() cdktf.StringMap
	ForceDelete() interface{}
	SetForceDelete(val interface{})
	ForceDeleteInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HadoopConfig() GoogleDataprocJobHadoopConfigOutputReference
	HadoopConfigInput() *GoogleDataprocJobHadoopConfig
	HiveConfig() GoogleDataprocJobHiveConfigOutputReference
	HiveConfigInput() *GoogleDataprocJobHiveConfig
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
	// The tree node.
	Node() constructs.Node
	PigConfig() GoogleDataprocJobPigConfigOutputReference
	PigConfigInput() *GoogleDataprocJobPigConfig
	Placement() GoogleDataprocJobPlacementOutputReference
	PlacementInput() *GoogleDataprocJobPlacement
	PrestoConfig() GoogleDataprocJobPrestoConfigOutputReference
	PrestoConfigInput() *GoogleDataprocJobPrestoConfig
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
	PysparkConfig() GoogleDataprocJobPysparkConfigOutputReference
	PysparkConfigInput() *GoogleDataprocJobPysparkConfig
	// Experimental.
	RawOverrides() interface{}
	Reference() GoogleDataprocJobReferenceOutputReference
	ReferenceInput() *GoogleDataprocJobReference
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	Scheduling() GoogleDataprocJobSchedulingOutputReference
	SchedulingInput() *GoogleDataprocJobScheduling
	SparkConfig() GoogleDataprocJobSparkConfigOutputReference
	SparkConfigInput() *GoogleDataprocJobSparkConfig
	SparksqlConfig() GoogleDataprocJobSparksqlConfigOutputReference
	SparksqlConfigInput() *GoogleDataprocJobSparksqlConfig
	Status() GoogleDataprocJobStatusList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleDataprocJobTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutHadoopConfig(value *GoogleDataprocJobHadoopConfig)
	PutHiveConfig(value *GoogleDataprocJobHiveConfig)
	PutPigConfig(value *GoogleDataprocJobPigConfig)
	PutPlacement(value *GoogleDataprocJobPlacement)
	PutPrestoConfig(value *GoogleDataprocJobPrestoConfig)
	PutPysparkConfig(value *GoogleDataprocJobPysparkConfig)
	PutReference(value *GoogleDataprocJobReference)
	PutScheduling(value *GoogleDataprocJobScheduling)
	PutSparkConfig(value *GoogleDataprocJobSparkConfig)
	PutSparksqlConfig(value *GoogleDataprocJobSparksqlConfig)
	PutTimeouts(value *GoogleDataprocJobTimeouts)
	ResetForceDelete()
	ResetHadoopConfig()
	ResetHiveConfig()
	ResetId()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPigConfig()
	ResetPrestoConfig()
	ResetProject()
	ResetPysparkConfig()
	ResetReference()
	ResetRegion()
	ResetScheduling()
	ResetSparkConfig()
	ResetSparksqlConfig()
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

// The jsii proxy struct for GoogleDataprocJob
type jsiiProxy_GoogleDataprocJob struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleDataprocJob) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) DriverControlsFilesUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverControlsFilesUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) DriverOutputResourceUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverOutputResourceUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) ForceDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) ForceDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) HadoopConfig() GoogleDataprocJobHadoopConfigOutputReference {
	var returns GoogleDataprocJobHadoopConfigOutputReference
	_jsii_.Get(
		j,
		"hadoopConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) HadoopConfigInput() *GoogleDataprocJobHadoopConfig {
	var returns *GoogleDataprocJobHadoopConfig
	_jsii_.Get(
		j,
		"hadoopConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) HiveConfig() GoogleDataprocJobHiveConfigOutputReference {
	var returns GoogleDataprocJobHiveConfigOutputReference
	_jsii_.Get(
		j,
		"hiveConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) HiveConfigInput() *GoogleDataprocJobHiveConfig {
	var returns *GoogleDataprocJobHiveConfig
	_jsii_.Get(
		j,
		"hiveConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) PigConfig() GoogleDataprocJobPigConfigOutputReference {
	var returns GoogleDataprocJobPigConfigOutputReference
	_jsii_.Get(
		j,
		"pigConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) PigConfigInput() *GoogleDataprocJobPigConfig {
	var returns *GoogleDataprocJobPigConfig
	_jsii_.Get(
		j,
		"pigConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) Placement() GoogleDataprocJobPlacementOutputReference {
	var returns GoogleDataprocJobPlacementOutputReference
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) PlacementInput() *GoogleDataprocJobPlacement {
	var returns *GoogleDataprocJobPlacement
	_jsii_.Get(
		j,
		"placementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) PrestoConfig() GoogleDataprocJobPrestoConfigOutputReference {
	var returns GoogleDataprocJobPrestoConfigOutputReference
	_jsii_.Get(
		j,
		"prestoConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) PrestoConfigInput() *GoogleDataprocJobPrestoConfig {
	var returns *GoogleDataprocJobPrestoConfig
	_jsii_.Get(
		j,
		"prestoConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) PysparkConfig() GoogleDataprocJobPysparkConfigOutputReference {
	var returns GoogleDataprocJobPysparkConfigOutputReference
	_jsii_.Get(
		j,
		"pysparkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) PysparkConfigInput() *GoogleDataprocJobPysparkConfig {
	var returns *GoogleDataprocJobPysparkConfig
	_jsii_.Get(
		j,
		"pysparkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) Reference() GoogleDataprocJobReferenceOutputReference {
	var returns GoogleDataprocJobReferenceOutputReference
	_jsii_.Get(
		j,
		"reference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) ReferenceInput() *GoogleDataprocJobReference {
	var returns *GoogleDataprocJobReference
	_jsii_.Get(
		j,
		"referenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) Scheduling() GoogleDataprocJobSchedulingOutputReference {
	var returns GoogleDataprocJobSchedulingOutputReference
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) SchedulingInput() *GoogleDataprocJobScheduling {
	var returns *GoogleDataprocJobScheduling
	_jsii_.Get(
		j,
		"schedulingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) SparkConfig() GoogleDataprocJobSparkConfigOutputReference {
	var returns GoogleDataprocJobSparkConfigOutputReference
	_jsii_.Get(
		j,
		"sparkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) SparkConfigInput() *GoogleDataprocJobSparkConfig {
	var returns *GoogleDataprocJobSparkConfig
	_jsii_.Get(
		j,
		"sparkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) SparksqlConfig() GoogleDataprocJobSparksqlConfigOutputReference {
	var returns GoogleDataprocJobSparksqlConfigOutputReference
	_jsii_.Get(
		j,
		"sparksqlConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) SparksqlConfigInput() *GoogleDataprocJobSparksqlConfig {
	var returns *GoogleDataprocJobSparksqlConfig
	_jsii_.Get(
		j,
		"sparksqlConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) Status() GoogleDataprocJobStatusList {
	var returns GoogleDataprocJobStatusList
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) Timeouts() GoogleDataprocJobTimeoutsOutputReference {
	var returns GoogleDataprocJobTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJob) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_job google_dataproc_job} Resource.
func NewGoogleDataprocJob(scope constructs.Construct, id *string, config *GoogleDataprocJobConfig) GoogleDataprocJob {
	_init_.Initialize()

	if err := validateNewGoogleDataprocJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocJob{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocJob.GoogleDataprocJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_job google_dataproc_job} Resource.
func NewGoogleDataprocJob_Override(g GoogleDataprocJob, scope constructs.Construct, id *string, config *GoogleDataprocJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocJob.GoogleDataprocJob",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJob)SetForceDelete(val interface{}) {
	if err := j.validateSetForceDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDelete",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJob)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJob)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJob)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJob)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJob)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJob)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJob)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTF code for importing a GoogleDataprocJob resource upon running "cdktf plan <stack-name>".
func GoogleDataprocJob_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleDataprocJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDataprocJob.GoogleDataprocJob",
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
func GoogleDataprocJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDataprocJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDataprocJob.GoogleDataprocJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDataprocJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDataprocJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDataprocJob.GoogleDataprocJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDataprocJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDataprocJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDataprocJob.GoogleDataprocJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleDataprocJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleDataprocJob.GoogleDataprocJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleDataprocJob) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleDataprocJob) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleDataprocJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocJob) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocJob) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocJob) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocJob) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocJob) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocJob) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleDataprocJob) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocJob) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDataprocJob) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleDataprocJob) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDataprocJob) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleDataprocJob) PutHadoopConfig(value *GoogleDataprocJobHadoopConfig) {
	if err := g.validatePutHadoopConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHadoopConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocJob) PutHiveConfig(value *GoogleDataprocJobHiveConfig) {
	if err := g.validatePutHiveConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHiveConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocJob) PutPigConfig(value *GoogleDataprocJobPigConfig) {
	if err := g.validatePutPigConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPigConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocJob) PutPlacement(value *GoogleDataprocJobPlacement) {
	if err := g.validatePutPlacementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPlacement",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocJob) PutPrestoConfig(value *GoogleDataprocJobPrestoConfig) {
	if err := g.validatePutPrestoConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPrestoConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocJob) PutPysparkConfig(value *GoogleDataprocJobPysparkConfig) {
	if err := g.validatePutPysparkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPysparkConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocJob) PutReference(value *GoogleDataprocJobReference) {
	if err := g.validatePutReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putReference",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocJob) PutScheduling(value *GoogleDataprocJobScheduling) {
	if err := g.validatePutSchedulingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScheduling",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocJob) PutSparkConfig(value *GoogleDataprocJobSparkConfig) {
	if err := g.validatePutSparkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSparkConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocJob) PutSparksqlConfig(value *GoogleDataprocJobSparksqlConfig) {
	if err := g.validatePutSparksqlConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSparksqlConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocJob) PutTimeouts(value *GoogleDataprocJobTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocJob) ResetForceDelete() {
	_jsii_.InvokeVoid(
		g,
		"resetForceDelete",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJob) ResetHadoopConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetHadoopConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJob) ResetHiveConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetHiveConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJob) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJob) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJob) ResetPigConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPigConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJob) ResetPrestoConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPrestoConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJob) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJob) ResetPysparkConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPysparkConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJob) ResetReference() {
	_jsii_.InvokeVoid(
		g,
		"resetReference",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJob) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJob) ResetScheduling() {
	_jsii_.InvokeVoid(
		g,
		"resetScheduling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJob) ResetSparkConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSparkConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJob) ResetSparksqlConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSparksqlConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJob) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocJob) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocJob) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


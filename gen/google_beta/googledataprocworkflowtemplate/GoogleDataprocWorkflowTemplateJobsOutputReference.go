package googledataprocworkflowtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataprocworkflowtemplate/internal"
)

type GoogleDataprocWorkflowTemplateJobsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	HadoopJob() GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference
	HadoopJobInput() *GoogleDataprocWorkflowTemplateJobsHadoopJob
	HiveJob() GoogleDataprocWorkflowTemplateJobsHiveJobOutputReference
	HiveJobInput() *GoogleDataprocWorkflowTemplateJobsHiveJob
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	PigJob() GoogleDataprocWorkflowTemplateJobsPigJobOutputReference
	PigJobInput() *GoogleDataprocWorkflowTemplateJobsPigJob
	PrerequisiteStepIds() *[]*string
	SetPrerequisiteStepIds(val *[]*string)
	PrerequisiteStepIdsInput() *[]*string
	PrestoJob() GoogleDataprocWorkflowTemplateJobsPrestoJobOutputReference
	PrestoJobInput() *GoogleDataprocWorkflowTemplateJobsPrestoJob
	PysparkJob() GoogleDataprocWorkflowTemplateJobsPysparkJobOutputReference
	PysparkJobInput() *GoogleDataprocWorkflowTemplateJobsPysparkJob
	Scheduling() GoogleDataprocWorkflowTemplateJobsSchedulingOutputReference
	SchedulingInput() *GoogleDataprocWorkflowTemplateJobsScheduling
	SparkJob() GoogleDataprocWorkflowTemplateJobsSparkJobOutputReference
	SparkJobInput() *GoogleDataprocWorkflowTemplateJobsSparkJob
	SparkRJob() GoogleDataprocWorkflowTemplateJobsSparkRJobOutputReference
	SparkRJobInput() *GoogleDataprocWorkflowTemplateJobsSparkRJob
	SparkSqlJob() GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference
	SparkSqlJobInput() *GoogleDataprocWorkflowTemplateJobsSparkSqlJob
	StepId() *string
	SetStepId(val *string)
	StepIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutHadoopJob(value *GoogleDataprocWorkflowTemplateJobsHadoopJob)
	PutHiveJob(value *GoogleDataprocWorkflowTemplateJobsHiveJob)
	PutPigJob(value *GoogleDataprocWorkflowTemplateJobsPigJob)
	PutPrestoJob(value *GoogleDataprocWorkflowTemplateJobsPrestoJob)
	PutPysparkJob(value *GoogleDataprocWorkflowTemplateJobsPysparkJob)
	PutScheduling(value *GoogleDataprocWorkflowTemplateJobsScheduling)
	PutSparkJob(value *GoogleDataprocWorkflowTemplateJobsSparkJob)
	PutSparkRJob(value *GoogleDataprocWorkflowTemplateJobsSparkRJob)
	PutSparkSqlJob(value *GoogleDataprocWorkflowTemplateJobsSparkSqlJob)
	ResetHadoopJob()
	ResetHiveJob()
	ResetLabels()
	ResetPigJob()
	ResetPrerequisiteStepIds()
	ResetPrestoJob()
	ResetPysparkJob()
	ResetScheduling()
	ResetSparkJob()
	ResetSparkRJob()
	ResetSparkSqlJob()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocWorkflowTemplateJobsOutputReference
type jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) HadoopJob() GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference {
	var returns GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference
	_jsii_.Get(
		j,
		"hadoopJob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) HadoopJobInput() *GoogleDataprocWorkflowTemplateJobsHadoopJob {
	var returns *GoogleDataprocWorkflowTemplateJobsHadoopJob
	_jsii_.Get(
		j,
		"hadoopJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) HiveJob() GoogleDataprocWorkflowTemplateJobsHiveJobOutputReference {
	var returns GoogleDataprocWorkflowTemplateJobsHiveJobOutputReference
	_jsii_.Get(
		j,
		"hiveJob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) HiveJobInput() *GoogleDataprocWorkflowTemplateJobsHiveJob {
	var returns *GoogleDataprocWorkflowTemplateJobsHiveJob
	_jsii_.Get(
		j,
		"hiveJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) PigJob() GoogleDataprocWorkflowTemplateJobsPigJobOutputReference {
	var returns GoogleDataprocWorkflowTemplateJobsPigJobOutputReference
	_jsii_.Get(
		j,
		"pigJob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) PigJobInput() *GoogleDataprocWorkflowTemplateJobsPigJob {
	var returns *GoogleDataprocWorkflowTemplateJobsPigJob
	_jsii_.Get(
		j,
		"pigJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) PrerequisiteStepIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"prerequisiteStepIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) PrerequisiteStepIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"prerequisiteStepIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) PrestoJob() GoogleDataprocWorkflowTemplateJobsPrestoJobOutputReference {
	var returns GoogleDataprocWorkflowTemplateJobsPrestoJobOutputReference
	_jsii_.Get(
		j,
		"prestoJob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) PrestoJobInput() *GoogleDataprocWorkflowTemplateJobsPrestoJob {
	var returns *GoogleDataprocWorkflowTemplateJobsPrestoJob
	_jsii_.Get(
		j,
		"prestoJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) PysparkJob() GoogleDataprocWorkflowTemplateJobsPysparkJobOutputReference {
	var returns GoogleDataprocWorkflowTemplateJobsPysparkJobOutputReference
	_jsii_.Get(
		j,
		"pysparkJob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) PysparkJobInput() *GoogleDataprocWorkflowTemplateJobsPysparkJob {
	var returns *GoogleDataprocWorkflowTemplateJobsPysparkJob
	_jsii_.Get(
		j,
		"pysparkJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) Scheduling() GoogleDataprocWorkflowTemplateJobsSchedulingOutputReference {
	var returns GoogleDataprocWorkflowTemplateJobsSchedulingOutputReference
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) SchedulingInput() *GoogleDataprocWorkflowTemplateJobsScheduling {
	var returns *GoogleDataprocWorkflowTemplateJobsScheduling
	_jsii_.Get(
		j,
		"schedulingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) SparkJob() GoogleDataprocWorkflowTemplateJobsSparkJobOutputReference {
	var returns GoogleDataprocWorkflowTemplateJobsSparkJobOutputReference
	_jsii_.Get(
		j,
		"sparkJob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) SparkJobInput() *GoogleDataprocWorkflowTemplateJobsSparkJob {
	var returns *GoogleDataprocWorkflowTemplateJobsSparkJob
	_jsii_.Get(
		j,
		"sparkJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) SparkRJob() GoogleDataprocWorkflowTemplateJobsSparkRJobOutputReference {
	var returns GoogleDataprocWorkflowTemplateJobsSparkRJobOutputReference
	_jsii_.Get(
		j,
		"sparkRJob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) SparkRJobInput() *GoogleDataprocWorkflowTemplateJobsSparkRJob {
	var returns *GoogleDataprocWorkflowTemplateJobsSparkRJob
	_jsii_.Get(
		j,
		"sparkRJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) SparkSqlJob() GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference {
	var returns GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference
	_jsii_.Get(
		j,
		"sparkSqlJob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) SparkSqlJobInput() *GoogleDataprocWorkflowTemplateJobsSparkSqlJob {
	var returns *GoogleDataprocWorkflowTemplateJobsSparkSqlJob
	_jsii_.Get(
		j,
		"sparkSqlJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) StepId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stepId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) StepIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stepIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataprocWorkflowTemplateJobsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleDataprocWorkflowTemplateJobsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocWorkflowTemplateJobsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleDataprocWorkflowTemplateJobsOutputReference_Override(g GoogleDataprocWorkflowTemplateJobsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference)SetPrerequisiteStepIds(val *[]*string) {
	if err := j.validateSetPrerequisiteStepIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prerequisiteStepIds",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference)SetStepId(val *string) {
	if err := j.validateSetStepIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stepId",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) PutHadoopJob(value *GoogleDataprocWorkflowTemplateJobsHadoopJob) {
	if err := g.validatePutHadoopJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHadoopJob",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) PutHiveJob(value *GoogleDataprocWorkflowTemplateJobsHiveJob) {
	if err := g.validatePutHiveJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHiveJob",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) PutPigJob(value *GoogleDataprocWorkflowTemplateJobsPigJob) {
	if err := g.validatePutPigJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPigJob",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) PutPrestoJob(value *GoogleDataprocWorkflowTemplateJobsPrestoJob) {
	if err := g.validatePutPrestoJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPrestoJob",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) PutPysparkJob(value *GoogleDataprocWorkflowTemplateJobsPysparkJob) {
	if err := g.validatePutPysparkJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPysparkJob",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) PutScheduling(value *GoogleDataprocWorkflowTemplateJobsScheduling) {
	if err := g.validatePutSchedulingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScheduling",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) PutSparkJob(value *GoogleDataprocWorkflowTemplateJobsSparkJob) {
	if err := g.validatePutSparkJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSparkJob",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) PutSparkRJob(value *GoogleDataprocWorkflowTemplateJobsSparkRJob) {
	if err := g.validatePutSparkRJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSparkRJob",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) PutSparkSqlJob(value *GoogleDataprocWorkflowTemplateJobsSparkSqlJob) {
	if err := g.validatePutSparkSqlJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSparkSqlJob",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) ResetHadoopJob() {
	_jsii_.InvokeVoid(
		g,
		"resetHadoopJob",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) ResetHiveJob() {
	_jsii_.InvokeVoid(
		g,
		"resetHiveJob",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) ResetPigJob() {
	_jsii_.InvokeVoid(
		g,
		"resetPigJob",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) ResetPrerequisiteStepIds() {
	_jsii_.InvokeVoid(
		g,
		"resetPrerequisiteStepIds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) ResetPrestoJob() {
	_jsii_.InvokeVoid(
		g,
		"resetPrestoJob",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) ResetPysparkJob() {
	_jsii_.InvokeVoid(
		g,
		"resetPysparkJob",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) ResetScheduling() {
	_jsii_.InvokeVoid(
		g,
		"resetScheduling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) ResetSparkJob() {
	_jsii_.InvokeVoid(
		g,
		"resetSparkJob",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) ResetSparkRJob() {
	_jsii_.InvokeVoid(
		g,
		"resetSparkRJob",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) ResetSparkSqlJob() {
	_jsii_.InvokeVoid(
		g,
		"resetSparkSqlJob",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package googledataprocworkflowtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataprocworkflowtemplate/internal"
)

type GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference interface {
	cdktf.ComplexObject
	ArchiveUris() *[]*string
	SetArchiveUris(val *[]*string)
	ArchiveUrisInput() *[]*string
	Args() *[]*string
	SetArgs(val *[]*string)
	ArgsInput() *[]*string
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
	FileUris() *[]*string
	SetFileUris(val *[]*string)
	FileUrisInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataprocWorkflowTemplateJobsHadoopJob
	SetInternalValue(val *GoogleDataprocWorkflowTemplateJobsHadoopJob)
	JarFileUris() *[]*string
	SetJarFileUris(val *[]*string)
	JarFileUrisInput() *[]*string
	LoggingConfig() GoogleDataprocWorkflowTemplateJobsHadoopJobLoggingConfigOutputReference
	LoggingConfigInput() *GoogleDataprocWorkflowTemplateJobsHadoopJobLoggingConfig
	MainClass() *string
	SetMainClass(val *string)
	MainClassInput() *string
	MainJarFileUri() *string
	SetMainJarFileUri(val *string)
	MainJarFileUriInput() *string
	Properties() *map[string]*string
	SetProperties(val *map[string]*string)
	PropertiesInput() *map[string]*string
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
	PutLoggingConfig(value *GoogleDataprocWorkflowTemplateJobsHadoopJobLoggingConfig)
	ResetArchiveUris()
	ResetArgs()
	ResetFileUris()
	ResetJarFileUris()
	ResetLoggingConfig()
	ResetMainClass()
	ResetMainJarFileUri()
	ResetProperties()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference
type jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) ArchiveUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"archiveUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) ArchiveUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"archiveUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) FileUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fileUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) FileUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fileUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) InternalValue() *GoogleDataprocWorkflowTemplateJobsHadoopJob {
	var returns *GoogleDataprocWorkflowTemplateJobsHadoopJob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) JarFileUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jarFileUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) JarFileUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jarFileUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) LoggingConfig() GoogleDataprocWorkflowTemplateJobsHadoopJobLoggingConfigOutputReference {
	var returns GoogleDataprocWorkflowTemplateJobsHadoopJobLoggingConfigOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) LoggingConfigInput() *GoogleDataprocWorkflowTemplateJobsHadoopJobLoggingConfig {
	var returns *GoogleDataprocWorkflowTemplateJobsHadoopJobLoggingConfig
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) MainClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) MainClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) MainJarFileUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainJarFileUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) MainJarFileUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainJarFileUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) Properties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) PropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocWorkflowTemplateJobsHadoopJobOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference_Override(g GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference)SetArchiveUris(val *[]*string) {
	if err := j.validateSetArchiveUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archiveUris",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference)SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference)SetFileUris(val *[]*string) {
	if err := j.validateSetFileUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileUris",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference)SetInternalValue(val *GoogleDataprocWorkflowTemplateJobsHadoopJob) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference)SetJarFileUris(val *[]*string) {
	if err := j.validateSetJarFileUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jarFileUris",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference)SetMainClass(val *string) {
	if err := j.validateSetMainClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mainClass",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference)SetMainJarFileUri(val *string) {
	if err := j.validateSetMainJarFileUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mainJarFileUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference)SetProperties(val *map[string]*string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) PutLoggingConfig(value *GoogleDataprocWorkflowTemplateJobsHadoopJobLoggingConfig) {
	if err := g.validatePutLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) ResetArchiveUris() {
	_jsii_.InvokeVoid(
		g,
		"resetArchiveUris",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		g,
		"resetArgs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) ResetFileUris() {
	_jsii_.InvokeVoid(
		g,
		"resetFileUris",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) ResetJarFileUris() {
	_jsii_.InvokeVoid(
		g,
		"resetJarFileUris",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) ResetMainClass() {
	_jsii_.InvokeVoid(
		g,
		"resetMainClass",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) ResetMainJarFileUri() {
	_jsii_.InvokeVoid(
		g,
		"resetMainJarFileUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


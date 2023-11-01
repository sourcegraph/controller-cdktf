package googledatapipelinepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatapipelinepipeline/internal"
)

type GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference interface {
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
	ContainerSpecGcsPath() *string
	SetContainerSpecGcsPath(val *string)
	ContainerSpecGcsPathInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Environment() GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentOutputReference
	EnvironmentInput() *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironment
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameter
	SetInternalValue(val *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameter)
	JobName() *string
	SetJobName(val *string)
	JobNameInput() *string
	LaunchOptions() *map[string]*string
	SetLaunchOptions(val *map[string]*string)
	LaunchOptionsInput() *map[string]*string
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransformNameMappings() *map[string]*string
	SetTransformNameMappings(val *map[string]*string)
	TransformNameMappingsInput() *map[string]*string
	Update() interface{}
	SetUpdate(val interface{})
	UpdateInput() interface{}
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
	PutEnvironment(value *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironment)
	ResetContainerSpecGcsPath()
	ResetEnvironment()
	ResetLaunchOptions()
	ResetParameters()
	ResetTransformNameMappings()
	ResetUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference
type jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ContainerSpecGcsPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerSpecGcsPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ContainerSpecGcsPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerSpecGcsPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) Environment() GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentOutputReference {
	var returns GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentOutputReference
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) EnvironmentInput() *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironment {
	var returns *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironment
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) InternalValue() *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameter {
	var returns *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) JobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) JobNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) LaunchOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"launchOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) LaunchOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"launchOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) TransformNameMappings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"transformNameMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) TransformNameMappingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"transformNameMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) Update() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) UpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewGoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataPipelinePipeline.GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference_Override(g GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataPipelinePipeline.GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetContainerSpecGcsPath(val *string) {
	if err := j.validateSetContainerSpecGcsPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerSpecGcsPath",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetInternalValue(val *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetJobName(val *string) {
	if err := j.validateSetJobNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobName",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetLaunchOptions(val *map[string]*string) {
	if err := j.validateSetLaunchOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchOptions",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetTransformNameMappings(val *map[string]*string) {
	if err := j.validateSetTransformNameMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transformNameMappings",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetUpdate(val interface{}) {
	if err := j.validateSetUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) PutEnvironment(value *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironment) {
	if err := g.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ResetContainerSpecGcsPath() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerSpecGcsPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		g,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ResetLaunchOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetLaunchOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ResetTransformNameMappings() {
	_jsii_.InvokeVoid(
		g,
		"resetTransformNameMappings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		g,
		"resetUpdate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package googledatapipelinepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatapipelinepipeline/internal"
)

type GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference interface {
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
	GcsPath() *string
	SetGcsPath(val *string)
	GcsPathInput() *string
	InternalValue() *GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequest
	SetInternalValue(val *GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequest)
	LaunchParameters() GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersOutputReference
	LaunchParametersInput() *GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParameters
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValidateOnly() interface{}
	SetValidateOnly(val interface{})
	ValidateOnlyInput() interface{}
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
	PutLaunchParameters(value *GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParameters)
	ResetGcsPath()
	ResetLaunchParameters()
	ResetLocation()
	ResetValidateOnly()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference
type jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GcsPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GcsPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) InternalValue() *GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequest {
	var returns *GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequest
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) LaunchParameters() GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersOutputReference {
	var returns GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersOutputReference
	_jsii_.Get(
		j,
		"launchParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) LaunchParametersInput() *GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParameters {
	var returns *GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParameters
	_jsii_.Get(
		j,
		"launchParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ValidateOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ValidateOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateOnlyInput",
		&returns,
	)
	return returns
}


func NewGoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataPipelinePipeline.GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference_Override(g GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataPipelinePipeline.GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference)SetGcsPath(val *string) {
	if err := j.validateSetGcsPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcsPath",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference)SetInternalValue(val *GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequest) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference)SetValidateOnly(val interface{}) {
	if err := j.validateSetValidateOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validateOnly",
		val,
	)
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) PutLaunchParameters(value *GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParameters) {
	if err := g.validatePutLaunchParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLaunchParameters",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ResetGcsPath() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ResetLaunchParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetLaunchParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ResetValidateOnly() {
	_jsii_.InvokeVoid(
		g,
		"resetValidateOnly",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


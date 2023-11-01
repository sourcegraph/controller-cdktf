package googledatapipelinepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatapipelinepipeline/internal"
)

type GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference interface {
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
	InternalValue() *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequest
	SetInternalValue(val *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequest)
	LaunchParameter() GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference
	LaunchParameterInput() *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameter
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
	PutLaunchParameter(value *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameter)
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

// The jsii proxy struct for GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference
type jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) InternalValue() *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequest {
	var returns *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequest
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) LaunchParameter() GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference {
	var returns GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference
	_jsii_.Get(
		j,
		"launchParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) LaunchParameterInput() *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameter {
	var returns *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameter
	_jsii_.Get(
		j,
		"launchParameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) ValidateOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) ValidateOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateOnlyInput",
		&returns,
	)
	return returns
}


func NewGoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataPipelinePipeline.GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference_Override(g GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataPipelinePipeline.GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference)SetInternalValue(val *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequest) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference)SetValidateOnly(val interface{}) {
	if err := j.validateSetValidateOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validateOnly",
		val,
	)
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) PutLaunchParameter(value *GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameter) {
	if err := g.validatePutLaunchParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLaunchParameter",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) ResetValidateOnly() {
	_jsii_.InvokeVoid(
		g,
		"resetValidateOnly",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


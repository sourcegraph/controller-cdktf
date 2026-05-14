package googlecesevaluation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesevaluation/internal"
)

type GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference interface {
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
	DisplayName() *string
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponse
	SetInternalValue(val *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponse)
	Response() *map[string]*string
	SetResponse(val *map[string]*string)
	ResponseInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Tool() *string
	SetTool(val *string)
	ToolInput() *string
	ToolsetTool() GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseToolsetToolOutputReference
	ToolsetToolInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseToolsetTool
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
	PutToolsetTool(value *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseToolsetTool)
	ResetId()
	ResetResponse()
	ResetTool()
	ResetToolsetTool()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference
type jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) InternalValue() *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponse {
	var returns *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponse
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) Response() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"response",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) ResponseInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"responseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) Tool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) ToolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) ToolsetTool() GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseToolsetToolOutputReference {
	var returns GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseToolsetToolOutputReference
	_jsii_.Get(
		j,
		"toolsetTool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) ToolsetToolInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseToolsetTool {
	var returns *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseToolsetTool
	_jsii_.Get(
		j,
		"toolsetToolInput",
		&returns,
	)
	return returns
}


func NewGoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesEvaluation.GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference_Override(g GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesEvaluation.GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference)SetInternalValue(val *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponse) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference)SetResponse(val *map[string]*string) {
	if err := j.validateSetResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"response",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference)SetTool(val *string) {
	if err := j.validateSetToolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tool",
		val,
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) PutToolsetTool(value *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseToolsetTool) {
	if err := g.validatePutToolsetToolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putToolsetTool",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) ResetResponse() {
	_jsii_.InvokeVoid(
		g,
		"resetResponse",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) ResetTool() {
	_jsii_.InvokeVoid(
		g,
		"resetTool",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) ResetToolsetTool() {
	_jsii_.InvokeVoid(
		g,
		"resetToolsetTool",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package googlecesevaluation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesevaluation/internal"
)

type GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference interface {
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
	InternalValue() *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponse
	SetInternalValue(val *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponse)
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
	ToolsetTool() GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseToolsetToolOutputReference
	ToolsetToolInput() *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseToolsetTool
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
	PutToolsetTool(value *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseToolsetTool)
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

// The jsii proxy struct for GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference
type jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) InternalValue() *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponse {
	var returns *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponse
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) Response() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"response",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) ResponseInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"responseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) Tool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) ToolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) ToolsetTool() GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseToolsetToolOutputReference {
	var returns GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseToolsetToolOutputReference
	_jsii_.Get(
		j,
		"toolsetTool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) ToolsetToolInput() *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseToolsetTool {
	var returns *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseToolsetTool
	_jsii_.Get(
		j,
		"toolsetToolInput",
		&returns,
	)
	return returns
}


func NewGoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesEvaluation.GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference_Override(g GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesEvaluation.GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference)SetInternalValue(val *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponse) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference)SetResponse(val *map[string]*string) {
	if err := j.validateSetResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"response",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference)SetTool(val *string) {
	if err := j.validateSetToolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tool",
		val,
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) PutToolsetTool(value *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseToolsetTool) {
	if err := g.validatePutToolsetToolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putToolsetTool",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) ResetResponse() {
	_jsii_.InvokeVoid(
		g,
		"resetResponse",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) ResetTool() {
	_jsii_.InvokeVoid(
		g,
		"resetTool",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) ResetToolsetTool() {
	_jsii_.InvokeVoid(
		g,
		"resetToolsetTool",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package googlecesevaluation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesevaluation/internal"
)

type GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponse
	SetInternalValue(val *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponse)
	Response() *map[string]*string
	SetResponse(val *map[string]*string)
	ResponseInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Tool() *string
	SetTool(val *string)
	ToolInput() *string
	ToolsetTool() GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseToolsetToolOutputReference
	ToolsetToolInput() *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseToolsetTool
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutToolsetTool(value *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseToolsetTool)
	ResetId()
	ResetResponse()
	ResetTool()
	ResetToolsetTool()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference
type jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) InternalValue() *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponse {
	var returns *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponse
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) Response() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"response",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) ResponseInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"responseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) Tool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) ToolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) ToolsetTool() GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseToolsetToolOutputReference {
	var returns GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseToolsetToolOutputReference
	_jsii_.Get(
		j,
		"toolsetTool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) ToolsetToolInput() *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseToolsetTool {
	var returns *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseToolsetTool
	_jsii_.Get(
		j,
		"toolsetToolInput",
		&returns,
	)
	return returns
}


func NewGoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesEvaluation.GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference_Override(g GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesEvaluation.GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference)SetInternalValue(val *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponse) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference)SetResponse(val *map[string]*string) {
	if err := j.validateSetResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"response",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference)SetTool(val *string) {
	if err := j.validateSetToolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tool",
		val,
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) PutToolsetTool(value *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseToolsetTool) {
	if err := g.validatePutToolsetToolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putToolsetTool",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) ResetResponse() {
	_jsii_.InvokeVoid(
		g,
		"resetResponse",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) ResetTool() {
	_jsii_.InvokeVoid(
		g,
		"resetTool",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) ResetToolsetTool() {
	_jsii_.InvokeVoid(
		g,
		"resetToolsetTool",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package googlecesevaluation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesevaluation/internal"
)

type GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference interface {
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
	ExpectedToolCall() GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationExpectedToolCallOutputReference
	ExpectedToolCallInput() *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationExpectedToolCall
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectation
	SetInternalValue(val *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectation)
	MockToolResponse() GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference
	MockToolResponseInput() *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponse
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
	PutExpectedToolCall(value *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationExpectedToolCall)
	PutMockToolResponse(value *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponse)
	ResetExpectedToolCall()
	ResetMockToolResponse()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference
type jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) ExpectedToolCall() GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationExpectedToolCallOutputReference {
	var returns GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationExpectedToolCallOutputReference
	_jsii_.Get(
		j,
		"expectedToolCall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) ExpectedToolCallInput() *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationExpectedToolCall {
	var returns *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationExpectedToolCall
	_jsii_.Get(
		j,
		"expectedToolCallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) InternalValue() *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectation {
	var returns *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) MockToolResponse() GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference {
	var returns GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponseOutputReference
	_jsii_.Get(
		j,
		"mockToolResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) MockToolResponseInput() *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponse {
	var returns *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponse
	_jsii_.Get(
		j,
		"mockToolResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesEvaluation.GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference_Override(g GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesEvaluation.GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference)SetInternalValue(val *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) PutExpectedToolCall(value *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationExpectedToolCall) {
	if err := g.validatePutExpectedToolCallParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExpectedToolCall",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) PutMockToolResponse(value *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponse) {
	if err := g.validatePutMockToolResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMockToolResponse",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) ResetExpectedToolCall() {
	_jsii_.InvokeVoid(
		g,
		"resetExpectedToolCall",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) ResetMockToolResponse() {
	_jsii_.InvokeVoid(
		g,
		"resetMockToolResponse",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


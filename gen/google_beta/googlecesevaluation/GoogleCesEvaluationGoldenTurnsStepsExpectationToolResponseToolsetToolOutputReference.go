package googlecesevaluation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesevaluation/internal"
)

type GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetTool
	SetInternalValue(val *GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetTool)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ToolId() *string
	SetToolId(val *string)
	ToolIdInput() *string
	Toolset() *string
	SetToolset(val *string)
	ToolsetInput() *string
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
	ResetToolId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference
type jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) InternalValue() *GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetTool {
	var returns *GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetTool
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) ToolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) ToolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) Toolset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) ToolsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolsetInput",
		&returns,
	)
	return returns
}


func NewGoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesEvaluation.GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference_Override(g GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesEvaluation.GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference)SetInternalValue(val *GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetTool) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference)SetToolId(val *string) {
	if err := j.validateSetToolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toolId",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference)SetToolset(val *string) {
	if err := j.validateSetToolsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toolset",
		val,
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) ResetToolId() {
	_jsii_.InvokeVoid(
		g,
		"resetToolId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseToolsetToolOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


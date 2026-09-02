package googlecesapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesapp/internal"
)

type GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference interface {
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
	ExpectationLevelMetricsThresholds() GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsExpectationLevelMetricsThresholdsOutputReference
	ExpectationLevelMetricsThresholdsInput() *GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsExpectationLevelMetricsThresholds
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholds
	SetInternalValue(val *GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholds)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TurnLevelMetricsThresholds() GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsTurnLevelMetricsThresholdsOutputReference
	TurnLevelMetricsThresholdsInput() *GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsTurnLevelMetricsThresholds
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
	PutExpectationLevelMetricsThresholds(value *GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsExpectationLevelMetricsThresholds)
	PutTurnLevelMetricsThresholds(value *GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsTurnLevelMetricsThresholds)
	ResetExpectationLevelMetricsThresholds()
	ResetTurnLevelMetricsThresholds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference
type jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) ExpectationLevelMetricsThresholds() GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsExpectationLevelMetricsThresholdsOutputReference {
	var returns GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsExpectationLevelMetricsThresholdsOutputReference
	_jsii_.Get(
		j,
		"expectationLevelMetricsThresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) ExpectationLevelMetricsThresholdsInput() *GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsExpectationLevelMetricsThresholds {
	var returns *GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsExpectationLevelMetricsThresholds
	_jsii_.Get(
		j,
		"expectationLevelMetricsThresholdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) InternalValue() *GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholds {
	var returns *GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholds
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) TurnLevelMetricsThresholds() GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsTurnLevelMetricsThresholdsOutputReference {
	var returns GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsTurnLevelMetricsThresholdsOutputReference
	_jsii_.Get(
		j,
		"turnLevelMetricsThresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) TurnLevelMetricsThresholdsInput() *GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsTurnLevelMetricsThresholds {
	var returns *GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsTurnLevelMetricsThresholds
	_jsii_.Get(
		j,
		"turnLevelMetricsThresholdsInput",
		&returns,
	)
	return returns
}


func NewGoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference_Override(g GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference)SetInternalValue(val *GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholds) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) PutExpectationLevelMetricsThresholds(value *GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsExpectationLevelMetricsThresholds) {
	if err := g.validatePutExpectationLevelMetricsThresholdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExpectationLevelMetricsThresholds",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) PutTurnLevelMetricsThresholds(value *GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsTurnLevelMetricsThresholds) {
	if err := g.validatePutTurnLevelMetricsThresholdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTurnLevelMetricsThresholds",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) ResetExpectationLevelMetricsThresholds() {
	_jsii_.InvokeVoid(
		g,
		"resetExpectationLevelMetricsThresholds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) ResetTurnLevelMetricsThresholds() {
	_jsii_.InvokeVoid(
		g,
		"resetTurnLevelMetricsThresholds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


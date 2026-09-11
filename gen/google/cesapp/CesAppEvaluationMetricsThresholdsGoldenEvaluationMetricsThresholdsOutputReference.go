package cesapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/cesapp/internal"
)

type CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference interface {
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
	ExpectationLevelMetricsThresholds() CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsExpectationLevelMetricsThresholdsOutputReference
	ExpectationLevelMetricsThresholdsInput() *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsExpectationLevelMetricsThresholds
	// Experimental.
	Fqn() *string
	InternalValue() *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholds
	SetInternalValue(val *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholds)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TurnLevelMetricsThresholds() CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsTurnLevelMetricsThresholdsOutputReference
	TurnLevelMetricsThresholdsInput() *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsTurnLevelMetricsThresholds
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
	PutExpectationLevelMetricsThresholds(value *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsExpectationLevelMetricsThresholds)
	PutTurnLevelMetricsThresholds(value *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsTurnLevelMetricsThresholds)
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

// The jsii proxy struct for CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference
type jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) ExpectationLevelMetricsThresholds() CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsExpectationLevelMetricsThresholdsOutputReference {
	var returns CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsExpectationLevelMetricsThresholdsOutputReference
	_jsii_.Get(
		j,
		"expectationLevelMetricsThresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) ExpectationLevelMetricsThresholdsInput() *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsExpectationLevelMetricsThresholds {
	var returns *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsExpectationLevelMetricsThresholds
	_jsii_.Get(
		j,
		"expectationLevelMetricsThresholdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) InternalValue() *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholds {
	var returns *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholds
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) TurnLevelMetricsThresholds() CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsTurnLevelMetricsThresholdsOutputReference {
	var returns CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsTurnLevelMetricsThresholdsOutputReference
	_jsii_.Get(
		j,
		"turnLevelMetricsThresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) TurnLevelMetricsThresholdsInput() *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsTurnLevelMetricsThresholds {
	var returns *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsTurnLevelMetricsThresholds
	_jsii_.Get(
		j,
		"turnLevelMetricsThresholdsInput",
		&returns,
	)
	return returns
}


func NewCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference {
	_init_.Initialize()

	if err := validateNewCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cesApp.CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference_Override(c CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cesApp.CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference)SetInternalValue(val *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholds) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) PutExpectationLevelMetricsThresholds(value *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsExpectationLevelMetricsThresholds) {
	if err := c.validatePutExpectationLevelMetricsThresholdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putExpectationLevelMetricsThresholds",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) PutTurnLevelMetricsThresholds(value *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsTurnLevelMetricsThresholds) {
	if err := c.validatePutTurnLevelMetricsThresholdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTurnLevelMetricsThresholds",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) ResetExpectationLevelMetricsThresholds() {
	_jsii_.InvokeVoid(
		c,
		"resetExpectationLevelMetricsThresholds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) ResetTurnLevelMetricsThresholds() {
	_jsii_.InvokeVoid(
		c,
		"resetTurnLevelMetricsThresholds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


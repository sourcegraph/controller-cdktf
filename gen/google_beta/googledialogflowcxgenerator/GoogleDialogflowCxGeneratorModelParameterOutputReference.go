package googledialogflowcxgenerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowcxgenerator/internal"
)

type GoogleDialogflowCxGeneratorModelParameterOutputReference interface {
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
	InternalValue() *GoogleDialogflowCxGeneratorModelParameter
	SetInternalValue(val *GoogleDialogflowCxGeneratorModelParameter)
	MaxDecodeSteps() *float64
	SetMaxDecodeSteps(val *float64)
	MaxDecodeStepsInput() *float64
	Temperature() *float64
	SetTemperature(val *float64)
	TemperatureInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TopK() *float64
	SetTopK(val *float64)
	TopKInput() *float64
	TopP() *float64
	SetTopP(val *float64)
	TopPInput() *float64
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
	ResetMaxDecodeSteps()
	ResetTemperature()
	ResetTopK()
	ResetTopP()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxGeneratorModelParameterOutputReference
type jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) InternalValue() *GoogleDialogflowCxGeneratorModelParameter {
	var returns *GoogleDialogflowCxGeneratorModelParameter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) MaxDecodeSteps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDecodeSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) MaxDecodeStepsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDecodeStepsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) Temperature() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"temperature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) TemperatureInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"temperatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) TopK() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topK",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) TopKInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topKInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) TopP() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topP",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) TopPInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topPInput",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowCxGeneratorModelParameterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowCxGeneratorModelParameterOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxGeneratorModelParameterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxGenerator.GoogleDialogflowCxGeneratorModelParameterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxGeneratorModelParameterOutputReference_Override(g GoogleDialogflowCxGeneratorModelParameterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxGenerator.GoogleDialogflowCxGeneratorModelParameterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference)SetInternalValue(val *GoogleDialogflowCxGeneratorModelParameter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference)SetMaxDecodeSteps(val *float64) {
	if err := j.validateSetMaxDecodeStepsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDecodeSteps",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference)SetTemperature(val *float64) {
	if err := j.validateSetTemperatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"temperature",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference)SetTopK(val *float64) {
	if err := j.validateSetTopKParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topK",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference)SetTopP(val *float64) {
	if err := j.validateSetTopPParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topP",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) ResetMaxDecodeSteps() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxDecodeSteps",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) ResetTemperature() {
	_jsii_.InvokeVoid(
		g,
		"resetTemperature",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) ResetTopK() {
	_jsii_.InvokeVoid(
		g,
		"resetTopK",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) ResetTopP() {
	_jsii_.InvokeVoid(
		g,
		"resetTopP",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxGeneratorModelParameterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


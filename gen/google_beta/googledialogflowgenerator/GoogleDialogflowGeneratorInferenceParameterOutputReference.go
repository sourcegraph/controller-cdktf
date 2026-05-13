package googledialogflowgenerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowgenerator/internal"
)

type GoogleDialogflowGeneratorInferenceParameterOutputReference interface {
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
	InternalValue() *GoogleDialogflowGeneratorInferenceParameter
	SetInternalValue(val *GoogleDialogflowGeneratorInferenceParameter)
	MaxOutputTokens() *float64
	SetMaxOutputTokens(val *float64)
	MaxOutputTokensInput() *float64
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
	ResetMaxOutputTokens()
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

// The jsii proxy struct for GoogleDialogflowGeneratorInferenceParameterOutputReference
type jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) InternalValue() *GoogleDialogflowGeneratorInferenceParameter {
	var returns *GoogleDialogflowGeneratorInferenceParameter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) MaxOutputTokens() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxOutputTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) MaxOutputTokensInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxOutputTokensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) Temperature() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"temperature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) TemperatureInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"temperatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) TopK() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topK",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) TopKInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topKInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) TopP() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topP",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) TopPInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topPInput",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowGeneratorInferenceParameterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowGeneratorInferenceParameterOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowGeneratorInferenceParameterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowGenerator.GoogleDialogflowGeneratorInferenceParameterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowGeneratorInferenceParameterOutputReference_Override(g GoogleDialogflowGeneratorInferenceParameterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowGenerator.GoogleDialogflowGeneratorInferenceParameterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference)SetInternalValue(val *GoogleDialogflowGeneratorInferenceParameter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference)SetMaxOutputTokens(val *float64) {
	if err := j.validateSetMaxOutputTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxOutputTokens",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference)SetTemperature(val *float64) {
	if err := j.validateSetTemperatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"temperature",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference)SetTopK(val *float64) {
	if err := j.validateSetTopKParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topK",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference)SetTopP(val *float64) {
	if err := j.validateSetTopPParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topP",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) ResetMaxOutputTokens() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxOutputTokens",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) ResetTemperature() {
	_jsii_.InvokeVoid(
		g,
		"resetTemperature",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) ResetTopK() {
	_jsii_.InvokeVoid(
		g,
		"resetTopK",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) ResetTopP() {
	_jsii_.InvokeVoid(
		g,
		"resetTopP",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorInferenceParameterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


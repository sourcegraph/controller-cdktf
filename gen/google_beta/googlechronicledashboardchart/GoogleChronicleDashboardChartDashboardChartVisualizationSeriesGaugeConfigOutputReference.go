package googlechronicledashboardchart

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlechronicledashboardchart/internal"
)

type GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference interface {
	cdktn.ComplexObject
	BaseValue() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigBaseValueOutputReference
	BaseValueInput() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigBaseValue
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
	InternalValue() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfig
	SetInternalValue(val *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfig)
	LimitValue() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigLimitValueOutputReference
	LimitValueInput() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigLimitValue
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThresholdValues() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigThresholdValuesList
	ThresholdValuesInput() interface{}
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
	PutBaseValue(value *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigBaseValue)
	PutLimitValue(value *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigLimitValue)
	PutThresholdValues(value interface{})
	ResetBaseValue()
	ResetLimitValue()
	ResetThresholdValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference
type jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) BaseValue() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigBaseValueOutputReference {
	var returns GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigBaseValueOutputReference
	_jsii_.Get(
		j,
		"baseValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) BaseValueInput() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigBaseValue {
	var returns *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigBaseValue
	_jsii_.Get(
		j,
		"baseValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) InternalValue() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfig {
	var returns *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) LimitValue() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigLimitValueOutputReference {
	var returns GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigLimitValueOutputReference
	_jsii_.Get(
		j,
		"limitValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) LimitValueInput() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigLimitValue {
	var returns *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigLimitValue
	_jsii_.Get(
		j,
		"limitValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) ThresholdValues() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigThresholdValuesList {
	var returns GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigThresholdValuesList
	_jsii_.Get(
		j,
		"thresholdValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) ThresholdValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thresholdValuesInput",
		&returns,
	)
	return returns
}


func NewGoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleDashboardChart.GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference_Override(g GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleDashboardChart.GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference)SetInternalValue(val *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) PutBaseValue(value *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigBaseValue) {
	if err := g.validatePutBaseValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBaseValue",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) PutLimitValue(value *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigLimitValue) {
	if err := g.validatePutLimitValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLimitValue",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) PutThresholdValues(value interface{}) {
	if err := g.validatePutThresholdValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putThresholdValues",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) ResetBaseValue() {
	_jsii_.InvokeVoid(
		g,
		"resetBaseValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) ResetLimitValue() {
	_jsii_.InvokeVoid(
		g,
		"resetLimitValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) ResetThresholdValues() {
	_jsii_.InvokeVoid(
		g,
		"resetThresholdValues",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


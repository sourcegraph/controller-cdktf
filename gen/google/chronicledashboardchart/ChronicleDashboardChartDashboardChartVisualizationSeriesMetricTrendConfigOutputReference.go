package chronicledashboardchart

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/chronicledashboardchart/internal"
)

type ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	InternalValue() *ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfig
	SetInternalValue(val *ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfig)
	MetricDisplayTrend() *string
	SetMetricDisplayTrend(val *string)
	MetricDisplayTrendInput() *string
	MetricFormat() *string
	SetMetricFormat(val *string)
	MetricFormatInput() *string
	MetricTrendType() *string
	SetMetricTrendType(val *string)
	MetricTrendTypeInput() *string
	ShowMetricTrend() any
	SetShowMetricTrend(val any)
	ShowMetricTrendInput() any
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
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	ResetMetricDisplayTrend()
	ResetMetricFormat()
	ResetMetricTrendType()
	ResetShowMetricTrend()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference
type jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) InternalValue() *ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfig {
	var returns *ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) MetricDisplayTrend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricDisplayTrend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) MetricDisplayTrendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricDisplayTrendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) MetricFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) MetricFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) MetricTrendType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricTrendType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) MetricTrendTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricTrendTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) ShowMetricTrend() any {
	var returns any
	_jsii_.Get(
		j,
		"showMetricTrend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) ShowMetricTrendInput() any {
	var returns any
	_jsii_.Get(
		j,
		"showMetricTrendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.chronicleDashboardChart.ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference_Override(c ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.chronicleDashboardChart.ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) SetInternalValue(val *ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) SetMetricDisplayTrend(val *string) {
	if err := j.validateSetMetricDisplayTrendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricDisplayTrend",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) SetMetricFormat(val *string) {
	if err := j.validateSetMetricFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricFormat",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) SetMetricTrendType(val *string) {
	if err := j.validateSetMetricTrendTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricTrendType",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) SetShowMetricTrend(val any) {
	if err := j.validateSetShowMetricTrendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showMetricTrend",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) ResetMetricDisplayTrend() {
	_jsii_.InvokeVoid(
		c,
		"resetMetricDisplayTrend",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) ResetMetricFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetMetricFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) ResetMetricTrendType() {
	_jsii_.InvokeVoid(
		c,
		"resetMetricTrendType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) ResetShowMetricTrend() {
	_jsii_.InvokeVoid(
		c,
		"resetShowMetricTrend",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		c,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

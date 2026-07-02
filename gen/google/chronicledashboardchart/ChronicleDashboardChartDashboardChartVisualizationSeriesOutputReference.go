package chronicledashboardchart

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/chronicledashboardchart/internal"
)

type ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference interface {
	cdktf.ComplexObject
	AreaStyle() ChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference
	AreaStyleInput() *ChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyle
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
	DataLabel() ChronicleDashboardChartDashboardChartVisualizationSeriesDataLabelOutputReference
	DataLabelInput() *ChronicleDashboardChartDashboardChartVisualizationSeriesDataLabel
	Encode() ChronicleDashboardChartDashboardChartVisualizationSeriesEncodeOutputReference
	EncodeInput() *ChronicleDashboardChartDashboardChartVisualizationSeriesEncode
	Field() *string
	SetField(val *string)
	FieldInput() *string
	// Experimental.
	Fqn() *string
	GaugeConfig() ChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference
	GaugeConfigInput() *ChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfig
	InternalValue() any
	SetInternalValue(val any)
	ItemColors() ChronicleDashboardChartDashboardChartVisualizationSeriesItemColorsOutputReference
	ItemColorsInput() *ChronicleDashboardChartDashboardChartVisualizationSeriesItemColors
	ItemStyle() ChronicleDashboardChartDashboardChartVisualizationSeriesItemStyleOutputReference
	ItemStyleInput() *ChronicleDashboardChartDashboardChartVisualizationSeriesItemStyle
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	MetricTrendConfig() ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference
	MetricTrendConfigInput() *ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfig
	Radius() *[]*string
	SetRadius(val *[]*string)
	RadiusInput() *[]*string
	SeriesName() *string
	SetSeriesName(val *string)
	SeriesNameInput() *string
	SeriesStackStrategy() *string
	SetSeriesStackStrategy(val *string)
	SeriesStackStrategyInput() *string
	SeriesType() *string
	SetSeriesType(val *string)
	SeriesTypeInput() *string
	SeriesUniqueValue() *string
	SetSeriesUniqueValue(val *string)
	SeriesUniqueValueInput() *string
	ShowBackground() any
	SetShowBackground(val any)
	ShowBackgroundInput() any
	ShowSymbol() any
	SetShowSymbol(val any)
	ShowSymbolInput() any
	Stack() *string
	SetStack(val *string)
	StackInput() *string
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
	PutAreaStyle(value *ChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyle)
	PutDataLabel(value *ChronicleDashboardChartDashboardChartVisualizationSeriesDataLabel)
	PutEncode(value *ChronicleDashboardChartDashboardChartVisualizationSeriesEncode)
	PutGaugeConfig(value *ChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfig)
	PutItemColors(value *ChronicleDashboardChartDashboardChartVisualizationSeriesItemColors)
	PutItemStyle(value *ChronicleDashboardChartDashboardChartVisualizationSeriesItemStyle)
	PutMetricTrendConfig(value *ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfig)
	ResetAreaStyle()
	ResetDataLabel()
	ResetEncode()
	ResetField()
	ResetGaugeConfig()
	ResetItemColors()
	ResetItemStyle()
	ResetLabel()
	ResetMetricTrendConfig()
	ResetRadius()
	ResetSeriesName()
	ResetSeriesStackStrategy()
	ResetSeriesType()
	ResetSeriesUniqueValue()
	ResetShowBackground()
	ResetShowSymbol()
	ResetStack()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference
type jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) AreaStyle() ChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference {
	var returns ChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference
	_jsii_.Get(
		j,
		"areaStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) AreaStyleInput() *ChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyle {
	var returns *ChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyle
	_jsii_.Get(
		j,
		"areaStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) DataLabel() ChronicleDashboardChartDashboardChartVisualizationSeriesDataLabelOutputReference {
	var returns ChronicleDashboardChartDashboardChartVisualizationSeriesDataLabelOutputReference
	_jsii_.Get(
		j,
		"dataLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) DataLabelInput() *ChronicleDashboardChartDashboardChartVisualizationSeriesDataLabel {
	var returns *ChronicleDashboardChartDashboardChartVisualizationSeriesDataLabel
	_jsii_.Get(
		j,
		"dataLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) Encode() ChronicleDashboardChartDashboardChartVisualizationSeriesEncodeOutputReference {
	var returns ChronicleDashboardChartDashboardChartVisualizationSeriesEncodeOutputReference
	_jsii_.Get(
		j,
		"encode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) EncodeInput() *ChronicleDashboardChartDashboardChartVisualizationSeriesEncode {
	var returns *ChronicleDashboardChartDashboardChartVisualizationSeriesEncode
	_jsii_.Get(
		j,
		"encodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) Field() *string {
	var returns *string
	_jsii_.Get(
		j,
		"field",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) FieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GaugeConfig() ChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference {
	var returns ChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference
	_jsii_.Get(
		j,
		"gaugeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GaugeConfigInput() *ChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfig {
	var returns *ChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfig
	_jsii_.Get(
		j,
		"gaugeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ItemColors() ChronicleDashboardChartDashboardChartVisualizationSeriesItemColorsOutputReference {
	var returns ChronicleDashboardChartDashboardChartVisualizationSeriesItemColorsOutputReference
	_jsii_.Get(
		j,
		"itemColors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ItemColorsInput() *ChronicleDashboardChartDashboardChartVisualizationSeriesItemColors {
	var returns *ChronicleDashboardChartDashboardChartVisualizationSeriesItemColors
	_jsii_.Get(
		j,
		"itemColorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ItemStyle() ChronicleDashboardChartDashboardChartVisualizationSeriesItemStyleOutputReference {
	var returns ChronicleDashboardChartDashboardChartVisualizationSeriesItemStyleOutputReference
	_jsii_.Get(
		j,
		"itemStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ItemStyleInput() *ChronicleDashboardChartDashboardChartVisualizationSeriesItemStyle {
	var returns *ChronicleDashboardChartDashboardChartVisualizationSeriesItemStyle
	_jsii_.Get(
		j,
		"itemStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) MetricTrendConfig() ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference {
	var returns ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference
	_jsii_.Get(
		j,
		"metricTrendConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) MetricTrendConfigInput() *ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfig {
	var returns *ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfig
	_jsii_.Get(
		j,
		"metricTrendConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) Radius() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"radius",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) RadiusInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"radiusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SeriesName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"seriesName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SeriesNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"seriesNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SeriesStackStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"seriesStackStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SeriesStackStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"seriesStackStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SeriesType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"seriesType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SeriesTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"seriesTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SeriesUniqueValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"seriesUniqueValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SeriesUniqueValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"seriesUniqueValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ShowBackground() any {
	var returns any
	_jsii_.Get(
		j,
		"showBackground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ShowBackgroundInput() any {
	var returns any
	_jsii_.Get(
		j,
		"showBackgroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ShowSymbol() any {
	var returns any
	_jsii_.Get(
		j,
		"showSymbol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ShowSymbolInput() any {
	var returns any
	_jsii_.Get(
		j,
		"showSymbolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) Stack() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) StackInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleDashboardChartDashboardChartVisualizationSeriesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.chronicleDashboardChart.ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference_Override(c ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.chronicleDashboardChart.ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SetField(val *string) {
	if err := j.validateSetFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"field",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SetRadius(val *[]*string) {
	if err := j.validateSetRadiusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"radius",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SetSeriesName(val *string) {
	if err := j.validateSetSeriesNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"seriesName",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SetSeriesStackStrategy(val *string) {
	if err := j.validateSetSeriesStackStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"seriesStackStrategy",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SetSeriesType(val *string) {
	if err := j.validateSetSeriesTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"seriesType",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SetSeriesUniqueValue(val *string) {
	if err := j.validateSetSeriesUniqueValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"seriesUniqueValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SetShowBackground(val any) {
	if err := j.validateSetShowBackgroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showBackground",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SetShowSymbol(val any) {
	if err := j.validateSetShowSymbolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showSymbol",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SetStack(val *string) {
	if err := j.validateSetStackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stack",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) PutAreaStyle(value *ChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyle) {
	if err := c.validatePutAreaStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAreaStyle",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) PutDataLabel(value *ChronicleDashboardChartDashboardChartVisualizationSeriesDataLabel) {
	if err := c.validatePutDataLabelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDataLabel",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) PutEncode(value *ChronicleDashboardChartDashboardChartVisualizationSeriesEncode) {
	if err := c.validatePutEncodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEncode",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) PutGaugeConfig(value *ChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfig) {
	if err := c.validatePutGaugeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGaugeConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) PutItemColors(value *ChronicleDashboardChartDashboardChartVisualizationSeriesItemColors) {
	if err := c.validatePutItemColorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putItemColors",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) PutItemStyle(value *ChronicleDashboardChartDashboardChartVisualizationSeriesItemStyle) {
	if err := c.validatePutItemStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putItemStyle",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) PutMetricTrendConfig(value *ChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfig) {
	if err := c.validatePutMetricTrendConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMetricTrendConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetAreaStyle() {
	_jsii_.InvokeVoid(
		c,
		"resetAreaStyle",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetDataLabel() {
	_jsii_.InvokeVoid(
		c,
		"resetDataLabel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetEncode() {
	_jsii_.InvokeVoid(
		c,
		"resetEncode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetField() {
	_jsii_.InvokeVoid(
		c,
		"resetField",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetGaugeConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetGaugeConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetItemColors() {
	_jsii_.InvokeVoid(
		c,
		"resetItemColors",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetItemStyle() {
	_jsii_.InvokeVoid(
		c,
		"resetItemStyle",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetLabel() {
	_jsii_.InvokeVoid(
		c,
		"resetLabel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetMetricTrendConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetMetricTrendConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetRadius() {
	_jsii_.InvokeVoid(
		c,
		"resetRadius",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetSeriesName() {
	_jsii_.InvokeVoid(
		c,
		"resetSeriesName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetSeriesStackStrategy() {
	_jsii_.InvokeVoid(
		c,
		"resetSeriesStackStrategy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetSeriesType() {
	_jsii_.InvokeVoid(
		c,
		"resetSeriesType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetSeriesUniqueValue() {
	_jsii_.InvokeVoid(
		c,
		"resetSeriesUniqueValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetShowBackground() {
	_jsii_.InvokeVoid(
		c,
		"resetShowBackground",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetShowSymbol() {
	_jsii_.InvokeVoid(
		c,
		"resetShowSymbol",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetStack() {
	_jsii_.InvokeVoid(
		c,
		"resetStack",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

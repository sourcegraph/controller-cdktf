package googlechronicledashboardchart

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlechronicledashboardchart/internal"
)

type GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference interface {
	cdktf.ComplexObject
	AreaStyle() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference
	AreaStyleInput() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyle
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
	DataLabel() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesDataLabelOutputReference
	DataLabelInput() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesDataLabel
	Encode() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesEncodeOutputReference
	EncodeInput() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesEncode
	Field() *string
	SetField(val *string)
	FieldInput() *string
	// Experimental.
	Fqn() *string
	GaugeConfig() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference
	GaugeConfigInput() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfig
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ItemColors() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesItemColorsOutputReference
	ItemColorsInput() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesItemColors
	ItemStyle() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesItemStyleOutputReference
	ItemStyleInput() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesItemStyle
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	MetricTrendConfig() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference
	MetricTrendConfigInput() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfig
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
	ShowBackground() interface{}
	SetShowBackground(val interface{})
	ShowBackgroundInput() interface{}
	ShowSymbol() interface{}
	SetShowSymbol(val interface{})
	ShowSymbolInput() interface{}
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
	PutAreaStyle(value *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyle)
	PutDataLabel(value *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesDataLabel)
	PutEncode(value *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesEncode)
	PutGaugeConfig(value *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfig)
	PutItemColors(value *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesItemColors)
	PutItemStyle(value *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesItemStyle)
	PutMetricTrendConfig(value *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfig)
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
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference
type jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) AreaStyle() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference {
	var returns GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference
	_jsii_.Get(
		j,
		"areaStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) AreaStyleInput() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyle {
	var returns *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyle
	_jsii_.Get(
		j,
		"areaStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) DataLabel() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesDataLabelOutputReference {
	var returns GoogleChronicleDashboardChartDashboardChartVisualizationSeriesDataLabelOutputReference
	_jsii_.Get(
		j,
		"dataLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) DataLabelInput() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesDataLabel {
	var returns *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesDataLabel
	_jsii_.Get(
		j,
		"dataLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) Encode() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesEncodeOutputReference {
	var returns GoogleChronicleDashboardChartDashboardChartVisualizationSeriesEncodeOutputReference
	_jsii_.Get(
		j,
		"encode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) EncodeInput() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesEncode {
	var returns *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesEncode
	_jsii_.Get(
		j,
		"encodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) Field() *string {
	var returns *string
	_jsii_.Get(
		j,
		"field",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) FieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GaugeConfig() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference {
	var returns GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfigOutputReference
	_jsii_.Get(
		j,
		"gaugeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GaugeConfigInput() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfig {
	var returns *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfig
	_jsii_.Get(
		j,
		"gaugeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ItemColors() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesItemColorsOutputReference {
	var returns GoogleChronicleDashboardChartDashboardChartVisualizationSeriesItemColorsOutputReference
	_jsii_.Get(
		j,
		"itemColors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ItemColorsInput() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesItemColors {
	var returns *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesItemColors
	_jsii_.Get(
		j,
		"itemColorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ItemStyle() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesItemStyleOutputReference {
	var returns GoogleChronicleDashboardChartDashboardChartVisualizationSeriesItemStyleOutputReference
	_jsii_.Get(
		j,
		"itemStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ItemStyleInput() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesItemStyle {
	var returns *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesItemStyle
	_jsii_.Get(
		j,
		"itemStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) MetricTrendConfig() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference {
	var returns GoogleChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfigOutputReference
	_jsii_.Get(
		j,
		"metricTrendConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) MetricTrendConfigInput() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfig {
	var returns *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfig
	_jsii_.Get(
		j,
		"metricTrendConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) Radius() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"radius",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) RadiusInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"radiusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SeriesName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"seriesName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SeriesNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"seriesNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SeriesStackStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"seriesStackStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SeriesStackStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"seriesStackStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SeriesType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"seriesType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SeriesTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"seriesTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SeriesUniqueValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"seriesUniqueValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) SeriesUniqueValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"seriesUniqueValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ShowBackground() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showBackground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ShowBackgroundInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showBackgroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ShowSymbol() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showSymbol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ShowSymbolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showSymbolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) Stack() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) StackInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleChronicleDashboardChart.GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference_Override(g GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleChronicleDashboardChart.GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference)SetField(val *string) {
	if err := j.validateSetFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"field",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference)SetRadius(val *[]*string) {
	if err := j.validateSetRadiusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"radius",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference)SetSeriesName(val *string) {
	if err := j.validateSetSeriesNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"seriesName",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference)SetSeriesStackStrategy(val *string) {
	if err := j.validateSetSeriesStackStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"seriesStackStrategy",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference)SetSeriesType(val *string) {
	if err := j.validateSetSeriesTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"seriesType",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference)SetSeriesUniqueValue(val *string) {
	if err := j.validateSetSeriesUniqueValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"seriesUniqueValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference)SetShowBackground(val interface{}) {
	if err := j.validateSetShowBackgroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showBackground",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference)SetShowSymbol(val interface{}) {
	if err := j.validateSetShowSymbolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showSymbol",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference)SetStack(val *string) {
	if err := j.validateSetStackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stack",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) PutAreaStyle(value *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyle) {
	if err := g.validatePutAreaStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAreaStyle",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) PutDataLabel(value *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesDataLabel) {
	if err := g.validatePutDataLabelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataLabel",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) PutEncode(value *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesEncode) {
	if err := g.validatePutEncodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncode",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) PutGaugeConfig(value *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesGaugeConfig) {
	if err := g.validatePutGaugeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGaugeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) PutItemColors(value *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesItemColors) {
	if err := g.validatePutItemColorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putItemColors",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) PutItemStyle(value *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesItemStyle) {
	if err := g.validatePutItemStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putItemStyle",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) PutMetricTrendConfig(value *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesMetricTrendConfig) {
	if err := g.validatePutMetricTrendConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMetricTrendConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetAreaStyle() {
	_jsii_.InvokeVoid(
		g,
		"resetAreaStyle",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetDataLabel() {
	_jsii_.InvokeVoid(
		g,
		"resetDataLabel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetEncode() {
	_jsii_.InvokeVoid(
		g,
		"resetEncode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetField() {
	_jsii_.InvokeVoid(
		g,
		"resetField",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetGaugeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGaugeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetItemColors() {
	_jsii_.InvokeVoid(
		g,
		"resetItemColors",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetItemStyle() {
	_jsii_.InvokeVoid(
		g,
		"resetItemStyle",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetLabel() {
	_jsii_.InvokeVoid(
		g,
		"resetLabel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetMetricTrendConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetMetricTrendConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetRadius() {
	_jsii_.InvokeVoid(
		g,
		"resetRadius",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetSeriesName() {
	_jsii_.InvokeVoid(
		g,
		"resetSeriesName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetSeriesStackStrategy() {
	_jsii_.InvokeVoid(
		g,
		"resetSeriesStackStrategy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetSeriesType() {
	_jsii_.InvokeVoid(
		g,
		"resetSeriesType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetSeriesUniqueValue() {
	_jsii_.InvokeVoid(
		g,
		"resetSeriesUniqueValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetShowBackground() {
	_jsii_.InvokeVoid(
		g,
		"resetShowBackground",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetShowSymbol() {
	_jsii_.InvokeVoid(
		g,
		"resetShowSymbol",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ResetStack() {
	_jsii_.InvokeVoid(
		g,
		"resetStack",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


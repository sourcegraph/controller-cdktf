package chronicledashboardchart

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/chronicledashboardchart/internal"
)

type ChronicleDashboardChartDashboardChartVisualizationOutputReference interface {
	cdktf.ComplexObject
	Button() ChronicleDashboardChartDashboardChartVisualizationButtonOutputReference
	ButtonInput() *ChronicleDashboardChartDashboardChartVisualizationButton
	ColumnDefs() ChronicleDashboardChartDashboardChartVisualizationColumnDefsList
	ColumnDefsInput() any
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
	GoogleMapsConfig() ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigOutputReference
	GoogleMapsConfigInput() *ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfig
	GroupingType() *string
	SetGroupingType(val *string)
	GroupingTypeInput() *string
	InternalValue() *ChronicleDashboardChartDashboardChartVisualization
	SetInternalValue(val *ChronicleDashboardChartDashboardChartVisualization)
	Legends() ChronicleDashboardChartDashboardChartVisualizationLegendsList
	LegendsInput() any
	Markdown() ChronicleDashboardChartDashboardChartVisualizationMarkdownOutputReference
	MarkdownInput() *ChronicleDashboardChartDashboardChartVisualizationMarkdown
	Series() ChronicleDashboardChartDashboardChartVisualizationSeriesList
	SeriesColumn() *[]*string
	SetSeriesColumn(val *[]*string)
	SeriesColumnInput() *[]*string
	SeriesInput() any
	TableConfig() ChronicleDashboardChartDashboardChartVisualizationTableConfigOutputReference
	TableConfigInput() *ChronicleDashboardChartDashboardChartVisualizationTableConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThresholdColoringEnabled() any
	SetThresholdColoringEnabled(val any)
	ThresholdColoringEnabledInput() any
	Tooltip() ChronicleDashboardChartDashboardChartVisualizationTooltipOutputReference
	TooltipInput() *ChronicleDashboardChartDashboardChartVisualizationTooltip
	VisualMaps() ChronicleDashboardChartDashboardChartVisualizationVisualMapsList
	VisualMapsInput() any
	XAxes() ChronicleDashboardChartDashboardChartVisualizationXAxesList
	XAxesInput() any
	YAxes() ChronicleDashboardChartDashboardChartVisualizationYAxesList
	YAxesInput() any
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
	PutButton(value *ChronicleDashboardChartDashboardChartVisualizationButton)
	PutColumnDefs(value any)
	PutGoogleMapsConfig(value *ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfig)
	PutLegends(value any)
	PutMarkdown(value *ChronicleDashboardChartDashboardChartVisualizationMarkdown)
	PutSeries(value any)
	PutTableConfig(value *ChronicleDashboardChartDashboardChartVisualizationTableConfig)
	PutTooltip(value *ChronicleDashboardChartDashboardChartVisualizationTooltip)
	PutVisualMaps(value any)
	PutXAxes(value any)
	PutYAxes(value any)
	ResetButton()
	ResetColumnDefs()
	ResetGoogleMapsConfig()
	ResetGroupingType()
	ResetLegends()
	ResetMarkdown()
	ResetSeries()
	ResetSeriesColumn()
	ResetTableConfig()
	ResetThresholdColoringEnabled()
	ResetTooltip()
	ResetVisualMaps()
	ResetXAxes()
	ResetYAxes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleDashboardChartDashboardChartVisualizationOutputReference
type jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) Button() ChronicleDashboardChartDashboardChartVisualizationButtonOutputReference {
	var returns ChronicleDashboardChartDashboardChartVisualizationButtonOutputReference
	_jsii_.Get(
		j,
		"button",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ButtonInput() *ChronicleDashboardChartDashboardChartVisualizationButton {
	var returns *ChronicleDashboardChartDashboardChartVisualizationButton
	_jsii_.Get(
		j,
		"buttonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ColumnDefs() ChronicleDashboardChartDashboardChartVisualizationColumnDefsList {
	var returns ChronicleDashboardChartDashboardChartVisualizationColumnDefsList
	_jsii_.Get(
		j,
		"columnDefs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ColumnDefsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"columnDefsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) GoogleMapsConfig() ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigOutputReference {
	var returns ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigOutputReference
	_jsii_.Get(
		j,
		"googleMapsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) GoogleMapsConfigInput() *ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfig {
	var returns *ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfig
	_jsii_.Get(
		j,
		"googleMapsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) GroupingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) GroupingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) InternalValue() *ChronicleDashboardChartDashboardChartVisualization {
	var returns *ChronicleDashboardChartDashboardChartVisualization
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) Legends() ChronicleDashboardChartDashboardChartVisualizationLegendsList {
	var returns ChronicleDashboardChartDashboardChartVisualizationLegendsList
	_jsii_.Get(
		j,
		"legends",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) LegendsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"legendsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) Markdown() ChronicleDashboardChartDashboardChartVisualizationMarkdownOutputReference {
	var returns ChronicleDashboardChartDashboardChartVisualizationMarkdownOutputReference
	_jsii_.Get(
		j,
		"markdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) MarkdownInput() *ChronicleDashboardChartDashboardChartVisualizationMarkdown {
	var returns *ChronicleDashboardChartDashboardChartVisualizationMarkdown
	_jsii_.Get(
		j,
		"markdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) Series() ChronicleDashboardChartDashboardChartVisualizationSeriesList {
	var returns ChronicleDashboardChartDashboardChartVisualizationSeriesList
	_jsii_.Get(
		j,
		"series",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) SeriesColumn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"seriesColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) SeriesColumnInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"seriesColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) SeriesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"seriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) TableConfig() ChronicleDashboardChartDashboardChartVisualizationTableConfigOutputReference {
	var returns ChronicleDashboardChartDashboardChartVisualizationTableConfigOutputReference
	_jsii_.Get(
		j,
		"tableConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) TableConfigInput() *ChronicleDashboardChartDashboardChartVisualizationTableConfig {
	var returns *ChronicleDashboardChartDashboardChartVisualizationTableConfig
	_jsii_.Get(
		j,
		"tableConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ThresholdColoringEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"thresholdColoringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ThresholdColoringEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"thresholdColoringEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) Tooltip() ChronicleDashboardChartDashboardChartVisualizationTooltipOutputReference {
	var returns ChronicleDashboardChartDashboardChartVisualizationTooltipOutputReference
	_jsii_.Get(
		j,
		"tooltip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) TooltipInput() *ChronicleDashboardChartDashboardChartVisualizationTooltip {
	var returns *ChronicleDashboardChartDashboardChartVisualizationTooltip
	_jsii_.Get(
		j,
		"tooltipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) VisualMaps() ChronicleDashboardChartDashboardChartVisualizationVisualMapsList {
	var returns ChronicleDashboardChartDashboardChartVisualizationVisualMapsList
	_jsii_.Get(
		j,
		"visualMaps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) VisualMapsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"visualMapsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) XAxes() ChronicleDashboardChartDashboardChartVisualizationXAxesList {
	var returns ChronicleDashboardChartDashboardChartVisualizationXAxesList
	_jsii_.Get(
		j,
		"xAxes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) XAxesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"xAxesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) YAxes() ChronicleDashboardChartDashboardChartVisualizationYAxesList {
	var returns ChronicleDashboardChartDashboardChartVisualizationYAxesList
	_jsii_.Get(
		j,
		"yAxes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) YAxesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"yAxesInput",
		&returns,
	)
	return returns
}

func NewChronicleDashboardChartDashboardChartVisualizationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ChronicleDashboardChartDashboardChartVisualizationOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleDashboardChartDashboardChartVisualizationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.chronicleDashboardChart.ChronicleDashboardChartDashboardChartVisualizationOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleDashboardChartDashboardChartVisualizationOutputReference_Override(c ChronicleDashboardChartDashboardChartVisualizationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.chronicleDashboardChart.ChronicleDashboardChartDashboardChartVisualizationOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) SetGroupingType(val *string) {
	if err := j.validateSetGroupingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupingType",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) SetInternalValue(val *ChronicleDashboardChartDashboardChartVisualization) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) SetSeriesColumn(val *[]*string) {
	if err := j.validateSetSeriesColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"seriesColumn",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) SetThresholdColoringEnabled(val any) {
	if err := j.validateSetThresholdColoringEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdColoringEnabled",
		val,
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) PutButton(value *ChronicleDashboardChartDashboardChartVisualizationButton) {
	if err := c.validatePutButtonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putButton",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) PutColumnDefs(value any) {
	if err := c.validatePutColumnDefsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putColumnDefs",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) PutGoogleMapsConfig(value *ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfig) {
	if err := c.validatePutGoogleMapsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGoogleMapsConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) PutLegends(value any) {
	if err := c.validatePutLegendsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLegends",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) PutMarkdown(value *ChronicleDashboardChartDashboardChartVisualizationMarkdown) {
	if err := c.validatePutMarkdownParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMarkdown",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) PutSeries(value any) {
	if err := c.validatePutSeriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSeries",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) PutTableConfig(value *ChronicleDashboardChartDashboardChartVisualizationTableConfig) {
	if err := c.validatePutTableConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTableConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) PutTooltip(value *ChronicleDashboardChartDashboardChartVisualizationTooltip) {
	if err := c.validatePutTooltipParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTooltip",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) PutVisualMaps(value any) {
	if err := c.validatePutVisualMapsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVisualMaps",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) PutXAxes(value any) {
	if err := c.validatePutXAxesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putXAxes",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) PutYAxes(value any) {
	if err := c.validatePutYAxesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putYAxes",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetButton() {
	_jsii_.InvokeVoid(
		c,
		"resetButton",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetColumnDefs() {
	_jsii_.InvokeVoid(
		c,
		"resetColumnDefs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetGoogleMapsConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetGoogleMapsConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetGroupingType() {
	_jsii_.InvokeVoid(
		c,
		"resetGroupingType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetLegends() {
	_jsii_.InvokeVoid(
		c,
		"resetLegends",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetMarkdown() {
	_jsii_.InvokeVoid(
		c,
		"resetMarkdown",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetSeries() {
	_jsii_.InvokeVoid(
		c,
		"resetSeries",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetSeriesColumn() {
	_jsii_.InvokeVoid(
		c,
		"resetSeriesColumn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetTableConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetTableConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetThresholdColoringEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetThresholdColoringEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetTooltip() {
	_jsii_.InvokeVoid(
		c,
		"resetTooltip",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetVisualMaps() {
	_jsii_.InvokeVoid(
		c,
		"resetVisualMaps",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetXAxes() {
	_jsii_.InvokeVoid(
		c,
		"resetXAxes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetYAxes() {
	_jsii_.InvokeVoid(
		c,
		"resetYAxes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

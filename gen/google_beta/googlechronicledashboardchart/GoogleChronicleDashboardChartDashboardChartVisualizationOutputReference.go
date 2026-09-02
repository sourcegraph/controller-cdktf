package googlechronicledashboardchart

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlechronicledashboardchart/internal"
)

type GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference interface {
	cdktn.ComplexObject
	Button() GoogleChronicleDashboardChartDashboardChartVisualizationButtonOutputReference
	ButtonInput() *GoogleChronicleDashboardChartDashboardChartVisualizationButton
	ColumnDefs() GoogleChronicleDashboardChartDashboardChartVisualizationColumnDefsList
	ColumnDefsInput() interface{}
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
	GoogleMapsConfig() GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigOutputReference
	GoogleMapsConfigInput() *GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfig
	GroupingType() *string
	SetGroupingType(val *string)
	GroupingTypeInput() *string
	InternalValue() *GoogleChronicleDashboardChartDashboardChartVisualization
	SetInternalValue(val *GoogleChronicleDashboardChartDashboardChartVisualization)
	Legends() GoogleChronicleDashboardChartDashboardChartVisualizationLegendsList
	LegendsInput() interface{}
	Markdown() GoogleChronicleDashboardChartDashboardChartVisualizationMarkdownOutputReference
	MarkdownInput() *GoogleChronicleDashboardChartDashboardChartVisualizationMarkdown
	Series() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesList
	SeriesColumn() *[]*string
	SetSeriesColumn(val *[]*string)
	SeriesColumnInput() *[]*string
	SeriesInput() interface{}
	TableConfig() GoogleChronicleDashboardChartDashboardChartVisualizationTableConfigOutputReference
	TableConfigInput() *GoogleChronicleDashboardChartDashboardChartVisualizationTableConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThresholdColoringEnabled() interface{}
	SetThresholdColoringEnabled(val interface{})
	ThresholdColoringEnabledInput() interface{}
	Tooltip() GoogleChronicleDashboardChartDashboardChartVisualizationTooltipOutputReference
	TooltipInput() *GoogleChronicleDashboardChartDashboardChartVisualizationTooltip
	VisualMaps() GoogleChronicleDashboardChartDashboardChartVisualizationVisualMapsList
	VisualMapsInput() interface{}
	XAxes() GoogleChronicleDashboardChartDashboardChartVisualizationXAxesList
	XAxesInput() interface{}
	YAxes() GoogleChronicleDashboardChartDashboardChartVisualizationYAxesList
	YAxesInput() interface{}
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
	PutButton(value *GoogleChronicleDashboardChartDashboardChartVisualizationButton)
	PutColumnDefs(value interface{})
	PutGoogleMapsConfig(value *GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfig)
	PutLegends(value interface{})
	PutMarkdown(value *GoogleChronicleDashboardChartDashboardChartVisualizationMarkdown)
	PutSeries(value interface{})
	PutTableConfig(value *GoogleChronicleDashboardChartDashboardChartVisualizationTableConfig)
	PutTooltip(value *GoogleChronicleDashboardChartDashboardChartVisualizationTooltip)
	PutVisualMaps(value interface{})
	PutXAxes(value interface{})
	PutYAxes(value interface{})
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
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference
type jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) Button() GoogleChronicleDashboardChartDashboardChartVisualizationButtonOutputReference {
	var returns GoogleChronicleDashboardChartDashboardChartVisualizationButtonOutputReference
	_jsii_.Get(
		j,
		"button",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ButtonInput() *GoogleChronicleDashboardChartDashboardChartVisualizationButton {
	var returns *GoogleChronicleDashboardChartDashboardChartVisualizationButton
	_jsii_.Get(
		j,
		"buttonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ColumnDefs() GoogleChronicleDashboardChartDashboardChartVisualizationColumnDefsList {
	var returns GoogleChronicleDashboardChartDashboardChartVisualizationColumnDefsList
	_jsii_.Get(
		j,
		"columnDefs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ColumnDefsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnDefsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) GoogleMapsConfig() GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigOutputReference {
	var returns GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigOutputReference
	_jsii_.Get(
		j,
		"googleMapsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) GoogleMapsConfigInput() *GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfig {
	var returns *GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfig
	_jsii_.Get(
		j,
		"googleMapsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) GroupingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) GroupingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) InternalValue() *GoogleChronicleDashboardChartDashboardChartVisualization {
	var returns *GoogleChronicleDashboardChartDashboardChartVisualization
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) Legends() GoogleChronicleDashboardChartDashboardChartVisualizationLegendsList {
	var returns GoogleChronicleDashboardChartDashboardChartVisualizationLegendsList
	_jsii_.Get(
		j,
		"legends",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) LegendsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"legendsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) Markdown() GoogleChronicleDashboardChartDashboardChartVisualizationMarkdownOutputReference {
	var returns GoogleChronicleDashboardChartDashboardChartVisualizationMarkdownOutputReference
	_jsii_.Get(
		j,
		"markdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) MarkdownInput() *GoogleChronicleDashboardChartDashboardChartVisualizationMarkdown {
	var returns *GoogleChronicleDashboardChartDashboardChartVisualizationMarkdown
	_jsii_.Get(
		j,
		"markdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) Series() GoogleChronicleDashboardChartDashboardChartVisualizationSeriesList {
	var returns GoogleChronicleDashboardChartDashboardChartVisualizationSeriesList
	_jsii_.Get(
		j,
		"series",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) SeriesColumn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"seriesColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) SeriesColumnInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"seriesColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) SeriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"seriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) TableConfig() GoogleChronicleDashboardChartDashboardChartVisualizationTableConfigOutputReference {
	var returns GoogleChronicleDashboardChartDashboardChartVisualizationTableConfigOutputReference
	_jsii_.Get(
		j,
		"tableConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) TableConfigInput() *GoogleChronicleDashboardChartDashboardChartVisualizationTableConfig {
	var returns *GoogleChronicleDashboardChartDashboardChartVisualizationTableConfig
	_jsii_.Get(
		j,
		"tableConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ThresholdColoringEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thresholdColoringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ThresholdColoringEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thresholdColoringEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) Tooltip() GoogleChronicleDashboardChartDashboardChartVisualizationTooltipOutputReference {
	var returns GoogleChronicleDashboardChartDashboardChartVisualizationTooltipOutputReference
	_jsii_.Get(
		j,
		"tooltip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) TooltipInput() *GoogleChronicleDashboardChartDashboardChartVisualizationTooltip {
	var returns *GoogleChronicleDashboardChartDashboardChartVisualizationTooltip
	_jsii_.Get(
		j,
		"tooltipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) VisualMaps() GoogleChronicleDashboardChartDashboardChartVisualizationVisualMapsList {
	var returns GoogleChronicleDashboardChartDashboardChartVisualizationVisualMapsList
	_jsii_.Get(
		j,
		"visualMaps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) VisualMapsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visualMapsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) XAxes() GoogleChronicleDashboardChartDashboardChartVisualizationXAxesList {
	var returns GoogleChronicleDashboardChartDashboardChartVisualizationXAxesList
	_jsii_.Get(
		j,
		"xAxes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) XAxesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xAxesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) YAxes() GoogleChronicleDashboardChartDashboardChartVisualizationYAxesList {
	var returns GoogleChronicleDashboardChartDashboardChartVisualizationYAxesList
	_jsii_.Get(
		j,
		"yAxes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) YAxesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"yAxesInput",
		&returns,
	)
	return returns
}


func NewGoogleChronicleDashboardChartDashboardChartVisualizationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleDashboardChartDashboardChartVisualizationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleDashboardChart.GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleChronicleDashboardChartDashboardChartVisualizationOutputReference_Override(g GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleDashboardChart.GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference)SetGroupingType(val *string) {
	if err := j.validateSetGroupingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupingType",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference)SetInternalValue(val *GoogleChronicleDashboardChartDashboardChartVisualization) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference)SetSeriesColumn(val *[]*string) {
	if err := j.validateSetSeriesColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"seriesColumn",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference)SetThresholdColoringEnabled(val interface{}) {
	if err := j.validateSetThresholdColoringEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdColoringEnabled",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) PutButton(value *GoogleChronicleDashboardChartDashboardChartVisualizationButton) {
	if err := g.validatePutButtonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putButton",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) PutColumnDefs(value interface{}) {
	if err := g.validatePutColumnDefsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putColumnDefs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) PutGoogleMapsConfig(value *GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfig) {
	if err := g.validatePutGoogleMapsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGoogleMapsConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) PutLegends(value interface{}) {
	if err := g.validatePutLegendsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLegends",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) PutMarkdown(value *GoogleChronicleDashboardChartDashboardChartVisualizationMarkdown) {
	if err := g.validatePutMarkdownParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMarkdown",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) PutSeries(value interface{}) {
	if err := g.validatePutSeriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSeries",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) PutTableConfig(value *GoogleChronicleDashboardChartDashboardChartVisualizationTableConfig) {
	if err := g.validatePutTableConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTableConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) PutTooltip(value *GoogleChronicleDashboardChartDashboardChartVisualizationTooltip) {
	if err := g.validatePutTooltipParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTooltip",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) PutVisualMaps(value interface{}) {
	if err := g.validatePutVisualMapsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVisualMaps",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) PutXAxes(value interface{}) {
	if err := g.validatePutXAxesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putXAxes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) PutYAxes(value interface{}) {
	if err := g.validatePutYAxesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putYAxes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetButton() {
	_jsii_.InvokeVoid(
		g,
		"resetButton",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetColumnDefs() {
	_jsii_.InvokeVoid(
		g,
		"resetColumnDefs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetGoogleMapsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGoogleMapsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetGroupingType() {
	_jsii_.InvokeVoid(
		g,
		"resetGroupingType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetLegends() {
	_jsii_.InvokeVoid(
		g,
		"resetLegends",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetMarkdown() {
	_jsii_.InvokeVoid(
		g,
		"resetMarkdown",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetSeries() {
	_jsii_.InvokeVoid(
		g,
		"resetSeries",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetSeriesColumn() {
	_jsii_.InvokeVoid(
		g,
		"resetSeriesColumn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetTableConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetTableConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetThresholdColoringEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetThresholdColoringEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetTooltip() {
	_jsii_.InvokeVoid(
		g,
		"resetTooltip",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetVisualMaps() {
	_jsii_.InvokeVoid(
		g,
		"resetVisualMaps",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetXAxes() {
	_jsii_.InvokeVoid(
		g,
		"resetXAxes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ResetYAxes() {
	_jsii_.InvokeVoid(
		g,
		"resetYAxes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package googlechronicledashboardchart

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlechronicledashboardchart/internal"
)

type GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference interface {
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
	FitData() interface{}
	SetFitData(val interface{})
	FitDataInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPosition
	SetInternalValue(val *GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPosition)
	LatitudeValue() *float64
	SetLatitudeValue(val *float64)
	LatitudeValueInput() *float64
	LongitudeValue() *float64
	SetLongitudeValue(val *float64)
	LongitudeValueInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ZoomScaleValue() *float64
	SetZoomScaleValue(val *float64)
	ZoomScaleValueInput() *float64
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
	ResetFitData()
	ResetLatitudeValue()
	ResetLongitudeValue()
	ResetZoomScaleValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference
type jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) FitData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fitData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) FitDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fitDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) InternalValue() *GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPosition {
	var returns *GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPosition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) LatitudeValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latitudeValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) LatitudeValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latitudeValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) LongitudeValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longitudeValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) LongitudeValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longitudeValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) ZoomScaleValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"zoomScaleValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) ZoomScaleValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"zoomScaleValueInput",
		&returns,
	)
	return returns
}


func NewGoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleDashboardChart.GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference_Override(g GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleDashboardChart.GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference)SetFitData(val interface{}) {
	if err := j.validateSetFitDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fitData",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference)SetInternalValue(val *GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPosition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference)SetLatitudeValue(val *float64) {
	if err := j.validateSetLatitudeValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latitudeValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference)SetLongitudeValue(val *float64) {
	if err := j.validateSetLongitudeValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longitudeValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference)SetZoomScaleValue(val *float64) {
	if err := j.validateSetZoomScaleValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoomScaleValue",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) ResetFitData() {
	_jsii_.InvokeVoid(
		g,
		"resetFitData",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) ResetLatitudeValue() {
	_jsii_.InvokeVoid(
		g,
		"resetLatitudeValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) ResetLongitudeValue() {
	_jsii_.InvokeVoid(
		g,
		"resetLongitudeValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) ResetZoomScaleValue() {
	_jsii_.InvokeVoid(
		g,
		"resetZoomScaleValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


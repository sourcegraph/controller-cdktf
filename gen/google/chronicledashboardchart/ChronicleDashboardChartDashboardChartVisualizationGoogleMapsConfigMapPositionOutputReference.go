package chronicledashboardchart

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/chronicledashboardchart/internal"
)

type ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference interface {
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
	InternalValue() *ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPosition
	SetInternalValue(val *ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPosition)
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

// The jsii proxy struct for ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference
type jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) FitData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fitData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) FitDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fitDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) InternalValue() *ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPosition {
	var returns *ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPosition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) LatitudeValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latitudeValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) LatitudeValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latitudeValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) LongitudeValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longitudeValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) LongitudeValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longitudeValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) ZoomScaleValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"zoomScaleValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) ZoomScaleValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"zoomScaleValueInput",
		&returns,
	)
	return returns
}


func NewChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleDashboardChart.ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference_Override(c ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleDashboardChart.ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference)SetFitData(val interface{}) {
	if err := j.validateSetFitDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fitData",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference)SetInternalValue(val *ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPosition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference)SetLatitudeValue(val *float64) {
	if err := j.validateSetLatitudeValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latitudeValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference)SetLongitudeValue(val *float64) {
	if err := j.validateSetLongitudeValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longitudeValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference)SetZoomScaleValue(val *float64) {
	if err := j.validateSetZoomScaleValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoomScaleValue",
		val,
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) ResetFitData() {
	_jsii_.InvokeVoid(
		c,
		"resetFitData",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) ResetLatitudeValue() {
	_jsii_.InvokeVoid(
		c,
		"resetLatitudeValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) ResetLongitudeValue() {
	_jsii_.InvokeVoid(
		c,
		"resetLongitudeValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) ResetZoomScaleValue() {
	_jsii_.InvokeVoid(
		c,
		"resetZoomScaleValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPositionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


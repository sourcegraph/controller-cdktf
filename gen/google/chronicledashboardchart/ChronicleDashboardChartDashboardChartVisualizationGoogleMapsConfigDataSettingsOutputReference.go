package chronicledashboardchart

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/chronicledashboardchart/internal"
)

type ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference interface {
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
	CountColumn() *string
	SetCountColumn(val *string)
	CountColumnInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettings
	SetInternalValue(val *ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettings)
	LatitudeColumn() *string
	SetLatitudeColumn(val *string)
	LatitudeColumnInput() *string
	LongitudeColumn() *string
	SetLongitudeColumn(val *string)
	LongitudeColumnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	ResetCountColumn()
	ResetLatitudeColumn()
	ResetLongitudeColumn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference
type jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) CountColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) CountColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) InternalValue() *ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettings {
	var returns *ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) LatitudeColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latitudeColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) LatitudeColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latitudeColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) LongitudeColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"longitudeColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) LongitudeColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"longitudeColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleDashboardChart.ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference_Override(c ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleDashboardChart.ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference)SetCountColumn(val *string) {
	if err := j.validateSetCountColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"countColumn",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference)SetInternalValue(val *ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference)SetLatitudeColumn(val *string) {
	if err := j.validateSetLatitudeColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latitudeColumn",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference)SetLongitudeColumn(val *string) {
	if err := j.validateSetLongitudeColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longitudeColumn",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) ResetCountColumn() {
	_jsii_.InvokeVoid(
		c,
		"resetCountColumn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) ResetLatitudeColumn() {
	_jsii_.InvokeVoid(
		c,
		"resetLatitudeColumn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) ResetLongitudeColumn() {
	_jsii_.InvokeVoid(
		c,
		"resetLongitudeColumn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


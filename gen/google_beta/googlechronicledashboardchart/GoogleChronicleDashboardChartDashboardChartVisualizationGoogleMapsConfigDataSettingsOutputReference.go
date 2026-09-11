package googlechronicledashboardchart

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlechronicledashboardchart/internal"
)

type GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference interface {
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
	InternalValue() *GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettings
	SetInternalValue(val *GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettings)
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

// The jsii proxy struct for GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference
type jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) CountColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) CountColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) InternalValue() *GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettings {
	var returns *GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) LatitudeColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latitudeColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) LatitudeColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latitudeColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) LongitudeColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"longitudeColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) LongitudeColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"longitudeColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleDashboardChart.GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference_Override(g GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleDashboardChart.GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference)SetCountColumn(val *string) {
	if err := j.validateSetCountColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"countColumn",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference)SetInternalValue(val *GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference)SetLatitudeColumn(val *string) {
	if err := j.validateSetLatitudeColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latitudeColumn",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference)SetLongitudeColumn(val *string) {
	if err := j.validateSetLongitudeColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longitudeColumn",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) ResetCountColumn() {
	_jsii_.InvokeVoid(
		g,
		"resetCountColumn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) ResetLatitudeColumn() {
	_jsii_.InvokeVoid(
		g,
		"resetLatitudeColumn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) ResetLongitudeColumn() {
	_jsii_.InvokeVoid(
		g,
		"resetLongitudeColumn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package googlechronicledashboardchart

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlechronicledashboardchart/internal"
)

type GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference interface {
	cdktn.ComplexObject
	Color() *string
	SetColor(val *string)
	ColorInput() *string
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
	InternalValue() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyle
	SetInternalValue(val *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyle)
	Opacity() *float64
	SetOpacity(val *float64)
	OpacityInput() *float64
	Origin() *string
	SetOrigin(val *string)
	OriginInput() *string
	ShadowBlur() *float64
	SetShadowBlur(val *float64)
	ShadowBlurInput() *float64
	ShadowColor() *string
	SetShadowColor(val *string)
	ShadowColorInput() *string
	ShadowOffsetX() *float64
	SetShadowOffsetX(val *float64)
	ShadowOffsetXInput() *float64
	ShadowOffsetY() *float64
	SetShadowOffsetY(val *float64)
	ShadowOffsetYInput() *float64
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
	ResetColor()
	ResetOpacity()
	ResetOrigin()
	ResetShadowBlur()
	ResetShadowColor()
	ResetShadowOffsetX()
	ResetShadowOffsetY()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference
type jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) Color() *string {
	var returns *string
	_jsii_.Get(
		j,
		"color",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) ColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) InternalValue() *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyle {
	var returns *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyle
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) Opacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"opacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) OpacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"opacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) Origin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) OriginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) ShadowBlur() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowBlur",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) ShadowBlurInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowBlurInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) ShadowColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shadowColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) ShadowColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shadowColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) ShadowOffsetX() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowOffsetX",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) ShadowOffsetXInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowOffsetXInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) ShadowOffsetY() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowOffsetY",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) ShadowOffsetYInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowOffsetYInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleDashboardChart.GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference_Override(g GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleDashboardChart.GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference)SetColor(val *string) {
	if err := j.validateSetColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"color",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference)SetInternalValue(val *GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyle) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference)SetOpacity(val *float64) {
	if err := j.validateSetOpacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"opacity",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference)SetOrigin(val *string) {
	if err := j.validateSetOriginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"origin",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference)SetShadowBlur(val *float64) {
	if err := j.validateSetShadowBlurParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shadowBlur",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference)SetShadowColor(val *string) {
	if err := j.validateSetShadowColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shadowColor",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference)SetShadowOffsetX(val *float64) {
	if err := j.validateSetShadowOffsetXParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shadowOffsetX",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference)SetShadowOffsetY(val *float64) {
	if err := j.validateSetShadowOffsetYParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shadowOffsetY",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) ResetColor() {
	_jsii_.InvokeVoid(
		g,
		"resetColor",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) ResetOpacity() {
	_jsii_.InvokeVoid(
		g,
		"resetOpacity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) ResetOrigin() {
	_jsii_.InvokeVoid(
		g,
		"resetOrigin",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) ResetShadowBlur() {
	_jsii_.InvokeVoid(
		g,
		"resetShadowBlur",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) ResetShadowColor() {
	_jsii_.InvokeVoid(
		g,
		"resetShadowColor",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) ResetShadowOffsetX() {
	_jsii_.InvokeVoid(
		g,
		"resetShadowOffsetX",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) ResetShadowOffsetY() {
	_jsii_.InvokeVoid(
		g,
		"resetShadowOffsetY",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartDashboardChartVisualizationSeriesAreaStyleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


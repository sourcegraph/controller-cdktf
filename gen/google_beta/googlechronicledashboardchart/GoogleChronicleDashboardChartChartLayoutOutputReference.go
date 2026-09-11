package googlechronicledashboardchart

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlechronicledashboardchart/internal"
)

type GoogleChronicleDashboardChartChartLayoutOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleChronicleDashboardChartChartLayout
	SetInternalValue(val *GoogleChronicleDashboardChartChartLayout)
	SpanX() *float64
	SetSpanX(val *float64)
	SpanXInput() *float64
	SpanY() *float64
	SetSpanY(val *float64)
	SpanYInput() *float64
	StartX() *float64
	SetStartX(val *float64)
	StartXInput() *float64
	StartY() *float64
	SetStartY(val *float64)
	StartYInput() *float64
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
	ResetStartX()
	ResetStartY()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleDashboardChartChartLayoutOutputReference
type jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) InternalValue() *GoogleChronicleDashboardChartChartLayout {
	var returns *GoogleChronicleDashboardChartChartLayout
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) SpanX() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spanX",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) SpanXInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spanXInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) SpanY() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spanY",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) SpanYInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spanYInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) StartX() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startX",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) StartXInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startXInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) StartY() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startY",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) StartYInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startYInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleChronicleDashboardChartChartLayoutOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleChronicleDashboardChartChartLayoutOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleDashboardChartChartLayoutOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleDashboardChart.GoogleChronicleDashboardChartChartLayoutOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleChronicleDashboardChartChartLayoutOutputReference_Override(g GoogleChronicleDashboardChartChartLayoutOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleDashboardChart.GoogleChronicleDashboardChartChartLayoutOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference)SetInternalValue(val *GoogleChronicleDashboardChartChartLayout) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference)SetSpanX(val *float64) {
	if err := j.validateSetSpanXParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spanX",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference)SetSpanY(val *float64) {
	if err := j.validateSetSpanYParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spanY",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference)SetStartX(val *float64) {
	if err := j.validateSetStartXParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startX",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference)SetStartY(val *float64) {
	if err := j.validateSetStartYParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startY",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) ResetStartX() {
	_jsii_.InvokeVoid(
		g,
		"resetStartX",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) ResetStartY() {
	_jsii_.InvokeVoid(
		g,
		"resetStartY",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleChronicleDashboardChartChartLayoutOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


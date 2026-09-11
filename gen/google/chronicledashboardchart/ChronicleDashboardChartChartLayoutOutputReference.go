package chronicledashboardchart

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/chronicledashboardchart/internal"
)

type ChronicleDashboardChartChartLayoutOutputReference interface {
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
	InternalValue() *ChronicleDashboardChartChartLayout
	SetInternalValue(val *ChronicleDashboardChartChartLayout)
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

// The jsii proxy struct for ChronicleDashboardChartChartLayoutOutputReference
type jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) InternalValue() *ChronicleDashboardChartChartLayout {
	var returns *ChronicleDashboardChartChartLayout
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) SpanX() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spanX",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) SpanXInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spanXInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) SpanY() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spanY",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) SpanYInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spanYInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) StartX() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startX",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) StartXInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startXInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) StartY() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startY",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) StartYInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startYInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewChronicleDashboardChartChartLayoutOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleDashboardChartChartLayoutOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleDashboardChartChartLayoutOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleDashboardChart.ChronicleDashboardChartChartLayoutOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleDashboardChartChartLayoutOutputReference_Override(c ChronicleDashboardChartChartLayoutOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleDashboardChart.ChronicleDashboardChartChartLayoutOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference)SetInternalValue(val *ChronicleDashboardChartChartLayout) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference)SetSpanX(val *float64) {
	if err := j.validateSetSpanXParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spanX",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference)SetSpanY(val *float64) {
	if err := j.validateSetSpanYParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spanY",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference)SetStartX(val *float64) {
	if err := j.validateSetStartXParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startX",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference)SetStartY(val *float64) {
	if err := j.validateSetStartYParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startY",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) ResetStartX() {
	_jsii_.InvokeVoid(
		c,
		"resetStartX",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) ResetStartY() {
	_jsii_.InvokeVoid(
		c,
		"resetStartY",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleDashboardChartChartLayoutOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


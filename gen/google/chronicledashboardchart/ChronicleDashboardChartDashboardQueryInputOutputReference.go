package chronicledashboardchart

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/chronicledashboardchart/internal"
)

type ChronicleDashboardChartDashboardQueryInputOutputReference interface {
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
	InternalValue() *ChronicleDashboardChartDashboardQueryInput
	SetInternalValue(val *ChronicleDashboardChartDashboardQueryInput)
	RelativeTime() ChronicleDashboardChartDashboardQueryInputRelativeTimeOutputReference
	RelativeTimeInput() *ChronicleDashboardChartDashboardQueryInputRelativeTime
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeWindow() ChronicleDashboardChartDashboardQueryInputTimeWindowOutputReference
	TimeWindowInput() *ChronicleDashboardChartDashboardQueryInputTimeWindow
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
	PutRelativeTime(value *ChronicleDashboardChartDashboardQueryInputRelativeTime)
	PutTimeWindow(value *ChronicleDashboardChartDashboardQueryInputTimeWindow)
	ResetRelativeTime()
	ResetTimeWindow()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleDashboardChartDashboardQueryInputOutputReference
type jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) InternalValue() *ChronicleDashboardChartDashboardQueryInput {
	var returns *ChronicleDashboardChartDashboardQueryInput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) RelativeTime() ChronicleDashboardChartDashboardQueryInputRelativeTimeOutputReference {
	var returns ChronicleDashboardChartDashboardQueryInputRelativeTimeOutputReference
	_jsii_.Get(
		j,
		"relativeTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) RelativeTimeInput() *ChronicleDashboardChartDashboardQueryInputRelativeTime {
	var returns *ChronicleDashboardChartDashboardQueryInputRelativeTime
	_jsii_.Get(
		j,
		"relativeTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) TimeWindow() ChronicleDashboardChartDashboardQueryInputTimeWindowOutputReference {
	var returns ChronicleDashboardChartDashboardQueryInputTimeWindowOutputReference
	_jsii_.Get(
		j,
		"timeWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) TimeWindowInput() *ChronicleDashboardChartDashboardQueryInputTimeWindow {
	var returns *ChronicleDashboardChartDashboardQueryInputTimeWindow
	_jsii_.Get(
		j,
		"timeWindowInput",
		&returns,
	)
	return returns
}


func NewChronicleDashboardChartDashboardQueryInputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleDashboardChartDashboardQueryInputOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleDashboardChartDashboardQueryInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleDashboardChart.ChronicleDashboardChartDashboardQueryInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleDashboardChartDashboardQueryInputOutputReference_Override(c ChronicleDashboardChartDashboardQueryInputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleDashboardChart.ChronicleDashboardChartDashboardQueryInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference)SetInternalValue(val *ChronicleDashboardChartDashboardQueryInput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) PutRelativeTime(value *ChronicleDashboardChartDashboardQueryInputRelativeTime) {
	if err := c.validatePutRelativeTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRelativeTime",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) PutTimeWindow(value *ChronicleDashboardChartDashboardQueryInputTimeWindow) {
	if err := c.validatePutTimeWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeWindow",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) ResetRelativeTime() {
	_jsii_.InvokeVoid(
		c,
		"resetRelativeTime",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) ResetTimeWindow() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeWindow",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardQueryInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


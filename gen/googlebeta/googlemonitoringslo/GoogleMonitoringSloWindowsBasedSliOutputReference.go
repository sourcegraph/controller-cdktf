package googlemonitoringslo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlemonitoringslo/internal"
)

type GoogleMonitoringSloWindowsBasedSliOutputReference interface {
	cdktf.ComplexObject
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
	GoodBadMetricFilter() *string
	SetGoodBadMetricFilter(val *string)
	GoodBadMetricFilterInput() *string
	GoodTotalRatioThreshold() GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference
	GoodTotalRatioThresholdInput() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThreshold
	InternalValue() *GoogleMonitoringSloWindowsBasedSli
	SetInternalValue(val *GoogleMonitoringSloWindowsBasedSli)
	MetricMeanInRange() GoogleMonitoringSloWindowsBasedSliMetricMeanInRangeOutputReference
	MetricMeanInRangeInput() *GoogleMonitoringSloWindowsBasedSliMetricMeanInRange
	MetricSumInRange() GoogleMonitoringSloWindowsBasedSliMetricSumInRangeOutputReference
	MetricSumInRangeInput() *GoogleMonitoringSloWindowsBasedSliMetricSumInRange
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WindowPeriod() *string
	SetWindowPeriod(val *string)
	WindowPeriodInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
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
	PutGoodTotalRatioThreshold(value *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThreshold)
	PutMetricMeanInRange(value *GoogleMonitoringSloWindowsBasedSliMetricMeanInRange)
	PutMetricSumInRange(value *GoogleMonitoringSloWindowsBasedSliMetricSumInRange)
	ResetGoodBadMetricFilter()
	ResetGoodTotalRatioThreshold()
	ResetMetricMeanInRange()
	ResetMetricSumInRange()
	ResetWindowPeriod()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleMonitoringSloWindowsBasedSliOutputReference
type jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) GoodBadMetricFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"goodBadMetricFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) GoodBadMetricFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"goodBadMetricFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) GoodTotalRatioThreshold() GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference {
	var returns GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference
	_jsii_.Get(
		j,
		"goodTotalRatioThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) GoodTotalRatioThresholdInput() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThreshold {
	var returns *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThreshold
	_jsii_.Get(
		j,
		"goodTotalRatioThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) InternalValue() *GoogleMonitoringSloWindowsBasedSli {
	var returns *GoogleMonitoringSloWindowsBasedSli
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) MetricMeanInRange() GoogleMonitoringSloWindowsBasedSliMetricMeanInRangeOutputReference {
	var returns GoogleMonitoringSloWindowsBasedSliMetricMeanInRangeOutputReference
	_jsii_.Get(
		j,
		"metricMeanInRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) MetricMeanInRangeInput() *GoogleMonitoringSloWindowsBasedSliMetricMeanInRange {
	var returns *GoogleMonitoringSloWindowsBasedSliMetricMeanInRange
	_jsii_.Get(
		j,
		"metricMeanInRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) MetricSumInRange() GoogleMonitoringSloWindowsBasedSliMetricSumInRangeOutputReference {
	var returns GoogleMonitoringSloWindowsBasedSliMetricSumInRangeOutputReference
	_jsii_.Get(
		j,
		"metricSumInRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) MetricSumInRangeInput() *GoogleMonitoringSloWindowsBasedSliMetricSumInRange {
	var returns *GoogleMonitoringSloWindowsBasedSliMetricSumInRange
	_jsii_.Get(
		j,
		"metricSumInRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) WindowPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) WindowPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowPeriodInput",
		&returns,
	)
	return returns
}


func NewGoogleMonitoringSloWindowsBasedSliOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleMonitoringSloWindowsBasedSliOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleMonitoringSloWindowsBasedSliOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleMonitoringSlo.GoogleMonitoringSloWindowsBasedSliOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleMonitoringSloWindowsBasedSliOutputReference_Override(g GoogleMonitoringSloWindowsBasedSliOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleMonitoringSlo.GoogleMonitoringSloWindowsBasedSliOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference)SetGoodBadMetricFilter(val *string) {
	if err := j.validateSetGoodBadMetricFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"goodBadMetricFilter",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference)SetInternalValue(val *GoogleMonitoringSloWindowsBasedSli) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference)SetWindowPeriod(val *string) {
	if err := j.validateSetWindowPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"windowPeriod",
		val,
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) PutGoodTotalRatioThreshold(value *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThreshold) {
	if err := g.validatePutGoodTotalRatioThresholdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGoodTotalRatioThreshold",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) PutMetricMeanInRange(value *GoogleMonitoringSloWindowsBasedSliMetricMeanInRange) {
	if err := g.validatePutMetricMeanInRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMetricMeanInRange",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) PutMetricSumInRange(value *GoogleMonitoringSloWindowsBasedSliMetricSumInRange) {
	if err := g.validatePutMetricSumInRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMetricSumInRange",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) ResetGoodBadMetricFilter() {
	_jsii_.InvokeVoid(
		g,
		"resetGoodBadMetricFilter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) ResetGoodTotalRatioThreshold() {
	_jsii_.InvokeVoid(
		g,
		"resetGoodTotalRatioThreshold",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) ResetMetricMeanInRange() {
	_jsii_.InvokeVoid(
		g,
		"resetMetricMeanInRange",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) ResetMetricSumInRange() {
	_jsii_.InvokeVoid(
		g,
		"resetMetricSumInRange",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) ResetWindowPeriod() {
	_jsii_.InvokeVoid(
		g,
		"resetWindowPeriod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


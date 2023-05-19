package googlemonitoringslo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlemonitoringslo/internal"
)

type GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference interface {
	cdktf.ComplexObject
	BadServiceFilter() *string
	SetBadServiceFilter(val *string)
	BadServiceFilterInput() *string
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
	GoodServiceFilter() *string
	SetGoodServiceFilter(val *string)
	GoodServiceFilterInput() *string
	InternalValue() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatio
	SetInternalValue(val *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatio)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TotalServiceFilter() *string
	SetTotalServiceFilter(val *string)
	TotalServiceFilterInput() *string
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
	ResetBadServiceFilter()
	ResetGoodServiceFilter()
	ResetTotalServiceFilter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference
type jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) BadServiceFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"badServiceFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) BadServiceFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"badServiceFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) GoodServiceFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"goodServiceFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) GoodServiceFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"goodServiceFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) InternalValue() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatio {
	var returns *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatio
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) TotalServiceFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalServiceFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) TotalServiceFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalServiceFilterInput",
		&returns,
	)
	return returns
}


func NewGoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference{}

	_jsii_.Create(
		"google-beta.googleMonitoringSlo.GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference_Override(g GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleMonitoringSlo.GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference)SetBadServiceFilter(val *string) {
	if err := j.validateSetBadServiceFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"badServiceFilter",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference)SetGoodServiceFilter(val *string) {
	if err := j.validateSetGoodServiceFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"goodServiceFilter",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference)SetInternalValue(val *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatio) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference)SetTotalServiceFilter(val *string) {
	if err := j.validateSetTotalServiceFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalServiceFilter",
		val,
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) ResetBadServiceFilter() {
	_jsii_.InvokeVoid(
		g,
		"resetBadServiceFilter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) ResetGoodServiceFilter() {
	_jsii_.InvokeVoid(
		g,
		"resetGoodServiceFilter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) ResetTotalServiceFilter() {
	_jsii_.InvokeVoid(
		g,
		"resetTotalServiceFilter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


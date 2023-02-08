package googlemonitoringslo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlemonitoringslo/internal"
)

type GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference interface {
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
	DistributionCut() GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceDistributionCutOutputReference
	DistributionCutInput() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceDistributionCut
	// Experimental.
	Fqn() *string
	GoodTotalRatio() GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference
	GoodTotalRatioInput() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatio
	InternalValue() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformance
	SetInternalValue(val *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformance)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutDistributionCut(value *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceDistributionCut)
	PutGoodTotalRatio(value *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatio)
	ResetDistributionCut()
	ResetGoodTotalRatio()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference
type jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) DistributionCut() GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceDistributionCutOutputReference {
	var returns GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceDistributionCutOutputReference
	_jsii_.Get(
		j,
		"distributionCut",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) DistributionCutInput() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceDistributionCut {
	var returns *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceDistributionCut
	_jsii_.Get(
		j,
		"distributionCutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) GoodTotalRatio() GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference {
	var returns GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatioOutputReference
	_jsii_.Get(
		j,
		"goodTotalRatio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) GoodTotalRatioInput() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatio {
	var returns *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatio
	_jsii_.Get(
		j,
		"goodTotalRatioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) InternalValue() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformance {
	var returns *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformance
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference{}

	_jsii_.Create(
		"google-beta.googleMonitoringSlo.GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference_Override(g GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleMonitoringSlo.GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference)SetInternalValue(val *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformance) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) PutDistributionCut(value *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceDistributionCut) {
	if err := g.validatePutDistributionCutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDistributionCut",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) PutGoodTotalRatio(value *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatio) {
	if err := g.validatePutGoodTotalRatioParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGoodTotalRatio",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) ResetDistributionCut() {
	_jsii_.InvokeVoid(
		g,
		"resetDistributionCut",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) ResetGoodTotalRatio() {
	_jsii_.InvokeVoid(
		g,
		"resetGoodTotalRatio",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


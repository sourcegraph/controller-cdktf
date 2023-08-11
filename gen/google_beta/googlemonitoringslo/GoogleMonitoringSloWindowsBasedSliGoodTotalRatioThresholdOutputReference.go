package googlemonitoringslo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlemonitoringslo/internal"
)

type GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference interface {
	cdktf.ComplexObject
	BasicSliPerformance() GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference
	BasicSliPerformanceInput() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformance
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
	InternalValue() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThreshold
	SetInternalValue(val *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThreshold)
	Performance() GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference
	PerformanceInput() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformance
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Threshold() *float64
	SetThreshold(val *float64)
	ThresholdInput() *float64
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
	PutBasicSliPerformance(value *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformance)
	PutPerformance(value *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformance)
	ResetBasicSliPerformance()
	ResetPerformance()
	ResetThreshold()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference
type jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) BasicSliPerformance() GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference {
	var returns GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference
	_jsii_.Get(
		j,
		"basicSliPerformance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) BasicSliPerformanceInput() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformance {
	var returns *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformance
	_jsii_.Get(
		j,
		"basicSliPerformanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) InternalValue() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThreshold {
	var returns *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThreshold
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) Performance() GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference {
	var returns GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceOutputReference
	_jsii_.Get(
		j,
		"performance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) PerformanceInput() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformance {
	var returns *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformance
	_jsii_.Get(
		j,
		"performanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) Threshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) ThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}


func NewGoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleMonitoringSlo.GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference_Override(g GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleMonitoringSlo.GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference)SetInternalValue(val *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThreshold) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference)SetThreshold(val *float64) {
	if err := j.validateSetThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) PutBasicSliPerformance(value *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformance) {
	if err := g.validatePutBasicSliPerformanceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBasicSliPerformance",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) PutPerformance(value *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformance) {
	if err := g.validatePutPerformanceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPerformance",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) ResetBasicSliPerformance() {
	_jsii_.InvokeVoid(
		g,
		"resetBasicSliPerformance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) ResetPerformance() {
	_jsii_.InvokeVoid(
		g,
		"resetPerformance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) ResetThreshold() {
	_jsii_.InvokeVoid(
		g,
		"resetThreshold",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


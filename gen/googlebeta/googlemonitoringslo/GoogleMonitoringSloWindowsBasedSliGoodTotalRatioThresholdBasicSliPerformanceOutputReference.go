package googlemonitoringslo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlemonitoringslo/internal"
)

type GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference interface {
	cdktf.ComplexObject
	Availability() GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceAvailabilityOutputReference
	AvailabilityInput() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceAvailability
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
	InternalValue() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformance
	SetInternalValue(val *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformance)
	Latency() GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceLatencyOutputReference
	LatencyInput() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceLatency
	Location() *[]*string
	SetLocation(val *[]*string)
	LocationInput() *[]*string
	Method() *[]*string
	SetMethod(val *[]*string)
	MethodInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *[]*string
	SetVersion(val *[]*string)
	VersionInput() *[]*string
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
	PutAvailability(value *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceAvailability)
	PutLatency(value *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceLatency)
	ResetAvailability()
	ResetLatency()
	ResetLocation()
	ResetMethod()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference
type jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) Availability() GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceAvailabilityOutputReference {
	var returns GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceAvailabilityOutputReference
	_jsii_.Get(
		j,
		"availability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) AvailabilityInput() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceAvailability {
	var returns *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceAvailability
	_jsii_.Get(
		j,
		"availabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) InternalValue() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformance {
	var returns *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformance
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) Latency() GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceLatencyOutputReference {
	var returns GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceLatencyOutputReference
	_jsii_.Get(
		j,
		"latency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) LatencyInput() *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceLatency {
	var returns *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceLatency
	_jsii_.Get(
		j,
		"latencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) Location() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) LocationInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) Method() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) MethodInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) Version() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) VersionInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewGoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleMonitoringSlo.GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference_Override(g GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleMonitoringSlo.GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference)SetInternalValue(val *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformance) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference)SetLocation(val *[]*string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference)SetMethod(val *[]*string) {
	if err := j.validateSetMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference)SetVersion(val *[]*string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) PutAvailability(value *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceAvailability) {
	if err := g.validatePutAvailabilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAvailability",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) PutLatency(value *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceLatency) {
	if err := g.validatePutLatencyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLatency",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) ResetAvailability() {
	_jsii_.InvokeVoid(
		g,
		"resetAvailability",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) ResetLatency() {
	_jsii_.InvokeVoid(
		g,
		"resetLatency",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		g,
		"resetMethod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


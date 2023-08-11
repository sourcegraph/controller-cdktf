package googlecomputeregionbackendservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeregionbackendservice/internal"
)

type GoogleComputeRegionBackendServiceOutlierDetectionOutputReference interface {
	cdktf.ComplexObject
	BaseEjectionTime() GoogleComputeRegionBackendServiceOutlierDetectionBaseEjectionTimeOutputReference
	BaseEjectionTimeInput() *GoogleComputeRegionBackendServiceOutlierDetectionBaseEjectionTime
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
	ConsecutiveErrors() *float64
	SetConsecutiveErrors(val *float64)
	ConsecutiveErrorsInput() *float64
	ConsecutiveGatewayFailure() *float64
	SetConsecutiveGatewayFailure(val *float64)
	ConsecutiveGatewayFailureInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnforcingConsecutiveErrors() *float64
	SetEnforcingConsecutiveErrors(val *float64)
	EnforcingConsecutiveErrorsInput() *float64
	EnforcingConsecutiveGatewayFailure() *float64
	SetEnforcingConsecutiveGatewayFailure(val *float64)
	EnforcingConsecutiveGatewayFailureInput() *float64
	EnforcingSuccessRate() *float64
	SetEnforcingSuccessRate(val *float64)
	EnforcingSuccessRateInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeRegionBackendServiceOutlierDetection
	SetInternalValue(val *GoogleComputeRegionBackendServiceOutlierDetection)
	Interval() GoogleComputeRegionBackendServiceOutlierDetectionIntervalOutputReference
	IntervalInput() *GoogleComputeRegionBackendServiceOutlierDetectionInterval
	MaxEjectionPercent() *float64
	SetMaxEjectionPercent(val *float64)
	MaxEjectionPercentInput() *float64
	SuccessRateMinimumHosts() *float64
	SetSuccessRateMinimumHosts(val *float64)
	SuccessRateMinimumHostsInput() *float64
	SuccessRateRequestVolume() *float64
	SetSuccessRateRequestVolume(val *float64)
	SuccessRateRequestVolumeInput() *float64
	SuccessRateStdevFactor() *float64
	SetSuccessRateStdevFactor(val *float64)
	SuccessRateStdevFactorInput() *float64
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
	PutBaseEjectionTime(value *GoogleComputeRegionBackendServiceOutlierDetectionBaseEjectionTime)
	PutInterval(value *GoogleComputeRegionBackendServiceOutlierDetectionInterval)
	ResetBaseEjectionTime()
	ResetConsecutiveErrors()
	ResetConsecutiveGatewayFailure()
	ResetEnforcingConsecutiveErrors()
	ResetEnforcingConsecutiveGatewayFailure()
	ResetEnforcingSuccessRate()
	ResetInterval()
	ResetMaxEjectionPercent()
	ResetSuccessRateMinimumHosts()
	ResetSuccessRateRequestVolume()
	ResetSuccessRateStdevFactor()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRegionBackendServiceOutlierDetectionOutputReference
type jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) BaseEjectionTime() GoogleComputeRegionBackendServiceOutlierDetectionBaseEjectionTimeOutputReference {
	var returns GoogleComputeRegionBackendServiceOutlierDetectionBaseEjectionTimeOutputReference
	_jsii_.Get(
		j,
		"baseEjectionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) BaseEjectionTimeInput() *GoogleComputeRegionBackendServiceOutlierDetectionBaseEjectionTime {
	var returns *GoogleComputeRegionBackendServiceOutlierDetectionBaseEjectionTime
	_jsii_.Get(
		j,
		"baseEjectionTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ConsecutiveErrors() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"consecutiveErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ConsecutiveErrorsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"consecutiveErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ConsecutiveGatewayFailure() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"consecutiveGatewayFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ConsecutiveGatewayFailureInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"consecutiveGatewayFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) EnforcingConsecutiveErrors() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enforcingConsecutiveErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) EnforcingConsecutiveErrorsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enforcingConsecutiveErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) EnforcingConsecutiveGatewayFailure() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enforcingConsecutiveGatewayFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) EnforcingConsecutiveGatewayFailureInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enforcingConsecutiveGatewayFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) EnforcingSuccessRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enforcingSuccessRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) EnforcingSuccessRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enforcingSuccessRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) InternalValue() *GoogleComputeRegionBackendServiceOutlierDetection {
	var returns *GoogleComputeRegionBackendServiceOutlierDetection
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) Interval() GoogleComputeRegionBackendServiceOutlierDetectionIntervalOutputReference {
	var returns GoogleComputeRegionBackendServiceOutlierDetectionIntervalOutputReference
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) IntervalInput() *GoogleComputeRegionBackendServiceOutlierDetectionInterval {
	var returns *GoogleComputeRegionBackendServiceOutlierDetectionInterval
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) MaxEjectionPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEjectionPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) MaxEjectionPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEjectionPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) SuccessRateMinimumHosts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRateMinimumHosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) SuccessRateMinimumHostsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRateMinimumHostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) SuccessRateRequestVolume() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRateRequestVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) SuccessRateRequestVolumeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRateRequestVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) SuccessRateStdevFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRateStdevFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) SuccessRateStdevFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRateStdevFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionBackendServiceOutlierDetectionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeRegionBackendServiceOutlierDetectionOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionBackendServiceOutlierDetectionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeRegionBackendService.GoogleComputeRegionBackendServiceOutlierDetectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionBackendServiceOutlierDetectionOutputReference_Override(g GoogleComputeRegionBackendServiceOutlierDetectionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeRegionBackendService.GoogleComputeRegionBackendServiceOutlierDetectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference)SetConsecutiveErrors(val *float64) {
	if err := j.validateSetConsecutiveErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consecutiveErrors",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference)SetConsecutiveGatewayFailure(val *float64) {
	if err := j.validateSetConsecutiveGatewayFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consecutiveGatewayFailure",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference)SetEnforcingConsecutiveErrors(val *float64) {
	if err := j.validateSetEnforcingConsecutiveErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforcingConsecutiveErrors",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference)SetEnforcingConsecutiveGatewayFailure(val *float64) {
	if err := j.validateSetEnforcingConsecutiveGatewayFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforcingConsecutiveGatewayFailure",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference)SetEnforcingSuccessRate(val *float64) {
	if err := j.validateSetEnforcingSuccessRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforcingSuccessRate",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference)SetInternalValue(val *GoogleComputeRegionBackendServiceOutlierDetection) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference)SetMaxEjectionPercent(val *float64) {
	if err := j.validateSetMaxEjectionPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxEjectionPercent",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference)SetSuccessRateMinimumHosts(val *float64) {
	if err := j.validateSetSuccessRateMinimumHostsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successRateMinimumHosts",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference)SetSuccessRateRequestVolume(val *float64) {
	if err := j.validateSetSuccessRateRequestVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successRateRequestVolume",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference)SetSuccessRateStdevFactor(val *float64) {
	if err := j.validateSetSuccessRateStdevFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successRateStdevFactor",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) PutBaseEjectionTime(value *GoogleComputeRegionBackendServiceOutlierDetectionBaseEjectionTime) {
	if err := g.validatePutBaseEjectionTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBaseEjectionTime",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) PutInterval(value *GoogleComputeRegionBackendServiceOutlierDetectionInterval) {
	if err := g.validatePutIntervalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInterval",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ResetBaseEjectionTime() {
	_jsii_.InvokeVoid(
		g,
		"resetBaseEjectionTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ResetConsecutiveErrors() {
	_jsii_.InvokeVoid(
		g,
		"resetConsecutiveErrors",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ResetConsecutiveGatewayFailure() {
	_jsii_.InvokeVoid(
		g,
		"resetConsecutiveGatewayFailure",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ResetEnforcingConsecutiveErrors() {
	_jsii_.InvokeVoid(
		g,
		"resetEnforcingConsecutiveErrors",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ResetEnforcingConsecutiveGatewayFailure() {
	_jsii_.InvokeVoid(
		g,
		"resetEnforcingConsecutiveGatewayFailure",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ResetEnforcingSuccessRate() {
	_jsii_.InvokeVoid(
		g,
		"resetEnforcingSuccessRate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ResetInterval() {
	_jsii_.InvokeVoid(
		g,
		"resetInterval",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ResetMaxEjectionPercent() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxEjectionPercent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ResetSuccessRateMinimumHosts() {
	_jsii_.InvokeVoid(
		g,
		"resetSuccessRateMinimumHosts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ResetSuccessRateRequestVolume() {
	_jsii_.InvokeVoid(
		g,
		"resetSuccessRateRequestVolume",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ResetSuccessRateStdevFactor() {
	_jsii_.InvokeVoid(
		g,
		"resetSuccessRateStdevFactor",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


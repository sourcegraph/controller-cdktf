package googleappengineflexibleappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleappengineflexibleappversion/internal"
)

type GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference interface {
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
	CoolDownPeriod() *string
	SetCoolDownPeriod(val *string)
	CoolDownPeriodInput() *string
	CpuUtilization() GoogleAppEngineFlexibleAppVersionAutomaticScalingCpuUtilizationOutputReference
	CpuUtilizationInput() *GoogleAppEngineFlexibleAppVersionAutomaticScalingCpuUtilization
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DiskUtilization() GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference
	DiskUtilizationInput() *GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilization
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleAppEngineFlexibleAppVersionAutomaticScaling
	SetInternalValue(val *GoogleAppEngineFlexibleAppVersionAutomaticScaling)
	MaxConcurrentRequests() *float64
	SetMaxConcurrentRequests(val *float64)
	MaxConcurrentRequestsInput() *float64
	MaxIdleInstances() *float64
	SetMaxIdleInstances(val *float64)
	MaxIdleInstancesInput() *float64
	MaxPendingLatency() *string
	SetMaxPendingLatency(val *string)
	MaxPendingLatencyInput() *string
	MaxTotalInstances() *float64
	SetMaxTotalInstances(val *float64)
	MaxTotalInstancesInput() *float64
	MinIdleInstances() *float64
	SetMinIdleInstances(val *float64)
	MinIdleInstancesInput() *float64
	MinPendingLatency() *string
	SetMinPendingLatency(val *string)
	MinPendingLatencyInput() *string
	MinTotalInstances() *float64
	SetMinTotalInstances(val *float64)
	MinTotalInstancesInput() *float64
	NetworkUtilization() GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference
	NetworkUtilizationInput() *GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization
	RequestUtilization() GoogleAppEngineFlexibleAppVersionAutomaticScalingRequestUtilizationOutputReference
	RequestUtilizationInput() *GoogleAppEngineFlexibleAppVersionAutomaticScalingRequestUtilization
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
	PutCpuUtilization(value *GoogleAppEngineFlexibleAppVersionAutomaticScalingCpuUtilization)
	PutDiskUtilization(value *GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilization)
	PutNetworkUtilization(value *GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization)
	PutRequestUtilization(value *GoogleAppEngineFlexibleAppVersionAutomaticScalingRequestUtilization)
	ResetCoolDownPeriod()
	ResetDiskUtilization()
	ResetMaxConcurrentRequests()
	ResetMaxIdleInstances()
	ResetMaxPendingLatency()
	ResetMaxTotalInstances()
	ResetMinIdleInstances()
	ResetMinPendingLatency()
	ResetMinTotalInstances()
	ResetNetworkUtilization()
	ResetRequestUtilization()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference
type jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) CoolDownPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coolDownPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) CoolDownPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coolDownPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) CpuUtilization() GoogleAppEngineFlexibleAppVersionAutomaticScalingCpuUtilizationOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionAutomaticScalingCpuUtilizationOutputReference
	_jsii_.Get(
		j,
		"cpuUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) CpuUtilizationInput() *GoogleAppEngineFlexibleAppVersionAutomaticScalingCpuUtilization {
	var returns *GoogleAppEngineFlexibleAppVersionAutomaticScalingCpuUtilization
	_jsii_.Get(
		j,
		"cpuUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) DiskUtilization() GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference
	_jsii_.Get(
		j,
		"diskUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) DiskUtilizationInput() *GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilization {
	var returns *GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilization
	_jsii_.Get(
		j,
		"diskUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) InternalValue() *GoogleAppEngineFlexibleAppVersionAutomaticScaling {
	var returns *GoogleAppEngineFlexibleAppVersionAutomaticScaling
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) MaxConcurrentRequests() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) MaxConcurrentRequestsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentRequestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) MaxIdleInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) MaxIdleInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) MaxPendingLatency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxPendingLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) MaxPendingLatencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxPendingLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) MaxTotalInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTotalInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) MaxTotalInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTotalInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) MinIdleInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIdleInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) MinIdleInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIdleInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) MinPendingLatency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minPendingLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) MinPendingLatencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minPendingLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) MinTotalInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTotalInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) MinTotalInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTotalInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) NetworkUtilization() GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference
	_jsii_.Get(
		j,
		"networkUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) NetworkUtilizationInput() *GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization {
	var returns *GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization
	_jsii_.Get(
		j,
		"networkUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) RequestUtilization() GoogleAppEngineFlexibleAppVersionAutomaticScalingRequestUtilizationOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionAutomaticScalingRequestUtilizationOutputReference
	_jsii_.Get(
		j,
		"requestUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) RequestUtilizationInput() *GoogleAppEngineFlexibleAppVersionAutomaticScalingRequestUtilization {
	var returns *GoogleAppEngineFlexibleAppVersionAutomaticScalingRequestUtilization
	_jsii_.Get(
		j,
		"requestUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleAppEngineFlexibleAppVersion.GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference_Override(g GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleAppEngineFlexibleAppVersion.GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetCoolDownPeriod(val *string) {
	if err := j.validateSetCoolDownPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coolDownPeriod",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetInternalValue(val *GoogleAppEngineFlexibleAppVersionAutomaticScaling) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetMaxConcurrentRequests(val *float64) {
	if err := j.validateSetMaxConcurrentRequestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentRequests",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetMaxIdleInstances(val *float64) {
	if err := j.validateSetMaxIdleInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIdleInstances",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetMaxPendingLatency(val *string) {
	if err := j.validateSetMaxPendingLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPendingLatency",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetMaxTotalInstances(val *float64) {
	if err := j.validateSetMaxTotalInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTotalInstances",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetMinIdleInstances(val *float64) {
	if err := j.validateSetMinIdleInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minIdleInstances",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetMinPendingLatency(val *string) {
	if err := j.validateSetMinPendingLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minPendingLatency",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetMinTotalInstances(val *float64) {
	if err := j.validateSetMinTotalInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTotalInstances",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) PutCpuUtilization(value *GoogleAppEngineFlexibleAppVersionAutomaticScalingCpuUtilization) {
	if err := g.validatePutCpuUtilizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCpuUtilization",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) PutDiskUtilization(value *GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilization) {
	if err := g.validatePutDiskUtilizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDiskUtilization",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) PutNetworkUtilization(value *GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization) {
	if err := g.validatePutNetworkUtilizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkUtilization",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) PutRequestUtilization(value *GoogleAppEngineFlexibleAppVersionAutomaticScalingRequestUtilization) {
	if err := g.validatePutRequestUtilizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRequestUtilization",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetCoolDownPeriod() {
	_jsii_.InvokeVoid(
		g,
		"resetCoolDownPeriod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetDiskUtilization() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskUtilization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetMaxConcurrentRequests() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxConcurrentRequests",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetMaxIdleInstances() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxIdleInstances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetMaxPendingLatency() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxPendingLatency",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetMaxTotalInstances() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxTotalInstances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetMinIdleInstances() {
	_jsii_.InvokeVoid(
		g,
		"resetMinIdleInstances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetMinPendingLatency() {
	_jsii_.InvokeVoid(
		g,
		"resetMinPendingLatency",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetMinTotalInstances() {
	_jsii_.InvokeVoid(
		g,
		"resetMinTotalInstances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetNetworkUtilization() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkUtilization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetRequestUtilization() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestUtilization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


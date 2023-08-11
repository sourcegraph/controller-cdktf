package googlecomputeregionautoscaler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeregionautoscaler/internal"
)

type GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference interface {
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
	CooldownPeriod() *float64
	SetCooldownPeriod(val *float64)
	CooldownPeriodInput() *float64
	CpuUtilization() GoogleComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationOutputReference
	CpuUtilizationInput() *GoogleComputeRegionAutoscalerAutoscalingPolicyCpuUtilization
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeRegionAutoscalerAutoscalingPolicy
	SetInternalValue(val *GoogleComputeRegionAutoscalerAutoscalingPolicy)
	LoadBalancingUtilization() GoogleComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutputReference
	LoadBalancingUtilizationInput() *GoogleComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization
	MaxReplicas() *float64
	SetMaxReplicas(val *float64)
	MaxReplicasInput() *float64
	Metric() GoogleComputeRegionAutoscalerAutoscalingPolicyMetricList
	MetricInput() interface{}
	MinReplicas() *float64
	SetMinReplicas(val *float64)
	MinReplicasInput() *float64
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	ScaleDownControl() GoogleComputeRegionAutoscalerAutoscalingPolicyScaleDownControlOutputReference
	ScaleDownControlInput() *GoogleComputeRegionAutoscalerAutoscalingPolicyScaleDownControl
	ScaleInControl() GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference
	ScaleInControlInput() *GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControl
	ScalingSchedules() GoogleComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesList
	ScalingSchedulesInput() interface{}
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
	PutCpuUtilization(value *GoogleComputeRegionAutoscalerAutoscalingPolicyCpuUtilization)
	PutLoadBalancingUtilization(value *GoogleComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization)
	PutMetric(value interface{})
	PutScaleDownControl(value *GoogleComputeRegionAutoscalerAutoscalingPolicyScaleDownControl)
	PutScaleInControl(value *GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControl)
	PutScalingSchedules(value interface{})
	ResetCooldownPeriod()
	ResetCpuUtilization()
	ResetLoadBalancingUtilization()
	ResetMetric()
	ResetMode()
	ResetScaleDownControl()
	ResetScaleInControl()
	ResetScalingSchedules()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference
type jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) CooldownPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldownPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) CooldownPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldownPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) CpuUtilization() GoogleComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationOutputReference {
	var returns GoogleComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationOutputReference
	_jsii_.Get(
		j,
		"cpuUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) CpuUtilizationInput() *GoogleComputeRegionAutoscalerAutoscalingPolicyCpuUtilization {
	var returns *GoogleComputeRegionAutoscalerAutoscalingPolicyCpuUtilization
	_jsii_.Get(
		j,
		"cpuUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) InternalValue() *GoogleComputeRegionAutoscalerAutoscalingPolicy {
	var returns *GoogleComputeRegionAutoscalerAutoscalingPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) LoadBalancingUtilization() GoogleComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutputReference {
	var returns GoogleComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutputReference
	_jsii_.Get(
		j,
		"loadBalancingUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) LoadBalancingUtilizationInput() *GoogleComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization {
	var returns *GoogleComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization
	_jsii_.Get(
		j,
		"loadBalancingUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) MaxReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) MaxReplicasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) Metric() GoogleComputeRegionAutoscalerAutoscalingPolicyMetricList {
	var returns GoogleComputeRegionAutoscalerAutoscalingPolicyMetricList
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) MetricInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) MinReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) MinReplicasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) ScaleDownControl() GoogleComputeRegionAutoscalerAutoscalingPolicyScaleDownControlOutputReference {
	var returns GoogleComputeRegionAutoscalerAutoscalingPolicyScaleDownControlOutputReference
	_jsii_.Get(
		j,
		"scaleDownControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) ScaleDownControlInput() *GoogleComputeRegionAutoscalerAutoscalingPolicyScaleDownControl {
	var returns *GoogleComputeRegionAutoscalerAutoscalingPolicyScaleDownControl
	_jsii_.Get(
		j,
		"scaleDownControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) ScaleInControl() GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference {
	var returns GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference
	_jsii_.Get(
		j,
		"scaleInControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) ScaleInControlInput() *GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControl {
	var returns *GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControl
	_jsii_.Get(
		j,
		"scaleInControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) ScalingSchedules() GoogleComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesList {
	var returns GoogleComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesList
	_jsii_.Get(
		j,
		"scalingSchedules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) ScalingSchedulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingSchedulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionAutoscalerAutoscalingPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeRegionAutoscaler.GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference_Override(g GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeRegionAutoscaler.GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference)SetCooldownPeriod(val *float64) {
	if err := j.validateSetCooldownPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cooldownPeriod",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference)SetInternalValue(val *GoogleComputeRegionAutoscalerAutoscalingPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference)SetMaxReplicas(val *float64) {
	if err := j.validateSetMaxReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxReplicas",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference)SetMinReplicas(val *float64) {
	if err := j.validateSetMinReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minReplicas",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) PutCpuUtilization(value *GoogleComputeRegionAutoscalerAutoscalingPolicyCpuUtilization) {
	if err := g.validatePutCpuUtilizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCpuUtilization",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) PutLoadBalancingUtilization(value *GoogleComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization) {
	if err := g.validatePutLoadBalancingUtilizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLoadBalancingUtilization",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) PutMetric(value interface{}) {
	if err := g.validatePutMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMetric",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) PutScaleDownControl(value *GoogleComputeRegionAutoscalerAutoscalingPolicyScaleDownControl) {
	if err := g.validatePutScaleDownControlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScaleDownControl",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) PutScaleInControl(value *GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControl) {
	if err := g.validatePutScaleInControlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScaleInControl",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) PutScalingSchedules(value interface{}) {
	if err := g.validatePutScalingSchedulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScalingSchedules",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) ResetCooldownPeriod() {
	_jsii_.InvokeVoid(
		g,
		"resetCooldownPeriod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) ResetCpuUtilization() {
	_jsii_.InvokeVoid(
		g,
		"resetCpuUtilization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) ResetLoadBalancingUtilization() {
	_jsii_.InvokeVoid(
		g,
		"resetLoadBalancingUtilization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) ResetMetric() {
	_jsii_.InvokeVoid(
		g,
		"resetMetric",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		g,
		"resetMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) ResetScaleDownControl() {
	_jsii_.InvokeVoid(
		g,
		"resetScaleDownControl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) ResetScaleInControl() {
	_jsii_.InvokeVoid(
		g,
		"resetScaleInControl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) ResetScalingSchedules() {
	_jsii_.InvokeVoid(
		g,
		"resetScalingSchedules",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


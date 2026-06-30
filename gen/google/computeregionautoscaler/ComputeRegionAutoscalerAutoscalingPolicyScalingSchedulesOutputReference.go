package computeregionautoscaler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeregionautoscaler/internal"
)

type ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Disabled() any
	SetDisabled(val any)
	DisabledInput() any
	DurationSec() *float64
	SetDurationSec(val *float64)
	DurationSecInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	MinRequiredReplicas() *float64
	SetMinRequiredReplicas(val *float64)
	MinRequiredReplicasInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Schedule() *string
	SetSchedule(val *string)
	ScheduleInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	ResetDescription()
	ResetDisabled()
	ResetTimeZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference
type jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) Disabled() any {
	var returns any
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) DisabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) DurationSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) DurationSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) MinRequiredReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRequiredReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) MinRequiredReplicasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRequiredReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) ScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func NewComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionAutoscaler.ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference_Override(c ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionAutoscaler.ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) SetDisabled(val any) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) SetDurationSec(val *float64) {
	if err := j.validateSetDurationSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"durationSec",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) SetMinRequiredReplicas(val *float64) {
	if err := j.validateSetMinRequiredReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minRequiredReplicas",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) SetSchedule(val *string) {
	if err := j.validateSetScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) ResetDisabled() {
	_jsii_.InvokeVoid(
		c,
		"resetDisabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		c,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

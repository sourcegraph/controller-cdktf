package escalationpath

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/incident/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/incident/escalationpath/internal"
)

type EscalationPathPathIfElseThenPathNotifyChannelOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	Targets() EscalationPathPathIfElseThenPathNotifyChannelTargetsList
	TargetsInput() any
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeToAckIntervalCondition() *string
	SetTimeToAckIntervalCondition(val *string)
	TimeToAckIntervalConditionInput() *string
	TimeToAckSeconds() *float64
	SetTimeToAckSeconds(val *float64)
	TimeToAckSecondsInput() *float64
	TimeToAckWeekdayIntervalConfigId() *string
	SetTimeToAckWeekdayIntervalConfigId(val *string)
	TimeToAckWeekdayIntervalConfigIdInput() *string
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
	PutTargets(value any)
	ResetTimeToAckIntervalCondition()
	ResetTimeToAckSeconds()
	ResetTimeToAckWeekdayIntervalConfigId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EscalationPathPathIfElseThenPathNotifyChannelOutputReference
type jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) Targets() EscalationPathPathIfElseThenPathNotifyChannelTargetsList {
	var returns EscalationPathPathIfElseThenPathNotifyChannelTargetsList
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) TargetsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"targetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) TimeToAckIntervalCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeToAckIntervalCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) TimeToAckIntervalConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeToAckIntervalConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) TimeToAckSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeToAckSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) TimeToAckSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeToAckSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) TimeToAckWeekdayIntervalConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeToAckWeekdayIntervalConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) TimeToAckWeekdayIntervalConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeToAckWeekdayIntervalConfigIdInput",
		&returns,
	)
	return returns
}

func NewEscalationPathPathIfElseThenPathNotifyChannelOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EscalationPathPathIfElseThenPathNotifyChannelOutputReference {
	_init_.Initialize()

	if err := validateNewEscalationPathPathIfElseThenPathNotifyChannelOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-incident.escalationPath.EscalationPathPathIfElseThenPathNotifyChannelOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEscalationPathPathIfElseThenPathNotifyChannelOutputReference_Override(e EscalationPathPathIfElseThenPathNotifyChannelOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-incident.escalationPath.EscalationPathPathIfElseThenPathNotifyChannelOutputReference",
		[]any{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) SetTimeToAckIntervalCondition(val *string) {
	if err := j.validateSetTimeToAckIntervalConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeToAckIntervalCondition",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) SetTimeToAckSeconds(val *float64) {
	if err := j.validateSetTimeToAckSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeToAckSeconds",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) SetTimeToAckWeekdayIntervalConfigId(val *string) {
	if err := j.validateSetTimeToAckWeekdayIntervalConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeToAckWeekdayIntervalConfigId",
		val,
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) PutTargets(value any) {
	if err := e.validatePutTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTargets",
		[]any{value},
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) ResetTimeToAckIntervalCondition() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeToAckIntervalCondition",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) ResetTimeToAckSeconds() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeToAckSeconds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) ResetTimeToAckWeekdayIntervalConfigId() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeToAckWeekdayIntervalConfigId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		e,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathNotifyChannelOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

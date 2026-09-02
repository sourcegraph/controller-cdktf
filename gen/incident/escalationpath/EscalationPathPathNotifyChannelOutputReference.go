package escalationpath

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/incident/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/incident/escalationpath/internal"
)

type EscalationPathPathNotifyChannelOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Targets() EscalationPathPathNotifyChannelTargetsList
	TargetsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutTargets(value interface{})
	ResetTimeToAckIntervalCondition()
	ResetTimeToAckSeconds()
	ResetTimeToAckWeekdayIntervalConfigId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EscalationPathPathNotifyChannelOutputReference
type jsiiProxy_EscalationPathPathNotifyChannelOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) Targets() EscalationPathPathNotifyChannelTargetsList {
	var returns EscalationPathPathNotifyChannelTargetsList
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) TargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) TimeToAckIntervalCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeToAckIntervalCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) TimeToAckIntervalConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeToAckIntervalConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) TimeToAckSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeToAckSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) TimeToAckSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeToAckSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) TimeToAckWeekdayIntervalConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeToAckWeekdayIntervalConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) TimeToAckWeekdayIntervalConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeToAckWeekdayIntervalConfigIdInput",
		&returns,
	)
	return returns
}


func NewEscalationPathPathNotifyChannelOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) EscalationPathPathNotifyChannelOutputReference {
	_init_.Initialize()

	if err := validateNewEscalationPathPathNotifyChannelOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EscalationPathPathNotifyChannelOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-incident.escalationPath.EscalationPathPathNotifyChannelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEscalationPathPathNotifyChannelOutputReference_Override(e EscalationPathPathNotifyChannelOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-incident.escalationPath.EscalationPathPathNotifyChannelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference)SetTimeToAckIntervalCondition(val *string) {
	if err := j.validateSetTimeToAckIntervalConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeToAckIntervalCondition",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference)SetTimeToAckSeconds(val *float64) {
	if err := j.validateSetTimeToAckSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeToAckSeconds",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathNotifyChannelOutputReference)SetTimeToAckWeekdayIntervalConfigId(val *string) {
	if err := j.validateSetTimeToAckWeekdayIntervalConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeToAckWeekdayIntervalConfigId",
		val,
	)
}

func (e *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) PutTargets(value interface{}) {
	if err := e.validatePutTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTargets",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) ResetTimeToAckIntervalCondition() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeToAckIntervalCondition",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) ResetTimeToAckSeconds() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeToAckSeconds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) ResetTimeToAckWeekdayIntervalConfigId() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeToAckWeekdayIntervalConfigId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathNotifyChannelOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


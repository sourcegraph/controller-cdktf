package escalationpath

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/incident/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/incident/escalationpath/internal"
)

type EscalationPathPathIfElseThenPathLevelOutputReference interface {
	cdktf.ComplexObject
	AckMode() *string
	SetAckMode(val *string)
	AckModeInput() *string
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
	RoundRobinConfig() EscalationPathPathIfElseThenPathLevelRoundRobinConfigOutputReference
	RoundRobinConfigInput() interface{}
	Targets() EscalationPathPathIfElseThenPathLevelTargetsList
	TargetsInput() interface{}
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
	PutRoundRobinConfig(value *EscalationPathPathIfElseThenPathLevelRoundRobinConfig)
	PutTargets(value interface{})
	ResetAckMode()
	ResetRoundRobinConfig()
	ResetTimeToAckIntervalCondition()
	ResetTimeToAckSeconds()
	ResetTimeToAckWeekdayIntervalConfigId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EscalationPathPathIfElseThenPathLevelOutputReference
type jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) AckMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ackMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) AckModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ackModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) RoundRobinConfig() EscalationPathPathIfElseThenPathLevelRoundRobinConfigOutputReference {
	var returns EscalationPathPathIfElseThenPathLevelRoundRobinConfigOutputReference
	_jsii_.Get(
		j,
		"roundRobinConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) RoundRobinConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"roundRobinConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) Targets() EscalationPathPathIfElseThenPathLevelTargetsList {
	var returns EscalationPathPathIfElseThenPathLevelTargetsList
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) TargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) TimeToAckIntervalCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeToAckIntervalCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) TimeToAckIntervalConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeToAckIntervalConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) TimeToAckSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeToAckSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) TimeToAckSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeToAckSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) TimeToAckWeekdayIntervalConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeToAckWeekdayIntervalConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) TimeToAckWeekdayIntervalConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeToAckWeekdayIntervalConfigIdInput",
		&returns,
	)
	return returns
}


func NewEscalationPathPathIfElseThenPathLevelOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EscalationPathPathIfElseThenPathLevelOutputReference {
	_init_.Initialize()

	if err := validateNewEscalationPathPathIfElseThenPathLevelOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-incident.escalationPath.EscalationPathPathIfElseThenPathLevelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEscalationPathPathIfElseThenPathLevelOutputReference_Override(e EscalationPathPathIfElseThenPathLevelOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-incident.escalationPath.EscalationPathPathIfElseThenPathLevelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference)SetAckMode(val *string) {
	if err := j.validateSetAckModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ackMode",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference)SetTimeToAckIntervalCondition(val *string) {
	if err := j.validateSetTimeToAckIntervalConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeToAckIntervalCondition",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference)SetTimeToAckSeconds(val *float64) {
	if err := j.validateSetTimeToAckSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeToAckSeconds",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference)SetTimeToAckWeekdayIntervalConfigId(val *string) {
	if err := j.validateSetTimeToAckWeekdayIntervalConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeToAckWeekdayIntervalConfigId",
		val,
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) PutRoundRobinConfig(value *EscalationPathPathIfElseThenPathLevelRoundRobinConfig) {
	if err := e.validatePutRoundRobinConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRoundRobinConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) PutTargets(value interface{}) {
	if err := e.validatePutTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTargets",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) ResetAckMode() {
	_jsii_.InvokeVoid(
		e,
		"resetAckMode",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) ResetRoundRobinConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetRoundRobinConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) ResetTimeToAckIntervalCondition() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeToAckIntervalCondition",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) ResetTimeToAckSeconds() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeToAckSeconds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) ResetTimeToAckWeekdayIntervalConfigId() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeToAckWeekdayIntervalConfigId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathLevelOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


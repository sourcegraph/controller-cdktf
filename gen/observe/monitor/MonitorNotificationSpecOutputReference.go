package monitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/observe/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/observe/monitor/internal"
)

type MonitorNotificationSpecOutputReference interface {
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
	Importance() *string
	SetImportance(val *string)
	ImportanceInput() *string
	InternalValue() *MonitorNotificationSpec
	SetInternalValue(val *MonitorNotificationSpec)
	Merge() *string
	SetMerge(val *string)
	MergeInput() *string
	NotifyOnClose() interface{}
	SetNotifyOnClose(val interface{})
	NotifyOnCloseInput() interface{}
	NotifyOnReminder() cdktn.IResolvable
	ReminderFrequency() *string
	SetReminderFrequency(val *string)
	ReminderFrequencyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	ResetImportance()
	ResetMerge()
	ResetNotifyOnClose()
	ResetReminderFrequency()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorNotificationSpecOutputReference
type jsiiProxy_MonitorNotificationSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference) Importance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference) ImportanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference) InternalValue() *MonitorNotificationSpec {
	var returns *MonitorNotificationSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference) Merge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"merge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference) MergeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference) NotifyOnClose() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnClose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference) NotifyOnCloseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnCloseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference) NotifyOnReminder() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"notifyOnReminder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference) ReminderFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reminderFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference) ReminderFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reminderFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorNotificationSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MonitorNotificationSpecOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorNotificationSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorNotificationSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-observe.monitor.MonitorNotificationSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorNotificationSpecOutputReference_Override(m MonitorNotificationSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-observe.monitor.MonitorNotificationSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference)SetImportance(val *string) {
	if err := j.validateSetImportanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importance",
		val,
	)
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference)SetInternalValue(val *MonitorNotificationSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference)SetMerge(val *string) {
	if err := j.validateSetMergeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"merge",
		val,
	)
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference)SetNotifyOnClose(val interface{}) {
	if err := j.validateSetNotifyOnCloseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyOnClose",
		val,
	)
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference)SetReminderFrequency(val *string) {
	if err := j.validateSetReminderFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reminderFrequency",
		val,
	)
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorNotificationSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorNotificationSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorNotificationSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorNotificationSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorNotificationSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorNotificationSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorNotificationSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorNotificationSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorNotificationSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorNotificationSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorNotificationSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorNotificationSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorNotificationSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorNotificationSpecOutputReference) ResetImportance() {
	_jsii_.InvokeVoid(
		m,
		"resetImportance",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorNotificationSpecOutputReference) ResetMerge() {
	_jsii_.InvokeVoid(
		m,
		"resetMerge",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorNotificationSpecOutputReference) ResetNotifyOnClose() {
	_jsii_.InvokeVoid(
		m,
		"resetNotifyOnClose",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorNotificationSpecOutputReference) ResetReminderFrequency() {
	_jsii_.InvokeVoid(
		m,
		"resetReminderFrequency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorNotificationSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorNotificationSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


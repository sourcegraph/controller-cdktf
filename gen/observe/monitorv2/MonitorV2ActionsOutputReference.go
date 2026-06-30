package monitorv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/observe/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/observe/monitorv2/internal"
)

type MonitorV2ActionsOutputReference interface {
	cdktf.ComplexObject
	Action() MonitorV2ActionsActionOutputReference
	ActionInput() *MonitorV2ActionsAction
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
	Conditions() MonitorV2ActionsConditionsOutputReference
	ConditionsInput() *MonitorV2ActionsConditions
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	Levels() *[]*string
	SetLevels(val *[]*string)
	LevelsInput() *[]*string
	Oid() *string
	SetOid(val *string)
	OidInput() *string
	SendEndNotifications() any
	SetSendEndNotifications(val any)
	SendEndNotificationsInput() any
	SendRemindersInterval() *string
	SetSendRemindersInterval(val *string)
	SendRemindersIntervalInput() *string
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
	PutAction(value *MonitorV2ActionsAction)
	PutConditions(value *MonitorV2ActionsConditions)
	ResetAction()
	ResetConditions()
	ResetLevels()
	ResetOid()
	ResetSendEndNotifications()
	ResetSendRemindersInterval()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorV2ActionsOutputReference
type jsiiProxy_MonitorV2ActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) Action() MonitorV2ActionsActionOutputReference {
	var returns MonitorV2ActionsActionOutputReference
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) ActionInput() *MonitorV2ActionsAction {
	var returns *MonitorV2ActionsAction
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) Conditions() MonitorV2ActionsConditionsOutputReference {
	var returns MonitorV2ActionsConditionsOutputReference
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) ConditionsInput() *MonitorV2ActionsConditions {
	var returns *MonitorV2ActionsConditions
	_jsii_.Get(
		j,
		"conditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) Levels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"levels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) LevelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"levelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) Oid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) OidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) SendEndNotifications() any {
	var returns any
	_jsii_.Get(
		j,
		"sendEndNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) SendEndNotificationsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"sendEndNotificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) SendRemindersInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sendRemindersInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) SendRemindersIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sendRemindersIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewMonitorV2ActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MonitorV2ActionsOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorV2ActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorV2ActionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-observe.monitorV2.MonitorV2ActionsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMonitorV2ActionsOutputReference_Override(m MonitorV2ActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-observe.monitorV2.MonitorV2ActionsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) SetLevels(val *[]*string) {
	if err := j.validateSetLevelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"levels",
		val,
	)
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) SetOid(val *string) {
	if err := j.validateSetOidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oid",
		val,
	)
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) SetSendEndNotifications(val any) {
	if err := j.validateSetSendEndNotificationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendEndNotifications",
		val,
	)
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) SetSendRemindersInterval(val *string) {
	if err := j.validateSetSendRemindersIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendRemindersInterval",
		val,
	)
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorV2ActionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) PutAction(value *MonitorV2ActionsAction) {
	if err := m.validatePutActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAction",
		[]any{value},
	)
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) PutConditions(value *MonitorV2ActionsConditions) {
	if err := m.validatePutConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putConditions",
		[]any{value},
	)
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		m,
		"resetAction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) ResetConditions() {
	_jsii_.InvokeVoid(
		m,
		"resetConditions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) ResetLevels() {
	_jsii_.InvokeVoid(
		m,
		"resetLevels",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) ResetOid() {
	_jsii_.InvokeVoid(
		m,
		"resetOid",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) ResetSendEndNotifications() {
	_jsii_.InvokeVoid(
		m,
		"resetSendEndNotifications",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) ResetSendRemindersInterval() {
	_jsii_.InvokeVoid(
		m,
		"resetSendRemindersInterval",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		m,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2ActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

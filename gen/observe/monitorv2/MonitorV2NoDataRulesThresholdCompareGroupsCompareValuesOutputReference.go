package monitorv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/observe/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/observe/monitorv2/internal"
)

type MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference interface {
	cdktf.ComplexObject
	CompareFn() *string
	SetCompareFn(val *string)
	CompareFnInput() *string
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
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValueBool() any
	SetValueBool(val any)
	ValueBoolInput() any
	ValueDuration() *[]*string
	SetValueDuration(val *[]*string)
	ValueDurationInput() *[]*string
	ValueFloat64() *[]*float64
	SetValueFloat64(val *[]*float64)
	ValueFloat64Input() *[]*float64
	ValueInt64() *[]*float64
	SetValueInt64(val *[]*float64)
	ValueInt64Input() *[]*float64
	ValueString() *[]*string
	SetValueString(val *[]*string)
	ValueStringInput() *[]*string
	ValueTimestamp() *[]*string
	SetValueTimestamp(val *[]*string)
	ValueTimestampInput() *[]*string
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
	ResetValueBool()
	ResetValueDuration()
	ResetValueFloat64()
	ResetValueInt64()
	ResetValueString()
	ResetValueTimestamp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference
type jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) CompareFn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compareFn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) CompareFnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compareFnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ValueBool() any {
	var returns any
	_jsii_.Get(
		j,
		"valueBool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ValueBoolInput() any {
	var returns any
	_jsii_.Get(
		j,
		"valueBoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ValueDuration() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valueDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ValueDurationInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valueDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ValueFloat64() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"valueFloat64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ValueFloat64Input() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"valueFloat64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ValueInt64() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"valueInt64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ValueInt64Input() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"valueInt64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ValueString() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valueString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ValueStringInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valueStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ValueTimestamp() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valueTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ValueTimestampInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valueTimestampInput",
		&returns,
	)
	return returns
}

func NewMonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-observe.monitorV2.MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference_Override(m MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-observe.monitorV2.MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) SetCompareFn(val *string) {
	if err := j.validateSetCompareFnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compareFn",
		val,
	)
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) SetValueBool(val any) {
	if err := j.validateSetValueBoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueBool",
		val,
	)
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) SetValueDuration(val *[]*string) {
	if err := j.validateSetValueDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueDuration",
		val,
	)
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) SetValueFloat64(val *[]*float64) {
	if err := j.validateSetValueFloat64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueFloat64",
		val,
	)
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) SetValueInt64(val *[]*float64) {
	if err := j.validateSetValueInt64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueInt64",
		val,
	)
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) SetValueString(val *[]*string) {
	if err := j.validateSetValueStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueString",
		val,
	)
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) SetValueTimestamp(val *[]*string) {
	if err := j.validateSetValueTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueTimestamp",
		val,
	)
}

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ResetValueBool() {
	_jsii_.InvokeVoid(
		m,
		"resetValueBool",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ResetValueDuration() {
	_jsii_.InvokeVoid(
		m,
		"resetValueDuration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ResetValueFloat64() {
	_jsii_.InvokeVoid(
		m,
		"resetValueFloat64",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ResetValueInt64() {
	_jsii_.InvokeVoid(
		m,
		"resetValueInt64",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ResetValueString() {
	_jsii_.InvokeVoid(
		m,
		"resetValueString",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ResetValueTimestamp() {
	_jsii_.InvokeVoid(
		m,
		"resetValueTimestamp",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsCompareValuesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

package monitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/observe/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/observe/monitor/internal"
)

type MonitorRuleChangeOutputReference interface {
	cdktn.ComplexObject
	AggregateFunction() *string
	SetAggregateFunction(val *string)
	AggregateFunctionInput() *string
	BaselineTime() *string
	SetBaselineTime(val *string)
	BaselineTimeInput() *string
	ChangeType() *string
	SetChangeType(val *string)
	ChangeTypeInput() *string
	CompareFunction() *string
	SetCompareFunction(val *string)
	CompareFunctionInput() *string
	CompareValue() *float64
	SetCompareValue(val *float64)
	CompareValueInput() *float64
	CompareValues() *[]*float64
	SetCompareValues(val *[]*float64)
	CompareValuesInput() *[]*float64
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
	InternalValue() *MonitorRuleChange
	SetInternalValue(val *MonitorRuleChange)
	LookbackTime() *string
	SetLookbackTime(val *string)
	LookbackTimeInput() *string
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
	ResetAggregateFunction()
	ResetChangeType()
	ResetCompareValue()
	ResetCompareValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorRuleChangeOutputReference
type jsiiProxy_MonitorRuleChangeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference) AggregateFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregateFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference) AggregateFunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregateFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference) BaselineTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baselineTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference) BaselineTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baselineTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference) ChangeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference) ChangeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference) CompareFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compareFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference) CompareFunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compareFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference) CompareValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"compareValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference) CompareValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"compareValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference) CompareValues() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"compareValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference) CompareValuesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"compareValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference) InternalValue() *MonitorRuleChange {
	var returns *MonitorRuleChange
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference) LookbackTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookbackTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference) LookbackTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookbackTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorRuleChangeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MonitorRuleChangeOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorRuleChangeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorRuleChangeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-observe.monitor.MonitorRuleChangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorRuleChangeOutputReference_Override(m MonitorRuleChangeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-observe.monitor.MonitorRuleChangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference)SetAggregateFunction(val *string) {
	if err := j.validateSetAggregateFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregateFunction",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference)SetBaselineTime(val *string) {
	if err := j.validateSetBaselineTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baselineTime",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference)SetChangeType(val *string) {
	if err := j.validateSetChangeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"changeType",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference)SetCompareFunction(val *string) {
	if err := j.validateSetCompareFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compareFunction",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference)SetCompareValue(val *float64) {
	if err := j.validateSetCompareValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compareValue",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference)SetCompareValues(val *[]*float64) {
	if err := j.validateSetCompareValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compareValues",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference)SetInternalValue(val *MonitorRuleChange) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference)SetLookbackTime(val *string) {
	if err := j.validateSetLookbackTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lookbackTime",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleChangeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorRuleChangeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleChangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorRuleChangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MonitorRuleChangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorRuleChangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorRuleChangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorRuleChangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorRuleChangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorRuleChangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorRuleChangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorRuleChangeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleChangeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MonitorRuleChangeOutputReference) ResetAggregateFunction() {
	_jsii_.InvokeVoid(
		m,
		"resetAggregateFunction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorRuleChangeOutputReference) ResetChangeType() {
	_jsii_.InvokeVoid(
		m,
		"resetChangeType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorRuleChangeOutputReference) ResetCompareValue() {
	_jsii_.InvokeVoid(
		m,
		"resetCompareValue",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorRuleChangeOutputReference) ResetCompareValues() {
	_jsii_.InvokeVoid(
		m,
		"resetCompareValues",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorRuleChangeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MonitorRuleChangeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


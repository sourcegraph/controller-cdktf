package monitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/observe/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/observe/monitor/internal"
)

type MonitorRuleLogOutputReference interface {
	cdktf.ComplexObject
	CompareFunction() *string
	SetCompareFunction(val *string)
	CompareFunctionInput() *string
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
	ExpressionSummary() *string
	SetExpressionSummary(val *string)
	ExpressionSummaryInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorRuleLog
	SetInternalValue(val *MonitorRuleLog)
	LogStageId() *string
	SetLogStageId(val *string)
	LogStageIdInput() *string
	LookbackTime() *string
	SetLookbackTime(val *string)
	LookbackTimeInput() *string
	SourceLogDataset() *string
	SetSourceLogDataset(val *string)
	SourceLogDatasetInput() *string
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
	ResetCompareValues()
	ResetExpressionSummary()
	ResetLogStageId()
	ResetSourceLogDataset()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorRuleLogOutputReference
type jsiiProxy_MonitorRuleLogOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorRuleLogOutputReference) CompareFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compareFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleLogOutputReference) CompareFunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compareFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleLogOutputReference) CompareValues() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"compareValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleLogOutputReference) CompareValuesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"compareValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleLogOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleLogOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleLogOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleLogOutputReference) ExpressionSummary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionSummary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleLogOutputReference) ExpressionSummaryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionSummaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleLogOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleLogOutputReference) InternalValue() *MonitorRuleLog {
	var returns *MonitorRuleLog
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleLogOutputReference) LogStageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleLogOutputReference) LogStageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleLogOutputReference) LookbackTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookbackTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleLogOutputReference) LookbackTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookbackTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleLogOutputReference) SourceLogDataset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceLogDataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleLogOutputReference) SourceLogDatasetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceLogDatasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleLogOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleLogOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorRuleLogOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorRuleLogOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorRuleLogOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorRuleLogOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-observe.monitor.MonitorRuleLogOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorRuleLogOutputReference_Override(m MonitorRuleLogOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-observe.monitor.MonitorRuleLogOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorRuleLogOutputReference)SetCompareFunction(val *string) {
	if err := j.validateSetCompareFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compareFunction",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleLogOutputReference)SetCompareValues(val *[]*float64) {
	if err := j.validateSetCompareValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compareValues",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleLogOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleLogOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleLogOutputReference)SetExpressionSummary(val *string) {
	if err := j.validateSetExpressionSummaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expressionSummary",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleLogOutputReference)SetInternalValue(val *MonitorRuleLog) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleLogOutputReference)SetLogStageId(val *string) {
	if err := j.validateSetLogStageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logStageId",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleLogOutputReference)SetLookbackTime(val *string) {
	if err := j.validateSetLookbackTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lookbackTime",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleLogOutputReference)SetSourceLogDataset(val *string) {
	if err := j.validateSetSourceLogDatasetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceLogDataset",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleLogOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleLogOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorRuleLogOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleLogOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorRuleLogOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleLogOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorRuleLogOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorRuleLogOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorRuleLogOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorRuleLogOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorRuleLogOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorRuleLogOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorRuleLogOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleLogOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleLogOutputReference) ResetCompareValues() {
	_jsii_.InvokeVoid(
		m,
		"resetCompareValues",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorRuleLogOutputReference) ResetExpressionSummary() {
	_jsii_.InvokeVoid(
		m,
		"resetExpressionSummary",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorRuleLogOutputReference) ResetLogStageId() {
	_jsii_.InvokeVoid(
		m,
		"resetLogStageId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorRuleLogOutputReference) ResetSourceLogDataset() {
	_jsii_.InvokeVoid(
		m,
		"resetSourceLogDataset",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorRuleLogOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleLogOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package monitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/observe/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/observe/monitor/internal"
)

type MonitorRuleOutputReference interface {
	cdktf.ComplexObject
	Change() MonitorRuleChangeOutputReference
	ChangeInput() *MonitorRuleChange
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
	Count() MonitorRuleCountOutputReference
	CountInput() *MonitorRuleCount
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Facet() MonitorRuleFacetOutputReference
	FacetInput() *MonitorRuleFacet
	// Experimental.
	Fqn() *string
	GroupByGroup() MonitorRuleGroupByGroupList
	GroupByGroupInput() interface{}
	InternalValue() *MonitorRule
	SetInternalValue(val *MonitorRule)
	Log() MonitorRuleLogOutputReference
	LogInput() *MonitorRuleLog
	Promote() MonitorRulePromoteOutputReference
	PromoteInput() *MonitorRulePromote
	SourceColumn() *string
	SetSourceColumn(val *string)
	SourceColumnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Threshold() MonitorRuleThresholdOutputReference
	ThresholdInput() *MonitorRuleThreshold
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
	PutChange(value *MonitorRuleChange)
	PutCount(value *MonitorRuleCount)
	PutFacet(value *MonitorRuleFacet)
	PutGroupByGroup(value interface{})
	PutLog(value *MonitorRuleLog)
	PutPromote(value *MonitorRulePromote)
	PutThreshold(value *MonitorRuleThreshold)
	ResetChange()
	ResetCount()
	ResetFacet()
	ResetGroupByGroup()
	ResetLog()
	ResetPromote()
	ResetSourceColumn()
	ResetThreshold()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorRuleOutputReference
type jsiiProxy_MonitorRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorRuleOutputReference) Change() MonitorRuleChangeOutputReference {
	var returns MonitorRuleChangeOutputReference
	_jsii_.Get(
		j,
		"change",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) ChangeInput() *MonitorRuleChange {
	var returns *MonitorRuleChange
	_jsii_.Get(
		j,
		"changeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) Count() MonitorRuleCountOutputReference {
	var returns MonitorRuleCountOutputReference
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) CountInput() *MonitorRuleCount {
	var returns *MonitorRuleCount
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) Facet() MonitorRuleFacetOutputReference {
	var returns MonitorRuleFacetOutputReference
	_jsii_.Get(
		j,
		"facet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) FacetInput() *MonitorRuleFacet {
	var returns *MonitorRuleFacet
	_jsii_.Get(
		j,
		"facetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) GroupByGroup() MonitorRuleGroupByGroupList {
	var returns MonitorRuleGroupByGroupList
	_jsii_.Get(
		j,
		"groupByGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) GroupByGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupByGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) InternalValue() *MonitorRule {
	var returns *MonitorRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) Log() MonitorRuleLogOutputReference {
	var returns MonitorRuleLogOutputReference
	_jsii_.Get(
		j,
		"log",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) LogInput() *MonitorRuleLog {
	var returns *MonitorRuleLog
	_jsii_.Get(
		j,
		"logInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) Promote() MonitorRulePromoteOutputReference {
	var returns MonitorRulePromoteOutputReference
	_jsii_.Get(
		j,
		"promote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) PromoteInput() *MonitorRulePromote {
	var returns *MonitorRulePromote
	_jsii_.Get(
		j,
		"promoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) SourceColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) SourceColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) Threshold() MonitorRuleThresholdOutputReference {
	var returns MonitorRuleThresholdOutputReference
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleOutputReference) ThresholdInput() *MonitorRuleThreshold {
	var returns *MonitorRuleThreshold
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}


func NewMonitorRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorRuleOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-observe.monitor.MonitorRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorRuleOutputReference_Override(m MonitorRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-observe.monitor.MonitorRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleOutputReference)SetInternalValue(val *MonitorRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleOutputReference)SetSourceColumn(val *string) {
	if err := j.validateSetSourceColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceColumn",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorRuleOutputReference) PutChange(value *MonitorRuleChange) {
	if err := m.validatePutChangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putChange",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorRuleOutputReference) PutCount(value *MonitorRuleCount) {
	if err := m.validatePutCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCount",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorRuleOutputReference) PutFacet(value *MonitorRuleFacet) {
	if err := m.validatePutFacetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFacet",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorRuleOutputReference) PutGroupByGroup(value interface{}) {
	if err := m.validatePutGroupByGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putGroupByGroup",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorRuleOutputReference) PutLog(value *MonitorRuleLog) {
	if err := m.validatePutLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putLog",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorRuleOutputReference) PutPromote(value *MonitorRulePromote) {
	if err := m.validatePutPromoteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPromote",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorRuleOutputReference) PutThreshold(value *MonitorRuleThreshold) {
	if err := m.validatePutThresholdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putThreshold",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorRuleOutputReference) ResetChange() {
	_jsii_.InvokeVoid(
		m,
		"resetChange",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorRuleOutputReference) ResetCount() {
	_jsii_.InvokeVoid(
		m,
		"resetCount",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorRuleOutputReference) ResetFacet() {
	_jsii_.InvokeVoid(
		m,
		"resetFacet",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorRuleOutputReference) ResetGroupByGroup() {
	_jsii_.InvokeVoid(
		m,
		"resetGroupByGroup",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorRuleOutputReference) ResetLog() {
	_jsii_.InvokeVoid(
		m,
		"resetLog",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorRuleOutputReference) ResetPromote() {
	_jsii_.InvokeVoid(
		m,
		"resetPromote",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorRuleOutputReference) ResetSourceColumn() {
	_jsii_.InvokeVoid(
		m,
		"resetSourceColumn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorRuleOutputReference) ResetThreshold() {
	_jsii_.InvokeVoid(
		m,
		"resetThreshold",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MonitorRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


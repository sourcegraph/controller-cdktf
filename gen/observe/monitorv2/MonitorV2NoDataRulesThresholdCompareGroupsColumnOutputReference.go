package monitorv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/observe/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/observe/monitorv2/internal"
)

type MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference interface {
	cdktf.ComplexObject
	ColumnPath() MonitorV2NoDataRulesThresholdCompareGroupsColumnColumnPathOutputReference
	ColumnPathInput() *MonitorV2NoDataRulesThresholdCompareGroupsColumnColumnPath
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
	InternalValue() *MonitorV2NoDataRulesThresholdCompareGroupsColumn
	SetInternalValue(val *MonitorV2NoDataRulesThresholdCompareGroupsColumn)
	LinkColumn() MonitorV2NoDataRulesThresholdCompareGroupsColumnLinkColumnOutputReference
	LinkColumnInput() *MonitorV2NoDataRulesThresholdCompareGroupsColumnLinkColumn
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
	PutColumnPath(value *MonitorV2NoDataRulesThresholdCompareGroupsColumnColumnPath)
	PutLinkColumn(value *MonitorV2NoDataRulesThresholdCompareGroupsColumnLinkColumn)
	ResetColumnPath()
	ResetLinkColumn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference
type jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) ColumnPath() MonitorV2NoDataRulesThresholdCompareGroupsColumnColumnPathOutputReference {
	var returns MonitorV2NoDataRulesThresholdCompareGroupsColumnColumnPathOutputReference
	_jsii_.Get(
		j,
		"columnPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) ColumnPathInput() *MonitorV2NoDataRulesThresholdCompareGroupsColumnColumnPath {
	var returns *MonitorV2NoDataRulesThresholdCompareGroupsColumnColumnPath
	_jsii_.Get(
		j,
		"columnPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) InternalValue() *MonitorV2NoDataRulesThresholdCompareGroupsColumn {
	var returns *MonitorV2NoDataRulesThresholdCompareGroupsColumn
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) LinkColumn() MonitorV2NoDataRulesThresholdCompareGroupsColumnLinkColumnOutputReference {
	var returns MonitorV2NoDataRulesThresholdCompareGroupsColumnLinkColumnOutputReference
	_jsii_.Get(
		j,
		"linkColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) LinkColumnInput() *MonitorV2NoDataRulesThresholdCompareGroupsColumnLinkColumn {
	var returns *MonitorV2NoDataRulesThresholdCompareGroupsColumnLinkColumn
	_jsii_.Get(
		j,
		"linkColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-observe.monitorV2.MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference_Override(m MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-observe.monitorV2.MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference",
		[]any{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) SetInternalValue(val *MonitorV2NoDataRulesThresholdCompareGroupsColumn) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) PutColumnPath(value *MonitorV2NoDataRulesThresholdCompareGroupsColumnColumnPath) {
	if err := m.validatePutColumnPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putColumnPath",
		[]any{value},
	)
}

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) PutLinkColumn(value *MonitorV2NoDataRulesThresholdCompareGroupsColumnLinkColumn) {
	if err := m.validatePutLinkColumnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putLinkColumn",
		[]any{value},
	)
}

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) ResetColumnPath() {
	_jsii_.InvokeVoid(
		m,
		"resetColumnPath",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) ResetLinkColumn() {
	_jsii_.InvokeVoid(
		m,
		"resetLinkColumn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (m *jsiiProxy_MonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

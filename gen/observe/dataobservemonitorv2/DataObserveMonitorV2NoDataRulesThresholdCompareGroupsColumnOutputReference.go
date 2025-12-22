package dataobservemonitorv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/observe/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/observe/dataobservemonitorv2/internal"
)

type DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference interface {
	cdktf.ComplexObject
	ColumnPath() DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnColumnPathList
	ColumnPathInput() interface{}
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
	LinkColumn() DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnLinkColumnList
	LinkColumnInput() interface{}
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
	PutColumnPath(value interface{})
	PutLinkColumn(value interface{})
	ResetColumnPath()
	ResetLinkColumn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference
type jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) ColumnPath() DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnColumnPathList {
	var returns DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnColumnPathList
	_jsii_.Get(
		j,
		"columnPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) ColumnPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) LinkColumn() DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnLinkColumnList {
	var returns DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnLinkColumnList
	_jsii_.Get(
		j,
		"linkColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) LinkColumnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linkColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference {
	_init_.Initialize()

	if err := validateNewDataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-observe.dataObserveMonitorV2.DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference_Override(d DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-observe.dataObserveMonitorV2.DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) PutColumnPath(value interface{}) {
	if err := d.validatePutColumnPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putColumnPath",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) PutLinkColumn(value interface{}) {
	if err := d.validatePutLinkColumnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLinkColumn",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) ResetColumnPath() {
	_jsii_.InvokeVoid(
		d,
		"resetColumnPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) ResetLinkColumn() {
	_jsii_.InvokeVoid(
		d,
		"resetLinkColumn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareGroupsColumnOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


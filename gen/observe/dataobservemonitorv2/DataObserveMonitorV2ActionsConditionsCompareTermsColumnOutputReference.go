package dataobservemonitorv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/observe/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/observe/dataobservemonitorv2/internal"
)

type DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference interface {
	cdktf.ComplexObject
	ColumnPath() DataObserveMonitorV2ActionsConditionsCompareTermsColumnColumnPathList
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
	LinkColumn() DataObserveMonitorV2ActionsConditionsCompareTermsColumnLinkColumnList
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

// The jsii proxy struct for DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference
type jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) ColumnPath() DataObserveMonitorV2ActionsConditionsCompareTermsColumnColumnPathList {
	var returns DataObserveMonitorV2ActionsConditionsCompareTermsColumnColumnPathList
	_jsii_.Get(
		j,
		"columnPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) ColumnPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) LinkColumn() DataObserveMonitorV2ActionsConditionsCompareTermsColumnLinkColumnList {
	var returns DataObserveMonitorV2ActionsConditionsCompareTermsColumnLinkColumnList
	_jsii_.Get(
		j,
		"linkColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) LinkColumnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linkColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference {
	_init_.Initialize()

	if err := validateNewDataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-observe.dataObserveMonitorV2.DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference_Override(d DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-observe.dataObserveMonitorV2.DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) PutColumnPath(value interface{}) {
	if err := d.validatePutColumnPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putColumnPath",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) PutLinkColumn(value interface{}) {
	if err := d.validatePutLinkColumnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLinkColumn",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) ResetColumnPath() {
	_jsii_.InvokeVoid(
		d,
		"resetColumnPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) ResetLinkColumn() {
	_jsii_.InvokeVoid(
		d,
		"resetLinkColumn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataObserveMonitorV2ActionsConditionsCompareTermsColumnOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


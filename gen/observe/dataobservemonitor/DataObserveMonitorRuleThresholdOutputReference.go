package dataobservemonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/observe/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/observe/dataobservemonitor/internal"
)

type DataObserveMonitorRuleThresholdOutputReference interface {
	cdktn.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ThresholdAggFunction() *string
	SetThresholdAggFunction(val *string)
	ThresholdAggFunctionInput() *string
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
	ResetCompareFunction()
	ResetCompareValues()
	ResetLookbackTime()
	ResetThresholdAggFunction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataObserveMonitorRuleThresholdOutputReference
type jsiiProxy_DataObserveMonitorRuleThresholdOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) CompareFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compareFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) CompareFunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compareFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) CompareValues() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"compareValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) CompareValuesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"compareValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) LookbackTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookbackTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) LookbackTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookbackTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) ThresholdAggFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thresholdAggFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) ThresholdAggFunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thresholdAggFunctionInput",
		&returns,
	)
	return returns
}


func NewDataObserveMonitorRuleThresholdOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataObserveMonitorRuleThresholdOutputReference {
	_init_.Initialize()

	if err := validateNewDataObserveMonitorRuleThresholdOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataObserveMonitorRuleThresholdOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-observe.dataObserveMonitor.DataObserveMonitorRuleThresholdOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataObserveMonitorRuleThresholdOutputReference_Override(d DataObserveMonitorRuleThresholdOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-observe.dataObserveMonitor.DataObserveMonitorRuleThresholdOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference)SetCompareFunction(val *string) {
	if err := j.validateSetCompareFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compareFunction",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference)SetCompareValues(val *[]*float64) {
	if err := j.validateSetCompareValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compareValues",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference)SetLookbackTime(val *string) {
	if err := j.validateSetLookbackTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lookbackTime",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference)SetThresholdAggFunction(val *string) {
	if err := j.validateSetThresholdAggFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdAggFunction",
		val,
	)
}

func (d *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) ResetCompareFunction() {
	_jsii_.InvokeVoid(
		d,
		"resetCompareFunction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) ResetCompareValues() {
	_jsii_.InvokeVoid(
		d,
		"resetCompareValues",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) ResetLookbackTime() {
	_jsii_.InvokeVoid(
		d,
		"resetLookbackTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) ResetThresholdAggFunction() {
	_jsii_.InvokeVoid(
		d,
		"resetThresholdAggFunction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorRuleThresholdOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


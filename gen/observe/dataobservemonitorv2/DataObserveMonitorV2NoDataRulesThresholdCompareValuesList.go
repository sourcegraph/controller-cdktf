package dataobservemonitorv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/observe/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/observe/dataobservemonitorv2/internal"
)

type DataObserveMonitorV2NoDataRulesThresholdCompareValuesList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataObserveMonitorV2NoDataRulesThresholdCompareValuesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataObserveMonitorV2NoDataRulesThresholdCompareValuesList
type jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareValuesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareValuesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareValuesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareValuesList) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareValuesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareValuesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareValuesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

func NewDataObserveMonitorV2NoDataRulesThresholdCompareValuesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataObserveMonitorV2NoDataRulesThresholdCompareValuesList {
	_init_.Initialize()

	if err := validateNewDataObserveMonitorV2NoDataRulesThresholdCompareValuesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareValuesList{}

	_jsii_.Create(
		"@cdktf/provider-observe.dataObserveMonitorV2.DataObserveMonitorV2NoDataRulesThresholdCompareValuesList",
		[]any{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataObserveMonitorV2NoDataRulesThresholdCompareValuesList_Override(d DataObserveMonitorV2NoDataRulesThresholdCompareValuesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-observe.dataObserveMonitorV2.DataObserveMonitorV2NoDataRulesThresholdCompareValuesList",
		[]any{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareValuesList) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareValuesList) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareValuesList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareValuesList) SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareValuesList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]any{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareValuesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareValuesList) Get(index *float64) DataObserveMonitorV2NoDataRulesThresholdCompareValuesOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataObserveMonitorV2NoDataRulesThresholdCompareValuesOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]any{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareValuesList) Resolve(_context cdktf.IResolveContext) any {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		d,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorV2NoDataRulesThresholdCompareValuesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

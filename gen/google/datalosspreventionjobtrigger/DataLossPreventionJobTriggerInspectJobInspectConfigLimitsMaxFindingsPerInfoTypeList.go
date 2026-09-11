package datalosspreventionjobtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/datalosspreventionjobtrigger/internal"
)

type DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList
type jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList {
	_init_.Initialize()

	if err := validateNewDataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList{}

	_jsii_.Create(
		"@cdktn/provider-google.dataLossPreventionJobTrigger.DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList_Override(d DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataLossPreventionJobTrigger.DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList) Get(index *float64) DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


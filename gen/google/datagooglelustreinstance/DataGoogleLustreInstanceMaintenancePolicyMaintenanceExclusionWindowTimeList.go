package datagooglelustreinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/datagooglelustreinstance/internal"
)

type DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(index *float64) DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList
type jsiiProxy_DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList {
	_init_.Initialize()

	if err := validateNewDataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleLustreInstance.DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList_Override(d DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleLustreInstance.DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList) Get(index *float64) DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


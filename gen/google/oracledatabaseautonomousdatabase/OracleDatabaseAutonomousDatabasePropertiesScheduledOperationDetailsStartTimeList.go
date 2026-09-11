package oracledatabaseautonomousdatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/oracledatabaseautonomousdatabase/internal"
)

type OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(index *float64) OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList
type jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewOracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList {
	_init_.Initialize()

	if err := validateNewOracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList{}

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseAutonomousDatabase.OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewOracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList_Override(o OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseAutonomousDatabase.OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := o.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		o,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList) Get(index *float64) OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeOutputReference {
	if err := o.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeOutputReference

	_jsii_.Invoke(
		o,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsStartTimeList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


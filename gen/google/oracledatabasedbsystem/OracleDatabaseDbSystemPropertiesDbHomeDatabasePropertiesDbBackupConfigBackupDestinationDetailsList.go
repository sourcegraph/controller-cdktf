package oracledatabasedbsystem

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/oracledatabasedbsystem/internal"
)

type OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList interface {
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
	Get(index *float64) OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList
type jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

func NewOracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList {
	_init_.Initialize()

	if err := validateNewOracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList{}

	_jsii_.Create(
		"@cdktf/provider-google.oracleDatabaseDbSystem.OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList",
		[]any{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewOracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList_Override(o OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.oracleDatabaseDbSystem.OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList",
		[]any{terraformResource, terraformAttribute, wrapsSet},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList) SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := o.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		o,
		"allWithMapKey",
		[]any{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList) Get(index *float64) OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsOutputReference {
	if err := o.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsOutputReference

	_jsii_.Invoke(
		o,
		"get",
		[]any{index},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList) Resolve(_context cdktf.IResolveContext) any {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		o,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

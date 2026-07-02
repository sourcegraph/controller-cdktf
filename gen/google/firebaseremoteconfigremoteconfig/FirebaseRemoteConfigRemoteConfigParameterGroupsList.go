package firebaseremoteconfigremoteconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/firebaseremoteconfigremoteconfig/internal"
)

type FirebaseRemoteConfigRemoteConfigParameterGroupsList interface {
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
	Get(index *float64) FirebaseRemoteConfigRemoteConfigParameterGroupsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirebaseRemoteConfigRemoteConfigParameterGroupsList
type jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsList) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

func NewFirebaseRemoteConfigRemoteConfigParameterGroupsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FirebaseRemoteConfigRemoteConfigParameterGroupsList {
	_init_.Initialize()

	if err := validateNewFirebaseRemoteConfigRemoteConfigParameterGroupsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsList{}

	_jsii_.Create(
		"@cdktf/provider-google.firebaseRemoteConfigRemoteConfig.FirebaseRemoteConfigRemoteConfigParameterGroupsList",
		[]any{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFirebaseRemoteConfigRemoteConfigParameterGroupsList_Override(f FirebaseRemoteConfigRemoteConfigParameterGroupsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.firebaseRemoteConfigRemoteConfig.FirebaseRemoteConfigRemoteConfigParameterGroupsList",
		[]any{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsList) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsList) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsList) SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := f.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		f,
		"allWithMapKey",
		[]any{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsList) Get(index *float64) FirebaseRemoteConfigRemoteConfigParameterGroupsOutputReference {
	if err := f.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns FirebaseRemoteConfigRemoteConfigParameterGroupsOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]any{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsList) Resolve(_context cdktf.IResolveContext) any {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		f,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

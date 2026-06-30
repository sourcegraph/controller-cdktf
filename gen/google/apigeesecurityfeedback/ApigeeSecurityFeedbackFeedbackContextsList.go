package apigeesecurityfeedback

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/apigeesecurityfeedback/internal"
)

type ApigeeSecurityFeedbackFeedbackContextsList interface {
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
	Get(index *float64) ApigeeSecurityFeedbackFeedbackContextsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApigeeSecurityFeedbackFeedbackContextsList
type jsiiProxy_ApigeeSecurityFeedbackFeedbackContextsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ApigeeSecurityFeedbackFeedbackContextsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityFeedbackFeedbackContextsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityFeedbackFeedbackContextsList) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityFeedbackFeedbackContextsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityFeedbackFeedbackContextsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityFeedbackFeedbackContextsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

func NewApigeeSecurityFeedbackFeedbackContextsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ApigeeSecurityFeedbackFeedbackContextsList {
	_init_.Initialize()

	if err := validateNewApigeeSecurityFeedbackFeedbackContextsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigeeSecurityFeedbackFeedbackContextsList{}

	_jsii_.Create(
		"@cdktf/provider-google.apigeeSecurityFeedback.ApigeeSecurityFeedbackFeedbackContextsList",
		[]any{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewApigeeSecurityFeedbackFeedbackContextsList_Override(a ApigeeSecurityFeedbackFeedbackContextsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.apigeeSecurityFeedback.ApigeeSecurityFeedbackFeedbackContextsList",
		[]any{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_ApigeeSecurityFeedbackFeedbackContextsList) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApigeeSecurityFeedbackFeedbackContextsList) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApigeeSecurityFeedbackFeedbackContextsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApigeeSecurityFeedbackFeedbackContextsList) SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_ApigeeSecurityFeedbackFeedbackContextsList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := a.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		a,
		"allWithMapKey",
		[]any{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeSecurityFeedbackFeedbackContextsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeSecurityFeedbackFeedbackContextsList) Get(index *float64) ApigeeSecurityFeedbackFeedbackContextsOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ApigeeSecurityFeedbackFeedbackContextsOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]any{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeSecurityFeedbackFeedbackContextsList) Resolve(_context cdktf.IResolveContext) any {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		a,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeSecurityFeedbackFeedbackContextsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

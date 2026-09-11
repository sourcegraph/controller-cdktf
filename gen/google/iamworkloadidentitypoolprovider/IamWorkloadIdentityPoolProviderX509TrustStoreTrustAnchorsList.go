package iamworkloadidentitypoolprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/iamworkloadidentitypoolprovider/internal"
)

type IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList interface {
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
	Get(index *float64) IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList
type jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewIamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList {
	_init_.Initialize()

	if err := validateNewIamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList{}

	_jsii_.Create(
		"@cdktn/provider-google.iamWorkloadIdentityPoolProvider.IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewIamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList_Override(i IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.iamWorkloadIdentityPoolProvider.IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		i,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := i.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		i,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList) Get(index *float64) IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsOutputReference {
	if err := i.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsOutputReference

	_jsii_.Invoke(
		i,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


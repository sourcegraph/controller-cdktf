package networkservicesmulticastdomaingroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/networkservicesmulticastdomaingroup/internal"
)

type NetworkServicesMulticastDomainGroupStateList interface {
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
	Get(index *float64) NetworkServicesMulticastDomainGroupStateOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkServicesMulticastDomainGroupStateList
type jsiiProxy_NetworkServicesMulticastDomainGroupStateList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_NetworkServicesMulticastDomainGroupStateList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesMulticastDomainGroupStateList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesMulticastDomainGroupStateList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesMulticastDomainGroupStateList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesMulticastDomainGroupStateList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

func NewNetworkServicesMulticastDomainGroupStateList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) NetworkServicesMulticastDomainGroupStateList {
	_init_.Initialize()

	if err := validateNewNetworkServicesMulticastDomainGroupStateListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkServicesMulticastDomainGroupStateList{}

	_jsii_.Create(
		"@cdktf/provider-google.networkServicesMulticastDomainGroup.NetworkServicesMulticastDomainGroupStateList",
		[]any{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewNetworkServicesMulticastDomainGroupStateList_Override(n NetworkServicesMulticastDomainGroupStateList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkServicesMulticastDomainGroup.NetworkServicesMulticastDomainGroupStateList",
		[]any{terraformResource, terraformAttribute, wrapsSet},
		n,
	)
}

func (j *jsiiProxy_NetworkServicesMulticastDomainGroupStateList) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesMulticastDomainGroupStateList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesMulticastDomainGroupStateList) SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (n *jsiiProxy_NetworkServicesMulticastDomainGroupStateList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := n.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		n,
		"allWithMapKey",
		[]any{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesMulticastDomainGroupStateList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesMulticastDomainGroupStateList) Get(index *float64) NetworkServicesMulticastDomainGroupStateOutputReference {
	if err := n.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns NetworkServicesMulticastDomainGroupStateOutputReference

	_jsii_.Invoke(
		n,
		"get",
		[]any{index},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesMulticastDomainGroupStateList) Resolve(_context cdktf.IResolveContext) any {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		n,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesMulticastDomainGroupStateList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

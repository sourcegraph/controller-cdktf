package gkeonprembaremetaladmincluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/gkeonprembaremetaladmincluster/internal"
)

type GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList interface {
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
	Get(index *float64) GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList
type jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList {
	_init_.Initialize()

	if err := validateNewGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList{}

	_jsii_.Create(
		"@cdktn/provider-google.gkeonpremBareMetalAdminCluster.GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList_Override(g GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.gkeonpremBareMetalAdminCluster.GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := g.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		g,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList) Get(index *float64) GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


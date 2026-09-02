package gkehubfleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/gkehubfleet/internal"
)

type GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList interface {
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
	Get(index *float64) GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList
type jsiiProxy_GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList {
	_init_.Initialize()

	if err := validateNewGkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList{}

	_jsii_.Create(
		"@cdktn/provider-google.gkeHubFleet.GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList_Override(g GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.gkeHubFleet.GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (g *jsiiProxy_GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList) Get(index *float64) GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


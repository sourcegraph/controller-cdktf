package googlecomputeurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeurlmap/internal"
)

type GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList
type jsiiProxy_GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList {
	_init_.Initialize()

	if err := validateNewGoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList{}

	_jsii_.Create(
		"google-beta.googleComputeUrlMap.GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList_Override(g GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeUrlMap.GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList) Get(index *float64) GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


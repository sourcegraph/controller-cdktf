package connectuserhierarchygroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/aws/connectuserhierarchygroup/internal"
)

type ConnectUserHierarchyGroupHierarchyPathLevelTwoList interface {
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
	Get(index *float64) ConnectUserHierarchyGroupHierarchyPathLevelTwoOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectUserHierarchyGroupHierarchyPathLevelTwoList
type jsiiProxy_ConnectUserHierarchyGroupHierarchyPathLevelTwoList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_ConnectUserHierarchyGroupHierarchyPathLevelTwoList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyGroupHierarchyPathLevelTwoList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyGroupHierarchyPathLevelTwoList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyGroupHierarchyPathLevelTwoList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyGroupHierarchyPathLevelTwoList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewConnectUserHierarchyGroupHierarchyPathLevelTwoList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ConnectUserHierarchyGroupHierarchyPathLevelTwoList {
	_init_.Initialize()

	if err := validateNewConnectUserHierarchyGroupHierarchyPathLevelTwoListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectUserHierarchyGroupHierarchyPathLevelTwoList{}

	_jsii_.Create(
		"@cdktn/provider-aws.connectUserHierarchyGroup.ConnectUserHierarchyGroupHierarchyPathLevelTwoList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewConnectUserHierarchyGroupHierarchyPathLevelTwoList_Override(c ConnectUserHierarchyGroupHierarchyPathLevelTwoList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.connectUserHierarchyGroup.ConnectUserHierarchyGroupHierarchyPathLevelTwoList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_ConnectUserHierarchyGroupHierarchyPathLevelTwoList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectUserHierarchyGroupHierarchyPathLevelTwoList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ConnectUserHierarchyGroupHierarchyPathLevelTwoList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_ConnectUserHierarchyGroupHierarchyPathLevelTwoList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := c.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		c,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserHierarchyGroupHierarchyPathLevelTwoList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserHierarchyGroupHierarchyPathLevelTwoList) Get(index *float64) ConnectUserHierarchyGroupHierarchyPathLevelTwoOutputReference {
	if err := c.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ConnectUserHierarchyGroupHierarchyPathLevelTwoOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserHierarchyGroupHierarchyPathLevelTwoList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserHierarchyGroupHierarchyPathLevelTwoList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


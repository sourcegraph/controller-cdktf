package cesappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/cesappversion/internal"
)

type CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList interface {
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
	Get(index *float64) CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList
type jsiiProxy_CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList {
	_init_.Initialize()

	if err := validateNewCesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList{}

	_jsii_.Create(
		"@cdktn/provider-google.cesAppVersion.CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList_Override(c CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cesAppVersion.CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (c *jsiiProxy_CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList) Get(index *float64) CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackOutputReference {
	if err := c.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CesAppVersionSnapshotGuardrailsCodeCallbackBeforeAgentCallbackList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


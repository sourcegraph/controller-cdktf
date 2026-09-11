package emrinstancefleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/aws/emrinstancefleet/internal"
)

type EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList interface {
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
	Get(index *float64) EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList
type jsiiProxy_EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList {
	_init_.Initialize()

	if err := validateNewEmrInstanceFleetLaunchSpecificationsOnDemandSpecificationListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList{}

	_jsii_.Create(
		"@cdktn/provider-aws.emrInstanceFleet.EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList_Override(e EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.emrInstanceFleet.EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := e.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		e,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList) Get(index *float64) EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationOutputReference {
	if err := e.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


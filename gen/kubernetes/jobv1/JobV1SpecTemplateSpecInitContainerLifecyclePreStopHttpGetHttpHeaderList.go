package jobv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/jobv1/internal"
)

type JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList interface {
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
	Get(index *float64) JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList
type jsiiProxy_JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewJobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList {
	_init_.Initialize()

	if err := validateNewJobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList{}

	_jsii_.Create(
		"@cdktn/provider-kubernetes.jobV1.JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewJobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList_Override(j JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-kubernetes.jobV1.JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		j,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := j.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		j,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList) Get(index *float64) JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderOutputReference {
	if err := j.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderOutputReference

	_jsii_.Invoke(
		j,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := j.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeaderList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


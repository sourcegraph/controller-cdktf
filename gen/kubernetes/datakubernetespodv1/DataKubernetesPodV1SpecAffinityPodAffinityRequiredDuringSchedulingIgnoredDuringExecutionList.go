package datakubernetespodv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/datakubernetespodv1/internal"
)

type DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList interface {
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
	Get(index *float64) DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList
type jsiiProxy_DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList {
	_init_.Initialize()

	if err := validateNewDataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList{}

	_jsii_.Create(
		"@cdktn/provider-kubernetes.dataKubernetesPodV1.DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList_Override(d DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-kubernetes.dataKubernetesPodV1.DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList) Get(index *float64) DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


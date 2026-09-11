package horizontalpodautoscalerv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/horizontalpodautoscalerv2/internal"
)

type HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList interface {
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
	Get(index *float64) HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList
type jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewHorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList {
	_init_.Initialize()

	if err := validateNewHorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList{}

	_jsii_.Create(
		"@cdktn/provider-kubernetes.horizontalPodAutoscalerV2.HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewHorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList_Override(h HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-kubernetes.horizontalPodAutoscalerV2.HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		h,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := h.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		h,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList) Get(index *float64) HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyOutputReference {
	if err := h.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyOutputReference

	_jsii_.Invoke(
		h,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := h.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


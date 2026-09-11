package autoscalingpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/aws/autoscalingpolicy/internal"
)

type AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList interface {
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
	Get(index *float64) AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList
type jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList {
	_init_.Initialize()

	if err := validateNewAutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList{}

	_jsii_.Create(
		"@cdktn/provider-aws.autoscalingPolicy.AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList_Override(a AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.autoscalingPolicy.AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := a.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		a,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList) Get(index *float64) AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


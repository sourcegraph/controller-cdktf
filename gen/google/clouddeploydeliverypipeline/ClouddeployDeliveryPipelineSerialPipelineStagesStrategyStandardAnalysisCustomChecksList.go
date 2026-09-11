package clouddeploydeliverypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/clouddeploydeliverypipeline/internal"
)

type ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList interface {
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
	Get(index *float64) ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList
type jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList {
	_init_.Initialize()

	if err := validateNewClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList{}

	_jsii_.Create(
		"@cdktn/provider-google.clouddeployDeliveryPipeline.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList_Override(c ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.clouddeployDeliveryPipeline.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList) Get(index *float64) ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksOutputReference {
	if err := c.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisCustomChecksList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


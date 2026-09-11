package googleclouddeploydeliverypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleclouddeploydeliverypipeline/internal"
)

type GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList interface {
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
	Get(index *float64) GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList
type jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList {
	_init_.Initialize()

	if err := validateNewGoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleClouddeployDeliveryPipeline.GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList_Override(g GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleClouddeployDeliveryPipeline.GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList) Get(index *float64) GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentVerifyConfigTasksList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


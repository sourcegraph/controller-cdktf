package googlevertexaiendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevertexaiendpoint/internal"
)

type GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(index *float64) GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList
type jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiEndpoint.GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList_Override(g GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiEndpoint.GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList) Get(index *float64) GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


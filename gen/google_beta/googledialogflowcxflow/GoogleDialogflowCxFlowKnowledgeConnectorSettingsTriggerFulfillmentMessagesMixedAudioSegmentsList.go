package googledialogflowcxflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowcxflow/internal"
)

type GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList interface {
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
	Get(index *float64) GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList
type jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowCxFlow.GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList_Override(g GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowCxFlow.GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList) Get(index *float64) GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioSegmentsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


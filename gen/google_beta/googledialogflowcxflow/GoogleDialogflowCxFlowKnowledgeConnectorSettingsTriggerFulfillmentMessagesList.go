package googledialogflowcxflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowcxflow/internal"
)

type GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList interface {
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
	Get(index *float64) GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList
type jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowCxFlow.GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList_Override(g GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowCxFlow.GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList) Get(index *float64) GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


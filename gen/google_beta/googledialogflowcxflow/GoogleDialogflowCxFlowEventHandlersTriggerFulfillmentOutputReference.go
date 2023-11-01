package googledialogflowcxflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowcxflow/internal"
)

type GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ConditionalCases() GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentConditionalCasesList
	ConditionalCasesInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDialogflowCxFlowEventHandlersTriggerFulfillment
	SetInternalValue(val *GoogleDialogflowCxFlowEventHandlersTriggerFulfillment)
	Messages() GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentMessagesList
	MessagesInput() interface{}
	ReturnPartialResponses() interface{}
	SetReturnPartialResponses(val interface{})
	ReturnPartialResponsesInput() interface{}
	SetParameterActions() GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentSetParameterActionsList
	SetParameterActionsInput() interface{}
	Tag() *string
	SetTag(val *string)
	TagInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Webhook() *string
	SetWebhook(val *string)
	WebhookInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutConditionalCases(value interface{})
	PutMessages(value interface{})
	PutSetParameterActions(value interface{})
	ResetConditionalCases()
	ResetMessages()
	ResetReturnPartialResponses()
	ResetSetParameterActions()
	ResetTag()
	ResetWebhook()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference
type jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) ConditionalCases() GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentConditionalCasesList {
	var returns GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentConditionalCasesList
	_jsii_.Get(
		j,
		"conditionalCases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) ConditionalCasesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionalCasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) InternalValue() *GoogleDialogflowCxFlowEventHandlersTriggerFulfillment {
	var returns *GoogleDialogflowCxFlowEventHandlersTriggerFulfillment
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) Messages() GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentMessagesList {
	var returns GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentMessagesList
	_jsii_.Get(
		j,
		"messages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) MessagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) ReturnPartialResponses() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnPartialResponses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) ReturnPartialResponsesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnPartialResponsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) SetParameterActions() GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentSetParameterActionsList {
	var returns GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentSetParameterActionsList
	_jsii_.Get(
		j,
		"setParameterActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) SetParameterActionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setParameterActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) Tag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) TagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) Webhook() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) WebhookInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookInput",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxFlow.GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference_Override(g GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxFlow.GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference)SetInternalValue(val *GoogleDialogflowCxFlowEventHandlersTriggerFulfillment) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference)SetReturnPartialResponses(val interface{}) {
	if err := j.validateSetReturnPartialResponsesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnPartialResponses",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference)SetTag(val *string) {
	if err := j.validateSetTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tag",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference)SetWebhook(val *string) {
	if err := j.validateSetWebhookParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhook",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) PutConditionalCases(value interface{}) {
	if err := g.validatePutConditionalCasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConditionalCases",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) PutMessages(value interface{}) {
	if err := g.validatePutMessagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMessages",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) PutSetParameterActions(value interface{}) {
	if err := g.validatePutSetParameterActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSetParameterActions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) ResetConditionalCases() {
	_jsii_.InvokeVoid(
		g,
		"resetConditionalCases",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) ResetMessages() {
	_jsii_.InvokeVoid(
		g,
		"resetMessages",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) ResetReturnPartialResponses() {
	_jsii_.InvokeVoid(
		g,
		"resetReturnPartialResponses",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) ResetSetParameterActions() {
	_jsii_.InvokeVoid(
		g,
		"resetSetParameterActions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) ResetTag() {
	_jsii_.InvokeVoid(
		g,
		"resetTag",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) ResetWebhook() {
	_jsii_.InvokeVoid(
		g,
		"resetWebhook",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowEventHandlersTriggerFulfillmentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


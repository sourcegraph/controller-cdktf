package googledialogflowcxflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowcxflow/internal"
)

type GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference interface {
	cdktf.ComplexObject
	Channel() *string
	SetChannel(val *string)
	ChannelInput() *string
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
	ConversationSuccess() GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccessOutputReference
	ConversationSuccessInput() *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccess
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EndInteraction() GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesEndInteractionList
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KnowledgeInfoCard() GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCardOutputReference
	KnowledgeInfoCardInput() *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCard
	LiveAgentHandoff() GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoffOutputReference
	LiveAgentHandoffInput() *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoff
	MixedAudio() GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioList
	OutputAudioText() GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioTextOutputReference
	OutputAudioTextInput() *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioText
	Payload() *string
	SetPayload(val *string)
	PayloadInput() *string
	PlayAudio() GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudioOutputReference
	PlayAudioInput() *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudio
	TelephonyTransferCall() GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCallOutputReference
	TelephonyTransferCallInput() *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCall
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Text() GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTextOutputReference
	TextInput() *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesText
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
	PutConversationSuccess(value *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccess)
	PutKnowledgeInfoCard(value *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCard)
	PutLiveAgentHandoff(value *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoff)
	PutOutputAudioText(value *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioText)
	PutPlayAudio(value *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudio)
	PutTelephonyTransferCall(value *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCall)
	PutText(value *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesText)
	ResetChannel()
	ResetConversationSuccess()
	ResetKnowledgeInfoCard()
	ResetLiveAgentHandoff()
	ResetOutputAudioText()
	ResetPayload()
	ResetPlayAudio()
	ResetTelephonyTransferCall()
	ResetText()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference
type jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) Channel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ConversationSuccess() GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccessOutputReference {
	var returns GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccessOutputReference
	_jsii_.Get(
		j,
		"conversationSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ConversationSuccessInput() *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccess {
	var returns *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccess
	_jsii_.Get(
		j,
		"conversationSuccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) EndInteraction() GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesEndInteractionList {
	var returns GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesEndInteractionList
	_jsii_.Get(
		j,
		"endInteraction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) KnowledgeInfoCard() GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCardOutputReference {
	var returns GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCardOutputReference
	_jsii_.Get(
		j,
		"knowledgeInfoCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) KnowledgeInfoCardInput() *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCard {
	var returns *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCard
	_jsii_.Get(
		j,
		"knowledgeInfoCardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) LiveAgentHandoff() GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoffOutputReference {
	var returns GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoffOutputReference
	_jsii_.Get(
		j,
		"liveAgentHandoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) LiveAgentHandoffInput() *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoff {
	var returns *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoff
	_jsii_.Get(
		j,
		"liveAgentHandoffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) MixedAudio() GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioList {
	var returns GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioList
	_jsii_.Get(
		j,
		"mixedAudio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) OutputAudioText() GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioTextOutputReference {
	var returns GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioTextOutputReference
	_jsii_.Get(
		j,
		"outputAudioText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) OutputAudioTextInput() *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioText {
	var returns *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioText
	_jsii_.Get(
		j,
		"outputAudioTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) Payload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PayloadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PlayAudio() GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudioOutputReference {
	var returns GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudioOutputReference
	_jsii_.Get(
		j,
		"playAudio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PlayAudioInput() *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudio {
	var returns *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudio
	_jsii_.Get(
		j,
		"playAudioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) TelephonyTransferCall() GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCallOutputReference {
	var returns GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCallOutputReference
	_jsii_.Get(
		j,
		"telephonyTransferCall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) TelephonyTransferCallInput() *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCall {
	var returns *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCall
	_jsii_.Get(
		j,
		"telephonyTransferCallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) Text() GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTextOutputReference {
	var returns GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTextOutputReference
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) TextInput() *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesText {
	var returns *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesText
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxFlow.GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference_Override(g GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxFlow.GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference)SetChannel(val *string) {
	if err := j.validateSetChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channel",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference)SetPayload(val *string) {
	if err := j.validateSetPayloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payload",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PutConversationSuccess(value *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccess) {
	if err := g.validatePutConversationSuccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConversationSuccess",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PutKnowledgeInfoCard(value *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCard) {
	if err := g.validatePutKnowledgeInfoCardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKnowledgeInfoCard",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PutLiveAgentHandoff(value *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoff) {
	if err := g.validatePutLiveAgentHandoffParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLiveAgentHandoff",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PutOutputAudioText(value *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioText) {
	if err := g.validatePutOutputAudioTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOutputAudioText",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PutPlayAudio(value *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudio) {
	if err := g.validatePutPlayAudioParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPlayAudio",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PutTelephonyTransferCall(value *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCall) {
	if err := g.validatePutTelephonyTransferCallParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTelephonyTransferCall",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PutText(value *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesText) {
	if err := g.validatePutTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putText",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetChannel() {
	_jsii_.InvokeVoid(
		g,
		"resetChannel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetConversationSuccess() {
	_jsii_.InvokeVoid(
		g,
		"resetConversationSuccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetKnowledgeInfoCard() {
	_jsii_.InvokeVoid(
		g,
		"resetKnowledgeInfoCard",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetLiveAgentHandoff() {
	_jsii_.InvokeVoid(
		g,
		"resetLiveAgentHandoff",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetOutputAudioText() {
	_jsii_.InvokeVoid(
		g,
		"resetOutputAudioText",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetPayload() {
	_jsii_.InvokeVoid(
		g,
		"resetPayload",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetPlayAudio() {
	_jsii_.InvokeVoid(
		g,
		"resetPlayAudio",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetTelephonyTransferCall() {
	_jsii_.InvokeVoid(
		g,
		"resetTelephonyTransferCall",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetText() {
	_jsii_.InvokeVoid(
		g,
		"resetText",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


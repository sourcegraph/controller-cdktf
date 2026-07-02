package dialogflowcxpage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/dialogflowcxpage/internal"
)

type DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference interface {
	cdktf.ComplexObject
	Channel() *string
	SetChannel(val *string)
	ChannelInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ConversationSuccess() DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesConversationSuccessOutputReference
	ConversationSuccessInput() *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesConversationSuccess
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	LiveAgentHandoff() DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoffOutputReference
	LiveAgentHandoffInput() *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoff
	OutputAudioText() DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputAudioTextOutputReference
	OutputAudioTextInput() *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputAudioText
	Payload() *string
	SetPayload(val *string)
	PayloadInput() *string
	PlayAudio() DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesPlayAudioOutputReference
	PlayAudioInput() *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesPlayAudio
	TelephonyTransferCall() DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCallOutputReference
	TelephonyTransferCallInput() *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCall
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Text() DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesTextOutputReference
	TextInput() *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesText
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutConversationSuccess(value *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesConversationSuccess)
	PutLiveAgentHandoff(value *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoff)
	PutOutputAudioText(value *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputAudioText)
	PutPlayAudio(value *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesPlayAudio)
	PutTelephonyTransferCall(value *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCall)
	PutText(value *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesText)
	ResetChannel()
	ResetConversationSuccess()
	ResetLiveAgentHandoff()
	ResetOutputAudioText()
	ResetPayload()
	ResetPlayAudio()
	ResetTelephonyTransferCall()
	ResetText()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference
type jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) Channel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) ChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) ConversationSuccess() DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesConversationSuccessOutputReference {
	var returns DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesConversationSuccessOutputReference
	_jsii_.Get(
		j,
		"conversationSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) ConversationSuccessInput() *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesConversationSuccess {
	var returns *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesConversationSuccess
	_jsii_.Get(
		j,
		"conversationSuccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) LiveAgentHandoff() DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoffOutputReference {
	var returns DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoffOutputReference
	_jsii_.Get(
		j,
		"liveAgentHandoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) LiveAgentHandoffInput() *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoff {
	var returns *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoff
	_jsii_.Get(
		j,
		"liveAgentHandoffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) OutputAudioText() DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputAudioTextOutputReference {
	var returns DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputAudioTextOutputReference
	_jsii_.Get(
		j,
		"outputAudioText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) OutputAudioTextInput() *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputAudioText {
	var returns *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputAudioText
	_jsii_.Get(
		j,
		"outputAudioTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) Payload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) PayloadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) PlayAudio() DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesPlayAudioOutputReference {
	var returns DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesPlayAudioOutputReference
	_jsii_.Get(
		j,
		"playAudio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) PlayAudioInput() *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesPlayAudio {
	var returns *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesPlayAudio
	_jsii_.Get(
		j,
		"playAudioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) TelephonyTransferCall() DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCallOutputReference {
	var returns DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCallOutputReference
	_jsii_.Get(
		j,
		"telephonyTransferCall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) TelephonyTransferCallInput() *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCall {
	var returns *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCall
	_jsii_.Get(
		j,
		"telephonyTransferCallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) Text() DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesTextOutputReference {
	var returns DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesTextOutputReference
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) TextInput() *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesText {
	var returns *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesText
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}

func NewDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxPage.DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference_Override(d DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxPage.DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) SetChannel(val *string) {
	if err := j.validateSetChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channel",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) SetPayload(val *string) {
	if err := j.validateSetPayloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payload",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) PutConversationSuccess(value *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesConversationSuccess) {
	if err := d.validatePutConversationSuccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConversationSuccess",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) PutLiveAgentHandoff(value *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoff) {
	if err := d.validatePutLiveAgentHandoffParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLiveAgentHandoff",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) PutOutputAudioText(value *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputAudioText) {
	if err := d.validatePutOutputAudioTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOutputAudioText",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) PutPlayAudio(value *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesPlayAudio) {
	if err := d.validatePutPlayAudioParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPlayAudio",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) PutTelephonyTransferCall(value *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCall) {
	if err := d.validatePutTelephonyTransferCallParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTelephonyTransferCall",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) PutText(value *DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesText) {
	if err := d.validatePutTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putText",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) ResetChannel() {
	_jsii_.InvokeVoid(
		d,
		"resetChannel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) ResetConversationSuccess() {
	_jsii_.InvokeVoid(
		d,
		"resetConversationSuccess",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) ResetLiveAgentHandoff() {
	_jsii_.InvokeVoid(
		d,
		"resetLiveAgentHandoff",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) ResetOutputAudioText() {
	_jsii_.InvokeVoid(
		d,
		"resetOutputAudioText",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) ResetPayload() {
	_jsii_.InvokeVoid(
		d,
		"resetPayload",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) ResetPlayAudio() {
	_jsii_.InvokeVoid(
		d,
		"resetPlayAudio",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) ResetTelephonyTransferCall() {
	_jsii_.InvokeVoid(
		d,
		"resetTelephonyTransferCall",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) ResetText() {
	_jsii_.InvokeVoid(
		d,
		"resetText",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		d,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

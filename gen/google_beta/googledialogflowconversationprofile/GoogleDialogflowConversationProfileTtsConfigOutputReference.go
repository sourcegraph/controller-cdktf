package googledialogflowconversationprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowconversationprofile/internal"
)

type GoogleDialogflowConversationProfileTtsConfigOutputReference interface {
	cdktn.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EffectsProfileId() *[]*string
	SetEffectsProfileId(val *[]*string)
	EffectsProfileIdInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDialogflowConversationProfileTtsConfig
	SetInternalValue(val *GoogleDialogflowConversationProfileTtsConfig)
	Pitch() *float64
	SetPitch(val *float64)
	PitchInput() *float64
	SpeakingRate() *float64
	SetSpeakingRate(val *float64)
	SpeakingRateInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Voice() GoogleDialogflowConversationProfileTtsConfigVoiceOutputReference
	VoiceInput() *GoogleDialogflowConversationProfileTtsConfigVoice
	VolumeGainDb() *float64
	SetVolumeGainDb(val *float64)
	VolumeGainDbInput() *float64
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutVoice(value *GoogleDialogflowConversationProfileTtsConfigVoice)
	ResetEffectsProfileId()
	ResetPitch()
	ResetSpeakingRate()
	ResetVoice()
	ResetVolumeGainDb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowConversationProfileTtsConfigOutputReference
type jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) EffectsProfileId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"effectsProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) EffectsProfileIdInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"effectsProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) InternalValue() *GoogleDialogflowConversationProfileTtsConfig {
	var returns *GoogleDialogflowConversationProfileTtsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) Pitch() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pitch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) PitchInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pitchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) SpeakingRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"speakingRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) SpeakingRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"speakingRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) Voice() GoogleDialogflowConversationProfileTtsConfigVoiceOutputReference {
	var returns GoogleDialogflowConversationProfileTtsConfigVoiceOutputReference
	_jsii_.Get(
		j,
		"voice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) VoiceInput() *GoogleDialogflowConversationProfileTtsConfigVoice {
	var returns *GoogleDialogflowConversationProfileTtsConfigVoice
	_jsii_.Get(
		j,
		"voiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) VolumeGainDb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeGainDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) VolumeGainDbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeGainDbInput",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowConversationProfileTtsConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowConversationProfileTtsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowConversationProfileTtsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowConversationProfile.GoogleDialogflowConversationProfileTtsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowConversationProfileTtsConfigOutputReference_Override(g GoogleDialogflowConversationProfileTtsConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowConversationProfile.GoogleDialogflowConversationProfileTtsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference)SetEffectsProfileId(val *[]*string) {
	if err := j.validateSetEffectsProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"effectsProfileId",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference)SetInternalValue(val *GoogleDialogflowConversationProfileTtsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference)SetPitch(val *float64) {
	if err := j.validateSetPitchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pitch",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference)SetSpeakingRate(val *float64) {
	if err := j.validateSetSpeakingRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"speakingRate",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference)SetVolumeGainDb(val *float64) {
	if err := j.validateSetVolumeGainDbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeGainDb",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) PutVoice(value *GoogleDialogflowConversationProfileTtsConfigVoice) {
	if err := g.validatePutVoiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVoice",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) ResetEffectsProfileId() {
	_jsii_.InvokeVoid(
		g,
		"resetEffectsProfileId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) ResetPitch() {
	_jsii_.InvokeVoid(
		g,
		"resetPitch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) ResetSpeakingRate() {
	_jsii_.InvokeVoid(
		g,
		"resetSpeakingRate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) ResetVoice() {
	_jsii_.InvokeVoid(
		g,
		"resetVoice",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) ResetVolumeGainDb() {
	_jsii_.InvokeVoid(
		g,
		"resetVolumeGainDb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileTtsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


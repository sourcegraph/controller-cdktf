package googledialogflowenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowenvironment/internal"
)

type GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Language() *string
	SetLanguage(val *string)
	LanguageInput() *string
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
	Voice() GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsVoiceOutputReference
	VoiceInput() *GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsVoice
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
	PutVoice(value *GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsVoice)
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

// The jsii proxy struct for GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference
type jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) EffectsProfileId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"effectsProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) EffectsProfileIdInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"effectsProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) Language() *string {
	var returns *string
	_jsii_.Get(
		j,
		"language",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) LanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) Pitch() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pitch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) PitchInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pitchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) SpeakingRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"speakingRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) SpeakingRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"speakingRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) Voice() GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsVoiceOutputReference {
	var returns GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsVoiceOutputReference
	_jsii_.Get(
		j,
		"voice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) VoiceInput() *GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsVoice {
	var returns *GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsVoice
	_jsii_.Get(
		j,
		"voiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) VolumeGainDb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeGainDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) VolumeGainDbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeGainDbInput",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowEnvironment.GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference_Override(g GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowEnvironment.GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference)SetEffectsProfileId(val *[]*string) {
	if err := j.validateSetEffectsProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"effectsProfileId",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference)SetLanguage(val *string) {
	if err := j.validateSetLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"language",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference)SetPitch(val *float64) {
	if err := j.validateSetPitchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pitch",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference)SetSpeakingRate(val *float64) {
	if err := j.validateSetSpeakingRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"speakingRate",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference)SetVolumeGainDb(val *float64) {
	if err := j.validateSetVolumeGainDbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeGainDb",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) PutVoice(value *GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsVoice) {
	if err := g.validatePutVoiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVoice",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) ResetEffectsProfileId() {
	_jsii_.InvokeVoid(
		g,
		"resetEffectsProfileId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) ResetPitch() {
	_jsii_.InvokeVoid(
		g,
		"resetPitch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) ResetSpeakingRate() {
	_jsii_.InvokeVoid(
		g,
		"resetSpeakingRate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) ResetVoice() {
	_jsii_.InvokeVoid(
		g,
		"resetVoice",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) ResetVolumeGainDb() {
	_jsii_.InvokeVoid(
		g,
		"resetVolumeGainDb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


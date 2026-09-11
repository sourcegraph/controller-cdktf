package googledialogflowenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowenvironment/internal"
)

type GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference interface {
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
	EnableTextToSpeech() interface{}
	SetEnableTextToSpeech(val interface{})
	EnableTextToSpeechInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDialogflowEnvironmentTextToSpeechSettings
	SetInternalValue(val *GoogleDialogflowEnvironmentTextToSpeechSettings)
	OutputAudioEncoding() *string
	SetOutputAudioEncoding(val *string)
	OutputAudioEncodingInput() *string
	SampleRateHertz() *float64
	SetSampleRateHertz(val *float64)
	SampleRateHertzInput() *float64
	SynthesizeSpeechConfigs() GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsList
	SynthesizeSpeechConfigsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutSynthesizeSpeechConfigs(value interface{})
	ResetEnableTextToSpeech()
	ResetOutputAudioEncoding()
	ResetSampleRateHertz()
	ResetSynthesizeSpeechConfigs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference
type jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) EnableTextToSpeech() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTextToSpeech",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) EnableTextToSpeechInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTextToSpeechInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) InternalValue() *GoogleDialogflowEnvironmentTextToSpeechSettings {
	var returns *GoogleDialogflowEnvironmentTextToSpeechSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) OutputAudioEncoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputAudioEncoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) OutputAudioEncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputAudioEncodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) SampleRateHertz() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleRateHertz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) SampleRateHertzInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleRateHertzInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) SynthesizeSpeechConfigs() GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsList {
	var returns GoogleDialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsList
	_jsii_.Get(
		j,
		"synthesizeSpeechConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) SynthesizeSpeechConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"synthesizeSpeechConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowEnvironmentTextToSpeechSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowEnvironment.GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference_Override(g GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowEnvironment.GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference)SetEnableTextToSpeech(val interface{}) {
	if err := j.validateSetEnableTextToSpeechParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableTextToSpeech",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference)SetInternalValue(val *GoogleDialogflowEnvironmentTextToSpeechSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference)SetOutputAudioEncoding(val *string) {
	if err := j.validateSetOutputAudioEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputAudioEncoding",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference)SetSampleRateHertz(val *float64) {
	if err := j.validateSetSampleRateHertzParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleRateHertz",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) PutSynthesizeSpeechConfigs(value interface{}) {
	if err := g.validatePutSynthesizeSpeechConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSynthesizeSpeechConfigs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) ResetEnableTextToSpeech() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableTextToSpeech",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) ResetOutputAudioEncoding() {
	_jsii_.InvokeVoid(
		g,
		"resetOutputAudioEncoding",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) ResetSampleRateHertz() {
	_jsii_.InvokeVoid(
		g,
		"resetSampleRateHertz",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) ResetSynthesizeSpeechConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetSynthesizeSpeechConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowEnvironmentTextToSpeechSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


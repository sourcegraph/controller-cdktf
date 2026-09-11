package dialogflowenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/dialogflowenvironment/internal"
)

type DialogflowEnvironmentTextToSpeechSettingsOutputReference interface {
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
	InternalValue() *DialogflowEnvironmentTextToSpeechSettings
	SetInternalValue(val *DialogflowEnvironmentTextToSpeechSettings)
	OutputAudioEncoding() *string
	SetOutputAudioEncoding(val *string)
	OutputAudioEncodingInput() *string
	SampleRateHertz() *float64
	SetSampleRateHertz(val *float64)
	SampleRateHertzInput() *float64
	SynthesizeSpeechConfigs() DialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsList
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

// The jsii proxy struct for DialogflowEnvironmentTextToSpeechSettingsOutputReference
type jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) EnableTextToSpeech() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTextToSpeech",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) EnableTextToSpeechInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTextToSpeechInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) InternalValue() *DialogflowEnvironmentTextToSpeechSettings {
	var returns *DialogflowEnvironmentTextToSpeechSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) OutputAudioEncoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputAudioEncoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) OutputAudioEncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputAudioEncodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) SampleRateHertz() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleRateHertz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) SampleRateHertzInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleRateHertzInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) SynthesizeSpeechConfigs() DialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsList {
	var returns DialogflowEnvironmentTextToSpeechSettingsSynthesizeSpeechConfigsList
	_jsii_.Get(
		j,
		"synthesizeSpeechConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) SynthesizeSpeechConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"synthesizeSpeechConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDialogflowEnvironmentTextToSpeechSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DialogflowEnvironmentTextToSpeechSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowEnvironmentTextToSpeechSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowEnvironment.DialogflowEnvironmentTextToSpeechSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowEnvironmentTextToSpeechSettingsOutputReference_Override(d DialogflowEnvironmentTextToSpeechSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowEnvironment.DialogflowEnvironmentTextToSpeechSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference)SetEnableTextToSpeech(val interface{}) {
	if err := j.validateSetEnableTextToSpeechParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableTextToSpeech",
		val,
	)
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference)SetInternalValue(val *DialogflowEnvironmentTextToSpeechSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference)SetOutputAudioEncoding(val *string) {
	if err := j.validateSetOutputAudioEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputAudioEncoding",
		val,
	)
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference)SetSampleRateHertz(val *float64) {
	if err := j.validateSetSampleRateHertzParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleRateHertz",
		val,
	)
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) PutSynthesizeSpeechConfigs(value interface{}) {
	if err := d.validatePutSynthesizeSpeechConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSynthesizeSpeechConfigs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) ResetEnableTextToSpeech() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableTextToSpeech",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) ResetOutputAudioEncoding() {
	_jsii_.InvokeVoid(
		d,
		"resetOutputAudioEncoding",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) ResetSampleRateHertz() {
	_jsii_.InvokeVoid(
		d,
		"resetSampleRateHertz",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) ResetSynthesizeSpeechConfigs() {
	_jsii_.InvokeVoid(
		d,
		"resetSynthesizeSpeechConfigs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowEnvironmentTextToSpeechSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package googlecesapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesapp/internal"
)

type GoogleCesAppAudioProcessingConfigOutputReference interface {
	cdktf.ComplexObject
	AmbientSoundConfig() GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference
	AmbientSoundConfigInput() *GoogleCesAppAudioProcessingConfigAmbientSoundConfig
	BargeInConfig() GoogleCesAppAudioProcessingConfigBargeInConfigOutputReference
	BargeInConfigInput() *GoogleCesAppAudioProcessingConfigBargeInConfig
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
	// Experimental.
	Fqn() *string
	InactivityTimeout() *string
	SetInactivityTimeout(val *string)
	InactivityTimeoutInput() *string
	InternalValue() *GoogleCesAppAudioProcessingConfig
	SetInternalValue(val *GoogleCesAppAudioProcessingConfig)
	SynthesizeSpeechConfigs() GoogleCesAppAudioProcessingConfigSynthesizeSpeechConfigsList
	SynthesizeSpeechConfigsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutAmbientSoundConfig(value *GoogleCesAppAudioProcessingConfigAmbientSoundConfig)
	PutBargeInConfig(value *GoogleCesAppAudioProcessingConfigBargeInConfig)
	PutSynthesizeSpeechConfigs(value interface{})
	ResetAmbientSoundConfig()
	ResetBargeInConfig()
	ResetInactivityTimeout()
	ResetSynthesizeSpeechConfigs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesAppAudioProcessingConfigOutputReference
type jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) AmbientSoundConfig() GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference {
	var returns GoogleCesAppAudioProcessingConfigAmbientSoundConfigOutputReference
	_jsii_.Get(
		j,
		"ambientSoundConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) AmbientSoundConfigInput() *GoogleCesAppAudioProcessingConfigAmbientSoundConfig {
	var returns *GoogleCesAppAudioProcessingConfigAmbientSoundConfig
	_jsii_.Get(
		j,
		"ambientSoundConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) BargeInConfig() GoogleCesAppAudioProcessingConfigBargeInConfigOutputReference {
	var returns GoogleCesAppAudioProcessingConfigBargeInConfigOutputReference
	_jsii_.Get(
		j,
		"bargeInConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) BargeInConfigInput() *GoogleCesAppAudioProcessingConfigBargeInConfig {
	var returns *GoogleCesAppAudioProcessingConfigBargeInConfig
	_jsii_.Get(
		j,
		"bargeInConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) InactivityTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inactivityTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) InactivityTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inactivityTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) InternalValue() *GoogleCesAppAudioProcessingConfig {
	var returns *GoogleCesAppAudioProcessingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) SynthesizeSpeechConfigs() GoogleCesAppAudioProcessingConfigSynthesizeSpeechConfigsList {
	var returns GoogleCesAppAudioProcessingConfigSynthesizeSpeechConfigsList
	_jsii_.Get(
		j,
		"synthesizeSpeechConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) SynthesizeSpeechConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"synthesizeSpeechConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCesAppAudioProcessingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCesAppAudioProcessingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesAppAudioProcessingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesApp.GoogleCesAppAudioProcessingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesAppAudioProcessingConfigOutputReference_Override(g GoogleCesAppAudioProcessingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesApp.GoogleCesAppAudioProcessingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference)SetInactivityTimeout(val *string) {
	if err := j.validateSetInactivityTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inactivityTimeout",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference)SetInternalValue(val *GoogleCesAppAudioProcessingConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) PutAmbientSoundConfig(value *GoogleCesAppAudioProcessingConfigAmbientSoundConfig) {
	if err := g.validatePutAmbientSoundConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAmbientSoundConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) PutBargeInConfig(value *GoogleCesAppAudioProcessingConfigBargeInConfig) {
	if err := g.validatePutBargeInConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBargeInConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) PutSynthesizeSpeechConfigs(value interface{}) {
	if err := g.validatePutSynthesizeSpeechConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSynthesizeSpeechConfigs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) ResetAmbientSoundConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAmbientSoundConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) ResetBargeInConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetBargeInConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) ResetInactivityTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetInactivityTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) ResetSynthesizeSpeechConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetSynthesizeSpeechConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesAppAudioProcessingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


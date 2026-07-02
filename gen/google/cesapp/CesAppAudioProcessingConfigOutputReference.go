package cesapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cesapp/internal"
)

type CesAppAudioProcessingConfigOutputReference interface {
	cdktf.ComplexObject
	AmbientSoundConfig() CesAppAudioProcessingConfigAmbientSoundConfigOutputReference
	AmbientSoundConfigInput() *CesAppAudioProcessingConfigAmbientSoundConfig
	BargeInConfig() CesAppAudioProcessingConfigBargeInConfigOutputReference
	BargeInConfigInput() *CesAppAudioProcessingConfigBargeInConfig
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
	InternalValue() *CesAppAudioProcessingConfig
	SetInternalValue(val *CesAppAudioProcessingConfig)
	SynthesizeSpeechConfigs() CesAppAudioProcessingConfigSynthesizeSpeechConfigsList
	SynthesizeSpeechConfigsInput() any
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
	PutAmbientSoundConfig(value *CesAppAudioProcessingConfigAmbientSoundConfig)
	PutBargeInConfig(value *CesAppAudioProcessingConfigBargeInConfig)
	PutSynthesizeSpeechConfigs(value any)
	ResetAmbientSoundConfig()
	ResetBargeInConfig()
	ResetInactivityTimeout()
	ResetSynthesizeSpeechConfigs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesAppAudioProcessingConfigOutputReference
type jsiiProxy_CesAppAudioProcessingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CesAppAudioProcessingConfigOutputReference) AmbientSoundConfig() CesAppAudioProcessingConfigAmbientSoundConfigOutputReference {
	var returns CesAppAudioProcessingConfigAmbientSoundConfigOutputReference
	_jsii_.Get(
		j,
		"ambientSoundConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppAudioProcessingConfigOutputReference) AmbientSoundConfigInput() *CesAppAudioProcessingConfigAmbientSoundConfig {
	var returns *CesAppAudioProcessingConfigAmbientSoundConfig
	_jsii_.Get(
		j,
		"ambientSoundConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppAudioProcessingConfigOutputReference) BargeInConfig() CesAppAudioProcessingConfigBargeInConfigOutputReference {
	var returns CesAppAudioProcessingConfigBargeInConfigOutputReference
	_jsii_.Get(
		j,
		"bargeInConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppAudioProcessingConfigOutputReference) BargeInConfigInput() *CesAppAudioProcessingConfigBargeInConfig {
	var returns *CesAppAudioProcessingConfigBargeInConfig
	_jsii_.Get(
		j,
		"bargeInConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppAudioProcessingConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppAudioProcessingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppAudioProcessingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppAudioProcessingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppAudioProcessingConfigOutputReference) InactivityTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inactivityTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppAudioProcessingConfigOutputReference) InactivityTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inactivityTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppAudioProcessingConfigOutputReference) InternalValue() *CesAppAudioProcessingConfig {
	var returns *CesAppAudioProcessingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppAudioProcessingConfigOutputReference) SynthesizeSpeechConfigs() CesAppAudioProcessingConfigSynthesizeSpeechConfigsList {
	var returns CesAppAudioProcessingConfigSynthesizeSpeechConfigsList
	_jsii_.Get(
		j,
		"synthesizeSpeechConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppAudioProcessingConfigOutputReference) SynthesizeSpeechConfigsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"synthesizeSpeechConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppAudioProcessingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppAudioProcessingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCesAppAudioProcessingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CesAppAudioProcessingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCesAppAudioProcessingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesAppAudioProcessingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cesApp.CesAppAudioProcessingConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesAppAudioProcessingConfigOutputReference_Override(c CesAppAudioProcessingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesApp.CesAppAudioProcessingConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesAppAudioProcessingConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesAppAudioProcessingConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesAppAudioProcessingConfigOutputReference) SetInactivityTimeout(val *string) {
	if err := j.validateSetInactivityTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inactivityTimeout",
		val,
	)
}

func (j *jsiiProxy_CesAppAudioProcessingConfigOutputReference) SetInternalValue(val *CesAppAudioProcessingConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesAppAudioProcessingConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesAppAudioProcessingConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesAppAudioProcessingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppAudioProcessingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppAudioProcessingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppAudioProcessingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppAudioProcessingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppAudioProcessingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppAudioProcessingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppAudioProcessingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppAudioProcessingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppAudioProcessingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppAudioProcessingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppAudioProcessingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppAudioProcessingConfigOutputReference) PutAmbientSoundConfig(value *CesAppAudioProcessingConfigAmbientSoundConfig) {
	if err := c.validatePutAmbientSoundConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAmbientSoundConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_CesAppAudioProcessingConfigOutputReference) PutBargeInConfig(value *CesAppAudioProcessingConfigBargeInConfig) {
	if err := c.validatePutBargeInConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBargeInConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_CesAppAudioProcessingConfigOutputReference) PutSynthesizeSpeechConfigs(value any) {
	if err := c.validatePutSynthesizeSpeechConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSynthesizeSpeechConfigs",
		[]any{value},
	)
}

func (c *jsiiProxy_CesAppAudioProcessingConfigOutputReference) ResetAmbientSoundConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetAmbientSoundConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppAudioProcessingConfigOutputReference) ResetBargeInConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetBargeInConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppAudioProcessingConfigOutputReference) ResetInactivityTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetInactivityTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppAudioProcessingConfigOutputReference) ResetSynthesizeSpeechConfigs() {
	_jsii_.InvokeVoid(
		c,
		"resetSynthesizeSpeechConfigs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppAudioProcessingConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		c,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppAudioProcessingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

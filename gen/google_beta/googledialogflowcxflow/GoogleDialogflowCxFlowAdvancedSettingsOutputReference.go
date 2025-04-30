package googledialogflowcxflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowcxflow/internal"
)

type GoogleDialogflowCxFlowAdvancedSettingsOutputReference interface {
	cdktf.ComplexObject
	AudioExportGcsDestination() GoogleDialogflowCxFlowAdvancedSettingsAudioExportGcsDestinationOutputReference
	AudioExportGcsDestinationInput() *GoogleDialogflowCxFlowAdvancedSettingsAudioExportGcsDestination
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
	DtmfSettings() GoogleDialogflowCxFlowAdvancedSettingsDtmfSettingsOutputReference
	DtmfSettingsInput() *GoogleDialogflowCxFlowAdvancedSettingsDtmfSettings
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDialogflowCxFlowAdvancedSettings
	SetInternalValue(val *GoogleDialogflowCxFlowAdvancedSettings)
	LoggingSettings() GoogleDialogflowCxFlowAdvancedSettingsLoggingSettingsOutputReference
	LoggingSettingsInput() *GoogleDialogflowCxFlowAdvancedSettingsLoggingSettings
	SpeechSettings() GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference
	SpeechSettingsInput() *GoogleDialogflowCxFlowAdvancedSettingsSpeechSettings
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
	PutAudioExportGcsDestination(value *GoogleDialogflowCxFlowAdvancedSettingsAudioExportGcsDestination)
	PutDtmfSettings(value *GoogleDialogflowCxFlowAdvancedSettingsDtmfSettings)
	PutLoggingSettings(value *GoogleDialogflowCxFlowAdvancedSettingsLoggingSettings)
	PutSpeechSettings(value *GoogleDialogflowCxFlowAdvancedSettingsSpeechSettings)
	ResetAudioExportGcsDestination()
	ResetDtmfSettings()
	ResetLoggingSettings()
	ResetSpeechSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxFlowAdvancedSettingsOutputReference
type jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) AudioExportGcsDestination() GoogleDialogflowCxFlowAdvancedSettingsAudioExportGcsDestinationOutputReference {
	var returns GoogleDialogflowCxFlowAdvancedSettingsAudioExportGcsDestinationOutputReference
	_jsii_.Get(
		j,
		"audioExportGcsDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) AudioExportGcsDestinationInput() *GoogleDialogflowCxFlowAdvancedSettingsAudioExportGcsDestination {
	var returns *GoogleDialogflowCxFlowAdvancedSettingsAudioExportGcsDestination
	_jsii_.Get(
		j,
		"audioExportGcsDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) DtmfSettings() GoogleDialogflowCxFlowAdvancedSettingsDtmfSettingsOutputReference {
	var returns GoogleDialogflowCxFlowAdvancedSettingsDtmfSettingsOutputReference
	_jsii_.Get(
		j,
		"dtmfSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) DtmfSettingsInput() *GoogleDialogflowCxFlowAdvancedSettingsDtmfSettings {
	var returns *GoogleDialogflowCxFlowAdvancedSettingsDtmfSettings
	_jsii_.Get(
		j,
		"dtmfSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) InternalValue() *GoogleDialogflowCxFlowAdvancedSettings {
	var returns *GoogleDialogflowCxFlowAdvancedSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) LoggingSettings() GoogleDialogflowCxFlowAdvancedSettingsLoggingSettingsOutputReference {
	var returns GoogleDialogflowCxFlowAdvancedSettingsLoggingSettingsOutputReference
	_jsii_.Get(
		j,
		"loggingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) LoggingSettingsInput() *GoogleDialogflowCxFlowAdvancedSettingsLoggingSettings {
	var returns *GoogleDialogflowCxFlowAdvancedSettingsLoggingSettings
	_jsii_.Get(
		j,
		"loggingSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) SpeechSettings() GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference {
	var returns GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference
	_jsii_.Get(
		j,
		"speechSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) SpeechSettingsInput() *GoogleDialogflowCxFlowAdvancedSettingsSpeechSettings {
	var returns *GoogleDialogflowCxFlowAdvancedSettingsSpeechSettings
	_jsii_.Get(
		j,
		"speechSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowCxFlowAdvancedSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowCxFlowAdvancedSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxFlowAdvancedSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxFlow.GoogleDialogflowCxFlowAdvancedSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxFlowAdvancedSettingsOutputReference_Override(g GoogleDialogflowCxFlowAdvancedSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxFlow.GoogleDialogflowCxFlowAdvancedSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference)SetInternalValue(val *GoogleDialogflowCxFlowAdvancedSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) PutAudioExportGcsDestination(value *GoogleDialogflowCxFlowAdvancedSettingsAudioExportGcsDestination) {
	if err := g.validatePutAudioExportGcsDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAudioExportGcsDestination",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) PutDtmfSettings(value *GoogleDialogflowCxFlowAdvancedSettingsDtmfSettings) {
	if err := g.validatePutDtmfSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDtmfSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) PutLoggingSettings(value *GoogleDialogflowCxFlowAdvancedSettingsLoggingSettings) {
	if err := g.validatePutLoggingSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLoggingSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) PutSpeechSettings(value *GoogleDialogflowCxFlowAdvancedSettingsSpeechSettings) {
	if err := g.validatePutSpeechSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSpeechSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) ResetAudioExportGcsDestination() {
	_jsii_.InvokeVoid(
		g,
		"resetAudioExportGcsDestination",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) ResetDtmfSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetDtmfSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) ResetLoggingSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetLoggingSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) ResetSpeechSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetSpeechSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package googlecesapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesapp/internal"
)

type GoogleCesAppLoggingSettingsOutputReference interface {
	cdktn.ComplexObject
	AudioRecordingConfig() GoogleCesAppLoggingSettingsAudioRecordingConfigOutputReference
	AudioRecordingConfigInput() *GoogleCesAppLoggingSettingsAudioRecordingConfig
	BigqueryExportSettings() GoogleCesAppLoggingSettingsBigqueryExportSettingsOutputReference
	BigqueryExportSettingsInput() *GoogleCesAppLoggingSettingsBigqueryExportSettings
	CloudLoggingSettings() GoogleCesAppLoggingSettingsCloudLoggingSettingsOutputReference
	CloudLoggingSettingsInput() *GoogleCesAppLoggingSettingsCloudLoggingSettings
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
	ConversationLoggingSettings() GoogleCesAppLoggingSettingsConversationLoggingSettingsOutputReference
	ConversationLoggingSettingsInput() *GoogleCesAppLoggingSettingsConversationLoggingSettings
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCesAppLoggingSettings
	SetInternalValue(val *GoogleCesAppLoggingSettings)
	RedactionConfig() GoogleCesAppLoggingSettingsRedactionConfigOutputReference
	RedactionConfigInput() *GoogleCesAppLoggingSettingsRedactionConfig
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
	PutAudioRecordingConfig(value *GoogleCesAppLoggingSettingsAudioRecordingConfig)
	PutBigqueryExportSettings(value *GoogleCesAppLoggingSettingsBigqueryExportSettings)
	PutCloudLoggingSettings(value *GoogleCesAppLoggingSettingsCloudLoggingSettings)
	PutConversationLoggingSettings(value *GoogleCesAppLoggingSettingsConversationLoggingSettings)
	PutRedactionConfig(value *GoogleCesAppLoggingSettingsRedactionConfig)
	ResetAudioRecordingConfig()
	ResetBigqueryExportSettings()
	ResetCloudLoggingSettings()
	ResetConversationLoggingSettings()
	ResetRedactionConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesAppLoggingSettingsOutputReference
type jsiiProxy_GoogleCesAppLoggingSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) AudioRecordingConfig() GoogleCesAppLoggingSettingsAudioRecordingConfigOutputReference {
	var returns GoogleCesAppLoggingSettingsAudioRecordingConfigOutputReference
	_jsii_.Get(
		j,
		"audioRecordingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) AudioRecordingConfigInput() *GoogleCesAppLoggingSettingsAudioRecordingConfig {
	var returns *GoogleCesAppLoggingSettingsAudioRecordingConfig
	_jsii_.Get(
		j,
		"audioRecordingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) BigqueryExportSettings() GoogleCesAppLoggingSettingsBigqueryExportSettingsOutputReference {
	var returns GoogleCesAppLoggingSettingsBigqueryExportSettingsOutputReference
	_jsii_.Get(
		j,
		"bigqueryExportSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) BigqueryExportSettingsInput() *GoogleCesAppLoggingSettingsBigqueryExportSettings {
	var returns *GoogleCesAppLoggingSettingsBigqueryExportSettings
	_jsii_.Get(
		j,
		"bigqueryExportSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) CloudLoggingSettings() GoogleCesAppLoggingSettingsCloudLoggingSettingsOutputReference {
	var returns GoogleCesAppLoggingSettingsCloudLoggingSettingsOutputReference
	_jsii_.Get(
		j,
		"cloudLoggingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) CloudLoggingSettingsInput() *GoogleCesAppLoggingSettingsCloudLoggingSettings {
	var returns *GoogleCesAppLoggingSettingsCloudLoggingSettings
	_jsii_.Get(
		j,
		"cloudLoggingSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) ConversationLoggingSettings() GoogleCesAppLoggingSettingsConversationLoggingSettingsOutputReference {
	var returns GoogleCesAppLoggingSettingsConversationLoggingSettingsOutputReference
	_jsii_.Get(
		j,
		"conversationLoggingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) ConversationLoggingSettingsInput() *GoogleCesAppLoggingSettingsConversationLoggingSettings {
	var returns *GoogleCesAppLoggingSettingsConversationLoggingSettings
	_jsii_.Get(
		j,
		"conversationLoggingSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) InternalValue() *GoogleCesAppLoggingSettings {
	var returns *GoogleCesAppLoggingSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) RedactionConfig() GoogleCesAppLoggingSettingsRedactionConfigOutputReference {
	var returns GoogleCesAppLoggingSettingsRedactionConfigOutputReference
	_jsii_.Get(
		j,
		"redactionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) RedactionConfigInput() *GoogleCesAppLoggingSettingsRedactionConfig {
	var returns *GoogleCesAppLoggingSettingsRedactionConfig
	_jsii_.Get(
		j,
		"redactionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCesAppLoggingSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesAppLoggingSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesAppLoggingSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesAppLoggingSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesAppLoggingSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesAppLoggingSettingsOutputReference_Override(g GoogleCesAppLoggingSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesAppLoggingSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference)SetInternalValue(val *GoogleCesAppLoggingSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) PutAudioRecordingConfig(value *GoogleCesAppLoggingSettingsAudioRecordingConfig) {
	if err := g.validatePutAudioRecordingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAudioRecordingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) PutBigqueryExportSettings(value *GoogleCesAppLoggingSettingsBigqueryExportSettings) {
	if err := g.validatePutBigqueryExportSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBigqueryExportSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) PutCloudLoggingSettings(value *GoogleCesAppLoggingSettingsCloudLoggingSettings) {
	if err := g.validatePutCloudLoggingSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCloudLoggingSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) PutConversationLoggingSettings(value *GoogleCesAppLoggingSettingsConversationLoggingSettings) {
	if err := g.validatePutConversationLoggingSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConversationLoggingSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) PutRedactionConfig(value *GoogleCesAppLoggingSettingsRedactionConfig) {
	if err := g.validatePutRedactionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRedactionConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) ResetAudioRecordingConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAudioRecordingConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) ResetBigqueryExportSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryExportSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) ResetCloudLoggingSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudLoggingSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) ResetConversationLoggingSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetConversationLoggingSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) ResetRedactionConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetRedactionConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package cesapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cesapp/internal"
)

type CesAppLoggingSettingsOutputReference interface {
	cdktf.ComplexObject
	AudioRecordingConfig() CesAppLoggingSettingsAudioRecordingConfigOutputReference
	AudioRecordingConfigInput() *CesAppLoggingSettingsAudioRecordingConfig
	BigqueryExportSettings() CesAppLoggingSettingsBigqueryExportSettingsOutputReference
	BigqueryExportSettingsInput() *CesAppLoggingSettingsBigqueryExportSettings
	CloudLoggingSettings() CesAppLoggingSettingsCloudLoggingSettingsOutputReference
	CloudLoggingSettingsInput() *CesAppLoggingSettingsCloudLoggingSettings
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
	ConversationLoggingSettings() CesAppLoggingSettingsConversationLoggingSettingsOutputReference
	ConversationLoggingSettingsInput() *CesAppLoggingSettingsConversationLoggingSettings
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *CesAppLoggingSettings
	SetInternalValue(val *CesAppLoggingSettings)
	RedactionConfig() CesAppLoggingSettingsRedactionConfigOutputReference
	RedactionConfigInput() *CesAppLoggingSettingsRedactionConfig
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
	PutAudioRecordingConfig(value *CesAppLoggingSettingsAudioRecordingConfig)
	PutBigqueryExportSettings(value *CesAppLoggingSettingsBigqueryExportSettings)
	PutCloudLoggingSettings(value *CesAppLoggingSettingsCloudLoggingSettings)
	PutConversationLoggingSettings(value *CesAppLoggingSettingsConversationLoggingSettings)
	PutRedactionConfig(value *CesAppLoggingSettingsRedactionConfig)
	ResetAudioRecordingConfig()
	ResetBigqueryExportSettings()
	ResetCloudLoggingSettings()
	ResetConversationLoggingSettings()
	ResetRedactionConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesAppLoggingSettingsOutputReference
type jsiiProxy_CesAppLoggingSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference) AudioRecordingConfig() CesAppLoggingSettingsAudioRecordingConfigOutputReference {
	var returns CesAppLoggingSettingsAudioRecordingConfigOutputReference
	_jsii_.Get(
		j,
		"audioRecordingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference) AudioRecordingConfigInput() *CesAppLoggingSettingsAudioRecordingConfig {
	var returns *CesAppLoggingSettingsAudioRecordingConfig
	_jsii_.Get(
		j,
		"audioRecordingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference) BigqueryExportSettings() CesAppLoggingSettingsBigqueryExportSettingsOutputReference {
	var returns CesAppLoggingSettingsBigqueryExportSettingsOutputReference
	_jsii_.Get(
		j,
		"bigqueryExportSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference) BigqueryExportSettingsInput() *CesAppLoggingSettingsBigqueryExportSettings {
	var returns *CesAppLoggingSettingsBigqueryExportSettings
	_jsii_.Get(
		j,
		"bigqueryExportSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference) CloudLoggingSettings() CesAppLoggingSettingsCloudLoggingSettingsOutputReference {
	var returns CesAppLoggingSettingsCloudLoggingSettingsOutputReference
	_jsii_.Get(
		j,
		"cloudLoggingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference) CloudLoggingSettingsInput() *CesAppLoggingSettingsCloudLoggingSettings {
	var returns *CesAppLoggingSettingsCloudLoggingSettings
	_jsii_.Get(
		j,
		"cloudLoggingSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference) ConversationLoggingSettings() CesAppLoggingSettingsConversationLoggingSettingsOutputReference {
	var returns CesAppLoggingSettingsConversationLoggingSettingsOutputReference
	_jsii_.Get(
		j,
		"conversationLoggingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference) ConversationLoggingSettingsInput() *CesAppLoggingSettingsConversationLoggingSettings {
	var returns *CesAppLoggingSettingsConversationLoggingSettings
	_jsii_.Get(
		j,
		"conversationLoggingSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference) InternalValue() *CesAppLoggingSettings {
	var returns *CesAppLoggingSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference) RedactionConfig() CesAppLoggingSettingsRedactionConfigOutputReference {
	var returns CesAppLoggingSettingsRedactionConfigOutputReference
	_jsii_.Get(
		j,
		"redactionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference) RedactionConfigInput() *CesAppLoggingSettingsRedactionConfig {
	var returns *CesAppLoggingSettingsRedactionConfig
	_jsii_.Get(
		j,
		"redactionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCesAppLoggingSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CesAppLoggingSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewCesAppLoggingSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesAppLoggingSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cesApp.CesAppLoggingSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesAppLoggingSettingsOutputReference_Override(c CesAppLoggingSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesApp.CesAppLoggingSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference)SetInternalValue(val *CesAppLoggingSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesAppLoggingSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) PutAudioRecordingConfig(value *CesAppLoggingSettingsAudioRecordingConfig) {
	if err := c.validatePutAudioRecordingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAudioRecordingConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) PutBigqueryExportSettings(value *CesAppLoggingSettingsBigqueryExportSettings) {
	if err := c.validatePutBigqueryExportSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBigqueryExportSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) PutCloudLoggingSettings(value *CesAppLoggingSettingsCloudLoggingSettings) {
	if err := c.validatePutCloudLoggingSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCloudLoggingSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) PutConversationLoggingSettings(value *CesAppLoggingSettingsConversationLoggingSettings) {
	if err := c.validatePutConversationLoggingSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putConversationLoggingSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) PutRedactionConfig(value *CesAppLoggingSettingsRedactionConfig) {
	if err := c.validatePutRedactionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRedactionConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) ResetAudioRecordingConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetAudioRecordingConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) ResetBigqueryExportSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetBigqueryExportSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) ResetCloudLoggingSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudLoggingSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) ResetConversationLoggingSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetConversationLoggingSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) ResetRedactionConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetRedactionConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppLoggingSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package googledialogflowcxpage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowcxpage/internal"
)

type GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference interface {
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
	EndpointerSensitivity() *float64
	SetEndpointerSensitivity(val *float64)
	EndpointerSensitivityInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettings
	SetInternalValue(val *GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettings)
	Models() *map[string]*string
	SetModels(val *map[string]*string)
	ModelsInput() *map[string]*string
	NoSpeechTimeout() *string
	SetNoSpeechTimeout(val *string)
	NoSpeechTimeoutInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UseTimeoutBasedEndpointing() interface{}
	SetUseTimeoutBasedEndpointing(val interface{})
	UseTimeoutBasedEndpointingInput() interface{}
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
	ResetEndpointerSensitivity()
	ResetModels()
	ResetNoSpeechTimeout()
	ResetUseTimeoutBasedEndpointing()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference
type jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) EndpointerSensitivity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endpointerSensitivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) EndpointerSensitivityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endpointerSensitivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) InternalValue() *GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettings {
	var returns *GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) Models() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"models",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) ModelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"modelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) NoSpeechTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noSpeechTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) NoSpeechTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noSpeechTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) UseTimeoutBasedEndpointing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useTimeoutBasedEndpointing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) UseTimeoutBasedEndpointingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useTimeoutBasedEndpointingInput",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference_Override(g GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference)SetEndpointerSensitivity(val *float64) {
	if err := j.validateSetEndpointerSensitivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointerSensitivity",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference)SetInternalValue(val *GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference)SetModels(val *map[string]*string) {
	if err := j.validateSetModelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"models",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference)SetNoSpeechTimeout(val *string) {
	if err := j.validateSetNoSpeechTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSpeechTimeout",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference)SetUseTimeoutBasedEndpointing(val interface{}) {
	if err := j.validateSetUseTimeoutBasedEndpointingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useTimeoutBasedEndpointing",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) ResetEndpointerSensitivity() {
	_jsii_.InvokeVoid(
		g,
		"resetEndpointerSensitivity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) ResetModels() {
	_jsii_.InvokeVoid(
		g,
		"resetModels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) ResetNoSpeechTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetNoSpeechTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) ResetUseTimeoutBasedEndpointing() {
	_jsii_.InvokeVoid(
		g,
		"resetUseTimeoutBasedEndpointing",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package dialogflowcxflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/dialogflowcxflow/internal"
)

type DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference interface {
	cdktf.ComplexObject
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
	EndpointerSensitivity() *float64
	SetEndpointerSensitivity(val *float64)
	EndpointerSensitivityInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *DialogflowCxFlowAdvancedSettingsSpeechSettings
	SetInternalValue(val *DialogflowCxFlowAdvancedSettingsSpeechSettings)
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
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseTimeoutBasedEndpointing() any
	SetUseTimeoutBasedEndpointing(val any)
	UseTimeoutBasedEndpointingInput() any
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
	ResetEndpointerSensitivity()
	ResetModels()
	ResetNoSpeechTimeout()
	ResetUseTimeoutBasedEndpointing()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference
type jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) EndpointerSensitivity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endpointerSensitivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) EndpointerSensitivityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endpointerSensitivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) InternalValue() *DialogflowCxFlowAdvancedSettingsSpeechSettings {
	var returns *DialogflowCxFlowAdvancedSettingsSpeechSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) Models() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"models",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) ModelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"modelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) NoSpeechTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noSpeechTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) NoSpeechTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noSpeechTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) UseTimeoutBasedEndpointing() any {
	var returns any
	_jsii_.Get(
		j,
		"useTimeoutBasedEndpointing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) UseTimeoutBasedEndpointingInput() any {
	var returns any
	_jsii_.Get(
		j,
		"useTimeoutBasedEndpointingInput",
		&returns,
	)
	return returns
}

func NewDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxFlow.DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference_Override(d DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxFlow.DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) SetEndpointerSensitivity(val *float64) {
	if err := j.validateSetEndpointerSensitivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointerSensitivity",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) SetInternalValue(val *DialogflowCxFlowAdvancedSettingsSpeechSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) SetModels(val *map[string]*string) {
	if err := j.validateSetModelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"models",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) SetNoSpeechTimeout(val *string) {
	if err := j.validateSetNoSpeechTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSpeechTimeout",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) SetUseTimeoutBasedEndpointing(val any) {
	if err := j.validateSetUseTimeoutBasedEndpointingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useTimeoutBasedEndpointing",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) ResetEndpointerSensitivity() {
	_jsii_.InvokeVoid(
		d,
		"resetEndpointerSensitivity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) ResetModels() {
	_jsii_.InvokeVoid(
		d,
		"resetModels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) ResetNoSpeechTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetNoSpeechTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) ResetUseTimeoutBasedEndpointing() {
	_jsii_.InvokeVoid(
		d,
		"resetUseTimeoutBasedEndpointing",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		d,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

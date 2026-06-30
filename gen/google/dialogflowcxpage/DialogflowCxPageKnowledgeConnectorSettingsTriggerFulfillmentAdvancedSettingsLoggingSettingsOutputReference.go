package dialogflowcxpage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/dialogflowcxpage/internal"
)

type DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference interface {
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
	EnableConsentBasedRedaction() any
	SetEnableConsentBasedRedaction(val any)
	EnableConsentBasedRedactionInput() any
	EnableInteractionLogging() any
	SetEnableInteractionLogging(val any)
	EnableInteractionLoggingInput() any
	EnableStackdriverLogging() any
	SetEnableStackdriverLogging(val any)
	EnableStackdriverLoggingInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettings
	SetInternalValue(val *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettings)
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
	ResetEnableConsentBasedRedaction()
	ResetEnableInteractionLogging()
	ResetEnableStackdriverLogging()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference
type jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) EnableConsentBasedRedaction() any {
	var returns any
	_jsii_.Get(
		j,
		"enableConsentBasedRedaction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) EnableConsentBasedRedactionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableConsentBasedRedactionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) EnableInteractionLogging() any {
	var returns any
	_jsii_.Get(
		j,
		"enableInteractionLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) EnableInteractionLoggingInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableInteractionLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) EnableStackdriverLogging() any {
	var returns any
	_jsii_.Get(
		j,
		"enableStackdriverLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) EnableStackdriverLoggingInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableStackdriverLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) InternalValue() *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettings {
	var returns *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxPage.DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference_Override(d DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxPage.DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) SetEnableConsentBasedRedaction(val any) {
	if err := j.validateSetEnableConsentBasedRedactionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableConsentBasedRedaction",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) SetEnableInteractionLogging(val any) {
	if err := j.validateSetEnableInteractionLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableInteractionLogging",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) SetEnableStackdriverLogging(val any) {
	if err := j.validateSetEnableStackdriverLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableStackdriverLogging",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) SetInternalValue(val *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) ResetEnableConsentBasedRedaction() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableConsentBasedRedaction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) ResetEnableInteractionLogging() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableInteractionLogging",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) ResetEnableStackdriverLogging() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableStackdriverLogging",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

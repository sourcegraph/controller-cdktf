package dialogflowcxagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/dialogflowcxagent/internal"
)

type DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference interface {
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
	InternalValue() *DialogflowCxAgentAdvancedSettingsLoggingSettings
	SetInternalValue(val *DialogflowCxAgentAdvancedSettingsLoggingSettings)
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

// The jsii proxy struct for DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference
type jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) EnableConsentBasedRedaction() any {
	var returns any
	_jsii_.Get(
		j,
		"enableConsentBasedRedaction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) EnableConsentBasedRedactionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableConsentBasedRedactionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) EnableInteractionLogging() any {
	var returns any
	_jsii_.Get(
		j,
		"enableInteractionLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) EnableInteractionLoggingInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableInteractionLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) EnableStackdriverLogging() any {
	var returns any
	_jsii_.Get(
		j,
		"enableStackdriverLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) EnableStackdriverLoggingInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableStackdriverLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) InternalValue() *DialogflowCxAgentAdvancedSettingsLoggingSettings {
	var returns *DialogflowCxAgentAdvancedSettingsLoggingSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxAgent.DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference_Override(d DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxAgent.DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) SetEnableConsentBasedRedaction(val any) {
	if err := j.validateSetEnableConsentBasedRedactionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableConsentBasedRedaction",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) SetEnableInteractionLogging(val any) {
	if err := j.validateSetEnableInteractionLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableInteractionLogging",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) SetEnableStackdriverLogging(val any) {
	if err := j.validateSetEnableStackdriverLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableStackdriverLogging",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) SetInternalValue(val *DialogflowCxAgentAdvancedSettingsLoggingSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) ResetEnableConsentBasedRedaction() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableConsentBasedRedaction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) ResetEnableInteractionLogging() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableInteractionLogging",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) ResetEnableStackdriverLogging() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableStackdriverLogging",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

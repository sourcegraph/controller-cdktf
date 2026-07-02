package dialogflowcxpage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/dialogflowcxpage/internal"
)

type DialogflowCxPageFormParametersOutputReference interface {
	cdktf.ComplexObject
	AdvancedSettings() DialogflowCxPageFormParametersAdvancedSettingsOutputReference
	AdvancedSettingsInput() *DialogflowCxPageFormParametersAdvancedSettings
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
	DefaultValue() *string
	SetDefaultValue(val *string)
	DefaultValueInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EntityType() *string
	SetEntityType(val *string)
	EntityTypeInput() *string
	FillBehavior() DialogflowCxPageFormParametersFillBehaviorOutputReference
	FillBehaviorInput() *DialogflowCxPageFormParametersFillBehavior
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	IsList() any
	SetIsList(val any)
	IsListInput() any
	Redact() any
	SetRedact(val any)
	RedactInput() any
	Required() any
	SetRequired(val any)
	RequiredInput() any
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
	PutAdvancedSettings(value *DialogflowCxPageFormParametersAdvancedSettings)
	PutFillBehavior(value *DialogflowCxPageFormParametersFillBehavior)
	ResetAdvancedSettings()
	ResetDefaultValue()
	ResetDisplayName()
	ResetEntityType()
	ResetFillBehavior()
	ResetIsList()
	ResetRedact()
	ResetRequired()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxPageFormParametersOutputReference
type jsiiProxy_DialogflowCxPageFormParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) AdvancedSettings() DialogflowCxPageFormParametersAdvancedSettingsOutputReference {
	var returns DialogflowCxPageFormParametersAdvancedSettingsOutputReference
	_jsii_.Get(
		j,
		"advancedSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) AdvancedSettingsInput() *DialogflowCxPageFormParametersAdvancedSettings {
	var returns *DialogflowCxPageFormParametersAdvancedSettings
	_jsii_.Get(
		j,
		"advancedSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) DefaultValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) DefaultValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) EntityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) EntityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) FillBehavior() DialogflowCxPageFormParametersFillBehaviorOutputReference {
	var returns DialogflowCxPageFormParametersFillBehaviorOutputReference
	_jsii_.Get(
		j,
		"fillBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) FillBehaviorInput() *DialogflowCxPageFormParametersFillBehavior {
	var returns *DialogflowCxPageFormParametersFillBehavior
	_jsii_.Get(
		j,
		"fillBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) IsList() any {
	var returns any
	_jsii_.Get(
		j,
		"isList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) IsListInput() any {
	var returns any
	_jsii_.Get(
		j,
		"isListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) Redact() any {
	var returns any
	_jsii_.Get(
		j,
		"redact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) RedactInput() any {
	var returns any
	_jsii_.Get(
		j,
		"redactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) Required() any {
	var returns any
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) RequiredInput() any {
	var returns any
	_jsii_.Get(
		j,
		"requiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDialogflowCxPageFormParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DialogflowCxPageFormParametersOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxPageFormParametersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxPageFormParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxPage.DialogflowCxPageFormParametersOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDialogflowCxPageFormParametersOutputReference_Override(d DialogflowCxPageFormParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxPage.DialogflowCxPageFormParametersOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) SetDefaultValue(val *string) {
	if err := j.validateSetDefaultValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) SetEntityType(val *string) {
	if err := j.validateSetEntityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityType",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) SetIsList(val any) {
	if err := j.validateSetIsListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isList",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) SetRedact(val any) {
	if err := j.validateSetRedactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redact",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) SetRequired(val any) {
	if err := j.validateSetRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageFormParametersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) PutAdvancedSettings(value *DialogflowCxPageFormParametersAdvancedSettings) {
	if err := d.validatePutAdvancedSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAdvancedSettings",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) PutFillBehavior(value *DialogflowCxPageFormParametersFillBehavior) {
	if err := d.validatePutFillBehaviorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFillBehavior",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) ResetAdvancedSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetAdvancedSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) ResetDefaultValue() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) ResetEntityType() {
	_jsii_.InvokeVoid(
		d,
		"resetEntityType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) ResetFillBehavior() {
	_jsii_.InvokeVoid(
		d,
		"resetFillBehavior",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) ResetIsList() {
	_jsii_.InvokeVoid(
		d,
		"resetIsList",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) ResetRedact() {
	_jsii_.InvokeVoid(
		d,
		"resetRedact",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) ResetRequired() {
	_jsii_.InvokeVoid(
		d,
		"resetRequired",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

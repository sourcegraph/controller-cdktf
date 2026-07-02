package apihubplugin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/apihubplugin/internal"
)

type ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EnumOptions() ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList
	EnumOptionsInput() any
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() any
	SetInternalValue(val any)
	MultiSelectOptions() ApihubPluginConfigTemplateAdditionalConfigTemplateMultiSelectOptionsList
	MultiSelectOptionsInput() any
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
	ValidationRegex() *string
	SetValidationRegex(val *string)
	ValidationRegexInput() *string
	ValueType() *string
	SetValueType(val *string)
	ValueTypeInput() *string
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
	PutEnumOptions(value any)
	PutMultiSelectOptions(value any)
	ResetDescription()
	ResetEnumOptions()
	ResetMultiSelectOptions()
	ResetRequired()
	ResetValidationRegex()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference
type jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) EnumOptions() ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList {
	var returns ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList
	_jsii_.Get(
		j,
		"enumOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) EnumOptionsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enumOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) MultiSelectOptions() ApihubPluginConfigTemplateAdditionalConfigTemplateMultiSelectOptionsList {
	var returns ApihubPluginConfigTemplateAdditionalConfigTemplateMultiSelectOptionsList
	_jsii_.Get(
		j,
		"multiSelectOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) MultiSelectOptionsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"multiSelectOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) Required() any {
	var returns any
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) RequiredInput() any {
	var returns any
	_jsii_.Get(
		j,
		"requiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ValidationRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ValidationRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ValueType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ValueTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueTypeInput",
		&returns,
	)
	return returns
}

func NewApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference {
	_init_.Initialize()

	if err := validateNewApihubPluginConfigTemplateAdditionalConfigTemplateOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.apihubPlugin.ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference_Override(a ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.apihubPlugin.ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) SetRequired(val any) {
	if err := j.validateSetRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) SetValidationRegex(val *string) {
	if err := j.validateSetValidationRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validationRegex",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) SetValueType(val *string) {
	if err := j.validateSetValueTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueType",
		val,
	)
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) PutEnumOptions(value any) {
	if err := a.validatePutEnumOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnumOptions",
		[]any{value},
	)
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) PutMultiSelectOptions(value any) {
	if err := a.validatePutMultiSelectOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMultiSelectOptions",
		[]any{value},
	)
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ResetEnumOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetEnumOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ResetMultiSelectOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetMultiSelectOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ResetRequired() {
	_jsii_.InvokeVoid(
		a,
		"resetRequired",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ResetValidationRegex() {
	_jsii_.InvokeVoid(
		a,
		"resetValidationRegex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		a,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

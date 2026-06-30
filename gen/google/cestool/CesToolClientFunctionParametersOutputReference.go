package cestool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cestool/internal"
)

type CesToolClientFunctionParametersOutputReference interface {
	cdktf.ComplexObject
	AdditionalProperties() *string
	SetAdditionalProperties(val *string)
	AdditionalPropertiesInput() *string
	AnyOf() *string
	SetAnyOf(val *string)
	AnyOfInput() *string
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
	Default() *string
	SetDefault(val *string)
	DefaultInput() *string
	Defs() *string
	SetDefs(val *string)
	DefsInput() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Enum() *[]*string
	SetEnum(val *[]*string)
	EnumInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *CesToolClientFunctionParameters
	SetInternalValue(val *CesToolClientFunctionParameters)
	Items() *string
	SetItems(val *string)
	ItemsInput() *string
	Maximum() *float64
	SetMaximum(val *float64)
	MaximumInput() *float64
	MaxItems() *float64
	SetMaxItems(val *float64)
	MaxItemsInput() *float64
	Minimum() *float64
	SetMinimum(val *float64)
	MinimumInput() *float64
	MinItems() *float64
	SetMinItems(val *float64)
	MinItemsInput() *float64
	Nullable() any
	SetNullable(val any)
	NullableInput() any
	PrefixItems() *string
	SetPrefixItems(val *string)
	PrefixItemsInput() *string
	Properties() *string
	SetProperties(val *string)
	PropertiesInput() *string
	Ref() *string
	SetRef(val *string)
	RefInput() *string
	Required() *[]*string
	SetRequired(val *[]*string)
	RequiredInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UniqueItems() any
	SetUniqueItems(val any)
	UniqueItemsInput() any
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
	ResetAdditionalProperties()
	ResetAnyOf()
	ResetDefault()
	ResetDefs()
	ResetDescription()
	ResetEnum()
	ResetItems()
	ResetMaximum()
	ResetMaxItems()
	ResetMinimum()
	ResetMinItems()
	ResetNullable()
	ResetPrefixItems()
	ResetProperties()
	ResetRef()
	ResetRequired()
	ResetTitle()
	ResetUniqueItems()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesToolClientFunctionParametersOutputReference
type jsiiProxy_CesToolClientFunctionParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) AdditionalProperties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) AdditionalPropertiesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) AnyOf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anyOf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) AnyOfInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anyOfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) Default() *string {
	var returns *string
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) DefaultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) Defs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) DefsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) Enum() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) EnumInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) InternalValue() *CesToolClientFunctionParameters {
	var returns *CesToolClientFunctionParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) Items() *string {
	var returns *string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) ItemsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"itemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) Maximum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) MaximumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) MaxItems() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) MaxItemsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxItemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) Minimum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) MinimumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) MinItems() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) MinItemsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minItemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) Nullable() any {
	var returns any
	_jsii_.Get(
		j,
		"nullable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) NullableInput() any {
	var returns any
	_jsii_.Get(
		j,
		"nullableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) PrefixItems() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) PrefixItemsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixItemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) Properties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) PropertiesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) RefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) Required() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) RequiredInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) UniqueItems() any {
	var returns any
	_jsii_.Get(
		j,
		"uniqueItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) UniqueItemsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"uniqueItemsInput",
		&returns,
	)
	return returns
}

func NewCesToolClientFunctionParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CesToolClientFunctionParametersOutputReference {
	_init_.Initialize()

	if err := validateNewCesToolClientFunctionParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesToolClientFunctionParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cesTool.CesToolClientFunctionParametersOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesToolClientFunctionParametersOutputReference_Override(c CesToolClientFunctionParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesTool.CesToolClientFunctionParametersOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetAdditionalProperties(val *string) {
	if err := j.validateSetAdditionalPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalProperties",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetAnyOf(val *string) {
	if err := j.validateSetAnyOfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anyOf",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetDefault(val *string) {
	if err := j.validateSetDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"default",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetDefs(val *string) {
	if err := j.validateSetDefsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defs",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetEnum(val *[]*string) {
	if err := j.validateSetEnumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enum",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetInternalValue(val *CesToolClientFunctionParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetItems(val *string) {
	if err := j.validateSetItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"items",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetMaximum(val *float64) {
	if err := j.validateSetMaximumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximum",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetMaxItems(val *float64) {
	if err := j.validateSetMaxItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxItems",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetMinimum(val *float64) {
	if err := j.validateSetMinimumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimum",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetMinItems(val *float64) {
	if err := j.validateSetMinItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minItems",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetNullable(val any) {
	if err := j.validateSetNullableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullable",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetPrefixItems(val *string) {
	if err := j.validateSetPrefixItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixItems",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetProperties(val *string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetRef(val *string) {
	if err := j.validateSetRefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ref",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetRequired(val *[]*string) {
	if err := j.validateSetRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CesToolClientFunctionParametersOutputReference) SetUniqueItems(val any) {
	if err := j.validateSetUniqueItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uniqueItems",
		val,
	)
}

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) ResetAdditionalProperties() {
	_jsii_.InvokeVoid(
		c,
		"resetAdditionalProperties",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) ResetAnyOf() {
	_jsii_.InvokeVoid(
		c,
		"resetAnyOf",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) ResetDefault() {
	_jsii_.InvokeVoid(
		c,
		"resetDefault",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) ResetDefs() {
	_jsii_.InvokeVoid(
		c,
		"resetDefs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) ResetEnum() {
	_jsii_.InvokeVoid(
		c,
		"resetEnum",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) ResetItems() {
	_jsii_.InvokeVoid(
		c,
		"resetItems",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) ResetMaximum() {
	_jsii_.InvokeVoid(
		c,
		"resetMaximum",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) ResetMaxItems() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxItems",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) ResetMinimum() {
	_jsii_.InvokeVoid(
		c,
		"resetMinimum",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) ResetMinItems() {
	_jsii_.InvokeVoid(
		c,
		"resetMinItems",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) ResetNullable() {
	_jsii_.InvokeVoid(
		c,
		"resetNullable",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) ResetPrefixItems() {
	_jsii_.InvokeVoid(
		c,
		"resetPrefixItems",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		c,
		"resetProperties",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) ResetRef() {
	_jsii_.InvokeVoid(
		c,
		"resetRef",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) ResetRequired() {
	_jsii_.InvokeVoid(
		c,
		"resetRequired",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		c,
		"resetTitle",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) ResetUniqueItems() {
	_jsii_.InvokeVoid(
		c,
		"resetUniqueItems",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_CesToolClientFunctionParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

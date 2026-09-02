package googleapihubplugin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleapihubplugin/internal"
)

type GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EnumOptions() GoogleApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList
	EnumOptionsInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MultiSelectOptions() GoogleApihubPluginConfigTemplateAdditionalConfigTemplateMultiSelectOptionsList
	MultiSelectOptionsInput() interface{}
	Required() interface{}
	SetRequired(val interface{})
	RequiredInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ValidationRegex() *string
	SetValidationRegex(val *string)
	ValidationRegexInput() *string
	ValueType() *string
	SetValueType(val *string)
	ValueTypeInput() *string
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
	PutEnumOptions(value interface{})
	PutMultiSelectOptions(value interface{})
	ResetDescription()
	ResetEnumOptions()
	ResetMultiSelectOptions()
	ResetRequired()
	ResetValidationRegex()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference
type jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) EnumOptions() GoogleApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList {
	var returns GoogleApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList
	_jsii_.Get(
		j,
		"enumOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) EnumOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enumOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) MultiSelectOptions() GoogleApihubPluginConfigTemplateAdditionalConfigTemplateMultiSelectOptionsList {
	var returns GoogleApihubPluginConfigTemplateAdditionalConfigTemplateMultiSelectOptionsList
	_jsii_.Get(
		j,
		"multiSelectOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) MultiSelectOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiSelectOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) Required() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) RequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ValidationRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ValidationRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ValueType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ValueTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueTypeInput",
		&returns,
	)
	return returns
}


func NewGoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleApihubPlugin.GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference_Override(g GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleApihubPlugin.GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference)SetRequired(val interface{}) {
	if err := j.validateSetRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference)SetValidationRegex(val *string) {
	if err := j.validateSetValidationRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validationRegex",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference)SetValueType(val *string) {
	if err := j.validateSetValueTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueType",
		val,
	)
}

func (g *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) PutEnumOptions(value interface{}) {
	if err := g.validatePutEnumOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEnumOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) PutMultiSelectOptions(value interface{}) {
	if err := g.validatePutMultiSelectOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMultiSelectOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ResetEnumOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetEnumOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ResetMultiSelectOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetMultiSelectOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ResetRequired() {
	_jsii_.InvokeVoid(
		g,
		"resetRequired",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ResetValidationRegex() {
	_jsii_.InvokeVoid(
		g,
		"resetValidationRegex",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleApihubPluginConfigTemplateAdditionalConfigTemplateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


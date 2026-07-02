package datalosspreventioninspecttemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/datalosspreventioninspecttemplate/internal"
)

type DataLossPreventionInspectTemplateInspectConfigOutputReference interface {
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
	ContentOptions() *[]*string
	SetContentOptions(val *[]*string)
	ContentOptionsInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomInfoTypes() DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesList
	CustomInfoTypesInput() any
	ExcludeInfoTypes() any
	SetExcludeInfoTypes(val any)
	ExcludeInfoTypesInput() any
	// Experimental.
	Fqn() *string
	IncludeQuote() any
	SetIncludeQuote(val any)
	IncludeQuoteInput() any
	InfoTypes() DataLossPreventionInspectTemplateInspectConfigInfoTypesList
	InfoTypesInput() any
	InternalValue() *DataLossPreventionInspectTemplateInspectConfig
	SetInternalValue(val *DataLossPreventionInspectTemplateInspectConfig)
	Limits() DataLossPreventionInspectTemplateInspectConfigLimitsOutputReference
	LimitsInput() *DataLossPreventionInspectTemplateInspectConfigLimits
	MinLikelihood() *string
	SetMinLikelihood(val *string)
	MinLikelihoodInput() *string
	RuleSet() DataLossPreventionInspectTemplateInspectConfigRuleSetList
	RuleSetInput() any
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
	PutCustomInfoTypes(value any)
	PutInfoTypes(value any)
	PutLimits(value *DataLossPreventionInspectTemplateInspectConfigLimits)
	PutRuleSet(value any)
	ResetContentOptions()
	ResetCustomInfoTypes()
	ResetExcludeInfoTypes()
	ResetIncludeQuote()
	ResetInfoTypes()
	ResetLimits()
	ResetMinLikelihood()
	ResetRuleSet()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionInspectTemplateInspectConfigOutputReference
type jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) ContentOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"contentOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) ContentOptionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"contentOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) CustomInfoTypes() DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesList {
	var returns DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesList
	_jsii_.Get(
		j,
		"customInfoTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) CustomInfoTypesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"customInfoTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) ExcludeInfoTypes() any {
	var returns any
	_jsii_.Get(
		j,
		"excludeInfoTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) ExcludeInfoTypesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"excludeInfoTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) IncludeQuote() any {
	var returns any
	_jsii_.Get(
		j,
		"includeQuote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) IncludeQuoteInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeQuoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) InfoTypes() DataLossPreventionInspectTemplateInspectConfigInfoTypesList {
	var returns DataLossPreventionInspectTemplateInspectConfigInfoTypesList
	_jsii_.Get(
		j,
		"infoTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) InfoTypesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"infoTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) InternalValue() *DataLossPreventionInspectTemplateInspectConfig {
	var returns *DataLossPreventionInspectTemplateInspectConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) Limits() DataLossPreventionInspectTemplateInspectConfigLimitsOutputReference {
	var returns DataLossPreventionInspectTemplateInspectConfigLimitsOutputReference
	_jsii_.Get(
		j,
		"limits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) LimitsInput() *DataLossPreventionInspectTemplateInspectConfigLimits {
	var returns *DataLossPreventionInspectTemplateInspectConfigLimits
	_jsii_.Get(
		j,
		"limitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) MinLikelihood() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minLikelihood",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) MinLikelihoodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minLikelihoodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) RuleSet() DataLossPreventionInspectTemplateInspectConfigRuleSetList {
	var returns DataLossPreventionInspectTemplateInspectConfigRuleSetList
	_jsii_.Get(
		j,
		"ruleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) RuleSetInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ruleSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDataLossPreventionInspectTemplateInspectConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataLossPreventionInspectTemplateInspectConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionInspectTemplateInspectConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionInspectTemplate.DataLossPreventionInspectTemplateInspectConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataLossPreventionInspectTemplateInspectConfigOutputReference_Override(d DataLossPreventionInspectTemplateInspectConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionInspectTemplate.DataLossPreventionInspectTemplateInspectConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) SetContentOptions(val *[]*string) {
	if err := j.validateSetContentOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentOptions",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) SetExcludeInfoTypes(val any) {
	if err := j.validateSetExcludeInfoTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeInfoTypes",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) SetIncludeQuote(val any) {
	if err := j.validateSetIncludeQuoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeQuote",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) SetInternalValue(val *DataLossPreventionInspectTemplateInspectConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) SetMinLikelihood(val *string) {
	if err := j.validateSetMinLikelihoodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minLikelihood",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) PutCustomInfoTypes(value any) {
	if err := d.validatePutCustomInfoTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustomInfoTypes",
		[]any{value},
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) PutInfoTypes(value any) {
	if err := d.validatePutInfoTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInfoTypes",
		[]any{value},
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) PutLimits(value *DataLossPreventionInspectTemplateInspectConfigLimits) {
	if err := d.validatePutLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLimits",
		[]any{value},
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) PutRuleSet(value any) {
	if err := d.validatePutRuleSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRuleSet",
		[]any{value},
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) ResetContentOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetContentOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) ResetCustomInfoTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomInfoTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) ResetExcludeInfoTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeInfoTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) ResetIncludeQuote() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeQuote",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) ResetInfoTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetInfoTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) ResetLimits() {
	_jsii_.InvokeVoid(
		d,
		"resetLimits",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) ResetMinLikelihood() {
	_jsii_.InvokeVoid(
		d,
		"resetMinLikelihood",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) ResetRuleSet() {
	_jsii_.InvokeVoid(
		d,
		"resetRuleSet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

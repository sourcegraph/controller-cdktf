package dialogflowgenerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/dialogflowgenerator/internal"
)

type DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	Section() *string
	SetSection(val *string)
	SectionInput() *string
	Summary() *string
	SetSummary(val *string)
	SummaryInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference
type jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) Section() *string {
	var returns *string
	_jsii_.Get(
		j,
		"section",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) SectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) Summary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) SummaryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowGenerator.DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference_Override(d DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowGenerator.DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) SetSection(val *string) {
	if err := j.validateSetSectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"section",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) SetSummary(val *string) {
	if err := j.validateSetSummaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"summary",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestionSummarySectionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

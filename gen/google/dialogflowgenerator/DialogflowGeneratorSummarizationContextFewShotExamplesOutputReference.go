package dialogflowgenerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/dialogflowgenerator/internal"
)

type DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference interface {
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
	ConversationContext() DialogflowGeneratorSummarizationContextFewShotExamplesConversationContextOutputReference
	ConversationContextInput() *DialogflowGeneratorSummarizationContextFewShotExamplesConversationContext
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExtraInfo() *map[string]*string
	SetExtraInfo(val *map[string]*string)
	ExtraInfoInput() *map[string]*string
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	Output() DialogflowGeneratorSummarizationContextFewShotExamplesOutputOutputReference
	OutputInput() *DialogflowGeneratorSummarizationContextFewShotExamplesOutput
	SummarizationSectionList() DialogflowGeneratorSummarizationContextFewShotExamplesSummarizationSectionListStructOutputReference
	SummarizationSectionListInput() *DialogflowGeneratorSummarizationContextFewShotExamplesSummarizationSectionListStruct
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
	PutConversationContext(value *DialogflowGeneratorSummarizationContextFewShotExamplesConversationContext)
	PutOutput(value *DialogflowGeneratorSummarizationContextFewShotExamplesOutput)
	PutSummarizationSectionList(value *DialogflowGeneratorSummarizationContextFewShotExamplesSummarizationSectionListStruct)
	ResetConversationContext()
	ResetExtraInfo()
	ResetSummarizationSectionList()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference
type jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ConversationContext() DialogflowGeneratorSummarizationContextFewShotExamplesConversationContextOutputReference {
	var returns DialogflowGeneratorSummarizationContextFewShotExamplesConversationContextOutputReference
	_jsii_.Get(
		j,
		"conversationContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ConversationContextInput() *DialogflowGeneratorSummarizationContextFewShotExamplesConversationContext {
	var returns *DialogflowGeneratorSummarizationContextFewShotExamplesConversationContext
	_jsii_.Get(
		j,
		"conversationContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ExtraInfo() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ExtraInfoInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) Output() DialogflowGeneratorSummarizationContextFewShotExamplesOutputOutputReference {
	var returns DialogflowGeneratorSummarizationContextFewShotExamplesOutputOutputReference
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) OutputInput() *DialogflowGeneratorSummarizationContextFewShotExamplesOutput {
	var returns *DialogflowGeneratorSummarizationContextFewShotExamplesOutput
	_jsii_.Get(
		j,
		"outputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) SummarizationSectionList() DialogflowGeneratorSummarizationContextFewShotExamplesSummarizationSectionListStructOutputReference {
	var returns DialogflowGeneratorSummarizationContextFewShotExamplesSummarizationSectionListStructOutputReference
	_jsii_.Get(
		j,
		"summarizationSectionList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) SummarizationSectionListInput() *DialogflowGeneratorSummarizationContextFewShotExamplesSummarizationSectionListStruct {
	var returns *DialogflowGeneratorSummarizationContextFewShotExamplesSummarizationSectionListStruct
	_jsii_.Get(
		j,
		"summarizationSectionListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowGeneratorSummarizationContextFewShotExamplesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowGenerator.DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference_Override(d DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowGenerator.DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) SetExtraInfo(val *map[string]*string) {
	if err := j.validateSetExtraInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraInfo",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) PutConversationContext(value *DialogflowGeneratorSummarizationContextFewShotExamplesConversationContext) {
	if err := d.validatePutConversationContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConversationContext",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) PutOutput(value *DialogflowGeneratorSummarizationContextFewShotExamplesOutput) {
	if err := d.validatePutOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOutput",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) PutSummarizationSectionList(value *DialogflowGeneratorSummarizationContextFewShotExamplesSummarizationSectionListStruct) {
	if err := d.validatePutSummarizationSectionListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSummarizationSectionList",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ResetConversationContext() {
	_jsii_.InvokeVoid(
		d,
		"resetConversationContext",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ResetExtraInfo() {
	_jsii_.InvokeVoid(
		d,
		"resetExtraInfo",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ResetSummarizationSectionList() {
	_jsii_.InvokeVoid(
		d,
		"resetSummarizationSectionList",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

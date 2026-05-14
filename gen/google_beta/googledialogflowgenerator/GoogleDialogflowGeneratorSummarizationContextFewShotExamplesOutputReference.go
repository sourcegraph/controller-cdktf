package googledialogflowgenerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowgenerator/internal"
)

type GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference interface {
	cdktf.ComplexObject
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
	ConversationContext() GoogleDialogflowGeneratorSummarizationContextFewShotExamplesConversationContextOutputReference
	ConversationContextInput() *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesConversationContext
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Output() GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputOutputReference
	OutputInput() *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutput
	SummarizationSectionList() GoogleDialogflowGeneratorSummarizationContextFewShotExamplesSummarizationSectionListStructOutputReference
	SummarizationSectionListInput() *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesSummarizationSectionListStruct
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
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
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
	PutConversationContext(value *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesConversationContext)
	PutOutput(value *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutput)
	PutSummarizationSectionList(value *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesSummarizationSectionListStruct)
	ResetConversationContext()
	ResetExtraInfo()
	ResetSummarizationSectionList()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference
type jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ConversationContext() GoogleDialogflowGeneratorSummarizationContextFewShotExamplesConversationContextOutputReference {
	var returns GoogleDialogflowGeneratorSummarizationContextFewShotExamplesConversationContextOutputReference
	_jsii_.Get(
		j,
		"conversationContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ConversationContextInput() *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesConversationContext {
	var returns *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesConversationContext
	_jsii_.Get(
		j,
		"conversationContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ExtraInfo() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ExtraInfoInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) Output() GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputOutputReference {
	var returns GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputOutputReference
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) OutputInput() *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutput {
	var returns *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutput
	_jsii_.Get(
		j,
		"outputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) SummarizationSectionList() GoogleDialogflowGeneratorSummarizationContextFewShotExamplesSummarizationSectionListStructOutputReference {
	var returns GoogleDialogflowGeneratorSummarizationContextFewShotExamplesSummarizationSectionListStructOutputReference
	_jsii_.Get(
		j,
		"summarizationSectionList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) SummarizationSectionListInput() *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesSummarizationSectionListStruct {
	var returns *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesSummarizationSectionListStruct
	_jsii_.Get(
		j,
		"summarizationSectionListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowGenerator.GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference_Override(g GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowGenerator.GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference)SetExtraInfo(val *map[string]*string) {
	if err := j.validateSetExtraInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraInfo",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) PutConversationContext(value *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesConversationContext) {
	if err := g.validatePutConversationContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConversationContext",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) PutOutput(value *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutput) {
	if err := g.validatePutOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOutput",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) PutSummarizationSectionList(value *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesSummarizationSectionListStruct) {
	if err := g.validatePutSummarizationSectionListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSummarizationSectionList",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ResetConversationContext() {
	_jsii_.InvokeVoid(
		g,
		"resetConversationContext",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ResetExtraInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetExtraInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ResetSummarizationSectionList() {
	_jsii_.InvokeVoid(
		g,
		"resetSummarizationSectionList",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


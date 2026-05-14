package googlecesevaluation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesevaluation/internal"
)

type GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference interface {
	cdktf.ComplexObject
	AgentTransfer() GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksAgentTransferOutputReference
	AgentTransferInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksAgentTransfer
	Blob() GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksBlobOutputReference
	BlobInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksBlob
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
	// Experimental.
	Fqn() *string
	Image() GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksImageOutputReference
	ImageInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksImage
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Text() *string
	SetText(val *string)
	TextInput() *string
	ToolCall() GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolCallOutputReference
	ToolCallInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolCall
	ToolResponse() GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference
	ToolResponseInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponse
	UpdatedVariables() *map[string]*string
	SetUpdatedVariables(val *map[string]*string)
	UpdatedVariablesInput() *map[string]*string
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
	PutAgentTransfer(value *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksAgentTransfer)
	PutBlob(value *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksBlob)
	PutImage(value *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksImage)
	PutToolCall(value *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolCall)
	PutToolResponse(value *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponse)
	ResetAgentTransfer()
	ResetBlob()
	ResetImage()
	ResetText()
	ResetToolCall()
	ResetToolResponse()
	ResetUpdatedVariables()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference
type jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) AgentTransfer() GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksAgentTransferOutputReference {
	var returns GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksAgentTransferOutputReference
	_jsii_.Get(
		j,
		"agentTransfer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) AgentTransferInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksAgentTransfer {
	var returns *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksAgentTransfer
	_jsii_.Get(
		j,
		"agentTransferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) Blob() GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksBlobOutputReference {
	var returns GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksBlobOutputReference
	_jsii_.Get(
		j,
		"blob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) BlobInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksBlob {
	var returns *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksBlob
	_jsii_.Get(
		j,
		"blobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) Image() GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksImageOutputReference {
	var returns GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksImageOutputReference
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) ImageInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksImage {
	var returns *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksImage
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) Text() *string {
	var returns *string
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) TextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) ToolCall() GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolCallOutputReference {
	var returns GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolCallOutputReference
	_jsii_.Get(
		j,
		"toolCall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) ToolCallInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolCall {
	var returns *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolCall
	_jsii_.Get(
		j,
		"toolCallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) ToolResponse() GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference {
	var returns GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponseOutputReference
	_jsii_.Get(
		j,
		"toolResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) ToolResponseInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponse {
	var returns *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponse
	_jsii_.Get(
		j,
		"toolResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) UpdatedVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"updatedVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) UpdatedVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"updatedVariablesInput",
		&returns,
	)
	return returns
}


func NewGoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesEvaluation.GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference_Override(g GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesEvaluation.GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference)SetText(val *string) {
	if err := j.validateSetTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"text",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference)SetUpdatedVariables(val *map[string]*string) {
	if err := j.validateSetUpdatedVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedVariables",
		val,
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) PutAgentTransfer(value *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksAgentTransfer) {
	if err := g.validatePutAgentTransferParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAgentTransfer",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) PutBlob(value *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksBlob) {
	if err := g.validatePutBlobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBlob",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) PutImage(value *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksImage) {
	if err := g.validatePutImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putImage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) PutToolCall(value *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolCall) {
	if err := g.validatePutToolCallParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putToolCall",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) PutToolResponse(value *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksToolResponse) {
	if err := g.validatePutToolResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putToolResponse",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) ResetAgentTransfer() {
	_jsii_.InvokeVoid(
		g,
		"resetAgentTransfer",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) ResetBlob() {
	_jsii_.InvokeVoid(
		g,
		"resetBlob",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		g,
		"resetImage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) ResetText() {
	_jsii_.InvokeVoid(
		g,
		"resetText",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) ResetToolCall() {
	_jsii_.InvokeVoid(
		g,
		"resetToolCall",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) ResetToolResponse() {
	_jsii_.InvokeVoid(
		g,
		"resetToolResponse",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) ResetUpdatedVariables() {
	_jsii_.InvokeVoid(
		g,
		"resetUpdatedVariables",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseChunksOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


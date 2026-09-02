package googlecesevaluation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesevaluation/internal"
)

type GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference interface {
	cdktn.ComplexObject
	AgentTransfer() GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksAgentTransferOutputReference
	AgentTransferInput() *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksAgentTransfer
	Blob() GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksBlobOutputReference
	BlobInput() *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksBlob
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
	Image() GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksImageOutputReference
	ImageInput() *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksImage
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Text() *string
	SetText(val *string)
	TextInput() *string
	ToolCall() GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolCallOutputReference
	ToolCallInput() *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolCall
	ToolResponse() GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference
	ToolResponseInput() *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponse
	UpdatedVariables() *map[string]*string
	SetUpdatedVariables(val *map[string]*string)
	UpdatedVariablesInput() *map[string]*string
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
	PutAgentTransfer(value *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksAgentTransfer)
	PutBlob(value *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksBlob)
	PutImage(value *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksImage)
	PutToolCall(value *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolCall)
	PutToolResponse(value *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponse)
	ResetAgentTransfer()
	ResetBlob()
	ResetImage()
	ResetText()
	ResetToolCall()
	ResetToolResponse()
	ResetUpdatedVariables()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference
type jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) AgentTransfer() GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksAgentTransferOutputReference {
	var returns GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksAgentTransferOutputReference
	_jsii_.Get(
		j,
		"agentTransfer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) AgentTransferInput() *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksAgentTransfer {
	var returns *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksAgentTransfer
	_jsii_.Get(
		j,
		"agentTransferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) Blob() GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksBlobOutputReference {
	var returns GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksBlobOutputReference
	_jsii_.Get(
		j,
		"blob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) BlobInput() *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksBlob {
	var returns *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksBlob
	_jsii_.Get(
		j,
		"blobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) Image() GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksImageOutputReference {
	var returns GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksImageOutputReference
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) ImageInput() *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksImage {
	var returns *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksImage
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) Text() *string {
	var returns *string
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) TextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) ToolCall() GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolCallOutputReference {
	var returns GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolCallOutputReference
	_jsii_.Get(
		j,
		"toolCall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) ToolCallInput() *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolCall {
	var returns *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolCall
	_jsii_.Get(
		j,
		"toolCallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) ToolResponse() GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference {
	var returns GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponseOutputReference
	_jsii_.Get(
		j,
		"toolResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) ToolResponseInput() *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponse {
	var returns *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponse
	_jsii_.Get(
		j,
		"toolResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) UpdatedVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"updatedVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) UpdatedVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"updatedVariablesInput",
		&returns,
	)
	return returns
}


func NewGoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesEvaluation.GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference_Override(g GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesEvaluation.GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference)SetText(val *string) {
	if err := j.validateSetTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"text",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference)SetUpdatedVariables(val *map[string]*string) {
	if err := j.validateSetUpdatedVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedVariables",
		val,
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) PutAgentTransfer(value *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksAgentTransfer) {
	if err := g.validatePutAgentTransferParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAgentTransfer",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) PutBlob(value *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksBlob) {
	if err := g.validatePutBlobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBlob",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) PutImage(value *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksImage) {
	if err := g.validatePutImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putImage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) PutToolCall(value *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolCall) {
	if err := g.validatePutToolCallParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putToolCall",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) PutToolResponse(value *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksToolResponse) {
	if err := g.validatePutToolResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putToolResponse",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) ResetAgentTransfer() {
	_jsii_.InvokeVoid(
		g,
		"resetAgentTransfer",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) ResetBlob() {
	_jsii_.InvokeVoid(
		g,
		"resetBlob",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		g,
		"resetImage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) ResetText() {
	_jsii_.InvokeVoid(
		g,
		"resetText",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) ResetToolCall() {
	_jsii_.InvokeVoid(
		g,
		"resetToolCall",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) ResetToolResponse() {
	_jsii_.InvokeVoid(
		g,
		"resetToolResponse",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) ResetUpdatedVariables() {
	_jsii_.InvokeVoid(
		g,
		"resetUpdatedVariables",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package googlecesexample

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesexample/internal"
)

type GoogleCesExampleMessagesChunksOutputReference interface {
	cdktn.ComplexObject
	AgentTransfer() GoogleCesExampleMessagesChunksAgentTransferOutputReference
	AgentTransferInput() *GoogleCesExampleMessagesChunksAgentTransfer
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
	Image() GoogleCesExampleMessagesChunksImageOutputReference
	ImageInput() *GoogleCesExampleMessagesChunksImage
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
	ToolCall() GoogleCesExampleMessagesChunksToolCallOutputReference
	ToolCallInput() *GoogleCesExampleMessagesChunksToolCall
	ToolResponse() GoogleCesExampleMessagesChunksToolResponseOutputReference
	ToolResponseInput() *GoogleCesExampleMessagesChunksToolResponse
	UpdatedVariables() *string
	SetUpdatedVariables(val *string)
	UpdatedVariablesInput() *string
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
	PutAgentTransfer(value *GoogleCesExampleMessagesChunksAgentTransfer)
	PutImage(value *GoogleCesExampleMessagesChunksImage)
	PutToolCall(value *GoogleCesExampleMessagesChunksToolCall)
	PutToolResponse(value *GoogleCesExampleMessagesChunksToolResponse)
	ResetAgentTransfer()
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

// The jsii proxy struct for GoogleCesExampleMessagesChunksOutputReference
type jsiiProxy_GoogleCesExampleMessagesChunksOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) AgentTransfer() GoogleCesExampleMessagesChunksAgentTransferOutputReference {
	var returns GoogleCesExampleMessagesChunksAgentTransferOutputReference
	_jsii_.Get(
		j,
		"agentTransfer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) AgentTransferInput() *GoogleCesExampleMessagesChunksAgentTransfer {
	var returns *GoogleCesExampleMessagesChunksAgentTransfer
	_jsii_.Get(
		j,
		"agentTransferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) Image() GoogleCesExampleMessagesChunksImageOutputReference {
	var returns GoogleCesExampleMessagesChunksImageOutputReference
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) ImageInput() *GoogleCesExampleMessagesChunksImage {
	var returns *GoogleCesExampleMessagesChunksImage
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) Text() *string {
	var returns *string
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) TextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) ToolCall() GoogleCesExampleMessagesChunksToolCallOutputReference {
	var returns GoogleCesExampleMessagesChunksToolCallOutputReference
	_jsii_.Get(
		j,
		"toolCall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) ToolCallInput() *GoogleCesExampleMessagesChunksToolCall {
	var returns *GoogleCesExampleMessagesChunksToolCall
	_jsii_.Get(
		j,
		"toolCallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) ToolResponse() GoogleCesExampleMessagesChunksToolResponseOutputReference {
	var returns GoogleCesExampleMessagesChunksToolResponseOutputReference
	_jsii_.Get(
		j,
		"toolResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) ToolResponseInput() *GoogleCesExampleMessagesChunksToolResponse {
	var returns *GoogleCesExampleMessagesChunksToolResponse
	_jsii_.Get(
		j,
		"toolResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) UpdatedVariables() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) UpdatedVariablesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedVariablesInput",
		&returns,
	)
	return returns
}


func NewGoogleCesExampleMessagesChunksOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleCesExampleMessagesChunksOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesExampleMessagesChunksOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesExampleMessagesChunksOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesExample.GoogleCesExampleMessagesChunksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleCesExampleMessagesChunksOutputReference_Override(g GoogleCesExampleMessagesChunksOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesExample.GoogleCesExampleMessagesChunksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference)SetText(val *string) {
	if err := j.validateSetTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"text",
		val,
	)
}

func (j *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference)SetUpdatedVariables(val *string) {
	if err := j.validateSetUpdatedVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedVariables",
		val,
	)
}

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) PutAgentTransfer(value *GoogleCesExampleMessagesChunksAgentTransfer) {
	if err := g.validatePutAgentTransferParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAgentTransfer",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) PutImage(value *GoogleCesExampleMessagesChunksImage) {
	if err := g.validatePutImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putImage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) PutToolCall(value *GoogleCesExampleMessagesChunksToolCall) {
	if err := g.validatePutToolCallParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putToolCall",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) PutToolResponse(value *GoogleCesExampleMessagesChunksToolResponse) {
	if err := g.validatePutToolResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putToolResponse",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) ResetAgentTransfer() {
	_jsii_.InvokeVoid(
		g,
		"resetAgentTransfer",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		g,
		"resetImage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) ResetText() {
	_jsii_.InvokeVoid(
		g,
		"resetText",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) ResetToolCall() {
	_jsii_.InvokeVoid(
		g,
		"resetToolCall",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) ResetToolResponse() {
	_jsii_.InvokeVoid(
		g,
		"resetToolResponse",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) ResetUpdatedVariables() {
	_jsii_.InvokeVoid(
		g,
		"resetUpdatedVariables",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesExampleMessagesChunksOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


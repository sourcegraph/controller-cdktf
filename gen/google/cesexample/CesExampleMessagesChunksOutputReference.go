package cesexample

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cesexample/internal"
)

type CesExampleMessagesChunksOutputReference interface {
	cdktf.ComplexObject
	AgentTransfer() CesExampleMessagesChunksAgentTransferOutputReference
	AgentTransferInput() *CesExampleMessagesChunksAgentTransfer
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
	Image() CesExampleMessagesChunksImageOutputReference
	ImageInput() *CesExampleMessagesChunksImage
	InternalValue() any
	SetInternalValue(val any)
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
	ToolCall() CesExampleMessagesChunksToolCallOutputReference
	ToolCallInput() *CesExampleMessagesChunksToolCall
	ToolResponse() CesExampleMessagesChunksToolResponseOutputReference
	ToolResponseInput() *CesExampleMessagesChunksToolResponse
	UpdatedVariables() *string
	SetUpdatedVariables(val *string)
	UpdatedVariablesInput() *string
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
	PutAgentTransfer(value *CesExampleMessagesChunksAgentTransfer)
	PutImage(value *CesExampleMessagesChunksImage)
	PutToolCall(value *CesExampleMessagesChunksToolCall)
	PutToolResponse(value *CesExampleMessagesChunksToolResponse)
	ResetAgentTransfer()
	ResetImage()
	ResetText()
	ResetToolCall()
	ResetToolResponse()
	ResetUpdatedVariables()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesExampleMessagesChunksOutputReference
type jsiiProxy_CesExampleMessagesChunksOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) AgentTransfer() CesExampleMessagesChunksAgentTransferOutputReference {
	var returns CesExampleMessagesChunksAgentTransferOutputReference
	_jsii_.Get(
		j,
		"agentTransfer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) AgentTransferInput() *CesExampleMessagesChunksAgentTransfer {
	var returns *CesExampleMessagesChunksAgentTransfer
	_jsii_.Get(
		j,
		"agentTransferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) Image() CesExampleMessagesChunksImageOutputReference {
	var returns CesExampleMessagesChunksImageOutputReference
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) ImageInput() *CesExampleMessagesChunksImage {
	var returns *CesExampleMessagesChunksImage
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) Text() *string {
	var returns *string
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) TextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) ToolCall() CesExampleMessagesChunksToolCallOutputReference {
	var returns CesExampleMessagesChunksToolCallOutputReference
	_jsii_.Get(
		j,
		"toolCall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) ToolCallInput() *CesExampleMessagesChunksToolCall {
	var returns *CesExampleMessagesChunksToolCall
	_jsii_.Get(
		j,
		"toolCallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) ToolResponse() CesExampleMessagesChunksToolResponseOutputReference {
	var returns CesExampleMessagesChunksToolResponseOutputReference
	_jsii_.Get(
		j,
		"toolResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) ToolResponseInput() *CesExampleMessagesChunksToolResponse {
	var returns *CesExampleMessagesChunksToolResponse
	_jsii_.Get(
		j,
		"toolResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) UpdatedVariables() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) UpdatedVariablesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedVariablesInput",
		&returns,
	)
	return returns
}

func NewCesExampleMessagesChunksOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CesExampleMessagesChunksOutputReference {
	_init_.Initialize()

	if err := validateNewCesExampleMessagesChunksOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesExampleMessagesChunksOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cesExample.CesExampleMessagesChunksOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCesExampleMessagesChunksOutputReference_Override(c CesExampleMessagesChunksOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesExample.CesExampleMessagesChunksOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) SetText(val *string) {
	if err := j.validateSetTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"text",
		val,
	)
}

func (j *jsiiProxy_CesExampleMessagesChunksOutputReference) SetUpdatedVariables(val *string) {
	if err := j.validateSetUpdatedVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedVariables",
		val,
	)
}

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) PutAgentTransfer(value *CesExampleMessagesChunksAgentTransfer) {
	if err := c.validatePutAgentTransferParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAgentTransfer",
		[]any{value},
	)
}

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) PutImage(value *CesExampleMessagesChunksImage) {
	if err := c.validatePutImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putImage",
		[]any{value},
	)
}

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) PutToolCall(value *CesExampleMessagesChunksToolCall) {
	if err := c.validatePutToolCallParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putToolCall",
		[]any{value},
	)
}

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) PutToolResponse(value *CesExampleMessagesChunksToolResponse) {
	if err := c.validatePutToolResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putToolResponse",
		[]any{value},
	)
}

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) ResetAgentTransfer() {
	_jsii_.InvokeVoid(
		c,
		"resetAgentTransfer",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		c,
		"resetImage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) ResetText() {
	_jsii_.InvokeVoid(
		c,
		"resetText",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) ResetToolCall() {
	_jsii_.InvokeVoid(
		c,
		"resetToolCall",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) ResetToolResponse() {
	_jsii_.InvokeVoid(
		c,
		"resetToolResponse",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) ResetUpdatedVariables() {
	_jsii_.InvokeVoid(
		c,
		"resetUpdatedVariables",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_CesExampleMessagesChunksOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

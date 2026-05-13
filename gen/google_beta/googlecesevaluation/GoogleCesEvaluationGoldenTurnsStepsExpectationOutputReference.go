package googlecesevaluation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesevaluation/internal"
)

type GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference interface {
	cdktf.ComplexObject
	AgentResponse() GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseOutputReference
	AgentResponseInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponse
	AgentTransfer() GoogleCesEvaluationGoldenTurnsStepsExpectationAgentTransferOutputReference
	AgentTransferInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentTransfer
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
	InternalValue() *GoogleCesEvaluationGoldenTurnsStepsExpectation
	SetInternalValue(val *GoogleCesEvaluationGoldenTurnsStepsExpectation)
	MockToolResponse() GoogleCesEvaluationGoldenTurnsStepsExpectationMockToolResponseOutputReference
	MockToolResponseInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationMockToolResponse
	Note() *string
	SetNote(val *string)
	NoteInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ToolCall() GoogleCesEvaluationGoldenTurnsStepsExpectationToolCallOutputReference
	ToolCallInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationToolCall
	ToolResponse() GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseOutputReference
	ToolResponseInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponse
	UpdatedVariables() GoogleCesEvaluationGoldenTurnsStepsExpectationUpdatedVariablesOutputReference
	UpdatedVariablesInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationUpdatedVariables
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
	PutAgentResponse(value *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponse)
	PutAgentTransfer(value *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentTransfer)
	PutMockToolResponse(value *GoogleCesEvaluationGoldenTurnsStepsExpectationMockToolResponse)
	PutToolCall(value *GoogleCesEvaluationGoldenTurnsStepsExpectationToolCall)
	PutToolResponse(value *GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponse)
	PutUpdatedVariables(value *GoogleCesEvaluationGoldenTurnsStepsExpectationUpdatedVariables)
	ResetAgentResponse()
	ResetAgentTransfer()
	ResetMockToolResponse()
	ResetNote()
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

// The jsii proxy struct for GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference
type jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) AgentResponse() GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseOutputReference {
	var returns GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponseOutputReference
	_jsii_.Get(
		j,
		"agentResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) AgentResponseInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponse {
	var returns *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponse
	_jsii_.Get(
		j,
		"agentResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) AgentTransfer() GoogleCesEvaluationGoldenTurnsStepsExpectationAgentTransferOutputReference {
	var returns GoogleCesEvaluationGoldenTurnsStepsExpectationAgentTransferOutputReference
	_jsii_.Get(
		j,
		"agentTransfer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) AgentTransferInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentTransfer {
	var returns *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentTransfer
	_jsii_.Get(
		j,
		"agentTransferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) InternalValue() *GoogleCesEvaluationGoldenTurnsStepsExpectation {
	var returns *GoogleCesEvaluationGoldenTurnsStepsExpectation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) MockToolResponse() GoogleCesEvaluationGoldenTurnsStepsExpectationMockToolResponseOutputReference {
	var returns GoogleCesEvaluationGoldenTurnsStepsExpectationMockToolResponseOutputReference
	_jsii_.Get(
		j,
		"mockToolResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) MockToolResponseInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationMockToolResponse {
	var returns *GoogleCesEvaluationGoldenTurnsStepsExpectationMockToolResponse
	_jsii_.Get(
		j,
		"mockToolResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) Note() *string {
	var returns *string
	_jsii_.Get(
		j,
		"note",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) NoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) ToolCall() GoogleCesEvaluationGoldenTurnsStepsExpectationToolCallOutputReference {
	var returns GoogleCesEvaluationGoldenTurnsStepsExpectationToolCallOutputReference
	_jsii_.Get(
		j,
		"toolCall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) ToolCallInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationToolCall {
	var returns *GoogleCesEvaluationGoldenTurnsStepsExpectationToolCall
	_jsii_.Get(
		j,
		"toolCallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) ToolResponse() GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseOutputReference {
	var returns GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponseOutputReference
	_jsii_.Get(
		j,
		"toolResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) ToolResponseInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponse {
	var returns *GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponse
	_jsii_.Get(
		j,
		"toolResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) UpdatedVariables() GoogleCesEvaluationGoldenTurnsStepsExpectationUpdatedVariablesOutputReference {
	var returns GoogleCesEvaluationGoldenTurnsStepsExpectationUpdatedVariablesOutputReference
	_jsii_.Get(
		j,
		"updatedVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) UpdatedVariablesInput() *GoogleCesEvaluationGoldenTurnsStepsExpectationUpdatedVariables {
	var returns *GoogleCesEvaluationGoldenTurnsStepsExpectationUpdatedVariables
	_jsii_.Get(
		j,
		"updatedVariablesInput",
		&returns,
	)
	return returns
}


func NewGoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesEvaluationGoldenTurnsStepsExpectationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesEvaluation.GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference_Override(g GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesEvaluation.GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference)SetInternalValue(val *GoogleCesEvaluationGoldenTurnsStepsExpectation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference)SetNote(val *string) {
	if err := j.validateSetNoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"note",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) PutAgentResponse(value *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponse) {
	if err := g.validatePutAgentResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAgentResponse",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) PutAgentTransfer(value *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentTransfer) {
	if err := g.validatePutAgentTransferParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAgentTransfer",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) PutMockToolResponse(value *GoogleCesEvaluationGoldenTurnsStepsExpectationMockToolResponse) {
	if err := g.validatePutMockToolResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMockToolResponse",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) PutToolCall(value *GoogleCesEvaluationGoldenTurnsStepsExpectationToolCall) {
	if err := g.validatePutToolCallParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putToolCall",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) PutToolResponse(value *GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponse) {
	if err := g.validatePutToolResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putToolResponse",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) PutUpdatedVariables(value *GoogleCesEvaluationGoldenTurnsStepsExpectationUpdatedVariables) {
	if err := g.validatePutUpdatedVariablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUpdatedVariables",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) ResetAgentResponse() {
	_jsii_.InvokeVoid(
		g,
		"resetAgentResponse",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) ResetAgentTransfer() {
	_jsii_.InvokeVoid(
		g,
		"resetAgentTransfer",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) ResetMockToolResponse() {
	_jsii_.InvokeVoid(
		g,
		"resetMockToolResponse",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) ResetNote() {
	_jsii_.InvokeVoid(
		g,
		"resetNote",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) ResetToolCall() {
	_jsii_.InvokeVoid(
		g,
		"resetToolCall",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) ResetToolResponse() {
	_jsii_.InvokeVoid(
		g,
		"resetToolResponse",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) ResetUpdatedVariables() {
	_jsii_.InvokeVoid(
		g,
		"resetUpdatedVariables",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsExpectationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


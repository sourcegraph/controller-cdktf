package googlecesevaluation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesevaluation/internal"
)

type GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference interface {
	cdktn.ComplexObject
	Audio() *string
	SetAudio(val *string)
	AudioInput() *string
	Blob() GoogleCesEvaluationGoldenTurnsStepsUserInputBlobOutputReference
	BlobInput() *GoogleCesEvaluationGoldenTurnsStepsUserInputBlob
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
	Dtmf() *string
	SetDtmf(val *string)
	DtmfInput() *string
	Event() GoogleCesEvaluationGoldenTurnsStepsUserInputEventOutputReference
	EventInput() *GoogleCesEvaluationGoldenTurnsStepsUserInputEvent
	// Experimental.
	Fqn() *string
	Image() GoogleCesEvaluationGoldenTurnsStepsUserInputImageOutputReference
	ImageInput() *GoogleCesEvaluationGoldenTurnsStepsUserInputImage
	InternalValue() *GoogleCesEvaluationGoldenTurnsStepsUserInput
	SetInternalValue(val *GoogleCesEvaluationGoldenTurnsStepsUserInput)
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
	ToolResponses() GoogleCesEvaluationGoldenTurnsStepsUserInputToolResponsesOutputReference
	ToolResponsesInput() *GoogleCesEvaluationGoldenTurnsStepsUserInputToolResponses
	Variables() *map[string]*string
	SetVariables(val *map[string]*string)
	VariablesInput() *map[string]*string
	WillContinue() interface{}
	SetWillContinue(val interface{})
	WillContinueInput() interface{}
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
	PutBlob(value *GoogleCesEvaluationGoldenTurnsStepsUserInputBlob)
	PutEvent(value *GoogleCesEvaluationGoldenTurnsStepsUserInputEvent)
	PutImage(value *GoogleCesEvaluationGoldenTurnsStepsUserInputImage)
	PutToolResponses(value *GoogleCesEvaluationGoldenTurnsStepsUserInputToolResponses)
	ResetAudio()
	ResetBlob()
	ResetDtmf()
	ResetEvent()
	ResetImage()
	ResetText()
	ResetToolResponses()
	ResetVariables()
	ResetWillContinue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference
type jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) Audio() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) AudioInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) Blob() GoogleCesEvaluationGoldenTurnsStepsUserInputBlobOutputReference {
	var returns GoogleCesEvaluationGoldenTurnsStepsUserInputBlobOutputReference
	_jsii_.Get(
		j,
		"blob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) BlobInput() *GoogleCesEvaluationGoldenTurnsStepsUserInputBlob {
	var returns *GoogleCesEvaluationGoldenTurnsStepsUserInputBlob
	_jsii_.Get(
		j,
		"blobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) Dtmf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dtmf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) DtmfInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dtmfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) Event() GoogleCesEvaluationGoldenTurnsStepsUserInputEventOutputReference {
	var returns GoogleCesEvaluationGoldenTurnsStepsUserInputEventOutputReference
	_jsii_.Get(
		j,
		"event",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) EventInput() *GoogleCesEvaluationGoldenTurnsStepsUserInputEvent {
	var returns *GoogleCesEvaluationGoldenTurnsStepsUserInputEvent
	_jsii_.Get(
		j,
		"eventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) Image() GoogleCesEvaluationGoldenTurnsStepsUserInputImageOutputReference {
	var returns GoogleCesEvaluationGoldenTurnsStepsUserInputImageOutputReference
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) ImageInput() *GoogleCesEvaluationGoldenTurnsStepsUserInputImage {
	var returns *GoogleCesEvaluationGoldenTurnsStepsUserInputImage
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) InternalValue() *GoogleCesEvaluationGoldenTurnsStepsUserInput {
	var returns *GoogleCesEvaluationGoldenTurnsStepsUserInput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) Text() *string {
	var returns *string
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) TextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) ToolResponses() GoogleCesEvaluationGoldenTurnsStepsUserInputToolResponsesOutputReference {
	var returns GoogleCesEvaluationGoldenTurnsStepsUserInputToolResponsesOutputReference
	_jsii_.Get(
		j,
		"toolResponses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) ToolResponsesInput() *GoogleCesEvaluationGoldenTurnsStepsUserInputToolResponses {
	var returns *GoogleCesEvaluationGoldenTurnsStepsUserInputToolResponses
	_jsii_.Get(
		j,
		"toolResponsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) Variables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) VariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) WillContinue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"willContinue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) WillContinueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"willContinueInput",
		&returns,
	)
	return returns
}


func NewGoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesEvaluationGoldenTurnsStepsUserInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesEvaluation.GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference_Override(g GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesEvaluation.GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference)SetAudio(val *string) {
	if err := j.validateSetAudioParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audio",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference)SetDtmf(val *string) {
	if err := j.validateSetDtmfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dtmf",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference)SetInternalValue(val *GoogleCesEvaluationGoldenTurnsStepsUserInput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference)SetText(val *string) {
	if err := j.validateSetTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"text",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference)SetVariables(val *map[string]*string) {
	if err := j.validateSetVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variables",
		val,
	)
}

func (j *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference)SetWillContinue(val interface{}) {
	if err := j.validateSetWillContinueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"willContinue",
		val,
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) PutBlob(value *GoogleCesEvaluationGoldenTurnsStepsUserInputBlob) {
	if err := g.validatePutBlobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBlob",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) PutEvent(value *GoogleCesEvaluationGoldenTurnsStepsUserInputEvent) {
	if err := g.validatePutEventParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEvent",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) PutImage(value *GoogleCesEvaluationGoldenTurnsStepsUserInputImage) {
	if err := g.validatePutImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putImage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) PutToolResponses(value *GoogleCesEvaluationGoldenTurnsStepsUserInputToolResponses) {
	if err := g.validatePutToolResponsesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putToolResponses",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) ResetAudio() {
	_jsii_.InvokeVoid(
		g,
		"resetAudio",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) ResetBlob() {
	_jsii_.InvokeVoid(
		g,
		"resetBlob",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) ResetDtmf() {
	_jsii_.InvokeVoid(
		g,
		"resetDtmf",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) ResetEvent() {
	_jsii_.InvokeVoid(
		g,
		"resetEvent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		g,
		"resetImage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) ResetText() {
	_jsii_.InvokeVoid(
		g,
		"resetText",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) ResetToolResponses() {
	_jsii_.InvokeVoid(
		g,
		"resetToolResponses",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) ResetVariables() {
	_jsii_.InvokeVoid(
		g,
		"resetVariables",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) ResetWillContinue() {
	_jsii_.InvokeVoid(
		g,
		"resetWillContinue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesEvaluationGoldenTurnsStepsUserInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


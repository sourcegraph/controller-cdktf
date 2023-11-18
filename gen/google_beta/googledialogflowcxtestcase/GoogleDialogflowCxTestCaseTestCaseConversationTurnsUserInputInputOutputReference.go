package googledialogflowcxtestcase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowcxtestcase/internal"
)

type GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dtmf() GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputDtmfOutputReference
	DtmfInput() *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputDtmf
	Event() GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputEventOutputReference
	EventInput() *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputEvent
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInput
	SetInternalValue(val *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInput)
	LanguageCode() *string
	SetLanguageCode(val *string)
	LanguageCodeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Text() GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputTextOutputReference
	TextInput() *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputText
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
	PutDtmf(value *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputDtmf)
	PutEvent(value *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputEvent)
	PutText(value *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputText)
	ResetDtmf()
	ResetEvent()
	ResetLanguageCode()
	ResetText()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference
type jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) Dtmf() GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputDtmfOutputReference {
	var returns GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputDtmfOutputReference
	_jsii_.Get(
		j,
		"dtmf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) DtmfInput() *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputDtmf {
	var returns *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputDtmf
	_jsii_.Get(
		j,
		"dtmfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) Event() GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputEventOutputReference {
	var returns GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputEventOutputReference
	_jsii_.Get(
		j,
		"event",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) EventInput() *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputEvent {
	var returns *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputEvent
	_jsii_.Get(
		j,
		"eventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) InternalValue() *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInput {
	var returns *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) LanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) LanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) Text() GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputTextOutputReference {
	var returns GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputTextOutputReference
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) TextInput() *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputText {
	var returns *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputText
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxTestCase.GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference_Override(g GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxTestCase.GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference)SetInternalValue(val *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference)SetLanguageCode(val *string) {
	if err := j.validateSetLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageCode",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) PutDtmf(value *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputDtmf) {
	if err := g.validatePutDtmfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDtmf",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) PutEvent(value *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputEvent) {
	if err := g.validatePutEventParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEvent",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) PutText(value *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputText) {
	if err := g.validatePutTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putText",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) ResetDtmf() {
	_jsii_.InvokeVoid(
		g,
		"resetDtmf",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) ResetEvent() {
	_jsii_.InvokeVoid(
		g,
		"resetEvent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) ResetLanguageCode() {
	_jsii_.InvokeVoid(
		g,
		"resetLanguageCode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) ResetText() {
	_jsii_.InvokeVoid(
		g,
		"resetText",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


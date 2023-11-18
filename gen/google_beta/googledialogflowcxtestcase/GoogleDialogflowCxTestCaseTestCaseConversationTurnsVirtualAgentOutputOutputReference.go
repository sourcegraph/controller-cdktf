package googledialogflowcxtestcase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowcxtestcase/internal"
)

type GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference interface {
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
	CurrentPage() GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputCurrentPageOutputReference
	CurrentPageInput() *GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputCurrentPage
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutput
	SetInternalValue(val *GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutput)
	SessionParameters() *string
	SetSessionParameters(val *string)
	SessionParametersInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TextResponses() GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputTextResponsesList
	TextResponsesInput() interface{}
	TriggeredIntent() GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputTriggeredIntentOutputReference
	TriggeredIntentInput() *GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputTriggeredIntent
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
	PutCurrentPage(value *GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputCurrentPage)
	PutTextResponses(value interface{})
	PutTriggeredIntent(value *GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputTriggeredIntent)
	ResetCurrentPage()
	ResetSessionParameters()
	ResetTextResponses()
	ResetTriggeredIntent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference
type jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) CurrentPage() GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputCurrentPageOutputReference {
	var returns GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputCurrentPageOutputReference
	_jsii_.Get(
		j,
		"currentPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) CurrentPageInput() *GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputCurrentPage {
	var returns *GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputCurrentPage
	_jsii_.Get(
		j,
		"currentPageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) InternalValue() *GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutput {
	var returns *GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) SessionParameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) SessionParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) TextResponses() GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputTextResponsesList {
	var returns GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputTextResponsesList
	_jsii_.Get(
		j,
		"textResponses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) TextResponsesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"textResponsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) TriggeredIntent() GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputTriggeredIntentOutputReference {
	var returns GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputTriggeredIntentOutputReference
	_jsii_.Get(
		j,
		"triggeredIntent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) TriggeredIntentInput() *GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputTriggeredIntent {
	var returns *GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputTriggeredIntent
	_jsii_.Get(
		j,
		"triggeredIntentInput",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxTestCase.GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference_Override(g GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxTestCase.GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference)SetInternalValue(val *GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference)SetSessionParameters(val *string) {
	if err := j.validateSetSessionParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionParameters",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) PutCurrentPage(value *GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputCurrentPage) {
	if err := g.validatePutCurrentPageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCurrentPage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) PutTextResponses(value interface{}) {
	if err := g.validatePutTextResponsesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTextResponses",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) PutTriggeredIntent(value *GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputTriggeredIntent) {
	if err := g.validatePutTriggeredIntentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTriggeredIntent",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) ResetCurrentPage() {
	_jsii_.InvokeVoid(
		g,
		"resetCurrentPage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) ResetSessionParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetSessionParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) ResetTextResponses() {
	_jsii_.InvokeVoid(
		g,
		"resetTextResponses",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) ResetTriggeredIntent() {
	_jsii_.InvokeVoid(
		g,
		"resetTriggeredIntent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package googlevertexaireasoningengine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevertexaireasoningengine/internal"
)

type GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference interface {
	cdktn.ComplexObject
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
	DeveloperConnectSource() GoogleVertexAiReasoningEngineSpecSourceCodeSpecDeveloperConnectSourceOutputReference
	DeveloperConnectSourceInput() *GoogleVertexAiReasoningEngineSpecSourceCodeSpecDeveloperConnectSource
	// Experimental.
	Fqn() *string
	ImageSpec() GoogleVertexAiReasoningEngineSpecSourceCodeSpecImageSpecOutputReference
	ImageSpecInput() *GoogleVertexAiReasoningEngineSpecSourceCodeSpecImageSpec
	InlineSource() GoogleVertexAiReasoningEngineSpecSourceCodeSpecInlineSourceOutputReference
	InlineSourceInput() *GoogleVertexAiReasoningEngineSpecSourceCodeSpecInlineSource
	InternalValue() *GoogleVertexAiReasoningEngineSpecSourceCodeSpec
	SetInternalValue(val *GoogleVertexAiReasoningEngineSpecSourceCodeSpec)
	PythonSpec() GoogleVertexAiReasoningEngineSpecSourceCodeSpecPythonSpecOutputReference
	PythonSpecInput() *GoogleVertexAiReasoningEngineSpecSourceCodeSpecPythonSpec
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutDeveloperConnectSource(value *GoogleVertexAiReasoningEngineSpecSourceCodeSpecDeveloperConnectSource)
	PutImageSpec(value *GoogleVertexAiReasoningEngineSpecSourceCodeSpecImageSpec)
	PutInlineSource(value *GoogleVertexAiReasoningEngineSpecSourceCodeSpecInlineSource)
	PutPythonSpec(value *GoogleVertexAiReasoningEngineSpecSourceCodeSpecPythonSpec)
	ResetDeveloperConnectSource()
	ResetImageSpec()
	ResetInlineSource()
	ResetPythonSpec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference
type jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) DeveloperConnectSource() GoogleVertexAiReasoningEngineSpecSourceCodeSpecDeveloperConnectSourceOutputReference {
	var returns GoogleVertexAiReasoningEngineSpecSourceCodeSpecDeveloperConnectSourceOutputReference
	_jsii_.Get(
		j,
		"developerConnectSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) DeveloperConnectSourceInput() *GoogleVertexAiReasoningEngineSpecSourceCodeSpecDeveloperConnectSource {
	var returns *GoogleVertexAiReasoningEngineSpecSourceCodeSpecDeveloperConnectSource
	_jsii_.Get(
		j,
		"developerConnectSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) ImageSpec() GoogleVertexAiReasoningEngineSpecSourceCodeSpecImageSpecOutputReference {
	var returns GoogleVertexAiReasoningEngineSpecSourceCodeSpecImageSpecOutputReference
	_jsii_.Get(
		j,
		"imageSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) ImageSpecInput() *GoogleVertexAiReasoningEngineSpecSourceCodeSpecImageSpec {
	var returns *GoogleVertexAiReasoningEngineSpecSourceCodeSpecImageSpec
	_jsii_.Get(
		j,
		"imageSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) InlineSource() GoogleVertexAiReasoningEngineSpecSourceCodeSpecInlineSourceOutputReference {
	var returns GoogleVertexAiReasoningEngineSpecSourceCodeSpecInlineSourceOutputReference
	_jsii_.Get(
		j,
		"inlineSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) InlineSourceInput() *GoogleVertexAiReasoningEngineSpecSourceCodeSpecInlineSource {
	var returns *GoogleVertexAiReasoningEngineSpecSourceCodeSpecInlineSource
	_jsii_.Get(
		j,
		"inlineSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) InternalValue() *GoogleVertexAiReasoningEngineSpecSourceCodeSpec {
	var returns *GoogleVertexAiReasoningEngineSpecSourceCodeSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) PythonSpec() GoogleVertexAiReasoningEngineSpecSourceCodeSpecPythonSpecOutputReference {
	var returns GoogleVertexAiReasoningEngineSpecSourceCodeSpecPythonSpecOutputReference
	_jsii_.Get(
		j,
		"pythonSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) PythonSpecInput() *GoogleVertexAiReasoningEngineSpecSourceCodeSpecPythonSpec {
	var returns *GoogleVertexAiReasoningEngineSpecSourceCodeSpecPythonSpec
	_jsii_.Get(
		j,
		"pythonSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiReasoningEngine.GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference_Override(g GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiReasoningEngine.GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference)SetInternalValue(val *GoogleVertexAiReasoningEngineSpecSourceCodeSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) PutDeveloperConnectSource(value *GoogleVertexAiReasoningEngineSpecSourceCodeSpecDeveloperConnectSource) {
	if err := g.validatePutDeveloperConnectSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeveloperConnectSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) PutImageSpec(value *GoogleVertexAiReasoningEngineSpecSourceCodeSpecImageSpec) {
	if err := g.validatePutImageSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putImageSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) PutInlineSource(value *GoogleVertexAiReasoningEngineSpecSourceCodeSpecInlineSource) {
	if err := g.validatePutInlineSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInlineSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) PutPythonSpec(value *GoogleVertexAiReasoningEngineSpecSourceCodeSpecPythonSpec) {
	if err := g.validatePutPythonSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPythonSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) ResetDeveloperConnectSource() {
	_jsii_.InvokeVoid(
		g,
		"resetDeveloperConnectSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) ResetImageSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetImageSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) ResetInlineSource() {
	_jsii_.InvokeVoid(
		g,
		"resetInlineSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) ResetPythonSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetPythonSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


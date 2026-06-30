package vertexaireasoningengine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/vertexaireasoningengine/internal"
)

type VertexAiReasoningEngineSpecSourceCodeSpecOutputReference interface {
	cdktf.ComplexObject
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
	DeveloperConnectSource() VertexAiReasoningEngineSpecSourceCodeSpecDeveloperConnectSourceOutputReference
	DeveloperConnectSourceInput() *VertexAiReasoningEngineSpecSourceCodeSpecDeveloperConnectSource
	// Experimental.
	Fqn() *string
	ImageSpec() VertexAiReasoningEngineSpecSourceCodeSpecImageSpecOutputReference
	ImageSpecInput() *VertexAiReasoningEngineSpecSourceCodeSpecImageSpec
	InlineSource() VertexAiReasoningEngineSpecSourceCodeSpecInlineSourceOutputReference
	InlineSourceInput() *VertexAiReasoningEngineSpecSourceCodeSpecInlineSource
	InternalValue() *VertexAiReasoningEngineSpecSourceCodeSpec
	SetInternalValue(val *VertexAiReasoningEngineSpecSourceCodeSpec)
	PythonSpec() VertexAiReasoningEngineSpecSourceCodeSpecPythonSpecOutputReference
	PythonSpecInput() *VertexAiReasoningEngineSpecSourceCodeSpecPythonSpec
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
	PutDeveloperConnectSource(value *VertexAiReasoningEngineSpecSourceCodeSpecDeveloperConnectSource)
	PutImageSpec(value *VertexAiReasoningEngineSpecSourceCodeSpecImageSpec)
	PutInlineSource(value *VertexAiReasoningEngineSpecSourceCodeSpecInlineSource)
	PutPythonSpec(value *VertexAiReasoningEngineSpecSourceCodeSpecPythonSpec)
	ResetDeveloperConnectSource()
	ResetImageSpec()
	ResetInlineSource()
	ResetPythonSpec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiReasoningEngineSpecSourceCodeSpecOutputReference
type jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) DeveloperConnectSource() VertexAiReasoningEngineSpecSourceCodeSpecDeveloperConnectSourceOutputReference {
	var returns VertexAiReasoningEngineSpecSourceCodeSpecDeveloperConnectSourceOutputReference
	_jsii_.Get(
		j,
		"developerConnectSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) DeveloperConnectSourceInput() *VertexAiReasoningEngineSpecSourceCodeSpecDeveloperConnectSource {
	var returns *VertexAiReasoningEngineSpecSourceCodeSpecDeveloperConnectSource
	_jsii_.Get(
		j,
		"developerConnectSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) ImageSpec() VertexAiReasoningEngineSpecSourceCodeSpecImageSpecOutputReference {
	var returns VertexAiReasoningEngineSpecSourceCodeSpecImageSpecOutputReference
	_jsii_.Get(
		j,
		"imageSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) ImageSpecInput() *VertexAiReasoningEngineSpecSourceCodeSpecImageSpec {
	var returns *VertexAiReasoningEngineSpecSourceCodeSpecImageSpec
	_jsii_.Get(
		j,
		"imageSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) InlineSource() VertexAiReasoningEngineSpecSourceCodeSpecInlineSourceOutputReference {
	var returns VertexAiReasoningEngineSpecSourceCodeSpecInlineSourceOutputReference
	_jsii_.Get(
		j,
		"inlineSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) InlineSourceInput() *VertexAiReasoningEngineSpecSourceCodeSpecInlineSource {
	var returns *VertexAiReasoningEngineSpecSourceCodeSpecInlineSource
	_jsii_.Get(
		j,
		"inlineSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) InternalValue() *VertexAiReasoningEngineSpecSourceCodeSpec {
	var returns *VertexAiReasoningEngineSpecSourceCodeSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) PythonSpec() VertexAiReasoningEngineSpecSourceCodeSpecPythonSpecOutputReference {
	var returns VertexAiReasoningEngineSpecSourceCodeSpecPythonSpecOutputReference
	_jsii_.Get(
		j,
		"pythonSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) PythonSpecInput() *VertexAiReasoningEngineSpecSourceCodeSpecPythonSpec {
	var returns *VertexAiReasoningEngineSpecSourceCodeSpecPythonSpec
	_jsii_.Get(
		j,
		"pythonSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewVertexAiReasoningEngineSpecSourceCodeSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VertexAiReasoningEngineSpecSourceCodeSpecOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiReasoningEngineSpecSourceCodeSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiReasoningEngine.VertexAiReasoningEngineSpecSourceCodeSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiReasoningEngineSpecSourceCodeSpecOutputReference_Override(v VertexAiReasoningEngineSpecSourceCodeSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiReasoningEngine.VertexAiReasoningEngineSpecSourceCodeSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) SetInternalValue(val *VertexAiReasoningEngineSpecSourceCodeSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) PutDeveloperConnectSource(value *VertexAiReasoningEngineSpecSourceCodeSpecDeveloperConnectSource) {
	if err := v.validatePutDeveloperConnectSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putDeveloperConnectSource",
		[]any{value},
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) PutImageSpec(value *VertexAiReasoningEngineSpecSourceCodeSpecImageSpec) {
	if err := v.validatePutImageSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putImageSpec",
		[]any{value},
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) PutInlineSource(value *VertexAiReasoningEngineSpecSourceCodeSpecInlineSource) {
	if err := v.validatePutInlineSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putInlineSource",
		[]any{value},
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) PutPythonSpec(value *VertexAiReasoningEngineSpecSourceCodeSpecPythonSpec) {
	if err := v.validatePutPythonSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPythonSpec",
		[]any{value},
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) ResetDeveloperConnectSource() {
	_jsii_.InvokeVoid(
		v,
		"resetDeveloperConnectSource",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) ResetImageSpec() {
	_jsii_.InvokeVoid(
		v,
		"resetImageSpec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) ResetInlineSource() {
	_jsii_.InvokeVoid(
		v,
		"resetInlineSource",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) ResetPythonSpec() {
	_jsii_.InvokeVoid(
		v,
		"resetPythonSpec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		v,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecSourceCodeSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

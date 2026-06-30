package pubsubtopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/pubsubtopic/internal"
)

type PubsubTopicMessageTransformsOutputReference interface {
	cdktf.ComplexObject
	AiInference() PubsubTopicMessageTransformsAiInferenceOutputReference
	AiInferenceInput() *PubsubTopicMessageTransformsAiInference
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
	Disabled() any
	SetDisabled(val any)
	DisabledInput() any
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	JavascriptUdf() PubsubTopicMessageTransformsJavascriptUdfOutputReference
	JavascriptUdfInput() *PubsubTopicMessageTransformsJavascriptUdf
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
	PutAiInference(value *PubsubTopicMessageTransformsAiInference)
	PutJavascriptUdf(value *PubsubTopicMessageTransformsJavascriptUdf)
	ResetAiInference()
	ResetDisabled()
	ResetJavascriptUdf()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PubsubTopicMessageTransformsOutputReference
type jsiiProxy_PubsubTopicMessageTransformsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PubsubTopicMessageTransformsOutputReference) AiInference() PubsubTopicMessageTransformsAiInferenceOutputReference {
	var returns PubsubTopicMessageTransformsAiInferenceOutputReference
	_jsii_.Get(
		j,
		"aiInference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicMessageTransformsOutputReference) AiInferenceInput() *PubsubTopicMessageTransformsAiInference {
	var returns *PubsubTopicMessageTransformsAiInference
	_jsii_.Get(
		j,
		"aiInferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicMessageTransformsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicMessageTransformsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicMessageTransformsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicMessageTransformsOutputReference) Disabled() any {
	var returns any
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicMessageTransformsOutputReference) DisabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicMessageTransformsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicMessageTransformsOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicMessageTransformsOutputReference) JavascriptUdf() PubsubTopicMessageTransformsJavascriptUdfOutputReference {
	var returns PubsubTopicMessageTransformsJavascriptUdfOutputReference
	_jsii_.Get(
		j,
		"javascriptUdf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicMessageTransformsOutputReference) JavascriptUdfInput() *PubsubTopicMessageTransformsJavascriptUdf {
	var returns *PubsubTopicMessageTransformsJavascriptUdf
	_jsii_.Get(
		j,
		"javascriptUdfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicMessageTransformsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicMessageTransformsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewPubsubTopicMessageTransformsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PubsubTopicMessageTransformsOutputReference {
	_init_.Initialize()

	if err := validateNewPubsubTopicMessageTransformsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PubsubTopicMessageTransformsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.pubsubTopic.PubsubTopicMessageTransformsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPubsubTopicMessageTransformsOutputReference_Override(p PubsubTopicMessageTransformsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.pubsubTopic.PubsubTopicMessageTransformsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PubsubTopicMessageTransformsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicMessageTransformsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicMessageTransformsOutputReference) SetDisabled(val any) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicMessageTransformsOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicMessageTransformsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicMessageTransformsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PubsubTopicMessageTransformsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubTopicMessageTransformsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubTopicMessageTransformsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubTopicMessageTransformsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubTopicMessageTransformsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubTopicMessageTransformsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubTopicMessageTransformsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubTopicMessageTransformsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubTopicMessageTransformsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubTopicMessageTransformsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubTopicMessageTransformsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubTopicMessageTransformsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubTopicMessageTransformsOutputReference) PutAiInference(value *PubsubTopicMessageTransformsAiInference) {
	if err := p.validatePutAiInferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAiInference",
		[]any{value},
	)
}

func (p *jsiiProxy_PubsubTopicMessageTransformsOutputReference) PutJavascriptUdf(value *PubsubTopicMessageTransformsJavascriptUdf) {
	if err := p.validatePutJavascriptUdfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putJavascriptUdf",
		[]any{value},
	)
}

func (p *jsiiProxy_PubsubTopicMessageTransformsOutputReference) ResetAiInference() {
	_jsii_.InvokeVoid(
		p,
		"resetAiInference",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicMessageTransformsOutputReference) ResetDisabled() {
	_jsii_.InvokeVoid(
		p,
		"resetDisabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicMessageTransformsOutputReference) ResetJavascriptUdf() {
	_jsii_.InvokeVoid(
		p,
		"resetJavascriptUdf",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicMessageTransformsOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		p,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubTopicMessageTransformsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

package wafregionalxssmatchset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/wafregionalxssmatchset/internal"
)

type WafregionalXssMatchSetXssMatchTupleOutputReference interface {
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
	FieldToMatch() WafregionalXssMatchSetXssMatchTupleFieldToMatchOutputReference
	FieldToMatchInput() *WafregionalXssMatchSetXssMatchTupleFieldToMatch
	// Experimental.
	Fqn() *string
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
	TextTransformation() *string
	SetTextTransformation(val *string)
	TextTransformationInput() *string
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
	PutFieldToMatch(value *WafregionalXssMatchSetXssMatchTupleFieldToMatch)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WafregionalXssMatchSetXssMatchTupleOutputReference
type jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) FieldToMatch() WafregionalXssMatchSetXssMatchTupleFieldToMatchOutputReference {
	var returns WafregionalXssMatchSetXssMatchTupleFieldToMatchOutputReference
	_jsii_.Get(
		j,
		"fieldToMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) FieldToMatchInput() *WafregionalXssMatchSetXssMatchTupleFieldToMatch {
	var returns *WafregionalXssMatchSetXssMatchTupleFieldToMatch
	_jsii_.Get(
		j,
		"fieldToMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) TextTransformation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textTransformation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) TextTransformationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textTransformationInput",
		&returns,
	)
	return returns
}

func NewWafregionalXssMatchSetXssMatchTupleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WafregionalXssMatchSetXssMatchTupleOutputReference {
	_init_.Initialize()

	if err := validateNewWafregionalXssMatchSetXssMatchTupleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafregionalXssMatchSet.WafregionalXssMatchSetXssMatchTupleOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWafregionalXssMatchSetXssMatchTupleOutputReference_Override(w WafregionalXssMatchSetXssMatchTupleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafregionalXssMatchSet.WafregionalXssMatchSetXssMatchTupleOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) SetTextTransformation(val *string) {
	if err := j.validateSetTextTransformationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textTransformation",
		val,
	)
}

func (w *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) PutFieldToMatch(value *WafregionalXssMatchSetXssMatchTupleFieldToMatch) {
	if err := w.validatePutFieldToMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putFieldToMatch",
		[]any{value},
	)
}

func (w *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		w,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafregionalXssMatchSetXssMatchTupleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

package contactcenterinsightsqaquestion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/contactcenterinsightsqaquestion/internal"
)

type ContactCenterInsightsQaQuestionAnswerChoicesOutputReference interface {
	cdktf.ComplexObject
	BoolValue() any
	SetBoolValue(val any)
	BoolValueInput() any
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
	InternalValue() any
	SetInternalValue(val any)
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	NaValue() any
	SetNaValue(val any)
	NaValueInput() any
	NumValue() *float64
	SetNumValue(val *float64)
	NumValueInput() *float64
	Score() *float64
	SetScore(val *float64)
	ScoreInput() *float64
	StrValue() *string
	SetStrValue(val *string)
	StrValueInput() *string
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
	ResetBoolValue()
	ResetKey()
	ResetNaValue()
	ResetNumValue()
	ResetScore()
	ResetStrValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContactCenterInsightsQaQuestionAnswerChoicesOutputReference
type jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) BoolValue() any {
	var returns any
	_jsii_.Get(
		j,
		"boolValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) BoolValueInput() any {
	var returns any
	_jsii_.Get(
		j,
		"boolValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) NaValue() any {
	var returns any
	_jsii_.Get(
		j,
		"naValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) NaValueInput() any {
	var returns any
	_jsii_.Get(
		j,
		"naValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) NumValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) NumValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) Score() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"score",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ScoreInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) StrValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) StrValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewContactCenterInsightsQaQuestionAnswerChoicesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ContactCenterInsightsQaQuestionAnswerChoicesOutputReference {
	_init_.Initialize()

	if err := validateNewContactCenterInsightsQaQuestionAnswerChoicesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.contactCenterInsightsQaQuestion.ContactCenterInsightsQaQuestionAnswerChoicesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewContactCenterInsightsQaQuestionAnswerChoicesOutputReference_Override(c ContactCenterInsightsQaQuestionAnswerChoicesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.contactCenterInsightsQaQuestion.ContactCenterInsightsQaQuestionAnswerChoicesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) SetBoolValue(val any) {
	if err := j.validateSetBoolValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boolValue",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) SetNaValue(val any) {
	if err := j.validateSetNaValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"naValue",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) SetNumValue(val *float64) {
	if err := j.validateSetNumValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numValue",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) SetScore(val *float64) {
	if err := j.validateSetScoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"score",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) SetStrValue(val *string) {
	if err := j.validateSetStrValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strValue",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ResetBoolValue() {
	_jsii_.InvokeVoid(
		c,
		"resetBoolValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		c,
		"resetKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ResetNaValue() {
	_jsii_.InvokeVoid(
		c,
		"resetNaValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ResetNumValue() {
	_jsii_.InvokeVoid(
		c,
		"resetNumValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ResetScore() {
	_jsii_.InvokeVoid(
		c,
		"resetScore",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ResetStrValue() {
	_jsii_.InvokeVoid(
		c,
		"resetStrValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

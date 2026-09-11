package contactcenterinsightsqaquestion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/contactcenterinsightsqaquestion/internal"
)

type ContactCenterInsightsQaQuestionAnswerChoicesOutputReference interface {
	cdktn.ComplexObject
	BoolValue() interface{}
	SetBoolValue(val interface{})
	BoolValueInput() interface{}
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	NaValue() interface{}
	SetNaValue(val interface{})
	NaValueInput() interface{}
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
	ResetBoolValue()
	ResetKey()
	ResetNaValue()
	ResetNumValue()
	ResetScore()
	ResetStrValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContactCenterInsightsQaQuestionAnswerChoicesOutputReference
type jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) BoolValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"boolValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) BoolValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"boolValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) InternalValue() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) NaValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"naValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) NaValueInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContactCenterInsightsQaQuestionAnswerChoicesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ContactCenterInsightsQaQuestionAnswerChoicesOutputReference {
	_init_.Initialize()

	if err := validateNewContactCenterInsightsQaQuestionAnswerChoicesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.contactCenterInsightsQaQuestion.ContactCenterInsightsQaQuestionAnswerChoicesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewContactCenterInsightsQaQuestionAnswerChoicesOutputReference_Override(c ContactCenterInsightsQaQuestionAnswerChoicesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.contactCenterInsightsQaQuestion.ContactCenterInsightsQaQuestionAnswerChoicesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetBoolValue(val interface{}) {
	if err := j.validateSetBoolValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boolValue",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetNaValue(val interface{}) {
	if err := j.validateSetNaValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"naValue",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetNumValue(val *float64) {
	if err := j.validateSetNumValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numValue",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetScore(val *float64) {
	if err := j.validateSetScoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"score",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetStrValue(val *string) {
	if err := j.validateSetStrValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strValue",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
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

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
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

func (c *jsiiProxy_ContactCenterInsightsQaQuestionAnswerChoicesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
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


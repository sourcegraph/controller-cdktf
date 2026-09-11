package googlecontactcenterinsightsqaquestion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontactcenterinsightsqaquestion/internal"
)

type GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference interface {
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

// The jsii proxy struct for GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference
type jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) BoolValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"boolValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) BoolValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"boolValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) NaValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"naValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) NaValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"naValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) NumValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) NumValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) Score() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"score",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ScoreInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) StrValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) StrValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleContactCenterInsightsQaQuestion.GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference_Override(g GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleContactCenterInsightsQaQuestion.GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetBoolValue(val interface{}) {
	if err := j.validateSetBoolValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boolValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetNaValue(val interface{}) {
	if err := j.validateSetNaValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"naValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetNumValue(val *float64) {
	if err := j.validateSetNumValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetScore(val *float64) {
	if err := j.validateSetScoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"score",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetStrValue(val *string) {
	if err := j.validateSetStrValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ResetBoolValue() {
	_jsii_.InvokeVoid(
		g,
		"resetBoolValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		g,
		"resetKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ResetNaValue() {
	_jsii_.InvokeVoid(
		g,
		"resetNaValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ResetNumValue() {
	_jsii_.InvokeVoid(
		g,
		"resetNumValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ResetScore() {
	_jsii_.InvokeVoid(
		g,
		"resetScore",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ResetStrValue() {
	_jsii_.InvokeVoid(
		g,
		"resetStrValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionAnswerChoicesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


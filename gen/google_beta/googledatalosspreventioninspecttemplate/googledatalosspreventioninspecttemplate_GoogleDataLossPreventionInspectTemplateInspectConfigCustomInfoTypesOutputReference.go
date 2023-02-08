package googledatalosspreventioninspecttemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventioninspecttemplate/internal"
)

type GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference interface {
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
	Dictionary() GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionaryOutputReference
	DictionaryInput() *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionary
	ExclusionType() *string
	SetExclusionType(val *string)
	ExclusionTypeInput() *string
	// Experimental.
	Fqn() *string
	InfoType() GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesInfoTypeOutputReference
	InfoTypeInput() *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesInfoType
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Likelihood() *string
	SetLikelihood(val *string)
	LikelihoodInput() *string
	Regex() GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegexOutputReference
	RegexInput() *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegex
	StoredType() GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredTypeOutputReference
	StoredTypeInput() *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredType
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
	PutDictionary(value *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionary)
	PutInfoType(value *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesInfoType)
	PutRegex(value *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegex)
	PutStoredType(value *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredType)
	ResetDictionary()
	ResetExclusionType()
	ResetLikelihood()
	ResetRegex()
	ResetStoredType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference
type jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) Dictionary() GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionaryOutputReference {
	var returns GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionaryOutputReference
	_jsii_.Get(
		j,
		"dictionary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) DictionaryInput() *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionary {
	var returns *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionary
	_jsii_.Get(
		j,
		"dictionaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ExclusionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exclusionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ExclusionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exclusionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) InfoType() GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesInfoTypeOutputReference {
	var returns GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesInfoTypeOutputReference
	_jsii_.Get(
		j,
		"infoType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) InfoTypeInput() *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesInfoType {
	var returns *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesInfoType
	_jsii_.Get(
		j,
		"infoTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) Likelihood() *string {
	var returns *string
	_jsii_.Get(
		j,
		"likelihood",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) LikelihoodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"likelihoodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) Regex() GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegexOutputReference {
	var returns GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegexOutputReference
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) RegexInput() *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegex {
	var returns *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegex
	_jsii_.Get(
		j,
		"regexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) StoredType() GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredTypeOutputReference {
	var returns GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredTypeOutputReference
	_jsii_.Get(
		j,
		"storedType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) StoredTypeInput() *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredType {
	var returns *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredType
	_jsii_.Get(
		j,
		"storedTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataLossPreventionInspectTemplate.GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference_Override(g GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataLossPreventionInspectTemplate.GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference)SetExclusionType(val *string) {
	if err := j.validateSetExclusionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclusionType",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference)SetLikelihood(val *string) {
	if err := j.validateSetLikelihoodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"likelihood",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) PutDictionary(value *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionary) {
	if err := g.validatePutDictionaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDictionary",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) PutInfoType(value *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesInfoType) {
	if err := g.validatePutInfoTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInfoType",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) PutRegex(value *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegex) {
	if err := g.validatePutRegexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRegex",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) PutStoredType(value *GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredType) {
	if err := g.validatePutStoredTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStoredType",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ResetDictionary() {
	_jsii_.InvokeVoid(
		g,
		"resetDictionary",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ResetExclusionType() {
	_jsii_.InvokeVoid(
		g,
		"resetExclusionType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ResetLikelihood() {
	_jsii_.InvokeVoid(
		g,
		"resetLikelihood",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ResetRegex() {
	_jsii_.InvokeVoid(
		g,
		"resetRegex",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ResetStoredType() {
	_jsii_.InvokeVoid(
		g,
		"resetStoredType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


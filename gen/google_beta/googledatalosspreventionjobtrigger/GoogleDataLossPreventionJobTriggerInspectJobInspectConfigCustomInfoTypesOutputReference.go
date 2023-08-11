package googledatalosspreventionjobtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventionjobtrigger/internal"
)

type GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference interface {
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
	Dictionary() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryOutputReference
	DictionaryInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary
	ExclusionType() *string
	SetExclusionType(val *string)
	ExclusionTypeInput() *string
	// Experimental.
	Fqn() *string
	InfoType() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeOutputReference
	InfoTypeInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Likelihood() *string
	SetLikelihood(val *string)
	LikelihoodInput() *string
	Regex() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesRegexOutputReference
	RegexInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesRegex
	StoredType() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeOutputReference
	StoredTypeInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType
	SurrogateType() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeOutputReference
	SurrogateTypeInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType
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
	PutDictionary(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary)
	PutInfoType(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType)
	PutRegex(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesRegex)
	PutStoredType(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType)
	PutSurrogateType(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType)
	ResetDictionary()
	ResetExclusionType()
	ResetLikelihood()
	ResetRegex()
	ResetStoredType()
	ResetSurrogateType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference
type jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) Dictionary() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryOutputReference
	_jsii_.Get(
		j,
		"dictionary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) DictionaryInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary
	_jsii_.Get(
		j,
		"dictionaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) ExclusionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exclusionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) ExclusionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exclusionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) InfoType() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeOutputReference
	_jsii_.Get(
		j,
		"infoType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) InfoTypeInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType
	_jsii_.Get(
		j,
		"infoTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) Likelihood() *string {
	var returns *string
	_jsii_.Get(
		j,
		"likelihood",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) LikelihoodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"likelihoodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) Regex() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesRegexOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesRegexOutputReference
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) RegexInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesRegex {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesRegex
	_jsii_.Get(
		j,
		"regexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) StoredType() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeOutputReference
	_jsii_.Get(
		j,
		"storedType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) StoredTypeInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType
	_jsii_.Get(
		j,
		"storedTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) SurrogateType() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeOutputReference
	_jsii_.Get(
		j,
		"surrogateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) SurrogateTypeInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType
	_jsii_.Get(
		j,
		"surrogateTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataLossPreventionJobTrigger.GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference_Override(g GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataLossPreventionJobTrigger.GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference)SetExclusionType(val *string) {
	if err := j.validateSetExclusionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclusionType",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference)SetLikelihood(val *string) {
	if err := j.validateSetLikelihoodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"likelihood",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) PutDictionary(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary) {
	if err := g.validatePutDictionaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDictionary",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) PutInfoType(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType) {
	if err := g.validatePutInfoTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInfoType",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) PutRegex(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesRegex) {
	if err := g.validatePutRegexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRegex",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) PutStoredType(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType) {
	if err := g.validatePutStoredTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStoredType",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) PutSurrogateType(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType) {
	if err := g.validatePutSurrogateTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSurrogateType",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) ResetDictionary() {
	_jsii_.InvokeVoid(
		g,
		"resetDictionary",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) ResetExclusionType() {
	_jsii_.InvokeVoid(
		g,
		"resetExclusionType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) ResetLikelihood() {
	_jsii_.InvokeVoid(
		g,
		"resetLikelihood",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) ResetRegex() {
	_jsii_.InvokeVoid(
		g,
		"resetRegex",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) ResetStoredType() {
	_jsii_.InvokeVoid(
		g,
		"resetStoredType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) ResetSurrogateType() {
	_jsii_.InvokeVoid(
		g,
		"resetSurrogateType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package googledatalosspreventionjobtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventionjobtrigger/internal"
)

type GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference interface {
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
	Dictionary() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference
	DictionaryInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary
	ExcludeByHotword() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordOutputReference
	ExcludeByHotwordInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotword
	ExcludeInfoTypes() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesOutputReference
	ExcludeInfoTypesInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule
	SetInternalValue(val *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule)
	MatchingType() *string
	SetMatchingType(val *string)
	MatchingTypeInput() *string
	Regex() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexOutputReference
	RegexInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex
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
	PutDictionary(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary)
	PutExcludeByHotword(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotword)
	PutExcludeInfoTypes(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes)
	PutRegex(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex)
	ResetDictionary()
	ResetExcludeByHotword()
	ResetExcludeInfoTypes()
	ResetRegex()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference
type jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) Dictionary() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference
	_jsii_.Get(
		j,
		"dictionary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) DictionaryInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary
	_jsii_.Get(
		j,
		"dictionaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ExcludeByHotword() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordOutputReference
	_jsii_.Get(
		j,
		"excludeByHotword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ExcludeByHotwordInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotword {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotword
	_jsii_.Get(
		j,
		"excludeByHotwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ExcludeInfoTypes() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesOutputReference
	_jsii_.Get(
		j,
		"excludeInfoTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ExcludeInfoTypesInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes
	_jsii_.Get(
		j,
		"excludeInfoTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) InternalValue() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) MatchingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) MatchingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) Regex() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexOutputReference
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) RegexInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex
	_jsii_.Get(
		j,
		"regexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataLossPreventionJobTrigger.GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference_Override(g GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataLossPreventionJobTrigger.GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference)SetInternalValue(val *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference)SetMatchingType(val *string) {
	if err := j.validateSetMatchingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchingType",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) PutDictionary(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary) {
	if err := g.validatePutDictionaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDictionary",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) PutExcludeByHotword(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotword) {
	if err := g.validatePutExcludeByHotwordParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExcludeByHotword",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) PutExcludeInfoTypes(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) {
	if err := g.validatePutExcludeInfoTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExcludeInfoTypes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) PutRegex(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex) {
	if err := g.validatePutRegexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRegex",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ResetDictionary() {
	_jsii_.InvokeVoid(
		g,
		"resetDictionary",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ResetExcludeByHotword() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludeByHotword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ResetExcludeInfoTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludeInfoTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ResetRegex() {
	_jsii_.InvokeVoid(
		g,
		"resetRegex",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


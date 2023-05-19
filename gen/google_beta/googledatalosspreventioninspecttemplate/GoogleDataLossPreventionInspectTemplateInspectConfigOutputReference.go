package googledatalosspreventioninspecttemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventioninspecttemplate/internal"
)

type GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference interface {
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
	ContentOptions() *[]*string
	SetContentOptions(val *[]*string)
	ContentOptionsInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomInfoTypes() GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesList
	CustomInfoTypesInput() interface{}
	ExcludeInfoTypes() interface{}
	SetExcludeInfoTypes(val interface{})
	ExcludeInfoTypesInput() interface{}
	// Experimental.
	Fqn() *string
	IncludeQuote() interface{}
	SetIncludeQuote(val interface{})
	IncludeQuoteInput() interface{}
	InfoTypes() GoogleDataLossPreventionInspectTemplateInspectConfigInfoTypesList
	InfoTypesInput() interface{}
	InternalValue() *GoogleDataLossPreventionInspectTemplateInspectConfig
	SetInternalValue(val *GoogleDataLossPreventionInspectTemplateInspectConfig)
	Limits() GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference
	LimitsInput() *GoogleDataLossPreventionInspectTemplateInspectConfigLimits
	MinLikelihood() *string
	SetMinLikelihood(val *string)
	MinLikelihoodInput() *string
	RuleSet() GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetList
	RuleSetInput() interface{}
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
	PutCustomInfoTypes(value interface{})
	PutInfoTypes(value interface{})
	PutLimits(value *GoogleDataLossPreventionInspectTemplateInspectConfigLimits)
	PutRuleSet(value interface{})
	ResetContentOptions()
	ResetCustomInfoTypes()
	ResetExcludeInfoTypes()
	ResetIncludeQuote()
	ResetInfoTypes()
	ResetLimits()
	ResetMinLikelihood()
	ResetRuleSet()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference
type jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) ContentOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"contentOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) ContentOptionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"contentOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) CustomInfoTypes() GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesList {
	var returns GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesList
	_jsii_.Get(
		j,
		"customInfoTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) CustomInfoTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customInfoTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) ExcludeInfoTypes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeInfoTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) ExcludeInfoTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeInfoTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) IncludeQuote() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeQuote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) IncludeQuoteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeQuoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) InfoTypes() GoogleDataLossPreventionInspectTemplateInspectConfigInfoTypesList {
	var returns GoogleDataLossPreventionInspectTemplateInspectConfigInfoTypesList
	_jsii_.Get(
		j,
		"infoTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) InfoTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"infoTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) InternalValue() *GoogleDataLossPreventionInspectTemplateInspectConfig {
	var returns *GoogleDataLossPreventionInspectTemplateInspectConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) Limits() GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference {
	var returns GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference
	_jsii_.Get(
		j,
		"limits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) LimitsInput() *GoogleDataLossPreventionInspectTemplateInspectConfigLimits {
	var returns *GoogleDataLossPreventionInspectTemplateInspectConfigLimits
	_jsii_.Get(
		j,
		"limitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) MinLikelihood() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minLikelihood",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) MinLikelihoodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minLikelihoodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) RuleSet() GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetList {
	var returns GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetList
	_jsii_.Get(
		j,
		"ruleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) RuleSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionInspectTemplateInspectConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionInspectTemplateInspectConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataLossPreventionInspectTemplate.GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionInspectTemplateInspectConfigOutputReference_Override(g GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataLossPreventionInspectTemplate.GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference)SetContentOptions(val *[]*string) {
	if err := j.validateSetContentOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentOptions",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference)SetExcludeInfoTypes(val interface{}) {
	if err := j.validateSetExcludeInfoTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeInfoTypes",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference)SetIncludeQuote(val interface{}) {
	if err := j.validateSetIncludeQuoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeQuote",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference)SetInternalValue(val *GoogleDataLossPreventionInspectTemplateInspectConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference)SetMinLikelihood(val *string) {
	if err := j.validateSetMinLikelihoodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minLikelihood",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) PutCustomInfoTypes(value interface{}) {
	if err := g.validatePutCustomInfoTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCustomInfoTypes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) PutInfoTypes(value interface{}) {
	if err := g.validatePutInfoTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInfoTypes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) PutLimits(value *GoogleDataLossPreventionInspectTemplateInspectConfigLimits) {
	if err := g.validatePutLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLimits",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) PutRuleSet(value interface{}) {
	if err := g.validatePutRuleSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRuleSet",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) ResetContentOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetContentOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) ResetCustomInfoTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomInfoTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) ResetExcludeInfoTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludeInfoTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) ResetIncludeQuote() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeQuote",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) ResetInfoTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetInfoTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) ResetLimits() {
	_jsii_.InvokeVoid(
		g,
		"resetLimits",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) ResetMinLikelihood() {
	_jsii_.InvokeVoid(
		g,
		"resetMinLikelihood",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) ResetRuleSet() {
	_jsii_.InvokeVoid(
		g,
		"resetRuleSet",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


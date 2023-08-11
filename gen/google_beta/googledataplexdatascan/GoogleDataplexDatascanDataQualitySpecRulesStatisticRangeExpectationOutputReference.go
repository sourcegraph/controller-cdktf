package googledataplexdatascan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataplexdatascan/internal"
)

type GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectation
	SetInternalValue(val *GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectation)
	MaxValue() *string
	SetMaxValue(val *string)
	MaxValueInput() *string
	MinValue() *string
	SetMinValue(val *string)
	MinValueInput() *string
	Statistic() *string
	SetStatistic(val *string)
	StatisticInput() *string
	StrictMaxEnabled() interface{}
	SetStrictMaxEnabled(val interface{})
	StrictMaxEnabledInput() interface{}
	StrictMinEnabled() interface{}
	SetStrictMinEnabled(val interface{})
	StrictMinEnabledInput() interface{}
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
	ResetMaxValue()
	ResetMinValue()
	ResetStrictMaxEnabled()
	ResetStrictMinEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference
type jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) InternalValue() *GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectation {
	var returns *GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) MaxValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) MaxValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) MinValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) MinValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) Statistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) StatisticInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statisticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) StrictMaxEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictMaxEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) StrictMaxEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictMaxEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) StrictMinEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictMinEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) StrictMinEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictMinEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataplexDatascan.GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference_Override(g GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataplexDatascan.GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference)SetInternalValue(val *GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference)SetMaxValue(val *string) {
	if err := j.validateSetMaxValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference)SetMinValue(val *string) {
	if err := j.validateSetMinValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference)SetStatistic(val *string) {
	if err := j.validateSetStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statistic",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference)SetStrictMaxEnabled(val interface{}) {
	if err := j.validateSetStrictMaxEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strictMaxEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference)SetStrictMinEnabled(val interface{}) {
	if err := j.validateSetStrictMinEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strictMinEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) ResetMaxValue() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) ResetMinValue() {
	_jsii_.InvokeVoid(
		g,
		"resetMinValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) ResetStrictMaxEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetStrictMaxEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) ResetStrictMinEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetStrictMinEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


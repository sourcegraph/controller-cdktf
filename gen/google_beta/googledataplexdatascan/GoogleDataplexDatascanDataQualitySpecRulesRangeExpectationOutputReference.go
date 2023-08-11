package googledataplexdatascan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataplexdatascan/internal"
)

type GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference interface {
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
	InternalValue() *GoogleDataplexDatascanDataQualitySpecRulesRangeExpectation
	SetInternalValue(val *GoogleDataplexDatascanDataQualitySpecRulesRangeExpectation)
	MaxValue() *string
	SetMaxValue(val *string)
	MaxValueInput() *string
	MinValue() *string
	SetMinValue(val *string)
	MinValueInput() *string
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

// The jsii proxy struct for GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference
type jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) InternalValue() *GoogleDataplexDatascanDataQualitySpecRulesRangeExpectation {
	var returns *GoogleDataplexDatascanDataQualitySpecRulesRangeExpectation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) MaxValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) MaxValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) MinValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) MinValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) StrictMaxEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictMaxEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) StrictMaxEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictMaxEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) StrictMinEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictMinEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) StrictMinEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictMinEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataplexDatascan.GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference_Override(g GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataplexDatascan.GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference)SetInternalValue(val *GoogleDataplexDatascanDataQualitySpecRulesRangeExpectation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference)SetMaxValue(val *string) {
	if err := j.validateSetMaxValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference)SetMinValue(val *string) {
	if err := j.validateSetMinValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference)SetStrictMaxEnabled(val interface{}) {
	if err := j.validateSetStrictMaxEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strictMaxEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference)SetStrictMinEnabled(val interface{}) {
	if err := j.validateSetStrictMinEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strictMinEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) ResetMaxValue() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) ResetMinValue() {
	_jsii_.InvokeVoid(
		g,
		"resetMinValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) ResetStrictMaxEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetStrictMaxEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) ResetStrictMinEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetStrictMinEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


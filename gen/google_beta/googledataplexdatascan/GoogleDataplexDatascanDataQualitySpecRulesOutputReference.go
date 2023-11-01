package googledataplexdatascan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataplexdatascan/internal"
)

type GoogleDataplexDatascanDataQualitySpecRulesOutputReference interface {
	cdktf.ComplexObject
	Column() *string
	SetColumn(val *string)
	ColumnInput() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Dimension() *string
	SetDimension(val *string)
	DimensionInput() *string
	// Experimental.
	Fqn() *string
	IgnoreNull() interface{}
	SetIgnoreNull(val interface{})
	IgnoreNullInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	NonNullExpectation() GoogleDataplexDatascanDataQualitySpecRulesNonNullExpectationOutputReference
	NonNullExpectationInput() *GoogleDataplexDatascanDataQualitySpecRulesNonNullExpectation
	RangeExpectation() GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference
	RangeExpectationInput() *GoogleDataplexDatascanDataQualitySpecRulesRangeExpectation
	RegexExpectation() GoogleDataplexDatascanDataQualitySpecRulesRegexExpectationOutputReference
	RegexExpectationInput() *GoogleDataplexDatascanDataQualitySpecRulesRegexExpectation
	RowConditionExpectation() GoogleDataplexDatascanDataQualitySpecRulesRowConditionExpectationOutputReference
	RowConditionExpectationInput() *GoogleDataplexDatascanDataQualitySpecRulesRowConditionExpectation
	SetExpectation() GoogleDataplexDatascanDataQualitySpecRulesSetExpectationOutputReference
	SetExpectationInput() *GoogleDataplexDatascanDataQualitySpecRulesSetExpectation
	StatisticRangeExpectation() GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference
	StatisticRangeExpectationInput() *GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectation
	TableConditionExpectation() GoogleDataplexDatascanDataQualitySpecRulesTableConditionExpectationOutputReference
	TableConditionExpectationInput() *GoogleDataplexDatascanDataQualitySpecRulesTableConditionExpectation
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Threshold() *float64
	SetThreshold(val *float64)
	ThresholdInput() *float64
	UniquenessExpectation() GoogleDataplexDatascanDataQualitySpecRulesUniquenessExpectationOutputReference
	UniquenessExpectationInput() *GoogleDataplexDatascanDataQualitySpecRulesUniquenessExpectation
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
	PutNonNullExpectation(value *GoogleDataplexDatascanDataQualitySpecRulesNonNullExpectation)
	PutRangeExpectation(value *GoogleDataplexDatascanDataQualitySpecRulesRangeExpectation)
	PutRegexExpectation(value *GoogleDataplexDatascanDataQualitySpecRulesRegexExpectation)
	PutRowConditionExpectation(value *GoogleDataplexDatascanDataQualitySpecRulesRowConditionExpectation)
	PutSetExpectation(value *GoogleDataplexDatascanDataQualitySpecRulesSetExpectation)
	PutStatisticRangeExpectation(value *GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectation)
	PutTableConditionExpectation(value *GoogleDataplexDatascanDataQualitySpecRulesTableConditionExpectation)
	PutUniquenessExpectation(value *GoogleDataplexDatascanDataQualitySpecRulesUniquenessExpectation)
	ResetColumn()
	ResetDescription()
	ResetIgnoreNull()
	ResetName()
	ResetNonNullExpectation()
	ResetRangeExpectation()
	ResetRegexExpectation()
	ResetRowConditionExpectation()
	ResetSetExpectation()
	ResetStatisticRangeExpectation()
	ResetTableConditionExpectation()
	ResetThreshold()
	ResetUniquenessExpectation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataplexDatascanDataQualitySpecRulesOutputReference
type jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) Column() *string {
	var returns *string
	_jsii_.Get(
		j,
		"column",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) ColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) Dimension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) DimensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) IgnoreNull() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreNull",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) IgnoreNullInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreNullInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) NonNullExpectation() GoogleDataplexDatascanDataQualitySpecRulesNonNullExpectationOutputReference {
	var returns GoogleDataplexDatascanDataQualitySpecRulesNonNullExpectationOutputReference
	_jsii_.Get(
		j,
		"nonNullExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) NonNullExpectationInput() *GoogleDataplexDatascanDataQualitySpecRulesNonNullExpectation {
	var returns *GoogleDataplexDatascanDataQualitySpecRulesNonNullExpectation
	_jsii_.Get(
		j,
		"nonNullExpectationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) RangeExpectation() GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference {
	var returns GoogleDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference
	_jsii_.Get(
		j,
		"rangeExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) RangeExpectationInput() *GoogleDataplexDatascanDataQualitySpecRulesRangeExpectation {
	var returns *GoogleDataplexDatascanDataQualitySpecRulesRangeExpectation
	_jsii_.Get(
		j,
		"rangeExpectationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) RegexExpectation() GoogleDataplexDatascanDataQualitySpecRulesRegexExpectationOutputReference {
	var returns GoogleDataplexDatascanDataQualitySpecRulesRegexExpectationOutputReference
	_jsii_.Get(
		j,
		"regexExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) RegexExpectationInput() *GoogleDataplexDatascanDataQualitySpecRulesRegexExpectation {
	var returns *GoogleDataplexDatascanDataQualitySpecRulesRegexExpectation
	_jsii_.Get(
		j,
		"regexExpectationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) RowConditionExpectation() GoogleDataplexDatascanDataQualitySpecRulesRowConditionExpectationOutputReference {
	var returns GoogleDataplexDatascanDataQualitySpecRulesRowConditionExpectationOutputReference
	_jsii_.Get(
		j,
		"rowConditionExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) RowConditionExpectationInput() *GoogleDataplexDatascanDataQualitySpecRulesRowConditionExpectation {
	var returns *GoogleDataplexDatascanDataQualitySpecRulesRowConditionExpectation
	_jsii_.Get(
		j,
		"rowConditionExpectationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) SetExpectation() GoogleDataplexDatascanDataQualitySpecRulesSetExpectationOutputReference {
	var returns GoogleDataplexDatascanDataQualitySpecRulesSetExpectationOutputReference
	_jsii_.Get(
		j,
		"setExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) SetExpectationInput() *GoogleDataplexDatascanDataQualitySpecRulesSetExpectation {
	var returns *GoogleDataplexDatascanDataQualitySpecRulesSetExpectation
	_jsii_.Get(
		j,
		"setExpectationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) StatisticRangeExpectation() GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference {
	var returns GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference
	_jsii_.Get(
		j,
		"statisticRangeExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) StatisticRangeExpectationInput() *GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectation {
	var returns *GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectation
	_jsii_.Get(
		j,
		"statisticRangeExpectationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) TableConditionExpectation() GoogleDataplexDatascanDataQualitySpecRulesTableConditionExpectationOutputReference {
	var returns GoogleDataplexDatascanDataQualitySpecRulesTableConditionExpectationOutputReference
	_jsii_.Get(
		j,
		"tableConditionExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) TableConditionExpectationInput() *GoogleDataplexDatascanDataQualitySpecRulesTableConditionExpectation {
	var returns *GoogleDataplexDatascanDataQualitySpecRulesTableConditionExpectation
	_jsii_.Get(
		j,
		"tableConditionExpectationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) Threshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) ThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) UniquenessExpectation() GoogleDataplexDatascanDataQualitySpecRulesUniquenessExpectationOutputReference {
	var returns GoogleDataplexDatascanDataQualitySpecRulesUniquenessExpectationOutputReference
	_jsii_.Get(
		j,
		"uniquenessExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) UniquenessExpectationInput() *GoogleDataplexDatascanDataQualitySpecRulesUniquenessExpectation {
	var returns *GoogleDataplexDatascanDataQualitySpecRulesUniquenessExpectation
	_jsii_.Get(
		j,
		"uniquenessExpectationInput",
		&returns,
	)
	return returns
}


func NewGoogleDataplexDatascanDataQualitySpecRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleDataplexDatascanDataQualitySpecRulesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataplexDatascanDataQualitySpecRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataplexDatascan.GoogleDataplexDatascanDataQualitySpecRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleDataplexDatascanDataQualitySpecRulesOutputReference_Override(g GoogleDataplexDatascanDataQualitySpecRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataplexDatascan.GoogleDataplexDatascanDataQualitySpecRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference)SetColumn(val *string) {
	if err := j.validateSetColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"column",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference)SetDimension(val *string) {
	if err := j.validateSetDimensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimension",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference)SetIgnoreNull(val interface{}) {
	if err := j.validateSetIgnoreNullParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreNull",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference)SetThreshold(val *float64) {
	if err := j.validateSetThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) PutNonNullExpectation(value *GoogleDataplexDatascanDataQualitySpecRulesNonNullExpectation) {
	if err := g.validatePutNonNullExpectationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNonNullExpectation",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) PutRangeExpectation(value *GoogleDataplexDatascanDataQualitySpecRulesRangeExpectation) {
	if err := g.validatePutRangeExpectationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRangeExpectation",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) PutRegexExpectation(value *GoogleDataplexDatascanDataQualitySpecRulesRegexExpectation) {
	if err := g.validatePutRegexExpectationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRegexExpectation",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) PutRowConditionExpectation(value *GoogleDataplexDatascanDataQualitySpecRulesRowConditionExpectation) {
	if err := g.validatePutRowConditionExpectationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRowConditionExpectation",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) PutSetExpectation(value *GoogleDataplexDatascanDataQualitySpecRulesSetExpectation) {
	if err := g.validatePutSetExpectationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSetExpectation",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) PutStatisticRangeExpectation(value *GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectation) {
	if err := g.validatePutStatisticRangeExpectationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStatisticRangeExpectation",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) PutTableConditionExpectation(value *GoogleDataplexDatascanDataQualitySpecRulesTableConditionExpectation) {
	if err := g.validatePutTableConditionExpectationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTableConditionExpectation",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) PutUniquenessExpectation(value *GoogleDataplexDatascanDataQualitySpecRulesUniquenessExpectation) {
	if err := g.validatePutUniquenessExpectationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUniquenessExpectation",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) ResetColumn() {
	_jsii_.InvokeVoid(
		g,
		"resetColumn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) ResetIgnoreNull() {
	_jsii_.InvokeVoid(
		g,
		"resetIgnoreNull",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) ResetNonNullExpectation() {
	_jsii_.InvokeVoid(
		g,
		"resetNonNullExpectation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) ResetRangeExpectation() {
	_jsii_.InvokeVoid(
		g,
		"resetRangeExpectation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) ResetRegexExpectation() {
	_jsii_.InvokeVoid(
		g,
		"resetRegexExpectation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) ResetRowConditionExpectation() {
	_jsii_.InvokeVoid(
		g,
		"resetRowConditionExpectation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) ResetSetExpectation() {
	_jsii_.InvokeVoid(
		g,
		"resetSetExpectation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) ResetStatisticRangeExpectation() {
	_jsii_.InvokeVoid(
		g,
		"resetStatisticRangeExpectation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) ResetTableConditionExpectation() {
	_jsii_.InvokeVoid(
		g,
		"resetTableConditionExpectation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) ResetThreshold() {
	_jsii_.InvokeVoid(
		g,
		"resetThreshold",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) ResetUniquenessExpectation() {
	_jsii_.InvokeVoid(
		g,
		"resetUniquenessExpectation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


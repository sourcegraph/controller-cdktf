package contactcenterinsightsassessmentrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/contactcenterinsightsassessmentrule/internal"
)

type ContactCenterInsightsAssessmentRuleSampleRuleOutputReference interface {
	cdktf.ComplexObject
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
	ConversationFilter() *string
	SetConversationFilter(val *string)
	ConversationFilterInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dimension() *string
	SetDimension(val *string)
	DimensionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *ContactCenterInsightsAssessmentRuleSampleRule
	SetInternalValue(val *ContactCenterInsightsAssessmentRuleSampleRule)
	SamplePercentage() *float64
	SetSamplePercentage(val *float64)
	SamplePercentageInput() *float64
	SampleRow() *float64
	SetSampleRow(val *float64)
	SampleRowInput() *float64
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
	ResetConversationFilter()
	ResetDimension()
	ResetSamplePercentage()
	ResetSampleRow()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContactCenterInsightsAssessmentRuleSampleRuleOutputReference
type jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) ConversationFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conversationFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) ConversationFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conversationFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) Dimension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) DimensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) InternalValue() *ContactCenterInsightsAssessmentRuleSampleRule {
	var returns *ContactCenterInsightsAssessmentRuleSampleRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) SamplePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) SamplePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) SampleRow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleRow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) SampleRowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleRowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewContactCenterInsightsAssessmentRuleSampleRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContactCenterInsightsAssessmentRuleSampleRuleOutputReference {
	_init_.Initialize()

	if err := validateNewContactCenterInsightsAssessmentRuleSampleRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.contactCenterInsightsAssessmentRule.ContactCenterInsightsAssessmentRuleSampleRuleOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContactCenterInsightsAssessmentRuleSampleRuleOutputReference_Override(c ContactCenterInsightsAssessmentRuleSampleRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.contactCenterInsightsAssessmentRule.ContactCenterInsightsAssessmentRuleSampleRuleOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) SetConversationFilter(val *string) {
	if err := j.validateSetConversationFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conversationFilter",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) SetDimension(val *string) {
	if err := j.validateSetDimensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimension",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) SetInternalValue(val *ContactCenterInsightsAssessmentRuleSampleRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) SetSamplePercentage(val *float64) {
	if err := j.validateSetSamplePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samplePercentage",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) SetSampleRow(val *float64) {
	if err := j.validateSetSampleRowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleRow",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) ResetConversationFilter() {
	_jsii_.InvokeVoid(
		c,
		"resetConversationFilter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) ResetDimension() {
	_jsii_.InvokeVoid(
		c,
		"resetDimension",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) ResetSamplePercentage() {
	_jsii_.InvokeVoid(
		c,
		"resetSamplePercentage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) ResetSampleRow() {
	_jsii_.InvokeVoid(
		c,
		"resetSampleRow",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

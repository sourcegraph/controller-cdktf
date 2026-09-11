package contactcenterinsightsassessmentrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/contactcenterinsightsassessmentrule/internal"
)

type ContactCenterInsightsAssessmentRuleSampleRuleOutputReference interface {
	cdktn.ComplexObject
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
	ResetConversationFilter()
	ResetDimension()
	ResetSamplePercentage()
	ResetSampleRow()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContactCenterInsightsAssessmentRuleSampleRuleOutputReference
type jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContactCenterInsightsAssessmentRuleSampleRuleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ContactCenterInsightsAssessmentRuleSampleRuleOutputReference {
	_init_.Initialize()

	if err := validateNewContactCenterInsightsAssessmentRuleSampleRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.contactCenterInsightsAssessmentRule.ContactCenterInsightsAssessmentRuleSampleRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContactCenterInsightsAssessmentRuleSampleRuleOutputReference_Override(c ContactCenterInsightsAssessmentRuleSampleRuleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.contactCenterInsightsAssessmentRule.ContactCenterInsightsAssessmentRuleSampleRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference)SetConversationFilter(val *string) {
	if err := j.validateSetConversationFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conversationFilter",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference)SetDimension(val *string) {
	if err := j.validateSetDimensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimension",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference)SetInternalValue(val *ContactCenterInsightsAssessmentRuleSampleRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference)SetSamplePercentage(val *float64) {
	if err := j.validateSetSamplePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samplePercentage",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference)SetSampleRow(val *float64) {
	if err := j.validateSetSampleRowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleRow",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ContactCenterInsightsAssessmentRuleSampleRuleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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


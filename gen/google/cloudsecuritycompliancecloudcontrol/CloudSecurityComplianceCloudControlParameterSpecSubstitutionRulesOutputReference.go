package cloudsecuritycompliancecloudcontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cloudsecuritycompliancecloudcontrol/internal"
)

type CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference interface {
	cdktf.ComplexObject
	AttributeSubstitutionRule() CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesAttributeSubstitutionRuleOutputReference
	AttributeSubstitutionRuleInput() *CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesAttributeSubstitutionRule
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	PlaceholderSubstitutionRule() CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesPlaceholderSubstitutionRuleOutputReference
	PlaceholderSubstitutionRuleInput() *CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesPlaceholderSubstitutionRule
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
	PutAttributeSubstitutionRule(value *CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesAttributeSubstitutionRule)
	PutPlaceholderSubstitutionRule(value *CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesPlaceholderSubstitutionRule)
	ResetAttributeSubstitutionRule()
	ResetPlaceholderSubstitutionRule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference
type jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) AttributeSubstitutionRule() CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesAttributeSubstitutionRuleOutputReference {
	var returns CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesAttributeSubstitutionRuleOutputReference
	_jsii_.Get(
		j,
		"attributeSubstitutionRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) AttributeSubstitutionRuleInput() *CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesAttributeSubstitutionRule {
	var returns *CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesAttributeSubstitutionRule
	_jsii_.Get(
		j,
		"attributeSubstitutionRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) PlaceholderSubstitutionRule() CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesPlaceholderSubstitutionRuleOutputReference {
	var returns CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesPlaceholderSubstitutionRuleOutputReference
	_jsii_.Get(
		j,
		"placeholderSubstitutionRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) PlaceholderSubstitutionRuleInput() *CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesPlaceholderSubstitutionRule {
	var returns *CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesPlaceholderSubstitutionRule
	_jsii_.Get(
		j,
		"placeholderSubstitutionRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference {
	_init_.Initialize()

	if err := validateNewCloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudSecurityComplianceCloudControl.CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference_Override(c CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudSecurityComplianceCloudControl.CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) PutAttributeSubstitutionRule(value *CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesAttributeSubstitutionRule) {
	if err := c.validatePutAttributeSubstitutionRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAttributeSubstitutionRule",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) PutPlaceholderSubstitutionRule(value *CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesPlaceholderSubstitutionRule) {
	if err := c.validatePutPlaceholderSubstitutionRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPlaceholderSubstitutionRule",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) ResetAttributeSubstitutionRule() {
	_jsii_.InvokeVoid(
		c,
		"resetAttributeSubstitutionRule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) ResetPlaceholderSubstitutionRule() {
	_jsii_.InvokeVoid(
		c,
		"resetPlaceholderSubstitutionRule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

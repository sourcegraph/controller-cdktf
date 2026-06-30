package computesecuritypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/computesecuritypolicy/internal"
)

type ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Enable() any
	SetEnable(val any)
	EnableInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig
	SetInternalValue(val *ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig)
	RuleVisibility() *string
	SetRuleVisibility(val *string)
	RuleVisibilityInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThresholdConfigs() ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigThresholdConfigsList
	ThresholdConfigsInput() any
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
	PutThresholdConfigs(value any)
	ResetEnable()
	ResetRuleVisibility()
	ResetThresholdConfigs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference
type jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) Enable() any {
	var returns any
	_jsii_.Get(
		j,
		"enable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) EnableInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) InternalValue() *ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig {
	var returns *ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) RuleVisibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) RuleVisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleVisibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) ThresholdConfigs() ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigThresholdConfigsList {
	var returns ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigThresholdConfigsList
	_jsii_.Get(
		j,
		"thresholdConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) ThresholdConfigsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"thresholdConfigsInput",
		&returns,
	)
	return returns
}

func NewComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference {
	_init_.Initialize()

	if err := validateNewComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeSecurityPolicy.ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference_Override(c ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeSecurityPolicy.ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) SetEnable(val any) {
	if err := j.validateSetEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enable",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) SetInternalValue(val *ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) SetRuleVisibility(val *string) {
	if err := j.validateSetRuleVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleVisibility",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) PutThresholdConfigs(value any) {
	if err := c.validatePutThresholdConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putThresholdConfigs",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) ResetEnable() {
	_jsii_.InvokeVoid(
		c,
		"resetEnable",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) ResetRuleVisibility() {
	_jsii_.InvokeVoid(
		c,
		"resetRuleVisibility",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) ResetThresholdConfigs() {
	_jsii_.InvokeVoid(
		c,
		"resetThresholdConfigs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

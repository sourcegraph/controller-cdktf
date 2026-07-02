package computeorganizationsecuritypolicyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeorganizationsecuritypolicyrule/internal"
)

type ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	RequestCookie() ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList
	RequestCookieInput() any
	RequestHeader() ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeaderList
	RequestHeaderInput() any
	RequestQueryParam() ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList
	RequestQueryParamInput() any
	RequestUri() ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestUriList
	RequestUriInput() any
	TargetRuleIds() *[]*string
	SetTargetRuleIds(val *[]*string)
	TargetRuleIdsInput() *[]*string
	TargetRuleSet() *string
	SetTargetRuleSet(val *string)
	TargetRuleSetInput() *string
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
	PutRequestCookie(value any)
	PutRequestHeader(value any)
	PutRequestQueryParam(value any)
	PutRequestUri(value any)
	ResetRequestCookie()
	ResetRequestHeader()
	ResetRequestQueryParam()
	ResetRequestUri()
	ResetTargetRuleIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference
type jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) RequestCookie() ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList {
	var returns ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList
	_jsii_.Get(
		j,
		"requestCookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) RequestCookieInput() any {
	var returns any
	_jsii_.Get(
		j,
		"requestCookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) RequestHeader() ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeaderList {
	var returns ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeaderList
	_jsii_.Get(
		j,
		"requestHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) RequestHeaderInput() any {
	var returns any
	_jsii_.Get(
		j,
		"requestHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) RequestQueryParam() ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList {
	var returns ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList
	_jsii_.Get(
		j,
		"requestQueryParam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) RequestQueryParamInput() any {
	var returns any
	_jsii_.Get(
		j,
		"requestQueryParamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) RequestUri() ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestUriList {
	var returns ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestUriList
	_jsii_.Get(
		j,
		"requestUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) RequestUriInput() any {
	var returns any
	_jsii_.Get(
		j,
		"requestUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) TargetRuleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetRuleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) TargetRuleIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetRuleIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) TargetRuleSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRuleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) TargetRuleSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRuleSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference {
	_init_.Initialize()

	if err := validateNewComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeOrganizationSecurityPolicyRule.ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference_Override(c ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeOrganizationSecurityPolicyRule.ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) SetTargetRuleIds(val *[]*string) {
	if err := j.validateSetTargetRuleIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetRuleIds",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) SetTargetRuleSet(val *string) {
	if err := j.validateSetTargetRuleSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetRuleSet",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) PutRequestCookie(value any) {
	if err := c.validatePutRequestCookieParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestCookie",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) PutRequestHeader(value any) {
	if err := c.validatePutRequestHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestHeader",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) PutRequestQueryParam(value any) {
	if err := c.validatePutRequestQueryParamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestQueryParam",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) PutRequestUri(value any) {
	if err := c.validatePutRequestUriParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestUri",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) ResetRequestCookie() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestCookie",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) ResetRequestHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) ResetRequestQueryParam() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestQueryParam",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) ResetRequestUri() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestUri",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) ResetTargetRuleIds() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetRuleIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

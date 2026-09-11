package computeorganizationsecuritypolicyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeorganizationsecuritypolicyrule/internal"
)

type ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList
type jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList {
	_init_.Initialize()

	if err := validateNewComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList{}

	_jsii_.Create(
		"@cdktn/provider-google.computeOrganizationSecurityPolicyRule.ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList_Override(c ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeOrganizationSecurityPolicyRule.ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := c.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		c,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList) Get(index *float64) ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieOutputReference {
	if err := c.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


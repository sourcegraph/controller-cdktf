package computeorganizationsecuritypolicyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeorganizationsecuritypolicyrule/internal"
)

type ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList interface {
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
	Get(index *float64) ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList
type jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList {
	_init_.Initialize()

	if err := validateNewComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList{}

	_jsii_.Create(
		"@cdktn/provider-google.computeOrganizationSecurityPolicyRule.ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList_Override(c ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeOrganizationSecurityPolicyRule.ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList) Get(index *float64) ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamOutputReference {
	if err := c.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


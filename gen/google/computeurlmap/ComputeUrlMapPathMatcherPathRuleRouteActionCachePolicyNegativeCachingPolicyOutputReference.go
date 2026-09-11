package computeurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeurlmap/internal"
)

type ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference interface {
	cdktn.ComplexObject
	Code() *float64
	SetCode(val *float64)
	CodeInput() *float64
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
	Ttl() ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyTtlOutputReference
	TtlInput() *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyTtl
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
	PutTtl(value *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyTtl)
	ResetCode()
	ResetTtl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference
type jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) Code() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"code",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) CodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"codeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) Ttl() ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyTtlOutputReference {
	var returns ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyTtlOutputReference
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) TtlInput() *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyTtl {
	var returns *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyTtl
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}


func NewComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.computeUrlMap.ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference_Override(c ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeUrlMap.ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference)SetCode(val *float64) {
	if err := j.validateSetCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"code",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) PutTtl(value *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyTtl) {
	if err := c.validatePutTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTtl",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) ResetCode() {
	_jsii_.InvokeVoid(
		c,
		"resetCode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) ResetTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


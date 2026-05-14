package computeurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeurlmap/internal"
)

type ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference interface {
	cdktf.ComplexObject
	CacheBypassRequestHeaderNames() *[]*string
	SetCacheBypassRequestHeaderNames(val *[]*string)
	CacheBypassRequestHeaderNamesInput() *[]*string
	CacheKeyPolicy() ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyCacheKeyPolicyOutputReference
	CacheKeyPolicyInput() *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyCacheKeyPolicy
	CacheMode() *string
	SetCacheMode(val *string)
	CacheModeInput() *string
	ClientTtl() ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyClientTtlOutputReference
	ClientTtlInput() *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyClientTtl
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
	DefaultTtl() ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyDefaultTtlOutputReference
	DefaultTtlInput() *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyDefaultTtl
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicy
	SetInternalValue(val *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicy)
	MaxTtl() ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyMaxTtlOutputReference
	MaxTtlInput() *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyMaxTtl
	NegativeCaching() interface{}
	SetNegativeCaching(val interface{})
	NegativeCachingInput() interface{}
	NegativeCachingPolicy() ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyNegativeCachingPolicyList
	NegativeCachingPolicyInput() interface{}
	RequestCoalescing() interface{}
	SetRequestCoalescing(val interface{})
	RequestCoalescingInput() interface{}
	ServeWhileStale() ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyServeWhileStaleOutputReference
	ServeWhileStaleInput() *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyServeWhileStale
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
	PutCacheKeyPolicy(value *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyCacheKeyPolicy)
	PutClientTtl(value *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyClientTtl)
	PutDefaultTtl(value *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyDefaultTtl)
	PutMaxTtl(value *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyMaxTtl)
	PutNegativeCachingPolicy(value interface{})
	PutServeWhileStale(value *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyServeWhileStale)
	ResetCacheBypassRequestHeaderNames()
	ResetCacheKeyPolicy()
	ResetCacheMode()
	ResetClientTtl()
	ResetDefaultTtl()
	ResetMaxTtl()
	ResetNegativeCaching()
	ResetNegativeCachingPolicy()
	ResetRequestCoalescing()
	ResetServeWhileStale()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference
type jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) CacheBypassRequestHeaderNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheBypassRequestHeaderNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) CacheBypassRequestHeaderNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheBypassRequestHeaderNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) CacheKeyPolicy() ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyCacheKeyPolicyOutputReference {
	var returns ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyCacheKeyPolicyOutputReference
	_jsii_.Get(
		j,
		"cacheKeyPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) CacheKeyPolicyInput() *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyCacheKeyPolicy {
	var returns *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyCacheKeyPolicy
	_jsii_.Get(
		j,
		"cacheKeyPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) CacheMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) CacheModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ClientTtl() ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyClientTtlOutputReference {
	var returns ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyClientTtlOutputReference
	_jsii_.Get(
		j,
		"clientTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ClientTtlInput() *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyClientTtl {
	var returns *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyClientTtl
	_jsii_.Get(
		j,
		"clientTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) DefaultTtl() ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyDefaultTtlOutputReference {
	var returns ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyDefaultTtlOutputReference
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) DefaultTtlInput() *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyDefaultTtl {
	var returns *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyDefaultTtl
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) InternalValue() *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicy {
	var returns *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) MaxTtl() ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyMaxTtlOutputReference {
	var returns ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyMaxTtlOutputReference
	_jsii_.Get(
		j,
		"maxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) MaxTtlInput() *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyMaxTtl {
	var returns *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyMaxTtl
	_jsii_.Get(
		j,
		"maxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) NegativeCaching() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeCaching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) NegativeCachingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeCachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) NegativeCachingPolicy() ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyNegativeCachingPolicyList {
	var returns ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyNegativeCachingPolicyList
	_jsii_.Get(
		j,
		"negativeCachingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) NegativeCachingPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeCachingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) RequestCoalescing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestCoalescing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) RequestCoalescingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestCoalescingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ServeWhileStale() ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyServeWhileStaleOutputReference {
	var returns ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyServeWhileStaleOutputReference
	_jsii_.Get(
		j,
		"serveWhileStale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ServeWhileStaleInput() *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyServeWhileStale {
	var returns *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyServeWhileStale
	_jsii_.Get(
		j,
		"serveWhileStaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeUrlMap.ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference_Override(c ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeUrlMap.ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference)SetCacheBypassRequestHeaderNames(val *[]*string) {
	if err := j.validateSetCacheBypassRequestHeaderNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheBypassRequestHeaderNames",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference)SetCacheMode(val *string) {
	if err := j.validateSetCacheModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheMode",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference)SetInternalValue(val *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference)SetNegativeCaching(val interface{}) {
	if err := j.validateSetNegativeCachingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"negativeCaching",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference)SetRequestCoalescing(val interface{}) {
	if err := j.validateSetRequestCoalescingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestCoalescing",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) PutCacheKeyPolicy(value *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyCacheKeyPolicy) {
	if err := c.validatePutCacheKeyPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCacheKeyPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) PutClientTtl(value *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyClientTtl) {
	if err := c.validatePutClientTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putClientTtl",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) PutDefaultTtl(value *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyDefaultTtl) {
	if err := c.validatePutDefaultTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDefaultTtl",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) PutMaxTtl(value *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyMaxTtl) {
	if err := c.validatePutMaxTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMaxTtl",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) PutNegativeCachingPolicy(value interface{}) {
	if err := c.validatePutNegativeCachingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNegativeCachingPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) PutServeWhileStale(value *ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyServeWhileStale) {
	if err := c.validatePutServeWhileStaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServeWhileStale",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ResetCacheBypassRequestHeaderNames() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheBypassRequestHeaderNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ResetCacheKeyPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheKeyPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ResetCacheMode() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ResetClientTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetClientTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ResetMaxTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ResetNegativeCaching() {
	_jsii_.InvokeVoid(
		c,
		"resetNegativeCaching",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ResetNegativeCachingPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetNegativeCachingPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ResetRequestCoalescing() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestCoalescing",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ResetServeWhileStale() {
	_jsii_.InvokeVoid(
		c,
		"resetServeWhileStale",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


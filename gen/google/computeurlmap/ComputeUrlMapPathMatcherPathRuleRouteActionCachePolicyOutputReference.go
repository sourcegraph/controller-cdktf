package computeurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeurlmap/internal"
)

type ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference interface {
	cdktn.ComplexObject
	CacheBypassRequestHeaderNames() *[]*string
	SetCacheBypassRequestHeaderNames(val *[]*string)
	CacheBypassRequestHeaderNamesInput() *[]*string
	CacheKeyPolicy() ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyCacheKeyPolicyOutputReference
	CacheKeyPolicyInput() *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyCacheKeyPolicy
	CacheMode() *string
	SetCacheMode(val *string)
	CacheModeInput() *string
	ClientTtl() ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyClientTtlOutputReference
	ClientTtlInput() *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyClientTtl
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
	DefaultTtl() ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyDefaultTtlOutputReference
	DefaultTtlInput() *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyDefaultTtl
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicy
	SetInternalValue(val *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicy)
	MaxTtl() ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyMaxTtlOutputReference
	MaxTtlInput() *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyMaxTtl
	NegativeCaching() interface{}
	SetNegativeCaching(val interface{})
	NegativeCachingInput() interface{}
	NegativeCachingPolicy() ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyList
	NegativeCachingPolicyInput() interface{}
	RequestCoalescing() interface{}
	SetRequestCoalescing(val interface{})
	RequestCoalescingInput() interface{}
	ServeWhileStale() ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyServeWhileStaleOutputReference
	ServeWhileStaleInput() *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyServeWhileStale
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
	PutCacheKeyPolicy(value *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyCacheKeyPolicy)
	PutClientTtl(value *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyClientTtl)
	PutDefaultTtl(value *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyDefaultTtl)
	PutMaxTtl(value *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyMaxTtl)
	PutNegativeCachingPolicy(value interface{})
	PutServeWhileStale(value *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyServeWhileStale)
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
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference
type jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) CacheBypassRequestHeaderNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheBypassRequestHeaderNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) CacheBypassRequestHeaderNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheBypassRequestHeaderNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) CacheKeyPolicy() ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyCacheKeyPolicyOutputReference {
	var returns ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyCacheKeyPolicyOutputReference
	_jsii_.Get(
		j,
		"cacheKeyPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) CacheKeyPolicyInput() *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyCacheKeyPolicy {
	var returns *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyCacheKeyPolicy
	_jsii_.Get(
		j,
		"cacheKeyPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) CacheMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) CacheModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) ClientTtl() ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyClientTtlOutputReference {
	var returns ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyClientTtlOutputReference
	_jsii_.Get(
		j,
		"clientTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) ClientTtlInput() *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyClientTtl {
	var returns *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyClientTtl
	_jsii_.Get(
		j,
		"clientTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) DefaultTtl() ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyDefaultTtlOutputReference {
	var returns ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyDefaultTtlOutputReference
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) DefaultTtlInput() *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyDefaultTtl {
	var returns *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyDefaultTtl
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) InternalValue() *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicy {
	var returns *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) MaxTtl() ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyMaxTtlOutputReference {
	var returns ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyMaxTtlOutputReference
	_jsii_.Get(
		j,
		"maxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) MaxTtlInput() *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyMaxTtl {
	var returns *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyMaxTtl
	_jsii_.Get(
		j,
		"maxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) NegativeCaching() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeCaching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) NegativeCachingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeCachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) NegativeCachingPolicy() ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyList {
	var returns ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyList
	_jsii_.Get(
		j,
		"negativeCachingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) NegativeCachingPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeCachingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) RequestCoalescing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestCoalescing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) RequestCoalescingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestCoalescingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) ServeWhileStale() ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyServeWhileStaleOutputReference {
	var returns ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyServeWhileStaleOutputReference
	_jsii_.Get(
		j,
		"serveWhileStale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) ServeWhileStaleInput() *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyServeWhileStale {
	var returns *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyServeWhileStale
	_jsii_.Get(
		j,
		"serveWhileStaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.computeUrlMap.ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference_Override(c ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeUrlMap.ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference)SetCacheBypassRequestHeaderNames(val *[]*string) {
	if err := j.validateSetCacheBypassRequestHeaderNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheBypassRequestHeaderNames",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference)SetCacheMode(val *string) {
	if err := j.validateSetCacheModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheMode",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference)SetInternalValue(val *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference)SetNegativeCaching(val interface{}) {
	if err := j.validateSetNegativeCachingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"negativeCaching",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference)SetRequestCoalescing(val interface{}) {
	if err := j.validateSetRequestCoalescingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestCoalescing",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) PutCacheKeyPolicy(value *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyCacheKeyPolicy) {
	if err := c.validatePutCacheKeyPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCacheKeyPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) PutClientTtl(value *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyClientTtl) {
	if err := c.validatePutClientTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putClientTtl",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) PutDefaultTtl(value *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyDefaultTtl) {
	if err := c.validatePutDefaultTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDefaultTtl",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) PutMaxTtl(value *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyMaxTtl) {
	if err := c.validatePutMaxTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMaxTtl",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) PutNegativeCachingPolicy(value interface{}) {
	if err := c.validatePutNegativeCachingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNegativeCachingPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) PutServeWhileStale(value *ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyServeWhileStale) {
	if err := c.validatePutServeWhileStaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServeWhileStale",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) ResetCacheBypassRequestHeaderNames() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheBypassRequestHeaderNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) ResetCacheKeyPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheKeyPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) ResetCacheMode() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) ResetClientTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetClientTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) ResetMaxTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) ResetNegativeCaching() {
	_jsii_.InvokeVoid(
		c,
		"resetNegativeCaching",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) ResetNegativeCachingPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetNegativeCachingPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) ResetRequestCoalescing() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestCoalescing",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) ResetServeWhileStale() {
	_jsii_.InvokeVoid(
		c,
		"resetServeWhileStale",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


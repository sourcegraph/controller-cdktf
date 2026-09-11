package computeurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeurlmap/internal"
)

type ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference interface {
	cdktn.ComplexObject
	CacheBypassRequestHeaderNames() *[]*string
	SetCacheBypassRequestHeaderNames(val *[]*string)
	CacheBypassRequestHeaderNamesInput() *[]*string
	CacheKeyPolicy() ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference
	CacheKeyPolicyInput() *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicy
	CacheMode() *string
	SetCacheMode(val *string)
	CacheModeInput() *string
	ClientTtl() ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyClientTtlOutputReference
	ClientTtlInput() *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyClientTtl
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
	DefaultTtl() ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyDefaultTtlOutputReference
	DefaultTtlInput() *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyDefaultTtl
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicy
	SetInternalValue(val *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicy)
	MaxTtl() ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyMaxTtlOutputReference
	MaxTtlInput() *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyMaxTtl
	NegativeCaching() interface{}
	SetNegativeCaching(val interface{})
	NegativeCachingInput() interface{}
	NegativeCachingPolicy() ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyNegativeCachingPolicyList
	NegativeCachingPolicyInput() interface{}
	RequestCoalescing() interface{}
	SetRequestCoalescing(val interface{})
	RequestCoalescingInput() interface{}
	ServeWhileStale() ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyServeWhileStaleOutputReference
	ServeWhileStaleInput() *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyServeWhileStale
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
	PutCacheKeyPolicy(value *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicy)
	PutClientTtl(value *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyClientTtl)
	PutDefaultTtl(value *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyDefaultTtl)
	PutMaxTtl(value *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyMaxTtl)
	PutNegativeCachingPolicy(value interface{})
	PutServeWhileStale(value *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyServeWhileStale)
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

// The jsii proxy struct for ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference
type jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) CacheBypassRequestHeaderNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheBypassRequestHeaderNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) CacheBypassRequestHeaderNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheBypassRequestHeaderNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) CacheKeyPolicy() ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference {
	var returns ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference
	_jsii_.Get(
		j,
		"cacheKeyPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) CacheKeyPolicyInput() *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicy {
	var returns *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicy
	_jsii_.Get(
		j,
		"cacheKeyPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) CacheMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) CacheModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) ClientTtl() ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyClientTtlOutputReference {
	var returns ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyClientTtlOutputReference
	_jsii_.Get(
		j,
		"clientTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) ClientTtlInput() *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyClientTtl {
	var returns *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyClientTtl
	_jsii_.Get(
		j,
		"clientTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) DefaultTtl() ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyDefaultTtlOutputReference {
	var returns ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyDefaultTtlOutputReference
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) DefaultTtlInput() *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyDefaultTtl {
	var returns *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyDefaultTtl
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) InternalValue() *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicy {
	var returns *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) MaxTtl() ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyMaxTtlOutputReference {
	var returns ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyMaxTtlOutputReference
	_jsii_.Get(
		j,
		"maxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) MaxTtlInput() *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyMaxTtl {
	var returns *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyMaxTtl
	_jsii_.Get(
		j,
		"maxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) NegativeCaching() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeCaching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) NegativeCachingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeCachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) NegativeCachingPolicy() ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyNegativeCachingPolicyList {
	var returns ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyNegativeCachingPolicyList
	_jsii_.Get(
		j,
		"negativeCachingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) NegativeCachingPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeCachingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) RequestCoalescing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestCoalescing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) RequestCoalescingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestCoalescingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) ServeWhileStale() ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyServeWhileStaleOutputReference {
	var returns ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyServeWhileStaleOutputReference
	_jsii_.Get(
		j,
		"serveWhileStale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) ServeWhileStaleInput() *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyServeWhileStale {
	var returns *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyServeWhileStale
	_jsii_.Get(
		j,
		"serveWhileStaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.computeUrlMap.ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference_Override(c ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeUrlMap.ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference)SetCacheBypassRequestHeaderNames(val *[]*string) {
	if err := j.validateSetCacheBypassRequestHeaderNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheBypassRequestHeaderNames",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference)SetCacheMode(val *string) {
	if err := j.validateSetCacheModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheMode",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference)SetInternalValue(val *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference)SetNegativeCaching(val interface{}) {
	if err := j.validateSetNegativeCachingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"negativeCaching",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference)SetRequestCoalescing(val interface{}) {
	if err := j.validateSetRequestCoalescingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestCoalescing",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) PutCacheKeyPolicy(value *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicy) {
	if err := c.validatePutCacheKeyPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCacheKeyPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) PutClientTtl(value *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyClientTtl) {
	if err := c.validatePutClientTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putClientTtl",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) PutDefaultTtl(value *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyDefaultTtl) {
	if err := c.validatePutDefaultTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDefaultTtl",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) PutMaxTtl(value *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyMaxTtl) {
	if err := c.validatePutMaxTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMaxTtl",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) PutNegativeCachingPolicy(value interface{}) {
	if err := c.validatePutNegativeCachingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNegativeCachingPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) PutServeWhileStale(value *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyServeWhileStale) {
	if err := c.validatePutServeWhileStaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServeWhileStale",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) ResetCacheBypassRequestHeaderNames() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheBypassRequestHeaderNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) ResetCacheKeyPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheKeyPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) ResetCacheMode() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) ResetClientTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetClientTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) ResetMaxTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) ResetNegativeCaching() {
	_jsii_.InvokeVoid(
		c,
		"resetNegativeCaching",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) ResetNegativeCachingPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetNegativeCachingPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) ResetRequestCoalescing() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestCoalescing",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) ResetServeWhileStale() {
	_jsii_.InvokeVoid(
		c,
		"resetServeWhileStale",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


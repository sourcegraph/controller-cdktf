package computeurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeurlmap/internal"
)

type ComputeUrlMapDefaultRouteActionCachePolicyOutputReference interface {
	cdktf.ComplexObject
	CacheBypassRequestHeaderNames() *[]*string
	SetCacheBypassRequestHeaderNames(val *[]*string)
	CacheBypassRequestHeaderNamesInput() *[]*string
	CacheKeyPolicy() ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference
	CacheKeyPolicyInput() *ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicy
	CacheMode() *string
	SetCacheMode(val *string)
	CacheModeInput() *string
	ClientTtl() ComputeUrlMapDefaultRouteActionCachePolicyClientTtlOutputReference
	ClientTtlInput() *ComputeUrlMapDefaultRouteActionCachePolicyClientTtl
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
	DefaultTtl() ComputeUrlMapDefaultRouteActionCachePolicyDefaultTtlOutputReference
	DefaultTtlInput() *ComputeUrlMapDefaultRouteActionCachePolicyDefaultTtl
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeUrlMapDefaultRouteActionCachePolicy
	SetInternalValue(val *ComputeUrlMapDefaultRouteActionCachePolicy)
	MaxTtl() ComputeUrlMapDefaultRouteActionCachePolicyMaxTtlOutputReference
	MaxTtlInput() *ComputeUrlMapDefaultRouteActionCachePolicyMaxTtl
	NegativeCaching() any
	SetNegativeCaching(val any)
	NegativeCachingInput() any
	NegativeCachingPolicy() ComputeUrlMapDefaultRouteActionCachePolicyNegativeCachingPolicyList
	NegativeCachingPolicyInput() any
	RequestCoalescing() any
	SetRequestCoalescing(val any)
	RequestCoalescingInput() any
	ServeWhileStale() ComputeUrlMapDefaultRouteActionCachePolicyServeWhileStaleOutputReference
	ServeWhileStaleInput() *ComputeUrlMapDefaultRouteActionCachePolicyServeWhileStale
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
	PutCacheKeyPolicy(value *ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicy)
	PutClientTtl(value *ComputeUrlMapDefaultRouteActionCachePolicyClientTtl)
	PutDefaultTtl(value *ComputeUrlMapDefaultRouteActionCachePolicyDefaultTtl)
	PutMaxTtl(value *ComputeUrlMapDefaultRouteActionCachePolicyMaxTtl)
	PutNegativeCachingPolicy(value any)
	PutServeWhileStale(value *ComputeUrlMapDefaultRouteActionCachePolicyServeWhileStale)
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
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeUrlMapDefaultRouteActionCachePolicyOutputReference
type jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) CacheBypassRequestHeaderNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheBypassRequestHeaderNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) CacheBypassRequestHeaderNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheBypassRequestHeaderNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) CacheKeyPolicy() ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference {
	var returns ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference
	_jsii_.Get(
		j,
		"cacheKeyPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) CacheKeyPolicyInput() *ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicy {
	var returns *ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicy
	_jsii_.Get(
		j,
		"cacheKeyPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) CacheMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) CacheModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) ClientTtl() ComputeUrlMapDefaultRouteActionCachePolicyClientTtlOutputReference {
	var returns ComputeUrlMapDefaultRouteActionCachePolicyClientTtlOutputReference
	_jsii_.Get(
		j,
		"clientTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) ClientTtlInput() *ComputeUrlMapDefaultRouteActionCachePolicyClientTtl {
	var returns *ComputeUrlMapDefaultRouteActionCachePolicyClientTtl
	_jsii_.Get(
		j,
		"clientTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) DefaultTtl() ComputeUrlMapDefaultRouteActionCachePolicyDefaultTtlOutputReference {
	var returns ComputeUrlMapDefaultRouteActionCachePolicyDefaultTtlOutputReference
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) DefaultTtlInput() *ComputeUrlMapDefaultRouteActionCachePolicyDefaultTtl {
	var returns *ComputeUrlMapDefaultRouteActionCachePolicyDefaultTtl
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) InternalValue() *ComputeUrlMapDefaultRouteActionCachePolicy {
	var returns *ComputeUrlMapDefaultRouteActionCachePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) MaxTtl() ComputeUrlMapDefaultRouteActionCachePolicyMaxTtlOutputReference {
	var returns ComputeUrlMapDefaultRouteActionCachePolicyMaxTtlOutputReference
	_jsii_.Get(
		j,
		"maxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) MaxTtlInput() *ComputeUrlMapDefaultRouteActionCachePolicyMaxTtl {
	var returns *ComputeUrlMapDefaultRouteActionCachePolicyMaxTtl
	_jsii_.Get(
		j,
		"maxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) NegativeCaching() any {
	var returns any
	_jsii_.Get(
		j,
		"negativeCaching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) NegativeCachingInput() any {
	var returns any
	_jsii_.Get(
		j,
		"negativeCachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) NegativeCachingPolicy() ComputeUrlMapDefaultRouteActionCachePolicyNegativeCachingPolicyList {
	var returns ComputeUrlMapDefaultRouteActionCachePolicyNegativeCachingPolicyList
	_jsii_.Get(
		j,
		"negativeCachingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) NegativeCachingPolicyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"negativeCachingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) RequestCoalescing() any {
	var returns any
	_jsii_.Get(
		j,
		"requestCoalescing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) RequestCoalescingInput() any {
	var returns any
	_jsii_.Get(
		j,
		"requestCoalescingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) ServeWhileStale() ComputeUrlMapDefaultRouteActionCachePolicyServeWhileStaleOutputReference {
	var returns ComputeUrlMapDefaultRouteActionCachePolicyServeWhileStaleOutputReference
	_jsii_.Get(
		j,
		"serveWhileStale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) ServeWhileStaleInput() *ComputeUrlMapDefaultRouteActionCachePolicyServeWhileStale {
	var returns *ComputeUrlMapDefaultRouteActionCachePolicyServeWhileStale
	_jsii_.Get(
		j,
		"serveWhileStaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewComputeUrlMapDefaultRouteActionCachePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeUrlMapDefaultRouteActionCachePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewComputeUrlMapDefaultRouteActionCachePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeUrlMap.ComputeUrlMapDefaultRouteActionCachePolicyOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeUrlMapDefaultRouteActionCachePolicyOutputReference_Override(c ComputeUrlMapDefaultRouteActionCachePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeUrlMap.ComputeUrlMapDefaultRouteActionCachePolicyOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) SetCacheBypassRequestHeaderNames(val *[]*string) {
	if err := j.validateSetCacheBypassRequestHeaderNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheBypassRequestHeaderNames",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) SetCacheMode(val *string) {
	if err := j.validateSetCacheModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheMode",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) SetInternalValue(val *ComputeUrlMapDefaultRouteActionCachePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) SetNegativeCaching(val any) {
	if err := j.validateSetNegativeCachingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"negativeCaching",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) SetRequestCoalescing(val any) {
	if err := j.validateSetRequestCoalescingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestCoalescing",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) PutCacheKeyPolicy(value *ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicy) {
	if err := c.validatePutCacheKeyPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCacheKeyPolicy",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) PutClientTtl(value *ComputeUrlMapDefaultRouteActionCachePolicyClientTtl) {
	if err := c.validatePutClientTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putClientTtl",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) PutDefaultTtl(value *ComputeUrlMapDefaultRouteActionCachePolicyDefaultTtl) {
	if err := c.validatePutDefaultTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDefaultTtl",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) PutMaxTtl(value *ComputeUrlMapDefaultRouteActionCachePolicyMaxTtl) {
	if err := c.validatePutMaxTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMaxTtl",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) PutNegativeCachingPolicy(value any) {
	if err := c.validatePutNegativeCachingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNegativeCachingPolicy",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) PutServeWhileStale(value *ComputeUrlMapDefaultRouteActionCachePolicyServeWhileStale) {
	if err := c.validatePutServeWhileStaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServeWhileStale",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) ResetCacheBypassRequestHeaderNames() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheBypassRequestHeaderNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) ResetCacheKeyPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheKeyPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) ResetCacheMode() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) ResetClientTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetClientTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) ResetMaxTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) ResetNegativeCaching() {
	_jsii_.InvokeVoid(
		c,
		"resetNegativeCaching",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) ResetNegativeCachingPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetNegativeCachingPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) ResetRequestCoalescing() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestCoalescing",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) ResetServeWhileStale() {
	_jsii_.InvokeVoid(
		c,
		"resetServeWhileStale",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

package googlecomputeurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeurlmap/internal"
)

type GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference interface {
	cdktn.ComplexObject
	CacheBypassRequestHeaderNames() *[]*string
	SetCacheBypassRequestHeaderNames(val *[]*string)
	CacheBypassRequestHeaderNamesInput() *[]*string
	CacheKeyPolicy() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyCacheKeyPolicyOutputReference
	CacheKeyPolicyInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyCacheKeyPolicy
	CacheMode() *string
	SetCacheMode(val *string)
	CacheModeInput() *string
	ClientTtl() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyClientTtlOutputReference
	ClientTtlInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyClientTtl
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
	DefaultTtl() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyDefaultTtlOutputReference
	DefaultTtlInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyDefaultTtl
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicy
	SetInternalValue(val *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicy)
	MaxTtl() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyMaxTtlOutputReference
	MaxTtlInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyMaxTtl
	NegativeCaching() interface{}
	SetNegativeCaching(val interface{})
	NegativeCachingInput() interface{}
	NegativeCachingPolicy() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyNegativeCachingPolicyList
	NegativeCachingPolicyInput() interface{}
	RequestCoalescing() interface{}
	SetRequestCoalescing(val interface{})
	RequestCoalescingInput() interface{}
	ServeWhileStale() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyServeWhileStaleOutputReference
	ServeWhileStaleInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyServeWhileStale
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
	PutCacheKeyPolicy(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyCacheKeyPolicy)
	PutClientTtl(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyClientTtl)
	PutDefaultTtl(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyDefaultTtl)
	PutMaxTtl(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyMaxTtl)
	PutNegativeCachingPolicy(value interface{})
	PutServeWhileStale(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyServeWhileStale)
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

// The jsii proxy struct for GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference
type jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) CacheBypassRequestHeaderNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheBypassRequestHeaderNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) CacheBypassRequestHeaderNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheBypassRequestHeaderNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) CacheKeyPolicy() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyCacheKeyPolicyOutputReference {
	var returns GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyCacheKeyPolicyOutputReference
	_jsii_.Get(
		j,
		"cacheKeyPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) CacheKeyPolicyInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyCacheKeyPolicy {
	var returns *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyCacheKeyPolicy
	_jsii_.Get(
		j,
		"cacheKeyPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) CacheMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) CacheModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ClientTtl() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyClientTtlOutputReference {
	var returns GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyClientTtlOutputReference
	_jsii_.Get(
		j,
		"clientTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ClientTtlInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyClientTtl {
	var returns *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyClientTtl
	_jsii_.Get(
		j,
		"clientTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) DefaultTtl() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyDefaultTtlOutputReference {
	var returns GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyDefaultTtlOutputReference
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) DefaultTtlInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyDefaultTtl {
	var returns *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyDefaultTtl
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) InternalValue() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicy {
	var returns *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) MaxTtl() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyMaxTtlOutputReference {
	var returns GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyMaxTtlOutputReference
	_jsii_.Get(
		j,
		"maxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) MaxTtlInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyMaxTtl {
	var returns *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyMaxTtl
	_jsii_.Get(
		j,
		"maxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) NegativeCaching() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeCaching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) NegativeCachingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeCachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) NegativeCachingPolicy() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyNegativeCachingPolicyList {
	var returns GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyNegativeCachingPolicyList
	_jsii_.Get(
		j,
		"negativeCachingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) NegativeCachingPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeCachingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) RequestCoalescing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestCoalescing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) RequestCoalescingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestCoalescingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ServeWhileStale() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyServeWhileStaleOutputReference {
	var returns GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyServeWhileStaleOutputReference
	_jsii_.Get(
		j,
		"serveWhileStale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ServeWhileStaleInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyServeWhileStale {
	var returns *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyServeWhileStale
	_jsii_.Get(
		j,
		"serveWhileStaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeUrlMap.GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference_Override(g GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeUrlMap.GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference)SetCacheBypassRequestHeaderNames(val *[]*string) {
	if err := j.validateSetCacheBypassRequestHeaderNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheBypassRequestHeaderNames",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference)SetCacheMode(val *string) {
	if err := j.validateSetCacheModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheMode",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference)SetInternalValue(val *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference)SetNegativeCaching(val interface{}) {
	if err := j.validateSetNegativeCachingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"negativeCaching",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference)SetRequestCoalescing(val interface{}) {
	if err := j.validateSetRequestCoalescingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestCoalescing",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) PutCacheKeyPolicy(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyCacheKeyPolicy) {
	if err := g.validatePutCacheKeyPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCacheKeyPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) PutClientTtl(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyClientTtl) {
	if err := g.validatePutClientTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClientTtl",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) PutDefaultTtl(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyDefaultTtl) {
	if err := g.validatePutDefaultTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDefaultTtl",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) PutMaxTtl(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyMaxTtl) {
	if err := g.validatePutMaxTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaxTtl",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) PutNegativeCachingPolicy(value interface{}) {
	if err := g.validatePutNegativeCachingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNegativeCachingPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) PutServeWhileStale(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyServeWhileStale) {
	if err := g.validatePutServeWhileStaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServeWhileStale",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ResetCacheBypassRequestHeaderNames() {
	_jsii_.InvokeVoid(
		g,
		"resetCacheBypassRequestHeaderNames",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ResetCacheKeyPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetCacheKeyPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ResetCacheMode() {
	_jsii_.InvokeVoid(
		g,
		"resetCacheMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ResetClientTtl() {
	_jsii_.InvokeVoid(
		g,
		"resetClientTtl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ResetMaxTtl() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxTtl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ResetNegativeCaching() {
	_jsii_.InvokeVoid(
		g,
		"resetNegativeCaching",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ResetNegativeCachingPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetNegativeCachingPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ResetRequestCoalescing() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestCoalescing",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ResetServeWhileStale() {
	_jsii_.InvokeVoid(
		g,
		"resetServeWhileStale",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCachePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


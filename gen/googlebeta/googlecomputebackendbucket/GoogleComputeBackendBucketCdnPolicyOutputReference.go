package googlecomputebackendbucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlecomputebackendbucket/internal"
)

type GoogleComputeBackendBucketCdnPolicyOutputReference interface {
	cdktf.ComplexObject
	BypassCacheOnRequestHeaders() GoogleComputeBackendBucketCdnPolicyBypassCacheOnRequestHeadersList
	BypassCacheOnRequestHeadersInput() interface{}
	CacheKeyPolicy() GoogleComputeBackendBucketCdnPolicyCacheKeyPolicyOutputReference
	CacheKeyPolicyInput() *GoogleComputeBackendBucketCdnPolicyCacheKeyPolicy
	CacheMode() *string
	SetCacheMode(val *string)
	CacheModeInput() *string
	ClientTtl() *float64
	SetClientTtl(val *float64)
	ClientTtlInput() *float64
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
	DefaultTtl() *float64
	SetDefaultTtl(val *float64)
	DefaultTtlInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeBackendBucketCdnPolicy
	SetInternalValue(val *GoogleComputeBackendBucketCdnPolicy)
	MaxTtl() *float64
	SetMaxTtl(val *float64)
	MaxTtlInput() *float64
	NegativeCaching() interface{}
	SetNegativeCaching(val interface{})
	NegativeCachingInput() interface{}
	NegativeCachingPolicy() GoogleComputeBackendBucketCdnPolicyNegativeCachingPolicyList
	NegativeCachingPolicyInput() interface{}
	RequestCoalescing() interface{}
	SetRequestCoalescing(val interface{})
	RequestCoalescingInput() interface{}
	ServeWhileStale() *float64
	SetServeWhileStale(val *float64)
	ServeWhileStaleInput() *float64
	SignedUrlCacheMaxAgeSec() *float64
	SetSignedUrlCacheMaxAgeSec(val *float64)
	SignedUrlCacheMaxAgeSecInput() *float64
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
	PutBypassCacheOnRequestHeaders(value interface{})
	PutCacheKeyPolicy(value *GoogleComputeBackendBucketCdnPolicyCacheKeyPolicy)
	PutNegativeCachingPolicy(value interface{})
	ResetBypassCacheOnRequestHeaders()
	ResetCacheKeyPolicy()
	ResetCacheMode()
	ResetClientTtl()
	ResetDefaultTtl()
	ResetMaxTtl()
	ResetNegativeCaching()
	ResetNegativeCachingPolicy()
	ResetRequestCoalescing()
	ResetServeWhileStale()
	ResetSignedUrlCacheMaxAgeSec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeBackendBucketCdnPolicyOutputReference
type jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) BypassCacheOnRequestHeaders() GoogleComputeBackendBucketCdnPolicyBypassCacheOnRequestHeadersList {
	var returns GoogleComputeBackendBucketCdnPolicyBypassCacheOnRequestHeadersList
	_jsii_.Get(
		j,
		"bypassCacheOnRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) BypassCacheOnRequestHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassCacheOnRequestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) CacheKeyPolicy() GoogleComputeBackendBucketCdnPolicyCacheKeyPolicyOutputReference {
	var returns GoogleComputeBackendBucketCdnPolicyCacheKeyPolicyOutputReference
	_jsii_.Get(
		j,
		"cacheKeyPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) CacheKeyPolicyInput() *GoogleComputeBackendBucketCdnPolicyCacheKeyPolicy {
	var returns *GoogleComputeBackendBucketCdnPolicyCacheKeyPolicy
	_jsii_.Get(
		j,
		"cacheKeyPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) CacheMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) CacheModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) ClientTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) ClientTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) DefaultTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) DefaultTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) InternalValue() *GoogleComputeBackendBucketCdnPolicy {
	var returns *GoogleComputeBackendBucketCdnPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) MaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) MaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) NegativeCaching() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeCaching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) NegativeCachingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeCachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) NegativeCachingPolicy() GoogleComputeBackendBucketCdnPolicyNegativeCachingPolicyList {
	var returns GoogleComputeBackendBucketCdnPolicyNegativeCachingPolicyList
	_jsii_.Get(
		j,
		"negativeCachingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) NegativeCachingPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeCachingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) RequestCoalescing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestCoalescing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) RequestCoalescingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestCoalescingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) ServeWhileStale() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serveWhileStale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) ServeWhileStaleInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serveWhileStaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) SignedUrlCacheMaxAgeSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signedUrlCacheMaxAgeSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) SignedUrlCacheMaxAgeSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signedUrlCacheMaxAgeSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeBackendBucketCdnPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeBackendBucketCdnPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeBackendBucketCdnPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleComputeBackendBucket.GoogleComputeBackendBucketCdnPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeBackendBucketCdnPolicyOutputReference_Override(g GoogleComputeBackendBucketCdnPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleComputeBackendBucket.GoogleComputeBackendBucketCdnPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference)SetCacheMode(val *string) {
	if err := j.validateSetCacheModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheMode",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference)SetClientTtl(val *float64) {
	if err := j.validateSetClientTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTtl",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference)SetDefaultTtl(val *float64) {
	if err := j.validateSetDefaultTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTtl",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference)SetInternalValue(val *GoogleComputeBackendBucketCdnPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference)SetMaxTtl(val *float64) {
	if err := j.validateSetMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTtl",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference)SetNegativeCaching(val interface{}) {
	if err := j.validateSetNegativeCachingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"negativeCaching",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference)SetRequestCoalescing(val interface{}) {
	if err := j.validateSetRequestCoalescingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestCoalescing",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference)SetServeWhileStale(val *float64) {
	if err := j.validateSetServeWhileStaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serveWhileStale",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference)SetSignedUrlCacheMaxAgeSec(val *float64) {
	if err := j.validateSetSignedUrlCacheMaxAgeSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signedUrlCacheMaxAgeSec",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) PutBypassCacheOnRequestHeaders(value interface{}) {
	if err := g.validatePutBypassCacheOnRequestHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBypassCacheOnRequestHeaders",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) PutCacheKeyPolicy(value *GoogleComputeBackendBucketCdnPolicyCacheKeyPolicy) {
	if err := g.validatePutCacheKeyPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCacheKeyPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) PutNegativeCachingPolicy(value interface{}) {
	if err := g.validatePutNegativeCachingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNegativeCachingPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) ResetBypassCacheOnRequestHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetBypassCacheOnRequestHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) ResetCacheKeyPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetCacheKeyPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) ResetCacheMode() {
	_jsii_.InvokeVoid(
		g,
		"resetCacheMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) ResetClientTtl() {
	_jsii_.InvokeVoid(
		g,
		"resetClientTtl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) ResetMaxTtl() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxTtl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) ResetNegativeCaching() {
	_jsii_.InvokeVoid(
		g,
		"resetNegativeCaching",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) ResetNegativeCachingPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetNegativeCachingPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) ResetRequestCoalescing() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestCoalescing",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) ResetServeWhileStale() {
	_jsii_.InvokeVoid(
		g,
		"resetServeWhileStale",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) ResetSignedUrlCacheMaxAgeSec() {
	_jsii_.InvokeVoid(
		g,
		"resetSignedUrlCacheMaxAgeSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendBucketCdnPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


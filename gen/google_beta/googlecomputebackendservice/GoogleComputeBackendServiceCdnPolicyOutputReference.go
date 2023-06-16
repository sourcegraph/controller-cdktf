package googlecomputebackendservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputebackendservice/internal"
)

type GoogleComputeBackendServiceCdnPolicyOutputReference interface {
	cdktf.ComplexObject
	BypassCacheOnRequestHeaders() GoogleComputeBackendServiceCdnPolicyBypassCacheOnRequestHeadersList
	BypassCacheOnRequestHeadersInput() interface{}
	CacheKeyPolicy() GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference
	CacheKeyPolicyInput() *GoogleComputeBackendServiceCdnPolicyCacheKeyPolicy
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
	InternalValue() *GoogleComputeBackendServiceCdnPolicy
	SetInternalValue(val *GoogleComputeBackendServiceCdnPolicy)
	MaxTtl() *float64
	SetMaxTtl(val *float64)
	MaxTtlInput() *float64
	NegativeCaching() interface{}
	SetNegativeCaching(val interface{})
	NegativeCachingInput() interface{}
	NegativeCachingPolicy() GoogleComputeBackendServiceCdnPolicyNegativeCachingPolicyList
	NegativeCachingPolicyInput() interface{}
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
	PutCacheKeyPolicy(value *GoogleComputeBackendServiceCdnPolicyCacheKeyPolicy)
	PutNegativeCachingPolicy(value interface{})
	ResetBypassCacheOnRequestHeaders()
	ResetCacheKeyPolicy()
	ResetCacheMode()
	ResetClientTtl()
	ResetDefaultTtl()
	ResetMaxTtl()
	ResetNegativeCaching()
	ResetNegativeCachingPolicy()
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

// The jsii proxy struct for GoogleComputeBackendServiceCdnPolicyOutputReference
type jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) BypassCacheOnRequestHeaders() GoogleComputeBackendServiceCdnPolicyBypassCacheOnRequestHeadersList {
	var returns GoogleComputeBackendServiceCdnPolicyBypassCacheOnRequestHeadersList
	_jsii_.Get(
		j,
		"bypassCacheOnRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) BypassCacheOnRequestHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassCacheOnRequestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) CacheKeyPolicy() GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference {
	var returns GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference
	_jsii_.Get(
		j,
		"cacheKeyPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) CacheKeyPolicyInput() *GoogleComputeBackendServiceCdnPolicyCacheKeyPolicy {
	var returns *GoogleComputeBackendServiceCdnPolicyCacheKeyPolicy
	_jsii_.Get(
		j,
		"cacheKeyPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) CacheMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) CacheModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) ClientTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) ClientTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) DefaultTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) DefaultTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) InternalValue() *GoogleComputeBackendServiceCdnPolicy {
	var returns *GoogleComputeBackendServiceCdnPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) MaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) MaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) NegativeCaching() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeCaching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) NegativeCachingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeCachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) NegativeCachingPolicy() GoogleComputeBackendServiceCdnPolicyNegativeCachingPolicyList {
	var returns GoogleComputeBackendServiceCdnPolicyNegativeCachingPolicyList
	_jsii_.Get(
		j,
		"negativeCachingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) NegativeCachingPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeCachingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) ServeWhileStale() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serveWhileStale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) ServeWhileStaleInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serveWhileStaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) SignedUrlCacheMaxAgeSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signedUrlCacheMaxAgeSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) SignedUrlCacheMaxAgeSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signedUrlCacheMaxAgeSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeBackendServiceCdnPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeBackendServiceCdnPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeBackendServiceCdnPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference{}

	_jsii_.Create(
		"google-beta.googleComputeBackendService.GoogleComputeBackendServiceCdnPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeBackendServiceCdnPolicyOutputReference_Override(g GoogleComputeBackendServiceCdnPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeBackendService.GoogleComputeBackendServiceCdnPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference)SetCacheMode(val *string) {
	if err := j.validateSetCacheModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheMode",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference)SetClientTtl(val *float64) {
	if err := j.validateSetClientTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTtl",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference)SetDefaultTtl(val *float64) {
	if err := j.validateSetDefaultTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTtl",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference)SetInternalValue(val *GoogleComputeBackendServiceCdnPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference)SetMaxTtl(val *float64) {
	if err := j.validateSetMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTtl",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference)SetNegativeCaching(val interface{}) {
	if err := j.validateSetNegativeCachingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"negativeCaching",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference)SetServeWhileStale(val *float64) {
	if err := j.validateSetServeWhileStaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serveWhileStale",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference)SetSignedUrlCacheMaxAgeSec(val *float64) {
	if err := j.validateSetSignedUrlCacheMaxAgeSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signedUrlCacheMaxAgeSec",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) PutBypassCacheOnRequestHeaders(value interface{}) {
	if err := g.validatePutBypassCacheOnRequestHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBypassCacheOnRequestHeaders",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) PutCacheKeyPolicy(value *GoogleComputeBackendServiceCdnPolicyCacheKeyPolicy) {
	if err := g.validatePutCacheKeyPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCacheKeyPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) PutNegativeCachingPolicy(value interface{}) {
	if err := g.validatePutNegativeCachingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNegativeCachingPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) ResetBypassCacheOnRequestHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetBypassCacheOnRequestHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) ResetCacheKeyPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetCacheKeyPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) ResetCacheMode() {
	_jsii_.InvokeVoid(
		g,
		"resetCacheMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) ResetClientTtl() {
	_jsii_.InvokeVoid(
		g,
		"resetClientTtl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) ResetMaxTtl() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxTtl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) ResetNegativeCaching() {
	_jsii_.InvokeVoid(
		g,
		"resetNegativeCaching",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) ResetNegativeCachingPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetNegativeCachingPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) ResetServeWhileStale() {
	_jsii_.InvokeVoid(
		g,
		"resetServeWhileStale",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) ResetSignedUrlCacheMaxAgeSec() {
	_jsii_.InvokeVoid(
		g,
		"resetSignedUrlCacheMaxAgeSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


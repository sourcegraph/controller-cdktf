package googlecomputebackendservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputebackendservice/internal"
)

type GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference interface {
	cdktf.ComplexObject
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
	IncludeHost() interface{}
	SetIncludeHost(val interface{})
	IncludeHostInput() interface{}
	IncludeHttpHeaders() *[]*string
	SetIncludeHttpHeaders(val *[]*string)
	IncludeHttpHeadersInput() *[]*string
	IncludeNamedCookies() *[]*string
	SetIncludeNamedCookies(val *[]*string)
	IncludeNamedCookiesInput() *[]*string
	IncludeProtocol() interface{}
	SetIncludeProtocol(val interface{})
	IncludeProtocolInput() interface{}
	IncludeQueryString() interface{}
	SetIncludeQueryString(val interface{})
	IncludeQueryStringInput() interface{}
	InternalValue() *GoogleComputeBackendServiceCdnPolicyCacheKeyPolicy
	SetInternalValue(val *GoogleComputeBackendServiceCdnPolicyCacheKeyPolicy)
	QueryStringBlacklist() *[]*string
	SetQueryStringBlacklist(val *[]*string)
	QueryStringBlacklistInput() *[]*string
	QueryStringWhitelist() *[]*string
	SetQueryStringWhitelist(val *[]*string)
	QueryStringWhitelistInput() *[]*string
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
	ResetIncludeHost()
	ResetIncludeHttpHeaders()
	ResetIncludeNamedCookies()
	ResetIncludeProtocol()
	ResetIncludeQueryString()
	ResetQueryStringBlacklist()
	ResetQueryStringWhitelist()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference
type jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeHost() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeHostInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeHttpHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeHttpHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeHttpHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeHttpHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeNamedCookies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeNamedCookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeNamedCookiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeNamedCookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeProtocol() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeProtocolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeQueryString() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeQueryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeQueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeQueryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) InternalValue() *GoogleComputeBackendServiceCdnPolicyCacheKeyPolicy {
	var returns *GoogleComputeBackendServiceCdnPolicyCacheKeyPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) QueryStringBlacklist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringBlacklist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) QueryStringBlacklistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringBlacklistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) QueryStringWhitelist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringWhitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) QueryStringWhitelistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringWhitelistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeBackendService.GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference_Override(g GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeBackendService.GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetIncludeHost(val interface{}) {
	if err := j.validateSetIncludeHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeHost",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetIncludeHttpHeaders(val *[]*string) {
	if err := j.validateSetIncludeHttpHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeHttpHeaders",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetIncludeNamedCookies(val *[]*string) {
	if err := j.validateSetIncludeNamedCookiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeNamedCookies",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetIncludeProtocol(val interface{}) {
	if err := j.validateSetIncludeProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeProtocol",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetIncludeQueryString(val interface{}) {
	if err := j.validateSetIncludeQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeQueryString",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetInternalValue(val *GoogleComputeBackendServiceCdnPolicyCacheKeyPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetQueryStringBlacklist(val *[]*string) {
	if err := j.validateSetQueryStringBlacklistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringBlacklist",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetQueryStringWhitelist(val *[]*string) {
	if err := j.validateSetQueryStringWhitelistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringWhitelist",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ResetIncludeHost() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeHost",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ResetIncludeHttpHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeHttpHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ResetIncludeNamedCookies() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeNamedCookies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ResetIncludeProtocol() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeProtocol",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ResetIncludeQueryString() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeQueryString",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ResetQueryStringBlacklist() {
	_jsii_.InvokeVoid(
		g,
		"resetQueryStringBlacklist",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ResetQueryStringWhitelist() {
	_jsii_.InvokeVoid(
		g,
		"resetQueryStringWhitelist",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


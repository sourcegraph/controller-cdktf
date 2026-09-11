package googlecomputeregionbackendservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeregionbackendservice/internal"
)

type GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference interface {
	cdktn.ComplexObject
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
	IncludeNamedCookies() *[]*string
	SetIncludeNamedCookies(val *[]*string)
	IncludeNamedCookiesInput() *[]*string
	IncludeProtocol() interface{}
	SetIncludeProtocol(val interface{})
	IncludeProtocolInput() interface{}
	IncludeQueryString() interface{}
	SetIncludeQueryString(val interface{})
	IncludeQueryStringInput() interface{}
	InternalValue() *GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicy
	SetInternalValue(val *GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicy)
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
	ResetIncludeHost()
	ResetIncludeNamedCookies()
	ResetIncludeProtocol()
	ResetIncludeQueryString()
	ResetQueryStringBlacklist()
	ResetQueryStringWhitelist()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference
type jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeHost() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeHostInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeNamedCookies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeNamedCookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeNamedCookiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeNamedCookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeProtocol() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeProtocolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeQueryString() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeQueryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeQueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeQueryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) InternalValue() *GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicy {
	var returns *GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) QueryStringBlacklist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringBlacklist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) QueryStringBlacklistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringBlacklistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) QueryStringWhitelist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringWhitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) QueryStringWhitelistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringWhitelistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeRegionBackendService.GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference_Override(g GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeRegionBackendService.GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetIncludeHost(val interface{}) {
	if err := j.validateSetIncludeHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeHost",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetIncludeNamedCookies(val *[]*string) {
	if err := j.validateSetIncludeNamedCookiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeNamedCookies",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetIncludeProtocol(val interface{}) {
	if err := j.validateSetIncludeProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeProtocol",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetIncludeQueryString(val interface{}) {
	if err := j.validateSetIncludeQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeQueryString",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetInternalValue(val *GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetQueryStringBlacklist(val *[]*string) {
	if err := j.validateSetQueryStringBlacklistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringBlacklist",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetQueryStringWhitelist(val *[]*string) {
	if err := j.validateSetQueryStringWhitelistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringWhitelist",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ResetIncludeHost() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeHost",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ResetIncludeNamedCookies() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeNamedCookies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ResetIncludeProtocol() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeProtocol",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ResetIncludeQueryString() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeQueryString",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ResetQueryStringBlacklist() {
	_jsii_.InvokeVoid(
		g,
		"resetQueryStringBlacklist",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ResetQueryStringWhitelist() {
	_jsii_.InvokeVoid(
		g,
		"resetQueryStringWhitelist",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


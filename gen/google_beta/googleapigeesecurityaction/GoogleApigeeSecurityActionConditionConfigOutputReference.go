package googleapigeesecurityaction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleapigeesecurityaction/internal"
)

type GoogleApigeeSecurityActionConditionConfigOutputReference interface {
	cdktn.ComplexObject
	AccessTokens() *[]*string
	SetAccessTokens(val *[]*string)
	AccessTokensInput() *[]*string
	ApiKeys() *[]*string
	SetApiKeys(val *[]*string)
	ApiKeysInput() *[]*string
	ApiProducts() *[]*string
	SetApiProducts(val *[]*string)
	ApiProductsInput() *[]*string
	Asns() *[]*string
	SetAsns(val *[]*string)
	AsnsInput() *[]*string
	BotReasons() *[]*string
	SetBotReasons(val *[]*string)
	BotReasonsInput() *[]*string
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
	DeveloperApps() *[]*string
	SetDeveloperApps(val *[]*string)
	DeveloperAppsInput() *[]*string
	Developers() *[]*string
	SetDevelopers(val *[]*string)
	DevelopersInput() *[]*string
	// Experimental.
	Fqn() *string
	HttpMethods() *[]*string
	SetHttpMethods(val *[]*string)
	HttpMethodsInput() *[]*string
	InternalValue() *GoogleApigeeSecurityActionConditionConfig
	SetInternalValue(val *GoogleApigeeSecurityActionConditionConfig)
	IpAddressRanges() *[]*string
	SetIpAddressRanges(val *[]*string)
	IpAddressRangesInput() *[]*string
	RegionCodes() *[]*string
	SetRegionCodes(val *[]*string)
	RegionCodesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserAgents() *[]*string
	SetUserAgents(val *[]*string)
	UserAgentsInput() *[]*string
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
	ResetAccessTokens()
	ResetApiKeys()
	ResetApiProducts()
	ResetAsns()
	ResetBotReasons()
	ResetDeveloperApps()
	ResetDevelopers()
	ResetHttpMethods()
	ResetIpAddressRanges()
	ResetRegionCodes()
	ResetUserAgents()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleApigeeSecurityActionConditionConfigOutputReference
type jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) AccessTokens() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accessTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) AccessTokensInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accessTokensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) ApiKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) ApiKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) ApiProducts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiProducts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) ApiProductsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiProductsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) Asns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"asns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) AsnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"asnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) BotReasons() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"botReasons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) BotReasonsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"botReasonsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) DeveloperApps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"developerApps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) DeveloperAppsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"developerAppsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) Developers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"developers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) DevelopersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"developersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) HttpMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"httpMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) HttpMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"httpMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) InternalValue() *GoogleApigeeSecurityActionConditionConfig {
	var returns *GoogleApigeeSecurityActionConditionConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) IpAddressRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddressRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) IpAddressRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddressRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) RegionCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) RegionCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) UserAgents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userAgents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) UserAgentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userAgentsInput",
		&returns,
	)
	return returns
}


func NewGoogleApigeeSecurityActionConditionConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleApigeeSecurityActionConditionConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleApigeeSecurityActionConditionConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleApigeeSecurityAction.GoogleApigeeSecurityActionConditionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleApigeeSecurityActionConditionConfigOutputReference_Override(g GoogleApigeeSecurityActionConditionConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleApigeeSecurityAction.GoogleApigeeSecurityActionConditionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference)SetAccessTokens(val *[]*string) {
	if err := j.validateSetAccessTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessTokens",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference)SetApiKeys(val *[]*string) {
	if err := j.validateSetApiKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiKeys",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference)SetApiProducts(val *[]*string) {
	if err := j.validateSetApiProductsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiProducts",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference)SetAsns(val *[]*string) {
	if err := j.validateSetAsnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asns",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference)SetBotReasons(val *[]*string) {
	if err := j.validateSetBotReasonsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"botReasons",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference)SetDeveloperApps(val *[]*string) {
	if err := j.validateSetDeveloperAppsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"developerApps",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference)SetDevelopers(val *[]*string) {
	if err := j.validateSetDevelopersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"developers",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference)SetHttpMethods(val *[]*string) {
	if err := j.validateSetHttpMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethods",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference)SetInternalValue(val *GoogleApigeeSecurityActionConditionConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference)SetIpAddressRanges(val *[]*string) {
	if err := j.validateSetIpAddressRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressRanges",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference)SetRegionCodes(val *[]*string) {
	if err := j.validateSetRegionCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionCodes",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference)SetUserAgents(val *[]*string) {
	if err := j.validateSetUserAgentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userAgents",
		val,
	)
}

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) ResetAccessTokens() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessTokens",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) ResetApiKeys() {
	_jsii_.InvokeVoid(
		g,
		"resetApiKeys",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) ResetApiProducts() {
	_jsii_.InvokeVoid(
		g,
		"resetApiProducts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) ResetAsns() {
	_jsii_.InvokeVoid(
		g,
		"resetAsns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) ResetBotReasons() {
	_jsii_.InvokeVoid(
		g,
		"resetBotReasons",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) ResetDeveloperApps() {
	_jsii_.InvokeVoid(
		g,
		"resetDeveloperApps",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) ResetDevelopers() {
	_jsii_.InvokeVoid(
		g,
		"resetDevelopers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) ResetHttpMethods() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpMethods",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) ResetIpAddressRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetIpAddressRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) ResetRegionCodes() {
	_jsii_.InvokeVoid(
		g,
		"resetRegionCodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) ResetUserAgents() {
	_jsii_.InvokeVoid(
		g,
		"resetUserAgents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleApigeeSecurityActionConditionConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


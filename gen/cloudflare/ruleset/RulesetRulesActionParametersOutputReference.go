package ruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/ruleset/internal"
)

type RulesetRulesActionParametersOutputReference interface {
	cdktf.ComplexObject
	AutomaticHttpsRewrites() interface{}
	SetAutomaticHttpsRewrites(val interface{})
	AutomaticHttpsRewritesInput() interface{}
	Autominify() RulesetRulesActionParametersAutominifyList
	AutominifyInput() interface{}
	Bic() interface{}
	SetBic(val interface{})
	BicInput() interface{}
	BrowserTtl() RulesetRulesActionParametersBrowserTtlList
	BrowserTtlInput() interface{}
	Cache() interface{}
	SetCache(val interface{})
	CacheInput() interface{}
	CacheKey() RulesetRulesActionParametersCacheKeyList
	CacheKeyInput() interface{}
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
	Content() *string
	SetContent(val *string)
	ContentInput() *string
	ContentType() *string
	SetContentType(val *string)
	ContentTypeInput() *string
	CookieFields() *[]*string
	SetCookieFields(val *[]*string)
	CookieFieldsInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisableApps() interface{}
	SetDisableApps(val interface{})
	DisableAppsInput() interface{}
	DisableRailgun() interface{}
	SetDisableRailgun(val interface{})
	DisableRailgunInput() interface{}
	DisableZaraz() interface{}
	SetDisableZaraz(val interface{})
	DisableZarazInput() interface{}
	EdgeTtl() RulesetRulesActionParametersEdgeTtlList
	EdgeTtlInput() interface{}
	EmailObfuscation() interface{}
	SetEmailObfuscation(val interface{})
	EmailObfuscationInput() interface{}
	// Experimental.
	Fqn() *string
	FromList() RulesetRulesActionParametersFromListStructList
	FromListInput() interface{}
	FromValue() RulesetRulesActionParametersFromValueList
	FromValueInput() interface{}
	Headers() RulesetRulesActionParametersHeadersList
	HeadersInput() interface{}
	HostHeader() *string
	SetHostHeader(val *string)
	HostHeaderInput() *string
	HotlinkProtection() interface{}
	SetHotlinkProtection(val interface{})
	HotlinkProtectionInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	Increment() *float64
	SetIncrement(val *float64)
	IncrementInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MatchedData() RulesetRulesActionParametersMatchedDataList
	MatchedDataInput() interface{}
	Mirage() interface{}
	SetMirage(val interface{})
	MirageInput() interface{}
	OpportunisticEncryption() interface{}
	SetOpportunisticEncryption(val interface{})
	OpportunisticEncryptionInput() interface{}
	Origin() RulesetRulesActionParametersOriginList
	OriginErrorPagePassthru() interface{}
	SetOriginErrorPagePassthru(val interface{})
	OriginErrorPagePassthruInput() interface{}
	OriginInput() interface{}
	Overrides() RulesetRulesActionParametersOverridesList
	OverridesInput() interface{}
	Phases() *[]*string
	SetPhases(val *[]*string)
	PhasesInput() *[]*string
	Polish() *string
	SetPolish(val *string)
	PolishInput() *string
	Products() *[]*string
	SetProducts(val *[]*string)
	ProductsInput() *[]*string
	RequestFields() *[]*string
	SetRequestFields(val *[]*string)
	RequestFieldsInput() *[]*string
	RespectStrongEtags() interface{}
	SetRespectStrongEtags(val interface{})
	RespectStrongEtagsInput() interface{}
	Response() RulesetRulesActionParametersResponseList
	ResponseFields() *[]*string
	SetResponseFields(val *[]*string)
	ResponseFieldsInput() *[]*string
	ResponseInput() interface{}
	RocketLoader() interface{}
	SetRocketLoader(val interface{})
	RocketLoaderInput() interface{}
	Rules() *map[string]*string
	SetRules(val *map[string]*string)
	Ruleset() *string
	SetRuleset(val *string)
	RulesetInput() *string
	Rulesets() *[]*string
	SetRulesets(val *[]*string)
	RulesetsInput() *[]*string
	RulesInput() *map[string]*string
	SecurityLevel() *string
	SetSecurityLevel(val *string)
	SecurityLevelInput() *string
	ServerSideExcludes() interface{}
	SetServerSideExcludes(val interface{})
	ServerSideExcludesInput() interface{}
	ServeStale() RulesetRulesActionParametersServeStaleList
	ServeStaleInput() interface{}
	Sni() RulesetRulesActionParametersSniList
	SniInput() interface{}
	Ssl() *string
	SetSsl(val *string)
	SslInput() *string
	StatusCode() *float64
	SetStatusCode(val *float64)
	StatusCodeInput() *float64
	Sxg() interface{}
	SetSxg(val interface{})
	SxgInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uri() RulesetRulesActionParametersUriList
	UriInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutAutominify(value interface{})
	PutBrowserTtl(value interface{})
	PutCacheKey(value interface{})
	PutEdgeTtl(value interface{})
	PutFromList(value interface{})
	PutFromValue(value interface{})
	PutHeaders(value interface{})
	PutMatchedData(value interface{})
	PutOrigin(value interface{})
	PutOverrides(value interface{})
	PutResponse(value interface{})
	PutServeStale(value interface{})
	PutSni(value interface{})
	PutUri(value interface{})
	ResetAutomaticHttpsRewrites()
	ResetAutominify()
	ResetBic()
	ResetBrowserTtl()
	ResetCache()
	ResetCacheKey()
	ResetContent()
	ResetContentType()
	ResetCookieFields()
	ResetDisableApps()
	ResetDisableRailgun()
	ResetDisableZaraz()
	ResetEdgeTtl()
	ResetEmailObfuscation()
	ResetFromList()
	ResetFromValue()
	ResetHeaders()
	ResetHostHeader()
	ResetHotlinkProtection()
	ResetId()
	ResetIncrement()
	ResetMatchedData()
	ResetMirage()
	ResetOpportunisticEncryption()
	ResetOrigin()
	ResetOriginErrorPagePassthru()
	ResetOverrides()
	ResetPhases()
	ResetPolish()
	ResetProducts()
	ResetRequestFields()
	ResetRespectStrongEtags()
	ResetResponse()
	ResetResponseFields()
	ResetRocketLoader()
	ResetRules()
	ResetRuleset()
	ResetRulesets()
	ResetSecurityLevel()
	ResetServerSideExcludes()
	ResetServeStale()
	ResetSni()
	ResetSsl()
	ResetStatusCode()
	ResetSxg()
	ResetUri()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RulesetRulesActionParametersOutputReference
type jsiiProxy_RulesetRulesActionParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) AutomaticHttpsRewrites() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticHttpsRewrites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) AutomaticHttpsRewritesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticHttpsRewritesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Autominify() RulesetRulesActionParametersAutominifyList {
	var returns RulesetRulesActionParametersAutominifyList
	_jsii_.Get(
		j,
		"autominify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) AutominifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autominifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Bic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) BicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) BrowserTtl() RulesetRulesActionParametersBrowserTtlList {
	var returns RulesetRulesActionParametersBrowserTtlList
	_jsii_.Get(
		j,
		"browserTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) BrowserTtlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"browserTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Cache() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) CacheInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) CacheKey() RulesetRulesActionParametersCacheKeyList {
	var returns RulesetRulesActionParametersCacheKeyList
	_jsii_.Get(
		j,
		"cacheKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) CacheKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) CookieFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cookieFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) CookieFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cookieFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) DisableApps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) DisableAppsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAppsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) DisableRailgun() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRailgun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) DisableRailgunInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRailgunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) DisableZaraz() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableZaraz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) DisableZarazInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableZarazInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) EdgeTtl() RulesetRulesActionParametersEdgeTtlList {
	var returns RulesetRulesActionParametersEdgeTtlList
	_jsii_.Get(
		j,
		"edgeTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) EdgeTtlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"edgeTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) EmailObfuscation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailObfuscation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) EmailObfuscationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailObfuscationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) FromList() RulesetRulesActionParametersFromListStructList {
	var returns RulesetRulesActionParametersFromListStructList
	_jsii_.Get(
		j,
		"fromList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) FromListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fromListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) FromValue() RulesetRulesActionParametersFromValueList {
	var returns RulesetRulesActionParametersFromValueList
	_jsii_.Get(
		j,
		"fromValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) FromValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fromValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Headers() RulesetRulesActionParametersHeadersList {
	var returns RulesetRulesActionParametersHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) HostHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) HostHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) HotlinkProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hotlinkProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) HotlinkProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hotlinkProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Increment() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"increment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) IncrementInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"incrementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) MatchedData() RulesetRulesActionParametersMatchedDataList {
	var returns RulesetRulesActionParametersMatchedDataList
	_jsii_.Get(
		j,
		"matchedData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) MatchedDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchedDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Mirage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mirage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) MirageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mirageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) OpportunisticEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opportunisticEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) OpportunisticEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opportunisticEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Origin() RulesetRulesActionParametersOriginList {
	var returns RulesetRulesActionParametersOriginList
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) OriginErrorPagePassthru() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originErrorPagePassthru",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) OriginErrorPagePassthruInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originErrorPagePassthruInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) OriginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Overrides() RulesetRulesActionParametersOverridesList {
	var returns RulesetRulesActionParametersOverridesList
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) OverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Phases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"phases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) PhasesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"phasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Polish() *string {
	var returns *string
	_jsii_.Get(
		j,
		"polish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) PolishInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"polishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Products() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"products",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ProductsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"productsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) RequestFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) RequestFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) RespectStrongEtags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"respectStrongEtags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) RespectStrongEtagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"respectStrongEtagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Response() RulesetRulesActionParametersResponseList {
	var returns RulesetRulesActionParametersResponseList
	_jsii_.Get(
		j,
		"response",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ResponseFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"responseFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ResponseFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"responseFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) RocketLoader() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rocketLoader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) RocketLoaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rocketLoaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Rules() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Ruleset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) RulesetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rulesetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Rulesets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rulesets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) RulesetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rulesetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) RulesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"rulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SecurityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SecurityLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ServerSideExcludes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverSideExcludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ServerSideExcludesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverSideExcludesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ServeStale() RulesetRulesActionParametersServeStaleList {
	var returns RulesetRulesActionParametersServeStaleList
	_jsii_.Get(
		j,
		"serveStale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) ServeStaleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serveStaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Sni() RulesetRulesActionParametersSniList {
	var returns RulesetRulesActionParametersSniList
	_jsii_.Get(
		j,
		"sni",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SniInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sniInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Ssl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SslInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) StatusCode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statusCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) StatusCodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statusCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Sxg() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sxg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) SxgInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sxgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Uri() RulesetRulesActionParametersUriList {
	var returns RulesetRulesActionParametersUriList
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) UriInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewRulesetRulesActionParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) RulesetRulesActionParametersOutputReference {
	_init_.Initialize()

	if err := validateNewRulesetRulesActionParametersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_RulesetRulesActionParametersOutputReference{}

	_jsii_.Create(
		"cloudflare.ruleset.RulesetRulesActionParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewRulesetRulesActionParametersOutputReference_Override(r RulesetRulesActionParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.ruleset.RulesetRulesActionParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetAutomaticHttpsRewrites(val interface{}) {
	if err := j.validateSetAutomaticHttpsRewritesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticHttpsRewrites",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetBic(val interface{}) {
	if err := j.validateSetBicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bic",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetCache(val interface{}) {
	if err := j.validateSetCacheParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cache",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetContent(val *string) {
	if err := j.validateSetContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetContentType(val *string) {
	if err := j.validateSetContentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetCookieFields(val *[]*string) {
	if err := j.validateSetCookieFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cookieFields",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetDisableApps(val interface{}) {
	if err := j.validateSetDisableAppsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableApps",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetDisableRailgun(val interface{}) {
	if err := j.validateSetDisableRailgunParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableRailgun",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetDisableZaraz(val interface{}) {
	if err := j.validateSetDisableZarazParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableZaraz",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetEmailObfuscation(val interface{}) {
	if err := j.validateSetEmailObfuscationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailObfuscation",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetHostHeader(val *string) {
	if err := j.validateSetHostHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostHeader",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetHotlinkProtection(val interface{}) {
	if err := j.validateSetHotlinkProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hotlinkProtection",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetIncrement(val *float64) {
	if err := j.validateSetIncrementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"increment",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetMirage(val interface{}) {
	if err := j.validateSetMirageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mirage",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetOpportunisticEncryption(val interface{}) {
	if err := j.validateSetOpportunisticEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"opportunisticEncryption",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetOriginErrorPagePassthru(val interface{}) {
	if err := j.validateSetOriginErrorPagePassthruParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originErrorPagePassthru",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetPhases(val *[]*string) {
	if err := j.validateSetPhasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phases",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetPolish(val *string) {
	if err := j.validateSetPolishParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"polish",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetProducts(val *[]*string) {
	if err := j.validateSetProductsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"products",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetRequestFields(val *[]*string) {
	if err := j.validateSetRequestFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestFields",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetRespectStrongEtags(val interface{}) {
	if err := j.validateSetRespectStrongEtagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"respectStrongEtags",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetResponseFields(val *[]*string) {
	if err := j.validateSetResponseFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseFields",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetRocketLoader(val interface{}) {
	if err := j.validateSetRocketLoaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rocketLoader",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetRules(val *map[string]*string) {
	if err := j.validateSetRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetRuleset(val *string) {
	if err := j.validateSetRulesetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleset",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetRulesets(val *[]*string) {
	if err := j.validateSetRulesetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rulesets",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetSecurityLevel(val *string) {
	if err := j.validateSetSecurityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityLevel",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetServerSideExcludes(val interface{}) {
	if err := j.validateSetServerSideExcludesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverSideExcludes",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetSsl(val *string) {
	if err := j.validateSetSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssl",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetStatusCode(val *float64) {
	if err := j.validateSetStatusCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statusCode",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetSxg(val interface{}) {
	if err := j.validateSetSxgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sxg",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_RulesetRulesActionParametersOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutAutominify(value interface{}) {
	if err := r.validatePutAutominifyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAutominify",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutBrowserTtl(value interface{}) {
	if err := r.validatePutBrowserTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putBrowserTtl",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutCacheKey(value interface{}) {
	if err := r.validatePutCacheKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putCacheKey",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutEdgeTtl(value interface{}) {
	if err := r.validatePutEdgeTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putEdgeTtl",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutFromList(value interface{}) {
	if err := r.validatePutFromListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putFromList",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutFromValue(value interface{}) {
	if err := r.validatePutFromValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putFromValue",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutHeaders(value interface{}) {
	if err := r.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putHeaders",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutMatchedData(value interface{}) {
	if err := r.validatePutMatchedDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putMatchedData",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutOrigin(value interface{}) {
	if err := r.validatePutOriginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putOrigin",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutOverrides(value interface{}) {
	if err := r.validatePutOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putOverrides",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutResponse(value interface{}) {
	if err := r.validatePutResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putResponse",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutServeStale(value interface{}) {
	if err := r.validatePutServeStaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putServeStale",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutSni(value interface{}) {
	if err := r.validatePutSniParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSni",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) PutUri(value interface{}) {
	if err := r.validatePutUriParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putUri",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetAutomaticHttpsRewrites() {
	_jsii_.InvokeVoid(
		r,
		"resetAutomaticHttpsRewrites",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetAutominify() {
	_jsii_.InvokeVoid(
		r,
		"resetAutominify",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetBic() {
	_jsii_.InvokeVoid(
		r,
		"resetBic",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetBrowserTtl() {
	_jsii_.InvokeVoid(
		r,
		"resetBrowserTtl",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetCache() {
	_jsii_.InvokeVoid(
		r,
		"resetCache",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetCacheKey() {
	_jsii_.InvokeVoid(
		r,
		"resetCacheKey",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetContent() {
	_jsii_.InvokeVoid(
		r,
		"resetContent",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetContentType() {
	_jsii_.InvokeVoid(
		r,
		"resetContentType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetCookieFields() {
	_jsii_.InvokeVoid(
		r,
		"resetCookieFields",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetDisableApps() {
	_jsii_.InvokeVoid(
		r,
		"resetDisableApps",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetDisableRailgun() {
	_jsii_.InvokeVoid(
		r,
		"resetDisableRailgun",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetDisableZaraz() {
	_jsii_.InvokeVoid(
		r,
		"resetDisableZaraz",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetEdgeTtl() {
	_jsii_.InvokeVoid(
		r,
		"resetEdgeTtl",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetEmailObfuscation() {
	_jsii_.InvokeVoid(
		r,
		"resetEmailObfuscation",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetFromList() {
	_jsii_.InvokeVoid(
		r,
		"resetFromList",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetFromValue() {
	_jsii_.InvokeVoid(
		r,
		"resetFromValue",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		r,
		"resetHeaders",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetHostHeader() {
	_jsii_.InvokeVoid(
		r,
		"resetHostHeader",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetHotlinkProtection() {
	_jsii_.InvokeVoid(
		r,
		"resetHotlinkProtection",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetIncrement() {
	_jsii_.InvokeVoid(
		r,
		"resetIncrement",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetMatchedData() {
	_jsii_.InvokeVoid(
		r,
		"resetMatchedData",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetMirage() {
	_jsii_.InvokeVoid(
		r,
		"resetMirage",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetOpportunisticEncryption() {
	_jsii_.InvokeVoid(
		r,
		"resetOpportunisticEncryption",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetOrigin() {
	_jsii_.InvokeVoid(
		r,
		"resetOrigin",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetOriginErrorPagePassthru() {
	_jsii_.InvokeVoid(
		r,
		"resetOriginErrorPagePassthru",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetOverrides() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrides",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetPhases() {
	_jsii_.InvokeVoid(
		r,
		"resetPhases",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetPolish() {
	_jsii_.InvokeVoid(
		r,
		"resetPolish",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetProducts() {
	_jsii_.InvokeVoid(
		r,
		"resetProducts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetRequestFields() {
	_jsii_.InvokeVoid(
		r,
		"resetRequestFields",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetRespectStrongEtags() {
	_jsii_.InvokeVoid(
		r,
		"resetRespectStrongEtags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetResponse() {
	_jsii_.InvokeVoid(
		r,
		"resetResponse",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetResponseFields() {
	_jsii_.InvokeVoid(
		r,
		"resetResponseFields",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetRocketLoader() {
	_jsii_.InvokeVoid(
		r,
		"resetRocketLoader",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetRules() {
	_jsii_.InvokeVoid(
		r,
		"resetRules",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetRuleset() {
	_jsii_.InvokeVoid(
		r,
		"resetRuleset",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetRulesets() {
	_jsii_.InvokeVoid(
		r,
		"resetRulesets",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetSecurityLevel() {
	_jsii_.InvokeVoid(
		r,
		"resetSecurityLevel",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetServerSideExcludes() {
	_jsii_.InvokeVoid(
		r,
		"resetServerSideExcludes",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetServeStale() {
	_jsii_.InvokeVoid(
		r,
		"resetServeStale",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetSni() {
	_jsii_.InvokeVoid(
		r,
		"resetSni",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetSsl() {
	_jsii_.InvokeVoid(
		r,
		"resetSsl",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetStatusCode() {
	_jsii_.InvokeVoid(
		r,
		"resetStatusCode",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetSxg() {
	_jsii_.InvokeVoid(
		r,
		"resetSxg",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetUri() {
	_jsii_.InvokeVoid(
		r,
		"resetUri",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		r,
		"resetVersion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRulesActionParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package pagerule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/pagerule/internal"
)

type PageRuleActionsOutputReference interface {
	cdktf.ComplexObject
	AlwaysOnline() *string
	SetAlwaysOnline(val *string)
	AlwaysOnlineInput() *string
	AlwaysUseHttps() interface{}
	SetAlwaysUseHttps(val interface{})
	AlwaysUseHttpsInput() interface{}
	AutomaticHttpsRewrites() *string
	SetAutomaticHttpsRewrites(val *string)
	AutomaticHttpsRewritesInput() *string
	BrowserCacheTtl() *string
	SetBrowserCacheTtl(val *string)
	BrowserCacheTtlInput() *string
	BrowserCheck() *string
	SetBrowserCheck(val *string)
	BrowserCheckInput() *string
	BypassCacheOnCookie() *string
	SetBypassCacheOnCookie(val *string)
	BypassCacheOnCookieInput() *string
	CacheByDeviceType() *string
	SetCacheByDeviceType(val *string)
	CacheByDeviceTypeInput() *string
	CacheDeceptionArmor() *string
	SetCacheDeceptionArmor(val *string)
	CacheDeceptionArmorInput() *string
	CacheKeyFields() PageRuleActionsCacheKeyFieldsOutputReference
	CacheKeyFieldsInput() *PageRuleActionsCacheKeyFields
	CacheLevel() *string
	SetCacheLevel(val *string)
	CacheLevelInput() *string
	CacheOnCookie() *string
	SetCacheOnCookie(val *string)
	CacheOnCookieInput() *string
	CacheTtlByStatus() PageRuleActionsCacheTtlByStatusList
	CacheTtlByStatusInput() interface{}
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
	DisableApps() interface{}
	SetDisableApps(val interface{})
	DisableAppsInput() interface{}
	DisablePerformance() interface{}
	SetDisablePerformance(val interface{})
	DisablePerformanceInput() interface{}
	DisableRailgun() interface{}
	SetDisableRailgun(val interface{})
	DisableRailgunInput() interface{}
	DisableSecurity() interface{}
	SetDisableSecurity(val interface{})
	DisableSecurityInput() interface{}
	DisableZaraz() interface{}
	SetDisableZaraz(val interface{})
	DisableZarazInput() interface{}
	EdgeCacheTtl() *float64
	SetEdgeCacheTtl(val *float64)
	EdgeCacheTtlInput() *float64
	EmailObfuscation() *string
	SetEmailObfuscation(val *string)
	EmailObfuscationInput() *string
	ExplicitCacheControl() *string
	SetExplicitCacheControl(val *string)
	ExplicitCacheControlInput() *string
	ForwardingUrl() PageRuleActionsForwardingUrlOutputReference
	ForwardingUrlInput() *PageRuleActionsForwardingUrl
	// Experimental.
	Fqn() *string
	HostHeaderOverride() *string
	SetHostHeaderOverride(val *string)
	HostHeaderOverrideInput() *string
	InternalValue() *PageRuleActions
	SetInternalValue(val *PageRuleActions)
	IpGeolocation() *string
	SetIpGeolocation(val *string)
	IpGeolocationInput() *string
	Minify() PageRuleActionsMinifyList
	MinifyInput() interface{}
	Mirage() *string
	SetMirage(val *string)
	MirageInput() *string
	OpportunisticEncryption() *string
	SetOpportunisticEncryption(val *string)
	OpportunisticEncryptionInput() *string
	OriginErrorPagePassThru() *string
	SetOriginErrorPagePassThru(val *string)
	OriginErrorPagePassThruInput() *string
	Polish() *string
	SetPolish(val *string)
	PolishInput() *string
	ResolveOverride() *string
	SetResolveOverride(val *string)
	ResolveOverrideInput() *string
	RespectStrongEtag() *string
	SetRespectStrongEtag(val *string)
	RespectStrongEtagInput() *string
	ResponseBuffering() *string
	SetResponseBuffering(val *string)
	ResponseBufferingInput() *string
	RocketLoader() *string
	SetRocketLoader(val *string)
	RocketLoaderInput() *string
	SecurityLevel() *string
	SetSecurityLevel(val *string)
	SecurityLevelInput() *string
	ServerSideExclude() *string
	SetServerSideExclude(val *string)
	ServerSideExcludeInput() *string
	SortQueryStringForCache() *string
	SetSortQueryStringForCache(val *string)
	SortQueryStringForCacheInput() *string
	Ssl() *string
	SetSsl(val *string)
	SslInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrueClientIpHeader() *string
	SetTrueClientIpHeader(val *string)
	TrueClientIpHeaderInput() *string
	Waf() *string
	SetWaf(val *string)
	WafInput() *string
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
	PutCacheKeyFields(value *PageRuleActionsCacheKeyFields)
	PutCacheTtlByStatus(value interface{})
	PutForwardingUrl(value *PageRuleActionsForwardingUrl)
	PutMinify(value interface{})
	ResetAlwaysOnline()
	ResetAlwaysUseHttps()
	ResetAutomaticHttpsRewrites()
	ResetBrowserCacheTtl()
	ResetBrowserCheck()
	ResetBypassCacheOnCookie()
	ResetCacheByDeviceType()
	ResetCacheDeceptionArmor()
	ResetCacheKeyFields()
	ResetCacheLevel()
	ResetCacheOnCookie()
	ResetCacheTtlByStatus()
	ResetDisableApps()
	ResetDisablePerformance()
	ResetDisableRailgun()
	ResetDisableSecurity()
	ResetDisableZaraz()
	ResetEdgeCacheTtl()
	ResetEmailObfuscation()
	ResetExplicitCacheControl()
	ResetForwardingUrl()
	ResetHostHeaderOverride()
	ResetIpGeolocation()
	ResetMinify()
	ResetMirage()
	ResetOpportunisticEncryption()
	ResetOriginErrorPagePassThru()
	ResetPolish()
	ResetResolveOverride()
	ResetRespectStrongEtag()
	ResetResponseBuffering()
	ResetRocketLoader()
	ResetSecurityLevel()
	ResetServerSideExclude()
	ResetSortQueryStringForCache()
	ResetSsl()
	ResetTrueClientIpHeader()
	ResetWaf()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PageRuleActionsOutputReference
type jsiiProxy_PageRuleActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PageRuleActionsOutputReference) AlwaysOnline() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alwaysOnline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) AlwaysOnlineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alwaysOnlineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) AlwaysUseHttps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysUseHttps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) AlwaysUseHttpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysUseHttpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) AutomaticHttpsRewrites() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticHttpsRewrites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) AutomaticHttpsRewritesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticHttpsRewritesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) BrowserCacheTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"browserCacheTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) BrowserCacheTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"browserCacheTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) BrowserCheck() *string {
	var returns *string
	_jsii_.Get(
		j,
		"browserCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) BrowserCheckInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"browserCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) BypassCacheOnCookie() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bypassCacheOnCookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) BypassCacheOnCookieInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bypassCacheOnCookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) CacheByDeviceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheByDeviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) CacheByDeviceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheByDeviceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) CacheDeceptionArmor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheDeceptionArmor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) CacheDeceptionArmorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheDeceptionArmorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) CacheKeyFields() PageRuleActionsCacheKeyFieldsOutputReference {
	var returns PageRuleActionsCacheKeyFieldsOutputReference
	_jsii_.Get(
		j,
		"cacheKeyFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) CacheKeyFieldsInput() *PageRuleActionsCacheKeyFields {
	var returns *PageRuleActionsCacheKeyFields
	_jsii_.Get(
		j,
		"cacheKeyFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) CacheLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) CacheLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) CacheOnCookie() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheOnCookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) CacheOnCookieInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheOnCookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) CacheTtlByStatus() PageRuleActionsCacheTtlByStatusList {
	var returns PageRuleActionsCacheTtlByStatusList
	_jsii_.Get(
		j,
		"cacheTtlByStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) CacheTtlByStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheTtlByStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) DisableApps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) DisableAppsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAppsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) DisablePerformance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePerformance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) DisablePerformanceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePerformanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) DisableRailgun() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRailgun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) DisableRailgunInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRailgunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) DisableSecurity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableSecurity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) DisableSecurityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableSecurityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) DisableZaraz() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableZaraz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) DisableZarazInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableZarazInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) EdgeCacheTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"edgeCacheTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) EdgeCacheTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"edgeCacheTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) EmailObfuscation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailObfuscation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) EmailObfuscationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailObfuscationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) ExplicitCacheControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"explicitCacheControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) ExplicitCacheControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"explicitCacheControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) ForwardingUrl() PageRuleActionsForwardingUrlOutputReference {
	var returns PageRuleActionsForwardingUrlOutputReference
	_jsii_.Get(
		j,
		"forwardingUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) ForwardingUrlInput() *PageRuleActionsForwardingUrl {
	var returns *PageRuleActionsForwardingUrl
	_jsii_.Get(
		j,
		"forwardingUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) HostHeaderOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostHeaderOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) HostHeaderOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostHeaderOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) InternalValue() *PageRuleActions {
	var returns *PageRuleActions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) IpGeolocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipGeolocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) IpGeolocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipGeolocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) Minify() PageRuleActionsMinifyList {
	var returns PageRuleActionsMinifyList
	_jsii_.Get(
		j,
		"minify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) MinifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"minifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) Mirage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mirage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) MirageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mirageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) OpportunisticEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opportunisticEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) OpportunisticEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opportunisticEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) OriginErrorPagePassThru() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originErrorPagePassThru",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) OriginErrorPagePassThruInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originErrorPagePassThruInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) Polish() *string {
	var returns *string
	_jsii_.Get(
		j,
		"polish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) PolishInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"polishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) ResolveOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolveOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) ResolveOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolveOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) RespectStrongEtag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"respectStrongEtag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) RespectStrongEtagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"respectStrongEtagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) ResponseBuffering() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseBuffering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) ResponseBufferingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseBufferingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) RocketLoader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rocketLoader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) RocketLoaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rocketLoaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) SecurityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) SecurityLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) ServerSideExclude() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideExclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) ServerSideExcludeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideExcludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) SortQueryStringForCache() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortQueryStringForCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) SortQueryStringForCacheInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortQueryStringForCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) Ssl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) SslInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) TrueClientIpHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trueClientIpHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) TrueClientIpHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trueClientIpHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) Waf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsOutputReference) WafInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wafInput",
		&returns,
	)
	return returns
}


func NewPageRuleActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PageRuleActionsOutputReference {
	_init_.Initialize()

	if err := validateNewPageRuleActionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PageRuleActionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.pageRule.PageRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPageRuleActionsOutputReference_Override(p PageRuleActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.pageRule.PageRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetAlwaysOnline(val *string) {
	if err := j.validateSetAlwaysOnlineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alwaysOnline",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetAlwaysUseHttps(val interface{}) {
	if err := j.validateSetAlwaysUseHttpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alwaysUseHttps",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetAutomaticHttpsRewrites(val *string) {
	if err := j.validateSetAutomaticHttpsRewritesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticHttpsRewrites",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetBrowserCacheTtl(val *string) {
	if err := j.validateSetBrowserCacheTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"browserCacheTtl",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetBrowserCheck(val *string) {
	if err := j.validateSetBrowserCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"browserCheck",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetBypassCacheOnCookie(val *string) {
	if err := j.validateSetBypassCacheOnCookieParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bypassCacheOnCookie",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetCacheByDeviceType(val *string) {
	if err := j.validateSetCacheByDeviceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheByDeviceType",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetCacheDeceptionArmor(val *string) {
	if err := j.validateSetCacheDeceptionArmorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheDeceptionArmor",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetCacheLevel(val *string) {
	if err := j.validateSetCacheLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheLevel",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetCacheOnCookie(val *string) {
	if err := j.validateSetCacheOnCookieParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheOnCookie",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetDisableApps(val interface{}) {
	if err := j.validateSetDisableAppsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableApps",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetDisablePerformance(val interface{}) {
	if err := j.validateSetDisablePerformanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disablePerformance",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetDisableRailgun(val interface{}) {
	if err := j.validateSetDisableRailgunParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableRailgun",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetDisableSecurity(val interface{}) {
	if err := j.validateSetDisableSecurityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableSecurity",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetDisableZaraz(val interface{}) {
	if err := j.validateSetDisableZarazParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableZaraz",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetEdgeCacheTtl(val *float64) {
	if err := j.validateSetEdgeCacheTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeCacheTtl",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetEmailObfuscation(val *string) {
	if err := j.validateSetEmailObfuscationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailObfuscation",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetExplicitCacheControl(val *string) {
	if err := j.validateSetExplicitCacheControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"explicitCacheControl",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetHostHeaderOverride(val *string) {
	if err := j.validateSetHostHeaderOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostHeaderOverride",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetInternalValue(val *PageRuleActions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetIpGeolocation(val *string) {
	if err := j.validateSetIpGeolocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipGeolocation",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetMirage(val *string) {
	if err := j.validateSetMirageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mirage",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetOpportunisticEncryption(val *string) {
	if err := j.validateSetOpportunisticEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"opportunisticEncryption",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetOriginErrorPagePassThru(val *string) {
	if err := j.validateSetOriginErrorPagePassThruParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originErrorPagePassThru",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetPolish(val *string) {
	if err := j.validateSetPolishParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"polish",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetResolveOverride(val *string) {
	if err := j.validateSetResolveOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resolveOverride",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetRespectStrongEtag(val *string) {
	if err := j.validateSetRespectStrongEtagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"respectStrongEtag",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetResponseBuffering(val *string) {
	if err := j.validateSetResponseBufferingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseBuffering",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetRocketLoader(val *string) {
	if err := j.validateSetRocketLoaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rocketLoader",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetSecurityLevel(val *string) {
	if err := j.validateSetSecurityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityLevel",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetServerSideExclude(val *string) {
	if err := j.validateSetServerSideExcludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverSideExclude",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetSortQueryStringForCache(val *string) {
	if err := j.validateSetSortQueryStringForCacheParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sortQueryStringForCache",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetSsl(val *string) {
	if err := j.validateSetSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssl",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetTrueClientIpHeader(val *string) {
	if err := j.validateSetTrueClientIpHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trueClientIpHeader",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsOutputReference)SetWaf(val *string) {
	if err := j.validateSetWafParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waf",
		val,
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsOutputReference) PutCacheKeyFields(value *PageRuleActionsCacheKeyFields) {
	if err := p.validatePutCacheKeyFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCacheKeyFields",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) PutCacheTtlByStatus(value interface{}) {
	if err := p.validatePutCacheTtlByStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCacheTtlByStatus",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) PutForwardingUrl(value *PageRuleActionsForwardingUrl) {
	if err := p.validatePutForwardingUrlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putForwardingUrl",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) PutMinify(value interface{}) {
	if err := p.validatePutMinifyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMinify",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetAlwaysOnline() {
	_jsii_.InvokeVoid(
		p,
		"resetAlwaysOnline",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetAlwaysUseHttps() {
	_jsii_.InvokeVoid(
		p,
		"resetAlwaysUseHttps",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetAutomaticHttpsRewrites() {
	_jsii_.InvokeVoid(
		p,
		"resetAutomaticHttpsRewrites",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetBrowserCacheTtl() {
	_jsii_.InvokeVoid(
		p,
		"resetBrowserCacheTtl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetBrowserCheck() {
	_jsii_.InvokeVoid(
		p,
		"resetBrowserCheck",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetBypassCacheOnCookie() {
	_jsii_.InvokeVoid(
		p,
		"resetBypassCacheOnCookie",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetCacheByDeviceType() {
	_jsii_.InvokeVoid(
		p,
		"resetCacheByDeviceType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetCacheDeceptionArmor() {
	_jsii_.InvokeVoid(
		p,
		"resetCacheDeceptionArmor",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetCacheKeyFields() {
	_jsii_.InvokeVoid(
		p,
		"resetCacheKeyFields",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetCacheLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetCacheLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetCacheOnCookie() {
	_jsii_.InvokeVoid(
		p,
		"resetCacheOnCookie",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetCacheTtlByStatus() {
	_jsii_.InvokeVoid(
		p,
		"resetCacheTtlByStatus",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetDisableApps() {
	_jsii_.InvokeVoid(
		p,
		"resetDisableApps",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetDisablePerformance() {
	_jsii_.InvokeVoid(
		p,
		"resetDisablePerformance",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetDisableRailgun() {
	_jsii_.InvokeVoid(
		p,
		"resetDisableRailgun",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetDisableSecurity() {
	_jsii_.InvokeVoid(
		p,
		"resetDisableSecurity",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetDisableZaraz() {
	_jsii_.InvokeVoid(
		p,
		"resetDisableZaraz",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetEdgeCacheTtl() {
	_jsii_.InvokeVoid(
		p,
		"resetEdgeCacheTtl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetEmailObfuscation() {
	_jsii_.InvokeVoid(
		p,
		"resetEmailObfuscation",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetExplicitCacheControl() {
	_jsii_.InvokeVoid(
		p,
		"resetExplicitCacheControl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetForwardingUrl() {
	_jsii_.InvokeVoid(
		p,
		"resetForwardingUrl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetHostHeaderOverride() {
	_jsii_.InvokeVoid(
		p,
		"resetHostHeaderOverride",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetIpGeolocation() {
	_jsii_.InvokeVoid(
		p,
		"resetIpGeolocation",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetMinify() {
	_jsii_.InvokeVoid(
		p,
		"resetMinify",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetMirage() {
	_jsii_.InvokeVoid(
		p,
		"resetMirage",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetOpportunisticEncryption() {
	_jsii_.InvokeVoid(
		p,
		"resetOpportunisticEncryption",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetOriginErrorPagePassThru() {
	_jsii_.InvokeVoid(
		p,
		"resetOriginErrorPagePassThru",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetPolish() {
	_jsii_.InvokeVoid(
		p,
		"resetPolish",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetResolveOverride() {
	_jsii_.InvokeVoid(
		p,
		"resetResolveOverride",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetRespectStrongEtag() {
	_jsii_.InvokeVoid(
		p,
		"resetRespectStrongEtag",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetResponseBuffering() {
	_jsii_.InvokeVoid(
		p,
		"resetResponseBuffering",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetRocketLoader() {
	_jsii_.InvokeVoid(
		p,
		"resetRocketLoader",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetSecurityLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetSecurityLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetServerSideExclude() {
	_jsii_.InvokeVoid(
		p,
		"resetServerSideExclude",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetSortQueryStringForCache() {
	_jsii_.InvokeVoid(
		p,
		"resetSortQueryStringForCache",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetSsl() {
	_jsii_.InvokeVoid(
		p,
		"resetSsl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetTrueClientIpHeader() {
	_jsii_.InvokeVoid(
		p,
		"resetTrueClientIpHeader",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ResetWaf() {
	_jsii_.InvokeVoid(
		p,
		"resetWaf",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


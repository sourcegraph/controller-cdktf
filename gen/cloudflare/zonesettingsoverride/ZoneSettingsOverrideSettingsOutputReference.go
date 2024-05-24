package zonesettingsoverride

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/zonesettingsoverride/internal"
)

type ZoneSettingsOverrideSettingsOutputReference interface {
	cdktf.ComplexObject
	AlwaysOnline() *string
	SetAlwaysOnline(val *string)
	AlwaysOnlineInput() *string
	AlwaysUseHttps() *string
	SetAlwaysUseHttps(val *string)
	AlwaysUseHttpsInput() *string
	AutomaticHttpsRewrites() *string
	SetAutomaticHttpsRewrites(val *string)
	AutomaticHttpsRewritesInput() *string
	BinaryAst() *string
	SetBinaryAst(val *string)
	BinaryAstInput() *string
	Brotli() *string
	SetBrotli(val *string)
	BrotliInput() *string
	BrowserCacheTtl() *float64
	SetBrowserCacheTtl(val *float64)
	BrowserCacheTtlInput() *float64
	BrowserCheck() *string
	SetBrowserCheck(val *string)
	BrowserCheckInput() *string
	CacheLevel() *string
	SetCacheLevel(val *string)
	CacheLevelInput() *string
	ChallengeTtl() *float64
	SetChallengeTtl(val *float64)
	ChallengeTtlInput() *float64
	Ciphers() *[]*string
	SetCiphers(val *[]*string)
	CiphersInput() *[]*string
	CnameFlattening() *string
	SetCnameFlattening(val *string)
	CnameFlatteningInput() *string
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
	DevelopmentMode() *string
	SetDevelopmentMode(val *string)
	DevelopmentModeInput() *string
	EarlyHints() *string
	SetEarlyHints(val *string)
	EarlyHintsInput() *string
	EmailObfuscation() *string
	SetEmailObfuscation(val *string)
	EmailObfuscationInput() *string
	FilterLogsToCloudflare() *string
	SetFilterLogsToCloudflare(val *string)
	FilterLogsToCloudflareInput() *string
	Fonts() *string
	SetFonts(val *string)
	FontsInput() *string
	// Experimental.
	Fqn() *string
	H2Prioritization() *string
	SetH2Prioritization(val *string)
	H2PrioritizationInput() *string
	HotlinkProtection() *string
	SetHotlinkProtection(val *string)
	HotlinkProtectionInput() *string
	Http2() *string
	SetHttp2(val *string)
	Http2Input() *string
	Http3() *string
	SetHttp3(val *string)
	Http3Input() *string
	ImageResizing() *string
	SetImageResizing(val *string)
	ImageResizingInput() *string
	InternalValue() *ZoneSettingsOverrideSettings
	SetInternalValue(val *ZoneSettingsOverrideSettings)
	IpGeolocation() *string
	SetIpGeolocation(val *string)
	IpGeolocationInput() *string
	Ipv6() *string
	SetIpv6(val *string)
	Ipv6Input() *string
	LogToCloudflare() *string
	SetLogToCloudflare(val *string)
	LogToCloudflareInput() *string
	MaxUpload() *float64
	SetMaxUpload(val *float64)
	MaxUploadInput() *float64
	Minify() ZoneSettingsOverrideSettingsMinifyOutputReference
	MinifyInput() *ZoneSettingsOverrideSettingsMinify
	MinTlsVersion() *string
	SetMinTlsVersion(val *string)
	MinTlsVersionInput() *string
	Mirage() *string
	SetMirage(val *string)
	MirageInput() *string
	MobileRedirect() ZoneSettingsOverrideSettingsMobileRedirectOutputReference
	MobileRedirectInput() *ZoneSettingsOverrideSettingsMobileRedirect
	OpportunisticEncryption() *string
	SetOpportunisticEncryption(val *string)
	OpportunisticEncryptionInput() *string
	OpportunisticOnion() *string
	SetOpportunisticOnion(val *string)
	OpportunisticOnionInput() *string
	OrangeToOrange() *string
	SetOrangeToOrange(val *string)
	OrangeToOrangeInput() *string
	OriginErrorPagePassThru() *string
	SetOriginErrorPagePassThru(val *string)
	OriginErrorPagePassThruInput() *string
	OriginMaxHttpVersion() *string
	SetOriginMaxHttpVersion(val *string)
	OriginMaxHttpVersionInput() *string
	Polish() *string
	SetPolish(val *string)
	PolishInput() *string
	PrefetchPreload() *string
	SetPrefetchPreload(val *string)
	PrefetchPreloadInput() *string
	PrivacyPass() *string
	SetPrivacyPass(val *string)
	PrivacyPassInput() *string
	ProxyReadTimeout() *string
	SetProxyReadTimeout(val *string)
	ProxyReadTimeoutInput() *string
	PseudoIpv4() *string
	SetPseudoIpv4(val *string)
	PseudoIpv4Input() *string
	ResponseBuffering() *string
	SetResponseBuffering(val *string)
	ResponseBufferingInput() *string
	RocketLoader() *string
	SetRocketLoader(val *string)
	RocketLoaderInput() *string
	SecurityHeader() ZoneSettingsOverrideSettingsSecurityHeaderOutputReference
	SecurityHeaderInput() *ZoneSettingsOverrideSettingsSecurityHeader
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
	Tls12Only() *string
	SetTls12Only(val *string)
	Tls12OnlyInput() *string
	Tls13() *string
	SetTls13(val *string)
	Tls13Input() *string
	TlsClientAuth() *string
	SetTlsClientAuth(val *string)
	TlsClientAuthInput() *string
	TrueClientIpHeader() *string
	SetTrueClientIpHeader(val *string)
	TrueClientIpHeaderInput() *string
	UniversalSsl() *string
	SetUniversalSsl(val *string)
	UniversalSslInput() *string
	VisitorIp() *string
	SetVisitorIp(val *string)
	VisitorIpInput() *string
	Waf() *string
	SetWaf(val *string)
	WafInput() *string
	Webp() *string
	SetWebp(val *string)
	WebpInput() *string
	Websockets() *string
	SetWebsockets(val *string)
	WebsocketsInput() *string
	ZeroRtt() *string
	SetZeroRtt(val *string)
	ZeroRttInput() *string
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
	PutMinify(value *ZoneSettingsOverrideSettingsMinify)
	PutMobileRedirect(value *ZoneSettingsOverrideSettingsMobileRedirect)
	PutSecurityHeader(value *ZoneSettingsOverrideSettingsSecurityHeader)
	ResetAlwaysOnline()
	ResetAlwaysUseHttps()
	ResetAutomaticHttpsRewrites()
	ResetBinaryAst()
	ResetBrotli()
	ResetBrowserCacheTtl()
	ResetBrowserCheck()
	ResetCacheLevel()
	ResetChallengeTtl()
	ResetCiphers()
	ResetCnameFlattening()
	ResetDevelopmentMode()
	ResetEarlyHints()
	ResetEmailObfuscation()
	ResetFilterLogsToCloudflare()
	ResetFonts()
	ResetH2Prioritization()
	ResetHotlinkProtection()
	ResetHttp2()
	ResetHttp3()
	ResetImageResizing()
	ResetIpGeolocation()
	ResetIpv6()
	ResetLogToCloudflare()
	ResetMaxUpload()
	ResetMinify()
	ResetMinTlsVersion()
	ResetMirage()
	ResetMobileRedirect()
	ResetOpportunisticEncryption()
	ResetOpportunisticOnion()
	ResetOrangeToOrange()
	ResetOriginErrorPagePassThru()
	ResetOriginMaxHttpVersion()
	ResetPolish()
	ResetPrefetchPreload()
	ResetPrivacyPass()
	ResetProxyReadTimeout()
	ResetPseudoIpv4()
	ResetResponseBuffering()
	ResetRocketLoader()
	ResetSecurityHeader()
	ResetSecurityLevel()
	ResetServerSideExclude()
	ResetSortQueryStringForCache()
	ResetSsl()
	ResetTls12Only()
	ResetTls13()
	ResetTlsClientAuth()
	ResetTrueClientIpHeader()
	ResetUniversalSsl()
	ResetVisitorIp()
	ResetWaf()
	ResetWebp()
	ResetWebsockets()
	ResetZeroRtt()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZoneSettingsOverrideSettingsOutputReference
type jsiiProxy_ZoneSettingsOverrideSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) AlwaysOnline() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alwaysOnline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) AlwaysOnlineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alwaysOnlineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) AlwaysUseHttps() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alwaysUseHttps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) AlwaysUseHttpsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alwaysUseHttpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) AutomaticHttpsRewrites() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticHttpsRewrites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) AutomaticHttpsRewritesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticHttpsRewritesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) BinaryAst() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryAst",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) BinaryAstInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryAstInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Brotli() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brotli",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) BrotliInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brotliInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) BrowserCacheTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"browserCacheTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) BrowserCacheTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"browserCacheTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) BrowserCheck() *string {
	var returns *string
	_jsii_.Get(
		j,
		"browserCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) BrowserCheckInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"browserCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) CacheLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) CacheLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ChallengeTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"challengeTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ChallengeTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"challengeTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Ciphers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ciphers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) CiphersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ciphersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) CnameFlattening() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cnameFlattening",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) CnameFlatteningInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cnameFlatteningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) DevelopmentMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developmentMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) DevelopmentModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developmentModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) EarlyHints() *string {
	var returns *string
	_jsii_.Get(
		j,
		"earlyHints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) EarlyHintsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"earlyHintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) EmailObfuscation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailObfuscation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) EmailObfuscationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailObfuscationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) FilterLogsToCloudflare() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterLogsToCloudflare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) FilterLogsToCloudflareInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterLogsToCloudflareInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Fonts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fonts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) FontsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) H2Prioritization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"h2Prioritization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) H2PrioritizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"h2PrioritizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) HotlinkProtection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hotlinkProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) HotlinkProtectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hotlinkProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Http2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"http2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Http2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"http2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Http3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"http3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Http3Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"http3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ImageResizing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageResizing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ImageResizingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageResizingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) InternalValue() *ZoneSettingsOverrideSettings {
	var returns *ZoneSettingsOverrideSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) IpGeolocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipGeolocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) IpGeolocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipGeolocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Ipv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Ipv6Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) LogToCloudflare() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logToCloudflare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) LogToCloudflareInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logToCloudflareInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) MaxUpload() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUpload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) MaxUploadInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUploadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Minify() ZoneSettingsOverrideSettingsMinifyOutputReference {
	var returns ZoneSettingsOverrideSettingsMinifyOutputReference
	_jsii_.Get(
		j,
		"minify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) MinifyInput() *ZoneSettingsOverrideSettingsMinify {
	var returns *ZoneSettingsOverrideSettingsMinify
	_jsii_.Get(
		j,
		"minifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) MinTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) MinTlsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTlsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Mirage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mirage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) MirageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mirageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) MobileRedirect() ZoneSettingsOverrideSettingsMobileRedirectOutputReference {
	var returns ZoneSettingsOverrideSettingsMobileRedirectOutputReference
	_jsii_.Get(
		j,
		"mobileRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) MobileRedirectInput() *ZoneSettingsOverrideSettingsMobileRedirect {
	var returns *ZoneSettingsOverrideSettingsMobileRedirect
	_jsii_.Get(
		j,
		"mobileRedirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) OpportunisticEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opportunisticEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) OpportunisticEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opportunisticEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) OpportunisticOnion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opportunisticOnion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) OpportunisticOnionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opportunisticOnionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) OrangeToOrange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orangeToOrange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) OrangeToOrangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orangeToOrangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) OriginErrorPagePassThru() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originErrorPagePassThru",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) OriginErrorPagePassThruInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originErrorPagePassThruInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) OriginMaxHttpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originMaxHttpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) OriginMaxHttpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originMaxHttpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Polish() *string {
	var returns *string
	_jsii_.Get(
		j,
		"polish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) PolishInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"polishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) PrefetchPreload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefetchPreload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) PrefetchPreloadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefetchPreloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) PrivacyPass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privacyPass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) PrivacyPassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privacyPassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ProxyReadTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyReadTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ProxyReadTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyReadTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) PseudoIpv4() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pseudoIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) PseudoIpv4Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pseudoIpv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResponseBuffering() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseBuffering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResponseBufferingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseBufferingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) RocketLoader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rocketLoader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) RocketLoaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rocketLoaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) SecurityHeader() ZoneSettingsOverrideSettingsSecurityHeaderOutputReference {
	var returns ZoneSettingsOverrideSettingsSecurityHeaderOutputReference
	_jsii_.Get(
		j,
		"securityHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) SecurityHeaderInput() *ZoneSettingsOverrideSettingsSecurityHeader {
	var returns *ZoneSettingsOverrideSettingsSecurityHeader
	_jsii_.Get(
		j,
		"securityHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) SecurityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) SecurityLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ServerSideExclude() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideExclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ServerSideExcludeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideExcludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) SortQueryStringForCache() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortQueryStringForCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) SortQueryStringForCacheInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortQueryStringForCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Ssl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) SslInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Tls12Only() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tls12Only",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Tls12OnlyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tls12OnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Tls13() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tls13",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Tls13Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tls13Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) TlsClientAuth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsClientAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) TlsClientAuthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsClientAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) TrueClientIpHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trueClientIpHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) TrueClientIpHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trueClientIpHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) UniversalSsl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"universalSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) UniversalSslInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"universalSslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) VisitorIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visitorIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) VisitorIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visitorIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Waf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) WafInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wafInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Webp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) WebpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Websockets() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websockets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) WebsocketsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websocketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ZeroRtt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zeroRtt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ZeroRttInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zeroRttInput",
		&returns,
	)
	return returns
}


func NewZoneSettingsOverrideSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZoneSettingsOverrideSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewZoneSettingsOverrideSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZoneSettingsOverrideSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zoneSettingsOverride.ZoneSettingsOverrideSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZoneSettingsOverrideSettingsOutputReference_Override(z ZoneSettingsOverrideSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zoneSettingsOverride.ZoneSettingsOverrideSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetAlwaysOnline(val *string) {
	if err := j.validateSetAlwaysOnlineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alwaysOnline",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetAlwaysUseHttps(val *string) {
	if err := j.validateSetAlwaysUseHttpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alwaysUseHttps",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetAutomaticHttpsRewrites(val *string) {
	if err := j.validateSetAutomaticHttpsRewritesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticHttpsRewrites",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetBinaryAst(val *string) {
	if err := j.validateSetBinaryAstParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binaryAst",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetBrotli(val *string) {
	if err := j.validateSetBrotliParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"brotli",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetBrowserCacheTtl(val *float64) {
	if err := j.validateSetBrowserCacheTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"browserCacheTtl",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetBrowserCheck(val *string) {
	if err := j.validateSetBrowserCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"browserCheck",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetCacheLevel(val *string) {
	if err := j.validateSetCacheLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheLevel",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetChallengeTtl(val *float64) {
	if err := j.validateSetChallengeTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"challengeTtl",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetCiphers(val *[]*string) {
	if err := j.validateSetCiphersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ciphers",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetCnameFlattening(val *string) {
	if err := j.validateSetCnameFlatteningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cnameFlattening",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetDevelopmentMode(val *string) {
	if err := j.validateSetDevelopmentModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"developmentMode",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetEarlyHints(val *string) {
	if err := j.validateSetEarlyHintsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"earlyHints",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetEmailObfuscation(val *string) {
	if err := j.validateSetEmailObfuscationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailObfuscation",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetFilterLogsToCloudflare(val *string) {
	if err := j.validateSetFilterLogsToCloudflareParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterLogsToCloudflare",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetFonts(val *string) {
	if err := j.validateSetFontsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fonts",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetH2Prioritization(val *string) {
	if err := j.validateSetH2PrioritizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"h2Prioritization",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetHotlinkProtection(val *string) {
	if err := j.validateSetHotlinkProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hotlinkProtection",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetHttp2(val *string) {
	if err := j.validateSetHttp2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"http2",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetHttp3(val *string) {
	if err := j.validateSetHttp3Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"http3",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetImageResizing(val *string) {
	if err := j.validateSetImageResizingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageResizing",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetInternalValue(val *ZoneSettingsOverrideSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetIpGeolocation(val *string) {
	if err := j.validateSetIpGeolocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipGeolocation",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetIpv6(val *string) {
	if err := j.validateSetIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetLogToCloudflare(val *string) {
	if err := j.validateSetLogToCloudflareParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logToCloudflare",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetMaxUpload(val *float64) {
	if err := j.validateSetMaxUploadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUpload",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetMinTlsVersion(val *string) {
	if err := j.validateSetMinTlsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTlsVersion",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetMirage(val *string) {
	if err := j.validateSetMirageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mirage",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetOpportunisticEncryption(val *string) {
	if err := j.validateSetOpportunisticEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"opportunisticEncryption",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetOpportunisticOnion(val *string) {
	if err := j.validateSetOpportunisticOnionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"opportunisticOnion",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetOrangeToOrange(val *string) {
	if err := j.validateSetOrangeToOrangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orangeToOrange",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetOriginErrorPagePassThru(val *string) {
	if err := j.validateSetOriginErrorPagePassThruParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originErrorPagePassThru",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetOriginMaxHttpVersion(val *string) {
	if err := j.validateSetOriginMaxHttpVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originMaxHttpVersion",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetPolish(val *string) {
	if err := j.validateSetPolishParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"polish",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetPrefetchPreload(val *string) {
	if err := j.validateSetPrefetchPreloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefetchPreload",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetPrivacyPass(val *string) {
	if err := j.validateSetPrivacyPassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privacyPass",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetProxyReadTimeout(val *string) {
	if err := j.validateSetProxyReadTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyReadTimeout",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetPseudoIpv4(val *string) {
	if err := j.validateSetPseudoIpv4Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pseudoIpv4",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetResponseBuffering(val *string) {
	if err := j.validateSetResponseBufferingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseBuffering",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetRocketLoader(val *string) {
	if err := j.validateSetRocketLoaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rocketLoader",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetSecurityLevel(val *string) {
	if err := j.validateSetSecurityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityLevel",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetServerSideExclude(val *string) {
	if err := j.validateSetServerSideExcludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverSideExclude",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetSortQueryStringForCache(val *string) {
	if err := j.validateSetSortQueryStringForCacheParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sortQueryStringForCache",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetSsl(val *string) {
	if err := j.validateSetSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssl",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetTls12Only(val *string) {
	if err := j.validateSetTls12OnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tls12Only",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetTls13(val *string) {
	if err := j.validateSetTls13Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tls13",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetTlsClientAuth(val *string) {
	if err := j.validateSetTlsClientAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsClientAuth",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetTrueClientIpHeader(val *string) {
	if err := j.validateSetTrueClientIpHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trueClientIpHeader",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetUniversalSsl(val *string) {
	if err := j.validateSetUniversalSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"universalSsl",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetVisitorIp(val *string) {
	if err := j.validateSetVisitorIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visitorIp",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetWaf(val *string) {
	if err := j.validateSetWafParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waf",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetWebp(val *string) {
	if err := j.validateSetWebpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webp",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetWebsockets(val *string) {
	if err := j.validateSetWebsocketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"websockets",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference)SetZeroRtt(val *string) {
	if err := j.validateSetZeroRttParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zeroRtt",
		val,
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := z.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := z.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := z.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		z,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := z.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		z,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := z.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		z,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := z.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		z,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := z.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		z,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := z.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		z,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := z.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		z,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := z.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) PutMinify(value *ZoneSettingsOverrideSettingsMinify) {
	if err := z.validatePutMinifyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putMinify",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) PutMobileRedirect(value *ZoneSettingsOverrideSettingsMobileRedirect) {
	if err := z.validatePutMobileRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putMobileRedirect",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) PutSecurityHeader(value *ZoneSettingsOverrideSettingsSecurityHeader) {
	if err := z.validatePutSecurityHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putSecurityHeader",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetAlwaysOnline() {
	_jsii_.InvokeVoid(
		z,
		"resetAlwaysOnline",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetAlwaysUseHttps() {
	_jsii_.InvokeVoid(
		z,
		"resetAlwaysUseHttps",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetAutomaticHttpsRewrites() {
	_jsii_.InvokeVoid(
		z,
		"resetAutomaticHttpsRewrites",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetBinaryAst() {
	_jsii_.InvokeVoid(
		z,
		"resetBinaryAst",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetBrotli() {
	_jsii_.InvokeVoid(
		z,
		"resetBrotli",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetBrowserCacheTtl() {
	_jsii_.InvokeVoid(
		z,
		"resetBrowserCacheTtl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetBrowserCheck() {
	_jsii_.InvokeVoid(
		z,
		"resetBrowserCheck",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetCacheLevel() {
	_jsii_.InvokeVoid(
		z,
		"resetCacheLevel",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetChallengeTtl() {
	_jsii_.InvokeVoid(
		z,
		"resetChallengeTtl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetCiphers() {
	_jsii_.InvokeVoid(
		z,
		"resetCiphers",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetCnameFlattening() {
	_jsii_.InvokeVoid(
		z,
		"resetCnameFlattening",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetDevelopmentMode() {
	_jsii_.InvokeVoid(
		z,
		"resetDevelopmentMode",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetEarlyHints() {
	_jsii_.InvokeVoid(
		z,
		"resetEarlyHints",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetEmailObfuscation() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailObfuscation",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetFilterLogsToCloudflare() {
	_jsii_.InvokeVoid(
		z,
		"resetFilterLogsToCloudflare",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetFonts() {
	_jsii_.InvokeVoid(
		z,
		"resetFonts",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetH2Prioritization() {
	_jsii_.InvokeVoid(
		z,
		"resetH2Prioritization",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetHotlinkProtection() {
	_jsii_.InvokeVoid(
		z,
		"resetHotlinkProtection",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetHttp2() {
	_jsii_.InvokeVoid(
		z,
		"resetHttp2",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetHttp3() {
	_jsii_.InvokeVoid(
		z,
		"resetHttp3",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetImageResizing() {
	_jsii_.InvokeVoid(
		z,
		"resetImageResizing",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetIpGeolocation() {
	_jsii_.InvokeVoid(
		z,
		"resetIpGeolocation",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetIpv6() {
	_jsii_.InvokeVoid(
		z,
		"resetIpv6",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetLogToCloudflare() {
	_jsii_.InvokeVoid(
		z,
		"resetLogToCloudflare",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetMaxUpload() {
	_jsii_.InvokeVoid(
		z,
		"resetMaxUpload",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetMinify() {
	_jsii_.InvokeVoid(
		z,
		"resetMinify",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetMinTlsVersion() {
	_jsii_.InvokeVoid(
		z,
		"resetMinTlsVersion",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetMirage() {
	_jsii_.InvokeVoid(
		z,
		"resetMirage",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetMobileRedirect() {
	_jsii_.InvokeVoid(
		z,
		"resetMobileRedirect",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetOpportunisticEncryption() {
	_jsii_.InvokeVoid(
		z,
		"resetOpportunisticEncryption",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetOpportunisticOnion() {
	_jsii_.InvokeVoid(
		z,
		"resetOpportunisticOnion",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetOrangeToOrange() {
	_jsii_.InvokeVoid(
		z,
		"resetOrangeToOrange",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetOriginErrorPagePassThru() {
	_jsii_.InvokeVoid(
		z,
		"resetOriginErrorPagePassThru",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetOriginMaxHttpVersion() {
	_jsii_.InvokeVoid(
		z,
		"resetOriginMaxHttpVersion",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetPolish() {
	_jsii_.InvokeVoid(
		z,
		"resetPolish",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetPrefetchPreload() {
	_jsii_.InvokeVoid(
		z,
		"resetPrefetchPreload",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetPrivacyPass() {
	_jsii_.InvokeVoid(
		z,
		"resetPrivacyPass",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetProxyReadTimeout() {
	_jsii_.InvokeVoid(
		z,
		"resetProxyReadTimeout",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetPseudoIpv4() {
	_jsii_.InvokeVoid(
		z,
		"resetPseudoIpv4",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetResponseBuffering() {
	_jsii_.InvokeVoid(
		z,
		"resetResponseBuffering",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetRocketLoader() {
	_jsii_.InvokeVoid(
		z,
		"resetRocketLoader",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetSecurityHeader() {
	_jsii_.InvokeVoid(
		z,
		"resetSecurityHeader",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetSecurityLevel() {
	_jsii_.InvokeVoid(
		z,
		"resetSecurityLevel",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetServerSideExclude() {
	_jsii_.InvokeVoid(
		z,
		"resetServerSideExclude",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetSortQueryStringForCache() {
	_jsii_.InvokeVoid(
		z,
		"resetSortQueryStringForCache",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetSsl() {
	_jsii_.InvokeVoid(
		z,
		"resetSsl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetTls12Only() {
	_jsii_.InvokeVoid(
		z,
		"resetTls12Only",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetTls13() {
	_jsii_.InvokeVoid(
		z,
		"resetTls13",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetTlsClientAuth() {
	_jsii_.InvokeVoid(
		z,
		"resetTlsClientAuth",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetTrueClientIpHeader() {
	_jsii_.InvokeVoid(
		z,
		"resetTrueClientIpHeader",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetUniversalSsl() {
	_jsii_.InvokeVoid(
		z,
		"resetUniversalSsl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetVisitorIp() {
	_jsii_.InvokeVoid(
		z,
		"resetVisitorIp",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetWaf() {
	_jsii_.InvokeVoid(
		z,
		"resetWaf",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetWebp() {
	_jsii_.InvokeVoid(
		z,
		"resetWebp",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetWebsockets() {
	_jsii_.InvokeVoid(
		z,
		"resetWebsockets",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ResetZeroRtt() {
	_jsii_.InvokeVoid(
		z,
		"resetZeroRtt",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := z.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		z,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


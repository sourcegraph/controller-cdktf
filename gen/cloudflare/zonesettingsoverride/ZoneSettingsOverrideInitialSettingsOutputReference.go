package zonesettingsoverride

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/zonesettingsoverride/internal"
)

type ZoneSettingsOverrideInitialSettingsOutputReference interface {
	cdktf.ComplexObject
	AlwaysOnline() *string
	AlwaysUseHttps() *string
	AutomaticHttpsRewrites() *string
	BinaryAst() *string
	Brotli() *string
	BrowserCacheTtl() *float64
	BrowserCheck() *string
	CacheLevel() *string
	ChallengeTtl() *float64
	Ciphers() *[]*string
	CnameFlattening() *string
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
	EarlyHints() *string
	EmailObfuscation() *string
	FilterLogsToCloudflare() *string
	// Experimental.
	Fqn() *string
	H2Prioritization() *string
	HotlinkProtection() *string
	Http2() *string
	Http3() *string
	ImageResizing() *string
	InternalValue() *ZoneSettingsOverrideInitialSettings
	SetInternalValue(val *ZoneSettingsOverrideInitialSettings)
	IpGeolocation() *string
	Ipv6() *string
	LogToCloudflare() *string
	MaxUpload() *float64
	Minify() ZoneSettingsOverrideInitialSettingsMinifyList
	MinTlsVersion() *string
	Mirage() *string
	MobileRedirect() ZoneSettingsOverrideInitialSettingsMobileRedirectList
	OpportunisticEncryption() *string
	OpportunisticOnion() *string
	OrangeToOrange() *string
	OriginErrorPagePassThru() *string
	OriginMaxHttpVersion() *string
	Polish() *string
	PrefetchPreload() *string
	PrivacyPass() *string
	ProxyReadTimeout() *string
	PseudoIpv4() *string
	ResponseBuffering() *string
	RocketLoader() *string
	SecurityHeader() ZoneSettingsOverrideInitialSettingsSecurityHeaderList
	SecurityLevel() *string
	ServerSideExclude() *string
	SortQueryStringForCache() *string
	Ssl() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Tls12Only() *string
	Tls13() *string
	TlsClientAuth() *string
	TrueClientIpHeader() *string
	UniversalSsl() *string
	VisitorIp() *string
	Waf() *string
	Webp() *string
	Websockets() *string
	ZeroRtt() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZoneSettingsOverrideInitialSettingsOutputReference
type jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) AlwaysOnline() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alwaysOnline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) AlwaysUseHttps() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alwaysUseHttps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) AutomaticHttpsRewrites() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticHttpsRewrites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) BinaryAst() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryAst",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) Brotli() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brotli",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) BrowserCacheTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"browserCacheTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) BrowserCheck() *string {
	var returns *string
	_jsii_.Get(
		j,
		"browserCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) CacheLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) ChallengeTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"challengeTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) Ciphers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ciphers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) CnameFlattening() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cnameFlattening",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) DevelopmentMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developmentMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) EarlyHints() *string {
	var returns *string
	_jsii_.Get(
		j,
		"earlyHints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) EmailObfuscation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailObfuscation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) FilterLogsToCloudflare() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterLogsToCloudflare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) H2Prioritization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"h2Prioritization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) HotlinkProtection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hotlinkProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) Http2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"http2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) Http3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"http3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) ImageResizing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageResizing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) InternalValue() *ZoneSettingsOverrideInitialSettings {
	var returns *ZoneSettingsOverrideInitialSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) IpGeolocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipGeolocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) Ipv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) LogToCloudflare() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logToCloudflare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) MaxUpload() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUpload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) Minify() ZoneSettingsOverrideInitialSettingsMinifyList {
	var returns ZoneSettingsOverrideInitialSettingsMinifyList
	_jsii_.Get(
		j,
		"minify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) MinTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) Mirage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mirage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) MobileRedirect() ZoneSettingsOverrideInitialSettingsMobileRedirectList {
	var returns ZoneSettingsOverrideInitialSettingsMobileRedirectList
	_jsii_.Get(
		j,
		"mobileRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) OpportunisticEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opportunisticEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) OpportunisticOnion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opportunisticOnion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) OrangeToOrange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orangeToOrange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) OriginErrorPagePassThru() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originErrorPagePassThru",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) OriginMaxHttpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originMaxHttpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) Polish() *string {
	var returns *string
	_jsii_.Get(
		j,
		"polish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) PrefetchPreload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefetchPreload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) PrivacyPass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privacyPass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) ProxyReadTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyReadTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) PseudoIpv4() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pseudoIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) ResponseBuffering() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseBuffering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) RocketLoader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rocketLoader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) SecurityHeader() ZoneSettingsOverrideInitialSettingsSecurityHeaderList {
	var returns ZoneSettingsOverrideInitialSettingsSecurityHeaderList
	_jsii_.Get(
		j,
		"securityHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) SecurityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) ServerSideExclude() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideExclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) SortQueryStringForCache() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortQueryStringForCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) Ssl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) Tls12Only() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tls12Only",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) Tls13() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tls13",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) TlsClientAuth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsClientAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) TrueClientIpHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trueClientIpHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) UniversalSsl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"universalSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) VisitorIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visitorIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) Waf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) Webp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) Websockets() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websockets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) ZeroRtt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zeroRtt",
		&returns,
	)
	return returns
}


func NewZoneSettingsOverrideInitialSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ZoneSettingsOverrideInitialSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewZoneSettingsOverrideInitialSettingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference{}

	_jsii_.Create(
		"cloudflare.zoneSettingsOverride.ZoneSettingsOverrideInitialSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewZoneSettingsOverrideInitialSettingsOutputReference_Override(z ZoneSettingsOverrideInitialSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.zoneSettingsOverride.ZoneSettingsOverrideInitialSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		z,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference)SetInternalValue(val *ZoneSettingsOverrideInitialSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


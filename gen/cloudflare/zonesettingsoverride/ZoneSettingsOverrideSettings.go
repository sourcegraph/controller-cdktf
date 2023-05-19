package zonesettingsoverride


type ZoneSettingsOverrideSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#always_online ZoneSettingsOverride#always_online}.
	AlwaysOnline *string `field:"optional" json:"alwaysOnline" yaml:"alwaysOnline"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#always_use_https ZoneSettingsOverride#always_use_https}.
	AlwaysUseHttps *string `field:"optional" json:"alwaysUseHttps" yaml:"alwaysUseHttps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#automatic_https_rewrites ZoneSettingsOverride#automatic_https_rewrites}.
	AutomaticHttpsRewrites *string `field:"optional" json:"automaticHttpsRewrites" yaml:"automaticHttpsRewrites"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#binary_ast ZoneSettingsOverride#binary_ast}.
	BinaryAst *string `field:"optional" json:"binaryAst" yaml:"binaryAst"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#brotli ZoneSettingsOverride#brotli}.
	Brotli *string `field:"optional" json:"brotli" yaml:"brotli"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#browser_cache_ttl ZoneSettingsOverride#browser_cache_ttl}.
	BrowserCacheTtl *float64 `field:"optional" json:"browserCacheTtl" yaml:"browserCacheTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#browser_check ZoneSettingsOverride#browser_check}.
	BrowserCheck *string `field:"optional" json:"browserCheck" yaml:"browserCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#cache_level ZoneSettingsOverride#cache_level}.
	CacheLevel *string `field:"optional" json:"cacheLevel" yaml:"cacheLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#challenge_ttl ZoneSettingsOverride#challenge_ttl}.
	ChallengeTtl *float64 `field:"optional" json:"challengeTtl" yaml:"challengeTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#ciphers ZoneSettingsOverride#ciphers}.
	Ciphers *[]*string `field:"optional" json:"ciphers" yaml:"ciphers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#cname_flattening ZoneSettingsOverride#cname_flattening}.
	CnameFlattening *string `field:"optional" json:"cnameFlattening" yaml:"cnameFlattening"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#development_mode ZoneSettingsOverride#development_mode}.
	DevelopmentMode *string `field:"optional" json:"developmentMode" yaml:"developmentMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#early_hints ZoneSettingsOverride#early_hints}.
	EarlyHints *string `field:"optional" json:"earlyHints" yaml:"earlyHints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#email_obfuscation ZoneSettingsOverride#email_obfuscation}.
	EmailObfuscation *string `field:"optional" json:"emailObfuscation" yaml:"emailObfuscation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#filter_logs_to_cloudflare ZoneSettingsOverride#filter_logs_to_cloudflare}.
	FilterLogsToCloudflare *string `field:"optional" json:"filterLogsToCloudflare" yaml:"filterLogsToCloudflare"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#h2_prioritization ZoneSettingsOverride#h2_prioritization}.
	H2Prioritization *string `field:"optional" json:"h2Prioritization" yaml:"h2Prioritization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#hotlink_protection ZoneSettingsOverride#hotlink_protection}.
	HotlinkProtection *string `field:"optional" json:"hotlinkProtection" yaml:"hotlinkProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#http2 ZoneSettingsOverride#http2}.
	Http2 *string `field:"optional" json:"http2" yaml:"http2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#http3 ZoneSettingsOverride#http3}.
	Http3 *string `field:"optional" json:"http3" yaml:"http3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#image_resizing ZoneSettingsOverride#image_resizing}.
	ImageResizing *string `field:"optional" json:"imageResizing" yaml:"imageResizing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#ip_geolocation ZoneSettingsOverride#ip_geolocation}.
	IpGeolocation *string `field:"optional" json:"ipGeolocation" yaml:"ipGeolocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#ipv6 ZoneSettingsOverride#ipv6}.
	Ipv6 *string `field:"optional" json:"ipv6" yaml:"ipv6"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#log_to_cloudflare ZoneSettingsOverride#log_to_cloudflare}.
	LogToCloudflare *string `field:"optional" json:"logToCloudflare" yaml:"logToCloudflare"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#max_upload ZoneSettingsOverride#max_upload}.
	MaxUpload *float64 `field:"optional" json:"maxUpload" yaml:"maxUpload"`
	// minify block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#minify ZoneSettingsOverride#minify}
	Minify *ZoneSettingsOverrideSettingsMinify `field:"optional" json:"minify" yaml:"minify"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#min_tls_version ZoneSettingsOverride#min_tls_version}.
	MinTlsVersion *string `field:"optional" json:"minTlsVersion" yaml:"minTlsVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#mirage ZoneSettingsOverride#mirage}.
	Mirage *string `field:"optional" json:"mirage" yaml:"mirage"`
	// mobile_redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#mobile_redirect ZoneSettingsOverride#mobile_redirect}
	MobileRedirect *ZoneSettingsOverrideSettingsMobileRedirect `field:"optional" json:"mobileRedirect" yaml:"mobileRedirect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#opportunistic_encryption ZoneSettingsOverride#opportunistic_encryption}.
	OpportunisticEncryption *string `field:"optional" json:"opportunisticEncryption" yaml:"opportunisticEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#opportunistic_onion ZoneSettingsOverride#opportunistic_onion}.
	OpportunisticOnion *string `field:"optional" json:"opportunisticOnion" yaml:"opportunisticOnion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#orange_to_orange ZoneSettingsOverride#orange_to_orange}.
	OrangeToOrange *string `field:"optional" json:"orangeToOrange" yaml:"orangeToOrange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#origin_error_page_pass_thru ZoneSettingsOverride#origin_error_page_pass_thru}.
	OriginErrorPagePassThru *string `field:"optional" json:"originErrorPagePassThru" yaml:"originErrorPagePassThru"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#origin_max_http_version ZoneSettingsOverride#origin_max_http_version}.
	OriginMaxHttpVersion *string `field:"optional" json:"originMaxHttpVersion" yaml:"originMaxHttpVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#polish ZoneSettingsOverride#polish}.
	Polish *string `field:"optional" json:"polish" yaml:"polish"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#prefetch_preload ZoneSettingsOverride#prefetch_preload}.
	PrefetchPreload *string `field:"optional" json:"prefetchPreload" yaml:"prefetchPreload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#privacy_pass ZoneSettingsOverride#privacy_pass}.
	PrivacyPass *string `field:"optional" json:"privacyPass" yaml:"privacyPass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#proxy_read_timeout ZoneSettingsOverride#proxy_read_timeout}.
	ProxyReadTimeout *string `field:"optional" json:"proxyReadTimeout" yaml:"proxyReadTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#pseudo_ipv4 ZoneSettingsOverride#pseudo_ipv4}.
	PseudoIpv4 *string `field:"optional" json:"pseudoIpv4" yaml:"pseudoIpv4"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#response_buffering ZoneSettingsOverride#response_buffering}.
	ResponseBuffering *string `field:"optional" json:"responseBuffering" yaml:"responseBuffering"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#rocket_loader ZoneSettingsOverride#rocket_loader}.
	RocketLoader *string `field:"optional" json:"rocketLoader" yaml:"rocketLoader"`
	// security_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#security_header ZoneSettingsOverride#security_header}
	SecurityHeader *ZoneSettingsOverrideSettingsSecurityHeader `field:"optional" json:"securityHeader" yaml:"securityHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#security_level ZoneSettingsOverride#security_level}.
	SecurityLevel *string `field:"optional" json:"securityLevel" yaml:"securityLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#server_side_exclude ZoneSettingsOverride#server_side_exclude}.
	ServerSideExclude *string `field:"optional" json:"serverSideExclude" yaml:"serverSideExclude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#sort_query_string_for_cache ZoneSettingsOverride#sort_query_string_for_cache}.
	SortQueryStringForCache *string `field:"optional" json:"sortQueryStringForCache" yaml:"sortQueryStringForCache"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#ssl ZoneSettingsOverride#ssl}.
	Ssl *string `field:"optional" json:"ssl" yaml:"ssl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#tls_1_2_only ZoneSettingsOverride#tls_1_2_only}.
	Tls12Only *string `field:"optional" json:"tls12Only" yaml:"tls12Only"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#tls_1_3 ZoneSettingsOverride#tls_1_3}.
	Tls13 *string `field:"optional" json:"tls13" yaml:"tls13"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#tls_client_auth ZoneSettingsOverride#tls_client_auth}.
	TlsClientAuth *string `field:"optional" json:"tlsClientAuth" yaml:"tlsClientAuth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#true_client_ip_header ZoneSettingsOverride#true_client_ip_header}.
	TrueClientIpHeader *string `field:"optional" json:"trueClientIpHeader" yaml:"trueClientIpHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#universal_ssl ZoneSettingsOverride#universal_ssl}.
	UniversalSsl *string `field:"optional" json:"universalSsl" yaml:"universalSsl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#visitor_ip ZoneSettingsOverride#visitor_ip}.
	VisitorIp *string `field:"optional" json:"visitorIp" yaml:"visitorIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#waf ZoneSettingsOverride#waf}.
	Waf *string `field:"optional" json:"waf" yaml:"waf"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#webp ZoneSettingsOverride#webp}.
	Webp *string `field:"optional" json:"webp" yaml:"webp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#websockets ZoneSettingsOverride#websockets}.
	Websockets *string `field:"optional" json:"websockets" yaml:"websockets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#zero_rtt ZoneSettingsOverride#zero_rtt}.
	ZeroRtt *string `field:"optional" json:"zeroRtt" yaml:"zeroRtt"`
}


package ruleset


type RulesetRulesActionParameters struct {
	// Turn on or off Cloudflare Automatic HTTPS rewrites.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#automatic_https_rewrites Ruleset#automatic_https_rewrites}
	AutomaticHttpsRewrites interface{} `field:"optional" json:"automaticHttpsRewrites" yaml:"automaticHttpsRewrites"`
	// autominify block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#autominify Ruleset#autominify}
	Autominify interface{} `field:"optional" json:"autominify" yaml:"autominify"`
	// Inspect the visitor's browser for headers commonly associated with spammers and certain bots.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#bic Ruleset#bic}
	Bic interface{} `field:"optional" json:"bic" yaml:"bic"`
	// browser_ttl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#browser_ttl Ruleset#browser_ttl}
	BrowserTtl interface{} `field:"optional" json:"browserTtl" yaml:"browserTtl"`
	// Whether to cache if expression matches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#cache Ruleset#cache}
	Cache interface{} `field:"optional" json:"cache" yaml:"cache"`
	// cache_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#cache_key Ruleset#cache_key}
	CacheKey interface{} `field:"optional" json:"cacheKey" yaml:"cacheKey"`
	// Content of the custom error response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#content Ruleset#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Content-Type of the custom error response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#content_type Ruleset#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// List of cookie values to include as part of custom fields logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#cookie_fields Ruleset#cookie_fields}
	CookieFields *[]*string `field:"optional" json:"cookieFields" yaml:"cookieFields"`
	// Turn off all active Cloudflare Apps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#disable_apps Ruleset#disable_apps}
	DisableApps interface{} `field:"optional" json:"disableApps" yaml:"disableApps"`
	// Turn off railgun feature of the Cloudflare Speed app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#disable_railgun Ruleset#disable_railgun}
	DisableRailgun interface{} `field:"optional" json:"disableRailgun" yaml:"disableRailgun"`
	// Turn off zaraz feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#disable_zaraz Ruleset#disable_zaraz}
	DisableZaraz interface{} `field:"optional" json:"disableZaraz" yaml:"disableZaraz"`
	// edge_ttl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#edge_ttl Ruleset#edge_ttl}
	EdgeTtl interface{} `field:"optional" json:"edgeTtl" yaml:"edgeTtl"`
	// Turn on or off the Cloudflare Email Obfuscation feature of the Cloudflare Scrape Shield app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#email_obfuscation Ruleset#email_obfuscation}
	EmailObfuscation interface{} `field:"optional" json:"emailObfuscation" yaml:"emailObfuscation"`
	// from_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#from_list Ruleset#from_list}
	FromList interface{} `field:"optional" json:"fromList" yaml:"fromList"`
	// from_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#from_value Ruleset#from_value}
	FromValue interface{} `field:"optional" json:"fromValue" yaml:"fromValue"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#headers Ruleset#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Host Header that request origin receives.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#host_header Ruleset#host_header}
	HostHeader *string `field:"optional" json:"hostHeader" yaml:"hostHeader"`
	// Turn on or off the hotlink protection feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#hotlink_protection Ruleset#hotlink_protection}
	HotlinkProtection interface{} `field:"optional" json:"hotlinkProtection" yaml:"hotlinkProtection"`
	// Identifier of the action parameter to modify.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#id Ruleset#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#increment Ruleset#increment}.
	Increment *float64 `field:"optional" json:"increment" yaml:"increment"`
	// matched_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#matched_data Ruleset#matched_data}
	MatchedData interface{} `field:"optional" json:"matchedData" yaml:"matchedData"`
	// Turn on or off Cloudflare Mirage of the Cloudflare Speed app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#mirage Ruleset#mirage}
	Mirage interface{} `field:"optional" json:"mirage" yaml:"mirage"`
	// Turn on or off the Cloudflare Opportunistic Encryption feature of the Edge Certificates tab in the Cloudflare SSL/TLS app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#opportunistic_encryption Ruleset#opportunistic_encryption}
	OpportunisticEncryption interface{} `field:"optional" json:"opportunisticEncryption" yaml:"opportunisticEncryption"`
	// origin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#origin Ruleset#origin}
	Origin interface{} `field:"optional" json:"origin" yaml:"origin"`
	// Pass-through error page for origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#origin_error_page_passthru Ruleset#origin_error_page_passthru}
	OriginErrorPagePassthru interface{} `field:"optional" json:"originErrorPagePassthru" yaml:"originErrorPagePassthru"`
	// overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#overrides Ruleset#overrides}
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// Point in the request/response lifecycle where the ruleset will be created.
	//
	// Available values: `ddos_l4`, `ddos_l7`, `http_custom_errors`, `http_log_custom_fields`, `http_request_cache_settings`, `http_request_firewall_custom`, `http_request_firewall_managed`, `http_request_late_transform`, `http_request_late_transform_managed`, `http_request_main`, `http_request_origin`, `http_request_dynamic_redirect`, `http_request_redirect`, `http_request_sanitize`, `http_request_transform`, `http_response_firewall_managed`, `http_response_headers_transform`, `http_response_headers_transform_managed`, `magic_transit`, `http_ratelimit`, `http_request_sbfm`, `http_config_settings`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#phases Ruleset#phases}
	Phases *[]*string `field:"optional" json:"phases" yaml:"phases"`
	// Apply options from the Polish feature of the Cloudflare Speed app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#polish Ruleset#polish}
	Polish *string `field:"optional" json:"polish" yaml:"polish"`
	// Products to target with the actions. Available values: `bic`, `hot`, `ratelimit`, `securityLevel`, `uablock`, `waf`, `zonelockdown`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#products Ruleset#products}
	Products *[]*string `field:"optional" json:"products" yaml:"products"`
	// List of request headers to include as part of custom fields logging, in lowercase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#request_fields Ruleset#request_fields}
	RequestFields *[]*string `field:"optional" json:"requestFields" yaml:"requestFields"`
	// Respect strong ETags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#respect_strong_etags Ruleset#respect_strong_etags}
	RespectStrongEtags interface{} `field:"optional" json:"respectStrongEtags" yaml:"respectStrongEtags"`
	// response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#response Ruleset#response}
	Response interface{} `field:"optional" json:"response" yaml:"response"`
	// List of response headers to include as part of custom fields logging, in lowercase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#response_fields Ruleset#response_fields}
	ResponseFields *[]*string `field:"optional" json:"responseFields" yaml:"responseFields"`
	// Turn on or off Cloudflare Rocket Loader in the Cloudflare Speed app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#rocket_loader Ruleset#rocket_loader}
	RocketLoader interface{} `field:"optional" json:"rocketLoader" yaml:"rocketLoader"`
	// Map of managed WAF rule ID to comma-delimited string of ruleset rule IDs.
	//
	// Example: `rules = { "efb7b8c949ac4650a09736fc376e9aee" = "5de7edfa648c4d6891dc3e7f84534ffa,e3a567afc347477d9702d9047e97d760" }`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#rules Ruleset#rules}
	Rules *map[string]*string `field:"optional" json:"rules" yaml:"rules"`
	// Which ruleset ID to target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#ruleset Ruleset#ruleset}
	Ruleset *string `field:"optional" json:"ruleset" yaml:"ruleset"`
	// List of managed WAF rule IDs to target. Only valid when the `"action"` is set to skip.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#rulesets Ruleset#rulesets}
	Rulesets *[]*string `field:"optional" json:"rulesets" yaml:"rulesets"`
	// Control options for the Security Level feature from the Security app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#security_level Ruleset#security_level}
	SecurityLevel *string `field:"optional" json:"securityLevel" yaml:"securityLevel"`
	// Turn on or off the Server Side Excludes feature of the Cloudflare Scrape Shield app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#server_side_excludes Ruleset#server_side_excludes}
	ServerSideExcludes interface{} `field:"optional" json:"serverSideExcludes" yaml:"serverSideExcludes"`
	// serve_stale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#serve_stale Ruleset#serve_stale}
	ServeStale interface{} `field:"optional" json:"serveStale" yaml:"serveStale"`
	// sni block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#sni Ruleset#sni}
	Sni interface{} `field:"optional" json:"sni" yaml:"sni"`
	// Control options for the SSL feature of the Edge Certificates tab in the Cloudflare SSL/TLS app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#ssl Ruleset#ssl}
	Ssl *string `field:"optional" json:"ssl" yaml:"ssl"`
	// HTTP status code of the custom error response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#status_code Ruleset#status_code}
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
	// Turn on or off the SXG feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#sxg Ruleset#sxg}
	Sxg interface{} `field:"optional" json:"sxg" yaml:"sxg"`
	// uri block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#uri Ruleset#uri}
	Uri interface{} `field:"optional" json:"uri" yaml:"uri"`
	// Version of the ruleset to deploy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#version Ruleset#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}


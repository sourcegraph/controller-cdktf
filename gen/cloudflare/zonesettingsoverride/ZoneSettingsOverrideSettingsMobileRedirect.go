package zonesettingsoverride


type ZoneSettingsOverrideSettingsMobileRedirect struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#mobile_subdomain ZoneSettingsOverride#mobile_subdomain}.
	MobileSubdomain *string `field:"required" json:"mobileSubdomain" yaml:"mobileSubdomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#status ZoneSettingsOverride#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#strip_uri ZoneSettingsOverride#strip_uri}.
	StripUri interface{} `field:"required" json:"stripUri" yaml:"stripUri"`
}


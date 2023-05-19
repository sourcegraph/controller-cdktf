package zonesettingsoverride


type ZoneSettingsOverrideSettingsMinify struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#css ZoneSettingsOverride#css}.
	Css *string `field:"required" json:"css" yaml:"css"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#html ZoneSettingsOverride#html}.
	Html *string `field:"required" json:"html" yaml:"html"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#js ZoneSettingsOverride#js}.
	Js *string `field:"required" json:"js" yaml:"js"`
}


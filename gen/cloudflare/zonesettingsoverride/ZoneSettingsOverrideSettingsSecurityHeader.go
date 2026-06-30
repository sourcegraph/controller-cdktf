package zonesettingsoverride

type ZoneSettingsOverrideSettingsSecurityHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#enabled ZoneSettingsOverride#enabled}.
	Enabled any `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#include_subdomains ZoneSettingsOverride#include_subdomains}.
	IncludeSubdomains any `field:"optional" json:"includeSubdomains" yaml:"includeSubdomains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#max_age ZoneSettingsOverride#max_age}.
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#nosniff ZoneSettingsOverride#nosniff}.
	Nosniff any `field:"optional" json:"nosniff" yaml:"nosniff"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_settings_override#preload ZoneSettingsOverride#preload}.
	Preload any `field:"optional" json:"preload" yaml:"preload"`
}

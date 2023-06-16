package devicesettingspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DeviceSettingsPolicyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The account identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_settings_policy#account_id DeviceSettingsPolicy#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_settings_policy#name DeviceSettingsPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether to allow devices to leave the organization. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_settings_policy#allowed_to_leave DeviceSettingsPolicy#allowed_to_leave}
	AllowedToLeave interface{} `field:"optional" json:"allowedToLeave" yaml:"allowedToLeave"`
	// Whether to allow mode switch for this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_settings_policy#allow_mode_switch DeviceSettingsPolicy#allow_mode_switch}
	AllowModeSwitch interface{} `field:"optional" json:"allowModeSwitch" yaml:"allowModeSwitch"`
	// Whether to allow updates under this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_settings_policy#allow_updates DeviceSettingsPolicy#allow_updates}
	AllowUpdates interface{} `field:"optional" json:"allowUpdates" yaml:"allowUpdates"`
	// The amount of time in minutes to reconnect after having been disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_settings_policy#auto_connect DeviceSettingsPolicy#auto_connect}
	AutoConnect *float64 `field:"optional" json:"autoConnect" yaml:"autoConnect"`
	// The captive portal value for this policy. Defaults to `180`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_settings_policy#captive_portal DeviceSettingsPolicy#captive_portal}
	CaptivePortal *float64 `field:"optional" json:"captivePortal" yaml:"captivePortal"`
	// Whether the policy refers to the default account policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_settings_policy#default DeviceSettingsPolicy#default}
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// Whether to disable auto fallback for this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_settings_policy#disable_auto_fallback DeviceSettingsPolicy#disable_auto_fallback}
	DisableAutoFallback interface{} `field:"optional" json:"disableAutoFallback" yaml:"disableAutoFallback"`
	// Whether the policy is enabled (cannot be set for default policies). Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_settings_policy#enabled DeviceSettingsPolicy#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Whether to add Microsoft IPs to split tunnel exclusions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_settings_policy#exclude_office_ips DeviceSettingsPolicy#exclude_office_ips}
	ExcludeOfficeIps interface{} `field:"optional" json:"excludeOfficeIps" yaml:"excludeOfficeIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_settings_policy#id DeviceSettingsPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Wirefilter expression to match a device against when evaluating whether this policy should take effect for that device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_settings_policy#match DeviceSettingsPolicy#match}
	Match *string `field:"optional" json:"match" yaml:"match"`
	// The precedence of the policy. Lower values indicate higher precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_settings_policy#precedence DeviceSettingsPolicy#precedence}
	Precedence *float64 `field:"optional" json:"precedence" yaml:"precedence"`
	// The service mode. Available values: `1dot1`, `warp`, `proxy`, `posture_only`, `warp_tunnel_only`. Defaults to `warp`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_settings_policy#service_mode_v2_mode DeviceSettingsPolicy#service_mode_v2_mode}
	ServiceModeV2Mode *string `field:"optional" json:"serviceModeV2Mode" yaml:"serviceModeV2Mode"`
	// The port to use for the proxy service mode. Required when using `service_mode_v2_mode`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_settings_policy#service_mode_v2_port DeviceSettingsPolicy#service_mode_v2_port}
	ServiceModeV2Port *float64 `field:"optional" json:"serviceModeV2Port" yaml:"serviceModeV2Port"`
	// The support URL that will be opened when sending feedback.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_settings_policy#support_url DeviceSettingsPolicy#support_url}
	SupportUrl *string `field:"optional" json:"supportUrl" yaml:"supportUrl"`
	// Enablement of the ZT client switch lock.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_settings_policy#switch_locked DeviceSettingsPolicy#switch_locked}
	SwitchLocked interface{} `field:"optional" json:"switchLocked" yaml:"switchLocked"`
}


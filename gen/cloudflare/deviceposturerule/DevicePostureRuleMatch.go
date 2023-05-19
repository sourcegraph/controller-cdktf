package deviceposturerule


type DevicePostureRuleMatch struct {
	// The platform of the device. Available values: `windows`, `mac`, `linux`, `android`, `ios`, `chromeos`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_posture_rule#platform DevicePostureRule#platform}
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
}


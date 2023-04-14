package deviceposturerule


type DevicePostureRuleMatch struct {
	// The platform of the device. Available values: `windows`, `mac`, `linux`, `android`, `ios`, `chromeos`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#platform DevicePostureRule#platform}
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
}


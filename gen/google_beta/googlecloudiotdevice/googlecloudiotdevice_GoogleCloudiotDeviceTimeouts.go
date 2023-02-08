package googlecloudiotdevice


type GoogleCloudiotDeviceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudiot_device#create GoogleCloudiotDevice#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudiot_device#delete GoogleCloudiotDevice#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudiot_device#update GoogleCloudiotDevice#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


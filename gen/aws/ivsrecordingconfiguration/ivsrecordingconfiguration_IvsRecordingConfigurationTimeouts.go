package ivsrecordingconfiguration


type IvsRecordingConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ivs_recording_configuration#create IvsRecordingConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ivs_recording_configuration#delete IvsRecordingConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}


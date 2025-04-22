package googlememorystoreinstance


type GoogleMemorystoreInstanceZoneDistributionConfig struct {
	// Optional. Current zone distribution mode. Defaults to MULTI_ZONE.   Possible values:  MULTI_ZONE SINGLE_ZONE Possible values: ["MULTI_ZONE", "SINGLE_ZONE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_memorystore_instance#mode GoogleMemorystoreInstance#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Optional. Defines zone where all resources will be allocated with SINGLE_ZONE mode. Ignored for MULTI_ZONE mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_memorystore_instance#zone GoogleMemorystoreInstance#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}


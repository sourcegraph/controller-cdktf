package googlelustreinstance


type GoogleLustreInstanceDynamicTierOptions struct {
	// The dynamic tier mode of the instance. Possible values: DISABLED DEFAULT_CACHE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_lustre_instance#mode GoogleLustreInstance#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}


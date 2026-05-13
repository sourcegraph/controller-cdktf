package lustreinstance


type LustreInstanceDynamicTierOptions struct {
	// The dynamic tier mode of the instance. Possible values: DISABLED DEFAULT_CACHE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/lustre_instance#mode LustreInstance#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}


package zonelockdown


type ZoneLockdownConfigurations struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_lockdown#target ZoneLockdown#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_lockdown#value ZoneLockdown#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}


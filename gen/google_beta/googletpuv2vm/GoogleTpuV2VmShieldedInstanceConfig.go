package googletpuv2vm


type GoogleTpuV2VmShieldedInstanceConfig struct {
	// Defines whether the instance has Secure Boot enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#enable_secure_boot GoogleTpuV2Vm#enable_secure_boot}
	EnableSecureBoot interface{} `field:"required" json:"enableSecureBoot" yaml:"enableSecureBoot"`
}


package googlefilestoreinstance


type GoogleFilestoreInstancePerformanceConfig struct {
	// fixed_iops block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_filestore_instance#fixed_iops GoogleFilestoreInstance#fixed_iops}
	FixedIops *GoogleFilestoreInstancePerformanceConfigFixedIops `field:"optional" json:"fixedIops" yaml:"fixedIops"`
	// iops_per_tb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_filestore_instance#iops_per_tb GoogleFilestoreInstance#iops_per_tb}
	IopsPerTb *GoogleFilestoreInstancePerformanceConfigIopsPerTb `field:"optional" json:"iopsPerTb" yaml:"iopsPerTb"`
}


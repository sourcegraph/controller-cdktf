package googlealloydbinstance


type GoogleAlloydbInstanceMachineConfig struct {
	// The number of CPU's in the VM instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_instance#cpu_count GoogleAlloydbInstance#cpu_count}
	CpuCount *float64 `field:"optional" json:"cpuCount" yaml:"cpuCount"`
}


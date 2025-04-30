package googlefilestoreinstance


type GoogleFilestoreInstancePerformanceConfigIopsPerTb struct {
	// The instance max IOPS will be calculated by multiplying the capacity of the instance (TB) by max_iops_per_tb, and rounding to the nearest 1000.
	//
	// The instance max IOPS
	// will be changed dynamically based on the instance
	// capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_filestore_instance#max_iops_per_tb GoogleFilestoreInstance#max_iops_per_tb}
	MaxIopsPerTb *float64 `field:"optional" json:"maxIopsPerTb" yaml:"maxIopsPerTb"`
}


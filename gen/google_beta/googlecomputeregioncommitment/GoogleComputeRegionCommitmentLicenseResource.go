package googlecomputeregioncommitment


type GoogleComputeRegionCommitmentLicenseResource struct {
	// Any applicable license URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_commitment#license GoogleComputeRegionCommitment#license}
	License *string `field:"required" json:"license" yaml:"license"`
	// The number of licenses purchased.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_commitment#amount GoogleComputeRegionCommitment#amount}
	Amount *string `field:"optional" json:"amount" yaml:"amount"`
	// Specifies the core range of the instance for which this license applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_commitment#cores_per_license GoogleComputeRegionCommitment#cores_per_license}
	CoresPerLicense *string `field:"optional" json:"coresPerLicense" yaml:"coresPerLicense"`
}


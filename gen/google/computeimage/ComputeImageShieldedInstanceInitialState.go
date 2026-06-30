package computeimage

type ComputeImageShieldedInstanceInitialState struct {
	// dbs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_image#dbs ComputeImage#dbs}
	Dbs any `field:"optional" json:"dbs" yaml:"dbs"`
	// dbxs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_image#dbxs ComputeImage#dbxs}
	Dbxs any `field:"optional" json:"dbxs" yaml:"dbxs"`
	// keks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_image#keks ComputeImage#keks}
	Keks any `field:"optional" json:"keks" yaml:"keks"`
	// pk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_image#pk ComputeImage#pk}
	Pk *ComputeImageShieldedInstanceInitialStatePk `field:"optional" json:"pk" yaml:"pk"`
}

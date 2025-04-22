package googlecomputeimage


type GoogleComputeImageShieldedInstanceInitialState struct {
	// dbs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_image#dbs GoogleComputeImage#dbs}
	Dbs interface{} `field:"optional" json:"dbs" yaml:"dbs"`
	// dbxs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_image#dbxs GoogleComputeImage#dbxs}
	Dbxs interface{} `field:"optional" json:"dbxs" yaml:"dbxs"`
	// keks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_image#keks GoogleComputeImage#keks}
	Keks interface{} `field:"optional" json:"keks" yaml:"keks"`
	// pk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_image#pk GoogleComputeImage#pk}
	Pk *GoogleComputeImageShieldedInstanceInitialStatePk `field:"optional" json:"pk" yaml:"pk"`
}


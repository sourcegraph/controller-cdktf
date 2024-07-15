package googlecomputeurlmap


type GoogleComputeUrlMapTest struct {
	// Host portion of the URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#host GoogleComputeUrlMap#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// Path portion of the URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#path GoogleComputeUrlMap#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// The backend service or backend bucket link that should be matched by this test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#service GoogleComputeUrlMap#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Description of this test case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#description GoogleComputeUrlMap#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}


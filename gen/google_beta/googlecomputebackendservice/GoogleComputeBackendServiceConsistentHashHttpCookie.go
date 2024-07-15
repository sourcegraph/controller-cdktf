package googlecomputebackendservice


type GoogleComputeBackendServiceConsistentHashHttpCookie struct {
	// Name of the cookie.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_backend_service#name GoogleComputeBackendService#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Path to set for the cookie.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_backend_service#path GoogleComputeBackendService#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// ttl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_backend_service#ttl GoogleComputeBackendService#ttl}
	Ttl *GoogleComputeBackendServiceConsistentHashHttpCookieTtl `field:"optional" json:"ttl" yaml:"ttl"`
}


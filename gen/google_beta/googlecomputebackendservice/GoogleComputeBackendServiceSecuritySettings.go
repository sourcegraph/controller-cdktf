package googlecomputebackendservice


type GoogleComputeBackendServiceSecuritySettings struct {
	// ClientTlsPolicy is a resource that specifies how a client should authenticate connections to backends of a service.
	//
	// This resource itself does not affect
	// configuration unless it is attached to a backend service resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_backend_service#client_tls_policy GoogleComputeBackendService#client_tls_policy}
	ClientTlsPolicy *string `field:"required" json:"clientTlsPolicy" yaml:"clientTlsPolicy"`
	// A list of alternate names to verify the subject identity in the certificate.
	//
	// If specified, the client will verify that the server certificate's subject
	// alt name matches one of the specified values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_backend_service#subject_alt_names GoogleComputeBackendService#subject_alt_names}
	SubjectAltNames *[]*string `field:"required" json:"subjectAltNames" yaml:"subjectAltNames"`
}


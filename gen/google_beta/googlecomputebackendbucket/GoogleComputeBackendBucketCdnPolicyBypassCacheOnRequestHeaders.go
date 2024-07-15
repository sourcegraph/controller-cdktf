package googlecomputebackendbucket


type GoogleComputeBackendBucketCdnPolicyBypassCacheOnRequestHeaders struct {
	// The header field name to match on when bypassing cache. Values are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_backend_bucket#header_name GoogleComputeBackendBucket#header_name}
	HeaderName *string `field:"optional" json:"headerName" yaml:"headerName"`
}


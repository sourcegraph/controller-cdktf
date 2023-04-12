package googlecomputebackendbucket


type GoogleComputeBackendBucketCdnPolicyBypassCacheOnRequestHeaders struct {
	// The header field name to match on when bypassing cache. Values are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_backend_bucket#header_name GoogleComputeBackendBucket#header_name}
	HeaderName *string `field:"optional" json:"headerName" yaml:"headerName"`
}


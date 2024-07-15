package googlecomputefirewall


type GoogleComputeFirewallLogConfig struct {
	// This field denotes whether to include or exclude metadata for firewall logs. Possible values: ["EXCLUDE_ALL_METADATA", "INCLUDE_ALL_METADATA"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_firewall#metadata GoogleComputeFirewall#metadata}
	Metadata *string `field:"required" json:"metadata" yaml:"metadata"`
}


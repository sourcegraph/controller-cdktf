package googlecomputeinterconnect


type GoogleComputeInterconnectMacsec struct {
	// pre_shared_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_interconnect#pre_shared_keys GoogleComputeInterconnect#pre_shared_keys}
	PreSharedKeys interface{} `field:"required" json:"preSharedKeys" yaml:"preSharedKeys"`
	// If set to true, the Interconnect connection is configured with a should-secure MACsec security policy, that allows the Google router to fallback to cleartext traffic if the MKA session cannot be established.
	//
	// By default, the Interconnect
	// connection is configured with a must-secure security policy that drops all traffic
	// if the MKA session cannot be established with your router.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_interconnect#fail_open GoogleComputeInterconnect#fail_open}
	FailOpen interface{} `field:"optional" json:"failOpen" yaml:"failOpen"`
}


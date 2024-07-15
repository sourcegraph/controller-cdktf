package googlednsresponsepolicy


type GoogleDnsResponsePolicyNetworks struct {
	// The fully qualified URL of the VPC network to bind to. This should be formatted like 'https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_response_policy#network_url GoogleDnsResponsePolicy#network_url}
	NetworkUrl *string `field:"required" json:"networkUrl" yaml:"networkUrl"`
}


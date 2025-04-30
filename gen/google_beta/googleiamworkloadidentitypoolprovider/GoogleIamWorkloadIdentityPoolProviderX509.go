package googleiamworkloadidentitypoolprovider


type GoogleIamWorkloadIdentityPoolProviderX509 struct {
	// trust_store block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iam_workload_identity_pool_provider#trust_store GoogleIamWorkloadIdentityPoolProvider#trust_store}
	TrustStore *GoogleIamWorkloadIdentityPoolProviderX509TrustStore `field:"required" json:"trustStore" yaml:"trustStore"`
}


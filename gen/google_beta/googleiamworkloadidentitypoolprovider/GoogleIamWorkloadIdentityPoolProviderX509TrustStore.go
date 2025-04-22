package googleiamworkloadidentitypoolprovider


type GoogleIamWorkloadIdentityPoolProviderX509TrustStore struct {
	// trust_anchors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_iam_workload_identity_pool_provider#trust_anchors GoogleIamWorkloadIdentityPoolProvider#trust_anchors}
	TrustAnchors interface{} `field:"required" json:"trustAnchors" yaml:"trustAnchors"`
	// intermediate_cas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_iam_workload_identity_pool_provider#intermediate_cas GoogleIamWorkloadIdentityPoolProvider#intermediate_cas}
	IntermediateCas interface{} `field:"optional" json:"intermediateCas" yaml:"intermediateCas"`
}


package googleiamworkloadidentitypool


type GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundles struct {
	// trust_anchors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workload_identity_pool#trust_anchors GoogleIamWorkloadIdentityPool#trust_anchors}
	TrustAnchors interface{} `field:"required" json:"trustAnchors" yaml:"trustAnchors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workload_identity_pool#trust_domain GoogleIamWorkloadIdentityPool#trust_domain}.
	TrustDomain *string `field:"required" json:"trustDomain" yaml:"trustDomain"`
	// If set to True, the trust bundle will include the private ca managed identity regional root public certificates.
	//
	// ~> **Note** 'trust_default_shared_ca' is only supported for managed identity trust domain
	// resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workload_identity_pool#trust_default_shared_ca GoogleIamWorkloadIdentityPool#trust_default_shared_ca}
	TrustDefaultSharedCa interface{} `field:"optional" json:"trustDefaultSharedCa" yaml:"trustDefaultSharedCa"`
}


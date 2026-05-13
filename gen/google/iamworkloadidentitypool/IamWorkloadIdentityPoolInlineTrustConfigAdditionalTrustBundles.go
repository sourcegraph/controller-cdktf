package iamworkloadidentitypool


type IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundles struct {
	// trust_anchors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool#trust_anchors IamWorkloadIdentityPool#trust_anchors}
	TrustAnchors interface{} `field:"required" json:"trustAnchors" yaml:"trustAnchors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool#trust_domain IamWorkloadIdentityPool#trust_domain}.
	TrustDomain *string `field:"required" json:"trustDomain" yaml:"trustDomain"`
	// If set to True, the trust bundle will include the private ca managed identity regional root public certificates.
	//
	// ~> **Note** 'trust_default_shared_ca' is only supported for managed identity trust domain
	// resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool#trust_default_shared_ca IamWorkloadIdentityPool#trust_default_shared_ca}
	TrustDefaultSharedCa interface{} `field:"optional" json:"trustDefaultSharedCa" yaml:"trustDefaultSharedCa"`
}


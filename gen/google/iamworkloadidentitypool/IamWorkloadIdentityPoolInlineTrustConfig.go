package iamworkloadidentitypool

type IamWorkloadIdentityPoolInlineTrustConfig struct {
	// additional_trust_bundles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/iam_workload_identity_pool#additional_trust_bundles IamWorkloadIdentityPool#additional_trust_bundles}
	AdditionalTrustBundles any `field:"optional" json:"additionalTrustBundles" yaml:"additionalTrustBundles"`
}

package gkehubfeaturemembership

type GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent struct {
	// bundles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/gke_hub_feature_membership#bundles GkeHubFeatureMembership#bundles}
	Bundles any `field:"optional" json:"bundles" yaml:"bundles"`
	// template_library block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/gke_hub_feature_membership#template_library GkeHubFeatureMembership#template_library}
	TemplateLibrary *GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary `field:"optional" json:"templateLibrary" yaml:"templateLibrary"`
}

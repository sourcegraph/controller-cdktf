package gkehubfeaturemembership

type GkeHubFeatureMembershipConfigmanagementHierarchyController struct {
	// **DEPRECATED** Configuring Hierarchy Controller through the configmanagement feature is no longer recommended. Use https://github.com/kubernetes-sigs/hierarchical-namespaces instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/gke_hub_feature_membership#enabled GkeHubFeatureMembership#enabled}
	Enabled any `field:"optional" json:"enabled" yaml:"enabled"`
	// Whether hierarchical resource quota is enabled in this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/gke_hub_feature_membership#enable_hierarchical_resource_quota GkeHubFeatureMembership#enable_hierarchical_resource_quota}
	EnableHierarchicalResourceQuota any `field:"optional" json:"enableHierarchicalResourceQuota" yaml:"enableHierarchicalResourceQuota"`
	// Whether pod tree labels are enabled in this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/gke_hub_feature_membership#enable_pod_tree_labels GkeHubFeatureMembership#enable_pod_tree_labels}
	EnablePodTreeLabels any `field:"optional" json:"enablePodTreeLabels" yaml:"enablePodTreeLabels"`
}

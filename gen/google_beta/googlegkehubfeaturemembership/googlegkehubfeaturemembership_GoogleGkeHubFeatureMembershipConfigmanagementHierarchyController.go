package googlegkehubfeaturemembership


type GoogleGkeHubFeatureMembershipConfigmanagementHierarchyController struct {
	// Whether Hierarchy Controller is enabled in this cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_hub_feature_membership#enabled GoogleGkeHubFeatureMembership#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Whether hierarchical resource quota is enabled in this cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_hub_feature_membership#enable_hierarchical_resource_quota GoogleGkeHubFeatureMembership#enable_hierarchical_resource_quota}
	EnableHierarchicalResourceQuota interface{} `field:"optional" json:"enableHierarchicalResourceQuota" yaml:"enableHierarchicalResourceQuota"`
	// Whether pod tree labels are enabled in this cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_hub_feature_membership#enable_pod_tree_labels GoogleGkeHubFeatureMembership#enable_pod_tree_labels}
	EnablePodTreeLabels interface{} `field:"optional" json:"enablePodTreeLabels" yaml:"enablePodTreeLabels"`
}

package googlecomputerolloutplan


type GoogleComputeRolloutPlanWavesSelectorsResourceHierarchySelector struct {
	// Format: "folders/{folder_id}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_rollout_plan#included_folders GoogleComputeRolloutPlan#included_folders}
	IncludedFolders *[]*string `field:"optional" json:"includedFolders" yaml:"includedFolders"`
	// Format: "organizations/{organization_id}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_rollout_plan#included_organizations GoogleComputeRolloutPlan#included_organizations}
	IncludedOrganizations *[]*string `field:"optional" json:"includedOrganizations" yaml:"includedOrganizations"`
	// Format: "projects/{project_id}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_rollout_plan#included_projects GoogleComputeRolloutPlan#included_projects}
	IncludedProjects *[]*string `field:"optional" json:"includedProjects" yaml:"includedProjects"`
}


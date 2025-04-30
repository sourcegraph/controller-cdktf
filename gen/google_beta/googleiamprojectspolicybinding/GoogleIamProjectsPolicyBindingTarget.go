package googleiamprojectspolicybinding


type GoogleIamProjectsPolicyBindingTarget struct {
	// Required.
	//
	// Immutable. Full Resource Name of the principal set used for principal access boundary policy bindings.
	// Examples for each one of the following supported principal set types:
	// * Project:
	//   * '//cloudresourcemanager.googleapis.com/projects/PROJECT_NUMBER'
	//   * '//cloudresourcemanager.googleapis.com/projects/PROJECT_ID'
	// * Workload Identity Pool: '//iam.googleapis.com/projects/PROJECT_NUMBER/locations/LOCATION/workloadIdentityPools/WORKLOAD_POOL_ID'
	// It must be parent by the policy binding's parent (the project).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iam_projects_policy_binding#principal_set GoogleIamProjectsPolicyBinding#principal_set}
	PrincipalSet *string `field:"optional" json:"principalSet" yaml:"principalSet"`
}


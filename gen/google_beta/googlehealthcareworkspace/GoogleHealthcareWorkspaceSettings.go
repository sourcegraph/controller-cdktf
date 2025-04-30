package googlehealthcareworkspace


type GoogleHealthcareWorkspaceSettings struct {
	// Project IDs for data projects hosted in a workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_healthcare_workspace#data_project_ids GoogleHealthcareWorkspace#data_project_ids}
	DataProjectIds *[]*string `field:"required" json:"dataProjectIds" yaml:"dataProjectIds"`
}


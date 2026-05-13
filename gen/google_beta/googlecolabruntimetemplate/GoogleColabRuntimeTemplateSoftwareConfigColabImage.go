package googlecolabruntimetemplate


type GoogleColabRuntimeTemplateSoftwareConfigColabImage struct {
	// The release name of the NotebookRuntime Colab image, e.g. "py310". If not specified, detault to the latest release.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_colab_runtime_template#release_name GoogleColabRuntimeTemplate#release_name}
	ReleaseName *string `field:"optional" json:"releaseName" yaml:"releaseName"`
}


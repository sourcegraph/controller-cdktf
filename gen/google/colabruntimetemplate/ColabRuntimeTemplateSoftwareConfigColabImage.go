package colabruntimetemplate


type ColabRuntimeTemplateSoftwareConfigColabImage struct {
	// The release name of the NotebookRuntime Colab image, e.g. "py310". If not specified, detault to the latest release.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/colab_runtime_template#release_name ColabRuntimeTemplate#release_name}
	ReleaseName *string `field:"optional" json:"releaseName" yaml:"releaseName"`
}


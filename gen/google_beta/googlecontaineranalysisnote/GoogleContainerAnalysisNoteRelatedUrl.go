package googlecontaineranalysisnote


type GoogleContainerAnalysisNoteRelatedUrl struct {
	// Specific URL associated with the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_analysis_note#url GoogleContainerAnalysisNote#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Label to describe usage of the URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_analysis_note#label GoogleContainerAnalysisNote#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
}


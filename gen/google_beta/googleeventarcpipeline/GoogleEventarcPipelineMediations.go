package googleeventarcpipeline


type GoogleEventarcPipelineMediations struct {
	// transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_eventarc_pipeline#transformation GoogleEventarcPipeline#transformation}
	Transformation *GoogleEventarcPipelineMediationsTransformation `field:"optional" json:"transformation" yaml:"transformation"`
}


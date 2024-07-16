package googledataplextask


type GoogleDataplexTaskNotebookInfrastructureSpec struct {
	// batch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task#batch GoogleDataplexTask#batch}
	Batch *GoogleDataplexTaskNotebookInfrastructureSpecBatch `field:"optional" json:"batch" yaml:"batch"`
	// container_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task#container_image GoogleDataplexTask#container_image}
	ContainerImage *GoogleDataplexTaskNotebookInfrastructureSpecContainerImage `field:"optional" json:"containerImage" yaml:"containerImage"`
	// vpc_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task#vpc_network GoogleDataplexTask#vpc_network}
	VpcNetwork *GoogleDataplexTaskNotebookInfrastructureSpecVpcNetwork `field:"optional" json:"vpcNetwork" yaml:"vpcNetwork"`
}


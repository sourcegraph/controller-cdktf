package appmeshvirtualservice


type AppmeshVirtualServiceSpec struct {
	// provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/appmesh_virtual_service#provider AppmeshVirtualService#provider}
	Provider *AppmeshVirtualServiceSpecProvider `field:"optional" json:"provider" yaml:"provider"`
}


package appmeshvirtualrouter


type AppmeshVirtualRouterSpec struct {
	// listener block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/appmesh_virtual_router#listener AppmeshVirtualRouter#listener}
	Listener interface{} `field:"required" json:"listener" yaml:"listener"`
}


package appmeshvirtualrouter


type AppmeshVirtualRouterSpecListener struct {
	// port_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/appmesh_virtual_router#port_mapping AppmeshVirtualRouter#port_mapping}
	PortMapping *AppmeshVirtualRouterSpecListenerPortMapping `field:"required" json:"portMapping" yaml:"portMapping"`
}


package replicationcontroller


type ReplicationControllerSpecTemplateSpecContainerSecurityContextCapabilities struct {
	// Added capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/replication_controller#add ReplicationController#add}
	Add *[]*string `field:"optional" json:"add" yaml:"add"`
	// Removed capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/replication_controller#drop ReplicationController#drop}
	Drop *[]*string `field:"optional" json:"drop" yaml:"drop"`
}


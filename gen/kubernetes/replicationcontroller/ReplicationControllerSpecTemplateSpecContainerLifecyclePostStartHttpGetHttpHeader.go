package replicationcontroller


type ReplicationControllerSpecTemplateSpecContainerLifecyclePostStartHttpGetHttpHeader struct {
	// The header field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/replication_controller#name ReplicationController#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The header field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/replication_controller#value ReplicationController#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


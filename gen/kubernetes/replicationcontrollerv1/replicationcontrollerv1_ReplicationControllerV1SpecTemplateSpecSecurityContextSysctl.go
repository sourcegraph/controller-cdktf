package replicationcontrollerv1


type ReplicationControllerV1SpecTemplateSpecSecurityContextSysctl struct {
	// Name of a property to set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller_v1#name ReplicationControllerV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of a property to set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller_v1#value ReplicationControllerV1#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}


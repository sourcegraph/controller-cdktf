package replicationcontrollerv1


type ReplicationControllerV1SpecTemplateSpecVolumeRbdSecretRef struct {
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/replication_controller_v1#name ReplicationControllerV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/replication_controller_v1#namespace ReplicationControllerV1#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}


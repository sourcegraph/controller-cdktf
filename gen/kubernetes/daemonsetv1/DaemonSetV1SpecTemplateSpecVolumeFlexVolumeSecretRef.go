package daemonsetv1


type DaemonSetV1SpecTemplateSpecVolumeFlexVolumeSecretRef struct {
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/daemon_set_v1#name DaemonSetV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/daemon_set_v1#namespace DaemonSetV1#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}


package daemonset


type DaemonsetSpecTemplateSpecVolumeProjectedSourcesSecret struct {
	// items block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/daemonset#items Daemonset#items}
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// Name of the secret in the pod's namespace to use. More info: http://kubernetes.io/docs/user-guide/volumes#secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/daemonset#name Daemonset#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Optional: Specify whether the Secret or it's keys must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/daemonset#optional Daemonset#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}


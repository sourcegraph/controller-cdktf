package daemonsetv1


type DaemonSetV1SpecTemplateSpecVolumeNfs struct {
	// Path that is exported by the NFS server. More info: http://kubernetes.io/docs/user-guide/volumes#nfs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/daemon_set_v1#path DaemonSetV1#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Server is the hostname or IP address of the NFS server. More info: http://kubernetes.io/docs/user-guide/volumes#nfs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/daemon_set_v1#server DaemonSetV1#server}
	Server *string `field:"required" json:"server" yaml:"server"`
	// Whether to force the NFS export to be mounted with read-only permissions. Defaults to false. More info: http://kubernetes.io/docs/user-guide/volumes#nfs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/daemon_set_v1#read_only DaemonSetV1#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}


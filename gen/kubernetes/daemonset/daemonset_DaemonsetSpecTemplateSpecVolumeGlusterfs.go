package daemonset


type DaemonsetSpecTemplateSpecVolumeGlusterfs struct {
	// The endpoint name that details Glusterfs topology. More info: http://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#endpoints_name Daemonset#endpoints_name}
	EndpointsName *string `field:"required" json:"endpointsName" yaml:"endpointsName"`
	// The Glusterfs volume path. More info: http://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#path Daemonset#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Whether to force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: http://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#read_only Daemonset#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}


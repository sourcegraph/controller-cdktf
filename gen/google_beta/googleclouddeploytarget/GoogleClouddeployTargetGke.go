package googleclouddeploytarget


type GoogleClouddeployTargetGke struct {
	// Information specifying a GKE Cluster. Format is `projects/{project_id}/locations/{location_id}/clusters/{cluster_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_clouddeploy_target#cluster GoogleClouddeployTarget#cluster}
	Cluster *string `field:"optional" json:"cluster" yaml:"cluster"`
	// Optional.
	//
	// If true, `cluster` is accessed using the private IP address of the control plane endpoint. Otherwise, the default IP address of the control plane endpoint is used. The default IP address is the private IP address for clusters with private control-plane endpoints and the public IP address otherwise. Only specify this option when `cluster` is a [private GKE cluster](https://cloud.google.com/kubernetes-engine/docs/concepts/private-cluster-concept).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_clouddeploy_target#internal_ip GoogleClouddeployTarget#internal_ip}
	InternalIp interface{} `field:"optional" json:"internalIp" yaml:"internalIp"`
	// Optional. If set, used to configure a [proxy](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/#proxy) to the Kubernetes server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_clouddeploy_target#proxy_url GoogleClouddeployTarget#proxy_url}
	ProxyUrl *string `field:"optional" json:"proxyUrl" yaml:"proxyUrl"`
}


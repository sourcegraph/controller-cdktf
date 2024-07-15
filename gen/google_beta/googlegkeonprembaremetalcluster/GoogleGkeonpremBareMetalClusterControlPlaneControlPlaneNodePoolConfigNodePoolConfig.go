package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfig struct {
	// The map of Kubernetes labels (key/value pairs) to be applied to each node.
	//
	// These will added in addition to any default label(s)
	// that Kubernetes may apply to the node. In case of conflict in
	// label keys, the applied set may differ depending on the Kubernetes
	// version -- it's best to assume the behavior is undefined and
	// conflicts should be avoided. For more information, including usage
	// and the valid values, see:
	//   http://kubernetes.io/v1.1/docs/user-guide/labels.html
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#labels GoogleGkeonpremBareMetalCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// node_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#node_configs GoogleGkeonpremBareMetalCluster#node_configs}
	NodeConfigs interface{} `field:"optional" json:"nodeConfigs" yaml:"nodeConfigs"`
	// Specifies the nodes operating system (default: LINUX).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#operating_system GoogleGkeonpremBareMetalCluster#operating_system}
	OperatingSystem *string `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// taints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#taints GoogleGkeonpremBareMetalCluster#taints}
	Taints interface{} `field:"optional" json:"taints" yaml:"taints"`
}


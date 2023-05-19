package clusterrole


type ClusterRoleMetadata struct {
	// An unstructured key value map stored with the clusterRole that may be used to store arbitrary metadata.
	//
	// More info: http://kubernetes.io/docs/user-guide/annotations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/cluster_role#annotations ClusterRole#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) the clusterRole.
	//
	// May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/cluster_role#labels ClusterRole#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Name of the clusterRole, must be unique. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/cluster_role#name ClusterRole#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}


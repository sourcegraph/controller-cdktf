package googleapigeeenvironment


type GoogleApigeeEnvironmentNodeConfig struct {
	// The maximum total number of gateway nodes that the is reserved for all instances that has the specified environment.
	//
	// If not specified, the default is determined by the
	// recommended maximum number of nodes for that gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_environment#max_node_count GoogleApigeeEnvironment#max_node_count}
	MaxNodeCount *string `field:"optional" json:"maxNodeCount" yaml:"maxNodeCount"`
	// The minimum total number of gateway nodes that the is reserved for all instances that has the specified environment.
	//
	// If not specified, the default is determined by the
	// recommended minimum number of nodes for that gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_environment#min_node_count GoogleApigeeEnvironment#min_node_count}
	MinNodeCount *string `field:"optional" json:"minNodeCount" yaml:"minNodeCount"`
}


package googlecontainercluster


type GoogleContainerClusterRbacBindingConfig struct {
	// Setting this to true will allow any ClusterRoleBinding and RoleBinding with subjects system:authenticated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#enable_insecure_binding_system_authenticated GoogleContainerCluster#enable_insecure_binding_system_authenticated}
	EnableInsecureBindingSystemAuthenticated interface{} `field:"optional" json:"enableInsecureBindingSystemAuthenticated" yaml:"enableInsecureBindingSystemAuthenticated"`
	// Setting this to true will allow any ClusterRoleBinding and RoleBinding with subjects system:anonymous or system:unauthenticated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#enable_insecure_binding_system_unauthenticated GoogleContainerCluster#enable_insecure_binding_system_unauthenticated}
	EnableInsecureBindingSystemUnauthenticated interface{} `field:"optional" json:"enableInsecureBindingSystemUnauthenticated" yaml:"enableInsecureBindingSystemUnauthenticated"`
}


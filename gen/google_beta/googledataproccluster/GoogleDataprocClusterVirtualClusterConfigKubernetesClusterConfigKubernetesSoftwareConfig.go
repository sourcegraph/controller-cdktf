package googledataproccluster


type GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig struct {
	// The components that should be installed in this Dataproc cluster.
	//
	// The key must be a string from the KubernetesComponent enumeration. The value is the version of the software to be installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#component_version GoogleDataprocCluster#component_version}
	ComponentVersion *map[string]*string `field:"required" json:"componentVersion" yaml:"componentVersion"`
	// The properties to set on daemon config files. Property keys are specified in prefix:property format, for example spark:spark.kubernetes.container.image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#properties GoogleDataprocCluster#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}


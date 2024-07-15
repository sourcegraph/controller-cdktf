package googledataproccluster


type GoogleDataprocClusterClusterConfigSoftwareConfig struct {
	// The Cloud Dataproc image version to use for the cluster - this controls the sets of software versions installed onto the nodes when you create clusters.
	//
	// If not specified, defaults to the latest version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#image_version GoogleDataprocCluster#image_version}
	ImageVersion *string `field:"optional" json:"imageVersion" yaml:"imageVersion"`
	// The set of optional components to activate on the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#optional_components GoogleDataprocCluster#optional_components}
	OptionalComponents *[]*string `field:"optional" json:"optionalComponents" yaml:"optionalComponents"`
	// A list of override and additional properties (key/value pairs) used to modify various aspects of the common configuration files used when creating a cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#override_properties GoogleDataprocCluster#override_properties}
	OverrideProperties *map[string]*string `field:"optional" json:"overrideProperties" yaml:"overrideProperties"`
}


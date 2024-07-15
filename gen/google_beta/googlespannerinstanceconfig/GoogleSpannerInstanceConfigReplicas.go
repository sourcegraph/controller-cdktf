package googlespannerinstanceconfig


type GoogleSpannerInstanceConfigReplicas struct {
	// If true, this location is designated as the default leader location where leader replicas are placed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_instance_config#default_leader_location GoogleSpannerInstanceConfigA#default_leader_location}
	DefaultLeaderLocation interface{} `field:"optional" json:"defaultLeaderLocation" yaml:"defaultLeaderLocation"`
	// The location of the serving resources, e.g. "us-central1".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_instance_config#location GoogleSpannerInstanceConfigA#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Indicates the type of replica.  See the [replica types documentation](https://cloud.google.com/spanner/docs/replication#replica_types) for more details. Possible values: ["READ_WRITE", "READ_ONLY", "WITNESS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_instance_config#type GoogleSpannerInstanceConfigA#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}


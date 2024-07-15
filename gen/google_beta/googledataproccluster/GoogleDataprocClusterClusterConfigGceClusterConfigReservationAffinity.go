package googledataproccluster


type GoogleDataprocClusterClusterConfigGceClusterConfigReservationAffinity struct {
	// Type of reservation to consume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#consume_reservation_type GoogleDataprocCluster#consume_reservation_type}
	ConsumeReservationType *string `field:"optional" json:"consumeReservationType" yaml:"consumeReservationType"`
	// Corresponds to the label key of reservation resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#key GoogleDataprocCluster#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Corresponds to the label values of reservation resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#values GoogleDataprocCluster#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}


package googledataprocjob


type GoogleDataprocJobPlacement struct {
	// The name of the cluster where the job will be submitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_dataproc_job#cluster_name GoogleDataprocJob#cluster_name}
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
}


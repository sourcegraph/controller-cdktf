package googledataprocjob


type GoogleDataprocJobPlacement struct {
	// The name of the cluster where the job will be submitted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_job#cluster_name GoogleDataprocJob#cluster_name}
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
}


package dataawsrdsclusters


type DataAwsRdsClustersFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/data-sources/rds_clusters#name DataAwsRdsClusters#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/data-sources/rds_clusters#values DataAwsRdsClusters#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}


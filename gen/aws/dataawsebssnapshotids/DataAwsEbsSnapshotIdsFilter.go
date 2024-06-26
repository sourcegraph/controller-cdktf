package dataawsebssnapshotids


type DataAwsEbsSnapshotIdsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/data-sources/ebs_snapshot_ids#name DataAwsEbsSnapshotIds#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/data-sources/ebs_snapshot_ids#values DataAwsEbsSnapshotIds#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}


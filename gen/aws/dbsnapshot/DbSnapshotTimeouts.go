package dbsnapshot


type DbSnapshotTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/db_snapshot#read DbSnapshot#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


package googlefilestoreinstance


type GoogleFilestoreInstanceFileShares struct {
	// File share capacity in GiB.
	//
	// This must be at least 1024 GiB
	// for the standard tier, or 2560 GiB for the premium tier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_filestore_instance#capacity_gb GoogleFilestoreInstance#capacity_gb}
	CapacityGb *float64 `field:"required" json:"capacityGb" yaml:"capacityGb"`
	// The name of the fileshare (16 characters or less).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_filestore_instance#name GoogleFilestoreInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// nfs_export_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_filestore_instance#nfs_export_options GoogleFilestoreInstance#nfs_export_options}
	NfsExportOptions interface{} `field:"optional" json:"nfsExportOptions" yaml:"nfsExportOptions"`
}


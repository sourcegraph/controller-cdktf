package googlednsrecordset


type GoogleDnsRecordSetRoutingPolicyPrimaryBackup struct {
	// backup_geo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_record_set#backup_geo GoogleDnsRecordSet#backup_geo}
	BackupGeo interface{} `field:"required" json:"backupGeo" yaml:"backupGeo"`
	// primary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_record_set#primary GoogleDnsRecordSet#primary}
	Primary *GoogleDnsRecordSetRoutingPolicyPrimaryBackupPrimary `field:"required" json:"primary" yaml:"primary"`
	// Specifies whether to enable fencing for backup geo queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_record_set#enable_geo_fencing_for_backups GoogleDnsRecordSet#enable_geo_fencing_for_backups}
	EnableGeoFencingForBackups interface{} `field:"optional" json:"enableGeoFencingForBackups" yaml:"enableGeoFencingForBackups"`
	// Specifies the percentage of traffic to send to the backup targets even when the primary targets are healthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_record_set#trickle_ratio GoogleDnsRecordSet#trickle_ratio}
	TrickleRatio *float64 `field:"optional" json:"trickleRatio" yaml:"trickleRatio"`
}


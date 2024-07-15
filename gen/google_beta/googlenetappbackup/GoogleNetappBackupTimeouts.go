package googlenetappbackup


type GoogleNetappBackupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_backup#create GoogleNetappBackup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_backup#delete GoogleNetappBackup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_backup#update GoogleNetappBackup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


package backupframework


type BackupFrameworkControlInputParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/backup_framework#name BackupFramework#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/backup_framework#value BackupFramework#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}


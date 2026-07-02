package backupselection

type BackupSelectionCondition struct {
	// string_equals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/backup_selection#string_equals BackupSelection#string_equals}
	StringEquals any `field:"optional" json:"stringEquals" yaml:"stringEquals"`
	// string_like block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/backup_selection#string_like BackupSelection#string_like}
	StringLike any `field:"optional" json:"stringLike" yaml:"stringLike"`
	// string_not_equals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/backup_selection#string_not_equals BackupSelection#string_not_equals}
	StringNotEquals any `field:"optional" json:"stringNotEquals" yaml:"stringNotEquals"`
	// string_not_like block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/backup_selection#string_not_like BackupSelection#string_not_like}
	StringNotLike any `field:"optional" json:"stringNotLike" yaml:"stringNotLike"`
}

package googlealloydbbackup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleAlloydbBackupConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the alloydb backup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_alloydb_backup#backup_id GoogleAlloydbBackup#backup_id}
	BackupId *string `field:"required" json:"backupId" yaml:"backupId"`
	// The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_alloydb_backup#cluster_name GoogleAlloydbBackup#cluster_name}
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The location where the alloydb backup should reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_alloydb_backup#location GoogleAlloydbBackup#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// User-provided description of the backup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_alloydb_backup#description GoogleAlloydbBackup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_alloydb_backup#encryption_config GoogleAlloydbBackup#encryption_config}
	EncryptionConfig *GoogleAlloydbBackupEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_alloydb_backup#id GoogleAlloydbBackup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User-defined labels for the alloydb backup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_alloydb_backup#labels GoogleAlloydbBackup#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_alloydb_backup#project GoogleAlloydbBackup#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_alloydb_backup#timeouts GoogleAlloydbBackup#timeouts}
	Timeouts *GoogleAlloydbBackupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


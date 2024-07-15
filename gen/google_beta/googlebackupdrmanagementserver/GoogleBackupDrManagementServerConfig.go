package googlebackupdrmanagementserver

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleBackupDrManagementServerConfig struct {
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
	// The location for the management server (management console).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_backup_dr_management_server#location GoogleBackupDrManagementServer#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of management server (management console).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_backup_dr_management_server#name GoogleBackupDrManagementServer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_backup_dr_management_server#networks GoogleBackupDrManagementServer#networks}
	Networks interface{} `field:"required" json:"networks" yaml:"networks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_backup_dr_management_server#id GoogleBackupDrManagementServer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_backup_dr_management_server#project GoogleBackupDrManagementServer#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_backup_dr_management_server#timeouts GoogleBackupDrManagementServer#timeouts}
	Timeouts *GoogleBackupDrManagementServerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The type of management server (management console). Default value: "BACKUP_RESTORE" Possible values: ["BACKUP_RESTORE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_backup_dr_management_server#type GoogleBackupDrManagementServer#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}


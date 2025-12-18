package grant

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GrantConfig struct {
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
	// The role to grant.
	//
	// Accepted values: `administrator`, `apitoken_creator`, `bookmark_manager`, `dashboard_creator`, `dashboard_editor`, `dashboard_viewer`, `dataset_accelerator`, `dataset_creator`, `dataset_editor`, `dataset_viewer`, `datastream_creator`, `datastream_editor`, `datastream_viewer`, `investigator_global`, `monitor_creator`, `monitor_editor`, `monitor_viewer`, `monitor_action_creator`, `monitor_global_muter`, `reference_table_creator`, `report_manager`, `service_account_creator`, `user_deleter`, `user_inviter`, `worksheet_creator`, `worksheet_editor`, `worksheet_viewer`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/grant#role Grant#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// OID of the subject. Must be a user or a group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/grant#subject Grant#subject}
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/grant#id Grant#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// qualifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/grant#qualifier Grant#qualifier}
	Qualifier *GrantQualifier `field:"optional" json:"qualifier" yaml:"qualifier"`
}


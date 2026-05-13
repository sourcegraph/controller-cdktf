package googlesqlprovisionscript

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleSqlProvisionScriptConfig struct {
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
	// The name of the Cloud SQL instance. Changing this forces the script to be run on the new instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_sql_provision_script#instance GoogleSqlProvisionScript#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// The SQL script to provision database resources.
	//
	// Its execution time limit is 30 s.
	// 				Changing this forces the script to be rerun. Make sure the script is idempotent.
	// 				You can use statements like "create if not exists …" or
	// 				"if not exists (select …) then … end if" to prevent existence-related errors. If it's not
	// 				possible to make a statement idempotent, you can run it once and then remove it from this script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_sql_provision_script#script GoogleSqlProvisionScript#script}
	Script *string `field:"required" json:"script" yaml:"script"`
	// The name of the database to which Terraform connects.
	//
	// Changing
	// 				this forces Terraform to connect to the new database and run the script. This argument is
	// 				required for Postgres instances. It's optional for MySQL, but some of your queries may require
	// 				a database. You can create and use a database in the script or explicitly reference a
	// 				google_sql_database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_sql_provision_script#database GoogleSqlProvisionScript#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// The deletion policy for the resources created by the script.
	//
	// The default is "ABANDON".
	// 				It must be "ABANDON" to allow Terraform to abandon the resources. If you want to delete resources, add statements
	// 				in the script such as "drop … if exists".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_sql_provision_script#deletion_policy GoogleSqlProvisionScript#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// The description of the provision script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_sql_provision_script#description GoogleSqlProvisionScript#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_sql_provision_script#id GoogleSqlProvisionScript#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the project in which the resource belongs.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_sql_provision_script#project GoogleSqlProvisionScript#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}


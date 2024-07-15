package googlealloydbuser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleAlloydbUserConfig struct {
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
	// Identifies the alloydb cluster. Must be in the format 'projects/{project}/locations/{location}/clusters/{cluster_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_user#cluster GoogleAlloydbUser#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The database role name of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_user#user_id GoogleAlloydbUser#user_id}
	UserId *string `field:"required" json:"userId" yaml:"userId"`
	// The type of this user. Possible values: ["ALLOYDB_BUILT_IN", "ALLOYDB_IAM_USER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_user#user_type GoogleAlloydbUser#user_type}
	UserType *string `field:"required" json:"userType" yaml:"userType"`
	// List of database roles this database user has.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_user#database_roles GoogleAlloydbUser#database_roles}
	DatabaseRoles *[]*string `field:"optional" json:"databaseRoles" yaml:"databaseRoles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_user#id GoogleAlloydbUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Password for this database user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_user#password GoogleAlloydbUser#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_user#timeouts GoogleAlloydbUser#timeouts}
	Timeouts *GoogleAlloydbUserTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


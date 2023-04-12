package defaultprivileges

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DefaultPrivilegesConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The database to grant default privileges for this role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/default_privileges#database DefaultPrivileges#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/default_privileges#object_type DefaultPrivileges#object_type}
	ObjectType *string `field:"required" json:"objectType" yaml:"objectType"`
	// Target role for which to alter default privileges.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/default_privileges#owner DefaultPrivileges#owner}
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The list of privileges to apply as default privileges.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/default_privileges#privileges DefaultPrivileges#privileges}
	Privileges *[]*string `field:"required" json:"privileges" yaml:"privileges"`
	// The name of the role to which grant default privileges on.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/default_privileges#role DefaultPrivileges#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/default_privileges#id DefaultPrivileges#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The database schema to set default privileges for this role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/default_privileges#schema DefaultPrivileges#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Permit the grant recipient to grant it to others.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/default_privileges#with_grant_option DefaultPrivileges#with_grant_option}
	WithGrantOption interface{} `field:"optional" json:"withGrantOption" yaml:"withGrantOption"`
}


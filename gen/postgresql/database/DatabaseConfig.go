package database

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseConfig struct {
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
	// The PostgreSQL database name to connect to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#name Database#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// If false then no one can connect to this database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#allow_connections Database#allow_connections}
	AllowConnections interface{} `field:"optional" json:"allowConnections" yaml:"allowConnections"`
	// How many concurrent connections can be made to this database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#connection_limit Database#connection_limit}
	ConnectionLimit *float64 `field:"optional" json:"connectionLimit" yaml:"connectionLimit"`
	// Character set encoding to use in the new database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#encoding Database#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#id Database#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If true, then this database can be cloned by any user with CREATEDB privileges.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#is_template Database#is_template}
	IsTemplate interface{} `field:"optional" json:"isTemplate" yaml:"isTemplate"`
	// Collation order (LC_COLLATE) to use in the new database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#lc_collate Database#lc_collate}
	LcCollate *string `field:"optional" json:"lcCollate" yaml:"lcCollate"`
	// Character classification (LC_CTYPE) to use in the new database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#lc_ctype Database#lc_ctype}
	LcCtype *string `field:"optional" json:"lcCtype" yaml:"lcCtype"`
	// The ROLE which owns the database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#owner Database#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// The name of the tablespace that will be associated with the new database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#tablespace_name Database#tablespace_name}
	TablespaceName *string `field:"optional" json:"tablespaceName" yaml:"tablespaceName"`
	// The name of the template from which to create the new database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#template Database#template}
	Template *string `field:"optional" json:"template" yaml:"template"`
}


package server

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServerConfig struct {
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
	// The name of the foreign-data wrapper that manages the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/server#fdw_name Server#fdw_name}
	FdwName *string `field:"required" json:"fdwName" yaml:"fdwName"`
	// The name of the foreign server to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/server#server_name Server#server_name}
	ServerName *string `field:"required" json:"serverName" yaml:"serverName"`
	// Automatically drop objects that depend on the server (such as user mappings), and in turn all objects that depend on those objects.
	//
	// Drop RESTRICT is the default
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/server#drop_cascade Server#drop_cascade}
	DropCascade interface{} `field:"optional" json:"dropCascade" yaml:"dropCascade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/server#id Server#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// This clause specifies the options for the server.
	//
	// The options typically define the connection details of the server, but the actual names and values are dependent on the server's foreign-data wrapper
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/server#options Server#options}
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// The user name of the new owner of the foreign server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/server#server_owner Server#server_owner}
	ServerOwner *string `field:"optional" json:"serverOwner" yaml:"serverOwner"`
	// Optional server type, potentially useful to foreign-data wrappers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/server#server_type Server#server_type}
	ServerType *string `field:"optional" json:"serverType" yaml:"serverType"`
	// Optional server version, potentially useful to foreign-data wrappers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/server#server_version Server#server_version}
	ServerVersion *string `field:"optional" json:"serverVersion" yaml:"serverVersion"`
}


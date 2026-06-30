package rbacstatement

type RbacStatementObject struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/rbac_statement#all RbacStatement#all}.
	All any `field:"optional" json:"all" yaml:"all"`
	// The Observe ID for a folder.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/rbac_statement#folder RbacStatement#folder}
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// The Observe ID for an object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/rbac_statement#id RbacStatement#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of object. Can be provided along with `type`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/rbac_statement#name RbacStatement#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// True to bind to objects owned by the user. Can be provided along with `type`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/rbac_statement#owner RbacStatement#owner}
	Owner any `field:"optional" json:"owner" yaml:"owner"`
	// The type of object such as dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/rbac_statement#type RbacStatement#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The Observe ID for a workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/rbac_statement#workspace RbacStatement#workspace}
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
}

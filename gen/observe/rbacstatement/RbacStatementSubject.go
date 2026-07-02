package rbacstatement

type RbacStatementSubject struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/rbac_statement#all RbacStatement#all}.
	All any `field:"optional" json:"all" yaml:"all"`
	// OID of a RBAC Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/rbac_statement#group RbacStatement#group}
	Group *string `field:"optional" json:"group" yaml:"group"`
	// OID of a user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/rbac_statement#user RbacStatement#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

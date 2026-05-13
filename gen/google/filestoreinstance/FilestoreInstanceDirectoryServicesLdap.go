package filestoreinstance


type FilestoreInstanceDirectoryServicesLdap struct {
	// The LDAP domain name in the format of 'my-domain.com'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/filestore_instance#domain FilestoreInstance#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// The servers names are used for specifying the LDAP servers names.
	//
	// The LDAP servers names can come with two formats:
	// 1. DNS name, for example: 'ldap.example1.com', 'ldap.example2.com'.
	// 2. IP address, for example: '10.0.0.1', '10.0.0.2', '10.0.0.3'.
	// All servers names must be in the same format: either all DNS names or all
	// IP addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/filestore_instance#servers FilestoreInstance#servers}
	Servers *[]*string `field:"required" json:"servers" yaml:"servers"`
	// The groups Organizational Unit (OU) is optional.
	//
	// This parameter is a hint
	// to allow faster lookup in the LDAP namespace. In case that this parameter
	// is not provided, Filestore instance will query the whole LDAP namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/filestore_instance#groups_ou FilestoreInstance#groups_ou}
	GroupsOu *string `field:"optional" json:"groupsOu" yaml:"groupsOu"`
	// The users Organizational Unit (OU) is optional.
	//
	// This parameter is a hint
	// to allow faster lookup in the LDAP namespace. In case that this parameter
	// is not provided, Filestore instance will query the whole LDAP namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/filestore_instance#users_ou FilestoreInstance#users_ou}
	UsersOu *string `field:"optional" json:"usersOu" yaml:"usersOu"`
}


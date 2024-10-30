package dataoktausers


type DataOktaUsersSearch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/users#comparison DataOktaUsers#comparison}.
	Comparison *string `field:"optional" json:"comparison" yaml:"comparison"`
	// A raw search expression string.
	//
	// This requires the search feature be on. Please see Okta documentation on their filter API for users. https://developer.okta.com/docs/api/resources/users#list-users-with-search
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/users#expression DataOktaUsers#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Property name to search for.
	//
	// This requires the search feature be on. Please see Okta documentation on their filter API for users. https://developer.okta.com/docs/api/resources/users#list-users-with-search
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/users#name DataOktaUsers#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/users#value DataOktaUsers#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}


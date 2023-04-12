package dataoktauser


type DataOktaUserSearch struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/d/user#comparison DataOktaUser#comparison}.
	Comparison *string `field:"optional" json:"comparison" yaml:"comparison"`
	// A raw search expression string.
	//
	// This requires the search feature be on. Please see Okta documentation on their filter API for users. https://developer.okta.com/docs/api/resources/users#list-users-with-search
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/d/user#expression DataOktaUser#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Property name to search for.
	//
	// This requires the search feature be on. Please see Okta documentation on their filter API for users. https://developer.okta.com/docs/api/resources/users#list-users-with-search
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/d/user#name DataOktaUser#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/d/user#value DataOktaUser#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}


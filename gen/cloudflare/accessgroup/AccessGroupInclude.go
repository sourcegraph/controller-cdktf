package accessgroup


type AccessGroupInclude struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#any_valid_service_token AccessGroup#any_valid_service_token}.
	AnyValidServiceToken interface{} `field:"optional" json:"anyValidServiceToken" yaml:"anyValidServiceToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#auth_method AccessGroup#auth_method}.
	AuthMethod *string `field:"optional" json:"authMethod" yaml:"authMethod"`
	// azure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#azure AccessGroup#azure}
	Azure interface{} `field:"optional" json:"azure" yaml:"azure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#certificate AccessGroup#certificate}.
	Certificate interface{} `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#common_name AccessGroup#common_name}.
	CommonName *string `field:"optional" json:"commonName" yaml:"commonName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#device_posture AccessGroup#device_posture}.
	DevicePosture *[]*string `field:"optional" json:"devicePosture" yaml:"devicePosture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#email AccessGroup#email}.
	Email *[]*string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#email_domain AccessGroup#email_domain}.
	EmailDomain *[]*string `field:"optional" json:"emailDomain" yaml:"emailDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#everyone AccessGroup#everyone}.
	Everyone interface{} `field:"optional" json:"everyone" yaml:"everyone"`
	// external_evaluation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#external_evaluation AccessGroup#external_evaluation}
	ExternalEvaluation *AccessGroupIncludeExternalEvaluation `field:"optional" json:"externalEvaluation" yaml:"externalEvaluation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#geo AccessGroup#geo}.
	Geo *[]*string `field:"optional" json:"geo" yaml:"geo"`
	// github block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#github AccessGroup#github}
	Github interface{} `field:"optional" json:"github" yaml:"github"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#group AccessGroup#group}.
	Group *[]*string `field:"optional" json:"group" yaml:"group"`
	// gsuite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#gsuite AccessGroup#gsuite}
	Gsuite interface{} `field:"optional" json:"gsuite" yaml:"gsuite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#ip AccessGroup#ip}.
	Ip *[]*string `field:"optional" json:"ip" yaml:"ip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#ip_list AccessGroup#ip_list}.
	IpList *[]*string `field:"optional" json:"ipList" yaml:"ipList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#login_method AccessGroup#login_method}.
	LoginMethod *[]*string `field:"optional" json:"loginMethod" yaml:"loginMethod"`
	// okta block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#okta AccessGroup#okta}
	Okta interface{} `field:"optional" json:"okta" yaml:"okta"`
	// saml block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#saml AccessGroup#saml}
	Saml interface{} `field:"optional" json:"saml" yaml:"saml"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#service_token AccessGroup#service_token}.
	ServiceToken *[]*string `field:"optional" json:"serviceToken" yaml:"serviceToken"`
}


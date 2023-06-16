package accesspolicy


type AccessPolicyExclude struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#any_valid_service_token AccessPolicy#any_valid_service_token}.
	AnyValidServiceToken interface{} `field:"optional" json:"anyValidServiceToken" yaml:"anyValidServiceToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#auth_method AccessPolicy#auth_method}.
	AuthMethod *string `field:"optional" json:"authMethod" yaml:"authMethod"`
	// azure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#azure AccessPolicy#azure}
	Azure interface{} `field:"optional" json:"azure" yaml:"azure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#certificate AccessPolicy#certificate}.
	Certificate interface{} `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#common_name AccessPolicy#common_name}.
	CommonName *string `field:"optional" json:"commonName" yaml:"commonName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#device_posture AccessPolicy#device_posture}.
	DevicePosture *[]*string `field:"optional" json:"devicePosture" yaml:"devicePosture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#email AccessPolicy#email}.
	Email *[]*string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#email_domain AccessPolicy#email_domain}.
	EmailDomain *[]*string `field:"optional" json:"emailDomain" yaml:"emailDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#everyone AccessPolicy#everyone}.
	Everyone interface{} `field:"optional" json:"everyone" yaml:"everyone"`
	// external_evaluation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#external_evaluation AccessPolicy#external_evaluation}
	ExternalEvaluation *AccessPolicyExcludeExternalEvaluation `field:"optional" json:"externalEvaluation" yaml:"externalEvaluation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#geo AccessPolicy#geo}.
	Geo *[]*string `field:"optional" json:"geo" yaml:"geo"`
	// github block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#github AccessPolicy#github}
	Github interface{} `field:"optional" json:"github" yaml:"github"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#group AccessPolicy#group}.
	Group *[]*string `field:"optional" json:"group" yaml:"group"`
	// gsuite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#gsuite AccessPolicy#gsuite}
	Gsuite interface{} `field:"optional" json:"gsuite" yaml:"gsuite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#ip AccessPolicy#ip}.
	Ip *[]*string `field:"optional" json:"ip" yaml:"ip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#ip_list AccessPolicy#ip_list}.
	IpList *[]*string `field:"optional" json:"ipList" yaml:"ipList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#login_method AccessPolicy#login_method}.
	LoginMethod *[]*string `field:"optional" json:"loginMethod" yaml:"loginMethod"`
	// okta block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#okta AccessPolicy#okta}
	Okta interface{} `field:"optional" json:"okta" yaml:"okta"`
	// saml block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#saml AccessPolicy#saml}
	Saml interface{} `field:"optional" json:"saml" yaml:"saml"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#service_token AccessPolicy#service_token}.
	ServiceToken *[]*string `field:"optional" json:"serviceToken" yaml:"serviceToken"`
}


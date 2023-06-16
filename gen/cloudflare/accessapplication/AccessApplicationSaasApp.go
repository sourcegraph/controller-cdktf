package accessapplication


type AccessApplicationSaasApp struct {
	// The service provider's endpoint that is responsible for receiving and parsing a SAML assertion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#consumer_service_url AccessApplication#consumer_service_url}
	ConsumerServiceUrl *string `field:"required" json:"consumerServiceUrl" yaml:"consumerServiceUrl"`
	// A globally unique name for an identity or service provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#sp_entity_id AccessApplication#sp_entity_id}
	SpEntityId *string `field:"required" json:"spEntityId" yaml:"spEntityId"`
	// The format of the name identifier sent to the SaaS application. Defaults to `email`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#name_id_format AccessApplication#name_id_format}
	NameIdFormat *string `field:"optional" json:"nameIdFormat" yaml:"nameIdFormat"`
}


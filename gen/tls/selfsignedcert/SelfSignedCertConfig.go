package selfsignedcert

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SelfSignedCertConfig struct {
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
	// List of key usages allowed for the issued certificate.
	//
	// Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`, `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`, `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`, `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`, `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/self_signed_cert#allowed_uses SelfSignedCert#allowed_uses}
	AllowedUses *[]*string `field:"required" json:"allowedUses" yaml:"allowedUses"`
	// Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong to. This can be read from a separate file using the [`file`](https://www.terraform.io/language/functions/file) interpolation function. Only an irreversible secure hash of the private key will be stored in the Terraform state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/self_signed_cert#private_key_pem SelfSignedCert#private_key_pem}
	PrivateKeyPem *string `field:"required" json:"privateKeyPem" yaml:"privateKeyPem"`
	// Number of hours, after initial issuing, that the certificate will remain valid for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/self_signed_cert#validity_period_hours SelfSignedCert#validity_period_hours}
	ValidityPeriodHours *float64 `field:"required" json:"validityPeriodHours" yaml:"validityPeriodHours"`
	// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/self_signed_cert#dns_names SelfSignedCert#dns_names}
	DnsNames *[]*string `field:"optional" json:"dnsNames" yaml:"dnsNames"`
	// The resource will consider the certificate to have expired the given number of hours before its actual expiry time.
	//
	// This can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the early renewal period. (default: `0`)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/self_signed_cert#early_renewal_hours SelfSignedCert#early_renewal_hours}
	EarlyRenewalHours *float64 `field:"optional" json:"earlyRenewalHours" yaml:"earlyRenewalHours"`
	// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/self_signed_cert#ip_addresses SelfSignedCert#ip_addresses}
	IpAddresses *[]*string `field:"optional" json:"ipAddresses" yaml:"ipAddresses"`
	// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/self_signed_cert#is_ca_certificate SelfSignedCert#is_ca_certificate}
	IsCaCertificate interface{} `field:"optional" json:"isCaCertificate" yaml:"isCaCertificate"`
	// Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/self_signed_cert#set_authority_key_id SelfSignedCert#set_authority_key_id}
	SetAuthorityKeyId interface{} `field:"optional" json:"setAuthorityKeyId" yaml:"setAuthorityKeyId"`
	// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/self_signed_cert#set_subject_key_id SelfSignedCert#set_subject_key_id}
	SetSubjectKeyId interface{} `field:"optional" json:"setSubjectKeyId" yaml:"setSubjectKeyId"`
	// subject block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/self_signed_cert#subject SelfSignedCert#subject}
	Subject *SelfSignedCertSubject `field:"optional" json:"subject" yaml:"subject"`
	// List of URIs for which a certificate is being requested (i.e. certificate subjects).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/self_signed_cert#uris SelfSignedCert#uris}
	Uris *[]*string `field:"optional" json:"uris" yaml:"uris"`
}


package locallysignedcert

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LocallySignedCertConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/locally_signed_cert#allowed_uses LocallySignedCert#allowed_uses}
	AllowedUses *[]*string `field:"required" json:"allowedUses" yaml:"allowedUses"`
	// Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/locally_signed_cert#ca_cert_pem LocallySignedCert#ca_cert_pem}
	CaCertPem *string `field:"required" json:"caCertPem" yaml:"caCertPem"`
	// Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/locally_signed_cert#ca_private_key_pem LocallySignedCert#ca_private_key_pem}
	CaPrivateKeyPem *string `field:"required" json:"caPrivateKeyPem" yaml:"caPrivateKeyPem"`
	// Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/locally_signed_cert#cert_request_pem LocallySignedCert#cert_request_pem}
	CertRequestPem *string `field:"required" json:"certRequestPem" yaml:"certRequestPem"`
	// Number of hours, after initial issuing, that the certificate will remain valid for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/locally_signed_cert#validity_period_hours LocallySignedCert#validity_period_hours}
	ValidityPeriodHours *float64 `field:"required" json:"validityPeriodHours" yaml:"validityPeriodHours"`
	// The resource will consider the certificate to have expired the given number of hours before its actual expiry time.
	//
	// This can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the early renewal period. (default: `0`)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/locally_signed_cert#early_renewal_hours LocallySignedCert#early_renewal_hours}
	EarlyRenewalHours *float64 `field:"optional" json:"earlyRenewalHours" yaml:"earlyRenewalHours"`
	// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/locally_signed_cert#is_ca_certificate LocallySignedCert#is_ca_certificate}
	IsCaCertificate interface{} `field:"optional" json:"isCaCertificate" yaml:"isCaCertificate"`
	// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/locally_signed_cert#set_subject_key_id LocallySignedCert#set_subject_key_id}
	SetSubjectKeyId interface{} `field:"optional" json:"setSubjectKeyId" yaml:"setSubjectKeyId"`
}


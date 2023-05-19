package certificatepack


type CertificatePackValidationRecords struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/certificate_pack#cname_name CertificatePack#cname_name}.
	CnameName *string `field:"optional" json:"cnameName" yaml:"cnameName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/certificate_pack#cname_target CertificatePack#cname_target}.
	CnameTarget *string `field:"optional" json:"cnameTarget" yaml:"cnameTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/certificate_pack#emails CertificatePack#emails}.
	Emails *[]*string `field:"optional" json:"emails" yaml:"emails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/certificate_pack#http_body CertificatePack#http_body}.
	HttpBody *string `field:"optional" json:"httpBody" yaml:"httpBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/certificate_pack#http_url CertificatePack#http_url}.
	HttpUrl *string `field:"optional" json:"httpUrl" yaml:"httpUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/certificate_pack#txt_name CertificatePack#txt_name}.
	TxtName *string `field:"optional" json:"txtName" yaml:"txtName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/certificate_pack#txt_value CertificatePack#txt_value}.
	TxtValue *string `field:"optional" json:"txtValue" yaml:"txtValue"`
}


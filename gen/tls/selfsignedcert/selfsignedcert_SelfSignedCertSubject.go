package selfsignedcert


type SelfSignedCertSubject struct {
	// Distinguished name: `CN`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tls/r/self_signed_cert#common_name SelfSignedCert#common_name}
	CommonName *string `field:"optional" json:"commonName" yaml:"commonName"`
	// Distinguished name: `C`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tls/r/self_signed_cert#country SelfSignedCert#country}
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Distinguished name: `L`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tls/r/self_signed_cert#locality SelfSignedCert#locality}
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// Distinguished name: `O`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tls/r/self_signed_cert#organization SelfSignedCert#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// Distinguished name: `OU`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tls/r/self_signed_cert#organizational_unit SelfSignedCert#organizational_unit}
	OrganizationalUnit *string `field:"optional" json:"organizationalUnit" yaml:"organizationalUnit"`
	// Distinguished name: `PC`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tls/r/self_signed_cert#postal_code SelfSignedCert#postal_code}
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// Distinguished name: `ST`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tls/r/self_signed_cert#province SelfSignedCert#province}
	Province *string `field:"optional" json:"province" yaml:"province"`
	// Distinguished name: `SERIALNUMBER`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tls/r/self_signed_cert#serial_number SelfSignedCert#serial_number}
	SerialNumber *string `field:"optional" json:"serialNumber" yaml:"serialNumber"`
	// Distinguished name: `STREET`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tls/r/self_signed_cert#street_address SelfSignedCert#street_address}
	StreetAddress *[]*string `field:"optional" json:"streetAddress" yaml:"streetAddress"`
}


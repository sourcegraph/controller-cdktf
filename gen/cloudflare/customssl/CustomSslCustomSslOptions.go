package customssl


type CustomSslCustomSslOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/custom_ssl#bundle_method CustomSsl#bundle_method}.
	BundleMethod *string `field:"optional" json:"bundleMethod" yaml:"bundleMethod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/custom_ssl#certificate CustomSsl#certificate}.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/custom_ssl#geo_restrictions CustomSsl#geo_restrictions}.
	GeoRestrictions *string `field:"optional" json:"geoRestrictions" yaml:"geoRestrictions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/custom_ssl#private_key CustomSsl#private_key}.
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/custom_ssl#type CustomSsl#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}


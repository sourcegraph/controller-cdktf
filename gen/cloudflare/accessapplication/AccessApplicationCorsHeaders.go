package accessapplication


type AccessApplicationCorsHeaders struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#allow_all_headers AccessApplication#allow_all_headers}.
	AllowAllHeaders interface{} `field:"optional" json:"allowAllHeaders" yaml:"allowAllHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#allow_all_methods AccessApplication#allow_all_methods}.
	AllowAllMethods interface{} `field:"optional" json:"allowAllMethods" yaml:"allowAllMethods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#allow_all_origins AccessApplication#allow_all_origins}.
	AllowAllOrigins interface{} `field:"optional" json:"allowAllOrigins" yaml:"allowAllOrigins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#allow_credentials AccessApplication#allow_credentials}.
	AllowCredentials interface{} `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#allowed_headers AccessApplication#allowed_headers}.
	AllowedHeaders *[]*string `field:"optional" json:"allowedHeaders" yaml:"allowedHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#allowed_methods AccessApplication#allowed_methods}.
	AllowedMethods *[]*string `field:"optional" json:"allowedMethods" yaml:"allowedMethods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#allowed_origins AccessApplication#allowed_origins}.
	AllowedOrigins *[]*string `field:"optional" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#max_age AccessApplication#max_age}.
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
}


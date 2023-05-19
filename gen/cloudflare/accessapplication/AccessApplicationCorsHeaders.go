package accessapplication


type AccessApplicationCorsHeaders struct {
	// Value to determine whether all HTTP headers are exposed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#allow_all_headers AccessApplication#allow_all_headers}
	AllowAllHeaders interface{} `field:"optional" json:"allowAllHeaders" yaml:"allowAllHeaders"`
	// Value to determine whether all methods are exposed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#allow_all_methods AccessApplication#allow_all_methods}
	AllowAllMethods interface{} `field:"optional" json:"allowAllMethods" yaml:"allowAllMethods"`
	// Value to determine whether all origins are permitted to make CORS requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#allow_all_origins AccessApplication#allow_all_origins}
	AllowAllOrigins interface{} `field:"optional" json:"allowAllOrigins" yaml:"allowAllOrigins"`
	// Value to determine if credentials (cookies, authorization headers, or TLS client certificates) are included with requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#allow_credentials AccessApplication#allow_credentials}
	AllowCredentials interface{} `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// List of HTTP headers to expose via CORS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#allowed_headers AccessApplication#allowed_headers}
	AllowedHeaders *[]*string `field:"optional" json:"allowedHeaders" yaml:"allowedHeaders"`
	// List of methods to expose via CORS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#allowed_methods AccessApplication#allowed_methods}
	AllowedMethods *[]*string `field:"optional" json:"allowedMethods" yaml:"allowedMethods"`
	// List of origins permitted to make CORS requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#allowed_origins AccessApplication#allowed_origins}
	AllowedOrigins *[]*string `field:"optional" json:"allowedOrigins" yaml:"allowedOrigins"`
	// The maximum time a preflight request will be cached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#max_age AccessApplication#max_age}
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
}


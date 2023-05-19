package pagesproject


type PagesProjectDeploymentConfigsPreview struct {
	// Use latest compatibility date for Pages Functions. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#always_use_latest_compatibility_date PagesProject#always_use_latest_compatibility_date}
	AlwaysUseLatestCompatibilityDate interface{} `field:"optional" json:"alwaysUseLatestCompatibilityDate" yaml:"alwaysUseLatestCompatibilityDate"`
	// Compatibility date used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#compatibility_date PagesProject#compatibility_date}
	CompatibilityDate *string `field:"optional" json:"compatibilityDate" yaml:"compatibilityDate"`
	// Compatibility flags used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#compatibility_flags PagesProject#compatibility_flags}
	CompatibilityFlags *[]*string `field:"optional" json:"compatibilityFlags" yaml:"compatibilityFlags"`
	// D1 Databases used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#d1_databases PagesProject#d1_databases}
	D1Databases *map[string]*string `field:"optional" json:"d1Databases" yaml:"d1Databases"`
	// Durable Object namespaces used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#durable_object_namespaces PagesProject#durable_object_namespaces}
	DurableObjectNamespaces *map[string]*string `field:"optional" json:"durableObjectNamespaces" yaml:"durableObjectNamespaces"`
	// Environment variables for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#environment_variables PagesProject#environment_variables}
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Fail open used for Pages Functions. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#fail_open PagesProject#fail_open}
	FailOpen interface{} `field:"optional" json:"failOpen" yaml:"failOpen"`
	// KV namespaces used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#kv_namespaces PagesProject#kv_namespaces}
	KvNamespaces *map[string]*string `field:"optional" json:"kvNamespaces" yaml:"kvNamespaces"`
	// R2 Buckets used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#r2_buckets PagesProject#r2_buckets}
	R2Buckets *map[string]*string `field:"optional" json:"r2Buckets" yaml:"r2Buckets"`
	// service_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#service_binding PagesProject#service_binding}
	ServiceBinding interface{} `field:"optional" json:"serviceBinding" yaml:"serviceBinding"`
	// Usage model used for Pages Functions. Defaults to `bundled`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#usage_model PagesProject#usage_model}
	UsageModel *string `field:"optional" json:"usageModel" yaml:"usageModel"`
}


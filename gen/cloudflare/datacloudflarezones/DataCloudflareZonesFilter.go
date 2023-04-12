package datacloudflarezones


type DataCloudflareZonesFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/d/zones#account_id DataCloudflareZones#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/d/zones#lookup_type DataCloudflareZones#lookup_type}.
	LookupType *string `field:"optional" json:"lookupType" yaml:"lookupType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/d/zones#match DataCloudflareZones#match}.
	Match *string `field:"optional" json:"match" yaml:"match"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/d/zones#name DataCloudflareZones#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/d/zones#paused DataCloudflareZones#paused}.
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/d/zones#status DataCloudflareZones#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}


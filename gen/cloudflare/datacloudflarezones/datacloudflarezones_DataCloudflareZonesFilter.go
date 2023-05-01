package datacloudflarezones


type DataCloudflareZonesFilter struct {
	// The account identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/d/zones#account_id DataCloudflareZones#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The type of search to perform for the `name` value when querying the zone API.
	//
	// Available values: `contains`, `exact`. Defaults to `exact`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/d/zones#lookup_type DataCloudflareZones#lookup_type}
	LookupType *string `field:"optional" json:"lookupType" yaml:"lookupType"`
	// A RE2 compatible regular expression to filter the	results.
	//
	// This is performed client side whereas the `name` and `lookup_type`	are performed on the Cloudflare server side.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/d/zones#match DataCloudflareZones#match}
	Match *string `field:"optional" json:"match" yaml:"match"`
	// A string value to search for.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/d/zones#name DataCloudflareZones#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Paused status of the zone to lookup. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/d/zones#paused DataCloudflareZones#paused}
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// Status of the zone to lookup.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/d/zones#status DataCloudflareZones#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}


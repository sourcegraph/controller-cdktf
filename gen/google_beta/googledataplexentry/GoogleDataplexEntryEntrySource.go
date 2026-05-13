package googledataplexentry


type GoogleDataplexEntryEntrySource struct {
	// ancestors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#ancestors GoogleDataplexEntry#ancestors}
	Ancestors interface{} `field:"optional" json:"ancestors" yaml:"ancestors"`
	// The time when the resource was created in the source system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#create_time GoogleDataplexEntry#create_time}
	CreateTime *string `field:"optional" json:"createTime" yaml:"createTime"`
	// A description of the data resource. Maximum length is 2,000 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#description GoogleDataplexEntry#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A user-friendly display name. Maximum length is 500 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#display_name GoogleDataplexEntry#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// User-defined labels.
	//
	// The maximum size of keys and values is 128 characters each.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#labels GoogleDataplexEntry#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The platform containing the source system. Maximum length is 64 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#platform GoogleDataplexEntry#platform}
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// The name of the resource in the source system. Maximum length is 4,000 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#resource GoogleDataplexEntry#resource}
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
	// The name of the source system. Maximum length is 64 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#system GoogleDataplexEntry#system}
	SystemAttribute *string `field:"optional" json:"systemAttribute" yaml:"systemAttribute"`
	// The time when the resource was last updated in the source system.
	//
	// If the entry exists in the system and its EntrySource has updateTime populated,
	// further updates to the EntrySource of the entry must provide incremental updates to its updateTime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#update_time GoogleDataplexEntry#update_time}
	UpdateTime *string `field:"optional" json:"updateTime" yaml:"updateTime"`
}


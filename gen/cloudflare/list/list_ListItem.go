package list


type ListItem struct {
	// value block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/list#value List#value}
	Value *ListItemValue `field:"required" json:"value" yaml:"value"`
	// An optional comment for the item.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/list#comment List#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}


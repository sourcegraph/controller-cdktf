package dataawsroutetables


type DataAwsRouteTablesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/data-sources/route_tables#name DataAwsRouteTables#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/data-sources/route_tables#values DataAwsRouteTables#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}


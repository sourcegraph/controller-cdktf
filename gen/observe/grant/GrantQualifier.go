package grant


type GrantQualifier struct {
	// OID of the object this grant applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/grant#oid Grant#oid}
	Oid *string `field:"optional" json:"oid" yaml:"oid"`
}


package waitingroomrules


type WaitingRoomRulesRules struct {
	// Action to perform in the ruleset rule. Available values: `bypass_waiting_room`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_rules#action WaitingRoomRules#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Criteria for an HTTP request to trigger the waiting room rule action.
	//
	// Uses the Firewall Rules expression language based on Wireshark display filters. Refer to the [Waiting Room Rules Docs](https://developers.cloudflare.com/waiting-room/additional-options/waiting-room-rules/bypass-rules/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_rules#expression WaitingRoomRules#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Brief summary of the waiting room rule and its intended use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_rules#description WaitingRoomRules#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the rule is enabled or disabled. Available values: `enabled`, `disabled`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_rules#status WaitingRoomRules#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}


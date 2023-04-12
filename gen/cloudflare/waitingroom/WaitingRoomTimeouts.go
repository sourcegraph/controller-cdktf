package waitingroom


type WaitingRoomTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/waiting_room#create WaitingRoom#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/waiting_room#update WaitingRoom#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


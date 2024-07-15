package googledialogflowcxpage


type GoogleDialogflowCxPageEntryFulfillmentMessagesTelephonyTransferCall struct {
	// Transfer the call to a phone number in E.164 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_page#phone_number GoogleDialogflowCxPage#phone_number}
	PhoneNumber *string `field:"required" json:"phoneNumber" yaml:"phoneNumber"`
}


package googledialogflowcxpage


type GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesPlayAudio struct {
	// URI of the audio clip.
	//
	// Dialogflow does not impose any validation on this value. It is specific to the client that reads it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_page#audio_uri GoogleDialogflowCxPage#audio_uri}
	AudioUri *string `field:"required" json:"audioUri" yaml:"audioUri"`
}


package googlefirebaseextensionsinstance


type GoogleFirebaseExtensionsInstanceConfigA struct {
	// The ref of the Extension from the Registry (e.g. publisher-id/awesome-extension).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_extensions_instance#extension_ref GoogleFirebaseExtensionsInstance#extension_ref}
	ExtensionRef *string `field:"required" json:"extensionRef" yaml:"extensionRef"`
	// Environment variables that may be configured for the Extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_extensions_instance#params GoogleFirebaseExtensionsInstance#params}
	Params *map[string]*string `field:"required" json:"params" yaml:"params"`
	// List of extension events selected by consumer that extension is allowed to emit, identified by their types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_extensions_instance#allowed_event_types GoogleFirebaseExtensionsInstance#allowed_event_types}
	AllowedEventTypes *[]*string `field:"optional" json:"allowedEventTypes" yaml:"allowedEventTypes"`
	// Fully qualified Eventarc resource name that consumers should use for event triggers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_extensions_instance#eventarc_channel GoogleFirebaseExtensionsInstance#eventarc_channel}
	EventarcChannel *string `field:"optional" json:"eventarcChannel" yaml:"eventarcChannel"`
	// The version of the Extension from the Registry (e.g. 1.0.3). If left blank, latest is assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_extensions_instance#extension_version GoogleFirebaseExtensionsInstance#extension_version}
	ExtensionVersion *string `field:"optional" json:"extensionVersion" yaml:"extensionVersion"`
	// Params whose values are only available at deployment time.
	//
	// Unlike other params, these will not be set as environment variables on
	// functions. See a full list of system parameters at
	// https://firebase.google.com/docs/extensions/publishers/parameters#system_parameters
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_extensions_instance#system_params GoogleFirebaseExtensionsInstance#system_params}
	SystemParams *map[string]*string `field:"optional" json:"systemParams" yaml:"systemParams"`
}


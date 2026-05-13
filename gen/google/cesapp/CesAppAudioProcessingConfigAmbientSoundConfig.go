package cesapp


type CesAppAudioProcessingConfigAmbientSoundConfig struct {
	// Ambient noise as a mono-channel, 16kHz WAV file stored in [Cloud Storage](https://cloud.google.com/storage). Note: Please make sure the CES service agent 'service-@gcp-sa-ces.iam.gserviceaccount.com' has 'storage.objects.get' permission to the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#gcs_uri CesApp#gcs_uri}
	GcsUri *string `field:"optional" json:"gcsUri" yaml:"gcsUri"`
	// Name of the prebuilt ambient sound.
	//
	// Valid values are: - "coffee_shop" - "keyboard" - "keypad" - "hum"
	// -"office_1" - "office_2" - "office_3"
	// -"room_1" - "room_2" - "room_3"
	// -"room_4" - "room_5" - "air_conditioner"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#prebuilt_ambient_sound CesApp#prebuilt_ambient_sound}
	PrebuiltAmbientSound *string `field:"optional" json:"prebuiltAmbientSound" yaml:"prebuiltAmbientSound"`
	// Volume gain (in dB) of the normal native volume supported by ambient noise, in the range [-96.0, 16.0]. If unset, or set to a value of 0.0 (dB), will play at normal native signal amplitude. A value of -6.0 (dB) will play at approximately half the amplitude of the normal native signal amplitude. A value of +6.0 (dB) will play at approximately twice the amplitude of the normal native signal amplitude. We strongly recommend not to exceed +10 (dB) as there's usually no effective increase in loudness for any value greater than that.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#volume_gain_db CesApp#volume_gain_db}
	VolumeGainDb *float64 `field:"optional" json:"volumeGainDb" yaml:"volumeGainDb"`
}


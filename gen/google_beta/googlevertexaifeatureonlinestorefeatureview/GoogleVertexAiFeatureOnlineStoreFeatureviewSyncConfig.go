package googlevertexaifeatureonlinestorefeatureview


type GoogleVertexAiFeatureOnlineStoreFeatureviewSyncConfig struct {
	// Cron schedule (https://en.wikipedia.org/wiki/Cron) to launch scheduled runs. To explicitly set a timezone to the cron tab, apply a prefix in the cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or "TZ=${IANA_TIME_ZONE}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store_featureview#cron GoogleVertexAiFeatureOnlineStoreFeatureview#cron}
	Cron *string `field:"optional" json:"cron" yaml:"cron"`
}


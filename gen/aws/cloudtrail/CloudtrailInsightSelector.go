package cloudtrail


type CloudtrailInsightSelector struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/cloudtrail#insight_type Cloudtrail#insight_type}.
	InsightType *string `field:"required" json:"insightType" yaml:"insightType"`
}


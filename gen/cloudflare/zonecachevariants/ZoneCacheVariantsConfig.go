package zonecachevariants

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZoneCacheVariantsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_cache_variants#zone_id ZoneCacheVariants#zone_id}.
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_cache_variants#avif ZoneCacheVariants#avif}.
	Avif *[]*string `field:"optional" json:"avif" yaml:"avif"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_cache_variants#bmp ZoneCacheVariants#bmp}.
	Bmp *[]*string `field:"optional" json:"bmp" yaml:"bmp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_cache_variants#gif ZoneCacheVariants#gif}.
	Gif *[]*string `field:"optional" json:"gif" yaml:"gif"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_cache_variants#id ZoneCacheVariants#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_cache_variants#jp2 ZoneCacheVariants#jp2}.
	Jp2 *[]*string `field:"optional" json:"jp2" yaml:"jp2"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_cache_variants#jpeg ZoneCacheVariants#jpeg}.
	Jpeg *[]*string `field:"optional" json:"jpeg" yaml:"jpeg"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_cache_variants#jpg ZoneCacheVariants#jpg}.
	Jpg *[]*string `field:"optional" json:"jpg" yaml:"jpg"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_cache_variants#jpg2 ZoneCacheVariants#jpg2}.
	Jpg2 *[]*string `field:"optional" json:"jpg2" yaml:"jpg2"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_cache_variants#png ZoneCacheVariants#png}.
	Png *[]*string `field:"optional" json:"png" yaml:"png"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_cache_variants#tif ZoneCacheVariants#tif}.
	Tif *[]*string `field:"optional" json:"tif" yaml:"tif"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_cache_variants#tiff ZoneCacheVariants#tiff}.
	Tiff *[]*string `field:"optional" json:"tiff" yaml:"tiff"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_cache_variants#webp ZoneCacheVariants#webp}.
	Webp *[]*string `field:"optional" json:"webp" yaml:"webp"`
}


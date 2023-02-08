package zonecachevariants

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/zonecachevariants/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_cache_variants cloudflare_zone_cache_variants}.
type ZoneCacheVariants interface {
	cdktf.TerraformResource
	Avif() *[]*string
	SetAvif(val *[]*string)
	AvifInput() *[]*string
	Bmp() *[]*string
	SetBmp(val *[]*string)
	BmpInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Gif() *[]*string
	SetGif(val *[]*string)
	GifInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Jp2() *[]*string
	SetJp2(val *[]*string)
	Jp2Input() *[]*string
	Jpeg() *[]*string
	SetJpeg(val *[]*string)
	JpegInput() *[]*string
	Jpg() *[]*string
	SetJpg(val *[]*string)
	Jpg2() *[]*string
	SetJpg2(val *[]*string)
	Jpg2Input() *[]*string
	JpgInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Png() *[]*string
	SetPng(val *[]*string)
	PngInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Tif() *[]*string
	SetTif(val *[]*string)
	Tiff() *[]*string
	SetTiff(val *[]*string)
	TiffInput() *[]*string
	TifInput() *[]*string
	Webp() *[]*string
	SetWebp(val *[]*string)
	WebpInput() *[]*string
	ZoneId() *string
	SetZoneId(val *string)
	ZoneIdInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAvif()
	ResetBmp()
	ResetGif()
	ResetId()
	ResetJp2()
	ResetJpeg()
	ResetJpg()
	ResetJpg2()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPng()
	ResetTif()
	ResetTiff()
	ResetWebp()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ZoneCacheVariants
type jsiiProxy_ZoneCacheVariants struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ZoneCacheVariants) Avif() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"avif",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) AvifInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"avifInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) Bmp() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bmp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) BmpInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bmpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) Gif() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gif",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) GifInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gifInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) Jp2() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jp2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) Jp2Input() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jp2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) Jpeg() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jpeg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) JpegInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jpegInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) Jpg() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jpg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) Jpg2() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jpg2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) Jpg2Input() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jpg2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) JpgInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jpgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) Png() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"png",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) PngInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pngInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) Tif() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tif",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) Tiff() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tiff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) TiffInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tiffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) TifInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tifInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) Webp() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"webp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) WebpInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"webpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneCacheVariants) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_cache_variants cloudflare_zone_cache_variants} Resource.
func NewZoneCacheVariants(scope constructs.Construct, id *string, config *ZoneCacheVariantsConfig) ZoneCacheVariants {
	_init_.Initialize()

	if err := validateNewZoneCacheVariantsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZoneCacheVariants{}

	_jsii_.Create(
		"cloudflare.zoneCacheVariants.ZoneCacheVariants",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_cache_variants cloudflare_zone_cache_variants} Resource.
func NewZoneCacheVariants_Override(z ZoneCacheVariants, scope constructs.Construct, id *string, config *ZoneCacheVariantsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.zoneCacheVariants.ZoneCacheVariants",
		[]interface{}{scope, id, config},
		z,
	)
}

func (j *jsiiProxy_ZoneCacheVariants)SetAvif(val *[]*string) {
	if err := j.validateSetAvifParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"avif",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariants)SetBmp(val *[]*string) {
	if err := j.validateSetBmpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bmp",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariants)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariants)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariants)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariants)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariants)SetGif(val *[]*string) {
	if err := j.validateSetGifParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gif",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariants)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariants)SetJp2(val *[]*string) {
	if err := j.validateSetJp2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jp2",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariants)SetJpeg(val *[]*string) {
	if err := j.validateSetJpegParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jpeg",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariants)SetJpg(val *[]*string) {
	if err := j.validateSetJpgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jpg",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariants)SetJpg2(val *[]*string) {
	if err := j.validateSetJpg2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jpg2",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariants)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariants)SetPng(val *[]*string) {
	if err := j.validateSetPngParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"png",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariants)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariants)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariants)SetTif(val *[]*string) {
	if err := j.validateSetTifParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tif",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariants)SetTiff(val *[]*string) {
	if err := j.validateSetTiffParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tiff",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariants)SetWebp(val *[]*string) {
	if err := j.validateSetWebpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webp",
		val,
	)
}

func (j *jsiiProxy_ZoneCacheVariants)SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ZoneCacheVariants_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZoneCacheVariants_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudflare.zoneCacheVariants.ZoneCacheVariants",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ZoneCacheVariants_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cloudflare.zoneCacheVariants.ZoneCacheVariants",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (z *jsiiProxy_ZoneCacheVariants) AddOverride(path *string, value interface{}) {
	if err := z.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (z *jsiiProxy_ZoneCacheVariants) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := z.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneCacheVariants) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := z.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneCacheVariants) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := z.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		z,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneCacheVariants) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := z.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		z,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneCacheVariants) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := z.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		z,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneCacheVariants) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := z.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		z,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneCacheVariants) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := z.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		z,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneCacheVariants) GetStringAttribute(terraformAttribute *string) *string {
	if err := z.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		z,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneCacheVariants) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := z.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		z,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneCacheVariants) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := z.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneCacheVariants) OverrideLogicalId(newLogicalId *string) {
	if err := z.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (z *jsiiProxy_ZoneCacheVariants) ResetAvif() {
	_jsii_.InvokeVoid(
		z,
		"resetAvif",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariants) ResetBmp() {
	_jsii_.InvokeVoid(
		z,
		"resetBmp",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariants) ResetGif() {
	_jsii_.InvokeVoid(
		z,
		"resetGif",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariants) ResetId() {
	_jsii_.InvokeVoid(
		z,
		"resetId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariants) ResetJp2() {
	_jsii_.InvokeVoid(
		z,
		"resetJp2",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariants) ResetJpeg() {
	_jsii_.InvokeVoid(
		z,
		"resetJpeg",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariants) ResetJpg() {
	_jsii_.InvokeVoid(
		z,
		"resetJpg",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariants) ResetJpg2() {
	_jsii_.InvokeVoid(
		z,
		"resetJpg2",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariants) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		z,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariants) ResetPng() {
	_jsii_.InvokeVoid(
		z,
		"resetPng",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariants) ResetTif() {
	_jsii_.InvokeVoid(
		z,
		"resetTif",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariants) ResetTiff() {
	_jsii_.InvokeVoid(
		z,
		"resetTiff",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariants) ResetWebp() {
	_jsii_.InvokeVoid(
		z,
		"resetWebp",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneCacheVariants) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneCacheVariants) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneCacheVariants) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneCacheVariants) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


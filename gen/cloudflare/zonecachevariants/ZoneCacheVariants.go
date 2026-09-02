package zonecachevariants

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/zonecachevariants/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_cache_variants cloudflare_zone_cache_variants}.
type ZoneCacheVariants interface {
	cdktn.TerraformResource
	Avif() *[]*string
	SetAvif(val *[]*string)
	AvifInput() *[]*string
	Bmp() *[]*string
	SetBmp(val *[]*string)
	BmpInput() *[]*string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
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
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Png() *[]*string
	SetPng(val *[]*string)
	PngInput() *[]*string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
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
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
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
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for ZoneCacheVariants
type jsiiProxy_ZoneCacheVariants struct {
	internal.Type__cdktnTerraformResource
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

func (j *jsiiProxy_ZoneCacheVariants) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
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

func (j *jsiiProxy_ZoneCacheVariants) Count() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_ZoneCacheVariants) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
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

func (j *jsiiProxy_ZoneCacheVariants) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
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

func (j *jsiiProxy_ZoneCacheVariants) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
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

func (j *jsiiProxy_ZoneCacheVariants) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
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


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_cache_variants cloudflare_zone_cache_variants} Resource.
func NewZoneCacheVariants(scope constructs.Construct, id *string, config *ZoneCacheVariantsConfig) ZoneCacheVariants {
	_init_.Initialize()

	if err := validateNewZoneCacheVariantsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZoneCacheVariants{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.zoneCacheVariants.ZoneCacheVariants",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_cache_variants cloudflare_zone_cache_variants} Resource.
func NewZoneCacheVariants_Override(z ZoneCacheVariants, scope constructs.Construct, id *string, config *ZoneCacheVariantsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.zoneCacheVariants.ZoneCacheVariants",
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

func (j *jsiiProxy_ZoneCacheVariants)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_ZoneCacheVariants)SetForEach(val cdktn.ITerraformIterator) {
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

func (j *jsiiProxy_ZoneCacheVariants)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
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

func (j *jsiiProxy_ZoneCacheVariants)SetProvider(val cdktn.TerraformProvider) {
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

// Generates CDKTN code for importing a ZoneCacheVariants resource upon running "cdktn plan <stack-name>".
func ZoneCacheVariants_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateZoneCacheVariants_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.zoneCacheVariants.ZoneCacheVariants",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
		"@cdktn/provider-cloudflare.zoneCacheVariants.ZoneCacheVariants",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZoneCacheVariants_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZoneCacheVariants_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.zoneCacheVariants.ZoneCacheVariants",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZoneCacheVariants_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZoneCacheVariants_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.zoneCacheVariants.ZoneCacheVariants",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ZoneCacheVariants_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-cloudflare.zoneCacheVariants.ZoneCacheVariants",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (z *jsiiProxy_ZoneCacheVariants) AddMoveTarget(moveTarget *string) {
	if err := z.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
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

func (z *jsiiProxy_ZoneCacheVariants) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := z.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

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

func (z *jsiiProxy_ZoneCacheVariants) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneCacheVariants) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := z.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (z *jsiiProxy_ZoneCacheVariants) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := z.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneCacheVariants) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := z.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		z,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneCacheVariants) MoveFromId(id *string) {
	if err := z.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveFromId",
		[]interface{}{id},
	)
}

func (z *jsiiProxy_ZoneCacheVariants) MoveTo(moveTarget *string, index interface{}) {
	if err := z.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (z *jsiiProxy_ZoneCacheVariants) MoveToId(id *string) {
	if err := z.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveToId",
		[]interface{}{id},
	)
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

func (z *jsiiProxy_ZoneCacheVariants) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := z.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
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

func (z *jsiiProxy_ZoneCacheVariants) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneCacheVariants) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toHclTerraform",
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

func (z *jsiiProxy_ZoneCacheVariants) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		z,
		"with",
		args,
		&returns,
	)

	return returns
}


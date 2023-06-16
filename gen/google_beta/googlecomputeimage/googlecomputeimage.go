package googlecomputeimage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeimage/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_image google_compute_image}.
type GoogleComputeImage interface {
	cdktf.TerraformResource
	ArchiveSizeBytes() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CreationTimestamp() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DiskSizeGb() *float64
	SetDiskSizeGb(val *float64)
	DiskSizeGbInput() *float64
	Family() *string
	SetFamily(val *string)
	FamilyInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GuestOsFeatures() GoogleComputeImageGuestOsFeaturesList
	GuestOsFeaturesInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImageEncryptionKey() GoogleComputeImageImageEncryptionKeyOutputReference
	ImageEncryptionKeyInput() *GoogleComputeImageImageEncryptionKey
	LabelFingerprint() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	Licenses() *[]*string
	SetLicenses(val *[]*string)
	LicensesInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	RawDisk() GoogleComputeImageRawDiskOutputReference
	RawDiskInput() *GoogleComputeImageRawDisk
	// Experimental.
	RawOverrides() interface{}
	SelfLink() *string
	SourceDisk() *string
	SetSourceDisk(val *string)
	SourceDiskInput() *string
	SourceImage() *string
	SetSourceImage(val *string)
	SourceImageInput() *string
	SourceSnapshot() *string
	SetSourceSnapshot(val *string)
	SourceSnapshotInput() *string
	StorageLocations() *[]*string
	SetStorageLocations(val *[]*string)
	StorageLocationsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeImageTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutGuestOsFeatures(value interface{})
	PutImageEncryptionKey(value *GoogleComputeImageImageEncryptionKey)
	PutRawDisk(value *GoogleComputeImageRawDisk)
	PutTimeouts(value *GoogleComputeImageTimeouts)
	ResetDescription()
	ResetDiskSizeGb()
	ResetFamily()
	ResetGuestOsFeatures()
	ResetId()
	ResetImageEncryptionKey()
	ResetLabels()
	ResetLicenses()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRawDisk()
	ResetSourceDisk()
	ResetSourceImage()
	ResetSourceSnapshot()
	ResetStorageLocations()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleComputeImage
type jsiiProxy_GoogleComputeImage struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputeImage) ArchiveSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"archiveSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) FamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"familyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) GuestOsFeatures() GoogleComputeImageGuestOsFeaturesList {
	var returns GoogleComputeImageGuestOsFeaturesList
	_jsii_.Get(
		j,
		"guestOsFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) GuestOsFeaturesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guestOsFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) ImageEncryptionKey() GoogleComputeImageImageEncryptionKeyOutputReference {
	var returns GoogleComputeImageImageEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"imageEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) ImageEncryptionKeyInput() *GoogleComputeImageImageEncryptionKey {
	var returns *GoogleComputeImageImageEncryptionKey
	_jsii_.Get(
		j,
		"imageEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) LabelFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) Licenses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"licenses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) LicensesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"licensesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) RawDisk() GoogleComputeImageRawDiskOutputReference {
	var returns GoogleComputeImageRawDiskOutputReference
	_jsii_.Get(
		j,
		"rawDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) RawDiskInput() *GoogleComputeImageRawDisk {
	var returns *GoogleComputeImageRawDisk
	_jsii_.Get(
		j,
		"rawDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) SourceDisk() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) SourceDiskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) SourceImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) SourceImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) SourceSnapshot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) SourceSnapshotInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) StorageLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"storageLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) StorageLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"storageLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) Timeouts() GoogleComputeImageTimeoutsOutputReference {
	var returns GoogleComputeImageTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImage) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_image google_compute_image} Resource.
func NewGoogleComputeImage(scope constructs.Construct, id *string, config *GoogleComputeImageConfig) GoogleComputeImage {
	_init_.Initialize()

	if err := validateNewGoogleComputeImageParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeImage{}

	_jsii_.Create(
		"google-beta.googleComputeImage.GoogleComputeImage",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_image google_compute_image} Resource.
func NewGoogleComputeImage_Override(g GoogleComputeImage, scope constructs.Construct, id *string, config *GoogleComputeImageConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeImage.GoogleComputeImage",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeImage)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImage)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImage)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImage)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImage)SetDiskSizeGb(val *float64) {
	if err := j.validateSetDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImage)SetFamily(val *string) {
	if err := j.validateSetFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"family",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImage)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImage)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImage)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImage)SetLicenses(val *[]*string) {
	if err := j.validateSetLicensesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenses",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImage)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImage)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImage)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImage)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImage)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImage)SetSourceDisk(val *string) {
	if err := j.validateSetSourceDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDisk",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImage)SetSourceImage(val *string) {
	if err := j.validateSetSourceImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceImage",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImage)SetSourceSnapshot(val *string) {
	if err := j.validateSetSourceSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceSnapshot",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImage)SetStorageLocations(val *[]*string) {
	if err := j.validateSetStorageLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageLocations",
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
func GoogleComputeImage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeImage_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeImage.GoogleComputeImage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeImage_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeImage_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeImage.GoogleComputeImage",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeImage_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeImage_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeImage.GoogleComputeImage",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeImage_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.googleComputeImage.GoogleComputeImage",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeImage) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeImage) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeImage) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeImage) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeImage) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeImage) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeImage) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeImage) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeImage) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeImage) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeImage) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeImage) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeImage) PutGuestOsFeatures(value interface{}) {
	if err := g.validatePutGuestOsFeaturesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGuestOsFeatures",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeImage) PutImageEncryptionKey(value *GoogleComputeImageImageEncryptionKey) {
	if err := g.validatePutImageEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putImageEncryptionKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeImage) PutRawDisk(value *GoogleComputeImageRawDisk) {
	if err := g.validatePutRawDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRawDisk",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeImage) PutTimeouts(value *GoogleComputeImageTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeImage) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeImage) ResetDiskSizeGb() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskSizeGb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeImage) ResetFamily() {
	_jsii_.InvokeVoid(
		g,
		"resetFamily",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeImage) ResetGuestOsFeatures() {
	_jsii_.InvokeVoid(
		g,
		"resetGuestOsFeatures",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeImage) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeImage) ResetImageEncryptionKey() {
	_jsii_.InvokeVoid(
		g,
		"resetImageEncryptionKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeImage) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeImage) ResetLicenses() {
	_jsii_.InvokeVoid(
		g,
		"resetLicenses",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeImage) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeImage) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeImage) ResetRawDisk() {
	_jsii_.InvokeVoid(
		g,
		"resetRawDisk",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeImage) ResetSourceDisk() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceDisk",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeImage) ResetSourceImage() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceImage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeImage) ResetSourceSnapshot() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceSnapshot",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeImage) ResetStorageLocations() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageLocations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeImage) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeImage) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeImage) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeImage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeImage) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


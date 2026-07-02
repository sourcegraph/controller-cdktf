package datakubernetesstorageclass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/datakubernetesstorageclass/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/data-sources/storage_class kubernetes_storage_class}.
type DataKubernetesStorageClass interface {
	cdktf.TerraformDataSource
	AllowedTopologies() DataKubernetesStorageClassAllowedTopologiesOutputReference
	AllowedTopologiesInput() *DataKubernetesStorageClassAllowedTopologies
	AllowVolumeExpansion() any
	SetAllowVolumeExpansion(val any)
	AllowVolumeExpansionInput() any
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Metadata() DataKubernetesStorageClassMetadataOutputReference
	MetadataInput() *DataKubernetesStorageClassMetadata
	MountOptions() *[]*string
	SetMountOptions(val *[]*string)
	MountOptionsInput() *[]*string
	// The tree node.
	Node() constructs.Node
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() any
	ReclaimPolicy() *string
	SetReclaimPolicy(val *string)
	ReclaimPolicyInput() *string
	StorageProvisioner() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	VolumeBindingMode() *string
	SetVolumeBindingMode(val *string)
	VolumeBindingModeInput() *string
	// Experimental.
	AddOverride(path *string, value any)
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutAllowedTopologies(value *DataKubernetesStorageClassAllowedTopologies)
	PutMetadata(value *DataKubernetesStorageClassMetadata)
	ResetAllowedTopologies()
	ResetAllowVolumeExpansion()
	ResetId()
	ResetMountOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
	ResetReclaimPolicy()
	ResetVolumeBindingMode()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for DataKubernetesStorageClass
type jsiiProxy_DataKubernetesStorageClass struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataKubernetesStorageClass) AllowedTopologies() DataKubernetesStorageClassAllowedTopologiesOutputReference {
	var returns DataKubernetesStorageClassAllowedTopologiesOutputReference
	_jsii_.Get(
		j,
		"allowedTopologies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) AllowedTopologiesInput() *DataKubernetesStorageClassAllowedTopologies {
	var returns *DataKubernetesStorageClassAllowedTopologies
	_jsii_.Get(
		j,
		"allowedTopologiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) AllowVolumeExpansion() any {
	var returns any
	_jsii_.Get(
		j,
		"allowVolumeExpansion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) AllowVolumeExpansionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowVolumeExpansionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) Metadata() DataKubernetesStorageClassMetadataOutputReference {
	var returns DataKubernetesStorageClassMetadataOutputReference
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) MetadataInput() *DataKubernetesStorageClassMetadata {
	var returns *DataKubernetesStorageClassMetadata
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) MountOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mountOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) MountOptionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mountOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) ReclaimPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reclaimPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) ReclaimPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reclaimPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) StorageProvisioner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProvisioner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) VolumeBindingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeBindingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesStorageClass) VolumeBindingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeBindingModeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/data-sources/storage_class kubernetes_storage_class} Data Source.
func NewDataKubernetesStorageClass(scope constructs.Construct, id *string, config *DataKubernetesStorageClassConfig) DataKubernetesStorageClass {
	_init_.Initialize()

	if err := validateNewDataKubernetesStorageClassParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataKubernetesStorageClass{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.dataKubernetesStorageClass.DataKubernetesStorageClass",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/data-sources/storage_class kubernetes_storage_class} Data Source.
func NewDataKubernetesStorageClass_Override(d DataKubernetesStorageClass, scope constructs.Construct, id *string, config *DataKubernetesStorageClassConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.dataKubernetesStorageClass.DataKubernetesStorageClass",
		[]any{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataKubernetesStorageClass) SetAllowVolumeExpansion(val any) {
	if err := j.validateSetAllowVolumeExpansionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowVolumeExpansion",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesStorageClass) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesStorageClass) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesStorageClass) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesStorageClass) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesStorageClass) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesStorageClass) SetMountOptions(val *[]*string) {
	if err := j.validateSetMountOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mountOptions",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesStorageClass) SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesStorageClass) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesStorageClass) SetReclaimPolicy(val *string) {
	if err := j.validateSetReclaimPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reclaimPolicy",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesStorageClass) SetVolumeBindingMode(val *string) {
	if err := j.validateSetVolumeBindingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeBindingMode",
		val,
	)
}

// Generates CDKTF code for importing a DataKubernetesStorageClass resource upon running "cdktf plan <stack-name>".
func DataKubernetesStorageClass_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataKubernetesStorageClass_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-kubernetes.dataKubernetesStorageClass.DataKubernetesStorageClass",
		"generateConfigForImport",
		[]any{scope, importToId, importFromId, provider},
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
func DataKubernetesStorageClass_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateDataKubernetesStorageClass_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-kubernetes.dataKubernetesStorageClass.DataKubernetesStorageClass",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataKubernetesStorageClass_IsTerraformDataSource(x any) *bool {
	_init_.Initialize()

	if err := validateDataKubernetesStorageClass_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-kubernetes.dataKubernetesStorageClass.DataKubernetesStorageClass",
		"isTerraformDataSource",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataKubernetesStorageClass_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateDataKubernetesStorageClass_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-kubernetes.dataKubernetesStorageClass.DataKubernetesStorageClass",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

func DataKubernetesStorageClass_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-kubernetes.dataKubernetesStorageClass.DataKubernetesStorageClass",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataKubernetesStorageClass) AddOverride(path *string, value any) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]any{path, value},
	)
}

func (d *jsiiProxy_DataKubernetesStorageClass) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesStorageClass) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesStorageClass) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesStorageClass) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesStorageClass) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesStorageClass) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesStorageClass) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesStorageClass) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesStorageClass) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesStorageClass) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesStorageClass) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (d *jsiiProxy_DataKubernetesStorageClass) PutAllowedTopologies(value *DataKubernetesStorageClassAllowedTopologies) {
	if err := d.validatePutAllowedTopologiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAllowedTopologies",
		[]any{value},
	)
}

func (d *jsiiProxy_DataKubernetesStorageClass) PutMetadata(value *DataKubernetesStorageClassMetadata) {
	if err := d.validatePutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMetadata",
		[]any{value},
	)
}

func (d *jsiiProxy_DataKubernetesStorageClass) ResetAllowedTopologies() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowedTopologies",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesStorageClass) ResetAllowVolumeExpansion() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowVolumeExpansion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesStorageClass) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesStorageClass) ResetMountOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetMountOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesStorageClass) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesStorageClass) ResetParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesStorageClass) ResetReclaimPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetReclaimPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesStorageClass) ResetVolumeBindingMode() {
	_jsii_.InvokeVoid(
		d,
		"resetVolumeBindingMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesStorageClass) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesStorageClass) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesStorageClass) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesStorageClass) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesStorageClass) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesStorageClass) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

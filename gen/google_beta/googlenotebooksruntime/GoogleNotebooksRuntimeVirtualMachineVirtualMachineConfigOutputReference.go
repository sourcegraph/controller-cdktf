package googlenotebooksruntime

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenotebooksruntime/internal"
)

type GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference interface {
	cdktf.ComplexObject
	AcceleratorConfig() GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigAcceleratorConfigOutputReference
	AcceleratorConfigInput() *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigAcceleratorConfig
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ContainerImages() GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList
	ContainerImagesInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataDisk() GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference
	DataDiskInput() *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk
	EncryptionConfig() GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigEncryptionConfigOutputReference
	EncryptionConfigInput() *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigEncryptionConfig
	// Experimental.
	Fqn() *string
	GuestAttributes() cdktf.StringMap
	InternalIpOnly() interface{}
	SetInternalIpOnly(val interface{})
	InternalIpOnlyInput() interface{}
	InternalValue() *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfig
	SetInternalValue(val *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfig)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	NicType() *string
	SetNicType(val *string)
	NicTypeInput() *string
	ReservedIpRange() *string
	SetReservedIpRange(val *string)
	ReservedIpRangeInput() *string
	ShieldedInstanceConfig() GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfig
	Subnet() *string
	SetSubnet(val *string)
	SubnetInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Zone() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutAcceleratorConfig(value *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigAcceleratorConfig)
	PutContainerImages(value interface{})
	PutDataDisk(value *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk)
	PutEncryptionConfig(value *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigEncryptionConfig)
	PutShieldedInstanceConfig(value *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfig)
	ResetAcceleratorConfig()
	ResetContainerImages()
	ResetEncryptionConfig()
	ResetInternalIpOnly()
	ResetLabels()
	ResetMetadata()
	ResetNetwork()
	ResetNicType()
	ResetReservedIpRange()
	ResetShieldedInstanceConfig()
	ResetSubnet()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference
type jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) AcceleratorConfig() GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigAcceleratorConfigOutputReference {
	var returns GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigAcceleratorConfigOutputReference
	_jsii_.Get(
		j,
		"acceleratorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) AcceleratorConfigInput() *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigAcceleratorConfig {
	var returns *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigAcceleratorConfig
	_jsii_.Get(
		j,
		"acceleratorConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ContainerImages() GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList {
	var returns GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList
	_jsii_.Get(
		j,
		"containerImages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ContainerImagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerImagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) DataDisk() GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference {
	var returns GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference
	_jsii_.Get(
		j,
		"dataDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) DataDiskInput() *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk {
	var returns *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk
	_jsii_.Get(
		j,
		"dataDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) EncryptionConfig() GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigEncryptionConfigOutputReference {
	var returns GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) EncryptionConfigInput() *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigEncryptionConfig {
	var returns *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) GuestAttributes() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"guestAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) InternalIpOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalIpOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) InternalIpOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalIpOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) InternalValue() *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfig {
	var returns *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) NicType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) NicTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ReservedIpRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservedIpRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ReservedIpRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservedIpRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ShieldedInstanceConfig() GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference {
	var returns GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ShieldedInstanceConfigInput() *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfig {
	var returns *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) Subnet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) SubnetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}


func NewGoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNotebooksRuntime.GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference_Override(g GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNotebooksRuntime.GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetInternalIpOnly(val interface{}) {
	if err := j.validateSetInternalIpOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalIpOnly",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetInternalValue(val *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetNicType(val *string) {
	if err := j.validateSetNicTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nicType",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetReservedIpRange(val *string) {
	if err := j.validateSetReservedIpRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedIpRange",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetSubnet(val *string) {
	if err := j.validateSetSubnetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnet",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) PutAcceleratorConfig(value *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigAcceleratorConfig) {
	if err := g.validatePutAcceleratorConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAcceleratorConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) PutContainerImages(value interface{}) {
	if err := g.validatePutContainerImagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putContainerImages",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) PutDataDisk(value *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk) {
	if err := g.validatePutDataDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataDisk",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) PutEncryptionConfig(value *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigEncryptionConfig) {
	if err := g.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) PutShieldedInstanceConfig(value *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfig) {
	if err := g.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetAcceleratorConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAcceleratorConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetContainerImages() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerImages",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetInternalIpOnly() {
	_jsii_.InvokeVoid(
		g,
		"resetInternalIpOnly",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetNicType() {
	_jsii_.InvokeVoid(
		g,
		"resetNicType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetReservedIpRange() {
	_jsii_.InvokeVoid(
		g,
		"resetReservedIpRange",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetSubnet() {
	_jsii_.InvokeVoid(
		g,
		"resetSubnet",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


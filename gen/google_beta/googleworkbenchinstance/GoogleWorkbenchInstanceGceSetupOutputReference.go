package googleworkbenchinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleworkbenchinstance/internal"
)

type GoogleWorkbenchInstanceGceSetupOutputReference interface {
	cdktf.ComplexObject
	AcceleratorConfigs() GoogleWorkbenchInstanceGceSetupAcceleratorConfigsList
	AcceleratorConfigsInput() interface{}
	BootDisk() GoogleWorkbenchInstanceGceSetupBootDiskOutputReference
	BootDiskInput() *GoogleWorkbenchInstanceGceSetupBootDisk
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
	ConfidentialInstanceConfig() GoogleWorkbenchInstanceGceSetupConfidentialInstanceConfigOutputReference
	ConfidentialInstanceConfigInput() *GoogleWorkbenchInstanceGceSetupConfidentialInstanceConfig
	ContainerImage() GoogleWorkbenchInstanceGceSetupContainerImageOutputReference
	ContainerImageInput() *GoogleWorkbenchInstanceGceSetupContainerImage
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataDisks() GoogleWorkbenchInstanceGceSetupDataDisksOutputReference
	DataDisksInput() *GoogleWorkbenchInstanceGceSetupDataDisks
	DisablePublicIp() interface{}
	SetDisablePublicIp(val interface{})
	DisablePublicIpInput() interface{}
	EnableIpForwarding() interface{}
	SetEnableIpForwarding(val interface{})
	EnableIpForwardingInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleWorkbenchInstanceGceSetup
	SetInternalValue(val *GoogleWorkbenchInstanceGceSetup)
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	NetworkInterfaces() GoogleWorkbenchInstanceGceSetupNetworkInterfacesList
	NetworkInterfacesInput() interface{}
	ServiceAccounts() GoogleWorkbenchInstanceGceSetupServiceAccountsList
	ServiceAccountsInput() interface{}
	ShieldedInstanceConfig() GoogleWorkbenchInstanceGceSetupShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *GoogleWorkbenchInstanceGceSetupShieldedInstanceConfig
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
	VmImage() GoogleWorkbenchInstanceGceSetupVmImageOutputReference
	VmImageInput() *GoogleWorkbenchInstanceGceSetupVmImage
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
	PutAcceleratorConfigs(value interface{})
	PutBootDisk(value *GoogleWorkbenchInstanceGceSetupBootDisk)
	PutConfidentialInstanceConfig(value *GoogleWorkbenchInstanceGceSetupConfidentialInstanceConfig)
	PutContainerImage(value *GoogleWorkbenchInstanceGceSetupContainerImage)
	PutDataDisks(value *GoogleWorkbenchInstanceGceSetupDataDisks)
	PutNetworkInterfaces(value interface{})
	PutServiceAccounts(value interface{})
	PutShieldedInstanceConfig(value *GoogleWorkbenchInstanceGceSetupShieldedInstanceConfig)
	PutVmImage(value *GoogleWorkbenchInstanceGceSetupVmImage)
	ResetAcceleratorConfigs()
	ResetBootDisk()
	ResetConfidentialInstanceConfig()
	ResetContainerImage()
	ResetDataDisks()
	ResetDisablePublicIp()
	ResetEnableIpForwarding()
	ResetMachineType()
	ResetMetadata()
	ResetNetworkInterfaces()
	ResetServiceAccounts()
	ResetShieldedInstanceConfig()
	ResetTags()
	ResetVmImage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleWorkbenchInstanceGceSetupOutputReference
type jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) AcceleratorConfigs() GoogleWorkbenchInstanceGceSetupAcceleratorConfigsList {
	var returns GoogleWorkbenchInstanceGceSetupAcceleratorConfigsList
	_jsii_.Get(
		j,
		"acceleratorConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) AcceleratorConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratorConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) BootDisk() GoogleWorkbenchInstanceGceSetupBootDiskOutputReference {
	var returns GoogleWorkbenchInstanceGceSetupBootDiskOutputReference
	_jsii_.Get(
		j,
		"bootDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) BootDiskInput() *GoogleWorkbenchInstanceGceSetupBootDisk {
	var returns *GoogleWorkbenchInstanceGceSetupBootDisk
	_jsii_.Get(
		j,
		"bootDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ConfidentialInstanceConfig() GoogleWorkbenchInstanceGceSetupConfidentialInstanceConfigOutputReference {
	var returns GoogleWorkbenchInstanceGceSetupConfidentialInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"confidentialInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ConfidentialInstanceConfigInput() *GoogleWorkbenchInstanceGceSetupConfidentialInstanceConfig {
	var returns *GoogleWorkbenchInstanceGceSetupConfidentialInstanceConfig
	_jsii_.Get(
		j,
		"confidentialInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ContainerImage() GoogleWorkbenchInstanceGceSetupContainerImageOutputReference {
	var returns GoogleWorkbenchInstanceGceSetupContainerImageOutputReference
	_jsii_.Get(
		j,
		"containerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ContainerImageInput() *GoogleWorkbenchInstanceGceSetupContainerImage {
	var returns *GoogleWorkbenchInstanceGceSetupContainerImage
	_jsii_.Get(
		j,
		"containerImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) DataDisks() GoogleWorkbenchInstanceGceSetupDataDisksOutputReference {
	var returns GoogleWorkbenchInstanceGceSetupDataDisksOutputReference
	_jsii_.Get(
		j,
		"dataDisks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) DataDisksInput() *GoogleWorkbenchInstanceGceSetupDataDisks {
	var returns *GoogleWorkbenchInstanceGceSetupDataDisks
	_jsii_.Get(
		j,
		"dataDisksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) DisablePublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) DisablePublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) EnableIpForwarding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIpForwarding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) EnableIpForwardingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIpForwardingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) InternalValue() *GoogleWorkbenchInstanceGceSetup {
	var returns *GoogleWorkbenchInstanceGceSetup
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) NetworkInterfaces() GoogleWorkbenchInstanceGceSetupNetworkInterfacesList {
	var returns GoogleWorkbenchInstanceGceSetupNetworkInterfacesList
	_jsii_.Get(
		j,
		"networkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) NetworkInterfacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ServiceAccounts() GoogleWorkbenchInstanceGceSetupServiceAccountsList {
	var returns GoogleWorkbenchInstanceGceSetupServiceAccountsList
	_jsii_.Get(
		j,
		"serviceAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ServiceAccountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ShieldedInstanceConfig() GoogleWorkbenchInstanceGceSetupShieldedInstanceConfigOutputReference {
	var returns GoogleWorkbenchInstanceGceSetupShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ShieldedInstanceConfigInput() *GoogleWorkbenchInstanceGceSetupShieldedInstanceConfig {
	var returns *GoogleWorkbenchInstanceGceSetupShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) VmImage() GoogleWorkbenchInstanceGceSetupVmImageOutputReference {
	var returns GoogleWorkbenchInstanceGceSetupVmImageOutputReference
	_jsii_.Get(
		j,
		"vmImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) VmImageInput() *GoogleWorkbenchInstanceGceSetupVmImage {
	var returns *GoogleWorkbenchInstanceGceSetupVmImage
	_jsii_.Get(
		j,
		"vmImageInput",
		&returns,
	)
	return returns
}


func NewGoogleWorkbenchInstanceGceSetupOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleWorkbenchInstanceGceSetupOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleWorkbenchInstanceGceSetupOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleWorkbenchInstance.GoogleWorkbenchInstanceGceSetupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleWorkbenchInstanceGceSetupOutputReference_Override(g GoogleWorkbenchInstanceGceSetupOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleWorkbenchInstance.GoogleWorkbenchInstanceGceSetupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference)SetDisablePublicIp(val interface{}) {
	if err := j.validateSetDisablePublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disablePublicIp",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference)SetEnableIpForwarding(val interface{}) {
	if err := j.validateSetEnableIpForwardingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIpForwarding",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference)SetInternalValue(val *GoogleWorkbenchInstanceGceSetup) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) PutAcceleratorConfigs(value interface{}) {
	if err := g.validatePutAcceleratorConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAcceleratorConfigs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) PutBootDisk(value *GoogleWorkbenchInstanceGceSetupBootDisk) {
	if err := g.validatePutBootDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBootDisk",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) PutConfidentialInstanceConfig(value *GoogleWorkbenchInstanceGceSetupConfidentialInstanceConfig) {
	if err := g.validatePutConfidentialInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConfidentialInstanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) PutContainerImage(value *GoogleWorkbenchInstanceGceSetupContainerImage) {
	if err := g.validatePutContainerImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putContainerImage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) PutDataDisks(value *GoogleWorkbenchInstanceGceSetupDataDisks) {
	if err := g.validatePutDataDisksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataDisks",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) PutNetworkInterfaces(value interface{}) {
	if err := g.validatePutNetworkInterfacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkInterfaces",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) PutServiceAccounts(value interface{}) {
	if err := g.validatePutServiceAccountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceAccounts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) PutShieldedInstanceConfig(value *GoogleWorkbenchInstanceGceSetupShieldedInstanceConfig) {
	if err := g.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) PutVmImage(value *GoogleWorkbenchInstanceGceSetupVmImage) {
	if err := g.validatePutVmImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVmImage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ResetAcceleratorConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetAcceleratorConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ResetBootDisk() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDisk",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ResetConfidentialInstanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetConfidentialInstanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ResetContainerImage() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerImage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ResetDataDisks() {
	_jsii_.InvokeVoid(
		g,
		"resetDataDisks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ResetDisablePublicIp() {
	_jsii_.InvokeVoid(
		g,
		"resetDisablePublicIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ResetEnableIpForwarding() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableIpForwarding",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ResetMachineType() {
	_jsii_.InvokeVoid(
		g,
		"resetMachineType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ResetNetworkInterfaces() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkInterfaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ResetServiceAccounts() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccounts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ResetVmImage() {
	_jsii_.InvokeVoid(
		g,
		"resetVmImage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleWorkbenchInstanceGceSetupOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


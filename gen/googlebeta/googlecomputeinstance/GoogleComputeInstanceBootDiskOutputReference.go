package googlecomputeinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlecomputeinstance/internal"
)

type GoogleComputeInstanceBootDiskOutputReference interface {
	cdktf.ComplexObject
	AutoDelete() interface{}
	SetAutoDelete(val interface{})
	AutoDeleteInput() interface{}
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeviceName() *string
	SetDeviceName(val *string)
	DeviceNameInput() *string
	DiskEncryptionKeyRaw() *string
	SetDiskEncryptionKeyRaw(val *string)
	DiskEncryptionKeyRawInput() *string
	DiskEncryptionKeySha256() *string
	// Experimental.
	Fqn() *string
	InitializeParams() GoogleComputeInstanceBootDiskInitializeParamsOutputReference
	InitializeParamsInput() *GoogleComputeInstanceBootDiskInitializeParams
	InternalValue() *GoogleComputeInstanceBootDisk
	SetInternalValue(val *GoogleComputeInstanceBootDisk)
	KmsKeySelfLink() *string
	SetKmsKeySelfLink(val *string)
	KmsKeySelfLinkInput() *string
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	Source() *string
	SetSource(val *string)
	SourceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutInitializeParams(value *GoogleComputeInstanceBootDiskInitializeParams)
	ResetAutoDelete()
	ResetDeviceName()
	ResetDiskEncryptionKeyRaw()
	ResetInitializeParams()
	ResetKmsKeySelfLink()
	ResetMode()
	ResetSource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeInstanceBootDiskOutputReference
type jsiiProxy_GoogleComputeInstanceBootDiskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) AutoDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) AutoDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) DeviceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) DeviceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) DiskEncryptionKeyRaw() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionKeyRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) DiskEncryptionKeyRawInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionKeyRawInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) DiskEncryptionKeySha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionKeySha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) InitializeParams() GoogleComputeInstanceBootDiskInitializeParamsOutputReference {
	var returns GoogleComputeInstanceBootDiskInitializeParamsOutputReference
	_jsii_.Get(
		j,
		"initializeParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) InitializeParamsInput() *GoogleComputeInstanceBootDiskInitializeParams {
	var returns *GoogleComputeInstanceBootDiskInitializeParams
	_jsii_.Get(
		j,
		"initializeParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) InternalValue() *GoogleComputeInstanceBootDisk {
	var returns *GoogleComputeInstanceBootDisk
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) KmsKeySelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeySelfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) KmsKeySelfLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeySelfLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeInstanceBootDiskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeInstanceBootDiskOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeInstanceBootDiskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeInstanceBootDiskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleComputeInstance.GoogleComputeInstanceBootDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeInstanceBootDiskOutputReference_Override(g GoogleComputeInstanceBootDiskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleComputeInstance.GoogleComputeInstanceBootDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference)SetAutoDelete(val interface{}) {
	if err := j.validateSetAutoDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDelete",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference)SetDeviceName(val *string) {
	if err := j.validateSetDeviceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceName",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference)SetDiskEncryptionKeyRaw(val *string) {
	if err := j.validateSetDiskEncryptionKeyRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncryptionKeyRaw",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference)SetInternalValue(val *GoogleComputeInstanceBootDisk) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference)SetKmsKeySelfLink(val *string) {
	if err := j.validateSetKmsKeySelfLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeySelfLink",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) PutInitializeParams(value *GoogleComputeInstanceBootDiskInitializeParams) {
	if err := g.validatePutInitializeParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInitializeParams",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) ResetAutoDelete() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoDelete",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) ResetDeviceName() {
	_jsii_.InvokeVoid(
		g,
		"resetDeviceName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) ResetDiskEncryptionKeyRaw() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskEncryptionKeyRaw",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) ResetInitializeParams() {
	_jsii_.InvokeVoid(
		g,
		"resetInitializeParams",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) ResetKmsKeySelfLink() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsKeySelfLink",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		g,
		"resetMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		g,
		"resetSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceBootDiskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


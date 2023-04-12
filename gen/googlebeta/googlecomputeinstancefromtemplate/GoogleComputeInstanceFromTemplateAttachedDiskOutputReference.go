package googlecomputeinstancefromtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlecomputeinstancefromtemplate/internal"
)

type GoogleComputeInstanceFromTemplateAttachedDiskOutputReference interface {
	cdktf.ComplexObject
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
	SetDiskEncryptionKeySha256(val *string)
	DiskEncryptionKeySha256Input() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ResetDeviceName()
	ResetDiskEncryptionKeyRaw()
	ResetDiskEncryptionKeySha256()
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

// The jsii proxy struct for GoogleComputeInstanceFromTemplateAttachedDiskOutputReference
type jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) DeviceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) DeviceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) DiskEncryptionKeyRaw() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionKeyRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) DiskEncryptionKeyRawInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionKeyRawInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) DiskEncryptionKeySha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionKeySha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) DiskEncryptionKeySha256Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionKeySha256Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) KmsKeySelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeySelfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) KmsKeySelfLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeySelfLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeInstanceFromTemplateAttachedDiskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleComputeInstanceFromTemplateAttachedDiskOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeInstanceFromTemplateAttachedDiskOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleComputeInstanceFromTemplate.GoogleComputeInstanceFromTemplateAttachedDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleComputeInstanceFromTemplateAttachedDiskOutputReference_Override(g GoogleComputeInstanceFromTemplateAttachedDiskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleComputeInstanceFromTemplate.GoogleComputeInstanceFromTemplateAttachedDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference)SetDeviceName(val *string) {
	if err := j.validateSetDeviceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceName",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference)SetDiskEncryptionKeyRaw(val *string) {
	if err := j.validateSetDiskEncryptionKeyRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncryptionKeyRaw",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference)SetDiskEncryptionKeySha256(val *string) {
	if err := j.validateSetDiskEncryptionKeySha256Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncryptionKeySha256",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference)SetKmsKeySelfLink(val *string) {
	if err := j.validateSetKmsKeySelfLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeySelfLink",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) ResetDeviceName() {
	_jsii_.InvokeVoid(
		g,
		"resetDeviceName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) ResetDiskEncryptionKeyRaw() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskEncryptionKeyRaw",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) ResetDiskEncryptionKeySha256() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskEncryptionKeySha256",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) ResetKmsKeySelfLink() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsKeySelfLink",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		g,
		"resetMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		g,
		"resetSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAttachedDiskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


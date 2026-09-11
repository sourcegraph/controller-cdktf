package backupdrrestoreworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/backupdrrestoreworkload/internal"
)

type BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference interface {
	cdktn.ComplexObject
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
	ExternalIp() *string
	SetExternalIp(val *string)
	ExternalIpInput() *string
	ExternalIpv6() *string
	SetExternalIpv6(val *string)
	ExternalIpv6Input() *string
	ExternalIpv6PrefixLength() *float64
	SetExternalIpv6PrefixLength(val *float64)
	ExternalIpv6PrefixLengthInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkTier() *string
	SetNetworkTier(val *string)
	NetworkTierInput() *string
	PublicPtrDomainName() *string
	SetPublicPtrDomainName(val *string)
	PublicPtrDomainNameInput() *string
	SetPublicPtr() interface{}
	SetSetPublicPtr(val interface{})
	SetPublicPtrInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	ResetExternalIp()
	ResetExternalIpv6()
	ResetExternalIpv6PrefixLength()
	ResetName()
	ResetNetworkTier()
	ResetPublicPtrDomainName()
	ResetSetPublicPtr()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference
type jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ExternalIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ExternalIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ExternalIpv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ExternalIpv6Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ExternalIpv6PrefixLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"externalIpv6PrefixLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ExternalIpv6PrefixLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"externalIpv6PrefixLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) NetworkTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) NetworkTierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) PublicPtrDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicPtrDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) PublicPtrDomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicPtrDomainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) SetPublicPtr() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setPublicPtr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) SetPublicPtrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setPublicPtrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.backupDrRestoreWorkload.BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference_Override(b BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.backupDrRestoreWorkload.BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetExternalIp(val *string) {
	if err := j.validateSetExternalIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalIp",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetExternalIpv6(val *string) {
	if err := j.validateSetExternalIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalIpv6",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetExternalIpv6PrefixLength(val *float64) {
	if err := j.validateSetExternalIpv6PrefixLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalIpv6PrefixLength",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetNetworkTier(val *string) {
	if err := j.validateSetNetworkTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkTier",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetPublicPtrDomainName(val *string) {
	if err := j.validateSetPublicPtrDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicPtrDomainName",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetSetPublicPtr(val interface{}) {
	if err := j.validateSetSetPublicPtrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"setPublicPtr",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ResetExternalIp() {
	_jsii_.InvokeVoid(
		b,
		"resetExternalIp",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ResetExternalIpv6() {
	_jsii_.InvokeVoid(
		b,
		"resetExternalIpv6",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ResetExternalIpv6PrefixLength() {
	_jsii_.InvokeVoid(
		b,
		"resetExternalIpv6PrefixLength",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		b,
		"resetName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ResetNetworkTier() {
	_jsii_.InvokeVoid(
		b,
		"resetNetworkTier",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ResetPublicPtrDomainName() {
	_jsii_.InvokeVoid(
		b,
		"resetPublicPtrDomainName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ResetSetPublicPtr() {
	_jsii_.InvokeVoid(
		b,
		"resetSetPublicPtr",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		b,
		"resetType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


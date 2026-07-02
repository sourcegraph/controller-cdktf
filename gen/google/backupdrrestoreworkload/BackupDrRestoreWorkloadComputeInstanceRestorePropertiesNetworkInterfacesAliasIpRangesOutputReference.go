package backupdrrestoreworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/backupdrrestoreworkload/internal"
)

type BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	IpCidrRange() *string
	SetIpCidrRange(val *string)
	IpCidrRangeInput() *string
	SubnetworkRangeName() *string
	SetSubnetworkRangeName(val *string)
	SubnetworkRangeNameInput() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetIpCidrRange()
	ResetSubnetworkRangeName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference
type jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) IpCidrRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipCidrRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) IpCidrRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipCidrRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) SubnetworkRangeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkRangeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) SubnetworkRangeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkRangeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference {
	_init_.Initialize()

	if err := validateNewBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.backupDrRestoreWorkload.BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference_Override(b BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.backupDrRestoreWorkload.BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) SetIpCidrRange(val *string) {
	if err := j.validateSetIpCidrRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipCidrRange",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) SetSubnetworkRangeName(val *string) {
	if err := j.validateSetSubnetworkRangeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetworkRangeName",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) ResetIpCidrRange() {
	_jsii_.InvokeVoid(
		b,
		"resetIpCidrRange",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) ResetSubnetworkRangeName() {
	_jsii_.InvokeVoid(
		b,
		"resetSubnetworkRangeName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		b,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRangesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

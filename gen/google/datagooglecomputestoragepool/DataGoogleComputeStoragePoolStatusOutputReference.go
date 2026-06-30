package datagooglecomputestoragepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/datagooglecomputestoragepool/internal"
)

type DataGoogleComputeStoragePoolStatusOutputReference interface {
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
	DiskCount() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataGoogleComputeStoragePoolStatus
	SetInternalValue(val *DataGoogleComputeStoragePoolStatus)
	LastResizeTimestamp() *string
	MaxTotalProvisionedDiskCapacityGb() *string
	PoolUsedCapacityBytes() *string
	PoolUsedIops() *string
	PoolUsedThroughput() *string
	PoolUserWrittenBytes() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TotalProvisionedDiskCapacityGb() *string
	TotalProvisionedDiskIops() *string
	TotalProvisionedDiskThroughput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGoogleComputeStoragePoolStatusOutputReference
type jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) DiskCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) InternalValue() *DataGoogleComputeStoragePoolStatus {
	var returns *DataGoogleComputeStoragePoolStatus
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) LastResizeTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastResizeTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) MaxTotalProvisionedDiskCapacityGb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTotalProvisionedDiskCapacityGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) PoolUsedCapacityBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolUsedCapacityBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) PoolUsedIops() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolUsedIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) PoolUsedThroughput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolUsedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) PoolUserWrittenBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolUserWrittenBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) TotalProvisionedDiskCapacityGb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalProvisionedDiskCapacityGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) TotalProvisionedDiskIops() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalProvisionedDiskIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) TotalProvisionedDiskThroughput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalProvisionedDiskThroughput",
		&returns,
	)
	return returns
}

func NewDataGoogleComputeStoragePoolStatusOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleComputeStoragePoolStatusOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleComputeStoragePoolStatusOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleComputeStoragePool.DataGoogleComputeStoragePoolStatusOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleComputeStoragePoolStatusOutputReference_Override(d DataGoogleComputeStoragePoolStatusOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleComputeStoragePool.DataGoogleComputeStoragePoolStatusOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) SetInternalValue(val *DataGoogleComputeStoragePoolStatus) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		d,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeStoragePoolStatusOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

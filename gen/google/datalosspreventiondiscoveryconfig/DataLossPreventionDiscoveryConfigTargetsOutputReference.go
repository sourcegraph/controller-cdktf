package datalosspreventiondiscoveryconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/datalosspreventiondiscoveryconfig/internal"
)

type DataLossPreventionDiscoveryConfigTargetsOutputReference interface {
	cdktf.ComplexObject
	BigQueryTarget() DataLossPreventionDiscoveryConfigTargetsBigQueryTargetOutputReference
	BigQueryTargetInput() *DataLossPreventionDiscoveryConfigTargetsBigQueryTarget
	CloudSqlTarget() DataLossPreventionDiscoveryConfigTargetsCloudSqlTargetOutputReference
	CloudSqlTargetInput() *DataLossPreventionDiscoveryConfigTargetsCloudSqlTarget
	CloudStorageTarget() DataLossPreventionDiscoveryConfigTargetsCloudStorageTargetOutputReference
	CloudStorageTargetInput() *DataLossPreventionDiscoveryConfigTargetsCloudStorageTarget
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
	OtherCloudTarget() DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference
	OtherCloudTargetInput() *DataLossPreventionDiscoveryConfigTargetsOtherCloudTarget
	SecretsTarget() DataLossPreventionDiscoveryConfigTargetsSecretsTargetOutputReference
	SecretsTargetInput() *DataLossPreventionDiscoveryConfigTargetsSecretsTarget
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
	PutBigQueryTarget(value *DataLossPreventionDiscoveryConfigTargetsBigQueryTarget)
	PutCloudSqlTarget(value *DataLossPreventionDiscoveryConfigTargetsCloudSqlTarget)
	PutCloudStorageTarget(value *DataLossPreventionDiscoveryConfigTargetsCloudStorageTarget)
	PutOtherCloudTarget(value *DataLossPreventionDiscoveryConfigTargetsOtherCloudTarget)
	PutSecretsTarget(value *DataLossPreventionDiscoveryConfigTargetsSecretsTarget)
	ResetBigQueryTarget()
	ResetCloudSqlTarget()
	ResetCloudStorageTarget()
	ResetOtherCloudTarget()
	ResetSecretsTarget()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionDiscoveryConfigTargetsOutputReference
type jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) BigQueryTarget() DataLossPreventionDiscoveryConfigTargetsBigQueryTargetOutputReference {
	var returns DataLossPreventionDiscoveryConfigTargetsBigQueryTargetOutputReference
	_jsii_.Get(
		j,
		"bigQueryTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) BigQueryTargetInput() *DataLossPreventionDiscoveryConfigTargetsBigQueryTarget {
	var returns *DataLossPreventionDiscoveryConfigTargetsBigQueryTarget
	_jsii_.Get(
		j,
		"bigQueryTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) CloudSqlTarget() DataLossPreventionDiscoveryConfigTargetsCloudSqlTargetOutputReference {
	var returns DataLossPreventionDiscoveryConfigTargetsCloudSqlTargetOutputReference
	_jsii_.Get(
		j,
		"cloudSqlTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) CloudSqlTargetInput() *DataLossPreventionDiscoveryConfigTargetsCloudSqlTarget {
	var returns *DataLossPreventionDiscoveryConfigTargetsCloudSqlTarget
	_jsii_.Get(
		j,
		"cloudSqlTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) CloudStorageTarget() DataLossPreventionDiscoveryConfigTargetsCloudStorageTargetOutputReference {
	var returns DataLossPreventionDiscoveryConfigTargetsCloudStorageTargetOutputReference
	_jsii_.Get(
		j,
		"cloudStorageTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) CloudStorageTargetInput() *DataLossPreventionDiscoveryConfigTargetsCloudStorageTarget {
	var returns *DataLossPreventionDiscoveryConfigTargetsCloudStorageTarget
	_jsii_.Get(
		j,
		"cloudStorageTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) OtherCloudTarget() DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference {
	var returns DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference
	_jsii_.Get(
		j,
		"otherCloudTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) OtherCloudTargetInput() *DataLossPreventionDiscoveryConfigTargetsOtherCloudTarget {
	var returns *DataLossPreventionDiscoveryConfigTargetsOtherCloudTarget
	_jsii_.Get(
		j,
		"otherCloudTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) SecretsTarget() DataLossPreventionDiscoveryConfigTargetsSecretsTargetOutputReference {
	var returns DataLossPreventionDiscoveryConfigTargetsSecretsTargetOutputReference
	_jsii_.Get(
		j,
		"secretsTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) SecretsTargetInput() *DataLossPreventionDiscoveryConfigTargetsSecretsTarget {
	var returns *DataLossPreventionDiscoveryConfigTargetsSecretsTarget
	_jsii_.Get(
		j,
		"secretsTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDataLossPreventionDiscoveryConfigTargetsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataLossPreventionDiscoveryConfigTargetsOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionDiscoveryConfigTargetsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionDiscoveryConfig.DataLossPreventionDiscoveryConfigTargetsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataLossPreventionDiscoveryConfigTargetsOutputReference_Override(d DataLossPreventionDiscoveryConfigTargetsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionDiscoveryConfig.DataLossPreventionDiscoveryConfigTargetsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) PutBigQueryTarget(value *DataLossPreventionDiscoveryConfigTargetsBigQueryTarget) {
	if err := d.validatePutBigQueryTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBigQueryTarget",
		[]any{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) PutCloudSqlTarget(value *DataLossPreventionDiscoveryConfigTargetsCloudSqlTarget) {
	if err := d.validatePutCloudSqlTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCloudSqlTarget",
		[]any{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) PutCloudStorageTarget(value *DataLossPreventionDiscoveryConfigTargetsCloudStorageTarget) {
	if err := d.validatePutCloudStorageTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCloudStorageTarget",
		[]any{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) PutOtherCloudTarget(value *DataLossPreventionDiscoveryConfigTargetsOtherCloudTarget) {
	if err := d.validatePutOtherCloudTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOtherCloudTarget",
		[]any{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) PutSecretsTarget(value *DataLossPreventionDiscoveryConfigTargetsSecretsTarget) {
	if err := d.validatePutSecretsTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecretsTarget",
		[]any{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) ResetBigQueryTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetBigQueryTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) ResetCloudSqlTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudSqlTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) ResetCloudStorageTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudStorageTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) ResetOtherCloudTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetOtherCloudTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) ResetSecretsTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

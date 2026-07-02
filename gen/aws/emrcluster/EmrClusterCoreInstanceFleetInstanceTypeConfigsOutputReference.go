package emrcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/emrcluster/internal"
)

type EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference interface {
	cdktf.ComplexObject
	BidPrice() *string
	SetBidPrice(val *string)
	BidPriceAsPercentageOfOnDemandPrice() *float64
	SetBidPriceAsPercentageOfOnDemandPrice(val *float64)
	BidPriceAsPercentageOfOnDemandPriceInput() *float64
	BidPriceInput() *string
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
	Configurations() EmrClusterCoreInstanceFleetInstanceTypeConfigsConfigurationsList
	ConfigurationsInput() any
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EbsConfig() EmrClusterCoreInstanceFleetInstanceTypeConfigsEbsConfigList
	EbsConfigInput() any
	// Experimental.
	Fqn() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() any
	SetInternalValue(val any)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeightedCapacity() *float64
	SetWeightedCapacity(val *float64)
	WeightedCapacityInput() *float64
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
	PutConfigurations(value any)
	PutEbsConfig(value any)
	ResetBidPrice()
	ResetBidPriceAsPercentageOfOnDemandPrice()
	ResetConfigurations()
	ResetEbsConfig()
	ResetWeightedCapacity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference
type jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) BidPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bidPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) BidPriceAsPercentageOfOnDemandPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bidPriceAsPercentageOfOnDemandPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) BidPriceAsPercentageOfOnDemandPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bidPriceAsPercentageOfOnDemandPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) BidPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bidPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) Configurations() EmrClusterCoreInstanceFleetInstanceTypeConfigsConfigurationsList {
	var returns EmrClusterCoreInstanceFleetInstanceTypeConfigsConfigurationsList
	_jsii_.Get(
		j,
		"configurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) ConfigurationsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"configurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) EbsConfig() EmrClusterCoreInstanceFleetInstanceTypeConfigsEbsConfigList {
	var returns EmrClusterCoreInstanceFleetInstanceTypeConfigsEbsConfigList
	_jsii_.Get(
		j,
		"ebsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) EbsConfigInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ebsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) WeightedCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightedCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) WeightedCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightedCapacityInput",
		&returns,
	)
	return returns
}

func NewEmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewEmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.emrCluster.EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference_Override(e EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.emrCluster.EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) SetBidPrice(val *string) {
	if err := j.validateSetBidPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bidPrice",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) SetBidPriceAsPercentageOfOnDemandPrice(val *float64) {
	if err := j.validateSetBidPriceAsPercentageOfOnDemandPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bidPriceAsPercentageOfOnDemandPrice",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) SetWeightedCapacity(val *float64) {
	if err := j.validateSetWeightedCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weightedCapacity",
		val,
	)
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) PutConfigurations(value any) {
	if err := e.validatePutConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putConfigurations",
		[]any{value},
	)
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) PutEbsConfig(value any) {
	if err := e.validatePutEbsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEbsConfig",
		[]any{value},
	)
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) ResetBidPrice() {
	_jsii_.InvokeVoid(
		e,
		"resetBidPrice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) ResetBidPriceAsPercentageOfOnDemandPrice() {
	_jsii_.InvokeVoid(
		e,
		"resetBidPriceAsPercentageOfOnDemandPrice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) ResetConfigurations() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigurations",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) ResetEbsConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetEbsConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) ResetWeightedCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetWeightedCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		e,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

package datagooglevmwareengineupgrades

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/datagooglevmwareengineupgrades/internal"
)

type DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference interface {
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
	EndDay() *string
	EndTime() DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsEndTimeList
	// Experimental.
	Fqn() *string
	InternalValue() *DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervals
	SetInternalValue(val *DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervals)
	StartDay() *string
	StartTime() DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsStartTimeList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference
type jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) EndDay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) EndTime() DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsEndTimeList {
	var returns DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsEndTimeList
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) InternalValue() *DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervals {
	var returns *DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervals
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) StartDay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) StartTime() DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsStartTimeList {
	var returns DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsStartTimeList
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.dataGoogleVmwareengineUpgrades.DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference_Override(d DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.dataGoogleVmwareengineUpgrades.DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference)SetInternalValue(val *DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervals) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleVmwareengineUpgradesUpgradesScheduleConstraintsDisallowedIntervalsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


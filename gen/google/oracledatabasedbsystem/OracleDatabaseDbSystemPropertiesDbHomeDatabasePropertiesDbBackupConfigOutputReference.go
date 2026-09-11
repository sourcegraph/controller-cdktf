package oracledatabasedbsystem

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/oracledatabasedbsystem/internal"
)

type OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference interface {
	cdktn.ComplexObject
	AutoBackupEnabled() interface{}
	SetAutoBackupEnabled(val interface{})
	AutoBackupEnabledInput() interface{}
	AutoFullBackupDay() *string
	SetAutoFullBackupDay(val *string)
	AutoFullBackupDayInput() *string
	AutoFullBackupWindow() *string
	SetAutoFullBackupWindow(val *string)
	AutoFullBackupWindowInput() *string
	AutoIncrementalBackupWindow() *string
	SetAutoIncrementalBackupWindow(val *string)
	AutoIncrementalBackupWindowInput() *string
	BackupDeletionPolicy() *string
	SetBackupDeletionPolicy(val *string)
	BackupDeletionPolicyInput() *string
	BackupDestinationDetails() OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList
	BackupDestinationDetailsInput() interface{}
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
	// Experimental.
	Fqn() *string
	InternalValue() *OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfig
	SetInternalValue(val *OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfig)
	RetentionPeriodDays() *float64
	SetRetentionPeriodDays(val *float64)
	RetentionPeriodDaysInput() *float64
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
	PutBackupDestinationDetails(value interface{})
	ResetAutoBackupEnabled()
	ResetAutoFullBackupDay()
	ResetAutoFullBackupWindow()
	ResetAutoIncrementalBackupWindow()
	ResetBackupDeletionPolicy()
	ResetBackupDestinationDetails()
	ResetRetentionPeriodDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference
type jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) AutoBackupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) AutoBackupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoBackupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) AutoFullBackupDay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoFullBackupDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) AutoFullBackupDayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoFullBackupDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) AutoFullBackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoFullBackupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) AutoFullBackupWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoFullBackupWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) AutoIncrementalBackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoIncrementalBackupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) AutoIncrementalBackupWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoIncrementalBackupWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) BackupDeletionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupDeletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) BackupDeletionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupDeletionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) BackupDestinationDetails() OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList {
	var returns OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetailsList
	_jsii_.Get(
		j,
		"backupDestinationDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) BackupDestinationDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupDestinationDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) InternalValue() *OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfig {
	var returns *OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) RetentionPeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) RetentionPeriodDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriodDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseDbSystem.OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference_Override(o OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseDbSystem.OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference)SetAutoBackupEnabled(val interface{}) {
	if err := j.validateSetAutoBackupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoBackupEnabled",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference)SetAutoFullBackupDay(val *string) {
	if err := j.validateSetAutoFullBackupDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoFullBackupDay",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference)SetAutoFullBackupWindow(val *string) {
	if err := j.validateSetAutoFullBackupWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoFullBackupWindow",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference)SetAutoIncrementalBackupWindow(val *string) {
	if err := j.validateSetAutoIncrementalBackupWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoIncrementalBackupWindow",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference)SetBackupDeletionPolicy(val *string) {
	if err := j.validateSetBackupDeletionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupDeletionPolicy",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference)SetInternalValue(val *OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference)SetRetentionPeriodDays(val *float64) {
	if err := j.validateSetRetentionPeriodDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionPeriodDays",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) PutBackupDestinationDetails(value interface{}) {
	if err := o.validatePutBackupDestinationDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putBackupDestinationDetails",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) ResetAutoBackupEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoBackupEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) ResetAutoFullBackupDay() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoFullBackupDay",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) ResetAutoFullBackupWindow() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoFullBackupWindow",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) ResetAutoIncrementalBackupWindow() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoIncrementalBackupWindow",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) ResetBackupDeletionPolicy() {
	_jsii_.InvokeVoid(
		o,
		"resetBackupDeletionPolicy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) ResetBackupDestinationDetails() {
	_jsii_.InvokeVoid(
		o,
		"resetBackupDestinationDetails",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) ResetRetentionPeriodDays() {
	_jsii_.InvokeVoid(
		o,
		"resetRetentionPeriodDays",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package oracledatabasedbsystem

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/oracledatabasedbsystem/internal"
)

type OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference interface {
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
	DatabaseManagementConfig() OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDatabaseManagementConfigOutputReference
	DatabaseManagementConfigInput() *OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDatabaseManagementConfig
	DbBackupConfig() OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference
	DbBackupConfigInput() *OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfig
	DbVersion() *string
	SetDbVersion(val *string)
	DbVersionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *OracleDatabaseDbSystemPropertiesDbHomeDatabaseProperties
	SetInternalValue(val *OracleDatabaseDbSystemPropertiesDbHomeDatabaseProperties)
	State() *string
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
	PutDatabaseManagementConfig(value *OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDatabaseManagementConfig)
	PutDbBackupConfig(value *OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfig)
	ResetDatabaseManagementConfig()
	ResetDbBackupConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference
type jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) DatabaseManagementConfig() OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDatabaseManagementConfigOutputReference {
	var returns OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDatabaseManagementConfigOutputReference
	_jsii_.Get(
		j,
		"databaseManagementConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) DatabaseManagementConfigInput() *OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDatabaseManagementConfig {
	var returns *OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDatabaseManagementConfig
	_jsii_.Get(
		j,
		"databaseManagementConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) DbBackupConfig() OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference {
	var returns OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigOutputReference
	_jsii_.Get(
		j,
		"dbBackupConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) DbBackupConfigInput() *OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfig {
	var returns *OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfig
	_jsii_.Get(
		j,
		"dbBackupConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) DbVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) DbVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) InternalValue() *OracleDatabaseDbSystemPropertiesDbHomeDatabaseProperties {
	var returns *OracleDatabaseDbSystemPropertiesDbHomeDatabaseProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewOracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.oracleDatabaseDbSystem.OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference_Override(o OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.oracleDatabaseDbSystem.OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference",
		[]any{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) SetDbVersion(val *string) {
	if err := j.validateSetDbVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbVersion",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) SetInternalValue(val *OracleDatabaseDbSystemPropertiesDbHomeDatabaseProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) PutDatabaseManagementConfig(value *OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDatabaseManagementConfig) {
	if err := o.validatePutDatabaseManagementConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDatabaseManagementConfig",
		[]any{value},
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) PutDbBackupConfig(value *OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfig) {
	if err := o.validatePutDbBackupConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDbBackupConfig",
		[]any{value},
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) ResetDatabaseManagementConfig() {
	_jsii_.InvokeVoid(
		o,
		"resetDatabaseManagementConfig",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) ResetDbBackupConfig() {
	_jsii_.InvokeVoid(
		o,
		"resetDbBackupConfig",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		o,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

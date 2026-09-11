package oracledatabasedbsystem

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/oracledatabasedbsystem/internal"
)

type OracleDatabaseDbSystemPropertiesDbHomeOutputReference interface {
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
	Database() OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference
	DatabaseInput() *OracleDatabaseDbSystemPropertiesDbHomeDatabase
	DbVersion() *string
	SetDbVersion(val *string)
	DbVersionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *OracleDatabaseDbSystemPropertiesDbHome
	SetInternalValue(val *OracleDatabaseDbSystemPropertiesDbHome)
	IsUnifiedAuditingEnabled() interface{}
	SetIsUnifiedAuditingEnabled(val interface{})
	IsUnifiedAuditingEnabledInput() interface{}
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
	PutDatabase(value *OracleDatabaseDbSystemPropertiesDbHomeDatabase)
	ResetDisplayName()
	ResetIsUnifiedAuditingEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OracleDatabaseDbSystemPropertiesDbHomeOutputReference
type jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) Database() OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference {
	var returns OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) DatabaseInput() *OracleDatabaseDbSystemPropertiesDbHomeDatabase {
	var returns *OracleDatabaseDbSystemPropertiesDbHomeDatabase
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) DbVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) DbVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) InternalValue() *OracleDatabaseDbSystemPropertiesDbHome {
	var returns *OracleDatabaseDbSystemPropertiesDbHome
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) IsUnifiedAuditingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isUnifiedAuditingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) IsUnifiedAuditingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isUnifiedAuditingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOracleDatabaseDbSystemPropertiesDbHomeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) OracleDatabaseDbSystemPropertiesDbHomeOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseDbSystemPropertiesDbHomeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseDbSystem.OracleDatabaseDbSystemPropertiesDbHomeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseDbSystemPropertiesDbHomeOutputReference_Override(o OracleDatabaseDbSystemPropertiesDbHomeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseDbSystem.OracleDatabaseDbSystemPropertiesDbHomeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference)SetDbVersion(val *string) {
	if err := j.validateSetDbVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbVersion",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference)SetInternalValue(val *OracleDatabaseDbSystemPropertiesDbHome) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference)SetIsUnifiedAuditingEnabled(val interface{}) {
	if err := j.validateSetIsUnifiedAuditingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isUnifiedAuditingEnabled",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) PutDatabase(value *OracleDatabaseDbSystemPropertiesDbHomeDatabase) {
	if err := o.validatePutDatabaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDatabase",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		o,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) ResetIsUnifiedAuditingEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetIsUnifiedAuditingEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


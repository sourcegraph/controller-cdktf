package oracledatabasedbsystem

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/oracledatabasedbsystem/internal"
)

type OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference interface {
	cdktf.ComplexObject
	AdminPassword() *string
	SetAdminPassword(val *string)
	AdminPasswordInput() *string
	CharacterSet() *string
	SetCharacterSet(val *string)
	CharacterSetInput() *string
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
	CreateTime() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DatabaseId() *string
	SetDatabaseId(val *string)
	DatabaseIdInput() *string
	DbHomeName() *string
	SetDbHomeName(val *string)
	DbHomeNameInput() *string
	DbName() *string
	SetDbName(val *string)
	DbNameInput() *string
	DbUniqueName() *string
	SetDbUniqueName(val *string)
	DbUniqueNameInput() *string
	// Experimental.
	Fqn() *string
	GcpOracleZone() *string
	SetGcpOracleZone(val *string)
	GcpOracleZoneInput() *string
	InternalValue() *OracleDatabaseDbSystemPropertiesDbHomeDatabase
	SetInternalValue(val *OracleDatabaseDbSystemPropertiesDbHomeDatabase)
	Name() *string
	NcharacterSet() *string
	SetNcharacterSet(val *string)
	NcharacterSetInput() *string
	OciUrl() *string
	OpsInsightsStatus() *string
	PluggableDatabaseId() *string
	SetPluggableDatabaseId(val *string)
	PluggableDatabaseIdInput() *string
	PluggableDatabaseName() *string
	SetPluggableDatabaseName(val *string)
	PluggableDatabaseNameInput() *string
	Properties() OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference
	PropertiesInput() *OracleDatabaseDbSystemPropertiesDbHomeDatabaseProperties
	TdeWalletPassword() *string
	SetTdeWalletPassword(val *string)
	TdeWalletPasswordInput() *string
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
	PutProperties(value *OracleDatabaseDbSystemPropertiesDbHomeDatabaseProperties)
	ResetCharacterSet()
	ResetDbHomeName()
	ResetDbName()
	ResetDbUniqueName()
	ResetGcpOracleZone()
	ResetNcharacterSet()
	ResetPluggableDatabaseId()
	ResetPluggableDatabaseName()
	ResetProperties()
	ResetTdeWalletPassword()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference
type jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) CharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) CharacterSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) DatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) DatabaseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) DbHomeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbHomeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) DbHomeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbHomeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) DbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) DbNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) DbUniqueName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbUniqueName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) DbUniqueNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbUniqueNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GcpOracleZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpOracleZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GcpOracleZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpOracleZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) InternalValue() *OracleDatabaseDbSystemPropertiesDbHomeDatabase {
	var returns *OracleDatabaseDbSystemPropertiesDbHomeDatabase
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) NcharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ncharacterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) NcharacterSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ncharacterSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) OpsInsightsStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsInsightsStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) PluggableDatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluggableDatabaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) PluggableDatabaseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluggableDatabaseIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) PluggableDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluggableDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) PluggableDatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluggableDatabaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) Properties() OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference {
	var returns OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) PropertiesInput() *OracleDatabaseDbSystemPropertiesDbHomeDatabaseProperties {
	var returns *OracleDatabaseDbSystemPropertiesDbHomeDatabaseProperties
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) TdeWalletPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tdeWalletPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) TdeWalletPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tdeWalletPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.oracleDatabaseDbSystem.OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference_Override(o OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.oracleDatabaseDbSystem.OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetAdminPassword(val *string) {
	if err := j.validateSetAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetCharacterSet(val *string) {
	if err := j.validateSetCharacterSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"characterSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetDatabaseId(val *string) {
	if err := j.validateSetDatabaseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseId",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetDbHomeName(val *string) {
	if err := j.validateSetDbHomeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbHomeName",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetDbName(val *string) {
	if err := j.validateSetDbNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbName",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetDbUniqueName(val *string) {
	if err := j.validateSetDbUniqueNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbUniqueName",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetGcpOracleZone(val *string) {
	if err := j.validateSetGcpOracleZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcpOracleZone",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetInternalValue(val *OracleDatabaseDbSystemPropertiesDbHomeDatabase) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetNcharacterSet(val *string) {
	if err := j.validateSetNcharacterSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ncharacterSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetPluggableDatabaseId(val *string) {
	if err := j.validateSetPluggableDatabaseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluggableDatabaseId",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetPluggableDatabaseName(val *string) {
	if err := j.validateSetPluggableDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluggableDatabaseName",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetTdeWalletPassword(val *string) {
	if err := j.validateSetTdeWalletPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tdeWalletPassword",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) PutProperties(value *OracleDatabaseDbSystemPropertiesDbHomeDatabaseProperties) {
	if err := o.validatePutPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putProperties",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ResetCharacterSet() {
	_jsii_.InvokeVoid(
		o,
		"resetCharacterSet",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ResetDbHomeName() {
	_jsii_.InvokeVoid(
		o,
		"resetDbHomeName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ResetDbName() {
	_jsii_.InvokeVoid(
		o,
		"resetDbName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ResetDbUniqueName() {
	_jsii_.InvokeVoid(
		o,
		"resetDbUniqueName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ResetGcpOracleZone() {
	_jsii_.InvokeVoid(
		o,
		"resetGcpOracleZone",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ResetNcharacterSet() {
	_jsii_.InvokeVoid(
		o,
		"resetNcharacterSet",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ResetPluggableDatabaseId() {
	_jsii_.InvokeVoid(
		o,
		"resetPluggableDatabaseId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ResetPluggableDatabaseName() {
	_jsii_.InvokeVoid(
		o,
		"resetPluggableDatabaseName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		o,
		"resetProperties",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ResetTdeWalletPassword() {
	_jsii_.InvokeVoid(
		o,
		"resetTdeWalletPassword",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


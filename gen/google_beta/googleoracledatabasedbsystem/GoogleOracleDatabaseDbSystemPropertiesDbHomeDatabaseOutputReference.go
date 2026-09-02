package googleoracledatabasedbsystem

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleoracledatabasedbsystem/internal"
)

type GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() *GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabase
	SetInternalValue(val *GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabase)
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
	Properties() GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference
	PropertiesInput() *GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseProperties
	TdeWalletPassword() *string
	SetTdeWalletPassword(val *string)
	TdeWalletPasswordInput() *string
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
	PutProperties(value *GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseProperties)
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
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference
type jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) CharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) CharacterSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) DatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) DatabaseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) DbHomeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbHomeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) DbHomeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbHomeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) DbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) DbNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) DbUniqueName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbUniqueName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) DbUniqueNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbUniqueNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GcpOracleZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpOracleZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GcpOracleZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpOracleZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) InternalValue() *GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabase {
	var returns *GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabase
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) NcharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ncharacterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) NcharacterSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ncharacterSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) OpsInsightsStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsInsightsStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) PluggableDatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluggableDatabaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) PluggableDatabaseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluggableDatabaseIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) PluggableDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluggableDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) PluggableDatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluggableDatabaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) Properties() GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference {
	var returns GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesOutputReference
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) PropertiesInput() *GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseProperties {
	var returns *GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseProperties
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) TdeWalletPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tdeWalletPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) TdeWalletPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tdeWalletPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseDbSystem.GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference_Override(g GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseDbSystem.GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetAdminPassword(val *string) {
	if err := j.validateSetAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetCharacterSet(val *string) {
	if err := j.validateSetCharacterSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"characterSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetDatabaseId(val *string) {
	if err := j.validateSetDatabaseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseId",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetDbHomeName(val *string) {
	if err := j.validateSetDbHomeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbHomeName",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetDbName(val *string) {
	if err := j.validateSetDbNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbName",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetDbUniqueName(val *string) {
	if err := j.validateSetDbUniqueNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbUniqueName",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetGcpOracleZone(val *string) {
	if err := j.validateSetGcpOracleZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcpOracleZone",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetInternalValue(val *GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabase) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetNcharacterSet(val *string) {
	if err := j.validateSetNcharacterSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ncharacterSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetPluggableDatabaseId(val *string) {
	if err := j.validateSetPluggableDatabaseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluggableDatabaseId",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetPluggableDatabaseName(val *string) {
	if err := j.validateSetPluggableDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluggableDatabaseName",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetTdeWalletPassword(val *string) {
	if err := j.validateSetTdeWalletPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tdeWalletPassword",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) PutProperties(value *GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseProperties) {
	if err := g.validatePutPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ResetCharacterSet() {
	_jsii_.InvokeVoid(
		g,
		"resetCharacterSet",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ResetDbHomeName() {
	_jsii_.InvokeVoid(
		g,
		"resetDbHomeName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ResetDbName() {
	_jsii_.InvokeVoid(
		g,
		"resetDbName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ResetDbUniqueName() {
	_jsii_.InvokeVoid(
		g,
		"resetDbUniqueName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ResetGcpOracleZone() {
	_jsii_.InvokeVoid(
		g,
		"resetGcpOracleZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ResetNcharacterSet() {
	_jsii_.InvokeVoid(
		g,
		"resetNcharacterSet",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ResetPluggableDatabaseId() {
	_jsii_.InvokeVoid(
		g,
		"resetPluggableDatabaseId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ResetPluggableDatabaseName() {
	_jsii_.InvokeVoid(
		g,
		"resetPluggableDatabaseName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ResetTdeWalletPassword() {
	_jsii_.InvokeVoid(
		g,
		"resetTdeWalletPassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


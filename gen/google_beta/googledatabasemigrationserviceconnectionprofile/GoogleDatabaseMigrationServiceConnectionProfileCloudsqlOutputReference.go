package googledatabasemigrationserviceconnectionprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatabasemigrationserviceconnectionprofile/internal"
)

type GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference interface {
	cdktf.ComplexObject
	CloudSqlId() *string
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
	InternalValue() *GoogleDatabaseMigrationServiceConnectionProfileCloudsql
	SetInternalValue(val *GoogleDatabaseMigrationServiceConnectionProfileCloudsql)
	PrivateIp() *string
	PublicIp() *string
	Settings() GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference
	SettingsInput() *GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettings
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
	PutSettings(value *GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettings)
	ResetSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference
type jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) CloudSqlId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudSqlId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) InternalValue() *GoogleDatabaseMigrationServiceConnectionProfileCloudsql {
	var returns *GoogleDatabaseMigrationServiceConnectionProfileCloudsql
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) PublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) Settings() GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference {
	var returns GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) SettingsInput() *GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettings {
	var returns *GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettings
	_jsii_.Get(
		j,
		"settingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference{}

	_jsii_.Create(
		"google-beta.googleDatabaseMigrationServiceConnectionProfile.GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference_Override(g GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDatabaseMigrationServiceConnectionProfile.GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference)SetInternalValue(val *GoogleDatabaseMigrationServiceConnectionProfileCloudsql) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) PutSettings(value *GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettings) {
	if err := g.validatePutSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) ResetSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


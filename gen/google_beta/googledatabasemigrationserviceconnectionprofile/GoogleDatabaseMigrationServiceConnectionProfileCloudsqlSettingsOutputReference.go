package googledatabasemigrationserviceconnectionprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatabasemigrationserviceconnectionprofile/internal"
)

type GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference interface {
	cdktf.ComplexObject
	ActivationPolicy() *string
	SetActivationPolicy(val *string)
	ActivationPolicyInput() *string
	AutoStorageIncrease() interface{}
	SetAutoStorageIncrease(val interface{})
	AutoStorageIncreaseInput() interface{}
	CmekKeyName() *string
	SetCmekKeyName(val *string)
	CmekKeyNameInput() *string
	Collation() *string
	SetCollation(val *string)
	CollationInput() *string
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
	DatabaseFlags() *map[string]*string
	SetDatabaseFlags(val *map[string]*string)
	DatabaseFlagsInput() *map[string]*string
	DatabaseVersion() *string
	SetDatabaseVersion(val *string)
	DatabaseVersionInput() *string
	DataDiskSizeGb() *string
	SetDataDiskSizeGb(val *string)
	DataDiskSizeGbInput() *string
	DataDiskType() *string
	SetDataDiskType(val *string)
	DataDiskTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettings
	SetInternalValue(val *GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettings)
	IpConfig() GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsIpConfigOutputReference
	IpConfigInput() *GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsIpConfig
	RootPassword() *string
	SetRootPassword(val *string)
	RootPasswordInput() *string
	RootPasswordSet() cdktf.IResolvable
	SourceId() *string
	SetSourceId(val *string)
	SourceIdInput() *string
	StorageAutoResizeLimit() *string
	SetStorageAutoResizeLimit(val *string)
	StorageAutoResizeLimitInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Tier() *string
	SetTier(val *string)
	TierInput() *string
	UserLabels() *map[string]*string
	SetUserLabels(val *map[string]*string)
	UserLabelsInput() *map[string]*string
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
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
	PutIpConfig(value *GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsIpConfig)
	ResetActivationPolicy()
	ResetAutoStorageIncrease()
	ResetCmekKeyName()
	ResetCollation()
	ResetDatabaseFlags()
	ResetDatabaseVersion()
	ResetDataDiskSizeGb()
	ResetDataDiskType()
	ResetIpConfig()
	ResetRootPassword()
	ResetStorageAutoResizeLimit()
	ResetTier()
	ResetUserLabels()
	ResetZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference
type jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ActivationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ActivationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) AutoStorageIncrease() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoStorageIncrease",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) AutoStorageIncreaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoStorageIncreaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) CmekKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cmekKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) CmekKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cmekKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) Collation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) CollationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) DatabaseFlags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"databaseFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) DatabaseFlagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"databaseFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) DatabaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) DatabaseVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) DataDiskSizeGb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDiskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) DataDiskSizeGbInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDiskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) DataDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) DataDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) InternalValue() *GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettings {
	var returns *GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) IpConfig() GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsIpConfigOutputReference {
	var returns GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsIpConfigOutputReference
	_jsii_.Get(
		j,
		"ipConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) IpConfigInput() *GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsIpConfig {
	var returns *GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsIpConfig
	_jsii_.Get(
		j,
		"ipConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) RootPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) RootPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) RootPasswordSet() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"rootPasswordSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) SourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) SourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) StorageAutoResizeLimit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAutoResizeLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) StorageAutoResizeLimitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAutoResizeLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) TierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) UserLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) UserLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


func NewGoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference{}

	_jsii_.Create(
		"google-beta.googleDatabaseMigrationServiceConnectionProfile.GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference_Override(g GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDatabaseMigrationServiceConnectionProfile.GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetActivationPolicy(val *string) {
	if err := j.validateSetActivationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activationPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetAutoStorageIncrease(val interface{}) {
	if err := j.validateSetAutoStorageIncreaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoStorageIncrease",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetCmekKeyName(val *string) {
	if err := j.validateSetCmekKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cmekKeyName",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetCollation(val *string) {
	if err := j.validateSetCollationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collation",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetDatabaseFlags(val *map[string]*string) {
	if err := j.validateSetDatabaseFlagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseFlags",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetDatabaseVersion(val *string) {
	if err := j.validateSetDatabaseVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetDataDiskSizeGb(val *string) {
	if err := j.validateSetDataDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetDataDiskType(val *string) {
	if err := j.validateSetDataDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataDiskType",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetInternalValue(val *GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetRootPassword(val *string) {
	if err := j.validateSetRootPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootPassword",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetSourceId(val *string) {
	if err := j.validateSetSourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceId",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetStorageAutoResizeLimit(val *string) {
	if err := j.validateSetStorageAutoResizeLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAutoResizeLimit",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetTier(val *string) {
	if err := j.validateSetTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tier",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetUserLabels(val *map[string]*string) {
	if err := j.validateSetUserLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userLabels",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) PutIpConfig(value *GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsIpConfig) {
	if err := g.validatePutIpConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIpConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetActivationPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetActivationPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetAutoStorageIncrease() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoStorageIncrease",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetCmekKeyName() {
	_jsii_.InvokeVoid(
		g,
		"resetCmekKeyName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetCollation() {
	_jsii_.InvokeVoid(
		g,
		"resetCollation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetDatabaseFlags() {
	_jsii_.InvokeVoid(
		g,
		"resetDatabaseFlags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetDatabaseVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetDatabaseVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetDataDiskSizeGb() {
	_jsii_.InvokeVoid(
		g,
		"resetDataDiskSizeGb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetDataDiskType() {
	_jsii_.InvokeVoid(
		g,
		"resetDataDiskType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetIpConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetIpConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetRootPassword() {
	_jsii_.InvokeVoid(
		g,
		"resetRootPassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetStorageAutoResizeLimit() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageAutoResizeLimit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetTier() {
	_jsii_.InvokeVoid(
		g,
		"resetTier",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetUserLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetUserLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetZone() {
	_jsii_.InvokeVoid(
		g,
		"resetZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


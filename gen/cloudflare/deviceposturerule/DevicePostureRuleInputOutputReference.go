package deviceposturerule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/deviceposturerule/internal"
)

type DevicePostureRuleInputOutputReference interface {
	cdktf.ComplexObject
	CheckDisks() *[]*string
	SetCheckDisks(val *[]*string)
	CheckDisksInput() *[]*string
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
	ComplianceStatus() *string
	SetComplianceStatus(val *string)
	ComplianceStatusInput() *string
	ConnectionId() *string
	SetConnectionId(val *string)
	ConnectionIdInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Exists() interface{}
	SetExists(val interface{})
	ExistsInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	Os() *string
	SetOs(val *string)
	OsDistroName() *string
	SetOsDistroName(val *string)
	OsDistroNameInput() *string
	OsDistroRevision() *string
	SetOsDistroRevision(val *string)
	OsDistroRevisionInput() *string
	OsInput() *string
	Overall() *string
	SetOverall(val *string)
	OverallInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	RequireAll() interface{}
	SetRequireAll(val interface{})
	RequireAllInput() interface{}
	Running() interface{}
	SetRunning(val interface{})
	RunningInput() interface{}
	SensorConfig() *string
	SetSensorConfig(val *string)
	SensorConfigInput() *string
	Sha256() *string
	SetSha256(val *string)
	Sha256Input() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Thumbprint() *string
	SetThumbprint(val *string)
	ThumbprintInput() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	VersionOperator() *string
	SetVersionOperator(val *string)
	VersionOperatorInput() *string
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
	ResetCheckDisks()
	ResetComplianceStatus()
	ResetConnectionId()
	ResetDomain()
	ResetEnabled()
	ResetExists()
	ResetId()
	ResetOperator()
	ResetOs()
	ResetOsDistroName()
	ResetOsDistroRevision()
	ResetOverall()
	ResetPath()
	ResetRequireAll()
	ResetRunning()
	ResetSensorConfig()
	ResetSha256()
	ResetThumbprint()
	ResetVersion()
	ResetVersionOperator()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DevicePostureRuleInputOutputReference
type jsiiProxy_DevicePostureRuleInputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) CheckDisks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"checkDisks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) CheckDisksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"checkDisksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) ComplianceStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) ComplianceStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) ConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) ConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) Exists() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) ExistsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"existsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) Os() *string {
	var returns *string
	_jsii_.Get(
		j,
		"os",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) OsDistroName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDistroName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) OsDistroNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDistroNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) OsDistroRevision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDistroRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) OsDistroRevisionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDistroRevisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) OsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) Overall() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) OverallInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) RequireAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) RequireAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) Running() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"running",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) RunningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) SensorConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sensorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) SensorConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sensorConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) Sha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) Sha256Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sha256Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) Thumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) ThumbprintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) VersionOperator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference) VersionOperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionOperatorInput",
		&returns,
	)
	return returns
}


func NewDevicePostureRuleInputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DevicePostureRuleInputOutputReference {
	_init_.Initialize()

	if err := validateNewDevicePostureRuleInputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DevicePostureRuleInputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.devicePostureRule.DevicePostureRuleInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDevicePostureRuleInputOutputReference_Override(d DevicePostureRuleInputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.devicePostureRule.DevicePostureRuleInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetCheckDisks(val *[]*string) {
	if err := j.validateSetCheckDisksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkDisks",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetComplianceStatus(val *string) {
	if err := j.validateSetComplianceStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complianceStatus",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetConnectionId(val *string) {
	if err := j.validateSetConnectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionId",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetExists(val interface{}) {
	if err := j.validateSetExistsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exists",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetOs(val *string) {
	if err := j.validateSetOsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"os",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetOsDistroName(val *string) {
	if err := j.validateSetOsDistroNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osDistroName",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetOsDistroRevision(val *string) {
	if err := j.validateSetOsDistroRevisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osDistroRevision",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetOverall(val *string) {
	if err := j.validateSetOverallParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overall",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetRequireAll(val interface{}) {
	if err := j.validateSetRequireAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireAll",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetRunning(val interface{}) {
	if err := j.validateSetRunningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"running",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetSensorConfig(val *string) {
	if err := j.validateSetSensorConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sensorConfig",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetSha256(val *string) {
	if err := j.validateSetSha256Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sha256",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetThumbprint(val *string) {
	if err := j.validateSetThumbprintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thumbprint",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_DevicePostureRuleInputOutputReference)SetVersionOperator(val *string) {
	if err := j.validateSetVersionOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionOperator",
		val,
	)
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ResetCheckDisks() {
	_jsii_.InvokeVoid(
		d,
		"resetCheckDisks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ResetComplianceStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetComplianceStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ResetConnectionId() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectionId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ResetDomain() {
	_jsii_.InvokeVoid(
		d,
		"resetDomain",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ResetExists() {
	_jsii_.InvokeVoid(
		d,
		"resetExists",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ResetOperator() {
	_jsii_.InvokeVoid(
		d,
		"resetOperator",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ResetOs() {
	_jsii_.InvokeVoid(
		d,
		"resetOs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ResetOsDistroName() {
	_jsii_.InvokeVoid(
		d,
		"resetOsDistroName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ResetOsDistroRevision() {
	_jsii_.InvokeVoid(
		d,
		"resetOsDistroRevision",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ResetOverall() {
	_jsii_.InvokeVoid(
		d,
		"resetOverall",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		d,
		"resetPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ResetRequireAll() {
	_jsii_.InvokeVoid(
		d,
		"resetRequireAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ResetRunning() {
	_jsii_.InvokeVoid(
		d,
		"resetRunning",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ResetSensorConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetSensorConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ResetSha256() {
	_jsii_.InvokeVoid(
		d,
		"resetSha256",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ResetThumbprint() {
	_jsii_.InvokeVoid(
		d,
		"resetThumbprint",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ResetVersionOperator() {
	_jsii_.InvokeVoid(
		d,
		"resetVersionOperator",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevicePostureRuleInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


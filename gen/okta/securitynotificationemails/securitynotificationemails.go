package securitynotificationemails

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/securitynotificationemails/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/security_notification_emails okta_security_notification_emails}.
type SecurityNotificationEmails interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ReportSuspiciousActivityEnabled() interface{}
	SetReportSuspiciousActivityEnabled(val interface{})
	ReportSuspiciousActivityEnabledInput() interface{}
	SendEmailForFactorEnrollmentEnabled() interface{}
	SetSendEmailForFactorEnrollmentEnabled(val interface{})
	SendEmailForFactorEnrollmentEnabledInput() interface{}
	SendEmailForFactorResetEnabled() interface{}
	SetSendEmailForFactorResetEnabled(val interface{})
	SendEmailForFactorResetEnabledInput() interface{}
	SendEmailForNewDeviceEnabled() interface{}
	SetSendEmailForNewDeviceEnabled(val interface{})
	SendEmailForNewDeviceEnabledInput() interface{}
	SendEmailForPasswordChangedEnabled() interface{}
	SetSendEmailForPasswordChangedEnabled(val interface{})
	SendEmailForPasswordChangedEnabledInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReportSuspiciousActivityEnabled()
	ResetSendEmailForFactorEnrollmentEnabled()
	ResetSendEmailForFactorResetEnabled()
	ResetSendEmailForNewDeviceEnabled()
	ResetSendEmailForPasswordChangedEnabled()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SecurityNotificationEmails
type jsiiProxy_SecurityNotificationEmails struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SecurityNotificationEmails) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) ReportSuspiciousActivityEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reportSuspiciousActivityEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) ReportSuspiciousActivityEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reportSuspiciousActivityEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) SendEmailForFactorEnrollmentEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendEmailForFactorEnrollmentEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) SendEmailForFactorEnrollmentEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendEmailForFactorEnrollmentEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) SendEmailForFactorResetEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendEmailForFactorResetEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) SendEmailForFactorResetEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendEmailForFactorResetEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) SendEmailForNewDeviceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendEmailForNewDeviceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) SendEmailForNewDeviceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendEmailForNewDeviceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) SendEmailForPasswordChangedEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendEmailForPasswordChangedEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) SendEmailForPasswordChangedEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendEmailForPasswordChangedEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityNotificationEmails) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/security_notification_emails okta_security_notification_emails} Resource.
func NewSecurityNotificationEmails(scope constructs.Construct, id *string, config *SecurityNotificationEmailsConfig) SecurityNotificationEmails {
	_init_.Initialize()

	if err := validateNewSecurityNotificationEmailsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityNotificationEmails{}

	_jsii_.Create(
		"@cdktf/provider-okta.securityNotificationEmails.SecurityNotificationEmails",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/security_notification_emails okta_security_notification_emails} Resource.
func NewSecurityNotificationEmails_Override(s SecurityNotificationEmails, scope constructs.Construct, id *string, config *SecurityNotificationEmailsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.securityNotificationEmails.SecurityNotificationEmails",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecurityNotificationEmails)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SecurityNotificationEmails)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecurityNotificationEmails)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecurityNotificationEmails)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SecurityNotificationEmails)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SecurityNotificationEmails)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecurityNotificationEmails)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SecurityNotificationEmails)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SecurityNotificationEmails)SetReportSuspiciousActivityEnabled(val interface{}) {
	if err := j.validateSetReportSuspiciousActivityEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportSuspiciousActivityEnabled",
		val,
	)
}

func (j *jsiiProxy_SecurityNotificationEmails)SetSendEmailForFactorEnrollmentEnabled(val interface{}) {
	if err := j.validateSetSendEmailForFactorEnrollmentEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendEmailForFactorEnrollmentEnabled",
		val,
	)
}

func (j *jsiiProxy_SecurityNotificationEmails)SetSendEmailForFactorResetEnabled(val interface{}) {
	if err := j.validateSetSendEmailForFactorResetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendEmailForFactorResetEnabled",
		val,
	)
}

func (j *jsiiProxy_SecurityNotificationEmails)SetSendEmailForNewDeviceEnabled(val interface{}) {
	if err := j.validateSetSendEmailForNewDeviceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendEmailForNewDeviceEnabled",
		val,
	)
}

func (j *jsiiProxy_SecurityNotificationEmails)SetSendEmailForPasswordChangedEnabled(val interface{}) {
	if err := j.validateSetSendEmailForPasswordChangedEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendEmailForPasswordChangedEnabled",
		val,
	)
}

// Generates CDKTF code for importing a SecurityNotificationEmails resource upon running "cdktf plan <stack-name>".
func SecurityNotificationEmails_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSecurityNotificationEmails_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.securityNotificationEmails.SecurityNotificationEmails",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SecurityNotificationEmails_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecurityNotificationEmails_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.securityNotificationEmails.SecurityNotificationEmails",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecurityNotificationEmails_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecurityNotificationEmails_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.securityNotificationEmails.SecurityNotificationEmails",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecurityNotificationEmails_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecurityNotificationEmails_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.securityNotificationEmails.SecurityNotificationEmails",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecurityNotificationEmails_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.securityNotificationEmails.SecurityNotificationEmails",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SecurityNotificationEmails) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SecurityNotificationEmails) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SecurityNotificationEmails) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityNotificationEmails) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityNotificationEmails) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityNotificationEmails) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityNotificationEmails) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityNotificationEmails) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityNotificationEmails) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityNotificationEmails) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityNotificationEmails) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityNotificationEmails) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityNotificationEmails) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SecurityNotificationEmails) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityNotificationEmails) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SecurityNotificationEmails) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SecurityNotificationEmails) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SecurityNotificationEmails) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SecurityNotificationEmails) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityNotificationEmails) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityNotificationEmails) ResetReportSuspiciousActivityEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetReportSuspiciousActivityEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityNotificationEmails) ResetSendEmailForFactorEnrollmentEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetSendEmailForFactorEnrollmentEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityNotificationEmails) ResetSendEmailForFactorResetEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetSendEmailForFactorResetEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityNotificationEmails) ResetSendEmailForNewDeviceEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetSendEmailForNewDeviceEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityNotificationEmails) ResetSendEmailForPasswordChangedEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetSendEmailForPasswordChangedEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityNotificationEmails) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityNotificationEmails) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityNotificationEmails) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityNotificationEmails) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


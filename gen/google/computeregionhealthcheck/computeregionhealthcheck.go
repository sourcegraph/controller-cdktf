package computeregionhealthcheck

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeregionhealthcheck/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_health_check google_compute_region_health_check}.
type ComputeRegionHealthCheck interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CheckIntervalSec() *float64
	SetCheckIntervalSec(val *float64)
	CheckIntervalSecInput() *float64
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	CreationTimestamp() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GrpcHealthCheck() ComputeRegionHealthCheckGrpcHealthCheckOutputReference
	GrpcHealthCheckInput() *ComputeRegionHealthCheckGrpcHealthCheck
	GrpcTlsHealthCheck() ComputeRegionHealthCheckGrpcTlsHealthCheckOutputReference
	GrpcTlsHealthCheckInput() *ComputeRegionHealthCheckGrpcTlsHealthCheck
	HealthCheckId() *float64
	HealthyThreshold() *float64
	SetHealthyThreshold(val *float64)
	HealthyThresholdInput() *float64
	Http2HealthCheck() ComputeRegionHealthCheckHttp2HealthCheckOutputReference
	Http2HealthCheckInput() *ComputeRegionHealthCheckHttp2HealthCheck
	HttpHealthCheck() ComputeRegionHealthCheckHttpHealthCheckOutputReference
	HttpHealthCheckInput() *ComputeRegionHealthCheckHttpHealthCheck
	HttpsHealthCheck() ComputeRegionHealthCheckHttpsHealthCheckOutputReference
	HttpsHealthCheckInput() *ComputeRegionHealthCheckHttpsHealthCheck
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogConfig() ComputeRegionHealthCheckLogConfigOutputReference
	LogConfigInput() *ComputeRegionHealthCheckLogConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	// Experimental.
	RawOverrides() any
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SelfLink() *string
	SslHealthCheck() ComputeRegionHealthCheckSslHealthCheckOutputReference
	SslHealthCheckInput() *ComputeRegionHealthCheckSslHealthCheck
	TcpHealthCheck() ComputeRegionHealthCheckTcpHealthCheckOutputReference
	TcpHealthCheckInput() *ComputeRegionHealthCheckTcpHealthCheck
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeRegionHealthCheckTimeoutsOutputReference
	TimeoutSec() *float64
	SetTimeoutSec(val *float64)
	TimeoutSecInput() *float64
	TimeoutsInput() any
	Type() *string
	UnhealthyThreshold() *float64
	SetUnhealthyThreshold(val *float64)
	UnhealthyThresholdInput() *float64
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value any)
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
	HasResourceMove() any
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
	MoveTo(moveTarget *string, index any)
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutGrpcHealthCheck(value *ComputeRegionHealthCheckGrpcHealthCheck)
	PutGrpcTlsHealthCheck(value *ComputeRegionHealthCheckGrpcTlsHealthCheck)
	PutHttp2HealthCheck(value *ComputeRegionHealthCheckHttp2HealthCheck)
	PutHttpHealthCheck(value *ComputeRegionHealthCheckHttpHealthCheck)
	PutHttpsHealthCheck(value *ComputeRegionHealthCheckHttpsHealthCheck)
	PutLogConfig(value *ComputeRegionHealthCheckLogConfig)
	PutSslHealthCheck(value *ComputeRegionHealthCheckSslHealthCheck)
	PutTcpHealthCheck(value *ComputeRegionHealthCheckTcpHealthCheck)
	PutTimeouts(value *ComputeRegionHealthCheckTimeouts)
	ResetCheckIntervalSec()
	ResetDescription()
	ResetGrpcHealthCheck()
	ResetGrpcTlsHealthCheck()
	ResetHealthyThreshold()
	ResetHttp2HealthCheck()
	ResetHttpHealthCheck()
	ResetHttpsHealthCheck()
	ResetId()
	ResetLogConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRegion()
	ResetSslHealthCheck()
	ResetTcpHealthCheck()
	ResetTimeouts()
	ResetTimeoutSec()
	ResetUnhealthyThreshold()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for ComputeRegionHealthCheck
type jsiiProxy_ComputeRegionHealthCheck struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeRegionHealthCheck) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) CheckIntervalSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"checkIntervalSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) CheckIntervalSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"checkIntervalSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) GrpcHealthCheck() ComputeRegionHealthCheckGrpcHealthCheckOutputReference {
	var returns ComputeRegionHealthCheckGrpcHealthCheckOutputReference
	_jsii_.Get(
		j,
		"grpcHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) GrpcHealthCheckInput() *ComputeRegionHealthCheckGrpcHealthCheck {
	var returns *ComputeRegionHealthCheckGrpcHealthCheck
	_jsii_.Get(
		j,
		"grpcHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) GrpcTlsHealthCheck() ComputeRegionHealthCheckGrpcTlsHealthCheckOutputReference {
	var returns ComputeRegionHealthCheckGrpcTlsHealthCheckOutputReference
	_jsii_.Get(
		j,
		"grpcTlsHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) GrpcTlsHealthCheckInput() *ComputeRegionHealthCheckGrpcTlsHealthCheck {
	var returns *ComputeRegionHealthCheckGrpcTlsHealthCheck
	_jsii_.Get(
		j,
		"grpcTlsHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) HealthCheckId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) HealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) HealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) Http2HealthCheck() ComputeRegionHealthCheckHttp2HealthCheckOutputReference {
	var returns ComputeRegionHealthCheckHttp2HealthCheckOutputReference
	_jsii_.Get(
		j,
		"http2HealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) Http2HealthCheckInput() *ComputeRegionHealthCheckHttp2HealthCheck {
	var returns *ComputeRegionHealthCheckHttp2HealthCheck
	_jsii_.Get(
		j,
		"http2HealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) HttpHealthCheck() ComputeRegionHealthCheckHttpHealthCheckOutputReference {
	var returns ComputeRegionHealthCheckHttpHealthCheckOutputReference
	_jsii_.Get(
		j,
		"httpHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) HttpHealthCheckInput() *ComputeRegionHealthCheckHttpHealthCheck {
	var returns *ComputeRegionHealthCheckHttpHealthCheck
	_jsii_.Get(
		j,
		"httpHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) HttpsHealthCheck() ComputeRegionHealthCheckHttpsHealthCheckOutputReference {
	var returns ComputeRegionHealthCheckHttpsHealthCheckOutputReference
	_jsii_.Get(
		j,
		"httpsHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) HttpsHealthCheckInput() *ComputeRegionHealthCheckHttpsHealthCheck {
	var returns *ComputeRegionHealthCheckHttpsHealthCheck
	_jsii_.Get(
		j,
		"httpsHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) LogConfig() ComputeRegionHealthCheckLogConfigOutputReference {
	var returns ComputeRegionHealthCheckLogConfigOutputReference
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) LogConfigInput() *ComputeRegionHealthCheckLogConfig {
	var returns *ComputeRegionHealthCheckLogConfig
	_jsii_.Get(
		j,
		"logConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) SslHealthCheck() ComputeRegionHealthCheckSslHealthCheckOutputReference {
	var returns ComputeRegionHealthCheckSslHealthCheckOutputReference
	_jsii_.Get(
		j,
		"sslHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) SslHealthCheckInput() *ComputeRegionHealthCheckSslHealthCheck {
	var returns *ComputeRegionHealthCheckSslHealthCheck
	_jsii_.Get(
		j,
		"sslHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) TcpHealthCheck() ComputeRegionHealthCheckTcpHealthCheckOutputReference {
	var returns ComputeRegionHealthCheckTcpHealthCheckOutputReference
	_jsii_.Get(
		j,
		"tcpHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) TcpHealthCheckInput() *ComputeRegionHealthCheckTcpHealthCheck {
	var returns *ComputeRegionHealthCheckTcpHealthCheck
	_jsii_.Get(
		j,
		"tcpHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) Timeouts() ComputeRegionHealthCheckTimeoutsOutputReference {
	var returns ComputeRegionHealthCheckTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) TimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) TimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) UnhealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheck) UnhealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThresholdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_health_check google_compute_region_health_check} Resource.
func NewComputeRegionHealthCheck(scope constructs.Construct, id *string, config *ComputeRegionHealthCheckConfig) ComputeRegionHealthCheck {
	_init_.Initialize()

	if err := validateNewComputeRegionHealthCheckParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionHealthCheck{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionHealthCheck.ComputeRegionHealthCheck",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_health_check google_compute_region_health_check} Resource.
func NewComputeRegionHealthCheck_Override(c ComputeRegionHealthCheck, scope constructs.Construct, id *string, config *ComputeRegionHealthCheckConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionHealthCheck.ComputeRegionHealthCheck",
		[]any{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheck) SetCheckIntervalSec(val *float64) {
	if err := j.validateSetCheckIntervalSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkIntervalSec",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheck) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheck) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheck) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheck) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheck) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheck) SetHealthyThreshold(val *float64) {
	if err := j.validateSetHealthyThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthyThreshold",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheck) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheck) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheck) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheck) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheck) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheck) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheck) SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheck) SetTimeoutSec(val *float64) {
	if err := j.validateSetTimeoutSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSec",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheck) SetUnhealthyThreshold(val *float64) {
	if err := j.validateSetUnhealthyThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unhealthyThreshold",
		val,
	)
}

// Generates CDKTF code for importing a ComputeRegionHealthCheck resource upon running "cdktf plan <stack-name>".
func ComputeRegionHealthCheck_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateComputeRegionHealthCheck_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionHealthCheck.ComputeRegionHealthCheck",
		"generateConfigForImport",
		[]any{scope, importToId, importFromId, provider},
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
func ComputeRegionHealthCheck_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateComputeRegionHealthCheck_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionHealthCheck.ComputeRegionHealthCheck",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeRegionHealthCheck_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateComputeRegionHealthCheck_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionHealthCheck.ComputeRegionHealthCheck",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeRegionHealthCheck_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateComputeRegionHealthCheck_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionHealthCheck.ComputeRegionHealthCheck",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func ComputeRegionHealthCheck_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.computeRegionHealthCheck.ComputeRegionHealthCheck",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeRegionHealthCheck) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) AddOverride(path *string, value any) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]any{path, value},
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionHealthCheck) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionHealthCheck) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionHealthCheck) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionHealthCheck) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionHealthCheck) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionHealthCheck) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionHealthCheck) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionHealthCheck) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionHealthCheck) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]any{id, provider},
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionHealthCheck) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]any{id},
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) MoveTo(moveTarget *string, index any) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]any{id},
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) PutGrpcHealthCheck(value *ComputeRegionHealthCheckGrpcHealthCheck) {
	if err := c.validatePutGrpcHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGrpcHealthCheck",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) PutGrpcTlsHealthCheck(value *ComputeRegionHealthCheckGrpcTlsHealthCheck) {
	if err := c.validatePutGrpcTlsHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGrpcTlsHealthCheck",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) PutHttp2HealthCheck(value *ComputeRegionHealthCheckHttp2HealthCheck) {
	if err := c.validatePutHttp2HealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttp2HealthCheck",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) PutHttpHealthCheck(value *ComputeRegionHealthCheckHttpHealthCheck) {
	if err := c.validatePutHttpHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpHealthCheck",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) PutHttpsHealthCheck(value *ComputeRegionHealthCheckHttpsHealthCheck) {
	if err := c.validatePutHttpsHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpsHealthCheck",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) PutLogConfig(value *ComputeRegionHealthCheckLogConfig) {
	if err := c.validatePutLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLogConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) PutSslHealthCheck(value *ComputeRegionHealthCheckSslHealthCheck) {
	if err := c.validatePutSslHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSslHealthCheck",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) PutTcpHealthCheck(value *ComputeRegionHealthCheckTcpHealthCheck) {
	if err := c.validatePutTcpHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTcpHealthCheck",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) PutTimeouts(value *ComputeRegionHealthCheckTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ResetCheckIntervalSec() {
	_jsii_.InvokeVoid(
		c,
		"resetCheckIntervalSec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ResetGrpcHealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetGrpcHealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ResetGrpcTlsHealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetGrpcTlsHealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ResetHealthyThreshold() {
	_jsii_.InvokeVoid(
		c,
		"resetHealthyThreshold",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ResetHttp2HealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetHttp2HealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ResetHttpHealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpHealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ResetHttpsHealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpsHealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ResetLogConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetLogConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ResetSslHealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetSslHealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ResetTcpHealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetTcpHealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ResetTimeoutSec() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeoutSec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ResetUnhealthyThreshold() {
	_jsii_.InvokeVoid(
		c,
		"resetUnhealthyThreshold",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheck) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionHealthCheck) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionHealthCheck) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

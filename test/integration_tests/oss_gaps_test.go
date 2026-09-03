package integration_tests

// Reasons passed to testdata.SkipIfOSS, collected here so that each gap is
// stated once and the set of things the SDK does not cover against plain OSS
// Conductor can be read (and counted) in one place.
//
// Keep these accurate. They are the only record of why a test does not run
// against OSS, and a stale one is worse than no comment: it tells the next
// person a feature is Enterprise-only when it may since have shipped in OSS,
// so the gate never gets revisited. When a claim is verified against the OSS
// server source rather than a live run, the note says so and names the class,
// so it can be re-checked without standing a server up.
//
// Whole tests are gated with testdata.SkipIfOSS(t, ossGapX). Individual
// assertion blocks inside an otherwise-shared test are gated with
// testdata.OSSGapSkipped(), which honors the same
// CONDUCTOR_INCLUDE_GATED_TESTS override. Prefer the narrow form: gate only
// the assertions that OSS actually fails, so the rest of the test keeps
// running. A few constants below are therefore referenced only from the
// comment on such a block rather than passed as an argument -- they are still
// the inventory entry for that gap.
const (
	// --- Orkes-Enterprise-only APIs: OSS registers no controller at all. ---

	ossGapApplicationsAPI = "the Applications/Authorization API (/applications) is Orkes-Enterprise-only, confirmed empirically (404 'No static resource api/applications') and by the absence of any such controller in the OSS server"
	ossGapUsersAPI        = "the Users/Authorization API (/users) is Orkes-Enterprise-only, confirmed empirically (404 'No static resource api/users') and by the absence of any such controller in the OSS server"
	ossGapGroupsAPI       = "the Groups/Authorization API (/groups) is Orkes-Enterprise-only, confirmed empirically (404 'No static resource api/groups') and by the absence of any such controller in the OSS server"
	ossGapSchemaAPI       = "the Schema API (/schema) is Orkes-Enterprise-only, confirmed empirically (404 'No static resource api/schema')"
	ossGapServiceRegistry = "the Service Registry API (/service-registry) is Orkes-Enterprise-only, confirmed empirically (404 'No static resource api/service-registry')"
	ossGapWebhooksAPI     = "the Webhooks API (/webhook) is Orkes-Enterprise-only, confirmed empirically (404 'No static resource api/webhook')"
	ossGapIntegrationsAPI = "the Integrations/Prompts APIs (/integrations, /prompts) are Orkes-Enterprise-only, confirmed empirically (404 'No static resource api/integrations|prompts')"

	// --- Tagging: an Orkes-Enterprise metadata extension throughout. ---
	//
	// OSS has no tag endpoints on any resource: no TagResource, and no
	// /tags sub-paths on MetadataResource, SchedulerResource or the
	// secrets/environment controllers.

	ossGapWorkflowDefTags  = "metadata tagging (/metadata/workflow/{name}/tags) is Orkes-Enterprise-only, confirmed empirically; OSS's MetadataResource has no /tags sub-paths"
	ossGapTaskDefTags      = "metadata tagging (/metadata/task/{name}/tags) is Orkes-Enterprise-only, confirmed empirically; OSS's MetadataResource has no /tags sub-paths"
	ossGapScheduleTags     = "schedule tagging (/scheduler/schedules/{name}/tags) is Orkes-Enterprise-only, confirmed empirically (404 'No static resource api/scheduler/schedules/.../tags'); the surrounding scheduler CRUD does work on OSS and is still exercised"
	ossGapEventHandlerTags = "OSS's event handler resource silently drops the Tags field (confirmed empirically: it always comes back empty); the rest of the event handler round-trip does work on OSS and is still asserted"
	ossGapEnvVarTags       = "environment variable tagging (/environment/{name}/tags) is Orkes-Enterprise-only: OSS's EnvironmentResource exposes only GET /environment and GET /environment/{key}"

	// --- Partially implemented on OSS: reads work, writes do not. ---
	//
	// Both of these were previously gated as "Orkes-Enterprise-only",
	// which was wrong -- see TestSecretResourceApiService and
	// TestGetAllEnvVariables for the read coverage that now runs on OSS.

	ossGapSecretWrites = "OSS ships only read-only SecretsDAO backends (env-var, noop), so secret writes return 501 rather than persisting, and its secrets controller has no /tags sub-paths; the read path is covered on OSS by TestSecretReads"
	ossGapEnvVarWrites = "OSS's EnvironmentResource exposes only GET /environment and GET /environment/{key} -- create/update/delete are Orkes-Enterprise-only; the read path is covered on OSS by TestGetAllEnvVariables"

	// --- Endpoints OSS does not implement. ---

	ossGapBulkDelete        = "POST /workflow/bulk/delete is not implemented by plain OSS Conductor, confirmed empirically (404 'No static resource api/workflow/bulk/delete'); OSS's WorkflowBulkResource offers DELETE /workflow/bulk/remove instead, which this SDK does not call"
	ossGapCorrelatedBatch   = "GET /workflow/correlated/batch is not implemented by plain OSS Conductor, confirmed empirically (404 'No static resource api/workflow/correlated/batch'); OSS only has the single-workflow /workflow/{name}/correlated/{correlationId}"
	ossGapWorkflowUpgrade   = "POST /workflow/{id}/upgrade is not implemented by plain OSS Conductor, confirmed empirically (404 'No static resource api/workflow/{id}/upgrade')"
	ossGapWorkflowJump      = "POST /workflow/{id}/jump/{taskRefName} is not implemented by plain OSS Conductor, confirmed empirically (404 'No static resource api/workflow/{id}/jump/{ref}')"
	ossGapWorkflowVariables = "PUT /workflow/{id}/variables is not implemented by plain OSS Conductor, confirmed empirically (404 'No static resource api/workflow/{id}/variables')"
	ossGapWorkflowState     = "PUT /workflow/{id}/state is not implemented by plain OSS Conductor, confirmed empirically (404 'No static resource api/workflow/{id}/state')"

	// --- Behavior differences rather than missing endpoints. ---
	//
	// These read as OSS-side defects rather than deliberate Enterprise
	// features. Each one is a candidate upstream bug report; link an issue
	// here as they get filed, so the gates get revisited instead of
	// hardening into permanent skips.

	ossGapSignalEndpoints = "signal endpoints (/tasks/{workflowId}/{status}/signal(/sync)) are broken on plain OSS Conductor: {status} resolves to the literal string \"signal\" regardless of the actual URL, confirmed empirically via a direct curl bypassing the SDK. OSS's TaskResource declares both POST /{workflowId}/{taskRefName}/{status} and POST /{workflowId}/{status}/signal, so a signal URL matches either pattern and gets routed to the former -- binding the trailing \"signal\" segment to {status}. Nothing in these tests is recoverable while that holds, so the whole test is gated rather than individual assertions"
	ossGapGetWorkflowTask = "the GET_WORKFLOW system task type does not exist on plain OSS Conductor: it has no WorkflowSystemTask bean there, so nothing ever polls that queue and the task stays SCHEDULED forever, confirmed empirically. (Note this is an absent task type, not a worker/config problem -- FORK/JOIN and the other async system tasks do execute on OSS.)"
	ossGapJavaScriptEval  = "the javascript evaluator on plain OSS Conductor doesn't support anonymous function expressions the way Orkes Enterprise does, confirmed empirically (SyntaxError parsing \"function() { return 'true'; }\")"
	ossGapIdempotencyKeys = "workflow start idempotency keys are not honored by plain OSS Conductor, confirmed empirically (a duplicate start with an existing idempotencyKey returns a new workflow ID instead of the existing one / no conflict error)"
	ossGapFailureWorkflow = "terminate-with-failure-workflow-trigger is not honored by plain OSS Conductor, confirmed empirically (conductor.failure_workflow output never gets set)"
	ossGapSkipTaskOutput  = "the skip-task endpoint's TaskOutput override is not applied by plain OSS Conductor, confirmed empirically (skipped task's outputData never contains the provided TaskOutput)"
	ossGapUpdateNotFound  = "OSS returns 200 instead of 404 for updates against a nonexistent workflow ID"
	ossGapSearchPaging    = "GET /workflow/search pagination (start>0) returns 0 results for this IN-operator query on plain OSS Conductor, confirmed empirically"

	// Both return-strategy gaps below cover the ReturnStrategy subtests and
	// the equivalent ExecuteAndGet* helpers, which fail the same way.
	ossGapBlockingWorkflow = "resolving to the blocking sub-workflow doesn't work on plain OSS Conductor, confirmed empirically (the returned TargetWorkflowId/tasks are the main workflow's instead of the blocked-on sub-workflow's)"
	ossGapBlockingTask     = "the BLOCKING_TASK / BLOCKING_TASK_INPUT return strategies 500 on plain OSS Conductor, confirmed empirically (error: {\"status\":500,\"message\":\"value\"})"
)

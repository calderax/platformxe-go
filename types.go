// =============================================================================
// (c) 2026 Caldera Technologies Ltd.
// Proprietary and confidential.
// Unauthorized copying or distribution is prohibited.
// =============================================================================

package platformxe

// APIResponse is the standard PlatformXe response envelope.
type APIResponse struct {
	Success bool                   `json:"success"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Error   *APIError              `json:"error,omitempty"`
}

// APIError represents an error from the API.
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ClientConfig holds configuration for the PlatformXe client.
type ClientConfig struct {
	APIKey   string
	BaseURL  string
	Timeout  int  // seconds, default 10
	Retries  int  // default 2
	FailOpen bool // default true
}

// -- Pagination --

// Pagination holds standard pagination metadata.
type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

// -- Delete Result --

// DeleteResult is the standard response for delete operations.
type DeleteResult struct {
	Deleted bool `json:"deleted"`
}

// -- Processor Configuration --

// ProcessorConfig represents the runtime processor configuration for a service.
type ProcessorConfig struct {
	ID             string                 `json:"id"`
	OrganizationID string                 `json:"organizationId"`
	Service        string                 `json:"service"`
	Enabled        bool                   `json:"enabled"`
	Config         map[string]interface{} `json:"config"`
	CreatedAt      string                 `json:"createdAt"`
	UpdatedAt      string                 `json:"updatedAt"`
}

// -- Health --

// HealthCheckResult is the response from the platform health endpoint.
type HealthCheckResult struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

// -- Messaging --

// EmailHealthResponse is the health status of the email service.
type EmailHealthResponse struct {
	ProviderOrder []string                       `json:"providerOrder"`
	Providers     map[string]EmailProviderStatus `json:"providers"`
	Queue         QueueStats                     `json:"queue"`
}

// EmailProviderStatus is the status of an individual email provider.
type EmailProviderStatus struct {
	Configured bool   `json:"configured"`
	State      string `json:"state"`
	Failures   int    `json:"failures"`
	Position   int    `json:"position"`
}

// SmsHealthResponse is the health status of SMS and WhatsApp services.
type SmsHealthResponse struct {
	Sms      ProviderHealthStatus `json:"sms"`
	Whatsapp ProviderHealthStatus `json:"whatsapp"`
}

// ProviderHealthStatus is the health status of a messaging provider.
type ProviderHealthStatus struct {
	Provider    string `json:"provider"`
	Status      string `json:"status"`
	LastChecked string `json:"lastChecked"`
}

// QueueStats holds email queue statistics.
type QueueStats struct {
	Pending    int `json:"pending"`
	Processing int `json:"processing"`
	Sent       int `json:"sent"`
	DeadLetter int `json:"deadLetter"`
}

// QueueProcessResult is the result of processing the email retry queue.
type QueueProcessResult struct {
	Processed    int `json:"processed"`
	Sent         int `json:"sent"`
	Failed       int `json:"failed"`
	DeadLettered int `json:"deadLettered"`
}

// SendMessageResult is the result of sending an email, SMS, or WhatsApp message.
type SendMessageResult struct {
	MessageID string `json:"messageId"`
}

// -- Webhooks --

// Webhook represents a webhook endpoint configuration.
type Webhook struct {
	ID             string   `json:"id"`
	OrganizationID string   `json:"organizationId"`
	Name           string   `json:"name"`
	URL            string   `json:"url"`
	Events         []string `json:"events"`
	IsActive       bool     `json:"isActive"`
	Secret         string   `json:"secret,omitempty"`
	CreatedAt      string   `json:"createdAt"`
	UpdatedAt      string   `json:"updatedAt"`
}

// WebhookTestResult is the result of a webhook test delivery.
type WebhookTestResult struct {
	Delivered    bool   `json:"delivered"`
	StatusCode   int    `json:"statusCode"`
	ResponseTime int    `json:"responseTime"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// WebhookListResponse wraps a list of webhooks.
type WebhookListResponse struct {
	Webhooks []Webhook `json:"webhooks"`
}

// -- Templates --

// Template represents an email/message template.
type Template struct {
	ID             string                   `json:"id"`
	OrganizationID string                   `json:"organizationId"`
	Name           string                   `json:"name"`
	Subject        string                   `json:"subject"`
	BodyBlocks     []map[string]interface{} `json:"bodyBlocks"`
	Variables      []string                 `json:"variables"`
	CreatedAt      string                   `json:"createdAt"`
	UpdatedAt      string                   `json:"updatedAt"`
}

// TemplateListResponse wraps a list of templates.
type TemplateListResponse struct {
	Templates []Template `json:"templates"`
}

// RenderResult is the result of rendering a template.
type RenderResult struct {
	Subject string `json:"subject"`
	HTML    string `json:"html"`
	Text   string `json:"text"`
}

// TemplateSendResult is the result of rendering and sending a template.
type TemplateSendResult struct {
	Rendered TemplateSendRendered `json:"rendered"`
	Send     TemplateSendInfo     `json:"send"`
}

// TemplateSendRendered holds the rendered fields of a template send.
type TemplateSendRendered struct {
	Subject string `json:"subject"`
}

// TemplateSendInfo holds the send result of a template send.
type TemplateSendInfo struct {
	MessageID string `json:"messageId,omitempty"`
	Provider  string `json:"provider"`
	Queued    bool   `json:"queued"`
}

// -- Sending Domains --

// SendingDomain represents a registered sending domain.
type SendingDomain struct {
	ID             string      `json:"id"`
	OrganizationID string      `json:"organizationId"`
	Domain         string      `json:"domain"`
	Status         string      `json:"status"`
	DNSRecords     []DNSRecord `json:"dnsRecords"`
	VerifiedAt     *string     `json:"verifiedAt"`
	CreatedAt      string      `json:"createdAt"`
	UpdatedAt      string      `json:"updatedAt"`
}

// DNSRecord represents a DNS record for domain verification.
type DNSRecord struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Value    string `json:"value"`
	Verified bool   `json:"verified"`
}

// DomainListResponse wraps a list of sending domains.
type DomainListResponse struct {
	Domains []SendingDomain `json:"domains"`
}

// DomainVerifyResult is the result of a domain verification check.
type DomainVerifyResult struct {
	Domain      SendingDomain `json:"domain"`
	AllVerified bool          `json:"allVerified"`
}

// -- Event Subscriptions --

// EventSubscription represents an event subscription.
type EventSubscription struct {
	ID             string   `json:"id"`
	OrganizationID string   `json:"organizationId"`
	Name           string   `json:"name"`
	WebhookURL     string   `json:"webhookUrl"`
	EventTypes     []string `json:"eventTypes"`
	Status         string   `json:"status"`
	Secret         string   `json:"secret,omitempty"`
	CreatedAt      string   `json:"createdAt"`
	UpdatedAt      string   `json:"updatedAt"`
}

// EventSubscriptionListResponse wraps a list of event subscriptions.
type EventSubscriptionListResponse struct {
	Subscriptions []EventSubscription `json:"subscriptions"`
}

// EventReplayResult is the result of replaying events.
type EventReplayResult struct {
	ReplayedCount int `json:"replayedCount"`
	SuccessCount  int `json:"successCount"`
	FailedCount   int `json:"failedCount"`
}

// EventLogEntry represents a single event log entry.
type EventLogEntry struct {
	ID        string                 `json:"id"`
	EventType string                 `json:"eventType"`
	SourceApp string                 `json:"sourceApp"`
	Payload   map[string]interface{} `json:"payload"`
	CreatedAt string                 `json:"createdAt"`
}

// EventLogResponse is the paginated event log response.
type EventLogResponse struct {
	Items []EventLogEntry `json:"items"`
	Total int             `json:"total"`
	Page  int             `json:"page"`
	Limit int             `json:"limit"`
}

// EmitEventResult is the result of emitting a custom event.
type EmitEventResult struct {
	EventType         string              `json:"eventType"`
	SourceApp         string              `json:"sourceApp"`
	TriggersEvaluated int                 `json:"triggersEvaluated"`
	TriggersMatched   int                 `json:"triggersMatched"`
	TriggersExecuted  int                 `json:"triggersExecuted"`
	Results           []TriggerExecResult `json:"results"`
}

// TriggerExecResult is the result of a single trigger execution.
type TriggerExecResult struct {
	TriggerID     string `json:"triggerId"`
	TriggerName   string `json:"triggerName"`
	Matched       bool   `json:"matched"`
	Executed      bool   `json:"executed"`
	ExecutionID   string `json:"executionId,omitempty"`
	SkippedReason string `json:"skippedReason,omitempty"`
	Error         string `json:"error,omitempty"`
}

// -- Storage --

// StorageFile represents a stored media file.
type StorageFile struct {
	ID               string `json:"id"`
	CallerService    string `json:"callerService"`
	Module           string `json:"module"`
	EntityID         string `json:"entityId"`
	OriginalFilename string `json:"originalFilename"`
	MimeType         string `json:"mimeType"`
	FileSize         int64  `json:"fileSize"`
	StorageProvider  string `json:"storageProvider"`
	PublicURL        string `json:"publicUrl"`
	DisplayOrder     int    `json:"displayOrder"`
	ModerationStatus string `json:"moderationStatus"`
	MediaType        string `json:"mediaType"`
	Width            *int   `json:"width,omitempty"`
	Height           *int   `json:"height,omitempty"`
	CreatedAt        string `json:"createdAt"`
}

// StorageFileListResponse wraps a list of storage files.
type StorageFileListResponse struct {
	Files []StorageFile `json:"files"`
}

// SignUploadResult is the result of signing an upload URL.
type SignUploadResult struct {
	UploadURL string `json:"uploadUrl"`
	FileID    string `json:"fileId"`
}

// ReorderFilesResult is the result of reordering files.
type ReorderFilesResult struct {
	Reordered bool `json:"reordered"`
}

// Document represents a fixed-storage document.
type Document struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description,omitempty"`
	Category     string `json:"category"`
	FileType     string `json:"fileType"`
	Space        string `json:"space"`
	FolderID     string `json:"folderId,omitempty"`
	OwnerType    string `json:"ownerType"`
	OwnerID      string `json:"ownerId"`
	AccessLevel  string `json:"accessLevel"`
	Visibility   string `json:"visibility"`
	FileURL      string `json:"fileUrl,omitempty"`
	FileName     string `json:"fileName,omitempty"`
	FileMimeType string `json:"fileMimeType,omitempty"`
	FileSize     int64  `json:"fileSize,omitempty"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

// DocumentListResponse wraps a list of documents.
type DocumentListResponse struct {
	Documents []Document `json:"documents"`
}

// DocumentFolder represents a storage folder.
type DocumentFolder struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	OwnerType     string `json:"ownerType"`
	OwnerID       string `json:"ownerId"`
	CallerService string `json:"callerService"`
	CreatedAt     string `json:"createdAt"`
}

// FolderListResponse wraps a list of folders.
type FolderListResponse struct {
	Folders []DocumentFolder `json:"folders"`
}

// -- Workflows --

// WorkflowTrigger represents a workflow automation trigger.
type WorkflowTrigger struct {
	ID               string                   `json:"id"`
	Name             string                   `json:"name"`
	Description      string                   `json:"description,omitempty"`
	Code             string                   `json:"code"`
	TriggerType      string                   `json:"triggerType"`
	EventType        string                   `json:"eventType,omitempty"`
	CronExpression   string                   `json:"cronExpression,omitempty"`
	Conditions       map[string]interface{}   `json:"conditions,omitempty"`
	EntityType       string                   `json:"entityType,omitempty"`
	Actions          []map[string]interface{} `json:"actions"`
	IsActive         bool                     `json:"isActive"`
	Priority         int                      `json:"priority"`
	CooldownMinutes  *int                     `json:"cooldownMinutes,omitempty"`
	MaxExecutionsDay *int                     `json:"maxExecutionsDay,omitempty"`
	SourceApp        string                   `json:"sourceApp,omitempty"`
	CreatedAt        string                   `json:"createdAt"`
	UpdatedAt        string                   `json:"updatedAt"`
}

// WorkflowListResponse wraps a list of workflow triggers.
type WorkflowListResponse struct {
	Workflows []WorkflowTrigger `json:"workflows"`
}

// WorkflowEvalResult is the result of evaluating workflows against an event.
type WorkflowEvalResult struct {
	Matched int                  `json:"matched"`
	Total   int                  `json:"total"`
	Results []WorkflowEvalDetail `json:"results"`
}

// WorkflowEvalDetail is the evaluation result for a single workflow trigger.
type WorkflowEvalDetail struct {
	TriggerID   string `json:"triggerId"`
	TriggerCode string `json:"triggerCode"`
	TriggerName string `json:"triggerName"`
	Matched     bool   `json:"matched"`
	Reason      string `json:"reason,omitempty"`
	ExecutionID string `json:"executionId,omitempty"`
	Error       string `json:"error,omitempty"`
}

// -- Permissions --

// PermissionRole represents a permission role.
type PermissionRole struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organizationId"`
	Name           string `json:"name"`
	Description    string `json:"description,omitempty"`
	Model          string `json:"model"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
}

// RoleListResponse wraps a list of permission roles.
type RoleListResponse struct {
	Roles []PermissionRole `json:"roles"`
}

// PermissionCheckResult is the result of a single permission check.
type PermissionCheckResult struct {
	Allowed     bool   `json:"allowed"`
	Source      string `json:"source"`
	RoleID      string `json:"roleId,omitempty"`
	EvaluatedAt string `json:"evaluatedAt"`
}

// BatchCheckResponse wraps a list of batch permission check results.
type BatchCheckResponse struct {
	Results []BatchCheckResult `json:"results"`
}

// BatchCheckResult is the result of a single check within a batch.
type BatchCheckResult struct {
	AdminID string `json:"adminId"`
	Path    string `json:"path"`
	Action  string `json:"action"`
	Allowed bool   `json:"allowed"`
	Reason  string `json:"reason,omitempty"`
}

// ResolvedCapabilities is the full resolved permission set for an admin.
type ResolvedCapabilities struct {
	AdminID      string               `json:"adminId"`
	Capabilities []string             `json:"capabilities"`
	Permissions  []ResolvedPermission `json:"permissions,omitempty"`
	TTL          int                  `json:"ttl"`
}

// ResolvedPermission represents a resolved module permission.
type ResolvedPermission struct {
	ModuleID string   `json:"moduleId"`
	Actions  []string `json:"actions"`
}

// RoleCapabilities holds capabilities for a role (Simple model).
type RoleCapabilities struct {
	RoleID       string   `json:"roleId"`
	Capabilities []string `json:"capabilities"`
}

// RoleModulePermissions holds module permissions for a role (Full model).
type RoleModulePermissions struct {
	RoleID  string             `json:"roleId"`
	Modules []ModulePermission `json:"modules"`
}

// ModulePermission represents a single module permission assignment.
type ModulePermission struct {
	ModuleID string   `json:"moduleId"`
	Actions  []string `json:"actions"`
}

// PermissionOverride represents an admin-level permission override.
type PermissionOverride struct {
	ID             string  `json:"id"`
	OrganizationID string  `json:"organizationId"`
	AdminID        string  `json:"adminId"`
	Path           string  `json:"path"`
	Action         string  `json:"action"`
	Effect         string  `json:"effect"`
	Reason         string  `json:"reason"`
	ExpiresAt      *string `json:"expiresAt,omitempty"`
	CreatedAt      string  `json:"createdAt"`
	CreatedBy      string  `json:"createdBy"`
}

// OverrideListResponse wraps a list of permission overrides.
type OverrideListResponse struct {
	Overrides []PermissionOverride `json:"overrides"`
}

// PermissionModule represents a registered permission module.
type PermissionModule struct {
	Key         string   `json:"key"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Actions     []string `json:"actions"`
}

// ModuleListResponse wraps a list of permission modules.
type ModuleListResponse struct {
	Modules []PermissionModule `json:"modules"`
}

// ResourcePolicy represents an ABAC resource policy.
type ResourcePolicy struct {
	ID             string                 `json:"id"`
	OrganizationID string                 `json:"organizationId"`
	Path           string                 `json:"path"`
	Action         string                 `json:"action"`
	Condition      map[string]interface{} `json:"condition"`
	Effect         string                 `json:"effect"`
	Priority       int                    `json:"priority,omitempty"`
	Description    string                 `json:"description,omitempty"`
	IsActive       bool                   `json:"isActive"`
	CreatedAt      string                 `json:"createdAt"`
	UpdatedAt      string                 `json:"updatedAt"`
}

// PolicyListResponse wraps a list of resource policies.
type PolicyListResponse struct {
	Policies []ResourcePolicy `json:"policies"`
}

// RelationshipTuple represents a ReBAC relationship tuple.
type RelationshipTuple struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organizationId"`
	ObjectType     string `json:"objectType"`
	ObjectID       string `json:"objectId"`
	Relation       string `json:"relation"`
	SubjectType    string `json:"subjectType"`
	SubjectID      string `json:"subjectId"`
	CreatedAt      string `json:"createdAt"`
}

// RelationshipListResponse wraps a list of relationship tuples.
type RelationshipListResponse struct {
	Relationships []RelationshipTuple `json:"relationships"`
}

// RelationshipUpdateResult is the result of a relationship batch operation.
type RelationshipUpdateResult struct {
	Created int `json:"created"`
	Deleted int `json:"deleted"`
	Errors  int `json:"errors"`
}

// PermissionAuditLog represents a permission decision audit log entry.
type PermissionAuditLog struct {
	ID        string                 `json:"id"`
	AdminID   string                 `json:"adminId"`
	Path      string                 `json:"path"`
	Action    string                 `json:"action"`
	Allowed   bool                   `json:"allowed"`
	Source    string                 `json:"source"`
	Context   map[string]interface{} `json:"context"`
	LatencyMs int                    `json:"latencyMs"`
	CreatedAt string                 `json:"createdAt"`
}

// AuditLogListResponse wraps a paginated list of audit logs.
type AuditLogListResponse struct {
	Items []PermissionAuditLog `json:"items"`
	Total int                  `json:"total"`
	Page  int                  `json:"page"`
	Limit int                  `json:"limit"`
}

// PermissionChangeLog represents a permission mutation change log entry.
type PermissionChangeLog struct {
	ID             string                 `json:"id"`
	OrganizationID string                 `json:"organizationId"`
	EntityType     string                 `json:"entityType"`
	EntityID       string                 `json:"entityId"`
	Changes        map[string]interface{} `json:"changes"`
	ChangedBy      string                 `json:"changedBy"`
	Timestamp      string                 `json:"timestamp"`
}

// ChangeLogListResponse wraps a paginated list of change logs.
type ChangeLogListResponse struct {
	Items []PermissionChangeLog `json:"items"`
	Total int                   `json:"total"`
	Page  int                   `json:"page"`
	Limit int                   `json:"limit"`
}

// AuditExportResult is the result of exporting audit logs.
type AuditExportResult struct {
	Decisions   []PermissionAuditLog  `json:"decisions"`
	Changes     []PermissionChangeLog `json:"changes"`
	PeriodStart string                `json:"periodStart,omitempty"`
	PeriodEnd   string                `json:"periodEnd,omitempty"`
	ExportedAt  string                `json:"exportedAt"`
}

// ShadowCheckResult is the result of a shadow permission check.
type ShadowCheckResult struct {
	AdminID        string `json:"adminId"`
	Path           string `json:"path"`
	Action         string `json:"action"`
	LocalDecision  bool   `json:"localDecision"`
	RemoteDecision bool   `json:"remoteDecision"`
	Match          bool   `json:"match"`
	Discrepancy    bool   `json:"discrepancy"`
}

// -- Federation --

// FederationGroup represents a federation group.
type FederationGroup struct {
	ID         string             `json:"id"`
	Name       string             `json:"name"`
	OwnerOrgID string             `json:"ownerOrgId"`
	Members    []FederationMember `json:"members"`
	CreatedAt  string             `json:"createdAt"`
}

// FederationGroupListResponse wraps a list of federation groups.
type FederationGroupListResponse struct {
	Groups []FederationGroup `json:"groups"`
}

// FederationMember represents a member of a federation group.
type FederationMember struct {
	OrganizationID string `json:"organizationId"`
	Prefix         string `json:"prefix"`
	JoinedAt       string `json:"joinedAt"`
}

// FederationMemberResult is the result of adding a member to a federation group.
type FederationMemberResult struct {
	OrganizationID string `json:"organizationId"`
	GroupID        string `json:"groupId"`
	Prefix         string `json:"prefix"`
	JoinedAt       string `json:"joinedAt"`
}

// FederationRemoveResult is the result of removing a member from a federation group.
type FederationRemoveResult struct {
	OrganizationID string `json:"organizationId"`
	GroupID        string `json:"groupId"`
	RemovedAt      string `json:"removedAt"`
}

// FederationPullResult is the result of pulling modules from a federation group.
type FederationPullResult struct {
	OrganizationID string                 `json:"organizationId"`
	Modules        []FederationModuleInfo `json:"modules"`
}

// FederationModuleInfo describes a module pulled from a federation group.
type FederationModuleInfo struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Paths []string `json:"paths"`
}

// FederationPushResult is the result of pushing permissions to federation members.
type FederationPushResult struct {
	OrganizationID string                `json:"organizationId"`
	AdminID        string                `json:"adminId"`
	Permissions    []FederationPermEntry `json:"permissions"`
	Status         string                `json:"status"`
	Error          string                `json:"error,omitempty"`
}

// FederationPermEntry describes a single permission entry in a federation push.
type FederationPermEntry struct {
	Path   string `json:"path"`
	Action string `json:"action"`
}

// FederationStatusResult is the result of querying federation group status.
type FederationStatusResult struct {
	Group       FederationGroupInfo `json:"group"`
	Members     []FederationMember  `json:"members"`
	RecentSyncs []FederationSync    `json:"recentSyncs"`
}

// FederationGroupInfo is a summary of a federation group.
type FederationGroupInfo struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	OwnerOrgID string `json:"ownerOrgId"`
	CreatedAt  string `json:"createdAt"`
}

// FederationSync represents a federation sync log entry.
type FederationSync struct {
	ID        string                 `json:"id"`
	GroupID   string                 `json:"groupId"`
	Type      string                 `json:"type"`
	Status    string                 `json:"status"`
	Timestamp string                 `json:"timestamp"`
	Details   map[string]interface{} `json:"details,omitempty"`
}

// -- Identity --

// IdentityResolveResult is the result of resolving an identity.
type IdentityResolveResult struct {
	Resolved          bool                   `json:"resolved"`
	LinkedIdentifiers []IdentityIdentifier   `json:"linkedIdentifiers"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
}

// IdentityIdentifier represents a linked identity identifier.
type IdentityIdentifier struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// IdentityVerifyResult is the result of verifying an identity.
type IdentityVerifyResult struct {
	Verified     bool                  `json:"verified"`
	MatchScore   float64               `json:"matchScore"`
	MatchDetails *IdentityMatchDetails `json:"matchDetails,omitempty"`
}

// IdentityMatchDetails holds name match details for identity verification.
type IdentityMatchDetails struct {
	FirstName IdentityFieldMatch `json:"firstName"`
	LastName  IdentityFieldMatch `json:"lastName"`
}

// IdentityFieldMatch represents the match result for a single field.
type IdentityFieldMatch struct {
	Match      bool    `json:"match"`
	Confidence float64 `json:"confidence"`
}

// IdentityLookupResult is the result of looking up an identity.
type IdentityLookupResult struct {
	Type      string                 `json:"type"`
	Value     string                 `json:"value"`
	Data      map[string]interface{} `json:"data"`
	Timestamp string                 `json:"timestamp"`
}

// IdentityProviderInfo describes an available identity provider.
type IdentityProviderInfo struct {
	Name         string `json:"name"`
	Status       string `json:"status"`
	LastChecked  string `json:"lastChecked"`
	ResponseTime *int   `json:"responseTime,omitempty"`
}

// IdentityProvidersResponse wraps a list of identity providers.
type IdentityProvidersResponse struct {
	Providers []IdentityProviderInfo `json:"providers"`
}

// -- OCR --

// IdentityDocVerifyResult is the result of OCR identity document verification.
type IdentityDocVerifyResult struct {
	Verified        bool              `json:"verified"`
	Confidence      float64           `json:"confidence"`
	Names           DocVerifyNames    `json:"names"`
	DocumentDetails map[string]string `json:"documentDetails,omitempty"`
}

// DocVerifyNames holds name matching details for document verification.
type DocVerifyNames struct {
	DocumentName string  `json:"documentName"`
	ProfileName  string  `json:"profileName"`
	MatchScore   float64 `json:"matchScore"`
}

// -- QR --

// QRCodeResult is the result of generating a QR code.
type QRCodeResult struct {
	QR     string `json:"qr"`
	Format string `json:"format"`
}

// -- Telemetry --

// TelemetryBatch is the full telemetry payload required by the API.
type TelemetryBatch struct {
	SourceApp      string            `json:"sourceApp"`
	PackageName    string            `json:"packageName"`
	PackageVersion string            `json:"packageVersion,omitempty"`
	Metrics        []TelemetryMetric `json:"metrics"`
	PeriodStart    string            `json:"periodStart"`
	PeriodEnd      string            `json:"periodEnd"`
}

// TelemetryMetric is a single metric in a telemetry batch.
type TelemetryMetric struct {
	MetricType string `json:"metricType"`
	MetricKey  string `json:"metricKey"`
	Count      int    `json:"count"`
}

// TelemetryResult is the result of sending a telemetry batch.
type TelemetryResult struct {
	Ingested    int    `json:"ingested"`
	SourceApp   string `json:"sourceApp"`
	PackageName string `json:"packageName"`
}

// -- Exports --

// ExportResult is the result of creating or querying a data export.
type ExportResult struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
	DataType    string `json:"dataType"`
	Format      string `json:"format"`
	DownloadURL string `json:"downloadUrl,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
}

// -- Usage --

// UsageSummary is the usage metrics response for a billing month.
type UsageSummary struct {
	Month               string         `json:"month"`
	Services            []ServiceUsage `json:"services"`
	TotalBillableEvents int            `json:"totalBillableEvents"`
}

// ServiceUsage represents usage for a single service.
type ServiceUsage struct {
	Service     string  `json:"service"`
	Used        int     `json:"used"`
	Limit       int     `json:"limit"`
	PercentUsed float64 `json:"percentUsed"`
}

// -- Threads --

// Thread represents a messaging thread.
type Thread struct {
	ID             string                 `json:"id"`
	OrganizationID string                 `json:"organizationId"`
	ChannelID      string                 `json:"channelId"`
	ChannelSlug    string                 `json:"channelSlug"`
	EntityID       string                 `json:"entityId"`
	Subject        string                 `json:"subject"`
	Status         string                 `json:"status"`
	Priority       string                 `json:"priority"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	Participants   []ThreadParticipantInfo `json:"participants"`
	CreatedAt      string                 `json:"createdAt"`
	UpdatedAt      string                 `json:"updatedAt"`
	ClosedAt       *string                `json:"closedAt,omitempty"`
}

// ThreadParticipantInfo represents a participant in a thread response.
type ThreadParticipantInfo struct {
	ID          string `json:"id"`
	Role        string `json:"role"`
	ExternalID  string `json:"externalId"`
	DisplayName string `json:"displayName"`
	AvatarURL   string `json:"avatarUrl,omitempty"`
	JoinedAt    string `json:"joinedAt"`
}

// ThreadListResponse wraps a paginated list of threads.
type ThreadListResponse struct {
	Threads []Thread `json:"threads"`
	Total   int      `json:"total"`
	Page    int      `json:"page"`
	Limit   int      `json:"limit"`
}

// ThreadMessage represents a message in a thread.
type ThreadMessage struct {
	ID               string   `json:"id"`
	ThreadID         string   `json:"threadId"`
	SenderExternalID string   `json:"senderExternalId"`
	SenderRole       string   `json:"senderRole"`
	Content          string   `json:"content"`
	Visibility       []string `json:"visibility"`
	IsSystem         bool     `json:"isSystem"`
	CreatedAt        string   `json:"createdAt"`
}

// ThreadCloseResult is the result of closing a thread.
type ThreadCloseResult struct {
	ThreadID string `json:"threadId"`
	Status   string `json:"status"`
	ClosedAt string `json:"closedAt"`
}

// ThreadReopenResult is the result of reopening a thread.
type ThreadReopenResult struct {
	ThreadID   string `json:"threadId"`
	Status     string `json:"status"`
	ReopenedAt string `json:"reopenedAt"`
}

// EntityEventResult is the result of forwarding an entity event.
type EntityEventResult struct {
	Evaluated   int    `json:"evaluated"`
	Matched     int    `json:"matched"`
	ActionsRun  int    `json:"actionsRun"`
	ThreadID    string `json:"threadId,omitempty"`
}

// MarkReadResult is the result of marking a thread as read.
type MarkReadResult struct {
	ThreadID string `json:"threadId"`
	ReadAt   string `json:"readAt"`
}

// InboxResponse wraps a paginated inbox response.
type InboxResponse struct {
	Threads []InboxThread `json:"threads"`
	Total   int           `json:"total"`
	Page    int           `json:"page"`
	Limit   int           `json:"limit"`
}

// InboxThread represents a thread in the inbox with unread count.
type InboxThread struct {
	Thread      Thread `json:"thread"`
	UnreadCount int    `json:"unreadCount"`
	LastMessage string `json:"lastMessage"`
}

// UnreadCountResult is the result of querying total unread count.
type UnreadCountResult struct {
	Count int `json:"count"`
}

// ThreadFlag represents a flag on a message.
type ThreadFlag struct {
	ID                  string `json:"id"`
	MessageID           string `json:"messageId"`
	Reason              string `json:"reason"`
	Note                string `json:"note,omitempty"`
	FlaggedByExternalID string `json:"flaggedByExternalId"`
	FlaggedByRole       string `json:"flaggedByRole"`
	CreatedAt           string `json:"createdAt"`
}

// FlagListResponse wraps a list of thread flags.
type FlagListResponse struct {
	Flags []ThreadFlag `json:"flags"`
}

// FlagMessageResult is the result of flagging a message.
type FlagMessageResult struct {
	FlagID    string `json:"flagId"`
	MessageID string `json:"messageId"`
	ThreadID  string `json:"threadId"`
}

// EscalateThreadResult is the result of escalating a thread.
type EscalateThreadResult struct {
	ThreadID    string `json:"threadId"`
	Priority    string `json:"priority"`
	EscalatedAt string `json:"escalatedAt"`
}

// EscalationConfig represents escalation configuration for a channel.
type EscalationConfig struct {
	Enabled    bool                   `json:"enabled"`
	Rules      map[string]interface{} `json:"rules,omitempty"`
	WebhookURL string                 `json:"webhookUrl,omitempty"`
}

// ThreadChannel represents a channel configuration.
type ThreadChannel struct {
	ID                string                 `json:"id"`
	OrganizationID    string                 `json:"organizationId"`
	Slug              string                 `json:"slug"`
	DisplayName       string                 `json:"displayName"`
	EntityType        string                 `json:"entityType"`
	ParticipantRoles  []string               `json:"participantRoles"`
	DefaultVisibility []string               `json:"defaultVisibility"`
	LifecycleRules    map[string]interface{} `json:"lifecycleRules"`
	EscalationConfig  map[string]interface{} `json:"escalationConfig"`
	WebhookURL        string                 `json:"webhookUrl"`
	IsActive          bool                   `json:"isActive"`
	CreatedAt         string                 `json:"createdAt"`
	UpdatedAt         string                 `json:"updatedAt"`
}

// ChannelListResponse wraps a list of thread channels.
type ChannelListResponse struct {
	Channels []ThreadChannel `json:"channels"`
}

// UpdateThreadInput is the input for updating a thread.
type UpdateThreadInput struct {
	Subject  *string                `json:"subject,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// MessageListResponse wraps a paginated list of thread messages.
type MessageListResponse struct {
	Messages []ThreadMessage `json:"messages"`
	Total    int             `json:"total"`
}

// AddParticipantInput is the input for adding a participant to a thread.
type AddParticipantInput struct {
	Role        string `json:"role"`
	ExternalID  string `json:"externalId"`
	DisplayName string `json:"displayName"`
	AvatarURL   string `json:"avatarUrl,omitempty"`
}

// UpdateParticipantInput is the input for updating a thread participant.
type UpdateParticipantInput struct {
	DisplayName *string `json:"displayName,omitempty"`
	AvatarURL   *string `json:"avatarUrl,omitempty"`
	IsMuted     *bool   `json:"isMuted,omitempty"`
}

// ThreadReadState represents the read state for a participant in a thread.
type ThreadReadState struct {
	ThreadID    string `json:"threadId"`
	ExternalID  string `json:"externalId"`
	Role        string `json:"role"`
	LastReadAt  string `json:"lastReadAt"`
}

// ReadStatesResponse wraps a list of thread read states.
type ReadStatesResponse struct {
	ReadStates []ThreadReadState `json:"readStates"`
}

// ReviewFlagInput is the input for reviewing a flagged message.
type ReviewFlagInput struct {
	Status     string `json:"status"`
	ReviewNote string `json:"reviewNote,omitempty"`
	ReviewedBy string `json:"reviewedBy,omitempty"`
}

// UploadFile describes a file for batch upload in Go.
type UploadFile struct {
	Filename    string `json:"filename"`
	ContentType string `json:"contentType"`
	Data        string `json:"data"` // base64-encoded
}

// BatchUploadResult is the result of a batch file upload.
type BatchUploadResult struct {
	Files     []StorageFile `json:"files"`
	Succeeded int           `json:"succeeded"`
	Failed    int           `json:"failed"`
}

// UploadResult is the result of a single file upload.
type UploadResult struct {
	File StorageFile `json:"file"`
}

// BatchQRResult is the result of batch QR code generation.
type BatchQRResult struct {
	Results        []QRCodeResult `json:"results"`
	TotalGenerated int            `json:"totalGenerated"`
	Errors         []string       `json:"errors"`
}

// DryRunResult is the result of a workflow dry-run.
type DryRunResult struct {
	TriggerID   string                 `json:"triggerId"`
	TriggerName string                 `json:"triggerName"`
	Matched     bool                   `json:"matched"`
	Reason      string                 `json:"reason,omitempty"`
	Actions     []map[string]interface{} `json:"actions,omitempty"`
}

// DeletionOverride represents a deletion override request/result.
type DeletionOverride struct {
	ID           string  `json:"id"`
	DocumentID   string  `json:"documentId"`
	RequestedBy  string  `json:"requestedBy"`
	Reason       string  `json:"reason"`
	Status       string  `json:"status"`
	ReviewedBy   *string `json:"reviewedBy,omitempty"`
	ReviewNote   *string `json:"reviewNote,omitempty"`
	CreatedAt    string  `json:"createdAt"`
	UpdatedAt    string  `json:"updatedAt"`
}

// PdfResult is a generic PDF generation result.
type PdfResult struct {
	URL      string `json:"url,omitempty"`
	FileID   string `json:"fileId,omitempty"`
	Filename string `json:"filename,omitempty"`
}

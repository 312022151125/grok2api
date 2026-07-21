# Graph Report - .  (2026-07-21)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 6009 nodes · 16192 edges · 257 communities (223 shown, 34 thin omitted)
- Extraction: 84% EXTRACTED · 16% INFERRED · 0% AMBIGUOUS · INFERRED: 2665 edges (avg confidence: 0.8)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `57c7b5df`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- Community 0
- Community 1
- Community 2
- Community 3
- Community 4
- Community 5
- Community 6
- Community 7
- Community 8
- Community 9
- Community 10
- Community 11
- Community 12
- Community 13
- Community 14
- Community 15
- Community 16
- Community 17
- Community 18
- Community 19
- Community 20
- Community 21
- Community 22
- Community 23
- Community 24
- Community 25
- Community 26
- Community 27
- Community 28
- Community 29
- Community 30
- Community 31
- Community 32
- Community 33
- Community 34
- Community 35
- Community 36
- Community 37
- Community 38
- Community 39
- Community 40
- Community 41
- Community 42
- Community 43
- Community 44
- Community 45
- Community 46
- Community 47
- Community 48
- Community 49
- Community 50
- Community 51
- Community 52
- Community 53
- Community 54
- Community 55
- Community 56
- Community 57
- Community 58
- Community 59
- Community 60
- Community 61
- Community 62
- Community 63
- Community 64
- Community 65
- Community 66
- Community 67
- Community 68
- Community 69
- Community 70
- Community 71
- Community 72
- Community 73
- Community 74
- Community 75
- Community 76
- Community 77
- Community 78
- Community 79
- Community 80
- Community 81
- Community 82
- Community 83
- Community 84
- Community 85
- Community 86
- Community 87
- Community 88
- Community 89
- Community 90
- Community 91
- Community 92
- Community 93
- Community 94
- Community 95
- Community 96
- Community 97
- Community 98
- Community 99
- Community 100
- Community 101
- Community 102
- Community 103
- Community 104
- Community 105
- Community 106
- Community 107
- Community 108
- Community 109
- Community 110
- Community 111
- Community 112
- Community 113
- Community 114
- Community 115
- Community 116
- Community 117
- Community 118
- Community 119
- Community 120
- Community 121
- Community 122
- Community 123
- Community 124
- Community 125
- Community 126
- Community 127
- Community 128
- Community 129
- Community 130
- Community 131
- Community 132
- Community 133
- Community 134
- Community 135
- Community 136
- Community 137
- Community 138
- Community 139
- Community 140
- Community 141
- Community 142
- Community 143
- Community 144
- Community 145
- Community 146
- Community 147
- Community 148
- Community 149
- Community 150
- Community 151
- Community 152
- Community 153
- Community 154
- Community 155
- Community 156
- Community 157
- Community 158
- Community 159
- Community 160
- Community 161
- Community 162
- Community 163
- Community 164
- Community 165
- Community 166
- Community 167
- Community 168
- Community 169
- Community 170
- Community 171
- Community 172
- Community 173
- Community 174
- Community 175
- Community 176
- Community 177
- Community 178
- Community 179
- Community 180
- Community 181
- Community 182
- Community 183
- Community 184
- Community 185
- Community 186
- Community 187
- Community 188
- Community 189
- Community 190
- Community 191
- Community 192
- Community 193
- Community 194
- Community 195
- Community 196
- Community 197
- Community 198
- Community 199
- Community 200
- Community 201
- Community 202
- Community 203
- Community 204
- Community 205
- Community 206
- Community 207
- Community 208
- Community 209
- Community 210
- Community 211
- Community 212
- Community 213
- Community 214
- Community 215
- Community 218
- Community 219
- Community 220
- Community 221
- Community 222
- Community 223
- Community 224
- Community 225
- Community 226
- Community 227
- Community 228
- Community 229
- Community 230
- Community 231
- Community 232
- Community 233
- Community 234
- Community 235
- Community 236
- Community 237
- Community 238
- Community 239
- Community 240
- Community 241
- Community 242
- Community 243
- Community 244
- Community 245
- Community 256

## God Nodes (most connected - your core abstractions)
1. `Credential` - 175 edges
2. `Provider` - 139 edges
3. `Service` - 123 edges
4. `Duration` - 123 edges
5. `cn()` - 119 edges
6. `OpenSQLite()` - 115 edges
7. `NewAccountRepository()` - 96 edges
8. `Response` - 89 edges
9. `NewCipher()` - 75 edges
10. `apiRequest()` - 66 edges

## Surprising Connections (you probably didn't know these)
- `ConvertResponseJSONWithOptions()` --calls--> `applyStopSequences()`  [INFERRED]
  backend/internal/infra/provider/conversation/response.go → backend/internal/infra/provider/conversation/messages_response.go
- `New()` --calls--> `newStartupState()`  [INFERRED]
  backend/internal/app/application.go → backend/internal/app/startup.go
- `New()` --calls--> `readinessSnapshot()`  [INFERRED]
  backend/internal/app/application.go → backend/internal/app/startup.go
- `New()` --calls--> `NewSelector()`  [INFERRED]
  backend/internal/app/application.go → backend/internal/application/gateway/selector.go
- `New()` --calls--> `NewServiceWithTickets()`  [INFERRED]
  backend/internal/app/application.go → backend/internal/application/media/service.go

## Import Cycles
- None detected.

## Communities (257 total, 34 thin omitted)

### Community 0 - "Community 0"
Cohesion: 0.06
Nodes (71): Badge(), BadgeProps, badgeVariants, DialogContent, DialogDescription, DialogFooter(), DialogHeader(), DialogOverlay (+63 more)

### Community 1 - "Community 1"
Cohesion: 0.03
Nodes (84): Calendar(), Message(), MessageAvatar(), MessageContent(), MessageFooter(), MessageGroup(), MessageHeader(), MessageScroller() (+76 more)

### Community 2 - "Community 2"
Cohesion: 0.04
Nodes (92): createModel(), CreateModelInput, decodeModelAccounts, decodeModelPage, decodeModelRoute, deleteModels(), listModelAccountOptions(), ListModelsInput (+84 more)

### Community 3 - "Community 3"
Cohesion: 0.05
Nodes (89): deleteModel(), acceptWebAccountTerms(), AccountBatchResultDTO, AccountCleanupStatus, AccountImportResultDTO, AccountProvider, AccountSummaryDTO, AccountSyncStrategy (+81 more)

### Community 4 - "Community 4"
Cohesion: 0.05
Nodes (64): accountReader, billingSynchronizer, identitySynchronizer, modelSynchronizer, providerPolicy, quotaSynchronizer, Result, Service (+56 more)

### Community 5 - "Community 5"
Cohesion: 0.04
Nodes (21): Bool, Context, Int64, Time, authRescueAdapter, credentialFailureImageAdapter, crossProviderBuildAdapter, crossProviderBuildOKAdapter (+13 more)

### Community 6 - "Community 6"
Cohesion: 0.05
Nodes (53): BuildSSOCookie(), ApplyChromiumClientHints(), Header, applyChromiumClientHints(), applyHeaders(), Header, Lease, Request (+45 more)

### Community 7 - "Community 7"
Cohesion: 0.05
Nodes (65): consoleEndpoint(), conversationErrorType(), Adapter, Once, ReadCloser, RWMutex, invalidConversationResponse(), jsonProviderResponse() (+57 more)

### Community 8 - "Community 8"
Cohesion: 0.07
Nodes (24): CredentialRefreshDueAt(), applyAccountStatusFilter(), currentWebTermsAcceptedAt(), Context, Database, DB, AccountRepository, Time (+16 more)

### Community 9 - "Community 9"
Cohesion: 0.07
Nodes (56): anyReplayCallKeyExists(), assistantContentEqual(), assistantParts(), comparableReplayCallIDs(), filterReplayItemsForInput(), RawMessage, insertReplayItems(), lastAssistantMessage() (+48 more)

### Community 10 - "Community 10"
Cohesion: 0.06
Nodes (55): ChartConfig, ChartContainer, ChartContext, ChartContextProps, ChartLegendContent, ChartTooltipContent, THEMES, DashboardDTO (+47 more)

### Community 11 - "Community 11"
Cohesion: 0.07
Nodes (30): Provider, applyAuditEgress(), applyMediaJobEgress(), primaryEgressScope(), Aliases(), consoleAlias(), AccountIdentityAdapter, AccountModelCapabilityNormalizer (+22 more)

### Community 12 - "Community 12"
Cohesion: 0.06
Nodes (62): Catalog(), buildWebChatPayload(), consumeUpstream(), normalizeOpenAIInput(), preflightUpstream(), buildDirectFileUploadBody(), buildImageEditPayload(), directFileUploadFallbackStatus() (+54 more)

### Community 13 - "Community 13"
Cohesion: 0.11
Nodes (48): applyDomainConfig(), Config, Context, Mutex, RWMutex, AuditConfig, BatchConfig, ClientKeyDefaultsConfig (+40 more)

### Community 14 - "Community 14"
Cohesion: 0.06
Nodes (46): deniedQuotaRefreshLock, AuditConfig, BatchConfig, ClientKeyDefaultsConfig, FrontendConfig, MediaConfig, ProviderBuildConfig, ProviderConsoleConfig (+38 more)

### Community 15 - "Community 15"
Cohesion: 0.06
Nodes (33): consoleSSOCodecAdapter, credentialRefreshAdapter, Context, Int64, Int64, decodeJWTClaims(), Time, importedCredentialExpiry() (+25 more)

### Community 16 - "Community 16"
Cohesion: 0.09
Nodes (11): webQuotaRefreshRequest, Service, AccountRepository, AuditRepository, Context, Group, Logger, Mutex (+3 more)

### Community 17 - "Community 17"
Cohesion: 0.09
Nodes (44): newStartupState(), T, TestReadinessKeepsBuildReadyWhenWebIsUnavailable(), TestReadinessRestoresPersistedCooldownWithoutUpstreamProbe(), TestReadinessStartupReportDoesNotExposeInternalErrors(), T, TestCleanupAccountsDeletesOnlySelectedCurrentStatuses(), TestCleanupAccountsRequiresStatus() (+36 more)

### Community 18 - "Community 18"
Cohesion: 0.07
Nodes (26): AuthStatus, AuthType, Billing, BillingHistoryEntry, BuildRouteMode, Credential, credentialMetadataAdapterStub, LinkedAccount (+18 more)

### Community 19 - "Community 19"
Cohesion: 0.09
Nodes (48): NormalizeTitle(), NormalizeURL(), sanitizeTitle(), FuzzNormalizeURL(), T, TestNormalizeTitle(), TestNormalizeURL(), unsafeTextRune() (+40 more)

### Community 20 - "Community 20"
Cohesion: 0.07
Nodes (38): AppShell(), documentation, navigation, AnonymousBoundary(), AuthBoundary(), Button, ButtonProps, buttonVariants (+30 more)

### Community 21 - "Community 21"
Cohesion: 0.09
Nodes (19): Context, Time, mediaJobAccountID(), mediaJobFromDomain(), mediaJobToDomain(), ticketToDomain(), Time, MediaAssetRepository (+11 more)

### Community 22 - "Community 22"
Cohesion: 0.11
Nodes (45): NewAccountRepository(), T, TestBuildSuperEntitledDefaultsFalseAndSurvivesUpsert(), TestBuildSuperEntitledNonBuildForcedFalse(), TestListRoutingCandidatesSharesEntitledBuildModels(), T, TestBatchDeleteModelRoutesAllowsRediscovery(), TestBuildPaidCapabilitiesAreSharedAcrossActiveSuperAccounts() (+37 more)

### Community 23 - "Community 23"
Cohesion: 0.07
Nodes (44): TooltipContent, accountLinks(), AccountNameCell(), identityDetails(), providerIcon, providerOrder, AccountQuota(), BuildQuota() (+36 more)

### Community 24 - "Community 24"
Cohesion: 0.11
Nodes (16): DeviceSession, Client, Context, Time, T, TestRedisRuntimeStoreIntegration(), NewConcurrencyLimiter(), NewDeviceSessionStore() (+8 more)

### Community 25 - "Community 25"
Cohesion: 0.08
Nodes (24): QuotaBreakdown, QuotaSource, QuotaWindow, Mutex, cleanVersion(), CurrentVersion(), toQuotaWindowDomain(), firstGRPCWebMessage() (+16 more)

### Community 26 - "Community 26"
Cohesion: 0.10
Nodes (38): boolValue(), firstValue(), numberValue(), optionalBool(), parseBilling(), parseSubscriptionTier(), planValues(), stringValue() (+30 more)

### Community 27 - "Community 27"
Cohesion: 0.10
Nodes (26): Header, ReadCloser, antiBotProviderResponse(), estimateTokens(), CancelFunc, Context, Header, Lease (+18 more)

### Community 28 - "Community 28"
Cohesion: 0.13
Nodes (22): Time, mapError(), toModelDomain(), ensureModelPublicIDNotAlias(), DB, preserveModelRouteAlias(), discoveredRouteDefaults(), findModelRoutesByPublicID() (+14 more)

### Community 29 - "Community 29"
Cohesion: 0.08
Nodes (27): Adapter, Client, Closer, Context, Mutex, Reader, Request, RoundTripper (+19 more)

### Community 30 - "Community 30"
Cohesion: 0.12
Nodes (44): containsToolType(), convertAssistantToolCalls(), convertChatContent(), convertChatMessages(), convertChatRequest(), convertChatToolChoice(), convertChatTools(), RawMessage (+36 more)

### Community 31 - "Community 31"
Cohesion: 0.09
Nodes (35): serverToolKey(), webServerToolName(), absoluteAssetURL(), appendCapturedImageURL(), collectCapturedImageURLs(), extractMarkdownImages(), firstInt(), firstString() (+27 more)

### Community 32 - "Community 32"
Cohesion: 0.17
Nodes (33): NewService(), Service, T, runQuotaRefreshWorkers(), TestBuildChatPermissionDenialDoesNotInvalidateVideoCredential(), testCipher(), TestCreateResponseFallsAcrossProviders(), TestCreateResponsePrefixedModelDoesNotHopProvider() (+25 more)

### Community 33 - "Community 33"
Cohesion: 0.06
Nodes (24): Time, accountCredentialModel, accountModelCapabilityModel, accountModelQuotaBlockModel, accountModelSyncStateModel, accountProviderLinkModel, adminModel, adminSessionModel (+16 more)

### Community 34 - "Community 34"
Cohesion: 0.09
Nodes (22): WebTier, earlierFuture(), effectiveQuotaMode(), AccountRepository, ConcurrencyLimiter, Context, Selector, Group (+14 more)

### Community 35 - "Community 35"
Cohesion: 0.11
Nodes (37): compatibilityBoundaryMessage(), decodeApplyPatchArguments(), dedupeNormalizedTools(), encodeToolOutput(), legacyShellAction(), legacyShellCommand(), normalizeApplyPatchOutputInput(), normalizeFunctionCallOutputInput() (+29 more)

### Community 36 - "Community 36"
Cohesion: 0.06
Nodes (30): IndicatorGeometry, Tabs(), TabsContent(), TabsList, TabsTrigger(), ModelRouteDTO, getSystemInfo(), AttemptButton() (+22 more)

### Community 37 - "Community 37"
Cohesion: 0.11
Nodes (37): chatFileExtension(), chatFileMIMEFromExtension(), chatFileName(), Addr, Config, Context, Header, Lease (+29 more)

### Community 38 - "Community 38"
Cohesion: 0.11
Nodes (20): AccountRepository, Capability, Context, Group, Logger, ModelRepository, Time, invalidInput() (+12 more)

### Community 39 - "Community 39"
Cohesion: 0.12
Nodes (40): responsesToolCompatibility, normalizeResponsesRequest(), T, TestResponsesBuild02101NativeAndUnsupportedToolMatrix(), TestResponsesCodexHistoryItemsAreStructuredOrVisible(), TestResponsesCustomGrammarDowngradesWithoutRejectingRequest(), TestResponsesCustomToolRequestHistoryAndJSONResponse(), TestResponsesCustomToolStreamUsesCustomEvents() (+32 more)

### Community 40 - "Community 40"
Cohesion: 0.12
Nodes (42): buildOpenAIResult(), collectServerTool(), newWebMessagesStream(), webMessagesSearchRequests(), newToolStreamSieve(), T, TestAnthropicWebSearchRequestConvertsForWebProvider(), TestBuildMessagesResultBoundsSearchSourcesAndTitles() (+34 more)

### Community 41 - "Community 41"
Cohesion: 0.11
Nodes (28): cleanupRateShard(), Context, Mutex, Time, maxEntriesPerShard(), NewConcurrencyLimiter(), NewDeviceSessionStore(), NewLockStore() (+20 more)

### Community 42 - "Community 42"
Cohesion: 0.13
Nodes (34): NewSelector(), stickySessionKey(), BenchmarkSelectorCandidatePlanning(), B, Context, Mutex, Selector, T (+26 more)

### Community 43 - "Community 43"
Cohesion: 0.07
Nodes (30): AppProviders(), router, AdminDTO, adminValidator, ApiError, ApiStreamEvent, AuthTokensDTO, authTokensValidator (+22 more)

### Community 44 - "Community 44"
Cohesion: 0.15
Nodes (20): copyHeaders(), decodeSingleJSON(), forceJSONBoolean(), Context, Header, Result, RouterGroup, Service (+12 more)

### Community 45 - "Community 45"
Cohesion: 0.08
Nodes (16): accountReaderStub, testConversationDefinition(), Adapter, Adapter, Adapter, Capability, Adapter, ConversationSurface (+8 more)

### Community 46 - "Community 46"
Cohesion: 0.18
Nodes (9): Success(), Handler, Context, RouterGroup, newAccountResponse(), newBillingResponse(), pagination(), parseIDs() (+1 more)

### Community 47 - "Community 47"
Cohesion: 0.12
Nodes (14): Context, fromEgressDomain(), Context, Database, NewEgressRepository(), toEgressDomain(), Context, consoleEgressRepositoryStub (+6 more)

### Community 48 - "Community 48"
Cohesion: 0.11
Nodes (18): QuotaRecoveryEvent, Context, Mutex, Time, Context, Mutex, Time, NewQuotaRecoveryQueue() (+10 more)

### Community 49 - "Community 49"
Cohesion: 0.11
Nodes (20): CursorResult, ListFilter, Period, Service, SummaryResult, SummaryUsage, decodeAuditCursor(), encodeAuditCursor() (+12 more)

### Community 50 - "Community 50"
Cohesion: 0.13
Nodes (36): T, TestConvertAnthropicClaudeCodeRequestToResponses(), TestConvertAnthropicMessagesIgnoresUnrepresentableTopK(), TestConvertAnthropicMessagesInlineSystemOnly(), TestConvertAnthropicMessagesInlineSystemRole(), TestConvertAnthropicMessagesRejectsUnknownRole(), TestConvertAnthropicMessagesRequestToResponses(), TestConvertAnthropicMessagesValidatesToolRelationships() (+28 more)

### Community 51 - "Community 51"
Cohesion: 0.10
Nodes (24): anthropicErrorJSON(), anthropicMessageID(), anthropicToolUseID(), anthropicUsage(), applyStopSequences(), messagesResponse(), normalizeAnthropicError(), normalizeAnthropicErrorType() (+16 more)

### Community 52 - "Community 52"
Cohesion: 0.14
Nodes (35): ConvertResponseJSONWithOptions(), assertSequentialContentBlocks(), contentBlockStartTypes(), T, TestAnthropicServerToolUseIDDoesNotCollideOnLongCommonSuffix(), TestAnthropicServerToolUseIDNormalizesPrefixedUntrustedID(), TestClientLowercaseWebSearchToolChoiceRemainsFunction(), TestClientWebSearchFunctionNotPromoted() (+27 more)

### Community 53 - "Community 53"
Cohesion: 0.14
Nodes (15): accountEventStream, accountSyncPipeline, buildConversionRequest, webConsoleSyncRequest, CancelFunc, Int64, Mutex, Once (+7 more)

### Community 54 - "Community 54"
Cohesion: 0.13
Nodes (32): cleanGatewayCompactionSummary(), expandGatewayCompactionHistory(), gatewayCompactionContinuation(), gatewayCompactionSummaryMessage(), isDegenerateGatewayCompactionSummary(), neutralizeGatewayCompactionTags(), newGatewayCompactionCodec(), prepareGatewayCompactionSample() (+24 more)

### Community 55 - "Community 55"
Cohesion: 0.12
Nodes (34): ChatMessage, ChatResponseResult, ChatStreamSnapshot, createChatResponse(), createVideo(), CreativeApiError, firstString(), generateImage() (+26 more)

### Community 56 - "Community 56"
Cohesion: 0.10
Nodes (23): RawMessage, isEmptyJSON(), mustJSON(), normalizeResponseFormat(), patchReasoningTextTypes(), responsesToolCompatibility, RawMessage, toolsOfType() (+15 more)

### Community 57 - "Community 57"
Cohesion: 0.09
Nodes (24): Capability, Resolve(), TestParseVideoConcatenatedJSONFixture(), TestParseVideoStreamFixture(), TestParseVideoStreamPreservesUpstreamStatus(), TestParseVideoStreamUsesModelResponseAttachment(), consumeVideoJSON(), consumeVideoSSE() (+16 more)

### Community 58 - "Community 58"
Cohesion: 0.09
Nodes (33): getSettings(), SettingsConfigDTO, updateSettings(), auditFlushDuration, byteSizeBytes(), byteSizeSchema, ByteSizeUnit, ByteSizeValue (+25 more)

### Community 59 - "Community 59"
Cohesion: 0.15
Nodes (16): fallbackScopes(), Context, Lease, EgressRepository, Group, Mutex, Request, Time (+8 more)

### Community 60 - "Community 60"
Cohesion: 0.11
Nodes (28): Context, RouterGroup, Service, Time, NewHandler(), newSettingsResponse(), T, TestSettingsDTOExcludesBrowserIdentityFields() (+20 more)

### Community 61 - "Community 61"
Cohesion: 0.17
Nodes (9): BatchProgressObserver, ImportedAccountObserver, ImportResult, WebAccountScriptOptions, WebConsoleSyncStrategy, offsetBatchProgress(), Service, Context (+1 more)

### Community 62 - "Community 62"
Cohesion: 0.16
Nodes (14): Capability, Time, cleanupExpiredBillingReservations(), decrementReservedUsage(), expiredBillingReservationCount(), Context, Database, DB (+6 more)

### Community 63 - "Community 63"
Cohesion: 0.14
Nodes (15): Context, Service, Header, Int64, Logger, MediaJobRepository, Mutex, ResponseRepository (+7 more)

### Community 64 - "Community 64"
Cohesion: 0.24
Nodes (31): Request, encryptFallbackTestJWT(), Adapter, T, newFallbackTestAdapter(), TestFallbackMarkerFailureStillReturnsSuccess(), TestForwardResponseForceBuildDoesNotFallbackOn403(), TestForwardResponseMarkedAccountStillNeedsCurrentBuild403() (+23 more)

### Community 65 - "Community 65"
Cohesion: 0.15
Nodes (20): chatToolCalls(), estimateToolCallTokens(), Builder, Writer, newWebID(), newWebStopFilter(), nullableWebString(), shouldEmitWebMessagesSearch() (+12 more)

### Community 66 - "Community 66"
Cohesion: 0.12
Nodes (26): CleanupStatus, ExportResult, IssueSummary, ListFilter, ProviderSummary, QuotaStatus, QuotaType, QuotaView (+18 more)

### Community 67 - "Community 67"
Cohesion: 0.13
Nodes (26): ErrorFrame, appendErrorFrames(), auditAccountID(), errorFrames(), errorUpstreamURL(), Closer, Header, ReadCloser (+18 more)

### Community 68 - "Community 68"
Cohesion: 0.10
Nodes (21): ClientKeyRepository, ConcurrencyLimiter, Context, Int64, RateLimiter, Time, invalidInput(), mapRepositoryError() (+13 more)

### Community 69 - "Community 69"
Cohesion: 0.10
Nodes (25): consoleProviderConfig(), AccountRepository, Adapter, Application, BatchConfig, Closer, Config, Context (+17 more)

### Community 70 - "Community 70"
Cohesion: 0.14
Nodes (18): startupReport, startupState, Timer, resetTimer(), AccountRepository, Application, Context, ModelRepository (+10 more)

### Community 71 - "Community 71"
Cohesion: 0.14
Nodes (18): cleanupThresholdBytes(), Context, Int64, Service, MediaAssetRepository, MediaJobRepository, MediaUploadTicketRepository, ReadCloser (+10 more)

### Community 72 - "Community 72"
Cohesion: 0.21
Nodes (27): T, TestAdminDeleteVideoJobsRemovesTerminalJobAssetAndTicket(), TestCleanupAllProtectedTerminatesWithoutDelete(), TestCleanupCapsExpiredTicketPruningPerInvocation(), TestCleanupDeletesOldestAssetsAtThreshold(), TestCleanupPagesPastProtectedAssetsAndDeletesLaterUnprotected(), TestCleanupPreservesMetadataWhenLocalObjectIsMissing(), TestCleanupPrunesExpiredUploadTickets() (+19 more)

### Community 73 - "Community 73"
Cohesion: 0.14
Nodes (19): errVideoTooLarge(), Context, Service, ReadCloser, Reader, Time, hashUploadToken(), isBodyTooLargeError() (+11 more)

### Community 74 - "Community 74"
Cohesion: 0.10
Nodes (26): dialContext(), Client, Conn, Context, newBuildClient(), Reader, T, isClosedNetworkError() (+18 more)

### Community 75 - "Community 75"
Cohesion: 0.17
Nodes (17): NewMediaPostProcessingError(), applyAppHeaders(), buildHeaders(), Header, Lease, decodeDirectFileUploadResponse(), decodeImageBlob(), Config (+9 more)

### Community 76 - "Community 76"
Cohesion: 0.14
Nodes (30): extractMetadata(), NewHandler(), T, TestCopyHeadersFiltersHopByHopAndUpstreamCookies(), TestCopyJSONForwardsBodyBeyondMetadataInspectionLimit(), TestCopyMediaAllowsExactLimit(), TestCopyMediaRejectsUnknownLengthOverflowWithoutWritingPastLimit(), TestCopyStreamRequiresProtocolTerminalEvent() (+22 more)

### Community 77 - "Community 77"
Cohesion: 0.13
Nodes (24): AlertDialogAction, AlertDialogCancel, AlertDialogContent, AlertDialogDescription, AlertDialogFooter(), AlertDialogHeader(), AlertDialogOverlay, AlertDialogTitle (+16 more)

### Community 78 - "Community 78"
Cohesion: 0.16
Nodes (20): coordinatedSessionRepository, NewService(), AdminSessionRepository, Mutex, T, newCoordinatedSessionRepository(), TestChangePasswordRevokesAllSessions(), TestConcurrentRefreshAllowsExactlyOneRotation() (+12 more)

### Community 79 - "Community 79"
Cohesion: 0.15
Nodes (27): AccessLog(), HandlerFunc, Logger, MaxBodyBytes(), RequestID(), SecurityHeaders(), T, TestMaxBodyBytesLimitsAllRequestBodies() (+19 more)

### Community 80 - "Community 80"
Cohesion: 0.13
Nodes (16): Service, Tokens, AdminRepository, AdminSessionRepository, Context, RateLimiter, Time, newAuditEventID() (+8 more)

### Community 81 - "Community 81"
Cohesion: 0.11
Nodes (16): Context, ReadCloser, Reader, T, Time, TestPersistRemoteVideoRetriesSameResultWithoutRegeneration(), TestRecoverVideoJobsRecordsDetachedAccountSnapshot(), TestRecoverVideoJobsRecordsFailedAuditWithEgress() (+8 more)

### Community 82 - "Community 82"
Cohesion: 0.11
Nodes (27): RawMessage, hasJSONValue(), newModelListItems(), officialVideoErrorCode(), parseOptionalVideoInteger(), parseVideoDuration(), validImageAspectRatio(), validImageEditSize() (+19 more)

### Community 83 - "Community 83"
Cohesion: 0.13
Nodes (17): Context, RouterGroup, Service, Time, NewHandler(), newModelResponse(), pagination(), parseIDs() (+9 more)

### Community 84 - "Community 84"
Cohesion: 0.07
Nodes (29): eslint, @eslint/js, eslint-plugin-react-hooks, eslint-plugin-react-refresh, devDependencies, eslint, @eslint/js, eslint-plugin-react-hooks (+21 more)

### Community 85 - "Community 85"
Cohesion: 0.14
Nodes (21): auditAttemptResponse, auditDetailResponse, auditErrorFrameResponse, auditResponse, Handler, pricingResponse, summaryRangeResponse, summaryResponse (+13 more)

### Community 86 - "Community 86"
Cohesion: 0.14
Nodes (22): validMediaStatus(), cloneSnapshot(), compareSemanticVersion(), Client, Context, Group, RWMutex, Time (+14 more)

### Community 87 - "Community 87"
Cohesion: 0.13
Nodes (14): rewriteLegacyShellAction(), decodeCustomToolInput(), encodeCustomToolArguments(), consumeCompatibleSSE(), customToolStreamPayload(), decodeToolSearchArguments(), responsesToolCompatibility, ReadCloser (+6 more)

### Community 88 - "Community 88"
Cohesion: 0.13
Nodes (9): newAnthropicStreamStopFilter(), consumeSSE(), Builder, streamConverter, Reader, Writer, newStreamConverter(), anthropicStreamStopFilter (+1 more)

### Community 89 - "Community 89"
Cohesion: 0.17
Nodes (23): Database, Int32, MediaAssetRepository, MediaUploadTicketRepository, Service, T, newUploadTestService(), openUploadTestDeps() (+15 more)

### Community 90 - "Community 90"
Cohesion: 0.17
Nodes (25): PricingResult, tokenPrice, tokenPriceRule, buildOfficialTokenPrices(), estimateJSONTokens(), EstimateOfficialCost(), EstimateOfficialImageCost(), EstimateOfficialImageEditCost() (+17 more)

### Community 91 - "Community 91"
Cohesion: 0.15
Nodes (18): Context, EgressRepository, NewService(), NormalizeProxyURL(), publicNode(), SanitizeCloudflareCookies(), T, TestBuildNodeAlwaysUsesProviderUserAgent() (+10 more)

### Community 92 - "Community 92"
Cohesion: 0.11
Nodes (22): AccountsPage, ApiDocsPage, AppShell, ClientKeysPage, CreativeConsolePage, DashboardPage, DeferredAccountsPage(), DeferredApiDocsPage() (+14 more)

### Community 93 - "Community 93"
Cohesion: 0.07
Nodes (26): compilerOptions, allowJs, allowSyntheticDefaultImports, esModuleInterop, forceConsistentCasingInFileNames, incremental, isolatedModules, jsx (+18 more)

### Community 94 - "Community 94"
Cohesion: 0.11
Nodes (7): billingStub, countingAdapter, identityAccountReaderStub, modelStub, quotaStub, Context, Mutex

### Community 95 - "Community 95"
Cohesion: 0.17
Nodes (18): applyAuditQuery(), billInsertedAudit(), createAuditAndBill(), Context, Database, DB, Summary, Time (+10 more)

### Community 96 - "Community 96"
Cohesion: 0.14
Nodes (17): claimString(), conversionStatus(), decodeBuildClaims(), firstValue(), Context, Time, Values, Adapter (+9 more)

### Community 97 - "Community 97"
Cohesion: 0.18
Nodes (16): Context, RouterGroup, Service, Time, NewHandler(), newKeyResponse(), pagination(), parseIDs() (+8 more)

### Community 98 - "Community 98"
Cohesion: 0.08
Nodes (26): class-variance-authority, clsx, date-fns, dependencies, class-variance-authority, clsx, date-fns, lucide-react (+18 more)

### Community 99 - "Community 99"
Cohesion: 0.15
Nodes (17): ReadCloser, Time, isRetryable(), isRetryableResponse(), isUpstreamStreamFailure(), parseFreeQuotaExhaustion(), parseRetryAfter(), rateLimitTeamFingerprint() (+9 more)

### Community 100 - "Community 100"
Cohesion: 0.15
Nodes (18): TestSyncAccountIdentityUsesWebBrowserIdentity(), T, grpcWebTrailerFrame(), TestAcceptTermsRequiresBothUpstreamSteps(), TestWebAccountBirthDateMapsOnlyAlreadySetResponse(), TestWebAccountSettingsMapUnauthorized(), TestWebAccountSettingsMatchCapturedProtocol(), TestWebAccountSettingsRejectsBodyTrailerFailure() (+10 more)

### Community 101 - "Community 101"
Cohesion: 0.14
Nodes (15): ClientKeyRepository, Context, T, Time, TestAuthenticateCachesUnlimitedKeyAndInvalidatesOnDisable(), TestAuthenticateDistinguishesRuntimeStoreFailures(), TestBillingLimitUsesAtomicReservations(), testCipher() (+7 more)

### Community 102 - "Community 102"
Cohesion: 0.22
Nodes (7): decodeVideoInput(), encodeVideoInput(), Context, Service, ReadCloser, waitVideoOutputRetry(), VideoInput

### Community 103 - "Community 103"
Cohesion: 0.17
Nodes (20): mediaPageQuery(), applyStableSort(), DB, stableSortSpec(), Time, IsValidSort(), sortSpec, AccountListFilter (+12 more)

### Community 104 - "Community 104"
Cohesion: 0.13
Nodes (22): accountBatchResponse, accountCleanupRequest, accountImportResponse, accountModelSynchronizer, accountResponse, accountSyncProgressor, accountTaskProgressResponse, accountTokenRefreshResponse (+14 more)

### Community 105 - "Community 105"
Cohesion: 0.21
Nodes (9): Admin, Session, Time, Context, Time, toAdminDomain(), toSessionDomain(), AdminRepository (+1 more)

### Community 106 - "Community 106"
Cohesion: 0.15
Nodes (14): adminResponse, changePasswordRequest, Handler, loginRequest, refreshRequest, tokenResponse, Context, Request (+6 more)

### Community 107 - "Community 107"
Cohesion: 0.19
Nodes (18): blockingAuditRepository, flakyAuditRepository, panicAuditRepository, summaryAuditRepository, NewService(), AuditRepository, Context, Int32 (+10 more)

### Community 108 - "Community 108"
Cohesion: 0.19
Nodes (5): Context, ReadCloser, Time, countingObjectStore, failingCommitStore

### Community 109 - "Community 109"
Cohesion: 0.12
Nodes (22): Providers(), DisplayUpstreamModel(), ExternalPublicID(), IsCanonicalPublicID(), NormalizePublicID(), NormalizeUpstreamModel(), PublicIDCandidates(), T (+14 more)

### Community 110 - "Community 110"
Cohesion: 0.22
Nodes (22): NewManager(), T, managerHasClientForNode(), TestAcquireCredentialRendersResinAccountAndOverridesNodeCookie(), TestAcquireIfConfiguredDoesNotChangeBuildDirectTransport(), TestBrowserRequestLeavesHeaderOrderingToTLSProfile(), TestBuildForbiddenDoesNotPoisonEgressNode(), TestClientCacheEvictsIdleEntriesAndEnforcesCapacity() (+14 more)

### Community 111 - "Community 111"
Cohesion: 0.16
Nodes (19): buildGatewayCompactionResponse(), extractCompactionSummary(), gatewayCompactionErrorIsTransient(), gatewayCompactionEventErrorIsTransient(), gatewayCompactionFailureProviderResponse(), gatewayCompactionHTTPFailure(), gatewayCompactionItemText(), gatewayCompactionStatusTransient() (+11 more)

### Community 112 - "Community 112"
Cohesion: 0.10
Nodes (8): SwaggerChatRequest, SwaggerImageEditRequest, SwaggerImageGenerationRequest, SwaggerImageReference, SwaggerMessage, SwaggerMessagesRequest, SwaggerResponsesRequest, SwaggerVideoGenerationRequest

### Community 113 - "Community 113"
Cohesion: 0.15
Nodes (18): accountProgressSynchronizerStub, accountSynchronizer, accountSynchronizerStub, Service, NewHandler(), Context, Result, T (+10 more)

### Community 114 - "Community 114"
Cohesion: 0.19
Nodes (4): View, Service, Context, mapRepositoryError()

### Community 115 - "Community 115"
Cohesion: 0.17
Nodes (13): Context, Logger, QuotaRecoveryQueue, RWMutex, Time, NewService(), Logger, T (+5 more)

### Community 116 - "Community 116"
Cohesion: 0.19
Nodes (14): Context, DB, AccountRepository, linkWebToBuildIfUnambiguous(), linkWebToConsole(), matchingConsoleSourceKey(), matchingWebSourceKey(), reconcileWebBuildByUserID() (+6 more)

### Community 117 - "Community 117"
Cohesion: 0.22
Nodes (15): Context, Mutex, Time, New(), T, TestCacheCoalescesConcurrentLoads(), TestCacheDeleteInvalidatesStoredValue(), TestCacheExpiresAndEvictsOldestEntry() (+7 more)

### Community 118 - "Community 118"
Cohesion: 0.15
Nodes (9): auditCursorPayload, Context, Engine, RouterGroup, parsePagination(), deleteImagesRequest, deleteVideosRequest, Handler (+1 more)

### Community 119 - "Community 119"
Cohesion: 0.22
Nodes (19): Service, T, Time, newCredentialRefreshTestService(), TestBatchRefreshTokensRefreshesOnlySelectedEligibleAccounts(), TestCredentialRefreshDueQueryStaysBoundedForLargePool(), TestCredentialRefreshFailureDistinguishesTransientAndPermanent(), TestCredentialRefreshSchedulerRefreshesOnlyDueAccounts() (+11 more)

### Community 120 - "Community 120"
Cohesion: 0.17
Nodes (16): fhttpResponseAsHTTP(), fromFHTTPResponse(), Conn, Context, Lease, Header, Request, newBrowserClient() (+8 more)

### Community 121 - "Community 121"
Cohesion: 0.30
Nodes (20): grokSessionID(), NewAdapter(), T, TestCredentialMetadataMarksOnlyNumericBotFlagOne(), TestForwardResponseDecodesExplicitGzipResponse(), TestForwardResponseDowngradesServerToolSearchBeforeUpstream(), TestForwardResponseInjectsPromptCacheKeyAfterChatConversion(), TestForwardResponseMapsClaudeCodeWebSearchEndToEnd() (+12 more)

### Community 122 - "Community 122"
Cohesion: 0.21
Nodes (18): applyParsedToolCalls(), appendParsedToolCall(), RawMessage, injectToolPrompt(), normalizeToolArguments(), parseFunctionTool(), parseJSONToolCalls(), parseToolCalls() (+10 more)

### Community 123 - "Community 123"
Cohesion: 0.21
Nodes (17): CLIProxyExportResult, buildCLIProxyAuthDocument(), BuildCLIProxyAuthExport(), cliproxyAuthFilename(), DisambiguateCLIProxyFilenames(), Time, marshalCLIProxyAuthDocument(), MarshalCLIProxyAuthFile() (+9 more)

### Community 124 - "Community 124"
Cohesion: 0.28
Nodes (16): alignedRange(), buildActivitySeries(), buildSeries(), Context, ModelUsage, Time, Usage, parsePeriod() (+8 more)

### Community 125 - "Community 125"
Cohesion: 0.15
Nodes (16): Context, dashboardRepositoryStub, Time, ModelUsage, Time, Context, dashboardRepositoryStub, Time (+8 more)

### Community 126 - "Community 126"
Cohesion: 0.16
Nodes (10): BuildRouteMode, Adapter, Context, isHTTPForbidden(), isXAIInferenceFallbackCapable(), normalizeBuildAPIPath(), normalizedBuildRouteMode(), shouldProbeXAIInferenceFallback() (+2 more)

### Community 127 - "Community 127"
Cohesion: 0.21
Nodes (9): RoutingCandidate, accountConcurrencyKey(), candidateScoreBetter(), concurrencySnapshotKey(), Context, Selector, Time, candidatePlan (+1 more)

### Community 128 - "Community 128"
Cohesion: 0.21
Nodes (15): containsAny(), extractUpstreamErrorMetadata(), firstNonEmptyFailure(), firstStringValue(), isAccountScopedForbidden(), isFreeQuotaExhaustion(), isModelQuotaExhaustion(), isPaidQuotaExhaustion() (+7 more)

### Community 129 - "Community 129"
Cohesion: 0.14
Nodes (6): Context, Int64, Mutex, buildCapabilityNormalizerAdapter, modelCapabilityAdapter, DeviceAuthorization

### Community 130 - "Community 130"
Cohesion: 0.15
Nodes (14): Adapter, Context, Fetch(), firstNonEmpty(), Context, Parse(), Context, Adapter (+6 more)

### Community 131 - "Community 131"
Cohesion: 0.21
Nodes (11): Context, RouterGroup, Service, Time, NewHandler(), newNodeResponse(), pathID(), Handler (+3 more)

### Community 132 - "Community 132"
Cohesion: 0.19
Nodes (11): buildConversionAdapter, conversionBatchRepository, AccountRepository, Context, Int64, T, Uint64, TestConvertAllWebAccountsToBuildProcessesMoreThanLegacyLimitInBatches() (+3 more)

### Community 133 - "Community 133"
Cohesion: 0.16
Nodes (11): AEAD, Config, Context, Database, Time, NewRuntimeSettingsRepository(), T, TestRuntimeSettingsRepositoryRoundTrip() (+3 more)

### Community 134 - "Community 134"
Cohesion: 0.12
Nodes (10): Logger, Mutex, Once, ReadCloser, Time, newGenerationTiming(), T, TestGenerationTimingLogsOnlyPhaseMetadata() (+2 more)

### Community 135 - "Community 135"
Cohesion: 0.22
Nodes (7): Time, Context, Database, Time, ResponseOwnership, WebResponseState, ResponseRepository

### Community 136 - "Community 136"
Cohesion: 0.22
Nodes (12): cloneReplayItems(), Context, Mutex, Time, NewReasoningReplayStore(), reasoningReplayMapKey(), T, TestReasoningReplayStoreEvictsOldest() (+4 more)

### Community 137 - "Community 137"
Cohesion: 0.21
Nodes (11): copyJSON(), copyMedia(), copyStream(), Reader, ResponseWriter, Usage, Writer, setResponseWriteDeadline() (+3 more)

### Community 138 - "Community 138"
Cohesion: 0.15
Nodes (6): DeviceStartResult, T, TestPreserveActiveQuotaWindowsUntilReset(), credentialRefreshBackoff(), Time, preserveActiveQuotaWindows()

### Community 139 - "Community 139"
Cohesion: 0.23
Nodes (11): Attempt, AttemptSource, EgressMode, Operation, Record, Summary, UsageSource, Time (+3 more)

### Community 140 - "Community 140"
Cohesion: 0.28
Nodes (16): defaultConfig(), Load(), T, TestDefaultConsoleProviderConfig(), TestDefaultGrokBuildClientVersionMatchesLocalBaseline(), TestDurationStringUsesCompactYAMLForm(), TestEffectivePublicAPIBaseURLPriority(), TestLoadAcceptsLegacyConsoleUserAgent() (+8 more)

### Community 141 - "Community 141"
Cohesion: 0.22
Nodes (10): Client, Context, Time, Values, newOAuthClient(), parseOAuthRetryAfter(), T, TestOAuthRefreshClassifiesPermanentAndTransientFailures() (+2 more)

### Community 142 - "Community 142"
Cohesion: 0.17
Nodes (10): Context, RouterGroup, Service, NewHandler(), Request, T, TestHandlerReturnsAndChecksVersion(), TestHandlerReturnsOnlyPublicFrontendConfig() (+2 more)

### Community 143 - "Community 143"
Cohesion: 0.21
Nodes (6): BuildConversionResult, BuildConversionStrategy, UpdateInput, BuildRouteMode, invalidInput(), normalizeIDs()

### Community 144 - "Community 144"
Cohesion: 0.18
Nodes (8): Mutex, RWMutex, Time, newAuthKeyCache(), newTouchTracker(), authKeyCache, cachedAuthKey, touchTracker

### Community 145 - "Community 145"
Cohesion: 0.27
Nodes (15): Adapter, T, newTestBuildVideoAdapter(), TestBuildVideoCreatePayloadImageOnlyEmptyPrompt(), TestBuildVideoCreatePayloadNoImageAndSingleR2URL(), TestBuildVideoDefinitionDeclaresVideoCapability(), TestGenerateVideoFailedStatusAndDownloadTrustedURL(), TestGenerateVideoMapsUnauthorizedAndUpstreamErrors() (+7 more)

### Community 146 - "Community 146"
Cohesion: 0.12
Nodes (15): aliases, components, hooks, lib, ui, utils, iconLibrary, rsc (+7 more)

### Community 147 - "Community 147"
Cohesion: 0.26
Nodes (13): Reader, Time, pendingWebAccountScriptOptions(), randomWebBirthDate(), T, TestPendingWebAccountScriptOptionsSkipsRecordedSteps(), TestRandomWebBirthDateRejectsUnavailableRandomSource(), TestRandomWebBirthDateUsesInclusiveAgeRange() (+5 more)

### Community 148 - "Community 148"
Cohesion: 0.28
Nodes (4): Context, ReadCloser, videoExtension(), LocalStore

### Community 149 - "Community 149"
Cohesion: 0.15
Nodes (9): assertSQLiteUniqueIndexes(), Database, T, TestInitializeSchemaUpgradesProviderChecksForConsole(), legacyEgressNodeModel, legacyModelRouteModel, legacyProviderAccountModel, legacyRequestAuditModel (+1 more)

### Community 150 - "Community 150"
Cohesion: 0.24
Nodes (10): Builder, responsesToolCompatibility, hashedToolAlias(), newResponsesToolCompatibility(), shortToolHash(), truncateToolAlias(), responsesRequestError, responsesStreamCall (+2 more)

### Community 151 - "Community 151"
Cohesion: 0.27
Nodes (13): AdminAuth(), bearerToken(), ClientAuth(), clientErrorCode(), clientErrorMessage(), clientErrorStatus(), Context, HandlerFunc (+5 more)

### Community 152 - "Community 152"
Cohesion: 0.38
Nodes (13): createWebAccountForScriptTest(), T, TestEnableWebNSFWAlwaysSetsBirthDateFirst(), TestEnableWebNSFWContinuesWhenBirthDateIsAlreadySet(), TestEnableWebNSFWDoesNotMarkUpstreamFailure(), TestEnableWebNSFWPersistsMarkerAfterClientCancellation(), TestRunAllWebAccountScriptsOnlyProcessesWebAccounts(), TestRunWebAccountScriptsRejectsEmptyPlan() (+5 more)

### Community 153 - "Community 153"
Cohesion: 0.24
Nodes (10): Capability, Context, Service, Result, firstError(), IsMediaPostProcessingError(), ImageEditInput, imageExecution (+2 more)

### Community 154 - "Community 154"
Cohesion: 0.29
Nodes (12): accountFromContext(), Context, RWMutex, recordSelection(), TraceFromContext(), WithAccount(), WithAccountIdentity(), WithCredential() (+4 more)

### Community 155 - "Community 155"
Cohesion: 0.32
Nodes (13): assertAccountFilterCount(), assertAuditFilterResult(), assertAuditSearchResult(), assertClientKeySearchCount(), assertModelSearchCount(), AccountRepository, AuditRepository, ClientKeyRepository (+5 more)

### Community 156 - "Community 156"
Cohesion: 0.41
Nodes (4): Context, Database, sqliteConstraintDefinition(), consoleConstraint

### Community 157 - "Community 157"
Cohesion: 0.29
Nodes (12): parseImportedCredentials(), T, TestMarshalCredentialsUsesImportDocument(), TestNormalizeResponsesRequest(), TestNormalizeResponsesRequestAddsEmptySummaryToEncryptedReasoning(), TestNormalizeResponsesRequestDoesNotInventPromptCacheKey(), TestNormalizeResponsesRequestFlattensJSONSchema(), TestNormalizeResponsesRequestPreservesExplicitPromptCacheKey() (+4 more)

### Community 158 - "Community 158"
Cohesion: 0.33
Nodes (12): Adapter, Context, Request, T, jsonHTTPResponse(), newReasoningRecoveryTestAdapter(), TestRecoverReasoningDecodeFailureDoesNotRetryOtherBadRequests(), TestRecoverReasoningDecodeFailurePreservesOriginalWhenRetryFails() (+4 more)

### Community 159 - "Community 159"
Cohesion: 0.36
Nodes (12): NewService(), T, TestGetAlignsDailyBucketsToCalendarDays(), TestGetBuildsStableBucketsAndEnrichedData(), TestGetCachesRepeatedAggregate(), TestGetRejectsInvalidTimezone(), TestGetRejectsUnknownPeriod(), TestGetUsesCalendarBoundariesAcrossDST() (+4 more)

### Community 160 - "Community 160"
Cohesion: 0.31
Nodes (12): assertMediaAssetIDs(), assertMediaJobIDs(), T, Time, TestAccountDeleteDetachesTerminalMediaJobsAndRejectsActiveJobs(), testMediaAsset(), TestMediaAssetRepositoryListMediaAssetsPaginatesAndCounts(), testMediaJob() (+4 more)

### Community 161 - "Community 161"
Cohesion: 0.21
Nodes (4): Context, egressRepositoryStub, imageAssetStoreRetryStub, imageAssetStoreStub

### Community 162 - "Community 162"
Cohesion: 0.30
Nodes (5): CredentialStartupReport, Service, Context, Timer, resetCredentialRefreshTimer()

### Community 163 - "Community 163"
Cohesion: 0.41
Nodes (3): Service, Context, WebAccountSettingsAdapter

### Community 164 - "Community 164"
Cohesion: 0.24
Nodes (9): defaultConfigPath(), parseOptions(), Run(), T, TestDefaultConfigPathFindsRepositoryRoot(), TestParseOptionsSupportsContainerListenOverride(), Logger, NewLogger() (+1 more)

### Community 165 - "Community 165"
Cohesion: 0.29
Nodes (9): configureDatabase(), Config, Context, DB, Database, gormConfig(), OpenPostgres(), T (+1 more)

### Community 166 - "Community 166"
Cohesion: 0.32
Nodes (9): chatResponse(), chatUsage(), parseResponse(), extractMessageAnnotations(), parsedResponse, responseContent, responseEnvelope, responseItem (+1 more)

### Community 167 - "Community 167"
Cohesion: 0.21
Nodes (7): extractCompletedPayloadFromSSE(), Context, ReadCloser, ReasoningReplay, patchCompletedOutput(), Buffer, replayCaptureBody

### Community 168 - "Community 168"
Cohesion: 0.33
Nodes (11): Time, Usage, toUsageDTO(), activityDTO, modelUsageDTO, providerUsageDTO, rangeDTO, resourcesDTO (+3 more)

### Community 169 - "Community 169"
Cohesion: 0.24
Nodes (9): Context, RouterGroup, Service, NewHandler(), T, TestHandlerRejectsInvalidPeriod(), TestHandlerRejectsInvalidTimezone(), TestHandlerReturnsDashboardContract() (+1 more)

### Community 170 - "Community 170"
Cohesion: 0.31
Nodes (4): webAccountSettingsAdapterStub, Context, Mutex, Time

### Community 171 - "Community 171"
Cohesion: 0.18
Nodes (11): AuditConfig, BatchConfig, ClientKeyDefaultsConfig, FrontendConfig, MediaConfig, RoutingConfig, ServerConfig, ProviderBuildConfig (+3 more)

### Community 172 - "Community 172"
Cohesion: 0.36
Nodes (8): Request, T, TestStickyLeaseDoesNotRetryAfterRequestWasWritten(), TestStickyLeaseDoesNotRetryUnsafeUpstreamOutcomes(), TestStickyLeaseKeepsNonReplayableResinResponseReadable(), TestStickyLeaseRetriesExplicitResinConnectResponse(), TestStickyLeaseRetriesSafeProxyConnectFailure(), scriptedRequestClient

### Community 173 - "Community 173"
Cohesion: 0.36
Nodes (10): T, TestAuditRepositoryAllowsRepeatedExternalRequestIDs(), TestAuditRepositoryAtomicallyRecordsClientBillingUsage(), TestAuditRepositoryBatchAndCursor(), TestAuditRepositoryNormalizesUntrustedUsage(), TestAuditRepositoryRoundTripsFailureAttempts(), TestAuditRepositorySettlesBillingReservationIdempotently(), TestAuditRepositorySummaryAppliesRangeAndGroupsPricingTier() (+2 more)

### Community 175 - "Community 175"
Cohesion: 0.35
Nodes (6): cloneItems(), Context, Time, hashKeyPart(), NewReasoningReplayStore(), ReasoningReplayStore

### Community 176 - "Community 176"
Cohesion: 0.18
Nodes (10): name, packageManager, private, scripts, build, dev, lint, preview (+2 more)

### Community 177 - "Community 177"
Cohesion: 0.18
Nodes (10): compilerOptions, allowImportingTsExtensions, composite, module, moduleResolution, noEmit, skipLibCheck, tsBuildInfoFile (+2 more)

### Community 178 - "Community 178"
Cohesion: 0.36
Nodes (8): digestUUID(), hexDigest(), resolveBuildSessionIdentity(), T, TestResolveBuildSessionIdentityIsStableAndTenantIsolated(), TestResolveBuildSessionIdentityPrefersExplicitKey(), TestResolveBuildSessionIdentitySeparatesAffinityFromUpstreamSession(), buildSessionIdentity

### Community 179 - "Community 179"
Cohesion: 0.44
Nodes (9): T, Time, testDashboardBoundaries(), TestDashboardRepositoryFillsTopModelsFromEnabledRoutes(), TestDashboardRepositoryRanksTopModels(), TestDashboardRepositorySnapshot(), testDashboardWindow(), timePointer() (+1 more)

### Community 180 - "Community 180"
Cohesion: 0.47
Nodes (9): assertMediaJobSQLContainsBuild(), assertMediaJobSQLLacksBuild(), Context, Database, T, mediaJobsTableSQL(), recreateLegacyMediaJobsTable(), TestInitializeSchemaUpgradesMediaJobChecksForBuild() (+1 more)

### Community 181 - "Community 181"
Cohesion: 0.27
Nodes (6): Bool, Context, Int32, fallbackMarkerStub, trackingFallbackMarker, uploadIssuerStub

### Community 182 - "Community 182"
Cohesion: 0.36
Nodes (9): T, TestInferWebTierFromUpstreamQuota(), TestParseCapturedWeeklyCreditsResponse(), TestParseCoarseWeeklyCreditsResponseRemainsUnavailable(), TestParseUnusedPreciseWeeklyCreditsResponse(), TestResolveWebTierUsesFreshWebQuotaOverStoredTier(), TestSyncQuotaCorrectsStoredSuperFromFreshWebQuota(), TestSyncQuotaFetchesWeeklyOnlyAfterPaidTierIsConfirmed() (+1 more)

### Community 183 - "Community 183"
Cohesion: 0.22
Nodes (6): HandlerFunc, Mutex, NewConcurrencyGate(), T, TestConcurrencyGateHotReloadsAndReleasesCapacity(), ConcurrencyGate

### Community 184 - "Community 184"
Cohesion: 0.36
Nodes (6): consoleIdentityAdapterStub, Context, T, TestSyncAccountIdentityDoesNotRepeatWhenEmailIsKnown(), TestSyncAccountIdentityLinksUniqueBuildWithoutSharingState(), TestSyncAccountIdentityUnauthorizedInvalidatesCurrentProviderAccount()

### Community 185 - "Community 185"
Cohesion: 0.31
Nodes (5): webAccountScriptActionsRequest, webAccountScriptsRequest, Handler, Context, uniqueWebAccountScriptIDs()

### Community 186 - "Community 186"
Cohesion: 0.39
Nodes (7): dashboardBucketExpression(), Context, Database, Time, NewDashboardRepository(), validateDashboardBoundaries(), DashboardRepository

### Community 187 - "Community 187"
Cohesion: 0.36
Nodes (7): asciiAlphaNumeric(), IsInternalHost(), T, TestValidateAllowsPublicHTTPSAndTrustedInternalSigner(), TestValidateRejectsUnsafeOrMalformedSigner(), Validate(), validServiceLabel()

### Community 188 - "Community 188"
Cohesion: 0.36
Nodes (7): extractPromptCacheSeed(), Header, normalizePromptCacheSeed(), promptCacheSeedFromUserID(), T, TestExtractPromptCacheSeedRejectsOversizedValues(), TestExtractPromptCacheSeedSupportsClaudeCodeForms()

### Community 189 - "Community 189"
Cohesion: 0.32
Nodes (4): blockingWebAccountSettingsAdapter, Context, Int32, Time

### Community 190 - "Community 190"
Cohesion: 0.32
Nodes (3): quotaCountingAdapter, Context, Int64

### Community 191 - "Community 191"
Cohesion: 0.46
Nodes (7): AccountRepository, Service, T, openAccountService(), TestAccountViewsIncludeBuildBotFlagMetadata(), TestBuildBotFlagSummaryUsesShortLivedCache(), TestUpdateBuildSuperEntitledBuildOnly()

### Community 192 - "Community 192"
Cohesion: 0.43
Nodes (7): T, TestBillingIsExhaustedForOnDemandCredits(), TestBillingIsPaidMatchesSQLSignals(), TestBillingPeriodEndMatchesExhaustedLimit(), TestBuildRouteModeValidation(), TestIsBuildSuper(), TestRoutingCandidateIsKnownFreeBuild()

### Community 193 - "Community 193"
Cohesion: 0.46
Nodes (7): createLinkedAccountTestCredential(), AccountRepository, Context, T, TestInitializeSchemaBackfillsStableWebEgressIdentity(), TestReconcileProviderLinksUsesOnlyHighConfidenceIdentity(), TestReconcileWebConsoleUsesUniqueUserIDAcrossDifferentSSOTokens()

### Community 194 - "Community 194"
Cohesion: 0.25
Nodes (5): ReadCloser, Request, RoundTripper, egressResponseBody, egressTransport

### Community 195 - "Community 195"
Cohesion: 0.33
Nodes (3): selectedQuotaAdapter, Context, Int64

### Community 196 - "Community 196"
Cohesion: 0.43
Nodes (5): cloneRequestBody(), Lease, Request, retryableResinResponse(), safeProxyConnectionFailure()

### Community 198 - "Community 198"
Cohesion: 0.38
Nodes (6): Error(), Context, stringValue(), errorBody, errorEnvelope, successEnvelope

### Community 199 - "Community 199"
Cohesion: 0.40
Nodes (4): T, TestEgressRepositorySortsInDatabase(), TestInitializeSchemaRemovesAndRejectsLegacyAllEgressNodes(), legacyAllEgressNode

### Community 200 - "Community 200"
Cohesion: 0.67
Nodes (5): assertForbiddenFieldsAbsent(), T, TestAppHeadersMatchStableBrowserFetchSignals(), TestCloudflareCookieWhitelistDropsBrowserIdentityFields(), TestWebHeadersOnlyUseSSOAndCloudflareCookies()

### Community 201 - "Community 201"
Cohesion: 0.47
Nodes (4): responseUpstreamURL(), T, TestResponseUpstreamURLHandlesResponsesWithoutRequestMetadata(), TestResponseUpstreamURLReturnsRequestURL()

### Community 202 - "Community 202"
Cohesion: 0.53
Nodes (5): frontendFile(), frontendRoot(), Engine, isBackendPath(), registerFrontend()

### Community 203 - "Community 203"
Cohesion: 0.33
Nodes (5): compilerOptions, baseUrl, paths, files, references

### Community 204 - "Community 204"
Cohesion: 0.40
Nodes (3): errorReader, T, TestCipherRoundTrip()

### Community 205 - "Community 205"
Cohesion: 0.40
Nodes (3): isResponsesCompactionRequest(), T, TestIsResponsesCompactionRequest()

### Community 206 - "Community 206"
Cohesion: 0.60
Nodes (4): T, TestParseImportedCredentialsAcceptsOneSSOTokenPerLine(), TestParseImportedCredentialsRejectsOversizedPlainToken(), TestWebCredentialJSONUsesCurrentDocumentShape()

### Community 207 - "Community 207"
Cohesion: 0.40
Nodes (4): imageStatsDTO, mediaAssetDTO, mediaJobDTO, videoStatsDTO

### Community 208 - "Community 208"
Cohesion: 0.67
Nodes (3): T, TestHTTPUpstreamFailureClassifiesBuildForbiddenBodies(), TestRetryableResponseHonorsUpstreamRetryVeto()

### Community 209 - "Community 209"
Cohesion: 0.67
Nodes (3): T, TestClientKeyBillingReservationsEnforceLimitAndExpire(), TestClientKeyUpdateDoesNotOverwriteConcurrentBillingState()

### Community 210 - "Community 210"
Cohesion: 0.25
Nodes (5): failingAdminRepository, rejectingRateLimiter, AdminRepository, Context, Time

### Community 211 - "Community 211"
Cohesion: 0.67
Nodes (3): T, TestRunWebAccountScriptsRejectsInvalidRequestsBeforeSSE(), TestRunWebAccountScriptsRejectsOversizedSelectionBeforeSSE()

### Community 213 - "Community 213"
Cohesion: 0.50
Nodes (3): T, TestRewriteAliasedModelAppliesOperationEffort(), rewriteAliasedModel()

## Knowledge Gaps
- **382 isolated node(s):** `github.com/chenyme/grok2api/backend`, `providerLinkRepository`, `providerPolicy`, `identitySynchronizer`, `billingReservationRepository` (+377 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **34 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `Duration` connect `Community 14` to `Community 128`, `Community 129`, `Community 4`, `Community 134`, `Community 7`, `Community 8`, `Community 136`, `Community 138`, `Community 9`, `Community 140`, `Community 13`, `Community 141`, `Community 15`, `Community 16`, `Community 24`, `Community 25`, `Community 27`, `Community 31`, `Community 162`, `Community 34`, `Community 41`, `Community 42`, `Community 47`, `Community 48`, `Community 49`, `Community 175`, `Community 179`, `Community 57`, `Community 59`, `Community 68`, `Community 69`, `Community 70`, `Community 71`, `Community 72`, `Community 75`, `Community 78`, `Community 79`, `Community 80`, `Community 96`, `Community 99`, `Community 102`, `Community 107`, `Community 111`, `Community 115`, `Community 117`, `Community 120`, `Community 124`?**
  _High betweenness centrality (0.088) - this node is a cross-community bridge._
- **Why does `Credential` connect `Community 18` to `Community 129`, `Community 130`, `Community 132`, `Community 5`, `Community 6`, `Community 7`, `Community 8`, `Community 138`, `Community 11`, `Community 15`, `Community 143`, `Community 16`, `Community 147`, `Community 152`, `Community 25`, `Community 154`, `Community 26`, `Community 27`, `Community 29`, `Community 31`, `Community 32`, `Community 34`, `Community 163`, `Community 38`, `Community 170`, `Community 42`, `Community 184`, `Community 57`, `Community 59`, `Community 61`, `Community 189`, `Community 190`, `Community 63`, `Community 193`, `Community 195`, `Community 67`, `Community 75`, `Community 81`, `Community 94`, `Community 96`, `Community 99`, `Community 102`, `Community 114`, `Community 119`, `Community 126`, `Community 127`?**
  _High betweenness centrality (0.067) - this node is a cross-community bridge._
- **Why does `Provider` connect `Community 11` to `Community 129`, `Community 132`, `Community 5`, `Community 134`, `Community 7`, `Community 135`, `Community 8`, `Community 15`, `Community 16`, `Community 17`, `Community 18`, `Community 25`, `Community 28`, `Community 29`, `Community 32`, `Community 34`, `Community 38`, `Community 170`, `Community 42`, `Community 45`, `Community 46`, `Community 178`, `Community 53`, `Community 184`, `Community 189`, `Community 190`, `Community 62`, `Community 66`, `Community 195`, `Community 81`, `Community 89`, `Community 94`, `Community 99`, `Community 100`, `Community 109`, `Community 114`, `Community 116`?**
  _High betweenness centrality (0.044) - this node is a cross-community bridge._
- **Are the 17 inferred relationships involving `Duration` (e.g. with `buildSeries()` and `TestCleanupAllProtectedTerminatesWithoutDelete()`) actually correct?**
  _`Duration` has 17 INFERRED edges - model-reasoned connections that need verification._
- **What connects `github.com/chenyme/grok2api/backend`, `providerLinkRepository`, `providerPolicy` to the rest of the system?**
  _382 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Community 0` be split into smaller, more focused modules?**
  _Cohesion score 0.0637915543575921 - nodes in this community are weakly interconnected._
- **Should `Community 1` be split into smaller, more focused modules?**
  _Cohesion score 0.03434343434343434 - nodes in this community are weakly interconnected._
# Credits

Response Types:

- <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#CreditCheckResponse">CreditCheckResponse</a>

Methods:

- <code title="post /v1/credits">client.Credits.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#CreditService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#CreditCheckResponse">CreditCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Logs

Response Types:

- <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#LogNewResponse">LogNewResponse</a>
- <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#LogGetSummaryResponse">LogGetSummaryResponse</a>

Methods:

- <code title="post /v1/usage">client.Logs.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#LogService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#LogNewParams">LogNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#LogNewResponse">LogNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/usage/summary">client.Logs.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#LogService.GetSummary">GetSummary</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#LogGetSummaryParams">LogGetSummaryParams</a>) (\*<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#LogGetSummaryResponse">LogGetSummaryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AccessToken

Response Types:

- <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#AccessTokenNewResponse">AccessTokenNewResponse</a>

Methods:

- <code title="post /v1/token">client.AccessToken.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#AccessTokenService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#AccessTokenNewParams">AccessTokenNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#AccessTokenNewResponse">AccessTokenNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# VerifyToken

Response Types:

- <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#VerifyTokenVerifyResponse">VerifyTokenVerifyResponse</a>

Methods:

- <code title="post /v1/token/verify">client.VerifyToken.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#VerifyTokenService.Verify">Verify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#VerifyTokenVerifyResponse">VerifyTokenVerifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# CamsKfintech

Response Types:

- <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#LinkedHolder">LinkedHolder</a>
- <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#Transaction">Transaction</a>
- <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#UnifiedResponse">UnifiedResponse</a>

Methods:

- <code title="post /v4/cams_kfintech/parse">client.CamsKfintech.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#CamsKfintechService.Parse">Parse</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#CamsKfintechParseParams">CamsKfintechParseParams</a>) (\*<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#UnifiedResponse">UnifiedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Cdsl

Methods:

- <code title="post /v4/cdsl/parse">client.Cdsl.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#CdslService.ParsePdf">ParsePdf</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#CdslParsePdfParams">CdslParsePdfParams</a>) (\*<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#UnifiedResponse">UnifiedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Fetch

Response Types:

- <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#CdslFetchRequestOtpResponse">CdslFetchRequestOtpResponse</a>
- <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#CdslFetchVerifyOtpResponse">CdslFetchVerifyOtpResponse</a>

Methods:

- <code title="post /v4/cdsl/fetch">client.Cdsl.Fetch.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#CdslFetchService.RequestOtp">RequestOtp</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#CdslFetchRequestOtpParams">CdslFetchRequestOtpParams</a>) (\*<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#CdslFetchRequestOtpResponse">CdslFetchRequestOtpResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v4/cdsl/fetch/{session_id}/verify">client.Cdsl.Fetch.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#CdslFetchService.VerifyOtp">VerifyOtp</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sessionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#CdslFetchVerifyOtpParams">CdslFetchVerifyOtpParams</a>) (\*<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#CdslFetchVerifyOtpResponse">CdslFetchVerifyOtpResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ContractNote

Response Types:

- <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#ContractNoteParseResponse">ContractNoteParseResponse</a>

Methods:

- <code title="post /v4/contract_note/parse">client.ContractNote.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#ContractNoteService.Parse">Parse</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#ContractNoteParseParams">ContractNoteParseParams</a>) (\*<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#ContractNoteParseResponse">ContractNoteParseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Inbox

Response Types:

- <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboxCheckConnectionStatusResponse">InboxCheckConnectionStatusResponse</a>
- <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboxConnectEmailResponse">InboxConnectEmailResponse</a>
- <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboxDisconnectEmailResponse">InboxDisconnectEmailResponse</a>
- <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboxListCasFilesResponse">InboxListCasFilesResponse</a>

Methods:

- <code title="post /v4/inbox/status">client.Inbox.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboxService.CheckConnectionStatus">CheckConnectionStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboxCheckConnectionStatusParams">InboxCheckConnectionStatusParams</a>) (\*<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboxCheckConnectionStatusResponse">InboxCheckConnectionStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v4/inbox/connect">client.Inbox.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboxService.ConnectEmail">ConnectEmail</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboxConnectEmailParams">InboxConnectEmailParams</a>) (\*<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboxConnectEmailResponse">InboxConnectEmailResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v4/inbox/disconnect">client.Inbox.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboxService.DisconnectEmail">DisconnectEmail</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboxDisconnectEmailParams">InboxDisconnectEmailParams</a>) (\*<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboxDisconnectEmailResponse">InboxDisconnectEmailResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v4/inbox/cas">client.Inbox.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboxService.ListCasFiles">ListCasFiles</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboxListCasFilesParams">InboxListCasFilesParams</a>) (\*<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboxListCasFilesResponse">InboxListCasFilesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Kfintech

Response Types:

- <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#KfintechGenerateCasResponse">KfintechGenerateCasResponse</a>

Methods:

- <code title="post /v4/kfintech/generate">client.Kfintech.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#KfintechService.GenerateCas">GenerateCas</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#KfintechGenerateCasParams">KfintechGenerateCasParams</a>) (\*<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#KfintechGenerateCasResponse">KfintechGenerateCasResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Nsdl

Methods:

- <code title="post /v4/nsdl/parse">client.Nsdl.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#NsdlService.Parse">Parse</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#NsdlParseParams">NsdlParseParams</a>) (\*<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#UnifiedResponse">UnifiedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Smart

Methods:

- <code title="post /v4/smart/parse">client.Smart.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#SmartService.ParseCasPdf">ParseCasPdf</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#SmartParseCasPdfParams">SmartParseCasPdfParams</a>) (\*<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#UnifiedResponse">UnifiedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# InboundEmail

Response Types:

- <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboundEmailNewResponse">InboundEmailNewResponse</a>
- <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboundEmailGetResponse">InboundEmailGetResponse</a>
- <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboundEmailListResponse">InboundEmailListResponse</a>
- <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboundEmailDeleteResponse">InboundEmailDeleteResponse</a>

Methods:

- <code title="post /v4/inbound-email">client.InboundEmail.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboundEmailService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboundEmailNewParams">InboundEmailNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboundEmailNewResponse">InboundEmailNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v4/inbound-email/{inbound_email_id}">client.InboundEmail.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboundEmailService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inboundEmailID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboundEmailGetResponse">InboundEmailGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v4/inbound-email">client.InboundEmail.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboundEmailService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboundEmailListParams">InboundEmailListParams</a>) (\*<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboundEmailListResponse">InboundEmailListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v4/inbound-email/{inbound_email_id}">client.InboundEmail.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboundEmailService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inboundEmailID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go">casparser</a>.<a href="https://pkg.go.dev/github.com/CASParser/cas-parser-go#InboundEmailDeleteResponse">InboundEmailDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

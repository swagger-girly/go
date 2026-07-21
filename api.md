# Shared Params Types

- <a href="https://pkg.go.dev/github.com/swagger-girly/go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go/shared#AddressParam">AddressParam</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go/shared#MoneyParam">MoneyParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/swagger-girly/go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go/shared#Address">Address</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go/shared#Money">Money</a>

# helloworldtestingggg

Response Types:

- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#SystemHealth">SystemHealth</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#GetRateLimitsResponse">GetRateLimitsResponse</a>

Methods:

- <code title="get /health">client.<a href="https://pkg.go.dev/github.com/swagger-girly/go#HelloworldtestinggggService.Health">Health</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#SystemHealth">SystemHealth</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /rate_limits">client.<a href="https://pkg.go.dev/github.com/swagger-girly/go#HelloworldtestinggggService.GetRateLimits">GetRateLimits</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#GetRateLimitsResponse">GetRateLimitsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Pet

Params Types:

- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetParam">PetParam</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetStatus">PetStatus</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#ConnectClientEventUnionParam">ConnectClientEventUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Pet">Pet</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetStatus">PetStatus</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetListFakePageResponse">PetListFakePageResponse</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetListUnpaginatedResponse">PetListUnpaginatedResponse</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetGetPremiumResponse">PetGetPremiumResponse</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetUploadImageResponse">PetUploadImageResponse</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#ConnectClientEvent">ConnectClientEvent</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#ConnectServerEvent">ConnectServerEvent</a>

Methods:

- <code title="post /pet">client.Pet.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetNewParams">PetNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/{petId}">client.Pet.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /pet">client.Pet.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetUpdateParams">PetUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet">client.Pet.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetListParams">PetListParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go/packages/pagination#CustomCursorPage">CustomCursorPage</a>[<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Pet">Pet</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /pet/{petId}">client.Pet.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /pet/findByStatus">client.Pet.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetService.FindByStatus">FindByStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetFindByStatusParams">PetFindByStatusParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/findByTags">client.Pet.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetService.FindByTags">FindByTags</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetFindByTagsParams">PetFindByTagsParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/fake-page">client.Pet.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetService.ListFakePage">ListFakePage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetListFakePageResponse">PetListFakePageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/fake-page-inferred">client.Pet.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetService.ListFakePageInferred">ListFakePageInferred</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go/packages/pagination#XFakeSinglePage">XFakeSinglePage</a>[<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Pet">Pet</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/unpaginated">client.Pet.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetService.ListUnpaginated">ListUnpaginated</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetListUnpaginatedParams">PetListUnpaginatedParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetListUnpaginatedResponse">PetListUnpaginatedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/{petId}/premium">client.Pet.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetService.GetPremium">GetPremium</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetGetPremiumResponse">PetGetPremiumResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /pet/{petId}">client.Pet.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetService.UpdateWithForm">UpdateWithForm</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetUpdateWithFormParams">PetUpdateWithFormParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /pet/{petId}/uploadImage">client.Pet.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetService.UploadImage">UploadImage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/builtin#io.Reader">io.Reader</a>, params <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetUploadImageParams">PetUploadImageParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetUploadImageResponse">PetUploadImageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/{petId}/status/stream">client.Pet.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetService.WatchStatus">WatchStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetWatchStatusParams">PetWatchStatusParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Files

Response Types:

- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Archive">Archive</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#File">File</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Fileslist">Fileslist</a>

Methods:

- <code title="put /files/{path}">client.Files.<a href="https://pkg.go.dev/github.com/swagger-girly/go#FileService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, path <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#FileUpdateParams">FileUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#File">File</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /files">client.Files.<a href="https://pkg.go.dev/github.com/swagger-girly/go#FileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Fileslist">Fileslist</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /files/{path}">client.Files.<a href="https://pkg.go.dev/github.com/swagger-girly/go#FileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, path <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /files/archive">client.Files.<a href="https://pkg.go.dev/github.com/swagger-girly/go#FileService.NewArchive">NewArchive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#FileNewArchiveParams">FileNewArchiveParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Archive">Archive</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /files/{path}">client.Files.<a href="https://pkg.go.dev/github.com/swagger-girly/go#FileService.Download">Download</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, path <a href="https://pkg.go.dev/builtin#string">string</a>) (\*http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /files">client.Files.<a href="https://pkg.go.dev/github.com/swagger-girly/go#FileService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#FileUploadParams">FileUploadParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#File">File</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /files/direct">client.Files.<a href="https://pkg.go.dev/github.com/swagger-girly/go#FileService.UploadDirect">UploadDirect</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/builtin#io.Reader">io.Reader</a>, params <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#FileUploadDirectParams">FileUploadDirectParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#File">File</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Profiles

Params Types:

- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#ProfileCreateRequestParam">ProfileCreateRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Profile">Profile</a>

Methods:

- <code title="post /profiles">client.Profiles.<a href="https://pkg.go.dev/github.com/swagger-girly/go#ProfileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#ProfileNewParams">ProfileNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Profile">Profile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /profiles/{profileId}">client.Profiles.<a href="https://pkg.go.dev/github.com/swagger-girly/go#ProfileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, profileID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Profile">Profile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /profiles/{profileId}">client.Profiles.<a href="https://pkg.go.dev/github.com/swagger-girly/go#ProfileService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, profileID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#ProfileUpdateParams">ProfileUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Profile">Profile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /profiles/legacy-search">client.Profiles.<a href="https://pkg.go.dev/github.com/swagger-girly/go#ProfileService.LegacySearch">LegacySearch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#ProfileLegacySearchParams">ProfileLegacySearchParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Profile">Profile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Adoptions

Response Types:

- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Application">Application</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#AdoptionGetDecisionResponse">AdoptionGetDecisionResponse</a>

Methods:

- <code title="post /adoptions/applications">client.Adoptions.<a href="https://pkg.go.dev/github.com/swagger-girly/go#AdoptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#AdoptionNewParams">AdoptionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Application">Application</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /adoptions/applications/{applicationId}">client.Adoptions.<a href="https://pkg.go.dev/github.com/swagger-girly/go#AdoptionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, applicationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Application">Application</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /adoptions/applications/{applicationId}/decision">client.Adoptions.<a href="https://pkg.go.dev/github.com/swagger-girly/go#AdoptionService.GetDecision">GetDecision</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, applicationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#AdoptionGetDecisionResponse">AdoptionGetDecisionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Policies

Params Types:

- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#AdoptionRuleParam">AdoptionRuleParam</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#ConditionValueUnionParam">ConditionValueUnionParam</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#RuleAttribute">RuleAttribute</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#AdoptionRule">AdoptionRule</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#ConditionValueUnion">ConditionValueUnion</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Policy">Policy</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#RuleAttribute">RuleAttribute</a>

Methods:

- <code title="post /adoptions/policies">client.Adoptions.Policies.<a href="https://pkg.go.dev/github.com/swagger-girly/go#AdoptionPolicyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#AdoptionPolicyNewParams">AdoptionPolicyNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Policy">Policy</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /adoptions/policies/{policyId}">client.Adoptions.Policies.<a href="https://pkg.go.dev/github.com/swagger-girly/go#AdoptionPolicyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, policyID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Policy">Policy</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /adoptions/policies/{policyId}">client.Adoptions.Policies.<a href="https://pkg.go.dev/github.com/swagger-girly/go#AdoptionPolicyService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, policyID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#AdoptionPolicyUpdateParams">AdoptionPolicyUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Policy">Policy</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /adoptions/policies">client.Adoptions.Policies.<a href="https://pkg.go.dev/github.com/swagger-girly/go#AdoptionPolicyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#AdoptionPolicyListParams">AdoptionPolicyListParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go/packages/pagination#CustomCursorPage">CustomCursorPage</a>[<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Policy">Policy</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Placements

Params Types:

- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PlacementEventUnionParam">PlacementEventUnionParam</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#TransferLegParam">TransferLegParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Placement">Placement</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PlacementEvent">PlacementEvent</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#TransferLeg">TransferLeg</a>

Methods:

- <code title="post /placements">client.Placements.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PlacementService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PlacementNewParams">PlacementNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Placement">Placement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /placements/{placementId}">client.Placements.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PlacementService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, placementID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Placement">Placement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /placements">client.Placements.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PlacementService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PlacementListParams">PlacementListParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go/packages/pagination#CustomCursorPage">CustomCursorPage</a>[<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Placement">Placement</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /placements/{placementId}/events">client.Placements.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PlacementService.RecordEvent">RecordEvent</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, placementID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PlacementRecordEventParams">PlacementRecordEventParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Placement">Placement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Veterinary

Params Types:

- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#VaccinationRecordParam">VaccinationRecordParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#MedicalSummary">MedicalSummary</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#VaccinationRecord">VaccinationRecord</a>

# Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetCreatedWebhookEvent">PetCreatedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetUpdatedWebhookEvent">PetUpdatedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetInventoryLowWebhookEvent">PetInventoryLowWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PetModerationWebhookEvent">PetModerationWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#StoreReportGeneratedWebhookEvent">StoreReportGeneratedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#AdoptionsPolicyChangedWebhookEvent">AdoptionsPolicyChangedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#PlacementEventRecordedWebhookEvent">PlacementEventRecordedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#ParsedWebhookEvent">ParsedWebhookEvent</a>

# Store

Response Types:

- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#StoreListInventoryResponse">StoreListInventoryResponse</a>

Methods:

- <code title="get /store/inventory">client.Store.<a href="https://pkg.go.dev/github.com/swagger-girly/go#StoreService.ListInventory">ListInventory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#StoreListInventoryResponse">StoreListInventoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Order

Response Types:

- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#StoreOrderNewResponse">StoreOrderNewResponse</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#StoreOrderGetResponse">StoreOrderGetResponse</a>

Methods:

- <code title="post /store/order">client.Store.Order.<a href="https://pkg.go.dev/github.com/swagger-girly/go#StoreOrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#StoreOrderNewParams">StoreOrderNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#StoreOrderNewResponse">StoreOrderNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /store/order/{orderId}">client.Store.Order.<a href="https://pkg.go.dev/github.com/swagger-girly/go#StoreOrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#StoreOrderGetResponse">StoreOrderGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /store/order/{orderId}">client.Store.Order.<a href="https://pkg.go.dev/github.com/swagger-girly/go#StoreOrderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Reports

Response Types:

- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#DailyInventory">DailyInventory</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#InventoryResponse">InventoryResponse</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Report">Report</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#ReportList">ReportList</a>

Methods:

- <code title="get /store/reports/{reportId}">client.Store.Reports.<a href="https://pkg.go.dev/github.com/swagger-girly/go#StoreReportService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, reportID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Report">Report</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /store/reports">client.Store.Reports.<a href="https://pkg.go.dev/github.com/swagger-girly/go#StoreReportService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#StoreReportListParams">StoreReportListParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go/packages/pagination#ReportCursorPage">ReportCursorPage</a>[<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#Report">Report</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /store/reports/{reportId}/embed">client.Store.Reports.<a href="https://pkg.go.dev/github.com/swagger-girly/go#StoreReportService.Embed">Embed</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, reportID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /store/reports/{reportId}/pause">client.Store.Reports.<a href="https://pkg.go.dev/github.com/swagger-girly/go#StoreReportService.Pause">Pause</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, reportID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#StoreReportPauseParams">StoreReportPauseParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Inventory

Methods:

- <code title="get /store/reports/{reportId}/inventory">client.Store.Reports.Inventory.<a href="https://pkg.go.dev/github.com/swagger-girly/go#StoreReportInventoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, reportID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#InventoryResponse">InventoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Daily

Methods:

- <code title="get /store/reports/{reportId}/inventory/daily/{date}">client.Store.Reports.Inventory.Daily.<a href="https://pkg.go.dev/github.com/swagger-girly/go#StoreReportInventoryDailyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, reportID <a href="https://pkg.go.dev/builtin#string">string</a>, date <a href="https://pkg.go.dev/builtin#time.Time">time.Time</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#DailyInventory">DailyInventory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# User

Response Types:

- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#UserNewResponse">UserNewResponse</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#UserGetResponse">UserGetResponse</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#UserNewWithListResponse">UserNewWithListResponse</a>

Methods:

- <code title="post /user">client.User.<a href="https://pkg.go.dev/github.com/swagger-girly/go#UserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#UserNewParams">UserNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#UserNewResponse">UserNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/{username}">client.User.<a href="https://pkg.go.dev/github.com/swagger-girly/go#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#UserGetResponse">UserGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /user/{username}">client.User.<a href="https://pkg.go.dev/github.com/swagger-girly/go#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#UserUpdateParams">UserUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /user/{username}">client.User.<a href="https://pkg.go.dev/github.com/swagger-girly/go#UserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /user/createWithList">client.User.<a href="https://pkg.go.dev/github.com/swagger-girly/go#UserService.NewWithList">NewWithList</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#UserNewWithListParams">UserNewWithListParams</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#UserNewWithListResponse">UserNewWithListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/login">client.User.<a href="https://pkg.go.dev/github.com/swagger-girly/go#UserService.Login">Login</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#UserLoginParams">UserLoginParams</a>) (\*<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/logout">client.User.<a href="https://pkg.go.dev/github.com/swagger-girly/go#UserService.Logout">Logout</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Media

Response Types:

- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#MediaBinaryAndJsonResponse">MediaBinaryAndJsonResponse</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#MediaJsonAndXmlResponse">MediaJsonAndXmlResponse</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#MediaJsonFirstResponse">MediaJsonFirstResponse</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#MediaSchemaOnJsonResponse">MediaSchemaOnJsonResponse</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#MediaTextFirstResponse">MediaTextFirstResponse</a>
- <a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#MediaVendorJsonResponse">MediaVendorJsonResponse</a>

Methods:

- <code title="get /media/binary-and-json">client.Media.<a href="https://pkg.go.dev/github.com/swagger-girly/go#MediaService.BinaryAndJson">BinaryAndJson</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#MediaBinaryAndJsonResponse">MediaBinaryAndJsonResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /media/json-and-xml">client.Media.<a href="https://pkg.go.dev/github.com/swagger-girly/go#MediaService.JsonAndXml">JsonAndXml</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#MediaJsonAndXmlResponse">MediaJsonAndXmlResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /media/json-first">client.Media.<a href="https://pkg.go.dev/github.com/swagger-girly/go#MediaService.JsonFirst">JsonFirst</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#MediaJsonFirstResponse">MediaJsonFirstResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /media/schema-on-json">client.Media.<a href="https://pkg.go.dev/github.com/swagger-girly/go#MediaService.SchemaOnJson">SchemaOnJson</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#MediaSchemaOnJsonResponse">MediaSchemaOnJsonResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /media/schema-on-text">client.Media.<a href="https://pkg.go.dev/github.com/swagger-girly/go#MediaService.SchemaOnText">SchemaOnText</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /media/text-first">client.Media.<a href="https://pkg.go.dev/github.com/swagger-girly/go#MediaService.TextFirst">TextFirst</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#MediaTextFirstResponse">MediaTextFirstResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /media/text-only">client.Media.<a href="https://pkg.go.dev/github.com/swagger-girly/go#MediaService.TextOnly">TextOnly</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /media/vendor-json">client.Media.<a href="https://pkg.go.dev/github.com/swagger-girly/go#MediaService.VendorJson">VendorJson</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/swagger-girly/go">helloworldtestingggg</a>.<a href="https://pkg.go.dev/github.com/swagger-girly/go#MediaVendorJsonResponse">MediaVendorJsonResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

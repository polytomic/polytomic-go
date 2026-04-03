# Reference
## BulkSync
<details><summary><code>client.BulkSync.List() -> *polytomic.BulkSyncListEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.BulkSyncListRequest{
        Active: polytomic.Bool(
            true,
        ),
    }
client.BulkSync.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**active:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.Create(request) -> *polytomic.BulkSyncResponseEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new Bulk Sync from a source to a destination (data warehouse, database, or cloud storage bucket like S3).

Bulk Syncs are used for the ELT pattern (Extract, Load, and Transform), where you want to sync un-transformed data to your data warehouses, databases, or cloud storage buckets like S3.

All of the functionality described in [the product
documentation](https://docs.polytomic.com/docs/bulk-syncs) is configurable via
the API.

Sample code examples:

- [Bulk sync (ELT) from Salesforce to S3](https://apidocs.polytomic.com/guides/code-examples/bulk-sync-elt-from-salesforce-to-s-3)
- [Bulk sync (ELT) from Salesforce to Snowflake](https://apidocs.polytomic.com/guides/code-examples/bulk-sync-elt-from-salesforce-to-snowflake)
- [Bulk sync (ELT) from HubSpot to PostgreSQL](https://apidocs.polytomic.com/guides/code-examples/bulk-sync-elt-from-hub-spot-to-postgre-sql)

## Connection specific configuration

The `destination_configuration` is integration-specific configuration for the
selected bulk sync destination. This includes settings such as the output schema
and is required when creating a new sync.

The `source_configuration` is optional. It allows configuration for how
Polytomic reads data from the source connection. This will not be available for
integrations that do not support additional configuration.

Consult the [connection configurations](https://apidocs.polytomic.com/2024-02-08/guides/configuring-your-connections/overview)
to see configurations for particular integrations (for example, [here](https://apidocs.polytomic.com/2024-02-08/guides/configuring-your-connections/connections/postgre-sql#source-1) is the available source configuration for the PostgreSQL bulk sync source).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.CreateBulkSyncRequest{
        DestinationConfiguration: map[string]any{
            "schema": "my_schema",
        },
        DestinationConnectionId: "248df4b7-aa70-47b8-a036-33ac447e668d",
        Name: "My Bulk Sync",
        Schedule: &polytomic.BulkSchedule{
            Frequency: polytomic.ScheduleFrequencyManual,
        },
        SourceConnectionId: "248df4b7-aa70-47b8-a036-33ac447e668d",
    }
client.BulkSync.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**active:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**automaticallyAddNewFields:** `*polytomic.BulkDiscover` 
    
</dd>
</dl>

<dl>
<dd>

**automaticallyAddNewObjects:** `*polytomic.BulkDiscover` 
    
</dd>
</dl>

<dl>
<dd>

**concurrencyLimit:** `*int` — Override the default concurrency limit for this sync.
    
</dd>
</dl>

<dl>
<dd>

**dataCutoffTimestamp:** `*time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**destinationConfiguration:** `map[string]any` 
    
</dd>
</dl>

<dl>
<dd>

**destinationConnectionId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**disableRecordTimestamps:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**discover:** `*bool` — DEPRECATED: Use automatically_add_new_objects/automatically_add_new_fields instead
    
</dd>
</dl>

<dl>
<dd>

**mode:** `*polytomic.BulkSyncMode` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**normalizeNames:** `*polytomic.BulkNormalizeNames` 
    
</dd>
</dl>

<dl>
<dd>

**organizationId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**policies:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**resyncConcurrencyLimit:** `*int` — Override the default resync concurrency limit for this sync.
    
</dd>
</dl>

<dl>
<dd>

**schedule:** `*polytomic.BulkSchedule` 
    
</dd>
</dl>

<dl>
<dd>

**schemas:** `[]*polytomic.V2CreateBulkSyncRequestSchemasItem` — List of schemas to sync; if omitted, all schemas will be selected for syncing.
    
</dd>
</dl>

<dl>
<dd>

**sourceConfiguration:** `map[string]any` 
    
</dd>
</dl>

<dl>
<dd>

**sourceConnectionId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.Get(Id) -> *polytomic.BulkSyncResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.BulkSyncGetRequest{
        RefreshSchemas: polytomic.Bool(
            true,
        ),
    }
client.BulkSync.Get(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**refreshSchemas:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.Update(Id, request) -> *polytomic.BulkSyncResponseEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> 📘 Updating schemas
>
> Schema updates can be performed using the [Update Bulk Sync Schemas](https://apidocs.polytomic.com/api-reference/bulk-sync/schemas/patch) endpoint.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.UpdateBulkSyncRequest{
        DestinationConfiguration: map[string]any{
            "schema": "my_schema",
        },
        DestinationConnectionId: "248df4b7-aa70-47b8-a036-33ac447e668d",
        Name: "My Bulk Sync",
        Schedule: &polytomic.BulkSchedule{
            Frequency: polytomic.ScheduleFrequencyManual,
        },
        SourceConnectionId: "248df4b7-aa70-47b8-a036-33ac447e668d",
    }
client.BulkSync.Update(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**active:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**automaticallyAddNewFields:** `*polytomic.BulkDiscover` 
    
</dd>
</dl>

<dl>
<dd>

**automaticallyAddNewObjects:** `*polytomic.BulkDiscover` 
    
</dd>
</dl>

<dl>
<dd>

**concurrencyLimit:** `*int` — Override the default concurrency limit for this sync.
    
</dd>
</dl>

<dl>
<dd>

**dataCutoffTimestamp:** `*time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**destinationConfiguration:** `map[string]any` 
    
</dd>
</dl>

<dl>
<dd>

**destinationConnectionId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**disableRecordTimestamps:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**discover:** `*bool` — DEPRECATED: Use automatically_add_new_objects/automatically_add_new_fields instead
    
</dd>
</dl>

<dl>
<dd>

**mode:** `*polytomic.BulkSyncMode` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**normalizeNames:** `*polytomic.BulkNormalizeNames` 
    
</dd>
</dl>

<dl>
<dd>

**organizationId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**policies:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**resyncConcurrencyLimit:** `*int` — Override the default resync concurrency limit for this sync.
    
</dd>
</dl>

<dl>
<dd>

**schedule:** `*polytomic.BulkSchedule` 
    
</dd>
</dl>

<dl>
<dd>

**schemas:** `[]*polytomic.V2UpdateBulkSyncRequestSchemasItem` — List of schemas to sync; if omitted, all schemas will be selected for syncing.
    
</dd>
</dl>

<dl>
<dd>

**sourceConfiguration:** `map[string]any` 
    
</dd>
</dl>

<dl>
<dd>

**sourceConnectionId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.Remove(Id) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.BulkSyncRemoveRequest{
        RefreshSchemas: polytomic.Bool(
            true,
        ),
    }
client.BulkSync.Remove(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**refreshSchemas:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.Activate(Id, request) -> *polytomic.ActivateSyncEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.ActivateSyncInput{
        Active: true,
    }
client.BulkSync.Activate(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*polytomic.ActivateSyncInput` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.Cancel(Id) -> *polytomic.CancelBulkSyncResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.BulkSync.Cancel(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The active execution of this bulk sync ID will be cancelled.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.Start(Id, request) -> *polytomic.BulkSyncExecutionEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.StartBulkSyncRequest{}
client.BulkSync.Start(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**fetchMode:** `*polytomic.BulkFetchMode` 
    
</dd>
</dl>

<dl>
<dd>

**resync:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**schemas:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**test:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.GetStatus(Id) -> *polytomic.BulkSyncStatusEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.BulkSync.GetStatus(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.GetSource(Id) -> *polytomic.BulkSyncSourceEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.BulkSyncGetSourceRequest{
        IncludeFields: polytomic.Bool(
            true,
        ),
    }
client.BulkSync.GetSource(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.GetDestination(Id) -> *polytomic.BulkSyncDestEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.BulkSync.GetDestination(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Connections
<details><summary><code>client.Connections.GetTypes() -> *polytomic.ConnectionTypeResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Connections.GetTypes(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.GetConnectionTypeSchema(Id) -> *polytomic.JsonschemaSchema</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Connections.GetConnectionTypeSchema(
        context.TODO(),
        "postgresql",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.GetTypeParameterValues(Type, request) -> *polytomic.ConnectionParameterValuesResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.GetConnectionTypeParameterValuesRequestSchema{
        Field: "field",
    }
client.Connections.GetTypeParameterValues(
        context.TODO(),
        "type",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**type_:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**connectionId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**field:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**parameters:** `map[string]any` 
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.List() -> *polytomic.ConnectionListResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Connections.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.Create(request) -> *polytomic.CreateConnectionResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.CreateConnectionRequestSchema{
        Configuration: map[string]any{
            "database": "example",
            "hostname": "postgres.example.com",
            "password": "********",
            "port": 5432,
            "username": "user",
        },
        Name: "My Postgres Connection",
        Type: "postgresql",
    }
client.Connections.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**configuration:** `map[string]any` 
    
</dd>
</dl>

<dl>
<dd>

**healthcheckInterval:** `*string` — Override interval for connection health checking.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**organizationId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**policies:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**redirectUrl:** `*string` — URL to redirect to after completing OAuth flow.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**validate:** `*bool` — Validate connection configuration.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.Connect(request) -> *polytomic.ConnectCardResponseEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new request for [Polytomic Connect](https://www.polytomic.com/connect).

This endpoint configures a Polytomic Connect request and returns the URL to
redirect users to. This allows embedding Polytomic connection authorization in
other applications.

See also:

- [Embedding authentication](https://apidocs.polytomic.com/2024-02-08/guides/embedding-authentication), a guide to using Polytomic Connect.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.ConnectCardRequest{
        Name: "Salesforce Connection",
        RedirectUrl: "redirect_url",
    }
client.Connections.Connect(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**connection:** `*string` — The id of an existing connection to update.
    
</dd>
</dl>

<dl>
<dd>

**dark:** `*bool` — Whether to use the dark theme for the Connect modal.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — Name of the new connection. Must be unique per organization.
    
</dd>
</dl>

<dl>
<dd>

**organizationId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**redirectUrl:** `string` — URL to redirect to after connection is created.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` — Connection type to create.
    
</dd>
</dl>

<dl>
<dd>

**whitelist:** `[]string` — List of connection types which are allowed to be created. Ignored if type is set.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.TestConnection(request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Tests a connection configuration.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.TestConnectionRequest{
        Configuration: map[string]any{
            "database": "example",
            "hostname": "postgres.example.com",
            "password": "password",
            "port": 5432,
            "username": "user",
        },
        Type: "postgresql",
    }
client.Connections.TestConnection(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**configuration:** `map[string]any` — Connection configuration to test.
    
</dd>
</dl>

<dl>
<dd>

**connectionId:** `*string` — Optional existing connection ID to use as a base for testing. The provided configuration will be merged over the stored configuration for this connection before testing.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `string` — The type of connection to test.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.Get(Id) -> *polytomic.ConnectionResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Connections.Get(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.Update(Id, request) -> *polytomic.CreateConnectionResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.UpdateConnectionRequestSchema{
        Configuration: map[string]any{
            "database": "example",
            "hostname": "postgres.example.com",
            "password": "********",
            "port": 5432,
            "username": "user",
        },
        Name: "My Postgres Connection",
    }
client.Connections.Update(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**configuration:** `map[string]any` 
    
</dd>
</dl>

<dl>
<dd>

**healthcheckInterval:** `*string` — Override interval for connection health checking.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**organizationId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**policies:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**reconnect:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**validate:** `*bool` — Validate connection configuration.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.Remove(Id) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.ConnectionsRemoveRequest{
        Force: polytomic.Bool(
            true,
        ),
    }
client.Connections.Remove(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**force:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.GetParameterValues(Id) -> *polytomic.ConnectionParameterValuesResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Connections.GetParameterValues(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.CreateSharedConnection(ParentConnectionId, request) -> *polytomic.V2CreateSharedConnectionResponseEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> 🚧 Requires partner key
>
> Shared connections can only be created by using [partner keys](https://apidocs.polytomic.com/guides/obtaining-api-keys#partner-keys).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.ApiRequest{
        ChildOrganizationId: "248df4b7-aa70-47b8-a036-33ac447e668d",
    }
client.Connections.CreateSharedConnection(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**parentConnectionId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**childOrganizationId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.ListSharedConnections(ParentConnectionId) -> *polytomic.ConnectionListResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Connections.ListSharedConnections(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**parentConnectionId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## QueryRunner
<details><summary><code>client.QueryRunner.RunQuery(ConnectionId, request) -> *polytomic.V4RunQueryEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Submit a query for asynchronous execution against the connection. The initial response may only contain the query task id and status. Poll GET /api/queries/{id} with the returned id to retrieve completion status, fields, and results.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.V4RunQueryRequest{
        Query: polytomic.String(
            "SELECT * FROM table",
        ),
    }
client.QueryRunner.RunQuery(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**connectionId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` — The query to execute against the connection.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.QueryRunner.GetQuery(Id) -> *polytomic.V4QueryResultsEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Fetch the latest status for a submitted query and, once complete, return fields and paginated results. Use the query id returned by POST /api/connections/{connection_id}/query.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.QueryRunnerGetQueryRequest{
        Page: polytomic.String(
            "page",
        ),
    }
client.QueryRunner.GetQuery(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Schemas
<details><summary><code>client.Schemas.UpsertField(ConnectionId, SchemaId, request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.UpsertSchemaFieldRequest{}
client.Schemas.UpsertField(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "public.users",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**connectionId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**schemaId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**fields:** `[]*polytomic.V4UserFieldRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Schemas.DeleteField(ConnectionId, SchemaId, FieldId) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Schemas.DeleteField(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "public.users",
        "first_name",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**connectionId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**schemaId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**fieldId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Schemas.SetPrimaryKeys(ConnectionId, SchemaId, request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.SetPrimaryKeysRequest{}
client.Schemas.SetPrimaryKeys(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "public.users",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**connectionId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**schemaId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**fields:** `[]*polytomic.SchemaPrimaryKeyOverrideInput` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Schemas.ResetPrimaryKeys(ConnectionId, SchemaId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete all primary key overrides for a schema. After this call the schema will use the primary keys detected from the source connection, if any.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Schemas.ResetPrimaryKeys(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "public.users",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**connectionId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**schemaId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Schemas.Refresh(Id) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Schemas.Refresh(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Schemas.GetStatus(Id) -> *polytomic.BulkSyncSourceStatusEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Polytomic periodically inspects the schemas for connections to discover new fields and update metadata. This endpoint returns the current inspection status.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Schemas.GetStatus(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Schemas.Get(Id, SchemaId) -> *polytomic.BulkSyncSourceSchemaEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Schemas.Get(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "public.users",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**schemaId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Schemas.GetRecords(Id, SchemaId) -> *polytomic.SchemaRecordsResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Schemas.GetRecords(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "public.users",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**schemaId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Models
<details><summary><code>client.Models.GetEnrichmentSource(Id) -> *polytomic.GetModelSyncSourceMetaEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.ModelsGetEnrichmentSourceRequest{}
client.Models.GetEnrichmentSource(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**params:** `map[string][]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Models.Post(ConnectionId, request) -> *polytomic.V2GetEnrichmentInputFieldsResponseEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

For a given connection and enrichment configuration, provides the valid sets of input fields.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.GetEnrichmentInputFieldsRequest{}
client.Models.Post(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**connectionId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**configuration:** `*polytomic.V2EnricherConfiguration` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Models.Preview(request) -> *polytomic.ModelResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.ModelsPreviewRequest{
        Async: polytomic.Bool(
            true,
        ),
        Body: &polytomic.CreateModelRequest{
            Configuration: map[string]any{
                "table": "public.users",
            },
            ConnectionId: "248df4b7-aa70-47b8-a036-33ac447e668d",
            Name: "Users",
        },
    }
client.Models.Preview(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**async:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*polytomic.CreateModelRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Models.List() -> *polytomic.ModelListResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Models.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Models.Create(request) -> *polytomic.ModelResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.ModelsCreateRequest{
        Async: polytomic.Bool(
            true,
        ),
        Body: &polytomic.CreateModelRequest{
            Configuration: map[string]any{
                "table": "public.users",
            },
            ConnectionId: "248df4b7-aa70-47b8-a036-33ac447e668d",
            Name: "Users",
        },
    }
client.Models.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**async:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*polytomic.CreateModelRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Models.Get(Id) -> *polytomic.ModelResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.ModelsGetRequest{
        Async: polytomic.Bool(
            true,
        ),
    }
client.Models.Get(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**async:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Models.Update(Id, request) -> *polytomic.ModelResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.UpdateModelRequest{
        Async: polytomic.Bool(
            false,
        ),
        Configuration: map[string]any{
            "table": "public.users",
        },
        ConnectionId: "248df4b7-aa70-47b8-a036-33ac447e668d",
        Name: "Users",
    }
client.Models.Update(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**async:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**additionalFields:** `[]*polytomic.ModelModelFieldRequest` 
    
</dd>
</dl>

<dl>
<dd>

**configuration:** `map[string]any` 
    
</dd>
</dl>

<dl>
<dd>

**connectionId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**enricher:** `*polytomic.Enrichment` 
    
</dd>
</dl>

<dl>
<dd>

**fields:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**identifier:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**labels:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**organizationId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**policies:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**refresh:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**relations:** `[]*polytomic.ModelRelation` 
    
</dd>
</dl>

<dl>
<dd>

**trackingColumns:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Models.Remove(Id) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.ModelsRemoveRequest{
        Async: polytomic.Bool(
            true,
        ),
    }
client.Models.Remove(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**async:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Models.Sample(Id) -> *polytomic.ModelSampleResponseEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns sample records from the model. The first ten records that the source provides will be returned after being enriched (if applicable). Synchronous requests must complete within 10s. If either querying or enrichment exceeds 10s, please use the async option.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.ModelsSampleRequest{
        Async: polytomic.Bool(
            true,
        ),
    }
client.Models.Sample(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**async:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ModelSync
<details><summary><code>client.ModelSync.GetSource(Id) -> *polytomic.GetModelSyncSourceMetaEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.ModelSyncGetSourceRequest{}
client.ModelSync.GetSource(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**params:** `map[string][]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ModelSync.GetSourceFields(Id) -> *polytomic.ModelFieldResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.ModelSyncGetSourceFieldsRequest{}
client.ModelSync.GetSourceFields(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**params:** `map[string][]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ModelSync.List() -> *polytomic.ListModelSyncResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.ModelSyncListRequest{
        Active: polytomic.Bool(
            true,
        ),
        Mode: polytomic.ModelSyncModeCreate.Ptr(),
        TargetConnectionId: polytomic.String(
            "0b155265-c537-44c9-9359-a3ceb468a4da",
        ),
    }
client.ModelSync.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**active:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**mode:** `*polytomic.ModelSyncMode` 
    
</dd>
</dl>

<dl>
<dd>

**targetConnectionId:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ModelSync.Create(request) -> *polytomic.ModelSyncResponseEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new sync from one or more models to a destination.

All of the functionality described in [the product
documentation](https://docs.polytomic.com/docs/sync-destinations) is
configurable via the API.

Guides:

- [Model sync (Reverse ETL) from Snowflake query to Salesforce](https://apidocs.polytomic.com/2024-02-08/guides/code-examples/model-sync-reverse-etl-from-snowflake-query-to-salesforce)
- [Joined model sync from Postgres, Airtable, and Stripe to Hubspot](https://apidocs.polytomic.com/2024-02-08/guides/code-examples/joined-model-sync-from-postgres-airtable-and-stripe-to-hubspot)

## Targets (Destinations)

Polytomic refers to a model sync's destination as the "target object", or
target. Target objects are identified by a connection ID and an object ID. You
can retrieve a list of all target objects for a connection using the [Get Target
Objects](./targets/list) endpoint.

The `target` object in the request specifies information about the sync destination.

```json
"target": {
    "connection_id": "248df4b7-aa70-47b8-a036-33ac447e668d",
    "object": "Users",
},
```

Some connections support additional configuration for targets. For example,
[Salesforce
connections](https://apidocs.polytomic.com/2024-02-08/guides/configuring-your-connections/connections/salesforce#target)
support optionally specifying the ingestion API to use. The target specific
options are passed as `configuration`; consult the [integration
guides](https://apidocs.polytomic.com/2024-02-08/guides/configuring-your-connections/overview)
for details about specific connection configurations.

### Creating a new target

Some integrations support creating a new target when creating a model sync. For
example, an ad audience or database table.

When creating a new target, `object` is omitted and `create` is specified
instead. The `create` property is an object containing integration specific
configuration for the new target.

```json
"target": {
    "connection_id": "248df4b7-aa70-47b8-a036-33ac447e668d",
    "create": {
        "name": "New audience",
        "type": "user_audience"
    }
},
```

The [Get Target List](./targets/list) endpoint returns information about whether
a connection supports target creation.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.CreateModelSyncRequest{
        Fields: []*polytomic.ModelSyncField{
            &polytomic.ModelSyncField{
                Target: "name",
            },
        },
        Mode: polytomic.ModelSyncModeCreate,
        Name: "Users Sync",
        Schedule: &polytomic.Schedule{},
        Target: &polytomic.Target{
            ConnectionId: "248df4b7-aa70-47b8-a036-33ac447e668d",
        },
    }
client.ModelSync.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**active:** `*bool` — Whether the sync is enabled and scheduled.
    
</dd>
</dl>

<dl>
<dd>

**encryptionPassphrase:** `*string` — Passphrase for encrypting the sync data.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `[]*polytomic.ModelSyncField` — Fields to sync from source to destination.
    
</dd>
</dl>

<dl>
<dd>

**filterLogic:** `*string` — Logical expression to combine filters.
    
</dd>
</dl>

<dl>
<dd>

**filters:** `[]*polytomic.Filter` — Filters to apply to the source data.
    
</dd>
</dl>

<dl>
<dd>

**identity:** `*polytomic.Identity` 
    
</dd>
</dl>

<dl>
<dd>

**mode:** `*polytomic.ModelSyncMode` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**onlyEnrichUpdates:** `*bool` — Whether to use enrichment models as a source of possible changes to sync. If true, only changes to the base models will cause a record to sync.
    
</dd>
</dl>

<dl>
<dd>

**organizationId:** `*string` — Organization ID for the sync; read-only with a partner key.
    
</dd>
</dl>

<dl>
<dd>

**overrideFields:** `[]*polytomic.ModelSyncField` — Values to set in the target unconditionally.
    
</dd>
</dl>

<dl>
<dd>

**overrides:** `[]*polytomic.Override` — Conditional value replacement for fields.
    
</dd>
</dl>

<dl>
<dd>

**policies:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**schedule:** `*polytomic.Schedule` 
    
</dd>
</dl>

<dl>
<dd>

**skipInitialBackfill:** `*bool` — Whether to skip the initial backfill of records; if true only records seen after the sync is enabled will be synced.
    
</dd>
</dl>

<dl>
<dd>

**syncAllRecords:** `*bool` — Whether to sync all records from the source, regardless of whether they've changed since the previous execution.
    
</dd>
</dl>

<dl>
<dd>

**target:** `*polytomic.Target` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ModelSync.GetScheduleOptions() -> *polytomic.ScheduleOptionResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ModelSync.GetScheduleOptions(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ModelSync.Get(Id) -> *polytomic.ModelSyncResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ModelSync.Get(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ModelSync.Update(Id, request) -> *polytomic.ModelSyncResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.UpdateModelSyncRequest{
        Fields: []*polytomic.ModelSyncField{
            &polytomic.ModelSyncField{
                Target: "name",
            },
        },
        Mode: polytomic.ModelSyncModeCreate,
        Name: "Users Sync",
        Schedule: &polytomic.Schedule{},
        Target: &polytomic.Target{
            ConnectionId: "248df4b7-aa70-47b8-a036-33ac447e668d",
        },
    }
client.ModelSync.Update(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**active:** `*bool` — Whether the sync is enabled and scheduled.
    
</dd>
</dl>

<dl>
<dd>

**encryptionPassphrase:** `*string` — Passphrase for encrypting the sync data.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `[]*polytomic.ModelSyncField` — Fields to sync from source to destination.
    
</dd>
</dl>

<dl>
<dd>

**filterLogic:** `*string` — Logical expression to combine filters.
    
</dd>
</dl>

<dl>
<dd>

**filters:** `[]*polytomic.Filter` — Filters to apply to the source data.
    
</dd>
</dl>

<dl>
<dd>

**identity:** `*polytomic.Identity` 
    
</dd>
</dl>

<dl>
<dd>

**mode:** `*polytomic.ModelSyncMode` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**onlyEnrichUpdates:** `*bool` — Whether to use enrichment models as a source of possible changes to sync. If true, only changes to the base models will cause a record to sync.
    
</dd>
</dl>

<dl>
<dd>

**organizationId:** `*string` — Organization ID for the sync; read-only with a partner key.
    
</dd>
</dl>

<dl>
<dd>

**overrideFields:** `[]*polytomic.ModelSyncField` — Values to set in the target unconditionally.
    
</dd>
</dl>

<dl>
<dd>

**overrides:** `[]*polytomic.Override` — Conditional value replacement for fields.
    
</dd>
</dl>

<dl>
<dd>

**policies:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**schedule:** `*polytomic.Schedule` 
    
</dd>
</dl>

<dl>
<dd>

**skipInitialBackfill:** `*bool` — Whether to skip the initial backfill of records; if true only records seen after the sync is enabled will be synced.
    
</dd>
</dl>

<dl>
<dd>

**syncAllRecords:** `*bool` — Whether to sync all records from the source, regardless of whether they've changed since the previous execution.
    
</dd>
</dl>

<dl>
<dd>

**target:** `*polytomic.Target` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ModelSync.Remove(Id) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ModelSync.Remove(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ModelSync.Activate(Id, request) -> *polytomic.ActivateSyncEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.ActivateSyncInput{
        Active: true,
    }
client.ModelSync.Activate(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*polytomic.ActivateSyncInput` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ModelSync.Cancel(Id) -> *polytomic.CancelModelSyncResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ModelSync.Cancel(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The active execution of this sync ID will be cancelled.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ModelSync.Start(Id, request) -> *polytomic.StartModelSyncResponseEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> 🚧 Force full resync
>
> Use caution when setting the `resync` parameter to `true`. This will force a full resync of the data from the source system. This can be a time-consuming operation and may impact the performance of the source system. It is recommended to only use this option when necessary.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.StartModelSyncRequest{}
client.ModelSync.Start(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**identities:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**resync:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**test:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ModelSync.GetStatus(Id) -> *polytomic.SyncStatusEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ModelSync.GetStatus(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Events
<details><summary><code>client.Events.List() -> *polytomic.EventsEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.EventsListRequest{
        OrganizationId: polytomic.String(
            "248df4b7-aa70-47b8-a036-33ac447e668d",
        ),
        Type: polytomic.String(
            "type",
        ),
        StartingAfter: polytomic.Time(
            polytomic.MustParseDateTime(
                "2020-01-01T00:00:00Z",
            ),
        ),
        EndingBefore: polytomic.Time(
            polytomic.MustParseDateTime(
                "2020-01-01T00:00:00Z",
            ),
        ),
        Limit: polytomic.Int(
            1,
        ),
    }
client.Events.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**startingAfter:** `*time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**endingBefore:** `*time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Events.GetTypes() -> *polytomic.EventTypesEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Events.GetTypes(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Jobs
<details><summary><code>client.Jobs.Get(Id, Type) -> *polytomic.JobResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Jobs.Get(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "createmodel",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**type_:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Identity
<details><summary><code>client.Identity.Get() -> *polytomic.GetIdentityResponseEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns information about the caller's identity.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Identity.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Notifications
<details><summary><code>client.Notifications.GetGlobalErrorSubscribers() -> *polytomic.V4GlobalErrorSubscribersResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Notifications.GetGlobalErrorSubscribers(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Notifications.SetGlobalErrorSubscribers(request) -> *polytomic.V4GlobalErrorSubscribersResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.V4GlobalErrorSubscribersRequest{}
client.Notifications.SetGlobalErrorSubscribers(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**emails:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organization
<details><summary><code>client.Organization.List() -> *polytomic.OrganizationsEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> 🚧 Requires partner key
>
> Organization endpoints are only accessible using [partner keys](https://apidocs.polytomic.com/guides/obtaining-api-keys#partner-keys).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organization.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.Create(request) -> *polytomic.OrganizationEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> 🚧 Requires partner key
>
> Organization endpoints are only accessible using [partner keys](https://apidocs.polytomic.com/guides/obtaining-api-keys#partner-keys).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.CreateOrganizationRequestSchema{
        Name: "My Organization",
    }
client.Organization.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**clientSecret:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**issuer:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**ssoDomain:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**ssoOrgId:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.Get(Id) -> *polytomic.OrganizationEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> 🚧 Requires partner key
>
> Organization endpoints are only accessible using [partner keys](https://apidocs.polytomic.com/guides/obtaining-api-keys#partner-keys).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organization.Get(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.Update(Id, request) -> *polytomic.OrganizationEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> 🚧 Requires partner key
>
> Organization endpoints are only accessible using [partner keys](https://apidocs.polytomic.com/guides/obtaining-api-keys#partner-keys).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.UpdateOrganizationRequestSchema{
        Name: "My Organization",
    }
client.Organization.Update(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**clientId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**clientSecret:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**issuer:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**ssoDomain:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**ssoOrgId:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organization.Remove(Id) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> 🚧 Requires partner key
>
> Organization endpoints are only accessible using [partner keys](https://apidocs.polytomic.com/guides/obtaining-api-keys#partner-keys).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organization.Remove(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users
<details><summary><code>client.Users.List(OrgId) -> *polytomic.ListUsersEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.List(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Create(OrgId, request) -> *polytomic.UserEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.CreateUserRequestSchema{
        Email: "mail@example.com",
    }
client.Users.Create(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**email:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**role:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**roleIds:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Get(Id, OrgId) -> *polytomic.UserEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Get(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**orgId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Update(Id, OrgId, request) -> *polytomic.UserEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.UpdateUserRequestSchema{
        Email: "mail@example.com",
    }
client.Users.Update(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**orgId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**email:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**role:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**roleIds:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Remove(Id, OrgId) -> *polytomic.UserEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Remove(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**orgId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.CreateApiKey(OrgId, Id) -> *polytomic.ApiKeyResponseEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> 🚧 Requires partner key
>
> User endpoints are only accessible using [partner keys](https://apidocs.polytomic.com/guides/obtaining-api-keys#partner-keys).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.UsersCreateApiKeyRequest{
        Force: polytomic.Bool(
            true,
        ),
    }
client.Users.CreateApiKey(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**force:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Webhooks
<details><summary><code>client.Webhooks.List() -> *polytomic.WebhookListEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Webooks can be set up using the webhook API endpoints. Currently, only one
webhook may be created per organization. The webhook will be called for events
in that organization.

Consult the [Events documentation](https://apidocs.polytomic.com/guides/events) for more information.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Webhooks.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Create(request) -> *polytomic.WebhookEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Webooks can be set up using the webhook API endpoints. Currently, only one
webhook may be created per organization. The webhook will be called for events
in that organization.

Consult the [Events documentation](https://apidocs.polytomic.com/guides/events) for more information.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.CreateWebhooksSchema{
        Endpoint: "https://example.com/webhook",
        Secret: "secret",
    }
client.Webhooks.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**endpoint:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**organizationId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**secret:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Get(Id) -> *polytomic.WebhookEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Webooks can be set up using the webhook API endpoints. Currently, only one
webhook may be created per organization. The webhook will be called for events
in that organization.

Consult the [Events documentation](https://apidocs.polytomic.com/guides/events) for more information.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Webhooks.Get(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Update(Id, request) -> *polytomic.WebhookEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Webooks can be set up using the webhook API endpoints. Currently, only one
webhook may be created per organization. The webhook will be called for events
in that organization.

Consult the [Events documentation](https://apidocs.polytomic.com/guides/events) for more information.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &polytomic.UpdateWebhooksSchema{
        Endpoint: "https://example.com/webhook",
        Secret: "secret",
    }
client.Webhooks.Update(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**endpoint:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**organizationId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**secret:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Remove(Id) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Webhooks.Remove(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Disable(Id) -> *polytomic.WebhookEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Webhooks.Disable(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Enable(Id) -> *polytomic.WebhookEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Webhooks.Enable(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## BulkSync Executions
<details><summary><code>client.BulkSync.Executions.ListStatus() -> *polytomic.ListBulkSyncExecutionStatusEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bulksync.ExecutionsListStatusRequest{
        All: polytomic.Bool(
            true,
        ),
        Active: polytomic.Bool(
            true,
        ),
    }
client.BulkSync.Executions.ListStatus(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**all:** `*bool` — Return the execution status of all syncs in the organization
    
</dd>
</dl>

<dl>
<dd>

**active:** `*bool` — Return the execution status of all active syncs in the organization
    
</dd>
</dl>

<dl>
<dd>

**syncId:** `*string` — Return the execution status of the specified sync; this may be supplied multiple times.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.Executions.List(Id) -> *polytomic.ListBulkSyncExecutionsEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bulksync.ExecutionsListRequest{
        PageToken: polytomic.String(
            "AmkYh8v0jR5B3kls2Qcc9y8MjrPmvR4CvaK7H0F4rEwqvg76K==",
        ),
        OnlyTerminal: polytomic.Bool(
            true,
        ),
        Ascending: polytomic.Bool(
            true,
        ),
        Limit: polytomic.Int(
            100,
        ),
    }
client.BulkSync.Executions.List(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**pageToken:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**onlyTerminal:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**ascending:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.Executions.Get(Id, ExecId) -> *polytomic.BulkSyncExecutionEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.BulkSync.Executions.Get(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**execId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.Executions.Cancel(Id, ExecId) -> *polytomic.CancelBulkSyncResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.BulkSync.Executions.Cancel(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The bulk sync ID.
    
</dd>
</dl>

<dl>
<dd>

**execId:** `string` — The execution ID to cancel.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.Executions.GetLogs(SyncId, ExecutionId) -> *polytomic.V4BulkSyncExecutionLogsEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.BulkSync.Executions.GetLogs(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**syncId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**executionId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.Executions.ExportLogs(SyncId, ExecutionId) -> *polytomic.V4ExportSyncLogsEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bulksync.ExecutionsExportLogsRequest{
        Notify: polytomic.Bool(
            true,
        ),
    }
client.BulkSync.Executions.ExportLogs(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**syncId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**executionId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**notify:** `*bool` — Send a notification to the user when the logs are ready for download.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## BulkSync Schemas
<details><summary><code>client.BulkSync.Schemas.List(Id) -> *polytomic.ListBulkSchema</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bulksync.SchemasListRequest{}
client.BulkSync.Schemas.List(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**filters:** `map[string]*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.Schemas.Patch(Id, request) -> *polytomic.ListBulkSchema</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bulksync.BulkSyncSchemasRequest{}
client.BulkSync.Schemas.Patch(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**schemas:** `[]*polytomic.BulkSchema` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.Schemas.Get(Id, SchemaId) -> *polytomic.BulkSchemaEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.BulkSync.Schemas.Get(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "Contact",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**schemaId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.Schemas.Update(Id, SchemaId, request) -> *polytomic.BulkSchemaEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bulksync.UpdateBulkSchema{}
client.BulkSync.Schemas.Update(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "contact",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**schemaId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**dataCutoffTimestamp:** `*time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**disableDataCutoff:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**fields:** `[]*polytomic.UpdateBulkField` 
    
</dd>
</dl>

<dl>
<dd>

**filters:** `[]*polytomic.BulkFilter` 
    
</dd>
</dl>

<dl>
<dd>

**partitionKey:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**trackingField:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**userOutputName:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.Schemas.Cancel(Id, SchemaId) -> *polytomic.CancelBulkSyncResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.BulkSync.Schemas.Cancel(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "schema_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The bulk sync ID.
    
</dd>
</dl>

<dl>
<dd>

**schemaId:** `string` — The schema ID to cancel for the bulk sync.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## BulkSync Schedules
<details><summary><code>client.BulkSync.Schedules.List(SyncId) -> *polytomic.SchedulesEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.BulkSync.Schedules.List(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**syncId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.Schedules.Create(SyncId, request) -> *polytomic.ScheduleEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bulksync.CreateScheduleRequest{
        Schedule: &polytomic.V4BulkSyncScheduleApi{
            Frequency: polytomic.ScheduleFrequencyManual,
        },
    }
client.BulkSync.Schedules.Create(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**syncId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**schedule:** `*polytomic.V4BulkSyncScheduleApi` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.Schedules.Get(SyncId, ScheduleId) -> *polytomic.ScheduleEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.BulkSync.Schedules.Get(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**syncId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**scheduleId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.Schedules.Update(SyncId, ScheduleId, request) -> *polytomic.ScheduleEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &bulksync.UpdateScheduleRequest{
        Schedule: &polytomic.V4BulkSyncScheduleApi{
            Frequency: polytomic.ScheduleFrequencyManual,
        },
    }
client.BulkSync.Schedules.Update(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**syncId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**scheduleId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**schedule:** `*polytomic.V4BulkSyncScheduleApi` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkSync.Schedules.Delete(SyncId, ScheduleId) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.BulkSync.Schedules.Delete(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**syncId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**scheduleId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ModelSync Targets
<details><summary><code>client.ModelSync.Targets.GetTarget(Id) -> *polytomic.GetConnectionMetaEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &modelsync.TargetsGetTargetRequest{
        Type: polytomic.String(
            "type",
        ),
        Search: polytomic.String(
            "search",
        ),
    }
client.ModelSync.Targets.GetTarget(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**search:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ModelSync.Targets.GetTargetFields(Id) -> *polytomic.TargetResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &modelsync.TargetsGetTargetFieldsRequest{
        Target: "database.table",
        Refresh: polytomic.Bool(
            false,
        ),
    }
client.ModelSync.Targets.GetTargetFields(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**target:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**refresh:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ModelSync.Targets.List(Id) -> *polytomic.V4TargetObjectsResponseEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns available model sync destinations for a connection.

If the connection supports creating new destinations, the `target_creation`
object will contain information on what properties are required to create the
target.

Target creation properties are all string values; the `enum` flag indicates if
the property has a fixed set of valid values. When `enum` is `true`, the [Target
Creation Property
Values](https://apidocs.polytomic.com/2024-02-08/api-reference/model-sync/targets/get-create-property)
endpoint can be used to retrieve the valid values.

## Sync modes

The sync mode determines which records are written to the destination for a
model sync. The `modes` array for a target object defines the `id` along with
what operations the mode supports.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ModelSync.Targets.List(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ModelSync.Targets.GetCreateProperty(Id, Property) -> *polytomic.V4TargetPropertyValuesEnvelope</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Connections which support creating new sync target objects (destinations) will
return `target_creation` with their [target object list](./list). This endpoint
will return possible values for properties where `enum` is `true`.

If the connection does not support creating new target objects, an HTTP 404 will
be returned.

The `values` array lists the valid options (and labels) for the property. Each
member of the `values` array has a `label` and `value`. For exaample,

```json
{
  "data": [
    {
      "id": "account",
      "title": "Account ID",
      "enum": true,
      "values": [
        {
          "value": "1234567::urn:li:organization:987654",
          "label": "Polytomic Inc. (1234567)"
        }
      ]
    }
  ]
}
```

The `value` for the selected option should be passed when [creating a
sync](https://apidocs.polytomic.com/2024-02-08/api-reference/model-sync/create).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ModelSync.Targets.GetCreateProperty(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "property",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**property:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ModelSync Executions
<details><summary><code>client.ModelSync.Executions.List(SyncId) -> *polytomic.ListExecutionResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &modelsync.ExecutionsListRequest{
        PageToken: polytomic.String(
            "AmkYh8v0jR5B3kls2Qcc9y8MjrPmvR4CvaK7H0F4rEwqvg76K==",
        ),
        OnlyCompleted: polytomic.Bool(
            true,
        ),
        Ascending: polytomic.Bool(
            true,
        ),
    }
client.ModelSync.Executions.List(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**syncId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**pageToken:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**onlyCompleted:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**ascending:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ModelSync.Executions.Get(SyncId, Id) -> *polytomic.GetExecutionResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ModelSync.Executions.Get(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**syncId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ModelSync.Executions.Update(SyncId, Id, request) -> *polytomic.GetExecutionResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &modelsync.UpdateExecutionRequest{
        Status: polytomic.ExecutionStatusCreated,
    }
client.ModelSync.Executions.Update(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**syncId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` — The ID of the execution to update.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*polytomic.ExecutionStatus` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ModelSync.Executions.GetLogUrls(SyncId, Id, Type) -> *polytomic.ExecutionLogsResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ModelSync.Executions.GetLogUrls(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        polytomic.V2ExecutionLogTypeRecords.Ptr(),
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**syncId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*polytomic.V2ExecutionLogType` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ModelSync.Executions.GetLogs(SyncId, Id, Type, Filename) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ModelSync.Executions.GetLogs(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        "0ecd09c1-b901-4d27-9053-f0367c427254",
        polytomic.V2ExecutionLogTypeRecords.Ptr(),
        "path/to/file.json",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**syncId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*polytomic.V2ExecutionLogType` 
    
</dd>
</dl>

<dl>
<dd>

**filename:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Permissions Policies
<details><summary><code>client.Permissions.Policies.List() -> *polytomic.ListPoliciesResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Permissions.Policies.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Permissions.Policies.Create(request) -> *polytomic.PolicyResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &permissions.CreatePolicyRequest{
        Name: "Custom",
    }
client.Permissions.Policies.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**organizationId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**policyActions:** `[]*polytomic.PolicyAction` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Permissions.Policies.Get(Id) -> *polytomic.PolicyResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Permissions.Policies.Get(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Permissions.Policies.Update(Id, request) -> *polytomic.PolicyResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &permissions.UpdatePolicyRequest{
        Name: "Custom",
    }
client.Permissions.Policies.Update(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**organizationId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**policyActions:** `[]*polytomic.PolicyAction` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Permissions.Policies.Remove(Id) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Permissions.Policies.Remove(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Permissions Roles
<details><summary><code>client.Permissions.Roles.List() -> *polytomic.RoleListResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Permissions.Roles.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Permissions.Roles.Create(request) -> *polytomic.RoleResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &permissions.CreateRoleRequest{
        Name: "Custom",
    }
client.Permissions.Roles.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**organizationId:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Permissions.Roles.Get(Id) -> *polytomic.RoleResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Permissions.Roles.Get(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Permissions.Roles.Update(Id, request) -> *polytomic.RoleResponseEnvelope</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &permissions.UpdateRoleRequest{
        Name: "Custom",
    }
client.Permissions.Roles.Update(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**organizationId:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Permissions.Roles.Remove(Id) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Permissions.Roles.Remove(
        context.TODO(),
        "248df4b7-aa70-47b8-a036-33ac447e668d",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>


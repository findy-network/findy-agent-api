# API Message types

- [API Message types](#api-message-types)
  - [Requests](#requests)
    - [Protocol starters](#protocol-starters)
      - [Connect](#connect)
      - [Basic message](#basic-message)
      - [Issue credential](#issue-credential)
      - [Request proof](#request-proof)
    - [Protocol continuators](#protocol-continuators)
      - [Continue credential](#continue-credential)
      - [Continue proof](#continue-proof)
      - [Continue verify proof](#continue-verify-proof)
    - [Other requests](#other-requests)
      - [Register edge endpoint](#register-edge-endpoint)
      - [Generate connection invitation](#generate-connection-invitation)
      - [Create schema](#create-schema)
      - [Create credential definition](#create-credential-definition)
      - [List available tasks](#list-available-tasks)
      - [Query task status](#query-task-status)
  - [Notifications](#notifications)
    - [Task status](#task-status)
    - [User action](#user-action)
    - [Verify proof](#verify-proof)

## Requests

### Protocol starters

All protocol starters return payload with task id. This id can be used afterwards to query task details with task status request.

Response:
```json
{
  "id": <task_id>,
}
```

Response example:

```json
{
  "id": "92396557-2ea8-4dd5-88dd-c3c3f12f37dc"
}
```

#### Connect

Request:
```json
{
  "info": <our_label>,
  "invitation": <aries_invitation>
}
```

Request example:

```json
{
  "info": "api_test_integration",
  "invitation": {
    "@id": "92396557-2ea8-4dd5-88dd-c3c3f12f37dc",
    "@type": "did:sov:BzCbsNYhMrjHiqZDTUASHg;spec/connections/1.0/invitation",
    "label": "api_test_integration_edge",
    "recipientKeys": [
      "Hsmne5a5sWCixojpmtBPZyiTKHSWuFBMxiL88CZZRPpr"
    ],
    "serviceEndpoint": "http://findy-agent:8080/a2a/XxoR86t8Kgbdt2X9vo2zZx/XxoR86t8Kgbdt2X9vo2zZx/VGSj1kwsdAX1a3eBeyK7hk"
  }
}
```


#### Basic message

Request:
```json
{
  "name": <connection_id>,
  "info": <message_content>,
}
```

Request example:
```json
{
  "info": "Long live Findy!",
  "name": "92396557-2ea8-4dd5-88dd-c3c3f12f37dc"
}
```


#### Issue credential

Request:
```json
{
  "credAttributes": <list_of_cred_attributes>,
  "credDefId": <cred_def_id>,
  "name": <connection_id>
}
```

Request example:
```json
{
  "credAttributes": [
    {
      "name": "email",
      "value": "miss.moneypenny@mi6.com"
    }
  ],
  "credDefId": "FLzG1vv1SBCgRJciNc8Dvw:3:CL:230:api_test_1590998835283783",
  "name": "73967c51-acbe-463c-889c-313d554210dd"
}
```

#### Request proof

Request:
```json
{
  "name": <connection_id>,
  "proofAttributes": <list_of_proof_attributes>
}
```

Request example:
```json
{
  "name": "2b26ea60-7c92-46a8-9c61-8017c9252273",
  "proofAttributes": [
    {
      "credDefId": "FLzG1vv1SBCgRJciNc8Dvw:3:CL:236:api_test_1590998856800660",
      "name": "email"
    }
  ]
}
```

### Protocol continuators

#### Continue credential

Accept or decline credential offer.

Request:
```json
{
  "id": <task_id>,
  "ready": <accept_or_not>
}
```

Request example:
```json
{
  "id": "08aad7bc-6afb-483d-a4b2-7d74294df204",
  "ready": true
}
```

#### Continue proof

Accept or decline proof request.

Request:
```json
{
  "id": <task_id>,
  "ready": <accept_or_not>
}
```

Request example:

#### Continue verify proof

### Other requests

#### Register edge endpoint

Request:
```json
{
  "rcvr_endp": <URL>,
  "rcvr_key": <verkey>,
}
```

Request example:
```json
{
  "rcvr_endp": "http://localhost:8000",
  "rcvr_key": "23xNKg56LppUTrHmUJFWNXaU2nB2bpbkMjL2CydubZPV"
}
```

Response:
```json
{}
```

#### Generate connection invitation

Request:
```json
{
  "info": <our_label>
}
```

Request example:
```json
{
  "info": "api_test_integration"
}
```

Response:
```json
{
  "id": <task_id>,
  "invitation": <aries_invitation>
}
```

Response example:
```json
{
  "id": "8ac93ecd-6c79-428c-bb95-7108ea8eee94",
  "invitation": {
    "@id": "8ac93ecd-6c79-428c-bb95-7108ea8eee94",
    "@type": "did:sov:BzCbsNYhMrjHiqZDTUASHg;spec/connections/1.0/invitation",
    "label": "api_test_integration",
    "recipientKeys": [
      "E9pGiTt9DokYWzix5WnxQi5TzPn8VvKqtjoKnfrbpkxR"
    ],
    "serviceEndpoint": "http://findy-agent:8080/a2a/R8aFZgV8iKNZXERPkXH4Fe/R8aFZgV8iKNZXERPkXH4Fe/2vqWVgkUXDB2dvXn2W5wYc"
  }
}
```

#### Create schema

Request:
```json
{
  "schema": {
    "attrs": <attribute_list>,
    "name": <schema_name>,
    "version": <schema_version>,
  }
}
```

Request example:
```json
{
  "schema": {
    "attrs": [
      "email"
    ],
    "name": "Email_1590998670580903",
    "version": "1.0"
  }
}
```
Response:
```json
{
  "schema": {
    "attrs": <attribute_list>,
    "name": <schema_name>,
    "id": <schema_id>,
    "version": <schema_version>,
  }
}
```

Response example:
```json
{
  "schema": {
    "attrs": [
      "email"
    ],
    "id": "FLzG1vv1SBCgRJciNc8Dvw:2:Email_1590999925230172:1.0",
    "name": "Email_1590999925230172",
    "version": "1.0"
  }
}
```


#### Create credential definition

Request:
```json
{
  "info": <cred_def_tag>,
  "schema": {
    "id": <schema_id>
  }
}
```

Request example:
```json
{
  "info": "api_test_1590998801028093",
  "schema": {
    "id": "FLzG1vv1SBCgRJciNc8Dvw:2:Email_1590998800910987:1.0"
  }
}
```

Response:
```json
{
  "id": <cred_def_id>
}
```

Response example:
```json
{
  "id": "FLzG1vv1SBCgRJciNc8Dvw:3:CL:242:api_test_1590999925372799"
}
```

#### List available tasks

Request:
```json
{
  "info": <mobile_device_notification_token>,
  "timestamp": <timestamp_for_filtering>
}
```

Request example:
```json
{
  "info": "device_token",
  "timestamp": 1590998797748
}
```

Response:
```json
{
  "body": <list_of_task_ids>
}
```

Response example:
```json
{
  "body": [
    "47600ac9-3514-4385-a80a-7feca58feee9",
    "5aec9132-a1db-4780-84a7-5a9f7641695f"
  ]
}
```

#### Query task status

See full payload example in [notifications](#task-status) section.

Request:
```json
{
  "id": <task_id>
}
```

Request example:
```json
{
  "id": "812cf4da-a105-429c-9028-90e1629abc24"
}
```

Response:
```json
{
  "body": <task_status_payload>
}
```

Response example:
```json
{
  "body": {
    "id": "89638403-616b-40e9-ac60-7c04c9ca2983",
    "name": "",
    "payload": null,
    "pendingUserAction": false,
    "status": "waiting",
    "timestamp": 1590999918093,
    "type": "connections"
  }
}
```


## Notifications

### Task status

Payload:
```json
{
  "body": {
    "id": <task_id>,
    "name": <connection_id>,
    "payload": <task_data>,
    "pendingUserAction": <user_action_indicator>,
    "status": <status>,
    "timestamp": <timestamp_ms>,
    "type": <protocol_type>
  }
}
```

Payload example:
```json
{
  "body": {
    "id": "e834ec92-5570-48ea-99f0-618e9a384f07",
    "name": "e834ec92-5570-48ea-99f0-618e9a384f07",
    "payload": {
      "myDid": "JKvuyoijqVafto2skVapK6",
      "name": "e834ec92-5570-48ea-99f0-618e9a384f07",
      "theirDid": "ByszCmhr54Ka1xPJhj8RHt",
      "theirEndpoint": "http://findy-agent:8080/a2a/XxoR86t8Kgbdt2X9vo2zZx/XxoR86t8Kgbdt2X9vo2zZx/ByszCmhr54Ka1xPJhj8RHt",
      "theirLabel": "api_test_integration_edge"
    },
    "pendingUserAction": false,
    "status": "ready",
    "timestamp": 1590999967698,
    "type": "connections"
  }
}
```

Task payload varies depending on the protocol type.

`connections` payload:
```json
{
  "myDid": <my_did_in_this_pairwise>,
  "name": <connection_id>,
  "theirDid": <their_did_in_this_pairwise>,
  "theirEndpoint": <their_endpoint>,
  "theirLabel": <their_label>
}
```
`connections` example:
```json
{
  "myDid": "JKvuyoijqVafto2skVapK6",
  "name": "e834ec92-5570-48ea-99f0-618e9a384f07",
  "theirDid": "ByszCmhr54Ka1xPJhj8RHt",
  "theirEndpoint": "http://findy-agent:8080/a2a/XxoR86t8Kgbdt2X9vo2zZx/XxoR86t8Kgbdt2X9vo2zZx/ByszCmhr54Ka1xPJhj8RHt",
  "theirLabel": "api_test_integration_edge"
}
```

`issue-credential` payload:
```json
{
  "attributes": <cred_attributes>,
  "credDefId": <cred_def_id>,
  "schemaId": <schema_id>
}
```
`issue-credential` example:
```json
{
  "attributes": [
    {
      "name": "email",
      "value": "miss.moneypenny@mi6.com"
    }
  ],
  "credDefId": "FLzG1vv1SBCgRJciNc8Dvw:3:CL:251:api_test_1590999970693227",
  "schemaId": "FLzG1vv1SBCgRJciNc8Dvw:2:Email_1590999968042980:1.0"
}
```

`basicmessage` payload:
```json
{
  "delivered": <delivered_status>,
  "message": <message_content>,
  "pairwise": <connection_id>,
  "sentByMe": <sent_or_received>
}
```
`basicmessage` example:
```json
{
  "delivered": true,
  "message": "And prosper!",
  "pairwise": "89638403-616b-40e9-ac60-7c04c9ca2983",
  "sentByMe": false
}
```

`present-proof` payload:
```json
{
  "attributes": <proof_attributes>
}
```
`present-proof` example:
```json
{
  "attributes": [
    {
      "credDefId": "FLzG1vv1SBCgRJciNc8Dvw:3:CL:251:api_test_1590999970693227",
      "name": "email"
    }
  ]
}
```


### User action

User action message has similar payload as [task status](#task-status). In this case `pendingUserAction`field is true.

### Verify proof
Payload:
```json
{
  "nonce": <task_id>,
  "proofValues": <proof_values>
}
```

Example: 
```json
{
  "nonce": "fdded2ba-4fd0-4241-9d8a-78ca00b888c3",
  "proofValues": [{
    "credDefId": "FLzG1vv1SBCgRJciNc8Dvw:3:CL:25:api_test_1591008590864548",
    "name": "email",
    "value": "miss.moneypenny@mi6.com"
  }]
}
```

# API description

## Message structure

The protocol message is wrapped to an encrypted JSON envelope. The actual JSON message is parsed based on the type string.

![Message structure](./protocol_message.png?raw=true 'Protocol message')

## Process

When sending requests or responding to received messages, the actual API message is first encrypted and then wrapped to encrypted envelope. `"application/x-binary"`should be used as the HTTP request Content-Type. 

![Send message](./send_message.png?raw=true 'Send message')

When messages are received, first is decrypted the envelope. After that the message itself can be decrypted and parsed.

![Receive message](./receive_message.png?raw=true 'Receive message')

## Messages

Description of the API messages is found below. Message payload and example can be found clicking message id.

### Common API requests

| ID                                                                       | Type                                      | Purpose                                                                                          |
| ------------------------------------------------------------------------ | ----------------------------------------- | ------------------------------------------------------------------------------------------------ |
| [`AttachEndpoint`](./docs/messages.md#register-edge-endpoint)            | `/attach/1.0/api_endp`                    | Register service HTTP endpoint to agency. Used primarily by service agents with static endpoint. |
| [`ConnectInvitation`](./docs/messages.md#generate-connection-invitation) | `/connections/1.0/invitation`             | Generate connection invitation for other edge agents.                                            |
| [`Connect`](./docs/messages.md#connect)                                  | `/connections/1.0/create`                 | Start pairwise protocol with given connection invitation.                                        |
| [`SendBasicMessage`](./docs/messages.md#basic-message)                   | `/basicmessage/1.0/send`                  | Send basic message to pairwise connection.                                                       |
| [`CredentialOfferAccept`](./docs/messages.md#continue-credential)        | `/protocol/1.0/continue-issue-credential` | Accept or deny received credential offer.                                                        |
| [`ProofRequest`](./docs/messages.md#request-proof)                       | `/present-proof/1.0/request`              | Send proof request to pairwise connection.                                                       |
| [`ProofRequestAccept`](./docs/messages.md#continue-proof)                | `/protocol/1.0/continue-present-proof`    | Accept or deny received proof request.                                                           |
| [`TaskStatus`](./docs/messages.md#query-task-status)                     | `/task/1.0/status`                        | Fetch status for ongoing agency task.                                                            |
| [`TaskList`](./docs/messages.md#list-available-tasks)                    | `/task/1.0/list`                          | Fetch available task ids.                                                                        |

### API requests - service agents

Some functionality is limited only for agents that have a public DID in the ledger.

| ID                                                                 | Type                                | Purpose                                             |
| ------------------------------------------------------------------ | ----------------------------------- | --------------------------------------------------- |
| [`CreateSchema`](./docs/messages.md#create-schema)                 | `/schema/1.0/create`                | Create credential schema and save it to ledger.     |
| [`CreateCredDef`](./docs/messages.md#create-credential-definition) | `/credential_definition/1.0/create` | Create credential definition and save it to ledger. |
| [`CredentialOffer`](./docs/messages.md#issue-credential)           | `/issue-credential/1.0/propose`     | Issue credential to pairwise connection.            |

### Common notifications

| ID                                                         | Type                               | Purpose                            |
| ---------------------------------------------------------- | ---------------------------------- | ---------------------------------- |
| [`Task`](./docs/messages.md#task-status)                   | `/notify/1.0/status`               | Task status notification.          |
| [`UserAction`](./docs/messages.md#user-action)             | `/notify/1.0/user-action`          | User action required notification. |
| [`AcceptValues`](./docs/messages.md#continue-verify-proof) | `/present_proof/1.0/accept_values` | Verify proof notification.         |


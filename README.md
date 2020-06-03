# findy-agent API

This documentation describes the API of how findy-agent clients i.e. edge agents communicate with findy-agent over the network.

**Note! This is the initial beta version of the API and it is likely to change with no backward compatibility in the coming findy-agent versions.**

## Protocol

The API is based on a proprietary protocol over HTTP(S) implemented for findy-agent communication. The protocol is inspired by the early [indy-sdk](https://github.com/hyperledger/indy-sdk) agent communication protocol. Some design decisions are based on legacy functionality that is nowadays obsolete and therefore the protocol is likely to change in the future.

To utilize the API, clients need to onboard to findy-agent. Onboarding to findy-agent can be done either dynamically with message exchange or exporting the edge wallet with findy-agent tools and importing the wallet file to the edge environment. Edge and cloud agents exchange public keys in the onboarding process, and thereafter API communication is encrypted with these keys.

The encryption is implemented with indy-sdk [anon_crypt](https://github.com/hyperledger/indy-sdk/blob/adfdec0ddaee158060f822c8f0810d8f286ae7ac/libindy/include/indy_crypto.h#L251) functionality and therefore indy-sdk is currently a required dependency for all API clients.

## Description

API requests and message packaging are described as more detailed [here](docs/api.md).

The notifications from findy-agent to edge can be received either with a webhook or a socket. The same message protocol is utilized in the notifications.


## Reference implementations

* findy-api-test (rust)
* [findy-agent-cli](https://github.com/findy-network/findy-agent-cli) (golang)
* [findy-issuer-api](https://github.com/findy-network/findy-issuer-api) (node.js)
* [findy-wallet-ios](https://github.com/findy-network/findy-wallet-ios) (swift)

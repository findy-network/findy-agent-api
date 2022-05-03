# Findy Agency Service API

[![test](https://github.com/findy-network/findy-agent-api/actions/workflows/test.yml/badge.svg?branch=dev)](https://github.com/findy-network/findy-agent-api/actions/workflows/test.yml)

## Getting Started

Findy Agency is a collection of services ([Core](https://github.com/findy-network/findy-agent),
[Auth](https://github.com/findy-network/findy-agent-auth),
[Vault](https://github.com/findy-network/findy-agent-vault) and
[Web Wallet](https://github.com/findy-network/findy-wallet-pwa)) that provide
full SSI agency along with a web wallet for individuals.
To start experimenting with Findy Agency we recommend you to start with
[the documentation](https://findy-network.github.io/) and
[set up the agency to your localhost environment](https://github.com/findy-network/findy-wallet-pwa/tree/dev/tools/env#agency-setup-for-local-development).

- [Documentation](https://findy-network.github.io/)
- [Instructions for starting agency in Docker containers](https://github.com/findy-network/findy-wallet-pwa/tree/dev/tools/env#agency-setup-for-local-development)

## Project

The gRPC API v1 includes interfaces for:

- **operation services**: log levels, statistics, health checking, etc.
- **agency services**: new agent on boarding, data hooks, etc.
- **clould agent services**: notifications, responses, etc.
- **protocol services**: starting, running, responses, etc.

More detailed documentation for each of them can be found from `protoc` files in
`idl` directory. The same documentation is available in the target languages
compiled from the IDL files.

## Typical Findy Agency Service Setup

The current implementation of the Findy Agency has three main services:

1. `findy-agent` includes the protocol engine which runs Aries compatible agent
   protocols and cloud agents.
2. `findy-agent-auth` includes a standard WebAuthn server. The
   `findy-agent-auth` repo also includes a headless FIDO2 authenticator for CLIs
   and service agents.
3. `findy-agent-vault` is a service for wallet data. The current API is in
   GraphQL. The future versions probably include gRPC API as well.

## Client Stubs And Helpers

Both [findy-common-go](https://github.com/findy-network/findy-common-go) and
[findy-common-ts](https://github.com/findy-network/findy-common-ts) include
client stubs and some convenient helpers for the use of the API.

## Service Implementations

The most of the services are implemented by the findy core agency in
[`findy-agent`](https://github.com/findy-network/findy-agent).

## Client Reference Implementations

- [findy-common-go](https://github.com/findy-network/findy-common-go) (golang)
- [findy-agent-cli](https://github.com/findy-network/findy-agent-cli) (golang)
- [findy-issuer-tool](https://github.com/findy-network/findy-issuer-tool) (node.js)

## gRPC Style Guide

We follow [the style guide](https://developers.google.com/protocol-buffers/docs/style)
with the following special rules:

- treat acronyms like numbers in the style guide.

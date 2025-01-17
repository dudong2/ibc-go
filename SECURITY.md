# Security

> **IMPORTANT**: If you find a security issue, you can contact our team directly at
security@interchain.berlin, or report it to our [bug bounty program](https://hackerone.com/cosmos) on HackerOne. *DO NOT* open a public issue on the repository.

## Bug Bounty

As part of our [Coordinated Vulnerability Disclosure Policy](https://tendermint.com/security), we operate a
[bug bounty program](https://hackerone.com/cosmos) with Hacker One.

See the policy linked above for more details on submissions and rewards and read
this [blog post](https://blog.cosmos.network/bug-bounty-program-for-tendermint-cosmos-833c67693586) for the program scope.

The following is a list of examples of the kinds of bugs we're most interested
in for the IBC Golang repository. Please refer to the corresponding repositories for vulnerabilities on the [LBM SDK](https://github.com/line/lbm-sdk/blob/main/SECURITY.md) and [Ostracon](https://github.com/line/ostracon/blob/main/SECURITY.md) repositories.

### IBC Core

- [`02-client`](https://github.com/line/ibc-go/tree/main/modules/core/02-client)
- [`03-connection`](https://github.com/line/ibc-go/tree/main/modules/core/03-connection)
- [`04-channel`](https://github.com/line/ibc-go/tree/main/modules/core/04-channel)
- [`05-port`](https://github.com/line/ibc-go/tree/main/modules/core/05-port)
- [`23-commitment`](https://github.com/line/ibc-go/tree/main/modules/core/23-commitment)
- [`24-host`](https://github.com/line/ibc-go/tree/main/modules/core/24-host)

### IBC Applications

- [`transfer`](https://github.com/line/ibc-go/tree/main/modules/apps/transfer)

### Light Clients

- [`99-ostracon`](https://github.com/line/ibc-go/tree/main/modules/light-clients/99-ostracon)

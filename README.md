# CrowdNFT
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme) ![GitHub issues](https://img.shields.io/github/issues/Cameronketchem/CEN3031-Group91)

OpenSea NFT aggregator enabling public pooling/crowdfunding and ownership of NFTs. Users browse NFTs and create crowdfunding campaigns; if the fund manages to outbid other buyers before the auction's expiration, then a new token is launched and distributed according to each user's relative contribution. The token represents a share of ownership (and subsequently a portion of the voting power) of the NFT. Users can trade their ownership tokens on various exchanges, or use them to vote on decisions regarding their NFT.

Group 91 project repository for CEN3031 at the University of Florida.

## Background
NFTs are non-fungible tokens, meaning tokens/contracts on a blockchain network that are distinguishable from one another. They can be analogously compared to baseball cards: each card is individual and its value is not a denomination of some other card, i.e 10 common cards are not the value-equivalent of a rarer card - you can sell those 10 cards to buy the rarer card, but besides that each card has its own unique value. Fungible tokens, on the other hand, all represent the same value or a denomination of said value: e.g if I collect four quarters (coins) then I have the value-equivalent of a dollar - the four quarters and $1 dollar bill have the exact same value and utility.

Blockchain accounts on the Ethereum network (or on other Ethereum forks) can either be controlled by a private key (meaning they are owned by an individual) or via a contract. If a contract is launched for an address, then all activity for that account (including access to all funds) is controlled via programmable functions.

Our project is a web app that serves as a marketplace for NFTs. Except instead of unique bidders competing against each other, our app will allow users to pool their money together and make bids to purchase that NFT communally. Users will send their funds to a contract, and the contract will be programmed to make bids on their behalf. If the bid is won, then that programmable contract will take ownership of the NFT and distribute a token (equivalent to a share in a public company) representing ownership of that NFT. On the other hand, the contract will refund all funds to their respective owners if it fails to win the auction. The project is inspired by an autonomous contract that was used to bid on a [sale of a copy of the US constitution](https://dailyhodl.com/2021/11/21/crypto-enthusiasts-raise-47000000-worth-of-ethereum-in-effort-to-buy-rare-copy-of-us-constitution/). By using decentralized networks, we hope to enable public ownership of assets with no potentially untrustworthy intermediaries; all decisions will be made via democratic voting.

This will be achieved by first fetching NFT auctions (via API calls) from the [OpenSea](https://opensea.io/) marketplace and displaying them on a web interface. The interface will display popular NFTs that are already being crowdfunded, as well as other NFTs which the user can select and start a crowdfunding campaign for. Each user will be represented by a wallet address, which they'll be able to use to verify themselves via browser extensions such as [MetaMask](https://metamask.io/). Most browser wallets can be interfaced via a standardized API called [Web3](https://web3js.readthedocs.io/en/v1.8.1/), we will use this API to request payments from the user, as well as to verify their identity and keep track of NFTs they already own. Users will be able to manage their bids, see and transact with their NFT ownership shares, and launch crowdfunding campaigns.

We'll either be launching our own testing network or use a publicly available testnet to avoid incurring real-life costs. We'll be using GoLang, Solidity, and SQLite in the backend, and AngularJS in the frontend. 

## Install
This project requires [NodeJS](https://nodejs.org/en/) and [Git](https://git-scm.com). Also, make sure you have `build-essential` (Linux) or `xcode-select` (MacOS).
```
sudo apt install build-essential
// or...
xcode-select --install
```

Next, begin by cloning the repository:
```
https://github.com/Cameronketchem/CEN3031-Group91.git
```

Navigate to the project directory and install the project dependencies:
```
cd CEN3031-Group91 && npm run install-dependencies
```

## Usage
<b> Dev environment: </b>
To launch the development environment, run the following in your terminal:
```
npm start
```
The frontend will then be accessible via `http://localhost:4200` and the backend will point to `http://localhost:8080`.

## Maintainers
<b> Frontend: </b><br>
Cameron Ketchem (@Cameronketchem)<br>
Sebastian Sarmiento (@sebassarmiento)

<b> Backend: </b><br>
Mateo Slivka (@MATE023)<br>
Michail Zeipekki (@zeim839)

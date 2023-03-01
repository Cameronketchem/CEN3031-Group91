## Progress we have made on the front-end:
We implemented a feed page that displays NFT's on our website. The NFTs are fetched dynamically from the backend's API using the `api/assets/` route. Additionally, we added a search bar that is not currently operational, but will also utililize the API (in future sprints) to search the databse with specified keywords. There is now a connect wallet button that checks if the user has the metamask extension installed and then connects your wallet to the current session. Additionally, the front-end now has two Cypress specs that will be explained later in the document.

- Implementing API with frontend. See [PR](https://github.com/Cameronketchem/CEN3031-Group91/pull/26) and [ISSUE](https://github.com/Cameronketchem/CEN3031-Group91/issues/14)
- Search bar and displaying NFTs on homepage: [PR](https://github.com/Cameronketchem/CEN3031-Group91/pull/15)
- Connecting frontend with the Metamask browser extension: [Issue](https://github.com/Cameronketchem/CEN3031-Group91/issues/8) and [PR0](https://github.com/Cameronketchem/CEN3031-Group91/pull/11), [PR1](https://github.com/Cameronketchem/CEN3031-Group91/pull/16)
- Display an error when user doesn't have Metamask installed: [PR](https://github.com/Cameronketchem/CEN3031-Group91/pull/18)

## Progress we have made on the back-end:
We implemented a simple API that utilizes GET request to complete functions that include getting a user, asset, or contract by ID and getting a batch of users or assets that ID's immediately follows the ID you provide. In addition to that we implemeted the Executor contract and the ERC721 sale contract using Solidity Contracts. There is also now a blockchain package to interact with these contracts via the blockchain. We also fixed an error in testing our query.

- Creating the preliminary API: [Issue](https://github.com/Cameronketchem/CEN3031-Group91/issues/7) and [PR](https://github.com/Cameronketchem/CEN3031-Group91/pull/12)
  - API routes unit testing: [Issue](https://github.com/Cameronketchem/CEN3031-Group91/issues/13) and [PR](https://github.com/Cameronketchem/CEN3031-Group91/pull/23)
- Fixed an error where the API was not displaying an error when a query did not exist: [PR](https://github.com/Cameronketchem/CEN3031-Group91/pull/9)
- Wrote blockchain smart contracts [Issue](https://github.com/Cameronketchem/CEN3031-Group91/issues/10) and [PR](https://github.com/Cameronketchem/CEN3031-Group91/pull/22)
  - NFT Sale contract: an NFT that also has API calls for setting price and selling the item.
  - Executor contract: an executor is a blockchain account that can own an NFT. This contract can collect contributions and purchase an NFT automatically once the required funds have been collected.

## Front-End Testing:
One of the Cypress test specs tests whether the connect wallet button and the crowd nft button works. The tests will return as passed if clicking on the buttons will cause the pages to be traversed from home to connect wallet and vice versa. The other test spec is used to test the Metamask connect button. It simulates whether the user has the extension or not and checks to make sure that either the alert appears or the Metamask pops up based on the situation.

See [cypress/e2e/connect-wallet.component.cy.js](https://github.com/Cameronketchem/CEN3031-Group91/blob/main/cypress/e2e/connect-wallet.component.cy.js)
and [cypress/e2e/home.component.cy.js](https://github.com/Cameronketchem/CEN3031-Group91/blob/main/cypress/e2e/home.component.cy.js).

## Back-End Testing:
The back end consists of a database package that manages a local SQLite database instance, a datastore package which can modify the database (using the database package), and a routes package that handles user HTTP requests by fetching data from the datastore package. The database and datastore packages were implemented in Sprint 1, and their respective unit tests were committed in [PR#4](https://github.com/Cameronketchem/CEN3031-Group91/pull/4) (these were documented in our Sprint 1 video).

For Sprint 2, we implemented the routes package and set up testing in [PR#23](https://github.com/Cameronketchem/CEN3031-Group91/pull/23). The unit tests can also be accessed via [server/routes/request_test.go](https://github.com/Cameronketchem/CEN3031-Group91/blob/main/server/routes/request_test.go). Our unit tests cover all requests documented in the API documentation (see below). 

## Back-End API Documentation

Our API gives the users of it the ability to get data of users, assets, and contracts with ease. Utilizing our API functions users can get single users or assets or a batch of users or assets.

### ROUTE (GET): `localhost:8080/api/user/USER_ID`
Sending a get request to this route will return either a JSON object for the specified USER_ID parameter, or an error indicating the user doesn't exist. The response data looks like this:
```json
{
    "user_id":1,
    "wallet_pubkey":"0xe452b6b42df7cac9415df4ea4adee5f0c05104c2",
    "profile_pic_url":"https://i.seadn.io/gae/1-b5P_fRbu2sKQblbXTapKrmAi-jBGzQ-eudUYVS9Ncs1nr697WDcVNLuUAiSbqUx1j1ZszjybJ9CB6Lu747xPQsNFBG3CahWL_stdI?auto=format\u0026w=1000",
    "bio":"Lorem ipsum dolor sit amet"
}
```

Error response:
```json
{"error":"No users with that ID"}
```
If a user id which does not exist was queried. 

### ROUTE (GET): `localhost:8080/api/asset/ASSET_ID`
Works the same as above. This route returns data for an asset - an asset is an NFT. The response data looks like this:
```json
{
    "asset_id":1,
    "exec_contract":1,
    "asset_contract":1,
    "img_preview":"https://i.seadn.io/gae/0IY-IVqFKdcTHRARxGayiXz5qu7t38c9HzSerTpN-lEsLA7vj4rxYkK58iVmd3vPFblNcaQqP12Mr9xhUgM1cIsb7ngLl0oIMgQ?auto=format\u0026w=1000",
    "executor":3,
    "desc":"lorem ipsum dolor sit amet",
    "price":250
}
```
Error response:
```json
{"error":"No assets with that ID"}
```
If an asset id which does not exist was queried. 

### ROUTE (GET): `localhost:8080/api/assets/LAST_ID`
This function is for use in the landing page - or for when you want to retrieve and display multiple assets instantaneously. The route sends batches of 20 assets at the time, sorted by their ID, where all IDs are greater than the LAST_ID parameter. So if you specify LAST_ID to be -1, you get the first 20 assets listed in the database. You then specify the last asset ID that you're aware of and you'll retrieve the next 20 assets in the list. This route will return 20 assets at most (but may return less). It may also return an empty array if there are no assets after a certain last known ID. 

Example response:
```json
[
    {
         "asset_id":98,
         "exec_contract":98,
         "asset_contract":98,
         "img_preview":"https://i.seadn.io/gae/22BHI00rhGTKf7EgMIOpCWhNxF7cC1TlAXKslMJksJfxICN4oNWjkMMiXioM6TR6zuyvY65Y84pjLsx3FaQdTCR99cp83axGlVX7ag?auto=format\u0026w=1000",
         "executor":2,
         "desc":"lorem ipsum dolor sit amet",
         "price":250
    },
    {
         "asset_id":99,
         "exec_contract":99,
         "asset_contract":99,
         "img_preview":"https://i.seadn.io/gae/22BHI00rhGTKf7EgMIOpCWhNxF7cC1TlAXKslMJksJfxICN4oNWjkMMiXioM6TR6zuyvY65Y84pjLsx3FaQdTCR99cp83axGlVX7ag?auto=format\u0026w=1000",
         "executor":0,
         "desc":"lorem ipsum dolor sit amet",
         "price":250
    }
]
```

If no users exist past a specified LAST_ID, then the route returns an empty array, like so:
```json
[]
```

### ROUTE (GET): `localhost:8080/api/users/LAST_ID`
Same as above, except for users. The response data is an array of user objects.

Example response:
```json
[
  {
    "user_id":5,
    "wallet_pubkey":"0x2e1d90501c3173367ecc6a409fb1b588bf3c16a5",
    "profile_pic_url":"https://i.seadn.io/gae/aiJ-oha65kuGWR0BNZa-ocK_qW1kHewfbRxdAuPoSaKP7qSXVd4m9Ew1HFPZX3zTkRxyyOp1OV-X6LdiruJWFltcLIwBvRREatQg?auto=format\u0026w=1000",
    "bio":"Lorem ipsum dolor sit amet"
  },
  {
    "user_id":6,
    "wallet_pubkey":"0xd18b57f34e7800aea3bee1a235ac0a1966180fdf",
    "profile_pic_url":"https://i.seadn.io/gae/ae60XvBE8jGGyR8D5hrjIvb8Tftz0BZ6aDRM1EBh2QSfHVaQhUgNITXqJtSXW9Vw74PAKb0kfLmp2sMjkyT4D6fD-XGt4-tLzHpyN2w?auto=format\u0026w=1000",
    "bio":"Lorem ipsum dolor sit amet"
  }
]
```

If no assets exist past a specified LAST_ID, then the route returns an empty array, like so:
```json
[]
```

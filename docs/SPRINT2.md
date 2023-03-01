**Progress we have made on the front-end:** 
We implemented a feed page that displays nft's on our website. For the time being this is hardcoded but as future functionality is introduced we expect to be able to call an api/db to fetch nft's data and display them. Additionally, we added a search bar that is not currently opertationable, but will also utililize the api to search the databse with specified keywords. There is now a connect wallet button that checks if the user has the metamask extension installed and then connects your wallet to the current session. Additionally, the front-end now has two Cypress specs that will be explained later in the document.

**Progress we have made on the back-end:** 
We implemented a simple api that utilizes GET request to complete functions that include getting a user, asset, or contract by ID and getting a batch of users or assets that ID's immediately follows the ID you provide. In addition to that we implemeted the Executor contract and the ERC721 sale contract using Solidity Contracts. There is also now a blockchain package to interact with these contracts via the blockchain. We also fixed an error in testing our query.

**Front-End Testing:**
One of the Cypress test specs tests whether the connect wallet button and the crowd nft button works. The tests will return as passed if clicking on the buttons will cause the pages to be traversed from home to connect wallet and vice versa. The other test spec is used to test the Metamask connect button. It simulates whether the user has the extension or not and checks to make sure that either the alert appears or the Metamask pops up based on the situation.

**Back-End Testing:**
We added testing for our api routes. These test the getUserById, getAssetById, getContractById, getBatchUsers, and getBatchAssets. We also already had implemeted testing for our query functions QueryAssetById, QueryContractById, QueryUserById, QueryContribById, InsertAsset, InsertContract, InsertUser, InserContrib, GetUsers and Get Assets. We also have tests for our Store struct. These include tests for the functions OpenStore, OpenStoreErasable, Init, DropAll and Seed. Finally we have tests for our types and converting them to JSON objects. These include tests for the functions AssetToJSON, UserToJSON, ContractToJSON, ContribToJSON. 

**Back-End API Documentation**

Our API gives the users of it the ability to get data of users, assets, and contracts with ease. Utilizing our API functions users can get single users or assets or a batch of users or assets.

## ROUTE (GET): `localhost:8080/api/user/USER_ID`
Sending a get request to this route will return either a JSON object for the specified USER_ID parameter, or an error indicating the user doesn't exist. The response data looks like this:
```json
{
    "user_id":1,
    "wallet_pubkey":"0xe452b6b42df7cac9415df4ea4adee5f0c05104c2",
    "profile_pic_url":"https://i.seadn.io/gae/1-b5P_fRbu2sKQblbXTapKrmAi-jBGzQ-eudUYVS9Ncs1nr697WDcVNLuUAiSbqUx1j1ZszjybJ9CB6Lu747xPQsNFBG3CahWL_stdI?auto=format\u0026w=1000",
    "bio":"Lorem ipsum dolor sit amet"
}
```

## ROUTE (GET): `localhost:8080/api/asset/ASSET_ID`
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

## ROUTE (GET): `localhost:8080/api/assets/LAST_ID`
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

## ROUTE (GET): `localhost:8080/api/users/LAST_ID`
Same as above, except for users. The response data is an array of users. 
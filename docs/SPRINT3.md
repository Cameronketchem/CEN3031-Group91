## Progress we have made on the front-end:

### Profile creation component
PR: [33](https://github.com/Cameronketchem/CEN3031-Group91/pull/33).

This code creates the profile creation component which facilitates any user that wishes to make an account using their wallet id. The code in account service extracts the nonce and user address to find the users walletID which is then used to query the API to see if the user has an account

<img width="620" alt="Screenshot 2023-03-29 at 22 13 52" src="https://user-images.githubusercontent.com/50573884/228710466-dbf905a1-616d-4d1b-bc27-46fd27172c68.png">

If they don't have an account they are navigated to the profile creation page where they will submit a profile picture and a bio if they would like. If they don't enter values a default bio and picture url will be given to their account. This data is then inputted into an account object which will then be posted to the API.

<img width="623" alt="Screenshot 2023-03-29 at 22 14 31" src="https://user-images.githubusercontent.com/50573884/228710536-04f72148-9193-41e2-ac3d-b7c7fdd4bdd3.png">

Other small changes involved in this commit come from the app component. All the change is if the hasAccount variable in the account service is true then the connect wallet button becomes a view profile button.

<img width="616" alt="Screenshot 2023-03-29 at 22 15 08" src="https://user-images.githubusercontent.com/50573884/228710625-b938f2a1-e725-48cf-bc61-320e0af89a99.png">

### Search bar
PR: [35](https://github.com/Cameronketchem/CEN3031-Group91/pull/35)

A feature for NFTs was made that would allow users to click on a user from the gallery view in order to individually view the NFT and see more information about it.  We also implemented the search bar feature so that it will actually be able to search specific NFTs.  Given an asset id in the search bar the website will navigate to a url specific to the asset.  This would display the name, description and price of the NFT.  If the NFT ID is invalid then the website will return that it is invalid.  

<img width="615" alt="Screenshot 2023-03-29 at 22 15 49" src="https://user-images.githubusercontent.com/50573884/228710886-b2fa76bb-1635-4a00-aef9-b3f5387cadd1.png">

### Load more button
PR: [37](https://github.com/Cameronketchem/CEN3031-Group91/pull/37).

Another feature implemented  in this sprint was added to the home screen component.  In the gallery view of the nfts users are given the ability to extend the view to load more nfts to view.  This feature would be activated with a plus button that the user can click on.

<img width="612" alt="Screenshot 2023-03-29 at 22 17 48" src="https://user-images.githubusercontent.com/50573884/228711006-081b14d8-cc4d-48ff-bdd2-26b12c1823bc.png">

Unit Tests:
Here is an example of some of the unit tests that are generated automatically with each component.

<img width="530" alt="Screenshot 2023-03-29 at 22 18 20" src="https://user-images.githubusercontent.com/50573884/228711062-b8cb0319-fd11-4778-8028-0793bcad9feb.png">

This checks that the variable app.title will = CrowdNFT.  This app.title variable will always be displayed in the top left corner.

<img width="517" alt="image" src="https://user-images.githubusercontent.com/50573884/228711126-3434fffe-ec0d-48fc-9a75-92b989ab2df2.png">

## Progress we have made on the back-end:

This sprint, we finalized the backend database schemas, extended our API to return related 
tables by querying foreign keys, and set up the authentication mechanism for the project.

With respect to the database schemas, we added a nonce and wallet address field to the USERS
table, for the purpose of authentication users (the user needs to sign their address with a 
private key in order to sign up). This was done in [PR30](https://github.com/Cameronketchem/CEN3031-Group91/pull/30).

Next, a new Go package was created called 'auth' to track basic utility functions that were
needed for authenticating users - namely createJWT and VerifySignature. These are detailed
in [PR30](https://github.com/Cameronketchem/CEN3031-Group91/pull/30). Additionally, tests were
created in `server/auth/auth_test.go` that test verifying a JWT and verifying a signature. 

Then, this package was used in the `server/routes/auth.go` file to write the `getNonce`,
`postSignUp`, and `postLogIn` route callbacks. The authentication process begins with the
getNonce route, which checks the database to see whether the requested user exists; 
if he does, then the nonce is returned (and an error if he doesn't). The postSignup
route manages signing the user up. It first checks whether a user with the specified address
exists, then verifies the sent signature to ensure that the user is who he says he is (if he
can sign a message with his private key, then we know he owns his public key, i.e his wallet
address). Finally, the signIn route expects a signed nonce (the nonce is retrieved from
`getNonce`) which, if present, is verified. Then, a JWT is assigned to the user via an
http authentication header. So long as the user carries this header, then we'll know that he's
been authenticated by the backend. This is all likewise implemented in [PR30](https://github.com/Cameronketchem/CEN3031-Group91/pull/30).

Finally, we expanded the response for asset objects called from our api. Before it would just call the asset object with integers in place of other objects that the asset contains. Now the api call includes contract objects for the ExecContract and AssetContract fields and a user object for the Executor field.
For example before when calling /api/asset/ID we get the following:
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

Now it returns a field looking like this
```
{
	"asset_contract":
	{
    	"contract_id":26,
    	"addr":"0x123456",
    	"is_exec":true,
    	"is_finished":false
	},
   "asset_id":26,
	"desc":"lorem ipsum dolor sit amet",
	"exec_contract":
	{
    	"contract_id":26,
    	"addr":"0x123456",
    	"is_exec":true,
    	"is_finished":false
	},
	"executor":
	{
    	"user_id":0,
    	"wallet_pubkey":"",
    	"profile_pic_url":"",
    	"bio":""
	},
	"img_preview":""https://i.seadn.io/gcs/files/842aa976f39848327cd9046baf532d44.png?auto=format&w=1000",
	"price":250,
}
```
This works in a similar manner when calling /api/assets/lastID, the assets will return with the objects within them. This was added in [PR
Unit tests are also adjusted for these two api routes to ensure it is working correctly.

## Back-End API Unit Tests
The back end consists of a database package that manages a local SQLite database instance, a datastore package which can modify the database (using the database package), and a routes package that handles user HTTP requests by fetching data from the datastore package. The database and datastore packages were implemented in Sprint 1, and their respective unit tests were committed in [PR#4](https://github.com/Cameronketchem/CEN3031-Group91/pull/4) (these were documented in our Sprint 1 video).

For Sprint 2, we implemented the routes package and set up testing in [PR#23](https://github.com/Cameronketchem/CEN3031-Group91/pull/23). The unit tests can also be accessed via [server/routes/request_test.go](https://github.com/Cameronketchem/CEN3031-Group91/blob/main/server/routes/request_test.go). 

In Sprint 3 we changed api/asset/ID and api/assets/lastID unit tests in [PR#31](https://github.com/Cameronketchem/CEN3031-Group91/pull/31). This will check the newly formatted api response for objects within the assets.The can be accessed at the same place as before via [server/routes/request_test.go](https://github.com/Cameronketchem/CEN3031-Group91/blob/main/server/routes/request_test.go). Additionally, testing for the authentication mechanism was implemented in `server/auth/auth_test.go` as part of [PR30](https://github.com/Cameronketchem/CEN3031-Group91/pull/30). Our unit tests cover all requests documented in the API documentation (see below). 


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
"asset_contract":
{
"contract_id":1,
"addr":"0x123456",
"is_exec":true,
"is_finished":false
},
"asset_id":1,
"desc":"lorem ipsum dolor sit amet",
"exec_contract":
{
"contract_id":1,
“addr":"0x123456",
"is_exec":true,
"is_finished":false
},
"executor":
{
"user_id":3,
"addr":"0xf3a50c34b5182d0f1f28d7c066f2d7ecfbfdaa47",
"nonce":"12345",
"profile_pic_url":"https://i.seadn.io/gae/4x8YT2iLCCVdHcHJ-_ekgS80_ahT3WxmVceqvhVL5aFLIeCRHMa0vPb3XIdhN7Vpr80XEDnLUPbmUaQ6Y75ddgBQIm5YoHO8MYzDLg?auto=format\u0026w=1000",
"bio":"Lorem ipsum dolor sit amet"
},
"img_preview":"https://i.seadn.io/gae/0IY-IVqFKdcTHRARxGayiXz5qu7t38c9HzSerTpN-lEsLA7vj4rxYkK58iVmd3vPFblNcaQqP12Mr9xhUgM1cIsb7ngLl0oIMgQ?auto=format\u0026w=1000",
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
{
"asset_contract":
{
"contract_id":98,
"addr":"0x123456",
"is_exec":true,
"is_finished":false
},
"asset_id":98,
"desc":"lorem ipsum dolor sit amet",
"exec_contract":
{
"contract_id":98,
"addr":"0x123456"
,"is_exec":true,
"Is_finished":false
},
"executor":
{
"user_id":2,
"Addr":"0xa3e5ad28a494d23453d9a44a2e5081ccf4c430d6",
"Nonce":"12345",
"profile_pic_url":"https://i.seadn.io/gae/JflFtdPgS415-OHBMY9Uixo6CDOm3mTjH6AYpV9IjddtIwyxtaqz0kLnP7RotRtFX5SwU1h1pHlZ71fA2WO081hVgeu2DLsoU_-A?auto=format\u0026w=1000",
"bio":"Lorem ipsum dolor sit amet"
},
"img_preview":"https://i.seadn.io/gae/22BHI00rhGTKf7EgMIOpCWhNxF7cC1TlAXKslMJksJfxICN4oNWjkMMiXioM6TR6zuyvY65Y84pjLsx3FaQdTCR99cp83axGlVX7ag?auto=format\u0026w=1000",
"Price":250
},
{
"asset_contract":
{
"contract_id":99
,"addr":"0x123456"
,"is_exec":true
,"is_finished":false
},
"asset_id":99,
"desc":"lorem ipsum dolor sit amet",
"exec_contract":
{
"contract_id":99,
"addr":"0x123456",
"is_exec":true,
"is_finished":false
},
"executor":
{
"user_id":0,
"addr":"0xbe14eb1ffca54861d3081560110a45f4a1a9e9c5",
"nonce":"12345",
"profile_pic_url":"https://i.seadn.io/gae/AL_lAfiHmqpFk72vHLLTEO8zuRWWYsqRH2uiCld3UuCo8ZAqTah5PwwhtmfFYvlGlTmLh9M7blNUK8kUqzXGN-Lk4hRtUPqAVL9W?auto=format\u0026w=1000",
"bio":"Lorem ipsum dolor sit amet"
},
"img_preview":"https://i.seadn.io/gae/22BHI00rhGTKf7EgMIOpCWhNxF7cC1TlAXKslMJksJfxICN4oNWjkMMiXioM6TR6zuyvY65Y84pjLsx3FaQdTCR99cp83axGlVX7ag?auto=format\u0026w=1000",
"Price":250
}
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
```

If no assets exist past a specified LAST_ID, then the route returns an empty array, like so:
```json
[]
```


# Front End:

## Features added:

### PR:[48](https://github.com/Cameronketchem/CEN3031-Group91/pull/48)
In this Sprint our main implementation involved adding more improvements 
to the home component of our application.  We originally had laid out a 
search bar in the previous sprint that acted more as a placeholder but we 
have now implemented code that will allow users to both search for 
specific NFTs and any other users registered with the website.  Using this 
search feature will redirect users to individual pages of that NFT or 
user.  Below is a snippet of what the html code of this feature looks 
like.

<img width="622" alt="image" 
src="https://user-images.githubusercontent.com/50573884/233229052-53d2d846-d3b4-4d1d-a2b4-63f818425ded.png">

As you can see in this html we use a dropdown menu to select between 
searching for NFTs and searching for users.  
### 
PR:[50](https://github.com/Cameronketchem/CEN3031-Group91/pull/50/files)

<img width="615" alt="image" 
src="https://user-images.githubusercontent.com/50573884/233229138-971e35f0-df28-471d-8b25-28f863d519ad.png">

This is the Javascript code for when a user clicks search.  First the 
value of the search category is checked and subsequently the user is 
redirected to a nftcard page or a user page with the given ID.  It also 
checks to make sure the user has actually entered a value or else the 
function will not be called ensuring the user isn’t sent to a blank page. 

We also implemented a feature so that when users are redirected to either 
the nft view page or profile view page the specific information of that 
entity’s ID is given.  If the ID is invalid such that the ID does not 
exist in the database a message will be displayed informing the user. Here 
is the code snippet for the user page. The two pages are coded very 
similarly so it does not seem relevant to display both.

<img width="615" alt="image" 
src="https://user-images.githubusercontent.com/50573884/233229228-50499612-bc38-403a-a1d0-92ce498b4ab9.png">

In order to verify the ID is valid within our database a call is made to 
our API searching for the specific user ID that was provided.  We then 
check if data was given back to us and if so that data is pushed into a 
userData object that is displayed in the html file.

<img width="611" alt="image" 
src="https://user-images.githubusercontent.com/50573884/233229273-6ed2fbd7-c7d5-4ca7-bed2-8914fe5213ed.png">

### PR:[56](https://github.com/Cameronketchem/CEN3031-Group91/pull/56)
Now shifting back to the home component page in the last sprint we had 
implemented a gallery view that allowed users to view multiple NFTs on one 
page with the ability to load more with the click of a button.  This has 
now been updated to include a user gallery view.  Similar to how the 
search function allowed a choice between NFTs and Users it is also 
possible to see all the users on the site on one page.  Both display their 
own sets of data for the user to view with the ability to click on the 
specific card to be redirected directly to the entity’s individual page.  
The code showcasing this is below.

<img width="625" alt="image" 
src="https://user-images.githubusercontent.com/50573884/233229313-0bc26482-37e0-418c-9861-bb820d28bb48.png">

This first snippet of code is the implementation of the user picking to 
search for users or NFTs.

<img width="622" alt="image" 
src="https://user-images.githubusercontent.com/50573884/233229367-17d81937-291b-4c8d-bb66-eb7635c5b3cb.png">

This snippet shows the two separate feed categories for the gallery view.  
The bottom line of code allows the user to view more entities if 
available.  If there are no longer any NFTs or users within the database a 
message will display informing the user.  

<img width="327" alt="image" 
src="https://user-images.githubusercontent.com/50573884/233229409-b9200797-7a89-4b33-b08e-5c60c993b89a.png">

This is code taken from the home component’s typescript file and is called 
after an api call to get more data.  We use a try catch here so that if we 
receive an error we can assume it is due to the fact that there are no 
more entities to load.  This will make the noMoreData variable true and in 
turn make the message informing the user that they cannot load any more 
entities.

### PR[57](https://github.com/Cameronketchem/CEN3031-Group91/pull/57)
A feature we thought would be helpful was implementation of sorting for 
the gallery view.  The options to sort vary from high to low and include 
the option to sort by the attributes price and ID for NFTs and just id for 
users.

<img width="624" alt="image" 
src="https://user-images.githubusercontent.com/50573884/233229477-02c190bb-29e7-452f-b2d7-c828e05e71d4.png">

Here is the code for the sorting.  As you can see when the user selects an 
option the onSortByChange function is called which will run a sorting 
algorithm through the database to rearrange all the entities.

<img width="621" alt="image" 
src="https://user-images.githubusercontent.com/50573884/233229516-bcfd1301-c98c-4a68-a8db-bed821486226.png">

This is the sort algorithm for the NFTs.

## Front End Cypress Tests:
### PR[51](https://github.com/Cameronketchem/CEN3031-Group91/pull/51)
Many of the Cypress tests are the same as before with a few additions.

<img width="621" alt="image" 
src="https://user-images.githubusercontent.com/50573884/233229621-965b5803-a65a-46aa-9c6f-16d04c7cf539.png">

This initial test ensures that within the home component page if the 
connect wallet button is pressed the user should be redirected to that 
component's specific url.

<img width="622" alt="image" 
src="https://user-images.githubusercontent.com/50573884/233229643-b348fe1a-e405-41cc-8304-2002f5a9df3d.png">

This next test does the opposite.  Upon clicking the title of the website 
“CrowdNFT” the user should be redirected back to the main page of the app 
and the url is checked to ensure this is so.

<img width="622" alt="image" 
src="https://user-images.githubusercontent.com/50573884/233229670-2ade1005-46e9-4a48-99c0-4e1184c95e08.png">

Next this code is used to test the functionality of the connect wallet 
button on the connect wallet page.  This code essentially simulates the 
existence of ethereum within the window which indicates that the user has 
the Metamask extension downloaded.  Finally, the last step is checking 
that the Metamask sign in pop up appears.

<img width="621" alt="image" 
src="https://user-images.githubusercontent.com/50573884/233229701-ecd17efd-68d4-48cd-b198-dff132106aa4.png">

This test simulates the case where a user does not have the Metamask 
extension installed.  So ethereum is set to undefined and upon clicking 
the connect wallet button a pop up will appear instructing the user to 
install Metamask.

## Front End Unit Tests:
Some unit tests are present within all components since it should be 
ensured that they all have this same functionality so I will cover those 
first and then dive into component specific tests after.  

<img width="454" alt="image" 
src="https://user-images.githubusercontent.com/50573884/233229729-41354b95-cc93-4d33-95cc-c89b1af08f44.png">

This first test present in every component essentially checks that the 
Angular app has been built without any errors.  It visits the 
application’s url and checks that the app actually exists at the URL.

<img width="616" alt="image" 
src="https://user-images.githubusercontent.com/50573884/233229762-f78d5390-1b24-46d3-836e-5285357e5fa9.png">

Next ensuring once again that the app has rendered properly this test 
checks first that the app component contains a tag <aT> and then checks 
that within the confines of this tag the title of the app “CrowdNFT” is 
present.

<img width="614" alt="image" 
src="https://user-images.githubusercontent.com/50573884/233229786-99355360-ec3b-4786-83ad-3f9173e901f2.png">
  
This code checks that the option variable which controls the search 
functionality has a value NFT in order to allow searching of NFT IDs.

<img width="622" alt="image" 
src="https://user-images.githubusercontent.com/50573884/233229835-74984f76-6ffa-47f0-bfea-c0ed11279e79.png">
  
This code fully tests the search functionality by first typing an NFT ID.  
It then selects the NFT option in the search bar and clicks search.  This 
test passes if the user is properly redirected to the corresponding 
nftcard page which in this case is “nftcard/1”.  Something of note for 
both users and NFTs their respective pages are either user/”ID” or 
nftcard/”ID”.

<img width="610" alt="image" 
src="https://user-images.githubusercontent.com/50573884/233229873-4f302eb9-c951-4964-ac76-6a567aba3c4b.png">
  
This code tests the same thing except for the user side.  First entering 
the user ID and then searching it to be redirected to the ID’s user page.  
The following tests were all derived from the home component.  I will now 
move on to the nftcard page.

<img width="617" alt="image" 
src="https://user-images.githubusercontent.com/50573884/233229895-5cf36649-0130-488f-94a9-195e8177c18f.png">
  
This test is used for when a user enters an invalid NFT ID.  In this case 
the ID 1200000.  The test first moves to the window under the URL of the 
nftcard/1200000.  From there it makes an API call to ensure that the ID is 
in fact invalid.  It then checks that the message for not finding an NFT 
is displayed.

### PR[54](https://github.com/Cameronketchem/CEN3031-Group91/pull/54)

<img width="611" alt="image" 
src="https://user-images.githubusercontent.com/50573884/233229931-f6a58b54-989f-41f6-9790-951d4bbe0791.png">
  
This test does the same as before except visiting with a validID.  An API 
call will be made to verify the ID is valid and then the tests checks that 
image data gets displayed for that NFT.
Next the profile view page tests will be shown.

<img width="626" alt="image" 
src="https://user-images.githubusercontent.com/50573884/233229964-71fd7f1d-749f-41b0-9e40-98f5b72b12f6.png">
  
These two tests are very similar to the ones that are used for the NFT 
page.  The only difference being a URL change.  The process is still the 
same: the url with this ID is passed into the application that ID is taken 
and checked with the API and depending on whether it exists in the 
database a error message will appear or the user profile’s data will 
appear.

# Backend:

### Security
We added some additional security measures into the api, starting with 
throttle, or rate limiter, that sets restrictions on the amount of 
requests at a time. In 
PR[46](https://github.com/Cameronketchem/CEN3031-Group91/pull/46) we 
implemented a rate limiter as a middleware function that will only allow 
at most 10 requests per second. If it reaches this limit it will return an 
error displaying “rate limit exceeded”. This was done using the 
“juju/ratelimit” package from github, and specifically using the token 
bucket algorithm. The token bucket algorithm being used represents the 
requests as tokens and fills the bucket in at a constant rate, the bucket 
has a limit, in our case being 10. One request represents one token being 
taken out of the bucket and if all ten are taken before the bucket is 
filled again, it stops handling the requests for the time being. The rate 
limiter’s primary purpose is to prevent distributed denial of service, or 
DDOS, attacks, and helps ensure the backend is working efficiently. 
Additional unit testing for the rate limiter was implemented in the same 
pull request, and will be covered more indepthly in the tests section.

Another security measure implemented in 
PR[53](https://github.com/Cameronketchem/CEN3031-Group91/pull/53) was the 
use of http headers that are concerned with security. 

```go
	api.Router.Use(
   	 RateLimiterMiddleware(limiter),
   	 func(c *gin.Context) {

   		 //Allows site administators to control resources allowed 
to be loaded for a page.
   		 //Helps protect from XSS attacks. Default src-directve 
set for the fetch directives of each resource.
   		 c.Writer.Header().Set("Content-Security-Policy", 
"default-src 'self'")

   		 //Avoid click-jacking attacks by only including a page in 
<frame> that is served by the same place as the page.
   		 c.Writer.Header().Set("X-Frame-Options", "SAMEORIGIN")

   		 //Browser remembers site should only be accessed by HTTPS 
for 1 year.
   		 c.Writer.Header().Set("Strict-Transport-Security", 
"max-age=31536000; includeSubDomains")

   		 c.Next()
   	 })
```

The header, “Content-Security-Policy”, is concerned with determining if a 
dynamic resource is allowed to load not. Cross-site scripting attacks 
happen when malicious code is injected into a software and can be used to 
take users data, and this header helps prevent that. All the resources, 
scripts, fonts, etc., have their fetch directives set to the default so 
they only allow trusted sources.

The header “X-Frame-Options”, prevents click-jacking attacks. A 
click-jacking attack happens when a transparent layer, likely a button, is 
placed on a webpage and affects people who unsuspectingly click on it. 
Using “SAMEORIGIN” a page can only be displayed if all frames, <frame>, 
<iframe>, <embed>, and <object> tagged, are from the same origin of the 
page itself. This helps prevent foreign objects from being inserted into 
the page.

The header “Strict-Transport-Security” forces the site to only be 
accessible by https for a specified amount of time, in this case 
31,536,000 seconds, or 1 year. By configuring this here the HTTP 
connection is more secure against man-in-the-middle attacks, rather than 
just redirecting from HTTP to HTTPS on the server. This is also enforced 
for sub-domains.

### Authenticated routes and interfacing the blockchain

In PR[55](https://github.com/Cameronketchem/CEN3031-Group91/pull/55), we 
used the authentication mechanism implemented in Sprint 3 to create 
authenticated routes. Specifically, we created the `POST: /api/asset` 
route, which would allow a registered and logged-in user to submit a new 
NFT onto the app. A new middleware function was written to check that the 
user is logged-in in `server/routes/auth.go`, which retrieves the request 
header and checks to see if an authentication bearer token is present. The 
token is then verified using the authentication package set up in sprint 
3. 

Additionally, new routes were added to query real-time properties of 
objects in our API. Since blockchain NFTs are decentralized, we could not 
query or store the price or status of an NFT in our database (it can be 
changed outside of the application). We thus added 4 new routes to query 
the asset price, asset status, contributions to the crowdfund from 
specific blockchain addresses, and the total amount pooled in the sale 
fund. This was accomplished by compiling our solidity contracts into 
Golang using the ‘solc’ compiler, creating Go packages for each respective 
contract, and then creating instances of contract “callers” which, when 
given a dialable ethereum node, could fetch data dynamically from the 
blockchain. Here is an example of a route that fetches from the 
blockchain:

```go
// server/routes/request.go:300
func getAssetActive(store *data.Store, cont *blockchain.Executor, c 
*gin.Context) {
	// Get ID as integer from string.
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter 
expected integer"})
		return
	}

	// Verify asset exists.
	asset, err := store.QueryAssetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No assets 
with that ID"})
		return
	}

	// Get asset state.
	assetCnt, err := store.QueryContractById(asset.ExecContract)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot find the asset's assosciated 
contract"})
		return
	}

            // Create a contract caller and query for sale status
	addr := assetCnt.Addr
	contex, err := 
executor.NewExecutorCaller(common.HexToAddress(addr), cont.Client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server cannot reach asset contract"})
		return
	}

            // Query for sale status by calling the SaleActive() func.
	active, err := contex.SaleActive(&bind.CallOpts{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server cannot retrieve asset state"})
		return
	}

	c.JSON(http.StatusOK, active)
}
```

## Back End Tests

As mentioned before, in 
PR[46](https://github.com/Cameronketchem/CEN3031-Group91/pull/46) the rate 
limiter was implemented with unit testing. The unit test is displayed 
below.
```go
func TestRateLimiting(t *testing.T) {
    // Gin router with rate limiter middleware
    r := gin.Default()
    limiter := ratelimit.NewBucketWithQuantum(
   	 time.Second,
   	 10,
   	 10,
    )
    r.Use(RateLimiterMiddleware(limiter))

    // Endpoint handler function that just returns a simple message
    r.GET("/api/user/0", func(c *gin.Context) {
   	 c.JSON(http.StatusOK, gin.H{"message": "Request handled"})
    })

    // Test server using the Gin router
    server := httptest.NewServer(r)
    defer server.Close()

    // Send 11 requests, bucket only should allow 10
    for i := 0; i < 11; i++ {
   	 resp, err := http.Get(server.URL + "/api/user/0")
   	 if err != nil {
   		 t.Fatalf("Failed to send request: %v", err)
   	 }
   	 if resp.StatusCode == http.StatusTooManyRequests {
   		 // Rate limit exceeded, passes test
   		 return
   	 }
   	 if resp.StatusCode != http.StatusOK {
   		 t.Fatalf("Unexpected status code: %d", resp.StatusCode)
   	 }
    }

    // The rate limit was not exceeded so the test fails
    t.Fatalf("Rate limit was not exceeded")
}
```

This test ensures that the rate limiter is enforcing the policy of not 
allowing more than 10 requests per second. It sets up a handler function 
and test server to request from, then iterates 11 times sending one 
request each time. The anticipated outcome is the rate limiter stopping 
the 11th request from being handled.
                       
## Updated API Reference

Our API gives the users of it the ability to get data of users, assets, 
and contracts with ease. Utilizing our API functions users can get single 
users or assets or a batch of users or assets.

### ROUTE (GET): `localhost:8080/api/user/USER_ID`
Sending a get request to this route will return either a JSON object for 
the specified USER_ID parameter, or an error indicating the user doesn't 
exist. The response data looks like this:
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
Works the same as above. This route returns data for an asset - an asset 
is an NFT. The response data looks like this:
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
This function is for use in the landing page - or for when you want to 
retrieve and display multiple assets instantaneously. The route sends 
batches of 20 assets at the time, sorted by their ID, where all IDs are 
greater than the LAST_ID parameter. So if you specify LAST_ID to be -1, 
you get the first 20 assets listed in the database. You then specify the 
last asset ID that you're aware of and you'll retrieve the next 20 assets 
in the list. This route will return 20 assets at most (but may return 
less). It may also return an empty array if there are no assets after a 
certain last known ID. 

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

If no users exist past a specified LAST_ID, then the route returns an 
empty array, like so:
```json
[]
```

### ROUTE (GET): `localhost:8080/api/users/LAST_ID`
Same as above, except for users. The response data is an array of user 
objects.

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

If no assets exist past a specified LAST_ID, then the route returns an 
empty array, like so:
```json
[]
```
      
### ROUTE (GET): `localhost:8080/api/auth/:addr`
Retrieves the nonce of a registered user, where `addr` is the wallet 
address of the user. Example response:
```json
"0g3h18ghj3j1h26"
```
      
If the user doesnt exist, then:
```json
      {"error": "no account with this address"}
```
      
### ROUTE (POST): `localhost:8080/api/auth/signup`
Sign up a new user. Input schema (encoding: json):
```json
{
      "addr": "WALLET ADDRESS", 
      "profile_pic_url": "URL TO PROFILE PIC", 
      "bio": "hello world!", 
      "sig": "signature of above data signed with the private key which 
yields addr"
}
```
      
### ROUTE (POST): `localhost:8080/api/auth/login`
Logs a user into the app, assigning an auth header to future requests 
(headers must be preserved accross API calls). Input schema (encoding: 
json):
```json
{"addr": "WALLET ADDRESS", "sig": "signature of account nonce retrieved 
from /api/auth/addr"}
```
      
### ROUTE (GET): `localhost:8080/api/:id/price`
Retrieves the price of the asset with id `:id`. The response is a float in 
JSON encoding. If the ID does not exist, then an error is issued:
```json
{"error":"No assets with that ID"}
```

### ROUTE (GET): `localhost:8080/api/:id/active`
Retrieves the status of the asset with id `:id`. The response is a bool in 
JSON encoding. If the ID does not exist, then an error is issued:
```json
{"error":"No assets with that ID"}
```
      
### ROUTE (GET): `localhost:8080/api/:id/contrib/:addr`
Retrieves the amount contributed by `:addr` to the sale of `:id`. The 
response is a float in JSON encoding. If the ID or addr does not exist, 
then an error is issued:
```json
{"error":"No assets with that ID"}
```

### ROUTE (GET): `localhost:8080/api/:id/balance`
Retrieves the account balance of the asset with id `:id`. The response is 
a bool in JSON encoding. If the ID does not exist, then an error is 
issued:
```json
{"error":"No assets with that ID"}
```




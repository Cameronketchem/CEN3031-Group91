package data

import (
	"context"
	"database/sql"
	"github.com/Cameronketchem/CEN3031-Group91/server/auth"
	"github.com/Cameronketchem/CEN3031-Group91/server/blockchain"
	"github.com/Cameronketchem/CEN3031-Group91/server/blockchain/erc721sale"
	"github.com/Cameronketchem/CEN3031-Group91/server/blockchain/executor"
	"github.com/Cameronketchem/CEN3031-Group91/server/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"math/big"
	"math/rand"
)

// Seed the database for sprint 1 demo. Contracts and
// assets are not fully functional.

var demoProfilePics = []string{
	"https://i.seadn.io/gae/AL_lAfiHmqpFk72vHLLTEO8zuRWWYsqRH2uiCld3UuCo8ZAqTah5PwwhtmfFYvlGlTmLh9M7blNUK8kUqzXGN-Lk4hRtUPqAVL9W?auto=format&w=1000",
	"https://i.seadn.io/gae/1-b5P_fRbu2sKQblbXTapKrmAi-jBGzQ-eudUYVS9Ncs1nr697WDcVNLuUAiSbqUx1j1ZszjybJ9CB6Lu747xPQsNFBG3CahWL_stdI?auto=format&w=1000",
	"https://i.seadn.io/gae/JflFtdPgS415-OHBMY9Uixo6CDOm3mTjH6AYpV9IjddtIwyxtaqz0kLnP7RotRtFX5SwU1h1pHlZ71fA2WO081hVgeu2DLsoU_-A?auto=format&w=1000",
	"https://i.seadn.io/gae/4x8YT2iLCCVdHcHJ-_ekgS80_ahT3WxmVceqvhVL5aFLIeCRHMa0vPb3XIdhN7Vpr80XEDnLUPbmUaQ6Y75ddgBQIm5YoHO8MYzDLg?auto=format&w=1000",
	"https://i.seadn.io/gae/Tlx4PPgYv6tzNkYL36JTogDPrzHg_ANqWiWoLn54Cl1YnkXZl4OyKa_3JvLgxTTEouXabAiBMX8oV9aJKVdXHRwv3eraY7xAMYgAtxY?auto=format&w=1000",
	"https://i.seadn.io/gae/aiJ-oha65kuGWR0BNZa-ocK_qW1kHewfbRxdAuPoSaKP7qSXVd4m9Ew1HFPZX3zTkRxyyOp1OV-X6LdiruJWFltcLIwBvRREatQg?auto=format&w=1000",
	"https://i.seadn.io/gae/ae60XvBE8jGGyR8D5hrjIvb8Tftz0BZ6aDRM1EBh2QSfHVaQhUgNITXqJtSXW9Vw74PAKb0kfLmp2sMjkyT4D6fD-XGt4-tLzHpyN2w?auto=format&w=1000",
}

// Test ganache wallets
var demoUserWallets = []string{
	"0x04a8Ee6EC6679Ed99c6c93d46c0D1C96615090B1",
	"0x6118E7163185b92B06706A545aDaF096A9F9F6cB",
	"0xb21bB1464112bE56feB1Dcb8216Ff2Dd2fd6bd85",
	"0xf8A2a70fBD171DE55b3ce25Bddd81295932197DF",
	"0xe836E9d3fe7c8BAE7C52839b0aDEA69eF82423c9",
	"0xa3e75AB84902DA606D6a6f95291233886c6b5999",
	"0x569F7E8a48ad2dA6257cEfd5BB58943443a6b3C8",
}

// Source URLs of demo nfts.
var demoNftPics = []string{
	"https://i.seadn.io/gcs/files/842aa976f39848327cd9046baf532d44.png?auto=format&w=1000",
	"https://i.seadn.io/gcs/files/7f91126dfecbcfd2ae66ba82159c1b98.png?auto=format&w=1000",
	"https://i.seadn.io/gcs/files/3042add0271c1379044a013f12970e56.png?auto=format&w=1000",
	"https://i.seadn.io/gcs/files/85d6beea8e73f3e72171322253545c0a.png?auto=format&w=1000",
	"https://i.seadn.io/gcs/files/de87979660003b7122dec5d55696c3d4.png?auto=format&w=1000",
	"https://i.seadn.io/gcs/files/d5dd64b64fd9d57a9906ae848b99e67f.png?auto=format&w=1000",
	"https://i.seadn.io/gcs/files/b5b80e31ae895c3624abe3ce8ad07202.jpg?auto=format&w=1000",
	"https://i.seadn.io/gae/jcC9jRoj6ZcgEf9Oq4KvijkIQQxJaX6nXhbKZZJ1DA8r7k7OZAmsfQaDBEu5oLS3vKAqEchCVHdWwR6FaKxaXFvlTwCC0tyWS86bH_Y?auto=format&w=1000",
	"https://i.seadn.io/gae/22BHI00rhGTKf7EgMIOpCWhNxF7cC1TlAXKslMJksJfxICN4oNWjkMMiXioM6TR6zuyvY65Y84pjLsx3FaQdTCR99cp83axGlVX7ag?auto=format&w=1000",
	"https://i.seadn.io/gae/gVXbl75qhKf1WAfINmsyS22A93N1LbXsIFFc8DjZUgY0aKfcc0F1XY9bmotuLY-G1sg03hFBNkVWKapGhfQingFrmJITziIBEN_XCg?auto=format&w=1000",
	"https://i.seadn.io/gcs/files/4269802df3544ba14d7fb0fef8119651.jpg?auto=format&w=1000",
	"https://i.seadn.io/gae/0IY-IVqFKdcTHRARxGayiXz5qu7t38c9HzSerTpN-lEsLA7vj4rxYkK58iVmd3vPFblNcaQqP12Mr9xhUgM1cIsb7ngLl0oIMgQ?auto=format&w=1000",
	"https://i.seadn.io/gae/v-q10IlTd7LP6dd8Af5TWFPswA0hNDLeag93WsjQXYCR8--Tnm9o3zBYg1_TcQZTCGCAa5ZwSnNSJrlAPOXiDe52zj4n9FiJaKPMLyk?auto=format&w=1000",
}

// Demo seed transactions. These do not reference real contracts yet.
var demoSeedTxs = []string{
	`INSERT INTO users(user_id, addr, nonce, profile_pic_url, bio) VALUES(?, ?, ?, ?, ?)`,
	`INSERT INTO contracts(contract_id, addr, is_exec, is_finished) VALUES(?, ?, ?, ?)`,
	`INSERT INTO assets(asset_id, exec_contract, asset_contract, img_preview, executor, desc, price) VALUES(?, ?, ?, ?, ?, ?, ?)`,
	`INSERT INTO contributions(contrib_id, exec_contract, user, amt, status) VALUES(?, ?, ?, ?, ?)`,
}

func demoSeed(ctx context.Context, tx *sql.Tx) error {
	// Prepare expressions

	userExpr, err := tx.Prepare(demoSeedTxs[0])
	if err != nil {
		return err
	}
	defer userExpr.Close()

	contractExpr, err := tx.Prepare(demoSeedTxs[1])
	if err != nil {
		return err
	}
	defer contractExpr.Close()

	assetExpr, err := tx.Prepare(demoSeedTxs[2])
	if err != nil {
		return err
	}
	defer assetExpr.Close()

	contribExpr, err := tx.Prepare(demoSeedTxs[3])
	if err != nil {
		return err
	}
	defer contribExpr.Close()

	// Insert users
	for i := 0; i < 7; i++ {
		nonce, _ := auth.Nonce(20)
		_, err := userExpr.Exec(i, demoUserWallets[i],
			nonce, demoProfilePics[i], "Lorem Ipsum dolor sit amet")
		if err != nil {
			return err
		}
	}

	// Load .env
	err = godotenv.Load()
	if err != nil {
		return err
	}

	// Create blockchain contract executor.
	bc := blockchain.NewExecutor(utils.GetVar("PRIV_KEY"),
		utils.GetVar("NODE_ADDR"))

	// contract
	ci := 0

	// Insert contracts & assets
	for i := 0; i < 100; i++ {

		// Asset properties
		imgPreview := demoNftPics[rand.Intn(13)]
		exec := rand.Intn(7)
		desc := "Lorem ipsum dolor sit amet"
		price := rand.Intn(300)

		// Convert price into wei integer
		priceflt := big.NewFloat(float64(price))
		factor := big.NewFloat(1000000000000000000) // 10^18
		priceflt.Mul(priceflt, factor)

		priceint := new(big.Int)
		priceflt.Int(priceint)

		// Address needs to be in common.Address
		address := common.HexToAddress(demoUserWallets[exec])

		// Create asset contract.
		assetAddr, _, err := bc.DeployERC721Contract(erc721sale.ERC721SALE,
			"Test", priceint, address)

		if err != nil {
			return err
		}

		_, err = contractExpr.Exec(ci, assetAddr.Hex(), false, false)
		if err != nil {
			return err
		}
		ci++

		// Create executor contract
		execAddr, _, err := bc.DeployERC20ExecutorContract(executor.ERC20EXECUTOR,
			"Test", assetAddr)

		if err != nil {
			return err
		}

		_, err = contractExpr.Exec(ci, execAddr.Hex(), true, false)
		if err != nil {
			return err
		}
		ci++

		// Finalize asset
		_, err = assetExpr.Exec(i, ci-1, ci-2, imgPreview, exec, desc, price)
		if err != nil {
			return err
		}
	}

	// Insert contributions
	for i := 0; i < 100; i++ {
		_, err = contribExpr.Exec(i, rand.Intn(100), rand.Intn(7), 50, 1)
		if err != nil {
			return err
		}
	}

	return nil
}

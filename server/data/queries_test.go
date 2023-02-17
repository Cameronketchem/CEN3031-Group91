package data

import (
	"github.com/Cameronketchem/CEN3031-Group91/server/utils/db"
	"github.com/stretchr/testify/require"
	"testing"
)

func initTestStore(t *testing.T) Store {
	str, err := OpenStore("fn.db", true)
	require.NoError(t, err)

	err = str.Init(db.SynchronousModeExtra, false)
	require.NoError(t, err)

	err = str.Seed()
	require.NoError(t, err)

	return str
}

func TestQueryAssetByID(t *testing.T) {
	str := initTestStore(t)
	defer str.Close()

	for i := 0; i < 100; i++ {
		_, err := str.QueryAssetById(i)
		require.NoError(t, err)
	}

  // Should show error if AssetID doesnt exist
  asset, err := str.QueryAssetById(1000)
  require.Error(t, err)
  require.Equal(t, asset.AssetID, 0)
}

func TestQueryContractByID(t *testing.T) {
	str := initTestStore(t)
	defer str.Close()

	for i := 0; i < 100; i++ {
		_, err := str.QueryContractById(i)
		require.NoError(t, err)
	}

  // Should show error if ContractID doesnt exist
  contract, err := str.QueryContractById(1000)
  require.Error(t, err)
  require.Equal(t, contract.ContractID, 0)
}

func TestQueryUserByID(t *testing.T) {
	str := initTestStore(t)
	defer str.Close()

	for i := 0; i < 7; i++ {
		_, err := str.QueryUserById(i)
		require.NoError(t, err)
	}

  // Should show error if UserID doesnt exist
  user, err := str.QueryUserById(1000)
  require.Error(t, err)
  require.Equal(t, user.UserID, 0)
}

func TestQueryContribByID(t *testing.T) {
	str := initTestStore(t)
	defer str.Close()

	for i := 0; i < 100; i++ {
		_, err := str.QueryContribById(i)
		require.NoError(t, err)
	}

  // Should show error if ContribID doesnt exist
  contrib, err := str.QueryContribById(1000)
  require.Error(t, err)
  require.Equal(t, contrib.ContribID, 0)
}

func TestInsertAsset(t *testing.T) {
	str := initTestStore(t)
	defer str.Close()

	for i := 0; i < 100; i++ {
		asset := Asset{i, i, i, "website.com", i, "lorem ipsum", 253.44}
		err := str.InsertAsset(asset)
		require.NoError(t, err)
	}
}

func TestInsertContract(t *testing.T) {
	str := initTestStore(t)
	defer str.Close()

	for i := 0; i < 100; i++ {
		contract := Contract{i, "0xFFFFFFFFFF", false, false}
		err := str.InsertContract(contract)
		require.NoError(t, err)
	}
}

func TestInsertUser(t *testing.T) {
	str := initTestStore(t)
	defer str.Close()

	for i := 0; i < 100; i++ {
		user := User{i, "0xFFAA99999", "website.com", "lorem ipsum"}
		err := str.InsertUser(user)
		require.NoError(t, err)
	}
}

func TestInsertContrib(t *testing.T) {
	str := initTestStore(t)
	defer str.Close()

	for i := 0; i < 100; i++ {
		contrib := Contribution{i, i, i, 253.4, ContribPending}
		err := str.InsertContrib(contrib)
		require.NoError(t, err)
	}
}

func TestGetUsers(t *testing.T) {
	str := initTestStore(t)
	defer str.Close()

	// Get first 20 users (seed database only has 7)
	users, err := str.GetUsers(-1)
	require.NoError(t, err)
	require.LessOrEqual(t, len(users), 20)
	require.Equal(t, users[0].UserID, 0)

	// Expect no users
	usersB, err := str.GetUsers(20)
	require.NoError(t, err)
	require.Equal(t, len(usersB), 0)
}

func TestGetAssets(t *testing.T) {
	str := initTestStore(t)
	defer str.Close()

	// Get first 20 assets
	assets, err := str.GetAssets(-1)
	require.NoError(t, err)
	require.LessOrEqual(t, len(assets), 20)
	require.Equal(t, assets[0].AssetID, 0)

	// Get 5 more sets of 20
	for i := 1; i < 5; i++ {
		index := i * 20
		assets, err = str.GetAssets(index)
		require.NoError(t, err)
		require.LessOrEqual(t, len(assets), 20)
		require.Equal(t, assets[0].AssetID, index+1)
	}

	// Should return nothing
	assets, err = str.GetAssets(101)
	require.NoError(t, err)
	require.LessOrEqual(t, len(assets), 20)
	require.Equal(t, len(assets), 0)
}

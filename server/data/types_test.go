package data

import "testing"

func TestAssetToJSON(t *testing.T) {
	expect := `{"asset_id":1,"exec_contract":1,"asset_contract":1,"img_preview":"http://website.com","executor":1,"desc":"lorem ipsum","price":150.36}`

	asset := Asset{
		1, 1, 1, "http://website.com", 1, "lorem ipsum", 150.36,
	}

	json, err := asset.ToJSON()
	if err != nil {
		t.Errorf("Asset ToJSON returned serialization error")
	}

	if json != expect {
		t.Errorf("Asset JSON output doesn't match expected output")
	}
}

func TestContractToJSON(t *testing.T) {
	expect := `{"contract_id":1,"addr":"1","is_exec":true,"is_finished":false}`
	contract := Contract{1, "1", true, false}
	json, err := contract.ToJSON()
	if err != nil {
		t.Errorf("Contract ToJSON returned serialization error")
	}

	if json != expect {
		t.Errorf("Contract JSON output doesn't match expected output")
	}
}

func TestUserToJSON(t *testing.T) {
	expect := `{"user_id":1,"wallet_pubkey":"1","profile_pic_url":"1","bio":"1"}`
	user := User{1, "1", "1", "1"}
	json, err := user.ToJSON()
	if err != nil {
		t.Errorf("User ToJSON returned serialization error")
	}

	if json != expect {
		t.Errorf("User JSON output doesn't match expected output")
	}
}

func TestContribtoJSON(t *testing.T) {
	expect := `{"contrib_id":1,"exec_contract":1,"user":1,"amt":150.33,"status":0}`
	contrib := Contribution{1, 1, 1, 150.33, ContribPending}
	json, err := contrib.ToJSON()
	if err != nil {
		t.Errorf("Contrib ToJSON returned serialization error")
	}

	if json != expect {
		t.Errorf("Contrib JSON output doesn't match expected output")
	}
}

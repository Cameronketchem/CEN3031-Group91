package data

// TODO: Queries for updating tables

import (
	"context"
	"database/sql"
)

// Query a table by ID and return its representative type.

func (str *Store) QueryAssetById(id int) (Asset, error) {
	assetctx := context.WithValue(context.Background(), "asset", &Asset{})
	assetctx = context.WithValue(assetctx, "id", id)
	err := str.Handler.Rdb.AtomicContext(assetctx, func(ctx context.Context, tx *sql.Tx) error {
		asset := ctx.Value("asset").(*Asset)
		id := ctx.Value("id").(int)

		err := tx.QueryRow("SELECT * FROM assets WHERE asset_id=? LIMIT 1", id).Scan(&(*asset).AssetID,
			&(*asset).ExecContract, &(*asset).AssetContract, &(*asset).ImgPreview, &(*asset).Executor,
			&(*asset).Desc, &(*asset).Price)

		return err
	})

	if err != nil {
		return Asset{}, err
	}

	return *assetctx.Value("asset").(*Asset), nil
}

func (str *Store) QueryContractById(id int) (Contract, error) {
	cntrctx := context.WithValue(context.Background(), "contract", &Contract{})
	cntrctx = context.WithValue(cntrctx, "id", id)
	err := str.Handler.Rdb.AtomicContext(cntrctx, func(ctx context.Context, tx *sql.Tx) error {
		contr := ctx.Value("contract").(*Contract)
		id := ctx.Value("id").(int)

		err := tx.QueryRow("SELECT * FROM contracts WHERE contract_id=? LIMIT 1", id).Scan(&(*contr).ContractID,
			&(*contr).Addr, &(*contr).IsExec, &(*contr).IsFinished)

		return err
	})

	if err != nil {
		return Contract{}, err
	}

	return *cntrctx.Value("contract").(*Contract), nil
}

func (str *Store) QueryContractByAddr(addr string) (Contract, error) {
	cntrctx := context.WithValue(context.Background(), "contract", &Contract{})
	cntrctx = context.WithValue(cntrctx, "addr", addr)
	err := str.Handler.Rdb.AtomicContext(cntrctx, func(ctx context.Context, tx *sql.Tx) error {
		contr := ctx.Value("contract").(*Contract)
		addr := ctx.Value("addr").(string)

		err := tx.QueryRow("SELECT * FROM contracts WHERE addr=? LIMIT 1", addr).Scan(&(*contr).ContractID,
			&(*contr).Addr, &(*contr).IsExec, &(*contr).IsFinished)

		return err
	})

	if err != nil {
		return Contract{}, err
	}

	return *cntrctx.Value("contract").(*Contract), nil
}

func (str *Store) QueryUserById(id int) (User, error) {
	usrctx := context.WithValue(context.Background(), "user", &User{})
	usrctx = context.WithValue(usrctx, "id", id)
	err := str.Handler.Rdb.AtomicContext(usrctx, func(ctx context.Context, tx *sql.Tx) error {
		usr := ctx.Value("user").(*User)
		id := ctx.Value("id").(int)

		err := tx.QueryRow("SELECT * FROM users WHERE user_id=? LIMIT 1", id).Scan(&(*usr).UserID,
			&(*usr).Addr, &(*usr).Nonce, &(*usr).ProfilePicUrl, &(*usr).Bio)

		return err
	})

	if err != nil {
		return User{}, err
	}

	return *usrctx.Value("user").(*User), nil
}

func (str *Store) QueryUserByAddr(addr string) (User, error) {
	usrctx := context.WithValue(context.Background(), "user", &User{})
	usrctx = context.WithValue(usrctx, "addr", addr)
	err := str.Handler.Rdb.AtomicContext(usrctx, func(ctx context.Context, tx *sql.Tx) error {
		usr := ctx.Value("user").(*User)
		addr := ctx.Value("addr").(string)

		err := tx.QueryRow("SELECT * FROM users WHERE addr=? LIMIT 1", addr).Scan(&(*usr).UserID,
			&(*usr).Addr, &(*usr).Nonce, &(*usr).ProfilePicUrl, &(*usr).Bio)

		return err
	})

	if err != nil {
		return User{}, err
	}

	return *usrctx.Value("user").(*User), nil
}

func (str *Store) QueryContribById(id int) (Contribution, error) {
	contribctx := context.WithValue(context.Background(), "contribution", &Contribution{})
	contribctx = context.WithValue(contribctx, "id", id)
	err := str.Handler.Rdb.AtomicContext(contribctx, func(ctx context.Context, tx *sql.Tx) error {
		contr := ctx.Value("contribution").(*Contribution)
		id := ctx.Value("id").(int)

		err := tx.QueryRow("SELECT * FROM contributions WHERE contrib_id=? LIMIT 1", id).Scan(&(*contr).ContribID,
			&(*contr).ExecContract, &(*contr).User, &(*contr).Amt, &(*contr).Status)

		return err
	})

	if err != nil {
		return Contribution{}, err
	}

	return *contribctx.Value("contribution").(*Contribution), nil
}

// SQL statements for inserting into a table
var insertExpr = []string{
	`INSERT INTO assets(exec_contract, asset_contract, img_preview, executor, desc, price)
   VALUES(?, ?, ?, ?, ?, ?)`,
	`INSERT INTO contracts(addr, is_exec, is_finished)
   VALUES(?, ?, ?)`,
	`INSERT INTO users(addr, nonce, profile_pic_url, bio)
   VALUES(?, ?, ?, ?)`,
	`INSERT INTO contributions(exec_contract, user, amt, status)
   VALUES(?, ?, ?, ?)`,
}

func (str *Store) InsertAsset(asset Asset) error {
	assetctx := context.WithValue(context.Background(), "asset", asset)
	err := str.Handler.Wdb.AtomicContext(assetctx, func(ctx context.Context, tx *sql.Tx) error {
		asset := ctx.Value("asset").(Asset)

		stmt, err := tx.Prepare(insertExpr[0])
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(asset.ExecContract, asset.AssetContract,
			asset.ImgPreview, asset.Executor, asset.Desc, asset.Price)

		return err
	})

	return err
}

func (str *Store) InsertContract(contr Contract) error {
	contrctx := context.WithValue(context.Background(), "contract", contr)
	err := str.Handler.Wdb.AtomicContext(contrctx, func(ctx context.Context, tx *sql.Tx) error {
		contr := ctx.Value("contract").(Contract)

		stmt, err := tx.Prepare(insertExpr[1])
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(contr.Addr, contr.IsExec, contr.IsFinished)
		return err
	})

	return err
}

func (str *Store) InsertUser(usr User) error {
	userctx := context.WithValue(context.Background(), "user", usr)
	err := str.Handler.Wdb.AtomicContext(userctx, func(ctx context.Context, tx *sql.Tx) error {
		usr := ctx.Value("user").(User)

		stmt, err := tx.Prepare(insertExpr[2])
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(usr.Addr, usr.Nonce, usr.ProfilePicUrl, usr.Bio)
		return err
	})

	return err
}

func (str *Store) UpdateUserNonce(id int, nonce string) error {
	usrctx := context.WithValue(context.Background(), "id", id)
	usrctx = context.WithValue(usrctx, "nonce", nonce)
	err := str.Handler.Wdb.AtomicContext(usrctx, func(ctx context.Context, tx *sql.Tx) error {
		usr := ctx.Value("id").(int)
		nce := ctx.Value("nonce").(string)

		stmt, err := tx.Prepare("UPDATE users SET nonce=? WHERE user_id=?")
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(nce, usr)
		return err
	})

	return err
}

func (str *Store) InsertContrib(contrib Contribution) error {
	contrctx := context.WithValue(context.Background(), "contribution", contrib)
	err := str.Handler.Wdb.AtomicContext(contrctx, func(ctx context.Context, tx *sql.Tx) error {
		contrib := ctx.Value("contribution").(Contribution)

		stmt, err := tx.Prepare(insertExpr[3])
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(contrib.ExecContract, contrib.User, contrib.Amt, contrib.Status)
		return err
	})

	return err
}

// Paginated GET queries

// Returns 20 users, ordered by id, whose IDs are all greater than lastID.
// Use lastID=-1 for the first 20 users.
func (str *Store) GetUsers(lastID int) ([]User, error) {
	userctx := context.WithValue(context.Background(), "users", &([]User{}))
	userctx = context.WithValue(userctx, "lastID", lastID)
	err := str.Handler.Rdb.AtomicContext(userctx, func(ctx context.Context, tx *sql.Tx) error {
		users := ctx.Value("users").(*([]User))
		lastID := ctx.Value("lastID").(int)

		rows, err := tx.Query("SELECT * FROM users WHERE user_id > ? ORDER BY user_id ASC LIMIT 20", lastID)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			usr := User{}
			err = rows.Scan(&usr.UserID, &usr.Addr, &usr.Nonce, &usr.ProfilePicUrl, &usr.Bio)
			if err != nil {
				return err
			}

			*users = append(*users, usr)
		}

		return nil
	})

	if err != nil {
		return []User{}, err
	}

	return *userctx.Value("users").(*([]User)), nil
}

// Return 20 assets, ordered by id, whose IDs are all greater than lastID.
// Use lastID=-1 for the first 20 assets.
func (str *Store) GetAssets(lastID int) ([]Asset, error) {
	assetctx := context.WithValue(context.Background(), "assets", &([]Asset{}))
	assetctx = context.WithValue(assetctx, "lastID", lastID)
	err := str.Handler.Rdb.AtomicContext(assetctx, func(ctx context.Context, tx *sql.Tx) error {
		assets := ctx.Value("assets").(*([]Asset))
		lastID := ctx.Value("lastID").(int)

		rows, err := tx.Query("SELECT * FROM assets WHERE asset_id > ? ORDER BY asset_id ASC LIMIT 20", lastID)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			asset := Asset{}
			err = rows.Scan(&asset.AssetID, &asset.ExecContract, &asset.AssetContract,
				&asset.ImgPreview, &asset.Executor, &asset.Desc, &asset.Price)

			if err != nil {
				return err
			}

			*assets = append(*assets, asset)
		}

		return nil
	})

	if err != nil {
		return []Asset{}, err
	}

	return *assetctx.Value("assets").(*([]Asset)), nil
}

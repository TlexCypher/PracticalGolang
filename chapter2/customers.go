//特定の構造体を引数として受け取り、その構造体のフィールドを使って条件判定等を行っている場合
//レシーバ変数としてその構造体を受け取るように実装を変更すると良い

package main

import "time"

type Customer struct {
	ActiveFlag bool
	Name       string
	Age        uint16
}

// 配列やsliceに対しても定義型を宣言することができる
type Customers []Customer

// コレクションの取得操作は型にまとめておくと見通しがよくなる
// メソッドチェーン等を使えるので、コレクションを取得する操作は戻り値もレシーバ変数と同じ型にしておくと良い
func (c Customers) ActiveCustomers() Customers {
	resp := make([]Customer, 0, len(c))
	for _, v := range c {
		if v.ActiveFlag {
			resp = append(resp, v)
		}
	}
	return resp
}

// e.g)
func (c Customers) RequiredFollows() Customers {
	//return c.activeCustomer().expires(time.Now().AddDate(0, 1, 0)).sortByExpiredAt()
	return nil
}

func (c Customers) activeCustomer() Customers {
	//契約が有効なユーザに絞り込む
	return nil
}

func (c Customers) expires(end time.Time) Customers {
	//endのにちじ以降に契約が執行するユーザに絞り込む
	return nil
}

func (c Customers) sortByExpiredAt() Customers {
	//契約期限日で昇順にソートする
	return nil
}

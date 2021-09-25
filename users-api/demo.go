package main

type Account struct {
	Id    int
	Name  string
	Email string
}

type HistoryAccount struct {
	Id          int
	AccountId   int
	LastUpdated string
	Action      string
}

type AccountHistoryItem struct {
	AccountId   int
	HistoryId   int
	Name        string
	Email       string
	LastUpdated string
	Action      string
}

// VD: có 1000 records(Account) & 3000 records(HistoryAccount)
func GetAllHistoryByAccount(id int) ([]AccountHistoryItem, error) {
	// cách cơ bản nhất là đi for
	items := []Account{}
	for i := 0; i < 1000; i++ {
		item := Account{Id: i, Name: "test", Email: "test"}
		items = append(items, item)
	}

	items2 := []HistoryAccount{}
	for i := 0; i < 2000; i++ {
		item2 := HistoryAccount{Id: i, AccountId: 1, LastUpdated: "test", Action: "test"}
		items2 = append(items2, item2)
	}
}

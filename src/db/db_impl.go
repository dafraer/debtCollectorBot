package db

// TODO THIS PACKAGE IS FULLY WRONG CHANGE IT BY ADDING DB
func GetDebtInfo() ([]Debt, error) {
	s := make([]Debt, 4)
	var a, b, c, d Debt
	a.DebtorUsername = "@dafraer"
	a.OwnerUsername = "@fiodop"
	a.Amount = 100
	a.Currency = "RUB"
	a.Language = "en"

	b.DebtorUsername = "@fiodop"
	b.OwnerUsername = "@dafraer"
	b.Amount = 230
	b.Currency = "RUB"
	b.Language = "en"

	c.DebtorUsername = "@kirooshi"
	c.OwnerUsername = "@Fleitas_Tobias14"
	c.Amount = 22
	c.Currency = "GEL"
	c.Language = "en"

	d.DebtorUsername = "@Fleitas_Tobias14"
	d.OwnerUsername = "@dafraer"
	d.Amount = 1456
	d.Currency = "GEL"
	d.Language = "en"
	s[0] = a
	s[1] = b
	s[2] = c
	s[3] = d
	return s, nil
}

func UpdateDebtInfo() error {
	return nil
}

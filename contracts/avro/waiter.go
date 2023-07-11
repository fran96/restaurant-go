package contracts

type makeDrinks struct {
	OrderID    int      `avro:"orderID" json:"orderID"`
	DrinkItems []string `avro:"drinkItems" json:"drinkItems"`
}

type drinksCompleted struct {
	OrderID int `avro:"orderID" json:"orderID"`
}

type OrderCompleted struct {
	OrderID string `avro:"orderID" json:"orderID"`
}

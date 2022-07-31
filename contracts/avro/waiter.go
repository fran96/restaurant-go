package contracts

var MakeDrinksSchema = `{
	"type": "record",
	"name": "makeFood",
	"namespace": "data.avro",
	"fields" : [
		{
			"name": "orderID", 
			"type": "int"
		},
		{
			"name": "drinkItems", 
			"type": {
				"type": "array",
				"items": "string"
			}
		}
	]
}`

type makeDrinks struct {
	OrderID    int      `avro:"orderID" json:"orderID"`
	DrinkItems []string `avro:"drinkItems" json:"drinkItems"`
}

type drinksCompleted struct {
	OrderID int `avro:"orderID" json:"orderID"`
}

type orderCompleted struct {
	OrderID int `avro:"orderID" json:"orderID"`
}

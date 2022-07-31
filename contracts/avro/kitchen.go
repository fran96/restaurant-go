package contracts

var MakeFoodSchema = `{
	"type": "record",
	"name": "makeFood",
	"namespace": "data.avro",
	"fields" : [
		{
			"name": "orderID", 
			"type": "int"
		},
		{
			"name": "foodItems", 
			"type": {
				"type": "array",
				"items": "string"
			}
		}
	]
}`

type MakeFood struct {
	OrderID   int      `avro:"orderID" json:"orderID"`
	FoodItems []string `avro:"foodItems" json:"foodItems"`
}

type foodCompleted struct {
	OrderID int `avro:"orderID" json:"orderID"`
}

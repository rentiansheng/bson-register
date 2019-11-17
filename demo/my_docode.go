package main

import (
	"fmt"

	_ "bson-register/register"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	fmt.Println("offical bson interface")
	mapDecodeResult()
	fmt.Println("offical bson struct anonymouse")
	structDecodeResult()

}

func mapDecodeResult() {
	type testSuit struct {
		name interface{}
		data interface{}
	}
	testDataArr := []testSuit{
		testSuit{
			name: "map",
			data: map[string]interface{}{"str": "str"},
		},

		testSuit{
			name: "str_arr",
			// bson.Unmarshal silce type is primitive.A{"item1", "item2"}
			data: map[string]interface{}{"str_arr": []string{"item1", "item2"}},
		},
		testSuit{
			name: "int_arr",
			// bson.Unmarshal silce type is primitive.A{1, 1}
			data: map[string]interface{}{"str_arr": []int{1, 2}},
		},

		testSuit{
			name: "map_suit",
			data: map[string]interface{}{"map_suit": map[string]interface{}{
				// bson.Unmarshal silce type is primitive.A{"map_suit_item1", "map_suit_item2"}
				"str_arr": []string{"map_suit_item1", "map_suit_item2"},
				// bson.Unmarshal silce type is primitive.A{1,2,3}
				"int":          []int{1, 2, 3},
				"map_suit_map": map[string]interface{}{"str": "str", "bool": "bool"},
			}},
		},
	}

	for _, item := range testDataArr {
		byteArr, err := bson.Marshal(item.data)
		if err != nil {
			panic(err.Error())
		}
		newData := make(map[string]interface{}, 0)
		err = bson.Unmarshal(byteArr, &newData)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(fmt.Sprintf("%#v", newData))
	}
}

func structDecodeResult() {
	type SubStruct struct {
		Str string `bson:"str"`
	}

	type testStruct struct {
		SubStruct
	}

	byteArr, err := bson.Marshal(map[string]string{"str": "str_val"})
	if err != nil {
		panic(err.Error())
	}
	newData := &testStruct{}
	err = bson.Unmarshal(byteArr, &newData)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(fmt.Sprintf("%#v", newData))

}

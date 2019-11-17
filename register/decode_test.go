package register

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func TestDecode(t *testing.T) {

	type testSuit struct {
		name interface{}
		data interface{}
	}

	testDataArr := []testSuit{
		testSuit{
			name: "str",
			data: map[string]interface{}{"str": "str"},
		},
		testSuit{
			name: "bool",
			data: map[string]interface{}{"bool": "bool"},
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
			name: "map_str_str",
			data: map[string]string{"str_arr": "map_str_str"},
		},
		testSuit{
			name: "map_str_str_arr",
			// bson.Unmarshal silce type is primitive.A{"map_str_str_arr"}
			data: map[string][]string{"str_arr": []string{"map_str_str_arr"}},
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
	for _, testData := range testDataArr {
		byteArr, err := bson.Marshal(testData.data)
		require.NoError(t, err)

		newData := make(map[string]interface{}, 0)
		err = bson.Unmarshal(byteArr, &newData)
		require.NoError(t, err)

		equalJSON(t, newData, testData.data)

	}

}

func TestStructField(t *testing.T) {

	byteArr, err := bson.Marshal(mapData)
	require.NoError(t, err)

	newData := TestStruct{}
	err = bson.Unmarshal(byteArr, &newData)
	require.NoError(t, err)

	equalStructJSON(t, newData, mapData)

}

func TestStructAnonymous(t *testing.T) {

	type AnonymousStruct struct {
		TestStruct
	}

	byteArr, err := bson.Marshal(mapData)
	require.NoError(t, err)

	newData := AnonymousStruct{}
	err = bson.Unmarshal(byteArr, &newData)
	require.NoError(t, err)

	equalStructJSON(t, newData, mapData)

	/*
		// wait todo
		type AnonymousPtrStruct struct {
			*TestStruct
		}

		newPtrData := AnonymousPtrStruct{}
		err = bson.Unmarshal(byteArr, &newPtrData)
		require.NoError(t, err)
		// equalStructJSON(t, newPtrData, mapData)
	*/

}

func TestStructLowerAnonymous(t *testing.T) {
	/*
		type lowerStruct struct {
			Str string `bson:"str" json:"str"`
		}

		type AnonymousStruct struct {
			lowerStruct
		}

		// wait todo
		type AnonymousPtrStruct struct {
			*lowerStruct
		}

		byteArr, err := bson.Marshal(mapData)
		require.NoError(t, err)

		newData := AnonymousStruct{}
		err = bson.Unmarshal(byteArr, &newData)
		require.NoError(t, err)

		// wait TODO
		//equalStructJSON(t, newData, mapData)

		newPtrData := AnonymousPtrStruct{}
		err = bson.Unmarshal(byteArr, &newPtrData)
		require.NoError(t, err)
		// wait TODO
		// equalStructJSON(t, newPtrData, mapData)
	*/
}

type TestStruct struct {
	Int          int                    `json:"int" bson:"int"`
	Str          string                 `json:"str" bson:"str"`
	Tag          string                 `json:"tag" bson:"tag,test"`
	StrArr       []string               `json:"str_arr" bson:"str_arr"`
	IntArr       []int                  `json:"int_arr" bson:"int_arr"`
	MapStrStr    map[string]string      `json:"map_str_str" bson:"map_str_str"`
	MapStr       map[string]interface{} `json:"map_str" bson:"map_str"`
	MapStrStrArr map[string][]string    `json:"map_str_str_arr" bson:"map_str_str_arr"`
	Bool         bool                   `json:"bool" bson:"bool"`
}

var (
	mapData = map[string]interface{}{
		"int":             1,
		"str":             "str",
		"tag":             "tag test",
		"str_arr":         []string{"item1", "item2"},
		"int_arr":         []int{1, 2, 3},
		"map_str_str":     map[string]string{"key1": "val1", "key2": "val2"},
		"map_str":         map[string]interface{}{"map_int": 1, "map_str": "str", "map_bool": true, "map_str_arr": []string{"map_str_arr_item1"}},
		"map_str_str_arr": map[string][]string{"map_str_str_arr_1": []string{"str1", "str2"}},
		"bool":            true,
	}
)

func equalStructJSON(t *testing.T, actual, expected interface{}) {
	actualByteArr, err := json.Marshal(actual)
	require.NoError(t, err)

	newData := make(map[string]interface{}, 0)
	err = json.Unmarshal(actualByteArr, &newData)
	require.NoError(t, err)

	equalJSON(t, newData, expected)
}

func equalJSON(t *testing.T, actual, expected interface{}) {
	actualByteArr, err := json.Marshal(actual)
	require.NoError(t, err)

	expectedByteArr, err := json.Marshal(expected)
	require.NoError(t, err)

	require.Equal(t, string(actualByteArr), string(expectedByteArr))
}

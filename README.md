# mongodb-go-driver-bson-register
bson decode interface priority use slice,map


The mongodb golang official driver bson uses the interface at the top, and puts the slice, map into an object called primitive.D (golang []interface type).

eg: 

``` json

// db data info 

{"str":"str"}
{"str_arr":["item11","item2"]}


// driver retrun  info

primitive.D{primitive.E{Key:"str", Value:"str"}}
primitive.D{primitive.E{Key:"str_arr", Value:primitive.A{"item1", "item2"}}}

```

### my implemented

##### 1. support interface decode 
Support when the incoming object is interface, the slice resolves to the privitive.A type (glang []interface{}), and the map object resolves to map[string]interface{}

#### 2. support anonymouse decode 





### result Compared


#### interface
official result :

``` 
primitive.D{primitive.E{Key:"str", Value:"str"}}
primitive.D{primitive.E{Key:"str_arr", Value:primitive.A{"item1", "item2"}}}
primitive.D{primitive.E{Key:"str_arr", Value:primitive.A{1, 2}}}
primitive.D{primitive.E{Key:"map_suit", Value:primitive.D{primitive.E{Key:"str_arr", Value:primitive.A{"map_suit_item1", "map_suit_item2"}}, primitive.E{Key:"int", Value:primitive.A{1, 2, 3}}, primitive.E{Key:"map_suit_map", Value:primitive.D{primitive.E{Key:"str", Value:"str"}, primitive.E{Key:"bool", Value:"bool"}}}}}}

```

my implemented result :

```

primitive.D{primitive.E{Key:"str", Value:"str"}}
primitive.D{primitive.E{Key:"str_arr", Value:primitive.A{"item1", "item2"}}}
primitive.D{primitive.E{Key:"str_arr", Value:primitive.A{1, 2}}}
primitive.D{primitive.E{Key:"map_suit", Value:primitive.D{primitive.E{Key:"str_arr", Value:primitive.A{"map_suit_item1", "map_suit_item2"}}, primitive.E{Key:"int", Value:primitive.A{1, 2, 3}}, primitive.E{Key:"map_suit_map", Value:primitive.D{primitive.E{Key:"str", Value:"str"}, primitive.E{Key:"bool", Value:"bool"}}}}}}

```

#### struct anonymouse property 



official result :

```
&main.testStruct{SubStruct:main.SubStruct{Str:""}}

```

my implemented result:

```
&main.testStruct{SubStruct:main.SubStruct{Str:"str_val"}}

```




### Specific examples result dispay


- demo/my_docode.go
- demo/offical_decode.go
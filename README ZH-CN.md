# mongodb-go-driver-bson-register
bson decode interface priority use slice,map


mongodb golang 官方driver中bson在顶层使用interface的时候，会将slice, map 放到一个叫primitive.D的对象中（golang []interface类型）。

示例: 

``` json

// 数据库中数据

{"str":"str"}
{"str_arr":["item11","item2"]}


// driver 返回数据

primitive.D{primitive.E{Key:"str", Value:"str"}}
primitive.D{primitive.E{Key:"str_arr", Value:primitive.A{"item1", "item2"}}}

```

### 我实现的功能

##### 1. 解析到interface类型支持
支持当传入对象是interface的时候，slice 解析到 privitive.A 类型（glang []interface{}）,map 对象解析到map[string]interface{}

#### 2. 支持struct匿名属性





### 结果对比


#### interface

官方driver结果:

``` 
primitive.D{primitive.E{Key:"str", Value:"str"}}
primitive.D{primitive.E{Key:"str_arr", Value:primitive.A{"item1", "item2"}}}
primitive.D{primitive.E{Key:"str_arr", Value:primitive.A{1, 2}}}
primitive.D{primitive.E{Key:"map_suit", Value:primitive.D{primitive.E{Key:"str_arr", Value:primitive.A{"map_suit_item1", "map_suit_item2"}}, primitive.E{Key:"int", Value:primitive.A{1, 2, 3}}, primitive.E{Key:"map_suit_map", Value:primitive.D{primitive.E{Key:"str", Value:"str"}, primitive.E{Key:"bool", Value:"bool"}}}}}}

```

我实现的register结果 :

```

primitive.D{primitive.E{Key:"str", Value:"str"}}
primitive.D{primitive.E{Key:"str_arr", Value:primitive.A{"item1", "item2"}}}
primitive.D{primitive.E{Key:"str_arr", Value:primitive.A{1, 2}}}
primitive.D{primitive.E{Key:"map_suit", Value:primitive.D{primitive.E{Key:"str_arr", Value:primitive.A{"map_suit_item1", "map_suit_item2"}}, primitive.E{Key:"int", Value:primitive.A{1, 2, 3}}, primitive.E{Key:"map_suit_map", Value:primitive.D{primitive.E{Key:"str", Value:"str"}, primitive.E{Key:"bool", Value:"bool"}}}}}}

```

#### struct  匿名属性



官方driver结果 :

```
&main.testStruct{SubStruct:main.SubStruct{Str:""}}

```

我实现的register结果 :

```
&main.testStruct{SubStruct:main.SubStruct{Str:"str_val"}}

```




### 结果对比样例文件

- demo/my_docode.go
- demo/offical_decode.go
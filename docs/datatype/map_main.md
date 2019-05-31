map就是键值对(key->value)，在其他语言中通常叫hash或者字典

## **完整写法**

---

```text
var m map[key_type]value_type
m = make(map[key_type]value_type)
```

例子

```text
var m map[string]int
m = make(map[string]int)
```

!!! note "简写"
	```text
	var m map[key_type]value_type = make(map[key_type]value_type)
	```

	例子

	```text
	var m map[string]int = make(make[string]int)
	```

!!! warning "key和value的类型要求"
	key必须是支持==和!=比较的类型，比如int、string等

	value可以是任意类型，因此value可以嵌套

## **短声明写法**

---

```text
m := make(map[key_type]value_type)
```

例子

```text
m := make(map[string]int)
```

## **value嵌套**

以下表示声明一个没有任何key-value的map，其中key类型为int，value类型为map[int]{string}。另外，value如果使用嵌套，则不需要使用make()

```go
a := map[int]map[int]string{}
fmt.Println(a)
```

输出

```
map[]
```

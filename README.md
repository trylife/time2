time2
======

### 

```bash
go get -u github.com:trylife/time2
```

### 设置起始时间

#### 设置当前时间为起始时间

```go
time2.Now() 
```

#### 载入某起始时间

```go
time2.Load(time.Now)
```

### 根据起始时间推算其他时间

#### 获取当天0点时间

```go
time2.Now().CurrentDayStart()
```

#### 获取当天零点时间戳

```go
time2.Now().CurrentDayStart().Unix()
```
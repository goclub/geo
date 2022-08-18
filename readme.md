---
permalink: /
sidebarBasedOnContent: true
---

# goclub/geo

Golang 地理位置,支持 mongo goclub/sql WGS84 GCJ02 BD09

[![Go Reference](https://pkg.go.dev/badge/github.com/goclub/geo.svg)](https://pkg.go.dev/github.com/goclub/geo)

## install

```shell
go get github.com/goclub/geo
```

```go
import xgeo "github.com/goclub/geo"
```

## WGS84

```go
point := xgeo.WGS84{121.48294,31.2328}
// 转换为gcj02坐标
point.GCJ02().LatCommaLngString()
```

## Point

支持 mongo json mysql

```go
point := xgeo.NewPoint(xgeo.WGS84{121.48294,31.2328}) // WGS84{经度,纬度}
// 转换为wgs84坐标
point.WGS84()
```
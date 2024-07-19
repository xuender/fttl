# fttl

[![Action][action-svg]][action-url]
[![Report Card][goreport-svg]][goreport-url]
[![Lines of code][lines-svg]][lines-url]
[![godoc][godoc-svg]][godoc-url]
[![License][license-svg]][license-url]

✨ **`xuender/fttl` is a time to live cache based on file system.**

## 🚀 Install

```shell
go install github.com/xuender/fttl@latest
```

## 💡 Usage

### base

```go
fdb := fttl.New(filepath.Join(os.TempDir(), "base"))
defer fdb.Close()

fdb.Put([]byte("key"), []byte("value"))

val, _ := fdb.Get([]byte("key"))
fmt.Println(string(val))
fmt.Println(fdb.Has([]byte("key")))

fdb.Delete([]byte("key"))
```

### ttl

```go
fdb.PutTTL([]byte("key"), []byte("value"), time.Minute, time.Second)
```

## 👤 Contributors

![Contributors][contributors-svg]

## 📝 License

© ender, 2024~time.Now

[MIT LICENSE][license-url]

[action-url]: https://github.com/xuender/fttl/actions
[action-svg]: https://github.com/xuender/fttl/workflows/Go/badge.svg

[goreport-url]: https://goreportcard.com/report/github.com/xuender/fttl
[goreport-svg]: https://goreportcard.com/badge/github.com/xuender/fttl

[godoc-url]: https://godoc.org/github.com/xuender/fttl
[godoc-svg]: https://godoc.org/github.com/xuender/fttl?status.svg

[license-url]: https://github.com/xuender/fttl/blob/master/LICENSE
[license-svg]: https://img.shields.io/badge/license-MIT-blue.svg

[contributors-svg]: https://contrib.rocks/image?repo=xuender/fttl

[lines-svg]: https://sloc.xyz/github/xuender/fttl
[lines-url]: https://github.com/boyter/scc

# List vs Map

データ構造としてListを用いた場合とMapを用いた場合で、パフォーマンスにどのような影響があるのかを確認するためのサンプルです。

## Overview

下記のような仕様を満たす関数を作成します。
なお、先頭にIDを含む文字列配列を「レコード」と呼ぶこととします。

**関数仕様: 2つのレコード配列を引数として受け取り、IDが等しいレコードの要素をマージする。**

```go
func(d1, d2 [][]string) ([][]string, error)
```

- 簡単のため、ふたつのレコード配列に存在するIDは一致するものとする。
- ただし、ふたつの配列内のレコードの順序が同じとは限らない。

### Example

Input1

```
[
  ["id1", "val1-1", "val1-2"],
  ["id2", "val2-1", "val2-2"],
  ...
]
```

Input2

```
[
  ["id1", "val1-3", "val1-4"],
  ["id2", "val2-3", "val2-4"],
  ...
]
```

Output

```
[
  ["id1", "val1-1", "val1-2", "val1-3", "val1-4"],
  ["id2", "val2-1", "val2-2", "val2-3", "val2-4"],
  ...
]
```

## Implements

- `SliceSample`: List構造（Slice）を用いた実装。
- `MapSample`: Map構造を用いた実装。

## Functional test

-  `TestSliceSample`, `TestMapSample`にて確認。どちらも機能仕様を満たす。

## Benchmark test

- `BenchmarkSliceSampleXXX`, `BenchmarkMapSampleXXX`にて確認可能。
- ベンチマーク結果を比較する場合は、`BenchmarkSample`を利用して確認可能。

### Compare benchmark

1. `BenchmarkSample`内の`benchmarkSample`の引数（`length`, `fn`）を変更し、それぞれのベンチマーク結果を保存する。

```cmd
$ go test -count 5 -test.bench BenchmarkSample > out/benchSlice10.log
$ go test -count 5 -test.bench BenchmarkSample > out/benchMap10.log
$ go test -count 5 -test.bench BenchmarkSample > out/benchSlice100.log
$ go test -count 5 -test.bench BenchmarkSample > out/benchMap100.log
$ go test -count 5 -test.bench BenchmarkSample > out/benchSlice1000.log
$ go test -count 5 -test.bench BenchmarkSample > out/benchMap1000.log
```

2. `length`が同じベンチマーク結果を比較する。

```cmd
$ benchstat out/benchSlice10.log out/benchMap10.log > out/compare10.log
$ benchstat out/benchSlice100.log out/benchMap100.log > out/compare100.log
$ benchstat out/benchSlice1000.log out/benchMap1000.log > out/compare1000.log
```

* Required `go get rsc.io/benchstat`

**Warning**

ここのサンプルでは、Mapを用いた方がパフォーマンス向上が見込まれる例を挙げていますが、必ずしもMapが優れているというわけではありません。

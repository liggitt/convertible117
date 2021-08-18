```
gvm use go1.16 && go test ./ -v -count=1
Now using version go1.16
=== RUN   TestConvertible
--- PASS: TestConvertible (0.00s)
PASS
ok  	github.com/liggitt/convertible117	0.629s


gvm use go1.17 && go test ./ -v -count=1
Now using version go1.17
=== RUN   TestConvertible
    convertible_test.go:24: !convertible
--- FAIL: TestConvertible (0.00s)
FAIL
FAIL	github.com/liggitt/convertible117	0.566s
FAIL
```

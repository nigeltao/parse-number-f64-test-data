# `parse_number_f64` Test Data

This repository contains test data for `parse_number_f64` implementations, also
known as `StringToDouble` or `strtod`, which convert from an ASCII string to a
64-bit value (IEEE 754 double-precision floating point).

Most of the `data/*.txt` files were derived by running
`script/extract-numbery-strings.go` on various repositories or zip files,
listed further below. Their contents look like:

    3FF0000000000000 1
    3FF4000000000000 1.25
    3FF6666666666666 1.4
    405EDD2F1A9FBE77 123.456
    4088A80000000000 789
    7FF0000000000000 123.456e789

In this case, the final line's `float64` representation is infinity. The
largest finite `float64` is approximately `1.8e+308`.


## Data

In the `data` directory:

- `freetype-2-7.txt` was extracted from [Freetype](https://www.freetype.org/)
  2.7
- `google-double-conversion.txt` was extracted from
  [google/double-conversion](https://github.com/google/double-conversion)
- `google-wuffs.txt` was extracted from
  [google/wuffs](https://github.com/google/wuffs)
- `ibm-fpgen.txt` was extracted from IBM's
  [IEEE 754R test suite](https://www.research.ibm.com/haifa/projects/verification/fpgen/test_suite_download.shtml)
- `lemire-fast-double-parser.txt` was extracted from
  [lemire/fast\_double\_parser](https://github.com/lemire/fast_double_parser)
- `tencent-rapidjson.txt` was extracted from
  [Tencent/rapidjson](https://github.com/Tencent/rapidjson)
- `ulfjack-ryu.txt` was extracted from
  [ulfjack/ryu](https://github.com/ulfjack/ryu)


### remyoudompheng/fptest

The `data/remyoudompheng-fptest-?.txt` files were created by running
`go test -test.run=TestTortureAtof64` in the
[remyoudompheng/fptest](https://github.com/remyoudompheng/fptest) repository
(with the following patch), `sort`ing  and `uniq`ing the resultant
`TestTortureAtof64.txt` file and then using `sed` to split what would be a 137
MiB file into multiple (million line) files:

```diff
    diff --git a/torture_test.go b/torture_test.go
    index 87ba7e7..59887ff 100644
    --- a/torture_test.go
    +++ b/torture_test.go
    @@ -1,8 +1,11 @@
     package fptest

     import (
    +       "bufio"
            "bytes"
    +       "fmt"
            "math"
    +       "os"
            "strconv"
            "testing"

    @@ -124,6 +127,11 @@ func TestTortureShortest32(t *testing.T) {
     }

     func TestTortureAtof64(t *testing.T) {
    +       tmpFile, _ := os.Create("/tmp/TestTortureAtof64.txt")
    +       defer tmpFile.Close()
    +       tmpWriter := bufio.NewWriter(tmpFile)
    +       defer tmpWriter.Flush()
    +
            count := 0
            buf := make([]byte, 64)
            roundUp := false
    @@ -140,6 +148,7 @@ func TestTortureAtof64(t *testing.T) {
                            t.Errorf("could not parse %q: %s", s, err)
                            return
                    }
    +               fmt.Fprintf(tmpWriter, "%016X %s\n", math.Float64bits(z), s)
                    expect := x
                    if roundUp {
                            expect = y
```


## Users

Programs that use this test data set:

- `script/manual-test-parse-number-f64.cc` in
  [google/wuffs](https://github.com/google/wuffs)

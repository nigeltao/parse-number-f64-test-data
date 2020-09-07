# `parse_number_f64` Test Data

This repository contains test data for `parse_number_f64` implementations, also
known as `StringToDouble` or `strtod`, which convert from an ASCII string to a
64-bit value (IEEE 754 double-precision floating point).

The `data/*.txt` files were derived by running
`script/extract-numbery-strings.go` on various repositories, listed further
below. Their contents look like:

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
- `lemire-fast-double-parser.txt` was extracted from
  [lemire/fast\_double\_parser](https://github.com/lemire/fast_double_parser)
- `tencent-rapidjson.txt` was extracted from
  [Tencent/rapidjson](https://github.com/Tencent/rapidjson)
- `ulfjack-ryu.txt` was extracted from
  [ulfjack/ryu](https://github.com/ulfjack/ryu)


## Users

Programs that use this test data set:

- `script/manual-test-parse-number-f64.cc` in
  [google/wuffs](https://github.com/google/wuffs)

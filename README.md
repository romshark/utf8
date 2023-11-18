# Experimental encoding/utf8

⚠️ Experimental, don't use in production! ⚠️

This is a fork of the [encoding/utf8](https://github.com/golang/go/tree/631a6c2abfb2cf7a877ea80f98c504fc4e0077be/src/unicode/utf8) Go standard library package with experimental performance optimizations.

## Benchmark Results

**Ryzen 7 5700X**
```
goos: linux
goarch: amd64
pkg: utf8validbench
cpu: AMD Ryzen 7 5700X 8-Core Processor             
                │ amd64_old.txt │            amd64_new.txt            │
                │    sec/op     │   sec/op     vs base                │
/ascii_16b.txt      2.626n ± 1%   2.624n ± 1%        ~ (p=0.565 n=16)
/ascii_2k.txt       138.2n ± 0%   115.2n ± 0%  -16.68% (p=0.000 n=16)
/unicode_2k.txt     1.620µ ± 1%   1.045µ ± 2%  -35.49% (p=0.000 n=16)
/wiki_ukr.html     1253.1µ ± 0%   694.9µ ± 1%  -44.54% (p=0.000 n=16)
/wiki_jap.html     1956.7µ ± 0%   878.4µ ± 0%  -55.11% (p=0.000 n=16)
/ascii_1mib.txt     72.45µ ± 1%   49.24µ ± 1%  -32.03% (p=0.000 n=16)
geomean             6.862µ        4.601µ       -32.95%
```

**Apple M1 on Sonoma 14.1.1**
```
goos: darwin
goarch: arm64
pkg: utf8validbench
                │  m1_old.txt  │             m1_new.txt              │
                │    sec/op    │   sec/op     vs base                │
/ascii_16b.txt     3.415n ± 0%   2.484n ± 0%  -27.26% (p=0.000 n=16)
/ascii_2k.txt      178.0n ± 0%   130.4n ± 0%  -26.74% (p=0.000 n=16)
/unicode_2k.txt    2.190µ ± 0%   1.353µ ± 0%  -38.22% (p=0.000 n=16)
/wiki_ukr.html    1592.0µ ± 0%   905.8µ ± 0%  -43.10% (p=0.000 n=16)
/wiki_jap.html     2.411m ± 0%   1.102m ± 0%  -54.28% (p=0.000 n=16)
/ascii_1mib.txt    86.24µ ± 0%   55.09µ ± 0%  -36.12% (p=0.000 n=16)
geomean            8.723µ        5.375µ       -38.39%
```

**Raspberry Pi 2 Model B**
```
goos: linux
goarch: arm
pkg: utf8validbench
                │ raspi_old.txt │            raspi_new.txt            │
                │    sec/op     │   sec/op     vs base                │
/ascii_16b.txt      187.0n ± 1%   172.5n ± 1%   -7.73% (p=0.000 n=16)
/ascii_2k.txt       18.90µ ± 1%   17.27µ ± 3%   -8.62% (p=0.000 n=16)
/unicode_2k.txt     52.53µ ± 3%   65.95µ ± 2%  +25.53% (p=0.000 n=16)
/wiki_ukr.html      32.16m ± 4%   28.93m ± 3%  -10.05% (p=0.000 n=16)
/wiki_jap.html      57.98m ± 2%   49.27m ± 3%  -15.03% (p=0.000 n=16)
/ascii_1mib.txt    10.380m ± 2%   9.281m ± 2%  -10.59% (p=0.000 n=16)
geomean             391.4µ        370.8µ        -5.26%
```
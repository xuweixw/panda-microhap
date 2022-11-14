## panda-microhap
---
### 1 Compile and Run
```
go build -o panda-microhap Microhaplotype
./panda-microhap -v
```

```bash
awk -F '\t' '{print $8}' mh-panda.vcf | 
    awk -F ';' '{print $25}' |
    awk -F '=' '{print $2}' | 
    awk '{s[$1]++} END {for (i in s) print i, s[i]}'
```
2-14 SNPs in a microhap
```bash
2   364049  182024.5
3   133937  44645.7
4   60241   15060.3
5   29073   5814.6
6   15276   2546
7   7924    1132
8   4056    507
9   2457    273
10  1190    119
11  506 46
12  276 23
13  130 13
14  56  4
```
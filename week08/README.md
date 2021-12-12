
```shell
edis-benchmark -d n -t get,set
```

###SET
| n    | 执行次数和耗时                                   | 每秒请求次数                        |
|------|-------------------------------------------|-------------------------------|
| 10   | 100000 requests completed in 0.69 seconds | 145772.59 requests per second |
| 20   | 100000 requests completed in 0.70 seconds | 142857.14 requests per second |
| 50   | 100000 requests completed in 0.70 seconds | 143061.52 requests per second |
| 100  | 100000 requests completed in 0.69 seconds | 144508.67 requests per second |
| 200  | 100000 requests completed in 0.69 seconds | 145772.59 requests per second |
| 1000 | 100000 requests completed in 0.69 seconds | 144092.22 requests per second |
| 5000 | 100000 requests completed in 0.72 seconds | 138696.25 requests per second |

###GET

| n    | 执行次数和耗时                                   | 每秒请求次数                        |
|------|-------------------------------------------|-------------------------------|
| 10   | 100000 requests completed in 0.66 seconds | 151515.14 requests per second |
| 20   | 100000 requests completed in 0.67 seconds | 149031.30 requests per second |
| 50   | 100000 requests completed in 0.67 seconds | 149253.73 requests per second |
| 100  | 100000 requests completed in 0.66 seconds | 150602.42 requests per second |
| 200  | 100000 requests completed in 0.66 seconds | 152207.00 requests per second |
| 1000 | 100000 requests completed in 0.65 seconds | 153846.16 requests per second |
| 5000 | 100000 requests completed in 0.70 seconds | 142247.52 requests per second |
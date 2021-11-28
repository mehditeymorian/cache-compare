# Cache Databases Comparison
This is an application that can perform set, get, and delete key-value operation for 3 different cache database including Redis, MongoDB, and TiKV.
The end goal is to compare these databases to see how they perform against large number of requests.
Load test is done using [Apache Benchmark](https://httpd.apache.org/docs/2.4/programs/ab.html) from within the cluster using [ab-test](https://github.com/mehditeymorian/ab-test).

# Result of Load Testing
Load test is done by sending 10k request with 2 different concurrency rate. Connection time, Processing time, Waiting time, and total time is compared for databases.

### Result of Load Test with concurrency rate of 100
![load test c100](https://github.com/mehditeymorian/cache-compare/blob/master/results/c100.png)
              
### Result of Load Test with concurrency rate of 1000
![load test c1000](https://github.com/mehditeymorian/cache-compare/blob/master/results/c1000.png)

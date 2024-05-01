# HTTP Request Limiter

This http request limiter were implemented as a middleware package - rlm.

# Default params:
 * Maximum Requests per second = 5 r/sec
 * Block time for too much request = 120 sec
 * Block time increment = 0 sec

# Ways to run
```
go get -u github.com/doxanocap/pkg
```

### If you want run project with following params
where:
 * A => input maximum requests per second (int)
 * B => input block time for too much request (int)
 * C => block time increment (int)
   

# TCP
Each line will return the sum length of all the history line txt.<br>
send 'quit' to exit. <br>
send 'reset' to reset.

```
nc localhost 8000
```

# external API
It will return the sum length of all the history line txt.
```  
http://localhost:8001?txt=<txt>
```

Reset the sum length = 0
``` 
http://localhost:8001/reset
``` 


# test api requests per second
``` 
go test -benchmem -run=none -bench BenchmarkSumLen
``` 

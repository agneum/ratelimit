## Rate limit CLI

The util is an analog of xargs with a rate limit and parallel commands launching.

### Usage: 
```ratelimit --rate <N> --inflight <P> <command...>```

* --rate: максимальное кол-во запусков команды в секунду. По умолчанию, 1.
* --inflight: максимальное кол-во параллельно запущенных команд. По умолчанию, 1.
* <command...>: команда для запуска, {} в команде заменяется на строчку из stdin.

### Examples:
```  $ for i in {1..60} ; do echo $i ; done | ./ratelimit --rate 15 --inflight 1 echo {}```

``` $ (echo 1 ; sleep 3 ; echo 2 ; echo 3) | ./ratelimit --rate 1 --inflight 2 echo {}```
# Using app

`go run main.go` and press `CTRL+C` a few seconds in to interrupt.

3 workers will be running and each will print a line to indicate when they notice you requested shutdown of app.

## Example Output

```Text
$ go run main.go
2017/07/21 11:02:20 received message from 0
2017/07/21 11:02:21 received message from 1
2017/07/21 11:02:22 received message from 2
2017/07/21 11:02:23 received message from 0
2017/07/21 11:02:24 received message from 1
^C interruption received!
2017/07/21 11:02:25 worker 2 noticed termination signal and shut down
2017/07/21 11:02:26 worker 0 noticed termination signal and shut down
2017/07/21 11:02:27 received message from 1
2017/07/21 11:02:30 worker 1 noticed termination signal and shut down
2017/07/21 11:02:30 program terminated
```

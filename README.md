# Using app
`go run main.go` and press `CTRL+C` a few seconds in to interrupt.

3 workers will be running and each will print a line to indicate when they notice you requested shutdown of app.

## Example Output
```
$ go run main.go
2016-06-06 10:18:30.480309672 -0500 CDT
2016-06-06 10:18:31.480872359 -0500 CDT
2016-06-06 10:18:32.481427664 -0500 CDT
2016-06-06 10:18:33.481259987 -0500 CDT
2016-06-06 10:18:34.482368667 -0500 CDT
2016-06-06 10:18:35.481544601 -0500 CDT
2016-06-06 10:18:36.482315864 -0500 CDT
2016-06-06 10:18:37.485142585 -0500 CDT
2016-06-06 10:18:38.482697632 -0500 CDT
2016-06-06 10:18:39.4831977 -0500 CDT
^C interruption received!
Worker 1 saw stop signal
Worker 2 saw stop signal
2016-06-06 10:18:42.483336815 -0500 CDT
Worker 0 saw stop signal
Program terminated
```

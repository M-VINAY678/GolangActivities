## How to run Code
* Open the Terminal,then move to activity directory and then enter "go run main.go"
* You can see the port number where the project is deployed
* Open any browser and enter "http://localhost:8085/process", here I used 8085 port number, you can replace it by replacing in the http server.
* You can see your results in browser and some in terminal.

## How did we use context package to manage timeouts and cancellations:
* In general every http request has it's own context, so we will create a derived context for the request context using:
  * ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second), here you can add your own requried time.
* In general when parent context get canceled , then automatically child context gets canceled.
  * We handle it using cxt.Done() method, then we can check error using ctx.Error() method by using it we can log the error.
  * when their is timeout, we get error :  "context.DeadlineExceeded"
  * When user is canceled then , we get error : "context.Canceled"
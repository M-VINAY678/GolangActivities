## Simulataneous Request
* Used go routines for simulataneous execution .
* For communication between goroutine we used channels .
* As we need max five simulataneously request, I used a channel of empty struct, limiting for loops for 5 requests.
* Used waitgroup for co-ordination of goroutines and main go routine waits until completion of all requests.
## Code 1
`go run main.go`

4 tasks were scheduled in order but
task 3 ran 1st as it has least amount of running time and then task1, task4, task2 based on their running time


<img width="454" alt="Screenshot 2023-11-16 at 11 11 55 AM" src="https://github.com/kritika0598/go-concurrency/assets/30694412/38be2828-9e05-4d71-96bc-3e6a327e5cce">

## Code 2
`go run fork-join/no-join-point/main.go`

It printed `done waiting, main exits` and did not print work func statement

## Code 3 (Wait group)
`go run fork-join/wg-join-point/main.go`

<img width="587" alt="Screenshot 2023-11-16 at 11 26 33 AM" src="https://github.com/kritika0598/go-concurrency/assets/30694412/dd72b16e-a588-4348-a326-81a7bbcb01f7">

## Code 4 (Using Channels)
`go run fork-join/channel-join-point/main.go`

<img width="603" alt="Screenshot 2023-11-16 at 11 29 46 AM" src="https://github.com/kritika0598/go-concurrency/assets/30694412/92666936-4f0b-4f48-9b63-b66af397af0d">

## Run Benchmarks
`cd benchmarks`
`go test -bench=.`


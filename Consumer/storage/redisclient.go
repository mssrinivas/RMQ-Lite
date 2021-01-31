import (
    "context"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// Consumer writes task data to redis with a TTL
func writeData() {

}


// TaskDispatcher fetches data from redis and updates triggered tasks
func GetTaskData() {

}


func TaskStorageClient() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    fmt.Println(rdb.)
}
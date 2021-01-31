# Instructions and Implementation details

## Producer

### Modules
    * A basic webserver (included in the starter code) 
    * Implement a producer that creates "n" number of task messages using multiple threads (goroutines) of struct and performs a http request to 
      Broker on the defined endpoints.
      ( task_name, task_type, task_update_time, scheduled_time, periodicity )
      n - is a command line argument from the user or a fixed constant defined in the code.  
       
## Broker

### Modules 
    * A basic webserver (included in the starter code) 
    
#### Work-Queue
    * Define work-queues ( buffered channels ) and assign each work-queue with a binding id.
    * Implement a service in request.go and requestprocessor.go to handles requests from Producer and redirect the requests to Exchange module.

#### Exchange module
    * Exchange module stores the queue information and their associated binding id's. Based on the binding id's forwards the request to respective queue. (DIRECT STRATEGY)

#### Dispatcher 
    * Implement a dispatcher service that continuously polls a task from the queue and sends the respective task to the consumer via a http POST request. This has         to be a async call as an acknowledge message is sent by Consumer for the dispatched task and pop it from the queue.
    * If the dispatch request is unsuccessful, then retry the same request for a different consumer with exponential backoff.

## Consumer
    * A basic webserver (included in the starter code) 

#### Consumer
    * Implememt multiple consumers (multiple buffered channels) and assign multiple queues based on round-robin strategy.
    * Handle the requests from Broker service and once the task is sucessfully consumed by the respective consumer and then send back an acknowledge message.

#### Storage
    * Redis Client (included in the starter code)
    * Consumer writes task data to redis with a TTL (Time to live) .

## Task Dispatcher

### Modules

#### Scheduler
    * TaskDispatcher runs a thread that is invoked every 1 minute to check for tasks in Redis to be triggered as per scheduled_time. 
      (Update the status of task from Scheduled to Complete)

#### Data Cleanup
    * Implement Data Cleanup module to delete the Completed tasks from Redis storage older than 30 days of updated time.

## Redis Storage
    * Follow the steps in the link provided to start a redis server locally. 
      ( https://redis.io/topics/quickstart )

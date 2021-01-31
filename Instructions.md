# Instructions and Implementation details

## Producer

### Modules
    * A basic webserver (included in the starter code) 
    * Program should accept the count of tasks 
    * Implement a producer that creates task messages of struct ( task_name, task_type, task_updat_time, scheduled_time, periodicity )
    * 
        
       

## Broker

### Modules 
    * A basic webserver (included in the starter code) 
    * Define work-queues and assign each work-queue with a binding id.
    * Service that handles requests from Producer and redirects the requests to Exchange module
    * Exchange module 

## Consumer

    * Redis Client (included in the starter code)
    * Consumer writes task data to redis with a TTL
    * 

## Task Dispatcher
    * TaskDispatcher runs a thread that is invoked every 1 minute to check for tasks in Redis to be triggered. (Update the status of task from Scheduled to Complete)
    * TaskDispatcher implements "delete tasks" module that deletes Completed tasks from Redis storage that are older than 30 days of updated time.

## Redis Storage
    * Follow the steps in the link provided to start a redis server locally. ( https://redis.io/topics/quickstart )

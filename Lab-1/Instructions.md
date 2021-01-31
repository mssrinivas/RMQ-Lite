# Instructions and Implementation details

## Producer

### Modules
    * A basic webserver (included in the starter code) 
    * Program should accept the count of tasks 
    * Implement a simple producer that creates task messages of struct ( task_name, task_type, task_updat_time, scheduled_time, periodicity )

## Consumer
    * A basic webserver (included in the starter code) 

#### Consumer
    * Implememt multiple consumers (multiple buffered channels) and assign multiple queues based on round-robin strategy.
    * Handle the requests from Broker service and once the task is sucessfully consumed by the respective consumer and then send back an acknowledge message.
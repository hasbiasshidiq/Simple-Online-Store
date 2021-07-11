# SIMPLE ONLINE STORE - Race Condition Demo

## Overview
This repo demonstrate simulation of race conditions on simple online store. We also include proof of concept of this occurrence and its solution.

## Race Condition in Flash Sale Event
In database operations, race conditions often occur when two or more threads access the same row. The second thread reads data in the row immediately after identical reading process carried out by the first thread. Then the two threads perform concurrent operations and may write data with unwanted values to the database.

This can also happen in online stores, especially on the day of flash sales. On that day, multiple requests hit API at the same time. There are most likely multiple reading and writing db processes on identical row with overlaping timelapse. This can result in abnormal inventory changes. The figure below describe how it occurs.

![alt text](https://github.com/hasbiasshidiq/Simple-Online-Store/blob/main/images/Race-Condition.png?raw=true)

On the picture above, user A orders 4 sneakers while user B also orders 4 sneakers. Both processes read a shared row on inventory table almost simultaneously and return 5 available sneakers. This is the beginning of race condition because both processes use shared variable to make a decision for the next process. The decision is to check whether there are more available items than the ordered items. This checking mechanism returns a condition that allows both process to write database inventory. Process A reduces the inventory database with initial 5 minus 4 so that the total inventory in the associated row becomes 1. Followed by process B which performs write operation on the same row. The writing process here reduces the current available items which is 1 so that the amount of inventory is -3 (1-4). As we know, this value is abnormal.


## Row Level Locking as Solution
To overcome this problem, row level locking is carried out by the process that comes first, that is process A. As shown in the picture below

![alt text](https://github.com/hasbiasshidiq/Simple-Online-Store/blob/main/images/Race-Condition-with-Lock.png?raw=true)

Associated row will be locked until process A finishes doing transactions marked with db commit. Next, process B read db and gets 1 available items. Checking process returns an unsuccessful response because number of available items is less than ordered items. That way there is no writing process to database so data integrity is assured. Row level locking in SQL operations can be added with the term `FOR UPDATE` at the end of a select operation statement. 


## How to run the application
For how to run the application, you can refer this documentation [here](https://github.com/hasbiasshidiq/Simple-Online-Store/blob/main/README-demo.md)
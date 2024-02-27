# BotanDB
In memory cache for go.

## Description
Sometimes, when developing projects, we need caching.
But it is not always possible to use external infrastructure.
Due to this, the simplest and most effective way is ***in memory caching***.

### Example
https://github.com/DolenL3/botanTask or test files in this repository.

### Approach
I made a decision to use partitioned maps to store data, reasoning behind it is:  
- As the final size of stored data is unknown, letting a large map evacuate may be an expensive operation,
by partitioning it we assure that the size of evacuated maps is smaller, thus making the operation easier.
- Different partitions of partitioned map work independetly, leading to less locking during concurent operations.
- Time complexity of maps is difficult to establish, I will state that it is the trivial O(1) in the best case,
although it isn't true.

### Settings
When creating a new client, you can choose the number of partitions and frequency of GC.
#### Garbage collector frequency
Garbage collector frequency, is a frequency, at which garbage collectoin happens.  
Every time it happens, it locks one partition after another, so starting it too oftem will be bad for any other operations.
Starting it too rarely, will leave a lot of unused (inaccessible) key-value pairs in the maps.  
Choose the setting wisely.
#### Number of partitions
Number of partitions is the number of maps in use.  
Best values are unresearched.

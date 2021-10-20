# big-o-notation-go

Examples of algorithms and explanation for each Big O Notation category.

Some examples are based in [this](https://www.youtube.com/watch?v=Mo4vesaut8g) video. If you didn't manage to understand Big O Notation yet, would be
an excellent choice to watch it.

## Running it
In order to run these scripts, you'll need to have [Go](https://golang.org/dl/) (preferably 1.16) installed.

You can run an example by typing `go run example.go` to any script in the [categories](categories) folder.

## Categories

### Constant time:
#### O(1)

An algorithm is said to be constant time (also written as O(1) time) if the value of T(n) 
is bounded by a value that does not depend on the size of the input. For example, accessing any
single element in an array takes constant time as only one operation has to be performed to locate it. 
In a similar manner, finding the minimal value in an array sorted in ascending order; it is the first element. ([source](https://en.wikipedia.org/wiki/Time_complexity#Constant_time))

Example: [constant_time.go](categories/constant_time.go)
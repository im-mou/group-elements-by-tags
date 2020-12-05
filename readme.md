# Sort elements by tags

This algorithm help group diferents elements with repeating tags into the minimum number of groups neceseary.

It can do this taking the tags with highest usage or the lowest.

The script takes an input file as input in which the first line indicates the number of tags in total and after that each line (index) belongs to the element (index) and it contains the tags folowed by the letter "G" (although it can be any letter).

ex.:<br/>
8<br/>
G1,G6<br/>
G1,G7<br/>
G1,G6,G8<br/>
G2,G7<br/>
G3,G6,G7<br/>
G3,G6,G8<br/>
G4,G7<br/>
G4,G8<br/>


To execute the script run the following command in you command line tool:

### Criterion: Max
```shell
go run ./main.go assignments.txt max
```

### Criterion: Min
```shell
go run ./main.go assignments.txt mix
```

## Result
The result should look something like this (max criterion):

It indicates the element number and the tag group it belongs to.

T1 -> G6<br/>
T2 -> G7<br/>
T3 -> G6<br/>
T4 -> G7<br/>
T5 -> G6<br/>
T6 -> G6<br/>
T7 -> G7<br/>
T8 -> G8<br/>
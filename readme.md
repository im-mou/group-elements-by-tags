# Sort tasks by tags

This algorithm help group diferents tasks with repeating tags into the minimum number of tag groups neceseary.

It can do this taking the tags with highest usage or the lowest.

The script takes an input file as input in which the first line indicates the number of tags in total and after that each line (index) belongs to the taks and it contains the tags folowed by the letter "G" (although it can be any letter).

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
The result should look something like this:

It indicates the task number and the group/tag group it belongs to.

T1 -> G1<br/>
T2 -> G1<br/>
T3 -> G1<br/>
T4 -> G2<br/>
T5 -> G3<br/>
T6 -> G3<br/>
T7 -> G4<br/>
T8 -> G4<br/>
# go-TW-ID-Generator
Taiwan ID Checker &amp; Generator

# Usage

```
% go run main.go --verify A123456789
A123456789 : PASS, time cost: 23.449µs

% go run main.go --verify A123456780
A123456780 : FAIL, time cost: 25.528µs

% time go run main.go --cpu 4
...
taskID: 0, taskCount:    1000000, totalCount:  510000000, timeCost: 57m56.195527594s, mem: 336 MB
taskID: 2, taskCount:    1000000, totalCount:  511000000, timeCost: 58m12.348436403s, mem: 286 MB
taskID: 1, taskCount:    1000000, totalCount:  512000000, timeCost: 58m17.794134191s, mem: 360 MB
taskID: 0, taskCount:    1000000, totalCount:  513000000, timeCost: 58m18.106069808s, mem: 380 MB
taskID: 2, taskCount:    1000000, totalCount:  514000000, timeCost: 58m33.170090521s, mem: 247 MB
taskID: 1, taskCount:    1000000, totalCount:  515000000, timeCost: 58m39.34318826s, mem: 308 MB
taskID: 0, taskCount:    1000000, totalCount:  516000000, timeCost: 58m39.902875471s, mem: 370 MB
taskID: 2, taskCount:    1000000, totalCount:  517000000, timeCost: 58m50.616989581s, mem: 133 MB
taskID: 1, taskCount:    1000000, totalCount:  518000000, timeCost: 58m55.655774417s, mem: 67 MB
taskID: 2, taskCount:    1000000, totalCount:  519000000, timeCost: 59m3.85291194s, mem: 111 MB
taskID: 2, taskCount:    1000000, totalCount:  520000000, timeCost: 59m14.91097006s, mem: 111 MB
Total verified IDs count: 520000000 , time cost: 59m18.802491746s
go run main.go --cpu 4  9320.09s user 3332.58s system 355% cpu 59:19.56 total

% ls -lah /tmp/ids*
-rw-------  1 user  wheel   1.4G  8 27 08:41 /tmp/ids-0.txt
-rw-------  1 user  wheel   1.4G  8 27 08:41 /tmp/ids-1.txt
-rw-------  1 user  wheel   1.4G  8 27 08:42 /tmp/ids-2.txt
-rw-------  1 user  wheel   1.0G  8 27 08:28 /tmp/ids-3.txt

% wc -l /tmp/ids-* 
 140000000 /tmp/ids-0.txt
 140000000 /tmp/ids-1.txt
 140000000 /tmp/ids-2.txt
 100000000 /tmp/ids-3.txt
 520000000 total
```

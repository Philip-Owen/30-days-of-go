# GO PROJECT #1 Tiny Port Status Reporter

Build a small CLI program that takes a list of network interfaces and prints a simple health report.

Start with the data hard-coded in the program.

example data:
```
Ethernet0  up    100000
Ethernet4  down  100000
Ethernet8  up    25000
Ethernet12 up    10000
```

Output should look like:
```
Interface Status Report
-----------------------
Ethernet0  : UP   - 100G
Ethernet4  : DOWN - 100G
Ethernet8  : UP   - 25G
Ethernet12 : UP   - 10G

4 interfaces checked
3 up
1 down
```

## Requirements

Your program should track each interface's:
- name
- operational status
- speed in Mbps
- Loop through the interfaces
- Print a human-readable status for each
- Count how many interfaces are up
- Count how many are down
- Convert speeds such as 100000 into 100G for display

## Stretch goal

Have the program print a warning when an interface is down:

```
Ethernet4 : DOWN - 100G [WARNING]
```
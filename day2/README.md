# GO PROJECT #2 Interface Utilization Monitor

You’re given a snapshot of traffic flowing through several switch interfaces. Your program should calculate each interface’s utilization percentage and classify its current load.

example input:
```
| Interface  | Capacity | Current traffic |
| ---------- | -------: | --------------: |
| Ethernet0  | 100 Gbps |         42 Gbps |
| Ethernet4  | 100 Gbps |         91 Gbps |
| Ethernet8  |  25 Gbps |         18 Gbps |
| Ethernet12 |  10 Gbps |          2 Gbps |
| Ethernet16 |  10 Gbps |        9.7 Gbps |

```
Use the same parallel-array approach

## Requirements
For every interface:

Calculate its utilization percentage.
1. Assign a state:
    - NORMAL: less than 70%
    - WARNING: 70% through 89%
    - CRITICAL: 90% or higher
2. Print a report showing:
    - interface name
    - utilization percentage
    - classification
3. At the bottom, print:
    - number of interfaces checked
    - number in NORMAL
    - number in WARNING
    - number in CRITICAL

example output:
```
Interface Utilization Report
----------------------------

Ethernet0   42.0%   NORMAL
Ethernet4   91.0%   CRITICAL
...
Number of interfaces checked: 5
2 NORMAL
```

## Stretch Goal

Find the most heavily utilized interface and print it's utilization

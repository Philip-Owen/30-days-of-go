# GO PROJECT #4 Build an Interface Inventory with Structs

Use a data model that actually represents an interface as one thing.

Start with the same general kind of network data:
example data
```
Ethernet0,up,100000
Ethernet4,down,100000
Ethernet8,up,25000
Ethernet12,up,10000
Ethernet16,down,10000
```

Your program should parse that raw text and turn each valid line into an interface record containing:

name
status
speed

Then store all of those records together and generate a report.

Your output should include each interface plus a summary showing:

total valid interfaces
interfaces up
interfaces down
total combined interface capacity

So with the sample data above, the total capacity would be 245 Gbps.
## Requirements

Specifically explore:
- defining your own struct
- creating values of that struct
- storing multiple struct values together
- accessing struct fields
- parsing raw text into structured data

Invalid rows should be skipped rather than crashing the program.

## Stretch Goal

Print all the names of interfaces that are currently down.

example
`Down interfaces: Ethernet4, Ethernet16`
# GO PROJECT #6 Interface State Change Detector

You have two snapshots of interface state: an older snapshot and a newer snapshot.

Example:
```
OLD:
Ethernet0,up,100000
Ethernet4,up,100000
Ethernet8,down,25000
Ethernet12,up,10000
Ethernet16,up,10000

NEW:
Ethernet0,up,100000
Ethernet4,down,100000
Ethernet8,up,25000
Ethernet12,up,10000
Ethernet20,up,25000
```

Compare them and produce a change report.

### What should you detect?

There are four interesting possibilities:

- Interface exists in both snapshots and nothing changed
- Interface exists in both snapshots but its status changed
- Interface existed before but is now missing
- Interface didn't exist before but is now new

## Report requirements

Print something like:
```
Interface Change Report
-----------------------

Ethernet4 changed: up -> down
Ethernet8 changed: down -> up
Ethernet16 removed
Ethernet20 added
```

Then a summary:
```
2 status changes
1 added
1 removed
2 unchanged
```

## Requirements

Use structs.

Parse each snapshot independently.

Use maps keyed by interface name for the comparison.

Sseparate functions for:
- parsing a snapshot into useful structured data
- comparing the two snapshots

### Important constraint

Don’t count a removed interface as a status change.

Likewise, an added interface isn’t “changed from nothing to up.” Treat added, removed, changed, and unchanged as distinct categories.

## Stretch Goal #1
Also detect speed changes.

If both status and speed change, report both.

## Stretch Goal #2

Create a struct representing a detected change.

Instead of immediately printing while comparing, build a collection of changes first and then have another part of the program print the report.
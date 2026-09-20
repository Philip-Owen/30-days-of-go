# GO PROJECT #5 Interface Inventory Validator

Take a raw interface inventory, parse it into structs, and then validate the inventory for problems.

raw input example: 
```
Ethernet0,up,100000
Ethernet4,down,100000
Ethernet8,up,25000
Ethernet12,up,10000
Ethernet4,up,100000
Ethernet16,down,10000
Ethernet20,up,25000
```

The program should parse valid rows into Interface structs just like Project #4

Then validate the resulting inventory and produce a report.

## Requirements

For each valid interface, print its:
- name
- normalized status
- human-readable speed

Then detect whether any interface name appears more than once.

Your summary should include:
- total valid interfaces
- total unique interface names
- number of duplicates found
- interfaces up
- interfaces down
- total capacity of unique interfaces only

That last requirement matters. If Ethernet4 appears twice, don't count its 100Gb capacity twice.

### Duplicate Behavior

When you encounter a duplicate, report it somehow. Example `Duplicate interface detected: Ethernet4`. The goal is simply to detect that the inventory contains conflicting entries.

## Stretch Goal

After processing everything, print the duplicated interface names only once. 

If this happened:
```
Ethernet4,up,100000
Ethernet4,down,100000
Ethernet4,up,100000
```

The final duplicate summary should still contain Ethernet4 only once

## Extra Stretch Goal

If a duplicate interface has conflicting information, call that out.

example:
```
Ethernet4,down,100000
Ethernet4,up,100000
```
could produce something like `Conflict detected for Ethernet4`

You decide what counts as a conflict: status, speed, or both.
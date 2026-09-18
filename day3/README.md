# GO PROJECT #3 Parse Interface Status Text

Your program gets a small block of interface status text that looks like command output:
```
Ethernet0,up,100000
Ethernet4,down,100000
Ethernet8,up,25000
Ethernet12,up,10000
Ethernet16,down,10000
```

Treat this as raw text input inside the program

Your job is to parse it and produce a summary report.

For each line, extract:

interface name
status
speed

Then print a report showing each interface and its parsed values, plus totals for:

interfaces checked
interfaces up
interfaces down

Also reuse the speed conversion idea from Project #1 so speeds display as human-readable values like 100Gb, 25Gb, and 10Gb.

## Requirements

Figure out how to:
- split the full text into individual lines
- split each line into its individual fields
- loop over the parsed data
- count interface states
- call at least one helper function

One important wrinkle: assume the input could contain an empty line somewhere. Your program should skip it instead of trying to parse it.

## Stretch Goal

Add one malformed line such as:
`Ethernet20,up`

The program should detect that it doesn't contain the expected number of fields and print something like:
`Skipping malformed line`

Then continue processing the rest of the input instead of crashing.
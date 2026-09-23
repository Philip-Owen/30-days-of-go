# GO PROJECT #7 Desired State Compliance Checker

Build on Project #6, but change the question. This project asks “Does the current state match what I intended?”.

Two datasets:
```
DESIRED STATE

Ethernet0,up,100000
Ethernet4,up,100000
Ethernet8,up,25000
Ethernet12,down,10000
Ethernet16,up,10000

ACTUAL STATE

Ethernet0,up,100000
Ethernet4,down,100000
Ethernet8,up,10000
Ethernet12,down,10000
Ethernet20,up,25000
```

Your program should determine whether each interface is compliant with the desired state.

### The possible outcomes

For this project, classify each interface into one of these categories:
- COMPLIANT — exists in both and matches desired status and speed
- STATUS_MISMATCH — interface exists, but status is wrong
- SPEED_MISMATCH — interface exists, but speed is wrong
- MULTIPLE_MISMATCHES — both status and speed differ
- MISSING — desired interface does not exist in actual state
- UNEXPECTED — actual interface exists but isn’t in desired state

So for the sample:

- Ethernet0 → compliant
- Ethernet4 → status mismatch
- Ethernet8 → speed mismatch
- Ethernet12 → compliant
- Ethernet16 → missing
- Ethernet20 → unexpected


### Main Report 

Print a report for every interface that appears in either dataset.
```
Interface Compliance Report
---------------------------

Ethernet0   COMPLIANT
Ethernet4   STATUS_MISMATCH
Ethernet8   SPEED_MISMATCH
Ethernet12  COMPLIANT
Ethernet16  MISSING
Ethernet20  UNEXPECTED
```

Then print a summary with counts for each category.

## Requirements

Reuse the parsing/data-model ideas you’ve already developed.

Use maps for comparing desired and actual state.

Assume the input is syntactically valid, but still normalize whitespace and status casing at the boundary.

Create some sort of structure representing the result of checking one interface. Don’t immediately print while comparing.

So the general flow should be:

parse → compare → build compliance results → print report

## Stretch Goal 1

For mismatch records, print the expected and actual values.

## Stretch Goal 2

At the end, print an overall result:

`COMPLIANCE CHECK: PASS`

only if every desired interface is compliant and there are no unexpected interfaces.

Otherwise:

`COMPLIANCE CHECK: FAIL`

## Stretch Goal 3

Calculate a compliance percentage based only on the desired interfaces.

If 3 out of 5 desired interfaces are fully compliant:

`Compliance: 60.0%`

Unexpected interfaces should not increase the denominator.
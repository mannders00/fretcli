# About
fretcli is an interactive fretboard note guessing quiz in the CLI.

It is simply an infinite loop of asking the user to guess the currently displayed note, and displaying a new one when the correct one is provided.

It can be configured with any variation of valid tuning, number of strings, and amount of frets displayed.

# Usage
```
go run main.go -h
Usage of fretcli:
  -frets int
        Amount of frets (default 12)
  -tuning string
        Guitar tuning and amount of strings (default "EADGBE")
```

# Guess Note
```
E ---| - | - | - | - | - | - | - | - | - | - | - | - |---
B ---| - | - | - | - | - | - | - | - | - | - | - | - |---
G ---| - | - | - | - | - | - | - | - | - | - | - | - |---
D ---| - | - | - | - | ? | - | - | - | - | - | - | - |---
A ---| - | - | - | - | - | - | - | - | - | - | - | - |---
E ---| - | - | - | - | - | - | - | - | - | - | - | - |---
               3       5       7       9          12

Enter Note:
```

```
E ---| - | - | - | - | - | - | - | - | - | - | - | - |---
B ---| - | - | - | - | - | - | - | - | - | - | - | - |---
G ---| - | - | - | - | - | - | - | - | - | - | - | - |---
D ---| - | - | - | - | G | - | - | - | - | - | - | - |---
A ---| - | - | - | - | - | - | - | - | - | - | - | - |---
E ---| - | - | - | - | - | - | - | - | - | - | - | - |---
               3       5       7       9          12

Correct!
```

# Configure Tuning / Fretboard Length
```
matt@mbp:~/Development/Sandbox/fretcli
$ ./fretcli -frets 7 -tuning ADADGBE

E       ---| - | - | - | - | - | - | - |---
B       ---| - | - | - | - | - | - | - |---
G       ---| - | - | - | - | - | - | - |---
D       ---| - | - | - | - | - | - | - |---
A       ---| - | - | - | - | ? | - | - |---
D       ---| - | - | - | - | - | - | - |---
A       ---| - | - | - | - | - | - | - |---
                     3       5       7
Enter note:
```

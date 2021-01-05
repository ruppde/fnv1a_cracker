fnv1a cracker in go. original version by @cybercdh from https://github.com/cybercdh/hacks/tree/master/sunburst

Performance is 4 times faster than the native go fnv1a lib on short strings. On large strings it's still 20% faster. Benchmarks in the bench folder.

This code does the XOR used in the sunburst malware in the end, so drop that if you want to use it for something else.

## Suggested Usage

Pipe the output of some other tool or file, e.g.

```cat processes.txt | go run main_optimized_lower_only.go```

or you can feed it a single process name on the command line:

```go run main_optimized_lower_only.go someprocess```

or you can feed it the output of another tool like e.g. john:

```john --incremental:LowerNum --stdout | go run main_optimized_lower_only.go```
(or prince mode, markov, ...)

Only feed lowercase strings.

## Example Output


```
14605870802367138151 : klifaa.sys
14076049386512171026 : cbk7.sys
6377259881933999545 : bhdrvx86.sys
18308913934889626123 : avgtpx86.sys
8144351048979905053 : bhdrvx64.sys
6976547106435497215 : avgtpx64.sys
15180155266747728342 : cyprotectdrv64.sys
13544031715334011032 : groundling64.sys
```

You need to ensure you have `hardcoded_hashes_left.txt` in the same dir as this code. This code was written without care for production quality or error checking.

Results by Team Hashcat in https://docs.google.com/spreadsheets/d/1u0_Df5OMsdzZcTkBDiaAtObbIOkMa5xbeXdKk_k0vWs/edit#gid=0


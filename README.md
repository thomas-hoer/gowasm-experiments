# go webassembly performance measurement

This is a performance comparison between the WebAssembly output of Go and plain JavaScript. It is based on a great example when people search for code to start their own work with GoLang and WASM. This repository compares the original [bouncy](https://stdiopt.github.io/gowasm-experiments/bouncy) example from stdiopt against a [plain JavaScript port](https://thomas-hoermann.de/gowasm-experiments).

The tests are done with Go 1.16.2, Chrome 89 and Windows 10 on a Intel Atom 1,6 GHz. 

## Running the local webserver

```sh
$ go run serve.go
```

## Test 1 (Go 1.13 vs Go 1.16)

The test compares the original wasm binaries based on Go 1.13 against assemblies of the same code compiled with Go 1.16.
All Tests use the following settings
- Speed 160
- Number of dots 1000
- Size 6
- Lines Unchecked
- Dashed Unchecked

## Results for Test 1

| Go Version | FPS |
| ---------- |----:|
| Go 1.13    | 3.7 |
| Go 1.16    | 4.6 |

We can se, that even on recompiling the same code with the latest compiler can improve the performance.

## Test 2 (WASM vs JS)

This test analyses the overhead that the communication between wasm and js costs. 
All the tests uses the same configuration as Test 1 and are compiled with Go 1.16.2. Other than before, the code that uses the cursor position was removed, so the results remain compareable.

## Results for Test 2

| Name              | FPS  | Duration in ms | Duration only syscall | syscall/js | n=1000 | time in ms per syscall/js | arguments | n=1000 | 2.5 * syscall+arg | time per arg|
| ----------------- | ---: | -------------: | --------------------- | ---------- | -----: | ------------------------: | --------- | -----: | --------------: | ----------: |
| wasm_only         |  5.1 |         196.08 |                155.92 | n*8+8      | 80008  |                     0.019 |    n*10+7 |  10007 |           30027 |   0.0051925 |
| wasm_reduce_calls |  5.4 |         185.19 |                145.03 | n*7+9      | 70009  |                     0.021 |     n*9+7 |   9007 |           26529 |   0.0054665 |
| helper_js         | 10.0 |         100.00 |                 59.84 | n*1+8      | 10008  |                     0.059 |     n*9+8 |   9008 |           11528 |   0.0051908 |
| helper2_js        | 10.3 |          97.09 |                 56.93 | n+1+8      | 10008  |                     0.056 |     n*8+7 |   8007 |           10527 |   0.0054077 |
| js_only           | 24.9 |          40.16 |                  0.00 | 0          |     0  |                     0.000 |         0 |      0 |               0 |   0.0000000 |

The results show that a function argument needs about 5 μs to be transfered and a function call or property needs about 12.5 μs.

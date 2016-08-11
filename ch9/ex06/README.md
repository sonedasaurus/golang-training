# Result 

## GOMAXPROCS=1 go test -bench .
```
PASS
BenchmarkDrawParallel	      30	  45189395 ns/op
ok  	_/Users/soneda/workspace/golang-training.bk/ch9/ex06	1.414s
PASS
BenchmarkDrawParallel	      30	  44662815 ns/op
ok  	_/Users/soneda/workspace/golang-training.bk/ch9/ex06	1.402s
PASS
BenchmarkDrawParallel	      30	  44604390 ns/op
ok  	_/Users/soneda/workspace/golang-training.bk/ch9/ex06	1.396s
```
## GOMAXPROCS=2 go test -bench .
```
PASS
BenchmarkDrawParallel-2	      50	  26784519 ns/op
ok  	_/Users/soneda/workspace/golang-training.bk/ch9/ex06	1.381s
PASS
BenchmarkDrawParallel-2	      50	  26503221 ns/op
ok  	_/Users/soneda/workspace/golang-training.bk/ch9/ex06	1.367s
PASS
BenchmarkDrawParallel-2	      50	  26620820 ns/op
ok  	_/Users/soneda/workspace/golang-training.bk/ch9/ex06	1.374s
```
## GOMAXPROCS=3 go test -bench .
```
PASS
BenchmarkDrawParallel-3	      50	  27175866 ns/op
ok  	_/Users/soneda/workspace/golang-training.bk/ch9/ex06	1.401s
PASS
BenchmarkDrawParallel-3	      50	  28444132 ns/op
ok  	_/Users/soneda/workspace/golang-training.bk/ch9/ex06	1.464s
PASS
BenchmarkDrawParallel-3	      50	  26164489 ns/op
ok  	_/Users/soneda/workspace/golang-training.bk/ch9/ex06	1.350s
```
## GOMAXPROCS=4 go test -bench .
```
PASS
BenchmarkDrawParallel-4	      50	  27813629 ns/op
ok  	_/Users/soneda/workspace/golang-training.bk/ch9/ex06	1.431s
PASS
BenchmarkDrawParallel-4	      50	  28128998 ns/op
ok  	_/Users/soneda/workspace/golang-training.bk/ch9/ex06	1.446s
PASS
BenchmarkDrawParallel-4	      50	  29164293 ns/op
ok  	_/Users/soneda/workspace/golang-training.bk/ch9/ex06	1.499s
```
## GOMAXPROCS=5 go test -bench .
```
PASS
BenchmarkDrawParallel-5	      50	  30081290 ns/op
ok  	_/Users/soneda/workspace/golang-training.bk/ch9/ex06	1.545s
PASS
BenchmarkDrawParallel-5	      50	  32209249 ns/op
ok  	_/Users/soneda/workspace/golang-training.bk/ch9/ex06	1.652s
PASS
BenchmarkDrawParallel-5	      50	  31266867 ns/op
ok  	_/Users/soneda/workspace/golang-training.bk/ch9/ex06	1.604s
```

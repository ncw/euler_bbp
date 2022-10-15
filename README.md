# Programs and paper to demonstrate Euler's BBP-type formulae for pi

This repository contains a paper about Euler's BBP-type forumulae
which can be used to compute binary digits from pi without calculating
the previous digits and programs calculate pi using them.

- `euler_bbp.py` - simple program to calculate hex digits of pi
- `euler_bbp.go` - faster (100x) program to do the same
- `check_bbp_equations.py` - program to check the equations in the paper work

Source for the paper

- `Makefile`
- `euler_bbp.tex`
- `refs.bib`

## Results

```
$ python euler_bbp.py
bbp_original         hex digits of pi from digit         10 5a308d31319 in 0.000s
bbp_euler            hex digits of pi from digit         10 5a308d31319 in 0.000s
bbp_euler2           hex digits of pi from digit         10 5a308d313198 in 0.000s
bbp_original         hex digits of pi from digit        100 c29b7c97c50 in 0.000s
bbp_euler            hex digits of pi from digit        100 c29b7c97c50 in 0.001s
bbp_euler2           hex digits of pi from digit        100 c29b7c97c50 in 0.000s
bbp_original         hex digits of pi from digit       1000 349f1c09b0 in 0.004s
bbp_euler            hex digits of pi from digit       1000 349f1c09b0 in 0.005s
bbp_euler2           hex digits of pi from digit       1000 349f1c09b0 in 0.002s
bbp_original         hex digits of pi from digit      10000 68ac8fcfb in 0.030s
bbp_euler            hex digits of pi from digit      10000 68ac8fcfb in 0.046s
bbp_euler2           hex digits of pi from digit      10000 68ac8fcfb in 0.019s
bbp_original         hex digits of pi from digit     100000 535ea16c in 0.340s
bbp_euler            hex digits of pi from digit     100000 535ea16c in 0.536s
bbp_euler2           hex digits of pi from digit     100000 535ea16c in 0.222s
bbp_original         hex digits of pi from digit    1000000 26c65e52 in 3.885s
bbp_euler            hex digits of pi from digit    1000000 26c65e5 in 6.115s
bbp_euler2           hex digits of pi from digit    1000000 26c65e52 in 2.627s
bbp_original         hex digits of pi from digit   10000000 17af586 in 42.293s
bbp_euler            hex digits of pi from digit   10000000 17af586 in 66.215s
bbp_euler2           hex digits of pi from digit   10000000 17af586 in 29.003s
bbp_original         hex digits of pi from digit  100000000 ecb840 in 463.304s
bbp_euler            hex digits of pi from digit  100000000 ecb840 in 752.472s
bbp_euler2           hex digits of pi from digit  100000000 ecb840 in 327.756s
bbp_original         hex digits of pi from digit 1000000000 85895 in 9379.175s
```

```
$ go run euler_bbp.go
  original hex digits of pi from digit           10 5a308d31319 in 26.87µs
     euler hex digits of pi from digit           10 5a308d31319 in 9.498µs
    euler2 hex digits of pi from digit           10 5a308d313198 in 18.084µs
  original hex digits of pi from digit          100 c29b7c97c50 in 19.587µs
     euler hex digits of pi from digit          100 c29b7c97c50 in 32.26µs
    euler2 hex digits of pi from digit          100 c29b7c97c50 in 27.682µs
  original hex digits of pi from digit         1000 349f1c09b0 in 248.003µs
     euler hex digits of pi from digit         1000 349f1c09b0 in 369.24µs
    euler2 hex digits of pi from digit         1000 349f1c09b0 in 131.676µs
  original hex digits of pi from digit        10000 68ac8fcfb in 2.174754ms
     euler hex digits of pi from digit        10000 68ac8fcfb in 3.45151ms
    euler2 hex digits of pi from digit        10000 68ac8fcfb in 1.37503ms
  original hex digits of pi from digit       100000 535ea16c in 26.643839ms
     euler hex digits of pi from digit       100000 535ea16c in 20.853328ms
    euler2 hex digits of pi from digit       100000 535ea16c in 16.504332ms
  original hex digits of pi from digit      1000000 26c65e52 in 55.597885ms
     euler hex digits of pi from digit      1000000 26c65e5 in 75.05774ms
    euler2 hex digits of pi from digit      1000000 26c65e52 in 160.25139ms
  original hex digits of pi from digit     10000000 17af586 in 480.289261ms
     euler hex digits of pi from digit     10000000 17af586 in 698.296118ms
    euler2 hex digits of pi from digit     10000000 17af586 in 362.285909ms
  original hex digits of pi from digit    100000000 ecb840 in 5.287894989s
     euler hex digits of pi from digit    100000000 ecb840 in 8.886897248s
    euler2 hex digits of pi from digit    100000000 ecb840 in 4.047852407s
  original hex digits of pi from digit   1000000000 85895 in 1m6.804017733s
     euler hex digits of pi from digit   1000000000 85895 in 1m41.747295359s
    euler2 hex digits of pi from digit   1000000000 85895 in 44.864847315s
  original hex digits of pi from digit  10000000000 921c in 13m16.405684459s
     euler hex digits of pi from digit  10000000000 921c in 19m46.698153514s
    euler2 hex digits of pi from digit  10000000000 921c in 8m36.587514529s
  original hex digits of pi from digit 100000000000 c9c in 2h28m30.473717759s
     euler hex digits of pi from digit 100000000000 c9c in 3h43m8.795869344s
    euler2 hex digits of pi from digit 100000000000 c9c in 1h36m58.833018519s
```

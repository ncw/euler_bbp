"""
This double checks the equations in the paper to check they actually
do calculate Pi
"""

import math
import gmpy2                    # for indefinite precision arithmetic

def bbp_original(n):
    """
    Produce a term of the original BBP formula
    """
    m = gmpy2.mpfr(16) ** (-n)
    return (
        + m * 4 / (8*n + 1)
        - m * 2 / (8*n + 4)
        - m * 1 / (8*n + 5)
        - m * 1 / (8*n + 6)
    )

def bbp_euler(n):
    """
    Produce a term of Euler's first BBP formula
    """
    m = gmpy2.mpfr(-4) ** (-n)
    return (
        + m * 2 / (4*n + 1)
        + m * 2 / (4*n + 2)
        + m * 1 / (4*n + 3)
    )

def bbp_euler2(n):
    """
    Produce a term of Euler's second BBP formula
    """
    m = (gmpy2.mpfr(-(2**30)) ** (-n)) / (2**26)
    return (
        + m * (2**27) / (20*n+1)
        + m * (2**25) / (12*n+1)
        + m * (2**25) / (10*n+1)
        + m * (2**24) / (20*n+3)
        + m * (2**22) / (6*n+1)
        - m * (2**20) / (60*n+15)
        - m * (2**19) / (10*n+3)
        - m * (2**18) / (20*n+7)
        - m * (2**15) / (12*n+5)
        + m * (2**15) / (20*n+9)
        + m * (2**12) / (30*n+15)
        + m * (2**12) / (20*n+11)
        - m * (2**10) / (12*n+7)
        - m * (2**9) / (20*n+13) 
        - m * (2**7) / (10*n+7)
        - m * (2**5) / (60*n+45)
        + m * (2**3) / (20*n+17)
        + m * (2**2) / (6*n+5)
        + m * 2 / (10*n+9)
        + m * 1 / (12*n+11)
        + m * 1 / (20*n+19)
    )

def test_series(term, precision):
    """
    Test the series defined by the function term to precision bits
    """
    gmpy2.get_context().precision = precision
    pi = gmpy2.const_pi()
    epsilon = gmpy2.mpfr(2)**(-int(precision*.99)+2) # must get to within this precision of pi
    name = term.__name__
    s = gmpy2.mpfr(0)
    multiplier = gmpy2.mpfr(1.0)
    for n in range(precision):
        s += term(n)
        #print(n,s)
        if abs(s-pi) < epsilon:
            print(f"{name} at precision {precision} bits: OK after {n} iterations")
            return
    print(f"{name} at precision {precision} bits: FAILED after {precision} iterations")
    
if __name__ == "__main__":
    for precision in (100,1000,10000,100000):
        test_series(bbp_original, precision)
        test_series(bbp_euler, precision)
        test_series(bbp_euler2, precision)

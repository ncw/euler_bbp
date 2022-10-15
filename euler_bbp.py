"""
This program calculates hexadecimal digits of Pi without calculating
the previous digits using the Euler BBP formulae.
"""

import math
import time

def term(x, k):
    """
    Calculate a term of the BBP series mod 1

    2**x / k mod 1
    """
    # If exponent is non-negative use modular arithmetic
    if x >= 0:
        return pow(2, x, k) / k
    # otherwise use normal floating point arithmetic
    return (2.0 ** x) / k

def bbp_original(m, n):
    """
    Produce a term of the original BBP formula multiplied by 2**m and
    mod 1
    """
    exponent = m - 4*n
    return (
        + term(exponent + 2, 8*n + 1)
        - term(exponent + 1, 8*n + 4)
        - term(exponent, 8*n + 5)
        - term(exponent, 8*n + 6)
    ) % 1

def bbp_euler(m, n):
    """
    Produce a term of Euler's first BBP formula multiplied by 2**m and
    mod 1
    """
    exponent = m - 2*n
    s = (
        + term(exponent + 1, 4*n + 1)
        + term(exponent + 1, 4*n + 2)
        + term(exponent, 4*n + 3)
    ) % 1
    if n % 2 != 0:
        s = -s
    return s

def bbp_euler2(m, n):
    """
    Produce a term of Euler's second BBP formula multiplied by 2**m
    and mod 1
    """
    exponent = m - 30*n - 26
    s = (
        + term(exponent + 27, 20*n + 1)
        + term(exponent + 25, 12*n + 1)
        + term(exponent + 25, 10*n + 1)
        + term(exponent + 24, 20*n + 3)
        + term(exponent + 22, 6*n + 1)
        - term(exponent + 20, 60*n + 15)
        - term(exponent + 19, 10*n + 3)
        - term(exponent + 18, 20*n + 7)
        - term(exponent + 15, 12*n + 5)
        + term(exponent + 15, 20*n + 9)
        + term(exponent + 12, 30*n + 15)
        + term(exponent + 12, 20*n + 11)
        - term(exponent + 10, 12*n + 7)
        - term(exponent + 9, 20*n + 13) 
        - term(exponent + 7, 10*n + 7)
        - term(exponent + 5, 60*n + 45)
        + term(exponent + 3, 20*n + 17)
        + term(exponent + 2, 6*n + 5)
        + term(exponent + 1, 10*n + 9)
        + term(exponent, 12*n + 11)
        + term(exponent, 20*n + 19)
    )
    if n % 2 != 0:
        s = -s
    return s

def bbp(m, formula, bits_per_iteration, terms):
    """
    Calculate the m-th hex digit of pi using the formula passed in.

    bits_per_iteration is used to calculate the stopping point and
    terms is used to estimate the error.
    """
    t0 = time.time()
    name = formula.__name__
    # python floats have this many bits of precision
    precision = 53
    # this means
    assert 1.0 + 2.0 **(-precision) == 1.0
    # Stop adding terms when they are less than this
    s = 0.0
    # start from the m-th hex digit counting from 1 not 0
    mh = 4*(m-1)

    # Calculate the number of iterations needed
    iterations = (mh + precision)//bits_per_iteration + 1
    for n in range(iterations+1):
        term = formula(mh, n)
        s = (s + term) % 1
        #print(n,s,term)
    dt = time.time() - t0
    # Estimate the accuracy
    #
    # We've done n * terms additions. We might expect to have
    # an error of 1/2 * (2**-precision) per addition here which would
    # make a crude estimate of the maximum error to be:
    max_error = n * terms / 2 * (2**-precision)
    ok_bits = -math.log2(max_error)
    ok_digits = int(ok_bits / 4)
    digits = int(s * (2**(4*ok_digits)))
    print(f"{name:20} hex digits of pi from digit {m:9} {digits:0{ok_digits}x} in {dt:.3f}s")
    
if __name__ == "__main__":
    for log_m in range(1,10):
        m = 10**log_m
        bbp(m, bbp_original, 4, 4)
        bbp(m, bbp_euler, 2, 3)
        bbp(m, bbp_euler2, 30, 21)

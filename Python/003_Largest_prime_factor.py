number = 600851475143


def isPrime(n):
    for i in range(2, n):
        if n % i == 0 and i != n:
            return False
    return True


def LargestPrimeFactor(n):
    _n, lpf = n, 0
    for i in range(2, n):
        if isPrime(i) and n % i == 0:
            _n, lpf = _n/i, i
            print(i, end=', ')
        if _n == 1:
            break
    print(lpf)
    return


LargestPrimeFactor(number)

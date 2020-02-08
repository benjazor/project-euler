def fibo(n1, n2, res):
    if n1 > 4000000:
        return res
    return fibo(n2, n2 + n1, res + n2 if n2 % 2 == 0 else res)

print( fibo(1, 1, 0) )
print(4613732)

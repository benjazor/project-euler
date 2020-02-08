def is_palindrome(n):
    # Reverse the string
    return str(n) == str(n)[::-1]


big = 0


for i in range(1, 1000):
    for j in range(i, 1000):
        x = i * j
        big = x if is_palindrome(x) and x > big else big


print(big)


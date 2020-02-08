'''
2520 is the smallest number that can be divided by
each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly
divisible by all of the numbers from 1 to 20?
'''

# VERSION 1 : Works 'great' for 10 but too slow for 20
def smallest_multiple(number):
    mult = 0


    while True:
        print(mult)
        mult += 1
        isOk = True
        for i in range(1, number):
            if mult % i != 0:
                isOk = False
                break
        if isOk:
            break

    return mult


# print( smallest_multiple(20) )

def isPrime(number):
    for i in range(2, number):
        if number % i == 0:
            return False
    return True

# Version 2: Uses prime numbers, works fine for 20 and 100 but 1000 is too slow
def smallestMultiple(number):
    # Take all the prime numbers
    primeNumbers = []
    for i in range(2, number):
        if isPrime(i):
            primeNumbers.append(i)

    # Multiply all the prime numbers
    multiple = 1
    for primeNumber in primeNumbers:
        multiple *= primeNumber

    # Find the smallest multiple
    i = 0
    while True:
        i += 1
        newMultiple = multiple * i

        # Check if the number is a multiple of all numbers bellow
        isOk = True
        for j in range(1, number):
            if newMultiple % j != 0:
                isOk = False
                break
        if isOk:
            return newMultiple

print( smallestMultiple(1000) )

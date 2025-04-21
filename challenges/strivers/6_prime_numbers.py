#  Given an integer N, check whether it is prime or not.
# A prime number is a number that is only divisible by 1 and itself and the total number
# of divisors is 2

def primeNumbers(n: int) -> bool:
    return len([num for num in range(1, n + 1) if n % num == 0]) == 2


if __name__ == "__main__":
    print(primeNumbers(23))

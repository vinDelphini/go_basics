# Given an integer N, return the number of digits in N.

def countDigits(num: int) -> int:

    if 10 < num >= 0 :
        return 1

    if num < 0:
        num = -num

    count = 0
    while num > 0:
        count += 1
        num //= 10

    return count

if __name__=="__main__":
    print(countDigits(12345))

# Given an integer N, return true if it is a palindrome else return false

def isPalindrome(num: int) -> int:
    out = 0
    temp_num = num
    while num > 0:
        digit = num % 10
        out = out * 10 + digit
        num //= 10
    print(out)
    return temp_num == out


if __name__ == "__main__":
    print(isPalindrome(121))

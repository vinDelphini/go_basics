# Given an integer N, return true it is an Armstrong number otherwise return false.

def armstrong(num: int) -> bool:

    out = 0
    temp_num = num
    power = len(str(num))

    while num > 0:
        digit = num % 10
        out += digit ** power
        num //= 10

    return out == temp_num

if __name__ == "__main__":
    print(armstrong(153))

def bubble_sort(arr):
    n = len(arr)
    for i in range(n):
        for j in range(0, n-i-1):
            if arr[j] > arr[j+1]:
                arr[j], arr[j+1] = arr[j+1], arr[j]

def main():
    print("Enter the number of elements:")
    n = int(input())

    arr = []
    print("Enter the elements (one per line):")
    for _ in range(n):
        element = int(input())
        arr.append(element)

    bubble_sort(arr)

    print("Sorted array:")
    for i in arr:
        print(i)

if __name__ == "__main__":
    main()


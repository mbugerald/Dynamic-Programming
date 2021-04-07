def binary_search(numbers, target):
    
    first = 0   # Constant time value assigment.
    last = len(numbers) - 1  # Constant time value assigment.

    # What causes the algorithm to grow
    while first <= last:
        
        # Constant time
        midpoint = (first + last)//2  # round down to the nearest whole number.

        # Constant time
        if numbers[midpoint] == target:
            return midpoint
        
        # Constant time
        elif numbers[midpoint] < target:
            first = midpoint + 1
        
        # Constant time
        else: 
            last = midpoint - 1

    return None


def verify_binary_search(value):
	if value is not None:
		print("Target found at index: ", value)
	else:
		print("Target not found in the numbers.")

numbers = [1,2,3,4,5,6,7,8,9,10]

result = binary_search(numbers, 12)
verify_binary_search(result)

result = binary_search(numbers, 6)
verify_binary_search(result)
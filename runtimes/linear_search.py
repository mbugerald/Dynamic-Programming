def linear_search(list, target):
	"""
	Return the index position of the target if found, else returns None
	"""

	for i in range(0, len(list)):
		if list[i] == target:
			return i
	return None

def verify_linear_search(value):
	if value is not None:
		print("Target found at index: ", value)
	else:
		print("Target not found in the list.")

numbers = [1,2,3,4,5,6,7,8,9,10]

result = linear_search(numbers, 12)
verify_linear_search(result)

result = linear_search(numbers, 6)
verify_linear_search(result)
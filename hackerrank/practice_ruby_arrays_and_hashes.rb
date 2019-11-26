##
# Initialization
array = Array.new
array_1 = Array.new(1)
array_2 = Array.new(2, 10)

##
# Index, Part 1
def element_at(arr, index)
  # return the element of the Array variable `arr` at the position `index`
  # arr.at(index) # or
  # arr[index]
  arr[index]
end

def inclusive_range(arr, start_pos, end_pos)
  # return the elements of the Array variable `arr` between the start_pos and end_pos (both inclusive)
  arr[start_pos, end_pos]
end

def non_inclusive_range(arr, start_pos, end_pos)
  # return the elements of the Array variable `arr`, start_pos inclusive and end_pos exclusive
  arr[start_pos...end_pos]
end

def start_and_length(arr, start_pos, length)
  # return `length` elements of the Array variable `arr` starting from `start_pos`
  arr[start_pos, length]
end

##
# Index, Part 2
def neg_pos(arr, index)
  # return the element of the array at the position `index` from the end of the list
  # Clue : arr[-index]
  arr[-index]
end

def first_element(arr)
  # return the first element of the array
  # arr.first
  arr.first
end

def last_element(arr)
  # return the last element of the array
  # arr.last
  arr.last
end

def first_n(arr, n)
  # return the first n elements of the array
  arr.first(n)
end

def drop_n(arr, n)
  # drop the first n elements of the array and return the rest
  arr.drop(n)
end

##
# Addition
def end_arr_add(arr, element)
  # Add `element` to the end of the Array variable `arr` and return `arr`
  arr.push(element)
end

def begin_arr_add(arr, element)
  # Add `element` to the beginning of the Array variable `arr` and return `arr`
  arr.unshift(element)
end

def index_arr_add(arr, index, element)
  # Add `element` at position `index` to the Array variable `arr` and return `arr`
  arr.insert(index, element)
  arr
end

def index_arr_multiple_add(arr, index)
  # add any two elements to the arr at the index
  arr.insert(index, 8, 9)
end

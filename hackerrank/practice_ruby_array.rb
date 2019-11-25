##
# Deletion
def end_arr_delete(arr)
  # delete the element from the end of the array and return the deleted element
  arr.pop
end

def start_arr_delete(arr)
  # delete the element at the beginning of the array and return the deleted element
  arr.shift
end

def delete_at_arr(arr, index)
  # delete the element at the position #index
  arr.delete_at(index)
end

def delete_all(arr, val)
  # delete all the elements of the array where element = val
  arr.delete(val)
end

##
# Selection
def select_arr(arr)
  # select and return all odd numbers from the Array variable `arr`
  arr.select {|a| a.odd?}
end

def reject_arr(arr)
  # reject all elements which are divisible by 3
  arr.select {|a| a % 3 != 0}
end

def delete_arr(arr)
  # delete all negative elements
  arr.select {|a| a >= 0}
end

def keep_arr(arr)
  # keep all non negative elements ( >= 0)
  arr.select {|a| a >= 0}
end

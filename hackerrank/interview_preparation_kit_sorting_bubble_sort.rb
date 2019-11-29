##
# Bubble Sort
#
#!/bin/ruby

require 'json'
require 'stringio'

# Complete the countSwaps function below.
def countSwaps(a)
  valid_a = a.sort
  swap_cnt = 0

  until a == valid_a do
    a.reverse_each.with_index do |elem, i|
      i = a.size - i - 1

      if elem > valid_a[i] && elem > a[i+1]
        a = a - [elem]
        a.insert(i+1, elem)
        swap_cnt += 1
      end
    end
  end

  puts "Array is sorted in #{swap_cnt} swaps."
  puts "First Element: #{a[0]}"
  puts "Last Element: #{a[-1]}"
end

n = gets.to_i

a = gets.rstrip.split(' ').map(&:to_i)

countSwaps a

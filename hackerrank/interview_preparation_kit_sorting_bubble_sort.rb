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

  until a == valid_a
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

##
# Mark and Toys
#
#!/bin/ruby

require 'json'
require 'stringio'

# Complete the maximumToys function below.
def maximumToys(prices, k)
  result = []

  buyable_prices = prices.select { |price| k >= price }
  1.upto(buyable_prices.size) do |i|
    result += buyable_prices
      .each_cons(i)
      .to_a
      .select { |c| k >= c.sum }
  end

  result.sort_by { |r| r.size }.last.size
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

nk = gets.rstrip.split

n = nk[0].to_i

k = nk[1].to_i

prices = gets.rstrip.split(' ').map(&:to_i)

result = maximumToys(prices, k)

fptr.write(result)
fptr.write("\n")

fptr.close()

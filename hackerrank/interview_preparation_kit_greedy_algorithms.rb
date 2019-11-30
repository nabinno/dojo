##
# Greedy Florist
#
#!/bin/ruby

require 'json'
require 'stringio'

# Complete the getMinimumCost function below.
def getMinimumCost(k, c)
  flowers = c.sort.reverse
  previous_purchases = Hash.new{0}
  total_price = 0

  flowers.each_slice(k).to_a.each do |sliced_flowers|
    sliced_flowers.each_with_index do |flower_price, j|
      total_price += (1 + previous_purchases[j]) * flower_price
      previous_purchases[j] += 1
    end
  end

  total_price
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

nk = gets.rstrip.split

n = nk[0].to_i

k = nk[1].to_i

c = gets.rstrip.split(' ').map(&:to_i)

minimumCost = getMinimumCost(k, c)

fptr.write(minimumCost)
fptr.write("\n")

fptr.close()

##
# 2D Array-DS
#
#!/bin/ruby

require 'json'
require 'stringio'

# Complete the hourglassSum function below.
def hourglassSum(arr)
  rc_arr = [
    [0, 0, 0, 0],
    [0, 0, 0, 0],
    [0, 0, 0, 0],
    [0, 0, 0, 0]
  ]

  arr.each_with_index do |row, i|
    row.each_with_index do |cell, j|
      rc_arr[i][j] += cell if (0..3).include?(j) && (0..3).include?(i)
      rc_arr[i][j-1] += cell if (0..3).include?(j-1) && (0..3).include?(i)
      rc_arr[i][j-2] += cell if (0..3).include?(j-2) && (0..3).include?(i)

      rc_arr[i-1][j-1] += cell if (0..3).include?(j-1) && (0..3).include?(i-1)

      rc_arr[i-2][j] += cell if (0..3).include?(j) && (0..3).include?(i-2)
      rc_arr[i-2][j-1] += cell if (0..3).include?(j-1) && (0..3).include?(i-2)
      rc_arr[i-2][j-2] += cell if (0..3).include?(j-2) && (0..3).include?(i-2)
    end
  end

  rc_arr.flatten.max
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

arr = Array.new(6)

6.times do |i|
  arr[i] = gets.rstrip.split(' ').map(&:to_i)
end

result = hourglassSum(arr)

fptr.write(result)
fptr.write("\n")

fptr.close()

##
# Arrays: Left Rotation
#
#!/bin/ruby

require 'json'
require 'stringio'

# Complete the rotLeft function below.
def rotLeft(a, d)
  size = a.size
  d.times do
    first = a.shift
    a.insert(size-1, first)
  end
  a
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

nd = gets.rstrip.split

n = nd[0].to_i

d = nd[1].to_i

a = gets.rstrip.split(' ').map(&:to_i)

result = rotLeft(a, d)

fptr.write(result.join(" "))
fptr.write("\n")

fptr.close()

##
# New Year Chaos
#
#!/bin/ruby

require 'json'
require 'stringio'

# Complete the minimumBribes function below.
def minimumBribes(q)
  moves = 0
  r = (0..q.size-1).to_a

  until q == (1..q.size).to_a do
    q.map { |a| a - 1 }
      .reverse_each.with_index do |person, i|
      i = q.size - i - 1
      if person - i > 2
        puts "Too chaotic"
        return
      end

      if person > r[i] && person > q[person] -1
        moves += person - r[i]
        q = q - [person + 1]
        q.insert(person, person + 1)
      end
    end
  end

  puts moves
end

t = gets.to_i

t.times do |t_itr|
  n = gets.to_i

  q = gets.rstrip.split(' ').map(&:to_i)

  minimumBribes(q)
end

##
# Minimum Swaps 2
#
#!/bin/ruby

require 'json'
require 'stringio'

# Complete the minimumSwaps function below.
def minimumSwaps(arr)
  swap_cnt = 0
  valid_arr = (1..arr.size).to_a

  until arr == valid_arr do
    arr.each_with_index do |elem, i|
      if elem > valid_arr[i]
        target = arr.index(i+1)

        arr[i] = valid_arr[i]
        arr[target] = elem
        swap_cnt += 1
      end
    end
  end

  swap_cnt
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

n = gets.to_i

arr = gets.rstrip.split(' ').map(&:to_i)

res = minimumSwaps(arr)

fptr.write(res)
fptr.write("\n")

fptr.close()

##
# Array Manipulation
#
#!/bin/ruby

require 'json'
require 'stringio'

# Complete the arrayManipulation function below.
def arrayManipulation(n, queries)
  arr = Array.new(n, 0)

  queries.each do |query|
    left = query[0] - 1
    right = query[1] - 1
    summand = query[2]

    arr[(left..right)] = arr[(left..right)].map { |a| a + summand }
  end

  arr.max
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

nm = gets.rstrip.split

n = nm[0].to_i

m = nm[1].to_i

queries = Array.new(m)

m.times do |i|
  queries[i] = gets.rstrip.split(' ').map(&:to_i)
end

result = arrayManipulation(n, queries)

fptr.write(result)
fptr.write("\n")

fptr.close()

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
  bribes_cnt = 0;
  min = q.size;

  q.size.times do |i|
    if (q[i] - i > 3)
      puts 'Too chaotic'
      break
    end

    if (q[i] > i+1)
      bribes_cnt += (q[i]-(i+1))
      # puts "[#{i}] #{bribes_cnt}, #{min}" if min < q.size
      puts bribes_cnt if min < q.size
    else 
      if (min > q[i])
        min = q[i]
      elsif (q[i] != min)
        bribes_cnt += 1
      end

      # puts "#{bribes_cnt}, #{min}"
      # break
    end
  end
end

t = gets.to_i

t.times do |t_itr|
  n = gets.to_i

  q = gets.rstrip.split(' ').map(&:to_i)

  minimumBribes(q)
end

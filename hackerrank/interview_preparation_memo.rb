##
# Count words 
#
require 'json'
require 'stringio'

def countWords(ss)
  rc = Hash.new{0}

  ss.each do |s|
    words = s.split(" ")

    words.each do |word|
      rc[word] += 1 if /[A-Z0-9]/.match(word[0])
    end
  end

  rc.keys.size
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

ss = []

2.times do
  ss += [gets.to_s.rstrip]
end

result = countWords(ss)

fptr.write(result)
fptr.write("\n")

fptr.close()

##
# Remove integer from array
#
require 'json'
require 'stringio'

def callfunc(as)
  rc = as.dup
  shifted_a = nil
  shifted_cnt = 0
  skip = 0
  
  as.each_with_index do |a, index|
    if index % 2 == 0 && a != shifted_a
      rc.delete_at(index+shifted_cnt+skip)
      shifted_cnt -= 1
    elsif a == shifted_a
      skip += 1
    end

    shifted_a = a
  end

  rc.join(",")
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

as = gets.rstrip.split(',').map(&:to_i)

result = callfunc(as)

fptr.write(result)
fptr.write("\n")

fptr.close()

##
# Average
#
require 'json'
require 'stringio'

def callfunc(csv)
  rc = [0, 0, 0, 0]

  csv.each_with_index do |row,i|
    next if i == 0

    rc = rc.map.with_index { |a,j| a += row[j].to_i }
  end

  rc.map { |a| a/4.to_f.round }
end

csv = [
  ["hello", "hello", "hi", "world"],
  [20, 30, 11, 80],
  [10, 32, 18, 40],
  [22, 34, 30, 40],
  [23, 39, 12, 90],
]
result = callfunc(csv)

puts csv[0].join(",")
puts result.join(",")

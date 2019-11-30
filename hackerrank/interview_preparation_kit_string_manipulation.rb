##
# Making Anagrams
#
#!/bin/ruby

require 'json'
require 'stringio'

# Complete the makeAnagram function below.
def makeAnagram(a, b)
  dict = Hash.new{0}
  cnt = 0

  a.chars.each {|n| dict[n] += 1}
  b.chars.each {|n| dict[n] -= 1}
  
  dict.values.each {|v| cnt += v.abs}
  cnt
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

a = gets.to_s.rstrip

b = gets.to_s.rstrip

res = makeAnagram(a, b)

fptr.write(res)
fptr.write("\n")

fptr.close()

##
# Alternating Characters
#
#!/bin/ruby

require 'json'
require 'stringio'

# Complete the alternatingCharacters function below.
def alternatingCharacters(s)
  previous_char = ""
  cnt = 0

  s.chars do |c|
    if c == previous_char
      cnt += 1
    end

    previous_char = c
  end

  cnt
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

q = gets.to_i

q.times do |q_itr|
  s = gets.to_s.rstrip

  result = alternatingCharacters s

  fptr.write result
  fptr.write "\n"
end

fptr.close()

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

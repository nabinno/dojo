##
# Sock Merchant
#
#!/bin/ruby

require 'json'
require 'stringio'

# Complete the sockMerchant function below.
def sockMerchant(n, ar)
  ar.group_by {|b| b}
    .values
    .map {|b| b.size/2}
    .sum
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

n = gets.to_i

ar = gets.rstrip.split(' ').map(&:to_i)

result = sockMerchant(n, ar)

fptr.write(result)
fptr.write("\n")

fptr.close()

##
# Couting Valleys
#
#!/bin/ruby

require 'json'
require 'stringio'

# Complete the countingValleys function below.
def countingValleys(n, s)
  high = 0
  valley = 0

  s.split("")
    .map(&lambda do |t|
       case t
       when "U" then 1
       when "D" then -1
       end
     end)
    .each(&lambda do |i|
       if high == 0
         high = high + i
         valley += 1 if high < 0
       else
         high = high + i
       end
     end)
  valley
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

n = gets.to_i

s = gets.to_s.rstrip

result = countingValleys n, s

fptr.write result
fptr.write "\n"

fptr.close()

##
# Jumping on the Clouds
#
#!/bin/ruby

require 'json'
require 'stringio'

# Complete the jumpingOnClouds function below.
def jumpingOnClouds(c)
  rc = 0
  jumpable = 0
  c.each_with_index do |d,i|
    jumpable = if d == 0
                 jumpable < 3 ? jumpable + 1 : 2
               else
                 0
               end
    rc += 1 if jumpable < 3 && d == 0 && i != 0
  end
  rc
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

n = gets.to_i

c = gets.rstrip.split(' ').map(&:to_i)

result = jumpingOnClouds(c)

fptr.write(result)
fptr.write("\n")

fptr.close()

##
# Repeated String
#
#!/bin/ruby

require 'json'
require 'stringio'

# Complete the repeatedString function below.
def repeatedString(s, n)
  counting_s = "a"
  return n if s.size == 1 && s == counting_s

  coefficient = n / s.size
  surplus = n % s.size

  s.count(counting_s) * coefficient + s[0,surplus].count(counting_s)
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

s = gets.to_s.rstrip

n = gets.to_i

result = repeatedString(s, n)

fptr.write(result)
fptr.write("\n")

fptr.close()


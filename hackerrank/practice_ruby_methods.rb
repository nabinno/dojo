##
# Arguments
def take(arr, ind=1)
  arr[ind,arr.size]
end

##
# Variable Arguments
def full_name(*args)
  args.join(" ")
end

##
# Keyword Arguments
def convert_temp(temp, input_scale: 'celsius', output_scale: 'celsius')
  case input_scale
  when 'celsius'
    case output_scale
    when 'celsius' then temp
    when 'kelvin' then temp + 273.15
    when 'fahrenheit' then (temp * 9/5.to_f) + 32
    end
  when 'kelvin'
    case output_scale
    when 'celsius' then temp - 273.15
    when 'kelvin' then temp
    when 'fahrenheit' then (temp - 273.15) * 9/5.to_f + 32
    end
  when 'fahrenheit'
    case output_scale
    when 'celsius' then (temp - 32) * 5/9.to_f
    when 'kelvin' then (temp - 32) * 5/9.to_f + 273.15
    when 'fahrenheit' then temp
    end
  end
end

##
# Blocks
def factorial(n)
  rc = (1..n).reduce(1, :*)
  yield(rc)
end

n = gets.to_i
factorial(n) do |rc|
  puts rc
end

##
# Proc
def square_of_sum (my_array, proc_square, proc_sum)
    sum = proc_sum.call(my_array)
    proc_square.call(sum)
end

proc_square_number = proc {|num| num**2}
proc_sum_array = proc {|arr| arr.sum}
my_array = gets.split().map(&:to_i)

puts square_of_sum(my_array, proc_square_number, proc_sum_array)

##
# Lambdas
square = lambda {|num| num**2} # Write a lambda which takes an integer and square it
plus_one = lambda {|num| num + 1} # Write a lambda which takes an integer and increment it by 1
into_2 = lambda {|num| num * 2} # Write a lambda which takes an integer and multiply it by 2
adder = lambda {|*num| num.sum} # Write a lambda which takes two integers and adds them
values_only = lambda {|h| h.values} # Write a lambda which takes a hash and returns an array of hash values

input_number_1 = gets.to_i
input_number_2 = gets.to_i
input_hash = eval(gets)

a = square.(input_number_1); b = plus_one.(input_number_2);c = into_2.(input_number_1); 
d = adder.(input_number_1, input_number_2);e = values_only.(input_hash)

p a; p b; p c; p d; p e

##
# Closures
def block_message_printer
  message = "Welcome to Block Message Printer"
  if block_given?
    yield
  end
  puts "But in this function/method message is :: #{message}"
end
message = gets
block_message_printer { puts "This message remembers message :: #{message}" }

def proc_message_printer(my_proc)
  message = "Welcome to Proc Message Printer"
  my_proc.call
  puts "But in this function/method message is :: #{message}"
end
my_proc = proc { puts "This message remembers message :: #{message}" }
proc_message_printer(my_proc)

def lambda_message_printer(my_lambda)
  message = "Welcome to Lambda Message Printer"
  my_lambda.call
  puts "But in this function/method message is :: #{message}"
end
my_lambda = -> { puts "This message remembers message :: #{message}" }
lambda_message_printer(my_lambda)    

##
# Partial Applications
combination = lambda do |n|
  lambda do |r|
    a, b = r, n-r
    a, b = b, a if a < b  # a is the larger
    numer = (a+1..n).inject(1) { |t,v| t*v }  # n!/r!
    denom = (2..b).inject(1) { |t,v| t*v }    # (n-r)!
    numer/denom
  end
end
n = gets.to_i
r = gets.to_i
nCr = combination.(n)
puts nCr.(r)

##
# Currying
power_function = -> (x, z) {
  (x) ** z
}

base = gets.to_i
raise_to_power = power_function.curry.(base)

power = gets.to_i
puts raise_to_power.(power)

##
# Introduction
class Object
  require 'prime'

  def prime?(num)
    Prime.prime?(num)
  end
end

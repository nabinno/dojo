##
# Lazy Evaluation
require 'prime'
N = gets.to_i
rc = []
Prime.each do |pr|
  rc << pr if pr.to_s == pr.to_s.reverse
  break if rc.size >= N
end
puts "#{rc}"

##
# each_with_index
def skip_animals(animals, skip)
  animals.map.with_index(&lambda do |a,i|
            next if i < skip
            "#{i}:#{a}"
          end)
         .compact
end

##
# collect
def rot13(secret_messages)
  secret_messages.collect(&lambda do |msg|
                    msg.split("")
                       .collect(&lambda do |ch|
                          if /^[a-z]$/ === ch
                            ((ch.ord + 13 - 'a'.ord) % 26 + 'a'.ord).chr
                          elsif /^[A-Z]$/ === ch
                            ((ch.ord + 13 - 'A'.ord) % 26 + 'A'.ord).chr
                          else
                            ch
                          end
                        end)
                       .join("")
                 end)
end

##
# reduce
def sum_terms(n)
  return unless n >= 0

  (1..n).map {|x| x**2+1 }
        .reduce(0,:+)
end

##
# 'any', 'all', 'none', and 'find'
def func_any(hash)
  # Check and return true if any key object within the hash is of the type Integer
  # If not found, return false.
  hash.any? {|k,v| k.is_a?(Integer)}
end

def func_all(hash)
  # Check and return true if all the values within the hash are Integers and are < 10
  # If not all values satisfy this, return false.
  hash.all? {|k,v| v.is_a?(Integer) && v < 10}
end

def func_none(hash)
  # Check and return true if none of the values within the hash are nil
  # If any value contains nil, return false.
  hash.none? {|k,v| v.nil?}
end

def func_find(hash)
  # Check and return the first object that satisfies either of the following properties:
  #   1. There is a [key, value] pair where the key and value are both Integers and the value is < 20 
  #   2. There is a [key, value] pair where the key and value are both Strings and the value starts with `a`.
  hash.find do |k,v|
    b1 = k.is_a?(Integer) && v.is_a?(Integer) && v < 20
    b2 = k.is_a?(String) && v.is_a?(String) && v[0] == "a"
    b1 || b2
  end
end

##
# group_by
def group_by_marks(marks, pass_marks)
  marks.group_by {|k,v| v >= pass_marks ? "Passed" : "Failed"}
end

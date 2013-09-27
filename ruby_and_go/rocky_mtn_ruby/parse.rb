"string".to_i  # 0

class MyClass; end

puts MyClass.new.to_i # NoMethodError

# m is true
m = MyClass.new
if m
  puts "m is true"
else
  puts "m is false"
end

class Swan
  def quack
    puts "Swan quack!"
  end
end

class Duck
  def quack
    puts "Duck quack!"
  end
end

class Chicken
  def quack
    puts "Chicken quack!"
  end
end

[Swan.new, Duck.new, Chicken.new].each do |bird|
  bird.quack
end

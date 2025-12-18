def worker(name)
  puts "Worker #{name} starting"

  sleep 2
  
  puts "Worker #{name} done"
end

threads = []
["A", "B", "C"].each do |name|
  threads << Thread.new { worker(name) }
end

threads.each(&:join)
def worker(name)
  puts "Worker #{name} starting"
  
  # This looks synchronous but Ruby's thread scheduler
  # can switch to other threads during I/O.
  sleep 2
  
  puts "Worker #{name} done"
end

threads = []
["A", "B", "C"].each do |name|
  threads << Thread.new { worker(name) }
end

threads.each(&:join)
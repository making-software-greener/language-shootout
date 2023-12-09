def bubble_sort(arr)
  n = arr.length
  loop do
    swapped = false

    (n-1).times do |i|
      if arr[i] > arr[i+1]
        arr[i], arr[i+1] = arr[i+1], arr[i]
        swapped = true
      end
    end

    break unless swapped
  end
  arr
end

puts "Enter the number of elements:"
n = gets.to_i

arr = []
puts "Enter #{n} integers (one per line):"
n.times do
  arr << gets.to_i
end

sorted_array = bubble_sort(arr)

puts "Sorted array:"
puts sorted_array


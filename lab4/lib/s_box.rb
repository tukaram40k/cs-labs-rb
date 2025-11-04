require_relative 's_boxes'

class SBox
  attr_accessor :random

  def initialize
    @random = false
  end

  def set_random
    @random = true
  end

  def rand_48bit
    rand(2**48).to_s(2).rjust(48, '0')
  end

  def print_table(table, index)
    puts "used table S#{index}:"
    table.each_with_index do |row, i|
      puts "  #{row.map { |n| n.to_s.rjust(2) }.join(' ')}"
    end
  end

  def sbox_sub(input_bits)
    input_bits = rand_48bit if @random

    unless input_bits.is_a?(String) && input_bits =~ /\A[01]{48}\z/
      raise ArgumentError, "input must be a 48-bit binary string"
    end

    output = ""
    input_bits.chars.each_slice(6).with_index(1) do |chunk_arr, idx|
      bits = chunk_arr.join
      row_bits = bits[0] + bits[5]
      col_bits = bits[1..4]
      row = row_bits.to_i(2)
      col = col_bits.to_i(2)
      sbox = S_BOXES[idx - 1]
      val = sbox[row][col]
      val_bin = val.to_s(2).rjust(4, '0')

      puts "-" * 60
      puts "chunk #{idx}: #{bits}"
      print_table(sbox, idx)
      puts "output: #{val_bin}"
      output << val_bin

    end
    puts "-" * 60
    puts "final output: #{output}"
    output
  end
end

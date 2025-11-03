require_relative 'lib/s_box'

sbox = SBox.new

input = "111000110010101011001010101010100110101011011110"

sbox.sbox_sub(input)

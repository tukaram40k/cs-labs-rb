require_relative 'lib/s_box'

input = "111000110010101011001010101010100110101011011110"

sbox = SBox.new
sbox.set_random
sbox.sbox_sub(input)

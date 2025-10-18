class PlayfairOperations
  attr_reader :letters

  def initialize
    @letters = %w[A Ă B C D E F G H I Î J K L M N O P Q R S Ș T Ț U V W X Y Z]
  end

  def encrypt(key, text)
    return "Invalid characters in key or text" unless valid_input?(key, text)
    return "Key must be 7+ letters" unless key.length >= 7

    matrix = build_matrix(key)
    digraphs = prepare_text(text)
    ciphertext = ""

    digraphs.each do |digraph|
      row1, col1 = find_position(matrix, digraph[0])
      row2, col2 = find_position(matrix, digraph[1])

      if row1 == row2
        ciphertext += matrix[row1][(col1 + 1) % 6] + matrix[row2][(col2 + 1) % 6]
      elsif col1 == col2
        ciphertext += matrix[(row1 + 1) % 5][col1] + matrix[(row2 + 1) % 5][col2]
      else
        ciphertext += matrix[row1][col2] + matrix[row2][col1]
      end
    end

    ciphertext
  end

  def decrypt(key, ciphertext)
    return "Invalid characters in key or ciphertext" unless valid_input?(key, ciphertext)
    return "Key must be 7+ letters" unless key.length >= 7

    matrix = build_matrix(key)
    digraphs = prepare_text(ciphertext)
    plaintext = ""

    digraphs.each do |digraph|
      row1, col1 = find_position(matrix, digraph[0])
      row2, col2 = find_position(matrix, digraph[1])

      if row1 == row2
        plaintext += matrix[row1][(col1 - 1) % 6] + matrix[row2][(col2 - 1) % 6]
      elsif col1 == col2
        plaintext += matrix[(row1 - 1) % 5][col1] + matrix[(row2 - 1) % 5][col2]
      else
        plaintext += matrix[row1][col2] + matrix[row2][col1]
      end
    end

    plaintext
  end

  private

  def valid_input?(key, text)
    (key + text).upcase.split('').all? { |c| @letters.include?(c) }
  end

  def build_matrix(key)
    key = key.upcase.gsub(/[^#{@letters.join}]/, '')
    key_chars = key.chars.uniq
    remaining = @letters - key_chars
    matrix = (key_chars + remaining).each_slice(6).to_a
  end

  def prepare_text(text)
    text = text.upcase.gsub(/[^#{@letters.join}]/, '')
    digraphs = []
    i = 0
    while i < text.length
      if i + 1 == text.length || text[i] == text[i + 1]
        digraphs << [text[i], 'X']
        i += 1
      else
        digraphs << [text[i], text[i + 1]]
        i += 2
      end
    end
    digraphs
  end

  def find_position(matrix, char)
    matrix.each_with_index do |row, i|
      j = row.index(char)
      return [i, j] if j
    end
  end
end